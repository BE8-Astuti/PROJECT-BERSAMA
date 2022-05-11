package transaction

import (
	"fmt"
	"net/http"
	middlewares "together/be8/delivery/middleware"
	"together/be8/delivery/view"
	transV "together/be8/delivery/view/transaction"
	"together/be8/entities"
	"together/be8/repository/transaction"
	"together/be8/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ControlTrans struct {
	repo     transaction.RepoTrans
	midtrans utils.ConfigMidtrans
	valid    *validator.Validate
}

func NewRepoTrans(Repo transaction.RepoTrans, validate *validator.Validate, midtrans utils.ConfigMidtrans) *ControlTrans {
	return &ControlTrans{
		repo:     Repo,
		valid:    validate,
		midtrans: midtrans,
	}
}

// CREATE NEW TRANSACTION
func (t *ControlTrans) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var InsertTransaction transV.InsertTransaction
		if err := c.Bind(&InsertTransaction); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}
		if err := t.valid.Struct(&InsertTransaction); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		NewTransaction := entities.Transaction{UserID: uint(UserID), Address: InsertTransaction.Address}
		fmt.Println(NewTransaction)
		result, err := t.repo.CreateTransaction(NewTransaction)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		SnapRedirectUrl := t.midtrans.CreateTransaction(result.OrderID, int64(result.TotalBill))
		if SnapRedirectUrl == nil {
			log.Warn("Failured Get Redirect Url")
			return c.JSON(http.StatusNoContent, transV.StatusErrorSnap())
		}

		return c.JSON(http.StatusCreated, transV.StatusCreate(result.OrderID, SnapRedirectUrl))
	}
}

// GET ALL TRANSACTION
func (t *ControlTrans) GetAllTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := t.repo.GetAllTransaction(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, transV.StatusGetAllOk(result))
	}
}

// GET TRANSACTION DETAILS GET
func (t *ControlTrans) GetTransactionDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderID := c.Param("order_id")
		UserID := middlewares.ExtractTokenUserId(c)

		result, err := t.repo.GetTransactionDetail(uint(UserID), orderID)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, transV.StatusTransactionDetail(result))
	}
}

// PAY TRANSACTION
func (t *ControlTrans) PayTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID := c.Param("order_id")
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := t.repo.PayTransaction(uint(UserID), OrderID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		resTrans := transV.RespondTransaction{
			OrderID:       result.OrderID,
			TotalBill:     result.TotalBill,
			PaymentMethod: result.PaymentMethod,
			Address:       result.Address,
			Status:        result.Status,
			CreatedAt:     result.CreatedAt,
		}
		return c.JSON(http.StatusOK, transV.StatusPayTrans(resTrans))
	}
}

// CANCEL TRANSACTION
func (t *ControlTrans) CancelTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID := c.Param("order_id")
		UserID := middlewares.ExtractTokenUserId(c)
		err := t.repo.CancelTransaction(uint(UserID), OrderID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, transV.StatusCancelTrans())
	}
}
