import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LoginComponent } from './login.component';
import { LoginService } from './login.service';
import { Router } from '@angular/router';
import { User } from '../user.model';
import { of } from 'rxjs';
import { FormsModule } from '@angular/forms';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let mockLoginService: jasmine.SpyObj<LoginService>;
  let mockRouter: jasmine.SpyObj<Router>;

  beforeEach(async () => {
    mockLoginService = jasmine.createSpyObj('LoginService', ['login']);
    mockRouter = jasmine.createSpyObj('Router', ['navigate']);

    await TestBed.configureTestingModule({
      declarations: [ LoginComponent ],
      imports:[FormsModule],
      providers: [
        { provide: LoginService, useValue: mockLoginService },
        { provide: Router, useValue: mockRouter }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should login', () => {
    const user: User = { id: "1", email: 'tester@example.com',name:"tester",password:"password123",clubs:"wece" };
    mockLoginService.login.and.returnValue(of(user));

    component.email = 'tester@example.com';
    component.password = 'password123';
    component.onSubmit();

    expect(mockLoginService.login).toHaveBeenCalledWith('tester@example.com', 'password123');
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/profile', { User: user }]);
  });

  it('should show error message on incorrect creds', () => {
    mockLoginService.login.and.throwError('Incorrect Username or Password');

    component.email = 'tester@example.com';
    component.password = 'password122';
    component.onSubmit();

    expect(mockLoginService.login).toHaveBeenCalledWith('tester@example.com', 'password122');
    expect(window.alert).toHaveBeenCalledWith('Incorrect Username or Password');
  });

  it('should navigate to register page', () => {
    component.onRegister();

    expect(mockRouter.navigate).toHaveBeenCalledWith(['/register']);
  });
});
