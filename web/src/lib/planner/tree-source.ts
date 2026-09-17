// web/src/lib/planner/tree-source.ts
// Pure wording for where the planner's talent trees came from. Split out of config.ts so
// this logic carries no `import.meta.env` dependency of its own: config.ts's API_BASE_URL
// reads import.meta.env.PUBLIC_API_BASE_URL at module scope, which is undefined outside
// Vite (Playwright's plain Node/TS loader, for one), so anything sharing a module with it
// cannot be imported there either. tests/e2e/planner.spec.ts and src/pages/_planner.test.ts
// both import treeSourceNotice directly from here for that reason.

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
