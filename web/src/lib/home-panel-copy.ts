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
import { scoreUnitWord, type SlotScoreUnit } from './bis/tank-view';
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
 * and §3.B.3's "Your upgrades" panel, now that `MeCharacter.build.gear`/`.talents` exist
 * (the API lane's own extension to the contract, landed alongside this lane) and
 * `lib/home/upgrades.ts`/`talent-delta.ts` compute the real comparison. Three states remain
 * genuinely unavailable, each named by its own real reason rather than a generic spinner
 * (tenet 8, never a fabricated figure):
 *
 *   - no spec at all (`character.spec` undefined) -- nothing to compare against;
 *   - a spec, but no addon gear/talent export yet (`build.gear`/`build.talents` undefined);
 *   - a spec and an export, but no published BiS file/band for it yet.
 */
export const homeHeroCardsCopy = {
  bestInSlotLabel: 'Best in slot',
  /** §4's "no spec yet" state, named once and shared by the Best in slot and Talents cards
   *  (both are blocked on the identical missing fact). */
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
  bestInSlotNoListYet: 'No best in slot list published yet for this spec and band.',
  /** §3.B.2: `N upgrades` zero-upgrades figure/line, the "don't hide a win" case. */
  bestInSlotFigure: (upgradeCount: number): string =>
    `${upgradeCount} upgrade${upgradeCount === 1 ? '' : 's'}`,
  bestInSlotAllMatchFigure: 'Best in slot',
  bestInSlotAllMatchLine: (bandLabel: string): string => `every slot matches the ${bandLabel} list`,
  /** Fix round 1 item A.1: `notSimCheckedCount` names how many upgrade slots this band's own
   *  figure excludes (a weapon slot with no sim-verified alternative, `lib/home/upgrades.ts`'s
   *  own `notSimChecked`) -- said here rather than silently dropped, so the total always adds
   *  up to a number a player can audit against the table below it. */
  bestInSlotUpgradesLine: (
    bandLabel: string,
    totalGain: number,
    notSimCheckedCount: number,
    unit: SlotScoreUnit = 'dps',
  ): string => {
    const base = `in your band, ${bandLabel} · +${totalGain.toFixed(1)} ${scoreUnitWord(unit)} together`;
    if (notSimCheckedCount === 0) return base;
    const clause = notSimCheckedCount === 1 ? 'one slot' : `${notSimCheckedCount} slots`;
    return `${base}, ${clause} not sim-checked`;
  },
  talentsLabel: 'Talents',
  talentsNotAvailable:
    'Not available yet: no talent export for this character. Open the addon once to send it.',
  talentsNoListYet: 'No talent build published yet for this spec and band.',
  talentsOptimizedFigure: 'Optimized',
  talentsUnoptimizedFigure: 'Unoptimized',
  /** Fix round 1 item A.3: names whose points the denominator counts (the band's own build),
   *  not just a bare number next to an unrelated band label. */
  talentsLine: (pointsDiffer: number, bandTalentPoints: number, bandLabel: string): string =>
    `${pointsDiffer} of the ${bandTalentPoints} points in the ${bandLabel} build differ · compare in the planner`,
  simulatorLabel: 'Simulator',
  /** §3.B.2's "Y.Y in band best in slot" clause, appended to the visitor's own saved-sim
   *  line only once the band's own `set_dps` is known -- never a pairing the band file does
   *  not actually publish. */
  simulatorBandSuffix: (bandSetDps: number): string => `${bandSetDps.toFixed(1)} in band best in slot`,
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
  youWearThis: 'you wear this',
  noItemEquipped: 'No item equipped',
  unknownItem: 'Unknown item',
  /** `Full list for Marksmanship 20 to 29` -- the panel's own header aside link. */
  fullListLink: (specName: string, bandLabel: string): string => `Full list for ${specName} ${bandLabel}`,
  alreadyBestInSlotPrefix: 'Already best in slot:',
  moreSlots: (count: number): string => `${count} more slot${count === 1 ? '' : 's'}`,
  everySlotMatches: (bandLabel: string): string =>
    `Every slot in the ${bandLabel} list already matches what you wear.`,
} as const;
