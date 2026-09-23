// Every hand-off link this lane builds, in one place, in exactly the four URL forms the
// spec names (docs/superpowers/specs/2026-09-21-one-product-design.md section 2): a pasted
// or carried FS1 code into the planner or the simulator, a fight-sourced character (whole
// fight or one combatant) into the simulator, and an instance preselected on the drops
// tool. No other module in this lane's scope should compose one of these paths by hand —
// that is what lets the coordinator retarget every link at Lane B's
// `current-character.ts` (`plannerHrefFor`/`simHrefFor`) in this one file at merge time.
import { defaultSimState, simSearch, withSimState } from './sim/url';

export function plannerCodeHref(code: string): string {
  return `/planner?code=${encodeURIComponent(code)}`;
}

export function simCodeHref(code: string): string {
  return `/sim?code=${encodeURIComponent(code)}`;
}

/**
 * `<report>:<fight>` when `guid` is omitted (the existing whole-fight "Sim this fight"
 * link's own ref shape), extended to `<report>:<fight>:<guid>` when it is given — the
 * combatant GUID exactly as the report's roster carries it. Lane B's `fromLoggedFight`
 * reads the optional third part and falls back to today's first-dps rule when absent.
 */
export function simFightHref(reportId: string, fightIndex: number, guid?: string): string {
  const ref = guid === undefined ? `${reportId}:${fightIndex}` : `${reportId}:${fightIndex}:${guid}`;
  return `/sim?source=fight&ref=${encodeURIComponent(ref)}`;
}

export function simDropsHref(instanceSlug: string): string {
  return `/sim/drops?instance=${encodeURIComponent(instanceSlug)}`;
}

/** A stored character, by key (spec 2026-09-22 §3.4): loads through the sim's own
 *  `?source=armory&ref=` bootstrap (`store.svelte.ts`'s `bootstrapSource`), which resolves
 *  whichever source the API actually holds -- `'addon'` or `'blizzard'` -- with no extra
 *  fetch needed to build this href, unlike `simCodeHref`/`plannerCodeHref` which need the
 *  FS1 string itself. Used where a whole list of characters needs a working "Open in
 *  simulator" link without one `sim-input` fetch per row (`CharacterRowLink.svelte`,
 *  `LandingState.svelte`'s own row link already builds the identical URL by hand). */
export function simArmoryHref(characterKey: string): string {
  return `/sim${simSearch(withSimState(defaultSimState(), { source: 'armory', ref: characterKey }))}`;
}
