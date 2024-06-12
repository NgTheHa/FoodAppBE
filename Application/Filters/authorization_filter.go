package Filters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizationFilter struct {
	UserTypes []int
}

func NewAuthorizationFilter(userTypes ...int) *AuthorizationFilter {
	return &AuthorizationFilter{
		UserTypes: userTypes,
	}
}

func (f *AuthorizationFilter) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		userTypeClaim, exist := c.Get("userType")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Không có quyền"})
			c.Abort()
			return
		}

		userType, ok := userTypeClaim.(int)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Lỗi khi xác định loại người dùng"})
			c.Abort()
			return
		}

		if !f.containsUserType(userType) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Không có quyền"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (f *AuthorizationFilter) containsUserType(userType int) bool {
	for _, t := range f.UserTypes {
		if t == userType {
			return true
		}
	}
	return false
}
