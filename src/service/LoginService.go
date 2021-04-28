package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email string
	password string
}

func StaticLoginService() *loginInformation {
	return &loginInformation{
		email: "test@mail.com",
		password: "test",
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
