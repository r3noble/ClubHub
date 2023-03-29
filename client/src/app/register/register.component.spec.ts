import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import appRoutes, { AppRoutingModule } from '../app-routing.module';
import { RegisterComponent } from './register.component';



describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterComponent ],
      imports: [FormsModule, AppRoutingModule,RouterModule, HttpClientModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

   it('should create', () => {
     expect(component).toBeTruthy();
   });

  it('should navigate to home page on cancel', () => {
    const router = TestBed.inject(Router);
    const spy = spyOn(router, 'navigate');
    component.onCancel();
    expect(spy).toHaveBeenCalledWith(['']);
  });
});
