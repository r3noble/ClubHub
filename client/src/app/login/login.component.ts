import { Component } from '@angular/core';
import { LoginService } from './login.service';
import { Router } from '@angular/router';
import { User } from '../user.model';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  email: string = "";
  password: string = "";
  errorMessage: string = "";
  baseUrl: any;


  constructor(private LoginService: LoginService,  public router: Router) {}

  onSubmit() {
    this.LoginService.login(this.email, this.password)
      .subscribe(
        (user: User) => {

          this.router.navigate(['/profile', {User: user}]);
        },
        (error) => {
          alert('Incorrect Username or Password');
          console.log(error);
        }
      );
  }

  onRegister() {
    this.router.navigate(['/register' ]);
  }
}
