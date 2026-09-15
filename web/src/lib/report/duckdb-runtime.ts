// web/src/lib/report/duckdb-runtime.ts
// The runtime half of scripts/duckdb-runtime.mjs: the two DuckDB engine modules are over
// Cloudflare's 25 MiB static-asset cap, so they are served out of the LOGS bucket by
// src/worker.ts instead of out of dist/. This module is the one place the route, the
// bucket keys and the version they both carry are written down, shared by the browser
// (query.ts) and the Worker (worker.ts).
//
// It deliberately holds nothing else: worker.ts imports it, so anything added here is
// added to the Worker bundle too.

/**
 * The pinned @duckdb/duckdb-wasm. duckdb-runtime.test.ts holds this to package.json's
 * exact pin, which is the version scripts/duckdb-runtime.mjs stages and uploads, so a
 * bump that forgets one half fails the unit tests rather than the Queries view.
 */
export const DUCKDB_WASM_VERSION = '1.32.0';

/** The Worker route. Everything under it is answered from the bucket or 404s. */
export const DUCKDB_RUNTIME_PREFIX = '/duckdb-runtime/';

/** The two files this route serves. Nothing else is reachable through it. */
export const DUCKDB_RUNTIME_MODULES = ['duckdb-eh.wasm', 'duckdb-mvp.wasm'] as const;

/** The static half: version-pinned under /duckdb/, so public/_headers can cache it for a year. */
export const DUCKDB_ASSET_PREFIX = `/duckdb/${DUCKDB_WASM_VERSION}`;

/** Where the module lives in the browser. Version-pinned, so it is immutable. */
export function duckdbRuntimeUrl(file: (typeof DUCKDB_RUNTIME_MODULES)[number]): string {
  return `${DUCKDB_RUNTIME_PREFIX}${DUCKDB_WASM_VERSION}/${file}`;
}

/**
 * The bucket key a request path names, or null for anything else under the prefix -- an
 * older version, a made-up filename, a traversal attempt. Null is a 404: these are two
 * known files at one known version, so there is nothing else to fall back to.
 */
export function duckdbRuntimeKey(pathname: string): string | null {
  if (!pathname.startsWith(DUCKDB_RUNTIME_PREFIX)) return null;
  const rest = pathname.slice(DUCKDB_RUNTIME_PREFIX.length);
  const file = DUCKDB_RUNTIME_MODULES.find((name) => rest === `${DUCKDB_WASM_VERSION}/${name}`);
  return file === undefined ? null : `runtime/duckdb/${DUCKDB_WASM_VERSION}/${file}`;
}
