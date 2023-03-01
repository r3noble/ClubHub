import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { FormBuilder, FormGroup, Validators } from '@angular/forms';


@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css']
})
export class RegisterComponent {

    name: string = "";
    username: string = "";
    email: string = "";
    password: string = "";  
    rpassword: string = ""; 
    major: string = "";



    registerForm: FormGroup;

    constructor(private formBuilder: FormBuilder) {
        this.registerForm = this.formBuilder.group({
        username: ['', [Validators.required]]
    });
  }

  onSpaceKeydown(event: any) {
    if (event.code === 'Space') {
      event.preventDefault();
      const input = event.target as HTMLInputElement;
      const value = input.value;
      const newValue = value.replace(/ /g, '_');
      input.value = newValue;
    }
  }

    onSubmit() {
        // Handle form submission
      }
    
      get f() {
        return this.registerForm.controls;
      }
}   

