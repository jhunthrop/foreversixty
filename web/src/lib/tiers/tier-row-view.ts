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

/** Each row's bar on its list: 0 at 90% of the lowest spec's share, 1 at the top. */
export function barFraction(
  row: Pick<TierRow, 'fraction'>,
  rows: readonly Pick<TierRow, 'fraction'>[],
): number {
  const floor = Math.min(...rows.map((r) => r.fraction)) * BAR_FLOOR;
  return (row.fraction - floor) / (1 - floor);
}

function isLeader(row: TierRow): boolean {
  return row.ranked && row.rank === 1 && row.gapPercent === 0;
}

function bestOf(on: boolean | undefined, title: string): Pick<RankFigure, 'best' | 'bestTitle'> {
  return on === true ? { best: true, bestTitle: title } : {};
}

/** The sorted figure, then on a tank its effective health and threat, each green where best. The
 *  column heads carry the units. */
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
  if (row.role !== 'tank' || row.effectiveHealth === undefined || row.threat === undefined)
    return [sorted, phoneGap(row)];
  return [
    sorted,
    {
      value: formatMetric(row.role, row.effectiveHealth),
      label: tiersCopy.figureLabels.effectiveHealth,
      testid: 'tier-effective-health',
      ...bestOf(row.isBestEffectiveHealth, title),
    },
    {
      value: formatMetric(row.role, row.threat),
      label: tiersCopy.figureLabels.threat,
      testid: 'tier-threat',
      ...bestOf(row.isBestThreat, title),
    },
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

/** "−7.3%" behind the top spec; the top spec itself reads "best". */
function gapOf(row: TierRow): string {
  return isLeader(row) ? tiersCopy.leaderGap : gapText(row);
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
    },
    key: specViewKey(row.classSlug, row.specSlug),
    nameTestid: 'tier-spec-link',
    testid: `tier-row-${row.classSlug}-${row.specSlug}`,
    ...noteOf(row),
  };
}

/** The line under the name: a tied healer's partner (so a shared rank explains itself), and the
 *  less-certain stat weights, joined when a row has both. */
function noteOf(row: TierRow): Pick<TierRankRowProps, 'note'> {
  const tie = tiersCopy.healerTie(row);
  const notes: { text: string; title: string }[] = [
    ...(tie === null ? [] : [tie]),
    ...(row.lowConfidence ? [{ text: tiersCopy.lowConfidence, title: tiersCopy.lowConfidenceTitle }] : []),
  ];
  if (notes.length === 0) return {};
  return {
    note: { text: notes.map((n) => n.text).join(' · '), title: notes.map((n) => n.title).join('. ') },
  };
}
