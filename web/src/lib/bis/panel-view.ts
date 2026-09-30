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
import { statLabelForSpec } from '../sim/weights-display';
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
import type {
  BisAlternative,
  BisFile,
  BisSlot,
  BisStatWeight,
  ChangedSlot,
  Faction,
  ItemDetail,
  LootQuestsFile,
} from './types';
import type { ItemTooltipModel } from '../items/tooltip';

/** The two ranker-own `swap_note` templates (`sim/cmd/leveling-bis/report.go`'s own
 *  `fmt.Sprintf` calls): "confirmed by the sim against X (id N): kept the pick, A vs B set
 *  DPS" and "beat the scored pick X (id N) in the sim: A vs B set DPS" -- in both, `A` is
 *  this row's own pick's measured DPS and `B` the named alternative's (verified directly
 *  against report.go's own argument order: `sw.BaselineDPS, sw.SwapDPS` when the pick was
 *  kept, `sw.SwapDPS, sw.BaselineDPS` when the pick IS the promoted swap winner -- either
 *  way the row's own current pick comes first). Checked against all 580 `swap_note` strings
 *  in the 2026-09-29 build (`data/builds/1.60.1.70009/bis/*.json`): every one matches. */
const SWAP_NOTE_PATTERN =
  /^(?:confirmed by the sim against|beat the scored pick) (.+?) \(id \d+\)(?:: kept the pick, |\s+in the sim: )(\d+(?:\.\d+)?) vs (\d+(?:\.\d+)?) set DPS$/;

export interface ParsedSwapNote {
  /** The named alternative's item name (never this row's own pick -- see the pattern's own
   *  doc for which side of "A vs B" each belongs to). */
  itemName: string;
  /** This row's own pick's measured DPS. */
  pickDps: number;
  /** The named alternative's measured DPS. */
  altDps: number;
}

/** Parses a `BisSlot.swap_note` sentence into the two numbers a player can read --
 *  undefined for any text that does not match `SWAP_NOTE_PATTERN` (a `swap_note` format
 *  this function does not know, e.g. the verify-error fallback string), so a caller falls
 *  back to the raw note rather than fabricate a reading of it (tenet 8). */
export function parseSwapNote(note: string): ParsedSwapNote | undefined {
  const match = SWAP_NOTE_PATTERN.exec(note);
  if (match === null) return undefined;
  const [, itemName, pickDps, altDps] = match;
  return { itemName: itemName!, pickDps: Number(pickDps), altDps: Number(altDps) };
}

/** The same two `swap_note` templates' common prefix, WITHOUT the trailing DPS clause --
 *  bis-ranker-integrity-5's own `dpsComparisonPhrase` (`report.go`) now prints "A vs B set
 *  DPS" only when A (the primary measurement) matches the band's own finished `set_dps`,
 *  else "+N.N DPS over it", so `SWAP_NOTE_PATTERN` above (which requires the "vs...set DPS"
 *  suffix) no longer matches every real `swap_note`. This pattern only needs the runner-up's
 *  own name -- the number now comes from the row's own `dps_delta` instead (see
 *  `evidenceLineFor`) -- so it stops right after "kept the pick," / "in the sim:" and never
 *  cares which suffix format follows. */
const RUNNER_UP_NAME_PATTERN =
  /^(?:confirmed by the sim against|beat the scored pick) (.+?) \(id \d+\)(?:: kept the pick,|\s+in the sim:)/;

/** The named runner-up's item name off a `swap_note` sentence, regardless of which DPS
 *  suffix `dpsComparisonPhrase` chose -- `undefined` for a `swap_note` this page does not
 *  recognise at all (e.g. the verify-error fallback string), same discipline as
 *  `parseSwapNote`. */
function runnerUpNameFromSwapNote(note: string): string | undefined {
  return RUNNER_UP_NAME_PATTERN.exec(note)?.[1];
}

/** The row's own evidence line (spec: fourth wow-player sweep, day 3, item 1; delta
 *  preference added bis-ranker-integrity-5) -- in player words when `swap_note` matches a
 *  known ranker template, the raw note otherwise (never nothing: a `swap_note` the page
 *  cannot parse is still real evidence, just unparsed).
 *
 *  Prefers the row's own `dps_delta` when present: that number is always real and always
 *  consistent with itself (`BisSlot.dps_delta`'s own doc), unlike `swap_note`'s two absolute
 *  numbers, which are each a full-set snapshot at the moment THIS slot was decided and can
 *  silently disagree with the band's own header `set_dps` -- never publish "X vs Y DPS" from
 *  the old parse once a trustworthy single number exists. Falls back to the full two-number
 *  parse (`parseSwapNote`) only when `dps_delta` is absent (a file published before this
 *  field existed). */
function evidenceLineFor(
  swapNote: string | undefined,
  dpsDelta: number | null | undefined,
): string | undefined {
  if (swapNote === undefined) return undefined;
  if (dpsDelta !== undefined && dpsDelta !== null) {
    const runnerUpName = runnerUpNameFromSwapNote(swapNote);
    if (runnerUpName !== undefined) return bisCopy.evidenceLineDelta(runnerUpName, dpsDelta);
  }
  const parsed = parseSwapNote(swapNote);
  return parsed === undefined
    ? swapNote
    : bisCopy.evidenceLine(parsed.itemName, parsed.pickDps, parsed.altDps);
}

/** The main pick's own verified-glyph title -- the default "confirmed by a Top Gear pass"
 *  copy unless the pick carries `sim_dps` and no `swap_note` (a sim-decided row with no
 *  swap narrative to tell -- a trinket/proc/weapon-pair tournament winner), which names its
 *  own real number instead (spec item 1's second rule). */
function verifiedGlyphTitleFor(row: BisSlot): string | undefined {
  return row.swap_note === undefined && row.sim_dps !== undefined
    ? bisCopy.simDpsVerifiedTitle(row.sim_dps)
    : undefined;
}

/** `empty_reason` -> the honest, player-worded line for why this slot has no pick (spec
 *  item 3) -- `no_sourced_item` and any value this page does not recognise both fall back
 *  to today's plain `noKnownSourceForSlot` line, never a fabricated reason. Never called for
 *  the off-hand-under-a-two-hander case, which `buildRowView` special-cases first with its
 *  own copy regardless of what `empty_reason` says (that rule predates this field and stays
 *  the more specific, more player-legible one of the two). */
function emptyReasonLabel(reason: string | undefined): string {
  switch (reason) {
    case 'no_dps_value':
      return bisCopy.emptyReasonNoDpsValue;
    case 'effect_not_modelled':
      return bisCopy.emptyReasonEffectNotModelled;
    default:
      return bisCopy.noKnownSourceForSlot;
  }
}

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
  /** `BisAlternative.verified` -- true for the one alternative (per slot, at most) the
   *  ranker's own verify pass actually simmed against the pick, whose `dpsDelta` is real
   *  sim output rather than a score estimate (wow-player fix round 1: label it with the
   *  same check glyph a verified main pick gets, never a second unlabelled number). */
  verified?: boolean;
  /** "ilvl 24 · needs 21" (band above the alternative's own required level) or "ilvl 24"
   *  (band at or above it) -- undefined only when the alternative's id has no tooltip
   *  model at all (wow-player fix round 1: read straight off the model already resolved
   *  below, no second item lookup). */
  metaLabel?: string;
  /** `BisAlternative.effect_unmodelled` -- see `RowView.effectUnmodelled`'s own doc; an
   *  alternative can carry the identical flag for the identical reason. */
  effectUnmodelled?: boolean;
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
  /** The row's own evidence line, in player words -- see `evidenceLineFor`'s own doc.
   *  Undefined when the pick carries no `swap_note` at all (most rows: nothing to show). */
  evidenceLine?: string;
  /** Override for the main pick's verified-glyph `title` -- see `verifiedGlyphTitleFor`'s
   *  own doc. Undefined keeps the glyph's own default title. */
  verifiedGlyphTitle?: string;
  /** `BisSlot.effect_unmodelled` -- true when this pick's proc/use effect is not
   *  simulated, so its DPS number is stats-only (fourth wow-player sweep, day 3: "Serenity
   *  Field" beating a real combat trinket by less than its own blind spot). */
  effectUnmodelled?: boolean;
  /** `BisSlot.sim_status === 'not_in_sim'` -- true when this build's sim database does not
   *  carry the pick at all, so it stays ranked by stat weights alone with no full sim run
   *  behind it (spec addendum 3, §D). Main-pick-only: `AlternativeView` has no counterpart
   *  -- `BisAlternative` never carries `sim_status` (`report.go`'s `alternativeRow`). */
  notSimChecked?: boolean;
  /** `BisSlot.low_value` -- true for a weapon row where no sourced candidate scored above
   *  zero and the ranker published the best-by-item-level fallback instead of an empty
   *  slot. */
  lowValue?: boolean;
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
  band: number,
): AlternativeView {
  const badgeLabel = sourceBadgeLabel({ source_kind: alt.source_kind }, faction);
  const cell = resolveSourceCell(
    { item_id: alt.item_id, source_kind: alt.source_kind },
    faction,
    lootFile,
    badgeLabel,
  );
  const model = tooltipFor(alt.item_id);
  return {
    itemId: alt.item_id,
    itemName: alt.item_name,
    model,
    sourceKind: cell.kind,
    sourceDetail: describeSourceCell(cell),
    dpsDelta: alt.dps_delta,
    verified: alt.verified,
    metaLabel:
      model === undefined
        ? undefined
        : bisCopy.alternativeMetaLabel(model.itemLevel, model.requiredLevel, band),
    effectUnmodelled: alt.effect_unmodelled,
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
  band: number,
): RowView {
  if (isMissingSlot(row) || !hasKnownSource(row)) {
    // A missing slot (`isMissingSlot`) never published an `empty_reason` at all -- the
    // pipeline's own `{ slot, missing: true }` shape carries nothing but the slot name --
    // so only a present-but-unsourced `BisSlot` has one to read.
    const emptyReason = isMissingSlot(row) ? undefined : row.empty_reason;
    const emptyCopy =
      row.slot === 'off_hand' && mainHandTwoHanded
        ? bisCopy.twoHanderEquippedLabel
        : emptyReasonLabel(emptyReason);
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
      buildAlternativeView(alt, faction, lootFile, tooltipFor, band),
    ),
    replacedName: replacedBySlot.get(row.slot),
    evidenceLine: evidenceLineFor(row.swap_note, row.dps_delta),
    verifiedGlyphTitle: verifiedGlyphTitleFor(row),
    effectUnmodelled: row.effect_unmodelled,
    notSimChecked: row.sim_status === 'not_in_sim',
    lowValue: row.low_value,
  };
}

/** `melee_haste`/`spell_haste` -- the only two haste ids a spec's own `weight_stats` ever
 *  carries (`sim/cmd/leveling-bis/report.go`'s own `isHasteStat`, mirrored here client-side
 *  for the same reason: vanilla haste has no rating conversion in this ruleset and is not
 *  comparable point-for-point against a primary/rating stat -- owner correction,
 *  2026-09-30). Never a table row on this rail; see `hasteCaptionFor` below. */
function isHasteStat(stat: string): boolean {
  return stat === 'melee_haste' || stat === 'spell_haste';
}

interface ScaleFactors {
  scaleFactor: number;
  dpsPerPoint?: number;
  scaleError: number;
}

/** `scale_factor`/`dps_per_point`/`scale_error` for every row in `weights`, one `ScaleFactors`
 *  per stat id, plus the stat id every `scaleFactor` was normalized against (`""` when none
 *  qualified). Reads the ranker's own published fields
 *  (`sim/cmd/leveling-bis/report.go`'s own `normalizeScaleFactors`) when present; otherwise
 *  computes the identical numbers here (this lane's brief: "the rail must fall back
 *  gracefully on a JSON that lacks the new fields"), mirroring that function's own rule
 *  exactly -- the significant, non-haste row with the largest `weight` is the anchor; a band
 *  with none (every row insignificant) publishes every `scaleFactor`/`scaleError` as 0 rather
 *  than divide by zero or invent an anchor from noise. */
function computeScaleFactors(
  weights: readonly BisStatWeight[],
  referenceDpsPerPoint: number | null,
  publishedAnchorStat: string | null,
): { byStat: ReadonlyMap<string, ScaleFactors>; anchorStat: string } {
  if (weights.some((w) => w.scale_factor !== undefined)) {
    const byStat = new Map<string, ScaleFactors>(
      weights.map((w) => [
        w.stat,
        { scaleFactor: w.scale_factor ?? 0, dpsPerPoint: w.dps_per_point, scaleError: w.scale_error ?? 0 },
      ]),
    );
    return { byStat, anchorStat: publishedAnchorStat ?? '' };
  }

  let anchorStat = '';
  let anchorWeight = 0;
  for (const w of weights) {
    if ((w.insignificant ?? false) || isHasteStat(w.stat)) continue;
    if (w.weight > anchorWeight) {
      anchorWeight = w.weight;
      anchorStat = w.stat;
    }
  }
  const byStat = new Map<string, ScaleFactors>(
    weights.map((w) => [
      w.stat,
      {
        scaleFactor: anchorStat === '' ? 0 : w.weight / anchorWeight,
        dpsPerPoint: referenceDpsPerPoint === null ? undefined : w.weight * referenceDpsPerPoint,
        scaleError: anchorStat === '' ? 0 : (w.error ?? 0) / anchorWeight,
      },
    ]),
  );
  return { byStat, anchorStat };
}

/** One scale-rail table row, ready to render -- never a haste stat (see `isHasteStat`). */
export interface ScaleRow {
  stat: string;
  /** `statLabelForSpec(stat, spec)`, plus a quiet " rating" suffix for a rating-family row
   *  (`unit === 'rating'`) -- "Crit" -> "Crit rating", so the row's own face reads as
   *  per-RATING-POINT even on a phone with no hover. */
  label: string;
  /** Per point of stat, normalized so the band's own top per-point stat reads exactly `1`
   *  (`ScaleFactors.scaleFactor`, published or computed -- see `computeScaleFactors`). */
  scaleFactor: number;
  /** Absolute DPS per point of this stat -- undefined when the band carries no
   *  `reference_dps_per_point` (never fabricated). */
  dpsPerPoint?: number;
  /** `error`, on the same normalized divisor as `scaleFactor`. */
  scaleError: number;
  /** `false` greys this row out on the page and shows `bisCopy.weightsNotSignificant`
   *  instead of its own scale factor -- the row itself is never dropped (tenet 4). */
  significant: boolean;
  /** 4-100: `%` width for the row's horizontal bar (never 0), scaled against the loudest
   *  scale factor in THIS table (haste is never a table row, so that loudest value is
   *  always the anchor stat's own `1`). */
  barPercent: number;
  /** The rating-family row's own hover override for the `<li>`'s native `title` attribute:
   *  "14 Crit rating = 1% Crit". Undefined for every non-rating row. */
  ratingFactorTitle?: string;
}

/** Every per-point row in `weights` (haste excluded -- `isHasteStat`), sorted strictly
 *  descending by `scaleFactor`, from the already-computed `byStat` (`computeScaleFactors`). */
function buildScaleRows(
  weights: readonly BisStatWeight[],
  spec: string,
  byStat: ReadonlyMap<string, ScaleFactors>,
): ScaleRow[] {
  const perPoint = weights.filter((w) => !isHasteStat(w.stat));
  const maxScale = Math.max(...perPoint.map((w) => byStat.get(w.stat)?.scaleFactor ?? 0), 0.0001);
  return perPoint
    .map((w): ScaleRow => {
      const factors = byStat.get(w.stat) ?? { scaleFactor: 0, scaleError: 0 };
      const label = statLabelForSpec(w.stat, spec);
      return {
        stat: w.stat,
        label: w.unit === 'rating' ? `${label} rating` : label,
        scaleFactor: factors.scaleFactor,
        dpsPerPoint: factors.dpsPerPoint,
        scaleError: factors.scaleError,
        significant: !(w.insignificant ?? false),
        barPercent: Math.min(100, Math.max(4, (factors.scaleFactor / maxScale) * 100)),
        ratingFactorTitle:
          w.rating_factor !== undefined ? bisCopy.weightsRatingFactorLine(label, w.rating_factor) : undefined,
      };
    })
    .sort((a, b) => b.scaleFactor - a.scaleFactor);
}

/** Haste's own one-line caption (owner correction, 2026-09-30, after player review: "haste
 *  is not a table row... It goes in the caption") -- `undefined` when this spec's own
 *  `weights` carries no haste stat at all (a caster spec with `spell_haste` still gets one;
 *  a spec with neither gets none), or when neither the published `band.haste_scale_factor`
 *  nor the client-side fallback (`byStat`) has a number to show (no trustworthy anchor this
 *  band -- see `computeScaleFactors`'s own doc). Prefers the published field so the page
 *  never has to recompute what the ranker already measured; falls back to the exact same
 *  divisor `buildScaleRows` used otherwise. */
function hasteCaptionFor(
  weights: readonly BisStatWeight[],
  publishedHasteScaleFactor: number | null,
  byStat: ReadonlyMap<string, ScaleFactors>,
): string | undefined {
  const hasteRow = weights.find((w) => isHasteStat(w.stat));
  if (hasteRow === undefined) return undefined;
  const scaleFactor = publishedHasteScaleFactor ?? byStat.get(hasteRow.stat)?.scaleFactor;
  if (scaleFactor === undefined) return undefined;
  return bisCopy.weightsHasteCaption(scaleFactor, hasteRow.insignificant ?? false);
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
  /** `BisBand.set_dps_partial` -- true when `setDps` excludes at least one pick's own
   *  contribution (a `notSimChecked` row this same band's `rows` also carries). */
  setDpsPartial: boolean;
  /** Count of this band's own `rows` carrying `notSimChecked: true` -- always computed
   *  from the SAME `rows` array the paperdoll renders, never a second backend counter, so
   *  `setDpsPartialNote`'s number always matches a row a player can find (spec addendum 3,
   *  §E). Meaningful only when `setDpsPartial` is true. */
  setDpsPartialCount: number;
  race: string;
  talentPoints: number;
  /** The scale-rail table, sorted descending by `scaleFactor`, haste never among them
   *  (`isHasteStat`, `buildScaleRows`). Empty when the band carries `weights_reason`. */
  scaleRows: ScaleRow[];
  /** The rail's own first line (`bisCopy.weightsScaleNote`/`weightsUnmeasuredLine`) --
   *  always a string, never fabricated: names the band's own top stat when one was found,
   *  states plainly when the sweep could not be trusted otherwise. */
  scaleNoteLine: string;
  /** Haste's own one-line caption (`hasteCaptionFor`) -- `undefined` when this spec carries
   *  no haste weight_stat, or the band's own weights could not be trusted at all. */
  hasteCaptionLine: string | undefined;
}

export interface PanelViewDeps {
  itemDetails: ReadonlyMap<number, ItemDetail>;
  loot: LootFile & Partial<LootQuestsFile>;
  tooltipFor: (itemId: number) => ItemTooltipModel | undefined;
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
        band,
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
    const weightsReason = bandData.weights_reason ?? null;
    const { byStat, anchorStat } =
      weightsReason !== null
        ? { byStat: new Map<string, ScaleFactors>(), anchorStat: '' }
        : computeScaleFactors(bandData.weights, referenceDpsPerPoint, bandData.scale_reference_stat ?? null);
    const scaleRows = weightsReason !== null ? [] : buildScaleRows(bandData.weights, deps.spec, byStat);
    const scaleNoteLine =
      weightsReason !== null
        ? bisCopy.weightsUnmeasuredLine
        : anchorStat === ''
          ? bisCopy.weightsUnmeasuredLine
          : bisCopy.weightsScaleNote(statLabelForSpec(anchorStat, deps.spec));
    const hasteCaptionLine =
      weightsReason !== null
        ? undefined
        : hasteCaptionFor(bandData.weights, bandData.haste_scale_factor ?? null, byStat);
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
        setDpsPartial: bandData.set_dps_partial ?? false,
        setDpsPartialCount: rows.filter((r) => r.notSimChecked).length,
        race: bandData.race,
        talentPoints: bandData.talent_points,
        scaleRows,
        scaleNoteLine,
        hasteCaptionLine,
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
