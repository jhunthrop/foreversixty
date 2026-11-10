// web/src/lib/planner/editor-trees.ts
// The planner's trees as ui/TalentEditor draws them (spec 2026-11-08 sections 4.4 and 5.4,
// decisions P2 to P5, ruling P4): every cell's state in the game's colours, with a character's
// level cap greying the unspent cells (P3) and its tooltip saying why; a locked cell's tooltip
// naming what it needs (P5); the gold dot and band line on a cell whose rank differs from a
// compared band; and the planned-ahead mark on a cell holding points past the character's level.
import type { EditorCell, EditorTree, TalentCellState } from '../ui/talent-editor';
import { SPECS } from '../sim/specs';
import type { BandDiffView } from './band-compare';
import { connectorsFor, gridViewBox } from './connectors';
import { bandCompareCopy, plannerCopy } from './copy';
import { pointsAvailableAt, ranksByTalent } from './derive';
import { gridCells, gridSize } from './grid';
import { canAddPoint, type TalentIndex } from './rules';
import { cellState } from './styles';
import { MAX_POINTS } from './types';

/** The talent frame's tiers: every class's trees fill seven. */
export const GAME_TIERS = 7;

const STATE: Record<ReturnType<typeof cellState>, TalentCellState> = {
  available: 'avail',
  filled: 'part',
  maxed: 'max',
  locked: 'lock',
};

export interface EditorTreesInput {
  index: TalentIndex;
  order: readonly number[];
  classSlug: string;
  /** The character's level, or null: its points stop there (P3) and the rest are planned (P4). */
  characterLevel: number | null;
  readOnly: boolean;
  bandDiff: BandDiffView | null;
  /** A data file's URL (`dataUrl(build, path)`). */
  asset: (path: string) => string;
}

/** A tree's head icon: its spec's icon from the generated spec list, '' when it has none. */
export function treeIconSrc(classSlug: string, treeIndex: number, asset: (path: string) => string): string {
  const icon = SPECS.find((row) => row.class_slug === classSlug && row.tree_index === treeIndex)?.icon;
  return icon === undefined ? '' : asset(`icons/${icon}.webp`);
}

export function editorTrees({
  index,
  order,
  classSlug,
  characterLevel,
  readOnly,
  bandDiff,
  asset,
}: EditorTreesInput): EditorTree[] {
  const ranks = ranksByTalent([...order]);
  const available = characterLevel === null ? MAX_POINTS : pointsAvailableAt(characterLevel);
  const atCap = characterLevel !== null && order.length >= available;
  const planned = new Map<number, number>();
  order.forEach((id, i) => {
    if (characterLevel !== null && i >= available) planned.set(id, (planned.get(id) ?? 0) + 1);
  });
  return index.trees.map((tree, treeIndex) => {
    // Every tree is drawn at the game frame's seven tiers (a trimmed tree too), so the Plate's
    // ready height is the one its loading reserve holds (spec 5.2).
    const size = { ...gridSize(tree), tiers: Math.max(GAME_TIERS, gridSize(tree).tiers) };
    const cells: EditorCell[] = gridCells(tree).map(({ talent, tier, column }) => {
      const rank = ranks.get(talent.id) ?? 0;
      const decision = readOnly
        ? { ok: false as const, reason: '' }
        : canAddPoint(index, [...order], talent.id);
      const capped = decision.ok && atCap;
      // The dot marks what a load would change; an empty build has nothing to change yet.
      const differs = order.length > 0 && (bandDiff?.diffTalentIds.has(talent.id) ?? false);
      const cell: EditorCell = {
        id: talent.id,
        tier,
        column,
        name: talent.name,
        iconSrc: asset(`icons/${talent.icon}.webp`),
        rank,
        max: talent.max_rank,
        state: STATE[cellState(rank, talent.max_rank, decision.ok && !capped)],
        differs,
        now: rank > 0 ? (talent.ranks[rank - 1]?.description ?? '') : '',
        next: rank < talent.max_rank ? (talent.ranks[rank]?.description ?? '') : '',
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
    const points = cells.reduce((sum, cell) => sum + cell.rank, 0);
    return {
      id: tree.id,
      name: tree.name,
      iconSrc: treeIconSrc(classSlug, treeIndex, asset),
      points,
      ...(tree.background ? { artSrc: asset(`trees/${tree.background}.webp`) } : {}),
      columns: size.columns,
      tiers: size.tiers,
      viewBox: gridViewBox(size.tiers, size.columns),
      connectors: connectorsFor(tree, ranks).map((c) => ({ d: c.d, met: c.met })),
      cells,
    };
  });
}
