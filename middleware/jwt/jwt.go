package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"wjjmjh/hermes/pkg/api_response"
	"wjjmjh/hermes/pkg/util/jwt_"
)

// JWT is jwt_ middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = api_response.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = api_response.INVALID_PARAMS
		} else {
			_, err := jwt_.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = api_response.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = api_response.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != api_response.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  api_response.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
