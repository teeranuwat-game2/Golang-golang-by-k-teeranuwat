package user

import (
	"context"
)

type Response struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type User struct {
	getUser  getUsername
	dbGetter dbGetter
}

func NewUser(getUser getUsername, dbGetter dbGetter) *User {
	return &User{
		getUser:  getUser,
		dbGetter: dbGetter,
	}
}

func (u *User) GetUsername(id string) string {
	return u.getUser.GetUsername(id)
}

func (u *User) GetUserFromDB(ctx context.Context, id int) (Response, error) {
	author, err := u.dbGetter.GetAuthor(ctx, int64(id))
	if err != nil {
		return Response{}, err
	}
	response := Response{
		ID:   author.ID,
		Name: author.Name,
		Bio:  author.Bio.String,
	}
	return response, nil
}
