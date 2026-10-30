// web/scripts/build-app-icon.mjs
// Renders the app icon set to apps/shared/brand/ from design/logo/foreversixty-mark.svg (the LX
// seal, vector, unmodified): the iOS AppIcon asset-catalog set and the Android adaptive icon
// layers. The seal is only scaled and placed, never redrawn.
// Run from the repo root: node web/scripts/build-app-icon.mjs
import { mkdirSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import sharp from 'sharp';

const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..', '..');
const MARK_SVG = join(ROOT, 'design/logo/foreversixty-mark.svg');
const OUT = join(ROOT, 'apps/shared/brand');
const SIZE = 1024;
/** tokens.css --color-bg, the near-black the site and the seal's own disc sit on. */
const GROUND = '#07090d';
/** The seal's ring spans 88% of its canvas. Android's adaptive mask keeps a centred circle of
 *  66/108 of the canvas, so the foreground seal is sized to sit inside it. */
const ANDROID_SEAL_RATIO = 0.6;

const renderSeal = (diameter) => sharp(MARK_SVG, { density: 1200 }).resize(diameter, diameter).png().toBuffer();

async function ios() {
  const dir = join(OUT, 'AppIcon.appiconset');
  mkdirSync(dir, { recursive: true });
  const seal = await renderSeal(SIZE);
  await sharp({ create: { width: SIZE, height: SIZE, channels: 3, background: GROUND } })
    .composite([{ input: seal }])
    .removeAlpha()
    .png()
    .toFile(join(dir, 'AppIcon-1024.png'));
  const contents = {
    images: [{ filename: 'AppIcon-1024.png', idiom: 'universal', platform: 'ios', size: '1024x1024' }],
    info: { author: 'xcode', version: 1 },
  };
  writeFileSync(join(dir, 'Contents.json'), `${JSON.stringify(contents, null, 2)}\n`);
}

async function android() {
  const dir = join(OUT, 'android');
  mkdirSync(dir, { recursive: true });
  const diameter = Math.round(SIZE * ANDROID_SEAL_RATIO / 0.88);
  const seal = await renderSeal(diameter);
  await sharp({ create: { width: SIZE, height: SIZE, channels: 4, background: { r: 0, g: 0, b: 0, alpha: 0 } } })
    .composite([{ input: seal, gravity: 'center' }])
    .png()
    .toFile(join(dir, 'ic_launcher_foreground.png'));
  await sharp({ create: { width: SIZE, height: SIZE, channels: 3, background: GROUND } })
    .png()
    .toFile(join(dir, 'ic_launcher_background.png'));
  writeFileSync(
    join(dir, 'ic_launcher.xml'),
    `<?xml version="1.0" encoding="utf-8"?>
<adaptive-icon xmlns:android="http://schemas.android.com/apk/res/android">
    <background android:drawable="@mipmap/ic_launcher_background" />
    <foreground android:drawable="@mipmap/ic_launcher_foreground" />
</adaptive-icon>
`,
  );
}

await ios();
await android();
