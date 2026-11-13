// web/src/lib/tiers/tier-row-view.ts
// A tier row as the ui/RankRow preset reads it (directions mock g-tiers, round 5): the spec's
// full name as the row's main link (its best in slot; a DPS row's figure opens the ranking sim
// run behind it), the accent bar, the figures under heads
// that name their units, and the gap to the top. Pure, so the row's text is tested once here and
// the page only places it.
import type { RankFigure, RankGap, RankNote } from '../ui/rank-list';
import { specViewKey } from './tier-callout';
import { formatMetric, gapText, specFullName, tiersCopy } from './tier-copy';
import { hrefsFor, type TierRow } from './tier-list';

export interface TierRankRowProps {
  rank: string;
  name: string;
  href: string;
  classSlug: string;
  fraction: number;
  barTone: 'accent';
  figures: RankFigure[];
  gap: RankGap;
  key: string;
  nameTestid: string;
  testid: string;
  note?: RankNote;
}

/** The bar starts at 90% of the lowest spec, so the field's spread shows (directions spec, Tier list). */
const BAR_FLOOR = 0.9;

/** Each row's bar on its list. A DPS or healer bar runs from 90% of the lowest spec's share to
 *  the top; a tank bar is the least damage taken over the row's own, from zero, so a third more
 *  damage taken reads as a bar a third shorter (three rows have no spread to stretch). */
export function barFraction(
  row: Pick<TierRow, 'fraction' | 'role'>,
  rows: readonly Pick<TierRow, 'fraction'>[],
): number {
  if (row.role === 'tank') return row.fraction;
  const floor = Math.min(...rows.map((r) => r.fraction)) * BAR_FLOOR;
  return (row.fraction - floor) / (1 - floor);
}

function isLeader(row: TierRow): boolean {
  return row.ranked && row.rank === 1 && row.gapPercent === 0;
}

function bestOf(on: boolean | undefined, title: string): Pick<RankFigure, 'best' | 'bestTitle'> {
  return on === true ? { best: true, bestTitle: title } : {};
}

/** The sorted figure, then on a tank its maximum health (once the band publishes it) and threat,
 *  each green where best, and the phone's vs #1. The column heads carry the units. */
export function tierFigures(row: TierRow, count: number): RankFigure[] {
  const title = tiersCopy.bestTitle(count, row.role);
  const value = formatMetric(row.role, row.metric);
  const sorted: RankFigure = {
    value,
    label: tiersCopy.figureLabels[row.role],
    testid: 'tier-number',
    ...bestOf(row.role === 'tank' && isLeader(row), title),
    ...simLinkOf(row, value),
  };
  if (row.role !== 'tank' || row.threat === undefined) return [sorted, phoneGap(row)];
  const health: RankFigure[] =
    row.health === undefined
      ? []
      : [
          {
            value: formatMetric(row.role, row.health),
            label: tiersCopy.figureLabels.health,
            title: tiersCopy.healthTitle,
            testid: 'tier-health',
            ...bestOf(row.isBestHealth, title),
          },
        ];
  return [
    sorted,
    ...health,
    {
      value: formatMetric(row.role, row.threat),
      label: tiersCopy.figureLabels.threat,
      testid: 'tier-threat',
      ...bestOf(row.isBestThreat, title),
    },
    phoneGap(row),
  ];
}

/** A DPS row's figure links the ranking sim run behind it (`simHref`; absent on tanks and
 *  healers, whose numbers the browser sim does not produce). */
function simLinkOf(row: TierRow, value: string): Pick<RankFigure, 'href' | 'title' | 'ariaLabel'> {
  if (row.simHref === undefined) return {};
  const spec = specFullName(row);
  return {
    href: row.simHref,
    title: tiersCopy.simLinkTitle(spec),
    ariaLabel: tiersCopy.simLinkLabel(spec, value),
  };
}

/** Below 768 the gap column goes: the row's vs #1 rides under its figure instead. */
function phoneGap(row: TierRow): RankFigure {
  return {
    value: gapOf(row),
    label: tiersCopy.gapLabel[row.role],
    bare: true,
    phoneOnly: true,
    testid: 'tier-gap-phone',
  };
}

/** "−7.3%" behind the top spec, "+53% taken" on a tank; the top spec itself reads "best", the
 *  least-hit tank "lowest". */
function gapOf(row: TierRow): string {
  if (!isLeader(row)) return gapText(row);
  return row.role === 'tank' ? tiersCopy.tankLeaderGap : tiersCopy.leaderGap;
}

/** The leader, for a tank gap's title. */
function leaderOf(rows: readonly TierRow[]): TierRow | undefined {
  return rows.find(isLeader);
}

export function rankRowProps(row: TierRow, rows: readonly TierRow[]): TierRankRowProps {
  const count = rows.length;
  return {
    rank: tiersCopy.rankLabel(row),
    name: specFullName(row),
    href: hrefsFor(row).bis,
    classSlug: row.classSlug,
    fraction: barFraction(row, rows),
    barTone: 'accent',
    figures: tierFigures(row, count),
    // The leader's gap reads "best": the column reads against it (the mock's "vs #1").
    gap: {
      text: gapOf(row),
      label: tiersCopy.gapLabel[row.role],
      leader: isLeader(row),
      ...gapTitleOf(row, rows),
    },
    key: specViewKey(row.classSlug, row.specSlug),
    nameTestid: 'tier-spec-link',
    testid: `tier-row-${row.classSlug}-${row.specSlug}`,
    ...noteOf(row),
  };
}

/** A tank's gap in words on hover; a tank's visible "+53% taken" stands alone on touch. */
function gapTitleOf(row: TierRow, rows: readonly TierRow[]): Pick<RankGap, 'title'> {
  const leader = leaderOf(rows);
  if (row.role !== 'tank' || leader === undefined || isLeader(row)) return {};
  return { title: tiersCopy.tankGapTitle(row.gapPercent, specFullName(leader)) };
}

/** The line under the name: a tied healer's partner (so a shared rank explains itself), and the
 *  less-certain stat weights, joined when a row has both. */
function noteOf(row: TierRow): Pick<TierRankRowProps, 'note'> {
  const tie = tiersCopy.rowTie(row);
  const notes: { text: string; title: string }[] = [
    ...(tie === null ? [] : [tie]),
    ...(row.lowConfidence ? [{ text: tiersCopy.lowConfidence, title: tiersCopy.lowConfidenceTitle }] : []),
  ];
  if (notes.length === 0) return {};
  return {
    note: { text: notes.map((n) => n.text).join(' · '), title: notes.map((n) => n.title).join('. ') },
  };
}
