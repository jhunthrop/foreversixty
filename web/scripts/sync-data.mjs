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
import { bisDirFor, bisSourceBuild } from './bis-source-build.mjs';

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
  // The leveling BiS lists (lane bis-all, docs/superpowers/specs/2026-09-28-leveling-bis-
  // design.md): one <spec>.json per written spec, published like loot.json once the
  // pipeline lands it. Optional the same way -- a build the data lane has not run the
  // ranking command for yet ships none, and /bis's own page falls back to this lane's
  // committed fixture (src/data/fixtures/bis/) instead of failing the sync. `copyBuild`
  // below fills in any spec this entry did not publish from src/data/fixtures/bis/ once
  // this loop is done (see `publishBisFixtures`), so the planner's BiS hover popover
  // (lib/bis/hover.ts) can fetch a spec's file the same way the static /bis page's own
  // loader (lib/bis/load.ts) already falls back for it.
  { name: 'bis', kind: 'dir', required: false },
  { name: 'enchants.json', kind: 'file', required: false },
  { name: 'suffixes.json', kind: 'file', required: false },
  // The item ids the engine's embedded SimDatabase actually carries (pipeline/simdb/
  // items.py's simdb_item_rows) -- every bulk candidate source filters against it. Optional
  // like the three above: a build the data lane has not regenerated this file for ships
  // none, and an absent file means "do not filter" (sim/sim-items.ts's loadSimItems).
  { name: 'simitems.json', kind: 'file', required: false },
  { name: 'itemnames.json', kind: 'file', required: false },
  { name: 'classes.json', kind: 'file', required: true },
  { name: 'races.json', kind: 'file', required: true },
  { name: 'combos.json', kind: 'file', required: true },
  { name: 'spellconst', kind: 'dir', required: false },
  { name: 'spells.json', kind: 'file', required: false },
  // The spell id to icon stem table the coach panel's icons read.
  { name: 'spellicons.json', kind: 'file', required: false },
  { name: 'simconsumes.json', kind: 'file', required: false },
  { name: 'simbuffs.json', kind: 'file', required: false },
  // The build's copy of the curated stat weights (pipeline/weights.py), which the addon
  // and the site's weight pickers both read. Optional: only builds the data lane has run
  // `pipeline weights` for ship one -- 1.15.9.69722 and forever-prebeta do not.
  { name: 'stat-weights.json', kind: 'file', required: false },
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
 * Names for the spell ids the sim engine reports that no class's spellconst or talents name:
 * world buffs and racials no class spellbook owns, and the engine ids the client's SpellName
 * table lacks (Ignite 12654, the Era ranks of talent procs). Curated beside the engine files
 * that use them, and read by sim/cmd/spec-breakdown's TestEveryPublishedResultActionHasAName
 * too, so one file says which ids the engine needs named.
 */
const ENGINE_SPELL_NAMES = fileURLToPath(
  new URL('../../data/curated/engine-spell-names.json', import.meta.url),
);

/**
 * How many files the client's whole spell and item name tables are split into
 * (simnames/spell/<id % n>.json, simnames/item/<id % n>.json). The page fetches the one shard
 * an id it cannot name lives in; 64 keeps each spell shard near 4 KB gzipped. Written into
 * simnames/_shared.json as `shards`, so the page never restates it.
 */
const NAME_SHARDS = 64;

/** The curated engine spell names, id to name. A missing or malformed file is fatal: the
 *  page would read "An unnamed spell" for every id it carries. */
async function engineSpellNames(file) {
  const curated = JSON.parse(await readFile(file, 'utf8'));
  const names = {};
  for (const [id, row] of Object.entries(curated?.spells ?? {})) {
    if (typeof row?.name !== 'string' || row.name === '') throw new Error(`${file}: spell ${id} has no name`);
    names[id] = row.name;
  }
  if (Object.keys(names).length === 0) throw new Error(`${file} names no spell`);
  return names;
}

/**
 * Writes `rows` (`{ id, name }`) as NAME_SHARDS files under `dir`, each the id-to-name table of
 * the ids that land in it.
 */
async function writeNameShards(dir, rows) {
  const shards = Array.from({ length: NAME_SHARDS }, () => ({}));
  for (const row of rows) shards[row.id % NAME_SHARDS][String(row.id)] = row.name;
  await mkdir(dir, { recursive: true });
  await Promise.all(
    shards.map((table, index) => writeFile(path.join(dir, `${index}.json`), JSON.stringify(table))),
  );
}

/**
 * WoW's own item-class ids (items.json's `class_id`): Consumable and Reagent, the two
 * classes a potion, elixir, food, bandage or mana gem lives in. Not every single-use
 * consumable this build ships lands in one of the two, though -- items.json misclassifies
 * a small, known set of mana gems as Trade Goods (class 4) instead -- so
 * ITEM_NAME_OVERRIDES below carries the ones production has actually reported.
 */
const ITEM_CLASS_CONSUMABLE = 0;
const ITEM_CLASS_REAGENT = 7;

/**
 * Item ids items.json puts outside ITEM_CLASS_CONSUMABLE/ITEM_CLASS_REAGENT despite being
 * ordinary single-use consumables (2026-09-21 result-page review round 2: production
 * reported item 5514, Mana Agate, unresolved -- items.json has it as class_id 4, Trade
 * Goods, alongside 6,000 genuine crafting reagents this table must NOT swallow whole).
 * Extend this the day another single-use item turns up misclassified the same way.
 */
const ITEM_NAME_OVERRIDES = {
  5514: 'Mana Agate',
};

/**
 * public/data/<build>/simnames/<class>.json: spell and item ids to display names, for the
 * simulator's results tables.
 *
 * The simulator's summary names abilities by engine action key (sim/adapter.ActionName), so
 * the web has to resolve them. spells.json is 31,754 rows and 1.8 MB -- far too big to
 * fetch on /sim -- so this prunes it to what a sim of that class can actually reference:
 * every spell in the class's spellconst file (falling back to spells.json only for a name
 * spellconst omits, never for membership), plus the class's own items for the item:<id>
 * rows. items/<class>.json carries no spell *id* for a proc or on-use -- its rows are
 * `{ id, name, icon, slot, quality, required_level, item_level, armor, stats, damage_min,
 * damage_max, speed, dps, two_hand, effect_text, set_id, unique }`, and `effect_text`
 * is the effect's prose ("Chance on hit: ..."), not the id of the spell behind it -- so a
 * trinket's triggered spell still cannot be resolved into an `item:<id>`-adjacent `spell`
 * row from here; if that spell is not itself one of the class's own spellconst entries,
 * the page reads it from the shards below. A build without spellconst (an older one, or one
 * the data lane has not regenerated) publishes nothing, and resolveActionName falls back to
 * its unnamed prose for everything.
 *
 * Also writes public/data/<build>/simnames/_shared.json: every class's own `spell` table
 * unioned into one (2026-09-21 result-page review, Defect 2), plus every class's own
 * TALENT spells (round 2 below), ENGINE_SPELL_NAMES above, `shards` (NAME_SHARDS), and a
 * `item` table of the build's consumables and reagents (round 2 below too). The
 * raid-buffed preset (RAID_BUFFS, web/src/lib/sim/settings.ts) can put an aura on the
 * player from ANY of the nine classes -- a mage running it gets a paladin's Blessing of
 * Kings and a warrior's Battle Shout same as its own Frostbolt -- and a build with no
 * class-of-its-own for that spell has nowhere else to resolve it from. Spell ids are one
 * global namespace (the engine's own ActionID.SpellID), so a later entry never overwrites
 * an earlier one that already named the same id; ordering only matters for which of two
 * sources "wins" a genuine collision, which this build has none of.
 *
 * Round 2 (2026-09-21, the same review's follow-up): production still showed four more
 * kinds of raw id once RAID_BUFFS itself was fixed --
 *   - a talent-granted PASSIVE proc (Flurry, a warrior talent): spellconst is pruned to a
 *     class's CASTABLE spellbook, so a proc with no cast bar entry of its own never
 *     reaches it even though it is unmistakably that class's -- every class's own
 *     talents/<class>.json (already shipped for the planner) names each talent's spell id
 *     per rank, and unioning that in closes exactly this gap;
 *   - a racial (Blood Fury): no class owns it and races.json carries no spell ids at all
 *     (ENGINE_SPELL_NAMES above, sourced from the engine directly);
 *   - a Darkmoon Faire world buff variant this file had not named yet (ENGINE_SPELL_NAMES
 *     above, all eight of Sayge's Dark Fortune);
 *   - and an ITEM action (a potion, a mana gem) -- items/<class>.json is gear only, so a
 *     consumable was never in any class's own item table at all. `item` below is
 *     items.json (the client's full table, same reason spells.json is excluded whole)
 *     filtered to Consumable and Reagent, plus ITEM_NAME_OVERRIDES for the rare item this
 *     build misclassifies outside both.
 *
 * Round 7 (2026-11-02, owner ruling 2): /sim still showed 36 rows as "An unnamed spell" --
 * item procs and generic actions no class table carries (Thrash, Chthonic Power, the wand's
 * Shoot, a set bonus's proc) and engine ids the client lacks (Ignite 12654). The client's
 * whole spell and item tables are now published as simnames/spell/<n>.json and
 * simnames/item/<n>.json (NAME_SHARDS files each, keyed by id modulo the count), which the
 * page fetches only for an id the class and shared tables cannot name; the engine's own ids
 * come from ENGINE_SPELL_NAMES.
 * @param {string} buildDir source data/builds/<build>
 * @param {string} outDir   public/data/<build>
 * @param {string} [engineNamesFile] the curated engine spell names (ENGINE_SPELL_NAMES)
 * @returns {Promise<string[]>} the simnames/<class>.json, simnames/_shared.json and shard
 *   directory paths written, relative to outDir
 */
export async function writeSimNames(buildDir, outDir, engineNamesFile = ENGINE_SPELL_NAMES) {
  let spells;
  try {
    spells = JSON.parse(await readFile(path.join(buildDir, 'spells.json'), 'utf8'));
  } catch {
    return [];
  }
  const nameById = new Map(spells.map((row) => [String(row.id), row.name]));

  /**
   * A class's own talent spells: every rank's `spell_id`, named by the talent's own
   * display name (a rank's own `description` is prose about what changed, not a name).
   * Missing or malformed talents/<class>.json publishes nothing extra -- the per-class
   * spellconst union above still runs -- the same "a build the data lane has not
   * regenerated this for yet loses nothing it already had" shape sim-items.ts's own
   * `loadOptional` callers use.
   */
  async function talentSpellNames(classSlug) {
    let talents;
    try {
      talents = JSON.parse(await readFile(path.join(buildDir, 'talents', `${classSlug}.json`), 'utf8'));
    } catch {
      return {};
    }
    const spell = {};
    for (const tree of talents.trees ?? []) {
      for (const talent of tree.talents ?? []) {
        for (const rank of talent.ranks ?? []) {
          if (rank.spell_id) spell[String(rank.spell_id)] = talent.name;
        }
        if (talent.spell_id) spell[String(talent.spell_id)] = talent.name;
      }
    }
    return spell;
  }

  /** The build's items.json rows; none when the build has no file. */
  async function clientItems() {
    try {
      return JSON.parse(await readFile(path.join(buildDir, 'items.json'), 'utf8'));
    } catch {
      return [];
    }
  }

  /**
   * The build's consumable and reagent item names, keyed by id, for every class at once --
   * a potion is nobody's gear, so no class's own items/<class>.json (equipment only) ever
   * carries one.
   */
  function sharedItemNames(items) {
    const item = { ...ITEM_NAME_OVERRIDES };
    for (const row of items) {
      if (row.class_id === ITEM_CLASS_CONSUMABLE || row.class_id === ITEM_CLASS_REAGENT) {
        item[String(row.id)] = row.name;
      }
    }
    return item;
  }

  let classFiles;
  try {
    classFiles = await readdir(path.join(buildDir, 'spellconst'));
  } catch {
    return [];
  }

  const written = [];
  const sharedSpell = await engineSpellNames(engineNamesFile);
  let buildVersion;
  await mkdir(path.join(outDir, 'simnames'), { recursive: true });
  for (const file of classFiles.filter((name) => name.endsWith('.json'))) {
    const slug = file.replace(/\.json$/, '');
    const constants = JSON.parse(await readFile(path.join(buildDir, 'spellconst', file), 'utf8'));
    buildVersion = constants.build;
    const spell = {};
    for (const [id, row] of Object.entries(constants.spells ?? {})) {
      spell[id] = row.name ?? nameById.get(id) ?? id;
      if (!(id in sharedSpell)) sharedSpell[id] = spell[id];
    }
    for (const [id, name] of Object.entries(await talentSpellNames(slug))) {
      if (!(id in sharedSpell)) sharedSpell[id] = name;
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
  if (written.length > 0) {
    const sharedTarget = path.join(outDir, 'simnames', '_shared.json');
    const items = await clientItems();
    await writeFile(
      sharedTarget,
      JSON.stringify({
        build: buildVersion,
        shards: NAME_SHARDS,
        spell: sharedSpell,
        item: sharedItemNames(items),
      }),
    );
    written.push('simnames/_shared.json');
    await writeNameShards(path.join(outDir, 'simnames', 'spell'), spells);
    await writeNameShards(path.join(outDir, 'simnames', 'item'), items);
    written.push('simnames/spell/', 'simnames/item/');
  }
  return written;
}

/**
 * Fills gaps in public/data/<build>/bis/ from this lane's own committed fixture
 * (src/data/fixtures/bis/) for any spec the SYNC_ENTRIES `bis` entry above did not already
 * publish -- lib/bis/load.ts's own precedence (a real data/builds/<build>/bis/<spec>.json
 * wins, else the committed fixture, per spec) applied to the fetchable copy under
 * public/data/ rather than to the static /bis page's build-time read. Runs for every
 * published build, fixture source and real alike: a real build the data lane has not run
 * the ranking command for yet still serves whatever fixture specs are committed, exactly
 * the way loadBisFile already does for the page itself.
 * @param {{ webRoot: string, publicDir: string }} options
 * @returns {Promise<string[]>} `bis/<spec>.json` paths filled in, relative to publicDir
 */
async function publishBisFixtures({ webRoot, publicDir }) {
  const fixtureDir = path.join(webRoot, 'src/data/fixtures/bis');
  let entries;
  try {
    entries = await readdir(fixtureDir);
  } catch {
    return [];
  }
  const bisDir = path.join(publicDir, 'bis');
  const filled = [];
  for (const name of entries.filter((entry) => entry.endsWith('.json'))) {
    const target = path.join(bisDir, name);
    if (await exists(target)) continue; // the real build already ranks this spec
    await mkdir(bisDir, { recursive: true });
    await cp(path.join(fixtureDir, name), target);
    filled.push(`bis/${name}`);
  }
  return filled;
}

/**
 * A real build the nightly has not ranked yet publishes the newest ranked build's bis/
 * (scripts/bis-source-build.mjs), so the planner's hover popover and the home upgrades panel
 * fetch the same files the static /bis pages read. Each file's own `build` field still names
 * the build it was ranked on. A no-op when the build has its own bis/ or the source is the
 * fixture.
 * @param {{ repoRoot: string, build: string, sourceDir: string, publicDir: string }} options
 * @returns {Promise<string[]>} `bis` when a fallback directory was copied
 */
async function publishFallbackBis({ repoRoot, build, sourceDir, publicDir }) {
  if (sourceDir !== path.join(repoRoot, 'data/builds', build)) return [];
  const source = bisSourceBuild(repoRoot, build);
  if (source === build) return [];
  await cp(bisDirFor(repoRoot, source), path.join(publicDir, 'bis'), { recursive: true });
  return [`bis (ranked on ${source})`];
}

/**
 * Resets public/data/<build> and copies SYNC_ENTRIES from sourceDir into it, then derives
 * public/data/<build>/simnames/<class>.json from the same sourceDir (see `writeSimNames`)
 * and fills any remaining bis/<spec>.json gaps from this lane's own fixture (see
 * `publishBisFixtures`). Throws when a required entry is missing from sourceDir. Returns
 * the names actually copied or written, SYNC_ENTRIES first.
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
  copied.push(...(await publishFallbackBis({ repoRoot, build, sourceDir, publicDir })));
  copied.push(...(await publishBisFixtures({ webRoot, publicDir })));

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
  const rootFiles = await publishRootFiles(webRoot);
  log.log(`sync-data: published ${rootFiles.map((name) => `/data/${name}`).join(', ')}`);
  log.log(`sync-data: published ${published.length} build(s) from ${source} data: ${published.join(', ')}`);

  return { build, source, sourceDir, usedFixture, copied, published, pruned, rootFiles };
}

/**
 * The build-independent files the apps fetch at the data root (docs/contracts/app-v1.md "Finding
 * the build"): /data/active-build.json names the build every icon, list and BiS URL hangs off,
 * and /data/dates.json carries the game dates. The site imports both at build time, so before
 * this they existed only in src/data and the CDN answered 404, leaving an app no build to read.
 */
export const ROOT_FILES = ['active-build.json', 'dates.json'];

async function publishRootFiles(webRoot) {
  const publicData = path.join(webRoot, 'public/data');
  await mkdir(publicData, { recursive: true });
  const published = [];
  for (const name of ROOT_FILES) {
    const source = path.join(webRoot, 'src/data', name);
    if (!(await exists(source))) continue;
    await cp(source, path.join(publicData, name));
    published.push(name);
  }
  return published;
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
