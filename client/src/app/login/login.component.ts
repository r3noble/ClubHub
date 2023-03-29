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
      (response)=> {
        // Redirect to the profile page if successful login
        this.router.navigate(['/profile', { user: "hello"}])
        alert(response);
       },
      (error) => {
        this.errorMessage = error;
        //this.router.navigate(['/profile', { username: this.username }])
        this.router.navigate(['/profile', { username: this.username}]);
      }
    );
  }

  onRegister() {
    this.router.navigate(['/register' ]);
  }
}
