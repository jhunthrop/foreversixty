// web/scripts/check-island-size.mjs
// Two size checks over dist/, both run by `npm run postbuild` so a pull request fails here
// rather than three steps later: the island budgets the spec sets, and the hard per-file
// limit Cloudflare enforces at deploy time.
import { readdir, readFile, stat } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';
import { gzipSync } from 'node:zlib';

// Two standalone islands, two budgets. The planner's 60 KB came from the Phase 1 spec; it is
// 66 KB since the character spine (spec 2026-09-25 section 4): the bar's identity, switch
// list and session read weigh about 3 KB gzipped on the real-data build (64.2 KB on CI on
// 2026-09-25 against 61.4 KB), and the guides and states lanes add a little more to the
// same island; 67 KB since the addon import round of 2026-09-26 (the planner switches
// class itself on a mismatched export and imports once the talents load, plus the
// unknown-item and character-name copy), which read 67,641 bytes on CI against 67,584.
// 70 KB since the web-item-tooltips lane (2026-09-28, rotation-accuracy program): every
// gear slot's icon is now its own ItemHover (tenet 2's "an item is never just a name"),
// which eagerly ships the hover/focus/tap open-close state machine and the edge-aware
// positioning math GearPanel's own BisSlotPopover trigger already paid for once, now
// duplicated for a second, independent tooltip. The heavy panel itself
// (ItemTooltip.svelte) stays code-split behind a dynamic import exactly like
// BisSlotPopover, so only ItemHover's small shell rides here. Measured 69,975 bytes
// gzipped on a fixture build (FOREVER_DATA=fixture) on 2026-09-28; 72 KB leaves headroom
// for the fixture/CI gap this file's own history already shows (the character spine line
// above is ~3 KB apart between the two).
const BUDGETS = [
  // 73 KB since 2026-09-29: the shared item tooltip now renders damage, speed, armor,
  // "Requires Level" and a capped source block the way the client does (tooltip-polish
  // lane), and the BiS popover normalises the ranker's alternatives; the growth was
  // measured at 73,951 bytes gzipped against the earlier 72 KB line.
  // 75 KB since 2026-09-30: the tooltip's PvP source line names the rank title and
  // faction (pvp-faction lane), measured at 75,793 bytes gzipped against the 73 KB line.
  // Debt: the tooltip copy tables should split out of the planner's boot chunk.
  // 77 KB since the lifecycle round of 2026-10-19: the selector's signed-in-empty state
  // (W.1), the card's send-to-addon state machine (W.2; the send itself loads on the first
  // press) and the unread-export panel (W.3) read 78,030 bytes on CI against 76,800.
  // 78 KB since the site alignment of 2026-10-29: the planner moved onto the shared PageHead,
  // Card and compact CardHead and its band compare gained the still-to-take line; CI read
  // 79,107 bytes against 78,848.
  // 82 KB since the shared dropdown of 2026-10-30: the planner's selects moved onto the shared
  // Select, whose closed-field and fine-pointer core ship in the boot chunk. The deploy's
  // real-data build read 82,747 bytes against 79,872 (the fixture build reads ~79.5 KB).
  // Debt: lazy-load the listbox core from the planner so its boot chunk drops back under 79 KB.
  { file: 'dist/planner-island.js', limitBytes: 82 * 1024 },
  { file: 'dist/report-island.js', limitBytes: 140 * 1024 },
  // The sim island is the planner's gear grid plus the report's tables plus a run control.
  // 90 KB gzipped is roughly twice what those parts weigh today and well under the report's
  // ceiling; sim.wasm is not counted, since it is a separate immutable asset fetched after
  // first paint and never on the LCP path.
  { file: 'dist/sim-island.js', limitBytes: 90 * 1024 },
  // The four combination tools: a candidate grid, an item search, a source picker and two
  // results views, each a lazily-imported chunk off one entry. 70 KB gzipped for the entry
  // is roughly twice what the shared parts (the character strip, the source switcher, the
  // run bar) weigh; the per-tool chunks are not counted here because none of them is on any
  // page's LCP path -- the shell paints its own skeleton first.
  { file: 'dist/sim-tools-island.js', limitBytes: 70 * 1024 },
];

// Rankings, Character and GuildShell are ordinary `client:load` Astro islands, not
// standalone Vite builds, so Astro emits them under dist/_astro/ with a content hash in
// the filename (e.g. Rankings.xBviRjNL.js) rather than the fixed name the two budgets
// above match on. GuildShell is the client-side router guild.astro and
// guild/[...path].astro mount (Task 11): it statically imports Guild, GuildClaim,
// GuildSettings and GuildJoin, so its chunk carries all four views' weight, not just the
// one plain-guild page's -- still ~7 KB gzipped today, comfortably under budget. None of
// their pages are in lighthouserc.json's collect.url, so nothing else in CI catches a
// regression here. 16 KB gzipped is enough headroom that legitimate UI growth will not
// flap the build, tight enough that an accidental heavy import (a chart library, a
// duplicated data module) still trips it well before it could threaten a Lighthouse score
// no test here measures directly.
const PAGE_ISLAND_BUDGETS = [
  { component: 'Rankings', limitBytes: 16 * 1024 },
  // Raised 16 -> 18 KB (paperdoll lane, 2026-10-31): the character page's Gear section became
  // the shared Paperdoll (design/specs/2026-10-31-paperdoll.md), which statically carries the
  // frame, its model and ItemHover. The spec's own §8 budget lets that frame add up to 4 KB
  // gzipped; measured, Character went from 14.7 KB (main 99d6cae) to 17.9 KB, +3.6 KB. The
  // section is on the page's first view for a character with gear, so lazy-loading it would
  // trade bytes for a second paint of the frame; 18 KB keeps the same ~0.1 KB slack per KB
  // the other page islands have and still trips an accidental heavy import.
  { component: 'Character', limitBytes: 18 * 1024 },
  { component: 'GuildShell', limitBytes: 16 * 1024 },
  // Lane bis-tooltip-host (2026-09-29): one BisTooltipHost per /bis/<class>/<spec> page
  // replaces roughly 840 per-row ItemHover islands with delegated document-level listeners
  // plus the same lazy-loaded ItemTooltip.svelte panel ItemHover already used -- no data,
  // no per-row state, so its own shell should stay close to ItemHover's own ~2 KB. 8 KB
  // gzipped leaves real headroom without hiding an accidental heavy import the way the
  // other two page-island budgets above already don't.
  { component: 'BisTooltipHost', limitBytes: 8 * 1024 },
  // The shared dropdown's lazy chunk (design/specs/2026-10-30-controls-charts.md §A11.2):
  // lib/ui/listbox.ts with its list renderer, the popover core and Floating UI, loaded after
  // first paint on a fine pointer only. The spec asked for 9 KB; it measures 10,221 bytes
  // gzipped on the fixture build of 2026-10-30, of which Floating UI's computePosition with
  // offset, shift and size is about 5.1 KB on its own (its autoUpdate and hide were replaced
  // by popover.ts's own watcher to save 1 KB). 10.5 KB holds it with the fixture/CI gap.
  { component: 'listbox', limitBytes: 10.5 * 1024 },
];

/**
 * Cloudflare Workers static assets refuse a single file over 25 MiB, on every plan, and
 * the refusal is the whole deploy rather than that one file. Nothing else here catches it:
 * vitest, Playwright and Lighthouse all serve dist/ off the disk, where a 34 MB file is
 * merely a 34 MB file. This is what stops the class of defect that put DuckDB's two
 * engine modules -- 32.7 and 37.5 MiB -- into dist/ in the first place (they are served
 * from R2 now; see scripts/duckdb-runtime.mjs).
 */
const CLOUDFLARE_ASSET_LIMIT_BYTES = 25 * 1024 * 1024;

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

let failed = false;

/** Every file under `dir`, depth first, as paths relative to dist/. */
async function* filesUnder(dir, prefix = '') {
  for (const entry of await readdir(dir, { withFileTypes: true })) {
    const relative = prefix === '' ? entry.name : `${prefix}/${entry.name}`;
    if (entry.isDirectory()) yield* filesUnder(path.join(dir, entry.name), relative);
    else yield relative;
  }
}

const distRoot = path.join(webRoot, 'dist');
let largest = { file: '', bytes: 0 };
for await (const file of filesUnder(distRoot)) {
  const { size } = await stat(path.join(distRoot, file));
  if (size > largest.bytes) largest = { file, bytes: size };
  if (size > CLOUDFLARE_ASSET_LIMIT_BYTES) {
    console.error(
      `dist/${file} is ${size} bytes, over Cloudflare's ${CLOUDFLARE_ASSET_LIMIT_BYTES} byte ` +
        'static-asset limit. `wrangler deploy` would refuse the whole site. Serve it from R2 ' +
        'through src/worker.ts instead, the way the DuckDB engine modules are.',
    );
    failed = true;
  }
}
console.log(`dist/: largest asset is dist/${largest.file} at ${largest.bytes} bytes`);

for (const budget of BUDGETS) {
  const bytes = await readFile(path.join(webRoot, budget.file));
  const gzipped = gzipSync(bytes, { level: 9 }).length;
  console.log(`${budget.file}: ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);
  if (gzipped > budget.limitBytes) {
    console.error(`${budget.file} is ${gzipped} bytes gzipped, over the ${budget.limitBytes} byte budget.`);
    failed = true;
  }
}

// The report charts' lazy chunk (controls and charts spec 2026-10-30 §B3, §B11): uPlot 1.6.32
// (about 22.5 KB gzipped) plus the plot, the tooltip and the brush that drive it. The spec's 30 KB
// cap was raised to 32 KB by the coordinator's ruling of 2026-10-30, on a measured 31,697 bytes:
// moving the tooltip and brush into the entry measured under 30 KB for the chunk, but Rollup then
// split a 19 KB shared chunk out of report-island.js, so the page would have loaded no less.
// The chunk's name is content-hashed, so it is found by uPlot's own class name.
const CHART_CHUNK = { prefix: 'report-island-', marker: 'u-cursor-x', limitBytes: 32 * 1024 };
{
  const distFiles = await readdir(path.join(webRoot, 'dist'));
  const chunks = [];
  for (const name of distFiles) {
    if (!name.startsWith(CHART_CHUNK.prefix) || !name.endsWith('.js')) continue;
    const bytes = await readFile(path.join(webRoot, 'dist', name));
    if (bytes.includes(CHART_CHUNK.marker)) chunks.push({ name, bytes });
  }
  if (chunks.length !== 1) {
    console.error(
      `Expected exactly one dist/${CHART_CHUNK.prefix}<hash>.js holding uPlot, found ${chunks.length}` +
        (chunks.length ? `: ${chunks.map((chunk) => chunk.name).join(', ')}` : '.'),
    );
    failed = true;
  } else {
    const [{ name, bytes }] = chunks;
    const gzipped = gzipSync(bytes, { level: 9 }).length;
    console.log(`dist/${name} (the chart chunk): ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);
    if (gzipped > CHART_CHUNK.limitBytes) {
      console.error(
        `dist/${name} (the chart chunk) is ${gzipped} bytes gzipped, over the ${CHART_CHUNK.limitBytes} byte budget.`,
      );
      failed = true;
    }
  }
}

const astroChunksDir = path.join(webRoot, 'dist/_astro');
const astroChunks = await readdir(astroChunksDir);
for (const budget of PAGE_ISLAND_BUDGETS) {
  const pattern = new RegExp(`^${budget.component}\\.[\\w-]+\\.js$`);
  const matches = astroChunks.filter((name) => pattern.test(name));
  if (matches.length !== 1) {
    console.error(
      `Expected exactly one dist/_astro/${budget.component}.<hash>.js chunk, found ${matches.length}` +
        (matches.length ? `: ${matches.join(', ')}` : '.'),
    );
    failed = true;
    continue;
  }
  const chunkPath = path.join(astroChunksDir, matches[0]);
  const bytes = await readFile(chunkPath);
  const gzipped = gzipSync(bytes, { level: 9 }).length;
  console.log(`dist/_astro/${matches[0]}: ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);
  if (gzipped > budget.limitBytes) {
    console.error(
      `dist/_astro/${matches[0]} is ${gzipped} bytes gzipped, over the ${budget.limitBytes} byte budget.`,
    );
    failed = true;
  }
}

if (failed) process.exit(1);
