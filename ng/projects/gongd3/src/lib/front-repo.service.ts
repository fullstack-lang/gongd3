import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { BarDB } from './bar-db'
import { BarService } from './bar.service'

import { KeyDB } from './key-db'
import { KeyService } from './key.service'

import { SerieDB } from './serie-db'
import { SerieService } from './serie.service'

import { ValueDB } from './value-db'
import { ValueService } from './value.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Bars_array = new Array<BarDB>(); // array of repo instances
  Bars = new Map<number, BarDB>(); // map of repo instances
  Bars_batch = new Map<number, BarDB>(); // same but only in last GET (for finding repo instances to delete)
  Keys_array = new Array<KeyDB>(); // array of repo instances
  Keys = new Map<number, KeyDB>(); // map of repo instances
  Keys_batch = new Map<number, KeyDB>(); // same but only in last GET (for finding repo instances to delete)
  Series_array = new Array<SerieDB>(); // array of repo instances
  Series = new Map<number, SerieDB>(); // map of repo instances
  Series_batch = new Map<number, SerieDB>(); // same but only in last GET (for finding repo instances to delete)
  Values_array = new Array<ValueDB>(); // array of repo instances
  Values = new Map<number, ValueDB>(); // map of repo instances
  Values_batch = new Map<number, ValueDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

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

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private barService: BarService,
    private keyService: KeyService,
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
  observableFrontRepo: [ // insertion point sub template 
    Observable<BarDB[]>,
    Observable<KeyDB[]>,
    Observable<SerieDB[]>,
    Observable<ValueDB[]>,
  ] = [ // insertion point sub template 
      this.barService.getBars(),
      this.keyService.getKeys(),
      this.serieService.getSeries(),
      this.valueService.getValues(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            bars_,
            keys_,
            series_,
            values_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var bars: BarDB[]
            bars = bars_ as BarDB[]
            var keys: KeyDB[]
            keys = keys_ as KeyDB[]
            var series: SerieDB[]
            series = series_ as SerieDB[]
            var values: ValueDB[]
            values = values_ as ValueDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Bars_array = bars

            // clear the map that counts Bar in the GET
            FrontRepoSingloton.Bars_batch.clear()

            bars.forEach(
              bar => {
                FrontRepoSingloton.Bars.set(bar.ID, bar)
                FrontRepoSingloton.Bars_batch.set(bar.ID, bar)
              }
            )

            // clear bars that are absent from the batch
            FrontRepoSingloton.Bars.forEach(
              bar => {
                if (FrontRepoSingloton.Bars_batch.get(bar.ID) == undefined) {
                  FrontRepoSingloton.Bars.delete(bar.ID)
                }
              }
            )

            // sort Bars_array array
            FrontRepoSingloton.Bars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Keys_array = keys

            // clear the map that counts Key in the GET
            FrontRepoSingloton.Keys_batch.clear()

            keys.forEach(
              key => {
                FrontRepoSingloton.Keys.set(key.ID, key)
                FrontRepoSingloton.Keys_batch.set(key.ID, key)
              }
            )

            // clear keys that are absent from the batch
            FrontRepoSingloton.Keys.forEach(
              key => {
                if (FrontRepoSingloton.Keys_batch.get(key.ID) == undefined) {
                  FrontRepoSingloton.Keys.delete(key.ID)
                }
              }
            )

            // sort Keys_array array
            FrontRepoSingloton.Keys_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Series_array = series

            // clear the map that counts Serie in the GET
            FrontRepoSingloton.Series_batch.clear()

            series.forEach(
              serie => {
                FrontRepoSingloton.Series.set(serie.ID, serie)
                FrontRepoSingloton.Series_batch.set(serie.ID, serie)
              }
            )

            // clear series that are absent from the batch
            FrontRepoSingloton.Series.forEach(
              serie => {
                if (FrontRepoSingloton.Series_batch.get(serie.ID) == undefined) {
                  FrontRepoSingloton.Series.delete(serie.ID)
                }
              }
            )

            // sort Series_array array
            FrontRepoSingloton.Series_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Values_array = values

            // clear the map that counts Value in the GET
            FrontRepoSingloton.Values_batch.clear()

            values.forEach(
              value => {
                FrontRepoSingloton.Values.set(value.ID, value)
                FrontRepoSingloton.Values_batch.set(value.ID, value)
              }
            )

            // clear values that are absent from the batch
            FrontRepoSingloton.Values.forEach(
              value => {
                if (FrontRepoSingloton.Values_batch.get(value.ID) == undefined) {
                  FrontRepoSingloton.Values.delete(value.ID)
                }
              }
            )

            // sort Values_array array
            FrontRepoSingloton.Values_array.sort((t1, t2) => {
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
                  let _key = FrontRepoSingloton.Keys.get(bar.XID.Int64)
                  if (_key) {
                    bar.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = FrontRepoSingloton.Keys.get(bar.YID.Int64)
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
            series.forEach(
              serie => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Key redeeming
                {
                  let _key = FrontRepoSingloton.Keys.get(serie.KeyID.Int64)
                  if (_key) {
                    serie.Key = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Bar.Set redeeming
                {
                  let _bar = FrontRepoSingloton.Bars.get(serie.Bar_SetDBID.Int64)
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
              }
            )
            values.forEach(
              value => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Serie.Values redeeming
                {
                  let _serie = FrontRepoSingloton.Series.get(value.Serie_ValuesDBID.Int64)
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

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.barService.getBars()
        ]).subscribe(
          ([ // insertion point sub template 
            bars,
          ]) => {
            // init the array
            FrontRepoSingloton.Bars_array = bars

            // clear the map that counts Bar in the GET
            FrontRepoSingloton.Bars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bars.forEach(
              bar => {
                FrontRepoSingloton.Bars.set(bar.ID, bar)
                FrontRepoSingloton.Bars_batch.set(bar.ID, bar)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field X redeeming
                {
                  let _key = FrontRepoSingloton.Keys.get(bar.XID.Int64)
                  if (_key) {
                    bar.X = _key
                  }
                }
                // insertion point for pointer field Y redeeming
                {
                  let _key = FrontRepoSingloton.Keys.get(bar.YID.Int64)
                  if (_key) {
                    bar.Y = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear bars that are absent from the GET
            FrontRepoSingloton.Bars.forEach(
              bar => {
                if (FrontRepoSingloton.Bars_batch.get(bar.ID) == undefined) {
                  FrontRepoSingloton.Bars.delete(bar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.keyService.getKeys()
        ]).subscribe(
          ([ // insertion point sub template 
            keys,
          ]) => {
            // init the array
            FrontRepoSingloton.Keys_array = keys

            // clear the map that counts Key in the GET
            FrontRepoSingloton.Keys_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            keys.forEach(
              key => {
                FrontRepoSingloton.Keys.set(key.ID, key)
                FrontRepoSingloton.Keys_batch.set(key.ID, key)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear keys that are absent from the GET
            FrontRepoSingloton.Keys.forEach(
              key => {
                if (FrontRepoSingloton.Keys_batch.get(key.ID) == undefined) {
                  FrontRepoSingloton.Keys.delete(key.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.serieService.getSeries()
        ]).subscribe(
          ([ // insertion point sub template 
            series,
          ]) => {
            // init the array
            FrontRepoSingloton.Series_array = series

            // clear the map that counts Serie in the GET
            FrontRepoSingloton.Series_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            series.forEach(
              serie => {
                FrontRepoSingloton.Series.set(serie.ID, serie)
                FrontRepoSingloton.Series_batch.set(serie.ID, serie)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Key redeeming
                {
                  let _key = FrontRepoSingloton.Keys.get(serie.KeyID.Int64)
                  if (_key) {
                    serie.Key = _key
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Bar.Set redeeming
                {
                  let _bar = FrontRepoSingloton.Bars.get(serie.Bar_SetDBID.Int64)
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
              }
            )

            // clear series that are absent from the GET
            FrontRepoSingloton.Series.forEach(
              serie => {
                if (FrontRepoSingloton.Series_batch.get(serie.ID) == undefined) {
                  FrontRepoSingloton.Series.delete(serie.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.valueService.getValues()
        ]).subscribe(
          ([ // insertion point sub template 
            values,
          ]) => {
            // init the array
            FrontRepoSingloton.Values_array = values

            // clear the map that counts Value in the GET
            FrontRepoSingloton.Values_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            values.forEach(
              value => {
                FrontRepoSingloton.Values.set(value.ID, value)
                FrontRepoSingloton.Values_batch.set(value.ID, value)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Serie.Values redeeming
                {
                  let _serie = FrontRepoSingloton.Series.get(value.Serie_ValuesDBID.Int64)
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
            FrontRepoSingloton.Values.forEach(
              value => {
                if (FrontRepoSingloton.Values_batch.get(value.ID) == undefined) {
                  FrontRepoSingloton.Values.delete(value.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
export function getSerieUniqueID(id: number): number {
  return 41 * id
}
export function getValueUniqueID(id: number): number {
  return 43 * id
}
