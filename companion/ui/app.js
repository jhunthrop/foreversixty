// companion/ui/app.js
// The whole window. It polls /api/status every two seconds and posts
// the three things the player can change. There is no framework and
// no build step: the page is four panels over one JSON document.
'use strict';

const api = (path, body) =>
  fetch(path, body === undefined
    ? { cache: 'no-store' }
    : { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(body) })
    .then((r) => r.json());

const $ = (id) => document.getElementById(id);
const text = (id, value) => { $(id).textContent = value; };

let editing = false;

function show(page) {
  for (const s of document.querySelectorAll('main section')) s.hidden = s.id !== page;
  for (const b of document.querySelectorAll('nav button')) {
    if (b.dataset.page === page) b.setAttribute('aria-current', 'page');
    else b.removeAttribute('aria-current');
  }
}

for (const b of document.querySelectorAll('nav button')) {
  b.addEventListener('click', () => show(b.dataset.page));
}

function renderStatus(s) {
  const p = s.pipeline;
  $('logging-dot').className = 'dot ' + (p.logging ? 'on' : 'off');
  text('logging-state', p.logging ? 'Logging' : 'Waiting for the game');
  text('logging-report', p.report_id || (p.report_key ? 'not uploaded yet' : '—'));
  text('logging-fights', p.fights);
  text('logging-queued', p.queued);
  text('engine-version', s.engine_version);

  const note = $('logging-note');
  if (!s.paired) {
    note.className = 'note bad';
    note.textContent = 'This device is not paired, so nothing is uploading. Open the Device page.';
  } else if (!s.installs.length) {
    note.className = 'note bad';
    note.textContent = 'No World of Warcraft folder was found. Add one on the Settings page.';
  } else {
    note.className = 'note';
    note.textContent = p.log_path ? 'Reading ' + p.log_path : 'Nothing is being written to the log yet.';
  }

  const installs = $('installs');
  installs.replaceChildren();
  if (!s.installs.length) {
    const empty = document.createElement('p');
    empty.className = 'note';
    empty.textContent = 'None found.';
    installs.append(empty);
  }
  for (const i of s.installs) {
    const row = document.createElement('div');
    row.className = 'row';
    const k = document.createElement('span');
    k.className = 'k';
    k.textContent = i.path;
    const v = document.createElement('span');
    v.className = 'v';
    if (!i.advanced_logging_known) v.textContent = 'advanced logging: unknown';
    else v.textContent = i.advanced_logging ? 'advanced logging: on' : 'advanced logging: OFF';
    row.append(k, v);
    installs.append(row);
    if (i.advanced_logging_known && !i.advanced_logging) {
      const help = document.createElement('p');
      help.className = 'note warn';
      help.textContent = 'Turn it on in ' + i.advanced_logging_help + '.';
      installs.append(help);
    }
  }
}

function renderReports(s) {
  const list = $('report-list');
  list.replaceChildren();
  if (!s.reports.length) {
    const li = document.createElement('li');
    li.className = 'note';
    li.textContent = 'No reports yet. They appear here the moment a fight is uploaded.';
    list.append(li);
    return;
  }
  for (const r of s.reports) {
    const li = document.createElement('li');
    const when = document.createElement('div');
    when.className = 'when';
    when.textContent = new Date(r.started_at).toLocaleString() +
      ' · ' + r.fights + (r.fights === 1 ? ' fight' : ' fights') +
      (r.done ? ' · complete' : ' · in progress');
    const title = document.createElement('div');
    if (r.url) {
      const a = document.createElement('a');
      a.href = r.url;
      a.textContent = r.zone || r.report_id;
      a.target = '_blank';
      a.rel = 'noreferrer';
      title.append(a);
    } else {
      title.textContent = (r.zone || r.key) + ' — not uploaded yet';
    }
    li.append(title, when);
    list.append(li);
  }
}

function renderDevice(s) {
  $('paired-dot').className = 'dot ' + (s.paired ? 'on' : 'bad');
  text('paired-state', s.paired ? 'Paired' : 'Not paired');
  text('token-backend', s.token_backend);
  text('api-base', s.api_base_url);
}

function renderSettings(s) {
  if (editing) return;
  $('visibility').value = s.visibility;
  $('wow-paths').value = s.installs.map((i) => i.path).join('\n');
  const select = $('logging-character');
  const chosen = s.logging_character
    ? [s.logging_character.region, s.logging_character.ruleset, s.logging_character.name].join('|')
    : '';
  select.replaceChildren();
  const none = document.createElement('option');
  none.value = '';
  none.textContent = 'Not set';
  select.append(none);
  for (const c of s.characters) {
    const o = document.createElement('option');
    o.value = [c.region, c.ruleset, c.name].join('|');
    o.textContent = c.name + ' · ' + c.ruleset + ' · ' + c.region;
    select.append(o);
  }
  select.value = chosen;
}

function render(s) {
  renderStatus(s);
  renderReports(s);
  renderDevice(s);
  renderSettings(s);
}

async function refresh() {
  try {
    const r = await api('api/status');
    if (r.ok) render(r.data);
  } catch (e) {
    const note = $('logging-note');
    note.className = 'note bad';
    note.textContent = 'The companion stopped answering: ' + e;
  }
}

function report(id, r) {
  const note = $(id);
  note.className = r.ok ? 'note' : 'note bad';
  note.textContent = r.ok ? 'Saved.' : r.error.message;
  if (r.ok) render(r.data);
}

$('pair').addEventListener('click', async () => {
  report('device-note', await api('api/pair', { code: $('pair-code').value.trim() }));
});

$('unpair').addEventListener('click', async () => {
  report('device-note', await api('api/unpair', {}));
});

for (const id of ['visibility', 'wow-paths', 'logging-character']) {
  $(id).addEventListener('input', () => { editing = true; });
}

$('save').addEventListener('click', async () => {
  const picked = $('logging-character').value;
  const [region, ruleset, name] = picked ? picked.split('|') : [];
  const r = await api('api/settings', {
    visibility: $('visibility').value,
    wow_paths: $('wow-paths').value.split('\n').map((s) => s.trim()).filter(Boolean),
    logging_character: picked ? { region, ruleset, name } : null,
  });
  editing = false;
  report('settings-note', r);
});

refresh();
setInterval(refresh, 2000);
