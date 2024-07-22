package userRouter

import (
	"timetracker/src/lib"
	userModel "timetracker/src/model/user"

	"github.com/gin-gonic/gin"
)

type AddUserBodyParam struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	PassportNumber string `json:"passportNumber"`
}

type AddUserRes struct {
	UserId int `json:"userId"`
}

func addUser(c *gin.Context) {
	var b AddUserBodyParam
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(400, "invalid body")
		return
	}

	u := userModel.AddUserParam{
		Name:           b.Name,
		Surname:        b.Surname,
		Patronymic:     b.Patronymic,
		Address:        b.Address,
		PassportNumber: b.PassportNumber,
	}

	res, err := userModel.AddUser(u)

	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	c.JSON(200, AddUserRes{res.UserId})
}
