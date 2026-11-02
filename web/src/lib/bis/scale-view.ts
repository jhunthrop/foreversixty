// web/src/lib/bis/scale-view.ts
// Stat weights as SimulationCraft prints scale factors (owner, 2026-09-30): per point of
// stat, normalized so the top significant per-point stat reads 1.00, sorted descending, with
// DPS per point and the sim error; haste is per 1% and sits in a caption, never a row. One
// module for every surface that shows weights (the BiS rail, the stat weights tool, a saved
// weights run), moved out of panel-view.ts so the tools island carries no BiS page code.
import { statLabelForSpec } from '../sim/weights-display';
import { bisCopy, tankCopy } from './copy';
import { rateUnitFor, type SlotScoreUnit } from './tank-view';
import type { BisStatWeight } from './types';

/** `melee_haste`/`spell_haste` -- the only two haste ids a spec's own `weight_stats` ever
 *  carries (`sim/cmd/leveling-bis/report.go`'s own `isHasteStat`, mirrored here client-side
 *  for the same reason: vanilla haste has no rating conversion in this ruleset and is not
 *  comparable point-for-point against a primary/rating stat -- owner correction,
 *  2026-09-30). Never a table row on this rail; see `hasteCaptionFor` below. */
export function isHasteStat(stat: string): boolean {
  return stat === 'melee_haste' || stat === 'spell_haste';
}

export interface ScaleFactors {
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
export function computeScaleFactors(
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
  /** The word after `dpsPerPoint`: "DPS", or "score" on a tank band (`dpsPerPoint` is then
   *  raw tank score per point). */
  perPointUnit: string;
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
export function buildScaleRows(
  weights: readonly BisStatWeight[],
  spec: string,
  byStat: ReadonlyMap<string, ScaleFactors>,
  perPointUnit: string,
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
        perPointUnit,
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
 *  divisor `buildScaleRows` used otherwise.
 *
 *  `hasteOnItems` (`band.haste_on_items`, defaulted to `true` by `normaliseBisFile` for an
 *  older file) decides the caption's own "not in the table" clause -- NEVER the haste row's
 *  own `insignificant` flag, which answers a different, statistical question and produced
 *  this caption's own second-draft bug: "Haste: 1.58 per 1%, per 1%" (owner fix, 2026-09-30,
 *  found on screenshot review -- see `bisCopy.weightsHasteCaption`'s own doc). */
export function hasteCaptionFor(
  weights: readonly BisStatWeight[],
  publishedHasteScaleFactor: number | null,
  byStat: ReadonlyMap<string, ScaleFactors>,
  hasteOnItems: boolean,
): string | undefined {
  const hasteRow = weights.find((w) => isHasteStat(w.stat));
  if (hasteRow === undefined) return undefined;
  const scaleFactor = publishedHasteScaleFactor ?? byStat.get(hasteRow.stat)?.scaleFactor;
  if (scaleFactor === undefined) return undefined;
  return bisCopy.weightsHasteCaption(scaleFactor, !hasteOnItems);
}

export function perPointUnitFor(unit: SlotScoreUnit): string {
  return unit === 'tank_score' ? tankCopy.scoreWord : rateUnitFor(unit);
}

/** The rail's first line, naming the unit a point of stat is worth in. */
export function scaleNoteFor(topLabel: string, unit: SlotScoreUnit): string {
  return unit === 'tank_score'
    ? tankCopy.weightsScaleNote(topLabel)
    : bisCopy.weightsScaleNote(topLabel, rateUnitFor(unit));
}

/** One weights table, ready to render: its rows, its first line and haste's caption. */
export interface ScaleView {
  rows: ScaleRow[];
  /** "Per point of stat, normalized to {top} = 1.00, ..." -- or the unmeasured line when no
   *  significant per-point stat could anchor the table. */
  noteLine: string;
  hasteCaption: string | undefined;
}

/** A table with nothing measured: no rows, the unmeasured line, no haste caption. */
export const UNMEASURED_SCALE_VIEW: ScaleView = {
  rows: [],
  noteLine: bisCopy.weightsUnmeasuredLine,
  hasteCaption: undefined,
};

export interface ScaleViewInput {
  weights: readonly BisStatWeight[];
  spec: string;
  unit: SlotScoreUnit;
  /** DPS per point of the weights' own reference stat (weight 1), when known. */
  referenceDpsPerPoint: number | null;
  publishedAnchorStat: string | null;
  publishedHasteScaleFactor: number | null;
  /** Whether any item the weights were measured for carries haste (the caption's clause). */
  hasteOnItems: boolean;
}

/** The SimulationCraft-form table for one set of weights (see this module's header). */
export function scaleViewFor(input: ScaleViewInput): ScaleView {
  const { byStat, anchorStat } = computeScaleFactors(
    input.weights,
    input.referenceDpsPerPoint,
    input.publishedAnchorStat,
  );
  return {
    rows: buildScaleRows(input.weights, input.spec, byStat, perPointUnitFor(input.unit)),
    noteLine:
      anchorStat === ''
        ? bisCopy.weightsUnmeasuredLine
        : scaleNoteFor(statLabelForSpec(anchorStat, input.spec), input.unit),
    hasteCaption: hasteCaptionFor(input.weights, input.publishedHasteScaleFactor, byStat, input.hasteOnItems),
  };
}
