import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class TopographyService {
    constructor(private http: HttpClient) { }

    getTopographyData(): Observable<any> {
        const topoDataURL =
            'https://cdn.jsdelivr.net/npm/us-atlas@3/counties-albers-10m.json';

        return this.http.get(topoDataURL);
    }
}
