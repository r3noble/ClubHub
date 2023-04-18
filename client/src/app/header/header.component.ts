import { Component } from '@angular/core';
import { LoginService } from '../login/login.service';
import { AuthService } from '../login/auth.service';
import { User } from '../user.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']

})
export class HeaderComponent {
  constructor (private authService: AuthService,private router : Router){}
  isLoggedIn(): boolean {
    return this.authService.isLoggedIn();
  }
  logout(): void {
    const user: User = {
      id: "0",
      name: "Please Login",
      email: "",
      password: "",
      clubs : ""
    }
     this.authService.setLoggedIn(false, user);
     this.router.navigate(['/login' ])
  }
}
