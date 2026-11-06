// web/src/lib/tiers/tier-row-view.ts
// A tier row as the ui/RankRow preset reads it (tiers spec 2026-11-06 §4.4): the name, the sub
// line, the figures with their units, the gap, the tie and its partners, and the four next steps.
// Pure, so the row's text is tested once here and the page only places it.
import type { RankFigure, RankGap, RankLink, RankNote } from '../ui/rank-list';
import { specViewKey } from './tier-callout';
import {
  formatMetric,
  gapText,
  rowName,
  rowSub,
  specFullName,
  tieText,
  tieTitle,
  tiersCopy,
} from './tier-copy';
import { hrefsFor, type TierRow } from './tier-list';

export interface TierRankRowProps {
  rank: string;
  name: string;
  href: string;
  classSlug: string;
  sub: string;
  fraction: number;
  figures: RankFigure[];
  gap: RankGap;
  links: RankLink[];
  key: string;
  nameTestid: string;
  testid: string;
  tie?: RankNote;
  note?: RankNote;
}

function isLeader(row: TierRow): boolean {
  return row.ranked && row.rank === 1 && row.gapPercent === 0;
}

function bestOf(on: boolean | undefined, title: string): Pick<RankFigure, 'best' | 'bestTitle'> {
  return on === true ? { best: true, bestTitle: title } : {};
}

/** The sorted figure, then on a tank its effective health and threat, each green where best. */
export function tierFigures(row: TierRow, count: number): RankFigure[] {
  const title = tiersCopy.bestTitle(count, row.role);
  const sorted: RankFigure = {
    value: formatMetric(row.role, row.metric),
    unit: tiersCopy.units[row.role],
    label: tiersCopy.figureLabels[row.role],
    testid: 'tier-number',
    ...bestOf(row.role === 'tank' && isLeader(row), title),
  };
  if (row.role !== 'tank' || row.effectiveHealth === undefined || row.threat === undefined) return [sorted];
  return [
    sorted,
    {
      value: formatMetric(row.role, row.effectiveHealth),
      unit: tiersCopy.units.effectiveHealth,
      label: tiersCopy.figureLabels.effectiveHealth,
      testid: 'tier-effective-health',
      ...bestOf(row.isBestEffectiveHealth, title),
    },
    {
      value: formatMetric(row.role, row.threat),
      unit: tiersCopy.units.threat,
      label: tiersCopy.figureLabels.threat,
      testid: 'tier-threat',
      ...bestOf(row.isBestThreat, title),
    },
  ];
}

/** BiS, Guide, Planner and, where /sim can rerun the number, the ranking Sim. */
export function tierLinks(row: TierRow): RankLink[] {
  const href = hrefsFor(row);
  const fullName = specFullName(row);
  const { bis, guide, planner, sim } = tiersCopy.rowLinks;
  const links: RankLink[] = [
    { label: bis, href: href.bis, primary: true, ariaLabel: `${bis}: ${fullName}` },
    { label: guide, href: href.guide, ariaLabel: `${guide}: ${fullName}` },
    {
      label: planner,
      href: href.planner,
      ariaLabel: `${planner}: ${fullName}`,
      title: tiersCopy.plannerTitle(fullName),
    },
  ];
  if (href.sim === undefined) return links;
  return [
    ...links,
    {
      label: sim,
      href: href.sim,
      ariaLabel: tiersCopy.simLinkLabel(fullName),
      title: tiersCopy.simLinkTitle(fullName),
      testid: 'tier-sim-link',
    },
  ];
}

export function rankRowProps(row: TierRow, count: number): TierRankRowProps {
  const tie = tieText(row);
  return {
    rank: tiersCopy.rankLabel(row),
    name: rowName(row),
    href: hrefsFor(row).bis,
    classSlug: row.classSlug,
    sub: rowSub(row),
    fraction: row.fraction,
    figures: tierFigures(row, count),
    gap: { text: gapText(row), label: tiersCopy.gapLabel[row.role], leader: isLeader(row) },
    links: tierLinks(row),
    key: specViewKey(row.classSlug, row.specSlug),
    nameTestid: 'tier-spec-link',
    testid: `tier-row-${row.classSlug}-${row.specSlug}`,
    ...(tie === null ? {} : { tie: { text: tie, title: tieTitle(row) } }),
    ...(row.lowConfidence
      ? { note: { text: tiersCopy.lowConfidence, title: tiersCopy.lowConfidenceTitle } }
      : {}),
  };
}
