// web/src/lib/sim/bulk-store-request.test.ts
// Unit tests for `buildRequest`'s weights branch, at the `BulkRequestDeps` level rather than
// through a full store + FS1 decode + wasm round trip (bulk-store.test.ts covers that end to
// end for the happy path). The healer-sim defect (BLOCKER 2) is a request that must never
// reach the pool at all -- a hand-built `deps` that would throw if `envelope()` or the pool
// were ever touched is the sharpest way to prove that.
import { describe, expect, it } from 'vitest';
import warriorTalentsJson from '../../fixtures/planner/talents/warrior.json';
import type { TalentFile } from '../planner/types';
import { buildRequest, type BulkRequestDeps } from './bulk-store-request';
import type { WeightsRequest } from './bulk-types';
import { bulkCopy, weightsUnsupportedSpec } from './copy';
import { PRECISION_ITERATIONS } from './precision';
import { defaultSettings } from './settings';
import { specLabel } from './spec-label';
import type { SimCharacter } from './character';
import type { SimPool } from './worker';

const warriorTalents = warriorTalentsJson as unknown as TalentFile;

function character(spec: string): SimCharacter {
  return {
    name: 'Test',
    spec,
    class_slug: 'warrior',
    race_slug: 'orc',
    talent_level: 60,
    tree_version: 'test',
    point_order: [],
    gear: {},
    buffs: [],
    consumables: [],
    source: { kind: 'addon', ref: '', captured_at: '' },
    gear_slots: [],
    professions: [],
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
  };
}

/** A deps object that throws if anything beyond the character/stats/precision getters this
 *  test cares about is touched -- `buildRequest`'s refusal must return before `envelope()`
 *  (and so before `getTalentFile`/`getSettings`) or the pool are ever reached. */
function deps(overrides: Partial<BulkRequestDeps> = {}): BulkRequestDeps {
  const unreached = (name: string) => () => {
    throw new Error(`buildRequest reached ${name}, but the refusal should have returned first`);
  };
  return {
    tool: 'weights',
    mode: null,
    getCharacter: () => character('warrior-fury'),
    getTalentFile: unreached('getTalentFile') as BulkRequestDeps['getTalentFile'],
    getSettings: unreached('getSettings') as BulkRequestDeps['getSettings'],
    getRows: () => [],
    getLocked: () => [],
    getLoadouts: () => [],
    getNamedSets: () => [],
    getPrecision: () => 'normal',
    getCap: () => 0,
    getConsumableIds: () => [],
    getStats: () => ['attack_power'],
    getReferenceStat: () => 'attack_power',
    setPrecision: () => {},
    setCap: () => {},
    setLocked: () => {},
    setLoadouts: () => {},
    setNamedSets: () => {},
    setStats: () => {},
    scheduleCount: () => {},
    poolOnce: unreached('poolOnce') as () => SimPool,
    ...overrides,
  };
}

describe('buildRequest: the weights honest refusal (sub-item 5)', () => {
  it('refuses a healer spec before the request ever reaches the engine', () => {
    const outcome = buildRequest(deps({ getCharacter: () => character('druid-restoration') }));
    expect(outcome).toEqual({ error: weightsUnsupportedSpec(specLabel('druid-restoration')) });
  });

  it('refuses a spec the site has never heard of the same way', () => {
    const outcome = buildRequest(deps({ getCharacter: () => character('nonesuch-spec') }));
    expect(outcome).toEqual({ error: weightsUnsupportedSpec(specLabel('nonesuch-spec')) });
  });

  it('names the spec, never a raw engine string or an internal spec id', () => {
    const outcome = buildRequest(deps({ getCharacter: () => character('druid-restoration') }));
    expect('error' in outcome && outcome.error).not.toMatch(/druid-restoration/);
    expect('error' in outcome && outcome.error).toContain('Restoration Druid');
  });

  it('still refuses an empty stats list for a supported spec, before the engine too', () => {
    const outcome = buildRequest(deps({ getStats: () => [] }));
    expect(outcome).toEqual({ error: bulkCopy.weightsNeedStats });
  });
});

describe('buildRequest: precision drives the weights iteration count (sub-item 1)', () => {
  // D45: a weights run used to send `PRECISION_ITERATIONS.normal` (3,000) no matter what
  // the precision control said, because the control was hidden and the code beneath it
  // never read `precision` at all. A full, real `envelope()` -- a real talent file and the
  // default settings, not the `unreached` stub above -- so this proves the iteration count
  // that would actually reach the engine, not just that a request was returned.
  const withPrecision = (id: 'fast' | 'normal' | 'high') =>
    deps({
      getPrecision: () => id,
      getTalentFile: () => warriorTalents,
      getSettings: () => defaultSettings(),
    });

  it.each([
    ['fast', PRECISION_ITERATIONS.fast],
    ['normal', PRECISION_ITERATIONS.normal],
    ['high', PRECISION_ITERATIONS.high],
  ] as const)('sends %s’s own fixed count, not always normal', (id, iterations) => {
    const outcome = buildRequest(withPrecision(id));
    expect('request' in outcome).toBe(true);
    expect((outcome as { request: WeightsRequest }).request.iterations).toBe(iterations);
  });

  it('fast and high genuinely differ, so the control has something to tighten (D45)', () => {
    const fast = buildRequest(withPrecision('fast')) as { request: WeightsRequest };
    const high = buildRequest(withPrecision('high')) as { request: WeightsRequest };
    expect(high.request.iterations).toBeGreaterThan(fast.request.iterations);
  });
});
