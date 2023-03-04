package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		v1.GET("/v1/bars", GetController().GetBars)
		v1.GET("/v1/bars/:id", GetController().GetBar)
		v1.POST("/v1/bars", GetController().PostBar)
		v1.PATCH("/v1/bars/:id", GetController().UpdateBar)
		v1.PUT("/v1/bars/:id", GetController().UpdateBar)
		v1.DELETE("/v1/bars/:id", GetController().DeleteBar)

		v1.GET("/v1/keys", GetController().GetKeys)
		v1.GET("/v1/keys/:id", GetController().GetKey)
		v1.POST("/v1/keys", GetController().PostKey)
		v1.PATCH("/v1/keys/:id", GetController().UpdateKey)
		v1.PUT("/v1/keys/:id", GetController().UpdateKey)
		v1.DELETE("/v1/keys/:id", GetController().DeleteKey)

		v1.GET("/v1/pies", GetController().GetPies)
		v1.GET("/v1/pies/:id", GetController().GetPie)
		v1.POST("/v1/pies", GetController().PostPie)
		v1.PATCH("/v1/pies/:id", GetController().UpdatePie)
		v1.PUT("/v1/pies/:id", GetController().UpdatePie)
		v1.DELETE("/v1/pies/:id", GetController().DeletePie)

		v1.GET("/v1/scatters", GetController().GetScatters)
		v1.GET("/v1/scatters/:id", GetController().GetScatter)
		v1.POST("/v1/scatters", GetController().PostScatter)
		v1.PATCH("/v1/scatters/:id", GetController().UpdateScatter)
		v1.PUT("/v1/scatters/:id", GetController().UpdateScatter)
		v1.DELETE("/v1/scatters/:id", GetController().DeleteScatter)

		v1.GET("/v1/series", GetController().GetSeries)
		v1.GET("/v1/series/:id", GetController().GetSerie)
		v1.POST("/v1/series", GetController().PostSerie)
		v1.PATCH("/v1/series/:id", GetController().UpdateSerie)
		v1.PUT("/v1/series/:id", GetController().UpdateSerie)
		v1.DELETE("/v1/series/:id", GetController().DeleteSerie)

		v1.GET("/v1/values", GetController().GetValues)
		v1.GET("/v1/values/:id", GetController().GetValue)
		v1.POST("/v1/values", GetController().PostValue)
		v1.PATCH("/v1/values/:id", GetController().UpdateValue)
		v1.PUT("/v1/values/:id", GetController().UpdateValue)
		v1.DELETE("/v1/values/:id", GetController().DeleteValue)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func(controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
