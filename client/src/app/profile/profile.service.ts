import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  private baseUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getProfile(name: string): Observable<any> {
    const url = `${this.baseUrl}/user/get/${name}`;
    return this.http.get<any>(url);
  }
}
