package users

type UserModel struct{}

// 此信息将写入鉴权中
type AuthUser struct {
	Id   int
	Name string
}
