// web/src/lib/planner/editor-trees.ts
// A class's talent trees as ui/TalentTrees draws them (spec 2026-11-08 sections 4.4 and 5.4,
// decisions P2 to P5, ruling P4; talent-tree lane 2026-11-10). `editorTrees`: the planner's, every
// cell's state in the game's colours, with a character's level cap greying the unspent cells (P3)
// and its tooltip saying why; a locked cell's tooltip naming what it needs (P5); the gold dot and
// band line on a cell whose rank differs from a compared band; the planned-ahead mark on a cell
// holding points past the character's level. `viewTrees`: any read-only build (a guide, the sim,
// /sim/talents) over the same frame, so every tree on the site is drawn from one builder.
import { SPECS } from '../sim/specs';
import { cellState, type TalentCellView, type TalentTreeView, viewState } from '../ui/talent-tree';
import type { BandDiffView } from './band-compare';
import { connectorsFor, gridViewBox } from './connectors';
import { bandCompareCopy, plannerCopy } from './copy';
import { pointsAvailableAt, ranksByTalent } from './derive';
import { gridCells, gridSize } from './grid';
import { canAddPoint, type TalentIndex } from './rules';
import type { Talent, TalentTree } from './types';
import { MAX_POINTS } from './types';

/** The talent frame's tiers: every class's trees fill seven. */
export const GAME_TIERS = 7;

/** A data file's URL (`dataUrl(build, path)`). */
type Asset = (path: string) => string;

export interface EditorTreesInput {
  index: TalentIndex;
  order: readonly number[];
  classSlug: string;
  /** The character's level, or null: its points stop there (P3) and the rest are planned (P4). */
  characterLevel: number | null;
  readOnly: boolean;
  bandDiff: BandDiffView | null;
  asset: Asset;
}

export interface ViewTreesInput {
  index: TalentIndex;
  /** Talent id → rank; a talent it does not name is untaken. */
  ranks: ReadonlyMap<number, number>;
  classSlug: string;
  asset: Asset;
  /** Carry every rank's text on the cells, so the trees open their own tooltips. False on a page
   *  whose tooltip host serves the text (a guide's bundle): the island's props stay small. */
  withText?: boolean;
}

/** A tree's head icon: its spec's icon from the generated spec list, '' when it has none. */
export function treeIconSrc(classSlug: string, treeIndex: number, asset: Asset): string {
  const icon = SPECS.find((row) => row.class_slug === classSlug && row.tree_index === treeIndex)?.icon;
  return icon === undefined ? '' : asset(`icons/${icon}.webp`);
}

/** A cell's place, icon and rank, the parts every tree shares. */
function baseCell(
  talent: Talent,
  tier: number,
  column: number,
  rank: number,
  asset: Asset,
  withText: boolean,
): Omit<TalentCellView, 'state'> {
  return {
    id: talent.id,
    tier,
    column,
    name: talent.name,
    icon: talent.icon,
    iconSrc: asset(`icons/${talent.icon}.webp`),
    rank,
    max: talent.max_rank,
    ...(withText ? { ranks: talent.ranks.map((row) => row.description) } : {}),
  };
}

/** The tree around its cells: head, art, the seven-tier grid and the arrows at these ranks. */
function treeFrame(
  tree: TalentTree,
  treeIndex: number,
  classSlug: string,
  ranks: ReadonlyMap<number, number>,
  cells: TalentCellView[],
  asset: Asset,
): TalentTreeView {
  // Every tree is drawn at the game frame's seven tiers (a trimmed tree too), so a tree's height
  // is the same everywhere and the planner's loading reserve holds it (spec 5.2).
  const size = { ...gridSize(tree), tiers: Math.max(GAME_TIERS, gridSize(tree).tiers) };
  const byId = new Map(tree.talents.map((talent) => [talent.id, talent]));
  return {
    id: tree.id,
    name: tree.name,
    iconSrc: treeIconSrc(classSlug, treeIndex, asset),
    points: cells.reduce((sum, cell) => sum + cell.rank, 0),
    ...(tree.background ? { artSrc: asset(`trees/${tree.background}.webp`) } : {}),
    columns: size.columns,
    tiers: size.tiers,
    viewBox: gridViewBox(size.tiers, size.columns),
    connectors: connectorsFor(tree, new Map(ranks)).map((c) => ({
      d: c.d,
      met: c.met,
      from: c.from,
      needs: byId.get(c.to)?.prereq_rank ?? 0,
    })),
    cells,
  };
}

/** A read-only build's trees: untaken greyed, partial green, maxed gold, arrows met in gold. */
export function viewTrees({
  index,
  ranks,
  classSlug,
  asset,
  withText = true,
}: ViewTreesInput): TalentTreeView[] {
  return index.trees.map((tree, treeIndex) => {
    const cells = gridCells(tree).map(({ talent, tier, column }) => {
      const rank = Math.min(talent.max_rank, ranks.get(talent.id) ?? 0);
      return {
        ...baseCell(talent, tier, column, rank, asset, withText),
        state: viewState(rank, talent.max_rank),
      };
    });
    return treeFrame(tree, treeIndex, classSlug, ranks, cells, asset);
  });
}

export function editorTrees({
  index,
  order,
  classSlug,
  characterLevel,
  readOnly,
  bandDiff,
  asset,
}: EditorTreesInput): TalentTreeView[] {
  const ranks = ranksByTalent([...order]);
  const available = characterLevel === null ? MAX_POINTS : pointsAvailableAt(characterLevel);
  const atCap = characterLevel !== null && order.length >= available;
  const planned = new Map<number, number>();
  order.forEach((id, i) => {
    if (characterLevel !== null && i >= available) planned.set(id, (planned.get(id) ?? 0) + 1);
  });
  return index.trees.map((tree, treeIndex) => {
    const cells: TalentCellView[] = gridCells(tree).map(({ talent, tier, column }) => {
      const rank = ranks.get(talent.id) ?? 0;
      const decision = readOnly
        ? { ok: false as const, reason: '' }
        : canAddPoint(index, [...order], talent.id);
      const capped = decision.ok && atCap;
      // The dot marks what a load would change; an empty build has nothing to change yet.
      const differs = order.length > 0 && (bandDiff?.diffTalentIds.has(talent.id) ?? false);
      const cell: TalentCellView = {
        ...baseCell(talent, tier, column, rank, asset, true),
        state: cellState(rank, talent.max_rank, decision.ok && !capped),
        differs,
        // A maxed talent says nothing more, as the client's tooltip: its rank line is the answer.
        req:
          rank >= talent.max_rank
            ? ''
            : capped && characterLevel !== null
              ? plannerCopy.noPointsLeft(characterLevel)
              : decision.ok
                ? ''
                : decision.reason,
        band:
          differs && bandDiff !== null
            ? bandCompareCopy.bandRankTooltipLine(
                bandDiff.bandLabel,
                bandDiff.bandRankById.get(talent.id) ?? 0,
                talent.max_rank,
              )
            : '',
      };
      if (capped) cell.plannable = true;
      if ((planned.get(talent.id) ?? 0) > 0 && characterLevel !== null) {
        cell.planned = true;
        cell.plannedLabel = plannerCopy.plannedCell(characterLevel);
      }
      return cell;
    });
    return treeFrame(tree, treeIndex, classSlug, ranks, cells, asset);
  });
}
