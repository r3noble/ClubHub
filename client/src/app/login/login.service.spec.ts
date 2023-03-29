import { TestBed } from '@angular/core/testing';
import { AppRoutingModule } from '../app-routing.module';
import { RouterModule } from '@angular/router';
import appRoutes from '../app-routing.module';
import { HttpClient, HttpClientModule, HttpErrorResponse } from '@angular/common/http';
import { LoginService } from './login.service';

describe('LoginService', () => {
  let service: LoginService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        HttpClientModule,
        AppRoutingModule,
        RouterModule.forRoot(appRoutes),
      ],
    });
    service = TestBed.inject(LoginService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
