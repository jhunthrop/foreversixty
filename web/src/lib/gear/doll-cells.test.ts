// web/src/lib/gear/doll-cells.test.ts
import { describe, expect, it } from 'vitest';
import type { BisBand, BisSlot } from '../bis/types';
import type { Item } from '../planner/types';
import { dollCellsFromSlots, wornGain, wornGainSummary } from './doll-cells';
import { bandPricer, UNPRICED } from './gain-pricer';
import type { PaperdollItemRef, PaperdollSlot } from './paperdoll';

const PICK: PaperdollItemRef = { id: 12640, name: 'Lionheart Helm', icon: 'inv_helmet_36', quality: 4 };

function worn(slot: PaperdollSlot['slot'], id: number, extra: Partial<PaperdollSlot> = {}): PaperdollSlot {
  return {
    slot,
    state: 'item',
    item: { id, name: `Item ${id}`, icon: 'inv_x', quality: 3, itemLevel: 62, twoHand: false },
    ...extra,
  };
}

describe('wornGain (§4.2 worn)', () => {
  it('an upgrade: the gain in the band’s unit, gold, against the pick', () => {
    expect(wornGain({ kind: 'gain', value: 1.84, unit: 'dps', pick: PICK })).toEqual({
      gain: {
        value: '+1.8',
        unit: 'DPS',
        basis: {
          prefix: 'vs',
          item: { id: 12640, name: 'Lionheart Helm', quality: 4, icon: 'inv_helmet_36' },
        },
        tone: 'gold',
      },
      badge: 'upgrade',
    });
    expect(wornGain({ kind: 'gain', value: 3, unit: 'hps', pick: PICK }).gain?.unit).toBe('HPS');
    expect(wornGain({ kind: 'gain', value: 3, unit: 'tank_score', pick: PICK }).gain?.unit).toBe('score');
  });

  it('a gain the rule cannot price says "Not sim-checked", never 0', () => {
    const { gain } = wornGain({ kind: 'gain', value: null, unit: 'dps', pick: PICK });
    expect(gain).toMatchObject({ value: '—', basis: { prefix: 'Not sim-checked' }, tone: 'muted' });
  });

  it('a gain under the tie floor is "Same" (doll-view isTie), not an upgrade', () => {
    expect(wornGain({ kind: 'gain', value: 0.04, unit: 'dps', pick: PICK })).toMatchObject({
      gain: { value: 'Same', tone: 'muted' },
    });
    expect(wornGain({ kind: 'gain', value: 0.04, unit: 'dps', pick: PICK }).badge).toBeUndefined();
  });

  it('"Best in slot" when worn, "Not scored" with the pick when unscored', () => {
    expect(wornGain({ kind: 'best' })).toEqual({
      gain: { value: 'Best in slot', tone: 'kill' },
      badge: 'worn',
    });
    expect(wornGain({ kind: 'unscored', pick: PICK }).gain).toMatchObject({
      value: 'Not scored',
      basis: { prefix: 'vs', item: { id: 12640 } },
    });
  });
});

describe('dollCellsFromSlots', () => {
  it('worn: the item, its item level then its enchant on line 2; the gain', () => {
    const slots = [
      worn('head', 1, {
        enchant: '+8 Stamina',
        upgrade: { kind: 'gain', value: 2, unit: 'dps', pick: PICK },
      }),
      worn('chest', 2),
    ];
    const [head, chest] = dollCellsFromSlots(slots, 'worn', 'warrior');
    expect(head).toMatchObject({
      slot: 'head',
      label: 'Head',
      item: { id: 1, name: 'Item 1', quality: 3, icon: 'inv_x' },
      source: 'Item level 62',
      enchant: '+8 Stamina',
      badge: 'upgrade',
      gain: { value: '+2.0' },
    });
    expect(chest).toMatchObject({ source: 'Item level 62' });
    expect(chest.enchant).toBeUndefined();
    expect(chest.gain).toBeUndefined();
  });

  it('report: the log item level and its enchant, and never a gain', () => {
    const slots = [
      worn('head', 1, {
        detail: 'Item level 66',
        enchant: 'Crusader',
        upgrade: { kind: 'gain', value: 2, unit: 'dps', pick: PICK },
      }),
    ];
    const [head] = dollCellsFromSlots(slots, 'report');
    expect(head).toMatchObject({ source: 'Item level 66', enchant: 'Crusader' });
    expect(head.gain).toBeUndefined();
  });

  it('an unknown item keeps the shipped copy; a plain item (failed file) reads common', () => {
    const [legs, ring] = dollCellsFromSlots(
      [
        { slot: 'legs', state: 'unknown', unknownId: 288887 },
        {
          slot: 'finger1',
          state: 'item',
          item: { id: 5, name: 'Item 5', icon: null, quality: null, itemLevel: null, twoHand: false },
        },
      ],
      'worn',
    );
    expect(legs).toMatchObject({
      item: { id: 288887, name: 'Legs · Item 288887' },
      unknown: true,
      source: 'Not in our item data yet',
    });
    expect(ring.item).toEqual({ id: 5, name: 'Item 5', quality: 1, icon: undefined });
    expect(ring.source).toBeUndefined();
  });

  it('an empty slot reads "Empty"; the off hand under a two-hander carries it greyed', () => {
    const slots: PaperdollSlot[] = [
      { slot: 'neck', state: 'empty' },
      { ...worn('main_hand', 9), item: { ...worn('main_hand', 9).item!, twoHand: true } },
      { slot: 'off_hand', state: 'empty' },
    ];
    const [neck, , off] = dollCellsFromSlots(slots, 'worn');
    expect(neck).toMatchObject({ item: null, emptyTitle: 'Empty' });
    expect(neck.coveredBy).toBeUndefined();
    expect(off).toMatchObject({ item: null, emptyTitle: 'Two-hand in main hand', coveredBy: { id: 9 } });
  });
});

describe('wornGainSummary', () => {
  it('sums the priced gains that are not ties, in their unit', () => {
    const gain = (value: number | null): PaperdollSlot =>
      worn('head', 1, { upgrade: { kind: 'gain', value, unit: 'dps', pick: PICK } });
    expect(wornGainSummary([gain(1.8), gain(2.1), gain(null), gain(0.01), worn('neck', 2)])).toBe(
      '+3.9 DPS across 2 slots',
    );
    expect(wornGainSummary([gain(1)])).toBe('+1.0 DPS across 1 slot');
    expect(wornGainSummary([gain(null)])).toBeNull();
  });
});

describe('bandPricer (lib/home/upgrades.ts slotGainFor)', () => {
  const slot = (row: Partial<BisSlot>): BisSlot => ({
    slot: 'head',
    item_id: 0,
    item_name: '',
    source: '',
    source_kind: 'dungeon',
    verified: true,
    ...row,
  });
  const band = {
    spec: 'warrior-fury',
    band: 60,
    faction: 'horde',
    race: 'orc',
    talents: '',
    talent_points: 51,
    weights: [{ stat: 'strength', weight: 2 }],
    slots: [
      slot({
        slot: 'head',
        item_id: 100,
        score: 50,
        alternatives: [{ item_id: 7, item_name: '', source_kind: '', source: '', dps_delta: -1.5 }],
      }),
      slot({ slot: 'finger1', item_id: 200, score: 30 }),
      slot({ slot: 'main_hand', item_id: 300 }),
    ],
    set_dps: 500,
    no_source_count: 0,
    coverage: {},
    new_at_band: [],
    weights_run_seconds: 0,
    verify_run_seconds: 0,
    reference_dps_per_point: 0.1,
  } as BisBand;
  const items = new Map<number, Item>([[8, { id: 8, stats: { strength: 10 }, armor: 0 } as unknown as Item]]);
  const pricer = bandPricer(band, (id) => items.get(id));

  it('reads a listed runner-up’s sim delta, else the score difference in DPS', () => {
    expect(pricer.unit).toBe('dps');
    expect(pricer.price('head', 100, 7)).toBe(1.5);
    expect(pricer.price('head', 100, 8)).toBeCloseTo(3);
    expect(pricer.price('head', 100, undefined)).toBeCloseTo(5);
  });

  it('prices a ring pick listed under the other ring slot', () => {
    expect(pricer.price('finger2', 200, undefined)).toBeCloseTo(3);
  });

  it('cannot price a weapon by stat weights, an unknown worn id, or a pick the band lacks', () => {
    expect(pricer.price('main_hand', 300, 8)).toBeNull();
    expect(pricer.price('head', 100, 999)).toBeNull();
    expect(pricer.price('head', 555, 8)).toBeNull();
    expect(UNPRICED.price('head', 100, 8)).toBeNull();
  });
});
