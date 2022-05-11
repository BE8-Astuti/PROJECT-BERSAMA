package product

import (
	"together/be8/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type ProdukRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *ProdukRepo {
	return &ProdukRepo{
		Db: db,
	}
}

func (pr *ProdukRepo) InsertProduk(newProduk entities.Product) (entities.Product, error) {
	if err := pr.Db.Create(&newProduk).Error; err != nil {
		log.Warn(err)
		return newProduk, err
	}

	log.Info()
	return newProduk, nil
}

func (pr *ProdukRepo) GetProdBySeller(UserID uint) ([]entities.Product, error) {
	var produks []entities.Product
	if err := pr.Db.Where("user_id = ?", UserID).Find(&produks).Error; err != nil {
		log.Warn("Eror Get Produk", err)
		return produks, err
	}
	return produks, nil
}

func (pr *ProdukRepo) GetProdbyID(id uint) (entities.Product, error) {
	var produk entities.Product
	if err := pr.Db.Where("id= ?", id).Find(&produk).Error; err != nil {
		log.Warn("Error Get By ID", err)
		return produk, err
	}
	return produk, nil
}

func (pr *ProdukRepo) GetProdByCategory(id int) ([]entities.Product, error) {
	var produks []entities.Product
	if err := pr.Db.Where("category_id= ?", id).Find(&produks).Error; err != nil {
		log.Warn("Eror Get Produk", err)
		return produks, err
	}
	return produks, nil
}

func (pr *ProdukRepo) UpdateProduk(id int, UpdateProduk entities.Product, UserID uint) (entities.Product, error) {
	var produks entities.Product

	if err := pr.Db.Where("id =? AND user_id =?", id, UserID).First(&produks).Updates(&UpdateProduk).Find(&produks).Error; err != nil {
		log.Warn("Update Address Error", err)
		return produks, err
	}

	return produks, nil
}

func (pr *ProdukRepo) DeleteProduk(id uint, UserID uint) error {

	var delete entities.Product
	if err := pr.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
