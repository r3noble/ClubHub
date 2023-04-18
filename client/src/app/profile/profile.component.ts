import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from '../login/auth.service';
import { User } from '../user.model';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  name: string  = "Please Login";
  email: string  = "";
  clubs: string = "No clubs joined yet!"

  constructor(private route: ActivatedRoute, private authService: AuthService) {}

  ngOnInit() {
    const user = this.authService.getUser();
    //  const user = JSON.parse(userString) as User;
      this.name = user.name;
      this.email = user.email;
      this.clubs = user.clubs;
      if ( this.clubs == "" || this.clubs == "No clubs joined!" ) {
        this.clubs = "No clubs joined!"
      }
      //alert(this.clubs);

     // this.name = "not workin";
  }

}
