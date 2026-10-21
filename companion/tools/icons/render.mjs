// companion/tools/icons/render.mjs
// Renders the Forever Sixty seal (web/public/favicon.svg) into the companion's
// icon assets: internal/icon/assets/seal-{16,24,32,48,64,128,256,512}.png,
// seal.ico (16-256, PNG entries) and seal.icns (16-512 with @2x).
//
// Regenerate (from the repo root, after `pnpm install` in web/):
//     node companion/tools/icons/render.mjs
// If web/node_modules lives elsewhere (a git worktree, say):
//     WEB_NODE_MODULES=/path/to/web/node_modules node companion/tools/icons/render.mjs
//
// The seal is the logo and is not redrawn here. The one change: at 32 px and
// below the ring's 36 tiny dots are sub-pixel noise, so those sizes render
// the seal without them.
import { createRequire } from 'node:module';
import { readFileSync, writeFileSync, mkdirSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const here = dirname(fileURLToPath(import.meta.url));
const repo = resolve(here, '../../..');
const nm = process.env.WEB_NODE_MODULES || join(repo, 'web/node_modules');
const require = createRequire(join(nm, 'noop.js'));
const { Resvg } = require('@resvg/resvg-js');

const svg = readFileSync(join(repo, 'web/public/favicon.svg'), 'utf8');
const calm = svg.replace(/<circle [^>]*r="0\.26"[^>]*\/>/g, '');
if (calm === svg) throw new Error('favicon.svg no longer has ring dots to drop; check the regex');

const png = (size) =>
  new Resvg(size <= 32 ? calm : svg, { fitTo: { mode: 'width', value: size } }).render().asPng();

const out = join(here, '../../internal/icon/assets');
mkdirSync(out, { recursive: true });

const sizes = [16, 24, 32, 48, 64, 128, 256, 512];
const cache = new Map();
const get = (s) => (cache.has(s) ? cache.get(s) : (cache.set(s, png(s)), cache.get(s)));
for (const s of sizes) writeFileSync(join(out, `seal-${s}.png`), get(s));

// ICO: ICONDIR, one ICONDIRENTRY per size, then the PNG blobs. Windows
// reads PNG entries at any size since Vista.
const icoSizes = [16, 24, 32, 48, 64, 128, 256];
const head = Buffer.alloc(6 + 16 * icoSizes.length);
head.writeUInt16LE(0, 0); head.writeUInt16LE(1, 2); head.writeUInt16LE(icoSizes.length, 4);
let offset = head.length;
icoSizes.forEach((s, i) => {
  const b = get(s), o = 6 + 16 * i;
  head[o] = s === 256 ? 0 : s; head[o + 1] = s === 256 ? 0 : s; head[o + 2] = 0; head[o + 3] = 0;
  head.writeUInt16LE(1, o + 4); head.writeUInt16LE(32, o + 6);
  head.writeUInt32LE(b.length, o + 8); head.writeUInt32LE(offset, o + 12);
  offset += b.length;
});
writeFileSync(join(out, 'seal.ico'), Buffer.concat([head, ...icoSizes.map(get)]));

// ICNS: 'icns', total length, then (type, length, PNG) elements, big-endian.
const icnsTypes = [
  ['icp4', 16], ['icp5', 32], ['icp6', 64], ['ic07', 128], ['ic08', 256], ['ic09', 512],
  ['ic10', 1024], // 512@2x
  ['ic11', 32],   // 16@2x
  ['ic12', 64],   // 32@2x
  ['ic13', 256],  // 128@2x
  ['ic14', 512],  // 256@2x
];
const elems = icnsTypes.map(([t, s]) => {
  const b = get(s), h = Buffer.alloc(8);
  h.write(t, 0, 'ascii'); h.writeUInt32BE(8 + b.length, 4);
  return Buffer.concat([h, b]);
});
const total = 8 + elems.reduce((n, e) => n + e.length, 0);
const ih = Buffer.alloc(8); ih.write('icns', 0, 'ascii'); ih.writeUInt32BE(total, 4);
writeFileSync(join(out, 'seal.icns'), Buffer.concat([ih, ...elems]));
console.log(`wrote ${sizes.length} PNGs, seal.ico, seal.icns to ${out}`);

// Tray icons that carry the headline's tone: the seal plus a status dot at
// its lower right (docs: companion premium spec, section 10.3). The dot has
// a ring in the page colour so it reads on a light taskbar and a dark one.
// "off" is the plain seal above and has no file of its own. 48 px is not in
// the spec's table; it scales the 32 px dot by 1.5.
const tones = { on: '#7bff5c', warn: '#e8883a', bad: '#ff6b5c' };
const dots = {
  16: { c: 12.5, r: 3, ring: 1 },
  24: { c: 18.5, r: 4.5, ring: 1.5 },
  32: { c: 25, r: 6, ring: 2 },
  48: { c: 37.5, r: 9, ring: 3 },
};
const withDot = (size, color) => {
  const d = dots[size];
  const base = get(size).toString('base64');
  const svg = `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="${size}" height="${size}" viewBox="0 0 ${size} ${size}">` +
    `<image x="0" y="0" width="${size}" height="${size}" xlink:href="data:image/png;base64,${base}"/>` +
    `<circle cx="${d.c}" cy="${d.c}" r="${d.r + d.ring}" fill="#07090d"/>` +
    `<circle cx="${d.c}" cy="${d.c}" r="${d.r}" fill="${color}"/></svg>`;
  return new Resvg(svg, { fitTo: { mode: 'width', value: size } }).render().asPng();
};
const trayCache = new Map();
const tray = (size, tone) => {
  const k = `${tone}-${size}`;
  if (!trayCache.has(k)) trayCache.set(k, withDot(size, tones[tone]));
  return trayCache.get(k);
};
const traySizes = [16, 24, 32, 48];
for (const tone of Object.keys(tones)) {
  const th = Buffer.alloc(6 + 16 * traySizes.length);
  th.writeUInt16LE(0, 0); th.writeUInt16LE(1, 2); th.writeUInt16LE(traySizes.length, 4);
  let off = th.length;
  traySizes.forEach((s, i) => {
    const b = tray(s, tone), o = 6 + 16 * i;
    th[o] = s; th[o + 1] = s; th[o + 2] = 0; th[o + 3] = 0;
    th.writeUInt16LE(1, o + 4); th.writeUInt16LE(32, o + 6);
    th.writeUInt32LE(b.length, o + 8); th.writeUInt32LE(off, o + 12);
    off += b.length;
  });
  writeFileSync(join(out, `seal-tray-${tone}.ico`), Buffer.concat([th, ...traySizes.map((s) => tray(s, tone))]));
  writeFileSync(join(out, `seal-tray-${tone}.png`), tray(32, tone));
}
console.log(`wrote seal-tray-{${Object.keys(tones).join(',')}}.ico and .png to ${out}`);
