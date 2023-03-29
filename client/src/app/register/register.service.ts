import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private http: HttpClient) { }

  registerUser(id: string, fullName: string, email: string, password: string) {
    const url = 'http://localhost:8080/api/addUser';
    const data = {
      id: "1",
      name: fullName,
      email: email,
      password: password
    };
    return this.http.post(url, data);
  }
}
