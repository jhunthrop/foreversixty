// The data the BiS parity mocks show, read from a built BiS page: node design/mocks/bis_facts.mjs [base URL]
// (default http://localhost:4321, a running `npm run preview` in web/). Writes design/mocks/data/bis-facts.json,
// which ws_bis.py renders the mocks from.
//
// Facts only: names, numbers, sources, labels, which slot sits in which group. The mock owns which blocks a
// region holds and where they sit (ws_bis.py); nothing structural is read here, so a block the build grows
// (a disclosure, a note, a metrics list) never reaches the mock. The point is that the mock and the build
// show the same data, so the parity gate compares like with like.
import { createRequire } from "node:module";
import { readFileSync, writeFileSync } from "node:fs";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");
const base = process.argv[2] ?? "http://localhost:4321";
const config = JSON.parse(readFileSync(new URL("./web/bis-parity.json", import.meta.url), "utf8"));
const OUT = new URL("./data/bis-facts.json", import.meta.url);

/** Runs in the page: every fact the mock renders. */
function readFacts() {
  const text = (el) => (el?.textContent ?? "").replace(/\s+/g, " ").trim();
  const q = (sel, root = document) => root.querySelector(sel);
  const qa = (sel, root = document) => [...root.querySelectorAll(sel)];
  const rarity = (el) => {
    const m = (el?.className ?? "").match(/text-rarity-(\w+)/);
    return m ? m[1] : "common";
  };
  const figureOf = (root) => {
    if (!root) return null;
    const spans = root.children;
    const label = q(".text-label", root);
    const valueSpan = [...spans].find((s) => s.classList.contains("tabular-nums"));
    const unit = valueSpan?.querySelector("span");
    const value = valueSpan ? text(valueSpan).replace(text(unit), "").trim() : "";
    const caption = [...spans].find((s) => s.classList.contains("text-[12px]"));
    return { label: text(label), value, unit: text(unit), caption: text(caption) };
  };
  const cell = (c) => {
    const flag = q(".ui-doll-flag", c);
    const name = q(".ui-doll-line1 > span:first-child", c);
    const basis = q(".ui-doll-basis", c);
    const basisItem = basis?.querySelector("span");
    const fig = q(".ui-doll-figure .ui-figure", c);
    return {
      slot: c.dataset.dollSlot,
      state: c.dataset.state,
      label: text(q(".ui-doll-label > span:first-child", c)),
      flag: flag ? { text: text(flag), tone: flag.dataset.tone ?? "" } : null,
      item: c.dataset.bisItem ? Number(c.dataset.bisItem) : null,
      name: text(name),
      quality: rarity(name),
      stats: text(q(".ui-doll-stats", c)),
      line2: text(q(".ui-doll-line2", c)),
      gain: fig ? figureOf(fig) : null,
      basis: basis
        ? { prefix: text(basis).replace(text(basisItem), "").trim(), item: text(basisItem), quality: rarity(basisItem) }
        : null,
    };
  };
  const answer = q('[data-testid="bis-this-set"]');
  const tank = qa('[data-testid^="bis-tank-figure-"]').map((f) => ({
    id: f.dataset.testid.replace("bis-tank-figure-", ""),
    label: text(q("dt", f)),
    value: text(q("dd", f)),
  }));
  const trust = q('[data-testid="bis-trust-line"]');
  const doll = q('[data-testid="bis-doll"]');
  const summary = q('[data-testid="bis-summary"]');
  const weights = q('[data-testid^="bis-weights-"][data-region^="section-"]');
  const guide = q('[data-testid="bis-guide-line"]');
  return {
    title: text(q('[data-testid="page-title"]')),
    meta: text(q('[data-testid="bis-page-sub"]')),
    answer: {
      figure: figureOf(q('[data-testid="bis-answer-figure"]')),
      tank,
      trust: text(trust).replace(text(q("a", trust)), "").replace(/\s*·\s*$/, ""),
      rankingSim: text(q('[data-testid="bis-ranking-sim-link"]', answer)),
      simLink: text(q('[data-testid="bis-sim-link"]', answer)),
      plannerLink: text(q('[data-testid="bis-planner-link"]', answer)),
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
    summary: {
      lead: text(summary?.firstChild),
      signIn: text(q('a[href="/login"]', summary)),
      after: (text(summary).split(text(q('a[href="/login"]', summary)))[1] ?? "").replace(text(q("button", summary)), "").trim(),
      next: text(q("button", summary)),
      view: qa('[data-testid="control-view"] [role="radio"]').map((r) => ({ label: text(r), on: r.getAttribute("aria-checked") === "true" })),
    },
    notes: qa('[data-region="notes"] p').map(text),
    doll: {
      label: q('[role="listbox"]', doll).getAttribute("aria-label"),
      view: doll.dataset.view,
      groups: qa(".ui-doll-group", doll).map((g) => ({
        title: text(q(".ui-doll-group-head > span:first-child", g)),
        count: text(q(".ui-doll-group-head > span:last-child", g)),
        cells: qa(".ui-doll-cell", g).map(cell),
      })),
      left: qa('[data-column="left"] .ui-doll-cell', doll).map(cell),
      right: qa('[data-column="right"] .ui-doll-cell', doll).map(cell),
    },
    weights: {
      title: text(q("h2, h3", weights)),
      sentence: text(q('[data-testid="bis-weights-sentence"]', weights)),
      fold: text(q("summary", weights)),
      addon: text(q(".bis-weights-line", weights)).replace(text(q(".bis-weights-line a", weights)), "").trim(),
      addonLink: text(q(".bis-weights-line a", weights)),
    },
    guide: { lead: text(guide).replace(text(q("a", guide)), "").trim(), link: text(q("a", guide)) },
  };
}

const browser = await chromium.launch();
const page = await browser.newPage({ viewport: config.viewport });
const facts = {};
for (const state of config.states) {
  await page.goto(base + state.path, { waitUntil: "networkidle" });
  await page.waitForSelector("html:not([data-bis-pending])", { state: "attached" });
  await page.waitForSelector("astro-island:has([data-testid=bis-page]):not([ssr])", { state: "attached" });
  // Every icon is deferred until it nears the viewport: scroll through so each cell names its icon.
  for (let y = 0; y < 4000; y += 600) {
    await page.evaluate((top) => window.scrollTo(0, top), y);
    await page.waitForTimeout(120);
  }
  const read = await page.evaluate(readFacts);
  // The item icons, from the page's own data (a fixture build ships no icon files, so the DOM shows frames).
  const items = await page.evaluate(async () => {
    const inline = JSON.parse(document.getElementById("bis-page-data")?.textContent ?? "{}").items ?? {};
    const res = await fetch(location.pathname + ".json");
    return { ...inline, ...(res.ok ? ((await res.json()).items ?? {}) : {}) };
  });
  const withIcon = (c) => ({ ...c, icon: c.item === null ? null : (items[c.item]?.icon ?? null) });
  read.doll.left = read.doll.left.map(withIcon);
  read.doll.right = read.doll.right.map(withIcon);
  read.doll.groups = read.doll.groups.map((g) => ({ ...g, cells: g.cells.map(withIcon) }));
  facts[state.id] = { ...state, ...read };
  console.log(state.id, facts[state.id].title, facts[state.id].doll.left.length + facts[state.id].doll.groups.length);
}
await browser.close();
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
