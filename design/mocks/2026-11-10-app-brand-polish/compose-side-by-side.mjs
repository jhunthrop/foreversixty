// design/mocks/2026-11-10-app-brand-polish/compose-side-by-side.mjs
// Lays each app tab root's capture beside the website's page head at 390 px (capture-web.mjs):
// website | iOS before | iOS after | Android after, one sheet per tab, at one height. The iOS
// "before" is the restyle lane's capture (design/reviews/2026-11-10-restyle-app).
// Run from the repo root: node design/mocks/2026-11-10-app-brand-polish/compose-side-by-side.mjs
import { existsSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const repo = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..', '..', '..');
const sharp = createRequire(path.join(repo, 'web', 'package.json'))('sharp');
const DIR = path.join(repo, 'design/reviews/2026-11-10-app-brand-polish');
const BEFORE = path.join(repo, 'design/reviews/2026-11-10-restyle-app');
const HEIGHT = 1400;
const GAP = 40;
const LABEL = 60;
/** Each tab with the web page whose head it echoes (Home's hero, a titled page's band). */
const SHEETS = { home: 'home', raid: 'bis', gear: 'bis', activity: 'guides' };

async function column(file, label) {
  const img = await sharp(file).resize({ height: HEIGHT }).png().toBuffer();
  const { width } = await sharp(img).metadata();
  return { img, width, label };
}

for (const [tab, web] of Object.entries(SHEETS)) {
  const sources = [
    [path.join(DIR, `web-${web}-390.png`), `Website /${web === 'home' ? '' : web} at 390`],
    [path.join(BEFORE, `${tab}-ios.png`), 'iOS before'],
    [path.join(DIR, `fade-${tab}-dark-ios.png`), 'iOS after'],
    [path.join(DIR, `fade-${tab}-dark-android.png`), 'Android after'],
  ].filter(([file]) => existsSync(file));
  const cols = await Promise.all(sources.map(([file, label]) => column(file, label)));
  const width = cols.reduce((sum, c) => sum + c.width + GAP, GAP);
  let x = GAP;
  const parts = [];
  const labels = [];
  for (const c of cols) {
    parts.push({ input: c.img, left: x, top: LABEL });
    labels.push(`<text x='${x}' y='40' font-family='Helvetica Neue, Arial' font-size='28' fill='#e8e2d6'>${c.label}</text>`);
    x += c.width + GAP;
  }
  const frame = `<svg xmlns='http://www.w3.org/2000/svg' width='${width}' height='${HEIGHT + LABEL + GAP}'><rect width='100%' height='100%' fill='#16130f'/>${labels.join('')}</svg>`;
  await sharp(Buffer.from(frame)).composite(parts).png().toFile(path.join(DIR, `side-by-side-${tab}.png`));
}
console.log(`wrote the side-by-side sheets to ${path.relative(repo, DIR)}`);
