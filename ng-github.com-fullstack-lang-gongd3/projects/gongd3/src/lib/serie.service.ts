// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { SerieAPI } from './serie-api'
import { Serie, CopySerieToSerieAPI } from './serie'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { KeyAPI } from './key-api'
import { ValueAPI } from './value-api'

@Injectable({
  providedIn: 'root'
})
export class SerieService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  SerieServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private seriesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.seriesUrl = origin + '/api/github.com/fullstack-lang/gongd3/go/v1/series';
  }

  /** GET series from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI[]> {
    return this.getSeries(GONG__StackPath, frontRepo)
  }
  getSeries(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<SerieAPI[]>(this.seriesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<SerieAPI[]>('getSeries', []))
      );
  }

  /** GET serie by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI> {
    return this.getSerie(id, GONG__StackPath, frontRepo)
  }
  getSerie(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.seriesUrl}/${id}`;
    return this.http.get<SerieAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched serie id=${id}`)),
      catchError(this.handleError<SerieAPI>(`getSerie id=${id}`))
    );
  }

  // postFront copy serie to a version with encoded pointers and post to the back
  postFront(serie: Serie, GONG__StackPath: string): Observable<SerieAPI> {
    let serieAPI = new SerieAPI
    CopySerieToSerieAPI(serie, serieAPI)
    const id = typeof serieAPI === 'number' ? serieAPI : serieAPI.ID
    const url = `${this.seriesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SerieAPI>(url, serieAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SerieAPI>('postSerie'))
    );
  }
  
  /** POST: add a new serie to the server */
  post(seriedb: SerieAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI> {
    return this.postSerie(seriedb, GONG__StackPath, frontRepo)
  }
  postSerie(seriedb: SerieAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SerieAPI>(this.seriesUrl, seriedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted seriedb id=${seriedb.ID}`)
      }),
      catchError(this.handleError<SerieAPI>('postSerie'))
    );
  }

  /** DELETE: delete the seriedb from the server */
  delete(seriedb: SerieAPI | number, GONG__StackPath: string): Observable<SerieAPI> {
    return this.deleteSerie(seriedb, GONG__StackPath)
  }
  deleteSerie(seriedb: SerieAPI | number, GONG__StackPath: string): Observable<SerieAPI> {
    const id = typeof seriedb === 'number' ? seriedb : seriedb.ID;
    const url = `${this.seriesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<SerieAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted seriedb id=${id}`)),
      catchError(this.handleError<SerieAPI>('deleteSerie'))
    );
  }

  // updateFront copy serie to a version with encoded pointers and update to the back
  updateFront(serie: Serie, GONG__StackPath: string): Observable<SerieAPI> {
    let serieAPI = new SerieAPI
    CopySerieToSerieAPI(serie, serieAPI)
    const id = typeof serieAPI === 'number' ? serieAPI : serieAPI.ID
    const url = `${this.seriesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<SerieAPI>(url, serieAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SerieAPI>('updateSerie'))
    );
  }

  /** PUT: update the seriedb on the server */
  update(seriedb: SerieAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI> {
    return this.updateSerie(seriedb, GONG__StackPath, frontRepo)
  }
  updateSerie(seriedb: SerieAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SerieAPI> {
    const id = typeof seriedb === 'number' ? seriedb : seriedb.ID;
    const url = `${this.seriesUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<SerieAPI>(url, seriedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated seriedb id=${seriedb.ID}`)
      }),
      catchError(this.handleError<SerieAPI>('updateSerie'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in SerieService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("SerieService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}