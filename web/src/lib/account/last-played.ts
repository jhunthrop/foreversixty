// web/src/lib/account/last-played.ts
// The account character played most recently: the newest `build.captured_at`, which an addon
// sync stamps with the capture and a Battle.net import with the character's last login
// (api/internal/bnetimport/build.go). A character the site holds no build for has no play time
// to judge by and never wins.
import type { MeCharacter } from './api';

function playedAt(character: MeCharacter): string {
  return character.build?.captured_at ?? '';
}

function tieRank(character: MeCharacter, preferKeys: readonly (string | undefined)[]): number {
  const at = preferKeys.indexOf(character.key);
  return at === -1 ? preferKeys.length : at;
}

/**
 * The character with the newest capture, or null when none has one. A tie goes to the first of
 * `preferKeys` among the tied (the restored character, then the main), then the higher level,
 * then list order.
 */
export function lastPlayedCharacter(
  characters: readonly MeCharacter[],
  preferKeys: readonly (string | undefined)[],
): MeCharacter | null {
  const played = characters.filter((character) => playedAt(character) !== '');
  if (played.length === 0) return null;
  const newest = played.reduce((best, next) => (playedAt(next) > playedAt(best) ? next : best));
  const tied = played.filter((character) => playedAt(character) === playedAt(newest));
  return tied.reduce((best, next) => {
    const byPreference = tieRank(next, preferKeys) - tieRank(best, preferKeys);
    if (byPreference !== 0) return byPreference < 0 ? next : best;
    return (next.level ?? 0) > (best.level ?? 0) ? next : best;
  });
}
