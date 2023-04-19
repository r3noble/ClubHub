import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { PublicprofileService } from './publicprofile.service';
import { User } from '../user.model';
import { CprofileService } from '../cprofile/cprofile.service';
import { Club } from '../club.model';
@Component({
  selector: 'app-publicprofile',
  templateUrl: './publicprofile.component.html',
  styleUrls: ['./publicprofile.component.css']
})
export class PublicprofileComponent implements OnInit {
  name: string ="";
  clubs:string="";
  email:string="";
  user: User | null = null;
  club:Club |null = null;

  constructor(private route: ActivatedRoute, private publicprofileservice: PublicprofileService, private cprofile:CprofileService) { }

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      this.name = params['name'];
      this.publicprofileservice.getpublicprofile(this.name).subscribe(
        (user: User) => {
          this.user = user;
          console.log(this.user.clubs);

          this.cprofile.getClubInfo(this.user?.clubs as string).subscribe(
            (club: Club) => {
              this.club = club;
              //console.log(this.club.VP);
            },
            (error) => {
              console.log(error);
            }
          );

        },
        (error) => {
          console.log(error);
        }

      );



    });
  }
}
