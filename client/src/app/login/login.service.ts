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
    const body = { username, password };
    const headers = new HttpHeaders().set("Access-Control-Allow-Origin", "*");
    const options = { headers: headers };
    return this.http.post('localhost:8080/api/login', body, options);

  }
}
