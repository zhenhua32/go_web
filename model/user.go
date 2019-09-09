package model

import (
	validator "gopkg.in/go-playground/validator.v9"
	"tzh.com/web/pkg/auth"
	"tzh.com/web/pkg/constvar"
)

// 定义用户的结构
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (*UserModel) TableName() string {
	return "tb_users"
}

// 创建新用户
func (u *UserModel) Create() error {
	return DB.Self.Create(u).Error
}

// 删除用户
func (u *UserModel) Delete(id uint) error {
	return DB.Self.Delete(u).Error
}

// 保存用户
func (u *UserModel) Save() error {
	return DB.Self.Save(u).Error
}

// 比较密码是否相同
func (u *UserModel) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

// 加密用户密码
func (u *UserModel) Encrypt() error {
	password, err := auth.Encrypt(u.Password)
	if err == nil {
		u.Password = password
	}
	return err
}

var validate *validator.Validate

// 验证字段
func (u *UserModel) Validate() error {
	validate = validator.New()
	return validate.Struct(u)
}

// 基于名字获取用户
func GetUserByName(username string) (*UserModel, error) {
	user := &UserModel{}
	result := DB.Self.Where("username = ?", username).First(user)
	return user, result.Error
}

// 获取用户的列表, 用户的总数
func ListUser(username string, offset, limit int) ([]*UserModel, uint, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint

	where := DB.Self.Where("username like ?", "%"+username+"%")

	// 统计用户的总数
	if result := where.Find(users).Count(&count); result.Error != nil {
		return users, count, result.Error
	}

	// 获取用户
	if result := where.Offset(offset).Limit(limit).Order("id desc").Find(users); result.Error != nil {
		return users, count, result.Error
	}

	return users, count, nil
}
