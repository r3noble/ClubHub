import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute, convertToParamMap } from '@angular/router';
import { AuthService } from '../login/auth.service';
import { ProfileComponent } from './profile.component';

describe('ProfileComponent', () => {
  let component: ProfileComponent;
  let fixture: ComponentFixture<ProfileComponent>;
  let authServiceSpy: jasmine.SpyObj<AuthService>;
  const user = { name: 'tester', email: 'tester@example.com', clubs: 'WECE', password:"password123" };

  beforeEach(async () => {
    const authServiceMock = jasmine.createSpyObj('AuthService', ['getUser']);
    authServiceMock.getUser.and.returnValue(user);

    await TestBed.configureTestingModule({
      declarations: [ProfileComponent],
      providers: [
        {
          provide: ActivatedRoute,
          useValue: {
            snapshot: {
              data: { User: user },
              paramMap: convertToParamMap({}),
            },
          },
        },
        { provide: AuthService, useValue: authServiceMock },
      ],
    }).compileComponents();

    authServiceSpy = TestBed.inject(AuthService) as jasmine.SpyObj<AuthService>;
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should set the name and email properties from the user object', () => {
    expect(component.name).toEqual(user.name);
    expect(component.email).toEqual(user.email);
  });

  it('should set the clubs property from the user object if it exists', () => {
    expect(component.clubs).toEqual(user.clubs);
  });

  it('should set the clubs property to "No clubs joined yet!" if no clubs joined', () => {
    authServiceSpy.getUser.and.returnValue({ name: 'tester', email: 'tester@example.com' , id:"1",clubs:"",password:"password123"});
    fixture.detectChanges();
    expect(component.clubs).toEqual('No clubs joined yet!');
  });
});
