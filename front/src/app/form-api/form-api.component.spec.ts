import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormApiComponent } from './form-api.component';

describe('FormApiComponent', () => {
  let component: FormApiComponent;
  let fixture: ComponentFixture<FormApiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FormApiComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FormApiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
