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


  constructor(private LoginService: LoginService,  public router: Router) {}

  onSubmit() {
    this.LoginService.login(this.username, this.password)
    .subscribe(
      ()=> {
        // Redirect to the profile page if successful login
        this.router.navigate(['/profile', { user: "hello"}])
        alert();
       },

    );
  }

  onRegister() {
    this.router.navigate(['/register' ]);
  }
}
