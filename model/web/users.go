package web

type CreateUser struct {
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname"`
	Password    string `json:"password" binding:"required"`
	Address     string `json:"address"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UpdateUser struct {
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname"`
	Password    string `json:"password" binding:"required"`
	Address     string `json:"address"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
