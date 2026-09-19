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
  it.skipIf(!mainGoPinned)('simSplit/simCombine cross the boundary as one JSON string, never a JS array', () => {
    const src = readFileSync(MAIN_GO, 'utf8');
    // simSplit: combine.Split returns []api.SimRequest, and the whole slice is marshalled
    // as one value -- not marshalled per-element into a []string.
    expect(src).toMatch(/func simSplit\([^)]*\)[\s\S]*?json\.Marshal\(parts\)/);
    // simCombine: args[0].String() is unmarshalled as one JSON array, not iterated per-arg.
    expect(src).toMatch(/var parts \[\]api\.SimResult[\s\S]*?json\.Unmarshal\(\[\]byte\(args\[0\]\.String\(\)\), &parts\)/);
    // simAbort answers a JSON object so a malformed call is distinguishable from "no such run".
    expect(src).toMatch(/Aborted bool `json:"aborted"`/);
    // Every one of the three fails the same way.
    expect(src).toMatch(/func errorJSON\(msg string\) string/);
  });
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
});
