// web/src/lib/guild/dense/member-week.ts
// A member's week on the Overview (design/specs/2026-11-09-guild-dense.md section 4): their
// attendance streak, their own line on the last two logged nights (figure, rating, percentile,
// deaths) and their loot against the guild's. Pure: built from the member's own profile, the two
// raid reads (a member's report carries only their own parses) and the home's roster.
import type { GuildRaidView, GuildRosterRow } from '../api';
import type { ProfileNight } from '../recruit-api';

/** Nights present in a row, counting back from the newest. */
export function streak(nights: readonly ProfileNight[]): number {
  let n = 0;
  for (const night of [...nights].sort((a, b) => b.starts_at.localeCompare(a.starts_at))) {
    if (night.outcome !== 'present' && night.outcome !== 'late') break;
    n += 1;
  }
  return n;
}

export interface MyNight {
  raidId: number;
  startsAt: string;
  boss: string;
  killMs: number | null;
  value: number | null;
  metric: string;
  rating: number | null;
  percentile: number | null;
  deaths: number;
}

/** The member's own line on one logged night: their best rated kill, and their deaths that night. */
export function myNight(view: GuildRaidView, keys: ReadonlySet<string>): MyNight | null {
  const report = view.report;
  if (report === null) return null;
  const mine = report.parses.filter((p) => keys.has(p.character_key));
  const best = mine.reduce<(typeof mine)[number] | null>(
    (top, p) => (top === null || (p.rating ?? 0) > (top.rating ?? 0) ? p : top),
    null,
  );
  const kill =
    report.kill_order.find((k) => best !== null && k.encounter_id === best.encounter_id) ??
    report.kill_order[0];
  const deaths = report.deaths
    .filter((d) => keys.has(d.character_key))
    .reduce((sum, d) => sum + d.on_kills + d.on_wipes, 0);
  return {
    raidId: view.id,
    startsAt: view.starts_at,
    boss: best?.encounter ?? kill?.name ?? '',
    killMs: kill?.kill_ms ?? null,
    value: best?.value ?? null,
    metric: best?.metric ?? 'dps',
    rating: best?.rating ?? null,
    percentile: best?.percentile ?? null,
    deaths,
  };
}

/** Items the guild won this tier per night a raider came, from the roster's attendance. */
export function guildPerNight(roster: readonly GuildRosterRow[], items: number): number {
  const nights = roster.reduce((sum, r) => sum + ((r.attendance_30 ?? r.attendance)?.present ?? 0), 0);
  return nights > 0 ? items / nights : 0;
}
