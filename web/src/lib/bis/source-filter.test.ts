import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { LOOT_KINDS } from '../sim/loot';
import {
  SOURCE_FILTER_KINDS,
  filterKindOfSource,
  filteredBand,
  filteredPick,
  pickKinds,
} from './source-filter';
import type { BisSlot, SourceFilterKind } from './types';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
// The ranker's own golden (sim/cmd/leveling-bis/source_filter_golden_test.go): published rows
// and what the Go resolver gives under each filter. The two resolvers must agree.
const GOLDEN = path.join(REPO_ROOT, 'sim/cmd/leveling-bis/testdata/source-filter-band.json');

interface GoldenBand {
  slots: BisSlot[];
  cases: { kinds: SourceFilterKind[]; want: Record<string, number> }[];
}

const kinds = (...list: SourceFilterKind[]): ReadonlySet<SourceFilterKind> => new Set(list);

function row(partial: Partial<BisSlot>): BisSlot {
  return {
    slot: 'head',
    item_id: 1,
    item_name: 'Pick',
    source: 'Deadmines',
    source_kind: 'dungeon',
    verified: true,
    ...partial,
  };
}

describe('filterKindOfSource', () => {
  it('maps every loot.json kind to one filter kind, the same table as the ranker', () => {
    expect(Object.fromEntries(LOOT_KINDS.map((kind) => [kind, filterKindOfSource(kind)]))).toEqual({
      raid: 'raid',
      dungeon: 'dungeon',
      world: 'world',
      world_drop: 'world',
      zone: 'world',
      vendor: 'vendor',
      crafted: 'crafted',
      rep: 'rep',
      pvp: 'pvp',
      quest: 'quest',
    });
    expect(new Set(LOOT_KINDS.map(filterKindOfSource))).toEqual(new Set(SOURCE_FILTER_KINDS));
  });

  it('reads a pick kinds from kinds, else from its source_kind', () => {
    expect(pickKinds({ kinds: ['dungeon', 'quest'], source_kind: 'quest' })).toEqual(['dungeon', 'quest']);
    expect(pickKinds({ source_kind: 'world_drop' })).toEqual(['world']);
  });
});

describe('filteredPick', () => {
  const head = row({
    by_kind: {
      quest: { item_id: 2, rank: 3, source: 'A Quest', dps_delta: -1.5 },
      world: { item_id: 3, rank: 2, source_kind: 'world_drop', source: 'World drop', dps_delta: -0.5 },
    },
  });

  it('keeps the pick while one of its kinds is on', () => {
    expect(filteredPick(head, kinds('dungeon', 'quest'))).toEqual({
      item_id: 1,
      two_hand: false,
      dps_delta: 0,
    });
  });

  it('takes the best-ranked entry among the kinds still on, with its own gap and source', () => {
    expect(filteredPick(head, kinds('quest', 'world'))).toEqual({
      item_id: 3,
      two_hand: false,
      source: 'World drop',
      source_kind: 'world_drop',
      dps_delta: -0.5,
    });
    expect(filteredPick(head, kinds('quest'))).toMatchObject({
      item_id: 2,
      source_kind: 'quest',
      dps_delta: -1.5,
    });
  });

  it('returns null with the slot for "No <slot> from these sources at this level"', () => {
    expect(filteredPick(head, kinds('raid', 'pvp'))).toBeNull();
    expect(filteredPick(row({ by_kind: undefined }), kinds())).toBeNull();
  });

  it('loads a file published before by_kind: the pick or nothing', () => {
    const old = row({ source_kind: 'vendor' });
    expect(filteredPick(old, kinds('vendor'))?.item_id).toBe(1);
    expect(filteredPick(old, kinds('dungeon'))).toBeNull();
  });

  it('never puts an off hand beside a two-hander, nor a ring the first slot already wears', () => {
    const off = row({ slot: 'off_hand', item_id: 7, source_kind: 'vendor' });
    expect(filteredPick(off, kinds('vendor'), { item_id: 9, two_hand: true })).toBeNull();
    const ring = row({
      slot: 'finger2',
      item_id: 20,
      source_kind: 'quest',
      not_with: [20],
      by_kind: {
        quest: {
          item_id: 21,
          rank: 1,
          source: 'Q',
          not_with: [21],
          next: { item_id: 22, rank: 4, source: 'Q2' },
        },
      },
    });
    expect(filteredPick(ring, kinds('quest'), { item_id: 30, two_hand: false })?.item_id).toBe(20);
    expect(filteredPick(ring, kinds('quest'), { item_id: 20, two_hand: false })?.item_id).toBe(21);
    expect(filteredPick(ring, kinds('quest'), { item_id: 21, two_hand: false })?.item_id).toBe(20);
  });
});

describe('filteredBand matches the ranker resolver', () => {
  const golden = JSON.parse(readFileSync(GOLDEN, 'utf8')) as Record<string, GoldenBand>;
  for (const [name, band] of Object.entries(golden)) {
    for (const testCase of band.cases) {
      it(`${name}: ${testCase.kinds.join('+') || 'nothing'} on`, () => {
        const got: Record<string, number> = {};
        for (const [slot, pick] of filteredBand(band.slots, new Set(testCase.kinds))) {
          if (pick !== null) got[slot] = pick.item_id;
        }
        expect(got).toEqual(testCase.want);
      });
    }
  }
});
