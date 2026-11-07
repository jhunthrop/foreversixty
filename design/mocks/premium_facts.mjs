// The copy and the page models the premium, account and checkout mocks render (design/specs/2026-11-08-premium-account.md):
//   node design/mocks/premium_facts.mjs [out.json]  ->  design/mocks/data/premium-facts.json (or out.json)
// Bundles the TypeScript modules themselves with esbuild (web/node_modules), never retyped:
//  - /premium: lib/billing/plans-page-copy.ts, plans.ts, plan-costs.ts, components/premium/model.ts (priceCell, buttonLabel,
//    costLine, trialParts, featureLines, checkoutHref), lib/billing/copy.ts, premium-source-copy.ts, premium-links.ts, and the
//    sample coach review (coach-sample.ts over src/data/premium-coach-sample.json; spell icons from the build's spellicons.json);
//  - /premium/checkout: lib/billing/copy.ts checkoutCopy, checkout-guild-copy.ts, premium-owned.ts premiumOwnedView;
//  - /account: lib/account/account-page-copy.ts, character-list-copy.ts, devices-copy.ts, signin-copy.ts, discord-link-copy.ts,
//    discord-role-copy.ts, notifications-copy.ts, lib/guild/copy.ts guildConsentCopy, lib/billing/account-plan-copy.ts; the models
//    accountPlanLine and guildRowPlanLine (plan-line.ts), guildPlanTag / notVerifiedMark / guildRowNextStep (guild-row-line.ts),
//    characterDescriptor, mainCharacter, setupAnswer (the head answer), deviceStatusLine, premiumRowValue / premiumRowLinked /
//    blockState (notifications.ts), offerStatusLine, relativeTime; the planner's shared-link copy for /b/:id.
// Every state's account is a fixture (web/tests/e2e/fixtures/home/me-34.json Felsnap, me-60.json Thandrel, me-new.json) with
// the billing, seat, guild and Discord fields of that state set here: DERIVED, marked per state below. The clock is pinned to
// 2026-11-08T10:00Z and the zone to UTC, so every date and "ago" is reproducible.
// CI's parity job has no src/data/generated: the generatedFallback plugin (as setup_facts.mjs) answers from src/fixtures/planner.
import { createRequire } from "node:module";
import { existsSync, mkdtempSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { basename, join } from "node:path";
import { pathToFileURL } from "node:url";

const WEB = new URL("../../web/", import.meta.url).pathname;
const REPO = new URL("../../", import.meta.url).pathname;
const require = createRequire(join(WEB, "package.json"));
const esbuild = require("esbuild");
const OUT = process.argv[2] ? pathToFileURL(process.argv[2]) : new URL("./data/premium-facts.json", import.meta.url);
const src = (p) => JSON.stringify(WEB + p);
const { build } = require(join(WEB, "src/data/active-build.json"));
const spellIcons = JSON.stringify(`${REPO}data/builds/${build}/spellicons.json`);

const entry = `
import { plansPageCopy } from ${src("src/lib/billing/plans-page-copy.ts")};
import { PLANS, PLAN_ORDER, INTERVALS, FOUNDER_LAST_DAY, formatPrice } from ${src("src/lib/billing/plans.ts")};
import { priceFor } from ${src("src/lib/billing/plan-costs.ts")};
import { priceCell, buttonLabel, costLine, trialParts, featureLines, isFirstOfAnchor, checkoutHref } from ${src("src/components/premium/model.ts")};
import { checkoutCopy, billingBlockCopy, premiumBlockCopy } from ${src("src/lib/billing/copy.ts")};
import { premiumSourceCopy } from ${src("src/lib/billing/premium-source-copy.ts")};
import { PREMIUM_LINKS } from ${src("src/lib/billing/premium-links.ts")};
import { checkoutGuildCopy } from ${src("src/lib/billing/checkout-guild-copy.ts")};
import { premiumOwnedView } from ${src("src/lib/billing/premium-owned.ts")};
import { accountPlanCopy } from ${src("src/lib/billing/account-plan-copy.ts")};
import { loadCoachSampleAnswer, coachSampleView } from ${src("src/lib/billing/coach-sample.ts")};
import { accountPageCopy } from ${src("src/lib/account/account-page-copy.ts")};
import { characterListCopy } from ${src("src/lib/account/character-list-copy.ts")};
import { devicesCopy, pairingEntryLabel } from ${src("src/lib/account/devices-copy.ts")};
import { accountSignInCopy } from ${src("src/lib/account/signin-copy.ts")};
import { discordLinkCopy } from ${src("src/lib/account/discord-link-copy.ts")};
import { discordRoleCopy } from ${src("src/lib/account/discord-role-copy.ts")};
import { notificationsCopy, premiumRowCopy } from ${src("src/lib/account/notifications-copy.ts")};
import { guildConsentCopy } from ${src("src/lib/guild/copy.ts")};
import { accountPlanLine, guildRowPlanLine, premiumSource } from ${src("src/lib/account/plan-line.ts")};
import { guildPlanTag, notVerifiedMark, guildRowNextStep } from ${src("src/lib/account/guild-row-line.ts")};
import { characterDescriptor } from ${src("src/lib/account/character-descriptor.ts")};
import { mainCharacter } from ${src("src/lib/account/main-character.ts")};
import { setupAnswer } from ${src("src/lib/setup/answer.ts")};
import { deviceStatusLine } from ${src("src/lib/reports/device-status.ts")};
import { premiumRowValue, premiumRowLinked, premiumRowVisible, blockState } from ${src("src/lib/account/notifications.ts")};
import { offerStatusLine } from ${src("src/lib/account/seat-offer-flow.ts")};
import { ACCOUNT_FAILED } from ${src("src/lib/account/api.ts")};
import { pairingCopy } from ${src("src/lib/reports/copy.ts")};
import { myReportsCopy } from ${src("src/lib/reports/my-reports-copy.ts")};
import { settingsWords } from ${src("src/lib/guild/settings-view.ts")};
import { plannerCopy } from ${src("src/lib/planner/copy.ts")};
import { READ_ONLY_REASON } from ${src("src/lib/planner/store.svelte.ts")};
import { relativeTime } from ${src("src/lib/dates.ts")};
import icons from ${spellIcons};
import me34 from ${src("tests/e2e/fixtures/home/me-34.json")};
import me60 from ${src("tests/e2e/fixtures/home/me-60.json")};
import meNew from ${src("tests/e2e/fixtures/home/me-new.json")};
import devicesFixture from ${src("tests/e2e/fixtures/setup/devices.json")};

const NOW = new Date('2026-11-08T10:00:00Z');
const Z = 'UTC';
const strip = (o) => JSON.parse(JSON.stringify(o, (k, v) => (typeof v === 'function' ? '[fn]' : v)));
const days = (n) => new Date(NOW.getTime() + n * 86400e3).toISOString();

// ---------------------------------------------------------------- /premium: every priced cell, from the card's own model
const priced = {};
for (const plan of PLAN_ORDER) for (const interval of INTERVALS) for (const founder of [true, false]) {
  priced[plan + ':' + interval + ':' + (founder ? 'founder' : 'list')] = {
    cell: priceCell(plan, interval, founder), button: buttonLabel(plan, interval, founder), cost: costLine(plan, interval, founder),
    href: checkoutHref(plan, interval), price: formatPrice(priceFor(plan, interval, founder), interval),
  };
}
const premiumLines = featureLines(plansPageCopy.premium.features).map((l, i, all) => ({ ...l, first: isFirstOfAnchor(all, i) }));
const iconOf = (spellId) => (spellId === null ? null : icons[String(spellId)] ?? null);
const coach = coachSampleView(loadCoachSampleAnswer(), iconOf);

// ---------------------------------------------------------------- the accounts, one per state (DERIVED from the fixtures)
const FREE_ENT = { personal_portal: false, server_sims: false, retention: false, multi_compare: false, history: false, notifications: false,
  officer_views: false, roster_check: false, log_coach: false, supporter_mark: false, billing: null };
const PREMIUM_ENT = { ...FREE_ENT, server_sims: true, retention: true, multi_compare: true, history: true, notifications: true, log_coach: true, supporter_mark: true, personal_portal: true };
const stripe = (o) => ({ source: 'stripe', plan: 'premium', status: 'active', current_period_end: days(361), cancel_at_period_end: false, founder: true, interval: 'yearly', amount_cents: 4000, ...o });
const felsnap = me34.characters[0];
const thandrel = me60.characters[0];
// Two more characters for the several-characters state: a Battle.net alt with no gear yet, and an addon-only alt (DERIVED).
const brisa = { key: 'us-normal-brisa', region: 'us', ruleset: 'normal', name: 'Brisa', class: 'Druid', realm: felsnap.realm, level: 19, faction: 'alliance', race: 'Night Elf', source: 'bnet' };
const korrin = { key: 'us-normal-korrin', region: 'us', ruleset: 'normal', name: 'Korrin', class: 'Priest', realm: felsnap.realm, level: 27, faction: 'alliance', race: 'Dwarf', spec: 'Holy', source: 'export',
  build: { source: 'addon', captured_at: days(-0.2), level: 27 } };
const olympus = me60.guilds[0];
const guildPlan = { tier: 'guild', you_hold_seat: false, scope: 'plan', status: 'active', current_period_end: days(330), cancel_at_period_end: false, billed_by: 'Thandrel', you_are_billing_contact: true };
const discordOff = { notifications: { connect: true, discord: null }, discord_link_available: true, discord_link: null };
const discordOn = { notifications: { connect: true, discord: { channel_id: '1183920045519', source: 'connected', guild_id: '7712', name: 'Forever Sixty', events: { sim_finished: true, upload_parsed: true, upload_private: false }, disabled_reason: null, last_error: null, failures: 0 } },
  discord_link_available: true, discord_link: { display_name: 'thandrel' } };
const base34 = { ...me34, entitlements: FREE_ENT, ...discordOff };
const base60 = { ...me60, entitlements: FREE_ENT, ...discordOff };
const ACCOUNTS = {
  // Felsnap, Battle.net, free, no device, Discord not linked.
  free: base34,
  // No character yet (me-new), signed in with Battle.net.
  new: { ...meNew, entitlements: FREE_ENT, ...discordOff },
  // Felsnap plus two alts: Battle.net and addon characters, Felsnap the chosen main; three devices.
  several: { ...base34, characters: [felsnap, brisa, korrin], main_character_key: felsnap.key },
  // Thandrel, own yearly Premium at the founder price, an officer billing Olympus XXVII on Guild; Discord linked, a channel connected.
  premium: { ...base60, entitlements: { ...PREMIUM_ENT, billing: stripe({}) }, server_sims: { runs_this_month: 4 }, ...discordOn,
    guilds: [{ ...olympus, consent: 'gear', plan: guildPlan, can: ['approve', 'remove', 'settings', 'award', 'billing'] }] },
  // Own monthly Premium at the founder price.
  monthly: { ...base34, entitlements: { ...PREMIUM_ENT, billing: stripe({ interval: 'monthly', amount_cents: 500, current_period_end: days(26) }) } },
  // Own Premium, cancelled: it ends at the period's end.
  ending: { ...base34, entitlements: { ...PREMIUM_ENT, billing: stripe({ cancel_at_period_end: true, current_period_end: days(40) }) } },
  // The owner's grant (forever-owner-account-grant): Premium with no end date.
  grant: { ...base34, entitlements: { ...PREMIUM_ENT, personal_portal: false, billing: { source: 'grant', plan: 'premium', status: 'active', current_period_end: null, cancel_at_period_end: false, founder: false, interval: null, amount_cents: null } } },
  // A seat from Olympus XXVII, given by Thandrel; Felsnap is a verified member there.
  seat: { ...base34, entitlements: { ...PREMIUM_ENT, personal_portal: false }, seat: { guild_name: olympus.name, given_by: 'Thandrel', since: days(-6), starts_at: null },
    guilds: [{ ...olympus, rank: 'Member', consent: 'gear', plan: { ...guildPlan, you_hold_seat: true, you_are_billing_contact: false }, can: [] }] },
  // A seat offered to an account that pays for its own Premium.
  offer: { ...base34, entitlements: { ...PREMIUM_ENT, billing: stripe({ current_period_end: days(120) }) },
    seat_offer: { guild_id: olympus.id, guild_name: olympus.name, from_name: 'Thandrel', offered_at: days(-1), own_ends_at: days(120), own_amount_cents: 4000, own_interval: 'yearly' } },
  // A seat that replaced the account's own plan, taken back 6 days ago.
  ended: { ...base34, entitlements: { ...FREE_ENT, personal_portal: true }, ended_seat: { guild_name: olympus.name, ended_at: days(-6), resubscribe_url: '/premium/checkout?plan=premium&interval=yearly&comeback=cb_9f2', amount_cents: 4000, interval: 'yearly' } },
  // Own Premium whose last payment failed.
  pastdue: { ...base34, entitlements: { ...PREMIUM_ENT, billing: stripe({ status: 'past_due', current_period_end: days(3) }) } },
  // A member (not verified) of a guild with no plan, and a guild whose claim this account started (waits for a second officer).
  claim: { ...base60, guilds: [
    { ...olympus, rank: 'Officer', consent: 'gear', verified: true, verified_count: 18, can: ['approve'], claim: { state: 'pending', yours: false, pending_yours: true, holder_name: null },
      subscribe: { step: 'claim_pending_yours', trial_used: false, available_at: null, reason: null, other_guild_name: null } },
    { id: 7, region: 'us', ruleset: 'normal', name: 'Pinefeather Lodge', slug: 'pinefeather-lodge', rank: 'Member', consent: 'roster', verified: false, verified_count: 4, faction: 'alliance', can: [],
      subscribe: { step: 'ask_leader', trial_used: null, available_at: null, reason: null, other_guild_name: null } }] },
};
const DEVICES = {
  none: [],
  one: devicesFixture,
  several: [...devicesFixture, { id: 'd2', name: 'STEAMDECK', platform: 'linux/amd64', kind: 'companion', created_at: days(-20), last_seen_at: days(-2) },
    { id: 'd3', name: 'MacBook Air', platform: 'darwin/arm64', kind: 'companion', created_at: days(-1), last_seen_at: null }],
};
const view = (me, key) => {
  const plan = accountPlanLine(me, NOW, Z);
  const main = mainCharacter(me.characters, me.main_character_key);
  return {
    key, source: premiumSource(me, NOW), plan,
    battletag: me.user.battletag,
    answer: setupAnswer(me, null, NOW),
    characters: me.characters.map((c) => ({ name: c.name, cls: (c.class ?? '').toLowerCase(), descriptor: characterDescriptor(c), guild: c.guild?.name ?? null,
      main: main?.key === c.key, source: c.build ? c.build.source : null, level: c.level ?? null })),
    imported: me.bnet_imported_at ? characterListCopy.importedFrom(relativeTime(new Date(me.bnet_imported_at), NOW)) : null,
    guilds: me.guilds.map((g) => ({ name: g.name, rank: g.rank ?? null, consent: g.consent ?? 'gear', tag: guildPlanTag(g), mark: notVerifiedMark(g), next: guildRowNextStep(g),
      planLine: g.plan ? guildRowPlanLine(g.plan, Z) : null, can: g.can ?? [], claim: g.claim ?? null, step: g.subscribe?.step ?? null,
      leaveLine: guildConsentCopy.leaveConfirmLine(g.name, me.characters.filter((c) => c.guild?.id === g.id).length || 1),
      leaveOfficer: (g.can ?? []).some((w) => ['approve', 'award', 'settings'].includes(w)) ? guildConsentCopy.leaveLosesOfficer(g.name) : null,
      claimLine: g.subscribe?.step === 'claim_pending_yours' ? checkoutGuildCopy.claimPendingYours(g.name) : null })),
    premiumRow: premiumRowVisible(me) ? { value: premiumRowValue(me), linked: premiumRowLinked(me) } : null,
    discord: { state: blockState(me).state, link: me.discord_link ?? null },
    serverSimsOn: premiumBlockCopy.rows.serverSims.on(me.server_sims?.runs_this_month),
    entitlements: me.entitlements,
    offer: me.seat_offer ? { title: accountPlanCopy.offer.title(me.seat_offer.guild_name), from: accountPlanCopy.offer.from(me.seat_offer.from_name, me.seat_offer.offered_at, Z),
      consequence: accountPlanCopy.offer.consequence(me.seat_offer.own_ends_at, me.seat_offer.own_amount_cents, me.seat_offer.own_interval, Z),
      accepted: offerStatusLine('accepted', me.seat_offer, Z), declined: offerStatusLine('declined', me.seat_offer, Z) } : null,
    owned: premiumSource(me, NOW) === 'none' ? null : premiumOwnedView(me, NOW, Z),
    signedInBanner: main ? accountPageCopy.signedInBanner(main.name) : accountPageCopy.signedInBannerEmpty(me.user.battletag, true),
  };
};
const accounts = Object.fromEntries(Object.entries(ACCOUNTS).map(([k, me]) => [k, view(me, k)]));
const devices = Object.fromEntries(Object.entries(DEVICES).map(([k, list]) => [k, { lines: list.map((d) => deviceStatusLine(d, NOW)), names: list.map((d) => d.name),
  entry: pairingEntryLabel(list.length, accountPageCopy.pairADevice), revoke: list.map((d) => devicesCopy.revokeQuestion(d.name)) }]));

export default {
  now: NOW.toISOString(), build: ${JSON.stringify(build)},
  plans: strip(plansPageCopy), plan_table: PLANS, plan_order: PLAN_ORDER, founder_last_day: FOUNDER_LAST_DAY,
  priced, premium_lines: premiumLines, trial: trialParts(), coach,
  checkout: strip(checkoutCopy), billing_block: strip(billingBlockCopy), premium_block: strip(premiumBlockCopy), source: strip(premiumSourceCopy), links: PREMIUM_LINKS,
  guild_checkout: strip(checkoutGuildCopy),
  guild_checkout_fns: {
    holdsClaim: checkoutGuildCopy.ready.holdsClaim(olympus.name),
    subscribe: checkoutGuildCopy.ready.subscribe(formatPrice(priceFor('guild', 'yearly', true), 'yearly')),
    verified: checkoutGuildCopy.facts.verified(18), noPlan: checkoutGuildCopy.facts.noPlan, onRoster: checkoutGuildCopy.facts.onRoster(31),
    askLeader: checkoutGuildCopy.askLeader('Pinefeather Lodge'), verified4: checkoutGuildCopy.facts.verified(4), onRoster4: checkoutGuildCopy.facts.onRoster(9),
    footPremium: checkoutGuildCopy.foot.premium(formatPrice(priceFor('premium', 'yearly', true), 'yearly')),
    chosen: checkoutCopy.chosenPlan('guild', 'yearly', true),
  },
  account_plan: strip(accountPlanCopy),
  seat_ended_line: accounts.ended.plan.ended,
  account: strip(accountPageCopy), characters_copy: strip(characterListCopy), devices_copy: strip(devicesCopy), signin: accountSignInCopy,
  discord_link: strip(discordLinkCopy), discord_role: discordRoleCopy, notifications: strip(notificationsCopy), premium_row: strip(premiumRowCopy),
  consent: strip(guildConsentCopy), account_failed: ACCOUNT_FAILED, pairing: strip(pairingCopy), my_reports: myReportsCopy,
  accounts, devices,
  guild_on_line: checkoutGuildCopy.onPlan.youPay(olympus.name, 'guild'), guild_settings: settingsWords.guildSettingsLink,
  shared: { note: plannerCopy.sharedLinkNote, fork: plannerCopy.fork, readOnly: READ_ONLY_REASON, shareTitle: plannerCopy.shareTitle },
};`;
const generatedFallback = {
  name: "generated-fallback",
  setup(b) {
    b.onResolve({ filter: /data\/generated\/[^/]+\.json$/ }, (args) => {
      const wanted = join(args.resolveDir, args.path);
      if (existsSync(wanted)) return { path: wanted };
      return { path: join(WEB, "src/fixtures/planner", basename(args.path)) };
    });
  },
};
const dir = mkdtempSync(join(tmpdir(), "premium-facts-"));
const outfile = join(dir, "facts.mjs");
await esbuild.build({ stdin: { contents: entry, resolveDir: WEB, loader: "ts" }, bundle: true, platform: "node", format: "esm",
  outfile, logLevel: "error", define: { "import.meta.env": "{}" }, loader: { ".json": "json", ".svelte": "empty" }, plugins: [generatedFallback] });
const facts = (await import(pathToFileURL(outfile).href)).default;
writeFileSync(OUT, JSON.stringify(facts, null, 1) + "\n");
console.log(OUT.pathname);
