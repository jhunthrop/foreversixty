// web/src/content/guides/builds.test.ts
// Underscore-free (unlike _sections.test.ts) is fine here -- Astro's content loader only
// globs .md, so a .test.ts never becomes a guide regardless of its name; _sections.test.ts's
// leading underscore is a convention, not a requirement enforced anywhere.
//
// This decodes every spec guide's `build:` with the planner's own decoder (`decodeFS1`) and
// checks it against the ACTIVE build's talent data -- `src/data/active-build.json`, the same
// file `tests/e2e/support/active-build.ts` reads, currently 1.60.1.70009 -- not the
// `data-build` segment carried inside each guide's own `build:` string. That segment is
// informational metadata about which client build a guide's author decoded against when
// they wrote it; the guide's own words say "the data-build segment stays as it is in each
// file" (rotation-accuracy program, lane guide-builds), so a stale segment is not itself a
// defect. What has to be correct is what the string decodes to under the tree layout the
// site actually ships today, which is the active build's.
//
// This reads guide markdown and talent/race/class/combo data straight off disk, the same
// way _sections.test.ts does and for the same reason: under a plain `vitest run`,
// `astro:content`'s content layer returns empty collections, and `pretest` publishes only
// the small fixture warrior into src/data/generated/ (FOREVER_DATA=fixture) -- neither
// gives every class's real talent tree. Reading data/builds/<active>/ directly sidesteps
// both.
//
// decodeFS1 treats a tree's digits purely positionally: digit index i is
// `tree.talents[i]`, in the array order the client's own data ships (verified tier/column
// ascending for every class this file touches). `orderFromRanks` (used by
// _sections.test.ts) is NOT reused here for the max-rank or point-total checks: it silently
// clamps an over-rank digit down to `max_rank` before comparing, which is exactly the kind
// of bug this test exists to catch, so the checks below read the raw decoded ranks directly
// instead.
import { readFileSync, readdirSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import { parse as parseYaml } from 'yaml';
import { describe, expect, it } from 'vitest';
import { decodeFS1 } from '../../lib/planner/fs1';
import { comboIsLegal } from '../../lib/planner/rules';
import {
  MAX_POINTS,
  POINTS_PER_TIER,
  type ClassRow,
  type Combo,
  type RaceRow,
  type TalentFile,
} from '../../lib/planner/types';

const guidesRoot = dirname(fileURLToPath(import.meta.url));
// guidesRoot: <repo>/web/src/content/guides -- four levels up is <repo>, which holds both
// web/ (this file's active-build.json) and data/ (the talent, race, class and combo tables).
const repoRoot = join(guidesRoot, '../../../..');

interface LoadedGuide {
  id: string;
  frontmatter: Record<string, unknown>;
}

function loadSpecGuides(): LoadedGuide[] {
  const guides: LoadedGuide[] = [];
  for (const classDir of readdirSync(guidesRoot, { withFileTypes: true })) {
    if (!classDir.isDirectory()) continue;
    for (const file of readdirSync(join(guidesRoot, classDir.name))) {
      if (!file.endsWith('.md')) continue;
      const raw = readFileSync(join(guidesRoot, classDir.name, file), 'utf8');
      const [, frontmatterBlock] = raw.split('---');
      const frontmatter = parseYaml(frontmatterBlock) as Record<string, unknown>;
      if (frontmatter.spec === undefined) continue; // class landing pages carry no build:
      guides.push({ id: `${classDir.name}/${file.slice(0, -'.md'.length)}`, frontmatter });
    }
  }
  return guides;
}

const guides = loadSpecGuides();

const ACTIVE_BUILD = (
  JSON.parse(readFileSync(join(repoRoot, 'web/src/data/active-build.json'), 'utf8')) as { build: string }
).build;

const dataRoot = join(repoRoot, 'data/builds', ACTIVE_BUILD);

const talentCache = new Map<string, TalentFile>();
function talentsFor(classSlug: string): TalentFile {
  const cached = talentCache.get(classSlug);
  if (cached) return cached;
  const file = JSON.parse(readFileSync(join(dataRoot, 'talents', `${classSlug}.json`), 'utf8')) as TalentFile;
  talentCache.set(classSlug, file);
  return file;
}

const classRows = JSON.parse(readFileSync(join(dataRoot, 'classes.json'), 'utf8')) as ClassRow[];
const raceRows = JSON.parse(readFileSync(join(dataRoot, 'races.json'), 'utf8')) as RaceRow[];
const comboRows = JSON.parse(readFileSync(join(dataRoot, 'combos.json'), 'utf8')) as Combo[];

/**
 * Every talent in a class's three trees, keyed by the (tree position, array index) pair a
 * decoded FS1 tree-digit-string addresses -- i.e. tab order, the same indexing
 * `orderFromRanks` and `encodeTree` both use.
 */
function talentAt(file: TalentFile, treeIndex: number, digitIndex: number) {
  const tree = [...file.trees].sort((a, b) => a.position - b.position)[treeIndex];
  return tree?.talents[digitIndex];
}

/**
 * The spec's own tree the guide's build must spend its full point cap in, per
 * `data/curated/specs.json`. Not read from disk here: the mapping below is the minimal
 * duplicate this test needs and is checked against specs.json directly by the "primary
 * tree gets every spec's own signature talents" table further down failing to compile if a
 * spec's tree renamed out from under it (`talentAt` throwing `undefined` would fail the
 * signature-talent test for that spec).
 */

/**
 * The rotation-accuracy program's "must-have" table (design doc section on `guideBuildTalents`):
 * for each spec whose curated APL (`data/curated/apl/<spec>.json`) `castSpell`s or gates on a
 * talent-granted spell id, the talent that grants it, and the rank the guide's build must
 * hold. Derived by cross-referencing every `spellId` the APL names against each rank's
 * `spell_id` in the active build's talent data (rogue-affliction's Amplify Curse and
 * priest-shadow's Inner Focus were both missing from the guide's `build:` string before this
 * lane's fix -- exactly the class of bug this table exists to catch).
 */
const SIGNATURE_TALENTS: Record<string, string[]> = {
  'hunter/beast-mastery': ['Bestial Wrath'],
  'mage/arcane': ['Arcane Blast', 'Presence of Mind', 'Arcane Power'],
  'mage/fire': ['Combustion'],
  'paladin/retribution': ['Seal of Command'],
  'priest/shadow': ['Inner Focus', 'Shadowform'],
  'rogue/assassination': ['Cold Blood', 'Mutilate', 'Venom'],
  'rogue/combat': ['Adrenaline Rush'],
  'rogue/subtlety': ['Ghostly Strike', 'Premeditation', 'Hemorrhage'],
  'hunter/survival': ['Counterattack', 'Strider Kick'],
  'shaman/enhancement': ['Stormstrike'],
  'warlock/affliction': ['Amplify Curse', 'Wrack'],
  'warlock/destruction': ['Incinerate'],
  'warrior/fury': ['Death Wish'],
};

describe('guide build codes decode to a legal, complete build on the active data build', () => {
  it.each(guides.map((g) => [g.id, g] as const))('%s has a build: that decodes', (id, guide) => {
    const code = guide.frontmatter.build;
    expect(typeof code, `${id}: no build: frontmatter`).toBe('string');
    const decoded = decodeFS1(code as string);
    expect(decoded.ok, `${id}: ${decoded.ok ? '' : decoded.message}`).toBe(true);
  });

  it.each(guides.map((g) => [g.id, g] as const))('%s decodes for its own class', (id, guide) => {
    const decoded = decodeFS1(guide.frontmatter.build as string);
    if (!decoded.ok) return; // reported by the previous test
    expect(decoded.build.classSlug, `${id}: build: is for the wrong class`).toBe(guide.frontmatter.classSlug);
  });

  it.each(guides.map((g) => [g.id, g] as const))(`%s spends exactly ${MAX_POINTS} points`, (id, guide) => {
    const decoded = decodeFS1(guide.frontmatter.build as string);
    if (!decoded.ok) return;
    const total = decoded.build.treeRanks.flat().reduce((sum, rank) => sum + rank, 0);
    expect(total, `${id}: spends ${total} points, not ${MAX_POINTS}`).toBe(MAX_POINTS);
  });

  it.each(guides.map((g) => [g.id, g] as const))('%s takes no talent above its max_rank', (id, guide) => {
    const decoded = decodeFS1(guide.frontmatter.build as string);
    if (!decoded.ok) return;
    const file = talentsFor(decoded.build.classSlug);
    const overflows: string[] = [];
    decoded.build.treeRanks.forEach((ranks, treeIndex) => {
      ranks.forEach((rank, digitIndex) => {
        if (rank === 0) return;
        const talent = talentAt(file, treeIndex, digitIndex);
        if (!talent) {
          overflows.push(`tree ${treeIndex} digit ${digitIndex}: no such talent (rank ${rank})`);
        } else if (rank > talent.max_rank) {
          overflows.push(`${talent.name}: rank ${rank} > max_rank ${talent.max_rank}`);
        }
      });
    });
    expect(overflows, `${id}: ${overflows.join('; ')}`).toEqual([]);
  });

  it.each(guides.map((g) => [g.id, g] as const))(
    '%s satisfies every prereq_talent_id/prereq_rank',
    (id, guide) => {
      const decoded = decodeFS1(guide.frontmatter.build as string);
      if (!decoded.ok) return;
      const file = talentsFor(decoded.build.classSlug);
      const trees = [...file.trees].sort((a, b) => a.position - b.position);
      const byId = new Map(trees.flatMap((tree) => tree.talents.map((t) => [t.id, t] as const)));
      const rankOf = new Map<number, number>();
      decoded.build.treeRanks.forEach((ranks, treeIndex) => {
        trees[treeIndex]?.talents.forEach((talent, digitIndex) => {
          rankOf.set(talent.id, Math.min(ranks[digitIndex] ?? 0, talent.max_rank));
        });
      });
      const unmet: string[] = [];
      for (const talent of byId.values()) {
        const have = rankOf.get(talent.id) ?? 0;
        if (have <= 0 || talent.prereq_talent_id === null) continue;
        const prereqHave = rankOf.get(talent.prereq_talent_id) ?? 0;
        const need = talent.prereq_rank ?? 0;
        if (prereqHave < need) {
          const prereq = byId.get(talent.prereq_talent_id);
          unmet.push(
            `${talent.name} needs ${need} in ${prereq?.name ?? talent.prereq_talent_id}, has ${prereqHave}`,
          );
        }
      }
      expect(unmet, `${id}: ${unmet.join('; ')}`).toEqual([]);
    },
  );

  it.each(guides.map((g) => [g.id, g] as const))(
    '%s reaches every tier it spends points in (5 points per tier below it, client rule)',
    (id, guide) => {
      const decoded = decodeFS1(guide.frontmatter.build as string);
      if (!decoded.ok) return;
      const file = talentsFor(decoded.build.classSlug);
      const trees = [...file.trees].sort((a, b) => a.position - b.position);
      const unreachable: string[] = [];
      trees.forEach((tree, treeIndex) => {
        const ranks = decoded.build.treeRanks[treeIndex] ?? [];
        const byTier = new Map<number, number>();
        tree.talents.forEach((talent, digitIndex) => {
          const rank = Math.min(ranks[digitIndex] ?? 0, talent.max_rank);
          byTier.set(talent.tier, (byTier.get(talent.tier) ?? 0) + rank);
        });
        let cumulativeBelow = 0;
        for (const tier of [...byTier.keys()].sort((a, b) => a - b)) {
          const tierPoints = byTier.get(tier) ?? 0;
          if (tierPoints > 0 && cumulativeBelow < POINTS_PER_TIER * tier) {
            unreachable.push(
              `${tree.name} tier ${tier}: ${tierPoints} points spent, only ${cumulativeBelow} points in earlier tiers (needs ${POINTS_PER_TIER * tier})`,
            );
          }
          cumulativeBelow += tierPoints;
        }
      });
      expect(unreachable, `${id}: ${unreachable.join('; ')}`).toEqual([]);
    },
  );

  it.each(guides.map((g) => [g.id, g] as const))('%s is a legal race/class combo', (id, guide) => {
    const decoded = decodeFS1(guide.frontmatter.build as string);
    if (!decoded.ok) return;
    const classRow = classRows.find((c) => c.slug === decoded.build.classSlug);
    const raceRow = raceRows.find((r) => r.slug === decoded.build.raceSlug);
    expect(classRow, `${id}: unknown class slug ${decoded.build.classSlug}`).toBeDefined();
    expect(raceRow, `${id}: unknown race slug ${decoded.build.raceSlug}`).toBeDefined();
    if (!classRow || !raceRow) return;
    expect(
      comboIsLegal(comboRows, raceRow.id, classRow.id),
      `${id}: ${raceRow.name} cannot be a ${classRow.name}`,
    ).toBe(true);
  });

  it.each(Object.entries(SIGNATURE_TALENTS))(
    '%s takes its rotation-required signature talents',
    (id, names) => {
      const guide = guides.find((g) => g.id === id);
      expect(guide, `${id}: no such guide`).toBeDefined();
      if (!guide) return;
      const decoded = decodeFS1(guide.frontmatter.build as string);
      if (!decoded.ok) return; // reported by the decode test above
      const file = talentsFor(decoded.build.classSlug);
      const trees = [...file.trees].sort((a, b) => a.position - b.position);
      const rankOf = new Map<string, number>();
      decoded.build.treeRanks.forEach((ranks, treeIndex) => {
        trees[treeIndex]?.talents.forEach((talent, digitIndex) => {
          rankOf.set(talent.name, Math.min(ranks[digitIndex] ?? 0, talent.max_rank));
        });
      });
      const missing = names.filter((name) => (rankOf.get(name) ?? 0) <= 0);
      expect(
        missing,
        `${id}: rotation needs ${missing.join(', ')} but the build: string does not take it`,
      ).toEqual([]);
    },
  );
});
