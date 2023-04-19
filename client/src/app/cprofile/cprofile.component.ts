import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CprofileService } from './cprofile.service';
import { User } from '../user.model';
import { Club } from '../club.model';
import { AuthService } from '../login/auth.service';
import { member } from '../member.model';

@Component({
  selector: 'app-cprofile',
  templateUrl: './cprofile.component.html',
  styleUrls: ['./cprofile.component.css']
})
export class CprofileComponent implements OnInit {
  name: string ="";
  club: Club | null = null;
  url:string="";
  member:member |null = null;

  constructor(private authService: AuthService, private route: ActivatedRoute, private router: Router, private cprofileService: CprofileService) { }

  isLoggedIn(): boolean {
    return this.authService.isLoggedIn();
  }

  onJoin(){
    this.cprofileService.joinClub(this.authService.getUser().id, this.name)
    .subscribe(
      (user: User) => {
        this.router.navigate(['/profile', {User: user}]);
      },
      (error) => {
        alert('You are already in this club!');
        console.log(error);
      }
    );

  }

  onCancel() {
    this.router.navigate(['/club']);
  }

  ngOnInit(): void {

    this.cprofileService.ismember(this.authService.getUser().id,this.name).subscribe(
      (member:member)=> {

      },
      (error) => {
        alert('You are already in this club!');
        console.log(error);
      }
    );
    this.route.params.subscribe(params => {
      this.name = params['name'];
      this.cprofileService.getClubInfo(this.name).subscribe(
        (club: Club) => {
          this.club = club;
          //this.url = club.calendar;
         // console.log(this.club.VP);
        },
        (error) => {
          console.log(error);
        }
      );
    });
  }
}
