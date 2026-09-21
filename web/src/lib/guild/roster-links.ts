// web/src/lib/guild/roster-links.ts
// The guild roster's one hand-off helper (spec section 4.1): prefers current-character.ts's
// plannerHrefFor/simHrefFor, falling back to handoff-links.ts's plainer
// plannerCodeHref/simCodeHref if the richer module is ever unavailable. Both exist on main
// today and, for a 'code'-sourced character at the default tab, produce identical strings
// (confirmed by this module's own test) -- the fallback exists so this is the one place
// that changes if that ever stops being true, per CharacterHandoffLinks.svelte's own
// compatibility rule for /account and /character/<key>. No other file in this lane builds
// one of these hrefs by hand.
import { plannerHrefFor, simHrefFor, type CurrentCharacter } from '../current-character';
import { plannerCodeHref, simCodeHref } from '../handoff-links';

function syntheticCurrent(code: string): CurrentCharacter {
  // Only `.source` and `.ref` are read by plannerHrefFor/simHrefFor's 'code' branch; the
  // other three fields exist only to satisfy CurrentCharacter's shape (they describe a
  // browser's own remembered pointer, which a guild roster row is not).
  return { source: 'code', ref: code, label: '', classSlug: '', savedAt: '' };
}

export function rosterPlannerHref(code: string): string {
  try {
    return plannerHrefFor(syntheticCurrent(code));
  } catch {
    return plannerCodeHref(code);
  }
}

export function rosterSimHref(code: string): string {
  try {
    return simHrefFor(syntheticCurrent(code));
  } catch {
    return simCodeHref(code);
  }
}
