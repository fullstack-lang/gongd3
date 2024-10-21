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
var __Serie__dummysDeclaration__ models.Serie
var __Serie_time__dummyDeclaration time.Duration

var mutexSerie sync.Mutex

// An SerieID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSerie updateSerie deleteSerie
type SerieID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SerieInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSerie updateSerie
type SerieInput struct {
	// The Serie to submit or modify
	// in: body
	Serie *orm.SerieAPI
}

// GetSeries
//
// swagger:route GET /series series getSeries
//
// # Get all series
//
// Responses:
// default: genericError
//
//	200: serieDBResponse
func (controller *Controller) GetSeries(c *gin.Context) {

	// source slice
	var serieDBs []orm.SerieDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSeries", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSerie.GetDB()

	_, err := db.Find(&serieDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	serieAPIs := make([]orm.SerieAPI, 0)

	// for each serie, update fields from the database nullable fields
	for idx := range serieDBs {
		serieDB := &serieDBs[idx]
		_ = serieDB
		var serieAPI orm.SerieAPI

		// insertion point for updating fields
		serieAPI.ID = serieDB.ID
		serieDB.CopyBasicFieldsToSerie_WOP(&serieAPI.Serie_WOP)
		serieAPI.SeriePointersEncoding = serieDB.SeriePointersEncoding
		serieAPIs = append(serieAPIs, serieAPI)
	}

	c.JSON(http.StatusOK, serieAPIs)
}

// PostSerie
//
// swagger:route POST /series series postSerie
//
// Creates a serie
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSerie(c *gin.Context) {

	mutexSerie.Lock()
	defer mutexSerie.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSeries", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSerie.GetDB()

	// Validate input
	var input orm.SerieAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create serie
	serieDB := orm.SerieDB{}
	serieDB.SeriePointersEncoding = input.SeriePointersEncoding
	serieDB.CopyBasicFieldsFromSerie_WOP(&input.Serie_WOP)

	_, err = db.Create(&serieDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSerie.CheckoutPhaseOneInstance(&serieDB)
	serie := backRepo.BackRepoSerie.Map_SerieDBID_SeriePtr[serieDB.ID]

	if serie != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), serie)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, serieDB)
}

// GetSerie
//
// swagger:route GET /series/{ID} series getSerie
//
// Gets the details for a serie.
//
// Responses:
// default: genericError
//
//	200: serieDBResponse
func (controller *Controller) GetSerie(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSerie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSerie.GetDB()

	// Get serieDB in DB
	var serieDB orm.SerieDB
	if _, err := db.First(&serieDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var serieAPI orm.SerieAPI
	serieAPI.ID = serieDB.ID
	serieAPI.SeriePointersEncoding = serieDB.SeriePointersEncoding
	serieDB.CopyBasicFieldsToSerie_WOP(&serieAPI.Serie_WOP)

	c.JSON(http.StatusOK, serieAPI)
}

// UpdateSerie
//
// swagger:route PATCH /series/{ID} series updateSerie
//
// # Update a serie
//
// Responses:
// default: genericError
//
//	200: serieDBResponse
func (controller *Controller) UpdateSerie(c *gin.Context) {

	mutexSerie.Lock()
	defer mutexSerie.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSerie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSerie.GetDB()

	// Validate input
	var input orm.SerieAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var serieDB orm.SerieDB

	// fetch the serie
	_, err := db.First(&serieDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	serieDB.CopyBasicFieldsFromSerie_WOP(&input.Serie_WOP)
	serieDB.SeriePointersEncoding = input.SeriePointersEncoding

	db, _ = db.Model(&serieDB)
	_, err = db.Updates(serieDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	serieNew := new(models.Serie)
	serieDB.CopyBasicFieldsToSerie(serieNew)

	// redeem pointers
	serieDB.DecodePointers(backRepo, serieNew)

	// get stage instance from DB instance, and call callback function
	serieOld := backRepo.BackRepoSerie.Map_SerieDBID_SeriePtr[serieDB.ID]
	if serieOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), serieOld, serieNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the serieDB
	c.JSON(http.StatusOK, serieDB)
}

// DeleteSerie
//
// swagger:route DELETE /series/{ID} series deleteSerie
//
// # Delete a serie
//
// default: genericError
//
//	200: serieDBResponse
func (controller *Controller) DeleteSerie(c *gin.Context) {

	mutexSerie.Lock()
	defer mutexSerie.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSerie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongd3/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSerie.GetDB()

	// Get model if exist
	var serieDB orm.SerieDB
	if _, err := db.First(&serieDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&serieDB)

	// get an instance (not staged) from DB instance, and call callback function
	serieDeleted := new(models.Serie)
	serieDB.CopyBasicFieldsToSerie(serieDeleted)

	// get stage instance from DB instance, and call callback function
	serieStaged := backRepo.BackRepoSerie.Map_SerieDBID_SeriePtr[serieDB.ID]
	if serieStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), serieStaged, serieDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
