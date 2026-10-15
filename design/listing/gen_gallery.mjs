// design/listing/gen_gallery.mjs
// Renders the store gallery plates for the Forever Sixty listing: each client capture the
// /setup page uses (web/public/addon/*.jpg) on a 1600x900 plate in the site's own dark and
// gold, with the surface's name and one line under it. Run from anywhere:
//   node design/listing/gen_gallery.mjs
// Output: design/listing/gallery/NN-<name>.png. Playwright comes from web/node_modules.
import { mkdir, readFile } from 'node:fs/promises';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const here = path.dirname(fileURLToPath(import.meta.url));
const repo = path.resolve(here, '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const { chromium } = require('@playwright/test');

const CAPTURES = path.join(repo, 'web', 'public', 'addon');
const FONTS = path.join(repo, 'web', 'src', 'assets', 'fonts');
const OUT = path.join(here, 'gallery');
const WIDTH = 1600;
const HEIGHT = 900;
/** The capture's box above the caption; a capture scales to fit it, up to 1.5x its pixels. */
const STAGE_HEIGHT = 600;
const STAGE_WIDTH = 1440;
const MAX_SCALE = 1.5;

/** Order and captions: design/listing/curseforge-description.md, "Gallery captions". */
const PLATES = [
  ['chip-tooltip-legend', 'The chip', 'One line beside the minimap: the talent to take, the bag item to equip, loot here for you, or what the ready check found missing.'],
  ['talent-marks-trees', 'Talent marks', "Your build's talents numbered in order, the next one ringed, a strip counting the points your level allows."],
  ['item-tooltip-ranged', 'Item tooltips', 'Whether an item beats what you wear and by how much, the next upgrade for the slot, where it drops, and two alternatives.'],
  ['bag-arrow', 'Paper doll and bags', 'A green arrow on the slot and the bag item that beats it; a gold corner on a slot worth chasing.'],
  ['character-panel-tab', 'The character-panel tab', "Your next five gear targets with their sources, your rotation at your level, and your build's progress."],
  ['player-tooltip-score', 'The player tooltip', "Every player's Forever Score, with the five parts behind it on Shift."],
  ['chat-badge-guild', 'Chat badges', "A sender's score in front of their line in guild, trade and whispers."],
  ['guild-tab', 'The guild tab', "Your roster's level, spec, gear, score, attendance and gear gap, sortable, live for members who run the addon."],
  ['group-panel-solo', 'The group panel', "Each member's spec, gear gap and score before the first pull; a ready check names who lacks flask, food, imbue or a class buff."],
  ['build-dialog-paste', 'Paste a build', 'The one dialog of its own: paste a build from foreversixty.gg and follow it.'],
  ['build-dialog-copy', 'Copy your character', "Ctrl+C sends your talents, gear and bags to the site's planner and simulator."],
  ['settings-switches', 'Settings', "Eight switches on the game's own Settings page, each with one line saying what it does."],
];

// A page set from a string has no file access, so fonts and captures travel inline.
const dataUrl = async (file, type) => `data:${type};base64,${(await readFile(file)).toString('base64')}`;
const font = (file) => dataUrl(path.join(FONTS, file), 'font/woff2');
const capture = (name) => dataUrl(path.join(CAPTURES, `${name}.jpg`), 'image/jpeg');

async function plateHtml(name, title, caption) {
  const [cinzel, barlow400, barlow600, shot] = await Promise.all([
    font('cinzel-700-subset.woff2'),
    font('barlow-400-subset.woff2'),
    font('barlow-600-subset.woff2'),
    capture(name),
  ]);
  return `<!doctype html>
<html lang="en"><head><meta charset="utf-8">
<style>
@font-face{font-family:Cinzel;src:url(${cinzel}) format('woff2');font-weight:700}
@font-face{font-family:Barlow;src:url(${barlow400}) format('woff2');font-weight:400}
@font-face{font-family:Barlow;src:url(${barlow600}) format('woff2');font-weight:600}
html,body{margin:0;width:${WIDTH}px;height:${HEIGHT}px;overflow:hidden}
body{background:#07090d;color:#e9e4d8;font-family:Barlow,sans-serif;display:grid;grid-template-rows:1fr auto;box-sizing:border-box;padding:64px 80px 56px;
  background-image:radial-gradient(ellipse 70% 60% at 50% 42%,rgba(229,185,85,.07),transparent 70%)}
.stage{display:grid;place-items:center;height:${STAGE_HEIGHT}px}
.shot{display:block;border:1px solid #262e40;border-radius:6px;box-shadow:0 24px 60px rgba(0,0,0,.65)}
.foot{display:grid;grid-template-columns:auto 1fr;gap:28px;align-items:center;margin-top:40px;padding-top:28px;border-top:1px solid #262e40}
.mark{display:flex;align-items:center;gap:12px;font-family:Cinzel,serif;font-weight:700;font-size:22px;letter-spacing:.1em;text-transform:uppercase;color:#f2eee4;white-space:nowrap}
.text{min-width:0}
h1{margin:0 0 6px;font-family:Barlow,sans-serif;font-weight:600;font-size:30px;line-height:1.1;color:#e5b955}
p{margin:0;font-size:20px;line-height:1.35;color:#b9b3a4;max-width:1100px}
</style></head><body>
<div class="stage"><img class="shot" src="${shot}" alt=""></div>
<div class="foot">
  <div class="mark"><svg width="34" height="34" viewBox="0 0 256 256" aria-hidden="true"><defs><linearGradient id="g" x1="0" y1="0" x2="0" y2="1"><stop offset="0" stop-color="#fbe7a1"/><stop offset=".45" stop-color="#e5b955"/><stop offset="1" stop-color="#a8762a"/></linearGradient></defs><circle cx="128" cy="128" r="112.64" fill="#0d111a"/><circle cx="128" cy="128" r="112.64" fill="none" stroke="url(#g)" stroke-width="12.8"/><text x="128" y="152" text-anchor="middle" font-family="Cinzel,serif" font-weight="700" font-size="72" fill="url(#g)">LX</text></svg>Forever Sixty</div>
  <div class="text"><h1>${title}</h1><p>${caption}</p></div>
</div>
</body></html>`;
}

async function main() {
  await mkdir(OUT, { recursive: true });
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: WIDTH, height: HEIGHT }, deviceScaleFactor: 1 });
  for (const [index, [name, title, caption]] of PLATES.entries()) {
    await page.setContent(await plateHtml(name, title, caption), { waitUntil: 'load' });
    await page.evaluate(
      ({ stageWidth, stageHeight, maxScale }) => {
        const shot = document.querySelector('.shot');
        if (!(shot instanceof HTMLImageElement) || shot.naturalWidth === 0) throw new Error('capture did not load');
        const scale = Math.min(maxScale, stageWidth / shot.naturalWidth, stageHeight / shot.naturalHeight);
        shot.style.width = `${Math.round(shot.naturalWidth * scale)}px`;
        shot.style.height = 'auto';
        return document.fonts.ready;
      },
      { stageWidth: STAGE_WIDTH, stageHeight: STAGE_HEIGHT, maxScale: MAX_SCALE },
    );
    const file = path.join(OUT, `${String(index + 1).padStart(2, '0')}-${name}.png`);
    await page.screenshot({ path: file, type: 'png' });
    process.stdout.write(`${path.relative(repo, file)}\n`);
  }
  await browser.close();
}

main().catch((error) => {
  process.stderr.write(`${error instanceof Error ? error.stack : String(error)}\n`);
  process.exitCode = 1;
});
