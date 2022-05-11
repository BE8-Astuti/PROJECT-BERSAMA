package entities

<<<<<<< HEAD
type Transaction struct {
=======
import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID        uint
	Address       string `json:"address"`
	PaymentMethod string `json:"paymentMethod"`
	TotalBill     int    `json:"totalBill"`
	Status        string `json:"status" gorm:"default:pending"`
	OrderID       string `json:"orderID"`
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
}
