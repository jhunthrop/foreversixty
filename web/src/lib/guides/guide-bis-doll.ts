// web/src/lib/guides/guide-bis-doll.ts
// The guide's Best in slot doll: the level 60 picks of each faction on the shared paperdoll
// (ui/PaperdollFrame, the frame /bis draws), beside the lane's slot list. The cells come from the
// lane's own rows (lib/guides/guide-bis.ts), so the doll and the list can never name different
// items; each pick carries its client tooltip for the frame's centre.
import { TWO_HAND_LABEL, type ItemTooltipModel } from '../items/tooltip';
import type { DollCell } from '../ui/paperdoll';
import type { ItemRef } from '../ui/primitives';
import type { Faction } from '../bis/types';
import type { GuideBis, GuideBisPanel, GuideBisRow } from './guide-bis';

/** One faction's level 60 doll. */
export interface GuideDollView {
  faction: Faction;
  /** The race the picks were ranked for, as a display name. */
  raceName: string;
  preRaid: boolean;
  cells: DollCell[];
}

export type TooltipLookup = (itemId: number) => ItemTooltipModel | undefined;

const DOLL_BAND = 60;

function refOf(row: GuideBisRow): ItemRef | null {
  if (row.empty || row.itemId === undefined) return null;
  return { id: row.itemId, name: row.itemName ?? '', quality: row.quality ?? 1, icon: row.icon };
}

/** The panel's rows as doll cells; an empty off hand under a two-hander shows it greyed. */
export function guideDollCells(panel: GuideBisPanel, tooltipFor: TooltipLookup): DollCell[] {
  const mainHand = panel.rows.find((row) => row.slot === 'main_hand');
  const mainRef = mainHand === undefined ? null : refOf(mainHand);
  const twoHander =
    mainRef !== null && tooltipFor(mainRef.id)?.slotLabel === TWO_HAND_LABEL ? mainRef : undefined;
  return panel.rows.map((row) => ({
    slot: row.slot,
    label: row.slotLabel,
    item: refOf(row),
    source: row.sourceDetail,
    sourceKind: row.sourceKind,
    empty: row.empty ? row.emptyCopy : undefined,
    coveredBy: row.slot === 'off_hand' && row.empty ? twoHander : undefined,
  }));
}

/** Each faction's level 60 doll, the guide's default faction first. */
export function guideDollViews(bis: GuideBis | undefined, tooltipFor: TooltipLookup): GuideDollView[] {
  if (bis === undefined) return [];
  return bis.panels
    .filter((panel) => panel.band === DOLL_BAND)
    .toSorted((a, b) => Number(b.faction === bis.defaultFaction) - Number(a.faction === bis.defaultFaction))
    .map((panel) => ({
      faction: panel.faction,
      raceName: panel.raceName,
      preRaid: panel.preRaid,
      cells: guideDollCells(panel, tooltipFor),
    }));
}

/** Every pick's tooltip, by item id, for the frame's centre (the two-hander's included). */
export function guideDollTooltips(
  views: readonly GuideDollView[],
  tooltipFor: TooltipLookup,
): Record<number, ItemTooltipModel> {
  const ids = new Set(
    views.flatMap((view) => view.cells.flatMap((cell) => (cell.item ? [cell.item.id] : []))),
  );
  return Object.fromEntries(
    [...ids].flatMap((id) => {
      const model = tooltipFor(id);
      return model === undefined ? [] : [[id, model] as const];
    }),
  );
}

/** The doll a reader sees: their character's faction when the guide has its picks, else the
 *  guide's default (the first view). */
export function dollForFaction(
  views: readonly GuideDollView[],
  faction: string | undefined,
): GuideDollView | undefined {
  return views.find((view) => view.faction === faction) ?? views[0];
}
