import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import { AppComponent } from './app.component';
import {Routes } from '@angular/router';
import { ProfileComponent } from './profile/profile.component';


const appRoutes: Routes = [
  {
  path : '',
  component: AppComponent

  },
  {
    path: 'profile',
    component: ProfileComponent
  }
];

export default appRoutes;

@NgModule({
  declarations: [],
  imports: [
    CommonModule,

  ],
  exports:[]
})
export class AppRoutingModule { }
