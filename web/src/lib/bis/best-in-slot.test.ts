// web/src/lib/bis/best-in-slot.test.ts
// W.12 ruling 6: one definition of "best in slot" for the character page's Gear section and the
// BiS card: the worn item is the band's listed pick, a pair matched either way round.
import { describe, expect, it } from 'vitest';
import { wearsListedBest } from './best-in-slot';

describe('wearsListedBest', () => {
  it('is the worn item being the listed pick', () => {
    expect(wearsListedBest('head', 5, { head: 5 })).toBe(true);
    expect(wearsListedBest('head', 5, { head: 6 })).toBe(false);
    expect(wearsListedBest('head', 5, {})).toBe(false);
  });
  it('matches a ring or trinket pair either way round', () => {
    expect(wearsListedBest('finger1', 7, { finger1: 8, finger2: 7 })).toBe(true);
    expect(wearsListedBest('trinket2', 9, { trinket1: 9 })).toBe(true);
    expect(wearsListedBest('head', 9, { trinket1: 9 })).toBe(false);
  });
});
