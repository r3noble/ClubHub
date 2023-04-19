import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import { AppComponent } from './app.component';
import {Routes } from '@angular/router';
import { ProfileComponent } from './profile/profile.component';
import { LoginComponent } from './login/login.component';
import { CalendarComponent } from './calendar/calendar.component';
import { RegisterComponent } from './register/register.component';
import { ClubComponent } from './club/club.component';
import { CprofileComponent } from './cprofile/cprofile.component';
import { CalComponent } from './cal/cal.component';
import { PublicprofileComponent } from './publicprofile/publicprofile.component';

const appRoutes: Routes = [
  {
  path : '',
  component: HomeComponent

  },
  {
    path: 'club',
    component: ClubComponent
  },
  {
    path: 'calendar',
    component: CalComponent
  },
  {
    path: 'profile',
    component: ProfileComponent
  },
  {
    path: 'profile/:User',
    component: ProfileComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'register',
    component: RegisterComponent
  },
  {
    path: 'cprofile',
    component: CprofileComponent
  },
  {
    path: 'cprofile/:name',
    component: CprofileComponent
  },
  {
    path: 'publicprofile',
    component: PublicprofileComponent
  },
  {
    path: 'publicprofile/:name',
    component: PublicprofileComponent
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
