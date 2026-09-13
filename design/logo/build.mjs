// Builds Forever Sixty logo SVGs from the real Cinzel outlines.
import { readFileSync, writeFileSync, mkdirSync } from 'node:fs';
import opentype from 'opentype.js';

const FONT_DIR = '/Users/jh/code/forever/.worktrees/phase-0-web/web/node_modules/@fontsource/cinzel/files/';
const OUT = '/Users/jh/code/forever/design/logo/';
mkdirSync(OUT, { recursive: true });

const toArrayBuffer = (buf) => buf.buffer.slice(buf.byteOffset, buf.byteOffset + buf.byteLength);
const bold = opentype.parse(toArrayBuffer(readFileSync(FONT_DIR + 'cinzel-latin-700-normal.woff')));
const black = opentype.parse(toArrayBuffer(readFileSync(FONT_DIR + 'cinzel-latin-800-normal.woff')));

const GOLD = { light: '#fbe7a1', mid: '#e5b955', deep: '#a8762a', edge: '#7a5420' };
const INK = '#f2eee4';
const BG = '#07090d';

function textPath(font, text, size, tracking = 0) {
  // tracking is in em units added after every glyph except the last
  const glyphs = font.stringToGlyphs(text);
  const scale = size / font.unitsPerEm;
  let x = 0;
  let d = '';
  glyphs.forEach((g, i) => {
    const p = g.getPath(x, 0, size);
    d += p.toPathData(2);
    x += g.advanceWidth * scale;
    if (i < glyphs.length - 1) {
      const next = glyphs[i + 1];
      x += (font.getKerningValue(g, next) || 0) * scale;
      x += tracking * size;
    }
  });
  const bb = opentype.Path.prototype.getBoundingBox.call({ commands: opentype.Path.prototype.commands, ...new opentype.Path() }, undefined) && null;
  return { d, advance: x };
}

function bbox(d) {
  const p = new opentype.Path();
  // cheap bbox from the path data numbers
  const nums = d.match(/-?\d+(?:\.\d+)?/g).map(Number);
  let minX = Infinity, minY = Infinity, maxX = -Infinity, maxY = -Infinity;
  for (let i = 0; i + 1 < nums.length; i += 2) {
    minX = Math.min(minX, nums[i]); maxX = Math.max(maxX, nums[i]);
    minY = Math.min(minY, nums[i + 1]); maxY = Math.max(maxY, nums[i + 1]);
  }
  void p;
  return { minX, minY, maxX, maxY, w: maxX - minX, h: maxY - minY };
}

const defs = (id) => `
  <defs>
    <linearGradient id="${id}" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0" stop-color="${GOLD.light}"/>
      <stop offset="0.45" stop-color="${GOLD.mid}"/>
      <stop offset="1" stop-color="${GOLD.deep}"/>
    </linearGradient>
  </defs>`;

function svg({ w, h, body, bg = null }) {
  return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${w} ${h}" width="${w}" height="${h}" role="img" aria-label="Forever Sixty">${bg ? `<rect width="${w}" height="${h}" fill="${bg}"/>` : ''}${body}</svg>\n`;
}

// ---- Mark: "60" in Cinzel Black, gold gradient with a thin bronze edge ----
const MARK_SIZE = 200;
const mark = textPath(black, '60', MARK_SIZE, -0.02);
const mb = bbox(mark.d);
const pad = 12;
const markW = Math.ceil(mb.w + pad * 2);
const markH = Math.ceil(mb.h + pad * 2);
const markTransform = `translate(${(pad - mb.minX).toFixed(2)} ${(pad - mb.minY).toFixed(2)})`;

const markGold = svg({
  w: markW, h: markH,
  body: `${defs('g')}<g transform="${markTransform}"><path d="${mark.d}" fill="url(#g)" stroke="${GOLD.edge}" stroke-width="1.5" stroke-linejoin="round" paint-order="stroke"/></g>`,
});
const markMono = svg({
  w: markW, h: markH,
  body: `<g transform="${markTransform}"><path d="${mark.d}" fill="${INK}"/></g>`,
});
const markMonoOnDark = svg({ w: markW, h: markH, bg: BG, body: `<g transform="${markTransform}"><path d="${mark.d}" fill="${INK}"/></g>` });

// ---- Wordmark: "FOREVER SIXTY" in Cinzel Bold, tracked, gold gradient ----
const WM_SIZE = 96;
const wm = textPath(bold, 'FOREVER SIXTY', WM_SIZE, 0.08);
const wb = bbox(wm.d);
const wmW = Math.ceil(wb.w + pad * 2);
const wmH = Math.ceil(wb.h + pad * 2);
const wmTransform = `translate(${(pad - wb.minX).toFixed(2)} ${(pad - wb.minY).toFixed(2)})`;
const wordmarkGold = svg({ w: wmW, h: wmH, body: `${defs('g')}<g transform="${wmTransform}"><path d="${wm.d}" fill="url(#g)"/></g>` });
const wordmarkMono = svg({ w: wmW, h: wmH, body: `<g transform="${wmTransform}"><path d="${wm.d}" fill="${INK}"/></g>` });

// ---- Horizontal lockup: mark 1.5x the wordmark cap height, vertically centered on the caps ----
const capScale = (wb.h * 1.5) / mb.h;
const gap = Math.round(wb.h * 0.5);
const markScaledH = mb.h * capScale;
const lockW = Math.ceil(mb.w * capScale + gap + wb.w + pad * 2);
const lockH = Math.ceil(markScaledH + pad * 2);
const wmY = pad + (markScaledH - wb.h) / 2;
const horizontal = svg({
  w: lockW, h: lockH,
  body: `${defs('g')}
  <g transform="translate(${pad} ${pad}) scale(${capScale.toFixed(4)}) translate(${(-mb.minX).toFixed(2)} ${(-mb.minY).toFixed(2)})"><path d="${mark.d}" fill="url(#g)" stroke="${GOLD.edge}" stroke-width="${(1.5 / capScale).toFixed(2)}" stroke-linejoin="round" paint-order="stroke"/></g>
  <g transform="translate(${(pad + mb.w * capScale + gap - wb.minX).toFixed(2)} ${(wmY - wb.minY).toFixed(2)})"><path d="${wm.d}" fill="url(#g)"/></g>`,
});

// ---- Stacked lockup: mark centered above wordmark ----
const stackMarkScale = (wb.w * 0.28) / mb.w;
const stackGap = Math.round(wb.h * 0.6);
const stackW = wmW;
const stackH = Math.ceil(mb.h * stackMarkScale + stackGap + wb.h + pad * 2);
const stacked = svg({
  w: stackW, h: stackH,
  body: `${defs('g')}
  <g transform="translate(${((stackW - mb.w * stackMarkScale) / 2).toFixed(2)} ${pad}) scale(${stackMarkScale.toFixed(4)}) translate(${(-mb.minX).toFixed(2)} ${(-mb.minY).toFixed(2)})"><path d="${mark.d}" fill="url(#g)" stroke="${GOLD.edge}" stroke-width="${(1.5 / stackMarkScale).toFixed(2)}" stroke-linejoin="round" paint-order="stroke"/></g>
  <g transform="translate(${(pad - wb.minX).toFixed(2)} ${(pad + mb.h * stackMarkScale + stackGap - wb.minY).toFixed(2)})"><path d="${wm.d}" fill="url(#g)"/></g>`,
});

// ---- Favicon: 32x32, off-white 60 on near-black, rounded square ----
const favScale = 24 / Math.max(mb.w, mb.h);
const favW = mb.w * favScale, favH = mb.h * favScale;
const favicon = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="32" height="32"><rect width="32" height="32" rx="6" fill="${BG}"/><g transform="translate(${((32 - favW) / 2).toFixed(2)} ${((32 - favH) / 2).toFixed(2)}) scale(${favScale.toFixed(4)}) translate(${(-mb.minX).toFixed(2)} ${(-mb.minY).toFixed(2)})"><path d="${mark.d}" fill="${INK}"/></g></svg>\n`;

// ---- Contact sheet for review ----
const sheetW = 1600, sheetH = 900;
const sheet = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${sheetW} ${sheetH}" width="${sheetW}" height="${sheetH}">
  <rect width="${sheetW}" height="${sheetH}" fill="${BG}"/>
  ${defs('g')}
  <g transform="translate(80 80)"><g transform="scale(${(560 / lockW).toFixed(4)})">${horizontal.replace(/^<svg[^>]*>|<\/svg>\n$/g, '')}</g></g>
  <g transform="translate(900 60)"><g transform="scale(${(560 / stackW).toFixed(4)})">${stacked.replace(/^<svg[^>]*>|<\/svg>\n$/g, '')}</g></g>
  <g transform="translate(80 380)"><g transform="scale(${(720 / wmW).toFixed(4)})">${wordmarkGold.replace(/^<svg[^>]*>|<\/svg>\n$/g, '')}</g></g>
  <g transform="translate(1040 360)"><g transform="scale(${(220 / markW).toFixed(4)})">${markGold.replace(/^<svg[^>]*>|<\/svg>\n$/g, '')}</g></g>
  <rect x="80" y="620" width="420" height="220" fill="#000"/>
  <g transform="translate(180 650)"><g transform="scale(${(160 / markH).toFixed(4)})">${markMono.replace(/^<svg[^>]*>|<\/svg>\n$/g, '')}</g></g>
  <rect x="540" y="620" width="420" height="220" fill="#f2eee4"/>
  <g transform="translate(640 650)"><g transform="scale(${(160 / markH).toFixed(4)})">${markMono.replace(/^<svg[^>]*>|<\/svg>\n$/g, '').replace(INK, BG)}</g></g>
  <g transform="translate(1040 640)">${favicon.replace(/^<svg[^>]*>|<\/svg>\n$/g, '').replace('<rect', '<rect transform="scale(4)"').replace('<g transform="', '<g transform="scale(4) ')}</g>
  <g transform="translate(1200 640)">${favicon.replace(/^<svg[^>]*>|<\/svg>\n$/g, '')}</g>
  <text x="80" y="600" fill="#9a9484" font-family="Barlow, Helvetica, Arial" font-size="18">Built from Cinzel Black (mark) and Cinzel Bold (wordmark) outlines. Left to right: horizontal lockup, stacked lockup, wordmark, mark, mono on black, mono on off-white, favicon at 4x and 1x.</text>
</svg>
`;

writeFileSync(OUT + 'foreversixty-mark.svg', markGold);
writeFileSync(OUT + 'foreversixty-mark-mono.svg', markMono);
writeFileSync(OUT + 'foreversixty-mark-mono-on-dark.svg', markMonoOnDark);
writeFileSync(OUT + 'foreversixty-wordmark.svg', wordmarkGold);
writeFileSync(OUT + 'foreversixty-wordmark-mono.svg', wordmarkMono);
writeFileSync(OUT + 'foreversixty-lockup-horizontal.svg', horizontal);
writeFileSync(OUT + 'foreversixty-lockup-stacked.svg', stacked);
writeFileSync(OUT + 'favicon.svg', favicon);
writeFileSync(OUT + 'contact-sheet.svg', sheet);
console.log({ markW, markH, wmW, wmH, lockW, lockH, stackW, stackH });
