package middleware

import (
	"net/http"
	"vette-tracker-services/internal/errors"
	"vette-tracker-services/internal/models"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only handle errors if there are any
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var response models.ErrorResponse

			switch e := err.(type) {
			case *errors.NotFoundError:
				response = models.ErrorResponse{
					Code:    http.StatusNotFound,
					Message: e.Error(),
				}
			case *errors.ValidationError:
				response = models.ErrorResponse{
					Code:    http.StatusBadRequest,
					Message: e.Error(),
				}
			case *errors.DatabaseError:
				response = models.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
					Details: e.Error(),
				}
			default:
				response = models.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				}
			}

			c.JSON(response.Code, response)
			c.Abort()
		}
	}
}
