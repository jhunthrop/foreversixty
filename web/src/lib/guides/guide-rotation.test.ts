// web/src/lib/guides/guide-rotation.test.ts
// The rotation cards read the build they describe (the raid build, else the level 60 build)
// and no other: its cooldown cuts, its cast times, its pre-pull timings.
import activeBuild from '../../data/active-build.json';
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { decodeBuildTrees, type DecodedTreeView } from './build-tree';
import { describedBuildTrees, guideRotation } from './guide-rotation';
import { gateWords, type AplGate } from '../bis/apl-gates';
import { bandTopLevel, loadRotations, rotationEntryFor, rotationLinesFor } from '../bis/rotation-view';
import {
  litTalentsOf,
  loadApl,
  markRaidBuildRows,
  raidBuildOnlyTalents,
  prepullCasts,
  spellLookupFor,
  talentCastTimeCuts,
  talentNamesOf,
  type GuideRole,
} from './rotation-lists';

const BUILD = activeBuild.build;
const GUIDES_DIR = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../content/guides');

function guideTrees(file: string): { leveling: DecodedTreeView[]; raid: DecodedTreeView[] | undefined } {
  const raw = readFileSync(path.join(GUIDES_DIR, file), 'utf8');
  const code = (key: string): string | undefined => new RegExp(`^${key}: '([^']+)'`, 'm').exec(raw)?.[1];
  const leveling = decodeBuildTrees(BUILD, code('build') ?? '') ?? [];
  const raidCode = code('raidBuild');
  const raid =
    raidCode === undefined || raidCode === code('build') ? undefined : decodeBuildTrees(BUILD, raidCode);
  return { leveling, raid };
}

function rankOf(trees: readonly DecodedTreeView[], name: string): number {
  return trees.flatMap((tree) => tree.talents).find((talent) => talent.name === name)?.rank ?? 0;
}

describe('guideRotation', () => {
  it('gives Marksmanship’s Rapid Fire the raid build’s cooldown, not the level 60 build’s Rapid Killing', () => {
    const { leveling, raid } = guideTrees('hunter/marksmanship.md');
    const trees = describedBuildTrees(leveling, raid);
    const rapidFire = guideRotation(BUILD, 'hunter', 'hunter-marksmanship', 'dps', trees)?.cooldowns.find(
      (item) => item.name === 'Rapid Fire',
    );
    const cut = rankOf(trees, 'Rapid Killing');
    expect(rapidFire?.clause.startsWith(`${5 - cut} min cooldown`)).toBe(true);
  });
});

interface GuideCard {
  label: string;
  classSlug: string;
  spec: string;
  role: GuideRole;
  trees: DecodedTreeView[];
}

/** Every guide's leveling build and, when it has its own, its raid build. */
function everyGuideBuild(): GuideCard[] {
  return readdirSync(GUIDES_DIR, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .flatMap((dir) =>
      readdirSync(path.join(GUIDES_DIR, dir.name))
        .filter((file) => file.endsWith('.md') && !file.startsWith('_') && file !== 'index.md')
        .flatMap((file): GuideCard[] => {
          const raw = readFileSync(path.join(GUIDES_DIR, dir.name, file), 'utf8');
          const role = (/^role: (\w+)/m.exec(raw)?.[1] ?? 'dps') as GuideRole;
          const spec = `${dir.name}-${file.slice(0, -'.md'.length)}`;
          const { leveling, raid } = guideTrees(`${dir.name}/${file}`);
          const base = { classSlug: dir.name, spec, role };
          return [
            { ...base, label: `${spec} level 60`, trees: leveling },
            ...(raid === undefined ? [] : [{ ...base, label: `${spec} raid`, trees: raid }]),
          ];
        }),
    );
}

describe('every guide card', () => {
  // Sim-accuracy round 18: a card describes one build, so a cast whose talent gate rules that
  // build out is not on it (Arcane's raid card listed the level 60 build's pre-pull Frostbolt,
  // gated on not knowing Heating Up, at the raid build's 3 s).
  it('lists no cast whose talent gate rules out the build it describes', () => {
    const cards = everyGuideBuild();
    expect(cards.length).toBeGreaterThan(30);
    for (const card of cards) {
      const lists = guideRotation(BUILD, card.classSlug, card.spec, card.role, card.trees);
      if (lists === undefined) continue;
      const facts = spellLookupFor(BUILD, card.classSlug);
      const talents = talentNamesOf(BUILD, card.classSlug);
      const lit = litTalentsOf(card.trees);
      const runs = (gates: readonly AplGate[]): boolean =>
        gates.every((gate) => {
          const name = gate.kind === 'known' ? facts(gate.spellId)?.name : undefined;
          return name === undefined || !talents.has(name) || lit.has(name) === gate.positive;
        });
      const casts = prepullCasts(loadApl(card.spec));
      for (const item of lists.opener) {
        const mine = casts.filter((cast) => cast.spellId === item.spellId);
        expect(
          mine.some((cast) => runs(cast.gates)),
          `${card.label} opener ${item.name}`,
        ).toBe(true);
      }
      const top = bandTopLevel(60);
      const entry = rotationEntryFor(loadRotations(BUILD, card.spec) ?? [], top)!;
      const kept = rotationLinesFor(entry, BUILD, card.classSlug, card.spec, top, runs);
      expect(lists.priority.length, `${card.label} priority`).toBe(kept.length);
    }
  });

  it('gives Arcane’s raid card Pyroblast alone, and its level 60 card the 2.5 s Frostbolt', () => {
    const { leveling, raid } = guideTrees('mage/arcane.md');
    const opener = (trees: DecodedTreeView[]) =>
      guideRotation(BUILD, 'mage', 'mage-arcane', 'dps', trees)?.opener.map(
        (item) => `${item.name}: ${item.clause}`,
      );
    expect(opener(raid!)).toEqual(['Pyroblast: 6 s before the pull']);
    expect(opener(leveling)).toEqual(['Frostbolt: 2.5 s before the pull']);
  });

  it('tags Arcane’s raid rows that need Heating Up as raid build only', () => {
    const { leveling, raid } = guideTrees('mage/arcane.md');
    const lists = guideRotation(BUILD, 'mage', 'mage-arcane', 'dps', raid!)!;
    const only = raidBuildOnlyTalents(
      talentNamesOf(BUILD, 'mage'),
      litTalentsOf(leveling),
      litTalentsOf(raid!),
    );
    const tagged = markRaidBuildRows(lists, only)
      .priority.filter((item) => item.raidBuildOnly === true)
      .map((item) => item.name);
    expect(tagged).toEqual(expect.arrayContaining(['Fire Blast', 'Fireball']));
    expect(tagged).not.toContain('Arcane Blast');
  });

  // Round 19: the card's rows are already the described build's own, so a talent it takes or
  // leaves is not worded as a condition ("Fireball: ..., with Heating Up" read as if Fireball
  // waited on the stacks it builds).
  it('words no talent gate the build it describes settles', () => {
    for (const card of everyGuideBuild()) {
      const lists = guideRotation(BUILD, card.classSlug, card.spec, card.role, card.trees);
      if (lists === undefined) continue;
      const facts = spellLookupFor(BUILD, card.classSlug);
      const talents = talentNamesOf(BUILD, card.classSlug);
      const worded = (gates: readonly AplGate[]): string[] =>
        gates.flatMap((gate) => {
          const name = gate.kind === 'known' ? facts(gate.spellId)?.name : undefined;
          return name !== undefined && talents.has(name) ? [gateWords(gate, name)] : [];
        });
      const casts = prepullCasts(loadApl(card.spec));
      const top = bandTopLevel(60);
      const entry = rotationEntryFor(loadRotations(BUILD, card.spec) ?? [], top)!;
      const lines = rotationLinesFor(entry, BUILD, card.classSlug, card.spec, top);
      const rows = [
        ...lists.opener.map((item) => ({
          item,
          words: casts.filter((cast) => cast.spellId === item.spellId).flatMap((cast) => worded(cast.gates)),
        })),
        ...lists.priority.map((item) => ({
          item,
          words: lines
            .filter((line) => line.spellId === item.spellId)
            .flatMap((line) => worded(line.gates ?? [])),
        })),
      ];
      for (const { item, words } of rows) {
        for (const word of words) {
          expect(
            item.clause.toLowerCase().endsWith(word.toLowerCase()),
            `${card.label} ${item.name}: ${item.clause}`,
          ).toBe(false);
        }
      }
    }
  });
});

describe('describedBuildTrees', () => {
  it('is the raid build when the guide has its own, else the level 60 build', () => {
    const leveling = [{ name: 'L', position: 0, talents: [], points: 0 }];
    const raid = [{ name: 'R', position: 0, talents: [], points: 0 }];
    expect(describedBuildTrees(leveling, raid)).toBe(raid);
    expect(describedBuildTrees(leveling, undefined)).toBe(leveling);
  });
});

describe('talentCastTimeCuts', () => {
  it('reads lasting cuts and not a proc’s', () => {
    const cuts = talentCastTimeCuts(
      BUILD,
      'mage',
      litTalentsOf([
        {
          talents: [
            { name: 'Improved Frostbolt', rank: 5 },
            { name: 'Heating Up', rank: 1 },
          ],
        },
      ]),
    );
    expect(cuts.get('Frostbolt')).toBe(500);
    expect(cuts.has('Pyroblast')).toBe(false);
  });
});
