package entity

type IAuthRepository interface{
	Login(input *Usuario)(interface{},error)
}