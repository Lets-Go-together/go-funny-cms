package auth

// 登陆需要的参数
type LoginParams struct {
	Account  string `validate:"required|email"`
	Password string `validate:"required"`
}

func validate() {

}
