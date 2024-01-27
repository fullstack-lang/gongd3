import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { BarDB } from './bar-db'
import { Bar, CopyBarDBToBar } from './bar'
import { BarService } from './bar.service'

import { KeyDB } from './key-db'
import { Key, CopyKeyDBToKey } from './key'
import { KeyService } from './key.service'

import { PieDB } from './pie-db'
import { Pie, CopyPieDBToPie } from './pie'
import { PieService } from './pie.service'

import { ScatterDB } from './scatter-db'
import { Scatter, CopyScatterDBToScatter } from './scatter'
import { ScatterService } from './scatter.service'

import { SerieDB } from './serie-db'
import { Serie, CopySerieDBToSerie } from './serie'
import { SerieService } from './serie.service'

import { ValueDB } from './value-db'
import { Value, CopyValueDBToValue } from './value'
import { ValueService } from './value.service'

export const StackType = "github.com/fullstack-lang/gongd3/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Bars = new Array<Bar>() // array of front instances
	map_ID_Bar = new Map<number, Bar>() // map of front instances

	array_Keys = new Array<Key>() // array of front instances
	map_ID_Key = new Map<number, Key>() // map of front instances

	array_Pies = new Array<Pie>() // array of front instances
	map_ID_Pie = new Map<number, Pie>() // map of front instances

	array_Scatters = new Array<Scatter>() // array of front instances
	map_ID_Scatter = new Map<number, Scatter>() // map of front instances

	array_Series = new Array<Serie>() // array of front instances
	map_ID_Serie = new Map<number, Serie>() // map of front instances

	array_Values = new Array<Value>() // array of front instances
	map_ID_Value = new Map<number, Value>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Bar':
				return this.array_Bars as unknown as Array<Type>
			case 'Key':
				return this.array_Keys as unknown as Array<Type>
			case 'Pie':
				return this.array_Pies as unknown as Array<Type>
			case 'Scatter':
				return this.array_Scatters as unknown as Array<Type>
			case 'Serie':
				return this.array_Series as unknown as Array<Type>
			case 'Value':
				return this.array_Values as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}
	
	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Bar':
				return this.map_ID_Bar as unknown as Map<number, Type>
			case 'Key':
				return this.map_ID_Key as unknown as Map<number, Type>
			case 'Pie':
				return this.map_ID_Pie as unknown as Map<number, Type>
			case 'Scatter':
				return this.map_ID_Scatter as unknown as Map<number, Type>
			case 'Serie':
				return this.map_ID_Serie as unknown as Map<number, Type>
			case 'Value':
				return this.map_ID_Value as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private barService: BarService,
		private keyService: KeyService,
		private pieService: PieService,
		private scatterService: ScatterService,
		private serieService: SerieService,
		private valueService: ValueService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<BarDB[]>,
		Observable<KeyDB[]>,
		Observable<PieDB[]>,
		Observable<ScatterDB[]>,
		Observable<SerieDB[]>,
		Observable<ValueDB[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.barService.getBars(this.GONG__StackPath, this.frontRepo),
			this.keyService.getKeys(this.GONG__StackPath, this.frontRepo),
			this.pieService.getPies(this.GONG__StackPath, this.frontRepo),
			this.scatterService.getScatters(this.GONG__StackPath, this.frontRepo),
			this.serieService.getSeries(this.GONG__StackPath, this.frontRepo),
			this.valueService.getValues(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.barService.getBars(this.GONG__StackPath, this.frontRepo),
			this.keyService.getKeys(this.GONG__StackPath, this.frontRepo),
			this.pieService.getPies(this.GONG__StackPath, this.frontRepo),
			this.scatterService.getScatters(this.GONG__StackPath, this.frontRepo),
			this.serieService.getSeries(this.GONG__StackPath, this.frontRepo),
			this.valueService.getValues(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						bars_,
						keys_,
						pies_,
						scatters_,
						series_,
						values_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var bars: BarDB[]
						bars = bars_ as BarDB[]
						var keys: KeyDB[]
						keys = keys_ as KeyDB[]
						var pies: PieDB[]
						pies = pies_ as PieDB[]
						var scatters: ScatterDB[]
						scatters = scatters_ as ScatterDB[]
						var series: SerieDB[]
						series = series_ as SerieDB[]
						var values: ValueDB[]
						values = values_ as ValueDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Bars = []
						this.frontRepo.map_ID_Bar.clear()

						bars.forEach(
							barDB => {
								let bar = new Bar
								this.frontRepo.array_Bars.push(bar)
								this.frontRepo.map_ID_Bar.set(barDB.ID, bar)
							}
						)

						// init the arrays
						this.frontRepo.array_Keys = []
						this.frontRepo.map_ID_Key.clear()

						keys.forEach(
							keyDB => {
								let key = new Key
								this.frontRepo.array_Keys.push(key)
								this.frontRepo.map_ID_Key.set(keyDB.ID, key)
							}
						)

						// init the arrays
						this.frontRepo.array_Pies = []
						this.frontRepo.map_ID_Pie.clear()

						pies.forEach(
							pieDB => {
								let pie = new Pie
								this.frontRepo.array_Pies.push(pie)
								this.frontRepo.map_ID_Pie.set(pieDB.ID, pie)
							}
						)

						// init the arrays
						this.frontRepo.array_Scatters = []
						this.frontRepo.map_ID_Scatter.clear()

						scatters.forEach(
							scatterDB => {
								let scatter = new Scatter
								this.frontRepo.array_Scatters.push(scatter)
								this.frontRepo.map_ID_Scatter.set(scatterDB.ID, scatter)
							}
						)

						// init the arrays
						this.frontRepo.array_Series = []
						this.frontRepo.map_ID_Serie.clear()

						series.forEach(
							serieDB => {
								let serie = new Serie
								this.frontRepo.array_Series.push(serie)
								this.frontRepo.map_ID_Serie.set(serieDB.ID, serie)
							}
						)

						// init the arrays
						this.frontRepo.array_Values = []
						this.frontRepo.map_ID_Value.clear()

						values.forEach(
							valueDB => {
								let value = new Value
								this.frontRepo.array_Values.push(value)
								this.frontRepo.map_ID_Value.set(valueDB.ID, value)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						bars.forEach(
							barDB => {
								let bar = this.frontRepo.map_ID_Bar.get(barDB.ID)
								CopyBarDBToBar(barDB, bar!, this.frontRepo)
							}
						)

						// fill up front objects
						keys.forEach(
							keyDB => {
								let key = this.frontRepo.map_ID_Key.get(keyDB.ID)
								CopyKeyDBToKey(keyDB, key!, this.frontRepo)
							}
						)

						// fill up front objects
						pies.forEach(
							pieDB => {
								let pie = this.frontRepo.map_ID_Pie.get(pieDB.ID)
								CopyPieDBToPie(pieDB, pie!, this.frontRepo)
							}
						)

						// fill up front objects
						scatters.forEach(
							scatterDB => {
								let scatter = this.frontRepo.map_ID_Scatter.get(scatterDB.ID)
								CopyScatterDBToScatter(scatterDB, scatter!, this.frontRepo)
							}
						)

						// fill up front objects
						series.forEach(
							serieDB => {
								let serie = this.frontRepo.map_ID_Serie.get(serieDB.ID)
								CopySerieDBToSerie(serieDB, serie!, this.frontRepo)
							}
						)

						// fill up front objects
						values.forEach(
							valueDB => {
								let value = this.frontRepo.map_ID_Value.get(valueDB.ID)
								CopyValueDBToValue(valueDB, value!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}
}

// insertion point for get unique ID per struct 
export function getBarUniqueID(id: number): number {
	return 31 * id
}
export function getKeyUniqueID(id: number): number {
	return 37 * id
}
export function getPieUniqueID(id: number): number {
	return 41 * id
}
export function getScatterUniqueID(id: number): number {
	return 43 * id
}
export function getSerieUniqueID(id: number): number {
	return 47 * id
}
export function getValueUniqueID(id: number): number {
	return 53 * id
}
