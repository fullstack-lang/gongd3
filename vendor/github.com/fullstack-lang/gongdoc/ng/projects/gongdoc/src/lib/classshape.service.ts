// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { ClassshapeDB } from './classshape-db';

// insertion point for imports
import { PositionDB } from './position-db'
import { ReferenceDB } from './reference-db'
import { ClassdiagramDB } from './classdiagram-db'

@Injectable({
  providedIn: 'root'
})
export class ClassshapeService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ClassshapeServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private classshapesUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.classshapesUrl = origin + '/api/github.com/fullstack-lang/gongdoc/go/v1/classshapes';
  }

  /** GET classshapes from the server */
  getClassshapes(): Observable<ClassshapeDB[]> {
    return this.http.get<ClassshapeDB[]>(this.classshapesUrl)
      .pipe(
        tap(_ => this.log('fetched classshapes')),
        catchError(this.handleError<ClassshapeDB[]>('getClassshapes', []))
      );
  }

  /** GET classshape by id. Will 404 if id not found */
  getClassshape(id: number): Observable<ClassshapeDB> {
    const url = `${this.classshapesUrl}/${id}`;
    return this.http.get<ClassshapeDB>(url).pipe(
      tap(_ => this.log(`fetched classshape id=${id}`)),
      catchError(this.handleError<ClassshapeDB>(`getClassshape id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new classshape to the server */
  postClassshape(classshapedb: ClassshapeDB): Observable<ClassshapeDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    classshapedb.Position = new PositionDB
    classshapedb.Reference = new ReferenceDB
    classshapedb.Fields = []
    classshapedb.Links = []
    let _Classdiagram_Classshapes_reverse = classshapedb.Classdiagram_Classshapes_reverse
    classshapedb.Classdiagram_Classshapes_reverse = new ClassdiagramDB

    return this.http.post<ClassshapeDB>(this.classshapesUrl, classshapedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        classshapedb.Classdiagram_Classshapes_reverse = _Classdiagram_Classshapes_reverse
        this.log(`posted classshapedb id=${classshapedb.ID}`)
      }),
      catchError(this.handleError<ClassshapeDB>('postClassshape'))
    );
  }

  /** DELETE: delete the classshapedb from the server */
  deleteClassshape(classshapedb: ClassshapeDB | number): Observable<ClassshapeDB> {
    const id = typeof classshapedb === 'number' ? classshapedb : classshapedb.ID;
    const url = `${this.classshapesUrl}/${id}`;

    return this.http.delete<ClassshapeDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted classshapedb id=${id}`)),
      catchError(this.handleError<ClassshapeDB>('deleteClassshape'))
    );
  }

  /** PUT: update the classshapedb on the server */
  updateClassshape(classshapedb: ClassshapeDB): Observable<ClassshapeDB> {
    const id = typeof classshapedb === 'number' ? classshapedb : classshapedb.ID;
    const url = `${this.classshapesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    classshapedb.Position = new PositionDB
    classshapedb.Reference = new ReferenceDB
    classshapedb.Fields = []
    classshapedb.Links = []
    let _Classdiagram_Classshapes_reverse = classshapedb.Classdiagram_Classshapes_reverse
    classshapedb.Classdiagram_Classshapes_reverse = new ClassdiagramDB

    return this.http.put<ClassshapeDB>(url, classshapedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        classshapedb.Classdiagram_Classshapes_reverse = _Classdiagram_Classshapes_reverse
        this.log(`updated classshapedb id=${classshapedb.ID}`)
      }),
      catchError(this.handleError<ClassshapeDB>('updateClassshape'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
