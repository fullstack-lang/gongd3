// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongd3/go/models"
	"github.com/fullstack-lang/gongd3/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Scatter__dummysDeclaration__ models.Scatter
var __Scatter_time__dummyDeclaration time.Duration

var mutexScatter sync.Mutex

// An ScatterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScatter updateScatter deleteScatter
type ScatterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ScatterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScatter updateScatter
type ScatterInput struct {
	// The Scatter to submit or modify
	// in: body
	Scatter *orm.ScatterAPI
}

// GetScatters
//
// swagger:route GET /scatters scatters getScatters
//
// # Get all scatters
//
// Responses:
// default: genericError
//
//	200: scatterDBResponse
func (controller *Controller) GetScatters(c *gin.Context) {

	// source slice
	var scatterDBs []orm.ScatterDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScatters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScatter.GetDB()

	query := db.Find(&scatterDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	scatterAPIs := make([]orm.ScatterAPI, 0)

	// for each scatter, update fields from the database nullable fields
	for idx := range scatterDBs {
		scatterDB := &scatterDBs[idx]
		_ = scatterDB
		var scatterAPI orm.ScatterAPI

		// insertion point for updating fields
		scatterAPI.ID = scatterDB.ID
		scatterDB.CopyBasicFieldsToScatter_WOP(&scatterAPI.Scatter_WOP)
		scatterAPI.ScatterPointersEncoding = scatterDB.ScatterPointersEncoding
		scatterAPIs = append(scatterAPIs, scatterAPI)
	}

	c.JSON(http.StatusOK, scatterAPIs)
}

// PostScatter
//
// swagger:route POST /scatters scatters postScatter
//
// Creates a scatter
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScatter(c *gin.Context) {

	mutexScatter.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScatters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScatter.GetDB()

	// Validate input
	var input orm.ScatterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create scatter
	scatterDB := orm.ScatterDB{}
	scatterDB.ScatterPointersEncoding = input.ScatterPointersEncoding
	scatterDB.CopyBasicFieldsFromScatter_WOP(&input.Scatter_WOP)

	query := db.Create(&scatterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScatter.CheckoutPhaseOneInstance(&scatterDB)
	scatter := backRepo.BackRepoScatter.Map_ScatterDBID_ScatterPtr[scatterDB.ID]

	if scatter != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), scatter)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, scatterDB)

	mutexScatter.Unlock()
}

// GetScatter
//
// swagger:route GET /scatters/{ID} scatters getScatter
//
// Gets the details for a scatter.
//
// Responses:
// default: genericError
//
//	200: scatterDBResponse
func (controller *Controller) GetScatter(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScatter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScatter.GetDB()

	// Get scatterDB in DB
	var scatterDB orm.ScatterDB
	if err := db.First(&scatterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var scatterAPI orm.ScatterAPI
	scatterAPI.ID = scatterDB.ID
	scatterAPI.ScatterPointersEncoding = scatterDB.ScatterPointersEncoding
	scatterDB.CopyBasicFieldsToScatter_WOP(&scatterAPI.Scatter_WOP)

	c.JSON(http.StatusOK, scatterAPI)
}

// UpdateScatter
//
// swagger:route PATCH /scatters/{ID} scatters updateScatter
//
// # Update a scatter
//
// Responses:
// default: genericError
//
//	200: scatterDBResponse
func (controller *Controller) UpdateScatter(c *gin.Context) {

	mutexScatter.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScatter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScatter.GetDB()

	// Validate input
	var input orm.ScatterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var scatterDB orm.ScatterDB

	// fetch the scatter
	query := db.First(&scatterDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	scatterDB.CopyBasicFieldsFromScatter_WOP(&input.Scatter_WOP)
	scatterDB.ScatterPointersEncoding = input.ScatterPointersEncoding

	query = db.Model(&scatterDB).Updates(scatterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	scatterNew := new(models.Scatter)
	scatterDB.CopyBasicFieldsToScatter(scatterNew)

	// redeem pointers
	scatterDB.DecodePointers(backRepo, scatterNew)

	// get stage instance from DB instance, and call callback function
	scatterOld := backRepo.BackRepoScatter.Map_ScatterDBID_ScatterPtr[scatterDB.ID]
	if scatterOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), scatterOld, scatterNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the scatterDB
	c.JSON(http.StatusOK, scatterDB)

	mutexScatter.Unlock()
}

// DeleteScatter
//
// swagger:route DELETE /scatters/{ID} scatters deleteScatter
//
// # Delete a scatter
//
// default: genericError
//
//	200: scatterDBResponse
func (controller *Controller) DeleteScatter(c *gin.Context) {

	mutexScatter.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScatter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScatter.GetDB()

	// Get model if exist
	var scatterDB orm.ScatterDB
	if err := db.First(&scatterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&scatterDB)

	// get an instance (not staged) from DB instance, and call callback function
	scatterDeleted := new(models.Scatter)
	scatterDB.CopyBasicFieldsToScatter(scatterDeleted)

	// get stage instance from DB instance, and call callback function
	scatterStaged := backRepo.BackRepoScatter.Map_ScatterDBID_ScatterPtr[scatterDB.ID]
	if scatterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), scatterStaged, scatterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexScatter.Unlock()
}
