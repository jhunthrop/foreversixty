// web/src/lib/planner/url.ts
// The planner's address mirrors the class and race on screen. The page reads `?class=`,
// `?race=` and `?code=` once on load (Planner.svelte's `fromQuery`); this is the other
// direction: after a class or race switch the query string is rewritten so a refresh, a
// bookmark or a copied address lands on the build the visitor was looking at, not the one
// the page opened on.

const CLASS = 'class';
const RACE = 'race';
const CODE = 'code';

/**
 * The query string the planner should carry for `classSlug` and `raceSlug`, or `null` when
 * the address already lands on the build on screen. `code` is the build the address loaded
 * (a `?code=`, or the current-character pointer a bare `/planner` restores): while its class
 * and race are still the ones on screen the address is left alone, so a refresh reloads that
 * build, talents and all. Once the visitor moves off it, class and race are written and any
 * `?code=` dropped, because a refresh would otherwise reload the code's build over the one
 * they chose. Every other parameter is left as it was.
 */
export function plannerSearchFor(
  search: string,
  classSlug: string,
  raceSlug: string,
  code: { classSlug: string; raceSlug: string } | null,
): string | null {
  const params = new URLSearchParams(search);
  if (code !== null && code.classSlug === classSlug && code.raceSlug === raceSlug) return null;

  const next = new URLSearchParams(search);
  next.delete(CODE);
  next.set(CLASS, classSlug);
  if (raceSlug === '') next.delete(RACE);
  else next.set(RACE, raceSlug);

  const rendered = next.toString();
  const current = params.toString();
  if (rendered === current) return null;
  return rendered === '' ? '' : `?${rendered}`;
}
