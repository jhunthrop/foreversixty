// web/src/lib/home/upgrades-parity.test.ts
// The site's half of the home answer's parity check with the API (api/internal/upgrades,
// which the apps' GET /v1/characters/{region}/{ruleset}/{name}/home reads). Both sides run
// their own comparison over one fixture character, api/internal/upgrades/testdata/parity/
// input.json (a real band and item rows of build 1.60.1.70658), and both pin the same answer,
// site.json, so a change to either rule fails until the other follows. UPDATE_PARITY=1
// rewrites site.json from this side.
import { readFileSync, writeFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import type { BisBand } from '../bis/types';
import type { Item } from '../planner/types';
import { upgradesFor, type UpgradesResult } from './upgrades';

const PARITY_DIR = fileURLToPath(
  new URL('../../../../api/internal/upgrades/testdata/parity/', import.meta.url),
);

interface ParityInput {
  band: BisBand;
  gear: Record<string, number>;
  /** The fixture's item rows; an item with no stats or armour leaves the key out. */
  items: (Omit<Item, 'stats' | 'armor'> & Partial<Pick<Item, 'stats' | 'armor'>>)[];
}

/** The answer both sides pin: every upgrade by slot, the already-best slots, the total and counts. */
function parityAnswer(result: UpgradesResult) {
  const round = (value: number): number => Math.round(value * 1e6) / 1e6;
  return {
    upgrades: [...result.upgrades]
      .sort((a, b) => a.slot.localeCompare(b.slot))
      .map((row) => ({
        slot: row.slot,
        pick_id: row.pick.item_id,
        worn_id: row.wornItemId ?? 0,
        worn_unknown: row.wornUnknown,
        gain: row.gainDps === null ? null : round(row.gainDps),
        not_sim_checked: row.notSimChecked,
      })),
    already_best: result.alreadyBis.map((row) => ({
      slot: row.slot,
      item_id: row.itemId,
      item_name: row.itemName,
    })),
    total_gain: round(result.totalGainDps),
    not_sim_checked_count: result.notSimCheckedCount,
    unit: result.scoreUnit,
  };
}

describe('the home answer matches the API on the parity fixture', () => {
  it('pins site.json', () => {
    const input = JSON.parse(readFileSync(`${PARITY_DIR}input.json`, 'utf8')) as ParityInput;
    const character = { build: { gear: input.gear } } as unknown as Pick<MeCharacter, 'build'>;
    const items = new Map(input.items.map((item) => [item.id, { stats: {}, armor: 0, ...item }]));
    const answer = parityAnswer(upgradesFor(character, input.band, items));
    if (process.env.UPDATE_PARITY === '1') {
      writeFileSync(`${PARITY_DIR}site.json`, `${JSON.stringify(answer, null, 1)}\n`);
    }
    expect(answer).toEqual(JSON.parse(readFileSync(`${PARITY_DIR}site.json`, 'utf8')));
  });
});
