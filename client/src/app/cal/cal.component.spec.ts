import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';

import { CalComponent } from './cal.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

describe('CalComponent', () => {
  let component: CalComponent;
  let fixture: ComponentFixture<CalComponent>;
  let httpTestingController: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [CalComponent],
      imports: [HttpClientTestingModule, FormsModule],
      providers: [RouterLink]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CalComponent);
    component = fixture.componentInstance;
    httpTestingController = TestBed.inject(HttpTestingController);
    fixture.detectChanges();
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
