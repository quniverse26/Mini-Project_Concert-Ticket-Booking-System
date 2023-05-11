package service

import (
	"fmt"
	"testing"

	"github.com/quniverse26/miniproject/dao"
	"github.com/quniverse26/miniproject/model"
	"github.com/stretchr/testify/assert"
)

type AdminService struct {
	adminDao dao.AdminDaoInterface
}

func (a AdminService) GetAdmin(id int) (model.Admin, error) {
	return a.adminDao.GetAdmin(id)
}

type AdminDaoMockFoundAdmin struct{}

func (a *AdminDaoMockFoundAdmin) GetAdmin(id int) (model.Admin, error) {
	return model.Admin{
		Name: "admin",
	}, nil
}

type AdminDaoMockNotFoundAdmin struct {}

func (a *AdminDaoMockNotFoundAdmin) GetAdmin(id int) (model.Admin, error)  {
	return model.Admin{}, fmt.Errorf("admin not found")
}

func Test_AdminService_GetAdmin(t *testing.T)  {
	type fields struct{
		adminDao dao.AdminDaoInterface
	}
	type args struct {
		id int
	}
	tests := []struct {
		Name string
		fields fields
		args args
		want model.Admin
		wantErr error
	}{
		{
		Name: `should return admin when found`,
		fields: fields{
			adminDao: &AdminDaoMockFoundAdmin{},
		},
		args: args{10},
		want: model.Admin{Name: "admin"},
	},
	{
		Name: `should return error when not found`,
		fields: fields{
			adminDao: &AdminDaoMockNotFoundAdmin{},
		},
		args: args{10},
		wantErr: fmt.Errorf("admin not found"),
	},
}

for _, tt := range tests {
	t.Run(tt.Name, func(t *testing.T) {
		a := AdminService {
			adminDao: tt.fields.adminDao,
		}
		got, err := a.GetAdmin(tt.args.id)
		assert.Equal(t, tt.want, got)
		assert.Equal(t, tt.wantErr, err)
	})
}
}