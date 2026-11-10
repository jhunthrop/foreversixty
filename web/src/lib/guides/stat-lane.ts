// web/src/lib/guides/stat-lane.ts
// The guide's stat model, read once for the At a glance column and the Stat priority lane
// (guides spec 2026-11-08 §4.3): the band 60 weights in the sim's order (`railStatRows`), the
// hit and expertise cap lines, the unit line, the rating note and the notes for the rows that
// carry no figure. Both blocks read this, so they can never disagree.
import { band60Weights } from './class-dps';
import { guidesCopy } from './copy';
import { guideExpertiseCapLine, guideHitCapLine, type GuideHitCapLine } from './hit-cap-line';
import { railStatRows, railStatValueText, type RailStatRow } from './rail-stats';
import { bandTalentHits } from './talent-hit';

type GuideRole = 'dps' | 'healer' | 'tank';

export interface StatLaneRow extends RailStatRow {
  /** What the weight cell prints: the figure, or "Not on any item", "Not significant", "No gain". */
  valueText: string;
  /** What the per-point cell prints; '' when the row has no figure. */
  perPointText: string;
}

export interface StatLane {
  role: GuideRole;
  rows: StatLaneRow[];
  hitCap?: GuideHitCapLine;
  expertiseCap?: { text: string };
  weightHead: string;
  perPointHead: string;
  /** "1 hit rating = 1.13 DPS at level 60." */
  unitLine?: string;
  ratingNote?: string;
  /** The notes under the table for rows with no figure. */
  notes: string[];
}

/** A per-point figure as the table prints it. */
export function perPointText(value: number | undefined): string {
  if (value === undefined) return '';
  return value >= 10 ? value.toFixed(1) : value.toFixed(2);
}

/** The stat model for a spec, or undefined when it has no band 60 weights (the lane then lists
 *  the guide's own `statPriority` order). */
export function guideStatLane(
  build: string,
  classSlug: string,
  spec: string,
  recommendedRaces: readonly string[],
  statPriority: readonly string[],
): StatLane | undefined {
  const band60 = band60Weights(build, spec, classSlug, [...recommendedRaces]);
  if (band60 === undefined) return undefined;
  const raw = railStatRows([...statPriority], band60.weights, spec, band60.haste);
  if (raw.length === 0) return undefined;
  const role: GuideRole = band60.role ?? 'dps';
  const rows = raw.map((row) => ({
    ...row,
    valueText: railStatValueText(row),
    perPointText: perPointText(row.perPoint),
  }));
  const valued = rows.filter((row) => row.reading === 'value');
  const talentHits = bandTalentHits(build, classSlug, band60.talents);
  const hitCap = guideHitCapLine(band60.hitToCap, talentHits);
  const expertiseCap = guideExpertiseCapLine(band60.expertiseToCap, talentHits);
  const unitRow = valued.find((row) => row.perPercent === undefined && row.perPoint !== undefined);
  const rated = valued.filter((row) => row.ratingFactor !== undefined);
  const insignificant = rows.filter((row) => row.reading === 'not-significant').map((row) => row.label);
  const noGain = rows.filter((row) => row.reading === 'no-gain').map((row) => row.label);
  const notes = [
    ...(insignificant.length > 0
      ? [guidesCopy.notSignificantNote(insignificant.join(' and '), insignificant.length > 1)]
      : []),
    ...(noGain.length > 0 ? [guidesCopy.noGainNote(noGain.join(' and '), noGain.length > 1, role)] : []),
    ...rows
      .filter((row) => row.reading === 'not-on-items')
      .map((row) => guidesCopy.notOnItemsNote(row.label, row.value?.toFixed(1), role)),
  ];
  return {
    role,
    rows,
    ...(hitCap === undefined ? {} : { hitCap }),
    ...(expertiseCap === undefined ? {} : { expertiseCap }),
    weightHead: guidesCopy.weightColumn,
    perPointHead: guidesCopy.perPointColumn(role),
    ...(unitRow === undefined
      ? {}
      : {
          unitLine: guidesCopy.unitLine(
            unitRow.label.toLowerCase(),
            perPointText(unitRow.perPoint),
            role,
            unitRow.ratingFactor !== undefined,
          ),
        }),
    ...(rated.length === 0
      ? {}
      : {
          ratingNote: guidesCopy.ratingNote(
            rated.map((row) => row.label).join(' and '),
            rated.length > 1,
            rated
              .map((row) => guidesCopy.ratingConversion(row.label.toLowerCase(), row.ratingFactor!))
              .join(' and '),
          ),
        }),
    notes,
  };
}
