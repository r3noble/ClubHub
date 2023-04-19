import { TestBed } from '@angular/core/testing';

import { RegisterService } from './register.service';

import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpTestingController } from '@angular/common/http/testing';
import { of, throwError } from 'rxjs';
// Other imports
import { HttpClient, HttpClientModule, HttpErrorResponse } from '@angular/common/http';
import { AppRoutingModule } from '../app-routing.module';
import { RouterModule } from '@angular/router';
import appRoutes from '../app-routing.module';
import { FormsModule } from '@angular/forms';


describe('HttpClient testing', () => {
  let httpClient: HttpClientTestingModule;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, FormsModule, HttpClientModule] // add HttpClientModule here
    });

    // Inject the http service and test controller for each test
    httpClient = TestBed.get(HttpClientTestingModule);
    httpTestingController = TestBed.get(HttpTestingController);
  });

  it('works', () => {});
});

describe('RegisterService', () => {
  let service: RegisterService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientModule], // add HttpClientModule here
    });
    service = TestBed.inject(RegisterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
