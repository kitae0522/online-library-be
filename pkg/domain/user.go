package domain

import "time"

type UserCreateUserProfileReq struct {
	UserUUID   string  `json:"userUUID"`
	ProfilePic *string `json:"profilePic"`
	Bio        *string `json:"bio"`
}

type UserGetUserProfileEntity struct {
	UserUUID   string    `json:"userUUID"`
	UserTag    string    `json:"userTag"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	ProfilePic string    `json:"profilePic"`
	Bio        string    `json:"bio"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserGetUserProfileRes struct {
	IsError    bool                      `json:"isError"`
	StatusCode int                       `json:"statusCode"`
	Message    string                    `json:"message"`
	Profile    *UserGetUserProfileEntity `json:"profile"`
}
