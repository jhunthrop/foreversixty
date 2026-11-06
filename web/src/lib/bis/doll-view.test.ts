import { describe, expect, it } from 'vitest';
import { dollCellFor, dollSummary, isTie, shortStats, statDelta, type DollViewDeps } from './doll-view';
import type { RowView } from './panel-view';
import type { ItemDetail } from './types';

const detail = (name: string, required: number, stats: Record<string, number>, ilvl = 25): ItemDetail => ({
  name,
  quality: 3,
  item_level: ilvl,
  required_level: required,
  icon: 'x',
  stats,
});
const DETAILS = new Map<number, ItemDetail>([
  [1, detail("Brawler's Leather Hood", 20, { stamina: 10, agility: 8 })],
  [2, detail("Trapper's Leather Hood", 18, { agility: 6, stamina: 5 })],
  [3, detail('Slime-encrusted Pads', 21, { mp5: 3, agility: 3 }, 25)],
]);
const ROW: RowView = {
  slot: 'head',
  slotLabel: 'Head',
  empty: false,
  itemId: 1,
  itemName: "Brawler's Leather Hood",
  sourceKind: 'crafted',
  sourceDetail: 'Crafted · Leatherworking',
  alternatives: [
    {
      itemId: 2,
      itemName: "Trapper's Leather Hood",
      sourceKind: 'dungeon',
      sourceDetail: 'x',
      dpsDelta: -1.52,
    },
  ],
};
const OUT: DollViewDeps = { details: DETAILS, unit: 'DPS' };

describe('isTie', () => {
  it('is the one rule: under 0.05 DPS is the same', () => {
    expect(isTie(0.04)).toBe(true);
    expect(isTie(-0.049)).toBe(true);
    expect(isTie(0.05)).toBe(false);
  });
});

describe('shortStats and statDelta', () => {
  it('abbreviates the first three stats in order', () => {
    expect(shortStats({ stamina: 10, agility: 8, spirit: 2, intellect: 1 })).toBe(
      '+10 Sta · +8 Agi · +2 Spi',
    );
    expect(shortStats({})).toBeUndefined();
  });

  it('spells the pick minus the worn item, largest first', () => {
    expect(statDelta({ agility: 8, stamina: 10 }, { mp5: 3, agility: 3 })).toBe('+10 Sta, +5 Agi, −3 Mp5');
  });
});

describe('dollCellFor, signed out', () => {
  const cell = dollCellFor(ROW, OUT);

  it('reads the gain in plain units against the next best, and names it on line 2', () => {
    expect(cell.gain).toEqual({ value: '+1.5', caption: 'DPS vs next best', tone: 'strong' });
    expect(cell.note?.lead).toBe('Next best:');
    expect(cell.note?.item?.name).toBe("Trapper's Leather Hood");
    expect(cell.note?.tail).toBe('−1.5 DPS');
  });

  it('flags the required level and carries the short stats and source', () => {
    expect(cell.flag).toEqual({ text: 'Level 20', tone: 'muted' });
    expect(cell.stats).toBe('+10 Sta · +8 Agi');
    expect(cell.source).toBe('Crafted · Leatherworking');
  });

  it('says "Same as next best", never "Tie"', () => {
    const tie = dollCellFor({ ...ROW, alternatives: [{ ...ROW.alternatives![0], dpsDelta: -0.02 }] }, OUT);
    expect(tie.gain?.caption).toBe('Same as next best');
    expect(tie.note?.tail).toBe('(same DPS)');
    expect(JSON.stringify(tie)).not.toMatch(/\bTie\b/);
  });

  it('keeps a tank band in its own unit', () => {
    expect(dollCellFor(ROW, { ...OUT, unit: 'score' }).gain?.caption).toBe('points vs next best');
  });

  it("explains an empty slot with the row's own copy", () => {
    const empty = dollCellFor(
      { slot: 'trinket1', slotLabel: 'Trinket', empty: true, emptyCopy: 'The first comes at 40.' },
      OUT,
    );
    expect(empty.item).toBeNull();
    expect(empty.empty).toBe('The first comes at 40.');
  });
});

describe('dollCellFor, signed in', () => {
  const character = { name: 'Felsnap', level: 19, worn: { head: 3 }, gains: { head: 0.8 } };
  const cell = dollCellFor(ROW, { ...OUT, character });

  it('compares with the worn item by name, in gold, with its item level and stat delta', () => {
    expect(cell.gain).toEqual({ value: '+0.8', caption: 'DPS vs your Slime-encrusted Pads', tone: 'gold' });
    expect(cell.badge).toBe('upgrade');
    expect(cell.note?.lead).toBe('You wear');
    expect(cell.note?.tail).toBe('(item level 25) · The pick: +10 Sta, +5 Agi, −3 Mp5');
  });

  it('flags the level against the character', () => {
    expect(cell.flag).toEqual({ text: 'At 20', tone: 'gold' });
    expect(dollCellFor(ROW, { ...OUT, character: { ...character, level: 24 } }).flag?.text).toBe(
      'Usable now',
    );
  });

  it('ticks a slot that already wears the pick', () => {
    const worn = dollCellFor(ROW, { ...OUT, character: { ...character, worn: { head: 1 } } });
    expect(worn.badge).toBe('worn');
    expect(worn.gain?.caption).toBe('You have this');
  });

  it('counts the upgrades and their sum in the summary line', () => {
    expect(dollSummary([cell], character, 'DPS')).toBe(
      'Felsnap (level 19) has 0 of 1 picks. 1 upgrade, +0.8 DPS available.',
    );
    expect(dollSummary([cell], undefined, 'DPS')).toBe('1 of 1 slots have a pick at this level.');
  });
});
