package user

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UserTestSuite struct {
	suite.Suite
	getUsername *MockgetUsername
	user        *User
}

func (s *UserTestSuite) SetupTest() {
	mController := gomock.NewController(s.T())
	s.getUsername = NewMockgetUsername(mController)
	s.user = NewUser(s.getUsername)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestGetUsernameWhenUserExistsThenReturnUsername() {
	s.getUsername.EXPECT().GetUsername("123").Return("john_doe")
	username := s.user.GetUsername("123")
	s.Equal("john_doe", username)
}
