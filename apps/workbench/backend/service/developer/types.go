package developer

type Developer struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"verifyPwd"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
}
