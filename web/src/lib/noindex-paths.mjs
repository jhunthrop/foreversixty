// web/src/lib/noindex-paths.mjs
// The pages a search engine should not index, in one list that both the layout (Base.astro writes
// `<meta name="robots" content="noindex">`) and the sitemap filter (astro.config.mjs) read. Each is
// personal or a step in a flow: signed out it is a sign-in prompt or a "checking" state, which
// Search Console reports as a soft 404, and signed in it is one person's own data.
//   /account           your account
//   /reports           your reports (each report is /reports/<id>, its own page)
//   /connect           the companion's Battle.net hand-off, read off a one-time code
//   /login             the sign-in step
//   /premium/checkout  the checkout step (/premium is the page that sells it)
// One more is a trust page, not a destination: /sources (how the source pills work) is linked from
// every footer and is mostly prose, so Search Console showed it soaking up impressions that belong
// to the guides and BiS pages (2026-11-11). It stays reachable from the footer and out of the index.
// Never list a page src/worker.ts serves as a shell for other URLs (guild-shell, character,
// rankings, report-shell, sim): a meta tag in the shell would reach every page it serves; the
// Worker sets x-robots-tag per URL there instead.

/** @type {readonly string[]} */
export const NOINDEX_PATHS = ['/account', '/reports', '/connect', '/login', '/premium/checkout', '/sources'];

/**
 * @param {string} path A site path, `/account`.
 * @returns {boolean}
 */
export function isNoindexPath(path) {
  return NOINDEX_PATHS.includes(path.replace(/\/+$/, '') || '/');
}
