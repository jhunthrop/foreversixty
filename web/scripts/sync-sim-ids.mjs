// web/scripts/sync-sim-ids.mjs
// sim/request/IDS.md -> web/src/data/generated/sim-ids.json, before dev, check, test and
// build. The file is gitignored like everything else under src/data/generated: it is
// derived, and a committed copy is a copy that goes stale.
//
// A missing IDS.md is fatal rather than an empty file: the settings panel would render
// with no rows at all and say nothing about why, which is the failure mode the whole
// "an id the engine cannot map is an error at the boundary" rule exists to avoid.
import { mkdir, readFile, writeFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { parseIdsMarkdown } from './ids-md.mjs';

const webRoot = path.join(path.dirname(fileURLToPath(import.meta.url)), '..');
const repoRoot = path.join(webRoot, '..');
const source = path.join(repoRoot, 'sim', 'request', 'IDS.md');
const target = path.join(webRoot, 'src', 'data', 'generated', 'sim-ids.json');

let markdown;
try {
  markdown = await readFile(source, 'utf8');
} catch (cause) {
  throw new Error(
    `sync-sim-ids: ${source} could not be read; the engine generates it with \`go run ./internal/genids\``,
    {
      cause,
    },
  );
}

const ids = parseIdsMarkdown(markdown);
if (ids.buffs.length === 0 || ids.consumables.length === 0) {
  throw new Error(`sync-sim-ids: ${source} carried no buff or consumable table`);
}

await mkdir(path.dirname(target), { recursive: true });
await writeFile(target, `${JSON.stringify(ids, null, 2)}\n`, 'utf8');
console.log(
  `sync-sim-ids: ${ids.buffs.length} buffs, ${ids.consumables.length} consumables, ` +
    `${ids.professions.length} professions, ${ids.worldBuffs.length} world buffs, ` +
    `${ids.stats.length} stats -> ${path.relative(webRoot, target)}`,
);
