// web/src/lib/guild/dense/fairness.ts
// The loot council's fairness table (design/specs/2026-11-09-guild-dense.md section 5.5): every
// raider ranked by the nights they came since their last item this tier, then by fewest items;
// raiders level on both share one row (a council reads "11 raiders, 5 nights, nothing yet").
import type { GuildRosterRow } from '../api';
import type { LootDecision, LootNights } from '../loot-council-api';
import type { GuildRecentAttendance } from '../overview-api';

/** Raiders level on nights and items share a row from this many. */
export const SHARE_ROW_FROM = 3;

export interface FairnessRaider {
  key: string;
  name: string;
  cls: string;
  spec: string;
  present: number;
  nights: number;
  items: number;
  last: { item: string; at: string } | null;
}

export interface FairnessRow {
  since: number;
  items: number;
  raiders: FairnessRaider[];
}

interface NightAward {
  raidId: number | null;
  decision: LootDecision;
}

function awardsByRaider(loot: LootNights): Map<string, NightAward[]> {
  const out = new Map<string, NightAward[]>();
  const seen = new Set<number>();
  for (const night of loot.nights)
    for (const decision of night.decisions) {
      if (seen.has(decision.award_id)) continue;
      seen.add(decision.award_id);
      const list = out.get(decision.character_key) ?? [];
      list.push({ raidId: night.raid_id, decision });
      out.set(decision.character_key, list);
    }
  for (const list of out.values()) list.sort((a, b) => a.decision.at.localeCompare(b.decision.at));
  return out;
}

/** The nights a raider came after the night of their last item (every night they came, with none). */
function nightsSince(cells: readonly string[], raidIds: readonly number[], lastRaid: number | null): number {
  const from = lastRaid === null ? 0 : raidIds.indexOf(lastRaid) + 1;
  return cells.slice(Math.max(0, from)).filter((c) => c === 'present').length;
}

/** Most overdue first; raiders level on both figures share a row when SHARE_ROW_FROM or more. */
export function fairnessRows(
  roster: readonly GuildRosterRow[],
  loot: LootNights,
  grid: GuildRecentAttendance,
): FairnessRow[] {
  const awards = awardsByRaider(loot);
  const raidIds = grid.raids.map((r) => r.id);
  const cellsByAccount = new Map(grid.rows.map((r) => [r.account_key, r.cells] as const));
  const buckets = new Map<string, FairnessRow>();
  for (const row of roster) {
    const cells = cellsByAccount.get(row.account_key ?? '') ?? [];
    const won = awards.get(row.character_key) ?? [];
    const lastAward = won.at(-1) ?? null;
    const raider: FairnessRaider = {
      key: row.character_key,
      name: row.name,
      cls: row.class ?? '',
      spec: row.spec_name ?? row.spec ?? '',
      present: cells.filter((c) => c === 'present').length,
      nights: cells.filter((c) => c !== 'none').length,
      items: won.length,
      last: lastAward === null ? null : { item: lastAward.decision.item_name, at: lastAward.decision.at },
    };
    const since = nightsSince(cells, raidIds, lastAward?.raidId ?? null);
    const id = `${since}:${raider.items}`;
    const bucket = buckets.get(id) ?? { since, items: raider.items, raiders: [] };
    bucket.raiders.push(raider);
    buckets.set(id, bucket);
  }
  const rows: FairnessRow[] = [];
  for (const bucket of [...buckets.values()].sort((a, b) => b.since - a.since || a.items - b.items)) {
    bucket.raiders.sort((a, b) => a.name.localeCompare(b.name));
    if (bucket.raiders.length >= SHARE_ROW_FROM) rows.push(bucket);
    else
      for (const raider of bucket.raiders)
        rows.push({ since: bucket.since, items: bucket.items, raiders: [raider] });
  }
  return rows;
}

/** Items awarded this tier per night attended, across the guild. */
export function guildItemsPerNight(
  roster: readonly GuildRosterRow[],
  loot: LootNights,
  grid: GuildRecentAttendance,
): number {
  const keys = new Set(roster.map((r) => r.account_key ?? ''));
  const attended = grid.rows
    .filter((r) => keys.has(r.account_key))
    .flatMap((r) => r.cells)
    .filter((c) => c === 'present').length;
  const items = [...awardsByRaider(loot).values()].reduce((sum, list) => sum + list.length, 0);
  return attended === 0 ? 0 : items / attended;
}
