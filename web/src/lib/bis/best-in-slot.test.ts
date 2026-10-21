// web/src/lib/bis/best-in-slot.test.ts
// W.12 ruling 6: one definition of "best in slot" for the character page's Gear section and the
// BiS card: the worn item is the band's listed pick, a pair matched either way round.
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE } from '../../fixtures/character-gear';
import { gearBestCount } from '../gear/panel';
import { wearsListedBest } from './best-in-slot';
import { slotCount } from './landing-answer';
import type { BisBand } from './types';

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

describe('the page and the card count the same slots', () => {
  it('counts the slots wearing their pick the same on both', () => {
    const gear: Record<string, number> = {};
    for (const slot of GEAR_FIXTURE) if (slot.worn !== null) gear[slot.slot] = slot.worn.id;
    const band = {
      slots: GEAR_FIXTURE.filter((slot) => slot.pick !== null).map((slot) => ({
        slot: slot.slot,
        item_id: slot.pick?.id,
        source: { kind: 'drop', text: 'x' },
      })),
    } as unknown as Pick<BisBand, 'slots'>;
    const page = gearBestCount(GEAR_FIXTURE);
    const card = slotCount(band, gear);
    expect(page.worn).toBeGreaterThan(0);
    expect(card.worn).toBe(page.worn);
    expect(card.total).toBe(page.total);
  });
});
