import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { of, Observable, throwError } from 'rxjs';



@Injectable({
  providedIn: 'root'
})
export class LoginService {
  constructor(private http: HttpClient) {}

  login(username: string, password: string) {
    const url = 'http://localhost:8080/api/login'
    const body = { username, password };

    return this.http.post(url, body);

  }
}
