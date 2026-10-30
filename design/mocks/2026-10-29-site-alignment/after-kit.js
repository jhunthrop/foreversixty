// design/mocks/2026-10-29-site-alignment/after-kit.js
// Runs inside the built page (page.addScriptTag) to draw the after-state the site-alignment spec
// asks for: the account pattern's Card (card-surface + CardHead + full-bleed body), the PageHead
// band, and the new glyphs. A mock, not product code: it restyles the real fixture page so the
// content, data and copy stay the site's own. Every value here is a token named in the spec.
(() => {
  const GLYPHS = {
    scale: '<path d="M12 3v18M7 21h10M5 7h14"/><path d="M5 7l-3 6a3 3 0 006 0L5 7zM19 7l-3 6a3 3 0 006 0l-3-6z"/>',
    bolt: '<path d="M13 2L4 14h7l-1 8 9-12h-7l1-8z"/>',
    list: '<path d="M9 6h11M9 12h11M9 18h11"/><path d="M4 6h.01M4 12h.01M4 18h.01"/>',
    pin: '<path d="M12 21s7-6.2 7-11.5A7 7 0 005 9.5C5 14.8 12 21 12 21z"/><circle cx="12" cy="9.5" r="2.5"/>',
    sword: '<path d="M14.5 17.5L3 6V3h3l11.5 11.5"/><path d="M13 19l6-6M16 16l4 4M19 21l2-2"/>',
    tree: '<circle cx="12" cy="5" r="2"/><circle cx="6" cy="19" r="2"/><circle cx="18" cy="19" r="2"/><path d="M12 7v5M12 12l-6 5M12 12l6 5"/>',
    book: '<path d="M4 5a2 2 0 012-2h13v16H6a2 2 0 00-2 2V5z"/><path d="M4 19a2 2 0 012-2h13"/>',
    upload: '<path d="M12 16V4M7 9l5-5 5 5"/><path d="M4 16v3a1 1 0 001 1h14a1 1 0 001-1v-3"/>',
    download: '<path d="M12 4v12M7 11l5 5 5-5"/><path d="M4 16v3a1 1 0 001 1h14a1 1 0 001-1v-3"/>',
    share: '<circle cx="6" cy="12" r="2.5"/><circle cx="18" cy="6" r="2.5"/><circle cx="18" cy="18" r="2.5"/><path d="M8.2 10.8l7.6-3.6M8.2 13.2l7.6 3.6"/>',
    layers: '<path d="M12 3l9 5-9 5-9-5 9-5z"/><path d="M3 13l9 5 9-5"/>',
    target: '<circle cx="12" cy="12" r="8"/><circle cx="12" cy="12" r="4"/><path d="M12 12h.01"/>',
    trophy: '<path d="M8 4h8v5a4 4 0 01-8 0V4z"/><path d="M8 6H4v1a4 4 0 004 4M16 6h4v1a4 4 0 01-4 4M12 13v4M8 21h8"/>',
    search: '<circle cx="11" cy="11" r="7"/><path d="M20 20l-4-4"/>',
    news: '<rect x="3" y="4" width="18" height="16" rx="2"/><path d="M7 8h10M7 12h10M7 16h6"/>',
    star: '<path d="M12 3.5l2.6 5.4 5.9.8-4.3 4.1 1 5.9L12 16.9l-5.2 2.8 1-5.9-4.3-4.1 5.9-.8L12 3.5z"/>',
    chart: '<path d="M4 4v16h16"/><path d="M8 15l4-5 3 3 5-7"/>',
    user: '<circle cx="12" cy="8" r="4"/><path d="M4 21a8 8 0 0116 0"/>',
    chevron: '<path d="M9 6l6 6-6 6"/>',
  };
  const svg = (name, size) =>
    `<svg width="${size}" height="${size}" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">${GLYPHS[name]}</svg>`;

  const style = document.createElement('style');
  style.textContent = `
    .x-card { background: var(--surface-card) !important; border: 1px solid var(--color-line-card) !important;
      border-radius: 12px !important; box-shadow: inset 0 1px 0 var(--card-highlight), var(--shadow-card) !important;
      padding: 0 !important; overflow: hidden; display: flex !important; flex-direction: column !important; gap: 0 !important; }
    .x-card.x-quiet { box-shadow: var(--shadow-card) !important; }
    .x-head { display: flex; align-items: center; gap: 12px; padding: 16px 20px; min-height: 68px; box-sizing: border-box; }
    .x-tile { width: 36px; height: 36px; flex: none; border-radius: 8px; background: var(--tile); color: var(--tile-ink);
      display: inline-flex; align-items: center; justify-content: center; }
    .x-tile.x-sm { width: 32px; height: 32px; }
    .x-title { font-family: var(--font-heading); font-weight: 800; font-size: 24px; line-height: 24px; letter-spacing: .01em;
      text-transform: uppercase; color: var(--color-strong, #f2eee4); margin: 0; }
    .x-micro { font: 700 11px/1 var(--font-body); letter-spacing: .14em; text-transform: uppercase; color: var(--color-muted, #9a9484); }
    .x-aside { margin-left: auto; font: 400 13px/18px var(--font-body); color: var(--color-muted, #9a9484); text-align: right; }
    .x-action { margin-left: auto; display: inline-flex; align-items: center; gap: 4px; height: 44px; font: 600 14px/1 var(--font-body);
      color: var(--gold, #e5b955); text-decoration: none; letter-spacing: 0; text-transform: none; }
    .x-body { border-top: 1px solid #1c2230; padding: 16px 20px 20px; display: flex; flex-direction: column; gap: 12px; }
    .x-body.x-bleed { padding: 0; gap: 0; }
    .x-row { display: flex; align-items: center; gap: 14px; min-height: 56px; padding: 10px 20px; border-top: 1px solid #1c2230;
      color: var(--color-text, #e9e4d8); text-decoration: none; box-sizing: border-box; }
    .x-row:first-child { border-top: 0; }
    .x-row .x-label { font: 600 15px/20px var(--font-body); color: #f2eee4; }
    .x-row .x-sub { font: 400 13px/18px var(--font-body); color: #9a9484; }
    .x-row .x-chev { margin-left: auto; color: #9a9484; }
    .x-band { background: var(--surface-band) !important; }
    .x-band::before, .x-band::after { display: none !important; }
    .x-hide { display: none !important; }
  `;
  document.head.appendChild(style);

  const norm = (s) => (s || '').replace(/\s+/g, ' ').trim().toLowerCase();
  /** The smallest element under `root` whose whole text is `text`. */
  function byText(text, root = document.querySelector('main'), sel = 'h1,h2,h3,h4,p,span,div,a,button,li') {
    const want = norm(text);
    const hits = [...root.querySelectorAll(sel)].filter((el) => norm(el.textContent) === want && el.offsetParent !== null);
    return hits.sort((a, b) => a.querySelectorAll('*').length - b.querySelectorAll('*').length)[0] ?? null;
  }
  /** The nearest bordered, rounded ancestor: the panel that holds `el`. */
  function boxOf(el) {
    for (let n = el; n && n !== document.body; n = n.parentElement) {
      const s = getComputedStyle(n);
      if (parseFloat(s.borderTopWidth) >= 1 && parseFloat(s.borderLeftWidth) >= 1 && parseFloat(s.borderTopLeftRadius) >= 4 && n.getBoundingClientRect().width > 200) return n;
    }
    return null;
  }
  function head({ glyph, title, label, aside, action, tone }) {
    const h = document.createElement('div');
    h.className = 'x-head';
    const tile = glyph ? `<span class="x-tile">${svg(glyph, 18)}</span>` : '';
    const lead = tone === 'image' ? '' : tile;
    const words = label
      ? `<div style="display:flex;flex-direction:column;gap:4px;min-width:0"><span class="x-micro">${label}</span><h2 class="x-title" style="text-transform:none">${title}</h2></div>`
      : `<h2 class="x-title">${title}</h2>`;
    const right = action
      ? `<a class="x-action" href="#">${action}${svg('chevron', 16)}</a>`
      : aside ? `<span class="x-aside">${aside}</span>` : '';
    h.innerHTML = lead + words + right;
    return h;
  }
  /** Turns `box` into the account Card: surface, a CardHead, the rest as its body. */
  function cardify(box, opts) {
    if (!box) return null;
    (opts.drop || []).forEach((el) => el && el.classList.add('x-hide'));
    if (opts.dropFirst && box.firstElementChild) box.firstElementChild.classList.add('x-hide');
    box.classList.add('x-card');
    if (opts.quiet) box.classList.add('x-quiet');
    const body = document.createElement('div');
    body.className = opts.bleed ? 'x-body x-bleed' : 'x-body';
    if (opts.inset) body.style.padding = '0 20px 8px';
    while (box.firstChild) body.appendChild(box.firstChild);
    box.appendChild(head(opts));
    box.appendChild(body);
    return box;
  }
  /** Wraps loose siblings (a bare section: its title and its content) into one new card. */
  function wrap(nodes, opts) {
    const box = document.createElement('section');
    nodes[0].parentElement.insertBefore(box, nodes[0]);
    nodes.forEach((n) => box.appendChild(n));
    return cardify(box, opts);
  }
  function band(el) { if (el) el.classList.add('x-band'); }
  function rows(container, items) {
    container.innerHTML = items
      .map((it) => `<a class="x-row" href="#"><span class="x-tile x-sm">${svg(it.glyph, 16)}</span><span style="display:flex;flex-direction:column;gap:2px"><span class="x-label">${it.label}</span>${it.sub ? `<span class="x-sub">${it.sub}</span>` : ''}</span><span class="x-chev">${svg('chevron', 16)}</span></a>`)
      .join('');
  }
  /** The CharacterHeader on the card surface: the line-card border, not the gold one. */
  function surfaceOf(el, quiet = false) {
    if (!el) return;
    el.style.setProperty('background', 'var(--surface-card)', 'important');
    el.style.setProperty('border', '1px solid var(--color-line-card)', 'important');
    el.style.setProperty('border-radius', '12px', 'important');
    el.style.setProperty('box-shadow', quiet ? 'var(--shadow-card)' : 'inset 0 1px 0 var(--card-highlight), var(--shadow-card)', 'important');
  }
  window.__align = { byText, boxOf, cardify, wrap, band, rows, surfaceOf, head, svg };
})();
