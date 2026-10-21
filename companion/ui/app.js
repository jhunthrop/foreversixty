// companion/ui/app.js
// The whole window. It polls /api/status every two seconds and posts
// what the player does. There is no framework and no build step: four
// pages over one JSON document. Every sentence the window owns is in
// copy.js; the rest arrives in the snapshot from internal/app/copy.go
// and is shown as it is. Nothing on screen is ever an error's text.
//
// A region is re-rendered only when the JSON behind it changed, so the
// two-second poll never costs a focused link its focus, a hover its
// state, or an open dropdown its list.
'use strict';

const POLL_MS = 2000;
const NOTE_MS = 4000;
const FLASH_MS = 4000;
// How long the page, headline and first-paint fades run (--dur-med in styles.css).
const ENTER_MS = 160;
// The first poll gets this long before the bar says the service is not answering.
const SLOW_FIRST_POLL_MS = 1500;
// "Updated to ..." stays in the footer for this long after a new version's first launch.
const UPDATED_MS = 30000;
// The Reports page lists the newest few; the rest are one link away on the site.
const REPORT_ROWS = 10;

const $ = (id) => document.getElementById(id);

// h builds an element: h('a', {class: 'x', href: '...'}, 'text', child).
function h(tag, props, ...kids) {
  const el = document.createElement(tag);
  for (const [k, v] of Object.entries(props || {})) {
    if (v === undefined || v === null || v === false) continue;
    if (k === 'class') el.className = v;
    else if (k === 'text') el.textContent = v;
    else if (k.startsWith('on')) el.addEventListener(k.slice(2), v);
    else el.setAttribute(k, v === true ? '' : v);
  }
  for (const kid of kids) {
    if (kid === undefined || kid === null || kid === false) continue;
    el.append(kid.nodeType ? kid : document.createTextNode(String(kid)));
  }
  return el;
}

// icon is an inline SVG on a 24 px grid, 2 px stroke, round caps, in the
// colour of the text around it. Nothing on screen is a text glyph or an
// emoji: the system font a glyph falls back to is not ours.
const ICONS = {
  check: ['M20 6 9 17l-5-5'],
  triangle: ['M10.3 3.9 1.8 18a2 2 0 0 0 1.7 3h17a2 2 0 0 0 1.7-3L13.7 3.9a2 2 0 0 0-3.4 0z', 'M12 9v4M12 17h.01'],
  file: ['M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z', 'M14 2v6h6M8 13h8M8 17h6'],
  chevron: ['m9 18 6-6-6-6'],
};
function icon(name, size, cls) {
  const NS = 'http://www.w3.org/2000/svg';
  const svg = document.createElementNS(NS, 'svg');
  svg.setAttribute('viewBox', '0 0 24 24');
  svg.setAttribute('width', String(size));
  svg.setAttribute('height', String(size));
  svg.setAttribute('class', 'icon' + (cls ? ' ' + cls : ''));
  svg.setAttribute('aria-hidden', 'true');
  for (const d of ICONS[name]) {
    const p = document.createElementNS(NS, 'path');
    p.setAttribute('d', d);
    svg.append(p);
  }
  return svg;
}

// ---- talking to the companion ----------------------------------------

// api answers {ok, data} or {ok: false, error: {code, message}}. A
// request that fails to arrive, or arrives as something other than the
// envelope, is the companion's own service being down; its message is
// ours, not the exception's.
async function api(path, body) {
  const down = { ok: false, error: { code: 'local', message: COPY.localDown } };
  try {
    const res = await fetch(path, body === undefined
      ? { cache: 'no-store' }
      : { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(body) });
    const r = await res.json();
    if (!r || typeof r.ok !== 'boolean') return down;
    if (!r.ok) r.error = { code: (r.error && r.error.code) || '', message: (r.error && r.error.message) || COPY.localDown };
    return r;
  } catch (e) {
    return down;
  }
}

// ---- small shared pieces ---------------------------------------------

const timers = new Map();
// note writes a message under a control. A success clears itself after
// four seconds; a failure stays until the next action.
function note(id, kind, text) {
  const el = $(id);
  clearTimeout(timers.get(id));
  el.className = 'note' + (kind === 'bad' ? ' bad' : kind === 'ok' ? ' ok' : '');
  el.replaceChildren();
  if (text && kind === 'ok') el.append(icon('check', 14));
  if (text) el.append(document.createTextNode(text));
  if (text && kind !== 'bad') timers.set(id, setTimeout(() => { el.replaceChildren(); }, NOTE_MS));
}

// busy disables a button and gives it a "doing it" label while fn runs.
async function busy(btn, label, fn) {
  const was = btn.textContent;
  btn.disabled = true;
  btn.textContent = label;
  try {
    return await fn();
  } finally {
    btn.disabled = false;
    btn.textContent = was;
  }
}

function dot(tone) { return h('span', { class: 'dot ' + tone, 'aria-hidden': 'true' }); }

// An external link: the real href stays (for the status bar and for
// copying), the click goes through the companion, which opens the
// system browser. The webview never opens a window of its own.
function ext(text, url, props) {
  return h('a', Object.assign({ href: url, 'data-open': '' }, props || {}), text);
}
document.addEventListener('click', (e) => {
  const a = e.target.closest && e.target.closest('a[data-open]');
  if (!a) return;
  e.preventDefault();
  openLink(a.getAttribute('href'));
});
async function openLink(url) {
  if (!url || url === '#') return;
  const r = await api('api/open', { url });
  if (!r.ok) showBar('tray-bar', r.error.message, 6000);
}
let barTimer = null;
function showBar(id, text, ms) {
  const bar = $(id);
  bar.hidden = false;
  if (id === 'poll-bar') {
    // The poll bar is the one with a mark: the companion's service is not answering.
    bar.replaceChildren(icon('triangle', 16), h('span', { text }));
  } else {
    const t = bar.querySelector('span') || bar;
    t.textContent = text;
  }
  clearTimeout(barTimer);
  if (ms) barTimer = setTimeout(() => { bar.hidden = true; }, ms);
}

// copiedLabel is the "Copied." a Copy link button says for two seconds,
// with the check before it, then goes back to its own words.
function copiedLabel(btn, was, copied) {
  btn.replaceChildren(icon('check', 12, 'ok'), copied);
  setTimeout(() => { btn.textContent = was; }, 2000);
}

async function copyText(text) {
  try {
    await navigator.clipboard.writeText(text);
    return true;
  } catch (e) {
    const ta = h('textarea', { 'aria-hidden': 'true', style: 'position:fixed;left:-1000px' });
    ta.value = text;
    document.body.append(ta);
    ta.select();
    let ok = false;
    try { ok = document.execCommand('copy'); } catch (e2) { ok = false; }
    ta.remove();
    return ok;
  }
}

// ---- regions: render only what changed -------------------------------

const cache = {};

// region renders into a container when the data behind it changed. A
// focused control inside it, and anything the player typed into an
// input marked data-keep, survive the render.
function region(name, container, data, render) {
  const key = JSON.stringify(data);
  if (cache[name] === key) return false;
  cache[name] = key;
  const active = document.activeElement;
  const fid = active && container.contains(active) ? active.getAttribute('data-fid') : null;
  const kept = {};
  for (const el of container.querySelectorAll('[data-keep]')) kept[el.getAttribute('data-keep')] = el.value;
  render(container);
  for (const el of container.querySelectorAll('[data-keep]')) {
    const v = kept[el.getAttribute('data-keep')];
    if (v !== undefined && !el.value) el.value = v;
  }
  if (fid) {
    const again = container.querySelector('[data-fid="' + fid + '"]');
    if (again) again.focus();
  }
  return true;
}

// ---- navigation ------------------------------------------------------

let page = null;
let snap = null;
const local = { pairOpen: false, confirmForget: false, forgot: false, flash: false, flashDrawn: false, copied: false, busyPair: false, starting: false, checkError: '', pollFailed: false };

// replay adds a class for one animation's length, so a class that was
// already there plays again.
function replay(el, cls, ms) {
  el.classList.remove(cls);
  void el.offsetWidth;
  el.classList.add(cls);
  setTimeout(() => el.classList.remove(cls), ms);
}

function show(name, focus) {
  const leaving = page;
  page = name;
  for (const s of document.querySelectorAll('main section')) s.hidden = s.id !== name;
  // The incoming page eases in; the outgoing one is simply gone, and a
  // poll that only changes data never gets here.
  if (leaving !== null && leaving !== name) replay($(name), 'enter', ENTER_MS);
  for (const b of document.querySelectorAll('nav button')) {
    if (b.dataset.page === name) b.setAttribute('aria-current', 'page');
    else b.removeAttribute('aria-current');
  }
  if (leaving === 'settings' && name !== 'settings') resetSettings();
  if (leaving === 'device' && name !== 'device') { local.confirmForget = false; local.forgot = false; local.flash = false; if (snap) renderDevice(snap); }
  if (focus) $(name).querySelector('h2').focus();
}

for (const b of document.querySelectorAll('nav button')) {
  b.textContent = COPY.nav[b.dataset.page];
  b.addEventListener('click', () => show(b.dataset.page, true));
}
for (const s of document.querySelectorAll('main section')) s.querySelector('h2').textContent = COPY.pages[s.id];

// ---- Status ----------------------------------------------------------

// The key of the headline last drawn: a change of state fades the card in,
// a poll that leaves the state alone never does.
let lastHeadlineKey = null;

// unavailableHeadline is what Status says while the companion's own
// service does not answer: grey, no button, the last numbers dimmed below.
function unavailableHeadline() {
  return { key: 'unavailable', tone: 'off', title: COPY.status.unavailableTitle, sub: COPY.status.unavailableSub };
}

function renderStatus(s) {
  const p = s.pipeline;
  const failed = local.pollFailed;
  const hd = failed ? unavailableHeadline() : s.headline;
  const reportURL = p.report_id ? s.site_base_url + '/reports/' + p.report_id : '';
  $('status-rows').classList.toggle('dim', failed);
  $('installs').classList.toggle('dim', failed);

  region('headline', $('headline'), { hd, reportURL }, (el) => {
    el.dataset.tone = hd.tone;
    el.dataset.key = hd.key;
    el.replaceChildren();
    // Connecting shows a spinner where every other state has its dot.
    const mark = hd.key === 'connecting' ? h('span', { class: 'spinner', 'aria-hidden': 'true' }) : dot(hd.tone);
    el.append(
      h('div', { class: 'head' }, mark, h('div', { class: 'title', text: hd.title })),
      h('p', { class: 'sub', text: hd.sub }),
    );
    const actions = [];
    if (hd.action) {
      // The account button is for signing in and nothing else.
      const account = hd.key === 'not_connected' || hd.key === 'signed_out';
      actions.push(h('button', { class: account ? 'account' : 'action primary big', type: 'button', 'data-fid': 'headline-action',
        onclick: () => show(hd.action.page, true) }, hd.action.label));
    }
    if (hd.key === 'logging' && reportURL) actions.push(ext(COPY.status.reportLogging, reportURL, { 'data-fid': 'headline-report' }));
    if (actions.length) el.append(h('div', { class: 'actions' }, ...actions));
    if (hd.key !== lastHeadlineKey) {
      lastHeadlineKey = hd.key;
      replay(el, 'enter', ENTER_MS);
    }
  });

  const rows = {
    report: p.report_id ? { zone: ((s.reports || []).find((r) => r.key === p.report_key) || {}).zone || '', credit: s.credit } : (p.report_key ? 'starting' : 'none'),
    // The words, not the time: the relative time moves on its own, and the
    // region only redraws when what it holds changes.
    lastSync: lastSyncRow(s.last_sync),
    fights: p.fights, queued: p.queued, addon: s.addon_line, stale: s.addon_stale, never: s.addon_never, help: s.addon_help, title: s.addon_seen_title,
    // The Check for builds now button and the line it leaves. The words are
    // computed here, like the Last sync row's, so they move on their own.
    paired: s.paired,
    checked: s.builds_checked ? FORMAT.checkedLine(s.builds_checked, new Date(), COPY.status, COPY.status.ago) : '',
    checking: checkBuilds.running,
    canCheck: checkBuilds.canPress(),
    checkError: local.checkError,
    failed,
  };
  region('status-rows', $('status-rows'), rows, (el) => {
    el.replaceChildren();
    if (rows.failed) el.append(h('div', { class: 'stale-cap', text: COPY.status.lastKnown }));
    let report;
    if (rows.report === 'none') report = h('span', { class: 'v', text: COPY.status.noReport });
    else if (rows.report === 'starting') report = h('span', { class: 'v', text: COPY.status.starting });
    else {
      // "In progress · credited to Bow Jackzon", or "... · not credited yet" in
      // the warning token with the one muted line that says what to do.
      const cr = FORMAT.currentReport(rows.report, COPY.status);
      report = h('span', { class: 'v' }, cr.lead + cr.separator,
        cr.credited ? cr.credit : h('span', { class: cr.noName ? 'muted-word' : 'warn', text: cr.credit }),
        cr.help ? h('span', { class: 'help13', text: cr.help }) : null);
    }
    el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.report }), report));
    // An open report with no fight yet says so, rather than a bare 0 that reads as lost;
    // with no report open the count is "None yet", like the Current report row above it.
    // Every count is in the body face with tabular digits.
    const noFights = p.fights === 0 && p.report_id && p.logging;
    el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.fights }),
      noFights
        ? h('span', { class: 'v' }, COPY.status.noFightsYet, h('span', { class: 'sub', text: COPY.status.noFightsHelp }))
        : h('span', { class: 'v', text: p.fights === 0 ? COPY.status.noReport : String(p.fights) })));
    if (p.queued > 0) {
      el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.queued }),
        h('span', { class: 'v', text: String(p.queued) })));
    }
    const v = h('span', { class: 'v' + (s.addon_stale ? ' amber' : '') + (s.addon_never ? ' muted' : ''), title: s.addon_seen_title || null, text: s.addon_line });
    // A sync over a day old, and no sync at all, each say what to do.
    if ((s.addon_stale || s.addon_never) && s.addon_help) v.append(h('span', { class: 'sub', text: s.addon_help }));
    if (rows.checked) v.append(h('span', { class: 'sub', text: rows.checked }));
    if (rows.checkError) v.append(h('span', { class: 'sub warn', text: rows.checkError }));
    if (rows.paired) {
      v.append(h('button', { class: 'ghost small check-builds', type: 'button', 'data-fid': 'check-builds',
        disabled: rows.checking || !rows.canCheck || rows.failed, onclick: onCheckBuilds },
      rows.checking ? h('span', { class: 'spinner small', 'aria-hidden': 'true' }) : null,
      rows.checking ? COPY.status.checking : COPY.status.checkBuilds));
    }
    el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.addon }), v));
    const ls = rows.lastSync;
    if (ls) {
      el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.lastSync }),
        h('span', { class: 'v' + (ls.tone ? ' ' + ls.tone : ''), title: ls.title || null, text: ls.text })));
    }
  });

  region('installs', $('installs'), { installs: s.installs, key: s.headline.key, budget: FORMAT.pathBudget(window.innerWidth, true) }, (el) => {
    el.replaceChildren();
    if (!s.installs.length) {
      el.append(h('p', { class: 'note', text: COPY.status.noFolders }));
      return;
    }
    for (const i of s.installs) el.append(installRow(i, s.headline.key));
  });
}

// lastSyncRow is the Last sync row's words, or null when the Addon row's
// install line says it all.
function lastSyncRow(ls) {
  if (!ls) return null;
  if (ls.state === 'sent') {
    return { text: COPY.status.lastSyncSent(ls.count, FORMAT.ago(ls.at, new Date(), COPY.status.ago)), title: ls.title };
  }
  if (ls.state === 'refused') return { text: ls.line, tone: 'amber' };
  if (ls.state === 'none') return { text: ls.line, tone: 'muted' };
  return null;
}

// checkBuilds is the Check for builds now button's state: one press at a
// time and one every half minute (FORMAT.checker).
const checkBuilds = FORMAT.checker(async () => {
  local.checkError = '';
  const r = await api('api/builds/check', {});
  if (r.ok) applySnapshot(r.data);
  else local.checkError = r.error.message;
}, Date.now);

async function onCheckBuilds() {
  const pressed = checkBuilds.press();
  if (snap) renderStatus(snap); // "Checking…" now, not when the poll next comes round
  await pressed;
  if (snap) renderStatus(snap);
}

function installRow(i, headlineKey) {
  const top = h('div', { class: 'top-line' },
    h('div', { class: 'name' }, i.flavor_name || i.flavor,
      i.watching ? h('span', { class: 'tag', title: COPY.status.watchingTitle, text: COPY.status.watching }) : null),
    advancedLabel(i, headlineKey));
  const row = h('div', { class: 'install' }, top,
    h('span', { class: 'played', text: FORMAT.lastPlayed(i.last_played, new Date(), COPY.status) }),
    h('span', { class: 'path', title: i.path, text: FORMAT.pathTruncate(i.path, FORMAT.pathBudget(window.innerWidth, true)) }));
  // The instruction lives in the headline when the headline is about
  // advanced logging; a row keeps its short line only for the watched
  // install, and an install the player is not logging has none.
  const adv = FORMAT.advancedView(i, headlineKey, COPY.status);
  if (adv.help) row.append(h('p', { class: 'help-line', text: adv.help }));
  return row;
}

function advancedLabel(i, headlineKey) {
  const adv = FORMAT.advancedView(i, headlineKey, COPY.status);
  if (adv.kind === 'unknown') return h('span', { class: 'adv', text: adv.text });
  if (adv.kind === 'on') return h('span', { class: 'adv on' }, icon('check', 14), adv.text);
  if (adv.kind === 'muted') return h('span', { class: 'adv', text: adv.text });
  return h('span', { class: 'adv off' }, icon('triangle', 14), adv.text);
}

// ---- Reports ---------------------------------------------------------

function renderReports(s) {
  const all = $('reports-all');
  if (all.getAttribute('href') !== s.site_base_url + '/logs') all.setAttribute('href', s.site_base_url + '/logs');
  $('report-panel').classList.toggle('dim', local.pollFailed);
  // Signed out, the one next step is signing in; the link would compete with it.
  all.hidden = !s.paired;

  const shown = s.reports.slice(0, REPORT_ROWS);
  $('reports-hint').hidden = !FORMAT.anyNotCredited(shown, COPY.status.credit);
  region('reports', $('report-panel'), { reports: s.reports, paired: s.paired, site: s.site_base_url }, (el) => {
    el.replaceChildren();
    if (!s.reports.length) {
      const box = h('div', { class: 'empty' });
      if (!s.paired) {
        box.append(h('p', { class: 'big-line', text: COPY.reports.unpairedTitle }),
          h('button', { class: 'account', type: 'button', onclick: () => show('device', true) }, COPY.reports.unpairedButton));
      } else {
        box.append(icon('file', 32, 'big-icon'), h('p', { class: 'big-line', text: COPY.reports.emptyTitle }),
          h('p', { class: 'note', text: COPY.reports.emptyBody }));
      }
      el.append(box);
      return;
    }
    const list = h('ul', { class: 'reports' });
    const now = new Date();
    for (const r of shown) list.append(reportRow(r, now));
    el.append(list);
  });
}

// reportRow: a marker, the zone and one meta line, and Copy link. The whole
// row is the link to the report; only the credit that is missing is loud.
function reportRow(r, now) {
  const zone = r.zone || COPY.reports.unnamed;
  const title = r.url ? ext(zone, r.url, { 'data-fid': 'report-' + r.key }) : h('span', { text: zone });
  // The marker column is always there so the rows line up: a live dot while
  // the report is open, a ring when it never reached the site, else empty.
  const marker = h('span', { class: 'marker' + (!r.done ? ' live' : (!r.url ? ' pending' : '')), 'aria-hidden': 'true' });
  const credit = COPY.status.credit;
  const cr = FORMAT.reportCredit(r, credit);
  const lead = credit.credited('');
  const parts = [FORMAT.when(r.started_at, now, COPY.reports), COPY.reports.fights(r.fights)];
  if (!r.done) parts.push(h('span', { class: 'live', text: COPY.reports.inProgress }));
  if (cr.kind === 'absent') {
    // The site has no such report yet, so there is nothing to credit: one muted phrase, no orange.
    parts.push(h('span', {}, COPY.reports.notOnSite + credit.separator + COPY.reports.notOnSiteWhen));
  } else if (cr.kind === 'credited') {
    parts.push(h('span', {}, lead, h('span', { class: 'who', text: cr.text.slice(lead.length) })));
  } else if (cr.kind === 'noname') {
    parts.push(h('span', {}, cr.text));
  } else {
    parts.push(h('span', { class: 'not-credited', title: credit.notYetHelp, text: cr.text }));
  }
  const meta = h('div', { class: 'when' });
  parts.forEach((p, n) => { if (n) meta.append(credit.separator); meta.append(p); });
  const li = h('li', {}, marker, h('div', { class: 'what' }, h('div', { class: 'zone' }, title), meta));
  if (r.url) {
    const btn = h('button', { class: 'ghost small', type: 'button', 'data-fid': 'copy-' + r.key }, COPY.reports.copyLink);
    btn.addEventListener('click', async () => {
      if (await copyText(r.url)) copiedLabel(btn, COPY.reports.copyLink, COPY.reports.copied);
    });
    li.append(btn);
  } else {
    li.append(h('span'));
  }
  return li;
}

// ---- Device ----------------------------------------------------------

function renderDevice(s) {
  const c = s.connect;
  const connecting = !s.paired && (c.state === 'waiting' || c.state === 'blocked');
  const data = {
    paired: s.paired, who: s.account_label, signedOut: s.signed_out ? s.signed_out.account_label || true : null, backend: s.token_backend, backendName: s.token_backend_name,
    api: s.api_base_url, name: s.device_name, site: s.site_base_url,
    c: { state: c.state, code: c.user_code, url: c.url, trouble: c.trouble, message: c.message, expires: c.expires_at },
    l: { open: local.pairOpen, confirm: local.confirmForget, forgot: local.forgot, flash: local.flash, copied: local.copied, busy: local.busyPair, starting: local.starting },
  };
  region('device', $('device-panel'), data, (el) => {
    el.replaceChildren();
    if (s.paired) el.append(...pairedView(s));
    else if (connecting) el.append(...signingInView(c));
    else el.append(...notConnectedView(s, c));
  });
  tickCountdown();
}

function notConnectedView(s, c) {
  const failed = c.state === 'failed';
  const body = [];
  if (failed && c.message) body.push(h('div', { class: 'reason', role: 'alert', text: c.message }));
  if (s.signed_out) body.push(h('div', { class: 'reason', role: 'alert', text: COPY.device.signedOutBar(s.signed_out.account_label) }));
  body.push(h('p', { class: 'lede', text: COPY.device.intro }));
  const btn = h('button', { class: 'account', type: 'button', id: 'connect', 'data-fid': 'connect',
    onclick: startConnect });
  btn.append(local.starting ? COPY.device.starting : (failed ? COPY.device.tryAgain : COPY.device.signIn));
  if (local.starting) btn.disabled = true;
  body.push(h('div', { class: 'cta' }, btn));
  body.push(h('p', { class: 'fine', text: COPY.device.signInNote }));

  const open = local.pairOpen || failed;
  const fb = h('div', { class: 'fallback' });
  if (!open) {
    fb.append(h('button', { class: 'link', type: 'button', 'data-fid': 'fallback-toggle',
      'aria-expanded': 'false', onclick: () => { local.pairOpen = true; renderDevice(snap); const i = $('pair-code'); if (i) i.focus(); } },
    COPY.device.fallbackToggle));
  } else {
    fb.append(pairForm());
  }
  body.push(fb);
  body.push(h('p', { class: 'note', id: 'device-note', role: 'status' }));
  return [
    h('div', { class: 'state-line' }, dot('bad'), COPY.device.notConnected),
    h('div', { class: 'device-body' }, ...body),
  ];
}

function pairForm() {
  const input = h('input', { id: 'pair-code', class: 'mono', 'data-keep': 'pair-code', 'data-fid': 'pair-code', autocomplete: 'off',
    spellcheck: 'false', placeholder: COPY.device.fallbackPlaceholder, 'aria-describedby': 'pair-help' });
  const submit = h('button', { class: 'action', type: 'submit', id: 'pair', 'data-fid': 'pair' },
    local.busyPair ? COPY.device.fallbackBusy : COPY.device.fallbackButton);
  if (local.busyPair) submit.disabled = true;
  return h('form', { id: 'pair-form', onsubmit: onPair },
    h('label', { for: 'pair-code', text: COPY.device.fallbackLabel }),
    input,
    h('p', { class: 'help', id: 'pair-help', text: COPY.device.fallbackHelp }),
    submit);
}

function signingInView(c) {
  const body = [
    // The browser-blocked state has its own sentence below, in place of this one.
    c.state === 'blocked' ? null : h('p', { class: 'lede', text: COPY.device.opened }),
    h('p', { class: 'lede', text: COPY.device.checkCode }),
    h('div', { class: 'code-plate' }, h('span', { class: 'code', text: c.user_code })),
    h('p', { class: 'fine expires', id: 'countdown', 'data-expires': c.expires_at || '' }),
    h('div', { class: 'expiry', 'aria-hidden': 'true' }, h('i', { id: 'expiry-fill' })),
  ];
  if (c.trouble) body.push(h('p', { class: 'note warn', role: 'status', text: COPY.device.trouble }));
  if (c.state === 'blocked') {
    body.push(h('p', { class: 'lede', text: COPY.device.blocked }));
    body.push(h('div', { class: 'link-field' },
      h('input', { class: 'mono', readonly: true, value: c.url, 'aria-label': COPY.device.linkLabel, onfocus: (e) => e.target.select() })));
  }
  const buttons = [];
  if (c.state === 'blocked') {
    buttons.push(h('button', { class: 'action primary', type: 'button', 'data-fid': 'copylink', onclick: async (e) => {
      const btn = e.currentTarget;
      if (await copyText(c.url)) copiedLabel(btn, COPY.device.copyLink, COPY.device.copied);
    } }, COPY.device.copyLink));
  } else {
    buttons.push(h('button', { class: 'action', type: 'button', 'data-fid': 'reopen', onclick: onReopen }, COPY.device.reopen));
  }
  buttons.push(h('button', { class: 'ghost', type: 'button', 'data-fid': 'cancel', onclick: onCancel }, COPY.device.cancel));
  body.push(h('div', { class: 'buttons' }, ...buttons));
  body.push(h('p', { class: 'note', id: 'device-note', role: 'status' }));
  return [
    h('div', { class: 'state-line' }, h('span', { class: 'spinner', 'aria-hidden': 'true' }),
      h('span', { role: 'status', text: COPY.device.waiting })),
    h('div', { class: 'device-body' }, ...body),
  ];
}

function pairedView(s) {
  const body = [];
  if (local.flash) {
    // The check draws once, on the render that first has the flash; a redraw
    // while the four seconds run shows it settled.
    const first = !local.flashDrawn;
    local.flashDrawn = true;
    body.push(h('p', { class: 'connected-flash' + (first ? ' draw' : ''), role: 'status' }, icon('check', 14), COPY.device.flash));
    if (first) replay($('device-panel'), 'flash', 1200);
  }
  // Where the key lives, in the player's words, with one muted line that says what
  // the store is. Never a warning colour: the file fallback still uploads.
  body.push(h('div', { class: 'row device-saved' }, h('span', { class: 'k', text: COPY.device.savedIn }),
    h('span', { class: 'v', text: s.token_backend_name },
      s.token_backend_note ? h('span', { class: 'sub', text: s.token_backend_note }) : null)));
  body.push(h('details', {}, h('summary', {}, icon('chevron', 14), COPY.device.details),
    h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.device.deviceName }), h('span', { class: 'v mono', text: s.device_name })),
    h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.device.api }), h('span', { class: 'v mono', text: s.api_base_url }))));
  if (local.confirmForget) {
    body.push(h('div', { class: 'ask', role: 'alertdialog', 'aria-label': COPY.device.forgetAsk },
      h('p', { text: COPY.device.forgetAsk }),
      h('div', { class: 'ask-buttons' },
        h('button', { class: 'action danger', type: 'button', id: 'forget-yes', 'data-fid': 'forget-yes', onclick: onForget }, COPY.device.forgetYes),
        h('button', { class: 'action', type: 'button', id: 'forget-no', 'data-fid': 'forget-no', onclick: () => cancelForget(true) }, COPY.device.forgetNo))));
  } else {
    body.push(h('div', { class: 'buttons' },
      h('button', { class: 'action', type: 'button', 'data-fid': 'my-reports', onclick: () => openLink(s.site_base_url + '/logs') }, COPY.device.openReports),
      h('button', { class: 'ghost', type: 'button', 'data-fid': 'forget', onclick: askForget }, COPY.device.forget)));
  }
  body.push(h('p', { class: 'note', id: 'device-note', role: 'status' }));
  return [
    h('div', { class: 'state-line' }, dot('on'), COPY.device.connectedAs(s.account_label)),
    h('div', { class: 'device-body' }, ...body),
  ];
}

function renderForgotten(s, revoked) {
  // After forgetting the page is the not-connected view, with one line
  // that says what did and did not happen: removed from the account, or
  // forgotten here only because the site could not be reached.
  const n = $('device-note');
  if (!n) return;
  n.className = 'note';
  if (revoked) { n.replaceChildren(COPY.device.forgottenRemoved); return; }
  n.replaceChildren(COPY.device.forgottenLocal, ext(COPY.device.forgottenLink, s.site_base_url + '/account'), COPY.device.forgottenLocalEnd);
}

function askForget() {
  local.confirmForget = true;
  renderDevice(snap);
  const no = $('forget-no');
  if (no) no.focus();
}

function cancelForget(refocus) {
  local.confirmForget = false;
  renderDevice(snap);
  if (refocus) { const b = document.querySelector('[data-fid="forget"]'); if (b) b.focus(); }
}

document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape' && local.confirmForget) { e.preventDefault(); cancelForget(true); }
});

async function onForget(e) {
  const r = await busy(e.target, COPY.device.forgetting, () => api('api/unpair', {}));
  local.confirmForget = false;
  if (!r.ok) { applySnapshot(snap); note('device-note', 'bad', r.error.message); return; }
  local.forgot = true;
  local.pairOpen = false;
  applySnapshot(r.data);
  renderForgotten(r.data, r.data.revoked);
}

async function startConnect() {
  local.starting = true;
  local.forgot = false;
  renderDevice(snap);
  const r = await api('api/connect', {});
  local.starting = false;
  if (!r.ok) { renderDevice(snap); note('device-note', 'bad', r.error.message); return; }
  applySnapshot(r.data);
}

async function onReopen(e) {
  const r = await busy(e.target, '...', () => api('api/connect/reopen', {}));
  if (!r.ok) { note('device-note', 'bad', r.error.message); return; }
  applySnapshot(r.data);
}

async function onCancel() {
  const r = await api('api/connect/cancel', {});
  if (!r.ok) { note('device-note', 'bad', r.error.message); return; }
  applySnapshot(r.data);
  const b = document.querySelector('[data-fid="connect"]');
  if (b) b.focus();
}

async function onPair(e) {
  e.preventDefault();
  const input = $('pair-code');
  const code = input.value.trim();
  if (!code) { note('device-note', 'bad', COPY.device.emptyCode); input.focus(); return; }
  local.busyPair = true;
  renderDevice(snap);
  const r = await api('api/pair', { code });
  local.busyPair = false;
  if (!r.ok) {
    renderDevice(snap);
    note('device-note', 'bad', r.error.message);
    const i = $('pair-code');
    if (i) i.focus();
    return;
  }
  const i = $('pair-code');
  if (i) i.value = '';
  local.pairOpen = false;
  applySnapshot(r.data);
}

// A countdown that ticks without re-rendering anything. The expiry bar
// under it is the share of the ten minutes left; its width eases over the
// second until the next tick, so it moves without a redraw.
function tickCountdown() {
  const el = $('countdown');
  if (!el || !el.dataset.expires) return;
  const now = new Date();
  const t = FORMAT.countdown(el.dataset.expires, now);
  el.textContent = COPY.device.countdown(t.m, t.s);
  el.classList.toggle('late', FORMAT.lastMinute(el.dataset.expires, now));
  const fill = $('expiry-fill');
  if (fill) fill.style.width = (FORMAT.expiryShare(el.dataset.expires, now) * 100) + '%';
}
setInterval(tickCountdown, 1000);

// ---- Settings --------------------------------------------------------

let baseline = null; // the saved values the form was last filled with; null before the first fill
let saving = false;

function charValue(s) {
  return s.logging_character ? FORMAT.characterValue(s.logging_character) : '';
}

function formValues() {
  return { visibility: $('visibility').value, character: $('logging-character').value };
}

function dirty() {
  if (baseline === null) return false;
  const v = formValues();
  return v.visibility !== baseline.visibility || v.character !== baseline.character;
}

function syncSave() {
  const strip = FORMAT.saveStrip(dirty(), saving, COPY.settings);
  $('save').disabled = !strip.enabled;
  $('save-state').textContent = strip.text;
  $('visibility-help').textContent = COPY.settings.visibilityHelp[$('visibility').value] || '';
}

// fillSettings puts the saved values into the form: on first load, and
// again when the player leaves the page with changes unsaved.
function fillSettings(s) {
  baseline = { visibility: s.visibility, character: charValue(s) };
  $('visibility').value = s.visibility;
  const select = $('logging-character');
  select.replaceChildren(h('option', { value: '', text: FORMAT.automaticLabel(s.detected_character, COPY.settings) }));
  for (const c of s.characters) {
    select.append(h('option', { value: FORMAT.characterValue(c), text: FORMAT.characterLabel(c, COPY.settings.rulesets) }));
  }
  select.value = baseline.character;
  if (select.value !== baseline.character) select.value = '';
  syncSave();
}

function resetSettings() {
  if (snap && dirty()) fillSettings(snap);
}

function renderSettings(s) {
  // The form is rewritten from the snapshot only when nothing has been
  // changed: unsaved edits are never overwritten by the poll.
  const formKey = { v: s.visibility, c: s.characters, l: charValue(s), d: s.detected_character };
  if (!dirty() && JSON.stringify(formKey) !== cache['settings-form']) {
    cache['settings-form'] = JSON.stringify(formKey);
    fillSettings(s);
  }

  region('folders', $('folder-list'), { installs: s.installs, pending: s.watch_pending, budget: FORMAT.pathBudget(window.innerWidth) }, (el) => {
    el.replaceChildren();
    if (!s.installs.length) {
      el.append(h('li', {}, h('span', { class: 'note', text: COPY.settings.noFolders })));
      return;
    }
    for (const i of s.installs) {
      // The name and when it was last played share a line, so the three
      // fields and Save stay together on a short window.
      const where = h('div', { class: 'where' },
        h('div', { class: 'head' }, h('span', { class: 'name', text: i.flavor_name || i.flavor }),
          h('span', { class: 'played', text: FORMAT.lastPlayed(i.last_played, new Date(), COPY.status) })),
        h('span', { class: 'path', title: i.path, text: FORMAT.pathTruncate(i.path, FORMAT.pathBudget(window.innerWidth)) }));
      const side = i.added
        ? h('button', { class: 'ghost small', type: 'button', 'data-fid': 'remove-' + i.path, onclick: (e) => onRemove(e.target, i.path) }, COPY.settings.remove)
        : h('span', { class: 'auto', text: COPY.settings.foundAutomatically });
      el.append(h('li', {}, where, side));
    }
  });
}

// A narrower window shortens the Settings paths; the poll would not notice a resize.
window.addEventListener('resize', () => { if (snap) { renderSettings(snap); renderStatus(snap); } });
$('visibility').addEventListener('change', syncSave);
$('logging-character').addEventListener('change', syncSave);

$('add-folder').addEventListener('click', async (e) => {
  note('folder-note', 'ok', '');
  const r = await busy(e.target, COPY.settings.adding, () => api('api/folders/pick', {}));
  if (!r.ok) { note('folder-note', 'bad', r.error.message); return; }
  const before = snap ? snap.installs.length : 0;
  applySnapshot(r.data);
  if (r.data.installs.length > before) note('folder-note', 'ok', r.data.watch_pending ? COPY.settings.folderAddedPending : COPY.settings.folderAdded);
});

async function onRemove(btn, path) {
  note('folder-note', 'ok', '');
  const r = await busy(btn, COPY.settings.removing, () => api('api/folders/remove', { path }));
  if (!r.ok) { note('folder-note', 'bad', r.error.message); return; }
  applySnapshot(r.data);
  note('folder-note', 'ok', COPY.settings.folderRemoved);
}

$('save').addEventListener('click', async () => {
  note('folder-note', 'ok', '');
  const v = formValues();
  const visibilityChanged = v.visibility !== baseline.visibility;
  const picked = v.character;
  const [region, ruleset, name] = picked ? picked.split('|') : [];
  saving = true;
  syncSave();
  $('save').textContent = COPY.settings.saving;
  const r = await api('api/settings', {
    visibility: v.visibility,
    logging_character: picked ? { region, ruleset, name } : null,
  });
  saving = false;
  $('save').textContent = COPY.settings.save;
  if (!r.ok) { syncSave(); note('settings-note', 'bad', r.error.message); return; }
  note('settings-note', 'ok', visibilityChanged
    ? COPY.settings.savedVisibility(COPY.settings.visibilitySaved[v.visibility])
    : COPY.settings.saved);
  baseline = { visibility: v.visibility, character: picked };
  applySnapshot(r.data);
  syncSave();
});

// ---- chrome: footer, bars, title -------------------------------------

// updatedSince is when the window first showed "Updated to ..." while it
// was on screen; the line stays for UPDATED_MS from then.
let updatedSince = null;
function updatedShowing(s) {
  if (!s.updated_to) return false;
  if (updatedSince === null && document.visibilityState === 'visible') updatedSince = Date.now();
  return updatedSince !== null && Date.now() - updatedSince < UPDATED_MS;
}

function renderChrome(s) {
  const updated = updatedShowing(s);
  region('footer', $('footer'), { v: s.app_version, u: s.update_staged, site: s.site_base_url, upd: updated ? s.updated_to : '' }, () => {
    // The version stays on screen after an update; only "just updated" is green.
    const left = $('foot-version');
    left.replaceChildren(COPY.footer.version(s.app_version));
    if (updated) left.append(COPY.footer.joint, h('span', { class: 'updated', text: COPY.footer.justUpdated }));
    $('foot-site').textContent = COPY.footer.site;
    $('foot-site').setAttribute('href', s.site_base_url);
    $('wordmark').setAttribute('href', s.site_base_url);
    $('update-bar').hidden = !s.update_staged;
    $('update-text').textContent = s.update_staged ? COPY.footer.updateReady : '';
  });
}

// The first close of the window: the shell calls this before it hides
// the window, or sets __trayNotice when it had to destroy it.
window.trayNotice = (text) => {
  $('tray-bar-text').textContent = text || COPY.trayNotice;
  $('tray-bar').hidden = false;
};
$('tray-bar-dismiss').addEventListener('click', () => { $('tray-bar').hidden = true; });
$('tray-bar-dismiss').textContent = COPY.dismiss;
if (window.__trayNotice) window.trayNotice(window.__trayNotice);

// The labels that never change are written once.
document.title = COPY.title;
$('folders-title').textContent = COPY.status.folders;
$('reports-all').textContent = COPY.reports.all;
$('visibility-label').textContent = COPY.settings.visibility;
for (const o of $('visibility').options) o.textContent = COPY.settings.visibilityOptions[o.value];
$('folders-label').textContent = COPY.settings.folders;
$('character-label').textContent = COPY.settings.character;
$('character-help').textContent = COPY.settings.characterHelp;
$('save').textContent = COPY.settings.save;
$('add-folder').textContent = COPY.settings.addFolder;
$('first-check').textContent = COPY.checking;
$('reports-hint').textContent = COPY.reports.notCreditedHint;

// ---- the poll --------------------------------------------------------

let wasPaired = null;

function applySnapshot(s) {
  const justPaired = wasPaired === false && s.paired;
  snap = s;
  wasPaired = s.paired;
  if (page === null) {
    // First paint: an unpaired companion opens on Device.
    $('first-check').hidden = true;
    $('main').hidden = false;
    replay($('main'), 'enter', ENTER_MS);
    show(s.paired ? 'status' : 'device', false);
  }
  renderChrome(s);
  renderStatus(s);
  renderReports(s);
  renderDevice(s);
  renderSettings(s);
  if (local.forgot && !s.paired) renderForgotten(s);
  if (justPaired) onConnected();
}

// Connecting is the one moment the companion changes pages for you:
// the Device page says so for four seconds, then Status shows the
// result. Away from the Device page, nothing moves.
function onConnected() {
  local.flash = true;
  local.flashDrawn = false;
  local.pairOpen = false;
  local.forgot = false;
  if (page === 'device') {
    renderDevice(snap);
    setTimeout(() => {
      local.flash = false;
      if (page === 'device') show('status', true);
      renderDevice(snap);
    }, FLASH_MS);
  } else {
    local.flash = false;
  }
}

async function refresh() {
  const r = await api('api/status');
  if (!r.ok) {
    showBar('poll-bar', COPY.localDown);
    $('first-check').hidden = true;
    // The cards keep the last numbers the companion sent, dimmed, under a
    // headline that says so.
    if (!local.pollFailed) {
      local.pollFailed = true;
      if (snap) { renderStatus(snap); renderReports(snap); }
    }
    return;
  }
  $('poll-bar').hidden = true;
  local.pollFailed = false;
  applySnapshot(r.data);
}

// If the first answer is slow the bar says so; there is no spinner on the page.
setTimeout(() => { if (!snap) { showBar('poll-bar', COPY.localDown); $('first-check').hidden = true; } }, SLOW_FIRST_POLL_MS);

// ---- native feel: no browser menu, no page zoom, nothing to drag -----

// The window is an app, not a web page: the browser's right-click menu,
// Ctrl-wheel and Ctrl-plus zoom and dragging a link or the seal are off.
// A text field or a selection keeps its menu so Copy and Paste still work.
document.addEventListener('contextmenu', (e) => {
  const t = e.target;
  const editable = t && t.closest && t.closest('input, select, textarea');
  if (!editable && !String(window.getSelection()).length) e.preventDefault();
});
document.addEventListener('keydown', (e) => {
  if ((e.ctrlKey || e.metaKey) && ['+', '-', '=', '0'].includes(e.key)) e.preventDefault();
});
document.addEventListener('wheel', (e) => { if (e.ctrlKey) e.preventDefault(); }, { passive: false });
document.addEventListener('dragstart', (e) => e.preventDefault());

refresh();
setInterval(refresh, POLL_MS);
