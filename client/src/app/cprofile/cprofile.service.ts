import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Club } from '../club.model';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CprofileService {
  constructor(private http: HttpClient) {}

  getClubInfo(name: string): Observable<Club> {
    const url = `http://localhost:8080/api/getClub/${name}`;

    return this.http.get<Club>(url).pipe(
      map((response: any) => {
        const club: Club = {
          name: response.name,
          president: response.president,
          VP: response.VP,
          treasurer: response.treasurer,
          about: response.about,
        };

        return club;
      })
    );
  }
}
