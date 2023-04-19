import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Club } from '../club.model';
import { map } from 'rxjs/operators';
import { User } from '../user.model';
import { Observable } from 'rxjs';
import { AuthService } from '../login/auth.service';
import { member } from '../member.model';

@Injectable({
  providedIn: 'root'
})
export class CprofileService {
  constructor(private http: HttpClient, private authService: AuthService) {}

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
          events: response.events,
        };



        return club;
      })
    );
  }
  ismember(id:string, name:string):Observable<member>{
    const url = `http://localhost:8080/api/getRole`;
    const body = { id, name };
    return this.http.post<Club>(url,body).pipe(
      map((response: any) => {
        const member: member = {
          id: response.string,
          name: response.name,
        };
        return member;
      })
    );
  }

  joinClub(id: string, name: string): Observable<User>
  {
    const url = `http://localhost:8080/api/joinClub`;
    const body = {
      id: id,
      name: name,
    };

    return this.http.post(url, body).pipe(
      map((response: any) => {
        const user: User = {
          id: response.id,
          name: response.name,
          email: response.email,
          password: response.password,
          clubs: response.clubs
        };
        this.authService.setLoggedIn(true,user);

        return user;
      })
    );
  }

}
