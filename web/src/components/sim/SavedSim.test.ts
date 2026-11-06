// web/src/components/sim/SavedSim.test.ts
// Task L8.1 fix round 1: the saved sim page's "Copy for the game". A static render
// (svelte/server, no jsdom), the setup SavedWeights.test.ts uses. The page's talent and item
// files only arrive after mount, so the render proves the button is there and disabled
// while the talent layout is missing; the code itself -- winning set for a bulk run, stored
// character for a run or weights result -- is `gameCodeForSimResult` (lib/sim/game-code.ts),
// the derivation the component calls, exercised here against the shared sim fixtures.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import bulkFixture from '../../fixtures/sim/bulk-result.json';
import fullFixture from '../../fixtures/sim/result-full.json';
import runFixture from '../../fixtures/sim/result.json';
import weightsFixture from '../../fixtures/sim/weights-result.json';
import itemsJson from '../../fixtures/planner/items/warrior.json';
import talentsJson from '../../fixtures/planner/talents/warrior.json';
import { addonCopy } from '../../lib/addon/copy';
import { decodeFSB1 } from '../../lib/addon/fsb1';
import { indexTalents } from '../../lib/planner/rules';
import type { Item, TalentFile } from '../../lib/planner/types';
import { gameCodeForSimResult } from '../../lib/sim/game-code';
import type { SimResult } from '../../lib/sim/types';
import SavedKind from './SavedKind.svelte';
import SavedSim from './SavedSim.svelte';

const runResult = runFixture as unknown as SimResult;
const fullResult = fullFixture as unknown as SimResult;
const bulkResult = bulkFixture as unknown as SimResult;
const weightsResult = weightsFixture as unknown as SimResult;
const items: ReadonlyMap<number, Item> = new Map(
  (itemsJson.items as unknown as Item[]).map((item) => [item.id, item]),
);
const talentIndex = indexTalents(talentsJson as unknown as TalentFile);

function decodedGear(code: string): { slot: string; itemId: number }[] {
  const decoded = decodeFSB1(code);
  if (!decoded.ok) throw new Error(decoded.message);
  return decoded.build.gear.map(({ slot, itemId }) => ({ slot, itemId }));
}

describe('SavedSim: which view a kind gets', () => {
  it('a run gets the one-page body: the table on the Plate, the spread and every open section', () => {
    const { body } = render(SavedSim, { props: { result: fullResult, onrerun: () => {} } });
    expect(body).toContain('sim-saved-run');
    expect(body).toContain('data-region="plate"');
    expect(body).toContain('sim-damage-card');
    expect(body).toContain('data-region="section-dps-spread"');
    for (const id of ['timeline', 'gear', 'talents', 'resource', 'buffs', 'rotation', 'settings']) {
      expect(body).toContain(`data-testid="sim-section-${id}"`);
    }
    for (const region of [
      'timeline',
      'gear',
      'talents',
      'rage',
      'buffs-and-debuffs',
      'this-run',
      'rotation',
    ]) {
      expect(body).toContain(`data-region="section-${region}"`);
    }
    expect(body).not.toContain('sim-section-nav');
  });

  it('an older run falls back without a broken section: no bins, no Saved', () => {
    const { body } = render(SavedSim, { props: { result: runResult, onrerun: () => {} } });
    expect(body).toContain('sim-spread-range');
    expect(body).toContain('This run kept the range, not the shape.');
    expect(body).not.toContain('sim-spread-plot');
    expect(body).not.toContain('sim-fact-saved');
  });

  it('a run with no sample and no gear keeps its sections standing', () => {
    const request = {
      ...fullResult.request,
      character: { ...fullResult.request.character, gear: [] },
    };
    const { body } = render(SavedSim, {
      props: { result: { ...fullResult, sample: undefined, request }, onrerun: () => {} },
    });
    expect(body).toContain('data-region="section-timeline"');
    expect(body).toContain('data-region="section-gear"');
    expect(body).not.toContain('The median run');
  });

  it('a bulk result keeps its own view: not the run head, its own chunk (SavedKind)', () => {
    const { body } = render(SavedSim, { props: { result: bulkResult, onrerun: () => {} } });
    expect(body).not.toContain('sim-saved-run');
    expect(body).not.toContain('data-region="head"');
    expect(render(SavedKind, { props: { result: bulkResult, onrerun: () => {} } }).body).toContain(
      'sim-saved-header',
    );
  });
});

describe('SavedSim: Copy for the game', () => {
  it('renders the button disabled while the talent layout has not loaded', () => {
    // A bulk result's page (SavedKind, SavedSim's own chunk for a bulk result).
    const { body } = render(SavedKind, { props: { result: bulkResult, onrerun: () => {} } });
    const match = /<button[^>]*data-testid="sim-saved-copy-for-game"[^>]*>[\s\S]*?<\/button>/.exec(body);
    if (match === null) throw new Error('copy-for-game button not found');
    expect(match[0]).toContain(addonCopy.copyForGame);
    expect(match[0]).toMatch(/\sdisabled/);
  });

  it('copies nothing until the talent layout is in hand', () => {
    expect(gameCodeForSimResult(runResult, '1', null, items)).toBe('');
  });

  it('copies a single run’s stored character', () => {
    const code = gameCodeForSimResult(runResult, '1', talentIndex, items);
    expect(decodedGear(code)).toEqual([
      { slot: 'head', itemId: 12640 },
      { slot: 'main_hand', itemId: 12784 },
    ]);
  });

  it('copies a bulk run’s winning set, not the gear it started from', () => {
    const code = gameCodeForSimResult(bulkResult, '1', talentIndex, items);
    const gear = decodedGear(code);
    // The fixture's leader swaps in Helm and Spaulders of Wrath over the equipped set.
    expect(gear.find((slot) => slot.slot === 'head')?.itemId).toBe(16963);
    expect(gear.find((slot) => slot.slot === 'shoulder')?.itemId).toBe(16966);
  });

  it('copies a weights run’s stored character: a weights result has no winning set', () => {
    const code = gameCodeForSimResult(weightsResult, '1', talentIndex, items);
    const expected = weightsResult.request.character.gear
      .filter((slot) => items.has(slot.item_id))
      .map((slot) => ({ slot: slot.slot, itemId: slot.item_id }));
    expect(decodedGear(code)).toEqual(expected);
  });
});
