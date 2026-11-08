// The copy and the QR the app beta mocks render (ws_app_beta.py): foreversixty.gg/app and the
// owner's /admin/testers.
//   node design/mocks/app_beta_facts.mjs [out.json]  ->  design/mocks/data/app-beta-facts.json (or out.json)
// Bundles the TypeScript modules themselves with esbuild (web/node_modules), never retyped:
//  - /app: lib/app-beta/copy.ts, and lib/app-beta/qr.ts for the same SVG the page inlines;
//  - /admin/testers: lib/app-beta/admin.ts over the testers tests/e2e/parity-app-beta.spec.ts serves.
import { createRequire } from "node:module";
import { mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/app-beta-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

const entry = `
import { appBetaCopy, PLATFORM_NAMES } from ${src("src/lib/app-beta/copy.ts")};
import { appQrSvg } from ${src("src/lib/app-beta/qr.ts")};
import { adminTestersCopy, countsByPlatform, PLATFORM_ORDER, testerRows } from ${src("src/lib/app-beta/admin.ts")};
import testers from ${src("tests/e2e/fixtures/app-beta/testers.json")};

const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const c = adminTestersCopy;
const counts = countsByPlatform(testers);

export default {
  app: {
    copy: strip(appBetaCopy),
    platforms: PLATFORM_NAMES,
    done: appBetaCopy.created('ios'),
    exists: appBetaCopy.exists('ios'),
    qr: await appQrSvg(),
  },
  admin: {
    copy: strip(c),
    meta: c.tableMeta(testers.length),
    buttons: PLATFORM_ORDER.map((p) => c.copyEmails(p, counts[p])),
    rows: testerRows(testers),
  },
};`;
// qrcode stays external (it requires node built-ins at load), so the bundle sits in web/ to resolve it.
const dir = mkdtempSync(join(WEB, "node_modules", ".app-beta-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", external: ["qrcode"], define: { "import.meta.env": "{}" }, loader: { ".json": "json", ".svelte": "empty" } });
const facts = (await import(pathToFileURL(outfile).href)).default;
rmSync(dir, { recursive: true, force: true });
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
