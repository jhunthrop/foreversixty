// web/src/lib/planner/config.ts
// Build-time constants for the planner. PUBLIC_API_BASE_URL is inlined by Vite into both
// the Astro build and the standalone island bundle, so the island works identically when
// the API serves it from /b/:id.

/** Where POST /v1/builds and /v1/subscribe live. */
export const API_BASE_URL: string = import.meta.env.PUBLIC_API_BASE_URL ?? 'https://api.foreversixty.gg';

/** The class the planner opens on when the query string does not say otherwise. */
export const DEFAULT_CLASS_SLUG = 'warrior';

/** The build id of the pre-beta data set, whose trees came from Wowhead, not a client. */
export const PREBETA_BUILD = 'forever-prebeta';

/**
 * What the planner says above the trees about where they came from. It names the
 * build rather than the expansion: the site serves several builds at once (a shared
 * link renders against the build it was saved on), so "Classic Era trees" was both
 * wrong and unanswerable once the beta client's trees shipped.
 */
export function treeSourceNotice(build: string): string {
  return build === PREBETA_BUILD
    ? 'Talent trees from a pre-beta Wowhead snapshot; the client’s own trees replace them at the beta.'
    : `Talent trees read from the game client, build ${build}.`;
}
