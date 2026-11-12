// The copy and figures the creator program's reference mocks render (design/mocks/ws_creators.py), from the pages' own
// modules, never retyped:
//   node design/mocks/creators_facts.mjs [out.json]  ->  design/mocks/data/creators-facts.json (or out.json)
// Bundles web/src/lib/creators/{copy,model}.ts with esbuild (from web/node_modules): `page` is the public /creators page
// (the program's sections and figures, the head's answer, the Apply path, the worked examples at today's prices through
// lib/creators/model.ts earningsRows); the rest are the program's pages over the sample creators the e2e and parity specs
// serve (web/tests/e2e/fixtures/creators/*.json). Their clock is 2026-12-09 (November closed, December open).
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/creators-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

const entry = `
import * as copy from ${src("src/lib/creators/copy.ts")};
import * as m from ${src("src/lib/creators/model.ts")};
import { founderOpen } from ${src("src/lib/billing/plans.ts")};
import creators from ${src("tests/e2e/fixtures/creators/creators.json")};
import page from ${src("tests/e2e/fixtures/creators/statements.json")};
const statements = page.statements;
import me from ${src("tests/e2e/fixtures/creators/me.json")};
import requests from ${src("tests/e2e/fixtures/creators/requests.json")};

const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
// The fixtures' clock (spec 0B): 9 Dec 2026, November closed and December open.
const CLOCK = new Date('2026-12-09T12:00:00Z');
const a = copy.adminCreatorsCopy, p = copy.payoutsCopy, st = copy.creatorStatsCopy, ob = copy.onboardingCopy;
const adminRows = creators.map((c) => ({
  name: c.name,
  sub: a.sub(c.slug, c.code, m.dayLabel(c.starts_on)),
  status: copy.stripeStatusCopy[c.stripe_status], tone: m.STATUS_TONE[c.stripe_status],
  edit: a.edit, figures: m.creatorFigures(c),
}));

const action = (s) => (m.canApprove(s) ? p.approve(m.money(s.total_cents)) : '');
const stRows = statements.map((s) => ({
  name: s.name, sub: m.statementStatusLine(s), action: action(s), lines: p.lines,
  notReady: s.status === 'blocked' ? p.notReady : '', figures: m.statementFigures(s),
}));
const first = statements[0];
const lines = first.lines.map(m.lineText);

export default {
  // The public page, as its build computes it (founder prices while the window is open).
  page: { program: strip(copy.programCopy), terms: copy.PROGRAM_TERMS, applyHref: copy.CREATORS_APPLY_PATH,
    examples: m.earningsRows(founderOpen(new Date())) },
  admin: {
    copy: strip(a), meta: a.tableMeta(creators.length), rows: adminRows, editTitle: a.editTitle(creators[0].name),
    fields: Object.values(a.fields).map((f) => strip(f)),
  },
  payouts: {
    copy: strip(p), off: p.payoutsOff, month: m.monthLabel(first.month), title: p.tableTitle(statements.length), meta: p.tableMeta(m.money(m.toPay(statements))),
    rows: stRows, linesTitle: p.linesTitle(first.name), lines,
    confirm: { title: p.confirmTitle(first.name, m.money(first.total_cents), m.monthLabel(first.month)), body: p.confirmBody(m.money(first.total_cents)), pay: p.confirmPay(m.money(first.total_cents)), cancel: p.cancel,
      balance: p.balanceLine(m.money(50000), m.money(first.total_cents)), balanceShortLine: p.balanceLine(m.money(12000), m.money(first.total_cents)), short: p.balanceShort },
  },
  onboarding: { copy: strip(ob), join: ob.inviteJoin(me.name) },
  requests: {
    copy: strip(copy.requestsCopy), meta: copy.requestsCopy.meta(requests.length),
    rows: requests.map((r) => ({ title: m.channelTitle(r.channel_url), sub: m.requestSub(r), about: r.about, url: r.channel_url, email: r.contact_email })),
  },
  program: strip(copy.programCopy),
  apply: {
    copy: strip(copy.applyCopy), counter: copy.applyCopy.counter(0, m.ABOUT_MAX),
    pending: copy.applyCopy.pending(m.dayLabel('2026-12-08T16:20:00Z')),
  },
  stats: {
    copy: strip(st), name: me.name, meta: st.meta(me.slug, me.code), figures: m.statsFigures(me),
    weeks: me.weeks.map((w) => m.weekRow(w, CLOCK)), filling: st.filling,
    status: copy.stripeStatusCopy[me.stripe_status], tone: m.STATUS_TONE[me.stripe_status],
    closed: m.closedMonthLine(me.month, CLOCK), paidSoFar: st.paidSoFar(m.money(me.payouts.reduce((n, p) => n + p.amount_cents, 0))), dashboard: me.stripe_status === 'ready' ? st.dashboard : '',
    payouts: me.payouts.map((x) => ({ text: m.monthLabel(x.month), amount: m.money(x.amount_cents) })),
  },
};`;
const dir = mkdtempSync(join(tmpdir(), "creators-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.env": "{}" }, loader: { ".json": "json", ".svelte": "empty" } });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
