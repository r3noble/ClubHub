import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms' ;
import { HttpClientModule, HttpClient} from '@angular/common/http' ;
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { ProfileComponent } from './profile/profile.component';
import  appRoutes  from './app-routing.module';
import { LoginComponent } from './login/login.component';
import { CalendarComponent } from './calendar/calendar.component';
import { RegisterComponent } from './register/register.component';
import { SchedulerModule } from 'smart-webcomponents-angular/scheduler';
import { StepperComponent } from './stepper/stepper.component';
import { ClubComponent } from './club/club.component';
import { CprofileComponent } from './cprofile/cprofile.component';
import { CalComponent } from './cal/cal.component';



@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    HeaderComponent,
    FooterComponent,
    ProfileComponent,
    LoginComponent,
    CalendarComponent,
    RegisterComponent,
    StepperComponent,
    ClubComponent,
    CprofileComponent,
    CalComponent
  ],


  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes),
    FormsModule,
    HttpClientModule,
    AppRoutingModule,
    SchedulerModule,

  ],
  exports: [
    HeaderComponent,
    FooterComponent
  ],
    schemas: [
      NO_ERRORS_SCHEMA
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
