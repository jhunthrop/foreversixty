// web/src/lib/guild/progression-view.ts
// The Progression tab's words (design/specs/2026-11-05-guild-week-one.md §3.4; batch 3, one panel a
// zone): the pre-raid empty state, each zone's bosses down, and a boss's one line, its first kill
// printed only when the API sent it. Pure, so the page never prints a field the API omitted.
import { nightDate } from '../dates';
import { counted } from '../plural';
import type { GuildProgressionEncounter } from './api';
import { raidsAreOpen, raidsOpenDay } from './week-one';

/** `/logs` covers both: how to log a raid, and the upload form. */
const LOGS_HREF = '/logs';

/** The tab with no pulls in the tier: before raids open it says when, after it says to upload. */
export function progressionEmpty(now: Date): { message: string; action: { label: string; href: string } } {
  if (raidsAreOpen(now)) {
    return {
      message: 'No pulls logged yet. Upload your first raid night and it shows here.',
      action: { label: 'Upload a log', href: LOGS_HREF },
    };
  }
  return {
    message: `Barrow Deeps, Hyjal Summit and Onyxia's Lair open ${raidsOpenDay()}. Progression starts with your first logged pull.`,
    action: { label: 'How to log a raid', href: LOGS_HREF },
  };
}

/** A zone's panel: its named bosses down and the bosses pulled, in the API's order. */
export interface ProgressionZone {
  zone: string;
  down: number;
  named: number;
  pulled: GuildProgressionEncounter[];
}

/** One panel per zone with a boss pulled: the zone's bosses down of its named bosses. */
export function progressionZones(encounters: readonly GuildProgressionEncounter[]): ProgressionZone[] {
  const zones = new Map<string, GuildProgressionEncounter[]>();
  for (const encounter of encounters)
    zones.set(encounter.zone, [...(zones.get(encounter.zone) ?? []), encounter]);
  return [...zones.entries()]
    .map(([zone, list]) => ({
      zone,
      down: list.filter((encounter) => encounter.kills > 0).length,
      named: list.length,
      pulled: list.filter((encounter) => encounter.pulls > 0),
    }))
    .filter((zone) => zone.pulled.length > 0);
}

/** `1 of 1 boss down`. */
export function bossesDownMeta(zone: Pick<ProgressionZone, 'down' | 'named'>): string {
  return `${zone.down} of ${counted(zone.named, 'boss', 'bosses')} down`;
}

/** `2 pulls · 1 wipe · 3.5 deaths a pull · first kill Wed 9 Dec`. */
export function bossLine(encounter: GuildProgressionEncounter, now: Date): string {
  const firstKill = encounter.first_kill_at;
  return [
    counted(encounter.pulls, 'pull'),
    counted(Math.max(0, encounter.pulls - encounter.kills), 'wipe'),
    `${encounter.deaths_per_pull.toFixed(1)} deaths a pull`,
    typeof firstKill === 'string' ? `first kill ${nightDate(new Date(firstKill), now)}` : '',
  ]
    .filter((part) => part !== '')
    .join(' · ');
}
