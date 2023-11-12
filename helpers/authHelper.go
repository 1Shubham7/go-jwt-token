package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUserId(c *gin.Context, userId string) (err error) {
	//  This is the match user function
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	// this means that user is USER not ADMIN and uid is not of the user. Because user can only access his id,
	// admin can access anyone's id
	if userId == "USER" && uid != userId {
		err = errors.New("You are not authorized to access this user")
	}
	err = CheckUserType(c, userType)
	return err

}