// web/src/lib/sim/combos.test.ts
import { describe, expect, it } from 'vitest';
import bulkResultJson from '../../fixtures/sim/bulk-result.json';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import setsJson from '../../fixtures/planner/sets.json';
import { addonStringFor } from './addon-export';
import { ranksFromTalentsString } from './character';
import {
  MINUS,
  anyKeepsSetBonus,
  canPlanCombo,
  collapsedComboCount,
  comboKey,
  comboRows,
  deltaLabel,
  exactTieGroups,
  gainLabel,
  gearForCombo,
  headlineFor,
  isEmptiedOffHand,
  keepsSetBonus,
  percentOf,
  planItHref,
  signedGainLabel,
  slotSummary,
  sourceNameOfCombo,
  substitutionChipLabel,
  substitutionLabel,
  winningGear,
  type ComboRow,
} from './combos';
import { bulkCopy } from './copy';
import { decodeFS1 } from '../planner/fs1';
import type { BulkResult, Combo } from './bulk-types';
import type { GearSlot } from './types';
import type { Item, ItemSet } from '../planner/types';

const result = bulkResultJson as unknown as BulkResult;
const items = new Map((itemsJson.items as unknown as Item[]).map((item) => [item.id, item]));
const sets = setsJson as unknown as ItemSet[];

describe('comboRows', () => {
  it('ranks best first and gives every member of a within-error group the same rank', () => {
    const rows = comboRows(result);
    // The fixture's 7 combos carry groups 0,0,1,1,2,2,2 (task-2 fix round 1 added a 7th
    // combo to the trailing group-2 tie). Rank is the 1-based index of a group's first
    // member, so the group-2 trio all share rank 5.
    expect(rows.map((row) => row.rank)).toEqual([1, 1, 3, 3, 5, 5, 5]);
    expect(rows[0].withinError).toBe(true);
    expect(rows[2].withinError).toBe(false);
  });

  it('carries the percent against the equipped set', () => {
    const rows = comboRows(result);
    expect(rows[0].percent).toBeCloseTo((41.2 / 1461.2) * 100, 6);
  });
});

/**
 * newcomer MAJOR (review.md:251-257): rings and trinkets are tried in both slots, so the
 * engine can emit the same combination twice, differing only by which finger/trinket slot
 * carries it. A row's identity is its substitution SET -- an `item` substitution keyed by
 * `item_id`/`enchant`/`suffix` and NOT `slot` -- sorted so the engine's emission order
 * cannot matter.
 */
describe('comboRows de-duplicates identical candidates at the source', () => {
  const finger1Band: Combo = {
    substitutions: [{ kind: 'item', slot: 'finger1', item_id: 19325, name: 'Band of Accuria' }],
    dps: result.equipped,
    delta: { mean: 10, stddev: 0, error: 1, min: 0, max: 0 },
    group: 0,
  };
  const finger2Band: Combo = {
    ...finger1Band,
    substitutions: [{ kind: 'item', slot: 'finger2', item_id: 19325, name: 'Band of Accuria' }],
  };

  it('collapses two combos identical but for slot (finger1 vs finger2) into one row', () => {
    const rows = comboRows({ ...result, combos: [finger1Band, finger2Band] });
    expect(rows).toHaveLength(1);
    expect(rows[0].combo).toBe(finger1Band);
  });

  it('keeps two combos with genuinely different item ids distinct', () => {
    const otherRing: Combo = {
      ...finger2Band,
      substitutions: [{ kind: 'item', slot: 'finger2', item_id: 19326, name: 'Some Other Ring' }],
    };
    const rows = comboRows({ ...result, combos: [finger1Band, otherRing] });
    expect(rows).toHaveLength(2);
  });

  it('collapses a two-substitution combo emitted in the two possible orders', () => {
    const orderA: Combo = {
      substitutions: [
        { kind: 'item', slot: 'head', item_id: 1, name: 'A' },
        { kind: 'item', slot: 'shoulder', item_id: 2, name: 'B' },
      ],
      dps: result.equipped,
      delta: { mean: 8, stddev: 0, error: 1, min: 0, max: 0 },
      group: 0,
    };
    const orderB: Combo = { ...orderA, substitutions: [orderA.substitutions[1], orderA.substitutions[0]] };
    const rows = comboRows({ ...result, combos: [orderA, orderB] });
    expect(rows).toHaveLength(1);
    expect(rows[0].combo).toBe(orderA);
  });

  it('renumbers ranks 1, 2, 3 with no hole after a collapse', () => {
    const nextGroup: Combo = {
      substitutions: [{ kind: 'item', slot: 'trinket1', item_id: 3, name: 'C' }],
      dps: result.equipped,
      delta: { mean: 5, stddev: 0, error: 1, min: 0, max: 0 },
      group: 1,
    };
    const lastGroup: Combo = {
      substitutions: [{ kind: 'item', slot: 'trinket2', item_id: 4, name: 'D' }],
      dps: result.equipped,
      delta: { mean: 2, stddev: 0, error: 1, min: 0, max: 0 },
      group: 2,
    };
    const rows = comboRows({ ...result, combos: [finger1Band, finger2Band, nextGroup, lastGroup] });
    expect(rows.map((row) => row.rank)).toEqual([1, 2, 3]);
  });

  it('keeps the leader’s rank when a within-error group loses a duplicate member', () => {
    const thirdMember: Combo = {
      substitutions: [{ kind: 'item', slot: 'trinket1', item_id: 5, name: 'E' }],
      dps: result.equipped,
      delta: { mean: 9.5, stddev: 0, error: 1, min: 0, max: 0 },
      group: 0,
    };
    const nextGroup: Combo = {
      substitutions: [{ kind: 'item', slot: 'trinket2', item_id: 6, name: 'F' }],
      dps: result.equipped,
      delta: { mean: 4, stddev: 0, error: 1, min: 0, max: 0 },
      group: 1,
    };
    const rows = comboRows({ ...result, combos: [finger1Band, finger2Band, thirdMember, nextGroup] });
    expect(rows.map((row) => row.rank)).toEqual([1, 1, 3]);
  });
});

describe('percentOf and deltaLabel', () => {
  it('reads a gain with its error and a sign, whole once the figure reaches 10', () => {
    expect(deltaLabel({ mean: 41.2, stddev: 0, error: 5.41, min: 0, max: 0 })).toBe('+41 ± 11');
    // Below 10, the magnitude keeps its decimal (dps D38: two builds 0.44 DPS apart must not
    // both read "+0"). 1.8 rounds to 1.8, which is still under 10, so it stays "1.8".
    expect(deltaLabel({ mean: -1.8, stddev: 0, error: 5.38, min: 0, max: 0 })).toBe('−1.8 ± 11');
  });

  it('keeps a decimal under 10 DPS so two builds 0.44 DPS apart do not both read "+0" (dps D38)', () => {
    expect(deltaLabel({ mean: 0.44, stddev: 0, error: 0, min: 0, max: 0 })).toBe('+0.4 ± 0');
    expect(deltaLabel({ mean: -0.44, stddev: 0, error: 0, min: 0, max: 0 })).toBe('−0.4 ± 0');
  });

  it('never rounds a non-zero band away to "0" -- the Droptimizer repro (tank-sim review)', () => {
    // 1.96 * 2 = 3.92, one decimal (below 10) via the same formatMargin the headline uses.
    // The mean goes through gainLabel too (dps D38), which keeps a decimal below 10 even
    // for a whole number like -3 -- the same "-3.0" the adjacent "keeps a decimal" test
    // documents, not a special case for this test's own round number.
    expect(deltaLabel({ mean: -3, stddev: 0, error: 2, min: 0, max: 0 })).toBe('−3.0 ± 3.9');
  });

  it('is zero percent against a zero baseline rather than infinite', () => {
    expect(percentOf(41.2, 0)).toBe(0);
  });
});

describe('gainLabel', () => {
  it('keeps one decimal under a magnitude of 10', () => {
    expect(gainLabel(0.44)).toBe('0.4');
  });

  it('renders a true zero as whole, not "0.0" — a real zero should not imply precision', () => {
    expect(gainLabel(0)).toBe('0');
  });

  it('rounds 9.95 past the 10 boundary first, then renders whole because the rounded value is not under 10', () => {
    expect(gainLabel(9.95)).toBe('10');
  });

  it('keeps 10.0 whole', () => {
    expect(gainLabel(10.0)).toBe('10');
  });

  it('drops the decimal and thousands-separates once the magnitude reaches 10', () => {
    expect(gainLabel(41.2)).toBe('41');
    expect(gainLabel(1234.2)).toBe('1,234');
  });

  /**
   * Final whole-branch review, Important 1: `gainLabel` documented a non-negative
   * precondition but never enforced it, and the BY SLOT column (ComboResults.svelte) called
   * it with a signed `SlotSummaryRow.gain` -- `gainLabel(-1234.2)` used to render "-1234.2",
   * a spurious decimal with the thousands separator lost, where the old code rendered
   * "-1,234". `gainLabel` now takes its own magnitude defensively, so a negative input can
   * never reintroduce that regression even if a future caller forgets `Math.abs` too.
   */
  it('takes the magnitude defensively -- a negative input formats the same as its absolute value', () => {
    expect(gainLabel(-1234.2)).toBe(gainLabel(1234.2));
    expect(gainLabel(-0.44)).toBe(gainLabel(0.44));
    expect(gainLabel(-1234.2)).toBe('1,234');
  });
});

/**
 * Final whole-branch review, Important 1: the BY SLOT gain cell's own sign decision, pulled
 * out of ComboResults.svelte's markup into a pure, tested function. `SlotSummaryRow.gain` is
 * `combo.delta.mean`, unconstrained in sign (a persona reviewer saw a -2 drop), unlike
 * `deltaLabel`'s `Estimate.mean`, which already went through `Math.abs` via `gainLabel`
 * before this fix. `MINUS`, not a hyphen -- the design system's rule for a negative figure.
 */
describe('signedGainLabel', () => {
  it('signs a negative gain at or above 10 with MINUS, whole and thousands-separated', () => {
    expect(signedGainLabel(-1234.2)).toBe(`${MINUS}1,234`);
  });

  it('signs a negative gain under 10 with MINUS and keeps its decimal', () => {
    expect(signedGainLabel(-1.8)).toBe(`${MINUS}1.8`);
  });

  it('signs a positive gain with a plus', () => {
    expect(signedGainLabel(41.2)).toBe('+41');
  });

  it('reads an unknown gain as an em dash', () => {
    expect(signedGainLabel(null)).toBe('—');
  });
});

/**
 * Final whole-branch review, Important 2: `store.combinations` is the engine's own
 * `simCount`, before `comboRows`' de-dupe (design 3.2's rings-and-trinkets-in-both-slots
 * rule) collapses duplicate placements into one row -- so the run bar's count can read
 * higher than the table's own row count with nothing on the page explaining the gap.
 * Controller ruling: explain the gap rather than recompute either number from the other.
 */
describe('collapsedComboCount', () => {
  it('is zero when nothing was collapsed', () => {
    expect(collapsedComboCount(result)).toBe(0);
  });

  it('counts exactly how many raw combos a finger1/finger2 duplicate collapsed into one row', () => {
    const finger1Band: Combo = {
      substitutions: [{ kind: 'item', slot: 'finger1', item_id: 19325, name: 'Band of Accuria' }],
      dps: result.equipped,
      delta: { mean: 10, stddev: 0, error: 1, min: 0, max: 0 },
      group: 0,
    };
    const finger2Band: Combo = {
      ...finger1Band,
      substitutions: [{ kind: 'item', slot: 'finger2', item_id: 19325, name: 'Band of Accuria' }],
    };
    expect(collapsedComboCount({ ...result, combos: [finger1Band, finger2Band] })).toBe(1);
  });
});

/**
 * Engine-lane rule 5's two-hander: the winner replaces main-hand and off-hand with one
 * weapon, which the engine reports as the new main hand PLUS
 * `{kind:"item", slot:"off_hand", item_id:0}`. The shared fixture has no off-hand, so the
 * case is built here from it rather than changed in a file eight other tests read.
 */
const twoHanded: BulkResult = {
  ...result,
  request: {
    ...result.request,
    character: {
      ...result.request.character,
      gear: [...result.request.character.gear, { slot: 'off_hand', item_id: 11684 }],
    },
  },
  combos: [
    {
      ...result.combos[0],
      substitutions: [
        { kind: 'item', slot: 'main_hand', item_id: 17182, name: 'Sulfuras' },
        { kind: 'item', slot: 'off_hand', item_id: 0, name: '<item removed>' },
      ],
    },
    ...result.combos.slice(1),
  ],
};

describe('winningGear', () => {
  it('writes the leader’s substitutions over the base character’s gear', () => {
    const gear = winningGear(result);
    expect(gear.find((slot) => slot.slot === 'head')?.item_id).toBe(16963);
    expect(gear.find((slot) => slot.slot === 'shoulder')?.item_id).toBe(16966);
    // untouched by the winner
    expect(gear.find((slot) => slot.slot === 'main_hand')?.item_id).toBe(12784);
  });

  it('leaves the base character’s own gear untouched', () => {
    const before = JSON.stringify(result.request.character.gear);
    winningGear(result);
    expect(JSON.stringify(result.request.character.gear)).toBe(before);
  });

  it('drops the slot a two-hander emptied rather than writing item 0 over it', () => {
    const gear = winningGear(twoHanded);
    expect(gear.find((slot) => slot.slot === 'main_hand')?.item_id).toBe(17182);
    // Not `{item_id: 0}`: nothing downstream defines 0 as "empty", so the slot is gone.
    expect(gear.some((slot) => slot.slot === 'off_hand')).toBe(false);
  });

  it('exports an addon string with no off-hand entry at all', () => {
    const addon = addonStringFor({
      dataBuild: '1.60.1',
      classSlug: 'warrior',
      raceSlug: 'orc',
      talents: '0-5530515-0',
      gear: winningGear(twoHanded),
    });
    expect(addon).toContain('main_hand=17182');
    expect(addon).not.toContain('off_hand');
    expect(addon).not.toContain('=0');
  });
});

describe('gearForCombo', () => {
  it('writes one substitution over the base gear, same as winningGear does for the leader', () => {
    const base: GearSlot[] = [{ slot: 'head', item_id: 1 }];
    const combo: Combo = {
      substitutions: [{ kind: 'item', slot: 'head', item_id: 2 }],
      dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      group: 0,
    };
    expect(gearForCombo(base, combo)).toEqual([{ slot: 'head', item_id: 2 }]);
  });

  it('agrees with winningGear on the leader combo', () => {
    expect(gearForCombo(result.request.character.gear, result.combos[0])).toEqual(winningGear(result));
  });

  it('removes an emptied off-hand rather than writing item_id 0, same rule as winningGear', () => {
    const base: GearSlot[] = [
      { slot: 'main_hand', item_id: 1 },
      { slot: 'off_hand', item_id: 2 },
    ];
    const combo: Combo = {
      substitutions: [
        { kind: 'item', slot: 'main_hand', item_id: 3 },
        { kind: 'item', slot: 'off_hand', item_id: 0 },
      ],
      dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      group: 0,
    };
    expect(gearForCombo(base, combo)).toEqual([{ slot: 'main_hand', item_id: 3 }]);
  });

  it('ignores a talents, set or consumes substitution, leaving the base gear untouched', () => {
    const base: GearSlot[] = [{ slot: 'head', item_id: 1 }];
    const combo: Combo = {
      substitutions: [
        { kind: 'talents', name: 'Deep Fury', talents: '0-5530515-' },
        { kind: 'set', name: 'Battlegear of Wrath' },
        { kind: 'consumes', name: 'flask_of_supreme_power' },
      ],
      dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      group: 0,
    };
    expect(gearForCombo(base, combo)).toEqual(base);
  });
});

/**
 * The fixture's own three non-item combos (task 7 fix round 1 added the talents and
 * set+consumes rows) are reused directly rather than re-built here: `comboRows(result)`'s
 * groups already carry one of each kind this predicate has to tell apart.
 */
describe('canPlanCombo', () => {
  const itemCombo = result.combos[0];
  const talentsCombo = result.combos.find((combo) =>
    combo.substitutions.every((sub) => sub.kind === 'talents'),
  )!;
  const setAndConsumesCombo = result.combos.find((combo) =>
    combo.substitutions.some((sub) => sub.kind === 'set'),
  )!;

  it('is true for a combo with an item substitution', () => {
    expect(canPlanCombo(itemCombo)).toBe(true);
  });

  it('is true for a talents-only combo -- its own talents string can ride onto the spec', () => {
    expect(canPlanCombo(talentsCombo)).toBe(true);
  });

  it('is false for a set (plus consumes) combo -- a named set carries no gear on the wire', () => {
    expect(canPlanCombo(setAndConsumesCombo)).toBe(false);
  });

  it('is false for a consumes-only combo -- consumables change neither gear nor talents', () => {
    const consumesOnly: Combo = {
      substitutions: [{ kind: 'consumes', name: 'flask_of_supreme_power' }],
      dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      group: 0,
    };
    expect(canPlanCombo(consumesOnly)).toBe(false);
  });

  it('is false for a set substitution even mixed with an item substitution -- the set’s own gear is still unknown', () => {
    const setPlusItem: Combo = {
      substitutions: [
        { kind: 'set', name: 'Battlegear of Wrath' },
        { kind: 'item', slot: 'trinket1', item_id: 13968 },
      ],
      dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
      group: 0,
    };
    expect(canPlanCombo(setPlusItem)).toBe(false);
  });
});

describe('planItHref', () => {
  const itemCombo = result.combos[0];
  const talentsCombo = result.combos.find((combo) =>
    combo.substitutions.every((sub) => sub.kind === 'talents'),
  )!;
  const setAndConsumesCombo = result.combos.find((combo) =>
    combo.substitutions.some((sub) => sub.kind === 'set'),
  )!;

  function decodedGearOf(href: string): { slot: string; itemId: number }[] {
    const code = decodeURIComponent(href.replace(/^\/planner\?code=/, ''));
    const decoded = decodeFS1(code);
    if (!decoded.ok) throw new Error(decoded.message);
    return decoded.build.gearSlots;
  }

  it('encodes an item row’s own substituted item into the planner link', () => {
    const href = planItHref(result, itemCombo, 'test-build');
    expect(href).not.toBeNull();
    const gear = decodedGearOf(href!);
    expect(gear.find((slot) => slot.slot === 'head')?.itemId).toBe(16963);
    expect(gear.find((slot) => slot.slot === 'shoulder')?.itemId).toBe(16966);
  });

  // FS1's own tree encoding writes a fully empty tree as a single "0" digit (`encodeTree`'s
  // trailing-zero trim falls back to "0" rather than the empty string), so a wholly-unspent
  // tree round-trips as `[0]`, not `[]`. Trailing zeros are trimmed on both sides before
  // comparing -- the same normalisation `encodeTree` itself already applies -- so this test
  // asserts what the talents STRING actually said, not an FS1 encoding artifact.
  function trimTrailingZeros(tree: readonly number[]): number[] {
    const trimmed = [...tree];
    while (trimmed.length > 0 && trimmed[trimmed.length - 1] === 0) trimmed.pop();
    return trimmed;
  }

  it('carries a talents-only row’s own talents string, with the base gear untouched', () => {
    const href = planItHref(result, talentsCombo, 'test-build');
    expect(href).not.toBeNull();
    const code = decodeURIComponent(href!.replace(/^\/planner\?code=/, ''));
    const decoded = decodeFS1(code);
    if (!decoded.ok) throw new Error(decoded.message);
    expect(decoded.build.treeRanks.map(trimTrailingZeros)).toEqual(
      ranksFromTalentsString(talentsCombo.substitutions[0].talents!).map(trimTrailingZeros),
    );
    expect(decoded.build.gearSlots.map((slot) => slot.itemId).sort()).toEqual(
      result.request.character.gear.map((slot) => slot.item_id).sort(),
    );
  });

  it('is null for a set (plus consumes) row -- no link opens gear the row never carried', () => {
    expect(planItHref(result, setAndConsumesCombo, 'test-build')).toBeNull();
  });
});

/**
 * Review fix round 1: the original `comboKey` (moved here unchanged from
 * `DropResults.svelte`) keyed on only the FIRST substitution, so the fixture's own
 * `combos[0]` (head+shoulder) and `combos[1]` (head alone) both rendered "head:16963" once
 * `comboKey` reached a component (`ComboResults.svelte`) whose rows can carry more than one
 * substitution -- harmless in `DropResults.svelte`, whose "drops" mode combos are always
 * single-substitution, but a real `data-testid` collision (and an e2e selector risk) once
 * shared. `web/tests/e2e/sim-drops.spec.ts` asserts `sim-drops-pin-finger1:19325`,
 * `sim-drops-pin-head:16963` and `sim-drops-pin-main_hand:12784` verbatim, so the plain
 * single-item case (no enchant, no suffix) must keep producing exactly `<slot>:<item_id>`.
 */
describe('comboKey', () => {
  it('keys a plain single-item row on slot:item_id, unchanged -- e2e selects these ids verbatim', () => {
    const headOnly = comboRows(result).find(
      (row) => row.combo.substitutions.length === 1 && row.combo.substitutions[0].slot === 'head',
    )!;
    expect(comboKey(headOnly)).toBe('head:16963');
  });

  it('gives combos[0] (head+shoulder) and combos[1] (head alone) different keys', () => {
    const rows = comboRows(result);
    const headAndShoulder = rows.find((row) => row.combo.substitutions.length === 2)!;
    const headOnly = rows.find(
      (row) => row.combo.substitutions.length === 1 && row.combo.substitutions[0].slot === 'head',
    )!;
    expect(comboKey(headAndShoulder)).not.toBe(comboKey(headOnly));
  });

  it('distinguishes two rows carrying the same item at different enchants', () => {
    const plain: ComboRow = {
      rank: 1,
      withinError: true,
      percent: 0,
      combo: {
        substitutions: [{ kind: 'item', slot: 'head', item_id: 16963 }],
        dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
        delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
        group: 0,
      },
    };
    const enchanted: ComboRow = {
      ...plain,
      combo: {
        ...plain.combo,
        substitutions: [{ kind: 'item', slot: 'head', item_id: 16963, enchant: 907 }],
      },
    };
    const differentlyEnchanted: ComboRow = {
      ...plain,
      combo: {
        ...plain.combo,
        substitutions: [{ kind: 'item', slot: 'head', item_id: 16963, enchant: 123 }],
      },
    };
    expect(comboKey(plain)).toBe('head:16963');
    expect(comboKey(enchanted)).not.toBe(comboKey(plain));
    expect(comboKey(enchanted)).not.toBe(comboKey(differentlyEnchanted));
  });

  it('distinguishes two talents-only rows with different loadouts', () => {
    const deepFury: ComboRow = {
      rank: 5,
      withinError: false,
      percent: 0,
      combo: {
        substitutions: [{ kind: 'talents', name: 'Deep Fury', talents: '0-5530515-' }],
        dps: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
        delta: { mean: 0, stddev: 0, error: 0, min: 0, max: 0 },
        group: 2,
      },
    };
    const arms: ComboRow = {
      ...deepFury,
      combo: {
        ...deepFury.combo,
        substitutions: [{ kind: 'talents', name: 'Deep Wounds', talents: '5530515-0-' }],
      },
    };
    expect(comboKey(deepFury)).not.toBe(comboKey(arms));
  });

  it('gives every row of the fixture result a unique key', () => {
    const rows = comboRows(result);
    const keys = new Set(rows.map((row) => comboKey(row)));
    expect(keys.size).toBe(rows.length);
  });
});

describe('isEmptiedOffHand', () => {
  it('is the rule-5 sentinel and nothing else', () => {
    expect(isEmptiedOffHand({ kind: 'item', slot: 'off_hand', item_id: 0 })).toBe(true);
    expect(isEmptiedOffHand({ kind: 'item', slot: 'off_hand', item_id: 11684 })).toBe(false);
    expect(isEmptiedOffHand({ kind: 'item', slot: 'main_hand', item_id: 0 })).toBe(false);
    expect(isEmptiedOffHand({ kind: 'talents', name: 'Deep Fury' })).toBe(false);
  });
});

describe('slotSummary', () => {
  it('names the winner’s item per slot and the gain that slot contributed alone', () => {
    const rows = slotSummary(result);
    const head = rows.find((row) => row.slot === 'head')!;
    expect(head.item_id).toBe(16963);
    expect(head.name).toBe('Helm of Wrath');
    // the single-substitution combo for that slot
    expect(head.gain).toBeCloseTo(38.6, 6);
    const shoulder = rows.find((row) => row.slot === 'shoulder')!;
    expect(shoulder.gain).toBeCloseTo(18.9, 6);
  });

  it('lists a slot the winner changed even when nothing measured it alone', () => {
    const trimmed: BulkResult = { ...result, combos: [result.combos[0]] };
    const rows = slotSummary(trimmed);
    expect(rows.map((row) => row.slot).sort()).toEqual(['head', 'shoulder']);
    expect(rows.every((row) => row.gain === null || typeof row.gain === 'number')).toBe(true);
  });

  // Engine-lane rule 5, the second leak (Task 16's re-review, routed to this task): a
  // two-hander replacing a main-plus-off-hand pair emits a second substitution
  // `{kind:"item", slot:"off_hand", item_id:0, name:"<item removed>"}`. SubstitutionChips
  // already renders this as an emptied slot; slotSummary fed the same substitution's raw
  // `name` straight through substitutionLabel, which prints "<item removed>" verbatim since
  // the field is non-empty. This is the "By slot" panel's own path onto the same sentinel.
  it('renders the emptied off-hand from a two-hander swap as an emptied slot, never the literal engine text', () => {
    const twoHander: BulkResult = {
      ...result,
      combos: [
        {
          substitutions: [
            { kind: 'item', slot: 'main_hand', item_id: 19351, name: 'Sulfuras, Hand of Ragnaros' },
            { kind: 'item', slot: 'off_hand', item_id: 0, name: '<item removed>' },
          ],
          dps: result.equipped,
          delta: result.equipped,
          group: 0,
        },
      ],
    };
    const rows = slotSummary(twoHander);
    const offHand = rows.find((row) => row.slot === 'off_hand')!;
    expect(offHand.name).toBe(bulkCopy.offHandEmptied);
    expect(offHand.name).not.toBe('<item removed>');
    const mainHand = rows.find((row) => row.slot === 'main_hand')!;
    expect(mainHand.name).toBe('Sulfuras, Hand of Ragnaros');
  });
});

describe('keepsSetBonus', () => {
  it('is false when the combination does not reach the piece count', () => {
    expect(keepsSetBonus(result.combos[0], result, items, sets, 4)).toBe(false);
  });

  it('is true at a piece count the combination does reach', () => {
    expect(keepsSetBonus(result.combos[0], result, items, sets, 2)).toBe(true);
  });
});

/**
 * dps D24: ComboResults.svelte used to print "Only combinations keeping a 4-piece set
 * bonus" over every result, including a character with no 4-piece set at all -- reading
 * like a filter silently throwing candidates away. `anyKeepsSetBonus` is what the page now
 * checks before showing the checkbox at all: whether it is even possible.
 */
describe('anyKeepsSetBonus', () => {
  it('is false at a piece count nothing in the result reaches (no 4-piece set exists here)', () => {
    expect(anyKeepsSetBonus(result, items, sets, 4)).toBe(false);
  });

  it('is true at a piece count some combo does reach', () => {
    expect(anyKeepsSetBonus(result, items, sets, 2)).toBe(true);
  });

  it('is false when the result has no combos at all', () => {
    expect(anyKeepsSetBonus({ ...result, combos: [] }, items, sets, 2)).toBe(false);
  });
});

/**
 * dps D36: three or four ranked rows for genuinely different items (three different necks,
 * none of them carrying a stat this spec's damage depends on) showed the identical delta to
 * the decimal, reading like a bug silently reporting one candidate under several rows.
 * `dedupedCombos` (above) only ever drops an EXACT repeat of the same item id, so distinct
 * items always keep distinct rows -- these are real ties, and `exactTieGroups` is what lets
 * the page say so instead of leaving the duplicate numbers unexplained.
 */
describe('exactTieGroups', () => {
  const estimate = { mean: -19, stddev: 0, error: 2.3, min: -30, max: -10 };
  const otherEstimate = { mean: -41, stddev: 0, error: 11, min: -60, max: -20 };

  function row(itemId: number, delta = estimate): ComboRow {
    return {
      rank: 1,
      combo: {
        substitutions: [{ kind: 'item', slot: 'neck', item_id: itemId, name: `Item ${itemId}` }],
        dps: delta,
        delta,
        group: 0,
      },
      withinError: true,
      percent: 0,
    };
  }

  it('groups rows whose delta is bit-identical, three different items included', () => {
    const rows = [row(1), row(2), row(3, otherEstimate)];
    expect(exactTieGroups(rows)).toEqual([[rows[0], rows[1]]]);
  });

  it('finds nothing when every row is its own number', () => {
    expect(exactTieGroups([row(1), row(2, otherEstimate)])).toEqual([]);
  });

  it('finds nothing in a list of one', () => {
    expect(exactTieGroups([row(1)])).toEqual([]);
  });
});

describe('headlineFor and substitutionLabel', () => {
  it('names the leader’s biggest change from the result’s own name field', () => {
    expect(headlineFor(result)).toBe('+41 DPS from Helm of Wrath');
  });

  // headlineFor splits deltaLabel(...) on a space and takes [0]; a sub-10 gain's decimal
  // must survive that split rather than being cut at the space inside "0.4".
  it('keeps a small gain’s decimal when splitting deltaLabel’s "+0.4 ± 0" on the space', () => {
    const smallGain: BulkResult = {
      ...result,
      combos: [{ ...result.combos[0], delta: { mean: 0.44, stddev: 0, error: 0, min: 0, max: 0 } }],
    };
    expect(headlineFor(smallGain)).toBe('+0.4 DPS from Helm of Wrath');
  });

  it('labels an item, a loadout, a set and a consumable list (contract 10.8)', () => {
    expect(substitutionLabel({ kind: 'item', slot: 'head', item_id: 16963, name: 'Helm of Wrath' })).toBe(
      'Helm of Wrath',
    );
    expect(substitutionLabel({ kind: 'talents', name: 'Deep Fury' })).toBe('Deep Fury');
    expect(substitutionLabel({ kind: 'set', name: 'AQ set' })).toBe('AQ set');
    expect(substitutionLabel({ kind: 'consumes', name: 'flask_of_supreme_power' })).toBe(
      'With flask of supreme power',
    );
  });

  it('falls back to the item id only when the engine sent no name', () => {
    expect(substitutionLabel({ kind: 'item', slot: 'head', item_id: 16963 })).toBe('Item 16963');
  });
});

describe('sourceNameOfCombo', () => {
  it('reads the boss off the substitution the engine copied it onto (contract 10.1 A6)', () => {
    expect(
      sourceNameOfCombo({
        substitutions: [{ kind: 'item', item_id: 1, source_name: 'Ragnaros' }],
        dps: result.equipped,
        delta: result.equipped,
        group: 0,
      }),
    ).toBe('Ragnaros');
  });
});

describe('substitutionChipLabel', () => {
  // Task 5 (newcomer MAJOR, review.md:360-363): SubstitutionChips.svelte used to carry this
  // exact string only in a `title`, which a phone can never hover to read. Folding the
  // source into the same string the chip already renders is what makes it tappable by
  // construction; this is the pure decision behind that, kept out of the component per the
  // lane's testable-decision rule.
  it('appends the source name when the substitution carries one', () => {
    expect(
      substitutionChipLabel({
        kind: 'item',
        slot: 'head',
        item_id: 16963,
        name: 'Helm of Wrath',
        source_name: 'Ragnaros',
      }),
    ).toBe('Helm of Wrath · Ragnaros');
  });

  it('is just the label when the substitution carries no source', () => {
    expect(substitutionChipLabel({ kind: 'talents', name: 'Deep Fury' })).toBe('Deep Fury');
  });

  it('ignores an empty source name the same way the title it replaces did', () => {
    expect(
      substitutionChipLabel({ kind: 'item', slot: 'head', item_id: 1, name: 'X', source_name: '' }),
    ).toBe('X');
  });
});
