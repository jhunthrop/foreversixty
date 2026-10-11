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
import SavedSim from './SavedSim.svelte';

const runResult = runFixture as unknown as SimResult;
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

describe('SavedSim: Copy for the game', () => {
  it('renders the button disabled while the talent layout has not loaded', () => {
    const { body } = render(SavedSim, { props: { result: runResult, onrerun: () => {} } });
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
