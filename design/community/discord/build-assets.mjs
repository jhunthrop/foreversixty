// design/community/discord/build-assets.mjs
// Builds the Discord images from the brand assets:
//   assets/icon.png       server icon, the Forever Sixty mark at 512 px
//   assets/premium.png    Premium role icon, the mark at 64 px
//   assets/crest-<class>.png  class role icons and onboarding emoji, the HD crests at 128 px
//   assets/<faction>.png  onboarding emoji, the faction logos at 128 px
//   assets/banner.png     server banner, 960x540: the stacked lockup, the line, the nine crests
// Needs dwebp and sips (macOS) and @playwright/test resolvable from web/ (or NODE_PATH).
import { execFileSync } from 'node:child_process';
import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { fileURLToPath } from 'node:url';

const here = (p) => fileURLToPath(new URL(p, import.meta.url));
const repo = (p) => here(`../../../${p}`);
const classes = JSON.parse(readFileSync(repo('web/src/data/classes.json'), 'utf8'));
const FACTIONS = ['alliance', 'horde'];
const BANNER = { width: 960, height: 540 };
const COLORS = { bg: '#07090d', gold: '#e5b955', muted: '#9a9484' };

function sipsResize(src, size, out) {
  execFileSync('sips', ['-Z', String(size), src, '--out', out], { stdio: 'ignore' });
}

function buildRasterIcons() {
  sipsResize(repo('design/logo/export/foreversixty-icon-1024.png'), 512, here('assets/icon.png'));
  sipsResize(repo('design/logo/export/foreversixty-icon-1024.png'), 64, here('assets/premium.png'));
  for (const { slug } of classes) {
    execFileSync('dwebp', [repo(`web/public/icons/hd/crests/${slug}.webp`), '-o', here(`assets/crest-${slug}.png`)], {
      stdio: 'ignore',
    });
  }
  for (const faction of FACTIONS) {
    execFileSync('dwebp', [repo(`web/public/icons/hd/faction/${faction}-logo-512.webp`), '-o', here(`assets/${faction}.png`)], {
      stdio: 'ignore',
    });
    sipsResize(here(`assets/${faction}.png`), 128, here(`assets/${faction}.png`));
  }
}

function bannerHtml() {
  const lockup = readFileSync(repo('design/logo/foreversixty-lockup-stacked.svg'), 'utf8');
  const font = readFileSync(repo('apps/shared/fonts/BarlowCondensed-ExtraBold.ttf')).toString('base64');
  const crests = classes
    .map(({ slug }) => {
      const png = readFileSync(here(`assets/crest-${slug}.png`)).toString('base64');
      return `<img src="data:image/png;base64,${png}" alt="">`;
    })
    .join('');
  return `<!doctype html><html><head><style>
@font-face{font-family:BC;src:url(data:font/ttf;base64,${font})}
html,body{margin:0;width:${BANNER.width}px;height:${BANNER.height}px;background:${COLORS.bg};overflow:hidden}
body{display:flex;flex-direction:column;align-items:center;justify-content:center;gap:26px;
background:radial-gradient(ellipse 60% 55% at 50% 42%,rgba(229,185,85,.16),transparent 70%),
radial-gradient(ellipse 120% 60% at 50% 120%,rgba(38,64,92,.55),transparent 70%),${COLORS.bg}}
.lockup svg{width:420px;height:auto;display:block}
.line{font:800 30px/1 BC,sans-serif;letter-spacing:.18em;text-transform:uppercase;color:${COLORS.gold}}
.crests{display:flex;gap:14px;margin-top:6px}.crests img{width:52px;height:52px}
</style></head><body><div class="lockup">${lockup}</div><div class="line">Play your class better</div>
<div class="crests">${crests}</div></body></html>`;
}

async function buildBanner() {
  const require = createRequire(repo('web/package.json'));
  const { chromium } = require('@playwright/test');
  const browser = await chromium.launch();
  try {
    const page = await browser.newPage({ viewport: BANNER });
    await page.setContent(bannerHtml(), { waitUntil: 'load' });
    await page.screenshot({ path: here('assets/banner.png') });
  } finally {
    await browser.close();
  }
}

buildRasterIcons();
await buildBanner();
console.log('assets built');
