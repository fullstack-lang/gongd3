package fullstack

import (
	// gongd3 stack for model analysis

	gongd3_controllers "gongd3/go/controllers"
	gongd3_models "gongd3/go/models"
	gongd3_orm "gongd3/go/orm"
	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "gongd3/ng/projects"
)

func Init(r *gin.Engine, filenames ...string) {

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	db_inMemory := gongd3_orm.SetupModels(&gongd3_models.Stage, false, filenames[0])

	// since gongd3sim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	gongd3_controllers.RegisterControllers(r)
}
