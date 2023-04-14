import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { of, Observable, throwError } from 'rxjs';
import { User } from '../user.model';
import { map } from 'rxjs/operators';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})

export class LoginService {
  private readonly AUTH_KEY = 'auth';
  constructor(private http: HttpClient, public router:Router) {}
  islog = false;

 public login(email: string, password: string): Observable<User> {
    const url = 'http://localhost:8080/api/login'
    const body = { email, password };


    return this.http.post(url, body).pipe(
      map((response: any) => {
        this.islog = true;
        const user: User = {
          id: response.id,
          name: response.name,
          email: response.email,
          password: response.password
        };

        localStorage.setItem(this.AUTH_KEY, JSON.stringify(user));
        return user;

      })
    );
  }

  public logout() {
    localStorage.removeItem(this.AUTH_KEY);
    this.islog = false;
    this.router.navigate(['']);

  }

  public isLoggedIn(): boolean {
    const userJson = localStorage.getItem(this.AUTH_KEY);
    return !!userJson;
  }

  public getUser(): User {
    const userJson = localStorage.getItem(this.AUTH_KEY);
    if (userJson != null) {
      return JSON.parse(userJson);
    }
    const user : User = {
      id: '0',
      name: "please login",
      email : "",
      password: "",

    }
    return user;
  }

}

