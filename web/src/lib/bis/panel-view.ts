// web/src/lib/bis/panel-view.ts
// The character-panel redesign's own view-model layer (lane bis-character-panel,
// 2026-09-29): every value `[spec].astro`'s paperdoll template reads, built once per band
// per faction from the raw `BisFile` the same way the list-first page's own `bandInfosFor`
// used to, but as a standalone module rather than page frontmatter -- `GearRow.astro` and
// the page both need the exact same `RowView`/`AlternativeView` shapes, and a page's own
// frontmatter cannot be `import`ed by a component the way a plain module can. Splitting
// this out also keeps the two files under the repo's own file-size guidance (many small,
// cohesive modules over one page that both computes and renders 17 slots x N bands).
import type { LootFile } from '../sim/loot';
import {
  buildWeightDisplayRows,
  refAbbrev,
  statLabelForSpec,
  type WeightDisplayRow,
} from '../sim/weights-display';
import { bisCopy } from './copy';
import {
  bandEntry,
  changedSinceBand,
  filledSlots,
  isMissingSlot,
  previousBandLevel,
  sourceBadgeLabel,
  type SlotRow,
} from './load';
import { describeSourceCell, hasKnownSource, resolveSourceCell, type SourceCell } from './source-cell';
import type { BisAlternative, BisFile, ChangedSlot, Faction, ItemDetail, LootQuestsFile } from './types';
import type { ItemTooltipModel } from '../items/tooltip';

/** One `BisAlternative` resolved to what the row shows: the same icon/name/source
 *  resolution a main pick gets, plus the gap-from-the-pick line (`bisCopy.
 *  alternativeGapLabel`, computed by the template from `dpsDelta` -- kept as the raw
 *  number here so a caller can format it however the surface needs). */
export interface AlternativeView {
  itemId: number;
  itemName: string;
  model?: ItemTooltipModel;
  sourceKind: SourceCell['kind'];
  sourceDetail: string;
  dpsDelta: number;
}

/** One slot's row: the source cell, the item itself, and (ruling 2) its up-to-three
 *  alternatives -- everything a `GearRow.astro` needs, resolved once here rather than in
 *  the template. */
export interface RowView {
  slot: string;
  empty: boolean;
  /** Set only when `empty` -- the exact copy the row shows (the ordinary "no known source"
   *  line, or the two-hander-equipped line for an off-hand slot ruling 4 carves out). */
  emptyCopy?: string;
  itemId?: number;
  itemName?: string;
  model?: ItemTooltipModel;
  itemLevel?: number;
  sourceKind?: SourceCell['kind'];
  sourceDetail?: string;
  keyStatsLine?: string;
  verified?: boolean;
  /** Always an array on a filled row (never undefined) -- empty when the pick has no
   *  alternatives, so the template's own "alternatives row absent: render nothing" rule
   *  (spec §4) is just `row.alternatives.length > 0`, no extra undefined check. */
  alternatives?: AlternativeView[];
  /** The item this slot's NEW pill replaced -- undefined at the file's first band (nothing
   *  was equipped before) or when the slot did not change this band. */
  replacedName?: string;
}

/** Resolves one `BisAlternative` the same way a main pick's source cell resolves --
 *  `resolveSourceCell`/`sourceBadgeLabel` only need `item_id`/`source_kind`, which an
 *  alternative carries just as a `BisSlot` does, so this is the identical resolution path,
 *  not a second one. */
function buildAlternativeView(
  alt: BisAlternative,
  faction: Faction,
  lootFile: LootFile & Partial<LootQuestsFile>,
  tooltipFor: (itemId: number) => ItemTooltipModel | undefined,
): AlternativeView {
  const badgeLabel = sourceBadgeLabel({ source_kind: alt.source_kind }, faction);
  const cell = resolveSourceCell(
    { item_id: alt.item_id, source_kind: alt.source_kind },
    faction,
    lootFile,
    badgeLabel,
  );
  return {
    itemId: alt.item_id,
    itemName: alt.item_name,
    model: tooltipFor(alt.item_id),
    sourceKind: cell.kind,
    sourceDetail: describeSourceCell(cell),
    dpsDelta: alt.dps_delta,
  };
}

function buildRowView(
  row: SlotRow,
  faction: Faction,
  itemDetailsMap: ReadonlyMap<number, ItemDetail>,
  lootFile: LootFile & Partial<LootQuestsFile>,
  replacedBySlot: ReadonlyMap<string, string | undefined>,
  tooltipFor: (itemId: number) => ItemTooltipModel | undefined,
  mainHandTwoHanded: boolean,
): RowView {
  if (isMissingSlot(row) || !hasKnownSource(row)) {
    const emptyCopy =
      row.slot === 'off_hand' && mainHandTwoHanded
        ? bisCopy.twoHanderEquippedLabel
        : bisCopy.noKnownSourceForSlot;
    return { slot: row.slot, empty: true, emptyCopy };
  }
  const badgeLabel = sourceBadgeLabel(row, faction);
  const cell = resolveSourceCell(row, faction, lootFile, badgeLabel);
  const model = tooltipFor(row.item_id);
  return {
    slot: row.slot,
    empty: false,
    itemId: row.item_id,
    itemName: row.item_name,
    model,
    itemLevel: itemDetailsMap.get(row.item_id)?.item_level,
    sourceKind: cell.kind,
    sourceDetail: describeSourceCell(cell),
    keyStatsLine: model && model.stats.length > 0 ? model.stats.slice(0, 4).join(', ') : undefined,
    verified: row.verified,
    alternatives: (row.alternatives ?? []).map((alt) =>
      buildAlternativeView(alt, faction, lootFile, tooltipFor),
    ),
    replacedName: replacedBySlot.get(row.slot),
  };
}

/** One weight-rail row, ready to render. */
export interface WeightBarRow {
  row: WeightDisplayRow;
  /** 4-100: `%` width for the row's horizontal bar (never 0 -- tenet 4's "nothing clipped"
   *  applied to a bar chart), scaled against the loudest non-reference weight in the band.
   *  Undefined for an insignificant row: its bar is omitted entirely (spec §3.3), not drawn
   *  faint, so a faded sliver never implies a real measurement. */
  barPercent: number;
  /** 1 point of this stat in raw DPS (`weight * band.reference_dps_per_point`) -- undefined
   *  when the band carries no `reference_dps_per_point` (never fabricated), for the
   *  reference row itself (always "= 1", never a DPS clause), or for an insignificant row
   *  (its value slot is `weightsNoEffect` instead). */
  dpsPerPoint?: number;
}

function weightBarsFor(
  weights: BisFile['bands'][number]['weights'],
  referenceStat: string,
  spec: string,
  referenceDpsPerPoint: number | null,
): WeightBarRow[] {
  const rows = buildWeightDisplayRows(
    weights.map((w) => ({
      stat: w.stat,
      weight: w.weight,
      error: w.error ?? 0,
      insignificant: w.insignificant ?? false,
    })),
    referenceStat,
    spec,
  );
  const maxWeight = Math.max(...rows.filter((r) => !r.isReference).map((r) => Math.abs(r.weight)), 0.0001);
  return rows.map((row) => ({
    row,
    barPercent: row.isReference ? 100 : Math.min(100, Math.max(4, (Math.abs(row.weight) / maxWeight) * 100)),
    dpsPerPoint:
      referenceDpsPerPoint === null || row.isReference || !row.significant
        ? undefined
        : row.weight * referenceDpsPerPoint,
  }));
}

/** One band's worth of `.paperdoll` data: every row, the centre column's numbers, and the
 *  weight rail -- the character panel's single source of truth, read straight by the
 *  template with no further resolution. */
export interface BandInfo {
  band: number;
  bandIndex: number;
  rows: RowView[];
  newSlots: ReadonlySet<string>;
  changed: ChangedSlot[] | undefined;
  previousBand: number | undefined;
  upgradesCount: number;
  setDps: number;
  dpsDelta: number | undefined;
  race: string;
  talentPoints: number;
  weightBars: WeightBarRow[];
  /** The reference stat's own short form (`refAbbrev`) for a significant row's "1 Agility =
   *  2.05 RAP" sentence -- computed once per band rather than once per row, since it never
   *  varies within a band. */
  refAbbrevText: string;
  /** The centre column's first line, above the weight list (spec §3.3): "1 <reference> =
   *  <n> DPS" when the band carries `reference_dps_per_point`, else the reference row's own
   *  `sentence` (today's `referenceSentence`, unchanged) -- never a fabricated DPS number. */
  referenceSentenceLine: string;
}

export interface PanelViewDeps {
  itemDetails: ReadonlyMap<number, ItemDetail>;
  loot: LootFile & Partial<LootQuestsFile>;
  tooltipFor: (itemId: number) => ItemTooltipModel | undefined;
  referenceStat: string;
  spec: string;
}

/** Every band this file covers for `faction`, in `bands`' own order -- the page's one call
 *  into this module, replacing the list-first page's own (now removed) `bandInfosFor`. */
export function bandInfosFor(
  file: BisFile,
  bands: readonly number[],
  faction: Faction,
  deps: PanelViewDeps,
): BandInfo[] {
  return bands.flatMap((band, bandIndex): BandInfo[] => {
    const bandData = bandEntry(file, band, faction);
    if (bandData === undefined) return [];
    const changed = changedSinceBand(file, band, faction);
    const replacedBySlot = new Map<string, string | undefined>(
      (changed ?? [])
        .filter((entry) => entry.after !== undefined)
        .map((entry) => [entry.slot, entry.before?.item_name]),
    );
    const slotRows = filledSlots(bandData);
    const mainHandRow = slotRows.find((r) => r.slot === 'main_hand');
    const mainHandModel =
      mainHandRow !== undefined && !isMissingSlot(mainHandRow) && hasKnownSource(mainHandRow)
        ? deps.tooltipFor(mainHandRow.item_id)
        : undefined;
    const mainHandTwoHanded = mainHandModel?.typeLabel === 'Two-Handed Weapon';
    const rows = slotRows.map((row) =>
      buildRowView(
        row,
        faction,
        deps.itemDetails,
        deps.loot,
        replacedBySlot,
        deps.tooltipFor,
        mainHandTwoHanded,
      ),
    );
    const newSlots = new Set(
      bandIndex === 0
        ? rows.filter((row) => !row.empty).map((row) => row.slot)
        : (changed ?? []).filter((entry) => entry.after !== undefined).map((entry) => entry.slot),
    );
    const previousBand = previousBandLevel(file, band);
    const previousSetDps =
      previousBand === undefined ? undefined : bandEntry(file, previousBand, faction)?.set_dps;
    const referenceDpsPerPoint = bandData.reference_dps_per_point ?? null;
    const weightBars = weightBarsFor(bandData.weights, deps.referenceStat, deps.spec, referenceDpsPerPoint);
    const referenceLabel = weightBars[0]?.row.label ?? statLabelForSpec(deps.referenceStat, deps.spec);
    const referenceSentenceLine =
      referenceDpsPerPoint === null
        ? (weightBars[0]?.row.sentence ?? referenceLabel)
        : bisCopy.weightsReferenceDpsLine(referenceLabel, referenceDpsPerPoint);
    return [
      {
        band,
        bandIndex,
        rows,
        newSlots,
        changed,
        previousBand,
        upgradesCount: changed?.length ?? 0,
        setDps: bandData.set_dps,
        dpsDelta: previousSetDps === undefined ? undefined : bandData.set_dps - previousSetDps,
        race: bandData.race,
        talentPoints: bandData.talent_points,
        weightBars,
        refAbbrevText: refAbbrev(referenceLabel),
        referenceSentenceLine,
      },
    ];
  });
}

/** Every model this band's rows and their alternatives, plus the "what changed" diff, name
 *  -- merged into `into` keyed by item id, the single deduplicated map `BisTooltipHost`
 *  needs. A `Map` write is idempotent, so an id named twice (an alternative that is also a
 *  later band's own pick) costs one entry, not two. */
export function collectModelsInto(
  into: Map<number, ItemTooltipModel>,
  bandInfos: readonly BandInfo[],
  tooltipFor: (itemId: number) => ItemTooltipModel | undefined,
): void {
  for (const info of bandInfos) {
    for (const row of info.rows) {
      if (row.itemId !== undefined && row.model !== undefined) into.set(row.itemId, row.model);
      for (const alt of row.alternatives ?? []) {
        if (alt.model !== undefined) into.set(alt.itemId, alt.model);
      }
    }
    for (const entry of info.changed ?? []) {
      if (entry.before !== undefined) {
        const model = tooltipFor(entry.before.item_id);
        if (model !== undefined) into.set(entry.before.item_id, model);
      }
      if (entry.after !== undefined) {
        const model = tooltipFor(entry.after.item_id);
        if (model !== undefined) into.set(entry.after.item_id, model);
      }
    }
  }
}
