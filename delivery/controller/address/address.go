package address

import (
	"net/http"
	"strconv"
	"together/be8/delivery/view"
	addressV "together/be8/delivery/view/address"
	"together/be8/entities"
	"together/be8/repository/address"

	"github.com/labstack/gommon/log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ControlAddress struct {
	Repo  address.RepoAddress
	Valid *validator.Validate
}

func NewControlAddress(NewAddr address.RepoAddress, validate *validator.Validate) *ControlAddress {
	return &ControlAddress{
		Repo:  NewAddr,
		Valid: validate,
	}
}

// METHOD Add New Address
func (r *ControlAddress) CreateAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert addressV.InsertAddress
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := r.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		NewAdd := entities.Address{
			// UserID:       1,
			Recipient:    Insert.Recipient,
			HP:           Insert.HP,
			Street:       Insert.Street,
			SubDistrict:  Insert.SubDistrict,
			UrbanVillage: Insert.UrbanVillage,
			City:         Insert.City,
			Zip:          Insert.Zip,
		}
		result, errCreate := r.Repo.CreateAddress(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, addressV.StatusCreate(result))
	}
}

// METHOD GET ALL ADDRESS
func (r *ControlAddress) GetAllAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := r.Repo.GetAllAddress()
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, addressV.StatusGetAllOk(result))
	}
}

// METHOD GET ADDRESS BY ID
func (r *ControlAddress) GetAddressID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idAddress, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, errGetAddressID := r.Repo.GetAddressID(uint(idAddress))
		if errGetAddressID != nil {
			log.Warn(errGetAddressID)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, addressV.StatusGetIdOk(result))
	}
}

// UPDATE ADDRESS BY ID
func (r *ControlAddress) UpdateAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update addressV.InsertAddress
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}
		id := c.Param("id")
		idAddress, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UpdateAddress := entities.Address{
			Recipient:    update.Recipient,
			HP:           update.HP,
			Street:       update.Street,
			UrbanVillage: update.UrbanVillage,
			SubDistrict:  update.SubDistrict,
			City:         update.City,
			Zip:          update.Zip,
		}

		result, errNotFound := r.Repo.UpdateAddress(uint(idAddress), UpdateAddress)
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, addressV.StatusUpdate(result))
	}
}

// DELETE ADDRESS BY ID
func (r *ControlAddress) DeleteAddress() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idAddress, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		errDelete := r.Repo.DeleteAddress(uint(idAddress))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
