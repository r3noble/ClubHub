import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CprofileService } from './cprofile.service';
import { Club } from '../club.model';

@Component({
  selector: 'app-cprofile',
  templateUrl: './cprofile.component.html',
  styleUrls: ['./cprofile.component.css']
})
export class CprofileComponent implements OnInit {
  name: string ="";
  club: Club | null = null;

  constructor(private route: ActivatedRoute, private router: Router, private cprofileService: CprofileService) { }

  onJoin(){
    
  }

  onCancel() {
    this.router.navigate(['/club']);
  }

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      this.name = params['name'];
      this.cprofileService.getClubInfo(this.name).subscribe(
        (club: Club) => {
          this.club = club;
          console.log(this.club.VP);
        },
        (error) => {
          console.log(error);
        }
      );
    });
  }
}
