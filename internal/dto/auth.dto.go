package dto

type AuthRegisterReq struct {
	UserTag  string `json:"userTag" validate:"required"`
	Name     string `json:"Name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type AuthRegisterRes struct {
	IsError    bool   `json:"isError"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type AuthLoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginRes struct {
	IsError    bool   `json:"isError"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}
