import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute } from '@angular/router';
import { of, throwError } from 'rxjs';

import { CprofileComponent } from './cprofile.component';
import { CprofileService } from './cprofile.service';

describe('CprofileComponent', () => {
  let component: CprofileComponent;
  let fixture: ComponentFixture<CprofileComponent>;
  let cprofileServiceSpy: jasmine.SpyObj<CprofileService>;

  beforeEach(async () => {
    const activatedRouteStub = {
      params: of({name: 'WECE'})
    };

    cprofileServiceSpy = jasmine.createSpyObj('CprofileService', ['getClubInfo']);

    await TestBed.configureTestingModule({
      declarations: [ CprofileComponent ],
      providers: [
        { provide: ActivatedRoute, useValue: activatedRouteStub },
        { provide: CprofileService, useValue: cprofileServiceSpy }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CprofileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call cprofileService.getClubInfo with the correct name', () => {
    expect(cprofileServiceSpy.getClubInfo).toHaveBeenCalledWith('WECE');
  });

  it('should set the club stuff when getClubInfo works', () => {
    const club = { name: 'WECE', VP: 'Sarah Schultz', president:"",treasurer:"",about:"" };
    cprofileServiceSpy.getClubInfo.and.returnValue(of(club));
    component.ngOnInit();
    expect(component.club).toEqual(club);
  });

  it('should log an error when getClubInfo fails', () => {
    const error = new Error('test-error');
    cprofileServiceSpy.getClubInfo.and.returnValue(throwError(error));
    spyOn(console, 'log');
    component.ngOnInit();
    expect(console.log).toHaveBeenCalledWith(error);
  });
});
