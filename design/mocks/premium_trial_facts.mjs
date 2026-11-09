// The copy and dates the free-trial mocks render (ws_premium_trial.py), from the trial module itself, never retyped:
//   node design/mocks/premium_trial_facts.mjs [out.json]  ->  design/mocks/data/premium-trial-facts.json (or out.json)
// Bundles web/src/lib/billing/trial-copy.ts (and plans.ts under it) with esbuild. The clock is the premium mocks' own,
// 2026-11-08T10:00Z (the founder window open), so a 7-day trial ends 15 Nov 2026 and a code's month 8 Dec 2026.
import { createRequire } from "node:module";
import { mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/premium-trial-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);

const entry = `
import * as t from ${src("src/lib/billing/trial-copy.ts")};
import { PLANS, founderOpen } from ${src("src/lib/billing/plans.ts")};

const START = new Date('2026-11-08T10:00:00Z');
const founder = founderOpen(START);
const LINK = 'https://foreversixty.gg/premium/cancel?t=…';
const ACCOUNT = 'https://foreversixty.gg/account#plan';
const RECEIPT = 'https://invoice.stripe.com/i/…';
const GUILD = 'Olympus XXVII';
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const terms = (plan, interval, code) => t.trialTerms(plan, interval, founder, code, START);
const shown = (x) => ({ ...x, endsAt: x.endsAt.toISOString() });

const mo = terms('premium', 'monthly', false), yr = terms('premium', 'yearly', false), code = terms('premium', 'monthly', true);
const guild = terms('guild', 'yearly', false), guildCode = terms('guild', 'yearly', true);
const remindAt = new Date(mo.endsAt.getTime() - t.reminderLeadDays(mo.days) * 86400000);
const nextAt = new Date(mo.endsAt.getTime()); nextAt.setUTCMonth(nextAt.getUTCMonth() + 1);

export default {
  copy: strip(t.trialCopy),
  founder,
  premium: { monthly: shown(mo), yearly: shown(yr), code: shown(code), name: PLANS.premium.name },
  lines: {
    monthly: t.trialLine(mo), yearly: t.trialLine(yr), code: t.trialLine(code), codeYearly: t.trialLine(terms('premium', 'yearly', true)),
    guild: t.trialLine(guild), guildCode: t.trialLine(guildCode),
  },
  chooser: {
    start: t.trialCopy.chooser.start(mo.days), startCode: t.trialCopy.chooser.start(code.days), subscribeYearly: t.trialCopy.chooser.subscribe(yr.price),
    subscribe: t.trialCopy.chooser.subscribe(mo.price), applied: t.trialCopy.chooser.codeApplied('GRIMBREW'),
  },
  status: {
    premium: t.trialStatusLine(mo.endDate, mo.price), premiumCanceled: t.trialCopy.account.canceled(mo.endDate),
    guild: t.trialStatusLine(guild.endDate, guild.price),
  },
  confirm: { body: t.trialCopy.confirm.body(mo.endDate), guildBody: t.trialCopy.confirm.guildBody(GUILD, guild.endDate) },
  guild: { name: GUILD, plan: PLANS.guild.name, terms: shown(guild), title: t.trialCopy.confirmation.guildTitle(GUILD) },
  email: {
    remindOn: t.trialDate(remindAt), lead: t.reminderLeadDays(mo.days),
    reminder: t.reminderEmail(mo.endDate, mo.price, LINK),
    guildReminder: t.guildReminderEmail(GUILD, PLANS.guild.name, guild.endDate, guild.price, LINK),
    notification: t.reminderNotification(mo.endDate, mo.price),
    receipt: t.receiptEmail(mo.price, mo.endDate, t.trialDate(nextAt), RECEIPT, ACCOUNT),
  },
};`;
const dir = mkdtempSync(join(tmpdir(), "premium-trial-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.env": "{}" }, loader: { ".json": "json", ".svelte": "empty" } });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
