// web/src/components/sim/tools/ComboResults.test.ts
// Task 7: one "Plan it" link per plannable row, each opening the planner with exactly that
// row's own substitution -- an item written into its slot, or a talents-only row's own
// talents string -- and no link at all on a row `canPlanCombo` (combos.ts) says has nothing
// a planner link can honestly open (a named set, alone or mixed with consumables). A static
// render (`svelte/server`), the same pattern `StatWeights.test.ts` and
// `DropResults.test.ts` already use; `ComboResults.svelte` had no test file before this
// task.
//
// The fixture (`bulk-result.json`, shared with `combos.test.ts`) carries all three of the
// shapes this component has to tell apart: a two-item combo, three single-item combos, a
// talents-only combo, and a set+consumes combo with no item at all -- so this exercises the
// real component wiring rather than a hand-picked happy path.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import bulkResultJson from '../../../fixtures/sim/bulk-result.json';
import itemsJson from '../../../fixtures/planner/items/warrior.json';
import setsJson from '../../../fixtures/planner/sets.json';
import { decodeFS1 } from '../../../lib/planner/fs1';
import type { Item, ItemSet } from '../../../lib/planner/types';
import type { BulkResult } from '../../../lib/sim/bulk-types';
import { ranksFromTalentsString } from '../../../lib/sim/character';
import { canPlanCombo, comboRows } from '../../../lib/sim/combos';
import { handoffCopy } from '../../../lib/sim/handoff-copy';
import { poolQualityCopy } from '../../../lib/sim/pool-quality-copy';
import type { Combo } from '../../../lib/sim/bulk-types';
import ComboResults from './ComboResults.svelte';

const result = bulkResultJson as unknown as BulkResult;
const items: ReadonlyMap<number, Item> = new Map(
  (itemsJson.items as unknown as Item[]).map((item) => [item.id, item]),
);
const sets = setsJson as unknown as ItemSet[];

function renderCombos(overrides: Partial<BulkResult> = {}): string {
  const { body } = render(ComboResults, {
    props: {
      result: { ...result, ...overrides },
      items,
      sets,
      treeVersion: 'test-build',
      partial: false,
      onsave: async () => null,
    },
  });
  return body;
}

/**
 * Each row's own `<div role="row" data-testid="sim-combo-row">…</div>`, in the same order
 * `comboRows(result)` produces them. Bounded to that one `</div>` rather than "up to the
 * next row" (or the end of the body): every cell inside a row is a `<span>`, never a nested
 * `<div>` (SubstitutionChips.svelte included), so the first `</div>` after a row's own
 * opening tag is always that row's own closing tag -- which matters for the LAST row, since
 * splitting on "the next row start" would otherwise fold the page's own trailing "Open in
 * planner" button (a real `href`) into it and hide a false positive.
 */
function rowSegments(body: string): string[] {
  const marker = 'data-testid="sim-combo-row"';
  const segments: string[] = [];
  let cursor = 0;
  for (;;) {
    const markerIndex = body.indexOf(marker, cursor);
    if (markerIndex === -1) break;
    const openIndex = body.lastIndexOf('<div', markerIndex);
    const closeIndex = body.indexOf('</div>', markerIndex);
    segments.push(body.slice(openIndex, closeIndex + '</div>'.length));
    cursor = closeIndex + 1;
  }
  return segments;
}

function firstHref(segment: string): string | null {
  const match = /href="([^"]+)"/.exec(segment);
  return match === null ? null : match[1];
}

function decodedGearOf(href: string): { slot: string; itemId: number }[] {
  const code = decodeURIComponent(href.replace(/^\/planner\?code=/, ''));
  const decoded = decodeFS1(code);
  if (!decoded.ok) throw new Error(decoded.message);
  return decoded.build.gearSlots;
}

function trimTrailingZeros(tree: readonly number[]): number[] {
  const trimmed = [...tree];
  while (trimmed.length > 0 && trimmed[trimmed.length - 1] === 0) trimmed.pop();
  return trimmed;
}

describe('ComboResults: "Plan it" per row (task 7)', () => {
  it('renders one row per combo, each carrying its own Plan it decision', () => {
    const body = renderCombos();
    const rows = comboRows(result);
    const segments = rowSegments(body);
    expect(segments).toHaveLength(rows.length);
  });

  it('gives every item-substitution row a Plan it link whose gear contains that row’s own item(s)', () => {
    const body = renderCombos();
    const rows = comboRows(result);
    const segments = rowSegments(body);
    rows.forEach((row, index) => {
      const itemSubs = row.combo.substitutions.filter((sub) => sub.kind === 'item');
      if (itemSubs.length === 0) return;
      const segment = segments[index];
      expect(segment).toContain(handoffCopy.planIt);
      const href = firstHref(segment);
      expect(href).not.toBeNull();
      const gear = decodedGearOf(href!);
      for (const sub of itemSubs) {
        expect(gear.find((slot) => slot.slot === sub.slot)?.itemId).toBe(sub.item_id);
      }
    });
  });

  it('carries the talents-only row’s own talents string into its Plan it link', () => {
    const body = renderCombos();
    const rows = comboRows(result);
    const segments = rowSegments(body);
    const talentsIndex = rows.findIndex((row) =>
      row.combo.substitutions.every((sub) => sub.kind === 'talents'),
    );
    expect(talentsIndex).toBeGreaterThanOrEqual(0);
    const href = firstHref(segments[talentsIndex]);
    expect(href).not.toBeNull();
    const code = decodeURIComponent(href!.replace(/^\/planner\?code=/, ''));
    const decoded = decodeFS1(code);
    if (!decoded.ok) throw new Error(decoded.message);
    const talents = rows[talentsIndex].combo.substitutions[0].talents!;
    expect(decoded.build.treeRanks.map(trimTrailingZeros)).toEqual(
      ranksFromTalentsString(talents).map(trimTrailingZeros),
    );
  });

  it('renders no Plan it link for the set+consumes row -- canPlanCombo says it has nothing to open', () => {
    const body = renderCombos();
    const rows = comboRows(result);
    const segments = rowSegments(body);
    const setIndex = rows.findIndex((row) => row.combo.substitutions.some((sub) => sub.kind === 'set'));
    expect(setIndex).toBeGreaterThanOrEqual(0);
    expect(canPlanCombo(rows[setIndex].combo)).toBe(false);
    expect(segments[setIndex]).not.toContain('sim-combo-plan-it-');
    expect(firstHref(segments[setIndex])).toBeNull();
  });

  it('gives the action column header an accessible name with no visible text', () => {
    const body = renderCombos();
    expect(body).toContain(`aria-label="${handoffCopy.planIt}"`);
  });

  /**
   * Review fix round 1: `comboKey` used to key on only a row's first substitution, so the
   * fixture's own `combos[0]` (head+shoulder) and `combos[1]` (head alone) rendered the
   * identical `data-testid="sim-combo-plan-it-head:16963"` -- this asserts the render
   * itself, not just `comboKey` in isolation, since it is the rendered id an e2e spec
   * would actually select by.
   */
  it('gives every plannable row a unique Plan it test id', () => {
    const body = renderCombos();
    const rows = comboRows(result);
    const plannableCount = rows.filter((row) => canPlanCombo(row.combo)).length;
    const ids = [...body.matchAll(/data-testid="(sim-combo-plan-it-[^"]+)"/g)].map((match) => match[1]);
    expect(ids).toHaveLength(plannableCount);
    expect(new Set(ids).size).toBe(ids.length);
  });
});

/**
 * newcomer round 4 (review.md:83-110) / dps D25-pattern: "too close to separate" used to
 * print under every result with more than one row, whatever the actual gap. Two synthetic
 * results -- a clear gap and an overlapping pair -- prove the note is now conditioned on
 * `closestOverlappingPair`, not just row count. Fixed at the component boundary, which is
 * also what SavedCombos.svelte (the saved read-only view) shares.
 */
describe('ComboResults: "too close to separate" only when it is true (newcomer r4)', () => {
  function combo(itemId: number, slot: string, name: string, mean: number, error: number): Combo {
    const estimate = { mean, stddev: error * 10, error, min: mean - error, max: mean + error };
    return {
      substitutions: [{ kind: 'item', slot, item_id: itemId, name }],
      dps: estimate,
      delta: estimate,
      group: 0,
    };
  }

  it('shows no note over a clear gap, matching the round-4 repro numbers', () => {
    const combos = [
      combo(16963, 'head', 'Your current build', 0, 1.6),
      combo(16966, 'shoulder', 'Build 1', -239, 1.5),
    ];
    const body = renderCombos({ combos });
    expect(body).not.toContain('sim-within-error-note');
  });

  it('shows the note, naming both rows, when an adjacent pair overlaps', () => {
    const combos = [
      combo(16963, 'head', 'Your current build', 100, 5),
      combo(16966, 'shoulder', 'Build 1', 98, 5),
    ];
    const body = renderCombos({ combos });
    expect(body).toContain('data-testid="sim-within-error-note"');
    expect(body).toContain(poolQualityCopy.withinErrorNoteNaming('Your current build', 'Build 1'));
  });
});

/**
 * newcomer round 4: the same pasted build read 681 ± 6.9 in the editor's own preview and
 * 400 in this table -- both correct (gear differs by design in talents mode), never
 * explained. The note fires only for a talents-mode result.
 */
describe('ComboResults: talents-mode gear-locked note (newcomer r4)', () => {
  it('shows the note when the request is talents mode', () => {
    const body = renderCombos({
      request: { ...result.request, bulk: { ...result.request.bulk!, mode: 'talents' } },
    });
    expect(body).toContain(poolQualityCopy.talentsGearLockedNote);
  });

  it('does not show the note for gear mode (the fixture’s own default)', () => {
    const body = renderCombos();
    expect(body).not.toContain(poolQualityCopy.talentsGearLockedNote);
  });
});
