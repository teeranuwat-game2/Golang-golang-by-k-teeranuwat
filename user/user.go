package user

type User struct {
	getUser getUsername
}

func NewUser(getUser getUsername) *User {
	return &User{
		getUser: getUser,
	}
}

func (u *User) GetUsername(id string) string {
	return u.getUser.GetUsername(id)
}
