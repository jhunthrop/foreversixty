// web/src/lib/sim/engine.test.ts
// Pins EngineModule -- and the fake that implements it -- to the shapes
// sim/cmd/wasm/main.go actually ships, since nothing tested engine.ts before the final
// whole-branch review's H1: simSplit and simCombine each cross the boundary as ONE JSON
// string (never a JS array), simAbort answers {"aborted": boolean}, and every failure of
// the three returns {"error": "..."}.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { createFakeEngine } from '../../fixtures/sim/engine-fake';
import { unwrapOrThrow, type EngineModule } from './engine';
import { DEFAULT_ENCOUNTER, type SimRequest, type SimResult } from './types';
import { ENGINE_VERSION } from './version';

// The engine lane's wasm entrypoint. Same existsSync/skipIf pattern as version.test.ts's
// ENGINEVER_GO pin: a skip that names its own condition, not a quiet pass on a missing file.
const MAIN_GO = path.resolve(import.meta.dirname, '../../../../sim/cmd/wasm/main.go');
// The shared JSON helpers moved to exports.go when the bulk exports landed; the assertions
// below read both files as one source.
const EXPORTS_GO = path.resolve(import.meta.dirname, '../../../../sim/cmd/wasm/exports.go');
const mainGoPinned = existsSync(MAIN_GO);

const request: SimRequest = {
  engine_version: ENGINE_VERSION,
  spec: 'warrior-fury',
  source: { kind: 'addon', ref: '', captured_at: '2026-09-14T10:00:00Z' },
  character: {
    name: 'Thrallgar',
    race: 'orc',
    class: 'warrior',
    level: 60,
    talents: '-5530515-',
    gear: [{ slot: 'head', item_id: 12640 }],
    buffs: [],
    consumes: [],
  },
  encounter: DEFAULT_ENCOUNTER,
  iterations: 3000,
  random_seed: 7,
};

describe('sim/cmd/wasm/main.go, as shipped', () => {
  it.skipIf(!mainGoPinned)(
    'simSplit/simCombine cross the boundary as one JSON string, never a JS array',
    () => {
      const src = readFileSync(MAIN_GO, 'utf8') + readFileSync(EXPORTS_GO, 'utf8');
      // simSplit: combine.Split returns []api.SimRequest, and the whole slice is marshalled
      // as one value -- not marshalled per-element into a []string.
      expect(src).toMatch(/func simSplit\([^)]*\)[\s\S]*?json\.Marshal\(parts\)/);
      // simCombine: args[0].String() is unmarshalled as one JSON array, not iterated per-arg.
      expect(src).toMatch(
        /var parts \[\]api\.SimResult[\s\S]*?json\.Unmarshal\(\[\]byte\(args\[0\]\.String\(\)\), &parts\)/,
      );
      // simAbort answers a JSON object so a malformed call is distinguishable from "no such run".
      expect(src).toMatch(/Aborted bool `json:"aborted"`/);
      // Every one of the three fails the same way.
      expect(src).toMatch(/func errorJSON\(msg string\) string/);
    },
  );
});

describe('EngineModule (fake)', () => {
  it('simSplit returns one JSON string encoding a SimRequest[]', () => {
    const engine: EngineModule = createFakeEngine({ tickMs: 0 });
    const splitJSON = engine.simSplit(JSON.stringify(request), 4);
    expect(typeof splitJSON).toBe('string');
    const parsed: unknown = JSON.parse(splitJSON);
    expect(Array.isArray(parsed)).toBe(true);
    const shards = parsed as SimRequest[];
    expect(shards).toHaveLength(4);
    expect(shards.reduce((sum, shard) => sum + shard.iterations, 0)).toBe(3000);
    expect(typeof shards[0].character).toBe('object');
  });

  it('simCombine takes one JSON string encoding a SimResult[]', async () => {
    const engine: EngineModule = createFakeEngine({ tickMs: 0 });
    const shards = (JSON.parse(engine.simSplit(JSON.stringify(request), 4)) as SimRequest[]).map((shard) =>
      JSON.stringify(shard),
    );
    const results = await Promise.all(shards.map((shard, i) => engine.simRun(shard, `s${i}`)));
    const combineInput = JSON.stringify(results.map((r) => JSON.parse(r) as SimResult));
    const combined = JSON.parse(engine.simCombine(combineInput)) as SimResult;
    expect(combined.iterations_run).toBe(3000);
  });

  it('simAbort returns {"aborted": boolean} JSON, false for an id nothing registered', () => {
    const engine: EngineModule = createFakeEngine({ tickMs: 0 });
    expect(JSON.parse(engine.simAbort('nobody'))).toEqual({ aborted: false });
  });
});

describe('unwrapOrThrow', () => {
  it('passes success JSON through unchanged, array or object', () => {
    expect(unwrapOrThrow('[1,2,3]')).toBe('[1,2,3]');
    expect(unwrapOrThrow('{"aborted":true}')).toBe('{"aborted":true}');
  });

  it('throws the engine\'s own message for an {"error"} envelope', () => {
    expect(() => unwrapOrThrow('{"error":"unknown buff: \\"battle-shout\\""}')).toThrow(
      'unknown buff: "battle-shout"',
    );
  });

  it('does not mistake a SimResult carrying no .error for a failure', () => {
    expect(() => unwrapOrThrow(JSON.stringify({ iterations_run: 400 }))).not.toThrow();
  });

  // Defect A: weightsJSON (sim/cmd/wasm/exports.go) fails a weights run the same way
  // simRun fails a shard -- a full SimResult JSON with `.error` set (failJSON), not a bare
  // `{"error": "..."}` -- and this generic check already throws on either shape, since it
  // only ever looks at the top-level `.error` field. The bug was never here: it was that
  // nothing called this function on simWeights' answer at all (see the test below).
  it("throws on a full SimResult carrying .error, the shape weightsJSON's failJSON sends", () => {
    const failedWeightsResult = JSON.stringify({
      request: { spec: 'warrior-fury' },
      error: 'iterations must be one of [500 3000 10000], got 60',
      summary: {},
    });
    expect(() => unwrapOrThrow(failedWeightsResult)).toThrow(
      'iterations must be one of [500 3000 10000], got 60',
    );
  });
});

describe("loadWasmEngine's simWeights binding (source pin)", () => {
  // Real wasm loading needs WebAssembly.instantiateStreaming and a live worker, which this
  // suite cannot run -- the same reason main.go/exports.go above are pinned by source rather
  // than exercised. Defect A: simWeights' promise resolves with a full SimResult carrying
  // `.error` on a refused or failed run (weightsJSON's own failJSON, sim/cmd/wasm/exports.go)
  // rather than rejecting -- unlike simValidate/simRank/simNeedsMore, its binding here used
  // to return that answer straight through, so a refused browser-lane weights run (a
  // "fast"-precision request outside the settings bar's closed iteration set, before that
  // was fixed in sim/api/envelope.go) resolved as an ordinary, empty result and the page
  // never learned the run had failed. Pinned so a future edit cannot quietly drop the unwrap.
  const ENGINE_TS = path.resolve(import.meta.dirname, 'engine.ts');

  it("runs simWeights' answer through unwrapOrThrow, the same as simValidate and simRank", () => {
    const src = readFileSync(ENGINE_TS, 'utf8');
    expect(src).toMatch(/simWeights:\s*async[\s\S]{0,120}unwrapOrThrow\(await globals\.simWeights!/);
  });
});
