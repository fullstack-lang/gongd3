package models

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(pathToFile string) error {

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName := UnmarshallGongstructStaging(assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_Bar = make(map[string]*Bar)
var __gong__map_Key = make(map[string]*Key)
var __gong__map_Pie = make(map[string]*Pie)
var __gong__map_Scatter = make(map[string]*Scatter)
var __gong__map_Serie = make(map[string]*Serie)
var __gong__map_Value = make(map[string]*Value)

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {
	astCoordinate := "\tAssignStmt: "
	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "Bar":
										instanceBar := (&Bar{Name: instanceName}).Stage()
										instance = any(instanceBar)
										__gong__map_Bar[identifier] = instanceBar
									case "Key":
										instanceKey := (&Key{Name: instanceName}).Stage()
										instance = any(instanceKey)
										__gong__map_Key[identifier] = instanceKey
									case "Pie":
										instancePie := (&Pie{Name: instanceName}).Stage()
										instance = any(instancePie)
										__gong__map_Pie[identifier] = instancePie
									case "Scatter":
										instanceScatter := (&Scatter{Name: instanceName}).Stage()
										instance = any(instanceScatter)
										__gong__map_Scatter[identifier] = instanceScatter
									case "Serie":
										instanceSerie := (&Serie{Name: instanceName}).Stage()
										instance = any(instanceSerie)
										__gong__map_Serie[identifier] = instanceSerie
									case "Value":
										instanceValue := (&Value{Name: instanceName}).Stage()
										instance = any(instanceValue)
										__gong__map_Value[identifier] = instanceValue
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "Bar":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Key":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Pie":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Scatter":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Serie":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Value":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "Bar":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Set":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Serie[targetIdentifier]
							__gong__map_Bar[identifier].Set =
								append(__gong__map_Bar[identifier].Set, target)
						}
					case "Key":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Pie":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Set":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Serie[targetIdentifier]
							__gong__map_Pie[identifier].Set =
								append(__gong__map_Pie[identifier].Set, target)
						}
					case "Scatter":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Set":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Serie[targetIdentifier]
							__gong__map_Scatter[identifier].Set =
								append(__gong__map_Scatter[identifier].Set, target)
						}
					case "Serie":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Values":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Value[targetIdentifier]
							__gong__map_Serie[identifier].Values =
								append(__gong__map_Serie[identifier].Values, target)
						}
					case "Value":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "Bar":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bar[identifier].Name = fielValue
				case "XMin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].XMin = fielValue
				case "XMax":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].XMax = fielValue
				case "YMin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].YMin = fielValue
				case "YMax":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].YMax = fielValue
				case "YLabelOffset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].YLabelOffset = fielValue
				case "XLabelOffset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].XLabelOffset = fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].Width = fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].Heigth = fielValue
				case "Margin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].Margin = fielValue
				}
			case "Key":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Key[identifier].Name = fielValue
				}
			case "Pie":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pie[identifier].Name = fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Pie[identifier].Width = fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Pie[identifier].Heigth = fielValue
				case "Margin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Pie[identifier].Margin = fielValue
				}
			case "Scatter":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Scatter[identifier].Name = fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Scatter[identifier].Width = fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Scatter[identifier].Heigth = fielValue
				case "Margin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Scatter[identifier].Margin = fielValue
				}
			case "Serie":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Serie[identifier].Name = fielValue
				}
			case "Value":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Value[identifier].Name = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "Bar":
				switch fieldName {
				// insertion point for field dependant code
				case "X":
					targetIdentifier := ident.Name
					__gong__map_Bar[identifier].X = __gong__map_Key[targetIdentifier]
				case "Y":
					targetIdentifier := ident.Name
					__gong__map_Bar[identifier].Y = __gong__map_Key[targetIdentifier]
				case "AutoDomainX":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].AutoDomainX = fielValue
				case "AutoDomainY":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].AutoDomainY = fielValue
				case "YLabelPresent":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].YLabelPresent = fielValue
				case "XLabelPresent":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].XLabelPresent = fielValue
				}
			case "Key":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Pie":
				switch fieldName {
				// insertion point for field dependant code
				case "X":
					targetIdentifier := ident.Name
					__gong__map_Pie[identifier].X = __gong__map_Key[targetIdentifier]
				case "Y":
					targetIdentifier := ident.Name
					__gong__map_Pie[identifier].Y = __gong__map_Key[targetIdentifier]
				}
			case "Scatter":
				switch fieldName {
				// insertion point for field dependant code
				case "X":
					targetIdentifier := ident.Name
					__gong__map_Scatter[identifier].X = __gong__map_Key[targetIdentifier]
				case "Y":
					targetIdentifier := ident.Name
					__gong__map_Scatter[identifier].Y = __gong__map_Key[targetIdentifier]
				case "Text":
					targetIdentifier := ident.Name
					__gong__map_Scatter[identifier].Text = __gong__map_Key[targetIdentifier]
				}
			case "Serie":
				switch fieldName {
				// insertion point for field dependant code
				case "Key":
					targetIdentifier := ident.Name
					__gong__map_Serie[identifier].Key = __gong__map_Key[targetIdentifier]
				}
			case "Value":
				switch fieldName {
				// insertion point for field dependant code
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "Bar":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Key":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Pie":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Scatter":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Serie":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Value":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
