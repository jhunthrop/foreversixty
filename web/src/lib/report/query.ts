// web/src/lib/report/query.ts
// Deep queries: DuckDB-WASM over the fight's own events.parquet, in the browser.
//
// Three rules, from spec sections 1 and 4:
//   never on page load    the import is dynamic and only the Queries view calls load()
//   from our own origin   the worker JS and the extension out of public/duckdb/, the two
//                         engine modules out of the LOGS bucket through the Worker's
//                         /duckdb-runtime/ route (they are over Cloudflare's 25 MiB
//                         static-asset cap; see scripts/duckdb-runtime.mjs). The
//                         package's getJsDelivrBundles() helper is deliberately unused
//   once per fight        the Parquet is fetched and registered once, then every brush,
//                         filter and drill-down is answered from memory
//
// The engine is an interface so the query layer is testable without a browser: DuckDB's
// Node build needs a web-worker shim and does not start under vitest. The real engine is
// exercised by tests/e2e/report-queries.spec.ts, in a browser, against the checked-in
// fixture Parquet.
import { DUCKDB_ASSET_PREFIX, duckdbRuntimeUrl } from './duckdb-runtime';
import type { TimeWindow } from './window';

/** The name the fight's bytes are registered under inside DuckDB. */
export const EVENTS_FILE = 'events.parquet';

/** The one table every query reads. The file is registered under this exact name. */
export const EVENTS_TABLE = `read_parquet('${EVENTS_FILE}')`;

/**
 * Milliseconds from the fight's start, which is the clock every other number on this page
 * runs on: the brush, the window presets, `?start=` and `?end=`, a death's `at_ms`, a
 * cast's `sequence`.
 *
 * The Parquet does not store that clock. `time_unix_nano`
 * (logs/engine/parquet/schema.go) is an absolute instant, so comparing it to a window
 * bound directly would quietly answer nothing at all. The fight's zero is its first
 * event: logs/engine/fight/fight.go's `newFight` takes `Start: e.Time` from the event it
 * is opened by and assigns that same event to the fight, so the minimum over the fight's
 * own file is exactly the instant every `_ms` in summary.json counts from.
 */
export const FIGHT_MS = `(time_unix_nano - (SELECT min(time_unix_nano) FROM ${EVENTS_TABLE})) / 1000000`;

/**
 * logs/engine/units's `NoGUID`: the placeholder the game writes in the advanced block
 * when the caster has no owner. A pet filter that tested `adv_owner_guid <> ''` alone
 * would match every player in the fight, because every one of them carries this.
 */
export const NO_OWNER_GUID = '0000000000000000';

export interface QueryResult {
  columns: string[];
  rows: unknown[][];
  elapsedMs: number;
  /**
   * How many rows the answer really had, when that is more than `rows` carries. Building
   * a million-row array in the main thread is how a query box freezes a tab, so the
   * engine materialises the first `MAX_ROWS` and says so rather than pretending.
   */
  total?: number;
}

export interface QueryEngine {
  /** Registers the bytes under `name` so SQL can read_parquet() them. */
  open(name: string, bytes: Uint8Array): Promise<void>;
  query(sql: string): Promise<QueryResult>;
  close(): Promise<void>;
}

export interface QueryTemplate {
  id: string;
  label: string;
  sql: string;
}

/**
 * The starting points. Every one of them answers something the summary cannot: a
 * per-ability split inside a window, a pet's own damage, the exact miss breakdown.
 * `:start` and `:end` are substituted with the current window by withWindow().
 *
 * Unlike the tables', these figures carry no `~`: window.ts scales a per-ability split by
 * the window's share of the actor's total, while a query measures the window itself.
 */
export const QUERY_TEMPLATES: readonly QueryTemplate[] = [
  {
    id: 'damage-by-ability',
    label: 'Damage by ability, exactly, in this window',
    sql: `SELECT source_name, spell_name, count(*) AS hits, sum(amount) AS total
FROM ${EVENTS_TABLE}
WHERE kind = 'damage' AND ${FIGHT_MS} BETWEEN :start AND :end
GROUP BY 1, 2
ORDER BY total DESC`,
  },
  {
    id: 'pet-damage',
    label: 'Pet damage, separated from its owner',
    sql: `SELECT source_name, adv_owner_guid, sum(amount) AS total
FROM ${EVENTS_TABLE}
WHERE kind = 'damage' AND adv_owner_guid NOT IN ('', '${NO_OWNER_GUID}')
  AND ${FIGHT_MS} BETWEEN :start AND :end
GROUP BY 1, 2
ORDER BY total DESC`,
  },
  {
    id: 'misses',
    label: 'Every miss, by type',
    sql: `SELECT dest_name, miss_type, count(*) AS n
FROM ${EVENTS_TABLE}
WHERE miss_type <> '' AND ${FIGHT_MS} BETWEEN :start AND :end
GROUP BY 1, 2
ORDER BY n DESC`,
  },
  {
    id: 'overheal',
    label: 'Healing and overhealing by spell',
    sql: `SELECT source_name, spell_name, sum(amount) AS healing, sum(overheal) AS overheal
FROM ${EVENTS_TABLE}
WHERE kind = 'heal' AND ${FIGHT_MS} BETWEEN :start AND :end
GROUP BY 1, 2
ORDER BY healing DESC`,
  },
  {
    id: 'event-counts',
    label: 'Which events this fight contains',
    sql: `SELECT event, count(*) AS n
FROM ${EVENTS_TABLE}
GROUP BY 1
ORDER BY n DESC`,
  },
];

export const NOT_A_READ = 'Queries here read the fight; they cannot change it.';
export const QUERY_TOO_SLOW =
  'That query ran for too long and was stopped. Narrow it with a WHERE or a LIMIT.';
export const QUERY_ABANDONED = 'That query was dropped when the view moved on.';
export const NOT_LOCAL = 'Queries here read this fight’s own file; they cannot fetch from another address.';

/**
 * Where scripts/sync-duckdb.mjs publishes the vendored DuckDB extensions: under the same
 * version-pinned prefix as the worker JS, so public/_headers can cache the lot for a year.
 * DuckDB builds `<repository>/<version>/<platform>/<name>.duckdb_extension.wasm` itself,
 * so this is the base and nothing else.
 */
export const EXTENSION_REPOSITORY = `${DUCKDB_ASSET_PREFIX}/extensions`;

/** A query that has not answered by then has its engine torn down and rebuilt. */
export const QUERY_TIMEOUT_MS = 30_000;

/**
 * How long a graceful shutdown is given before the worker is killed outright.
 *
 * DuckDB's own `connection.close()` posts a DISCONNECT message and waits for the worker
 * to answer it, and the worker answers on the same single-threaded message loop the query
 * is running on -- so the one shutdown that has to work, the one after a query wedged the
 * worker, is exactly the one that would hang forever. Past this the worker is terminated.
 */
export const ENGINE_CLOSE_TIMEOUT_MS = 2_000;

/**
 * How long the layer waits for an engine to shut down before it stops caring. Longer than
 * ENGINE_CLOSE_TIMEOUT_MS, so the real engine always finishes first and this only bites a
 * `load()` that returns something misbehaving: the point is that close() and the run
 * queue behind it can never be held hostage by an engine that will not die.
 */
export const TEARDOWN_TIMEOUT_MS = 5_000;

/** The most rows one answer materialises; `total` carries the real count. */
export const MAX_ROWS = 200;

/**
 * Waits for `work`, but never longer than `ms`, and never rejects. Used only for shutdown,
 * where the useful outcome is "stop waiting" rather than "report what went wrong": the
 * caller's next line kills the worker regardless.
 */
function withDeadline(work: Promise<unknown>, ms: number): Promise<void> {
  let timer: ReturnType<typeof setTimeout> | undefined;
  const deadline = new Promise<void>((resolve) => {
    timer = setTimeout(resolve, ms);
  });
  return Promise.race([
    work.then(
      () => undefined,
      () => undefined,
    ),
    deadline,
  ]).finally(() => clearTimeout(timer));
}

/**
 * One statement, and it must be a SELECT or a WITH. This is not a security boundary --
 * everything runs in the visitor's own tab over a file they can already download -- it is
 * a guardrail: DuckDB can COPY to a file and INSTALL extensions, and neither belongs in a
 * box labelled "ask a question about this fight".
 */
export function isReadOnlySql(sql: string): boolean {
  const trimmed = sql.trim().replace(/;+\s*$/, '');
  if (trimmed.includes(';')) return false;
  return /^(select|with)\b/i.test(trimmed);
}

/**
 * No address but this fight's own file.
 *
 * `read_parquet('https://…')` is the one way a query can still reach off this origin: it
 * makes DuckDB autoload `httpfs`, which is not among the extensions vendored into
 * public/duckdb/extensions, and then fetch from wherever the URL points. Neither belongs
 * on a site whose rule is that every byte comes from foreversixty.gg, so a query carrying
 * any URL scheme is refused before the engine is even built. The events file is already
 * registered in DuckDB's own filesystem under a bare name, so nothing legitimate here
 * needs one.
 */
export function isLocalOnlySql(sql: string): boolean {
  return !/[a-z][a-z0-9+.-]*:\/\//i.test(sql);
}

/** The one gate run() applies: the refusal to show, or null when the query may run. */
export function refuseSql(sql: string): string | null {
  if (!isReadOnlySql(sql)) return NOT_A_READ;
  if (!isLocalOnlySql(sql)) return NOT_LOCAL;
  return null;
}

export function withWindow(sql: string, window: TimeWindow): string {
  return sql.replaceAll(':start', String(window.startMs)).replaceAll(':end', String(window.endMs));
}

/** Arrow marks its 128-bit views with this shared symbol; see normalizeCell. */
const ARROW_BIG_NUM = Symbol.for('isArrowBigNum');

/**
 * One Arrow cell as something a table cell and JSON can both hold.
 *
 * Two shapes need the help. A 64-bit column arrives as a `bigint`, which `JSON.stringify`
 * refuses outright. A 128-bit one -- which is what `sum()` over a BIGINT produces --
 * arrives as Arrow's BigNum, a view over four 32-bit words whose default string form is
 * "1,2,0,0"; its own `toString` is the exact decimal.
 */
export function normalizeCell(value: unknown): unknown {
  if (value === null || value === undefined) return null;
  if (typeof value === 'bigint') {
    return Number.isSafeInteger(Number(value)) ? Number(value) : value.toString();
  }
  if (typeof value === 'string' || typeof value === 'number' || typeof value === 'boolean') {
    return value;
  }
  if (value instanceof Date) return value.toISOString();
  if (typeof value === 'object' && ARROW_BIG_NUM in value) return String(value);
  try {
    return JSON.stringify(value, (_key, item: unknown) =>
      typeof item === 'bigint' ? item.toString() : item,
    );
  } catch {
    return String(value);
  }
}

export interface QueryLayerOptions {
  load: () => Promise<QueryEngine>;
  fetchBytes: (url: string) => Promise<Uint8Array>;
  /** Overridden only by tests; production uses QUERY_TIMEOUT_MS. */
  timeoutMs?: number;
  /** Overridden only by tests; production uses TEARDOWN_TIMEOUT_MS. */
  teardownMs?: number;
}

export interface QueryLayer {
  /** True once the engine exists and a fight's events are registered. */
  readonly ready: boolean;
  run(eventsUrl: string, sql: string): Promise<QueryResult>;
  close(): Promise<void>;
}

/**
 * The lazy, single-instance owner of the engine.
 *
 * Everything it does is awaited and slow, so three things are deliberate rather than
 * incidental:
 *
 *   one engine      `loading` memoises the promise, not the engine, so two clicks that
 *                   start together share the build instead of racing into two live DuckDB
 *                   workers -- and only one of those two would ever be terminated.
 *   one at a time   runs are chained, so a fight switch between two runs cannot interleave
 *                   a registration of fight A's bytes with a query meant for fight B.
 *   nothing outlives close()  `generation` invalidates every run already in flight, and
 *                   close() awaits a build still in progress so the engine that arrives
 *                   after the view is gone is still terminated.
 */
export function createQueryLayer(options: QueryLayerOptions): QueryLayer {
  const timeoutMs = options.timeoutMs ?? QUERY_TIMEOUT_MS;
  const teardownMs = options.teardownMs ?? TEARDOWN_TIMEOUT_MS;
  let engine: QueryEngine | null = null;
  let loading: Promise<QueryEngine> | null = null;
  let openedUrl = '';
  let generation = 0;
  let tail: Promise<unknown> = Promise.resolve();

  const ignore = (): undefined => undefined;

  /** Runs jobs one after another, whether the one before it kept or broke its promise. */
  function serial<T>(job: () => Promise<T>): Promise<T> {
    const next = tail.then(job, job);
    tail = next.then(ignore, ignore);
    return next;
  }

  function engineOnce(): Promise<QueryEngine> {
    if (loading !== null) return loading;
    const attempt: Promise<QueryEngine> = options.load().then(
      (built) => {
        engine = built;
        return built;
      },
      (thrown: unknown) => {
        // Cleared, because a rejected promise stays rejected and `??=` would never
        // replace it. The download that fails here is the thirty-odd megabyte engine, the
        // one a flaky connection actually drops, and without this every later Run would
        // answer with the same stale error and no network activity at all. Only if it is
        // still the current attempt: a teardown may already have started a newer one.
        if (loading === attempt) loading = null;
        throw thrown;
      },
    );
    loading = attempt;
    return attempt;
  }

  async function teardown(): Promise<void> {
    generation += 1;
    const pending = loading;
    loading = null;
    engine = null;
    openedUrl = '';
    if (pending === null) return;
    // Awaited, not dropped: a build still in flight would otherwise resolve into a live
    // worker with nobody left holding a reference to terminate it. Bounded, because this
    // is awaited from inside the run queue on the timeout path: an engine that will not
    // shut down must not also stop every later query from starting. Giving up on the
    // await does not abandon the shutdown -- createDuckDbEngine's close() kills the
    // worker in a `finally` whether anyone is still listening or not.
    await withDeadline(
      pending.then((built) => built.close(), ignore),
      teardownMs,
    );
  }

  /** Throws if the layer was closed, or timed out, while the last await was outstanding. */
  function stillWanted(mine: number): void {
    if (mine !== generation) throw new Error(QUERY_ABANDONED);
  }

  async function execute(eventsUrl: string, sql: string, mine: number): Promise<QueryResult> {
    const active = await engineOnce();
    stillWanted(mine);
    if (openedUrl !== eventsUrl) {
      const bytes = await options.fetchBytes(eventsUrl);
      stillWanted(mine);
      await active.open(EVENTS_FILE, bytes);
      stillWanted(mine);
      openedUrl = eventsUrl;
    }
    const answer = await active.query(sql);
    stillWanted(mine);
    return answer;
  }

  /**
   * A hostile or merely careless query -- a cross join over every event -- can run inside
   * the worker for longer than anyone will wait, and the worker owns its own thread, so
   * nothing on the page can interrupt it. The timeout is the recovery: the run rejects
   * with something legible and the worker is terminated, so the next Run starts a fresh
   * engine rather than queueing behind a thread that is never coming back.
   */
  function withTimeout(work: Promise<QueryResult>): Promise<QueryResult> {
    let timer: ReturnType<typeof setTimeout> | undefined;
    const limit = new Promise<never>((_resolve, reject) => {
      timer = setTimeout(() => reject(new Error(QUERY_TOO_SLOW)), timeoutMs);
    });
    return Promise.race([work, limit]).finally(() => clearTimeout(timer));
  }

  return {
    get ready(): boolean {
      return engine !== null && openedUrl !== '';
    },
    run(eventsUrl: string, sql: string): Promise<QueryResult> {
      // Ahead of the queue on purpose: a statement that will never be allowed to run
      // should not wait behind a slow one to be told so, and nothing is built for it.
      const refusal = refuseSql(sql);
      if (refusal !== null) return Promise.reject(new Error(refusal));
      // Read now, not inside the job: the job is a microtask away, and a close() in
      // between -- the visitor leaving the Queries view the instant after pressing Run --
      // would otherwise be read as having happened before the run, and the layer would
      // build a whole engine for a component that is already gone.
      const mine = generation;
      return serial(async () => {
        try {
          stillWanted(mine);
          return await withTimeout(execute(eventsUrl, sql, mine));
        } catch (thrown) {
          if (thrown instanceof Error && thrown.message === QUERY_TOO_SLOW) await teardown();
          throw thrown;
        }
      });
    },
    close(): Promise<void> {
      return teardown();
    },
  };
}

/**
 * The real engine. The import is dynamic, so nothing here is in report-island.js: Vite
 * splits it into its own chunk that is fetched the first time someone opens Queries.
 */
export async function createDuckDbEngine(): Promise<QueryEngine> {
  const duckdb = await import('@duckdb/duckdb-wasm');

  // Safari versions still in use have no WebAssembly exceptions, so both builds are
  // published and the runtime picks. Both come from foreversixty.gg, never from a CDN: the
  // package's own getJsDelivrBundles() is what the site's "no third-party bytes" rule
  // exists to rule out. The worker is a static asset; the module is too large to be one
  // and is served from the LOGS bucket by src/worker.ts.
  const features = await duckdb.getPlatformFeatures();
  const bundle = features.wasmExceptions
    ? {
        module: duckdbRuntimeUrl('duckdb-eh.wasm'),
        worker: `${DUCKDB_ASSET_PREFIX}/duckdb-browser-eh.worker.js`,
      }
    : {
        module: duckdbRuntimeUrl('duckdb-mvp.wasm'),
        worker: `${DUCKDB_ASSET_PREFIX}/duckdb-browser-mvp.worker.js`,
      };

  // A plain string, not `new URL(..., import.meta.url)`: the second form asks Vite to
  // bundle a worker of its own, which would put DuckDB back inside the island's chunk.
  const worker = new Worker(bundle.worker, { type: 'classic' });
  const database = new duckdb.AsyncDuckDB(new duckdb.VoidLogger(), worker);
  await database.instantiate(bundle.module);
  const connection = await database.connect();

  // The one third-party request DuckDB makes on its own. Since 1.4 the Parquet reader is
  // not compiled into the engine -- it is a "known extension" the first read_parquet() in
  // a session fetches from https://extensions.duckdb.org -- which would break the site's
  // rule that every byte comes from foreversixty.gg, in the one view whose whole point is
  // that it does not. Turning autoload off instead just breaks the queries, so the signed
  // extension is vendored and served from here: DuckDB appends /<version>/<platform>/
  // <name>.duckdb_extension.wasm to whatever repository it is given, and
  // scripts/sync-duckdb.mjs publishes exactly that shape under the extension repository.
  //
  // location.origin, not a bare path: the setting is a repository URL, and the report
  // island runs on foreversixty.gg, on localhost under Playwright, and on a preview
  // deployment, so the origin has to come from the page rather than a constant.
  await connection.query(`SET custom_extension_repository='${location.origin}${EXTENSION_REPOSITORY}'`);

  return {
    async open(name: string, bytes: Uint8Array): Promise<void> {
      // Dropped first so switching fight replaces the registration rather than stacking a
      // second copy of two to ten megabytes in the WASM heap. The first fight has nothing
      // to drop, which is not an error worth surfacing.
      await database.dropFile(name).catch(() => null);
      await database.registerFileBuffer(name, bytes);
    },
    async query(sql: string): Promise<QueryResult> {
      const started = performance.now();
      const table = await connection.query(sql);
      const columns = table.schema.fields.map((field) => field.name);
      // Read positionally, by child vector, rather than through a row object keyed by
      // column name: two output columns can share a name, and a keyed row would drop one.
      const children = columns.map((_name, index) => table.getChildAt(index));
      const shown = Math.min(table.numRows, MAX_ROWS);
      const rows: unknown[][] = [];
      for (let index = 0; index < shown; index += 1) {
        rows.push(children.map((child) => normalizeCell(child?.get(index))));
      }
      return { columns, rows, elapsedMs: performance.now() - started, total: table.numRows };
    },
    async close(): Promise<void> {
      // The graceful half is best-effort and on a deadline. connection.close() posts
      // DISCONNECT and waits for the worker to answer on the same message loop a running
      // query occupies, so after a query wedged the worker -- the one case this path
      // exists to recover from -- it never returns. The kill is in a `finally` so it
      // happens on every route out of here, including that one.
      try {
        await withDeadline(connection.close(), ENGINE_CLOSE_TIMEOUT_MS);
      } finally {
        // AsyncDuckDB.terminate() is what actually calls worker.terminate(); the second
        // call is the belt for a database that has already forgotten its worker.
        await database.terminate().catch(() => undefined);
        worker.terminate();
      }
    },
  };
}
