package userRouter

import (
	"timetracker/src/lib"
	userModel "timetracker/src/model/user"

	"github.com/gin-gonic/gin"
)

type updateUserBodyParam struct {
	UserId               int    `json:"userId"`
	Name                 string `json:"name"`
	Surname              string `json:"surname"`
	Patronymic           string `json:"patronymic"`
	Address              string `json:"address"`
	PassportNumber       string `json:"passportNumber"`
	NameUpdate           bool   `json:"nameUpdate"`
	SurnameUpdate        bool   `json:"surnameUpdate"`
	PatronymicUpdate     bool   `json:"patronymicUpdate"`
	AddressUpdate        bool   `json:"addressUpdate"`
	PassportNumberUpdate bool   `json:"passportNumberUpdate"`
}

func updateUser(c *gin.Context) {
	var b updateUserBodyParam
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(400, "invalid body")
		return
	}

	u := userModel.UpdateUserUserParam{
		UserId:               b.UserId,
		Name:                 b.Name,
		Surname:              b.Surname,
		Patronymic:           b.Patronymic,
		Address:              b.Address,
		PassportNumber:       b.PassportNumber,
		NameUpdate:           b.NameUpdate,
		SurnameUpdate:        b.SurnameUpdate,
		PatronymicUpdate:     b.PatronymicUpdate,
		AddressUpdate:        b.AddressUpdate,
		PassportNumberUpdate: b.PassportNumberUpdate,
	}

	err := userModel.UpdateUser(u)
	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	c.JSON(200, "success")
}
