// web/scripts/bis-source-build.mjs
// Which data/builds/<build>/bis/ directory serves the BiS lists for the active build.
//
// The nightly ranker publishes data/builds/<build>/bis/ AFTER a new client build is
// activated, so for up to a day the active build has no BiS files at all. Rather than
// blank every /bis page, the home upgrades panel and the planner's hover popover, the
// site reads the newest OTHER build that does have them. Each file records the build it
// was ranked on in its own `build` field, so the data still says "ranked on <build>"
// without a word of prose. The moment the active build gains its own files they win and
// this module returns the active build unchanged.
//
// Shared by the build-time loader (src/lib/bis/load.ts), scripts/sync-data.mjs and the
// e2e support (tests/e2e/support/bis-file.ts) so all three agree on one answer.
import { existsSync, readdirSync } from 'node:fs';
import path from 'node:path';

/** Every published BiS file is `<spec>.json`; the sibling `<spec>.md` reports do not count. */
const BIS_FILE_SUFFIX = '.json';

/** `data/builds/<build>/bis`. */
export function bisDirFor(repoRoot, build) {
  return path.join(repoRoot, 'data/builds', build, 'bis');
}

/** Whether `data/builds/<build>/bis/` holds at least one published BiS file. */
export function buildHasBisFiles(repoRoot, build) {
  const dir = bisDirFor(repoRoot, build);
  return existsSync(dir) && readdirSync(dir).some((name) => name.endsWith(BIS_FILE_SUFFIX));
}

/** Dotted client-build ids compare numerically, field by field (70291 > 70009 > 69893). */
export function compareBuildsDescending(a, b) {
  const left = a.split('.').map(Number);
  const right = b.split('.').map(Number);
  for (let i = 0; i < Math.max(left.length, right.length); i++) {
    const delta = (right[i] ?? 0) - (left[i] ?? 0);
    if (delta !== 0) return delta;
  }
  return 0;
}

/**
 * The build whose bis/ directory serves `activeBuild`: the active build itself when it has
 * BiS files, else the newest other build that has them, else the active build (so callers
 * fall through to their fixture / empty state exactly as before). A build with no
 * data/builds directory at all never falls back.
 * @param {string} repoRoot
 * @param {string} activeBuild
 * @returns {string}
 */
export function bisSourceBuild(repoRoot, activeBuild) {
  const buildsDir = path.join(repoRoot, 'data/builds');
  // An id with no data/builds directory is not a build awaiting its ranking: leave it alone.
  if (!existsSync(path.join(buildsDir, activeBuild)) || buildHasBisFiles(repoRoot, activeBuild)) {
    return activeBuild;
  }
  const newest = readdirSync(buildsDir)
    .filter((name) => name !== activeBuild && buildHasBisFiles(repoRoot, name))
    .sort(compareBuildsDescending)[0];
  return newest ?? activeBuild;
}
