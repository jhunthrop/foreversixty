// web/src/lib/site/chunk-css.test.ts
// The postbuild guard (scripts/check-chunk-css.mjs) over a fixture dist: a chunk whose preload
// list names a stylesheet the build did not emit fails; one whose stylesheet exists passes.
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { missingChunkCss } from '../../../scripts/chunk-css.mjs';

let dist: string;

beforeEach(() => {
  dist = mkdtempSync(path.join(tmpdir(), 'chunk-css-'));
  mkdirSync(path.join(dist, '_astro'));
});

afterEach(() => {
  rmSync(dist, { recursive: true, force: true });
});

const chunk = (name: string, deps: readonly string[]): void =>
  writeFileSync(
    path.join(dist, '_astro', name),
    `const __vite__mapDeps=(i,m=__vite__mapDeps,d=(m.f||(m.f=${JSON.stringify(deps)})))=>i.map(i=>d[i]);`,
  );

describe('missingChunkCss', () => {
  it('names a preloaded stylesheet the build did not emit', () => {
    chunk('LogsLanding.abc.js', ['_astro/Account.def.js', '_astro/Account.xyz.css']);
    expect(missingChunkCss(dist)).toEqual([
      '_astro/LogsLanding.abc.js preloads _astro/Account.xyz.css, which the build did not emit',
    ]);
  });

  it('passes a chunk whose stylesheets exist, and one that preloads none', () => {
    writeFileSync(path.join(dist, '_astro', 'Report.xyz.css'), '.a{}');
    chunk('ReportView.abc.js', ['_astro/Report.def.js', '_astro/Report.xyz.css']);
    chunk('Plain.abc.js', ['_astro/Plain.def.js']);
    expect(missingChunkCss(dist)).toEqual([]);
  });

  it('passes a dist with no chunks at all', () => {
    rmSync(path.join(dist, '_astro'), { recursive: true });
    expect(missingChunkCss(dist)).toEqual([]);
  });
});
