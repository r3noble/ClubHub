import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LoginComponent } from './login.component';
import { LoginService } from './login.service';
import { RouterTestingModule } from '@angular/router/testing';
import { of, throwError } from 'rxjs';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let loginServiceSpy: jasmine.SpyObj<LoginService>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoginComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call loginService login method on submit', () => {
    const username = 'testUser';
    const password = 'testPassword';
    loginServiceSpy.login.and.returnValue(of(null));
    component.username = username;
    component.password = password;

    component.onSubmit();

    expect(loginServiceSpy.login).toHaveBeenCalledWith(username, password);
  });

  it('should navigate to profile page on successful login', () => {
    const username = 'testUser';
    const password = 'testPassword';
    loginServiceSpy.login.and.returnValue(of(null));
    component.username = username;
    component.password = password;
    spyOn(component.router, 'navigate');

    component.onSubmit();

    expect(component.router.navigate).toHaveBeenCalledWith([`${component.baseUrl}/profile?username=${username}`]);
  });

  it('should set error message and navigate to profile page on failed login', () => {
    const errorMessage = 'Invalid credentials';
    loginServiceSpy.login.and.returnValue(throwError({ message: errorMessage }));
    spyOn(component.router, 'navigate');

    component.onSubmit();

    expect(component.errorMessage).toEqual(errorMessage);
    expect(component.router.navigate).toHaveBeenCalledWith(['/profile']);
  });

  it('should navigate to register page on register button click', () => {
    spyOn(component.router, 'navigate');

    component.onRegister();

    expect(component.router.navigate).toHaveBeenCalledWith(['/register']);
  });
});
