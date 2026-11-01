// web/scripts/sync-raid-preset.mjs
// data/curated/presets.json's "raid" preset -> web/src/data/generated/raid-preset.json, before
// dev, check, test and build. The BiS ranker, the tier list and every sim tool run their
// level-60 numbers under that preset (sim/request/preset.go). The /sim page's Raid-buffed
// preset reads its buff and debuff ids from here, so a number on the page is simmed with
// the buffs the published number was. The file is gitignored like everything else under
// src/data/generated: it is derived, and a committed copy is a copy that goes stale.
//
// A missing or empty preset is fatal: a Raid-buffed preset with no buffs would still claim
// the label, and the page would sim a bare character under it without saying so.
import { mkdir, readFile, writeFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const webRoot = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = path.join(webRoot, '..');
const source = path.join(repoRoot, 'data', 'curated', 'presets.json');
const target = path.join(webRoot, 'src', 'data', 'generated', 'raid-preset.json');
const PRESET = 'raid';

/** The ids of one preset list, in the curated file's order. */
function idsOf(entries, list) {
  if (!Array.isArray(entries) || entries.length === 0) {
    throw new Error(`sync-raid-preset: ${source} preset "${PRESET}" has no ${list}`);
  }
  return entries.map((entry) => {
    if (typeof entry?.id !== 'string' || entry.id === '') {
      throw new Error(`sync-raid-preset: ${source} preset "${PRESET}" has a ${list} entry with no id`);
    }
    return entry.id;
  });
}

let curated;
try {
  curated = JSON.parse(await readFile(source, 'utf8'));
} catch (cause) {
  throw new Error(`sync-raid-preset: ${source} could not be read; data/curated owns the sim presets`, {
    cause,
  });
}
const preset = curated?.presets?.[PRESET];
if (preset === undefined) {
  throw new Error(`sync-raid-preset: ${source} has no "${PRESET}" preset`);
}

const out = { buffs: idsOf(preset.buffs, 'buffs'), debuffs: idsOf(preset.debuffs, 'debuffs') };
await mkdir(path.dirname(target), { recursive: true });
await writeFile(target, `${JSON.stringify(out, null, 2)}\n`, 'utf8');
console.log(
  `sync-raid-preset: ${out.buffs.length} buffs, ${out.debuffs.length} debuffs -> ${path.relative(webRoot, target)}`,
);
