// design/listing/gen_gallery.mjs
// Renders the store gallery plates for the Forever Sixty listing: each client capture on a
// 16:9 plate in the site's page head (the band art over the ground, Marcellus and Open Sans, the
// painted seal), with the surface's name and one line under it.
// Run from anywhere:
//   node design/listing/gen_gallery.mjs                    # every plate
//   node design/listing/gen_gallery.mjs --only bag-arrow   # one plate
// Output: design/listing/gallery/NN-<name>.png at 3200x1800 (a 1600x900 layout at device scale 2,
// so the caption text is rasterised at full resolution). Playwright comes from web/node_modules.
//
// Crispness rules (design/listing/RECAPTURE.md says how the owner takes the captures):
// - a capture is drawn at most 1:1, one capture pixel to one plate pixel, never enlarged;
// - a lossless PNG in design/listing/captures/ wins over the site's JPEG in web/public/addon/;
// - the capture sits on whole device pixels, so a 1:1 capture is copied, not resampled.
// Each plate prints its source, native size and how much of the stage it fills; a JPEG source
// or a capture too small to fill the stage is flagged RECAPTURE.
import { mkdir, readFile, stat } from 'node:fs/promises';
import { existsSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { INK } from './brand-ink.mjs';

const here = path.dirname(fileURLToPath(import.meta.url));
const repo = path.resolve(here, '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const { chromium } = require('@playwright/test');

/** Lossless captures first, then the site's JPEGs. */
const CAPTURE_SOURCES = [
  { dir: path.join(here, 'captures'), ext: 'png', type: 'image/png', lossless: true },
  { dir: path.join(repo, 'web', 'public', 'addon'), ext: 'jpg', type: 'image/jpeg', lossless: false },
];
const FONTS = path.join(repo, 'web', 'src', 'assets', 'fonts');
const OUT = path.join(here, 'gallery');
const CAPTIONS_FILE = path.join(here, 'curseforge-description.md');

/** Layout in CSS pixels; the PNG is DEVICE_SCALE times larger. */
const WIDTH = 1600;
const HEIGHT = 900;
const DEVICE_SCALE = 2;
/** The capture's box above the caption. */
const STAGE_WIDTH = 1440;
const STAGE_HEIGHT = 600;
/** A capture filling less than this share of the stage's width or height reads as small. */
const MIN_FILL = 0.6;
/**
 * Neither store publishes a gallery limit (checked 2026-11-03); this is our own ceiling so an
 * upload never fails on size. A plate over it fails the run.
 */
const MAX_PLATE_BYTES = 5 * 1024 * 1024;

/** Capture file names in gallery order (gen_captures.py renders each from the addon's mocks); titles and
 * captions come from CAPTIONS_FILE. */
const PLATES = [
  'chip',
  'window-me',
  'window-builds',
  'window-gear',
  'window-rotation',
  'window-guild',
  'window-officer',
  'bag-arrow',
  'quest-reward-marks',
  'player-tooltip',
];

/** Reads "N. **Title.** Caption." lines under "## Gallery captions". */
async function readCaptions() {
  const markdown = await readFile(CAPTIONS_FILE, 'utf8');
  const section = markdown.split(/^## Gallery captions$/m)[1]?.split(/^## /m)[0];
  if (!section) throw new Error(`no "## Gallery captions" section in ${CAPTIONS_FILE}`);
  const captions = [...section.matchAll(/^\d+\. \*\*(.+?)\.\*\* (.+)$/gm)].map(([, title, text]) => ({ title, text }));
  if (captions.length !== PLATES.length) {
    throw new Error(`${CAPTIONS_FILE} has ${captions.length} gallery captions; the generator has ${PLATES.length} plates`);
  }
  return captions;
}

function findCapture(name) {
  for (const source of CAPTURE_SOURCES) {
    const file = path.join(source.dir, `${name}.${source.ext}`);
    if (existsSync(file)) return { ...source, file };
  }
  throw new Error(`no capture for ${name} in ${CAPTURE_SOURCES.map((s) => s.dir).join(' or ')}`);
}

// A page set from a string has no file access, so fonts and captures travel inline.
const dataUrl = async (file, type) => `data:${type};base64,${(await readFile(file)).toString('base64')}`;
const font = (file) => dataUrl(path.join(FONTS, file), 'font/woff2');
const escapeHtml = (text) => text.replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]);

/** The site's page head behind the plate: the band art over the ground (web/public/art). */
const ART = path.join(repo, 'web', 'public', 'art');
/** The painted seal in the foot, at 48 CSS px (96 device px): the painted seal from 64 px up. */
const SEAL = path.join(repo, 'design', 'assets', 'logo', 'lx-seal-v4-128.png');

async function plateHtml(capture, { title, text }) {
  const [marcellus, openSans, ground, band, seal, shot] = await Promise.all([
    font('marcellus-400-subset.woff2'),
    font('open-sans-400-subset.woff2'),
    dataUrl(path.join(ART, 'ground.webp'), 'image/webp'),
    dataUrl(path.join(ART, 'band-2400.webp'), 'image/webp'),
    dataUrl(SEAL, 'image/png'),
    dataUrl(capture.file, capture.type),
  ]);
  return `<!doctype html>
<html lang="en"><head><meta charset="utf-8">
<style>
@font-face{font-family:Marcellus;src:url(${marcellus}) format('woff2')}
@font-face{font-family:'Open Sans';src:url(${openSans}) format('woff2')}
html,body{margin:0;width:${WIDTH}px;height:${HEIGHT}px;overflow:hidden}
body{position:relative;color:${INK.text};font-family:'Open Sans',sans-serif;display:grid;grid-template-rows:1fr auto;box-sizing:border-box;padding:64px 80px 52px;
  -webkit-font-smoothing:antialiased;background:url(${ground}) 0 0/1600px auto repeat,${INK.bg}}
body::before{content:'';position:absolute;inset:0 0 auto;height:62%;background:url(${band}) center top/cover;
  -webkit-mask-image:linear-gradient(#000 50%,transparent);mask-image:linear-gradient(#000 50%,transparent)}
.stage{position:relative;height:${STAGE_HEIGHT}px}
.frame{position:absolute;background:${INK.bg};border:1px solid ${INK.line};border-radius:6px;overflow:hidden}
.shot{display:block}
.foot{position:relative;display:grid;grid-template-columns:auto 1fr;gap:28px;align-items:start;margin-top:40px;padding-top:26px;border-top:1px solid ${INK.line}}
.mark{display:flex;align-items:center;gap:14px;font-family:Marcellus,serif;font-size:28px;color:${INK.cream};white-space:nowrap}
.mark img{width:48px;height:48px}
.text{min-width:0}
h1{margin:0 0 6px;font-family:Marcellus,serif;font-weight:400;font-size:32px;line-height:1.1;color:${INK.cream}}
p{margin:0;font-size:20px;line-height:1.4;color:${INK.mutedGround};max-width:1100px;min-height:calc(2 * 1.4em)}
</style></head><body>
<div class="stage"><div class="frame"><img class="shot" src="${shot}" alt=""></div></div>
<div class="foot">
  <div class="mark"><img src="${seal}" alt="">Forever Sixty</div>
  <div class="text"><h1>${escapeHtml(title)}</h1><p>${escapeHtml(text)}</p></div>
</div>
</body></html>`;
}

/**
 * Draws the capture at most one capture pixel per device pixel, centred in the stage on whole
 * CSS pixels. Runs in the page. The capture goes through a canvas sized to whole CSS pixels:
 * an <img> of odd device size gets a half CSS pixel and Chromium resamples it, where a canvas
 * at its own device size is copied as is. At 1:1 drawImage is an exact copy; a capture larger
 * than the stage shrinks with the high-quality filter.
 */
function placeCapture({ stageWidth, stageHeight, deviceScale }) {
  const shot = document.querySelector('.shot');
  const frame = document.querySelector('.frame');
  if (!(shot instanceof HTMLImageElement) || !(frame instanceof HTMLElement) || shot.naturalWidth === 0) {
    throw new Error('capture did not load');
  }
  const border = 1;
  const fit = Math.min(1, ((stageWidth - 2 * border) * deviceScale) / shot.naturalWidth, ((stageHeight - 2 * border) * deviceScale) / shot.naturalHeight);
  const deviceWidth = Math.floor(shot.naturalWidth * fit);
  const deviceHeight = Math.floor(shot.naturalHeight * fit);
  const cssWidth = Math.ceil(deviceWidth / deviceScale);
  const cssHeight = Math.ceil(deviceHeight / deviceScale);
  const canvas = document.createElement('canvas');
  canvas.width = cssWidth * deviceScale;
  canvas.height = cssHeight * deviceScale;
  canvas.style.cssText = `display:block;width:${cssWidth}px;height:${cssHeight}px`;
  const context = canvas.getContext('2d');
  if (!context) throw new Error('no 2d canvas');
  context.imageSmoothingEnabled = fit < 1;
  context.imageSmoothingQuality = 'high';
  context.drawImage(shot, 0, 0, deviceWidth, deviceHeight);
  shot.replaceWith(canvas);
  frame.style.left = `${Math.round((stageWidth - cssWidth) / 2 - border)}px`;
  frame.style.top = `${Math.round((stageHeight - cssHeight) / 2 - border)}px`;
  return document.fonts.ready.then(() => ({
    natural: [shot.naturalWidth, shot.naturalHeight],
    drawn: [deviceWidth, deviceHeight],
    fill: Math.max(deviceWidth / (stageWidth * deviceScale), deviceHeight / (stageHeight * deviceScale)),
  }));
}

function report(file, capture, placed, bytes) {
  const reasons = [];
  if (!capture.lossless) reasons.push('JPEG source');
  if (placed.fill < MIN_FILL) reasons.push(`fills ${Math.round(placed.fill * 100)}% of the stage`);
  const drawnScale = placed.drawn[0] / placed.natural[0];
  const line = [
    path.relative(repo, file),
    `source ${path.relative(repo, capture.file)} ${placed.natural.join('x')}`,
    `drawn ${placed.drawn.join('x')} (${drawnScale === 1 ? '1:1' : `${drawnScale.toFixed(2)}x down`})`,
    `${(bytes / 1024 / 1024).toFixed(2)} MB`,
    reasons.length ? `RECAPTURE: ${reasons.join(', ')}` : 'ok',
  ];
  process.stdout.write(`${line.join(' | ')}\n`);
}

/** `--only <name>` renders one plate (a preview while the other captures are still to shoot). */
const only = process.argv.includes('--only') ? process.argv[process.argv.indexOf('--only') + 1] : null;
if (only && !PLATES.includes(only)) throw new Error(`--only takes one of ${PLATES.join(', ')}`);

async function main() {
  const captions = await readCaptions();
  await mkdir(OUT, { recursive: true });
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: WIDTH, height: HEIGHT }, deviceScaleFactor: DEVICE_SCALE });
  for (const [index, name] of PLATES.entries()) {
    if (only && name !== only) continue;
    const capture = findCapture(name);
    await page.setContent(await plateHtml(capture, captions[index]), { waitUntil: 'load' });
    const placed = await page.evaluate(placeCapture, { stageWidth: STAGE_WIDTH, stageHeight: STAGE_HEIGHT, deviceScale: DEVICE_SCALE });
    const file = path.join(OUT, `${String(index + 1).padStart(2, '0')}-${name}.png`);
    await page.screenshot({ path: file, type: 'png', scale: 'device' });
    const { size } = await stat(file);
    if (size > MAX_PLATE_BYTES) throw new Error(`${file} is ${size} bytes, over the ${MAX_PLATE_BYTES}-byte ceiling`);
    report(file, capture, placed, size);
  }
  await browser.close();
}

main().catch((error) => {
  process.stderr.write(`${error instanceof Error ? error.stack : String(error)}\n`);
  process.exitCode = 1;
});
