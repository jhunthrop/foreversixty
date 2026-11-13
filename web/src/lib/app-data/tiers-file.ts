// web/src/lib/app-data/tiers-file.ts
// /data/{build}/tiers.json (Wave 3 contract W3-C1): the web's tier lists, already ranked. The
// order, gaps, letters, ties and every sentence come from the same functions /tiers renders
// with (`loadTierData`, `rankRole`, `tierSpecViews`, `readColumns`, `gapText`), so neither app
// sorts, gaps or letters anything (spec F9.4).
import type { BisRole, Faction } from '../bis/types';
import { readSpecCatalog } from '../bis/load';
import { tierSpecViews, specViewKey } from '../tiers/tier-callout';
import { formatMetric, gapText, readColumns, tiersCopy } from '../tiers/tier-copy';
import { TIER_ROLES, type TierRow } from '../tiers/tier-list';
import type { RoleLists, TierData } from '../tiers/tier-load';

export interface AppTierRow {
  rank: number;
  spec: string;
  class: string;
  spec_slug: string;
  /** `Fury Warrior` */
  name: string;
  race: string;
  icon: string | null;
  metric: number;
  /** `512.3` (the role's own rounding) and its unit word. */
  metric_text: string;
  unit: string;
  gap_percent: number;
  /** `Top`, `−2.8%`, `+15.2%`. */
  gap_text: string;
  /** 0..1, the bar's length. */
  fraction: number;
  /** `above`, `below`, `both` or null. */
  tie: string | null;
  /** S to D, null on tank and healer lists. */
  tier: string | null;
  low_confidence: boolean;
  /** The callout's sentences when this is the reader's spec. */
  title: string;
  detail: string;
}

export interface AppTierList {
  /** The web's "How to read this" columns, plain text, in the order the web prints them. */
  notes: string[];
  rows: AppTierRow[];
}

export interface AppTiersFile {
  build: string;
  /** The newest file's `generated_at` day, `2026-10-31`. */
  updated: string | null;
  preset_label: string | null;
  factions: Record<Faction, Record<BisRole, AppTierList>>;
}

const FACTIONS: readonly Faction[] = ['alliance', 'horde'];

function iconOf(spec: string): string | null {
  return readSpecCatalog().find((entry) => entry.spec === spec)?.icon ?? null;
}

function rowOf(row: TierRow, views: ReturnType<typeof tierSpecViews>): AppTierRow {
  const view = views[specViewKey(row.classSlug, row.specSlug)];
  if (view === undefined) throw new Error(`tiers file: no callout for ${row.spec}`);
  return {
    rank: row.rank,
    spec: row.spec,
    class: row.classSlug,
    spec_slug: row.specSlug,
    name: view.fullName,
    race: row.race,
    icon: iconOf(row.spec),
    metric: row.metric,
    metric_text: formatMetric(row.role, row.metric),
    unit: tiersCopy.units[row.role],
    gap_percent: row.gapPercent,
    gap_text: gapText(row),
    fraction: row.fraction,
    tie: row.tie,
    tier: row.tier,
    low_confidence: row.lowConfidence,
    title: view.title,
    detail: view.detail,
  };
}

function listsOf(
  build: string,
  data: TierData,
  faction: Faction,
  lists: RoleLists,
): Record<BisRole, AppTierList> {
  const views = tierSpecViews(lists);
  return Object.fromEntries(
    TIER_ROLES.map((role) => {
      const notes =
        lists[role].length === 0 || data.presetLabel === null
          ? []
          : readColumns({
              role,
              faction,
              presetLabel: data.presetLabel,
              healSeconds: data.healSeconds,
              healerBasis: data.healerBasis,
              raidFightSeconds: data.raidFightSeconds,
              healerOrderMeasured: data.healerOrderMeasured[faction],
              bossSwingSeconds: data.bossSwingSeconds,
            }).map((column) => column.body.map((segment) => segment.text).join(''));
      return [role, { notes, rows: lists[role].map((row) => rowOf(row, views)) }];
    }),
  ) as Record<BisRole, AppTierList>;
}

export function buildTiersFile(build: string, data: TierData): AppTiersFile {
  return {
    build,
    updated: data.updated === null ? null : data.updated.toISOString().slice(0, 10),
    preset_label: data.presetLabel,
    factions: Object.fromEntries(
      FACTIONS.map((faction) => [faction, listsOf(build, data, faction, data.lists[faction])]),
    ) as Record<Faction, Record<BisRole, AppTierList>>,
  };
}
