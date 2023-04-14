import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { of, Observable, throwError } from 'rxjs';
import { User } from '../user.model';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  constructor(private http: HttpClient) {}
  islog = false;

  login(username: string, password: string): Observable<User> {
    const url = 'http://localhost:8080/api/login'
    const body = { username, password };


    return this.http.post(url, body).pipe(
      map((response: any) => {
        this.islog = true;
        const user: User = {
          id: response.id,
          name: response.name,
          email: response.email,
          password: response.password
        };
        return user;
      })
    );
  }

}

