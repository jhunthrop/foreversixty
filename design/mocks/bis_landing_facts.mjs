// The data and copy the /bis landing mocks render (design/specs/2026-11-07-bis-landing.md):
//   node design/mocks/bis_landing_facts.mjs [out.json]  ->  design/mocks/data/bis-landing-facts.json (or out.json)
// Runs the landing's own models, bundled with esbuild from web/node_modules, never retyped:
//  - the page: loadLandingData (lib/bis/landing-build.ts: the classes and specs in game order, the band captions, the
//    raid label, the newest generated_at, the config), dayMonth for the Updated stamp, landingCopy;
//  - the answer, per account fixture: planAnswer over landingAnswerSpec, the count against the published BiS file
//    read off disk (the build fetches the same file), landingAction, and readyAnswer / emptyAnswer
//    (lib/bis/landing-answer-view.ts), with the clock this run reads (the build reads the same);
//  - the finished-sim line: finishedRunsFor, specLabel and serverRunsCopy on the Premium fixture.
// The accounts are the home fixtures and web/tests/e2e/fixtures/bis-landing (README.md there).
import { createRequire } from "node:module";
import { existsSync, mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/bis-landing-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

const entry = `
import activeBuild from ${src("src/data/active-build.json")};
import me24 from ${src("tests/e2e/fixtures/bis-landing/me-24.json")};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import me60 from ${src("tests/e2e/fixtures/home/me-60.json")};
import meNew from ${src("tests/e2e/fixtures/home/me-new.json")};
import meNoGear from ${src("tests/e2e/fixtures/bis-landing/me-nogear.json")};
import meSim from ${src("tests/e2e/fixtures/bis-landing/me-sim.json")};
import { loadLandingData } from ${src("src/lib/bis/landing-build.ts")};
import { loadBisFile } from ${src("src/lib/bis/load.ts")};
import { landingCopy } from ${src("src/lib/bis/landing-copy.ts")};
import { landingAction, landingAnswerSpec } from ${src("src/lib/bis/landing-answer.ts")};
import { emptyAnswer, landingAnswerCharacter, readyAnswer } from ${src("src/lib/bis/landing-answer-view.ts")};
import { characterBand } from ${src("src/lib/bis/landing-bands.ts")};
import { planAnswer } from ${src("src/lib/character-header/answer-plan.ts")};
import { finishedRunsFor } from ${src("src/lib/character-header/finished-sim-line.ts")};
import { dayMonth } from ${src("src/lib/guides/guide-header.ts")};
import { answerCopy } from ${src("src/lib/account/answer-copy.ts")};
import { serverRunsCopy } from ${src("src/lib/sim/server-runs-copy.ts")};
import { specLabel } from ${src("src/lib/sim/spec-label.ts")};

async function answerFor(me, config, now) {
  const character = landingAnswerCharacter(me, null);
  if (character === null) return { answer: emptyAnswer(me, false), band: null };
  const spec = landingAnswerSpec(config, async (build, key) => loadBisFile(key, build));
  const plan = planAnswer(character, spec);
  const fetched = plan === 'load' ? await spec.load(character).catch(() => 'failed') : undefined;
  const action = landingAction(character, config);
  return { answer: readyAnswer({ character, plan, fetched, action, now }), band: characterBand(character.level) ?? null,
           faction: character.faction ?? null, crestFrom: character.class ?? null };
}

export async function run() {
  const build = activeBuild.build;
  const data = loadLandingData(build);
  const now = new Date();
  const accounts = {};
  for (const [id, me] of Object.entries({ '24': me24, '34': me34, '60': me60, new: meNew, nogear: meNoGear, sim: meSim })) {
    accounts[id] = await answerFor(me, data.config, now);
  }
  accounts.error = { answer: emptyAnswer(null, true), band: null };
  const runs = finishedRunsFor(meSim);
  const newest = runs[0];
  const label = specLabel(newest.spec);
  return {
    build,
    updated: data.updated === null ? null : dayMonth(data.updated),
    raidLabel: data.raidLabel,
    bands: data.captions.map((c) => ({ band: c.band, label: c.label, summary: landingCopy.bandSummary(c.band === 60, data.raidLabel) })),
    classes: data.cards.map((card) => ({ slug: card.slug, name: card.name, color: card.color,
      specs: card.specs.map((s) => ({ key: s.key, classSlug: s.classSlug, specSlug: s.specSlug, name: s.name, role: s.role,
        icon: s.iconSrc.split('/').pop().replace(/\\.webp$/, '') })) })),
    accounts,
    me: { '24': me24, '34': me34, '60': me60, nogear: meNoGear, sim: meSim },
    finished: { prefix: serverRunsCopy.bandLineParts(label, newest.dps).prefix, figure: serverRunsCopy.bandLineParts(label, newest.dps).figure,
                open: serverRunsCopy.bandOpen, dismiss: serverRunsCopy.dismissLabel(label), more: runs.length - 1 },
    copy: { title: landingCopy.title, tierQuestion: landingCopy.tierQuestion,
            tierLink: landingCopy.tierLink, answerLabel: landingCopy.answerLabel, bandGroupLabel: landingCopy.bandGroupLabel,
            plateLabel: landingCopy.plateLabel, yours: landingCopy.yours, roleWords: landingCopy.roleWords,
            yourRange: Object.fromEntries(data.captions.map((c) => [c.band, landingCopy.yourRange(c.label)])),
            retry: answerCopy.retry },
  };
}`;
// CI's parity job never runs sync:data, so src/data/generated is absent there (guides_facts.mjs does the same).
const generatedFallback = {
  name: "generated-fallback",
  setup(build) {
    build.onResolve({ filter: /data\/generated\/[^/]+\.json$/ }, (args) => {
      const wanted = join(args.resolveDir, args.path);
      return { path: existsSync(wanted) ? wanted : join(WEB, "src/fixtures/planner", basename(args.path)) };
    });
  },
};
const dir = mkdtempSync(join(tmpdir(), "bis-landing-facts-"));
const outfile = join(dir, "facts.mjs");
const metaUrl = pathToFileURL(WEB + "src/lib/bis/load.ts").href;
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm", outfile,
  logLevel: "error", loader: { ".json": "json" }, plugins: [generatedFallback], define: { "import.meta.url": JSON.stringify(metaUrl), "import.meta.env": "{}" } });
const { run } = await import(pathToFileURL(outfile).href);
const facts = await run();
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
