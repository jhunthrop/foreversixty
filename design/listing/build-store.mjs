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
// caption in Marcellus, and the capture below it, never enlarged. Captions come from
// store-screens.json; the captures from capture-apps.sh (app-captures/<platform>/<capture>.png).
// Run from anywhere: node design/listing/build-store.mjs
// sharp and @playwright/test resolve from web/node_modules.
import { mkdirSync, readFileSync, rmSync, statSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { PNG, renderTile } from '../../web/scripts/brand-art.mjs';

const here = path.dirname(fileURLToPath(import.meta.url));
const repo = path.resolve(here, '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const sharp = require('sharp');
const { chromium } = require('@playwright/test');

const { screens } = JSON.parse(readFileSync(path.join(here, 'store-screens.json'), 'utf8'));

/** The token values the plates paint (design/tokens/tokens.json; a browser page with no site CSS). */
const INK = { bg: '#100a06', cream: '#f3eee2', cream2: '#fcedcf', text: '#e3dfd9', hairline: 'rgba(243, 238, 226, 0.14)' };
const TAGLINE = 'Play your class better.';
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

/** A plate: the band over the ground, the caption, and the capture at the plate's width less its
 *  margins, its corners rounded as a phone's screen. */
function plateHtml({ width, height }, caption, capture) {
  const margin = Math.round(width * 0.09);
  return `<!doctype html><meta charset="utf-8"><style>${FONTS}
html,body{margin:0;width:${width}px;height:${height}px;overflow:hidden;background:${INK.bg}}
.p{position:relative;width:100%;height:100%;background:url(${art('ground.webp')}) 0 0/1600px auto repeat,${INK.bg}}
.band{position:absolute;left:0;right:0;top:0;height:${Math.round(height * 0.42)}px;
background:url(${art('band-1200.webp')}) center top/cover;
-webkit-mask-image:linear-gradient(#000 55%,transparent);mask-image:linear-gradient(#000 55%,transparent)}
h1{position:relative;margin:0;padding:${Math.round(height * 0.06)}px ${margin}px 0;font:400 ${Math.round(width * 0.082)}px/1.15 Marcellus,serif;
color:${INK.cream};text-align:center;text-wrap:balance;text-shadow:0 1px 8px rgba(0,0,0,.55)}
.s{position:absolute;left:${margin}px;right:${margin}px;bottom:0;top:${Math.round(height * 0.24)}px;overflow:hidden;
border-radius:${Math.round(width * 0.08)}px ${Math.round(width * 0.08)}px 0 0;border:1px solid ${INK.hairline};border-bottom:0}
.s img{display:block;width:100%;height:auto}
</style><div class="p"><div class="band"></div><h1>${caption}</h1><div class="s"><img src="${capture}"></div></div>`;
}

/** The page `html` at `size`, as an opaque PNG at the page's device scale. */
async function render(page, html, size) {
  await page.setViewportSize({ width: size.width, height: size.height });
  await page.setContent(html, { waitUntil: 'load' });
  await page.evaluate(() => document.fonts.ready);
  return sharp(await page.screenshot({ type: 'png' })).flatten({ background: INK.bg }).png().toBuffer();
}

/** Fails when a capture would be enlarged: its pixels must cover the plate's screen box. */
async function checkCapture(file, store) {
  const { width } = await sharp(file).metadata();
  const box = Math.round(store.width * (1 - 2 * 0.09)) * store.scale;
  if (width < box) throw new Error(`${file} is ${width} px wide; the plate draws it at ${box}, so it would be enlarged`);
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
    for (const [index, { capture, caption }] of screens.entries()) {
      const source = path.join(here, 'app-captures', store.platform, `${capture}.png`);
      await checkCapture(source, store);
      const rel = `${dir}/${String(index + 1).padStart(2, '0')}-${capture}.jpg`;
      const plate = await render(page, plateHtml(store, caption, dataUri(source, 'image/png')), store);
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
img{width:150px;height:150px;filter:drop-shadow(0 4px 12px rgba(0,0,0,.5))}
h1{margin:0;font:400 76px/1.05 Marcellus,serif;color:${INK.cream2};letter-spacing:-1px;text-shadow:0 2px 12px rgba(0,0,0,.6)}
p{margin:12px 0 0;font:400 28px/1.3 'Open Sans',sans-serif;color:${INK.text};text-shadow:0 1px 8px rgba(0,0,0,.7)}
small{display:block;margin-top:6px;font:400 17px/1.3 'Open Sans',sans-serif;color:${INK.text};opacity:.85}
</style><div class="f"><img src="${seal}"><div><h1>Forever Sixty</h1><p>${TAGLINE}</p><small>${FAN_LINE}</small></div></div>`;
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
