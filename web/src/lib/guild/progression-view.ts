// web/src/lib/guild/progression-view.ts
// The Progression tab's words (design/specs/2026-11-05-guild-week-one.md §3.4): the pre-raid empty
// state, and an encounter's kill line, printed only when both its figures exist. Pure, so the page
// can never print `NaN:NaN` for a field the API omitted.
import { nightDate } from '../dates';
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

/** `3:45`, from a kill's milliseconds. */
function killTime(ms: number): string {
  const totalSeconds = Math.round(ms / 1000);
  return `${Math.floor(totalSeconds / 60)}:${String(totalSeconds % 60).padStart(2, '0')}`;
}

/** `First kill Wed 16 Dec · best kill 3:45`; null unless both fields are really there (the API omits
 *  them with no kill, so `!== null` would let `undefined` through). */
export function firstKillLine(
  encounter: Pick<GuildProgressionEncounter, 'first_kill_at' | 'best_kill_ms'>,
  now: Date,
): string | null {
  const { first_kill_at: firstKill, best_kill_ms: bestKill } = encounter;
  if (typeof firstKill !== 'string' || typeof bestKill !== 'number') return null;
  return `First kill ${nightDate(new Date(firstKill), now)} · best kill ${killTime(bestKill)}`;
}
