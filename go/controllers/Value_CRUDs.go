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
var __Value__dummysDeclaration__ models.Value
var __Value_time__dummyDeclaration time.Duration

var mutexValue sync.Mutex

// An ValueID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getValue updateValue deleteValue
type ValueID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ValueInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postValue updateValue
type ValueInput struct {
	// The Value to submit or modify
	// in: body
	Value *orm.ValueAPI
}

// GetValues
//
// swagger:route GET /values values getValues
//
// # Get all values
//
// Responses:
// default: genericError
//
//	200: valueDBResponse
func (controller *Controller) GetValues(c *gin.Context) {

	// source slice
	var valueDBs []orm.ValueDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetValues", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoValue.GetDB()

	_, err := db.Find(&valueDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	valueAPIs := make([]orm.ValueAPI, 0)

	// for each value, update fields from the database nullable fields
	for idx := range valueDBs {
		valueDB := &valueDBs[idx]
		_ = valueDB
		var valueAPI orm.ValueAPI

		// insertion point for updating fields
		valueAPI.ID = valueDB.ID
		valueDB.CopyBasicFieldsToValue_WOP(&valueAPI.Value_WOP)
		valueAPI.ValuePointersEncoding = valueDB.ValuePointersEncoding
		valueAPIs = append(valueAPIs, valueAPI)
	}

	c.JSON(http.StatusOK, valueAPIs)
}

// PostValue
//
// swagger:route POST /values values postValue
//
// Creates a value
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostValue(c *gin.Context) {

	mutexValue.Lock()
	defer mutexValue.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostValues", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoValue.GetDB()

	// Validate input
	var input orm.ValueAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create value
	valueDB := orm.ValueDB{}
	valueDB.ValuePointersEncoding = input.ValuePointersEncoding
	valueDB.CopyBasicFieldsFromValue_WOP(&input.Value_WOP)

	_, err = db.Create(&valueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoValue.CheckoutPhaseOneInstance(&valueDB)
	value := backRepo.BackRepoValue.Map_ValueDBID_ValuePtr[valueDB.ID]

	if value != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), value)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, valueDB)
}

// GetValue
//
// swagger:route GET /values/{ID} values getValue
//
// Gets the details for a value.
//
// Responses:
// default: genericError
//
//	200: valueDBResponse
func (controller *Controller) GetValue(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetValue", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoValue.GetDB()

	// Get valueDB in DB
	var valueDB orm.ValueDB
	if _, err := db.First(&valueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var valueAPI orm.ValueAPI
	valueAPI.ID = valueDB.ID
	valueAPI.ValuePointersEncoding = valueDB.ValuePointersEncoding
	valueDB.CopyBasicFieldsToValue_WOP(&valueAPI.Value_WOP)

	c.JSON(http.StatusOK, valueAPI)
}

// UpdateValue
//
// swagger:route PATCH /values/{ID} values updateValue
//
// # Update a value
//
// Responses:
// default: genericError
//
//	200: valueDBResponse
func (controller *Controller) UpdateValue(c *gin.Context) {

	mutexValue.Lock()
	defer mutexValue.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateValue", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoValue.GetDB()

	// Validate input
	var input orm.ValueAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var valueDB orm.ValueDB

	// fetch the value
	_, err := db.First(&valueDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	valueDB.CopyBasicFieldsFromValue_WOP(&input.Value_WOP)
	valueDB.ValuePointersEncoding = input.ValuePointersEncoding

	db, _ = db.Model(&valueDB)
	_, err = db.Updates(&valueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	valueNew := new(models.Value)
	valueDB.CopyBasicFieldsToValue(valueNew)

	// redeem pointers
	valueDB.DecodePointers(backRepo, valueNew)

	// get stage instance from DB instance, and call callback function
	valueOld := backRepo.BackRepoValue.Map_ValueDBID_ValuePtr[valueDB.ID]
	if valueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), valueOld, valueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the valueDB
	c.JSON(http.StatusOK, valueDB)
}

// DeleteValue
//
// swagger:route DELETE /values/{ID} values deleteValue
//
// # Delete a value
//
// default: genericError
//
//	200: valueDBResponse
func (controller *Controller) DeleteValue(c *gin.Context) {

	mutexValue.Lock()
	defer mutexValue.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteValue", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoValue.GetDB()

	// Get model if exist
	var valueDB orm.ValueDB
	if _, err := db.First(&valueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&valueDB)

	// get an instance (not staged) from DB instance, and call callback function
	valueDeleted := new(models.Value)
	valueDB.CopyBasicFieldsToValue(valueDeleted)

	// get stage instance from DB instance, and call callback function
	valueStaged := backRepo.BackRepoValue.Map_ValueDBID_ValuePtr[valueDB.ID]
	if valueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), valueStaged, valueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
