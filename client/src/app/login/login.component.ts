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
        (user: User) => {

          this.router.navigate(['/profile'], { queryParams: { user: JSON.stringify(user) }});
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
