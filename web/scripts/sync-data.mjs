// web/scripts/sync-data.mjs
// Copies one data/builds/<build>/ directory into web/ before dev, check, test and build:
//   public/data/<build>/…        fetched at runtime by the planner island
//   src/data/generated/*.json    imported at build time by src/pages/classes.astro
//
// The pipeline's manifest.json is the contract: every path it lists must exist on disk, or
// this script throws and names the missing files. Until the data/ plan emits the Phase 1
// per-class directories, a build whose manifest lists none of them falls back to the
// checked-in fixture at src/fixtures/planner — loudly, and never when CF_PAGES is set
// (the same deploy guard astro.config.mjs uses for placeholder community links).
import { access, cp, mkdir, readFile, rm, writeFile } from 'node:fs/promises';
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
  return JSON.parse(await readFile(file, 'utf8'));
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

  const buildDir = path.join(repoRoot, 'data/builds', build);
  const manifestFile = path.join(buildDir, 'manifest.json');
  const manifest = (await exists(manifestFile)) ? await readJson(manifestFile) : null;

  let sourceDir = buildDir;
  let usedFixture = false;

  if (manifest && hasPhaseOneData(manifest)) {
    const missing = await missingManifestFiles(manifest, buildDir);
    if (missing.length > 0) {
      throw new Error(
        `${path.relative(repoRoot, manifestFile)} lists files that are missing: ${missing.join(', ')}. ` +
          `Re-run the data pipeline for build ${build}.`,
      );
    }
  } else {
    if (!allowFixture) {
      throw new Error(
        `data/builds/${build} has no Phase 1 talent data (manifest.json lists no talents/*.json). ` +
          `Run the data pipeline before deploying; the fixture fallback is refused here.`,
      );
    }
    sourceDir = path.join(webRoot, 'src/fixtures/planner');
    usedFixture = true;
    log.warn(
      `sync-data: data/builds/${build} has no Phase 1 talent data; copying the fixture from ` +
        `src/fixtures/planner instead. Talent and item content on this build is placeholder data.`,
    );
  }

  const publicDir = path.join(webRoot, 'public/data', build);
  const generatedDir = path.join(webRoot, 'src/data/generated');
  await rm(publicDir, { recursive: true, force: true });
  await mkdir(publicDir, { recursive: true });
  await mkdir(generatedDir, { recursive: true });

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

  for (const name of PAGE_IMPORTS) {
    const body = await readFile(path.join(sourceDir, name), 'utf8');
    await writeFile(path.join(generatedDir, name), body, 'utf8');
  }

  log.log(
    `sync-data: build ${build} -> public/data/${build} (${copied.join(', ')})` +
      `${usedFixture ? ' [fixture]' : ''}`,
  );
  return { build, sourceDir, usedFixture, copied };
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
