import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import { AppComponent } from './app.component';
import {Routes } from '@angular/router';
import { ProfileComponent } from './profile/profile.component';
import { LoginComponent } from './login/login.component';
import { CalendarComponent } from './calendar/calendar.component';
import { RegisterComponent } from './register/register.component';


const appRoutes: Routes = [
  {
  path : '',
  component: HomeComponent

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
    path: 'calendar',
    component: CalendarComponent
  },
  {
    path: 'register',
    component: RegisterComponent
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
