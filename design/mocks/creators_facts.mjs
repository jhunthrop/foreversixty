// The public /creators page's words and figures, for the creators mock (design/mocks/ws_creators.py):
//   node design/mocks/creators_facts.mjs [out.json]  ->  design/mocks/data/creators-page-facts.json (or out.json)
// Reads web/src/lib/creators/copy.ts itself (bundled with esbuild from web/node_modules), never retyped: the program's
// sections, each section's figures, the head's answer, the apply section, the Apply link, the contents label and the
// worked examples (lib/creators/model.ts earningsRows at today's prices).
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/creators-page-facts.json", import.meta.url);

const entry = `
import { CREATORS_APPLY_EMAIL, CREATORS_APPLY_HREF, PROGRAM_TERMS, programCopy } from ${JSON.stringify(WEB + "src/lib/creators/copy.ts")};
import { earningsRows } from ${JSON.stringify(WEB + "src/lib/creators/model.ts")};
import { founderOpen } from ${JSON.stringify(WEB + "src/lib/billing/plans.ts")};
// The worked examples at today's prices, as the page's build computes them.
export const facts = { program: programCopy, terms: PROGRAM_TERMS, applyEmail: CREATORS_APPLY_EMAIL, applyHref: CREATORS_APPLY_HREF,
  examples: earningsRows(founderOpen(new Date())) };
`;
const outfile = join(mkdtempSync(join(tmpdir(), "creators-facts-")), "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm", outfile, logLevel: "error" });
const { facts } = await import(pathToFileURL(outfile).href);
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
