import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { fixtureResult, fixtureSpecs } from '../../test-support/sim-api';
import { ENGINE_VERSION, engineAssetUrl, engineLabel, isStale } from './version';

// The engine lane's Go module is not on main yet -- this tree carries only the data lane's
// sim/specs -- so the pin assertion runs when the file lands and reports as skipped until
// then. A skip that names its own condition is honest; a test that quietly passes on a
// missing file is the thing this assertion exists to prevent.
const ENGINEVER_GO = path.resolve(import.meta.dirname, '../../../../sim/enginever/version.go');
const enginePinned = existsSync(ENGINEVER_GO);

describe('ENGINE_VERSION', () => {
  it('is a short git sha', () => {
    expect(ENGINE_VERSION).toMatch(/^[0-9a-f]{7,12}$/);
  });

  it.skipIf(!enginePinned)(
    'is the same sha sim/enginever/version.go carries, because make engine-pin writes both',
    () => {
      const pinned = /const Version = "([0-9a-f]{7,12})"/.exec(readFileSync(ENGINEVER_GO, 'utf8'));
      expect(pinned).not.toBeNull();
      expect(ENGINE_VERSION).toBe(pinned![1]);
    },
  );
});

describe('engineAssetUrl', () => {
  it('addresses the immutable per-version directory', () => {
    expect(engineAssetUrl('sim.wasm')).toBe(`/_sim/${ENGINE_VERSION}/sim.wasm`);
    expect(engineAssetUrl('sim.js')).toBe(`/_sim/${ENGINE_VERSION}/sim.js`);
  });

  it('can address an older version, so a stored result stays readable', () => {
    expect(engineAssetUrl('sim.wasm', '6a1c2d9')).toBe('/_sim/6a1c2d9/sim.wasm');
  });
});

describe('isStale', () => {
  it('is false for the current engine and true for any other', () => {
    expect(isStale(ENGINE_VERSION)).toBe(false);
    expect(isStale('6a1c2d9')).toBe(true);
  });

  it('treats a missing version as stale rather than current', () => {
    expect(isStale('')).toBe(true);
  });
});

describe('engineLabel', () => {
  it('reads as a version, not as a hash', () => {
    expect(engineLabel('6a1c2d9')).toBe('Engine 6a1c2d9');
  });
});

describe('the fixtures', () => {
  it('are pinned to the current engine, since JSON cannot import the constant', () => {
    expect(fixtureResult.engine_version).toBe(ENGINE_VERSION);
    expect(fixtureResult.request.engine_version).toBe(ENGINE_VERSION);
    expect(fixtureResult.summary.engine_version).toBe(`sim:${ENGINE_VERSION}`);
    for (const row of fixtureSpecs) {
      if (row.engine_version !== '') expect(row.engine_version).toBe(ENGINE_VERSION);
    }
  });
});
