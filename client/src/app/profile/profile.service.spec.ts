import { TestBed } from '@angular/core/testing';

import { ProfileService } from './profile.service';

// Other imports
import { HttpClient, HttpClientModule, HttpErrorResponse } from '@angular/common/http';
import { AppRoutingModule } from '../app-routing.module';
import { RouterModule } from '@angular/router';
import appRoutes from '../app-routing.module';
import { FormsModule } from '@angular/forms';


describe('ProfileService', () => {
  let service: ProfileService;

  beforeEach(() => {
    TestBed.configureTestingModule({imports:[HttpClientModule,AppRoutingModule,RouterModule, FormsModule]});
    service = TestBed.inject(ProfileService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
