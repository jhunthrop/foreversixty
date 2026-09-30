// web/src/lib/home/get-set-up.ts
// The home page's "Get set up" banner, once signed in (review round 1 fix item 3): a
// returning visitor is shown only the step(s) `me` proves are not done yet, never the full
// three-step pitch a first-time visitor sees. "Signed in" is always true once this runs --
// HomeGetSetUp.svelte only calls this after `me !== null` -- so the one real variable is the
// addon; the companion has no signal on this page at all, so its own item is always shown
// and never carries a done mark.
import type { Me, MeCharacter } from '../account/api';
import { homeGetSetUpCopy } from '../home-landing-copy';
import { relativeTime } from '../sim/sources';

/** True once any of the visitor's characters carries an addon-sourced build -- the same
 *  fact `MeCharacter.build.source` records for the account page's own build pill. */
export function addonLinked(me: Me): boolean {
  return me.characters.some((character) => character.build?.source === 'addon');
}

/** The most recent addon-sourced build's own `captured_at`, across every character --
 *  undefined when none has one (mirrors `addonLinked`'s own check, so a caller never reads
 *  this without `addonLinked` already being true). Ties are broken by array order, the same
 *  "first one wins" rule every other home surface applies to a multi-character fact. */
function latestAddonSync(characters: readonly MeCharacter[]): string | undefined {
  const timestamps = characters
    .filter((character) => character.build?.source === 'addon')
    .map((character) => character.build!.captured_at);
  return timestamps.sort((a, b) => new Date(b).getTime() - new Date(a).getTime())[0];
}

/** The signed-in banner's one line: the exact collapsed sentence once the addon is linked
 *  too (home rebuild spec §3.B.6: carrying the same relative sync time the hero's own sync
 *  line shows, so the two can never disagree), or the addon step alongside the ever-present
 *  companion step otherwise. */
export function getSetUpLine(me: Me, now: Date = new Date()): string {
  if (addonLinked(me)) {
    const capturedAt = latestAddonSync(me.characters);
    const syncedAgo = capturedAt === undefined ? '' : relativeTime(capturedAt, now);
    return homeGetSetUpCopy.bothDone(syncedAgo);
  }
  const { addonRemaining, companionRemaining } = homeGetSetUpCopy;
  return `${addonRemaining} · ${companionRemaining} →`;
}
