// web/src/lib/guild/overview-view.ts
// The guild page's header band and Overview cards in a player's words (lifecycle review round 4,
// N8/N9): plurals from the one helper, days as words, and the counts only where the viewer can
// open what they count. Pure, so every rule is pinned without a page.
import { dayWord, nightDate, nightDateOnly } from '../dates';
import { counted } from '../plural';
import type { GuildViewerRole } from './api';

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

/** "Updated today": the day in the viewer's own zone. Empty for a stamp it cannot read. */
export function updatedLine(updatedAt: string, now: Date = new Date()): string {
  const word = dayWord(new Date(updatedAt), now);
  return word === '' ? '' : `Updated ${word}`;
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

/** "1 encounter down this tier: Onyxia." / "No encounter down this tier yet." */
export function encountersDownLine(names: readonly string[]): string {
  if (names.length === 0) return 'No encounter down this tier yet.';
  return `${counted(names.length, 'encounter')} down this tier: ${names.join(', ')}.`;
}

/** "1 raider · 4 waiting for approval · 0 below the rating floor". */
export function rosterCardParts(summary: {
  raiders: number;
  waiting_for_approval: number;
  below_rating_floor: number;
}): { raiders: string; waiting: string; belowFloor: string } {
  return {
    raiders: `${counted(summary.raiders, 'raider')} with gear on the site`,
    waiting: `${summary.waiting_for_approval} waiting for approval`,
    belowFloor: `${summary.below_rating_floor} below the rating floor`,
  };
}

/** "1 check failing", "3 checks failing". */
export function checksFailingLine(failing: number): string {
  return `${counted(failing, 'check')} failing`;
}

export const NO_PUBLIC_REPORTS = 'No public reports this week yet.';
export const NO_PUBLIC_RAID_NIGHTS = 'No public raid nights yet.';
