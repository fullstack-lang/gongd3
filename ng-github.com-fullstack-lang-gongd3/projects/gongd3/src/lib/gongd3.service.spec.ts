import { TestBed } from '@angular/core/testing';

import { Gongd3Service } from './gongd3.service';

describe('Gongd3Service', () => {
  let service: Gongd3Service;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Gongd3Service);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
