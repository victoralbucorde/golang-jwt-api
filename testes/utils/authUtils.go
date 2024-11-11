package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if userType != role {
		err = errors.New("unauthorized access")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("userType")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" || uid != userId {
		err = errors.New("unauthorized user")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}
