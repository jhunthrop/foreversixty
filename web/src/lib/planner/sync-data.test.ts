import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { existsSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { syncData } from '../../../scripts/sync-data.mjs';

const BUILD = '1.15.9.69722';
const silent = { log: () => {}, warn: () => {} };

let root: string;
let repoRoot: string;
let webRoot: string;

function writeJson(file: string, value: unknown): void {
  mkdirSync(path.dirname(file), { recursive: true });
  writeFileSync(file, `${JSON.stringify(value, null, 2)}\n`, 'utf8');
}

/** A web/ tree with active-build.json and a fixture build, matching the real layout. */
function scaffoldWeb(): void {
  writeJson(path.join(webRoot, 'src/data/active-build.json'), { build: BUILD });
  const fixture = path.join(webRoot, 'src/fixtures/planner');
  writeJson(path.join(fixture, 'talents/warrior.json'), { build: BUILD, class_slug: 'warrior', trees: [] });
  writeJson(path.join(fixture, 'items/warrior.json'), { build: BUILD, class_slug: 'warrior', items: [] });
  writeJson(path.join(fixture, 'sets.json'), []);
  writeJson(path.join(fixture, 'classes.json'), [{ id: 1, slug: 'warrior' }]);
  writeJson(path.join(fixture, 'races.json'), [{ id: 1, slug: 'human' }]);
  writeJson(path.join(fixture, 'combos.json'), [{ race_id: 1, class_id: 1, new_in_forever: false }]);
  mkdirSync(path.join(fixture, 'icons'), { recursive: true });
  writeFileSync(path.join(fixture, 'icons/fixture_a.webp'), 'x');
  writeJson(path.join(fixture, 'manifest.json'), { build: BUILD, fixture: true, files: {} });
}

/** A data/builds/<build>/ tree carrying real Phase 1 files. */
function scaffoldRealBuild(): string {
  const dir = path.join(repoRoot, 'data/builds', BUILD);
  writeJson(path.join(dir, 'talents/paladin.json'), { build: BUILD, class_slug: 'paladin', trees: [] });
  writeJson(path.join(dir, 'classes.json'), [{ id: 2, slug: 'paladin' }]);
  writeJson(path.join(dir, 'races.json'), [{ id: 5, slug: 'undead' }]);
  writeJson(path.join(dir, 'combos.json'), [{ race_id: 5, class_id: 2, new_in_forever: true }]);
  writeJson(path.join(dir, 'manifest.json'), {
    build: BUILD,
    files: {
      'talents/paladin.json': 'aa',
      'classes.json': 'bb',
      'races.json': 'cc',
      'combos.json': 'dd',
    },
  });
  return dir;
}

beforeEach(() => {
  root = mkdtempSync(path.join(tmpdir(), 'sync-data-'));
  repoRoot = path.join(root, 'repo');
  webRoot = path.join(repoRoot, 'web');
  scaffoldWeb();
});

afterEach(() => rmSync(root, { recursive: true, force: true }));

describe('syncData', () => {
  it('falls back to the fixture when the build has no Phase 1 files', async () => {
    const result = await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(result.usedFixture).toBe(true);
    expect(result.build).toBe(BUILD);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/warrior.json'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'icons/fixture_a.webp'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'combos.json'))).toBe(true);
  });

  it('writes the three page-imported files into src/data/generated', async () => {
    await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    for (const name of ['classes.json', 'races.json', 'combos.json']) {
      const file = path.join(webRoot, 'src/data/generated', name);
      expect(existsSync(file)).toBe(true);
      expect(() => JSON.parse(readFileSync(file, 'utf8'))).not.toThrow();
    }
  });

  it('refuses the fixture fallback when fixtures are not allowed', async () => {
    await expect(syncData({ repoRoot, webRoot, allowFixture: false, log: silent })).rejects.toThrow(
      /no Phase 1 talent data/,
    );
  });

  it('copies the real build when the manifest lists Phase 1 files', async () => {
    scaffoldRealBuild();
    const result = await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(result.usedFixture).toBe(false);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/paladin.json'))).toBe(true);
    expect(existsSync(path.join(webRoot, 'public/data', BUILD, 'talents/warrior.json'))).toBe(false);
  });

  it('fails loudly and names every manifest file missing from disk', async () => {
    const dir = scaffoldRealBuild();
    rmSync(path.join(dir, 'races.json'));
    await expect(syncData({ repoRoot, webRoot, allowFixture: true, log: silent })).rejects.toThrow(
      /manifest\.json lists files that are missing: races\.json/,
    );
  });

  it('fails when a required file is absent from the source directory', async () => {
    const dir = scaffoldRealBuild();
    rmSync(path.join(dir, 'combos.json'));
    writeJson(path.join(dir, 'manifest.json'), {
      build: BUILD,
      files: { 'talents/paladin.json': 'aa', 'classes.json': 'bb', 'races.json': 'cc' },
    });
    await expect(syncData({ repoRoot, webRoot, allowFixture: true, log: silent })).rejects.toThrow(
      /required entry combos\.json is missing/,
    );
  });

  it('replaces a stale public/data directory instead of merging into it', async () => {
    const stale = path.join(webRoot, 'public/data', BUILD, 'talents/stale.json');
    writeJson(stale, { gone: true });
    await syncData({ repoRoot, webRoot, allowFixture: true, log: silent });
    expect(existsSync(stale)).toBe(false);
  });
});
