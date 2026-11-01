// web/src/lib/character-header/gear-hint.ts
// Which shape the character header's loading state reserves. A character with gear on file draws
// a sync line under its descriptor and the page's figures; one without draws neither (the answer
// is the "add the addon" sentence), so the ready header is shorter. The header records which one
// it last drew here, and Base.astro's pre-paint script turns a remembered '1' into
// `<html data-ch-gear="1">`, so the skeleton takes the geared box only when the header that
// replaces it will be geared. With nothing remembered the skeleton takes the gearless box, the
// shape every character has until its first sync.
import { hasGearOnFile } from './sync';
import type { MeCharacter } from '../account/api';

/** The localStorage key; Base.astro's inline script reads it under the same name. */
export const HEADER_GEAR_HINT_KEY = 'fs.header.gear';

/** The value stored for a geared header; anything else reads as gearless. */
export const HEADER_GEARED = '1';
const HEADER_GEARLESS = '0';

/** Records the shape of the header just drawn for `character`, for the next page's skeleton. */
export function rememberHeaderGear(character: MeCharacter, storage: Pick<Storage, 'setItem'>): void {
  try {
    storage.setItem(HEADER_GEAR_HINT_KEY, hasGearOnFile(character) ? HEADER_GEARED : HEADER_GEARLESS);
  } catch {
    /* storage unavailable: the next skeleton takes the gearless box */
  }
}
