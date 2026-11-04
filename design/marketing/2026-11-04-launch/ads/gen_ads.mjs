// design/marketing/2026-11-04-launch/ads/gen_ads.mjs
// Renders the Reddit launch ads from the production captures in ./captures/ (capture_ads.mjs
// takes them). Same brand treatment as design/listing/gen_gallery.mjs: the site's dark and
// gold, Cinzel wordmark, Barlow text, a bordered frame around each capture.
//   node design/marketing/2026-11-04-launch/ads/capture_ads.mjs   # once, from production
//   node design/marketing/2026-11-04-launch/ads/gen_ads.mjs
// Output: ./out/<concept>-<width>x<height>.png at the exact pixel size Reddit asks for.
//
// Crispness: a capture is taken at device scale 2 and only ever drawn smaller, never enlarged,
// and always whole: the stack is fitted inside the stage, never clipped.
// Each line of output prints the drawn scale and the share of the plate that text covers;
// a scale over 1 or text over TEXT_SHARE_LIMIT fails the run.
import { mkdir, readFile } from 'node:fs/promises';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const here = path.dirname(fileURLToPath(import.meta.url));
const repo = path.resolve(here, '..', '..', '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const { chromium } = require('@playwright/test');

const FONTS = path.join(repo, 'web', 'src', 'assets', 'fonts');
const CAPTURES = path.join(here, 'captures');
const OUT = path.join(here, 'out');
/** Reddit's own guidance: keep text to about a fifth of the image. */
const TEXT_SHARE_LIMIT = 0.2;

const BRAND = { url: 'foreversixty.gg', tagline: 'Play your class better' };

const FORMATS = [
  { width: 1200, height: 628, layout: 'wide' },
  { width: 1080, height: 1080, layout: 'tall' },
  { width: 1080, height: 1350, layout: 'tall' },
];

const GUIDE_COPY = {
  headline: 'Every Forever spec. A real rotation, the talents that matter.',
  sub: 'Free guides, BiS and a sim for WoW Forever',
};

/**
 * `shots` per layout: the landscape guide plate drops the guide header so the cards draw larger.
 * `note` is a small caption under the capture that says what the capture shows.
 */
const guide = (name) => ({ ...GUIDE_COPY, shots: { wide: [`guide-${name}-cards`], tall: [`guide-${name}`] } });
const CONCEPTS = [
  { name: 'a-guides', ...guide('fury') },
  {
    name: 'b-bis',
    shots: { wide: ['bis-head', 'bis-list'], tall: ['bis-head', 'bis-list'] },
    headline: 'BiS for every Forever spec, leveling to raid',
    sub: 'Every slot at 20, 30, 40, 50 and 60, and where it drops',
  },
  {
    name: 'c-sim',
    shots: { wide: ['sim-dps', 'sim-doll'], tall: ['sim-dps', 'sim-doll'] },
    headline: 'Sim your own character in the browser',
    sub: 'Paste your gear or sign in. Free, and it runs on your machine',
    note: 'Shown: Fury BiS at 60, raid-buffed, Patchwerk, 3:00, one target',
  },
  { name: 'a-guides-rogue', ...guide('rogue') },
  { name: 'a-guides-shaman', ...guide('shaman') },
];

const dataUrl = async (file, type) => `data:${type};base64,${(await readFile(file)).toString('base64')}`;
const font = (file) => dataUrl(path.join(FONTS, file), 'font/woff2');
const escapeHtml = (text) => text.replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]);

const CREST = `<svg class="crest" viewBox="0 0 256 256" aria-hidden="true"><defs><linearGradient id="g" x1="0" y1="0" x2="0" y2="1"><stop offset="0" stop-color="#fbe7a1"/><stop offset=".45" stop-color="#e5b955"/><stop offset="1" stop-color="#a8762a"/></linearGradient></defs><circle cx="128" cy="128" r="112.64" fill="#0d111a"/><circle cx="128" cy="128" r="112.64" fill="none" stroke="url(#g)" stroke-width="12.8"/><text x="128" y="152" text-anchor="middle" font-family="Cinzel,serif" font-weight="700" font-size="72" fill="url(#g)">LX</text></svg>`;

/** Sizes per layout, in plate pixels. */
const LAYOUT_CSS = {
  wide: `
body{display:grid;grid-template-columns:360px minmax(0,1fr);gap:40px;padding:44px 44px 44px 52px}
.copy{display:flex;flex-direction:column;justify-content:space-between}
.mark{font-size:20px}.crest{width:34px;height:34px}
h1{font-size:44px}.sub{font-size:21px;margin-top:16px}
.foot{display:flex;flex-direction:column;gap:6px}`,
  tall: `
body{display:flex;flex-direction:column;padding:52px 48px 44px}
.copy{display:contents}
.top{order:0;display:flex;justify-content:space-between;align-items:center}
.lede{order:1;margin:34px 0 30px}
.stage{order:2;flex:1}
.foot{order:3;display:flex;justify-content:space-between;align-items:baseline;margin-top:28px;padding-top:22px;border-top:1px solid #262e40}
.mark{font-size:24px}.crest{width:40px;height:40px}
h1{font-size:58px}.sub{font-size:26px;margin-top:14px}`,
};

async function plateHtml(concept, format) {
  const [cinzel, barlow400, barlow600, condensed800, ...shots] = await Promise.all([
    font('cinzel-700-subset.woff2'),
    font('barlow-400-subset.woff2'),
    font('barlow-600-subset.woff2'),
    font('barlow-condensed-800-subset.woff2'),
    ...concept.shots[format.layout].map((name) => dataUrl(path.join(CAPTURES, `${name}.png`), 'image/png')),
  ]);
  const frames = shots.map((src) => `<div class="frame"><img class="shot" src="${src}" alt=""></div>`).join('');
  return `<!doctype html>
<html lang="en"><head><meta charset="utf-8">
<style>
@font-face{font-family:Cinzel;src:url(${cinzel}) format('woff2');font-weight:700}
@font-face{font-family:Barlow;src:url(${barlow400}) format('woff2');font-weight:400}
@font-face{font-family:Barlow;src:url(${barlow600}) format('woff2');font-weight:600}
@font-face{font-family:'Barlow Condensed';src:url(${condensed800}) format('woff2');font-weight:800}
html,body{margin:0;width:${format.width}px;height:${format.height}px;overflow:hidden}
body{box-sizing:border-box;background:#07090d;color:#e9e4d8;font-family:Barlow,sans-serif;
  -webkit-font-smoothing:antialiased;text-rendering:geometricPrecision;
  background-image:radial-gradient(ellipse 70% 60% at 60% 50%,rgba(229,185,85,.08),transparent 70%)}
.mark{display:flex;align-items:center;gap:12px;font-family:Cinzel,serif;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:#f2eee4;white-space:nowrap}
h1{margin:0;font-family:'Barlow Condensed',sans-serif;font-weight:800;line-height:1.02;color:#f2eee4;letter-spacing:.005em}
h1 em{font-style:normal;color:#e5b955}
.sub{margin:0;line-height:1.3;color:#b9b3a4}
.url{font-weight:600;color:#e5b955;font-size:22px;letter-spacing:.02em}
.tagline{font-size:17px;color:#9a9484;letter-spacing:.04em}
.stage{position:relative;min-height:0;min-width:0;overflow:hidden;display:flex;flex-direction:column;justify-content:center;align-items:center;gap:14px}
.frame{background:#07090d;border:1px solid #262e40;border-radius:8px;overflow:hidden;box-shadow:0 24px 60px rgba(0,0,0,.65);flex:none}
.shot{display:block}
.note{margin:0;font-size:15px;color:#9a9484;letter-spacing:.02em;flex:none}
h1{text-wrap:balance}
${LAYOUT_CSS[format.layout]}
</style></head><body>
<div class="copy">
  <div class="top"><div class="mark text">${CREST}Forever Sixty</div></div>
  <div class="lede"><h1 class="text">${headlineHtml(concept.headline)}</h1><p class="sub text">${escapeHtml(concept.sub)}</p></div>
  <div class="foot"><span class="tagline text">${BRAND.tagline}</span><span class="url text">${BRAND.url}</span></div>
</div>
<div class="stage">${frames}${concept.note ? `<p class="note text">${escapeHtml(concept.note)}</p>` : ''}</div>
</body></html>`;
}

/** A two-sentence headline sets its first sentence gold, on a line of its own. */
function headlineHtml(headline) {
  const split = headline.indexOf('. ');
  if (split === -1) return escapeHtml(headline);
  return `<em>${escapeHtml(headline.slice(0, split + 1))}</em><br>${escapeHtml(headline.slice(split + 2))}`;
}

/**
 * Fits the stacked captures inside the stage at one shared scale, never above 1, so nothing is
 * clipped or enlarged. Runs in the page.
 */
function placeCaptures() {
  const stage = document.querySelector('.stage');
  const images = [...document.querySelectorAll('.shot')];
  if (!stage || images.some((image) => image.naturalWidth === 0)) throw new Error('capture did not load');
  const border = 2;
  const gap = 14;
  const box = stage.getBoundingClientRect();
  const note = document.querySelector('.note');
  const reserved = note ? note.getBoundingClientRect().height + gap : 0;
  const widest = Math.max(...images.map((image) => image.naturalWidth));
  const tallest = images.reduce((sum, image) => sum + image.naturalHeight, 0);
  const room = box.height - reserved - (gap + border) * images.length;
  const scale = Math.min(1, (box.width - border) / widest, room / tallest);
  for (const image of images) {
    const width = Math.floor(image.naturalWidth * scale);
    const height = Math.floor(image.naturalHeight * scale);
    const canvas = document.createElement('canvas');
    canvas.width = width;
    canvas.height = height;
    canvas.style.cssText = `display:block;width:${width}px;height:${height}px`;
    const context = canvas.getContext('2d');
    context.imageSmoothingEnabled = scale < 1;
    context.imageSmoothingQuality = 'high';
    context.drawImage(image, 0, 0, width, height);
    image.replaceWith(canvas);
  }
  return document.fonts.ready.then(() => {
    const textArea = [...document.querySelectorAll('.text')]
      .filter((element) => element.offsetParent !== null)
      .reduce((sum, element) => {
        const r = element.getBoundingClientRect();
        return sum + r.width * r.height;
      }, 0);
    return { scale, textShare: textArea / (window.innerWidth * window.innerHeight) };
  });
}

async function main() {
  await mkdir(OUT, { recursive: true });
  const browser = await chromium.launch();
  for (const format of FORMATS) {
    const page = await browser.newPage({ viewport: { width: format.width, height: format.height }, deviceScaleFactor: 1 });
    for (const concept of CONCEPTS) {
      await page.setContent(await plateHtml(concept, format), { waitUntil: 'load' });
      const placed = await page.evaluate(placeCaptures);
      const file = path.join(OUT, `${concept.name}-${format.width}x${format.height}.png`);
      await page.screenshot({ path: file, type: 'png' });
      const verdict = placed.scale > 1 || placed.textShare > TEXT_SHARE_LIMIT ? 'FAIL' : 'ok';
      process.stdout.write(`${path.relative(repo, file)} | drawn ${placed.scale.toFixed(2)}x of the 2x capture | text ${(placed.textShare * 100).toFixed(1)}% | ${verdict}\n`);
      if (verdict === 'FAIL') process.exitCode = 1;
    }
    await page.close();
  }
  await browser.close();
}

main().catch((error) => {
  process.stderr.write(`${error instanceof Error ? error.stack : String(error)}\n`);
  process.exitCode = 1;
});
