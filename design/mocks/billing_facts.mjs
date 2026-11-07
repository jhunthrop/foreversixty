// The copy the billing-lane reference mocks render (ws_billing.py): the
// account's Discord paste form and the owner's /admin/links page.
//   node design/mocks/billing_facts.mjs [out.json]  ->  design/mocks/data/billing-facts.json (or out.json)
// Bundles the TypeScript modules themselves with esbuild (web/node_modules), never retyped, and none that reads
// src/data/generated (the CI parity job's sparse checkout has no generated data):
//  - Discord paste form: lib/account/notifications-copy.ts;
//  - /admin/links: lib/golinks/admin-copy.ts and funnel.ts over the links tests/e2e/admin-links.spec.ts serves.
// The clock is pinned to 2026-11-15T12:00Z (inside the founder window) and the zone to UTC.
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/billing-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

const entry = `
import { notificationsCopy } from ${src("src/lib/account/notifications-copy.ts")};
import { adminLinksCopy } from ${src("src/lib/golinks/admin-copy.ts")};
import { countLabel, funnelTotals, shareOf, sortedByClicks } from ${src("src/lib/golinks/funnel.ts")};

const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const link = (slug, target, label, f) => ({ slug, target, label, clicks: f[0], bot_clicks: f[1], sign_ins: f[2], linked: f[3], premium: f[4] });
const LINKS = [link('reddit', '/', 'Reddit', [40, 6, 10, 5, 1]), link('reddit-sim', '/sim', 'Reddit: sim post', [120, 2, 12, 9, 3]), link('wago', '/', 'Wago', [0, 0, 0, 0, 0])];
const c = adminLinksCopy;
const note = (part, whole, basis) => (whole > 0 ? shareOf(part, whole) + ' ' + basis : shareOf(part, whole));
const steps = (f) => [
  { label: c.columns.clicks, value: countLabel(f.clicks), note: f.clicks === 0 && f.bot_clicks === 0 ? c.noClicks : c.bots(f.bot_clicks) },
  { label: c.columns.signIns, value: countLabel(f.sign_ins), note: note(f.sign_ins, f.clicks, c.shareOf.signIns) },
  { label: c.columns.linked, value: countLabel(f.linked), note: note(f.linked, f.sign_ins, c.shareOf.linked) },
  { label: c.columns.premium, value: countLabel(f.premium), note: note(f.premium, f.linked, c.shareOf.premium) },
];
const sorted = sortedByClicks(LINKS);

export default {
  discord: strip(notificationsCopy),
  admin: {
    copy: strip(adminLinksCopy), meta: c.tableMeta(sorted.length),
    rows: sorted.map((r) => ({ slug: r.slug, target: r.target, label: r.label, steps: steps(r) })), totals: steps(funnelTotals(LINKS)),
  },
};`;
const dir = mkdtempSync(join(tmpdir(), "billing-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.env": "{}" }, loader: { ".json": "json", ".svelte": "empty" } });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
