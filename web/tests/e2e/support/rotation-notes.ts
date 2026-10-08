// web/tests/e2e/support/rotation-notes.ts
// The player-facing rotation notes a spec's drawer and /sim/specs card print, read from the
// curated priority list the same way scripts/sync-rotations.mjs derives them, so a retuned
// rotation changes the expectation with the data instead of breaking a pinned sentence.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { rotationNotesOf } from '../../../scripts/rotation-apl.mjs';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');

type CuratedApl = Parameters<typeof rotationNotesOf>[0];

export function curatedRotationNotes(spec: string): string[] {
  const file = path.join(REPO_ROOT, 'data/curated/apl', `${spec}.json`);
  return rotationNotesOf(JSON.parse(readFileSync(file, 'utf8')) as CuratedApl, spec);
}
