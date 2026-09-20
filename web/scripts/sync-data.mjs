// web/scripts/sync-data.mjs
// Copies data/builds/<build>/ directories into web/ before dev, check, test and build:
//   public/data/<build>/…        fetched at runtime by the planner island
//   src/data/generated/*.json    imported at build time by src/pages/classes.astro
//
// FOREVER_DATA picks the source explicitly. `real` (the default) publishes data/builds/ as
// described below. `fixture` publishes src/fixtures/planner regardless of what data/builds
// holds, so the unit tests and the browser suite run against the small, fixed two-tree
// warrior they were written for instead of whatever the data pipeline last emitted. Tests
// that assert on talent names, item ids and counts are only meaningful against data that
// does not move; the real data gets its own smoke suite (tests/e2e/real-data.spec.ts).
//
// Every build whose manifest lists Phase 1 talent data is published under public/data/, not
// only the active one. Shared builds are immutable and keyed by tree_version: /b/:id renders
// with the tree_version the build was saved against, and the island fetches
// /data/<tree_version>/… for it. Publishing only the active build would turn every share
// link made against an earlier build into "Talent data did not load" the moment
// active-build.json moved on. The retained set matches what the API accepts -- any
// tree_version it holds data for. src/data/generated/ still mirrors the active build alone,
// because those are the build-time imports for /classes.
//
// The pipeline's manifest.json is the contract: every path it lists must exist on disk, or
// this script throws and names the missing files. Under FOREVER_DATA=real, a build whose
// manifest lists none of the Phase 1 per-class directories falls back to the checked-in
// fixture at src/fixtures/planner — loudly. Neither that fallback nor an explicit
// FOREVER_DATA=fixture is allowed when CF_PAGES is set (the same deploy guard
// astro.config.mjs uses for placeholder community links).
import { access, cp, mkdir, readFile, readdir, rm, writeFile } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath, pathToFileURL } from 'node:url';

/** What gets copied into public/data/<build>/, in copy order. */
export const SYNC_ENTRIES = [
  { name: 'talents', kind: 'dir', required: true },
  { name: 'items', kind: 'dir', required: false },
  { name: 'icons', kind: 'dir', required: false },
  { name: 'trees', kind: 'dir', required: false },
  { name: 'sets.json', kind: 'file', required: false },
  // Contract 6 and 10.4: the Droptimizer's source tables and Top Gear's enchant and suffix
  // lists. Optional the same way sets.json is -- a build the data lane has not regenerated
  // ships none, and the pages say so rather than failing to render. simbuffs.json (the
  // buff/consumable name table) is already listed below, beside simconsumes.json.
  { name: 'loot.json', kind: 'file', required: false },
  { name: 'enchants.json', kind: 'file', required: false },
  { name: 'suffixes.json', kind: 'file', required: false },
  { name: 'classes.json', kind: 'file', required: true },
  { name: 'races.json', kind: 'file', required: true },
  { name: 'combos.json', kind: 'file', required: true },
  { name: 'spellconst', kind: 'dir', required: false },
  { name: 'spells.json', kind: 'file', required: false },
  { name: 'simconsumes.json', kind: 'file', required: false },
  { name: 'simbuffs.json', kind: 'file', required: false },
];

/** The files src/pages/classes.astro imports statically. */
export const PAGE_IMPORTS = ['classes.json', 'races.json', 'combos.json'];

/** The values FOREVER_DATA accepts. `real` is the default; see the file header. */
export const DATA_SOURCES = ['real', 'fixture'];

/**
 * Reads the FOREVER_DATA selector, defaulting to 'real'. Throws on any other value rather
 * than guessing: a typo that silently published real data would make the test suites
 * nondeterministic in exactly the way this selector exists to prevent.
 * @param {string | undefined} value
 * @returns {'real' | 'fixture'}
 */
export function resolveDataSource(value) {
  if (value === undefined || value === '') return 'real';
  if (!DATA_SOURCES.includes(value)) {
    throw new Error(`FOREVER_DATA must be one of ${DATA_SOURCES.join(', ')}; got ${JSON.stringify(value)}.`);
  }
  return value;
}

/** True once the manifest lists the Phase 1 per-class talent files. */
export function hasPhaseOneData(manifest) {
  return Object.keys(manifest?.files ?? {}).some((key) => key.startsWith('talents/'));
}

/** The data/builds/<build>/ directory names, sorted. Empty when data/builds/ is absent. */
async function listBuildDirs(repoRoot) {
  let entries;
  try {
    entries = await readdir(path.join(repoRoot, 'data/builds'), { withFileTypes: true });
  } catch {
    return [];
  }
  return entries
    .filter((entry) => entry.isDirectory())
    .map((entry) => entry.name)
    .sort();
}

/**
 * Splits data/builds/ into the builds worth publishing and the ones with nothing to publish.
 * `retained` is every build whose manifest lists Phase 1 talent data; `skipped` is the rest,
 * so the caller can name them.
 * @param {string} repoRoot
 * @returns {Promise<{ retained: string[], skipped: string[] }>}
 */
export async function partitionBuilds(repoRoot) {
  const retained = [];
  const skipped = [];
  for (const build of await listBuildDirs(repoRoot)) {
    const manifestFile = path.join(repoRoot, 'data/builds', build, 'manifest.json');
    const manifest = (await exists(manifestFile)) ? await readJson(manifestFile) : null;
    (hasPhaseOneData(manifest) ? retained : skipped).push(build);
  }
  return { retained, skipped };
}

async function exists(file) {
  try {
    await access(file);
    return true;
  } catch {
    return false;
  }
}

/** Manifest-listed paths with no file on disk, sorted. */
export async function missingManifestFiles(manifest, sourceDir) {
  const names = Object.keys(manifest?.files ?? {}).sort();
  const missing = [];
  for (const name of names) {
    if (!(await exists(path.join(sourceDir, name)))) missing.push(name);
  }
  return missing;
}

async function readJson(file) {
  const raw = await readFile(file, 'utf8');
  try {
    return JSON.parse(raw);
  } catch (cause) {
    throw new Error(`${file}: invalid JSON (${cause.message})`, { cause });
  }
}

/**
 * Decides which directory to copy from: the checked-in fixture when the caller asked for it,
 * otherwise the real build directory when its manifest lists Phase 1 talent data, otherwise
 * the fixture again as a fallback. Throws when the manifest is missing listed files, or when
 * the fixture is not allowed.
 * @param {{
 *   repoRoot: string,
 *   webRoot: string,
 *   build: string,
 *   source: 'real' | 'fixture',
 *   allowFixture: boolean,
 * }} options
 */
async function resolveSourceDir({ repoRoot, webRoot, build, source, allowFixture }) {
  const fixtureDir = path.join(webRoot, 'src/fixtures/planner');
  if (source === 'fixture') {
    if (!allowFixture) {
      throw new Error(
        `FOREVER_DATA=fixture is refused here: this build publishes the site, and fixture ` +
          `talent and item content is placeholder data. Unset FOREVER_DATA or set it to real.`,
      );
    }
    return { sourceDir: fixtureDir, usedFixture: true };
  }

  const buildDir = path.join(repoRoot, 'data/builds', build);
  const manifestFile = path.join(buildDir, 'manifest.json');
  const manifest = (await exists(manifestFile)) ? await readJson(manifestFile) : null;

  if (manifest && hasPhaseOneData(manifest)) {
    const missing = await missingManifestFiles(manifest, buildDir);
    if (missing.length > 0) {
      throw new Error(
        `${path.relative(repoRoot, manifestFile)} lists files that are missing: ${missing.join(', ')}. ` +
          `Re-run the data pipeline for build ${build}.`,
      );
    }
    return { sourceDir: buildDir, usedFixture: false };
  }

  if (!allowFixture) {
    throw new Error(
      `data/builds/${build} has no Phase 1 talent data (manifest.json lists no talents/*.json). ` +
        `Run the data pipeline before deploying; the fixture fallback is refused here.`,
    );
  }
  return { sourceDir: fixtureDir, usedFixture: true };
}

/**
 * public/data/<build>/simnames/<class>.json: spell and item ids to display names, for the
 * simulator's results tables.
 *
 * The simulator's summary names abilities by engine action key (sim/adapter.ActionName), so
 * the web has to resolve them. spells.json is 31,754 rows and 1.8 MB -- far too big to
 * fetch on /sim -- so this prunes it to what a sim of that class can actually reference:
 * every spell in the class's spellconst file (falling back to spells.json only for a name
 * spellconst omits, never for membership), plus the class's own items for the item:<id>
 * rows. items/<class>.json carries no spell, proc or on-use field -- its rows are only
 * `{ id, name, icon, slot, quality, required_level, item_level, armor, stats, set_id,
 * unique }` -- so a trinket's triggered spell is not pulled into `spell` from here; if that
 * spell is not itself one of the class's own spellconst entries, resolveActionName falls
 * back to the raw key for it, which is legible. A build without spellconst (an older one,
 * or one the data lane has not regenerated) publishes nothing, and resolveActionName falls
 * back to the key for everything.
 * @param {string} buildDir source data/builds/<build>
 * @param {string} outDir   public/data/<build>
 * @returns {Promise<string[]>} the simnames/<class>.json paths written, relative to outDir
 */
export async function writeSimNames(buildDir, outDir) {
  let spells;
  try {
    spells = JSON.parse(await readFile(path.join(buildDir, 'spells.json'), 'utf8'));
  } catch {
    return [];
  }
  const nameById = new Map(spells.map((row) => [String(row.id), row.name]));

  let classFiles;
  try {
    classFiles = await readdir(path.join(buildDir, 'spellconst'));
  } catch {
    return [];
  }

  const written = [];
  await mkdir(path.join(outDir, 'simnames'), { recursive: true });
  for (const file of classFiles.filter((name) => name.endsWith('.json'))) {
    const slug = file.replace(/\.json$/, '');
    const constants = JSON.parse(await readFile(path.join(buildDir, 'spellconst', file), 'utf8'));
    const spell = {};
    for (const [id, row] of Object.entries(constants.spells ?? {})) {
      spell[id] = row.name ?? nameById.get(id) ?? id;
    }

    const item = {};
    try {
      const items = JSON.parse(await readFile(path.join(buildDir, 'items', file), 'utf8'));
      for (const row of items.items ?? []) item[String(row.id)] = row.name;
    } catch {
      // A class with no normalized item list publishes spell names only.
    }

    const target = path.join(outDir, 'simnames', file);
    await writeFile(target, JSON.stringify({ build: constants.build, class_slug: slug, spell, item }));
    written.push(`simnames/${file}`);
  }
  return written;
}

/**
 * Resets public/data/<build> and copies SYNC_ENTRIES from sourceDir into it, then derives
 * public/data/<build>/simnames/<class>.json from the same sourceDir (see `writeSimNames`).
 * Throws when a required entry is missing from sourceDir. Returns the names actually
 * copied or written, SYNC_ENTRIES first.
 * @param {{ repoRoot: string, webRoot: string, build: string, sourceDir: string }} options
 */
async function copyBuild({ repoRoot, webRoot, build, sourceDir }) {
  const publicDir = path.join(webRoot, 'public/data', build);
  await rm(publicDir, { recursive: true, force: true });
  await mkdir(publicDir, { recursive: true });

  const copied = [];
  for (const entry of SYNC_ENTRIES) {
    const from = path.join(sourceDir, entry.name);
    if (!(await exists(from))) {
      if (entry.required) {
        throw new Error(
          `${path.relative(repoRoot, sourceDir)}: required entry ${entry.name} is missing. ` +
            `The planner cannot build without it.`,
        );
      }
      continue;
    }
    await cp(from, path.join(publicDir, entry.name), { recursive: entry.kind === 'dir' });
    copied.push(entry.name);
  }

  copied.push(...(await writeSimNames(sourceDir, publicDir)));

  return copied;
}

/**
 * Mirrors PAGE_IMPORTS from the active build's source directory into src/data/generated,
 * which src/pages/classes.astro imports at build time. Only the active build lands here.
 * @param {{ webRoot: string, sourceDir: string }} options
 */
async function writePageImports({ webRoot, sourceDir }) {
  const generatedDir = path.join(webRoot, 'src/data/generated');
  await mkdir(generatedDir, { recursive: true });
  for (const name of PAGE_IMPORTS) {
    const body = await readFile(path.join(sourceDir, name), 'utf8');
    await writeFile(path.join(generatedDir, name), body, 'utf8');
  }
}

/**
 * Publishes the retained builds other than the active one. Each already has Phase 1 data, so
 * the fixture fallback can never apply and is refused. Returns the build ids published.
 * @param {{ repoRoot: string, webRoot: string, builds: string[] }} options
 */
async function publishRetained({ repoRoot, webRoot, builds }) {
  const published = [];
  for (const build of builds) {
    const { sourceDir } = await resolveSourceDir({
      repoRoot,
      webRoot,
      build,
      source: 'real',
      allowFixture: false,
    });
    await copyBuild({ repoRoot, webRoot, build, sourceDir });
    published.push(build);
  }
  return published;
}

/**
 * Removes public/data/<build> directories this run did not publish. Without it, switching
 * FOREVER_DATA or retiring a build leaves a stale directory the island would still fetch.
 * @param {{ webRoot: string, published: string[] }} options
 * @returns {Promise<string[]>} the build ids removed
 */
async function pruneUnpublished({ webRoot, published }) {
  const publicData = path.join(webRoot, 'public/data');
  const keep = new Set(published);
  let entries;
  try {
    entries = await readdir(publicData, { withFileTypes: true });
  } catch {
    return [];
  }
  const removed = [];
  for (const entry of entries) {
    if (!entry.isDirectory() || keep.has(entry.name)) continue;
    await rm(path.join(publicData, entry.name), { recursive: true, force: true });
    removed.push(entry.name);
  }
  return removed;
}

/**
 * @param {{
 *   repoRoot: string,
 *   webRoot: string,
 *   source?: 'real' | 'fixture',
 *   allowFixture?: boolean,
 *   log?: { log: (...args: unknown[]) => void, warn: (...args: unknown[]) => void },
 * }} options
 */
export async function syncData({
  repoRoot,
  webRoot,
  source = 'real',
  allowFixture = true,
  log = console,
} = {}) {
  const activeFile = path.join(webRoot, 'src/data/active-build.json');
  const { build } = await readJson(activeFile);
  if (typeof build !== 'string' || build.length === 0) {
    throw new Error(`${activeFile} must contain a non-empty "build" string.`);
  }

  const { sourceDir, usedFixture } = await resolveSourceDir({
    repoRoot,
    webRoot,
    build,
    source,
    allowFixture,
  });
  if (usedFixture && source === 'real') {
    log.warn(
      `sync-data: data/builds/${build} has no Phase 1 talent data; copying the fixture from ` +
        `src/fixtures/planner instead. Talent and item content on this build is placeholder data.`,
    );
  }
  const copied = await copyBuild({ repoRoot, webRoot, build, sourceDir });
  await writePageImports({ webRoot, sourceDir });
  log.log(
    `sync-data: active build ${build} -> public/data/${build} (${copied.join(', ')})` +
      `${usedFixture ? ' [fixture]' : ''}`,
  );

  // The active build has already been published, with whatever fixture handling it needed.
  // Under FOREVER_DATA=fixture there is nothing else to retain: data/builds is not consulted
  // at all, so the fixture stands alone as the one published build.
  const published = [build];
  if (source === 'real') {
    const { retained, skipped } = await partitionBuilds(repoRoot);
    for (const other of skipped.filter((name) => name !== build)) {
      log.warn(
        `sync-data: data/builds/${other} has no Phase 1 talent data; not publishing it. ` +
          `Share links made against build ${other} will not find their talent data.`,
      );
    }
    published.push(
      ...(await publishRetained({
        repoRoot,
        webRoot,
        builds: retained.filter((name) => name !== build),
      })),
    );
  }
  const pruned = await pruneUnpublished({ webRoot, published });
  if (pruned.length > 0) log.log(`sync-data: removed stale public/data/${pruned.join(', ')}`);
  log.log(`sync-data: published ${published.length} build(s) from ${source} data: ${published.join(', ')}`);

  return { build, source, sourceDir, usedFixture, copied, published, pruned };
}

const invokedDirectly =
  process.argv[1] !== undefined && import.meta.url === pathToFileURL(process.argv[1]).href;

if (invokedDirectly) {
  const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
  await syncData({
    repoRoot: path.resolve(webRoot, '..'),
    webRoot,
    source: resolveDataSource(process.env.FOREVER_DATA),
    // CF_PAGES is set only in the deploy build (see .github/workflows/web.yml), where
    // shipping fixture talent data to foreversixty.gg would be a lie.
    allowFixture: !process.env.CF_PAGES,
  });
}
