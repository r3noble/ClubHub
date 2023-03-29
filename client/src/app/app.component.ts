import { Component } from '@angular/core';
import {HttpClientModule, HttpClient} from '@angular/common/http'




@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']

})


export class AppComponent {

  public title = 'ClubHub'

  //working with this later is how to integrate backend
  // 30:30 in https://www.youtube.com/watch?v=pHRHJCYBqxw
constructor(
  private http: HttpClient

) {}



}
