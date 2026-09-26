// web/src/lib/home/get-set-up.ts
// The home page's "Get set up" banner, once signed in (review round 1 fix item 3): a
// returning visitor is shown only the step(s) `me` proves are not done yet, never the full
// three-step pitch a first-time visitor sees. "Signed in" is always true once this runs --
// HomeGetSetUp.svelte only calls this after `me !== null` -- so the one real variable is the
// addon; the companion has no signal on this page at all, so its own item is always shown
// and never carries a done mark.
import type { Me } from '../account/api';
import { homeGetSetUpCopy } from '../home-landing-copy';

/** True once any of the visitor's characters carries an addon-sourced build -- the same
 *  fact `MeCharacter.build.source` records for the account page's own build pill. */
export function addonLinked(me: Me): boolean {
  return me.characters.some((character) => character.build?.source === 'addon');
}

/** The signed-in banner's one line: the exact collapsed sentence once the addon is linked
 *  too, or the addon step alongside the ever-present companion step otherwise. */
export function getSetUpLine(me: Me): string {
  if (addonLinked(me)) return homeGetSetUpCopy.bothDone;
  const { addonRemaining, companionRemaining } = homeGetSetUpCopy;
  return `${addonRemaining} · ${companionRemaining} →`;
}
