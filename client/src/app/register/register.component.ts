import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { FormBuilder, FormGroup, Validators } from '@angular/forms';


@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css']
})
export class RegisterComponent {

    constructor(private router: Router) { }
    
    name: string = "";
    username: string = "";
    email: string = "";
    password: string = "";  
    rpassword: string = ""; 
    major: string = "";

    onCancel() {
        this.router.navigate(['']);
      }
}   

