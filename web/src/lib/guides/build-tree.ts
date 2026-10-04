// web/src/lib/guides/build-tree.ts
// Build-time (not browser) talent decode for two server-rendered, non-interactive previews
// the spec guide rebuild adds (spec §4.C.1's rail thumbnail, §4.D's leveling band-strip
// thumbnail): the real, loadable tree `GuideBuildTree.svelte` renders client-side is still
// the only INTERACTIVE tree on the page (unchanged island) -- these two previews share its
// own decoder (`decodeFS1`, the FS1 format's own talent-rank grammar) and the same
// tier/column sort `GuideBuildTree`'s `TreeGrid` renders by, but read the class's talent
// file straight off disk the way `rotation-view.ts`/`spec-icon.ts` already do for the same
// reason (Astro's static pages run at build time in Node, never in a browser, so
// `planner/load.ts`'s `fetch`-based `loadTalents` cannot run here).
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { decodeFS1 } from '../planner/fs1';
import type { Talent, TalentFile } from '../planner/types';

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

function loadTalentFile(build: string, classSlug: string): TalentFile | undefined {
  const file = path.join(repoRoot(), 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(file)) return undefined;
  return JSON.parse(readFileSync(file, 'utf8')) as TalentFile;
}

export interface DecodedTalentView extends Talent {
  rank: number;
}

export interface DecodedTreeView {
  name: string;
  position: number;
  /** Sorted tier ascending, then column ascending -- the same order the FS1 code's own
   *  per-tree digit string assumes (`fs1.ts`'s own header) and `TreeGrid` renders by. */
  talents: DecodedTalentView[];
  points: number;
}

/**
 * Every tree of `code`'s own build, ranks decoded and attached -- `undefined` when the code
 * does not parse (`decodeFS1`) or this build carries no talent file for the code's own
 * class. A caller choosing a single "the spec's own tree" thumbnail (the rail's Build card)
 * reads the array back by `classSlug`-derived position, never a hardcoded tree name (see
 * `primaryTree` below).
 */
export function decodeBuildTrees(build: string, code: string): DecodedTreeView[] | undefined {
  const decoded = decodeFS1(code);
  if (!decoded.ok) return undefined;
  const talentFile = loadTalentFile(build, decoded.build.classSlug);
  if (talentFile === undefined) return undefined;
  return talentFile.trees.map((tree, index): DecodedTreeView => {
    const sorted = [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column);
    const ranks = decoded.build.treeRanks[index] ?? [];
    const talents = sorted.map((talent, talentIndex) => ({ ...talent, rank: ranks[talentIndex] ?? 0 }));
    return {
      name: tree.name,
      position: tree.position,
      talents,
      points: talents.reduce((sum, talent) => sum + talent.rank, 0),
    };
  });
}

/** The tree with the most points spent -- the guide's own spec tree, for the rail's Build
 *  card thumbnail (spec §4.C.1: "a thumbnail of the build's own tree", singular, fitting a
 *  120px-tall card -- the full three-column tree already renders below, unchanged, under
 *  "Talents and builds"). A tie (should not occur for a real spec build, which always
 *  commits most of its points to one tree) keeps the first tree in tree order. */
export function primaryTree(trees: readonly DecodedTreeView[]): DecodedTreeView | undefined {
  if (trees.length === 0) return undefined;
  return trees.reduce((best, tree) => (tree.points > best.points ? tree : best));
}

/** Every talent across all three trees with `rank > 0` (the leveling band strip's own
 *  thumbnail, spec §4.D finding 5) -- one flat, tree-unordered list the caller wraps as a
 *  row of icons; tree order is preserved among ties since `decodeBuildTrees` already
 *  returns trees in the file's own `position` order. */
export function litTalents(trees: readonly DecodedTreeView[]): DecodedTalentView[] {
  return trees.flatMap((tree) => tree.talents.filter((talent) => talent.rank > 0));
}

/**
 * A leveling-BiS band's own `talents` string (`BisBand.talents`, e.g.
 * `"35311103002000000-353211005050010051-000000000000000000"`) -- a different, simpler
 * grammar than an FS1 code's own tree field: three dash-separated, fixed-width digit
 * strings, one digit per talent, zero-padded rather than trailing-zero-trimmed (confirmed
 * against the real build's own BiS files). `undefined` when this build carries no talent
 * file for `classSlug`, or `talents` does not split into exactly three segments.
 */
export function decodeBandTalents(
  build: string,
  classSlug: string,
  talents: string,
): DecodedTreeView[] | undefined {
  const segments = talents.split('-');
  if (segments.length !== 3) return undefined;
  const talentFile = loadTalentFile(build, classSlug);
  if (talentFile === undefined) return undefined;
  return talentFile.trees.map((tree, index): DecodedTreeView => {
    const sorted = [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column);
    const digits = segments[index] ?? '';
    const ranked = sorted.map((talent, talentIndex) => ({
      ...talent,
      rank: digits[talentIndex] === undefined ? 0 : Number.parseInt(digits[talentIndex]!, 10) || 0,
    }));
    return {
      name: tree.name,
      position: tree.position,
      talents: ranked,
      points: ranked.reduce((sum, talent) => sum + talent.rank, 0),
    };
  });
}
