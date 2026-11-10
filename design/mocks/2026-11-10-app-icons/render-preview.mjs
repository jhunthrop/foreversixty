// design/mocks/2026-11-10-app-icons/render-preview.mjs
// The review sheet for the app icon variants (lane app-icon-variants, 2026-11-10). Every icon is
// drawn from the shipped asset itself: the Android drawables' pathData and the iOS PNGs.
// Run from the repo root: node design/mocks/2026-11-10-app-icons/render-preview.mjs
import { mkdirSync, readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const repo = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..', '..', '..');
const require = createRequire(path.join(repo, 'web', 'package.json'));
const sharp = require('sharp');

const OUT = path.join(repo, 'design/reviews/2026-11-10-app-icons');
const DRAWABLE = path.join(repo, 'apps/android/app/src/main/res/drawable');
const APPICON = path.join(repo, 'apps/shared/brand/ios/Brand.xcassets/AppIcon.appiconset');
const ACCENT = '#2C8A9E';
const FONT = "font-family='Helvetica Neue, Helvetica, Arial, sans-serif'";

/** A VectorDrawable's single path as an SVG drawn in `ink`, at `px` square. */
function drawable(name, ink, px, viewBox) {
  const xml = readFileSync(path.join(DRAWABLE, name), 'utf8');
  const vp = xml.match(/viewportWidth="([\d.]+)"/)[1];
  const d = xml.match(/pathData="([^"]+)"/)[1];
  const box = viewBox ?? `0 0 ${vp} ${vp}`;
  return `<svg xmlns='http://www.w3.org/2000/svg' width='${px}' height='${px}' viewBox='${box}'><path fill='${ink}' d='${d}'/></svg>`;
}

const png = (svg) => sharp(Buffer.from(svg)).png().toBuffer();
const text = (x, y, s, size = 22, fill = '#e8e2d6', weight = 400, anchor = 'start') =>
  `<text x='${x}' y='${y}' ${FONT} font-size='${size}' font-weight='${weight}' fill='${fill}' text-anchor='${anchor}'>${s}</text>`;

/** Section 1: the small icon at every density, in light and dark status bars and the shade. */
async function notificationSection(top) {
  const parts = [];
  const svg = [
    text(40, top + 34, 'Android notification small icon (ic_notification_seal, 24 dp)', 26, '#f3ead8', 600),
  ];
  const densities = [
    ['mdpi', 24],
    ['hdpi', 36],
    ['xhdpi', 48],
    ['xxhdpi', 72],
    ['xxxhdpi', 96],
  ];
  const bars = [
    ['Dark status bar', '#000000', '#ffffff'],
    ['Light status bar', '#f4f1ec', '#1f1f1f'],
  ];
  let y = top + 60;
  for (const [label, bg, ink] of bars) {
    svg.push(`<rect x='40' y='${y}' width='1520' height='150' rx='14' fill='${bg}'/>`);
    svg.push(text(64, y + 32, label, 18, ink === '#ffffff' ? '#9a9a9a' : '#5f5f5f'));
    let x = 64;
    for (const [name, px] of densities) {
      parts.push({
        input: await png(drawable('ic_notification_seal.xml', ink, px)),
        left: x,
        top: y + 118 - px,
      });
      svg.push(
        text(
          x + px / 2,
          y + 140,
          `${name} ${px}px`,
          14,
          ink === '#ffffff' ? '#9a9a9a' : '#5f5f5f',
          400,
          'middle',
        ),
      );
      x += Math.max(px, 90) + 40;
    }
    // A status bar row at xxhdpi: the clock, then the seal beside two system glyph stand-ins.
    const bx = 900;
    svg.push(text(bx, y + 96, '9:41', 42, ink, 500));
    parts.push({
      input: await png(drawable('ic_notification_seal.xml', ink, 72)),
      left: bx + 110,
      top: y + 46,
    });
    parts.push({
      input: await png(drawable('ic_notification_seal.xml', ink, 72)),
      left: bx + 190,
      top: y + 46,
    });
    svg.push(
      text(
        bx + 280,
        y + 96,
        '(two pushes, xxhdpi, actual size)',
        18,
        ink === '#ffffff' ? '#9a9a9a' : '#5f5f5f',
      ),
    );
    y += 170;
  }
  // The expanded shade header: the icon in the turquoise accent with the app name.
  for (const [i, [bg, fg, sub]] of [
    ['#1c1b1f', '#e6e1e5', '#a9a4ab'],
    ['#fbf8f4', '#1c1b1f', '#5f5b62'],
  ].entries()) {
    const x = 40 + i * 770;
    svg.push(`<rect x='${x}' y='${y}' width='750' height='170' rx='28' fill='${bg}'/>`);
    parts.push({
      input: await png(drawable('ic_notification_seal.xml', ACCENT, 48)),
      left: x + 32,
      top: y + 28,
    });
    svg.push(text(x + 96, y + 60, `Forever Sixty · Raids · now`, 22, ACCENT, 500));
    svg.push(text(x + 32, y + 112, 'Molten Core tonight, 19:30', 28, fg, 600));
    svg.push(text(x + 32, y + 148, 'Invites in 30 minutes. You are In as Holy.', 22, sub));
  }
  y += 190;
  return { svg: svg.join(''), parts, bottom: y };
}

/** Section 2: the themed launcher icon in Material You tints, light and dark. */
async function themedSection(top) {
  const svg = [
    text(
      40,
      top + 34,
      'Android themed icon (adaptive &lt;monochrome&gt;, 66 dp safe zone)',
      26,
      '#f3ead8',
      600,
    ),
  ];
  const parts = [];
  const tints = [
    ['Blue', '#d6e3ff', '#0a305f', '#284777', '#d6e3ff'],
    ['Green', '#c4efc1', '#0c3a14', '#25502a', '#c4efc1'],
    ['Rose', '#ffd9e2', '#3e0021', '#5b1f36', '#ffd9e2'],
    ['Sand', '#f6e0bd', '#3a2a08', '#52401d', '#f6e0bd'],
  ];
  const size = 150;
  // The launcher shows the middle 72 dp of the 108 dp canvas, masked to a circle.
  const crop = '18 18 72 72';
  let x = 40;
  const y = top + 60;
  for (const [name, lightBg, lightFg, darkBg, darkFg] of tints) {
    for (const [mode, bg, fg, panel] of [
      ['light', lightBg, lightFg, '#f2efe9'],
      ['dark', darkBg, darkFg, '#121316'],
    ]) {
      svg.push(`<rect x='${x}' y='${y}' width='180' height='230' rx='16' fill='${panel}'/>`);
      svg.push(`<circle cx='${x + 90}' cy='${y + 100}' r='${size / 2}' fill='${bg}'/>`);
      parts.push({
        input: await png(drawable('ic_launcher_monochrome.xml', fg, size, crop)),
        left: x + 15,
        top: y + 25,
      });
      svg.push(
        text(
          x + 90,
          y + 210,
          `${name}, ${mode}`,
          16,
          mode === 'light' ? '#3a3a3a' : '#c9c6c0',
          400,
          'middle',
        ),
      );
      x += 190;
    }
  }
  // The safe-zone proof: the full 108 dp canvas with the 66 dp circle drawn over it.
  const proof = 230;
  const px = 40;
  const py = y + 250;
  svg.push(`<rect x='${px}' y='${py}' width='${proof}' height='${proof}' fill='#2a2620'/>`);
  parts.push({
    input: await png(drawable('ic_launcher_monochrome.xml', '#e8e2d6', proof)),
    left: px,
    top: py,
  });
  svg.push(
    `<circle cx='${px + proof / 2}' cy='${py + proof / 2}' r='${(proof * 33) / 108}' fill='none' stroke='#e0574f' stroke-width='2' stroke-dasharray='6 5'/>`,
  );
  svg.push(
    text(
      px + proof + 24,
      py + 40,
      'Full 108 dp canvas; dashed circle is the 66 dp safe zone.',
      18,
      '#c9c6c0',
    ),
  );
  svg.push(text(px + proof + 24, py + 68, 'The seal is 64 dp across, centred.', 18, '#c9c6c0'));
  return { svg: svg.join(''), parts, bottom: py + proof + 30 };
}

/** The iOS icon mask: a continuous-corner square approximated by a 22.37% rounded rect. */
const squircle = (s) =>
  `<svg xmlns='http://www.w3.org/2000/svg' width='${s}' height='${s}'><rect width='${s}' height='${s}' rx='${s * 0.2237}' fill='#fff'/></svg>`;

async function iosIcon(appearance, s) {
  const file = {
    light: 'AppIcon-1024.png',
    dark: 'AppIcon-Dark-1024.png',
    tinted: 'AppIcon-Tinted-1024.png',
  }[appearance];
  const art = await sharp(path.join(APPICON, file))
    .resize(s, s)
    .ensureAlpha()
    .raw()
    .toBuffer({ resolveWithObject: true });
  let layer;
  if (appearance === 'tinted') {
    // The system's tint: the greyscale art's brightness becomes the tint colour's coverage.
    const { data } = art;
    const out = Buffer.alloc(s * s * 4);
    const tint = [0xe8, 0xb8, 0x6a];
    for (let i = 0; i < s * s; i++) {
      const a = (data[i * 4] / 255) * data[i * 4 + 3];
      out.set([...tint, Math.round(a)], i * 4);
    }
    layer = await sharp(out, { raw: { width: s, height: s, channels: 4 } })
      .png()
      .toBuffer();
  } else {
    layer = await sharp(art.data, { raw: art.info }).png().toBuffer();
  }
  // Dark and tinted icons sit on the system's dark backdrop.
  const backdrop =
    appearance === 'light'
      ? `<rect width='${s}' height='${s}' fill='#000'/>`
      : `<defs><linearGradient id='g' x1='0' y1='0' x2='0' y2='1'><stop offset='0' stop-color='${appearance === 'tinted' ? '#2e2a24' : '#2f2f31'}'/><stop offset='1' stop-color='${appearance === 'tinted' ? '#12100d' : '#141415'}'/></linearGradient></defs><rect width='${s}' height='${s}' fill='url(#g)'/>`;
  const base = await png(
    `<svg xmlns='http://www.w3.org/2000/svg' width='${s}' height='${s}'>${backdrop}</svg>`,
  );
  const icon = await sharp(base)
    .composite([{ input: layer }])
    .png()
    .toBuffer();
  return sharp(icon)
    .composite([{ input: Buffer.from(squircle(s)), blend: 'dest-in' }])
    .png()
    .toBuffer();
}

/** Section 3: the three iOS appearances on mock home screens, the painted light one first. */
async function iosSection(top) {
  const svg = [
    text(40, top + 34, 'iOS app icon appearances (AppIcon, single size 1024)', 26, '#f3ead8', 600),
  ];
  const parts = [];
  const screens = [
    ['light', 'Any (light): painted seal, unchanged', ['#cfd9e6', '#9fb3cc'], '#1d1d1f'],
    ['dark', 'Dark: flat seal on transparent', ['#1d2433', '#05070b'], '#f2f2f7'],
    ['tinted', 'Tinted: greyscale flat seal, amber tint', ['#3a2e1c', '#0d0a06'], '#f2e6cf'],
  ];
  const w = 500;
  const h = 540;
  for (const [i, [appearance, label, [g0, g1], ink]] of screens.entries()) {
    const x = 40 + i * (w + 20);
    const y = top + 60;
    svg.push(
      `<defs><linearGradient id='w${i}' x1='0' y1='0' x2='0.4' y2='1'><stop offset='0' stop-color='${g0}'/><stop offset='1' stop-color='${g1}'/></linearGradient></defs>`,
    );
    svg.push(`<rect x='${x}' y='${y}' width='${w}' height='${h}' rx='36' fill='url(#w${i})'/>`);
    svg.push(text(x + 28, y + 48, '9:41', 24, ink, 600));
    // A large view of the icon, then a home-screen row at the real 60 pt @2x size.
    parts.push({ input: await iosIcon(appearance, 260), left: x + (w - 260) / 2, top: y + 70 });
    for (const [j, col] of [0, 1, 2].entries()) {
      const cx = x + 70 + col * 140;
      if (j === 1) {
        parts.push({ input: await iosIcon(appearance, 120), left: cx, top: y + 360 });
        svg.push(text(cx + 60, y + 508, 'Forever Sixty', 18, ink, 500, 'middle'));
      } else {
        svg.push(
          `<rect x='${cx}' y='${y + 360}' width='120' height='120' rx='26.8' fill='${ink}' fill-opacity='0.12'/>`,
        );
      }
    }
    svg.push(text(x + 28, y + h + 34, label, 18, '#c9c6c0'));
  }
  return { svg: svg.join(''), parts, bottom: top + 60 + h + 60 };
}

async function main() {
  mkdirSync(OUT, { recursive: true });
  const sections = [];
  let top = 20;
  for (const build of [notificationSection, themedSection, iosSection]) {
    const section = await build(top);
    sections.push(section);
    top = section.bottom + 30;
  }
  const width = 1600;
  const frame = `<svg xmlns='http://www.w3.org/2000/svg' width='${width}' height='${top}'><rect width='100%' height='100%' fill='#16130f'/>${sections.map((s) => s.svg).join('')}</svg>`;
  await sharp(Buffer.from(frame))
    .composite(sections.flatMap((s) => s.parts))
    .png()
    .toFile(path.join(OUT, 'preview.png'));
  console.log(`wrote ${path.relative(repo, path.join(OUT, 'preview.png'))}`);
}

await main();
