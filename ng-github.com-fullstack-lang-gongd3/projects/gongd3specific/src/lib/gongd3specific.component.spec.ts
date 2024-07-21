import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Gongd3specificComponent } from './gongd3specific.component';

describe('Gongd3specificComponent', () => {
  let component: Gongd3specificComponent;
  let fixture: ComponentFixture<Gongd3specificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ Gongd3specificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Gongd3specificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
