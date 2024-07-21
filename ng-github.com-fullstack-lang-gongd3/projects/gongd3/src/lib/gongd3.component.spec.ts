import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Gongd3Component } from './gongd3.component';

describe('Gongd3Component', () => {
  let component: Gongd3Component;
  let fixture: ComponentFixture<Gongd3Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Gongd3Component]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(Gongd3Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
