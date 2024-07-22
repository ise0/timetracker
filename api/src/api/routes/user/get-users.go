package userRouter

import (
	"strconv"
	"timetracker/src/lib"
	userModel "timetracker/src/model/user"

	"github.com/gin-gonic/gin"
)

type getUsersRes struct {
	UserId         int    `json:"userId"`
	Name           string `json:"userName"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	PassportNumber string `json:"passportNumber"`
	Address        string `json:"address"`
}

func getUsers(c *gin.Context) {
	var page, show = 1, 3
	if c.Query("page") != "" {
		if v, err := strconv.Atoi(c.Query("page")); err != nil {
			c.JSON(400, "provide valid page")
		} else {
			page = v
		}
	}
	if c.Query("show") != "" {
		if v, err := strconv.Atoi(c.Query("show")); err != nil {
			c.JSON(400, `provide valid "show" parameter`)
		} else {
			show = v
		}
	}
	pagi := userModel.GetUsersPaginationParam{
		Limit:  show,
		Offset: show * (page - 1),
	}

	var filters = userModel.GetUsersFiltersParam{
		Name:           c.Query("userName"),
		Surname:        c.Query("surname"),
		Patronymic:     c.Query("patronymic"),
		Address:        c.Query("address"),
		PassportNumber: c.Query("passport"),
	}

	for _, v := range c.QueryArray("userId") {
		if n, err := strconv.Atoi(v); err == nil {
			filters.UserId = append(filters.UserId, n)
		}
	}

	r, err := userModel.GetUsers(pagi, filters)

	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	var res = make([]getUsersRes, len(r))
	for i, v := range r {
		res[i] = getUsersRes{v.UserId, v.Name, v.Surname, v.Patronymic, v.PassportNumber, v.Address}
	}

	c.JSON(200, res)
}
