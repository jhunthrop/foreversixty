// design/listing/build-store.mjs
// Builds every store listing image from the brand sources and the app captures, so none is placed
// by hand:
//   icons/listing-icon-400.png             CurseForge and Wago project icon and account avatar
//   app-store/icon-1024.png                App Store icon: the painted tile, opaque
//   app-store/screenshots/NN-<capture>.jpg iPhone 6.9" screenshots, 1320x2868 (--only app-store)
//   play/icon-512.png                      Play icon: the painted tile, opaque
//   play/feature-graphic.png               Play feature graphic, 1024x500, opaque
//   play/screenshots/NN-<capture>.jpg      phone screenshots, 1080x1920 (9:16) (--only play)
// Each screenshot is a plate in the site's page head: the band art fading into the ground, the
// caption in Marcellus, and the whole screen below it, never enlarged. Captions come from
// store-screens.json; the captures from capture-apps.sh (app-captures/<platform>/<capture>.png).
// Run from anywhere: node design/listing/build-store.mjs
// sharp and @playwright/test resolve from web/node_modules.
import { mkdirSync, readFileSync, rmSync, statSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { PNG, renderTile } from '../../web/scripts/brand-art.mjs';
import { INK } from './brand-ink.mjs';

const here = path.dirname(fileURLToPath(import.meta.url));
const repo = path.resolve(here, '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const sharp = require('sharp');
const { chromium } = require('@playwright/test');

const { screens } = JSON.parse(readFileSync(path.join(here, 'store-screens.json'), 'utf8'));

/** The feature graphic's line: what the app does, in the store's own words. */
const FEATURE_LINE = 'Raid sign-ups, the roster and your next upgrade';
const FAN_LINE = 'Fan-made player tool for WoW: Forever';

/** Each store's screenshot: its required pixel size, drawn as a CSS layout at `scale`.
 *  App Store: the 6.9" iPhone slot (1320x2868), which the store scales to every smaller iPhone.
 *  Play: 9:16 at 1080x1920, the size Play asks for to feature an app. */
const STORES = {
  'app-store': { platform: 'ios', width: 440, height: 956, scale: 3 },
  play: { platform: 'android', width: 360, height: 640, scale: 3 },
};
/** Upload ceilings: App Store screenshots and Play images both well under 8 MB; Play's icon 1 MB. */
const MAX_BYTES = { screenshot: 8 * 1024 * 1024, playIcon: 1024 * 1024, feature: 15 * 1024 * 1024 };
const JPEG = { quality: 92, mozjpeg: true, chromaSubsampling: '4:4:4' };
const FEATURE = { width: 1024, height: 500, scale: 2 };
const ICON = { listing: 400, appStore: 1024, play: 512 };

const dataUri = (file, type) => `data:${type};base64,${readFileSync(file).toString('base64')}`;
const font = (name) => dataUri(path.join(repo, 'apps/shared/fonts', name), 'font/ttf');
const FONTS = `
@font-face{font-family:Marcellus;src:url(${font('Marcellus-Regular.ttf')})}
@font-face{font-family:'Open Sans';src:url(${font('OpenSans-Regular.ttf')})}`;
const art = (name) => dataUri(path.join(repo, 'web/public/art', name), 'image/webp');

function out(rel) {
  const file = path.join(here, rel);
  mkdirSync(path.dirname(file), { recursive: true });
  return file;
}

function report(rel, limit) {
  const bytes = statSync(path.join(here, rel)).size;
  if (bytes > limit) throw new Error(`${rel} is ${bytes} bytes, over the store's ${limit}`);
  console.log(`${rel.padEnd(44)} ${String(Math.round(bytes / 1024)).padStart(5)} KB  ok`);
}

async function buildIcons() {
  await sharp(await renderTile(ICON.listing)).png(PNG).toFile(out('icons/listing-icon-400.png'));
  await sharp(await renderTile(ICON.appStore)).removeAlpha().png(PNG).toFile(out('app-store/icon-1024.png'));
  await sharp(await renderTile(ICON.play)).removeAlpha().png(PNG).toFile(out('play/icon-512.png'));
  report('icons/listing-icon-400.png', MAX_BYTES.playIcon);
  report('app-store/icon-1024.png', MAX_BYTES.screenshot);
  report('play/icon-512.png', MAX_BYTES.playIcon);
}

/** The plate's layout as shares of its height: the caption block, and the margin under the screen. */
const PLATE = { caption: 0.15, foot: 0.035 };

/** The screen's box on a plate: the whole capture (status bar to home indicator), fitted under the
 *  caption and centred, so nothing is cut through. */
function screenBox(store, capture) {
  const top = Math.round(store.height * PLATE.caption);
  const height = store.height - top - Math.round(store.height * PLATE.foot);
  return { top, height, width: Math.round((height * capture.width) / capture.height) };
}

/** A plate: the band over the ground, the caption, and the whole screen under it with a phone
 *  screen's rounded corners. */
function plateHtml(store, caption, capture, box) {
  const { width, height } = store;
  return `<!doctype html><meta charset="utf-8"><style>${FONTS}
html,body{margin:0;width:${width}px;height:${height}px;overflow:hidden;background:${INK.bg}}
.p{position:relative;width:100%;height:100%;background:url(${art('ground.webp')}) 0 0/1600px auto repeat,${INK.bg}}
.band{position:absolute;left:0;right:0;top:0;height:${Math.round(height * 0.4)}px;
background:url(${art('band-1200.webp')}) center top/cover;
-webkit-mask-image:linear-gradient(#000 55%,transparent);mask-image:linear-gradient(#000 55%,transparent)}
h1{position:absolute;left:0;right:0;top:0;height:${box.top}px;margin:0;padding:0 ${Math.round(width * 0.07)}px;box-sizing:border-box;
display:flex;align-items:center;justify-content:center;font:400 ${Math.round(width * 0.075)}px/1.15 Marcellus,serif;
color:${INK.cream};text-align:center;text-wrap:balance}
.s{position:absolute;top:${box.top}px;left:${Math.round((width - box.width) / 2)}px;width:${box.width}px;height:${box.height}px;
overflow:hidden;border-radius:${Math.round(box.width * 0.11)}px;border:1px solid ${INK.line}}
.s img{display:block;width:100%;height:100%}
</style><div class="p"><div class="band"></div><h1>${caption}</h1><div class="s"><img src="${capture.uri}"></div></div>`;
}

/** The page `html` at `size`, as an opaque PNG at the page's device scale. */
async function render(page, html, size) {
  await page.setViewportSize({ width: size.width, height: size.height });
  await page.setContent(html, { waitUntil: 'load' });
  await page.evaluate(() => document.fonts.ready);
  return sharp(await page.screenshot({ type: 'png' })).flatten({ background: INK.bg }).png().toBuffer();
}

/** The capture, with its size; fails when the plate would enlarge it. */
async function loadCapture(file, store) {
  const { width, height } = await sharp(file).metadata();
  const capture = { width, height, uri: dataUri(file, 'image/png') };
  const box = screenBox(store, capture);
  if (width < box.width * store.scale) {
    throw new Error(`${file} is ${width} px wide; the plate draws it at ${box.width * store.scale}, so it would be enlarged`);
  }
  return { capture, box };
}

/** `--only app-store` or `--only play` builds one store's screenshots. */
const only = process.argv.includes('--only') ? process.argv[process.argv.indexOf('--only') + 1] : null;
if (only && !STORES[only]) throw new Error(`--only takes ${Object.keys(STORES).join(' or ')}`);

async function buildScreenshots(browser) {
  for (const [name, store] of Object.entries(STORES)) {
    if (only && name !== only) continue;
    const dir = `${name}/screenshots`;
    rmSync(path.join(here, dir), { recursive: true, force: true });
    const page = await browser.newPage({ deviceScaleFactor: store.scale });
    const listed = screens.filter(({ platforms }) => !platforms || platforms.includes(store.platform));
    for (const [index, { capture: name, caption }] of listed.entries()) {
      const { capture, box } = await loadCapture(path.join(here, 'app-captures', store.platform, `${name}.png`), store);
      const rel = `${dir}/${String(index + 1).padStart(2, '0')}-${name}.jpg`;
      const plate = await render(page, plateHtml(store, caption, capture, box), store);
      writeFileSync(out(rel), await sharp(plate).jpeg(JPEG).toBuffer());
      const { width, height } = await sharp(path.join(here, rel)).metadata();
      if (width !== store.width * store.scale || height !== store.height * store.scale) {
        throw new Error(`${rel} is ${width}x${height}, not the store's size`);
      }
      report(rel, MAX_BYTES.screenshot);
    }
    await page.close();
  }
}

/** Play's feature graphic: the link preview's road under its veil, the painted seal, the name and
 *  the line, kept clear of the edges Play may crop. */
function featureHtml() {
  const seal = dataUri(path.join(repo, 'design/assets/logo/lx-seal-v4-256.png'), 'image/png');
  const road = dataUri(path.join(repo, 'web/public/art/og-art.jpg'), 'image/jpeg');
  return `<!doctype html><meta charset="utf-8"><style>${FONTS}
html,body{margin:0;width:${FEATURE.width}px;height:${FEATURE.height}px;overflow:hidden}
.f{width:100%;height:100%;display:flex;align-items:center;justify-content:center;gap:34px;
background:linear-gradient(rgba(7,29,46,.5),rgba(7,29,46,.5)),url(${road}) center/cover}
img{width:150px;height:150px;}
h1{margin:0;font:400 76px/1.05 Marcellus,serif;color:${INK.cream2};letter-spacing:-1px;text-shadow:0 2px 12px rgba(0,0,0,.6)}
p{margin:12px 0 0;font:400 28px/1.3 'Open Sans',sans-serif;color:${INK.text};text-shadow:0 1px 8px rgba(0,0,0,.7)}
small{display:block;margin-top:6px;font:400 17px/1.3 'Open Sans',sans-serif;color:${INK.text};opacity:.85}
</style><div class="f"><img src="${seal}"><div><h1>Forever Sixty</h1><p>${FEATURE_LINE}</p><small>${FAN_LINE}</small></div></div>`;
}

async function buildFeatureGraphic(browser) {
  const page = await browser.newPage({ deviceScaleFactor: FEATURE.scale });
  const big = await render(page, featureHtml(), FEATURE);
  await page.close();
  await sharp(big).resize(FEATURE.width, FEATURE.height, { kernel: 'lanczos3' }).removeAlpha().png(PNG).toFile(out('play/feature-graphic.png'));
  report('play/feature-graphic.png', MAX_BYTES.feature);
}

await buildIcons();
const browser = await chromium.launch();
try {
  await buildFeatureGraphic(browser);
  await buildScreenshots(browser);
} finally {
  await browser.close();
}
