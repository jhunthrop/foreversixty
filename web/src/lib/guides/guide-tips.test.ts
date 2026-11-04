import { existsSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { talentTipModel, tipFromBundle } from '../tooltips/game-tip';
import { buildFile, repoRoot } from './build-files';
import { decodeBuildTrees } from './build-tree';
import { lineTipKey, proseNames, racialKeyFor, specGuideData } from './guide-tips';
import { linkProse } from './prose-tips';

const build = activeBuild.build;
const GUIDES = path.join(repoRoot(), 'web/src/content/guides');
const hasData = existsSync(buildFile(build, 'spelltips', 'shared.json'));

/** Every spec guide on disk: its class, spec, races, build code and body. */
function specGuides(): { classSlug: string; spec: string; races: string[]; code?: string; body: string }[] {
  return readdirSync(GUIDES, { withFileTypes: true })
    .filter((dir) => dir.isDirectory())
    .flatMap((dir) =>
      readdirSync(path.join(GUIDES, dir.name))
        .filter((file) => file.endsWith('.md'))
        .map((file) => {
          const text = readFileSync(path.join(GUIDES, dir.name, file), 'utf8');
          const front = /^---\n([\s\S]*?)\n---\n([\s\S]*)$/.exec(text);
          const yaml = front?.[1] ?? '';
          const races = /recommendedRaces:\s*\[([^\]]*)\]/.exec(yaml)?.[1] ?? '';
          return {
            classSlug: dir.name,
            spec: file.replace(/\.md$/, ''),
            races: races
              .split(',')
              .map((race) => race.trim().replace(/['"]/g, ''))
              .filter(Boolean),
            code: /\nbuild:\s*['"]?([^'"\n]+)/.exec(`\n${yaml}`)?.[1],
            body: front?.[2] ?? '',
          };
        }),
    )
    .filter((guide) =>
      /\nspec:/.test(`\n${readFileSync(path.join(GUIDES, guide.classSlug, `${guide.spec}.md`), 'utf8')}`),
    );
}

describe.skipIf(!hasData)('the guide tooltip bundle, on the active build', () => {
  const guides = specGuides();

  it('finds the spec guides', () => {
    expect(guides.length).toBeGreaterThan(20);
  });

  it.each(guides.map((guide) => [`${guide.classSlug}-${guide.spec}`, guide] as const))(
    '%s: every talent, gear piece, enchant and consumable it shows opens a tooltip',
    (_label, guide) => {
      const spec = `${guide.classSlug}-${guide.spec}`;
      const { bis, extras, bundle } = specGuideData(build, guide.classSlug, spec, guide.races);
      const trees = guide.code === undefined ? [] : (decodeBuildTrees(build, guide.code) ?? []);
      for (const tree of trees) {
        for (const talent of tree.talents) {
          expect(tipFromBundle(bundle, `talent:${talent.id}:${talent.rank}`), talent.name).not.toBeNull();
        }
      }
      for (const row of (bis?.panels ?? []).flatMap((panel) => panel.rows)) {
        if (row.itemId !== undefined) expect(bundle.items[String(row.itemId)], row.itemName).toBeDefined();
      }
      for (const row of extras?.enchants ?? []) {
        expect(bundle.tips[`enchant:${row.id}`], row.name).toBeDefined();
      }
      // Every name the prose matcher can mark resolves in the same bundle.
      for (const [name, target] of proseNames(build, guide.classSlug, trees, bundle)) {
        if ('key' in target) expect(tipFromBundle(bundle, target.key), name).not.toBeNull();
        else expect(bundle.items[String(target.itemId)], name).toBeDefined();
      }
    },
  );

  it.each(guides.map((guide) => [`${guide.classSlug}-${guide.spec}`, guide] as const))(
    '%s: no tooltip prints a raw client token, and every one has its icon',
    (_label, guide) => {
      const spec = `${guide.classSlug}-${guide.spec}`;
      const { bundle } = specGuideData(build, guide.classSlug, spec, guide.races);
      const models = [
        ...Object.values(bundle.tips),
        ...Object.values(bundle.talents).flatMap((entry) =>
          Array.from({ length: entry.maxRank + 1 }, (_, rank) => talentTipModel(entry, rank)),
        ),
      ];
      for (const model of models) {
        expect(JSON.stringify(model), model.name).not.toMatch(/\$/);
        expect(model.icon, model.name).toBeDefined();
      }
      for (const item of Object.values(bundle.items)) {
        expect(JSON.stringify(item), item.name).not.toMatch(/\$/);
        expect(item.icon, item.name).not.toBe('');
      }
    },
  );

  it('marks the abilities a Fury warrior guide names, at their top rank', () => {
    const { bundle } = specGuideData(build, 'warrior', 'warrior-fury', ['orc']);
    const names = proseNames(build, 'warrior', [], bundle);
    const html = linkProse('<p>Bloodthirst on cooldown, Whirlwind, then Execute below 20%.</p>', names);
    const keys = [...html.matchAll(/data-game-tip="([^"]+)"/g)].map((m) => m[1]!);
    expect(keys).toHaveLength(3);
    for (const key of keys) {
      const tip = tipFromBundle(bundle, key);
      expect(tip?.description, key).not.toBe('');
    }
    expect(tipFromBundle(bundle, keys[2]!)?.subtitle).toMatch(/^Rank \d+$/);
  });

  it('opens a racial cooldown line on the racial, and a racial by its race-data name', () => {
    const { bundle } = specGuideData(build, 'warrior', 'warrior-fury', ['orc']);
    expect(lineTipKey(bundle, undefined, 'Blood Fury')).toBe('racial:Blood Fury');
    expect(racialKeyFor(bundle, 'blood fury')).toBe('racial:Blood Fury');
    expect(tipFromBundle(bundle, 'racial:Blood Fury')?.subtitle).toBe('Racial');
  });
});
