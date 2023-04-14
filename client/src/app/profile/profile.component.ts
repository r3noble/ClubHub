import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from './profile.service';
import { User } from '../user.model';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})

export class ProfileComponent implements OnInit {
  name: string | null = "please login";
  email: string | null = "";

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
    this.name = this.route.snapshot.queryParamMap.get('name');
    this.email = this.route.snapshot.queryParamMap.get('email');
  }
}


