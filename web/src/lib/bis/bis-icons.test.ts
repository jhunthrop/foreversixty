// web/src/lib/bis/bis-icons.test.ts
// Every item a published BiS list names (each pick, runner-up and per-source alternative) has its
// icon in the published build: a missing file is a broken image on /bis and the home example.
// The icons come from the pipeline (data/builds/<build>/icons) -- the fixture publish copies the
// ones the BiS lists name (scripts/sync-data.mjs publishBisIcons) -- so a gap here means the
// pipeline never generated that icon, never a reason to place a file by hand.
import { existsSync, readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { bisItemIds } from '../../../scripts/sync-data.mjs';
import { bisDirFor, bisSourceBuild } from '../../../scripts/bis-source-build.mjs';
import activeBuild from '../../data/active-build.json';

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const BUILD_DIR = path.join(ROOT, 'data/builds', activeBuild.build);
const PUBLISHED_ICONS = path.join(ROOT, 'web/public/data', activeBuild.build, 'icons');
const hasData = existsSync(path.join(BUILD_DIR, 'items')) && existsSync(PUBLISHED_ICONS);

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

function jsonFiles(dir: string): string[] {
  return readdirSync(dir)
    .filter((name) => name.endsWith('.json'))
    .map((name) => path.join(dir, name));
}

describe.skipIf(!hasData)('the published BiS lists', () => {
  it('have an icon for every item they name', () => {
    const iconOf = new Map<number, string>();
    for (const file of jsonFiles(path.join(BUILD_DIR, 'items'))) {
      for (const item of readJson<{ items: { id: number; icon: string }[] }>(file).items) {
        iconOf.set(item.id, item.icon);
      }
    }
    const gaps: string[] = [];
    for (const file of jsonFiles(bisDirFor(ROOT, bisSourceBuild(ROOT, activeBuild.build)))) {
      for (const id of bisItemIds(readJson<unknown>(file)) as Set<number>) {
        const icon = iconOf.get(id);
        if (icon === undefined) gaps.push(`${path.basename(file)}: item ${id} has no item row`);
        else if (!existsSync(path.join(PUBLISHED_ICONS, `${icon}.webp`))) {
          gaps.push(`${path.basename(file)}: item ${id} icon ${icon}.webp is missing`);
        }
      }
    }
    expect(gaps).toEqual([]);
  });
});
