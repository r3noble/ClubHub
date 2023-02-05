import { Component } from '@angular/core';
import {HttpClient} from '@angular/common/http'


interface clubItem {
  clubTitle: string
  title: string
  post:string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent {
  public clubTitle = ''
  public title = ''
  public post = ''


  public clubItems: clubItem[] = [
    {
      clubTitle: 'WECE',
      title: 'GBM',
      post: 'description'
    },
    {
      clubTitle: 'ACM',
      title: 'Resume Review',
      post: 'description'
    }
  ]

  //working with this later is how to integrate backend
  // 30:30 in https://www.youtube.com/watch?v=pHRHJCYBqxw
constructor(
  private httpclient: HttpClient
) {}

// this should be customized on another page, accessed by club owners
addEvent() {
  this.clubItems.push({
    //backend will automatically grab user's club
    clubTitle: 'User Club',
    title: this.title,
    post: this.post
  })
  this.title = ''
  this.post = ''
}
}
