// web/src/lib/guild/raid-zones.test.ts
import { describe, expect, it } from 'vitest';
import { timeZoneOptions } from './raid-zones';

const WINTER = new Date('2026-12-10T02:00:00Z');
const SUMMER = new Date('2027-07-10T02:00:00Z');

describe('timeZoneOptions', () => {
  it('labels a zone with its region name and its abbreviation at the instant', () => {
    const americas = timeZoneOptions(WINTER).find((group) => group.region === 'Americas');
    expect(americas?.options).toContainEqual({ value: 'America/Chicago', label: 'Central (Chicago) · CST' });
    const summer = timeZoneOptions(SUMMER).find((group) => group.region === 'Americas');
    expect(summer?.options).toContainEqual({ value: 'America/Chicago', label: 'Central (Chicago) · CDT' });
  });

  it('adds a zone the list lacks under Other, once', () => {
    const groups = timeZoneOptions(WINTER, ['Asia/Tokyo', 'Asia/Tokyo', 'America/Chicago']);
    const other = groups.find((group) => group.region === 'Other');
    expect(other?.options).toEqual([{ value: 'Asia/Tokyo', label: 'Asia/Tokyo · GMT+9' }]);
  });

  it('has no Other group when every zone is listed', () => {
    expect(timeZoneOptions(WINTER, ['Europe/London']).map((g) => g.region)).toEqual([
      'Americas',
      'Europe',
      'Oceania',
    ]);
  });
});
