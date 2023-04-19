import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PublicprofileComponent } from './publicprofile.component';

describe('PublicprofileComponent', () => {
  let component: PublicprofileComponent;
  let fixture: ComponentFixture<PublicprofileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PublicprofileComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PublicprofileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
