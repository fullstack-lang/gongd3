import { TestBed } from '@angular/core/testing';

import { Gongd3specificService } from './gongd3specific.service';

describe('Gongd3specificService', () => {
  let service: Gongd3specificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(Gongd3specificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
