// web/scripts/sync-data.mjs
// Copies data/builds/<build>/ directories into web/ before dev, check, test and build:
//   public/data/<build>/…        fetched at runtime by the planner island
//   src/data/generated/*.json    imported at build time by src/pages/classes.astro
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
// this script throws and names the missing files. Until the data/ plan emits the Phase 1
// per-class directories, a build whose manifest lists none of them falls back to the
// checked-in fixture at src/fixtures/planner — loudly, and never when CF_PAGES is set
// (the same deploy guard astro.config.mjs uses for placeholder community links).
import { access, cp, mkdir, readFile, readdir, rm, writeFile } from 'node:fs/promises';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath, pathToFileURL } from 'node:url';

/** What gets copied into public/data/<build>/, in copy order. */
export const SYNC_ENTRIES = [
  { name: 'talents', kind: 'dir', required: true },
  { name: 'items', kind: 'dir', required: false },
  { name: 'icons', kind: 'dir', required: false },
  { name: 'sets.json', kind: 'file', required: false },
  { name: 'classes.json', kind: 'file', required: true },
  { name: 'races.json', kind: 'file', required: true },
  { name: 'combos.json', kind: 'file', required: true },
];

/** The files src/pages/classes.astro imports statically. */
export const PAGE_IMPORTS = ['classes.json', 'races.json', 'combos.json'];

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
 * Decides which directory to copy from: the real build directory when its manifest lists
 * Phase 1 talent data, or the checked-in fixture otherwise. Throws when the manifest is
 * missing listed files, or when falling back to the fixture is not allowed.
 * @param {{
 *   repoRoot: string,
 *   webRoot: string,
 *   build: string,
 *   allowFixture: boolean,
 *   log: { log: (...args: unknown[]) => void, warn: (...args: unknown[]) => void },
 * }} options
 */
async function resolveSourceDir({ repoRoot, webRoot, build, allowFixture, log }) {
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
  log.warn(
    `sync-data: data/builds/${build} has no Phase 1 talent data; copying the fixture from ` +
      `src/fixtures/planner instead. Talent and item content on this build is placeholder data.`,
  );
  return { sourceDir: path.join(webRoot, 'src/fixtures/planner'), usedFixture: true };
}

/**
 * Resets public/data/<build> and copies SYNC_ENTRIES from sourceDir into it. Throws when a
 * required entry is missing from sourceDir. Returns the names actually copied.
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
 * @param {{
 *   repoRoot: string,
 *   webRoot: string,
 *   builds: string[],
 *   log: { log: (...args: unknown[]) => void, warn: (...args: unknown[]) => void },
 * }} options
 */
async function publishRetained({ repoRoot, webRoot, builds, log }) {
  const published = [];
  for (const build of builds) {
    const { sourceDir } = await resolveSourceDir({
      repoRoot,
      webRoot,
      build,
      allowFixture: false,
      log,
    });
    await copyBuild({ repoRoot, webRoot, build, sourceDir });
    published.push(build);
  }
  return published;
}

/**
 * @param {{
 *   repoRoot: string,
 *   webRoot: string,
 *   allowFixture?: boolean,
 *   log?: { log: (...args: unknown[]) => void, warn: (...args: unknown[]) => void },
 * }} options
 */
export async function syncData({ repoRoot, webRoot, allowFixture = true, log = console } = {}) {
  const activeFile = path.join(webRoot, 'src/data/active-build.json');
  const { build } = await readJson(activeFile);
  if (typeof build !== 'string' || build.length === 0) {
    throw new Error(`${activeFile} must contain a non-empty "build" string.`);
  }

  const { sourceDir, usedFixture } = await resolveSourceDir({ repoRoot, webRoot, build, allowFixture, log });
  const copied = await copyBuild({ repoRoot, webRoot, build, sourceDir });
  await writePageImports({ webRoot, sourceDir });
  log.log(
    `sync-data: active build ${build} -> public/data/${build} (${copied.join(', ')})` +
      `${usedFixture ? ' [fixture]' : ''}`,
  );

  // The active build has already been published, with whatever fixture handling it needed.
  const { retained, skipped } = await partitionBuilds(repoRoot);
  for (const other of skipped.filter((name) => name !== build)) {
    log.warn(
      `sync-data: data/builds/${other} has no Phase 1 talent data; not publishing it. ` +
        `Share links made against build ${other} will not find their talent data.`,
    );
  }
  const published = [
    build,
    ...(await publishRetained({
      repoRoot,
      webRoot,
      builds: retained.filter((name) => name !== build),
      log,
    })),
  ];
  log.log(`sync-data: published ${published.length} build(s): ${published.join(', ')}`);

  return { build, sourceDir, usedFixture, copied, published };
}

const invokedDirectly =
  process.argv[1] !== undefined && import.meta.url === pathToFileURL(process.argv[1]).href;

if (invokedDirectly) {
  const webRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
  await syncData({
    repoRoot: path.resolve(webRoot, '..'),
    webRoot,
    // CF_PAGES is set only in the deploy build (see .github/workflows/web.yml), where
    // shipping fixture talent data to foreversixty.gg would be a lie.
    allowFixture: !process.env.CF_PAGES,
  });
}
