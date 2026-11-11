// design/community/discord/build-assets.mjs
// Builds every Discord image from the brand sources, so none is placed by hand:
//   assets/icon.png       server icon: the live seal v4 at 512 px (owner 2026-11-11: the server
//                         icon stays the seal, not the tile), copied from LIVE.icon
//   assets/avatar.png     the bot's avatar (Developer Portal app icon): the same tile at 1024 px
//   assets/premium.png    Premium role icon: the flat seal's tiny bold drawing at 64 px (Discord
//                         shows it at about 20 px beside a name, the size that drawing is for)
//   assets/crest-<class>.png  class role icons and onboarding emoji: the HD crests at 128 px
//   assets/<faction>.png  onboarding emoji: the faction logos at 128 px
//   assets/banner.jpg     server banner, 1920x1080: the live seal v4 banner, copied from LIVE.banner
//   assets/splash.jpg     invite background, 1920x1080: the same with "Forever Sixty" in Marcellus
//                         (the layout is design/assets/art/discord-banner.html)
// Run from anywhere: node design/community/discord/build-assets.mjs
// sharp and @playwright/test resolve from web/node_modules. Deterministic for the PNGs; the banner
// is a browser render.
import { copyFileSync, mkdirSync, readFileSync, statSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import { fileURLToPath } from 'node:url';
import { PNG, renderFlatSeal, renderTile } from '../../../web/scripts/brand-art.mjs';

const here = (p) => fileURLToPath(new URL(p, import.meta.url));
const repo = (p) => here(`../../../${p}`);
const require = createRequire(repo('web/package.json'));
const sharp = require('sharp');
const classes = JSON.parse(readFileSync(repo('web/src/data/classes.json'), 'utf8'));

const FACTIONS = ['alliance', 'horde'];
/** Discord's sizes: the server icon at 512 (its recommended upload), the app icon at 1024, role
 *  icons at 64 (its minimum), emoji at 128 (shown at 48 at most). */
const SIZE = { icon: 512, avatar: 1024, roleIcon: 64, emoji: 128 };
/** The banner layout is 960x540 (Discord's minimum, 16:9) drawn at 2x. */
const BANNER = { width: 960, height: 540, scale: 2, quality: 86 };
/** Discord's upload ceilings: 10 MB for the banner and splash, 256 KB for a role icon or emoji. */
const MAX_BYTES = { banner: 10 * 1024 * 1024, roleIcon: 256 * 1024, emoji: 256 * 1024 };
const ASSETS = here('assets');

const out = (name) => `${ASSETS}/${name}`;

function checkSize(name, limit) {
  const bytes = statSync(out(name)).size;
  if (bytes > limit) throw new Error(`${name} is ${bytes} bytes, over Discord's ${limit}`);
  console.log(`${name.padEnd(20)} ${String(Math.round(bytes / 1024)).padStart(5)} KB  ok`);
}

/** What is live on the server and stays (owner 2026-11-11): the seal v4 icon and its banner. */
const LIVE = {
  icon: 'design/assets/logo/discord-icon-seal-v4-512.png',
  banner: 'design/assets/art/discord-banner-seal-v4.jpg',
};

async function buildIdentity() {
  copyFileSync(repo(LIVE.icon), out('icon.png'));
  await sharp(await renderTile(SIZE.avatar)).png(PNG).toFile(out('avatar.png'));
  await sharp(await renderFlatSeal(SIZE.roleIcon, 'tiny')).png(PNG).toFile(out('premium.png'));
  checkSize('icon.png', MAX_BYTES.banner);
  checkSize('avatar.png', MAX_BYTES.banner);
  checkSize('premium.png', MAX_BYTES.roleIcon);
}

async function webpToPng(src, name) {
  await sharp(repo(src)).resize(SIZE.emoji, SIZE.emoji, { kernel: 'lanczos3' }).png(PNG).toFile(out(name));
  checkSize(name, MAX_BYTES.emoji);
}

async function buildEmoji() {
  for (const { slug } of classes) await webpToPng(`web/public/icons/hd/crests/${slug}.webp`, `crest-${slug}.png`);
  for (const faction of FACTIONS) await webpToPng(`web/public/icons/hd/faction/${faction}-logo-512.webp`, `${faction}.png`);
}

const dataUri = (path, type) => `data:${type};base64,${readFileSync(path).toString('base64')}`;

/** Over the splash layout: a lighter veil, so the room reads at Discord's 240 px sidebar size. */
const VEIL = '.v{background:radial-gradient(ellipse at center,rgba(23,14,9,.1),rgba(23,14,9,.55))}';

/** The banner layout with its three files inlined, so the render needs no server. */
function bannerHtml(extraCss) {
  return readFileSync(repo('design/assets/art/discord-banner.html'), 'utf8')
    .replace('url(marcellus.woff2)', `url(${dataUri(repo('apps/shared/fonts/Marcellus-Regular.ttf'), 'font/ttf')})`)
    .replace('url(tavern.jpeg)', `url(${dataUri(repo('design/assets/art/discord-banner-tavern.jpeg'), 'image/jpeg')})`)
    .replace('src="seal.png"', `src="${dataUri(repo('design/assets/logo/lx-seal-v4-512.png'), 'image/png')}"`)
    .replace('</style>', `${extraCss}</style>`);
}

async function buildBanners() {
  const { chromium } = require('@playwright/test');
  const browser = await chromium.launch();
  try {
    const page = await browser.newPage({
      viewport: { width: BANNER.width, height: BANNER.height },
      deviceScaleFactor: BANNER.scale,
    });
    copyFileSync(repo(LIVE.banner), out('banner.jpg'));
    checkSize('banner.jpg', MAX_BYTES.banner);
    await page.setContent(bannerHtml(VEIL), { waitUntil: 'load' });
    await page.evaluate(() => document.fonts.ready);
    const shot = await page.locator('.b').screenshot({ type: 'png' });
    writeFileSync(out('splash.jpg'), await sharp(shot).jpeg({ quality: BANNER.quality, mozjpeg: true }).toBuffer());
    checkSize('splash.jpg', MAX_BYTES.banner);
  } finally {
    await browser.close();
  }
}

mkdirSync(ASSETS, { recursive: true });
await buildIdentity();
await buildEmoji();
await buildBanners();
