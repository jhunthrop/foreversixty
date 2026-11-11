// web/scripts/check-island-size.mjs
// Two size checks over dist/, both run by `npm run postbuild` so a pull request fails here
// rather than three steps later: the island budgets the spec sets, and a per-file budget
// below the hard limit Cloudflare enforces at deploy time.
import { readdir, readFile } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';
import { gzipSync } from 'node:zlib';
import { ASSET_BUDGET_BYTES, CLOUDFLARE_ASSET_LIMIT_BYTES, measureAssets } from './asset-budget.mjs';
import { startupClosure } from './island-closure.mjs';

// Every island also has a startup-closure budget (`closureLimitBytes`, owner-delegated
// 2026-11-07): the entry plus every chunk it imports statically, transitively, each gzipped on
// its own (scripts/island-closure.mjs), which the entry budget alone never saw: a heavy static
// import split into its own chunk left the entry small and the page just as slow. Each limit is
// the closure measured on the fixture build (FOREVER_DATA=fixture) of 2026-11-07 plus 10%:
// planner 98,497, report 147,156, sim 114,927, sim tools 85,455, Rankings 103,282, Character
// 125,590, GuildShell 128,665, BisTooltipHost 23,746, BisSpecPage 70,428, listbox 15,705 bytes.

/** Checks one island's closure; true when it is within its limit. */
async function checkClosure(label, entry, limitBytes) {
  if (limitBytes === undefined)
    throw new Error(`${label} has no closureLimitBytes: every island budget sets one.`);
  const { files, gzipped } = await startupClosure(entry, path.join(webRoot, 'dist'));
  console.log(`${label} startup closure: ${files.length} files, ${gzipped} bytes gzipped`);
  if (gzipped <= limitBytes) return true;
  console.error(
    `${label}'s startup closure is ${gzipped} bytes gzipped over ${files.length} files, over the ${limitBytes} ` +
      'byte budget. Load the new weight with a dynamic import(), or measure and raise the limit with a reason.',
  );
  return false;
}

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
  // 83 KB since sim-accuracy round 5 (2026-11-01): the Raid-buffed preset carries each spec's
  // published consumable set (sim/request/raid-consumes.json, about 0.7 KB gzipped) so the
  // planner's live figure and the /sim page match the published numbers; the real-data build
  // read 84,126 bytes against 83,968.
  // 86 KB since the redesign batch of 2026-11-09 (ship-fast mode): the planner's density pass
  // and the shared corner frame; the fixture build read 85,274 bytes against 84,992. Debt as above.
  // Closure 108,347 -> 111,000, batch L (2026-11-11): the shared Crest, CharacterName and crest.css (lane D) and the shared states, copy and retry (lane E) ride in every island's startup closure; measured on the box's fixture and real-data builds of the merged tree, which read the same, with the usual ~2% slack. Planner read 109,083.
  { file: 'dist/planner-island.js', limitBytes: 86 * 1024, closureLimitBytes: 111000 },
  // Closure 165,000 since batch L lane E (2026-11-11): the report's empties, errors and filter
  // popover moved onto the shared EmptyState, LoadError (the busy Button with its spinner) and
  // FilterPopover; the fixture build read 163,673 against 161,872. One copy of each, site-wide.
  { file: 'dist/report-island.js', limitBytes: 140 * 1024, closureLimitBytes: 165000 },
  // The sim island is the planner's gear grid plus the report's tables plus a run control.
  // 90 KB gzipped is roughly twice what those parts weigh today and well under the report's
  // ceiling; sim.wasm is not counted, since it is a separate immutable asset fetched after
  // first paint and never on the LCP path.
  { file: 'dist/sim-island.js', limitBytes: 90 * 1024, closureLimitBytes: 126420 },
  // The four combination tools: a candidate grid, an item search, a source picker and two
  // results views, each a lazily-imported chunk off one entry. 70 KB gzipped for the entry
  // is roughly twice what the shared parts (the character strip, the source switcher, the
  // run bar) weigh; the per-tool chunks are not counted here because none of them is on any
  // page's LCP path -- the shell paints its own skeleton first.
  { file: 'dist/sim-tools-island.js', limitBytes: 70 * 1024, closureLimitBytes: 94001 },
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
  // 2026-11-09 (lane P5): the rankings landing (top parses per boss, the spec leaders), the guild
  // boards on the ranked list and the designed empty states measured 113,999 on CI; 114,300 keeps
  // the trip wire. Loading the landing on its own did not help: its modules are shared.
  // Closure 118,000 since the redesign batch of 2026-11-09: the rankings density pass read 115,884.
  // Closure 118,000 -> 124,500, batch L (2026-11-11): the shared Crest, CharacterName and crest.css (lane D) and the shared states, copy and retry (lane E) ride in every island's startup closure; measured on the box's fixture and real-data builds of the merged tree, which read the same, with the usual ~2% slack. Rankings read 122,238.
  { component: 'Rankings', limitBytes: 16 * 1024, closureLimitBytes: 124500 },
  // Raised 16 -> 18 KB (paperdoll lane, 2026-10-31): the character page's Gear section became
  // the shared Paperdoll (design/specs/2026-10-31-paperdoll.md), which statically carries the
  // frame, its model and ItemHover. The spec's own §8 budget lets that frame add up to 4 KB
  // gzipped; measured, Character went from 14.7 KB (main 99d6cae) to 17.9 KB, +3.6 KB. The
  // section is on the page's first view for a character with gear, so lazy-loading it would
  // trade bytes for a second paint of the frame; 18 KB keeps the same ~0.1 KB slack per KB
  // the other page islands have and still trips an accidental heavy import.
  // +256 bytes (row-band sweep, 2026-11-01): the fight tables' marked and addressed rows are the
  // shared row band (lib/ui/row.ts), now a chunk of its own; the entry gains only the preload
  // entries naming it (+38 bytes gzipped measured, 18,424 -> 18,462), with the slack restored.
  // Closure 138,149 -> 140,000 (house formats, lane B, 2026-11-10): the page's dates, clock,
  // numbers, durations, parse figures and countdown now read the shared modules (lib/dates.ts with
  // the zoned calendar and the locale clock, lib/format/number.ts, lib/format/duration.ts,
  // lib/ui/countdown.ts) instead of local copies. Measured on the real-data build: 138,302 on
  // the base (already over) -> 139,451, +1,149 bytes, with the usual ~0.5 KB slack.
  // Closure 140,000 -> 140,900 (unify-icons on top of house formats, 2026-11-10): the page's icons
  // draw through ui/SpellIcon, ItemIcon and lib/ui/assets iconUrl. Measured 140,362 on the merged
  // fixture build, +362 over lane B's figure, with the usual ~0.5 KB slack.
  // Closure 140,900 -> 146,000, batch L (2026-11-11): the shared Crest, CharacterName and crest.css (lane D) and the shared states, copy and retry (lane E) ride in every island's startup closure; measured on the box's fixture and real-data builds of the merged tree, which read the same, with the usual ~2% slack. Character read 143,102.
  { component: 'Character', limitBytes: 18 * 1024 + 256, closureLimitBytes: 146000 },
  // +512 bytes (guild round 3, 2026-11-05): the guild page's first view now carries the week-one
  // card's raider step, the plan gate on the band, the unknown-guild panel and the landing's
  // Your guilds card. Measured 16,564 gzipped against 16,022 at round 2 (main moved too). Lazy
  // loading the Your guilds card and the not-found panel was tried: each saved under 100 bytes
  // and the card flashed in above the lookup, so the bytes stay; 16.5 KB still trips an
  // accidental heavy import.
  // +256 bytes (batch H, h-guild-tabs, 2026-11-10): measured 16,995 with lane/h-edges merged on main
  // (GuildLanding's restyle) before any guild-tab change, 17,020 after: the tabs' new lazy chunks
  // add their names to the shell's preload map. The tabs themselves stay lazy.
  { component: 'GuildShell', limitBytes: 16 * 1024 + 768, closureLimitBytes: 141532 },
  // Lane bis-tooltip-host (2026-09-29): one BisTooltipHost per /bis/<class>/<spec> page
  // replaces roughly 840 per-row ItemHover islands with delegated document-level listeners
  // plus the same lazy-loaded ItemTooltip.svelte panel ItemHover already used -- no data,
  // no per-row state, so its own shell should stay close to ItemHover's own ~2 KB. 8 KB
  // gzipped leaves real headroom without hiding an accidental heavy import the way the
  // other two page-island budgets above already don't.
  { component: 'BisTooltipHost', limitBytes: 8 * 1024, closureLimitBytes: 26121 },
  // The BiS page island (web layout spec 2026-11-05 §5, slice 2): the page head's answer, the
  // control bar, the paperdoll with its inline/side detail, the Sources filter and What to stack,
  // over one compact view model the page inlines as JSON. Measured 19,156 bytes gzipped on the
  // fixture build of 2026-11-06 (the entry; the compare code and the full data load lazily). 20 KB
  // keeps the page islands' slack and still trips an accidental heavy import.
  // Closure 77,471 -> 79,700 (lane P2, the paperdoll page, 2026-11-09): the frame's centre is the
  // selected slot's client tooltip, server-rendered, so ItemTooltip and its Overlay ride in the
  // startup closure (they were the host's lazy chunk only). Measured 78,626 bytes gzipped on the
  // real-data build; lazy-loading the tooltip would paint the centre empty and fill it a beat
  // later. The retired head answer, inspector and weight columns paid for most of it.
  // The catalogue full name (spec-name-feral-bear, batch A) adds a few bytes on top.
  // The logs lane (2026-11-09): ui/Section's `panel` mode (docs/design-system.md, the shared change
  // every Section page carries) brings Plate into the closure, about +116 gzipped on the real build.
  // 79,700 -> 80,600 (lane P9, the character page, 2026-11-09): the character island now draws the
  // same ui/PaperdollFrame, so the frame, its copy and drop-chance leave the BiS chunk for shared
  // chunks of their own (44 -> 46 files, the same code). Measured 80,065 bytes gzipped on the
  // real-data build (79,241 before): the cost is the two chunks' module boundaries, not new weight.
  // Closure 83,500 since the BiS paperdoll of 2026-11-09: the fixture build read 81,796.
  // Closure 83,500 -> 85,700 (unify-icons, 2026-11-10): the paperdoll and lists draw their icons
  // through the shared ui/SpellIcon and ItemIcon with iconUrl. Measured 85,177 on the merged
  // fixture build, with the usual ~0.5 KB slack.
  // Closure 85,700 -> 90,700, batch L (2026-11-11): the shared Crest, CharacterName and crest.css (lane D) and the shared states, copy and retry (lane E) ride in every island's startup closure; measured on the box's fixture and real-data builds of the merged tree, which read the same, with the usual ~2% slack. BisSpecPage read 88,966.
  { component: 'BisSpecPage', limitBytes: 20 * 1024, closureLimitBytes: 90700 },
  // The shared dropdown's lazy chunk (design/specs/2026-10-30-controls-charts.md §A11.2):
  // lib/ui/listbox.ts with its list renderer, the popover core and Floating UI, loaded after
  // first paint on a fine pointer only. The spec asked for 9 KB; it measures 10,221 bytes
  // gzipped on the fixture build of 2026-10-30, of which Floating UI's computePosition with
  // offset, shift and size is about 5.1 KB on its own (its autoUpdate and hide were replaced
  // by popover.ts's own watcher to save 1 KB). 10.5 KB holds it with the fixture/CI gap.
  // Closure 17,276 -> 17,800, batch L (2026-11-11): the shared Crest, CharacterName and crest.css (lane D) and the shared states, copy and retry (lane E) ride in every island's startup closure; measured on the box's fixture and real-data builds of the merged tree, which read the same, with the usual ~2% slack. listbox read 17,460.
  { component: 'listbox', limitBytes: 10.5 * 1024, closureLimitBytes: 17800 },
];

const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

let failed = false;

// Every dist/ file against the per-file budget below Cloudflare's 25 MiB deploy limit
// (scripts/asset-budget.mjs says why each number is what it is).
const distRoot = path.join(webRoot, 'dist');
const { largest, overBudget } = await measureAssets(distRoot, ASSET_BUDGET_BYTES);
for (const { file, bytes } of overBudget) {
  console.error(
    `dist/${file} is ${bytes} bytes, over the ${ASSET_BUDGET_BYTES} byte per-file budget ` +
      `(Cloudflare refuses the whole deploy at ${CLOUDFLARE_ASSET_LIMIT_BYTES}). Compress it the way ` +
      'sim.wasm.gz is, or serve it from R2 through src/worker.ts the way the DuckDB engine modules are.',
  );
  failed = true;
}
console.log(
  `dist/: largest asset is dist/${largest.file} at ${largest.bytes} bytes (budget ${ASSET_BUDGET_BYTES})`,
);

for (const budget of BUDGETS) {
  const bytes = await readFile(path.join(webRoot, budget.file));
  const gzipped = gzipSync(bytes, { level: 9 }).length;
  console.log(`${budget.file}: ${bytes.length} bytes raw, ${gzipped} bytes gzipped`);
  if (gzipped > budget.limitBytes) {
    console.error(`${budget.file} is ${gzipped} bytes gzipped, over the ${budget.limitBytes} byte budget.`);
    failed = true;
  }
  if (!(await checkClosure(budget.file, path.join(webRoot, budget.file), budget.closureLimitBytes)))
    failed = true;
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
  if (!(await checkClosure(`dist/_astro/${matches[0]}`, chunkPath, budget.closureLimitBytes))) failed = true;
}

if (failed) process.exit(1);
