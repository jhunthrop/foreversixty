// web/src/lib/planner/band-compare.ts
// Planner rebuild spec §4.D/§7: the per-cell adapter behind `BandCompare.svelte`. A band's
// own `talents` field (`bis/<spec>.json`'s `BisBand.talents`) is the same per-tree digit
// string convention `MeCharacter.build.talents.trees`/an FS1 code's tree field already use
// (one base-36 digit per talent, tab order) -- this module never invents a second encoding.
//
// Two small, pure functions, each reusing an existing, unchanged primitive rather than
// forking a second traversal of the tree:
//  - `diffAgainstBand` walks the same tree/talent shape `lib/home/talent-delta.ts`'s loop does,
//    summing the band's unmatched points per talent cell, keyed by talent id instead of string position, so it can also report which ids differ -- the per-cell
//    marker `TalentCell.svelte` needs (spec §4.E.4).
//  - `loadFromBand` reconstructs a legal point order for a band's own ranks with
//    `orderFromRanks` (fs1.ts), the exact same reconstruction an addon import already runs.
import type { TalentIndex } from './rules';
import { specKeyFor } from '../addon/score';
import { specDisplayName, specRow } from '../sim/spec-label';
import { bandBuildTitle } from './band-title';
import { orderFromRanks } from './fs1';
import { unmatchedBandPoints } from '../home/talent-delta';

/** One base-36 digit per talent, tab order -- the convention every tree-rank string on the
 *  site already shares (FS1, `MeCharacter.build.talents.trees`, a BiS band's `talents`). */
export function bandTreeRanksFrom(bandTalents: string): number[][] {
  return bandTalents.split('-').map((tree) => [...tree].map((digit) => Number.parseInt(digit, 36) || 0));
}

export interface BandDiff {
  /** The band's points the build has not matched (`Σ max(band - build, 0)`, the home card's
   *  `talentDeltaFor` rule), so it never exceeds the band's own point total. */
  diffCount: number;
  /** Talent ids whose rank differs from the band's own rank at that cell. */
  diffTalentIds: ReadonlySet<number>;
  /** The band's own rank at every talent id in the index, for the tooltip line. */
  bandRankById: ReadonlyMap<number, number>;
}

/**
 * `ranks` is the build's own `talentId -> rank` map (`PlannerStore.ranks`, from
 * `ranksByTalent`). Walks `index.trees[].talents[]` in the same tab order the digit string
 * is written in -- the same shape `talentDeltaFor`'s own loop walks, just by talent id
 * rather than by string position, so the ids differing can be kept alongside the sum.
 */
export function diffAgainstBand(
  index: TalentIndex,
  ranks: ReadonlyMap<number, number>,
  bandTalents: string,
): BandDiff {
  const bandTrees = bandTreeRanksFrom(bandTalents);
  let diffCount = 0;
  const diffTalentIds = new Set<number>();
  const bandRankById = new Map<number, number>();

  index.trees.forEach((tree, treeIndex) => {
    const bandDigits = bandTrees[treeIndex] ?? [];
    tree.talents.forEach((talent, talentIndex) => {
      const buildRank = ranks.get(talent.id) ?? 0;
      const bandRank = Math.min(bandDigits[talentIndex] ?? 0, talent.max_rank);
      bandRankById.set(talent.id, bandRank);
      diffCount += unmatchedBandPoints(buildRank, bandRank);
      if (buildRank !== bandRank) diffTalentIds.add(talent.id);
    });
  });

  return { diffCount, diffTalentIds, bandRankById };
}

/** `TreeGrid`/`TalentCell`'s own view of a loaded band diff (spec §4.E.4): the label for
 *  the tooltip's extra line, plus `BandDiff`'s own two maps. A thin alias rather than a
 *  new shape -- every field here is `BandDiff`'s own. */
export interface BandDiffView extends BandDiff {
  bandLabel: string;
}

export interface BandLoad {
  /** A legal point order reaching the band's own ranks, as far as they are reachable. */
  order: number[];
  /** Ranks the reconstruction could not legally reach (same meaning as an addon import's
   *  own dropped points -- `Planner.svelte`'s `codeNote` already has a message for this). */
  dropped: number[];
}

/** The band's own build, as a loadable point order -- `orderFromRanks`, unchanged, fed from
 *  a band's digit string instead of a decoded FS1 code (spec §4.D's "load mechanism"). */
export function loadFromBand(index: TalentIndex, bandTalents: string): BandLoad {
  const { order, dropped } = orderFromRanks(index, bandTreeRanksFrom(bandTalents));
  return { order, dropped };
}

export interface BandLoadAtLevel extends BandLoad {
  /** The points the band's whole build holds (before any cut); fewer in `order` means it was cut. */
  fullPoints: number;
}

/**
 * The band's build for a character at `characterLevel` (null: none selected). The point order
 * is cut at the character's level (points = level - 9) when that is fewer than the band's
 * whole build; a character past the band, or none, gets the whole build. A character too low
 * to have a point (below level 10) also gets the whole build rather than an empty one.
 */
export function loadFromBandAtLevel(
  index: TalentIndex,
  bandTalents: string,
  characterLevel: number | null,
): BandLoadAtLevel {
  const full = loadFromBand(index, bandTalents);
  const fullPoints = full.order.length;
  const allowed = (characterLevel ?? 9) - 9;
  const order = allowed > 0 ? full.order.slice(0, allowed) : full.order;
  return { ...full, order, fullPoints };
}

/** The title for a band build just applied to a draft of `classSlug` with `split` points per tree. */
export function bandTitleForDraft(bandLabel: string, classSlug: string, split: readonly number[]): string {
  const key = specKeyFor(classSlug, split);
  return bandBuildTitle(bandLabel, specRow(key) === null ? null : specDisplayName(key));
}

/** The band a visitor with no character and no points spent is shown: the level-60 build,
 *  the end state every other band leads to, so a first look at the planner answers "what does
 *  a finished build look like" instead of "level 9, nothing to compare". The first band with
 *  talents (10 to 19) is the other reading; it is not used because that band's build is a few
 *  points of a tree the visitor has not chosen, and its DPS says nothing about the class. */
export const VISITOR_DEFAULT_BAND = 60;

/** The band the panel compares against: the visitor default while nothing is spent and no
 *  character is selected, else the band of the level the build reaches. */
export function bandForPanel(
  pointsSpent: number,
  characterLevel: number | null,
  reachedBand: number,
): number {
  return pointsSpent === 0 && characterLevel === null ? VISITOR_DEFAULT_BAND : reachedBand;
}
