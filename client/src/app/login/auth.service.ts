import { Injectable } from "@angular/core";
import { User } from "../user.model";

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly AUTH_KEY = 'auth';
  private isLoggedInValue = false;

  constructor() {
    const authValue = localStorage.getItem(this.AUTH_KEY);
    if (authValue) {
      const authObject = JSON.parse(authValue);
      this.isLoggedInValue = authObject.isLoggedIn;
    }
  }

  isLoggedIn(): boolean {
    return this.isLoggedInValue;
  }

  setLoggedIn(value: boolean, user: User): void {
    this.isLoggedInValue = value;
    localStorage.setItem(
      this.AUTH_KEY,
      JSON.stringify({
        isLoggedIn: value,
        user,
      })
    );
  }

  public getUser(): User {
    const authValue = localStorage.getItem(this.AUTH_KEY) as string;
    const authObject = JSON.parse(authValue);
    return authObject.user;
  }
}
