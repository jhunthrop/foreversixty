// web/src/lib/current-character-bar.ts
// Pure logic for the spine bar's four doors and the class it resolves to display
// (CurrentCharacterBar.svelte's `spine` mode, spec 2026-09-25 section 4.1). Kept apart from
// the component so the precedence rules -- pointer wins, main is the fallback; Plan/Sim carry
// a restorable pointer when one exists, Rankings is pre-filtered by class and ruleset -- are
// unit-tested without a DOM.
import type { MeCharacter } from './account/api';
import { plannerHrefFor, simHrefFor, type CurrentCharacter } from './current-character';
import { defaultRankingsState, rankingsSearch } from './rankings/url';
import { classDisplayName } from './sim/spec-label';

/** The class the spine displays and links every door with: the resolved pointer character's
 *  class first, then the raw pointer's own `classSlug` (every source carries one, even a
 *  paste with no matching armory record), then the signed-in main's class, else null (no
 *  character known at all, or a main with no class on file yet from Battle.net). */
export function resolveSpineClassSlug(
  pointerCharacter: MeCharacter | null,
  current: CurrentCharacter | null,
  main: MeCharacter | null,
): string | null {
  if (pointerCharacter?.class !== undefined) return pointerCharacter.class.toLowerCase();
  if (current !== null) return current.classSlug;
  return main?.class?.toLowerCase() ?? null;
}

export interface SpineDoor {
  readonly id: 'plan' | 'sim' | 'logs' | 'rankings';
  readonly label: string;
  /** What a phone shows in place of `label` when the full label would not fit six doors
   *  on one 390px row: "Rankings" for "Rankings for Warrior". Absent when they are the same. */
  readonly shortLabel?: string;
  readonly href: string;
  readonly testid: string;
}

/**
 * The four door links (spec 4.1: "Plan, Sim, Logs, Rankings for {class}"). Plan and Sim
 * carry the pointer when one exists -- the exact same hrefs `plannerHrefFor`/`simHrefFor`
 * already build for the chip -- and fall back to a bare `?class=`/bare `/sim` link when
 * there is none (the main-fallback case; the destination page's own bootstrap then opens on
 * the signed-in main, per the planner and sim doors' own spec bullets). Logs never carries
 * the character in its URL: the logs page reads the same pointer/session directly. Rankings
 * is omitted entirely when no class is known at all -- a filter with nothing to filter by is
 * not a door.
 */
export function spineDoorsFor(
  classSlug: string | null,
  current: CurrentCharacter | null,
  region: string | null,
  ruleset: string | null,
): SpineDoor[] {
  if (classSlug === null) {
    return [
      { id: 'plan', label: 'Plan', href: '/planner', testid: 'current-character-bar-plan' },
      { id: 'sim', label: 'Sim', href: '/sim', testid: 'current-character-bar-sim' },
      { id: 'logs', label: 'Logs', href: '/logs', testid: 'current-character-bar-logs' },
    ];
  }
  const plan = current === null ? `/planner?class=${encodeURIComponent(classSlug)}` : plannerHrefFor(current);
  const sim = current === null ? '/sim' : simHrefFor(current);
  const rankingsQuery = rankingsSearch({
    ...defaultRankingsState(),
    class: classSlug,
    ruleset: region !== null && ruleset !== null ? ruleset : '',
  });
  return [
    { id: 'plan', label: 'Plan', href: plan, testid: 'current-character-bar-plan' },
    { id: 'sim', label: 'Sim', href: sim, testid: 'current-character-bar-sim' },
    { id: 'logs', label: 'Logs', href: '/logs', testid: 'current-character-bar-logs' },
    {
      id: 'rankings',
      label: `Rankings for ${classDisplayName(classSlug)}`,
      shortLabel: 'Rankings',
      href: `/rankings${rankingsQuery}`,
      testid: 'current-character-bar-rankings',
    },
  ];
}
