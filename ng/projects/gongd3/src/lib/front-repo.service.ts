import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { BarDB } from './bar-db'
import { BarService } from './bar.service'

import { KeyDB } from './key-db'
import { KeyService } from './key.service'

import { PieDB } from './pie-db'
import { PieService } from './pie.service'

import { ScatterDB } from './scatter-db'
import { ScatterService } from './scatter.service'

import { SerieDB } from './serie-db'
import { SerieService } from './serie.service'

import { ValueDB } from './value-db'
import { ValueService } from './value.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  Bars_array = new Array<BarDB>() // array of repo instances
  Bars = new Map<number, BarDB>() // map of repo instances
  Bars_batch = new Map<number, BarDB>() // same but only in last GET (for finding repo instances to delete)

  Keys_array = new Array<KeyDB>() // array of repo instances
  Keys = new Map<number, KeyDB>() // map of repo instances
  Keys_batch = new Map<number, KeyDB>() // same but only in last GET (for finding repo instances to delete)

  Pies_array = new Array<PieDB>() // array of repo instances
  Pies = new Map<number, PieDB>() // map of repo instances
  Pies_batch = new Map<number, PieDB>() // same but only in last GET (for finding repo instances to delete)

  Scatters_array = new Array<ScatterDB>() // array of repo instances
  Scatters = new Map<number, ScatterDB>() // map of repo instances
  Scatters_batch = new Map<number, ScatterDB>() // same but only in last GET (for finding repo instances to delete)

  Series_array = new Array<SerieDB>() // array of repo instances
  Series = new Map<number, SerieDB>() // map of repo instances
  Series_batch = new Map<number, SerieDB>() // same but only in last GET (for finding repo instances to delete)

  Values_array = new Array<ValueDB>() // array of repo instances
  Values = new Map<number, ValueDB>() // map of repo instances
  Values_batch = new Map<number, ValueDB>() // same but only in last GET (for finding repo instances to delete)


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'Bar':
        return this.Bars_array as unknown as Array<Type>
      case 'Key':
        return this.Keys_array as unknown as Array<Type>
      case 'Pie':
        return this.Pies_array as unknown as Array<Type>
      case 'Scatter':
        return this.Scatters_array as unknown as Array<Type>
      case 'Serie':
        return this.Series_array as unknown as Array<Type>
      case 'Value':
        return this.Values_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'Bar':
        return this.Bars_array as unknown as Map<number, Type>
      case 'Key':
        return this.Keys_array as unknown as Map<number, Type>
      case 'Pie':
        return this.Pies_array as unknown as Map<number, Type>
      case 'Scatter':
        return this.Scatters_array as unknown as Map<number, Type>
      case 'Serie':
        return this.Series_array as unknown as Map<number, Type>
      case 'Value':
        return this.Values_array as unknown as Map<number, Type>
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
  SourceStruct: string = ""  // The "Aclass"
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
      this.barService.getBars(this.GONG__StackPath),
      this.keyService.getKeys(this.GONG__StackPath),
      this.pieService.getPies(this.GONG__StackPath),
      this.scatterService.getScatters(this.GONG__StackPath),
      this.serieService.getSeries(this.GONG__StackPath),
      this.valueService.getValues(this.GONG__StackPath),
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
      this.barService.getBars(this.GONG__StackPath),
      this.keyService.getKeys(this.GONG__StackPath),
      this.pieService.getPies(this.GONG__StackPath),
      this.scatterService.getScatters(this.GONG__StackPath),
      this.serieService.getSeries(this.GONG__StackPath),
      this.valueService.getValues(this.GONG__StackPath),
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
            // init the array
            this.frontRepo.Bars_array = bars

            // clear the map that counts Bar in the GET
            this.frontRepo.Bars_batch.clear()

            bars.forEach(
              bar => {
                this.frontRepo.Bars.set(bar.ID, bar)
                this.frontRepo.Bars_batch.set(bar.ID, bar)
              }
            )

            // clear bars that are absent from the batch
            this.frontRepo.Bars.forEach(
              bar => {
                if (this.frontRepo.Bars_batch.get(bar.ID) == undefined) {
                  this.frontRepo.Bars.delete(bar.ID)
                }
              }
            )

            // sort Bars_array array
            this.frontRepo.Bars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Keys_array = keys

            // clear the map that counts Key in the GET
            this.frontRepo.Keys_batch.clear()

            keys.forEach(
              key => {
                this.frontRepo.Keys.set(key.ID, key)
                this.frontRepo.Keys_batch.set(key.ID, key)
              }
            )

            // clear keys that are absent from the batch
            this.frontRepo.Keys.forEach(
              key => {
                if (this.frontRepo.Keys_batch.get(key.ID) == undefined) {
                  this.frontRepo.Keys.delete(key.ID)
                }
              }
            )

            // sort Keys_array array
            this.frontRepo.Keys_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Pies_array = pies

            // clear the map that counts Pie in the GET
            this.frontRepo.Pies_batch.clear()

            pies.forEach(
              pie => {
                this.frontRepo.Pies.set(pie.ID, pie)
                this.frontRepo.Pies_batch.set(pie.ID, pie)
              }
            )

            // clear pies that are absent from the batch
            this.frontRepo.Pies.forEach(
              pie => {
                if (this.frontRepo.Pies_batch.get(pie.ID) == undefined) {
                  this.frontRepo.Pies.delete(pie.ID)
                }
              }
            )

            // sort Pies_array array
            this.frontRepo.Pies_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Scatters_array = scatters

            // clear the map that counts Scatter in the GET
            this.frontRepo.Scatters_batch.clear()

            scatters.forEach(
              scatter => {
                this.frontRepo.Scatters.set(scatter.ID, scatter)
                this.frontRepo.Scatters_batch.set(scatter.ID, scatter)
              }
            )

            // clear scatters that are absent from the batch
            this.frontRepo.Scatters.forEach(
              scatter => {
                if (this.frontRepo.Scatters_batch.get(scatter.ID) == undefined) {
                  this.frontRepo.Scatters.delete(scatter.ID)
                }
              }
            )

            // sort Scatters_array array
            this.frontRepo.Scatters_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Series_array = series

            // clear the map that counts Serie in the GET
            this.frontRepo.Series_batch.clear()

            series.forEach(
              serie => {
                this.frontRepo.Series.set(serie.ID, serie)
                this.frontRepo.Series_batch.set(serie.ID, serie)
              }
            )

            // clear series that are absent from the batch
            this.frontRepo.Series.forEach(
              serie => {
                if (this.frontRepo.Series_batch.get(serie.ID) == undefined) {
                  this.frontRepo.Series.delete(serie.ID)
                }
              }
            )

            // sort Series_array array
            this.frontRepo.Series_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Values_array = values

            // clear the map that counts Value in the GET
            this.frontRepo.Values_batch.clear()

            values.forEach(
              value => {
                this.frontRepo.Values.set(value.ID, value)
                this.frontRepo.Values_batch.set(value.ID, value)
              }
            )

            // clear values that are absent from the batch
            this.frontRepo.Values.forEach(
              value => {
                if (this.frontRepo.Values_batch.get(value.ID) == undefined) {
                  this.frontRepo.Values.delete(value.ID)
                }
              }
            )

            // sort Values_array array
            this.frontRepo.Values_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            bars.forEach(
              bar => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field X redeeming
                {
                  let _key = this.frontRepo.Keys.get(bar.XID.Int64)
                  if (_key) {
                    bar.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = this.frontRepo.Keys.get(bar.YID.Int64)
                  if (_key) {
                    bar.Y = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            keys.forEach(
              key => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            pies.forEach(
              pie => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field X redeeming
                {
                  let _key = this.frontRepo.Keys.get(pie.XID.Int64)
                  if (_key) {
                    pie.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = this.frontRepo.Keys.get(pie.YID.Int64)
                  if (_key) {
                    pie.Y = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            scatters.forEach(
              scatter => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field X redeeming
                {
                  let _key = this.frontRepo.Keys.get(scatter.XID.Int64)
                  if (_key) {
                    scatter.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = this.frontRepo.Keys.get(scatter.YID.Int64)
                  if (_key) {
                    scatter.Y = _key
                  }
                }
                // insertion point for pointer field Text redeeming
                {
                  let _key = this.frontRepo.Keys.get(scatter.TextID.Int64)
                  if (_key) {
                    scatter.Text = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            series.forEach(
              serie => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Key redeeming
                {
                  let _key = this.frontRepo.Keys.get(serie.KeyID.Int64)
                  if (_key) {
                    serie.Key = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Bar.Set redeeming
                {
                  let _bar = this.frontRepo.Bars.get(serie.Bar_SetDBID.Int64)
                  if (_bar) {
                    if (_bar.Set == undefined) {
                      _bar.Set = new Array<SerieDB>()
                    }
                    _bar.Set.push(serie)
                    if (serie.Bar_Set_reverse == undefined) {
                      serie.Bar_Set_reverse = _bar
                    }
                  }
                }
                // insertion point for slice of pointer field Pie.Set redeeming
                {
                  let _pie = this.frontRepo.Pies.get(serie.Pie_SetDBID.Int64)
                  if (_pie) {
                    if (_pie.Set == undefined) {
                      _pie.Set = new Array<SerieDB>()
                    }
                    _pie.Set.push(serie)
                    if (serie.Pie_Set_reverse == undefined) {
                      serie.Pie_Set_reverse = _pie
                    }
                  }
                }
                // insertion point for slice of pointer field Scatter.Set redeeming
                {
                  let _scatter = this.frontRepo.Scatters.get(serie.Scatter_SetDBID.Int64)
                  if (_scatter) {
                    if (_scatter.Set == undefined) {
                      _scatter.Set = new Array<SerieDB>()
                    }
                    _scatter.Set.push(serie)
                    if (serie.Scatter_Set_reverse == undefined) {
                      serie.Scatter_Set_reverse = _scatter
                    }
                  }
                }
              }
            )
            values.forEach(
              value => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Serie.Values redeeming
                {
                  let _serie = this.frontRepo.Series.get(value.Serie_ValuesDBID.Int64)
                  if (_serie) {
                    if (_serie.Values == undefined) {
                      _serie.Values = new Array<ValueDB>()
                    }
                    _serie.Values.push(value)
                    if (value.Serie_Values_reverse == undefined) {
                      value.Serie_Values_reverse = _serie
                    }
                  }
                }
              }
            )

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem 
            bars.forEach(
              bar => {
                // insertion point for sorting
                bar.Set?.sort((t1, t2) => {
                  if (t1.Bar_SetDBID_Index.Int64 > t2.Bar_SetDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Bar_SetDBID_Index.Int64 < t2.Bar_SetDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            keys.forEach(
              key => {
                // insertion point for sorting
              }
            )
            pies.forEach(
              pie => {
                // insertion point for sorting
                pie.Set?.sort((t1, t2) => {
                  if (t1.Pie_SetDBID_Index.Int64 > t2.Pie_SetDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Pie_SetDBID_Index.Int64 < t2.Pie_SetDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            scatters.forEach(
              scatter => {
                // insertion point for sorting
                scatter.Set?.sort((t1, t2) => {
                  if (t1.Scatter_SetDBID_Index.Int64 > t2.Scatter_SetDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Scatter_SetDBID_Index.Int64 < t2.Scatter_SetDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            series.forEach(
              serie => {
                // insertion point for sorting
                serie.Values?.sort((t1, t2) => {
                  if (t1.Serie_ValuesDBID_Index.Int64 > t2.Serie_ValuesDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Serie_ValuesDBID_Index.Int64 < t2.Serie_ValuesDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            values.forEach(
              value => {
                // insertion point for sorting
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // BarPull performs a GET on Bar of the stack and redeem association pointers 
  BarPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.barService.getBars(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            bars,
          ]) => {
            // init the array
            this.frontRepo.Bars_array = bars

            // clear the map that counts Bar in the GET
            this.frontRepo.Bars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bars.forEach(
              bar => {
                this.frontRepo.Bars.set(bar.ID, bar)
                this.frontRepo.Bars_batch.set(bar.ID, bar)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field X redeeming
                {
                  let _key = this.frontRepo.Keys.get(bar.XID.Int64)
                  if (_key) {
                    bar.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = this.frontRepo.Keys.get(bar.YID.Int64)
                  if (_key) {
                    bar.Y = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear bars that are absent from the GET
            this.frontRepo.Bars.forEach(
              bar => {
                if (this.frontRepo.Bars_batch.get(bar.ID) == undefined) {
                  this.frontRepo.Bars.delete(bar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // KeyPull performs a GET on Key of the stack and redeem association pointers 
  KeyPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.keyService.getKeys(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            keys,
          ]) => {
            // init the array
            this.frontRepo.Keys_array = keys

            // clear the map that counts Key in the GET
            this.frontRepo.Keys_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            keys.forEach(
              key => {
                this.frontRepo.Keys.set(key.ID, key)
                this.frontRepo.Keys_batch.set(key.ID, key)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear keys that are absent from the GET
            this.frontRepo.Keys.forEach(
              key => {
                if (this.frontRepo.Keys_batch.get(key.ID) == undefined) {
                  this.frontRepo.Keys.delete(key.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // PiePull performs a GET on Pie of the stack and redeem association pointers 
  PiePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pieService.getPies(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            pies,
          ]) => {
            // init the array
            this.frontRepo.Pies_array = pies

            // clear the map that counts Pie in the GET
            this.frontRepo.Pies_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            pies.forEach(
              pie => {
                this.frontRepo.Pies.set(pie.ID, pie)
                this.frontRepo.Pies_batch.set(pie.ID, pie)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field X redeeming
                {
                  let _key = this.frontRepo.Keys.get(pie.XID.Int64)
                  if (_key) {
                    pie.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = this.frontRepo.Keys.get(pie.YID.Int64)
                  if (_key) {
                    pie.Y = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear pies that are absent from the GET
            this.frontRepo.Pies.forEach(
              pie => {
                if (this.frontRepo.Pies_batch.get(pie.ID) == undefined) {
                  this.frontRepo.Pies.delete(pie.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // ScatterPull performs a GET on Scatter of the stack and redeem association pointers 
  ScatterPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.scatterService.getScatters(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            scatters,
          ]) => {
            // init the array
            this.frontRepo.Scatters_array = scatters

            // clear the map that counts Scatter in the GET
            this.frontRepo.Scatters_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            scatters.forEach(
              scatter => {
                this.frontRepo.Scatters.set(scatter.ID, scatter)
                this.frontRepo.Scatters_batch.set(scatter.ID, scatter)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field X redeeming
                {
                  let _key = this.frontRepo.Keys.get(scatter.XID.Int64)
                  if (_key) {
                    scatter.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = this.frontRepo.Keys.get(scatter.YID.Int64)
                  if (_key) {
                    scatter.Y = _key
                  }
                }
                // insertion point for pointer field Text redeeming
                {
                  let _key = this.frontRepo.Keys.get(scatter.TextID.Int64)
                  if (_key) {
                    scatter.Text = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear scatters that are absent from the GET
            this.frontRepo.Scatters.forEach(
              scatter => {
                if (this.frontRepo.Scatters_batch.get(scatter.ID) == undefined) {
                  this.frontRepo.Scatters.delete(scatter.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // SeriePull performs a GET on Serie of the stack and redeem association pointers 
  SeriePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.serieService.getSeries(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            series,
          ]) => {
            // init the array
            this.frontRepo.Series_array = series

            // clear the map that counts Serie in the GET
            this.frontRepo.Series_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            series.forEach(
              serie => {
                this.frontRepo.Series.set(serie.ID, serie)
                this.frontRepo.Series_batch.set(serie.ID, serie)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Key redeeming
                {
                  let _key = this.frontRepo.Keys.get(serie.KeyID.Int64)
                  if (_key) {
                    serie.Key = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Bar.Set redeeming
                {
                  let _bar = this.frontRepo.Bars.get(serie.Bar_SetDBID.Int64)
                  if (_bar) {
                    if (_bar.Set == undefined) {
                      _bar.Set = new Array<SerieDB>()
                    }
                    _bar.Set.push(serie)
                    if (serie.Bar_Set_reverse == undefined) {
                      serie.Bar_Set_reverse = _bar
                    }
                  }
                }
                // insertion point for slice of pointer field Pie.Set redeeming
                {
                  let _pie = this.frontRepo.Pies.get(serie.Pie_SetDBID.Int64)
                  if (_pie) {
                    if (_pie.Set == undefined) {
                      _pie.Set = new Array<SerieDB>()
                    }
                    _pie.Set.push(serie)
                    if (serie.Pie_Set_reverse == undefined) {
                      serie.Pie_Set_reverse = _pie
                    }
                  }
                }
                // insertion point for slice of pointer field Scatter.Set redeeming
                {
                  let _scatter = this.frontRepo.Scatters.get(serie.Scatter_SetDBID.Int64)
                  if (_scatter) {
                    if (_scatter.Set == undefined) {
                      _scatter.Set = new Array<SerieDB>()
                    }
                    _scatter.Set.push(serie)
                    if (serie.Scatter_Set_reverse == undefined) {
                      serie.Scatter_Set_reverse = _scatter
                    }
                  }
                }
              }
            )

            // clear series that are absent from the GET
            this.frontRepo.Series.forEach(
              serie => {
                if (this.frontRepo.Series_batch.get(serie.ID) == undefined) {
                  this.frontRepo.Series.delete(serie.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // ValuePull performs a GET on Value of the stack and redeem association pointers 
  ValuePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.valueService.getValues(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            values,
          ]) => {
            // init the array
            this.frontRepo.Values_array = values

            // clear the map that counts Value in the GET
            this.frontRepo.Values_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            values.forEach(
              value => {
                this.frontRepo.Values.set(value.ID, value)
                this.frontRepo.Values_batch.set(value.ID, value)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Serie.Values redeeming
                {
                  let _serie = this.frontRepo.Series.get(value.Serie_ValuesDBID.Int64)
                  if (_serie) {
                    if (_serie.Values == undefined) {
                      _serie.Values = new Array<ValueDB>()
                    }
                    _serie.Values.push(value)
                    if (value.Serie_Values_reverse == undefined) {
                      value.Serie_Values_reverse = _serie
                    }
                  }
                }
              }
            )

            // clear values that are absent from the GET
            this.frontRepo.Values.forEach(
              value => {
                if (this.frontRepo.Values_batch.get(value.ID) == undefined) {
                  this.frontRepo.Values.delete(value.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

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
