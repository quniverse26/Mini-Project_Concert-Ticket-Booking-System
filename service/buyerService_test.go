package service

import (
	"fmt"
	"testing"

	"github.com/quniverse26/miniproject/dao"
	"github.com/quniverse26/miniproject/model"
	"github.com/stretchr/testify/assert"
)

type BuyerService struct {
	buyerDao dao.BuyerDaoInterface
}

func (b BuyerService) GetBuyer(id int) (model.Buyer, error) {
	return b.buyerDao.GetBuyer(id)
}

type BuyerDaoMockFoundBuyer struct{}

func (b *BuyerDaoMockFoundBuyer) GetBuyer(id int) (model.Buyer, error) {
	return model.Buyer{
		Name: "Rangga",
	}, nil
}

type BuyerDaoMockNotFoundBuyer struct {}

func (b *BuyerDaoMockNotFoundBuyer) GetBuyer(id int) (model.Buyer, error)  {
	return model.Buyer{}, fmt.Errorf("buyer not found")
}

func Test_BuyerService_GetBuyer(t *testing.T)  {
	type fields struct{
		buyerDao dao.BuyerDaoInterface
	}
	type args struct {
		id int
	}
	tests := []struct {
		Name string
		fields fields
		args args
		want model.Buyer
		wantErr error
	}{
		{
		Name: `should return buyer when found`,
		fields: fields{
			buyerDao: &BuyerDaoMockFoundBuyer{},
		},
		args: args{10},
		want: model.Buyer{Name: "Rangga"},
	},
	{
		Name: `should return error when not found`,
		fields: fields{
			buyerDao: &BuyerDaoMockNotFoundBuyer{},
		},
		args: args{10},
		wantErr: fmt.Errorf("buyer not found"),
	},
}

for _, tt := range tests {
	t.Run(tt.Name, func(t *testing.T) {
		b := BuyerService {
			buyerDao: tt.fields.buyerDao,
		}
		got, err := b.GetBuyer(tt.args.id)
		assert.Equal(t, tt.want, got)
		assert.Equal(t, tt.wantErr, err)
	})
}
}