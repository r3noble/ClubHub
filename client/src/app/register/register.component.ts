import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RegisterService } from './register.service';


@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css']
})
export class RegisterComponent {

    constructor(public router: Router, private RegisterService: RegisterService) {}
    
    fullName: string = "";
    firstName: string = "";
    lastName: string = "";
    email: string = "";
    password: string = "";  
    rpassword: string = ""; 
    college: string = "";
    register: boolean = false;

    onRegister() {
        // Validate the form fields
        if(this.firstName == ""){
            alert("Please enter a first name.")
        }

        if(this.lastName == ""){
            alert("Please enter a last name.")
        }
        
        if (!this.email.endsWith("@ufl.edu")) {
            alert("Please enter a valid email ending in '@ufl.edu'.");
            return;
        }

        if(this.password == ""){
            alert("Please enter a password.")
        }

        if (this.password != this.rpassword) {
            alert("Passwords do not match. Please retype your password.");
            return;
        }

        if(this.college == ""){
            alert("Please select your college.")
        }

        const checkbox = document.getElementById('agree-checkbox') as HTMLInputElement;
        if (!checkbox.checked) {
            alert("Please agree to the terms and conditions before registering.");
            return;
        }
        else{
            // If all fields are valid, set the boolean variable to true
            this.register = true;
        }

        // If register boolean is tru then submit the form data to the server to register the user.
        if(this.register == true){
            this.fullName = this.firstName + " " + this.lastName;
            this.RegisterService.registerUser("1", this.fullName, this.email, this.password).subscribe(
                response => {
                  console.log(response);
                  // handle success
                  this.router.navigate(['']);
                },
                error => {
                  console.log(error);
                  // handle error
                }
              );            
        }

    }

    onCancel() {
        this.router.navigate(['']);
      }
}   

