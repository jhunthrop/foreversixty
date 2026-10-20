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
  el.textContent = text || '';
  if (text && kind !== 'bad') timers.set(id, setTimeout(() => { el.textContent = ''; }, NOTE_MS));
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
  const t = bar.querySelector('span') || bar;
  t.textContent = text;
  clearTimeout(barTimer);
  if (ms) barTimer = setTimeout(() => { bar.hidden = true; }, ms);
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
const local = { pairOpen: false, confirmForget: false, forgot: false, flash: false, copied: false, busyPair: false, starting: false };

function show(name, focus) {
  const leaving = page;
  page = name;
  for (const s of document.querySelectorAll('main section')) s.hidden = s.id !== name;
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

function renderStatus(s) {
  const p = s.pipeline;
  const hd = s.headline;
  const reportURL = p.report_id ? s.site_base_url + '/reports/' + p.report_id : '';

  region('headline', $('headline'), { hd, reportURL }, (el) => {
    el.replaceChildren();
    el.append(
      h('div', { class: 'head' }, dot(hd.tone), h('div', { class: 'title', text: hd.title })),
      h('p', { class: 'sub', text: hd.sub }),
    );
    const actions = [];
    if (hd.action) {
      actions.push(h('button', { class: 'action primary', type: 'button', 'data-fid': 'headline-action',
        onclick: () => show(hd.action.page, true) }, hd.action.label));
    }
    if (hd.key === 'logging' && reportURL) actions.push(ext(COPY.status.reportLogging, reportURL, { 'data-fid': 'headline-report' }));
    if (actions.length) el.append(h('div', { class: 'actions' }, ...actions));
  });

  const rows = {
    report: p.report_id ? { zone: ((s.reports || []).find((r) => r.key === p.report_key) || {}).zone || '' } : (p.report_key ? 'starting' : 'none'),
    fights: p.fights, queued: p.queued, addon: s.addon_line, stale: s.addon_stale, never: s.addon_never, help: s.addon_help, title: s.addon_seen_title,
  };
  region('status-rows', $('status-rows'), rows, (el) => {
    el.replaceChildren();
    let report;
    if (rows.report === 'none') report = h('span', { class: 'v', text: COPY.status.noReport });
    else if (rows.report === 'starting') report = h('span', { class: 'v', text: COPY.status.starting });
    else report = h('span', { class: 'v', text: rows.report.zone || COPY.status.inProgress });
    el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.report }), report));
    el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.fights }),
      h('span', { class: 'v mono', text: String(p.fights) })));
    if (p.queued > 0) {
      el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.queued }),
        h('span', { class: 'v mono', text: String(p.queued) })));
    }
    const v = h('span', { class: 'v' + (s.addon_stale ? ' amber' : '') + (s.addon_never ? ' muted' : ''), title: s.addon_seen_title || null, text: s.addon_line });
    if (s.addon_stale && s.addon_help) v.append(h('span', { class: 'sub', text: s.addon_help }));
    el.append(h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.status.rows.addon }), v));
  });

  region('installs', $('installs'), { installs: s.installs, key: hd.key }, (el) => {
    el.replaceChildren();
    if (!s.installs.length) {
      el.append(h('p', { class: 'note', text: COPY.status.noFolders }));
      return;
    }
    for (const i of s.installs) el.append(installRow(i, hd.key));
  });
}

function installRow(i, headlineKey) {
  const top = h('div', { class: 'top-line' },
    h('div', { class: 'name' }, i.flavor_name || i.flavor,
      i.watching ? h('span', { class: 'tag', title: COPY.status.watchingTitle, text: COPY.status.watching }) : null),
    advancedLabel(i));
  const row = h('div', { class: 'install' }, top,
    h('span', { class: 'played', text: FORMAT.lastPlayed(i.last_played, new Date(), COPY.status) }),
    h('span', { class: 'path', title: i.path, text: FORMAT.midTruncate(i.path, 62) }));
  // The instruction lives in the headline when the headline is about
  // advanced logging; a row only keeps its short line when it is not.
  if (i.advanced_logging_known && !i.advanced_logging && headlineKey !== 'advanced_off') {
    row.append(h('p', { class: 'help-line', text: i.advanced_logging_help }));
  }
  return row;
}

function advancedLabel(i) {
  if (!i.advanced_logging_known) {
    return h('span', { class: 'adv', text: COPY.status.advancedUnknown });
  }
  if (i.advanced_logging) {
    return h('span', { class: 'adv' }, h('span', { class: 'glyph', 'aria-hidden': 'true', text: '✓' }), COPY.status.advancedOn);
  }
  return h('span', { class: 'adv off' }, h('span', { class: 'glyph', 'aria-hidden': 'true', text: '⚠' }), COPY.status.advancedOff);
}

// ---- Reports ---------------------------------------------------------

function renderReports(s) {
  const all = $('reports-all');
  if (all.getAttribute('href') !== s.site_base_url + '/logs') all.setAttribute('href', s.site_base_url + '/logs');

  region('reports', $('report-panel'), { reports: s.reports, paired: s.paired, site: s.site_base_url }, (el) => {
    el.replaceChildren();
    if (!s.reports.length) {
      const box = h('div', { class: 'empty' });
      if (!s.paired) {
        box.append(h('p', { class: 'big-line', text: COPY.reports.unpairedTitle }),
          h('button', { class: 'action primary', type: 'button', onclick: () => show('device', true) }, COPY.reports.unpairedButton));
      } else {
        box.append(h('p', { class: 'big-line', text: COPY.reports.emptyTitle }),
          h('p', { class: 'note', text: COPY.reports.emptyBody }));
      }
      el.append(box);
      return;
    }
    const list = h('ul', { class: 'reports' });
    const now = new Date();
    for (const r of s.reports.slice(0, REPORT_ROWS)) list.append(reportRow(r, now));
    el.append(list);
  });
}

function reportRow(r, now) {
  const zone = r.zone || COPY.reports.unnamed;
  const title = r.url ? ext(zone, r.url, { 'data-fid': 'report-' + r.key }) : h('span', { text: zone });
  const meta = h('div', { class: 'when' },
    r.done ? null : dot('on'),
    [FORMAT.when(r.started_at, now, COPY.reports), COPY.reports.fights(r.fights)]
      .concat(r.done ? [] : [COPY.reports.inProgress])
      .concat(r.url ? [] : [COPY.reports.notUploaded]).join(' · '));
  const li = h('li', {}, h('div', { class: 'what' }, h('div', { class: 'zone' }, title), meta));
  if (r.url) {
    const btn = h('button', { class: 'ghost small', type: 'button', 'data-fid': 'copy-' + r.key }, COPY.reports.copyLink);
    btn.addEventListener('click', async () => {
      if (await copyText(r.url)) {
        btn.textContent = COPY.reports.copied;
        setTimeout(() => { btn.textContent = COPY.reports.copyLink; }, 2000);
      }
    });
    li.append(btn);
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
  const out = [
    h('div', { class: 'state-line' }, dot('bad'), COPY.device.notConnected),
  ];
  if (failed && c.message) out.push(h('div', { class: 'reason', role: 'alert', text: c.message }));
  if (s.signed_out) out.push(h('div', { class: 'reason', role: 'alert', text: COPY.device.signedOutBar(s.signed_out.account_label) }));
  out.push(h('p', { class: 'lede', text: COPY.device.intro }));
  const btn = h('button', { class: 'action primary big', type: 'button', id: 'connect', 'data-fid': 'connect',
    onclick: startConnect });
  btn.append(local.starting ? COPY.device.starting : (failed ? COPY.device.tryAgain : COPY.device.signIn));
  if (local.starting) btn.disabled = true;
  out.push(h('div', {}, btn));
  out.push(h('p', { class: 'fine', text: COPY.device.signInNote }));

  const open = local.pairOpen || failed;
  const fb = h('div', { class: 'fallback' });
  if (!open) {
    fb.append(h('button', { class: 'link', type: 'button', 'data-fid': 'fallback-toggle',
      'aria-expanded': 'false', onclick: () => { local.pairOpen = true; renderDevice(snap); const i = $('pair-code'); if (i) i.focus(); } },
    COPY.device.fallbackToggle));
  } else {
    fb.append(pairForm());
  }
  out.push(fb);
  out.push(h('p', { class: 'note', id: 'device-note', role: 'status' }));
  return out;
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
  const out = [
    h('div', { class: 'state-line' }, h('span', { class: 'spinner', 'aria-hidden': 'true' }),
      h('span', { role: 'status', text: COPY.device.waiting })),
    h('p', { class: 'lede' }, COPY.device.checkCode, h('br'), h('span', { class: 'code', text: c.user_code })),
    h('p', { class: 'fine' }, COPY.device.lasts + ' ', h('span', { id: 'countdown', 'data-expires': c.expires_at || '' })),
  ];
  if (c.trouble) out.push(h('p', { class: 'note warn', role: 'status', text: COPY.device.trouble }));
  if (c.state === 'blocked') {
    out.push(h('p', { class: 'lede', text: COPY.device.blocked }));
    out.push(h('div', { class: 'link-field' },
      h('input', { class: 'mono', readonly: true, value: c.url, 'aria-label': COPY.device.linkLabel, onfocus: (e) => e.target.select() })));
  }
  const buttons = [];
  if (c.state === 'blocked') {
    buttons.push(h('button', { class: 'action primary', type: 'button', 'data-fid': 'copylink', onclick: async (e) => {
      if (await copyText(c.url)) {
        e.target.textContent = COPY.device.copied;
        setTimeout(() => { e.target.textContent = COPY.device.copyLink; }, 2000);
      }
    } }, COPY.device.copyLink));
  } else {
    buttons.push(h('button', { class: 'action', type: 'button', 'data-fid': 'reopen', onclick: onReopen }, COPY.device.reopen));
  }
  buttons.push(h('button', { class: 'ghost', type: 'button', 'data-fid': 'cancel', onclick: onCancel }, COPY.device.cancel));
  out.push(h('div', { class: 'buttons' }, ...buttons));
  out.push(h('p', { class: 'note', id: 'device-note', role: 'status' }));
  return out;
}

function pairedView(s) {
  const out = [
    h('div', { class: 'state-line' }, dot('on'), COPY.device.connectedAs(s.account_label)),
  ];
  if (local.flash) out.push(h('p', { class: 'note ok', role: 'status', text: COPY.device.flash }));
  const file = s.token_backend === 'file';
  out.push(h('div', { class: 'row', style: 'margin-top:12px' }, h('span', { class: 'k', text: COPY.device.savedIn }),
    h('span', { class: 'v' + (file ? ' amber' : ''), text: s.token_backend_name },
      file && s.token_backend_note ? h('span', { class: 'sub', text: s.token_backend_note }) : null)));
  out.push(h('details', {}, h('summary', { text: COPY.device.details }),
    h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.device.deviceName }), h('span', { class: 'v mono', text: s.device_name })),
    h('div', { class: 'row' }, h('span', { class: 'k', text: COPY.device.api }), h('span', { class: 'v mono', text: s.api_base_url }))));
  if (local.confirmForget) {
    out.push(h('div', { class: 'ask', role: 'alertdialog', 'aria-label': COPY.device.forgetAsk },
      h('p', { text: COPY.device.forgetAsk }),
      h('div', { class: 'ask-buttons' },
        h('button', { class: 'action danger', type: 'button', id: 'forget-yes', 'data-fid': 'forget-yes', onclick: onForget }, COPY.device.forgetYes),
        h('button', { class: 'action', type: 'button', id: 'forget-no', 'data-fid': 'forget-no', onclick: () => cancelForget(true) }, COPY.device.forgetNo))));
  } else {
    out.push(h('div', { class: 'buttons' },
      h('button', { class: 'action', type: 'button', 'data-fid': 'my-reports', onclick: () => openLink(s.site_base_url + '/logs') }, COPY.device.openReports),
      h('button', { class: 'ghost', type: 'button', 'data-fid': 'forget', onclick: askForget }, COPY.device.forget)));
  }
  out.push(h('p', { class: 'note', id: 'device-note', role: 'status' }));
  return out;
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

// A countdown that ticks without re-rendering anything.
function tickCountdown() {
  const el = $('countdown');
  if (!el || !el.dataset.expires) return;
  const t = FORMAT.countdown(el.dataset.expires, new Date());
  el.textContent = COPY.device.countdown(t.m, t.s);
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
  $('save').disabled = saving || !dirty();
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

  region('folders', $('folder-list'), { installs: s.installs, pending: s.watch_pending }, (el) => {
    el.replaceChildren();
    if (!s.installs.length) {
      el.append(h('li', {}, h('span', { class: 'note', text: COPY.settings.noFolders })));
      return;
    }
    for (const i of s.installs) {
      const where = h('div', { class: 'where' },
        h('div', { class: 'name', text: i.flavor_name || i.flavor }),
        h('span', { class: 'played', text: FORMAT.lastPlayed(i.last_played, new Date(), COPY.status) }),
        h('span', { class: 'path', title: i.path, text: FORMAT.midTruncate(i.path, 52) }));
      const side = i.added
        ? h('button', { class: 'ghost small', type: 'button', 'data-fid': 'remove-' + i.path, onclick: (e) => onRemove(e.target, i.path) }, COPY.settings.remove)
        : h('span', { class: 'auto', text: COPY.settings.foundAutomatically });
      el.append(h('li', {}, where, side));
    }
  });
}

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

function renderChrome(s) {
  region('footer', $('footer'), { v: s.app_version, u: s.update_staged, site: s.site_base_url }, () => {
    $('foot-version').textContent = COPY.footer.version(s.app_version);
    $('foot-site').textContent = COPY.footer.site;
    $('foot-site').setAttribute('href', s.site_base_url);
    $('wordmark').setAttribute('href', s.site_base_url);
    $('foot-update').hidden = !s.update_staged;
    $('foot-update').textContent = s.update_staged ? COPY.footer.updateReady : '';
  });
  $('footer').hidden = false;
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

// ---- the poll --------------------------------------------------------

let wasPaired = null;

function applySnapshot(s) {
  const justPaired = wasPaired === false && s.paired;
  snap = s;
  wasPaired = s.paired;
  if (page === null) {
    // First paint: an unpaired companion opens on Device.
    $('main').hidden = false;
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
    return;
  }
  $('poll-bar').hidden = true;
  applySnapshot(r.data);
}

refresh();
setInterval(refresh, POLL_MS);
