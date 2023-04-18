import { TestBed } from '@angular/core/testing';

import { CprofileService } from './cprofile.service';

describe('CprofileService', () => {
  let service: CprofileService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CprofileService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
