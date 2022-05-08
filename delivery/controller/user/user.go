package user

import (
	"net/http"
	"strconv"
	jwt "together/be8/delivery/middleware"
	"together/be8/delivery/view"
	userview "together/be8/delivery/view/user"
	"together/be8/entities"
	ruser "together/be8/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  ruser.User
	Valid *validator.Validate
}

func New(repo ruser.User, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

func (uc *UserController) InsertUser(c echo.Context) error {
	var tmpUser userview.InsertUserRequest

	if err := c.Bind(&tmpUser); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	if err := uc.Valid.Struct(tmpUser); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	newUser := entities.User{Name: tmpUser.Name, Email: tmpUser.Email, Password: tmpUser.Password, Phone: tmpUser.Phone}
	res, err := uc.Repo.InsertUser(newUser)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, userview.SuccessInsert(res))
}

func (uc *UserController) GetAllUser(c echo.Context) error {

	res, err := uc.Repo.GetAllUser()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}

func (uc *UserController) GetUserbyID(c echo.Context) error {
	id := c.Param("id")

	convID, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "connot convert ID",
			"data":    nil,
		})
	}

	hasil, err := uc.Repo.GetUserID(convID)

	if err != nil {
		log.Warn(err)
		notFound := "data tidak ditemukan"
		if err.Error() == notFound {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())

	}

	log.Info("data user found")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "data user ditemukan",
		"status":  true,
		"data":    hasil,
	})

}

func (uc *UserController) UpdateUserID(c echo.Context) error {

	u := new(userview.UpdateUserRequest)

	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))

	hasil, err := uc.Repo.UpdateUser(id, u.Email)

	if err != nil {
		log.Warn(err)
		notFound := "data tidak ditemukan"
		if err.Error() == notFound {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())

	}

	log.Info("data user update")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "data user update",
		"status":  true,
		"data":    hasil,
	})

}
func (uc *UserController) Login(c echo.Context) error {
	param := userview.LoginRequest{}

	if err := c.Bind(&param); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	if err := uc.Valid.Struct(&param); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	hasil, err := uc.Repo.Login(param.Email, param.Password)

	if err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusNotFound, "Email atau Password tidak ditemukan")
	}

	res := userview.LoginResponse{}

	if res.Token == "" {
		token, _ := jwt.CreateToken(int(hasil.ID), (hasil.Name), (hasil.Email))
		res.Token = token
		return c.JSON(http.StatusOK, userview.LoginOK(res, "Berhasil login"))
	}

	return c.JSON(http.StatusOK, userview.LoginOK(res, "Berhasil login"))
}

func (uc *UserController) DeleteUserID(c echo.Context) error {
	id := c.Param("id")

	convID, _ := strconv.Atoi(id)
	res, err := uc.Repo.DeleteUser(convID)

	if err != nil {
		log.Warn(err)
		notFound := "data tidak dapat didelete"
		if err.Error() == notFound {
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())

	}
	log.Info("data user delete")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "data user delete",
		"status":  true,
		"data":    res,
	})
}
