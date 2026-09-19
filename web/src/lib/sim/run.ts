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
import type { CharacterSource, CharacterSpec, EncounterSpec, Estimate, SimRequest, SimResult } from './types';
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
  iterations: number;
  randomSeed?: number;
}

export interface RunUpdate {
  estimate: Estimate;
  iterationsDone: number;
  iterationsTotal: number;
}

export interface RunHandle {
  readonly callbackId: string;
  readonly result: Promise<SimResult>;
  cancel(): void;
}

export function buildSimRequest(input: RunInput): SimRequest {
  return {
    engine_version: ENGINE_VERSION,
    spec: input.spec,
    source: input.source,
    character: input.character,
    encounter: input.encounter,
    iterations: input.iterations,
    random_seed: input.randomSeed ?? 0,
  };
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
    const requestJSON = JSON.stringify(buildSimRequest(input));
    const shards = await pool.split(requestJSON, Math.min(pool.size, total));

    const latest = new Map<number, ShardProgress>();
    onUpdate({ estimate: EMPTY_ESTIMATE, iterationsDone: 0, iterationsTotal: total });

    const parts = await pool.run(shards, callbackId, (progress) => {
      latest.set(progress.shard, progress);
      onUpdate({ ...combineEstimate([...latest.values()]), iterationsTotal: total });
    });

    const combined = JSON.parse(await pool.combine(parts)) as SimResult;
    onUpdate({
      estimate: combined.dps,
      iterationsDone: combined.iterations_run,
      iterationsTotal: total,
    });

    return {
      ...combined,
      engine_version: ENGINE_VERSION,
      request: buildSimRequest(input),
      lane: 'browser',
      duration_ms: Math.round(now() - startedAt),
    };
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
