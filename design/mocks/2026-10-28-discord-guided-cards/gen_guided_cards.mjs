// Boards for design/specs/2026-10-28-discord-guided-cards.md: the /bis guided card at each step,
// the result it becomes, and the scheduled raid post with its day-before reminder, at the Discord
// desktop-client (640) and phone-client (390) widths. Mock only. It draws through the reviewers'
// render kit (web/tests/captures/discord-bot/discord-render.ts) and extends it here, in the mock,
// for what the kit lacks: a select showing its chosen option (Discord's default: true), a
// disabled select, one select opened at a time, and the reply chip over a reminder. The result
// frame is the real handler payload bis-top from the kit's payloads.json with the amendment's
// Change picks button added; the post frames start from the real post-open payload.
// Run this file with Node 22 or later (it strips the kit's TypeScript types); boards land here.
import { existsSync, readFileSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(HERE, '../../..');
const WEB = path.join(ROOT, 'web');
const KIT = path.join(WEB, 'tests/captures/discord-bot');
const { renderMessage } = await import(pathToFileURL(path.join(KIT, 'discord-render.ts')).href);
const { chromium } = await import(pathToFileURL(path.join(WEB, 'node_modules/playwright/index.mjs')).href);

const SITE = 'https://foreversixty.gg';
const SITE_URL = /^https:\/\/foreversixty\.gg\//;
const CLOCK = '2026-12-10T22:00:00Z';
const CLIENTS = [
  { name: 'desktop', width: 640 },
  { name: 'phone', width: 390 },
];
const GOLD = 15055189;
const HUNTER = 0xaad372;
const EPHEMERAL = 64;

// Sentinels the mock's post-pass turns into the states the kit does not draw.
const CHOSEN = '⁣';
const DISABLED = '⁤';
const OPEN = '⁥';

const CLASSES = [
  ['Druid', 'Balance, Feral, Feral Bear, Restoration'],
  ['Hunter', 'Beast Mastery, Marksmanship, Survival'],
  ['Mage', 'Arcane, Fire, Frost'],
  ['Paladin', 'Holy, Protection, Retribution'],
  ['Priest', 'Discipline, Holy, Shadow'],
  ['Rogue', 'Assassination, Combat, Subtlety'],
  ['Shaman', 'Elemental, Enhancement, Restoration'],
  ['Warlock', 'Affliction, Demonology, Destruction'],
  ['Warrior', 'Arms, Fury, Protection'],
];
const HUNTER_SPECS = [
  ['Beast Mastery', 'DPS'],
  ['Marksmanship', 'DPS'],
  ['Survival', 'DPS'],
];
const RANGES = [
  ['Level 20', 'Levels 20 to 29'],
  ['Level 30', 'Levels 30 to 39'],
  ['Level 40', 'Levels 40 to 49'],
  ['Level 50', 'Levels 50 to 59'],
  ['Level 60', 'Level 60, no buffs or consumables'],
  ['Phase 1 raid', 'Level 60, raid buffs and consumables'],
];

const options = (rows, chosen) =>
  rows.map(([label, description]) => ({
    label,
    value: label.toLowerCase().replace(/ /g, '-'),
    description: label === chosen ? CHOSEN + description : description,
    default: label === chosen,
  }));

/** A string select: chosen shows its option, disabled greys it, open lists it (one per frame). */
function select(id, placeholder, rows, { chosen, disabled = false, open = false } = {}) {
  let shown = chosen ? CHOSEN + chosen : placeholder;
  if (disabled) shown = DISABLED + placeholder;
  if (open) shown = OPEN + shown;
  return { type: 3, custom_id: id, placeholder: shown, disabled, options: disabled ? [] : options(rows, chosen) };
}
const row = (...components) => ({ type: 1, components });
const button = (label, style, id) => ({ type: 2, style, label, custom_id: id });
const link = (label, url) => ({ type: 2, style: 5, label, url });

const crest = (cls) => ({ url: SITE + '/icons/hd/crests/' + cls + '.webp' });

function card({ title, description, color = GOLD, thumb, classSel, specSel, rangeSel, buttons }) {
  const embed = { author: { name: 'Best in slot' }, title, description, color, footer: { text: 'Forever Sixty' }, timestamp: CLOCK };
  if (thumb) embed.thumbnail = crest(thumb);
  return {
    embeds: [embed],
    components: [row(classSel), row(specSel), row(rangeSel), row(...buttons)],
    flags: EPHEMERAL,
  };
}

const id = (part) => (state) => 'gbis:1:' + part + ':' + state;
const C = id('c');
const S = id('s');
const R = id('r');
const GO = id('go');
const ME = id('me');
const LINK = SITE + '/v1/discord/link/start';

const LEAD_UNLINKED = 'Choose a class, a spec and a range, then Show BiS. Link Discord and this starts from your character.';
const LEAD_READY = 'Ready. Show BiS, or change a pick.';

const picked = (state, cls, spec, range, open) => ({
  classSel: select(C(state), 'Choose a class', CLASSES, { chosen: cls }),
  specSel: select(S(state), 'Choose a spec', HUNTER_SPECS, { chosen: spec, open: open === 'spec' }),
  rangeSel: select(R(state), 'Choose a range', RANGES, { chosen: range, open: open === 'range' }),
});

const BOLD = '\x2a\x2a';
const bold = (s) => BOLD + s + BOLD;
const STALE_LINE = 'This card is out of date. Run ' + String.fromCharCode(47) + 'bis again.';
const NUDGE = bold('Choose a spec first.') + ' Then Show BiS.';
const MAIN_LEAD = 'Picked from your main, ' + bold('Zulmara') + '. Show BiS, or change a pick.';
const HUNTER_SLUG = 'hunter';
const P1 = 'Phase 1 raid';
const showRow = (state, linked) =>
  linked ? [button('Show BiS', 3, GO(state))] : [button('Show BiS', 3, GO(state)), link('Link Discord', LINK)];
const frame = (fid, title, message) => ({ id: fid, title, message });

const HUNTER_NAME = 'Hunter';
const hunterCard = (title, description, state, spec, range, open, linked) =>
  card(Object.assign(picked(state, HUNTER_NAME, spec, range, open), { title, description, color: HUNTER, thumb: HUNTER_SLUG, buttons: showRow(state, linked) }));

const PH_CLASS = 'Choose a class';
const PH_CLASS_FIRST = 'Choose a class first';
const PH_RANGE = 'Choose a range';
const STATE0 = '::p1:';
const unlinkedFirst = card({
  title: PH_CLASS,
  description: LEAD_UNLINKED,
  classSel: select(C(STATE0), PH_CLASS, CLASSES, { open: true }),
  specSel: select(S(STATE0), PH_CLASS_FIRST, HUNTER_SPECS, { disabled: true }),
  rangeSel: select(R(STATE0), PH_RANGE, RANGES, { chosen: P1 }),
  buttons: showRow(STATE0, false),
});

const BM = 'Beast Mastery';
const MM = 'Marksmanship';
const BM_P1 = 'Beast Mastery Hunter · level 60 Phase 1 raid';
const HUNTER_PICK = 'Hunter · choose a spec';
const S_CLASS = 'hunter::p1:';
const S_BM = 'hunter:beast-mastery:p1:';
const S_MM = 'hunter:marksmanship:50:';
const changed = hunterCard('Marksmanship Hunter · level 50', LEAD_READY, S_MM, MM, 'Level 50', '', true);
changed.components[3].components.push(button('Use Zulmara', 2, ME(S_MM)));
const stale = Object.assign(Object.create(null), { content: STALE_LINE, flags: EPHEMERAL });

const frames = [
  frame('bis-card-1-class', 'not linked: the card, the class list open', unlinkedFirst),
  frame('bis-card-2-spec', 'after the class pick: the spec list open', hunterCard(HUNTER_PICK, 'Now choose a Hunter spec.', S_CLASS, undefined, P1, 'spec', false)),
  frame('bis-card-3-range', 'after the spec pick: the range list open', hunterCard(BM_P1, LEAD_READY, S_BM, BM, P1, 'range', false)),
  frame('bis-card-4-ready', 'every pick made: Show BiS', hunterCard(BM_P1, LEAD_READY, S_BM, BM, P1, '', false)),
  frame('bis-card-nudge', 'Show BiS before a spec: the nudge, picks kept', hunterCard(HUNTER_PICK, NUDGE, S_CLASS, undefined, P1, '', false)),
  frame('bis-card-main', 'linked: the card picked from the main', hunterCard(BM_P1, MAIN_LEAD, S_BM, BM, P1, '', true)),
  frame('bis-card-main-changed', 'linked, a pick changed: Use Zulmara appears', changed),
  frame('bis-card-stale', 'a press on a card of an older shape', stale),
];

const payloads = JSON.parse(readFileSync(path.join(KIT, 'payloads.json'), 'utf8'));
const payload = (name) => structuredClone(payloads.find((p) => p.id === name).message);

// The result: the real handler's bis-top payload, plus Change picks before the link.
const result = payload('bis-top');
const [share, allSlots] = result.components[0].components;
result.components = [row(share, button('Change picks', 2, 'gbis:1:chg:hunter:beast-mastery:p1:'), allSlots)];
frames.push({ id: 'bis-result', title: 'Show BiS: the card becomes the typed answer, in place', message: result });

// The raid post the moment the raid is scheduled (Sat 5 Dec, 9:00 pm guild time): nobody yet.
const START = 1796954400;
const scheduled = payload('post-open');
scheduled.nonce = undefined;
scheduled.enforce_nonce = undefined;
scheduled.embeds[0].fields[2].value = 'None yet';
scheduled.embeds[0].timestamp = '2026-12-06T03:00:00Z';
frames.push({ id: 'post-scheduled', title: 'posted on scheduling', clock: '2026-12-06T03:00:00Z', message: scheduled });

// The reminder, 24 h before the start: a reply to the post, the same row.
const when = '<t:' + START + ':F> (<t:' + START + ':R>)';
const reminderText = "**Onyxia's Lair is tomorrow**, " + when + '. 33 of 40 signed up so far.';
frames.push({
  id: 'post-reminder',
  title: 'the day-before reminder, a reply to the post',
  clock: '2026-12-10T02:00:00Z',
  reply: "Forever Sixty  Onyxia's Lair · Thursday 10 December",
  message: { content: reminderText, components: payload('post-open').components },
});

// The reminder through a webhook (no bot): no buttons, a masked link to the raid page.
const raidURL = SITE + '/guild/us/normal/olympus-xxvii/raids/1';
frames.push({
  id: 'post-reminder-webhook',
  title: 'the day-before reminder through a webhook',
  clock: '2026-12-10T02:00:00Z',
  message: { content: reminderText + ' [Sign up on Forever Sixty](' + raidURL + ')' },
});

// The post itself on the reminder, when it is still the newest message: the Tomorrow state.
const TOMORROW_TITLE = "Tomorrow: Onyxia's Lair · Thursday 10 December";
const TOMORROW_DESC = 'The raid is tomorrow. Sign up with the buttons below.';
const REMINDER_CLOCK = '2026-12-10T02:00:00Z';
const tomorrow = payload('post-open');
tomorrow.nonce = undefined;
tomorrow.enforce_nonce = undefined;
tomorrow.embeds[0].title = TOMORROW_TITLE;
tomorrow.embeds[0].description = TOMORROW_DESC;
tomorrow.embeds[0].timestamp = REMINDER_CLOCK;
frames.push(Object.assign(frame('post-tomorrow', 'the reminder as an edit of the post', tomorrow), { clock: REMINDER_CLOCK }));

const EXTRA_CSS = '<style>' +
  '.select .chosen{color:#dbdee1}' +
  '.select.dis{opacity:.5;cursor:not-allowed}' +
  '.option.sel{background:#404249;border-radius:3px}' +
  '.option.sel b::after{content:"  \\2713";color:#00a8fc}' +
  '</style>';

const CLOSED_LIST = /<div class="select open"><span>(?!⁥)([^<]*)<\/span><span>⌃<\/span><\/div><div class="options">(?:<div class="option">.*?<\/div>)*<\/div>/g;

const OPEN_MARK = /<span>⁥/g;
const DISABLED_SELECT = /<div class="select(?: open)?"><span>⁤([^<]*)<\/span>/g;
const CHOSEN_VALUE = /<span>⁣([^<]*)<\/span>/g;
const CHOSEN_OPTION = /<div class="option"><b>([^<]*)<\/b><small>⁣/g;
const ICON_PATH = /^\/data\/([^/]+)\/(icons\/.+)/;
const SPAN = '<span>';
const SPAN_END = '</span>';
const STYLE_END = '</style>';
const closedSelect = (_m, label) => '<div class="select"><span>' + label + '</span><span>⌄</span></div>';
const disabledSelect = (_m, label) => '<div class="select dis"><span>' + label + SPAN_END;
const chosenValue = (_m, label) => '<span class="chosen">' + label + SPAN_END;
const chosenOption = (_m, label) => '<div class="option sel"><b>' + label + '</b><small>';

/** Draws what the kit lacks: the chosen value, the disabled select, one list open. */
const postPass = (html) =>
  html
    .replace(CLOSED_LIST, closedSelect)
    .replace(OPEN_MARK, SPAN)
    .replace(DISABLED_SELECT, disabledSelect)
    .replace(CHOSEN_VALUE, chosenValue)
    .replace(CHOSEN_OPTION, chosenOption)
    .replace(STYLE_END, STYLE_END + EXTRA_CSS);

const BUILDS = path.join(ROOT, 'data', 'builds');
const PUBLIC = path.join(WEB, 'public');
const iconFile = (match) => path.join(BUILDS, match[1], match[2]);
const localFile = (url) => (ICON_PATH.test(url.pathname) ? iconFile(ICON_PATH.exec(url.pathname)) : path.join(PUBLIC, url.pathname));

const serve = async (route) => {
  const file = localFile(new URL(route.request().url()));
  return existsSync(file) ? route.fulfill({ path: file, contentType: IMAGE_TYPE }) : route.abort();
};
const IMAGE_TYPE = 'image/webp';

const shoot = async (page, client, f) => {
  const entry = { id: f.id, row: ROW, title: f.title, surface: SURFACE, clock: f.clock ?? CLOCK, message: f.message };
  const html = postPass(renderMessage(entry, { width: client.width, selectOpen: true, usedCommand: f.reply }));
  const name = f.id + DASH + client.name;
  writeFileSync(path.join(HERE, name + HTML_EXT), html);
  await page.setViewportSize({ width: client.width, height: 600 });
  await page.setContent(html, { waitUntil: LOAD });
  await page.locator(MSG).screenshot({ path: path.join(HERE, name + PNG_EXT) });
};
const ROW = 'G';
const SURFACE = 'ephemeral';
const DASH = '-';
const HTML_EXT = '.html';
const PNG_EXT = '.png';
const LOAD = 'load';
const MSG = '#msg';

const browser = await chromium.launch();
const page = await browser.newPage({ deviceScaleFactor: 2 });
await page.route(SITE_URL, serve);
for (const client of CLIENTS) for (const f of frames) await shoot(page, client, f);
await browser.close();
console.log('wrote', frames.length * CLIENTS.length, 'boards to', HERE);
