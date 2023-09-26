package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	gongd3_go "github.com/fullstack-lang/gongd3/go"
	gongd3_fullstack "github.com/fullstack-lang/gongd3/go/fullstack"
	gongd3_models "github.com/fullstack-lang/gongd3/go/models"
	gongd3_orm "github.com/fullstack-lang/gongd3/go/orm"
	gongd3_probe "github.com/fullstack-lang/gongd3/go/probe"
	gongd3_static "github.com/fullstack-lang/gongd3/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gongd3_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gongd3/go/models", "main")
}

func main() {

	log.SetPrefix("gongd3: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongd3_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var stage *gongd3_models.StageStruct
	var backRepo *gongd3_orm.BackRepoStruct

	if *marshallOnCommit != "" {
		// persistence in a SQLite file on disk in memory
		stage, backRepo = gongd3_fullstack.NewStackInstance(r, "gongd3")
	} else {
		// persistence in a SQLite file on disk
		stage, backRepo = gongd3_fullstack.NewStackInstance(r, "gongd3", "./gongd3.db")
	}

	if *unmarshallFromCode != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		err := gongd3_models.ParseAstFile(stage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		stage.OnInitCommitCallback = hook
	}

	gongd3_probe.NewProbe(r, gongd3_go.GoModelsDir, gongd3_go.GoDiagramsDir, 
		*embeddedDiagrams,"gongd3", stage, backRepo)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
