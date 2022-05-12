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

// CREATE NEW TRANSACTION
func (t *TransDB) CreateTransaction(NewTransaction entities.Transaction) (entities.Transaction, error) {
	var Carts []entities.Cart

	if err := t.Db.Where("user_id=? AND to_buy='yes'", NewTransaction.UserID).Find(&Carts).Error; err != nil {
		return NewTransaction, err
	}
	if len(Carts) == 0 {
		return NewTransaction, errors.New("Data Not Found")
	}
	var totalbill int
	for _, v := range Carts {
		totalbill += v.Qty * v.Price
	}
	NewTransaction.OrderID = fmt.Sprintf("Order-%d%d", int(NewTransaction.UserID), int(Carts[0].ID))
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
		resTrans.TransDetail = transaction.RespondTransaction{
			OrderID:       v.OrderID,
			TotalBill:     v.TotalBill,
			PaymentMethod: v.PaymentMethod,
			Address:       v.Address,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
		}
		var carts []entities.Cart
		if err := t.Db.Unscoped().Where("order_id=?", v.OrderID).Order("name_seller").Find(&carts).Error; err != nil {
			log.Warn("Error Get All Transaction", err)
			return resAllTrans, errors.New("Access Database Error")
		}
		var AllProducts []transaction.ProductTransaction
		for _, v := range carts {
			Product := transaction.ProductTransaction{
				ProductID:   v.ProductID,
				NameSeller:  v.NameSeller,
				NameProduct: v.NameProduct,
				Qty:         v.Qty,
				Price:       v.Price,
				UrlProduct:  v.UrlProduct,
				SubTotal:    v.Qty * v.Price,
			}
			AllProducts = append(AllProducts, Product)
		}
		resTrans.Product = AllProducts
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

// FINISH PAYMENT
func (t *TransDB) FinishPayment(OrderID string, updateStatus entities.Transaction) (entities.Transaction, error) {
	var result entities.Transaction
	if err := t.Db.Where("order_id = ?", OrderID).Updates(&updateStatus).Find(&result).Error; err != nil {
		log.Warn(err)
		return updateStatus, errors.New("Access Database Error")
	}
	return result, nil
}
