// web/src/lib/bis/spec-icon.test.ts
// The spec-icon map is the generated spec list (`lib/sim/specs.ts`): every spec's icon is read
// off the active client (its tree's TalentTab icon, or its form spell's), and every surface that
// draws a spec icon reads it from there. The surfaces that render from Astro pages are held by
// tests/e2e/spec-icons.spec.ts; the data builders and the shared helpers are held here.
import { existsSync, readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { appGuideIndex, appGuides, appTiers } from '../app-data/app-files';
import { specIconSrc } from '../guild/raid-spec';
import { dataUrl } from '../planner/load';
import { SPEC_BY_KEY, SPEC_ICON_BUILD, SPECS } from '../sim/specs';
import { specIconByName } from '../ui/select-options';
import { loadLandingData } from './landing-build';
import { readSpecCatalog } from './load';
import { publishedIcon } from './spec-icon';

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const BUILD_DIR = path.join(ROOT, 'data/builds', activeBuild.build);
const ICONS = path.join(ROOT, 'web/public/data', activeBuild.build, 'icons');
const hasBuild = existsSync(path.join(BUILD_DIR, 'talents'));

interface CuratedRow {
  spec: string;
  form_spell?: number;
  icon?: string;
}

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

function clientIcon(row: CuratedRow, classSlug: string, treeIndex: number): string {
  if (row.form_spell !== undefined) {
    const spells = readJson<Record<string, string>>(path.join(BUILD_DIR, 'spellicons.json'));
    return spells[String(row.form_spell)];
  }
  const talents = readJson<{ trees: { position: number; icon: string }[] }>(
    path.join(BUILD_DIR, 'talents', `${classSlug}.json`),
  );
  const tree = talents.trees.find((t) => t.position === treeIndex);
  if (tree === undefined) throw new Error(`${row.spec}: no tree ${treeIndex}`);
  return tree.icon;
}

describe('the spec-icon map', () => {
  it('was generated from the active build', () => {
    expect(SPEC_ICON_BUILD).toBe(activeBuild.build);
  });

  it.skipIf(!hasBuild)("matches the client: each tree's TalentTab icon, or the form spell's", () => {
    const curated = readJson<CuratedRow[]>(path.join(ROOT, 'data/curated/specs.json'));
    for (const row of curated) {
      const spec = SPEC_BY_KEY[row.spec];
      expect(spec, row.spec).toBeDefined();
      expect(row.icon, `${row.spec}: the curated list names no icon`).toBeUndefined();
      expect(spec.icon, row.spec).toBe(clientIcon(row, spec.class_slug, spec.tree_index));
    }
  });

  it('draws Feral with Cat Form and Feral Bear with Dire Bear Form, apart', () => {
    const curated = readJson<CuratedRow[]>(path.join(ROOT, 'data/curated/specs.json'));
    const form = (spec: string) => curated.find((row) => row.spec === spec)?.form_spell;
    expect(form('druid-feral')).toBe(768);
    expect(form('druid-feral-bear')).toBe(9634);
    expect(SPEC_BY_KEY['druid-feral'].icon).not.toBe(SPEC_BY_KEY['druid-feral-bear'].icon);
  });

  it('gives every spec of a class its own icon', () => {
    for (const classSlug of new Set(SPECS.map((s) => s.class_slug))) {
      const icons = SPECS.filter((s) => s.class_slug === classSlug).map((s) => s.icon);
      expect(new Set(icons).size, classSlug).toBe(icons.length);
    }
  });

  it.skipIf(!existsSync(ICONS))('has art for every icon in the published build', () => {
    for (const spec of SPECS) expect(publishedIcon(activeBuild.build, spec.icon), spec.spec).toBe(spec.icon);
  });
});

describe('every surface resolves through the map', () => {
  it('the spec catalogue the pages read (BiS, guides, tiers, header, search, og)', () => {
    for (const entry of readSpecCatalog()) expect(entry.icon, entry.spec).toBe(SPEC_BY_KEY[entry.spec].icon);
  });

  it('the raid sign-up, rankings and select icons', () => {
    for (const spec of SPECS) {
      const src = dataUrl(activeBuild.build, `icons/${spec.icon}.webp`);
      expect(specIconSrc(spec), spec.spec).toBe(src);
      expect(specIconByName(spec.class_slug, spec.name), spec.spec).toBe(src);
    }
  });

  it.skipIf(!hasBuild)('the /bis landing cards', () => {
    for (const row of loadLandingData(activeBuild.build).cards.flatMap((card) => card.specs)) {
      expect(row.iconSrc, row.key).toBe(
        dataUrl(activeBuild.build, `icons/${SPEC_BY_KEY[row.key].icon}.webp`),
      );
    }
  });

  it.skipIf(!hasBuild)("the apps' guide, guide index and tier files", () => {
    for (const guide of appGuides(activeBuild.build))
      expect(guide.icon, guide.spec).toBe(SPEC_BY_KEY[guide.spec].icon);
    const index = appGuideIndex(activeBuild.build);
    for (const row of index.classes.flatMap((c) => c.specs)) {
      expect(row.icon, row.spec).toBe(SPEC_BY_KEY[row.spec].icon);
    }
    const tiers = appTiers(activeBuild.build);
    for (const roles of Object.values(tiers.factions)) {
      for (const list of Object.values(roles)) {
        for (const row of list.rows) expect(row.icon, row.spec).toBe(SPEC_BY_KEY[row.spec].icon);
      }
    }
  });

  it("the apps' bundled catalogue (web/src/data/specs.json)", () => {
    const bundled = readJson<{ spec: string; icon: string }[]>(path.join(ROOT, 'web/src/data/specs.json'));
    expect(bundled.map((row) => [row.spec, row.icon])).toEqual(SPECS.map((s) => [s.spec, s.icon]));
  });

  it('no web source picks a spec icon any other way', () => {
    // The retired paths: a tree's first talent stood in for its tab icon, and a hand map
    // overrode Feral. Anything that reads a tree or talent icon for a spec must go through
    // the map instead.
    const banned = [
      /specTabIcon/,
      /specCardIcon/,
      /SPEC_ICON_OVERRIDES/,
      /talents\[0\]\??\.icon/,
      /catformattack/,
    ];
    const offenders: string[] = [];
    const walk = (dir: string): void => {
      for (const entry of readdirSync(dir, { withFileTypes: true })) {
        const file = path.join(dir, entry.name);
        if (entry.isDirectory()) {
          if (!['fixtures', 'content'].includes(entry.name)) walk(file);
        } else if (/\.(ts|svelte|astro)$/.test(entry.name) && !entry.name.endsWith('.test.ts')) {
          const text = readFileSync(file, 'utf8');
          if (banned.some((re) => re.test(text))) offenders.push(path.relative(ROOT, file));
        }
      }
    };
    walk(path.join(ROOT, 'web/src'));
    expect(offenders).toEqual([]);
  });
});

describe('publishedIcon', () => {
  it('passes nothing through as nothing', () => {
    expect(publishedIcon(activeBuild.build, undefined)).toBeUndefined();
  });

  it('drops an icon this build does not publish', () => {
    expect(publishedIcon(activeBuild.build, 'no_such_icon_anywhere')).toBeUndefined();
  });
});
