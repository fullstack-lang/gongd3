import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import * as topojson_specification from 'topojson-specification'

@Injectable({
    providedIn: 'root',
})
export class TopographyService {
    constructor(private http: HttpClient) { }

    getTopographyData(source: string): Observable<any> {

        console.log("Topography service")

        return this.http.get(source);
    }
}