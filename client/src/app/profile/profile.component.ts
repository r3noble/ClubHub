import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from '../login/auth.service';
import { User } from '../user.model';
import { CprofileService } from '../cprofile/cprofile.service';
import { Club } from '../club.model';
@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  name: string  = "Please Login";
  email: string  = "";
  clubs: string = "No clubs joined yet!"
  club:Club | null = null;


  constructor(private route: ActivatedRoute, private authService: AuthService, private cprofile: CprofileService) {}

  ngOnInit() {
    const user = this.authService.getUser();
    //  const user = JSON.parse(userString) as User;
      this.name = user?.name ?? "No user found";
      this.email = user?.email;
      this.clubs = user?.clubs;
      if ( this.clubs == "" || this.clubs == "No clubs joined!" ) {
        this.clubs = "No clubs joined!"
      }

      this.cprofile.getClubInfo(user?.clubs as string).subscribe(
        (club: Club) => {
          this.club = club;
          //console.log(this.club.VP);
        },
        (error) => {
          console.log(error);
        }
      );
      //alert(this.clubs);

     // this.name = "not workin";
  }

}
