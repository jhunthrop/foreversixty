import { describe, expect, it } from 'vitest';
import pkg from '../../../package.json';
import {
  DUCKDB_ASSET_PREFIX,
  DUCKDB_RUNTIME_MODULES,
  DUCKDB_RUNTIME_PREFIX,
  DUCKDB_WASM_VERSION,
  duckdbRuntimeKey,
  duckdbRuntimeUrl,
} from './duckdb-runtime';

describe('the DuckDB runtime route', () => {
  // The version is written in two places that cannot import each other: this module, which
  // the browser and the Worker read, and scripts/duckdb-runtime.mjs, which stages and
  // uploads the bytes off the installed package. package.json's pin is exact, so it is
  // what ties the two together -- and a bump that misses this constant would point the
  // browser at a key nothing was ever uploaded to.
  it('carries the exact version package.json pins', () => {
    expect(pkg.dependencies['@duckdb/duckdb-wasm']).toBe(DUCKDB_WASM_VERSION);
  });

  it('addresses each module under the version, so the bytes can be cached forever', () => {
    expect(duckdbRuntimeUrl('duckdb-eh.wasm')).toBe(`/duckdb-runtime/${DUCKDB_WASM_VERSION}/duckdb-eh.wasm`);
    expect(duckdbRuntimeUrl('duckdb-mvp.wasm')).toBe(
      `/duckdb-runtime/${DUCKDB_WASM_VERSION}/duckdb-mvp.wasm`,
    );
    expect(DUCKDB_ASSET_PREFIX).toBe(`/duckdb/${DUCKDB_WASM_VERSION}`);
  });

  it('maps each served path onto its bucket key', () => {
    for (const file of DUCKDB_RUNTIME_MODULES) {
      expect(duckdbRuntimeKey(duckdbRuntimeUrl(file))).toBe(`runtime/duckdb/${DUCKDB_WASM_VERSION}/${file}`);
    }
  });

  it('refuses every other path under the prefix', () => {
    const refused = [
      DUCKDB_RUNTIME_PREFIX,
      `${DUCKDB_RUNTIME_PREFIX}${DUCKDB_WASM_VERSION}/`,
      `${DUCKDB_RUNTIME_PREFIX}${DUCKDB_WASM_VERSION}/duckdb-coi.wasm`,
      `${DUCKDB_RUNTIME_PREFIX}1.31.0/duckdb-eh.wasm`,
      `${DUCKDB_RUNTIME_PREFIX}${DUCKDB_WASM_VERSION}/../../reports/fixture2abcd/report.json`,
      `${DUCKDB_RUNTIME_PREFIX}${DUCKDB_WASM_VERSION}/duckdb-eh.wasm/extra`,
      '/duckdb/duckdb-eh.wasm',
    ];
    for (const pathname of refused) expect(duckdbRuntimeKey(pathname)).toBeNull();
  });
});
