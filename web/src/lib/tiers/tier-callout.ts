// web/src/lib/tiers/tier-callout.ts
// The signed-in answer on the tier list: the visitor's own spec, ranked, in the spec panel beside
// the list. The page embeds one ready-made view per spec and faction, and a small script picks
// the one that matches the visitor's selected character, so the text is written once, here, and
// tested. The apps' tiers file reads the same views (lib/app-data/tiers-file.ts).
import { classSlugFromName } from '../report/tree-sizes';
import type { BisRole } from '../bis/types';
import { bisBandPath } from '../bis/band-path';
import { MAX_BAND, MIN_BAND } from '../bis/band-levels';
import { ROLE_LABELS, ROLE_PATHS, formatMetric, specFullName, tiersCopy } from './tier-copy';
import type { TierRow } from './tier-list';
import { ordinal } from '../format/number';

/** What the page needs to answer for one spec. Every field is final text or a path. */
export interface TierSpecView {
  role: BisRole;
  classSlug: string;
  /** "Fury Warrior" */
  fullName: string;
  /** "Fury Warrior is 2nd of 20 DPS specs.": the apps' callout (lib/app-data/tiers-file.ts). */
  title: string;
  /** The spec panel's line: "Tier A · #2 · 1,004.2 DPS", "=2 · 912.4 HPS", or the figure alone
   *  on a list not ranked yet. */
  standing: string;
  /** "Tier A · 8.2% behind Marksmanship Hunter" */
  detail: string;
  bisHref: string;
  guideHref: string;
  /** The role's path, where an answer on another role's page points. */
  rolePath: string;
}

export type TierSpecViews = Record<string, TierSpecView>;

/** The key the script looks a character up by: `<class slug>/<spec slug>`. */
export function specViewKey(classSlug: string, specSlug: string): string {
  return `${classSlug}/${specSlug}`;
}

/** The lookup key for a signed-in character, or `undefined` while its class or spec is not
 *  known yet (never a guess at either). */
export function specViewKeyForCharacter(character: { class?: string; spec?: string }): string | undefined {
  if (character.class === undefined || character.spec === undefined) return undefined;
  return specViewKey(classSlugFromName(character.class), classSlugFromName(character.spec));
}

function detailFor(row: TierRow, top: TierRow): string {
  if (!row.ranked) return 'Ranks appear once the next ranking publishes the measured error.';
  if (row.role === 'healer' && row.tie !== null) return 'Tied';
  if (row.rank === 1) return row.role === 'tank' ? 'Takes the least damage.' : 'Level with the top spec.';
  const gap =
    row.role === 'tank'
      ? tiersCopy.tankGapTitle(row.gapPercent, specFullName(top))
      : `${row.gapPercent.toFixed(1)}% behind ${specFullName(top)}`;
  return row.tier === null ? gap : `Tier ${row.tier} · ${gap}`;
}

/** Where a row stands, in words: no place on a list not ranked yet, "tied" on a shared rank. */
function placeTitle(row: TierRow, fullName: string, count: number): string {
  const role = ROLE_LABELS[row.role];
  if (!row.ranked) {
    return `${fullName} is on the ${role} list; ranks appear once the next ranking publishes the measured error.`;
  }
  const tied = tiersCopy.sharesRank(row) ? 'tied ' : '';
  return `${fullName} is ${tied}${ordinal(row.rank)} of ${count} ${role} specs.`;
}

function standingOf(row: TierRow): string {
  const figure = `${formatMetric(row.role, row.metric)} ${tiersCopy.units[row.role]}`;
  if (!row.ranked) return figure;
  const place = tiersCopy.sharesRank(row) ? `=${row.rank}` : `#${row.rank}`;
  const tier = row.tier === null ? [] : [tiersCopy.tierLabel(row.tier)];
  return [...tier, place, figure].join(' · ');
}

/** One view per row of one faction's lists. */
export function tierSpecViews(lists: Record<BisRole, readonly TierRow[]>): TierSpecViews {
  const views: TierSpecViews = {};
  for (const [role, rows] of Object.entries(lists) as [BisRole, readonly TierRow[]][]) {
    const top = rows[0];
    if (top === undefined) continue;
    for (const row of rows) {
      const path = `${row.classSlug}/${row.specSlug}`;
      const fullName = specFullName(row);
      views[specViewKey(row.classSlug, row.specSlug)] = {
        role,
        classSlug: row.classSlug,
        fullName,
        title: placeTitle(row, fullName, rows.length),
        standing: standingOf(row),
        detail: detailFor(row, top),
        bisHref: bisBandPath(row.classSlug, row.specSlug, MAX_BAND, MIN_BAND),
        guideHref: `/guides/${path}`,
        rolePath: ROLE_PATHS[role],
      };
    }
  }
  return views;
}
