package controllers

import (
	"sync"

	gongd3_orm "gongd3/go/orm"

	"github.com/gin-gonic/gin"
)

// A Controller is the handler of all API REST calls matching the stack model
// It forwards API requests to the stack instance identified by the GONG_StackPath parameters in the request
// the stack instance is the BackRepo instance
type Controller struct {

	// Map_BackRepos is the map to the backRepo instance according to the stack instance path
	Map_BackRepos map[string]*gongd3_orm.BackRepoStruct
}

var instance *Controller
var once sync.Once

func Register(r *gin.Engine) {
	once.Do(func() {
		RegisterControllers(r)
	})
}

func GetController() *Controller {
	once.Do(func() {
		instance = &Controller{
			Map_BackRepos: make(map[string]*gongd3_orm.BackRepoStruct),
		}
		instance.Map_BackRepos[""] = &gongd3_orm.BackRepo
	})
	return instance
}
