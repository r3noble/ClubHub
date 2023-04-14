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
  username: string = "";
  password: string = "";
  errorMessage: string = "";
  baseUrl: any;


  constructor(private LoginService: LoginService,  public router: Router) {}

  onSubmit() {
    this.LoginService.login(this.username, this.password)
    .subscribe(
      response => {
        this.router.navigate(['/profile', { User: response}])

      },
      error => {
        alert(error);
        console.log(error);
        // handle error
      }
    );
  }

  onRegister() {
    this.router.navigate(['/register' ]);
  }
}
