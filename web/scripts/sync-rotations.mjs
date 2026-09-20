// web/scripts/sync-rotations.mjs
// data/curated/apl/<spec>.json -> web/src/data/generated/rotations.json, before dev, check,
// test and build. Mirrors sync-sim-ids.mjs's shape: the target directory is gitignored like
// everything else under src/data/generated/, because derived data is not committed.
//
// A missing data/curated/apl/ directory is fatal: that means the sync is misconfigured (the
// curated rotation lane's own layout moved), not that a spec has no rotation. A missing
// *file* for one spec, by contrast, is not fatal -- druid-restoration, paladin-holy and the
// other 5 non-dps specs simply have no curated APL and RotationCard never asks for one; a
// dps spec whose file has not landed yet reads the same way, with an empty step list rather
// than a build failure.
import { mkdir, readdir, readFile, writeFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { rotationNotesOf } from './rotation-apl.mjs';

const webRoot = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = path.join(webRoot, '..');
const sourceDir = path.join(repoRoot, 'data', 'curated', 'apl');
const target = path.join(webRoot, 'src', 'data', 'generated', 'rotations.json');

let entries;
try {
  entries = await readdir(sourceDir, { withFileTypes: true });
} catch (cause) {
  throw new Error(
    `sync-rotations: ${sourceDir} could not be read; the curated rotation data owns this directory`,
    { cause },
  );
}

/** @type {Record<string, string[]>} */
const rotations = {};
let specsWithNotes = 0;
for (const entry of entries) {
  if (!entry.isFile() || !entry.name.endsWith('.json')) continue;
  const spec = entry.name.slice(0, -'.json'.length);
  const raw = await readFile(path.join(sourceDir, entry.name), 'utf8');
  const notes = rotationNotesOf(JSON.parse(raw));
  rotations[spec] = notes;
  if (notes.length > 0) specsWithNotes += 1;
}

await mkdir(path.dirname(target), { recursive: true });
await writeFile(target, `${JSON.stringify(rotations, null, 2)}\n`, 'utf8');
console.log(
  `sync-rotations: ${Object.keys(rotations).length} curated specs, ${specsWithNotes} with rotation prose -> ` +
    `${path.relative(webRoot, target)}`,
);
