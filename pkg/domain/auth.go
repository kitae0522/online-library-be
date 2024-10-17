package domain

type DefaultRes struct {
	IsError    bool   `json:"isError"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type AuthRegisterReq struct {
	UserTag         string `json:"userTag" validate:"required"`
	Name            string `json:"Name" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
	Email           string `json:"email" validate:"required"`
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

type AuthPasswordResetReq struct {
	UserUUID           string `json:"userID" validate:"requierd"`
	OldPassword        string `json:"oldPassword" validate:"required"`
	NewPassword        string `json:"newPassword" validate:"required"`
	NewPasswordConfirm string `json:"newPasswordConfirm" validate:"required"`
}

type AuthWithdrawReq struct {
	UserUUID string `json:"userID" validate:"requierd"`
}
