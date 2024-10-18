package requests

import "mime/multipart"

type UserUpdateRequest struct {
	Username     string                `form:"username" binding:"required"`
	ProfileImage *multipart.FileHeader `form:"profileImage"`
}

type UserUpdateServiceRequest struct {
	Username     string
	ProfileImage string
}
