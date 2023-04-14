import { Component, OnChanges, OnInit } from '@angular/core';
import { Injectable } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from './profile.service';
import { User } from '../user.model';
import { LoginService } from '../login/login.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})

export class ProfileComponent implements OnInit {
  name: string  = "Please Login";
  email: string  = "";


  constructor(private route: ActivatedRoute, private loginService: LoginService, public router: Router) {
    if (!this.loginService.islog) {
      this.router.navigate(['/login']);
    }

  }

  ngOnInit() {
    const userString = this.route.snapshot.queryParamMap.get('user') as string;
    const user = JSON.parse(userString) as User;
    this.name = user.name;
    this.email = user.email;
   // this.name = "not workin";
  }
}


