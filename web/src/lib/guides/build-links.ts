// web/src/lib/guides/build-links.ts
// The two links under a spec guide's embedded tree (spec 5): "Load this build" is a plain
// /planner?code= link (current-character.ts's own plannerHrefFor builds the identical shape
// for a 'code' pointer); "Sim this build" goes through lib/sim/url.ts's own state builder
// (that module's header comment: a `?code=` link there is "never built by hand"). Neither
// link writes the current-character pointer itself -- Planner.svelte's writePlannerPointer
// and sim's sources.ts already do that unconditionally on any `?code=` load, so this needs no
// session read of its own (Global Constraint 3 stays satisfied by having nothing to read).
import { SIM_TABS, tabHref } from '../sim/tabs';
import { defaultSimState, withSimState } from '../sim/url';

export function loadBuildHref(code: string): string {
  return `/planner?code=${encodeURIComponent(code)}`;
}

export function simBuildHref(code: string): string {
  return tabHref(SIM_TABS[0].href, withSimState(defaultSimState(), { code }));
}
