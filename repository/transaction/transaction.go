package transaction

import (
	"errors"
	"fmt"
	"together/be8/delivery/view/transaction"
	"together/be8/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type TransDB struct {
	Db *gorm.DB
}

func NewTransDB(DB *gorm.DB) *TransDB {
	return &TransDB{
		Db: DB,
	}
}

func (t *TransDB) Shipment(UserID uint) (entities.Address, []entities.Cart, []string, error) {
	var Address entities.Address
	var Cart []entities.Cart
	Seller := []string{}
	if err := t.Db.Where("user_id=? AND address_default='yes'", UserID).First(&Address).Error; err != nil {
		return Address, Cart, Seller, err
	}
	if err := t.Db.Where("user_id=? AND to_buy='yes'", UserID).Find(&Cart).Error; err != nil {
		return Address, Cart, Seller, err
	}
	if err := t.Db.Table("carts").Where("user_id = ? AND to_buy='yes'", UserID).Select("name_seller").Distinct("name_seller").Order("created_at DESC").Find(&Seller).Error; err != nil {
		log.Warn("Error Get All Cart", err)
		return Address, Cart, Seller, err
	}
	return Address, Cart, Seller, nil
}

// CREATE NEW TRANSACTION
func (t *TransDB) CreateTransaction(NewTransaction entities.Transaction) (entities.Transaction, error) {
	address, carts, _, err := t.Shipment(NewTransaction.UserID)
	if err != nil {
		log.Warn(err)
		return entities.Transaction{}, err
	}
	var totalbill int
	for _, v := range carts {
		totalbill += v.Qty * v.Price
	}
	NewTransaction.Address = address.Street + " " + address.SubDistrict + " " + address.City
	NewTransaction.OrderID = fmt.Sprintf("Order-%d%d", int(NewTransaction.UserID), int(carts[0].ID))
	NewTransaction.TotalBill = totalbill
	var cart entities.Cart
	if err := t.Db.Table("carts").Where("to_buy='yes' AND deleted_at IS NULL").Update("order_id", NewTransaction.OrderID).Error; err != nil {
		log.Warn(err)
		return NewTransaction, err
	}
	if err := t.Db.Where("to_buy='yes'").Delete(&cart).Error; err != nil {
		log.Warn(err)
		return NewTransaction, err
	}

	if err := t.Db.Create(&NewTransaction).Error; err != nil {
		log.Warn(err)
		return NewTransaction, err
	}

	return NewTransaction, nil
}

// GET ALL Transaction IN DATABASE
func (t *TransDB) GetAllTransaction(UserID uint) ([]transaction.AllTrans, error) {
	var AllTransaction []entities.Transaction
	var resAllTrans []transaction.AllTrans

	if err := t.Db.Where("user_id = ?", UserID).Order("created_at DESC").Find(&AllTransaction).Error; err != nil {
		log.Warn("Error Get All Transaction", err)
		return resAllTrans, errors.New("Access Database Error")
	}

	for _, v := range AllTransaction {
		var resTrans transaction.AllTrans
		resTrans.TransDetail = v
		var cards []entities.Cart
		if err := t.Db.Unscoped().Where("order_id", v.OrderID).Find(&cards).Error; err != nil {
			log.Warn("Error Get All Transaction", err)
			return resAllTrans, errors.New("Access Database Error")
		}
		fmt.Println(cards)
		resTrans.Product = cards
		resAllTrans = append(resAllTrans, resTrans)
	}

	return resAllTrans, nil
}

// GET Transaction BY ID
func (t *TransDB) GetTransactionDetail(UserID uint, OrderID string) (transaction.AllTrans, error) {
	resAllTrans, err := t.GetAllTransaction(UserID)
	if err != nil {
		log.Warn(err)
		return transaction.AllTrans{}, err
	}
	for _, v := range resAllTrans {
		if v.TransDetail.OrderID == OrderID {
			return v, nil
		}
	}
	return transaction.AllTrans{}, errors.New("Get Transaction Error")
}

// UPDATE Transaction BY ID
func (t *TransDB) PayTransaction(UserID uint, OrderID string) (entities.Transaction, error) {
	var updated entities.Transaction

	if err := t.Db.Where("user_id =? AND order_id=?", UserID, OrderID).First(&updated).Update("status", "success").Error; err != nil {
		log.Warn("Pay Transaction Error", err)
		return updated, errors.New("Access Database Error")
	}
	return updated, nil
}

// DELETE Transaction BY ID
func (t *TransDB) CancelTransaction(UserID uint, OrderID string) error {
	var Cancel entities.Transaction
	if err := t.Db.Where("user_id = ? AND order_id = ?", UserID, OrderID).First(&Cancel).Update("status", "failured").Error; err != nil {
		log.Warn("Cancel Transaction Error")
		return err
	}
	return nil
}
