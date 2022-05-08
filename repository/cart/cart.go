package cart

import (
	"errors"
	"fmt"
	"together/be8/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type RepoCart struct {
	Db *gorm.DB
}

func NewRepoCart(DB *gorm.DB) *RepoCart {
	return &RepoCart{
		Db: DB,
	}
}

// CREATE NEW CART
func (r *RepoCart) CreateCart(NewCart entities.Cart) (entities.Cart, error) {
	if err := r.Db.Create(&NewCart).Error; err != nil {
		log.Warn(err)
		return NewCart, err
	}
	return NewCart, nil
}

// GET ALL Cart IN DATABASE
func (r *RepoCart) GetAllCart(UserID uint) ([]entities.Cart, []string, error) {
	var AllCart []entities.Cart

	test := []string{}
	fmt.Println(UserID)
	if err := r.Db.Table("carts").Where("user_id = ?", UserID).Select("name_seller").Distinct("name_seller").Order("created_at DESC").Find(&test).Error; err != nil {
		log.Warn("Error Get All Cart", err)
		return AllCart, test, errors.New("Access Database Error")
	}

	if err := r.Db.Where("user_id = ?", UserID).Find(&AllCart).Error; err != nil {
		log.Warn("Error Get All Cart", err)
		return AllCart, test, errors.New("Access Database Error")
	}
	return AllCart, test, nil
}

// GET Cart BY ID
func (r *RepoCart) GetCartID(id uint, UserID uint) (entities.Cart, error) {
	var Cart entities.Cart
	if err := r.Db.Where("id = ? AND user_id = ?", id, UserID).First(&Cart).Error; err != nil {
		log.Warn("Error Get Cart By ID", err)
		return Cart, errors.New("Access Database Error")
	}

	return Cart, nil
}

// UPDATE Cart BY ID
func (r *RepoCart) UpdateCart(id uint, updatedCart entities.Cart, UserID uint) (entities.Cart, error) {
	var updated entities.Cart

	if err := r.Db.Table("carts").Where("id =? AND user_id =?", id, UserID).Updates(&updatedCart).First(&updated).Error; err != nil {
		log.Warn("Update Cart Error", err)
		return updated, errors.New("Access Database Error")
	}

	return updated, nil
}

// DELETE Cart BY ID
func (r *RepoCart) DeleteCart(id uint, UserID uint) error {

	var delete entities.Cart
	if err := r.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}
