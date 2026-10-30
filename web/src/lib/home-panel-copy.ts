// web/src/lib/home-panel-copy.ts
// The home page's account-aware hero (spec 2026-09-22 §3.2, restyled into the hero by
// spec 2026-09-23 §2, rebuilt by the home rebuild spec 2026-09-30 §3.B): reference, not
// pitch -- one sentence, one button, signed out; signed in, the shared character header's
// answer words (design/specs/2026-10-11-character-header.md) and the two next-action cards
// under it (no switcher: the header selector is the one way to change character).

/** The signed-out block's `id` in `index.astro`: `HomeAccountPanel.svelte` reaches outside
 *  its own root to find it (same cross-island DOM-reach pattern as `SIM_TAB_ATTR` /
 *  `syncTabHrefs` in `lib/sim/tabs.ts`) and mark it `inert`/`aria-hidden` once signed in, so
 *  the duplicate "Sign in with Battle.net" link the grid-overlay CLS trick leaves behind it
 *  is never focusable or announced. One id, exported once, so the producer (index.astro)
 *  and the consumer (HomeAccountPanel.svelte) can never drift apart. */
export const HOME_SIGNED_OUT_ID = 'home-signed-out';

/** The attribute `HomeAccountPanel.svelte` puts on <html> when the pre-paint session hint
 *  turned out wrong, so global.css's hint-only collapse of the signed-out hero's cells
 *  (`html[data-session='1']:not([data-home-signed-out])`) stops applying. One name, exported
 *  once, shared with that stylesheet by its comment. */
export const HOME_SIGNED_OUT_ATTR = 'data-home-signed-out';

export const homePanelCopy = {
  signInButton: 'Sign in with Battle.net',
  /** The signed-in Simulator card's own empty/loading fallback (§3.B.2) -- reused verbatim
   *  from the pre-rebuild next-steps grid, since the spec's own Empty-state row for this
   *  card names these exact two strings. */
  noSimYet: 'No sim yet.',
  runAction: 'Run a sim',
  /** The shared character header's words on the home page (spec 2026-10-11 section 4). */
  headerAnswerLabel: 'Upgrades for your level range',
  headerNote: (bandLabel: string): string => `Best in slot for ${bandLabel}, together`,
  /** Said, not silently dropped: a slot with no sim-verified alternative is left out of the gain. */
  headerNotSimChecked: (count: number): string =>
    `${count === 1 ? 'one slot' : `${count} slots`} not sim-checked`,
  headerAllMatch: (bandLabel: string): string => `Every slot matches the ${bandLabel} list.`,
  headerNoGear: 'Add the addon to see which slots you can upgrade.',
  headerFailed: 'Your upgrades did not load.',
  headerAction: 'See your upgrades →',
} as const;

/**
 * Home rebuild spec §3.B.2's next-action cards (Talents / Simulator -- Best in slot is the
 * shared header's answer and action now) and §3.B.3's "Your upgrades" panel, now that
 * `MeCharacter.build.gear`/`.talents` exist (the API lane's own extension to the contract,
 * landed alongside this lane) and
 * `lib/home/upgrades.ts`/`talent-delta.ts` compute the real comparison. Three states remain
 * genuinely unavailable, each named by its own real reason rather than a generic spinner
 * (tenet 8, never a fabricated figure):
 *
 *   - no spec at all (`character.spec` undefined) -- nothing to compare against;
 *   - a spec, but no addon gear/talent export yet (`build.gear`/`build.talents` undefined);
 *   - a spec and an export, but no published BiS file/band for it yet.
 */
export const homeHeroCardsCopy = {
  /** §4's "no spec yet" state, named once and shared by the header, the Talents card and the
   *  upgrades panel (all blocked on the identical missing fact). */
  pickASpec: 'Pick a spec',
  pickASpecLine: 'Set a spec in the planner to see this.',
  /** Reworded (home rebuild spec §4, "Empty states stay honest") from the pre-contract
   *  "the addon does not send worn gear" -- now that the addon CAN send it, the honest gap
   *  is this one character's own export, not a site-wide limitation. Best in slot card and
   *  the "Your upgrades" panel name the identical real gap, so they read the same sentence
   *  rather than two different ways of saying it. */
  bestInSlotNotAvailable:
    'Not available yet: no gear export for this character. Open the addon once to send it.',
  /** A spec and a gear export both exist, but this build has no published BiS list for the
   *  spec/band pair yet -- distinct from `bestInSlotNotAvailable` (that names a gap in THIS
   *  character's own data; this one names a gap in the site's own published data). */
  bestInSlotNoListYet: 'No best in slot list published yet for this spec and level range.',
  talentsLabel: 'Talents',
  talentsNotAvailable:
    'Not available yet: no talent export for this character. Open the addon once to send it.',
  talentsNoListYet: 'No talent build published yet for this spec and level range.',
  talentsOptimizedFigure: 'Optimized',
  talentsUnoptimizedFigure: 'Unoptimized',
  /** Fix round 1 item A.3: names whose points the denominator counts (the band's own build),
   *  not just a bare number next to an unrelated band label. */
  talentsLine: (pointsDiffer: number, bandTalentPoints: number, bandLabel: string): string =>
    `${pointsDiffer} of the ${bandTalentPoints} points in the ${bandLabel} build differ · compare in the planner`,
  simulatorLabel: 'Simulator',
  /** The visitor's own latest saved sim, the unit on the figure: "26 DPS now". Shown only
   *  when that sim ran the band entry's own setup (`lib/home/sim-card.ts`). */
  simulatorNowFigure: (dps: string): string => `${dps} DPS now`,
  /** The band's own `set_dps`. Read after the figure it is comparable with:
   *  "26 DPS now · 37.5 DPS at band best in slot". */
  simulatorBandLine: (bandSetDps: number): string =>
    `${bandSetDps.toFixed(1)} DPS at best in slot for your level range`,
  /** With no comparable saved sim the band figure stands alone, as the card's figure. */
  simulatorBandOnlyFigure: (bandSetDps: number): string => `${bandSetDps.toFixed(1)} DPS`,
  simulatorBandOnlyLine: 'at best in slot for your level range',
} as const;

/**
 * §3.B.3's "Your upgrades" panel -- the full worn-vs-BiS table, once `lib/home/upgrades.ts`
 * has a real comparison to show. Shares `homeHeroCardsCopy`'s own not-yet-available/no-spec/
 * no-list sentences (the identical real gaps), so only the table's own rows/footer need
 * their own strings here.
 */
export const homeUpgradesCopy = {
  heading: 'Your upgrades',
  slotHeader: 'Slot',
  youWearHeader: 'You wear',
  bestInSlotHeader: 'Best in slot',
  gainHeader: 'Gain',
  noItemEquipped: 'No item equipped',
  unknownItem: 'Unknown item',
  /** `Full list for Marksmanship 20 to 29` -- the panel's own header aside link. */
  fullListLink: (specName: string, bandLabel: string): string => `Full list for ${specName} ${bandLabel}`,
  alreadyBestInSlotPrefix: 'Already best in slot:',
  moreSlots: (count: number): string => `${count} more slot${count === 1 ? '' : 's'}`,
  everySlotMatches: (bandLabel: string): string =>
    `Every slot in the ${bandLabel} list already matches what you wear.`,
} as const;
