import { describe, expect, it } from 'vitest';
import envelope from '../../fixtures/sim/envelope-v2.json';
import { simCopy } from './copy';
import { SIM_KINDS, requestKind } from './kind';
import type { SimRequest, SimResult } from './types';

// The fixture is the contract's own shapes as data. This single-level `as` (never through
// `unknown`) is the compile-time half of this test (`npx astro check`): TypeScript's JSON
// module inference widens the fixture's string literals (e.g. `source.kind: "addon"`) to
// plain `string`, so a direct assignment or `satisfies` would reject every literal-union
// field regardless of whether the fixture is correct. A single `as`, by contrast, checks
// type comparability rather than plain assignability, which still requires every field
// `SimRequest`/`SimResult` name to exist on the fixture with a comparable shape -- a
// renamed, missing or wrongly-shaped field still fails `astro check` (verified: renaming
// `request.spec` broke this cast with "Property 'spec' is missing"). The assertions below
// are the runtime half -- that the key names the fixture uses are the ones the code reads.
const fixture = envelope as { request: SimRequest; result: SimResult };

describe('requestKind', () => {
  it('is run for a request with neither bulk nor weights', () => {
    expect(requestKind({ bulk: undefined, weights: undefined })).toBe('run');
  });

  it('takes gear, talents and drops from bulk.mode', () => {
    for (const mode of ['gear', 'talents', 'drops'] as const) {
      expect(requestKind({ bulk: { ...fixture.request.bulk!, mode }, weights: undefined })).toBe(mode);
    }
  });

  it('is weights when a weights block is present', () => {
    expect(requestKind({ bulk: undefined, weights: { stats: ['crit'], reference: 'crit' } })).toBe('weights');
  });

  it('prefers bulk over weights so a malformed request never reports two kinds', () => {
    expect(
      requestKind({
        bulk: { ...fixture.request.bulk!, mode: 'gear' },
        weights: { stats: [], reference: '' },
      }),
    ).toBe('gear');
  });

  it('falls back to run for a bulk block with a mode nothing recognises', () => {
    expect(requestKind({ bulk: { ...fixture.request.bulk!, mode: 'nonsense' }, weights: undefined })).toBe(
      'run',
    );
  });

  it('lists the contract’s five kinds', () => {
    expect(SIM_KINDS).toEqual(['run', 'gear', 'talents', 'drops', 'weights']);
  });
});

describe('the amended envelope', () => {
  it('carries the new encounter fields under the contract’s names', () => {
    expect(fixture.request.encounter.style).toBe('light-movement');
    expect(fixture.request.encounter.movement).toEqual({
      interval_sec: 45,
      duration_sec: 5,
      kind: 'away',
    });
    expect(fixture.request.encounter.targets_over_time).toEqual([{ at_sec: 0, count: 1 }]);
    expect(fixture.request.encounter.target_level).toBe(63);
    expect(fixture.request.encounter.target_armor).toBe(0);
    expect(fixture.request.encounter.target_type).toBe('humanoid');
    expect(fixture.request.encounter.dummy).toBe(false);
    expect(fixture.request.target_error).toBe(0.005);
    expect(fixture.request.character.cooldowns).toEqual([{ id: 'spell:11305', at_sec: [0, 90] }]);
  });

  it('carries the new result fields under the contract’s names', () => {
    // Contract A12: an action key, never a name and never a spell id.
    expect(fixture.result.sample?.[0]).toEqual({
      at_ms: -1500,
      action: 'spell:11305',
      resources: { rage: 0 },
    });
    expect(fixture.result.stages).toEqual([{ iterations: 1000, combos: 4 }]);
    expect(fixture.result.combos?.[0].group).toBe(0);
    expect(fixture.result.weights?.[0]).toEqual({ stat: 'crit', weight: 1, error: 0.02 });
  });
});

describe('substitution kinds', () => {
  it('names all four, including contract 10.8’s consumes', () => {
    expect(Object.keys(simCopy.substitutionKindLabel).sort()).toEqual(['consumes', 'item', 'set', 'talents']);
  });
});
