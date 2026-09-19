// web/src/lib/sim/run.ts
// One run, end to end: build the envelope, split it, run every shard with its progress
// pooled into one figure, combine, and hand back a SimResult in the contract's shape.
//
// The summary is NOT built here. sim/adapter runs inside the wasm, so what comes out of
// simCombine already carries a summary.Summary built by the same Go code the server lane
// runs. Four fields are overwritten on the way out -- engine_version, request, lane and
// duration_ms -- because those are the browser's facts about the run, not the engine's.
//
// Every failure a player can cause or see becomes one of two SimRunErrors: a stop they
// asked for, or an engine that could not run this character.
import { simCopy } from './copy';
import { combineEstimate, EMPTY_ESTIMATE, type ShardProgress } from './estimate';
import { relativeError } from './precision';
import type { CharacterSource, CharacterSpec, EncounterSpec, Estimate, SimRequest, SimResult } from './types';
import { STEP_ITERATIONS_DEFAULT } from './types';
import { ENGINE_VERSION } from './version';
import type { SimPool } from './worker';

/**
 * A failure, with the engine's own words beside the generic sentence.
 *
 * sim/request refuses an unknown or ambiguous buff or consumable id rather than dropping it
 * -- "a sim that quietly ran without a world buff reports a DPS number that is wrong and
 * says nothing about why" -- and its message names the id. It does the same for a spec it
 * has no model for (ErrUnknownSpec), and sim/adapter for a duplicate summary row
 * (ErrDuplicateRow). None of those is worth paraphrasing, so `detail` carries the engine's
 * text verbatim and the UI shows it under `simCopy.failed`. Empty for a cancel, and for a
 * failure that carried no message.
 */
export class SimRunError extends Error {
  constructor(
    message: string,
    readonly cancelled: boolean,
    readonly detail: string,
    options: { cause?: unknown } = {},
  ) {
    super(message, options);
    this.name = 'SimRunError';
  }
}

/** The engine's own words for a failure, or empty when it gave none. */
function detailOf(cause: unknown): string {
  if (cause instanceof Error) return cause.message;
  return typeof cause === 'string' ? cause : '';
}

export interface RunInput {
  spec: string;
  source: CharacterSource;
  character: CharacterSpec;
  encounter: EncounterSpec;
  /** The count for a fixed run; the ceiling for a target-error run. */
  iterations: number;
  randomSeed?: number;
  /** > 0 turns this into a target-error run (contract 1.2). */
  targetError?: number;
  /** The size of one step of a target-error run. Ignored without `targetError`. */
  stepIterations?: number;
}

export interface RunUpdate {
  estimate: Estimate;
  iterationsDone: number;
  /** The fixed count, or the ceiling of a target-error run. */
  iterationsTotal: number;
  /** `error / mean`, for the "±41 DPS, 0.4%" half of the progress line. Zero before the first tick. */
  relativeError: number;
}

export interface RunHandle {
  readonly callbackId: string;
  readonly result: Promise<SimResult>;
  cancel(): void;
}

export function buildSimRequest(input: RunInput): SimRequest {
  const request: SimRequest = {
    engine_version: ENGINE_VERSION,
    spec: input.spec,
    source: input.source,
    character: input.character,
    encounter: input.encounter,
    iterations: input.iterations,
    random_seed: input.randomSeed ?? 0,
  };
  // Omitted rather than sent as 0: `target_error` is `omitempty` on the Go side, and a
  // request that carries the key with a zero in it reads, in the drawer and in a share
  // URL, as a deliberate choice rather than as today's fixed-count run.
  return (input.targetError ?? 0) > 0 ? { ...request, target_error: input.targetError } : request;
}

let runCounter = 0;

function nextCallbackId(): string {
  runCounter += 1;
  return `sim-${runCounter}`;
}

export function runSim(
  pool: SimPool,
  input: RunInput,
  onUpdate: (update: RunUpdate) => void,
  now: () => number = () => Date.now(),
): RunHandle {
  const callbackId = nextCallbackId();
  const total = input.iterations;
  let cancelled = false;

  async function execute(): Promise<SimResult> {
    const startedAt = now();
    const request = buildSimRequest(input);
    const requestJSON = JSON.stringify(request);
    const targetError = input.targetError ?? 0;
    const step = targetError > 0 ? Math.max(1, input.stepIterations ?? STEP_ITERATIONS_DEFAULT) : total;

    // Every shard result of every step, in order. simCombine pools the lot at the end, so
    // the figure a target-error run reports is the whole run's and not the last step's.
    const everyPart: string[] = [];
    let done = 0;

    onUpdate({ estimate: EMPTY_ESTIMATE, iterationsDone: 0, iterationsTotal: total, relativeError: 0 });

    for (;;) {
      // A stop that lands between two steps ends the run here rather than starting
      // another one; a stop mid-step still lets that step's shards finish (the pool's
      // abort is a request to the engine, not a guarantee it wins the race -- see the
      // header note above and store.svelte.ts's own stopRequested guard).
      if (cancelled) throw new Error('sim run: stopped between steps');

      const thisStep = Math.min(step, total - done);
      if (thisStep <= 0) break;

      // The step's own request: the same envelope with this step's count. The seed is left
      // alone -- 0 means "random", which is what makes each step independent, and a paired
      // run that pinned one gets the same pinned one every step, which is what pairing means.
      const stepJSON = JSON.stringify({ ...request, iterations: thisStep });
      const shards = await pool.split(stepJSON, Math.min(pool.size, thisStep));

      const latest = new Map<number, ShardProgress>();
      const before = done;
      const parts = await pool.run(shards, `${callbackId}-s${everyPart.length}`, (progress) => {
        latest.set(progress.shard, progress);
        const pooled = combineEstimate([...latest.values()]);
        onUpdate({
          estimate: pooled.estimate,
          iterationsDone: before + pooled.iterationsDone,
          iterationsTotal: total,
          relativeError: relativeError(pooled.estimate),
        });
      });

      everyPart.push(...parts);
      done += thisStep;

      const combinedJSON = await pool.combine(everyPart);
      const combined = JSON.parse(combinedJSON) as SimResult;
      onUpdate({
        estimate: combined.dps,
        iterationsDone: combined.iterations_run,
        iterationsTotal: total,
        relativeError: relativeError(combined.dps),
      });

      const finished = (): SimResult => ({
        ...combined,
        engine_version: ENGINE_VERSION,
        request,
        lane: 'browser',
        duration_ms: Math.round(now() - startedAt),
      });

      // The engine decides. Never a comparison written here: the page does not own a
      // stopping rule over an estimate it did not pool itself.
      if (targetError === 0 || done >= total) return finished();
      if (!(await pool.needsMore(combinedJSON, requestJSON))) return finished();
    }

    // Unreachable for any positive `total`; a zero-iteration request is refused at the
    // boundary long before it reaches here, and throwing says so rather than returning a
    // result nothing produced.
    throw new Error('sim run: no iterations to run');
  }

  const result = execute().catch((cause: unknown) => {
    throw new SimRunError(
      cancelled ? simCopy.stopped : simCopy.failed,
      cancelled,
      cancelled ? '' : detailOf(cause),
      { cause },
    );
  });

  return {
    callbackId,
    result,
    cancel() {
      cancelled = true;
      pool.abort(callbackId);
    },
  };
}
