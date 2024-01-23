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
var __Pie__dummysDeclaration__ models.Pie
var __Pie_time__dummyDeclaration time.Duration

var mutexPie sync.Mutex

// An PieID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPie updatePie deletePie
type PieID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PieInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPie updatePie
type PieInput struct {
	// The Pie to submit or modify
	// in: body
	Pie *orm.PieAPI
}

// GetPies
//
// swagger:route GET /pies pies getPies
//
// # Get all pies
//
// Responses:
// default: genericError
//
//	200: pieDBResponse
func (controller *Controller) GetPies(c *gin.Context) {

	// source slice
	var pieDBs []orm.PieDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPies", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPie.GetDB()

	query := db.Find(&pieDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pieAPIs := make([]orm.PieAPI, 0)

	// for each pie, update fields from the database nullable fields
	for idx := range pieDBs {
		pieDB := &pieDBs[idx]
		_ = pieDB
		var pieAPI orm.PieAPI

		// insertion point for updating fields
		pieAPI.ID = pieDB.ID
		pieDB.CopyBasicFieldsToPie_WOP(&pieAPI.Pie_WOP)
		pieAPI.PiePointersEncoding = pieDB.PiePointersEncoding
		pieAPIs = append(pieAPIs, pieAPI)
	}

	c.JSON(http.StatusOK, pieAPIs)
}

// PostPie
//
// swagger:route POST /pies pies postPie
//
// Creates a pie
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPie(c *gin.Context) {

	mutexPie.Lock()
	defer mutexPie.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPies", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPie.GetDB()

	// Validate input
	var input orm.PieAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pie
	pieDB := orm.PieDB{}
	pieDB.PiePointersEncoding = input.PiePointersEncoding
	pieDB.CopyBasicFieldsFromPie_WOP(&input.Pie_WOP)

	query := db.Create(&pieDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPie.CheckoutPhaseOneInstance(&pieDB)
	pie := backRepo.BackRepoPie.Map_PieDBID_PiePtr[pieDB.ID]

	if pie != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pie)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pieDB)
}

// GetPie
//
// swagger:route GET /pies/{ID} pies getPie
//
// Gets the details for a pie.
//
// Responses:
// default: genericError
//
//	200: pieDBResponse
func (controller *Controller) GetPie(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPie.GetDB()

	// Get pieDB in DB
	var pieDB orm.PieDB
	if err := db.First(&pieDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pieAPI orm.PieAPI
	pieAPI.ID = pieDB.ID
	pieAPI.PiePointersEncoding = pieDB.PiePointersEncoding
	pieDB.CopyBasicFieldsToPie_WOP(&pieAPI.Pie_WOP)

	c.JSON(http.StatusOK, pieAPI)
}

// UpdatePie
//
// swagger:route PATCH /pies/{ID} pies updatePie
//
// # Update a pie
//
// Responses:
// default: genericError
//
//	200: pieDBResponse
func (controller *Controller) UpdatePie(c *gin.Context) {

	mutexPie.Lock()
	defer mutexPie.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPie.GetDB()

	// Validate input
	var input orm.PieAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pieDB orm.PieDB

	// fetch the pie
	query := db.First(&pieDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pieDB.CopyBasicFieldsFromPie_WOP(&input.Pie_WOP)
	pieDB.PiePointersEncoding = input.PiePointersEncoding

	query = db.Model(&pieDB).Updates(pieDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pieNew := new(models.Pie)
	pieDB.CopyBasicFieldsToPie(pieNew)

	// redeem pointers
	pieDB.DecodePointers(backRepo, pieNew)

	// get stage instance from DB instance, and call callback function
	pieOld := backRepo.BackRepoPie.Map_PieDBID_PiePtr[pieDB.ID]
	if pieOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pieOld, pieNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pieDB
	c.JSON(http.StatusOK, pieDB)
}

// DeletePie
//
// swagger:route DELETE /pies/{ID} pies deletePie
//
// # Delete a pie
//
// default: genericError
//
//	200: pieDBResponse
func (controller *Controller) DeletePie(c *gin.Context) {

	mutexPie.Lock()
	defer mutexPie.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPie.GetDB()

	// Get model if exist
	var pieDB orm.PieDB
	if err := db.First(&pieDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&pieDB)

	// get an instance (not staged) from DB instance, and call callback function
	pieDeleted := new(models.Pie)
	pieDB.CopyBasicFieldsToPie(pieDeleted)

	// get stage instance from DB instance, and call callback function
	pieStaged := backRepo.BackRepoPie.Map_PieDBID_PiePtr[pieDB.ID]
	if pieStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pieStaged, pieDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
