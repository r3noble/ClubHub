import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FooterComponent } from './footer.component';
import appRoutes from '../app-routing.module';
import { AppRoutingModule } from '../app-routing.module';
import { RouterModule } from '@angular/router';

describe('FooterComponent', () => {
  let component: FooterComponent;
  let fixture: ComponentFixture<FooterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        FooterComponent,
       ],
       imports: [
        AppRoutingModule,
        RouterModule.forRoot(appRoutes),
        ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

   it('should create', () => {
     expect(component).toBeTruthy();
   });
});
