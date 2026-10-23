// web/src/lib/guides/guide-bis.ts
// The guide's "Best in slot" table: the same per-slot picks the /bis page lists, one block
// per level range, read from the spec's ranked BiS file through the same view builder /bis
// uses (`bandInfosFor`) so a guide and its BiS page can never name different items. The
// table is server-rendered from this; nothing here is typed per guide.
import { bisCopy } from '../bis/copy';
import { SLOT_DISPLAY_LABELS } from '../bis/slot-display-labels';
import { bandInfosFor, type BandInfo, type RowView } from '../bis/panel-view';
import { bandLevels, itemDetails, loadLootFile } from '../bis/load';
import { raceNameFor } from './race-table';
import { factionForFirstRace, rankedBisFile } from './class-dps';
import { readItemTooltipModel } from '../items/lookup-node';
import type { ItemTooltipModel } from '../items/tooltip';
import type { Slot } from '../planner/types';
import type { BisFile, Faction } from '../bis/types';

export interface GuideBisRow {
  slot: string;
  slotLabel: string;
  empty: boolean;
  /** Why an empty slot is empty, in the /bis page's own words. */
  emptyCopy?: string;
  itemId?: number;
  itemName?: string;
  /** Client quality 0..5, for the name colour and the icon border. */
  quality?: number;
  /** The icon file stem, absent when the build has none for the item. */
  icon?: string;
  /** `source-cell.ts`'s kind, for the small glyph. */
  sourceKind?: string;
  /** "Dungeon · Boss · 12%", "Crafted: Tailoring", "Quest: ..." as /bis words it. */
  sourceDetail?: string;
  /** The ranker did not run the pick through a full sim, only scored its stats. */
  notSimChecked: boolean;
  /** The client's item type line ("Two-Handed Weapon", "Shield"), for the enchant picks. */
  typeLabel?: string;
}

export interface GuideBisBand {
  band: number;
  /** "20 to 29", "60". */
  label: string;
  /** Anchor on the /bis page for this band, e.g. `band-alliance-60`. */
  bisAnchor: string;
  rows: GuideBisRow[];
  /** True when the level 60 picks carry no raid-tier item, so the table is the pre-raid list. */
  preRaid: boolean;
}

export interface GuideBis {
  faction: Faction;
  /** The race the band was ranked for, as a display name. */
  raceName: string;
  /** The band the table opens on. */
  defaultBand: number;
  bands: GuideBisBand[];
}

const PLACEHOLDER_ICON = 'inv_misc_questionmark';

function rowFor(row: RowView): GuideBisRow {
  const slotLabel = SLOT_DISPLAY_LABELS[row.slot as Slot] ?? row.slot;
  if (row.empty || row.itemId === undefined) {
    return { slot: row.slot, slotLabel, empty: true, emptyCopy: row.emptyCopy, notSimChecked: false };
  }
  const icon = row.model?.icon;
  return {
    slot: row.slot,
    slotLabel,
    empty: false,
    itemId: row.itemId,
    itemName: row.itemName,
    quality: row.model?.quality ?? 1,
    icon: icon === undefined || icon === PLACEHOLDER_ICON ? undefined : icon,
    sourceKind: row.sourceKind,
    sourceDetail: row.sourceDetail,
    notSimChecked: row.notSimChecked ?? false,
    typeLabel: row.model?.typeLabel,
  };
}

/** Whether a level 60 pick comes from a raid: when none does, the list is the pre-raid one. */
function isPreRaid(band: number, rows: readonly RowView[]): boolean {
  return band === 60 && !rows.some((row) => row.sourceKind === 'raid');
}

export function guideBisBandsFrom(infos: readonly BandInfo[], faction: Faction): GuideBisBand[] {
  return infos.map((info) => ({
    band: info.band,
    label: bisCopy.bandRangeLabel(info.band),
    bisAnchor: `band-${faction}-${info.band}`,
    rows: info.rows.map(rowFor),
    preRaid: isPreRaid(info.band, info.rows),
  }));
}

export function guideBisFrom(
  file: BisFile,
  faction: Faction,
  deps: Parameters<typeof bandInfosFor>[3],
): GuideBis | undefined {
  const infos = bandInfosFor(file, bandLevels(file), faction, deps);
  if (infos.length === 0) return undefined;
  const top = infos[infos.length - 1]!;
  return {
    faction,
    raceName: raceNameFor(top.race),
    defaultBand: top.band,
    bands: guideBisBandsFrom(infos, faction),
  };
}

/** The spec's picks for the faction its first recommended race plays, from the real
 *  published file only (never the /bis dev fixture); `undefined` when the spec has no ranked
 *  list yet, and the guide then shows no table at all rather than an empty one. */
export function loadGuideBis(
  build: string,
  classSlug: string,
  spec: string,
  recommendedRaces: readonly string[],
): GuideBis | undefined {
  const file = rankedBisFile(build, spec);
  const faction = factionForFirstRace(classSlug, recommendedRaces);
  if (file === undefined || faction === undefined) return undefined;
  const tooltipFor = (itemId: number): ItemTooltipModel | undefined =>
    readItemTooltipModel(build, classSlug, itemId) ?? undefined;
  return guideBisFrom(file, faction, {
    itemDetails: itemDetails(build, classSlug),
    loot: loadLootFile(build),
    tooltipFor,
    spec,
  });
}
