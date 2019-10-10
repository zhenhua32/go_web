package service

import (
	"fmt"
	"sync"

	"tzh.com/web/model"
	"tzh.com/web/util"
)

// 业务处理函数, 获取用户列表
func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 并行转换
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortID, err := util.GenShortID()
			if err != nil {
				errChan <- err
				return
			}

			// 更新数据时加锁, 保持一致性
			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			userList.IdMap[u.ID] = &model.UserInfo{
				ID:        u.ID,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortID),
				Password:  u.Password,
				CreatedAt: util.TimeToStr(&u.CreatedAt),
				UpdatedAt: util.TimeToStr(&u.UpdatedAt),
				DeletedAt: util.TimeToStr(u.DeletedAt),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	// 等待完成
	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
