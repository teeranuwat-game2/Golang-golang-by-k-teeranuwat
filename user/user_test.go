package user

import (
	"context"
	"dayone/gensql/postgresdb"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UserTestSuite struct {
	suite.Suite
	getUsername  *MockgetUsername
	mockdbGetter *MockdbGetter
	user         *User
}

func (s *UserTestSuite) SetupTest() {
	mController := gomock.NewController(s.T())
	s.getUsername = NewMockgetUsername(mController)
	s.mockdbGetter = NewMockdbGetter(mController)
	s.user = NewUser(s.getUsername, s.mockdbGetter)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestGetUsernameWhenUserExistsThenReturnUsername() {
	s.getUsername.EXPECT().GetUsername("123").Return("john_doe")
	username := s.user.GetUsername("123")
	s.Equal("john_doe", username)
}

func (s *UserTestSuite) TestGetUserAuthorsFromDatabase() {
	ctx := context.Background()
	s.mockdbGetter.EXPECT().GetAuthor(ctx, int64(1)).Return(postgresdb.Author{ID: 1, Name: "John Doe", Bio: pgtype.Text{
		String: "Author Bio",
		Valid:  true,
	}}, nil)

	reponse, err := s.user.GetUserFromDB(ctx, 1)
	s.NoError(err)
	s.Equal(Response{ID: 1, Name: "John Doe", Bio: "Author Bio"}, reponse)
}
