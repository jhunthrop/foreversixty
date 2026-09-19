// web/src/fixtures/sim/engine-fake.ts
// The contract's stub for sim.wasm, with the same four-function surface, so swapping the
// real one in is one environment variable and nothing else.
//
// Five behaviours it has to get right, because the UI can see all five:
//   * iterations split and sum exactly, including the remainder;
//   * progress arrives in ticks, so the DPS figure on screen really does move;
//   * the same seed gives the same numbers, so a paired run is paired;
//   * an abort rejects the pending run rather than resolving it;
//   * simCombine pools the partials the way the real one does, and the combined result
//     carries the fixture summary, so every report component mounts.
//
// The numbers are drawn from the fixture's own mean and standard deviation with a seeded
// generator rather than randomly, so a Playwright assertion on the DPS figure is stable.
import type { EngineModule, ProgressHandler } from '../../lib/sim/engine';
import type { SimRequest, SimResult } from '../../lib/sim/types';
import {
  type BulkRequest,
  type Combination,
  type Combo,
  type Precision,
  type Stage,
  type StageRequests,
  type Substitution,
  type WeightsRequest,
} from '../../lib/sim/bulk-types';
import type { Item } from '../../lib/planner/types';
import fixtureItemsJson from '../planner/items/warrior.json';
import fixtureResultJson from './result.json';
import sampleJson from './sample.json';

const fixture = fixtureResultJson as unknown as SimResult;
/** simdb's stand-in, for the item names contract 10.1 A6 puts on every substitution. */
const fixtureItems = fixtureItemsJson.items as unknown as Item[];
// Design 5.1's sample-iteration log needs a sample on every result the fixture engine
// hands back, so /sim's own e2e can exercise the tab without a second fixture wired
// through simSplit/simCombine.
const fixtureSample = sampleJson as unknown as SimResult['sample'];

export interface FakeEngineOptions {
  tickMs?: number;
  ticks?: number;
  /**
   * Makes every `simRun` reject with this exact message instead of running. Task 5's
   * "keeps the engine's own words" test uses it to prove the message survives every hop to
   * `SimRunError.detail` -- the real wasm answers an unknown buff id with
   * `request: unknown buff: "battle-shout"`, and that sentence is the only thing that
   * tells a player what to change.
   */
  failWith?: string;
}

function seeded(seed: number): () => number {
  let state = seed >>> 0;
  return () => {
    state = (state + 0x6d2b79f5) >>> 0;
    let t = state;
    t = Math.imul(t ^ (t >>> 15), t | 1);
    t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
  };
}

function normal(random: () => number, mean: number, stddev: number): number {
  const u = Math.max(random(), Number.EPSILON);
  const v = random();
  return mean + stddev * Math.sqrt(-2 * Math.log(u)) * Math.cos(2 * Math.PI * v);
}

function statsOf(samples: number[]): SimResult['dps'] & { n: number } {
  const n = samples.length;
  const mean = samples.reduce((a, b) => a + b, 0) / n;
  const variance = samples.reduce((a, b) => a + (b - mean) ** 2, 0) / n;
  const stddev = Math.sqrt(variance);
  return {
    n,
    mean,
    stddev,
    error: n > 0 ? stddev / Math.sqrt(n) : 0,
    min: Math.min(...samples),
    max: Math.max(...samples),
  };
}

const delay = (ms: number): Promise<void> =>
  ms <= 0 ? Promise.resolve() : new Promise((resolve) => setTimeout(resolve, ms));

/** The `{"error": "..."}` envelope simNeedsMore, simValidate and simCount all answer with
 * for input that is not JSON at all -- one place for the message-extraction logic. */
const errorEnvelope = (error: unknown): string =>
  JSON.stringify({ error: error instanceof Error ? error.message : String(error) });

/**
 * The iteration ladder, as contract 1.3 writes it. The real planner reads it from sim/bulk;
 * the fake repeats it because the fake IS the planner here, and the e2e suite asserts the
 * progress line names the right stage count. The final entry of each row matches
 * `BULK_FINAL_ITERATIONS` in `precision.ts` (3,000 fast/normal, 10,000 high); the earlier
 * stages are a fixture-only ladder, not a value that crosses the wasm boundary.
 */
const LADDER: Record<Precision, number[]> = {
  fast: [100, 1000, 3000],
  normal: [1000, 3000],
  high: [1000, 10000],
};

/** The fixture item file stands in for simdb, so a substitution comes back named. */
function itemName(itemId: number): string {
  return fixtureItems.find((item) => item.id === itemId)?.name ?? `Item ${itemId}`;
}

/**
 * The fake's expansion, following contract 10.1 A4: **mode decides**. `gear` takes the
 * product of every candidate group -- one group per slot, each group being "keep what is
 * equipped" plus that slot's candidates -- crossed with the consumable alternatives of A5;
 * `drops` and `talents` take one substitution at a time. The empty product (nothing
 * substituted) is the equipped set, which is `requests[0]` and not a combination, so it is
 * dropped.
 *
 * Substitutions come back NAMED (contract 10.1 A6): the real planner reads the name from
 * simdb, and the fake reads it from the fixture item file, so the page's "never re-join an
 * id to a name" rule is exercised rather than merely asserted.
 *
 * A candidate whose slot is "" fits more than one slot and only the item database can say
 * which; the fake has no database, so it resolves every one to `finger1`. Nothing in the
 * suite asserts on that resolution -- the real engine's own smoke test covers it.
 *
 * Consumable alternatives (A5) are crossed onto every gear combination, INCLUDING the one
 * where gear is left untouched: A5 says "each inner list replaces Consumes for that
 * combination", and a gear-untouched, consumes-changed run is a real, distinct combination
 * (the base request's own consumes are `[]` in the fixtures, so neither alternative list is
 * ever a no-op). There is therefore no "untouched base" left to subtract once a consumables
 * list is present.
 */
function expand(request: BulkRequest): Combination[] {
  const bulk = request.bulk;
  const slotOf = (slot: string): string => (slot === '' ? 'finger1' : slot);
  const itemSubs: Substitution[] = bulk.candidates
    .filter((candidate) => !(bulk.locked ?? []).includes(slotOf(candidate.slot)))
    .map((candidate) => ({
      kind: 'item',
      slot: slotOf(candidate.slot),
      item_id: candidate.item_id,
      enchant: candidate.enchant,
      suffix: candidate.suffix,
      origin: candidate.origin,
      name: itemName(candidate.item_id),
      source_name: candidate.source_name,
    }));

  let groups: Substitution[][];
  if (bulk.mode === 'gear') {
    // The product of one group per slot, each group led by "nothing" -- which is exactly
    // the set of subsets that use each slot at most once, in a stable order.
    const bySlot = new Map<string, Substitution[]>();
    for (const sub of itemSubs) {
      const existing = bySlot.get(sub.slot ?? '');
      if (existing === undefined) bySlot.set(sub.slot ?? '', [sub]);
      else existing.push(sub);
    }
    let product: Substitution[][] = [[]];
    for (const options of bySlot.values()) {
      product = product.flatMap((base) => [base, ...options.map((sub) => [...base, sub])]);
    }
    // Consumable alternatives multiply the WHOLE gear product, including the untouched
    // entry (contract 10.1 A5): every base, whether it changed gear or not, gets tried
    // against each alternative consumable list.
    const consumableLists = bulk.consumables ?? [];
    if (consumableLists.length > 0) {
      product = product.flatMap((base) =>
        consumableLists.map((list) => [...base, { kind: 'consumes', name: list.join(', ') } as Substitution]),
      );
    }
    groups = product.filter((group) => group.length > 0).sort((a, b) => a.length - b.length);
  } else {
    groups = itemSubs.map((sub) => [sub]);
  }

  for (const loadout of bulk.talents ?? []) {
    groups.push([{ kind: 'talents', name: loadout.name, talents: loadout.talents }]);
  }
  for (const set of bulk.sets ?? []) groups.push([{ kind: 'set', name: set.name }]);

  return groups.map((substitutions) => ({
    request: applySubstitutions(request, substitutions),
    substitutions,
  }));
}

/** The base character with the substitutions written into its gear, talents and consumes. */
function applySubstitutions(request: BulkRequest, substitutions: readonly Substitution[]): SimRequest {
  const gear = request.character.gear.map((slot) => ({ ...slot }));
  let talents = request.character.talents;
  let consumes = [...request.character.consumes];
  for (const sub of substitutions) {
    if (sub.kind === 'talents') {
      talents = sub.talents ?? talents;
      continue;
    }
    if (sub.kind === 'consumes') {
      // A5: the inner list REPLACES the character's consumes for this combination.
      consumes = (sub.name ?? '').split(', ').filter((id) => id !== '');
      continue;
    }
    if (sub.kind !== 'item') continue;
    const existing = gear.findIndex((slot) => slot.slot === sub.slot);
    const next = { slot: sub.slot!, item_id: sub.item_id!, enchant: sub.enchant, suffix: sub.suffix };
    if (existing >= 0) gear[existing] = next;
    else gear.push(next);
  }
  const { bulk: _bulk, ...base } = request;
  return { ...base, character: { ...request.character, gear, talents, consumes } };
}

/** The equipped set, with the bulk block stripped: it is a plain run like any other. */
function equippedRequest(request: BulkRequest, iterations: number): SimRequest {
  const { bulk: _bulk, ...base } = request;
  return { ...base, iterations };
}

function stageAt(request: BulkRequest, stage: number, combos: Combination[], ran: Stage[]): StageRequests {
  const iterations = LADDER[request.bulk.precision as Precision][stage - 1];
  return {
    stage,
    iterations,
    requests: [
      equippedRequest(request, iterations),
      ...combos.map((combo) => ({ ...combo.request, iterations })),
    ],
    combos,
    // Contract 10.1 A10: the ladder's history rides on the stage object across the
    // boundary, so nothing stateful lives in the engine between calls.
    ran: [...ran],
  };
}

export function createFakeEngine(options: FakeEngineOptions = {}): EngineModule {
  const tickMs = options.tickMs ?? 120;
  const ticks = options.ticks ?? 10;
  const failWith = options.failWith ?? '';
  const aborted = new Set<string>();
  // Tracks a run's callback id for exactly as long as simRun is in flight, so simAbort can
  // answer {"aborted": false} for an id nothing registered -- main.go's own distinction
  // between "you stopped it" and "you called this wrong".
  const active = new Set<string>();
  let progress: ProgressHandler = () => {};

  // A plain closure, not a method on the returned object: `simWeights` calls this directly
  // (task-3 fix round 1, finding 2) rather than through `this.simRun`, so it does not depend
  // on being invoked through the exact object reference `createFakeEngine` returns -- every
  // other cross-call in this file already goes through a local closure the same way.
  async function runSim(requestJSON: string, callbackId: string): Promise<string> {
    aborted.delete(callbackId);
    active.add(callbackId);
    try {
      if (failWith !== '') {
        await delay(tickMs);
        throw new Error(failWith);
      }
      const request = JSON.parse(requestJSON) as SimRequest;
      const random = seeded(request.random_seed * 7919 + request.iterations);
      const samples: number[] = [];
      const perTick = Math.max(1, Math.ceil(request.iterations / ticks));
      const startedAt = Date.now();

      while (samples.length < request.iterations) {
        await delay(tickMs);
        if (aborted.has(callbackId)) {
          aborted.delete(callbackId);
          throw new Error(`sim run ${callbackId} aborted`);
        }
        const upTo = Math.min(request.iterations, samples.length + perTick);
        while (samples.length < upTo) {
          samples.push(normal(random, fixture.dps.mean, fixture.dps.stddev));
        }
        const { n, ...dps } = statsOf(samples);
        progress(callbackId, JSON.stringify({ iterations_run: n, dps }));
      }

      const { n, ...dps } = statsOf(samples);
      return JSON.stringify({
        engine_version: request.engine_version,
        request,
        lane: 'browser',
        dps,
        iterations_run: n,
        duration_ms: Date.now() - startedAt,
        summary: fixture.summary,
        sample: fixtureSample,
      } satisfies SimResult);
    } finally {
      active.delete(callbackId);
    }
  }

  return {
    onProgress(handler) {
      progress = handler;
    },

    simSplit(requestJSON, n) {
      const request = JSON.parse(requestJSON) as SimRequest;
      const used = Math.max(1, Math.min(n, request.iterations));
      const base = Math.floor(request.iterations / used);
      const remainder = request.iterations % used;
      const parts: SimRequest[] = Array.from({ length: used }, (_, i) => ({
        ...request,
        iterations: base + (i < remainder ? 1 : 0),
        random_seed: request.random_seed === 0 ? i + 1 : request.random_seed * 1000 + i,
      }));
      return JSON.stringify(parts);
    },

    simRun: runSim,

    simCombine(resultsJSON) {
      const parts = JSON.parse(resultsJSON) as SimResult[];
      let n = 0;
      let sum = 0;
      let sumSquares = 0;
      let min = Number.POSITIVE_INFINITY;
      let max = Number.NEGATIVE_INFINITY;
      let durationMs = 0;
      for (const part of parts) {
        n += part.iterations_run;
        sum += part.dps.mean * part.iterations_run;
        sumSquares += part.iterations_run * (part.dps.stddev ** 2 + part.dps.mean ** 2);
        min = Math.min(min, part.dps.min);
        max = Math.max(max, part.dps.max);
        durationMs = Math.max(durationMs, part.duration_ms);
      }
      const mean = n === 0 ? 0 : sum / n;
      const variance = n === 0 ? 0 : Math.max(0, sumSquares / n - mean * mean);
      const stddev = Math.sqrt(variance);
      const first = parts[0];
      return JSON.stringify({
        engine_version: first.engine_version,
        request: { ...first.request, iterations: n },
        lane: 'browser',
        dps: {
          mean,
          stddev,
          error: n > 0 ? stddev / Math.sqrt(n) : 0,
          min: n === 0 ? 0 : min,
          max: n === 0 ? 0 : max,
        },
        iterations_run: n,
        duration_ms: durationMs,
        summary: fixture.summary,
        sample: fixtureSample,
      } satisfies SimResult);
    },

    simAbort(callbackId) {
      const registered = active.has(callbackId);
      if (registered) aborted.add(callbackId);
      return JSON.stringify({ aborted: registered });
    },

    simPlan(requestJSON) {
      const request = JSON.parse(requestJSON) as BulkRequest;
      const combos = expand(request);
      if (combos.length > request.bulk.cap) {
        return JSON.stringify({
          error: 'cap_exceeded',
          cap: request.bulk.cap,
          combinations: combos.length,
        });
      }
      return JSON.stringify(stageAt(request, 1, combos, []));
    },

    simRank(requestJSON, stageJSON, resultsJSON) {
      const request = JSON.parse(requestJSON) as BulkRequest;
      const stage = JSON.parse(stageJSON) as StageRequests;
      const results = JSON.parse(resultsJSON) as SimResult[];
      const equipped = results[0].dps;
      const scored = stage.combos
        .map((combo, index) => ({ combo, dps: results[index + 1].dps }))
        .sort((a, b) => b.dps.mean - a.dps.mean);

      const ladder = LADDER[request.bulk.precision as Precision];
      const ran: Stage[] = [
        ...(stage.ran ?? []),
        { iterations: stage.iterations, combos: stage.combos.length },
      ];

      if (stage.stage < ladder.length) {
        // The cut: the top half at every stage but the last. The real planner keeps a
        // quarter plus anything within two standard errors; the fake keeps a fixed
        // fraction so a test can predict how many requests the next stage carries.
        const keep = Math.max(1, Math.ceil(scored.length / 2));
        const survivors = scored.slice(0, keep).map((entry) => entry.combo);
        return JSON.stringify({ next: stageAt(request, stage.stage + 1, survivors, ran) });
      }

      // Within-error grouping: a run whose delta interval overlaps the leader's is group 0,
      // then each next non-overlapping run opens the next group.
      const combos: Combo[] = [];
      let group = 0;
      let boundary = Number.POSITIVE_INFINITY;
      for (const entry of scored) {
        const delta = {
          mean: entry.dps.mean - equipped.mean,
          stddev: 0,
          error: Math.hypot(entry.dps.error, equipped.error),
          min: 0,
          max: 0,
        };
        const high = delta.mean + 1.96 * delta.error;
        if (high < boundary) {
          if (Number.isFinite(boundary)) group += 1;
          boundary = delta.mean - 1.96 * delta.error;
        }
        combos.push({ substitutions: entry.combo.substitutions, dps: entry.dps, delta, group });
      }

      return JSON.stringify({
        result: {
          ...results[0],
          request,
          lane: 'browser',
          combos,
          equipped,
          stages: ran,
        } satisfies SimResult,
      });
    },

    async simWeights(requestJSON, callbackId) {
      const request = JSON.parse(requestJSON) as WeightsRequest;
      const base = await runSim(JSON.stringify({ ...request }), callbackId);
      const result = JSON.parse(base) as SimResult;
      // Seeded off the stat name, so the same request gives the same weights every time.
      const weights = request.weights.stats.map((stat) => {
        if (stat === request.weights.reference) return { stat, weight: 1, error: 0 };
        const random = seeded([...stat].reduce((sum, ch) => sum + ch.charCodeAt(0), 0));
        return {
          stat,
          weight: Math.round(random() * 3000) / 100,
          error: Math.round(random() * 200) / 100 + 0.01,
        };
      });
      return JSON.stringify({ ...result, request, weights });
    },

    simNeedsMore(resultJSON, requestJSON) {
      try {
        const result = JSON.parse(resultJSON) as SimResult;
        const request = JSON.parse(requestJSON) as SimRequest;
        const target = request.target_error ?? 0;
        // Four ways to be done, and the real engine agrees on all four: this was not a
        // target-error run; nothing has been measured, so a relative error is undefined;
        // the band is inside the target; or the ceiling is reached.
        const needsMore =
          target > 0 &&
          result.dps.mean > 0 &&
          result.dps.error / result.dps.mean > target &&
          result.iterations_run < request.iterations;
        return JSON.stringify({ needs_more: needsMore });
      } catch (error) {
        return errorEnvelope(error);
      }
    },

    simValidate(requestJSON) {
      let request: SimRequest;
      try {
        request = JSON.parse(requestJSON) as SimRequest;
      } catch (error) {
        return errorEnvelope(error);
      }
      // A short stand-in for api.SimRequest.Validate: the checks the drawer's own tests
      // exercise. The real wasm runs the whole thing, and the drawer renders whatever
      // fields come back, so a fake that refuses fewer things cannot make the page wrong.
      const errors: { field: string; message: string }[] = [];
      if (request.engine_version === undefined || request.engine_version === '') {
        errors.push({ field: 'engine_version', message: 'engine_version is required' });
      }
      if (request.spec === undefined || request.spec === '') {
        errors.push({ field: 'spec', message: 'spec is required' });
      }
      if (!(request.iterations > 0)) {
        errors.push({ field: 'iterations', message: 'iterations must be a positive number' });
      }
      const targets = request.encounter?.targets;
      if (targets !== undefined && (targets < 1 || targets > 10)) {
        errors.push({ field: 'encounter.targets', message: 'targets must be between 1 and 10' });
      }
      return JSON.stringify({ ok: errors.length === 0, errors });
    },

    simCount(requestJSON) {
      let request: SimRequest;
      try {
        request = JSON.parse(requestJSON) as SimRequest;
      } catch (error) {
        return errorEnvelope(error);
      }
      const bulk = request.bulk;
      if (bulk === undefined) return JSON.stringify({ combinations: 0 });
      // Reuses simPlan's own `expand()` rather than re-deriving the arithmetic (controller
      // ruling, task-3 fix round 1): simCount and simPlan must agree on what a request
      // expands to, because Task 15's live cap-notice UI and its client-side server-cap
      // gate both read simCount, and a fake that disagreed with simPlan would bake a wrong
      // number into that UI and its tests. This one call site is now the only place the
      // fake computes a combination count.
      const combinations = expand(request as BulkRequest).length;
      if (bulk.cap > 0 && combinations > bulk.cap) {
        return JSON.stringify({ error: 'cap_exceeded', cap: bulk.cap, combinations });
      }
      return JSON.stringify({ combinations });
    },
  };
}
