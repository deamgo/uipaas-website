package user

type User struct {
	UID            uint   `json:"id"`
	Username       string `json:"username" validate:"min=6,max=12"`
	InvitationCode string `json:"invitation_code" validate:"len=6"`
	Email          string `json:"email" validate:"email"`
	Password       string `json:"password" validate:"verifyPwd" field_error_info:"密码由大小写字母和数字组成，且至少8个字符"`
}
