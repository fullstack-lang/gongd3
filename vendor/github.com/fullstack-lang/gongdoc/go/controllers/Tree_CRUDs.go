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
var __Tree__dummysDeclaration__ models.Tree
var __Tree_time__dummyDeclaration time.Duration

// An TreeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTree updateTree deleteTree
type TreeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TreeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTree updateTree
type TreeInput struct {
	// The Tree to submit or modify
	// in: body
	Tree *orm.TreeAPI
}

// GetTrees
//
// swagger:route GET /trees trees getTrees
//
// # Get all trees
//
// Responses:
// default: genericError
//
//	200: treeDBResponse
func GetTrees(c *gin.Context) {
	db := orm.BackRepo.BackRepoTree.GetDB()

	// source slice
	var treeDBs []orm.TreeDB
	query := db.Find(&treeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	treeAPIs := make([]orm.TreeAPI, 0)

	// for each tree, update fields from the database nullable fields
	for idx := range treeDBs {
		treeDB := &treeDBs[idx]
		_ = treeDB
		var treeAPI orm.TreeAPI

		// insertion point for updating fields
		treeAPI.ID = treeDB.ID
		treeDB.CopyBasicFieldsToTree(&treeAPI.Tree)
		treeAPI.TreePointersEnconding = treeDB.TreePointersEnconding
		treeAPIs = append(treeAPIs, treeAPI)
	}

	c.JSON(http.StatusOK, treeAPIs)
}

// PostTree
//
// swagger:route POST /trees trees postTree
//
// Creates a tree
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func PostTree(c *gin.Context) {
	db := orm.BackRepo.BackRepoTree.GetDB()

	// Validate input
	var input orm.TreeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tree
	treeDB := orm.TreeDB{}
	treeDB.TreePointersEnconding = input.TreePointersEnconding
	treeDB.CopyBasicFieldsFromTree(&input.Tree)

	query := db.Create(&treeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	orm.BackRepo.BackRepoTree.CheckoutPhaseOneInstance(&treeDB)
	tree := (*orm.BackRepo.BackRepoTree.Map_TreeDBID_TreePtr)[treeDB.ID]

	if tree != nil {
		models.AfterCreateFromFront(&models.Stage, tree)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, treeDB)
}

// GetTree
//
// swagger:route GET /trees/{ID} trees getTree
//
// Gets the details for a tree.
//
// Responses:
// default: genericError
//
//	200: treeDBResponse
func GetTree(c *gin.Context) {
	db := orm.BackRepo.BackRepoTree.GetDB()

	// Get treeDB in DB
	var treeDB orm.TreeDB
	if err := db.First(&treeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var treeAPI orm.TreeAPI
	treeAPI.ID = treeDB.ID
	treeAPI.TreePointersEnconding = treeDB.TreePointersEnconding
	treeDB.CopyBasicFieldsToTree(&treeAPI.Tree)

	c.JSON(http.StatusOK, treeAPI)
}

// UpdateTree
//
// swagger:route PATCH /trees/{ID} trees updateTree
//
// # Update a tree
//
// Responses:
// default: genericError
//
//	200: treeDBResponse
func UpdateTree(c *gin.Context) {
	db := orm.BackRepo.BackRepoTree.GetDB()

	// Get model if exist
	var treeDB orm.TreeDB

	// fetch the tree
	query := db.First(&treeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.TreeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	treeDB.CopyBasicFieldsFromTree(&input.Tree)
	treeDB.TreePointersEnconding = input.TreePointersEnconding

	query = db.Model(&treeDB).Updates(treeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	treeNew := new(models.Tree)
	treeDB.CopyBasicFieldsToTree(treeNew)

	// get stage instance from DB instance, and call callback function
	treeOld := (*orm.BackRepo.BackRepoTree.Map_TreeDBID_TreePtr)[treeDB.ID]
	if treeOld != nil {
		models.AfterUpdateFromFront(&models.Stage, treeOld, treeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the treeDB
	c.JSON(http.StatusOK, treeDB)
}

// DeleteTree
//
// swagger:route DELETE /trees/{ID} trees deleteTree
//
// # Delete a tree
//
// default: genericError
//
//	200: treeDBResponse
func DeleteTree(c *gin.Context) {
	db := orm.BackRepo.BackRepoTree.GetDB()

	// Get model if exist
	var treeDB orm.TreeDB
	if err := db.First(&treeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&treeDB)

	// get an instance (not staged) from DB instance, and call callback function
	treeDeleted := new(models.Tree)
	treeDB.CopyBasicFieldsToTree(treeDeleted)

	// get stage instance from DB instance, and call callback function
	treeStaged := (*orm.BackRepo.BackRepoTree.Map_TreeDBID_TreePtr)[treeDB.ID]
	if treeStaged != nil {
		models.AfterDeleteFromFront(&models.Stage, treeStaged, treeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
