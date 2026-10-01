// web/src/lib/planner/band-compare.ts
// Planner rebuild spec §4.D/§7: the per-cell adapter behind `BandCompare.svelte`. A band's
// own `talents` field (`bis/<spec>.json`'s `BisBand.talents`) is the same per-tree digit
// string convention `MeCharacter.build.talents.trees`/an FS1 code's tree field already use
// (one base-36 digit per talent, tab order) -- this module never invents a second encoding.
//
// Two small, pure functions, each reusing an existing, unchanged primitive rather than
// forking a second traversal of the tree:
//  - `diffAgainstBand` walks the same tree/talent shape `lib/home/talent-delta.ts`'s
//    `talentDeltaFor` sums over (`Σ|build - band|` per talent cell), but keyed by talent id
//    instead of string position, so it can also report which ids differ -- the per-cell
//    marker `TalentCell.svelte` needs (spec §4.E.4).
//  - `loadFromBand` reconstructs a legal point order for a band's own ranks with
//    `orderFromRanks` (fs1.ts), the exact same reconstruction an addon import already runs.
import type { TalentIndex } from './rules';
import { orderFromRanks } from './fs1';

/** One base-36 digit per talent, tab order -- the convention every tree-rank string on the
 *  site already shares (FS1, `MeCharacter.build.talents.trees`, a BiS band's `talents`). */
export function bandTreeRanksFrom(bandTalents: string): number[][] {
  return bandTalents.split('-').map((tree) => [...tree].map((digit) => Number.parseInt(digit, 36) || 0));
}

export interface BandDiff {
  /** `Σ|build - band|` across every talent cell in every tree -- the same sum
   *  `talentDeltaFor` computes, generalised to talent ids. */
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
      const delta = Math.abs(buildRank - bandRank);
      if (delta > 0) {
        diffCount += delta;
        diffTalentIds.add(talent.id);
      }
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
