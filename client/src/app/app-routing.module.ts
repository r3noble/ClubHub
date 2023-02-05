import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import {Routes, RouterModule} from '@angular/router';
import { ProfileComponent } from './profile/profile.component';


const routes: Routes = [
  {
  path : '',
  component: HomeComponent

  },
  {
    path: 'profile',
    component: ProfileComponent
  }
];


@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule
  ],
  exports:[RouterModule]
})
export class AppRoutingModule { }
