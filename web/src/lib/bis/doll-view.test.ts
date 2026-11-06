import { describe, expect, it } from 'vitest';
import {
  changedCount,
  dollCellFor,
  dollSummary,
  filteredPicks,
  isTie,
  shortStats,
  sourceCounts,
  statDelta,
  whereGroups,
  dollCellsFor,
  type DollViewDeps,
} from './doll-view';
import type { BisBandModel, BisItem, BisRowModel } from './page-model';

const item = (name: string, req: number, stats: Record<string, number>, ilvl = 25): BisItem => ({
  name,
  quality: 3,
  ilvl,
  req,
  icon: 'x',
  stats,
});
const ITEMS: Record<number, BisItem> = {
  1: item("Brawler's Leather Hood", 20, { stamina: 10, agility: 8 }),
  2: item("Trapper's Leather Hood", 18, { agility: 6, stamina: 5 }),
  3: item('Slime-encrusted Pads', 21, { mp5: 3, agility: 3 }, 25),
  4: item("Serpent's Shoulders", 18, { agility: 5 }),
};
const ROW: BisRowModel = {
  slot: 'head',
  label: 'Head',
  id: 1,
  alts: [{ id: 2, delta: -1.52 }],
  sources: {
    1: { line: 'Crafted · Leatherworking', kind: 'crafted', place: 'Leatherworking' },
    2: { line: 'Lady Anacondra · Wailing Caverns (levels 17-24) · 40% drop', kind: 'dungeon', place: 'Wailing Caverns' },
  },
  filter: {
    slot: 'head',
    item_id: 1,
    source_kind: 'crafted',
    by_kind: { dungeon: { item_id: 2, rank: 1, source: 'x', source_kind: 'dungeon', dps_delta: -1.52 } },
  },
};
const SHOULDER: BisRowModel = {
  slot: 'shoulder',
  label: 'Shoulder',
  id: 4,
  alts: [],
  sources: { 4: { line: 'Wailing Caverns', kind: 'dungeon', place: 'Wailing Caverns' } },
  filter: { slot: 'shoulder', item_id: 4, source_kind: 'dungeon' },
};
const OUT: DollViewDeps = { items: ITEMS, unit: 'DPS' };
const MODEL = { band: 20, unit: 'DPS', rows: [ROW, SHOULDER] } as unknown as BisBandModel;

describe('isTie', () => {
  it('is the one rule: under 0.05 DPS is the same', () => {
    expect(isTie(0.04)).toBe(true);
    expect(isTie(-0.049)).toBe(true);
    expect(isTie(0.05)).toBe(false);
  });
});

describe('shortStats and statDelta', () => {
  it('abbreviates the first three stats in order', () => {
    expect(shortStats({ stamina: 10, agility: 8, spirit: 2, intellect: 1 })).toBe('+10 Sta · +8 Agi · +2 Spi');
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

  it('flags the required level and carries the short stats and full source', () => {
    expect(cell.flag).toEqual({ text: 'Level 20', tone: 'muted' });
    expect(cell.stats).toBe('+10 Sta · +8 Agi');
    expect(cell.source).toBe('Crafted · Leatherworking');
  });

  it('says "Same as next best", never "Tie"', () => {
    const tie = dollCellFor({ ...ROW, alts: [{ id: 2, delta: -0.02 }] }, OUT);
    expect(tie.gain?.caption).toBe('Same as next best');
    expect(tie.note?.tail).toBe('(same DPS)');
    expect(JSON.stringify(tie)).not.toMatch(/\bTie\b/);
  });

  it('keeps a tank band in its own unit', () => {
    expect(dollCellFor(ROW, { ...OUT, unit: 'score' }).gain?.caption).toBe('score vs next best');
  });

  it("explains an empty slot with the row's own copy", () => {
    const empty = dollCellFor(
      { slot: 'trinket1', label: 'Trinket', empty: 'The first comes at 40.', alts: [], sources: {}, filter: { slot: 'trinket1', item_id: 0, source_kind: '' } },
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
    expect(cell.note?.tail).toBe('(item level 25, +3 Mp5 · +3 Agi). The pick is +0.8 DPS, +10 Sta, +5 Agi, −3 Mp5.');
  });

  it('flags the level against the character', () => {
    expect(cell.flag).toEqual({ text: 'At 20', tone: 'gold' });
    expect(dollCellFor(ROW, { ...OUT, character: { ...character, level: 24 } }).flag?.text).toBe('Usable now');
  });

  it('ticks a slot that already wears the pick', () => {
    const worn = dollCellFor(ROW, { ...OUT, character: { ...character, worn: { head: 1 } } });
    expect(worn.badge).toBe('worn');
    expect(worn.gain?.caption).toBe('You have this');
  });

  it('counts the upgrades and their sum in the summary line', () => {
    expect(dollSummary([cell], character, 'DPS')).toBe('Felsnap (level 19) has 0 of 1 picks. 1 upgrade, +0.8 DPS available.');
    expect(dollSummary([cell], undefined, 'DPS')).toBe('1 of 1 slots have a pick at this level.');
  });
});

describe('source filters re-pick, never hide', () => {
  it('re-picks a crafted head from dungeon and states the gap to the best pick', () => {
    const picks = filteredPicks(MODEL, { on: ['dungeon'] });
    const cells = dollCellsFor(MODEL, OUT, picks, 'dungeon only');
    expect(cells[0]!.item?.id).toBe(2);
    expect(cells[0]!.gain).toEqual({ value: '−1.5', caption: 'DPS vs the best pick', tone: 'wipe' });
    expect(cells[0]!.note?.lead).toBe('Re-picked (dungeon only). Best overall:');
    expect(changedCount(MODEL, picks)).toBe(1);
  });

  it('keeps a slot with no option visible and explained', () => {
    const picks = filteredPicks(MODEL, { on: ['quest'] });
    const cell = dollCellsFor(MODEL, OUT, picks)[1]!;
    expect(cell.item).toBeNull();
    expect(cell.empty).toMatch(/^No option from the sources you chose\. Best overall: Serpent's Shoulders \(dungeon\)\.$/);
  });

  it('counts slots with an option per kind', () => {
    expect(sourceCounts(MODEL).dungeon).toBe(2);
    expect(sourceCounts(MODEL).crafted).toBe(1);
  });
});

describe('whereGroups', () => {
  it('groups by instance first, most pieces first', () => {
    const picks = filteredPicks(MODEL, { on: ['dungeon'] });
    const groups = whereGroups(MODEL, dollCellsFor(MODEL, OUT, picks));
    expect(groups).toEqual([{ key: 'dungeon:Wailing Caverns', title: 'Wailing Caverns', count: '2 pieces', slots: ['head', 'shoulder'] }]);
  });
});
