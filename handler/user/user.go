package user

import "tzh.com/web/model"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username" form:"username"`
	Offset   int    `json:"offset" form:"offset"`
	Limit    int    `json:"limit" form:"limit"`
}

type ListResponse struct {
	TotalCount uint              `json:"total_count"`
	UserList   []*model.UserInfo `json:"user_list"`
}
