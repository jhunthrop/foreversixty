// web/src/lib/site/removed-routes.test.ts
// L8.4 fix round 3 item 5: `npm run build`'s last postbuild step, check-no-removed-routes,
// failed on Windows with "dist/setup.html is missing" though the file was on disk. The script
// built its dist path from `new URL(...).pathname`, which on Windows is "/C:/.../dist/", and
// path.join turned that into "\C:\...\dist" -- a path that never exists, so every check read
// "absent". The check now resolves the directory with fileURLToPath and is a pure function of
// a directory, driven here against a temporary one.
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { pathToFileURL } from 'node:url';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { distDirOf, removedRouteFailures } from '../../../scripts/removed-routes.mjs';

let root: string;
let dist: string;

beforeEach(() => {
  root = mkdtempSync(path.join(tmpdir(), 'removed-routes-'));
  dist = path.join(root, 'dist');
  mkdirSync(dist);
});

afterEach(() => {
  rmSync(root, { recursive: true, force: true });
});

describe('distDirOf', () => {
  it('resolves dist/ beside scripts/ to a path the file system can open, on Windows too', () => {
    const script = pathToFileURL(path.join(root, 'scripts', 'check-no-removed-routes.mjs'));
    expect(path.resolve(distDirOf(script))).toBe(path.resolve(dist));
    writeFileSync(path.join(dist, 'setup.html'), '');
    expect(removedRouteFailures(distDirOf(script))).toEqual([]);
  });
});

describe('removedRouteFailures', () => {
  it('names a built /pricing page: the edge 301 owns that path', () => {
    writeFileSync(path.join(dist, 'setup.html'), '');
    writeFileSync(path.join(dist, 'pricing.html'), '');
    expect(removedRouteFailures(dist)).toEqual(['dist/pricing.html still exists']);
  });

  it('passes a dist with /setup and none of the removed routes', () => {
    writeFileSync(path.join(dist, 'setup.html'), '');
    writeFileSync(path.join(dist, 'about.html'), '');
    expect(removedRouteFailures(dist)).toEqual([]);
  });

  it('names a missing /setup', () => {
    expect(removedRouteFailures(dist)).toEqual([
      'dist/setup.html is missing -- /setup should replace /addon',
    ]);
  });

  it('names a removed page or a removed section that is still built', () => {
    writeFileSync(path.join(dist, 'setup.html'), '');
    writeFileSync(path.join(dist, 'addon.html'), '');
    mkdirSync(path.join(dist, 'zones'));
    writeFileSync(path.join(dist, 'zones', 'elwynn.html'), '');
    mkdirSync(path.join(dist, 'dungeons'));
    expect(removedRouteFailures(dist)).toEqual([
      'dist/addon.html still exists',
      'dist/zones/ still has files: elwynn.html',
    ]);
  });
});
