// web/src/lib/bis/landing-lead.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { LandingConfig } from './landing-config';
import { characterBand } from './landing-bands';
import { buildLeadView, excludedNote, excludedSlotNames, slotCount, syncLine } from './landing-lead';
import type { BisBand, BisSlot } from './types';

const NOW = new Date('2026-10-10T12:00:00Z');
const minutesAgo = (minutes: number): string => new Date(NOW.getTime() - minutes * 60_000).toISOString();
const daysAgo = (days: number): string => minutesAgo(days * 24 * 60);

const CONFIG: LandingConfig = {
  build: 'b1',
  classes: {
    warrior: { slug: 'warrior', name: 'Warrior', color: '#c69b6d', defaultSpecKey: 'warrior-arms' },
  },
  specs: {
    'warrior-arms': { classSlug: 'warrior', specSlug: 'arms', name: 'Arms' },
    'warrior-fury': { classSlug: 'warrior', specSlug: 'fury', name: 'Fury' },
  },
};

const SLOT_ORDER = [
  'head',
  'neck',
  'shoulder',
  'back',
  'chest',
  'wrist',
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
  'main_hand',
  'off_hand',
  'ranged',
] as const;

/** A band whose slots all have a pick (item id 100 + index) except the named ones. */
function bandWithout(...missing: string[]): Pick<BisBand, 'slots'> {
  const slots = SLOT_ORDER.map((slot, index) =>
    missing.includes(slot) ? { slot, verified: false } : { slot, item_id: 100 + index, verified: true },
  );
  return { slots: slots as unknown as BisSlot[] };
}

const wearing = (...slots: string[]): Record<string, number> =>
  Object.fromEntries(
    slots.map((slot) => [slot, 100 + SLOT_ORDER.indexOf(slot as (typeof SLOT_ORDER)[number])]),
  );

function character(overrides: Partial<MeCharacter> = {}): MeCharacter {
  return {
    key: 'us/normal/tester',
    region: 'us',
    ruleset: 'normal',
    name: 'Tester',
    class: 'Warrior',
    spec: 'Fury',
    level: 47,
    faction: 'alliance',
    build: { source: 'addon', captured_at: minutesAgo(12), gear: wearing('head', 'neck') },
    ...overrides,
  };
}

describe('characterBand', () => {
  it('rounds a level down into its band and clamps the top', () => {
    expect(characterBand(47)).toBe(40);
    expect(characterBand(20)).toBe(20);
    expect(characterBand(60)).toBe(60);
  });

  it('is undefined for a level under 20 or an unknown level', () => {
    expect(characterBand(12)).toBeUndefined();
    expect(characterBand(undefined)).toBeUndefined();
  });
});

describe('slotCount', () => {
  it('counts worn picks against the slots that have a pick, and names the ones that do not', () => {
    const count = slotCount(
      bandWithout('trinket1', 'trinket2'),
      wearing('head', 'neck', 'chest', 'legs', 'feet', 'wrist'),
    );
    expect(count).toEqual({ worn: 6, total: 15, excluded: ['Trinkets'] });
  });

  it('does not count an item worn in a slot where it is not the pick', () => {
    const count = slotCount(bandWithout(), { head: 999 });
    expect(count.worn).toBe(0);
    expect(count.total).toBe(17);
  });

  it('has nothing excluded when every slot has a pick', () => {
    expect(slotCount(bandWithout(), wearing('head')).excluded).toEqual([]);
  });
});

describe('excluded slot names and the note', () => {
  it('groups a missing pair and keeps paper-doll order', () => {
    expect(excludedSlotNames(['ranged', 'trinket2', 'trinket1', 'off_hand'])).toEqual([
      'Trinkets',
      'Off hand',
      'Ranged',
    ]);
    expect(excludedSlotNames(['finger1', 'finger2'])).toEqual(['Rings']);
  });

  it('names a lone half of a pair by its own label', () => {
    expect(excludedSlotNames(['trinket2'])).toEqual(['Trinket']);
  });

  it('reads the spec’s sentence, plural for a pair or several slots', () => {
    expect(excludedNote(['Trinkets'])).toBe('Trinkets have no pick at this band.');
    expect(excludedNote(['Off hand', 'Ranged'])).toBe('Off hand and Ranged have no pick at this band.');
    expect(excludedNote(['Trinkets', 'Off hand', 'Ranged'])).toBe(
      'Trinkets, Off hand and Ranged have no pick at this band.',
    );
  });

  it('agrees with a single named slot and says nothing when none is excluded', () => {
    expect(excludedNote(['Neck'])).toBe('Neck has no pick at this band.');
    expect(excludedNote([])).toBeNull();
  });
});

describe('syncLine', () => {
  it('reads the age of an addon sync with no stale tail while fresh', () => {
    expect(syncLine({ source: 'addon', captured_at: minutesAgo(12) }, NOW)).toEqual({
      text: 'Addon · 12 minutes ago',
      staleTail: null,
    });
  });

  it('adds the stale words, verbatim, past the stale limit (7 days with no sync history)', () => {
    expect(syncLine({ source: 'addon', captured_at: daysAgo(19) }, NOW)).toEqual({
      text: 'Addon · 19 days ago',
      staleTail: 'Log out or /reload in the game; the companion then syncs it',
    });
  });

  it('uses the player’s own median gap, never under the 3-day floor', () => {
    const build = { source: 'addon' as const, captured_at: daysAgo(4), median_sync_gap_sec: 43_200 };
    expect(syncLine(build, NOW)?.staleTail).not.toBeNull();
    expect(syncLine({ ...build, captured_at: daysAgo(2) }, NOW)?.staleTail).toBeNull();
  });

  it('is absent for a Battle.net build or no build', () => {
    expect(syncLine({ source: 'blizzard', captured_at: minutesAgo(5) }, NOW)).toBeNull();
    expect(syncLine(undefined, NOW)).toBeNull();
  });
});

describe('buildLeadView', () => {
  const input = (
    c: MeCharacter,
    bandEntry: Pick<BisBand, 'slots'> | null = bandWithout('trinket1', 'trinket2'),
  ) => ({
    character: c,
    config: CONFIG,
    bandEntry: (bandEntry ?? undefined) as BisBand | undefined,
    now: NOW,
  });

  it('opens the character’s own spec at its own band and faction, and counts the worn slots', () => {
    const view = buildLeadView(input(character({ faction: 'horde' })))!;
    expect(view.href).toBe('/bis/warrior/fury?faction=horde#band-horde-40');
    expect(view.buttonLabel).toBe('Fury gear for 40 to 49 →');
    expect(view.specClass).toBe('Fury Warrior');
    expect(view.count).toMatchObject({
      kind: 'count',
      worn: 2,
      total: 15,
      label: 'Best-in-slot gear you wear',
      note: 'Trinkets have no pick at this band.',
    });
  });

  it('names the top band plainly in the button', () => {
    const view = buildLeadView(input(character({ level: 60 }), bandWithout()))!;
    expect(view.buttonLabel).toBe('Fury gear for 60 →');
    expect(view.count).toMatchObject({ kind: 'count', total: 17, note: null });
  });

  it('relabels the count and keeps the sync words when the sync is stale', () => {
    const view = buildLeadView(
      input(character({ build: { source: 'addon', captured_at: daysAgo(19), gear: wearing('head') } })),
    )!;
    expect(view.count).toMatchObject({ kind: 'count', label: 'Gear you wore at the last sync' });
    expect(view.sync?.staleTail).not.toBeNull();
  });

  it('under 20: one sentence, no count, and the button opens the first band', () => {
    const view = buildLeadView(input(character({ level: 12 })))!;
    expect(view.count).toEqual({ kind: 'under20' });
    expect(view.buttonLabel).toBe('Fury gear for 20 to 29 →');
    expect(view.href).toBe('/bis/warrior/fury?faction=alliance#band-alliance-20');
  });

  it('no gear on file: the no-sync sentence, never a zero, and the button still works', () => {
    const view = buildLeadView(input(character({ build: undefined })))!;
    expect(view.count).toEqual({ kind: 'nosync' });
    expect(view.sync).toBeNull();
    expect(view.buttonLabel).toBe('Fury gear for 40 to 49 →');
  });

  it('a Battle.net build carries no worn items, so it is also the no-sync state', () => {
    const view = buildLeadView(
      input(character({ build: { source: 'blizzard', captured_at: minutesAgo(5) } })),
    )!;
    expect(view.count).toEqual({ kind: 'nosync' });
  });

  it('gear on file but no list for the band: no count rather than a made-up one', () => {
    expect(buildLeadView(input(character(), null))!.count).toEqual({ kind: 'none' });
  });

  it('a character with no spec yet opens the class’s first spec and is labelled by its class', () => {
    const view = buildLeadView(input(character({ spec: undefined })))!;
    expect(view.href).toBe('/bis/warrior/arms?faction=alliance#band-alliance-40');
    expect(view.buttonLabel).toBe('Warrior gear for 40 to 49 →');
    expect(view.specClass).toBe('Warrior');
  });

  it('is null for a class the site does not know', () => {
    expect(buildLeadView(input(character({ class: 'Monk' })))).toBeNull();
    expect(buildLeadView(input(character({ class: undefined })))).toBeNull();
  });
});
