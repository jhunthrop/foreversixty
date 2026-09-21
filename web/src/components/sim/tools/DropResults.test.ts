// web/src/components/sim/tools/DropResults.test.ts
// D48's own newcomer-review repro: a boss whose Substitution.SourceName never rode along
// (a result saved before contract 10.1 A6's field existed) used to fall back to the raw
// `drop:`-stripped origin id verbatim -- "dungeon:blackrock-spire:175245" as a boss name.
// A static-render check (svelte/server, no jsdom), the same pattern
// SubstitutionChips.test.ts uses for a component whose whole surface is a few props.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { decodeFS1 } from '../../../lib/planner/fs1';
import type { Item } from '../../../lib/planner/types';
import type { BulkResult, Combo } from '../../../lib/sim/bulk-types';
import { handoffCopy } from '../../../lib/sim/handoff-copy';
import DropResults from './DropResults.svelte';

function comboFor(origin: string, sourceName?: string): Combo {
  return {
    substitutions: [
      {
        kind: 'item',
        slot: 'head',
        item_id: 16963,
        origin,
        name: 'Helm of Wrath',
        source_name: sourceName,
      },
    ],
    dps: { mean: 1500, stddev: 0, error: 0, min: 0, max: 0 },
    delta: { mean: 30, stddev: 0, error: 0, min: 0, max: 0 },
    group: 0,
  };
}

/**
 * Task 7: every "Every upgrade" row now carries a Plan it link, built from
 * `result.request.character` -- so the fixture result needs a real (if minimal) character
 * on it, where before it needed none.
 */
function resultWith(combo: Combo): BulkResult {
  return {
    request: {
      character: {
        name: 'Fury',
        race: 'orc',
        class: 'warrior',
        level: 60,
        talents: '0-5530515-',
        gear: [{ slot: 'main_hand', item_id: 12784 }],
        buffs: [],
        consumes: [],
      },
    },
    combos: [combo],
    equipped: { mean: 1400, stddev: 0, error: 12, min: 0, max: 0 },
  } as unknown as BulkResult;
}

function renderDrops(result: BulkResult): string {
  const items: ReadonlyMap<number, Item> = new Map();
  const { body } = render(DropResults, {
    props: { result, items, treeVersion: 'test', untried: [], onpin: () => {} },
  });
  return body;
}

describe('DropResults', () => {
  it('names a boss from its source name when one rode along', () => {
    const body = renderDrops(resultWith(comboFor('drop:raid:molten-core:11502', 'Ragnaros')));
    expect(body).toContain('Ragnaros');
  });

  it('humanises the origin id, never showing it raw, for a result with no source name', () => {
    const body = renderDrops(resultWith(comboFor('drop:dungeon:blackrock-spire:175245')));
    expect(body).not.toContain('dungeon:blackrock-spire:175245');
    expect(body).not.toContain('drop:');
    expect(body).toContain('Dungeon Blackrock spire 175245');
  });

  it('renders a Plan it link on every upgrade row, targeting the planner with that row’s own item', () => {
    const body = renderDrops(resultWith(comboFor('drop:raid:molten-core:11502', 'Ragnaros')));
    expect(body).toContain('/planner?code=');
    expect(body).toContain(handoffCopy.planIt);
    const href = /href="(\/planner\?code=[^"]+)"/.exec(body)?.[1];
    expect(href).toBeDefined();
    const code = decodeURIComponent(href!.replace('/planner?code=', ''));
    const decoded = decodeFS1(code);
    if (!decoded.ok) throw new Error(decoded.message);
    expect(decoded.build.gearSlots.find((slot) => slot.slot === 'head')?.itemId).toBe(16963);
  });
});
