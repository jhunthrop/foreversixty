import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { decodeFS1 } from '../planner/fs1';
import { indexTalents } from '../planner/rules';
import type { TalentFile } from '../planner/types';
import { parseSimState } from '../sim/url';
import { bisSimHref } from './sim-link';
import type { BisBand, BisSlot } from './types';

const index = indexTalents(fixtureTalents as TalentFile);

function slot(name: string, itemId: number): BisSlot {
  return {
    slot: name,
    item_id: itemId,
    item_name: `i${itemId}`,
    source: 'X',
    source_kind: 'drop',
    verified: true,
  };
}

function band(overrides: Partial<BisBand>): BisBand {
  return {
    spec: 'warrior-fury',
    band: 60,
    faction: 'horde',
    race: 'troll',
    // Five points in Cruelty and five in Booming Voice, the two tier-0 Fury talents.
    talents: '0000000-5500000',
    talent_points: 10,
    weights: [],
    slots: [slot('head', 12640), slot('main_hand', 1728), slot('ranged', 12651)],
    set_dps: 772.5,
    no_source_count: 0,
    coverage: {},
    new_at_band: [],
    weights_run_seconds: 0,
    verify_run_seconds: 0,
    ...overrides,
  } as BisBand;
}

function decodedFrom(href: string) {
  expect(href.startsWith('/sim?')).toBe(true);
  const decoded = decodeFS1(parseSimState(href.slice('/sim'.length)).code);
  if (!decoded.ok) throw new Error(decoded.message);
  return decoded.build;
}

describe('bisSimHref', () => {
  it('carries the enchants the band published, so the simulator measures the same set', () => {
    const slots = [{ ...slot('head', 12640), enchant_id: 1506 }, slot('main_hand', 1728)];
    const gear = decodedFrom(bisSimHref(band({ slots }), index, '1.60.1.70790')).gearSlots;
    expect(gear).toContainEqual({ slot: 'head', itemId: 12640, enchant: 1506 });
    expect(gear).toContainEqual({ slot: 'main_hand', itemId: 1728 });
  });

  it('carries the band set as a code the simulator reads: class, race, level, talents and every pick', () => {
    const build = decodedFrom(bisSimHref(band({}), index, '1.60.1.70755'));
    expect(build.dataBuild).toBe('1.60.1.70755');
    expect(build.classSlug).toBe('warrior');
    expect(build.raceSlug).toBe('troll');
    expect(build.level).toBe(60);
    expect(build.treeRanks[1]).toEqual([5, 5]);
    expect(build.gear).toEqual({ head: 12640, main_hand: 1728, ranged: 12651 });
  });

  it("opens a leveling band at the band's own top level", () => {
    expect(decodedFrom(bisSimHref(band({ band: 30 }), index, 'b')).level).toBe(39);
  });

  it('leaves a slot with no pick out of the set', () => {
    const empty = { slot: 'neck', verified: false } as unknown as BisSlot;
    const build = decodedFrom(bisSimHref(band({ slots: [slot('head', 1), empty] }), index, 'b'));
    expect(build.gear).toEqual({ head: 1 });
  });

  it('keeps only the points the trees can reach, so the simulator never refuses the set', () => {
    // Flurry (Fury's last, tier-3 talent) with nothing above it: unreachable on these trees.
    const build = decodedFrom(bisSimHref(band({ talents: '0000000-5500001' }), index, 'b'));
    expect(build.treeRanks[1]).toEqual([5, 5]);
  });
});
