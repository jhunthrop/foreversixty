// web/src/lib/rankings/landing.ts
// The /rankings landing (coordinator round, 2026-11-09: "a raider's ranking tool, not a thin
// list"): each ranked encounter with its top three parses inline, and the leader of every spec on
// the raid boss with the most ranked kills. Healers lead by HPS, every other spec by DPS, as a
// raider reads a board. Pure functions over the boards the page reads; no record is made up, so
// a spec with no ranked kill is simply not on the list.
import { SPECS, type Spec, type SpecRole } from '../sim/specs';
import type { RankFigure } from '../ui/rank-list';
import type { RankedEncounter, RankingRow, RankingsPage } from './api';
import { rankingsCopy } from './copy';
import { percentileLabel, rankRowView, type RankRowView } from './rank-row-view';
import { seedFor, type RankingsSeed, type RankingsSnapshot } from './snapshot';

/** How many encounters the landing shows with their top parses, most kills first; the rest are a list. */
export const LANDING_ENCOUNTERS = 4;
/** The parses shown under each encounter. */
export const TOP_PARSES = 3;
/** The smallest raid: a boss at this size or over is a raid boss. */
export const RAID_SIZE = 10;

export type LandingMetric = 'dps' | 'hps';

/** A landing row: the board row as RankRow draws it, with two figures (the value, the percentile). */
export interface LandingRow extends RankRowView {
  key: string;
}

/** The encounter the spec leaders are read on: the raid boss with the most ranked kills, else the
 *  encounter with the most. Encounters arrive most kills first. */
export function leaderEncounter(encounters: readonly RankedEncounter[]): RankedEncounter | null {
  return encounters.find((encounter) => (encounter.size ?? 0) >= RAID_SIZE) ?? encounters[0] ?? null;
}

/** The spec a ranking row names, by its class and spec name. */
function specOf(className: string, specName: string): Spec | undefined {
  const classSlug = className.toLowerCase();
  return SPECS.find(
    (entry) => entry.class_slug === classSlug && entry.name.toLowerCase() === specName.toLowerCase(),
  );
}

/** A spec's role, from its class and spec name; DPS when the pair is unknown. */
export function specRole(className: string, specName: string): SpecRole {
  return specOf(className, specName)?.role ?? 'dps';
}

const UNIT: Readonly<Record<LandingMetric, string>> = { dps: 'DPS', hps: 'HPS' };

/** A board row with the value (its column head names the metric) and the percentile. */
export function landingRow(
  row: RankingRow,
  board: Pick<RankingsPage, 'total' | 'rows'>,
  metric: LandingMetric,
): LandingRow {
  const view = rankRowView(row, {
    total: board.total,
    top: board.rows[0]?.value ?? 0,
    unit: UNIT[metric],
    multiClass: true,
  });
  const { rankColor: parseColor, ...plain } = view;
  const value = view.figures[0];
  const figures: RankFigure[] = [
    // The column head names the metric (DPS, HPS): the figure carries no unit.
    ...(value === undefined ? [] : [{ ...value, unit: undefined }]),
    {
      value: percentileLabel(row.rank, board.total),
      label: rankingsCopy.percentileLabel,
      title: rankingsCopy.percentileLabel,
      bare: true,
      color: parseColor,
    },
  ];
  // The landing has no percentile legend: the colour sits on the percentile, the rank stays plain.
  return { ...plain, key: row.player.key, figures };
}

/** The first `count` rows of an encounter's board. */
export function topParses(
  board: Pick<RankingsPage, 'total' | 'rows'>,
  count: number = TOP_PARSES,
): LandingRow[] {
  return board.rows.slice(0, count).map((row) => landingRow(row, board, 'dps'));
}

/** The spec leaders, in two groups: damage by DPS, healing by HPS. */
export interface SpecLeaders {
  damage: LandingRow[];
  healing: LandingRow[];
}

/** A spec's full name as a player says it ("Frost Mage"), from its class and spec. */
export function specFullName(className: string, specName: string): string {
  return specOf(className, specName)?.full_name ?? `${specName} ${className}`;
}

/** One group: each spec's best row on the board, best first, numbered 1 to N within the group, the
 *  spec leading the row and the player under it; the row opens that spec's board on the boss. */
function group(
  board: Pick<RankingsPage, 'total' | 'rows'>,
  metric: LandingMetric,
  healers: boolean,
  specHref: (classSlug: string, specName: string) => string,
): LandingRow[] {
  const best = new Map<string, RankingRow>();
  for (const row of board.rows) {
    // A row whose log named no spec (the API leaves it out then) leads no spec.
    if (!row.player.spec) continue;
    if ((specRole(row.player.class, row.player.spec) === 'healer') !== healers) continue;
    const key = `${row.player.class.toLowerCase()}/${row.player.spec}`;
    if (!best.has(key)) best.set(key, row);
  }
  return [...best.values()]
    .sort((a, b) => b.value - a.value)
    .map((row, index) => {
      const view = landingRow(row, board, metric);
      const player = view.name;
      return {
        ...view,
        rank: String(index + 1),
        name: specFullName(row.player.class, row.player.spec),
        href: specHref(row.player.class.toLowerCase(), row.player.spec),
        sub: [player, row.guild?.name]
          .filter((part): part is string => part !== undefined && part !== '')
          .join(' · '),
      };
    });
}

/**
 * The best row of every spec on the boss: the damage specs (tanks among them, as a log site ranks
 * them) from the DPS board, the healers from the HPS board. Only the boards' first pages are read,
 * which hold every spec's best on today's boards.
 */
export function specLeaders(
  dps: Pick<RankingsPage, 'total' | 'rows'>,
  hps: Pick<RankingsPage, 'total' | 'rows'> | null,
  specHref: (classSlug: string, specName: string) => string,
): SpecLeaders {
  return {
    damage: group(dps, 'dps', false, specHref),
    healing: hps === null ? [] : group(hps, 'hps', true, specHref),
  };
}

/** The landing's part of the build's rankings snapshot: the encounters it lists, their top parses,
 *  and the leader encounter's boards for Leaders by spec. */
export function landingSeed(snapshot: RankingsSnapshot): RankingsSeed | null {
  return seedFor(snapshot, '', {
    shown: LANDING_ENCOUNTERS,
    topRows: TOP_PARSES,
    leader: leaderEncounter(snapshot.encounters)?.slug ?? null,
  });
}

/** One encounter's page's part of the snapshot: its board's first rows. */
export function encounterSeed(snapshot: RankingsSnapshot, slug: string): RankingsSeed | null {
  return slug === '' ? null : seedFor(snapshot, slug, { shown: 0, topRows: 0, leader: null });
}
