// web/src/lib/home/upgrades-layout.ts
// The height "Your upgrades" holds before its island mounts. The shell cannot know which state
// the island will render (a table, "every slot matches", no spec yet), so the island stores the
// height it settled at and the next visit's pre-paint script (index.astro) reserves exactly that.
// A first visit reserves the smallest real state, so an empty list never leaves a void.
import { writeReserve } from '../layout-reserve';

/** Read back by the inline script under the slot in index.astro. */
export const HOME_UPGRADES_RESERVE_KEY = 'fs.home.upgrades.reserve';

/** The custom property the script sets on the slot. */
export const HOME_UPGRADES_RESERVE_VAR = '--home-upgrades-reserve';

/** Eight full rows plus the head and foot stand under 800 px; above this a stored value is nonsense. */
export const HOME_UPGRADES_MAX_RESERVE_PX = 1600;

/** The smallest real state's card height (a one-line body under the Card head), measured at 390
 *  and 1440: the no-spec, no-gear and no-list states are all this tall. */
export const HOME_UPGRADES_FALLBACK_PX = 147;

/** Stores the settled card height for the next visit. */
export function storeUpgradesReserve(pixels: number): void {
  writeReserve(HOME_UPGRADES_RESERVE_KEY, pixels, HOME_UPGRADES_MAX_RESERVE_PX);
}
