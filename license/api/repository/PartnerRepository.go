package repository

import (
	"license/api/model"
	"license/storage"
)

type PartnerRepository struct {
}

func (partaner *PartnerRepository) AddPartner(p *model.Partner) (err error) {
	if err = storage.DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (partaner *PartnerRepository) ListAllPartner(p *[]model.Partner) (err error) {

	if err = storage.DB.Preload("Product").Find(&p).Error; err != nil {
		return err
	}
	return nil
}

func (partaner *PartnerRepository) AddProduct(p *model.Product) (err error) {
	if err = storage.DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}
