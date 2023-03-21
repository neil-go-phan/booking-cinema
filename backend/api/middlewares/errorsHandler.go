package middlewares

import (
	"booking-cinema-backend/helper"

	"github.com/gin-gonic/gin"
)

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}
// TODO: handle error 
func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if (len(c.Errors) != 0) {
			errResponse := helper.NewErrorReponseAndLog(c.Errors.Last())
			c.AbortWithStatusJSON(errResponse.HttpStatusCode, gin.H{"success": false, "message": errResponse.ErrorMessage})
			return
			// for i, ginErr := range c.Errors {
				// log error // push kanaba, log sentry,....
				// log.Println(i, ginErr)
				// return last error obj pushed
			// }
		}
	}
}