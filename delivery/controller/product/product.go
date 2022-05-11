package product

import (
	"strconv"
	middlewares "together/be8/delivery/middleware"
	"together/be8/delivery/view"

	vproduk "together/be8/delivery/view/product"

	"together/be8/repository/product"

	"net/http"
	"together/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProdukController struct {
	Repo product.RepoProduk

	Valid *validator.Validate
}

func New(NewAddp product.RepoProduk, validate *validator.Validate) *ProdukController {
	return &ProdukController{
		Repo:  NewAddp,
		Valid: validate,
	}
}

func (pc *ProdukController) InsertProd() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpProd vproduk.InsertProdukRequest

		if err := c.Bind(&tmpProd); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := pc.Valid.Struct(&tmpProd); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		id := middlewares.ExtractTokenUserId(c)
		newProd := entities.Product{
			UserID:      uint(id),
			CategoryID:  uint(tmpProd.CategoryID),
			NameSeller:  tmpProd.NameSeller,
			Name:        tmpProd.Name,
			Stock:       tmpProd.Stock,
			Price:       tmpProd.Price,
			Description: tmpProd.Description,
		}
		res, errInsert := pc.Repo.InsertProduk(newProd)

		if errInsert != nil {
			log.Warn(errInsert)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		response := vproduk.RespondProduct{
			ProductID:   res.ID,
			UserID:      res.UserID,
			CategoryID:  res.CategoryID,
			Name:        res.Name,
			NameSeller:  res.NameSeller,
			Stock:       res.Stock,
			Price:       res.Price,
			Description: res.Description,
		}
		log.Info("berhasil insert")
		return c.JSON(http.StatusCreated, vproduk.StatusCreate(response))
	}
}

func (pc *ProdukController) GetProdukbySeller() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		userid, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		result, errprodukID := pc.Repo.GetProdBySeller(uint(userid))
		if errprodukID != nil {
			log.Warn(errprodukID)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		var GetAll []vproduk.RespondProduct
		for _, v := range result {
			response := vproduk.RespondProduct{
				ProductID:   v.ID,
				UserID:      v.UserID,
				CategoryID:  v.CategoryID,
				Name:        v.Name,
				NameSeller:  v.NameSeller,
				Stock:       v.Stock,
				Price:       v.Price,
				Description: v.Description,
				Sold:        v.Sold,
			}
			GetAll = append(GetAll, response)
		}
		return c.JSON(http.StatusOK, vproduk.StatusGetIdOk(GetAll))
	}

}
func (pc *ProdukController) GetProdukByCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		categoryid, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		result, errprodukcate := pc.Repo.GetProdByCategory(categoryid)
		if errprodukcate != nil && len(result) == 0 {
			log.Warn(errprodukcate)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		var GetAll []vproduk.RespondProduct
		for _, v := range result {
			response := vproduk.RespondProduct{
				ProductID:   v.ID,
				UserID:      v.UserID,
				CategoryID:  v.CategoryID,
				Name:        v.Name,
				NameSeller:  v.NameSeller,
				Stock:       v.Stock,
				Price:       v.Price,
				Description: v.Description,
				Sold:        v.Sold,
			}
			GetAll = append(GetAll, response)
		}
		return c.JSON(http.StatusOK, vproduk.StatusGetAllOk(GetAll))
	}
}
func (pc *ProdukController) GetProID() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idcat, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, err := pc.Repo.GetProdbyID(uint(idcat))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		response := vproduk.RespondProduct{
			ProductID:   result.ID,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			Name:        result.Name,
			NameSeller:  result.NameSeller,
			Stock:       result.Stock,
			Price:       result.Price,
			Description: result.Description,
			Sold:        result.Sold,
		}
		return c.JSON(http.StatusOK, vproduk.StatusGetIdOk(response))
	}
}

func (pc *ProdukController) UpdateProduk() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idproduk, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		var update vproduk.UpdateProdukRequest
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		UpdateProduk := entities.Product{
			Name:        update.Name,
			Stock:       update.Stock,
			Price:       update.Price,
			Description: update.Description,
		}

		result, errNotFound := pc.Repo.UpdateProduk(idproduk, UpdateProduk, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		response := vproduk.RespondProduct{
			ProductID:   result.ID,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			Name:        result.Name,
			NameSeller:  result.NameSeller,
			Stock:       result.Stock,
			Price:       result.Price,
			Description: result.Description,
			Sold:        result.Sold,
		}
		return c.JSON(http.StatusOK, vproduk.StatusUpdate(response))
	}
}
func (pc *ProdukController) DeleteProduk() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		idproduk, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := pc.Repo.DeleteProduk(uint(idproduk), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
