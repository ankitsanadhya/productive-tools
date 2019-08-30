package service

import (
	"fmt"
	"license/api/model"
	"license/api/repository"
	"license/api/vo"

	"github.com/google/uuid"
)

type PartnerService struct {
}

func (ps *PartnerService) ListAllPartner(p *[]vo.Partner) (err error) {
	pr := new(repository.PartnerRepository)
	var partner []model.Partner
	err = pr.ListAllPartner(&partner)
	if err != nil {
		return
	}

	fmt.Println(partner)

	return
}
func (ps *PartnerService) AddPartner(p *vo.Partner) (err error) {

	pr := new(repository.PartnerRepository)
	partaner := new(model.Partner)
	partaner.PartnerID = uuid.New().String()
	partaner.Name = p.Name
	partaner.Description = p.Description
	partaner.Currency = p.Currency
	err = pr.AddPartner(partaner)
	if err == nil {
		p.PartnerID = partaner.PartnerID
	}
	return err
}

// Partner       Partner `gorm:"foreignkey:PartnerID"`
// 	PartnerID     string  `json:"partnerId,omitempty"`
// 	Name          string  `gorm:"type:varchar(100)" json:"name,omitempty"`
// 	Description   string  `gorm:"type:text" json:"description,omitempty"`
// 	DefaultPrice  float64 `gorm:"type:double;default null" json:"defaultPrice,omitempty"`
// 	IsDefault     bool    `gorm:"type:tinyint(1);not null;default:0" json:"isDefault,omitempty"`
// 	Specification string  `gorm:"type:text" json:"specifications,omitempty"`
func (ps *PartnerService) AddProduct(p *vo.Product, PartnerID string) (err error) {

	pr := new(repository.PartnerRepository)
	product := new(model.Product)
	product.PartnerID = PartnerID
	product.ProductID = uuid.New().String()
	product.Name = p.Name
	product.Description = p.Description
	product.DefaultPrice = p.DefaultPrice
	product.IsDefault = p.IsDefault
	product.Specification = p.Specification

	err = pr.AddProduct(product)
	if err == nil {
		p.ProductID = product.ProductID
	}
	return err
}

//var book []Models.Book
// func (partaner *PartnerRepository) ListAllPartner(p *[]model.Partner) (err error) {
