import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from '../user.model';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PublicprofileService {
  constructor(private http: HttpClient) {}

  getpublicprofile(name: string): Observable<User> {
    const url = `http://localhost:8080/api/getUserfromClub/${name}`;

    return this.http.get<User>(url).pipe(
      map((response: any) => {
        const user: User = {
          name: response.name,
          id: response.id,
          password: "",
          clubs : response.clubs,
          email:response.email
        };

        return user;
      })
    );
  }
}
