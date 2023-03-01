import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { of, Observable, throwError } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class LoginService {
  private loginUrl = 'http://localhost:8080/user/login';

  constructor(private http: HttpClient) {}

  login(username: string, password: string): Observable<any> { // Use the new Observable symbol
    const credentials = { username: username, password: password };
    return this.http.post(this.loginUrl, credentials).pipe(
      catchError((error) => {
        return throwError(error);
      })
    );
  }
}
