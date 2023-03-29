import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { of, Observable, throwError } from 'rxjs';


interface LoginResponse {
  // Define the shape of the response object here
  // For example, you might have a "token" property
  token: string;
}

@Injectable({
  providedIn: 'root'

})
export class LoginService {
  private loginUrl = 'http://localhost:8080/api/login';


  constructor(private http: HttpClient, ) {}


  login(username: string, password: string, ): Observable<any> {
    const headers = new HttpHeaders().set('Content-Type', 'application/json');
    const body = { username:username, password:password};
    return this.http.post(this.loginUrl, body, { headers }).pipe(

      catchError(error => {
        // Handle error
        return throwError(error);
      })
    );
  }
}
