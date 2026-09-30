// web/lighthouserc.push.cjs
//
// Push/PR Lighthouse pass: audits the smallest set of URLs that still exercises every
// distinct rule in lighthouserc.json's `ci.assert.assertMatrix`, instead of all 16. The full
// list keeps running, unchanged, on the nightly schedule (`npm run lhci`, no --config) --
// see .github/workflows/web.yml's `lhci` vs `lhci-nightly` jobs.
//
// lighthouserc.json has 7 assertMatrix entries and @lhci/utils applies every entry whose
// `matchingUrlPattern` matches a collected URL (web/README.md's Lighthouse section documents
// this: "every collected URL matches exactly one"). The 7 representative URLs below were
// picked by running each of lighthouserc.json's 16 URLs against its own assertMatrix and
// keeping one URL per distinct pattern that matched (verified 2026-09-30; re-run the same
// check below if either file changes):
//
//   node -e '
//     const c = require("./lighthouserc.json");
//     for (const u of c.ci.collect.url) {
//       const m = c.ci.assert.assertMatrix.filter(x => new RegExp(x.matchingUrlPattern).test(u));
//       console.log(u, "->", m.map(x => x.matchingUrlPattern).join(" | "));
//     }
//   '
//
// Dropped from the push list because another kept URL already carries their identical
// assertMatrix entry: guides/warrior/fury.html and bis.html (guides.html's default-pattern
// entry), setup.html and logs.html (planner.html's entry), and the five other
// sim/*.html pages (sim.html's entry). Nothing here drops a budget -- every one of the 7
// rules is still asserted on every push; the other 9 URLs only add page-specific coverage of
// a rule already being checked, which is what makes them fit for nightly instead.
//
// This file re-exports the base config's ci.collect/ci.assert/ci.upload wholesale -- the
// assertMatrix and its budgets are never duplicated, so changing a budget still means
// changing it in exactly one place: lighthouserc.json.
//
// LHCI_PUSH_SHARD (1-3, set by .github/workflows/web.yml's `lhci` matrix) narrows PUSH_URLS
// further, to one of PUSH_GROUPS below, so the three shards run in parallel instead of one
// job auditing all 7: same numberOfRuns, same assertions, no URL or assertion dropped --
// PUSH_GROUPS.flat() is exactly PUSH_URLS. Unset (a plain `npm run lhci:push`) audits all 7,
// same as before sharding.
// @lhci/cli's own config loader require()s `.js`/`.cjs` config files (web/README.md's
// Lighthouse budgets section); this file has to be requireable the same way, so it stays
// CommonJS rather than the ESM `import` the rest of web/ uses.
// eslint-disable-next-line @typescript-eslint/no-require-imports
const base = require('./lighthouserc.json');

// Grouped, not a flat list, so a shard is "one of these arrays" rather than a slice index
// that silently drifts if a URL is added or removed above it. Balanced 3/2/2 by page count,
// the closest three-way split of 7: each group's total Lighthouse cost (numberOfRuns x URLs)
// is otherwise about the same, since every URL here carries the same numberOfRuns: 3.
const PUSH_GROUPS = [
  [
    'http://localhost/index.html', // index pattern
    'http://localhost/guides.html', // default (catch-all) pattern
    'http://localhost/bis/hunter/marksmanship.html', // bis/<class>/<spec> pattern
  ],
  [
    'http://localhost/planner.html', // planner|logs|setup pattern
    'http://localhost/reports/fixture2abcd.html', // reports/* pattern
  ],
  [
    'http://localhost/sim.html', // sim(/specs|/gear|/talents|/drops|/weights)? pattern
    'http://localhost/sim/simfixtureab.html', // sim/<12-char fixture slug> pattern
  ],
];

const PUSH_URLS = PUSH_GROUPS.flat();

function urlsForShard(shard) {
  if (shard === undefined) return PUSH_URLS;
  const index = Number(shard) - 1;
  const group = PUSH_GROUPS[index];
  if (group === undefined) {
    throw new Error(`LHCI_PUSH_SHARD must be 1-${PUSH_GROUPS.length}; got '${shard}'`);
  }
  return group;
}

module.exports = {
  ...base,
  ci: {
    ...base.ci,
    collect: { ...base.ci.collect, url: urlsForShard(process.env.LHCI_PUSH_SHARD) },
  },
};
