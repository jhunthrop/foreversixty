// web/src/components/sim/tools/StatWeightsSpecScope.test.ts
// Final whole-branch review, Finding 2: GET /v1/specs carries no `weight_stats` column in
// production -- api/internal/sims/specs.go's own SpecFidelity struct has no such field, and
// Store.Specs never sets one. Only the shared test fixture (src/fixtures/sim/specs.json)
// hand-carried it, which hid that the picker's spec-scoping shipped inert: every real spec
// resolved to `undefined`, `pickableStatsFor` fell back to the full 17-entry WEIGHT_STATS,
// and the original D45 defect reproduced exactly.
//
// This exercises the real chain -- `createBulkStore` -> `weightStatsFor` -> `specRow` ->
// `pickableStatsFor` -- against an API response shaped like today's real one (no
// `weight_stats` field at all), through the actual `StatWeights.svelte` render, not a
// hand-mocked store. It also covers Finding 3: the "engine weighs" note must never disagree
// with what the picker actually offers, including the one shape `pickableStatsFor` treats as
// "not said" that a bare `!== undefined` check does not -- an explicit empty list.
//
// Separate file, not folded into StatWeights.test.ts: `svelte/server`'s render() is a
// non-DOM (string) renderer and breaks under jsdom (the environment `test-support/sim-api.ts`
// normally runs under, since its own reset() clears `document.cookie` for CSRF). Both calls
// this file makes -- `loadAddon` and `loadSpecs` -- are GET-only and never read that cookie,
// so this stubs and unstubs `fetch` directly rather than pulling in the CSRF-clearing reset()
// and the jsdom environment it requires, keeping this file's render on the same plain-Node
// footing as every other StatWeights render test.
import { render } from 'svelte/server';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { createBulkStore } from '../../../lib/sim/bulk-store.svelte';
import { WEIGHTS_STATS_FROM_ENGINE } from '../../../lib/sim/copy';
import { specRow } from '../../../lib/sim/spec-label';
import {
  createSimApi,
  envelope,
  fixtureSpecs,
  FIXTURE_DATA_BUILD,
  TEST_API,
} from '../../../test-support/sim-api';
import StatWeights from './StatWeights.svelte';

const FURY = `FS1:${FIXTURE_DATA_BUILD}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;
const api = createSimApi();

beforeEach(() => api.install());
afterEach(() => vi.unstubAllGlobals());

describe('StatWeights: weight_stats resolves through the curated spec list, and the note never disagrees with the picker (Findings 2 & 3)', () => {
  it('offers the spec’s own curated stats, and shows the note, even though the API row sent no weight_stats -- the real production shape', async () => {
    const store = createBulkStore({ tool: 'weights', treeVersion: FIXTURE_DATA_BUILD, apiBase: TEST_API });
    await store.loadAddon(FURY);
    await store.loadSpecs();
    const curated = specRow('warrior-fury')?.weight_stats;
    expect(curated?.length).toBeGreaterThan(0);
    expect(store.weightStats).toEqual(curated);

    const { body } = render(StatWeights, { props: { store, me: null } });
    expect(body).toContain(WEIGHTS_STATS_FROM_ENGINE);
    expect(body).toContain('data-testid="sim-weight-pick-melee_haste"');
    // mp5/spell_haste/feral_attack_power are retail-only and never in warrior-fury's own
    // curated list (D45); expertise and armor penetration do not exist on the 1.60 client
    // and left every curated list in 3d364a9.
    expect(body).not.toContain('data-testid="sim-weight-pick-mp5"');
    expect(body).not.toContain('data-testid="sim-weight-pick-expertise"');
  });

  it('never shows the note over the fallback list -- an explicit empty weight_stats falls back exactly like an absent one', async () => {
    const emptyRow = {
      ...fixtureSpecs.find((row) => row.spec === 'warrior-fury')!,
      weight_stats: [] as string[],
    };
    api.route({
      method: 'GET',
      pattern: /\/v1\/specs$/,
      respond: () => envelope({ specs: [emptyRow] }),
    });
    const store = createBulkStore({ tool: 'weights', treeVersion: FIXTURE_DATA_BUILD, apiBase: TEST_API });
    await store.loadAddon(FURY);
    await store.loadSpecs();
    expect(store.weightStats).toEqual([]);

    const { body } = render(StatWeights, { props: { store, me: null } });
    expect(body).not.toContain(WEIGHTS_STATS_FROM_ENGINE);
    // The full pinned vocabulary, the same fallback an absent column gets -- an empty list
    // must not be dressed up as a curated one either (D45's own defect, in miniature).
    expect(body).toContain('data-testid="sim-weight-pick-expertise"');
    expect(body).toContain('data-testid="sim-weight-pick-mp5"');
  });
});
