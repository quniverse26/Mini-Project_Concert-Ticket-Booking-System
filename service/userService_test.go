package service

import (
	"fmt"
	"testing"

	"github.com/quniverse26/miniproject/dao"
	"github.com/quniverse26/miniproject/model"
	"github.com/stretchr/testify/assert"
)

type UserService struct {
	userDao dao.UserDaoInterface
}

func (u UserService) GetUser(id int) (model.User, error) {
	return u.userDao.GetUser(id)
}

type UserDaoMockFoundUser struct{}

func (u *UserDaoMockFoundUser) GetUser(id int) (model.User, error) {
	return model.User{
		Name: "Rangga",
	}, nil
}

type UserDaoMockNotFoundUser struct {}

func (u *UserDaoMockNotFoundUser) GetUser(id int) (model.User, error)  {
	return model.User{}, fmt.Errorf("user not found")
}

func Test_UserService_GetUser(t *testing.T)  {
	type fields struct{
		userDao dao.UserDaoInterface
	}
	type args struct {
		id int
	}
	tests := []struct {
		Name string
		fields fields
		args args
		want model.User
		wantErr error
	}{
		{
		Name: `should return user when found`,
		fields: fields{
			userDao: &UserDaoMockFoundUser{},
		},
		args: args{10},
		want: model.User{Name: "Rangga"},
	},
	{
		Name: `should return error when not found`,
		fields: fields{
			userDao: &UserDaoMockNotFoundUser{},
		},
		args: args{10},
		wantErr: fmt.Errorf("user not found"),
	},
}

for _, tt := range tests {
	t.Run(tt.Name, func(t *testing.T) {
		u := UserService {
			userDao: tt.fields.userDao,
		}
		got, err := u.GetUser(tt.args.id)
		assert.Equal(t, tt.want, got)
		assert.Equal(t, tt.wantErr, err)
	})
}
}