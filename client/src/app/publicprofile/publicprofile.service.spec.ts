import { TestBed } from '@angular/core/testing';

import { PublicprofileService } from './publicprofile.service';

describe('PublicprofileService', () => {
  let service: PublicprofileService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PublicprofileService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
