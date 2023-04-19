import { TestBed } from '@angular/core/testing';
import { AppRoutingModule } from '../app-routing.module';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { HttpClient, HttpClientModule, HttpErrorResponse } from '@angular/common/http';

import { PublicprofileService } from './publicprofile.service';

describe('PublicprofileService', () => {
  let service: PublicprofileService;

  beforeEach(() => {
    TestBed.configureTestingModule({imports:[HttpClientModule,AppRoutingModule,RouterModule, FormsModule]});
    service = TestBed.inject(PublicprofileService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
