// web/src/lib/guild/overview-view.ts
// The guild page's header band and Overview cards in a player's words (lifecycle review round 4,
// N8/N9): plurals from the one helper, days as words, and the counts only where the viewer can
// open what they count. Pure, so every rule is pinned without a page.
import { nightDate, nightDateOnly } from '../dates';
import { counted } from '../plural';
import type { GuildViewerRole } from './api';
import type { Fact } from './fact-line';

/** "1 boss down · 2 pulls". */
export function bandCountsLine(killed: number, pulls: number): string {
  return `${counted(killed, 'boss', 'bosses')} down · ${counted(pulls, 'pull')}`;
}

/**
 * Whether the band prints boss and pull counts. A member or officer counts every report the
 * guild holds. A visitor counts what the public tabs list, so when the only report is
 * guild-only the band says nothing about bosses beside a tab reading "No public reports this
 * week yet".
 */
export function showBandCounts(role: GuildViewerRole, publicReportCount: number): boolean {
  return role !== 'public' || publicReportCount > 0;
}
/** The date of a report's `created_at` (an instant) for a raid-night row: "Tue 20 Oct", the same words
 *  the account's report row prints for the same night. */
export function raidNightDay(createdAt: string, now: Date = new Date()): string {
  return nightDate(new Date(createdAt), now);
}

/** A progression night's date, which the API sends as a calendar date. */
export function progressionNightDay(date: string, now: Date = new Date()): string {
  return date.length > 10 ? nightDate(new Date(date), now) : nightDateOnly(date, now);
}

/** The named encounters the guild has killed, in the progression read's order: the names every
 *  "down this tier" line lists, so an encounter only pulled is never named as down. */
export function downEncounterNames(encounters: readonly { name: string; kills: number }[]): string[] {
  return encounters.filter((encounter) => encounter.kills > 0).map((encounter) => encounter.name);
}

/** "1 of 1 named bosses down: Onyxia." / "No boss down this tier yet." */
export function bossesDownLine(names: readonly string[], named: number): string {
  if (names.length === 0) return 'No boss down this tier yet.';
  return `${names.length} of ${named} named bosses down: ${names.join(', ')}.`;
}

/** The same line before the progression read names the bosses: the home's own count. */
export function bossesDownCountLine(down: number, named: number): string {
  return down === 0 ? 'No boss down this tier yet.' : `${down} of ${named} named bosses down.`;
}

/** The Before raids open card shows while nothing is logged in the tier. */
export function showsWeekOne(role: GuildViewerRole, tier: { pulls_total: number } | undefined): boolean {
  return role !== 'public' && (tier === undefined || tier.pulls_total === 0);
}

/** The Roster card's line, "1 raider with gear on the site · 4 waiting for approval · 2 below the
 *  rating floor": a zero clause after the first is left out, like the roster's count line (guild
 *  loop round 3: "0 waiting · 0 below the rating floor"). Waiting reads gold, an officer's to-do. */
export function rosterCardFacts(summary: {
  raiders: number;
  waiting_for_approval: number;
  below_rating_floor: number;
}): Fact[] {
  const facts: Fact[] = [{ text: `${counted(summary.raiders, 'raider')} with gear on the site` }];
  if (summary.waiting_for_approval > 0)
    facts.push({ text: `${summary.waiting_for_approval} waiting for approval`, class: 'text-gold' });
  if (summary.below_rating_floor > 0)
    facts.push({ text: `${summary.below_rating_floor} below the rating floor` });
  return facts;
}

export const NO_PUBLIC_REPORTS = 'No public reports this week yet.';
/** The Progression card with nothing the viewer can open: progression is not weekly, so it never
 *  borrows the raid-nights card's `this week` (guild-page spec §4.A, guild loop round 3). */
export const NO_PUBLIC_LOGS = 'No public logs yet.';
export const NO_PUBLIC_RAID_NIGHTS = 'No public raid nights yet.';
