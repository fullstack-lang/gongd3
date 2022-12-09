import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { CountryDB } from './country-db'
import { CountryService } from './country.service'

import { HelloDB } from './hello-db'
import { HelloService } from './hello.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Countrys_array = new Array<CountryDB>(); // array of repo instances
  Countrys = new Map<number, CountryDB>(); // map of repo instances
  Countrys_batch = new Map<number, CountryDB>(); // same but only in last GET (for finding repo instances to delete)
  Hellos_array = new Array<HelloDB>(); // array of repo instances
  Hellos = new Map<number, HelloDB>(); // map of repo instances
  Hellos_batch = new Map<number, HelloDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private countryService: CountryService,
    private helloService: HelloService,
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
    Observable<CountryDB[]>,
    Observable<HelloDB[]>,
  ] = [ // insertion point sub template 
      this.countryService.getCountrys(),
      this.helloService.getHellos(),
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
            countrys_,
            hellos_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var countrys: CountryDB[]
            countrys = countrys_ as CountryDB[]
            var hellos: HelloDB[]
            hellos = hellos_ as HelloDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Countrys_array = countrys

            // clear the map that counts Country in the GET
            FrontRepoSingloton.Countrys_batch.clear()

            countrys.forEach(
              country => {
                FrontRepoSingloton.Countrys.set(country.ID, country)
                FrontRepoSingloton.Countrys_batch.set(country.ID, country)
              }
            )

            // clear countrys that are absent from the batch
            FrontRepoSingloton.Countrys.forEach(
              country => {
                if (FrontRepoSingloton.Countrys_batch.get(country.ID) == undefined) {
                  FrontRepoSingloton.Countrys.delete(country.ID)
                }
              }
            )

            // sort Countrys_array array
            FrontRepoSingloton.Countrys_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Hellos_array = hellos

            // clear the map that counts Hello in the GET
            FrontRepoSingloton.Hellos_batch.clear()

            hellos.forEach(
              hello => {
                FrontRepoSingloton.Hellos.set(hello.ID, hello)
                FrontRepoSingloton.Hellos_batch.set(hello.ID, hello)
              }
            )

            // clear hellos that are absent from the batch
            FrontRepoSingloton.Hellos.forEach(
              hello => {
                if (FrontRepoSingloton.Hellos_batch.get(hello.ID) == undefined) {
                  FrontRepoSingloton.Hellos.delete(hello.ID)
                }
              }
            )

            // sort Hellos_array array
            FrontRepoSingloton.Hellos_array.sort((t1, t2) => {
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
            countrys.forEach(
              country => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Hello redeeming
                {
                  let _hello = FrontRepoSingloton.Hellos.get(country.HelloID.Int64)
                  if (_hello) {
                    country.Hello = _hello
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            hellos.forEach(
              hello => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
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

  // CountryPull performs a GET on Country of the stack and redeem association pointers 
  CountryPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.countryService.getCountrys()
        ]).subscribe(
          ([ // insertion point sub template 
            countrys,
          ]) => {
            // init the array
            FrontRepoSingloton.Countrys_array = countrys

            // clear the map that counts Country in the GET
            FrontRepoSingloton.Countrys_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            countrys.forEach(
              country => {
                FrontRepoSingloton.Countrys.set(country.ID, country)
                FrontRepoSingloton.Countrys_batch.set(country.ID, country)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Hello redeeming
                {
                  let _hello = FrontRepoSingloton.Hellos.get(country.HelloID.Int64)
                  if (_hello) {
                    country.Hello = _hello
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear countrys that are absent from the GET
            FrontRepoSingloton.Countrys.forEach(
              country => {
                if (FrontRepoSingloton.Countrys_batch.get(country.ID) == undefined) {
                  FrontRepoSingloton.Countrys.delete(country.ID)
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

  // HelloPull performs a GET on Hello of the stack and redeem association pointers 
  HelloPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.helloService.getHellos()
        ]).subscribe(
          ([ // insertion point sub template 
            hellos,
          ]) => {
            // init the array
            FrontRepoSingloton.Hellos_array = hellos

            // clear the map that counts Hello in the GET
            FrontRepoSingloton.Hellos_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            hellos.forEach(
              hello => {
                FrontRepoSingloton.Hellos.set(hello.ID, hello)
                FrontRepoSingloton.Hellos_batch.set(hello.ID, hello)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear hellos that are absent from the GET
            FrontRepoSingloton.Hellos.forEach(
              hello => {
                if (FrontRepoSingloton.Hellos_batch.get(hello.ID) == undefined) {
                  FrontRepoSingloton.Hellos.delete(hello.ID)
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
export function getCountryUniqueID(id: number): number {
  return 31 * id
}
export function getHelloUniqueID(id: number): number {
  return 37 * id
}
