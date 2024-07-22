package userRouter

import (
	"fmt"
	"strconv"
	"timetracker/src/lib"
	userModel "timetracker/src/model/user"

	"github.com/gin-gonic/gin"
)

type getUserByPassportRes struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func getUserByPassport(c *gin.Context) {
	var errMsg string
	passportNumber, err := strconv.Atoi(c.Query("passportNumber"))
	if err != nil {
		errMsg = "provide valid passportNumber"
	}
	passportSerie, err := strconv.Atoi(c.Query("passportSerie"))
	if err != nil {
		errMsg = "provide valid passportSerie"
	}
	if errMsg != "" {
		c.JSON(400, errMsg)
		return
	}

	r, err := userModel.GetUserByPassport(fmt.Sprintf("%v %v", passportSerie, passportNumber))
	if err != nil {
		if e, ok := err.(lib.ApiError); ok {
			c.JSON(int(e.Code), e.Error())
		} else {
			c.JSON(500, "something went wrong")
		}
		return
	}

	c.JSON(200, getUserByPassportRes{r.Name, r.Surname, r.Patronymic, r.Address})
}
