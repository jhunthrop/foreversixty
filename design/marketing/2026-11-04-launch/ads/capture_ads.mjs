// design/marketing/2026-11-04-launch/ads/capture_ads.mjs
// Captures the production pages the Reddit ads show, at device scale 2, as lossless PNGs in
// ./captures/. Nothing on a page is edited: every number is what foreversixty.gg renders for a
// signed-out visitor. The sim capture runs the site's own Fury BiS set (level 60, Alliance)
// with the Fury guide's build, a public example, never a player's character.
//   node design/marketing/2026-11-04-launch/ads/capture_ads.mjs
// Playwright comes from web/node_modules (npm ci in web/ first).
import { mkdir } from 'node:fs/promises';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const here = path.dirname(fileURLToPath(import.meta.url));
const repo = path.resolve(here, '..', '..', '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const { chromium } = require('@playwright/test');

const SITE = 'https://foreversixty.gg';
const OUT = path.join(here, 'captures');
const VIEWPORT = { width: 1440, height: 900 };
/** BiS and the sim at a laptop width, so their cards pack denser and read larger on the plate. */
const NARROW_VIEWPORT = { width: 1024, height: 900 };
const DEVICE_SCALE = 2;
/** Concept A's guides: Fury first, then the Rogue and Shaman variants (talents changed in the latest client build). */
const GUIDES = [
  { slug: 'warrior/fury', name: 'fury' },
  { slug: 'rogue/combat', name: 'rogue' },
  { slug: 'shaman/elemental', name: 'shaman' },
];
/** Space kept around a cropped element, in CSS pixels. */
const PAD = 16;
/** The sim runs in the visitor's browser; 3,000 iterations finish in a few seconds. */
const SIM_TIMEOUT_MS = 90_000;

/** The guide build's own FS1 code (its "Sim this build" link) plus the BiS 60 Alliance set. */
const FURY_BUILD = 'FS1:1.60.1.70755:warrior:human:30323013002/051531025120100231/2';
const FURY_BIS_60 = [
  'head=12640', 'neck=13089', 'shoulder=16729', 'back=13340', 'chest=252483', 'wrist=13400',
  'hands=19049', 'waist=13142', 'legs=15065', 'feet=16711', 'finger1=17713', 'finger2=13217',
  'trinket1=249469', 'trinket2=249470', 'main_hand=1728', 'off_hand=279261', 'ranged=12651',
].join(',');

/** Union of the boxes of the given selectors, in page coordinates, padded. Runs in the page. */
function unionBox({ selectors, pad }) {
  const boxes = selectors.map((selector) => {
    const element = document.querySelector(selector);
    if (!element) throw new Error(`no element for ${selector}`);
    return element.getBoundingClientRect();
  });
  const left = Math.min(...boxes.map((b) => b.left)) - pad;
  const top = Math.min(...boxes.map((b) => b.top)) + window.scrollY - pad;
  const right = Math.max(...boxes.map((b) => b.right)) + pad;
  const bottom = Math.max(...boxes.map((b) => b.bottom)) + window.scrollY + pad;
  return { x: Math.floor(left), y: Math.floor(top), width: Math.ceil(right - left), height: Math.ceil(bottom - top) };
}

async function open(page, url) {
  await page.goto(url, { waitUntil: 'networkidle' });
  await page.evaluate(() => document.fonts.ready);
}

async function shoot(page, name, clip) {
  const file = path.join(OUT, `${name}.png`);
  await page.screenshot({ path: file, clip, fullPage: true, type: 'png', animations: 'disabled' });
  process.stdout.write(`${path.relative(repo, file)} ${clip.width * DEVICE_SCALE}x${clip.height * DEVICE_SCALE}\n`);
}

/**
 * Guide header, build and rotation cards. No hover tooltip: the guide's talent icons open no
 * tooltip on a mouse move at this width (tried 2026-11-03), so the plate shows the page at rest.
 */
async function captureGuide(page, { slug, name }) {
  await open(page, `${SITE}/guides/${slug}`);
  const box = await page.evaluate(unionBox, { selectors: ['h1', '[data-testid="guide-action-rail"]'], pad: PAD });
  // Build and Rotation only: the stat card on the right is the sim's job in concept C.
  const rotation = await page.evaluate(() => {
    const heading = [...document.querySelectorAll('[data-testid="guide-action-rail"] h2')].find((h) => h.textContent.trim() === 'Rotation');
    const card = heading?.closest('section');
    if (!card) throw new Error('no Rotation card');
    return card.getBoundingClientRect().right;
  });
  const clip = { ...box, width: Math.ceil(rotation + PAD - box.x) };
  await shoot(page, `guide-${name}`, clip);
  // The cards alone, for the landscape plate, where the header would shrink them.
  const rail = await page.evaluate(unionBox, { selectors: ['[data-testid="guide-action-rail"]'], pad: PAD });
  await shoot(page, `guide-${name}-cards`, { ...rail, width: clip.width });
}

/** Box of the elements whose own text is exactly one of `texts`, padded. Runs in the page. */
function textBox({ texts, pad }) {
  const hits = texts.map((text) => {
    const element = [...document.querySelectorAll('button,a,label,span,div')].find((e) => e.textContent.trim().toLowerCase() === text.toLowerCase() && e.offsetParent !== null);
    if (!element) throw new Error(`no visible element reading ${text}`);
    return element.getBoundingClientRect();
  });
  const left = Math.min(...hits.map((b) => b.left)) - pad;
  const top = Math.min(...hits.map((b) => b.top)) + window.scrollY - pad;
  const right = Math.max(...hits.map((b) => b.right)) + pad;
  const bottom = Math.max(...hits.map((b) => b.bottom)) + window.scrollY + pad;
  return { x: Math.floor(left), y: Math.floor(top), width: Math.ceil(right - left), height: Math.ceil(bottom - top) };
}

/** BiS at level 60, Alliance: the title with the level strip, then the Head and Neck rows. */
async function captureBis(page) {
  await page.setViewportSize(NARROW_VIEWPORT);
  await open(page, `${SITE}/bis/warrior/fury#band-alliance-60`);
  // At this viewport the faction pills wrap under the level strip; the head stops at the strip.
  const strip = await page.evaluate(textBox, { texts: ['20 to 29', '60'], pad: PAD });
  const title = await page.evaluate(() => {
    const h1 = [...document.querySelectorAll('h1')].find((e) => e.offsetParent !== null);
    if (!h1) throw new Error('no visible h1');
    const top = h1.getBoundingClientRect().top;
    const crest = [...document.querySelectorAll('img')].find((img) => {
      const r = img.getBoundingClientRect();
      return img.offsetParent !== null && r.width > 40 && Math.abs(r.top - top) < 80;
    });
    return { top: top + window.scrollY, left: crest ? crest.getBoundingClientRect().left : h1.getBoundingClientRect().left };
  });
  const top = Math.floor(title.top) - PAD;
  const slotHeader = await page.evaluate(textBox, { texts: ['Slot'], pad: PAD });
  const list = await page.evaluate(unionBox, {
    selectors: ['#band-alliance-60 [data-testid="bis-list-alliance-60"]'],
    pad: PAD,
  });
  const rows = await page.evaluate(unionBox, {
    selectors: ['#band-alliance-60 [data-testid="bis-slot-alliance-60-head"]', '#band-alliance-60 [data-testid="bis-slot-alliance-60-neck"]'],
    pad: 2,
  });
  const left = Math.min(list.x, Math.floor(title.left) - PAD);
  await shoot(page, 'bis-head', { x: left, y: top, width: list.x + list.width - left, height: strip.y + strip.height - top });
  await shoot(page, 'bis-list', { ...list, y: slotHeader.y, height: rows.y + rows.height - slotHeader.y });
}

/** A finished quick sim of the public Fury BiS 60 set: the DPS line and the paper doll. */
async function captureSim(page) {
  const code = encodeURIComponent(`${FURY_BUILD}:${FURY_BIS_60}`);
  await page.setViewportSize(NARROW_VIEWPORT);
  await open(page, `${SITE}/sim?code=${code}`);
  await page.getByRole('button', { name: /^run sim$/i }).first().click();
  await page.getByRole('button', { name: /^run again$/i }).first().waitFor({ timeout: SIM_TIMEOUT_MS });
  await page.waitForTimeout(800);
  const [doll, dps] = await page.evaluate(({ pad, minWidth }) => {
    const card = (text, minHeight) => {
      let node = [...document.querySelectorAll('h2,button')].find((e) => e.textContent.trim().toLowerCase() === text);
      while (node && (node.getBoundingClientRect().width < minWidth || node.getBoundingClientRect().height < minHeight)) node = node.parentElement;
      if (!node) throw new Error(`no card around ${text}`);
      const r = node.getBoundingClientRect();
      return { x: Math.floor(r.left - pad), y: Math.floor(r.top + window.scrollY - pad), width: Math.ceil(r.width + 2 * pad), height: Math.ceil(r.height + 2 * pad) };
    };
    return [card('character', 400), card('run again', 100)];
  }, { pad: PAD, minWidth: NARROW_VIEWPORT.width - 200 });
  // The doll down to its sixth row: the lower rows are mostly empty space on the left.
  // An item icon reaches about 7 px below its name and the next row starts 10 px after that.
  const sixthRow = await page.evaluate(textBox, { texts: ['Vambraces of the Sadist'], pad: 0 });
  await shoot(page, 'sim-doll', { ...doll, height: sixthRow.y + sixthRow.height + 12 - doll.y });
  await shoot(page, 'sim-dps', dps);
}

async function main() {
  await mkdir(OUT, { recursive: true });
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: VIEWPORT, deviceScaleFactor: DEVICE_SCALE, colorScheme: 'dark' });
  for (const guide of GUIDES) await captureGuide(page, guide);
  await captureBis(page);
  await captureSim(page);
  await browser.close();
}

main().catch((error) => {
  process.stderr.write(`${error instanceof Error ? error.stack : String(error)}\n`);
  process.exitCode = 1;
});
