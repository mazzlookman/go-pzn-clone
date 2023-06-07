package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-pzn-clone/formatter"
	"go-pzn-clone/middleware/auth"
	"go-pzn-clone/service"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(jwtAuth auth.JWTAuth, userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Getting the header value
		header := ctx.GetHeader("Authorization")
		if !strings.Contains(header, "Bearer") {
			apiResponse := formatter.APIResponse("You aren't Unauthorized", 401, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apiResponse)
			return
		}

		//Get token string from header value
		splitHeader := strings.Split(header, " ")
		token := splitHeader[1]

		//Validate token string
		validateToken, err := jwtAuth.ValidateToken(token)
		if err != nil {
			apiResponse := formatter.APIResponse("You aren't Unauthorized", 401, "error", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apiResponse)
			return
		}

		//Get payload/claim from token
		claims := validateToken.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		//Get user data based on userID from jwt payload
		findUserByID, err := userService.FindUserByID(userID)
		if err != nil {
			apiResponse := formatter.APIResponse("InternalServerError", http.StatusInternalServerError, "error", err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, apiResponse)
			return
		}

		//Set context which contain user data
		ctx.Set("currentUser", findUserByID)
	}
}
