import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { User } from '../user.model';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient, private authService: AuthService) {}

  login(email: string, password: string): Observable<User> {
    const url = 'http://localhost:8080/api/login'
    const body = { email, password };

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

  logout(): void {
    const user: User = {
      id: "0",
      name: "Please Login",
      email: "",
      password: "",
      clubs: ""
    }
    this.authService.setLoggedIn(false,user);
  }
}
