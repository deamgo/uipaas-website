package developer

type Developer struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"min=6,max=12"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"verifyPwd"`
	Status   int    `json:"status"`
}
