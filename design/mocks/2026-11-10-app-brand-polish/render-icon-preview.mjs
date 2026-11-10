// design/mocks/2026-11-10-app-brand-polish/render-icon-preview.mjs
// The review sheet for the painted app icon (lane app-brand-polish, 2026-11-10): the icon before
// (git ref BEFORE, the 30 October placeholder) and after (the files on disk), on iOS at its home
// screen sizes and on Android at 48 dp in every density and launcher mask, beside stand-in
// neighbours. Every icon is drawn from the shipped files themselves; nothing is redrawn here.
// Run from the repo root: node design/mocks/2026-11-10-app-brand-polish/render-icon-preview.mjs [BEFORE]
import { execFileSync } from 'node:child_process';
import { mkdirSync, readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const repo = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..', '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const sharp = require('sharp');

const BEFORE = process.argv[2] ?? '0637711a0';
const OUT = path.join(repo, 'design/reviews/2026-11-10-app-brand-polish');
const IOS_ICON = 'apps/shared/brand/ios/Brand.xcassets/AppIcon.appiconset/AppIcon-1024.png';
const ANDROID_BG = 'apps/shared/brand/android/ic_launcher_background.png';
const ANDROID_FG = 'apps/shared/brand/android/ic_launcher_foreground.png';
const FONT = "font-family='Helvetica Neue, Helvetica, Arial, sans-serif'";
/** iOS home-screen icon sizes in points, each at @2x and @3x: iPhone 60, iPad 76, iPad Pro 83.5. */
const IOS_POINTS = [60, 76, 83.5];
/** Android's 48 dp launcher icon at each density. */
const DENSITIES = [
  ['mdpi', 48],
  ['hdpi', 72],
  ['xhdpi', 96],
  ['xxhdpi', 144],
  ['xxxhdpi', 192],
];
/** Neutral stand-in neighbours: flat tiles in the colours common on a home screen, no marks. */
const NEIGHBOURS = ['#34c759', '#0a84ff', '#ff9f0a', '#f2f2f7', '#1c1c1e', '#ff375f'];

const file = (rel, ref) =>
  ref ? execFileSync('git', ['show', `${ref}:${rel}`], { cwd: repo, maxBuffer: 1 << 26 }) : readFileSync(path.join(repo, rel));
const text = (x, y, s, size = 20, fill = '#e8e2d6', weight = 400, anchor = 'start') =>
  `<text x='${x}' y='${y}' ${FONT} font-size='${size}' font-weight='${weight}' fill='${fill}' text-anchor='${anchor}'>${s}</text>`;

/** The iOS mask: a continuous-corner square approximated by a 22.37% rounded rect. */
const squircle = (s) =>
  `<svg xmlns='http://www.w3.org/2000/svg' width='${s}' height='${s}'><rect width='${s}' height='${s}' rx='${s * 0.2237}' fill='#fff'/></svg>`;
const circle = (s) =>
  `<svg xmlns='http://www.w3.org/2000/svg' width='${s}' height='${s}'><circle cx='${s / 2}' cy='${s / 2}' r='${s / 2}' fill='#fff'/></svg>`;
/** Android's rounded-square launcher mask (Pixel's is a squircle near 30%). */
const roundedSquare = (s) =>
  `<svg xmlns='http://www.w3.org/2000/svg' width='${s}' height='${s}'><rect width='${s}' height='${s}' rx='${s * 0.3}' fill='#fff'/></svg>`;

const masked = async (input, s, mask) =>
  sharp(await sharp(input).resize(s, s, { kernel: 'lanczos3' }).ensureAlpha().png().toBuffer())
    .composite([{ input: Buffer.from(mask(s)), blend: 'dest-in' }])
    .png()
    .toBuffer();

/** The adaptive icon as a launcher shows it: both layers, the middle 72 of 108 dp, at `s` px. */
async function adaptive(ref, s, mask) {
  const bg = file(ANDROID_BG, ref);
  const fg = file(ANDROID_FG, ref);
  const side = (await sharp(bg).metadata()).width;
  const crop = Math.round((side * 72) / 108);
  const at = Math.round((side - crop) / 2);
  const fgSized = await sharp(fg).resize(side, side).png().toBuffer();
  const stack = await sharp(bg).composite([{ input: fgSized }]).png().toBuffer();
  const view = await sharp(stack).extract({ left: at, top: at, width: crop, height: crop }).png().toBuffer();
  return masked(view, s, mask);
}

/** A home-screen row: two neighbours, the icon, two neighbours, at one size. */
function neighbours(svg, x, y, s, gap, mask, skip) {
  for (let i = 0; i < 4; i++) {
    const nx = x + (i < 2 ? i : i + 1) * (s + gap);
    const rx = mask === 'circle' ? s / 2 : s * (mask === 'squircle' ? 0.2237 : 0.3);
    if (i !== skip) svg.push(`<rect x='${nx}' y='${y}' width='${s}' height='${s}' rx='${rx}' fill='${NEIGHBOURS[(i + s) % NEIGHBOURS.length]}'/>`);
  }
  return x + 2 * (s + gap);
}

async function iosSection(top) {
  const svg = [text(40, top + 34, 'iOS, Any appearance: before (left) and after (right); size rows at 1/2.4 of device pixels', 26, '#f3ead8', 600)];
  const parts = [];
  let y = top + 60;
  for (const [ref, label, x0] of [[BEFORE, `Before (${BEFORE})`, 40], [null, 'After', 840]]) {
    const icon = file(IOS_ICON, ref);
    parts.push({ input: await masked(icon, 360, squircle), left: x0 + 200, top: y });
    svg.push(text(x0 + 380, y + 396, label, 20, '#c9c6c0', 500, 'middle'));
  }
  y += 430;
  // Every home-screen size at @2x and @3x on a wallpaper, between neighbours.
  for (const [wallLabel, wall, ink] of [['Light wallpaper', '#cfd9e6', '#1d1d1f'], ['Dark wallpaper', '#10141c', '#f2f2f7']]) {
    for (const [ref, x0] of [[BEFORE, 40], [null, 840]]) {
      svg.push(`<rect x='${x0}' y='${y}' width='760' height='300' rx='24' fill='${wall}'/>`);
      svg.push(text(x0 + 20, y + 30, wallLabel, 16, ink));
      let x = x0 + 20;
      for (const pt of IOS_POINTS) {
        for (const scale of [2, 3]) {
          const s = Math.round((pt * scale) / 2.4);
          parts.push({ input: await masked(file(IOS_ICON, ref), s, squircle), left: x, top: y + 60 });
          svg.push(text(x + s / 2, y + 60 + s + 22, `${pt}@${scale}x`, 13, ink, 400, 'middle'));
          x += s + 14;
        }
      }
      // The real iPhone size (60 pt @3x = 180 px) shown at half scale in a row of neighbours.
      const ix = neighbours(svg, x0 + 20, y + 190, 60, 26, 'squircle', -1);
      parts.push({ input: await masked(file(IOS_ICON, ref), 60, squircle), left: ix, top: y + 190 });
      svg.push(text(ix + 30, y + 270, 'Forever Sixty', 12, ink, 500, 'middle'));
    }
    y += 320;
  }
  return { svg: svg.join(''), parts, bottom: y };
}

async function androidSection(top) {
  const svg = [text(40, top + 34, 'Android adaptive icon, 48 dp at every density: before (left) and after (right)', 26, '#f3ead8', 600)];
  const parts = [];
  let y = top + 60;
  for (const [maskLabel, mask] of [['Circle mask', circle], ['Rounded-square mask', roundedSquare]]) {
    for (const [wallLabel, wall, ink] of [['Light', '#eef0f4', '#202124'], ['Dark', '#121316', '#e8eaed']]) {
      for (const [ref, x0] of [[BEFORE, 40], [null, 840]]) {
        svg.push(`<rect x='${x0}' y='${y}' width='760' height='300' rx='24' fill='${wall}'/>`);
        svg.push(text(x0 + 20, y + 30, `${maskLabel}, ${wallLabel.toLowerCase()} wallpaper`, 16, ink));
        let x = x0 + 20;
        for (const [name, px] of DENSITIES) {
          parts.push({ input: await adaptive(ref, px, mask), left: x, top: y + 50 });
          svg.push(text(x + px / 2, y + 50 + 192 + 26, `${name} ${px}px`, 13, ink, 400, 'middle'));
          x += px + 20;
        }
      }
      y += 320;
    }
  }
  // A launcher row at xxhdpi (48 dp = 144 px, shown at half) between neighbours, after only.
  svg.push(`<rect x='40' y='${y}' width='1560' height='150' rx='24' fill='#1f2a36'/>`);
  const ix = neighbours(svg, 60, y + 30, 72, 40, 'circle', -1);
  parts.push({ input: await adaptive(null, 72, circle), left: ix, top: y + 30 });
  svg.push(text(ix + 36, y + 128, 'Forever Sixty', 14, '#e8eaed', 500, 'middle'));
  return { svg: svg.join(''), parts, bottom: y + 170 };
}

async function main() {
  mkdirSync(OUT, { recursive: true });
  const sections = [];
  let top = 20;
  for (const build of [iosSection, androidSection]) {
    const section = await build(top);
    sections.push(section);
    top = section.bottom + 30;
  }
  const width = 1640;
  const frame = `<svg xmlns='http://www.w3.org/2000/svg' width='${width}' height='${top}'><rect width='100%' height='100%' fill='#16130f'/>${sections.map((s) => s.svg).join('')}</svg>`;
  await sharp(Buffer.from(frame))
    .composite(sections.flatMap((s) => s.parts))
    .png()
    .toFile(path.join(OUT, 'icon-preview.png'));
  console.log(`wrote ${path.relative(repo, path.join(OUT, 'icon-preview.png'))}`);
}

await main();
