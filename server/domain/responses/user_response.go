package responses

import "access-granting/domain/entities"

type UserResponse struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage"`
}

func NewUserResponse(user entities.User) UserResponse {
	return UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}
}
