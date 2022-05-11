package utils

type ConfigMidtrans interface {
	CreateTransaction(OrderID string, GrossAmt int64) map[string]interface{}
}
