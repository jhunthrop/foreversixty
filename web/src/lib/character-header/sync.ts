// web/src/lib/character-header/sync.ts
// The header's sync line: "Addon · 12 minutes ago", and past the stale limit the words
// "Log out in the game to sync" in ember. The stale limit is the selector's own rule
// (character-selector/stale.ts) and the source word its own (rows.ts): one source of truth each.
import type { MeCharacter } from '../account/api';
import { sourceWordOf } from '../character-selector/rows';
import { isStale } from '../character-selector/stale';
import { relativeTime } from '../dates';
import { characterHeaderCopy } from './copy';

export interface SyncView {
  /** "Addon", "Battle.net" or "Pasted export"; empty when nothing says where it came from. */
  source: string;
  /** "12 minutes ago". */
  age: string;
  /** The words past the age, in ember; null while the sync is fresh. */
  staleTail: string | null;
}

/** Worn gear on file: an unknown is not a zero, so an empty export counts as none. */
export function hasGearOnFile(character: MeCharacter): boolean {
  const gear = character.build?.gear;
  return gear !== undefined && Object.keys(gear).length > 0;
}

/** Only an addon sync is asked to "log out in the game"; Battle.net refreshes itself. */
export function isStaleAddonSync(character: MeCharacter, now: Date): boolean {
  return character.build?.source === 'addon' && isStale(character.build, now);
}

/** The sync line, or null when the site holds no gear for the character (the line goes). */
export function syncView(character: MeCharacter, now: Date): SyncView | null {
  const build = character.build;
  if (build === undefined || !hasGearOnFile(character)) return null;
  return {
    source: sourceWordOf(character),
    age: relativeTime(new Date(build.captured_at), now),
    staleTail: isStaleAddonSync(character, now) ? characterHeaderCopy.staleTail : null,
  };
}

/** "Addon · 12 minutes ago"; just the age when nothing says where the build came from. */
export function syncText(sync: Pick<SyncView, 'source' | 'age'>): string {
  return sync.source === '' ? sync.age : `${sync.source} \u00b7 ${sync.age}`;
}
