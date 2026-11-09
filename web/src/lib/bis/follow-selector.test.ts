// web/src/lib/bis/follow-selector.test.ts
import { describe, expect, it } from 'vitest';
import {
  bandFromParam,
  decideBisDefault,
  decideBisFollow,
  defaultSpecKeyFor,
  type BisPageState,
} from './follow-selector';

const PAGE: BisPageState = { specKey: 'hunter-beast-mastery', faction: 'alliance', band: 20 };

describe('decideBisFollow', () => {
  it('goes to the new character’s page when the class differs, with its band and faction', () => {
    expect(
      decideBisFollow(PAGE, { classSlug: 'warrior', specKey: 'warrior-arms', level: 20, faction: 'horde' }),
    ).toEqual({ kind: 'navigate', href: '/bis/warrior/arms?faction=horde' });
  });

  it('goes to the other spec’s page when only the spec differs', () => {
    expect(
      decideBisFollow(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-marksmanship',
        level: 45,
        faction: 'alliance',
      }),
    ).toEqual({ kind: 'navigate', href: '/bis/hunter/marksmanship/40-49' });
  });

  it('updates band and faction in place when class and spec match', () => {
    expect(
      decideBisFollow(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-beast-mastery',
        level: 34,
        faction: 'horde',
      }),
    ).toEqual({ kind: 'update', faction: 'horde', band: 30 });
  });

  it('stays when class, spec, band and faction already match', () => {
    expect(
      decideBisFollow(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-beast-mastery',
        level: 25,
        faction: 'alliance',
      }),
    ).toEqual({ kind: 'stay' });
  });

  it('sends a character with no known spec to the class’s default spec page', () => {
    const fallback = defaultSpecKeyFor('mage');
    expect(fallback).toBeDefined();
    const action = decideBisFollow(PAGE, { classSlug: 'mage', level: 20, faction: 'alliance' });
    expect(action.kind).toBe('navigate');
    expect(action).toMatchObject({ href: expect.stringContaining('/bis/mage/') as string });
  });

  it('does nothing for a pasted export that names no class', () => {
    expect(decideBisFollow(PAGE, { classSlug: '' })).toEqual({ kind: 'stay' });
  });

  it('uses the 60 band for a level 60 character', () => {
    expect(
      decideBisFollow(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-beast-mastery',
        level: 60,
        faction: 'alliance',
      }),
    ).toEqual({ kind: 'update', faction: 'alliance', band: 60 });
  });

  it('keeps the page’s band and faction for a class-only pointer on a different class', () => {
    expect(decideBisFollow(PAGE, { classSlug: 'warrior', specKey: 'warrior-fury' })).toEqual({
      kind: 'navigate',
      href: '/bis/warrior/fury',
    });
  });
});

describe('decideBisDefault', () => {
  it('opens on the selected character’s band and faction when it is this spec', () => {
    expect(
      decideBisDefault(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-beast-mastery',
        level: 34,
        faction: 'horde',
      }),
    ).toEqual({ kind: 'update', faction: 'horde', band: 30 });
  });

  it('keeps the band 20 list for a character under 20', () => {
    expect(
      decideBisDefault(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-beast-mastery',
        level: 12,
        faction: 'alliance',
      }),
    ).toEqual({ kind: 'stay' });
  });

  it('never leaves a page the player opened for another spec', () => {
    expect(
      decideBisDefault(PAGE, {
        classSlug: 'hunter',
        specKey: 'hunter-marksmanship',
        level: 45,
        faction: 'alliance',
      }),
    ).toEqual({ kind: 'stay' });
    expect(decideBisDefault(PAGE, { classSlug: 'warrior', level: 45 })).toEqual({ kind: 'stay' });
  });
});

describe('bandFromParam', () => {
  it('reads ?band=60 and a raw level as their band', () => {
    expect(bandFromParam('60')).toBe(60);
    expect(bandFromParam('34')).toBe(30);
    expect(bandFromParam('5')).toBe(20);
  });

  it('ignores anything that is not a level', () => {
    expect(bandFromParam(null)).toBeNull();
    expect(bandFromParam('raid')).toBeNull();
    expect(bandFromParam('600')).toBeNull();
  });
});
