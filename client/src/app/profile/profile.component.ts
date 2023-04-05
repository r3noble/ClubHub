import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from './profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})

export class ProfileComponent implements OnInit {
  username: string ="";
  profile: any;
  sub: any;

  constructor(
    private route: ActivatedRoute,
    private profileService: ProfileService
  ) {}

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.username = params['Name'];
      });
      console.log(this.username);
    this.profileService.getProfile(this.username).subscribe(data => {
      this.profile = data;
    });
  }
}


