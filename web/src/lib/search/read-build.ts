// web/src/lib/search/read-build.ts
// Build-time reads for the W3-C8 static files (`pages/data/[build]/…`): the active build's class
// items and talent files, straight off data/builds/<build>/ the way `lib/bis/load.ts` reads the
// BiS inputs. A build with no data yields empty lists, so a fixture checkout emits no files.
import { existsSync, readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import activeBuild from '../../data/active-build.json';
import type { TalentFile } from '../planner/types';
import { loadBisFile, loadLootFile, readSpecCatalog } from '../bis/load';
import type { ClassItemsInput } from './search-index';
import type { BisFile } from '../bis/types';

const REPO_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');

/** The build the site serves now; the W3-C8 files are emitted for it alone. */
export const ACTIVE_BUILD: string = activeBuild.build;

function buildDir(build: string): string {
  return path.join(REPO_ROOT, 'data/builds', build);
}

export function hasBuildData(build: string): boolean {
  return existsSync(path.join(buildDir(build), 'items')) && existsSync(path.join(buildDir(build), 'talents'));
}

function jsonFiles(dir: string): string[] {
  return existsSync(dir)
    ? readdirSync(dir)
        .filter((name) => name.endsWith('.json'))
        .sort()
    : [];
}

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

export function readClassItems(build: string): ClassItemsInput[] {
  const dir = path.join(buildDir(build), 'items');
  return jsonFiles(dir).map((name) => {
    const file = readJson<{ class_slug: string; items: ClassItemsInput['items'] }>(path.join(dir, name));
    return { classSlug: file.class_slug, items: file.items };
  });
}

export function readTalentFiles(build: string): TalentFile[] {
  const dir = path.join(buildDir(build), 'talents');
  return jsonFiles(dir).map((name) => readJson<TalentFile>(path.join(dir, name)));
}

export function readBisFiles(build: string): BisFile[] {
  return readSpecCatalog()
    .map((spec) => loadBisFile(spec.spec, build))
    .filter((file): file is BisFile => file !== null);
}

export { loadLootFile, readSpecCatalog };
