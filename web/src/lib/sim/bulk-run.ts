// web/src/lib/sim/bulk-run.ts
// The stage loop every combination tool runs on: simPlan, then each stage's requests
// through the pool, then simRank, then the next stage or the finished result.
//
// Nothing here is a statistic. Which combinations exist, which survive a cut, how many
// iterations each stage runs and which runs are within error of the leader are all decided
// by sim/bulk inside the wasm. This file moves JSON between the planner and the pool,
// counts what has finished, and gives the player a way to stop.
//
// Two shapes of the work are not the obvious ones, and both are load-bearing:
//
//   * a stage's requests go to `pool.run` directly. Contract 10.2 says a stage's request
//     array runs "as independent whole requests" -- `pool.run` IS that fan-out. A stage
//     request is one indivisible unit of work, so simSplit/simCombine around it would be
//     the identity and would push two extra messages per combination through worker 0 --
//     which is also a worker running combinations.
//   * a stage runs in chunks of the pool's own width. `pool.run` settles all-or-nothing and
//     aborts its siblings on the first failure, so one call per stage would throw away every
//     finished combination the moment the player pressed Stop. Chunking loses at most one
//     chunk, which is what makes "abort returns the partial" true rather than aspirational.
//
// `pool.count`, `pool.plan` and `pool.rank` are typed, not string-returning (worker.ts,
// task 4): their answers are consumed as the discriminated values they already are, never
// re-parsed as a second copy of JSON the pool parsed for us.
import { bulkCopy, simCopy } from './copy';
import {
  STAGES_BY_PRECISION,
  type BulkRequest,
  type BulkResult,
  type Combination,
  type Precision,
  type StageRequests,
  type WeightsRequest,
} from './bulk-types';
import type { SimResult } from './types';
import type { SimPool } from './worker';

/** The planner refused: the list is bigger than the lane allows, and by this much. */
export class BulkCapError extends Error {
  constructor(
    readonly cap: number,
    readonly combinations: number,
  ) {
    super(bulkCopy.capNotice(cap, combinations));
    this.name = 'BulkCapError';
  }
}

/**
 * A stop the player asked for, or an engine that could not run these combinations. `detail`
 * carries the engine's own words verbatim -- sim/request names the buff or consumable id it
 * refused, and that sentence is the only thing that says what to change.
 */
export class BulkRunError extends Error {
  constructor(
    message: string,
    readonly cancelled: boolean,
    readonly detail: string,
    options: { cause?: unknown } = {},
  ) {
    super(message, options);
    this.name = 'BulkRunError';
  }
}

export interface BulkProgress {
  stage: number;
  stages: number;
  combosDone: number;
  combosTotal: number;
}

export interface BulkRunHandle {
  readonly callbackId: string;
  readonly result: Promise<SimResult>;
  cancel(): void;
}

export function stageProgressLine(progress: BulkProgress): string {
  return bulkCopy.stageProgress(progress.stage, progress.stages, progress.combosDone, progress.combosTotal);
}

/** Runs of at most `size`, in order. A size of zero or less is one run of everything. */
export function chunk<T>(items: readonly T[], size: number): T[][] {
  if (items.length === 0) return [];
  if (size <= 0) return [[...items]];
  const out: T[][] = [];
  for (let i = 0; i < items.length; i += size) out.push(items.slice(i, i + size));
  return out;
}

function detailOf(cause: unknown): string {
  if (cause instanceof Error) return cause.message;
  return typeof cause === 'string' ? cause : '';
}

/**
 * The live combination count, for the run bar (`simCount`, contract 10.2). Throws
 * `BulkCapError` past the cap, because the button has to say what would exceed it and by
 * how much -- it never trims the list. `pool.count` already discriminates the cap refusal
 * from a genuine engine failure (the latter rejects the promise on its own).
 */
export async function countCombinations(pool: SimPool, request: BulkRequest): Promise<number> {
  const answer = await pool.count(JSON.stringify(request));
  if (!answer.ok) throw new BulkCapError(answer.cap, answer.combinations);
  return answer.combinations;
}

let bulkCounter = 0;

function nextCallbackId(): string {
  bulkCounter += 1;
  return `bulk-${bulkCounter}`;
}

/** The planner's first stage, or the cap refusal as `BulkCapError`. */
async function planStage(pool: SimPool, requestJSON: string): Promise<StageRequests> {
  const answer = await pool.plan(requestJSON);
  if (!answer.ok) throw new BulkCapError(answer.cap, answer.combinations);
  return answer.stage;
}

/**
 * What a stop has to hand back: the equipped run plus whatever combinations finished,
 * re-ranked by the wasm as if this were the final stage. The stage number is set to the
 * precision's last so `simRank` returns a `result` rather than another `next` -- ranking a
 * short list is still the planner's arithmetic, never this file's.
 */
function partialStage(stage: StageRequests, stages: number, finishedCombos: Combination[]): StageRequests {
  return {
    stage: stages,
    iterations: stage.iterations,
    requests: [stage.requests[0], ...finishedCombos.map((combo) => combo.request)],
    combos: finishedCombos,
    // The ladder's history travels with the stage (contract 10.1 A10), so a partial
    // result still says what each finished stage cost rather than starting the tally over.
    ran: stage.ran ?? [],
  };
}

export function runBulk(
  pool: SimPool,
  request: BulkRequest,
  onProgress: (progress: BulkProgress) => void,
  now: () => number = () => Date.now(),
): BulkRunHandle {
  const callbackId = nextCallbackId();
  // `BulkSpec.precision` is `string` on the wire (types.ts mirrors the Go tag loosely);
  // narrowed here the same way engine-fake.ts's own `LADDER` lookup does -- a request this
  // far along already passed `pool.plan`'s validation, which is where an unknown precision
  // would surface as a genuine engine error, not a silently wrong stage count.
  const stages = STAGES_BY_PRECISION[request.bulk.precision as Precision];
  const requestJSON = JSON.stringify(request);
  let cancelled = false;

  async function execute(): Promise<SimResult> {
    const startedAt = now();
    let stage = await planStage(pool, requestJSON);

    for (;;) {
      const combosTotal = stage.combos.length;
      onProgress({ stage: stage.stage, stages, combosDone: 0, combosTotal });

      // Index 0 is the equipped set and every later index is combos[index - 1]; the chunks
      // keep that order, so a partial rank can pair them back up by position.
      const results: string[] = [];
      let stopped = false;
      for (const part of chunk(stage.requests, pool.size)) {
        if (cancelled) {
          stopped = true;
          break;
        }
        const offset = results.length;
        const shardsJSON = part.map((r) => JSON.stringify(r));
        try {
          results.push(...(await pool.run(shardsJSON, `${callbackId}-s${stage.stage}-${offset}`, () => {})));
        } catch (cause) {
          if (!cancelled) throw cause;
          stopped = true;
          break;
        }
        onProgress({
          stage: stage.stage,
          stages,
          combosDone: Math.max(0, results.length - 1),
          combosTotal,
        });
      }

      if (stopped || cancelled) {
        // Nothing to hand back without the baseline: every delta is measured against it.
        if (results.length < 2) {
          throw new BulkRunError(simCopy.stopped, true, '');
        }
        const finished = stage.combos.slice(0, results.length - 1);
        const answer = await pool.rank(
          requestJSON,
          JSON.stringify(partialStage(stage, stages, finished)),
          `[${results.join(',')}]`,
        );
        const partial = answer.result as BulkResult | undefined;
        if (partial === undefined) throw new BulkRunError(simCopy.stopped, true, '');
        return finish(partial, startedAt, true);
      }

      const answer = await pool.rank(requestJSON, JSON.stringify(stage), `[${results.join(',')}]`);
      if (answer.result !== undefined) return finish(answer.result as BulkResult, startedAt, false);
      if (answer.next === undefined) throw new Error(bulkCopy.planFailed);
      stage = answer.next;
    }
  }

  /** The four facts the browser owns, not the engine, exactly as run.ts stamps them. */
  function finish(result: BulkResult, startedAt: number, aborted: boolean): SimResult {
    return {
      ...result,
      request,
      lane: 'browser',
      duration_ms: Math.round(now() - startedAt),
      ...(aborted ? { aborted: true } : {}),
    };
  }

  const result = execute().catch((cause: unknown) => {
    if (cause instanceof BulkCapError || cause instanceof BulkRunError) throw cause;
    throw new BulkRunError(
      cancelled ? simCopy.stopped : bulkCopy.bulkFailed,
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

/**
 * A stat-weights run. It is one wasm call rather than a stage loop -- the engine already
 * computes every weight in one pass -- so this is `runBulk`'s shape with the loop taken out.
 */
export function runWeightsRun(
  pool: SimPool,
  request: WeightsRequest,
  onTick: (iterationsDone: number) => void,
  now: () => number = () => Date.now(),
): BulkRunHandle {
  const callbackId = nextCallbackId();
  let cancelled = false;

  const result = (async (): Promise<SimResult> => {
    const startedAt = now();
    const json = await pool.weights(JSON.stringify(request), callbackId, (progress) =>
      onTick(progress.iterationsDone),
    );
    const finished = JSON.parse(json) as SimResult;
    return {
      ...finished,
      request,
      lane: 'browser',
      duration_ms: Math.round(now() - startedAt),
    };
  })().catch((cause: unknown) => {
    throw new BulkRunError(
      cancelled ? simCopy.stopped : bulkCopy.bulkFailed,
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
