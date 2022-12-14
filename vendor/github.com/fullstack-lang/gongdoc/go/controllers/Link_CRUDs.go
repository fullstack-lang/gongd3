// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Link__dummysDeclaration__ models.Link
var __Link_time__dummyDeclaration time.Duration

// An LinkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLink updateLink deleteLink
type LinkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LinkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLink updateLink
type LinkInput struct {
	// The Link to submit or modify
	// in: body
	Link *orm.LinkAPI
}

// GetLinks
//
// swagger:route GET /links links getLinks
//
// # Get all links
//
// Responses:
// default: genericError
//
//	200: linkDBResponse
func GetLinks(c *gin.Context) {
	db := orm.BackRepo.BackRepoLink.GetDB()

	// source slice
	var linkDBs []orm.LinkDB
	query := db.Find(&linkDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	linkAPIs := make([]orm.LinkAPI, 0)

	// for each link, update fields from the database nullable fields
	for idx := range linkDBs {
		linkDB := &linkDBs[idx]
		_ = linkDB
		var linkAPI orm.LinkAPI

		// insertion point for updating fields
		linkAPI.ID = linkDB.ID
		linkDB.CopyBasicFieldsToLink(&linkAPI.Link)
		linkAPI.LinkPointersEnconding = linkDB.LinkPointersEnconding
		linkAPIs = append(linkAPIs, linkAPI)
	}

	c.JSON(http.StatusOK, linkAPIs)
}

// PostLink
//
// swagger:route POST /links links postLink
//
// Creates a link
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostLink(c *gin.Context) {
	db := orm.BackRepo.BackRepoLink.GetDB()

	// Validate input
	var input orm.LinkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create link
	linkDB := orm.LinkDB{}
	linkDB.LinkPointersEnconding = input.LinkPointersEnconding
	linkDB.CopyBasicFieldsFromLink(&input.Link)

	query := db.Create(&linkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoLink.CheckoutPhaseOneInstance(&linkDB)
	link := (*orm.BackRepo.BackRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID]

	if link != nil {
		models.AfterCreateFromFront(&models.Stage, link)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, linkDB)
}

// GetLink
//
// swagger:route GET /links/{ID} links getLink
//
// Gets the details for a link.
//
// Responses:
// default: genericError
//
//	200: linkDBResponse
func GetLink(c *gin.Context) {
	db := orm.BackRepo.BackRepoLink.GetDB()

	// Get linkDB in DB
	var linkDB orm.LinkDB
	if err := db.First(&linkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var linkAPI orm.LinkAPI
	linkAPI.ID = linkDB.ID
	linkAPI.LinkPointersEnconding = linkDB.LinkPointersEnconding
	linkDB.CopyBasicFieldsToLink(&linkAPI.Link)

	c.JSON(http.StatusOK, linkAPI)
}

// UpdateLink
//
// swagger:route PATCH /links/{ID} links updateLink
//
// # Update a link
//
// Responses:
// default: genericError
//
//	200: linkDBResponse
func UpdateLink(c *gin.Context) {
	db := orm.BackRepo.BackRepoLink.GetDB()

	// Get model if exist
	var linkDB orm.LinkDB

	// fetch the link
	query := db.First(&linkDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.LinkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	linkDB.CopyBasicFieldsFromLink(&input.Link)
	linkDB.LinkPointersEnconding = input.LinkPointersEnconding

	query = db.Model(&linkDB).Updates(linkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	linkNew := new(models.Link)
	linkDB.CopyBasicFieldsToLink(linkNew)

	// get stage instance from DB instance, and call callback function
	linkOld := (*orm.BackRepo.BackRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID]
	if linkOld != nil {
		models.AfterUpdateFromFront(&models.Stage, linkOld, linkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the linkDB
	c.JSON(http.StatusOK, linkDB)
}

// DeleteLink
//
// swagger:route DELETE /links/{ID} links deleteLink
//
// # Delete a link
//
// default: genericError
//
//	200: linkDBResponse
func DeleteLink(c *gin.Context) {
	db := orm.BackRepo.BackRepoLink.GetDB()

	// Get model if exist
	var linkDB orm.LinkDB
	if err := db.First(&linkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&linkDB)

	// get an instance (not staged) from DB instance, and call callback function
	linkDeleted := new(models.Link)
	linkDB.CopyBasicFieldsToLink(linkDeleted)

	// get stage instance from DB instance, and call callback function
	linkStaged := (*orm.BackRepo.BackRepoLink.Map_LinkDBID_LinkPtr)[linkDB.ID]
	if linkStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, linkStaged, linkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
