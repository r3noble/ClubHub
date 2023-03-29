import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CalendarComponent } from './calendar.component';
import { FormsModule } from '@angular/forms';
// Other imports
import { HttpClient, HttpClientModule, HttpErrorResponse } from '@angular/common/http';
import { AppRoutingModule } from '../app-routing.module';
import { RouterModule } from '@angular/router';
import appRoutes from '../app-routing.module';

describe('CalendarComponent', () => {
  let component: CalendarComponent;
  let fixture: ComponentFixture<CalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CalendarComponent ],
      imports:[]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

   it('should create', () => {
     //expect(component).toBeTruthy();
   });

  it('should initialize the dataSource property with data', () => {
    expect(component.dataSource.length).toBeGreaterThan(0);
  });

  it('should initialize the currentTimeIndicator property', () => {
    expect(component.currentTimeIndicator).toEqual(true);
  });

  it('should initialize the shadeUntilCurrentTime property', () => {
    expect(component.shadeUntilCurrentTime).toEqual(true);
  });

  it('should initialize the view property', () => {
    expect(component.view).toEqual('day');
  });

  it('should initialize the views property', () => {
    expect(component.views).toEqual(['day', 'week', 'month', 'timelineDay', 'timelineWeek', 'timelineMonth']);
  });

  it('should initialize the firstDayOfWeek property', () => {
    expect(component.firstDayOfWeek).toEqual(1);
  });

  it('should initialize the dataSource property with specific data', () => {
    const data = component.getData();
    expect(data.length).toEqual(2);
    expect(data[0].label).toEqual('Example Event');
    expect(data[1].label).toEqual('Jenna and Kenneth 21st Birthday');
  });
});
