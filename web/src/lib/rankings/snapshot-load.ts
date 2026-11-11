// web/src/lib/rankings/snapshot-load.ts
// Reads the build's rankings snapshot (lib/rankings/snapshot.ts) from src/data/generated, where
// scripts/sync-rankings.mjs wrote it. Node only: the rankings pages call it at build time. A build
// with no snapshot file (the sync has not run) draws the boards the way it always did, from the API
// alone; a file that is there but malformed is an error, never a silently empty board.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { EMPTY_SNAPSHOT, type RankingsSnapshot } from './snapshot';

// From the repo root, as lib/bis/load.ts reaches data/: this module runs from src/lib/rankings in
// the tests and from dist/.prerender/chunks in a build, both four levels under the root.
const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
export const SNAPSHOT_FILE = path.join(REPO_ROOT, 'web/src/data/generated/rankings-snapshot.json');

function isSnapshot(value: unknown): value is RankingsSnapshot {
  if (typeof value !== 'object' || value === null) return false;
  const candidate = value as Partial<RankingsSnapshot>;
  return (
    Array.isArray(candidate.encounters) &&
    typeof candidate.boards === 'object' &&
    candidate.boards !== null &&
    typeof candidate.phase === 'string'
  );
}

export function loadRankingsSnapshot(file: string = SNAPSHOT_FILE): RankingsSnapshot {
  if (!existsSync(file)) return EMPTY_SNAPSHOT;
  const parsed: unknown = JSON.parse(readFileSync(file, 'utf8'));
  if (!isSnapshot(parsed)) throw new Error(`${file} is not a rankings snapshot (scripts/sync-rankings.mjs)`);
  return parsed;
}
