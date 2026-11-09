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
    2: {
      line: 'Lady Anacondra · Wailing Caverns (levels 17-24) · 40% drop',
      kind: 'dungeon',
      place: 'Wailing Caverns',
    },
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

  it('reads the gain in plain units and names the next best under it, leaving line 2 to the source', () => {
    expect(cell.gain?.value).toBe('+1.5');
    expect(cell.gain?.unit).toBe('DPS');
    expect(cell.gain?.tone).toBe('strong');
    expect(cell.gain?.basis?.prefix).toBe('vs');
    expect(cell.gain?.basis?.item?.name).toBe("Trapper's Leather Hood");
    expect(cell.note).toBeUndefined();
  });

  it('flags the required level and carries the short stats and full source', () => {
    expect(cell.flag).toEqual({ text: 'Level 20', tone: 'muted' });
    expect(cell.stats).toBe('+10 Sta · +8 Agi');
    expect(cell.source).toBe('Crafted · Leatherworking');
  });

  it('says "Same as next best", never "Tie"', () => {
    const tie = dollCellFor({ ...ROW, alts: [{ id: 2, delta: -0.02 }] }, OUT);
    expect(tie.gain?.value).toBe('Same');
    expect(JSON.stringify(tie)).not.toMatch(/\bTie\b/);
  });

  it('names the nearest alternative that does not tie when the runner-up ties', () => {
    const tie = dollCellFor(
      {
        ...ROW,
        alts: [
          { id: 2, delta: -0.02 },
          { id: 3, delta: -0.04 },
          { id: 4, delta: -1.1 },
          { id: 5, delta: -0.4 },
        ],
      },
      OUT,
    );
    expect(tie.gain?.basis).toEqual({
      prefix: 'next: −0.4 vs',
      item: { id: 5, name: 'Item 5', quality: 1, icon: undefined },
    });
    expect(JSON.stringify(tie.gain)).not.toContain("Trapper's Leather Hood");
  });

  it('keeps the gap of the next different in the band unit', () => {
    const tie = dollCellFor(
      {
        ...ROW,
        alts: [
          { id: 2, delta: 0 },
          { id: 4, delta: -2.25 },
        ],
      },
      { ...OUT, unit: 'score' },
    );
    expect(tie.gain?.basis?.item?.name).toBe("Serpent's Shoulders");
    expect(tie.gain?.basis?.prefix).toBe('next: −2.3 vs');
    expect(tie.gain?.value).toBe('Same');
  });

  it('says "No close alternative" when every runner-up ties', () => {
    const tie = dollCellFor(
      {
        ...ROW,
        alts: [
          { id: 2, delta: -0.01 },
          { id: 3, delta: 0.03 },
        ],
      },
      OUT,
    );
    expect(tie.gain?.basis).toEqual({ prefix: 'No close alternative' });
    expect(tie.gain?.value).toBe('Same');
  });

  it('keeps the row to two stats and leaves the enchant to the detail', () => {
    const enchanted = dollCellFor(
      { ...ROW, enchant: 'Lesser Arcanum of Voracity' },
      { ...OUT, items: { ...ITEMS, 1: item('Hood', 20, { stamina: 10, agility: 8, spirit: 2 }) } },
    );
    expect(enchanted.stats).toBe('+10 Sta · +8 Agi');
    expect(JSON.stringify(enchanted)).not.toContain('Voracity');
  });

  it('keeps a tank band in its own unit', () => {
    expect(dollCellFor(ROW, { ...OUT, unit: 'score' }).gain?.unit).toBe('score');
  });

  it("explains an empty slot with the row's own copy", () => {
    const empty = dollCellFor(
      {
        slot: 'trinket1',
        label: 'Trinket',
        empty: 'The first comes at 40.',
        alts: [],
        sources: {},
        filter: { slot: 'trinket1', item_id: 0, source_kind: '' },
      },
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
    expect(cell.gain?.value).toBe('+0.8');
    expect(cell.gain?.tone).toBe('gold');
    expect(cell.gain?.basis?.prefix).toBe('vs your');
    expect(cell.gain?.basis?.item?.name).toBe('Slime-encrusted Pads');
    expect(cell.badge).toBe('upgrade');
    expect(cell.note?.label).toBe('You wear');
    expect(cell.note?.tail).toBe(
      '(item level 25, +3 Mp5 · +3 Agi). The pick is +0.8 DPS, +10 Sta, +5 Agi, −3 Mp5.',
    );
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
    expect(worn.gain?.basis).toEqual({ prefix: 'You have this' });
  });

  it('counts the upgrades and their sum in the summary line', () => {
    expect(dollSummary([cell], character, 'DPS')).toBe(
      'Felsnap (level 19) has 0 of 1 picks. 1 upgrade, +0.8 DPS available.',
    );
    expect(dollSummary([cell], undefined, 'DPS')).toBe('1 of 1 slots have a pick at this level.');
    expect(dollSummary([cell], { ...character, gains: { [cell.slot]: null } }, 'DPS')).toBe(
      'Felsnap (level 19) has 0 of 1 picks. 1 upgrade, not sim-checked yet.',
    );
  });
});

describe('source filters re-pick, never hide', () => {
  it('re-picks a crafted head from dungeon and states the gap to the best pick', () => {
    const picks = filteredPicks(MODEL, { on: ['dungeon'] });
    const cells = dollCellsFor(MODEL, OUT, picks, 'dungeon only');
    expect(cells[0]!.item?.id).toBe(2);
    expect(cells[0]!.gain?.value).toBe('−1.5');
    expect(cells[0]!.gain?.tone).toBe('wipe');
    expect(cells[0]!.gain?.basis?.prefix).toBe('vs best:');
    expect(cells[0]!.note?.label).toBe('Re-picked');
    expect(cells[0]!.note?.lead).toBe('Re-picked (dungeon only). Best overall:');
    expect(changedCount(MODEL, picks)).toBe(1);
  });

  it('keeps a slot with no option visible and explained', () => {
    const picks = filteredPicks(MODEL, { on: ['quest'] });
    const cell = dollCellsFor(MODEL, OUT, picks)[1]!;
    expect(cell.item).toBeNull();
    expect(cell.empty).toMatch(
      /^No option from the sources you chose\. Best overall: Serpent's Shoulders \(dungeon\)\.$/,
    );
  });

  it("caps every kind at the leveling default's standing, not only the rep kind (P2 review, 2026-11-09)", () => {
    // Protector's Band: Silverwing Sentinels Honored, sold by their quartermaster at the same
    // standing, so its kinds are rep and vendor and both are above Friendly.
    const ring = {
      ...ROW,
      slot: 'finger1',
      id: 20439,
      filter: {
        slot: 'finger1',
        item_id: 20439,
        source_kind: 'rep',
        kinds: ['rep', 'vendor'],
        standing: 'honored',
        kind_standings: { rep: 'honored', vendor: 'honored' },
        by_kind: { dungeon: { item_id: 1156, rank: 2, source: 'The Deadmines: Gilnid' } },
      },
    } as unknown as BisRowModel;
    const model = { band: 20, unit: 'DPS', rows: [ring] } as unknown as BisBandModel;
    const all = ['dungeon', 'raid', 'world', 'quest', 'crafted', 'vendor', 'pvp'] as const;
    const picks = filteredPicks(model, { on: [...all], repCap: 'friendly' });
    expect(picks.get('finger1')?.item_id).toBe(1156);
  });

  it('never lets a by_kind entry above the cap stand in, whatever its kind', () => {
    const ring = {
      ...ROW,
      slot: 'finger1',
      id: 9001,
      filter: {
        slot: 'finger1',
        item_id: 9001,
        source_kind: 'rep',
        standing: 'revered',
        by_kind: {
          vendor: {
            item_id: 20439,
            rank: 2,
            source: 'Illiyana Moonblaze',
            standing: 'honored',
            next: { item_id: 2246, rank: 4, source: 'Some Vendor' },
          },
        },
      },
    } as unknown as BisRowModel;
    const model = { band: 20, unit: 'DPS', rows: [ring] } as unknown as BisBandModel;
    const all = ['dungeon', 'raid', 'world', 'quest', 'crafted', 'vendor', 'pvp'] as const;
    const picks = filteredPicks(model, { on: [...all], repCap: 'friendly' });
    expect(picks.get('finger1')?.item_id).toBe(2246);
  });

  it('keeps a vendor kind no standing gates under the cap', () => {
    const ring = {
      ...ROW,
      slot: 'finger1',
      id: 20440,
      filter: {
        slot: 'finger1',
        item_id: 20440,
        source_kind: 'rep',
        kinds: ['rep', 'vendor'],
        standing: 'honored',
        kind_standings: { rep: 'honored' },
        by_kind: { dungeon: { item_id: 1156, rank: 2, source: 'The Deadmines: Gilnid' } },
      },
    } as unknown as BisRowModel;
    const model = { band: 20, unit: 'DPS', rows: [ring] } as unknown as BisBandModel;
    const all = ['dungeon', 'raid', 'world', 'quest', 'crafted', 'vendor', 'pvp'] as const;
    const picks = filteredPicks(model, { on: [...all], repCap: 'friendly' });
    expect(picks.get('finger1')?.item_id).toBe(20440);
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
    expect(groups).toEqual([
      {
        key: 'dungeon:Wailing Caverns',
        title: 'Wailing Caverns',
        count: '2 pieces',
        slots: ['head', 'shoulder'],
      },
    ]);
  });
});
