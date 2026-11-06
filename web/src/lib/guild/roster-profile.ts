// web/src/lib/guild/roster-profile.ts
// Where a roster row's name leads: the member profile (web layout spec 2026-11-05 §7.2), for an
// officer any verified raider's, for a member only their own. Apart from roster-links.ts, whose
// hand-off helpers would weigh down the guild page's first chunk.
import { guildMemberHref } from '../characters';
import type { GuildRosterRow } from './api';

/**
 * A roster row's member profile (web layout spec 2026-11-05 §7.2): an officer opens anyone's by
 * account, a member only their own row's (as `me`); null when the row links nowhere.
 */
export function rosterProfileHref(
  row: Pick<GuildRosterRow, 'character_key' | 'account_key' | 'verified'>,
  viewer: { officer: boolean; myCharacterKey: string | null; guildPath: string },
): string | null {
  if (viewer.guildPath === '' || !row.verified) return null;
  if (row.character_key === viewer.myCharacterKey) return guildMemberHref(viewer.guildPath, 'me');
  if (!viewer.officer || row.account_key === undefined || !row.account_key.startsWith('u:')) return null;
  return guildMemberHref(viewer.guildPath, row.account_key);
}
