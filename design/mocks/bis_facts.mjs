// The data the BiS parity mocks show, read from a built BiS page:
//   node design/mocks/bis_facts.mjs <base URL> [out file]
// The base URL is required: the build to read, e.g. http://localhost:4500 for `npm run preview -- --port 4500` in web/.
// Writes the facts to [out file], by default design/mocks/data/bis-facts.json, which ws_bis.py renders the mocks from
// (FS_BIS_FACTS points ws_bis.py at another file). The parity spec runs this against its own preview at test time
// (web/tests/e2e/support/parity-mocks.ts, design/parity.md).
//
// Facts only: names, numbers, sources, labels, which slot sits in which group. The mock owns which blocks a
// region holds and where they sit (ws_bis.py); nothing structural is read here, so a block the build grows
// (a disclosure, a note, a metrics list) never reaches the mock. The point is that the mock and the build
// show the same data, so the parity gate compares like with like.
import { createRequire } from "node:module";
import { readFileSync, writeFileSync } from "node:fs";
import { resolve } from "node:path";

const USAGE = [
  "usage: node design/mocks/bis_facts.mjs <base URL> [out file]",
  "  in web/:          npm run build && npm run preview -- --port 4500",
  "  then, repo root:  node design/mocks/bis_facts.mjs http://localhost:4500 && python3 design/mocks/ws_bis.py",
].join("\n");

/** Stops with the usage: the exact command, never a guess at which build to read. */
function fail(why) {
  console.error(`bis_facts.mjs: ${why}\n${USAGE}`);
  process.exit(2);
}

/** The base URL, required and absolute: never whatever happens to hold a default port. */
function baseUrl(arg) {
  if (arg === undefined) fail("the base URL of a running build is required.");
  if (!URL.canParse(arg) || !/^https?:$/.test(new URL(arg).protocol)) fail(`"${arg}" is not an http(s) URL.`);
  return new URL(arg).href.replace(/\/$/, "");
}

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const base = baseUrl(process.argv[2]);
const config = JSON.parse(readFileSync(new URL("./web/bis-parity.json", import.meta.url), "utf8"));
const OUT = process.argv[3] ? resolve(process.argv[3]) : new URL("./data/bis-facts.json", import.meta.url);

/** Runs in the page: every fact the mock renders. */
function readFacts() {
  const text = (el) => (el?.textContent ?? "").replace(/\s+/g, " ").trim();
  const q = (sel, root = document) => root.querySelector(sel);
  const qa = (sel, root = document) => [...root.querySelectorAll(sel)];
  const rarity = (el) => {
    const m = (el?.className ?? "").match(/text-rarity-(\w+)/);
    return m ? m[1] : "common";
  };
  const doll = q('[data-testid="bis-doll"]');
  const frame = q(".pdf-frame", doll);
  const title = q(".pdf-title", frame);
  const slots = qa(".pdf-slot", frame).map((b) => ({
    slot: b.dataset.testid.replace("paperdoll-frame-slot-", ""),
    item: Number(b.dataset.bisItem),
    label: b.getAttribute("aria-label"),
    covered: q(".pdf-covered", b) !== null,
    selected: b.getAttribute("aria-pressed") === "true",
  }));
  const tip = q('.pdf-centre [data-testid="item-tooltip"]', frame);
  const side = (testid) => q(`[data-testid="${testid}"]`);
  const facts = (root) =>
    qa("dt", root).map((dt) => ({ label: text(dt), value: text(dt.nextElementSibling) }));
  const stack = side("bis-weights");
  const where = side("bis-where");
  const set = side("bis-set");
  return {
    title: text(q(".pdf-title-text", title)),
    meta: text(q(".pdf-meta", title)),
    figure: q(".pdf-figure", title)
      ? { value: text(q(".pdf-figure", title).firstChild), unit: text(q(".pdf-figure small", title)) }
      : null,
    frameLabel: doll.getAttribute("aria-label"),
    slots,
    centre: {
      tip: tip
        ? {
            name: tip.getAttribute("aria-label"),
            lines: qa(":scope > div:first-child > *", tip).map(text).filter(Boolean),
            tail: qa('[data-testid^="bis-slot-"]', tip).map(text),
            foot: text(q(":scope > .ui-overlay-rule-t", tip)),
          }
        : null,
      sim: text(q('.pdf-centre [data-testid="bis-sim-link"]', frame)),
    },
    controls: {
      label: q('[data-testid="bis-controls"]').getAttribute("aria-label"),
      groups: qa('[data-testid="bis-controls"] .ui-cbar-group:not([data-testid="bis-sources"])').map((g) => ({
        label: text(q(".ui-cbar-label", g)),
        kind: g.dataset.kind,
        options: qa('[role="radio"]', g).map((r) => ({
          label: text(r),
          on: r.getAttribute("aria-checked") === "true",
          icon: q("img", r) !== null,
        })),
        value: g.dataset.kind === "select" ? text(q("select option:checked", g)) : null,
      })),
      sources: text(q('[data-testid="bis-sources"] button')),
    },
    set: set ? { title: text(q("h2", set)), facts: facts(set), line: text(q(".bis-side-line", set)) } : null,
    stack: {
      title: text(q("h2", stack)),
      facts: facts(stack),
      lines: qa(".bis-side-line", stack).map(text),
      link: text(q('[data-testid="bis-weights-sim-link"]', stack)),
    },
    where: where
      ? {
          title: text(q("h2", where)),
          stops: qa(".bis-stop-link", where).map((p) => ({ title: text(q("a", p)), count: text(q("span", p)) })),
          action: text(q('[data-testid="bis-pre-raid-toggle"]', where)),
        }
      : null,
    itinerary: {
      label: q('[data-testid="bis-itinerary"]').getAttribute("aria-label"),
      stops: qa(".bis-stop").map((s) => ({
        title: text(q(".bis-stop-title", s)).replace(text(q(".bis-stop-range", s)), "").trim(),
        range: text(q(".bis-stop-range", s)),
        rows: qa(".ui-item-row", s).map((r) => {
          const name = q("[data-item-id] [data-testid^='item-hover-'], [data-testid^='item-hover-']", r);
          const fig = q(".ui-figure", r);
          return {
            item: Number(r.dataset.itemId),
            name: text(name),
            quality: rarity(name),
            line: text(q(".ui-item-row-src", r)),
            gain: fig ? text(fig) : "",
          };
        }),
      })),
    },
  };
}

const browser = await chromium.launch();
const page = await browser.newPage({ viewport: config.viewport });
const facts = {};
for (const state of config.states) {
  await page.goto(base + state.path, { waitUntil: "networkidle" });
  await page.waitForSelector("html:not([data-bis-pending])", { state: "attached" });
  await page.waitForSelector("astro-island:has([data-testid=bis-page]):not([ssr])", { state: "attached" });
  const read = await page.evaluate(readFacts);
  // The item icons, from the page's own data (a fixture build ships no icon files, so the DOM shows sockets).
  const items = await page.evaluate(async () => {
    const inline = JSON.parse(document.getElementById("bis-page-data")?.textContent ?? "{}").items ?? {};
    const res = await fetch(location.pathname + ".json");
    return { ...inline, ...(res.ok ? ((await res.json()).items ?? {}) : {}) };
  });
  const withItem = (c) => ({ ...c, icon: items[c.item]?.icon ?? null, quality: c.quality ?? items[c.item]?.quality ?? 1 });
  read.slots = read.slots.map(withItem);
  read.itinerary.stops = read.itinerary.stops.map((s) => ({ ...s, rows: s.rows.map(withItem) }));
  facts[state.id] = { ...state, ...read };
  console.log(state.id, read.title, read.slots.length, read.itinerary.stops.length);
}
await browser.close();
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
