// web/src/lib/bis/landing-answer.test.ts
import { afterEach, describe, expect, it, vi } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { LandingConfig } from './landing-config';
import { characterBand } from './landing-bands';
import {
  bisAnswerSpec,
  countAnswer,
  excludedNote,
  excludedSlotNames,
  landingAction,
  loadBisAnswer,
  slotCount,
} from './landing-answer';
import * as hover from './hover';
import type { BisBand, BisFile, BisSlot } from './types';

const NOW = new Date('2026-10-10T12:00:00Z');
const minutesAgo = (minutes: number): string => new Date(NOW.getTime() - minutes * 60_000).toISOString();

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
    expect(excludedNote(['Trinkets'])).toBe('Trinkets have no pick at this level range.');
    expect(excludedNote(['Off hand', 'Ranged'])).toBe(
      'Off hand and Ranged have no pick at this level range.',
    );
    expect(excludedNote(['Trinkets', 'Off hand', 'Ranged'])).toBe(
      'Trinkets, Off hand and Ranged have no pick at this level range.',
    );
  });

  it('agrees with a single named slot and says nothing when none is excluded', () => {
    expect(excludedNote(['Neck'])).toBe('Neck has no pick at this level range.');
    expect(excludedNote([])).toBeNull();
  });
});

describe('countAnswer', () => {
  it('is the worn count as one figure with its unit, a bar and the excluded-slot note', () => {
    const answer = countAnswer(bandWithout('trinket1', 'trinket2'), wearing('head', 'neck'));
    expect(answer).toEqual({
      kind: 'figures',
      label: 'Best-in-slot pieces you wear',
      staleLabel: 'Pieces you wore at the last sync',
      figures: [{ value: '2', unit: 'of 15 slots' }],
      bar: 2 / 15,
      note: { text: 'Trinkets have no pick at this level range.' },
    });
  });

  it('has no note when every slot has a pick', () => {
    expect(countAnswer(bandWithout(), wearing('head'))).not.toHaveProperty('note');
  });
});

describe('landingAction', () => {
  it('opens the character’s own spec at its own band and faction', () => {
    expect(landingAction(character({ faction: 'horde' }), CONFIG)).toEqual({
      label: 'Fury gear for 40 to 49 →',
      href: '/bis/warrior/fury?faction=horde#band-horde-40',
    });
  });

  it('names the top band plainly', () => {
    expect(landingAction(character({ level: 60 }), CONFIG).label).toBe('Fury gear for 60 →');
  });

  it('under 20 opens the first band', () => {
    expect(landingAction(character({ level: 12 }), CONFIG)).toEqual({
      label: 'Fury gear for 20 to 29 →',
      href: '/bis/warrior/fury?faction=alliance#band-alliance-20',
    });
  });

  it('a character with no spec yet opens the class’s first spec, labelled by its class', () => {
    expect(landingAction(character({ spec: undefined }), CONFIG)).toEqual({
      label: 'Warrior gear for 40 to 49 →',
      href: '/bis/warrior/arms?faction=alliance#band-alliance-40',
    });
  });

  it('points a class the site does not know at the grid', () => {
    expect(landingAction(character({ class: 'Monk' }), CONFIG)).toEqual({
      label: 'Pick a spec →',
      href: '#bis-classes',
    });
  });
});

describe('loadBisAnswer', () => {
  afterEach(() => vi.restoreAllMocks());

  const file = (band: Pick<BisBand, 'slots'>): BisFile =>
    ({ bands: [{ ...band, band: 40, faction: 'alliance' }] }) as unknown as BisFile;
  const options = { build: 'b1', specKeyOf: () => 'warrior-fury' };

  it('counts against the character’s band in its spec’s file', async () => {
    vi.spyOn(hover, 'fetchBisFile').mockResolvedValue(file(bandWithout()));
    vi.spyOn(hover, 'bandEntryFor').mockReturnValue(bandWithout() as BisBand);
    const answer = await loadBisAnswer(character(), options);
    expect(answer).toMatchObject({ kind: 'figures', figures: [{ value: '2', unit: 'of 17 slots' }] });
  });

  it('says there is no list rather than invent a count', async () => {
    vi.spyOn(hover, 'fetchBisFile').mockResolvedValue(null);
    expect(await loadBisAnswer(character(), options)).toEqual({
      kind: 'sentence',
      text: 'No best in slot list published yet for this spec and level range.',
    });
  });

  it('shows what the page chose when the character has no spec the page knows', async () => {
    const answer = await loadBisAnswer(character(), {
      ...options,
      specKeyOf: () => undefined,
      withoutList: { kind: 'none' },
    });
    expect(answer).toEqual({ kind: 'none' });
  });
});

describe('bisAnswerSpec', () => {
  it('starts at level 20, reads gear and carries the page’s sentences', () => {
    const spec = bisAnswerSpec({ build: 'b1', specKeyOf: () => undefined });
    expect(spec.levelFloor).toBe(20);
    expect(spec.needsGear).toBe(true);
    expect(spec.strings.belowLevelFloor).toBe('Best in slot starts at 20.');
    expect(spec.strings.noGear?.link).toEqual({ label: 'Get the addon', href: '/setup' });
  });
});
