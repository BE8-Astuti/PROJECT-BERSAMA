package user

import (
	"net/http"
	"strconv"
	middlewares "together/be8/delivery/middleware"
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

func (uc *UserController) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUser userview.InsertUserRequest

		if err := c.Bind(&tmpUser); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := uc.Valid.Struct(&tmpUser); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
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
}

// func (uc *UserController) GetAllUser(c echo.Context) error {

// 	res, err := uc.Repo.GetAllUser()

// 	if err != nil {
// 		log.Warn("masalah pada server")
// 		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
// 	}
// 	log.Info("berhasil get all data")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":    http.StatusOK,
// 		"message": "berhasil get all data",
// 		"status":  true,
// 		"data":    res,
// 	})
// }

func (uc *UserController) GetUserbyID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		// UserID := middlewares.ExtractTokenUserId(c)
		// UserID := middlewares.ExtractTokenUserId(c)
		// log.Debugf("id: %d,  user: %d ", convID, UserID)
		// if UserID != float64(convID) {
		// 	return c.JSON(http.StatusNotFound, view.NotFound())
		// }

		hasil, err := uc.Repo.GetUserID(convID)

		if err != nil {
			log.Warn()
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		return c.JSON(http.StatusOK, userview.StatusGetIdOk(hasil))

	}

}

func (uc *UserController) UpdateUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update userview.UpdateUserRequest

		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		// UserID := middlewares.ExtractTokenUserId(c)

		// if UserID != float64(id) {
		// 	return c.JSON(http.StatusNotFound, view.NotFound())
		// }
		UpdateEmail := entities.User{Email: update.Email}

		hasil, err := uc.Repo.UpdateUser(id, UpdateEmail.Email)

		if err != nil {
			log.Warn(err)
			notFound := "data tidak ditemukan"
			if err.Error() == notFound {
				return c.JSON(http.StatusNotFound, view.NotFound())
			}
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())

		}

		return c.JSON(http.StatusOK, userview.StatusUpdate(hasil))
	}

}
func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := userview.LoginRequest{}

		if err := c.Bind(&param); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := uc.Valid.Struct(&param); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}

		hasil, err := uc.Repo.Login(param.Email, param.Password)

		if err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		res := userview.LoginResponse{}

		if res.Token == "" {
			token, _ := middlewares.CreateToken(float64(hasil.ID), (hasil.Name), (hasil.Email))
			res.Token = token
			return c.JSON(http.StatusOK, userview.LoginOK(res, "Berhasil login"))
		}

		return c.JSON(http.StatusOK, userview.LoginOK(res, "Berhasil login"))
	}
}

func (uc *UserController) DeleteUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		convID, err := strconv.Atoi(id)

		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		// UserID := middlewares.ExtractTokenUserId(c)

		// if UserID != float64(convID) {
		// 	return c.JSON(http.StatusNotFound, view.NotFound())
		// }

		_, erro := uc.Repo.DeleteUser(convID)

		if erro != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
