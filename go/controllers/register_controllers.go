package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gongd3/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/gongd3/go")
	{ // insertion point for registrations
		v1.GET("/v1/bars", GetBars)
		v1.GET("/v1/bars/:id", GetBar)
		v1.POST("/v1/bars", PostBar)
		v1.PATCH("/v1/bars/:id", UpdateBar)
		v1.PUT("/v1/bars/:id", UpdateBar)
		v1.DELETE("/v1/bars/:id", DeleteBar)

		v1.GET("/v1/keys", GetKeys)
		v1.GET("/v1/keys/:id", GetKey)
		v1.POST("/v1/keys", PostKey)
		v1.PATCH("/v1/keys/:id", UpdateKey)
		v1.PUT("/v1/keys/:id", UpdateKey)
		v1.DELETE("/v1/keys/:id", DeleteKey)

		v1.GET("/v1/pies", GetPies)
		v1.GET("/v1/pies/:id", GetPie)
		v1.POST("/v1/pies", PostPie)
		v1.PATCH("/v1/pies/:id", UpdatePie)
		v1.PUT("/v1/pies/:id", UpdatePie)
		v1.DELETE("/v1/pies/:id", DeletePie)

		v1.GET("/v1/scatters", GetScatters)
		v1.GET("/v1/scatters/:id", GetScatter)
		v1.POST("/v1/scatters", PostScatter)
		v1.PATCH("/v1/scatters/:id", UpdateScatter)
		v1.PUT("/v1/scatters/:id", UpdateScatter)
		v1.DELETE("/v1/scatters/:id", DeleteScatter)

		v1.GET("/v1/series", GetSeries)
		v1.GET("/v1/series/:id", GetSerie)
		v1.POST("/v1/series", PostSerie)
		v1.PATCH("/v1/series/:id", UpdateSerie)
		v1.PUT("/v1/series/:id", UpdateSerie)
		v1.DELETE("/v1/series/:id", DeleteSerie)

		v1.GET("/v1/values", GetValues)
		v1.GET("/v1/values/:id", GetValue)
		v1.POST("/v1/values", PostValue)
		v1.PATCH("/v1/values/:id", UpdateValue)
		v1.PUT("/v1/values/:id", UpdateValue)
		v1.DELETE("/v1/values/:id", DeleteValue)

		v1.GET("/v1/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
