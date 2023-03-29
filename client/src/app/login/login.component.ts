import { Component } from '@angular/core';
import { LoginService } from './login.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username: string = "";
  password: string = "";
  errorMessage: string = "";
  baseUrl: any;

  constructor(private loginService: LoginService,  public router: Router) {}

  onSubmit() {
    this.loginService.login(this.username, this.password)
    .subscribe(
      () => {
        // Redirect to the profile page if successful login
        this.router.navigate(['/profile', { username: this.username }]);
      },
      (error) => {
        this.errorMessage = this.username;
        //this.router.navigate(['/profile', { username: this.username }])
        //this.router.navigate(['/profile']);
      }
    );
  }

  onRegister() {
    this.router.navigate(['/register' ]);
  }
}
