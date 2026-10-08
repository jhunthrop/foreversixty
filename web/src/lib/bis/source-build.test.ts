// web/src/lib/bis/source-build.test.ts
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { bisSourceBuild, compareBuildsDescending } from '../../../scripts/bis-source-build.mjs';

let repoRoot: string;

function addBuild(build: string, bisFiles: string[] = []): void {
  const dir = path.join(repoRoot, 'data/builds', build);
  mkdirSync(dir, { recursive: true });
  if (bisFiles.length === 0) return;
  mkdirSync(path.join(dir, 'bis'));
  for (const name of bisFiles) writeFileSync(path.join(dir, 'bis', name), '{}');
}

beforeEach(() => {
  repoRoot = mkdtempSync(path.join(tmpdir(), 'bis-source-'));
});

afterEach(() => {
  rmSync(repoRoot, { recursive: true, force: true });
});

describe('bisSourceBuild', () => {
  it('keeps the active build when it has its own BiS files', () => {
    addBuild('1.60.1.70009', ['hunter-marksmanship.json']);
    addBuild('1.60.1.70291', ['hunter-marksmanship.json']);
    expect(bisSourceBuild(repoRoot, '1.60.1.70291')).toBe('1.60.1.70291');
  });

  it('falls back to the newest other build with BiS files while the active build has none', () => {
    addBuild('1.60.1.69893', ['a.json']);
    addBuild('1.60.1.70009', ['a.json']);
    addBuild('1.60.1.70291');
    expect(bisSourceBuild(repoRoot, '1.60.1.70291')).toBe('1.60.1.70009');
  });

  it('does not count the markdown reports beside the JSON files', () => {
    addBuild('1.60.1.70009', ['a.json']);
    addBuild('1.60.1.70291', ['a.md']);
    expect(bisSourceBuild(repoRoot, '1.60.1.70291')).toBe('1.60.1.70009');
  });

  it('returns the active build when no build has BiS files', () => {
    addBuild('1.60.1.70291');
    expect(bisSourceBuild(repoRoot, '1.60.1.70291')).toBe('1.60.1.70291');
  });

  it('never falls back for an id that is not a build directory', () => {
    addBuild('1.60.1.70009', ['a.json']);
    expect(bisSourceBuild(repoRoot, 'no-such-build')).toBe('no-such-build');
  });
});

describe('compareBuildsDescending', () => {
  it('orders dotted build ids numerically, newest first', () => {
    const ordered = ['1.60.1.69893', '1.60.1.70291', '1.15.9.69722', '1.60.1.70009'].sort(
      compareBuildsDescending,
    );
    expect(ordered).toEqual(['1.60.1.70291', '1.60.1.70009', '1.60.1.69893', '1.15.9.69722']);
  });
});
