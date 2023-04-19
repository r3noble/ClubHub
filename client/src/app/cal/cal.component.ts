import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-cal',
  templateUrl: './cal.component.html',
  styleUrls: ['./cal.component.css']
})
export class CalComponent {
  events: Event[] = [];


  constructor(private http: HttpClient) {}

  //ngOnInit() {
  //  this.http.get<Event[]>('/api/getEvents').subscribe((events) => {
  //    this.events = events;
  //  });
 // }


}
