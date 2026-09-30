// web/src/lib/home-panel-copy.ts
// The home page's account-aware hero (spec 2026-09-22 §3.2, restyled into the hero by
// spec 2026-09-23 §2, rebuilt by the home rebuild spec 2026-09-30 §3.B): reference, not
// pitch -- one sentence, one button, signed out; the hero character, its descriptor, the
// three next-action cards and the Switch character panel, signed in.

/** The signed-out block's `id` in `index.astro`: `HomeAccountPanel.svelte` reaches outside
 *  its own root to find it (same cross-island DOM-reach pattern as `SIM_TAB_ATTR` /
 *  `syncTabHrefs` in `lib/sim/tabs.ts`) and mark it `inert`/`aria-hidden` once signed in, so
 *  the duplicate "Sign in with Battle.net" link the grid-overlay CLS trick leaves behind it
 *  is never focusable or announced. One id, exported once, so the producer (index.astro)
 *  and the consumer (HomeAccountPanel.svelte) can never drift apart. */
export const HOME_SIGNED_OUT_ID = 'home-signed-out';

export const homePanelCopy = {
  signInButton: 'Sign in with Battle.net',
  /** The signed-in Simulator card's own empty/loading fallback (§3.B.2) -- reused verbatim
   *  from the pre-rebuild next-steps grid, since the spec's own Empty-state row for this
   *  card names these exact two strings. */
  noSimYet: 'No sim yet.',
  runAction: 'Run',
  /** Home rebuild spec §3.B.1's signed-in eyebrow label (the mock's own "Sample" pill next
   *  to it is NOT reproduced here: that pill documents the mock's own stand-in fixture
   *  data, and this hero renders the visitor's real `/v1/me` character -- labelling real
   *  data "Sample" would be the exact fabrication tenet 8 rules out). */
  yourCharacterEyebrow: 'Your character',
  /** §3.B.1's sync line tail, after the mono relative-time figure: "{time} from the addon
   *  · gear and talents in sync" -- the "bags in sync" clause the mock shows is omitted,
   *  since no field on `MeCharacter` backs it yet. */
  syncedFromAddon: 'from the addon · gear and talents in sync',
  /** §4's stale-sync row (> 24h): one added line, no alarm colour. */
  reopenAddonToRefresh: 'Reopen the addon to refresh your gear.',
  /** §3.B.4's "Switch character" panel header. */
  switchCharacterLabel: 'Switch character',
  addOneCharacter: 'Add one',
} as const;

/**
 * Home rebuild spec §3.B.2's three next-action cards (Best in slot / Talents / Simulator)
 * and §3.B.3's "Your upgrades" panel. Best in slot, Your upgrades and Talents each name a
 * real, live comparison the site cannot compute yet -- no source joins a signed-in
 * character's actually-worn gear against a BiS list, and no function compares spent talents
 * against a band's recommended build -- so rather than a skeleton that shimmers forever
 * with nothing ever arriving behind it (which tells a screen-reader user "Loading" for a
 * fetch that will never resolve), these show one honest, settled line instead: labelled as
 * not yet available and naming the real reason, never a fabricated figure (tenet 8). Review
 * round 1's exact wording -- one sentence, not a label plus a second line. The Simulator
 * card carries no such gap (the visitor's own saved sim already exists, `lib/home/
 * next-steps.ts`'s `simCardLine`), so it alone keeps a real loading/ready/empty state.
 */
export const homeHeroCardsCopy = {
  bestInSlotLabel: 'Best in slot',
  /** Best in slot card and the "Your upgrades" panel name the identical real gap, so they
   *  read the same sentence rather than two different ways of saying it. */
  bestInSlotNotAvailable: 'Not available yet: the addon does not send worn gear.',
  talentsLabel: 'Talents',
  talentsNotAvailable: 'Not available yet: the addon does not send talents.',
  simulatorLabel: 'Simulator',
} as const;
