// web/scripts/build-app-icon.mjs
// Renders the app icon set to apps/shared/brand/ from design/logo/foreversixty-mark.svg (the LX
// seal, vector, unmodified): the iOS AppIcon asset-catalog set and the Android adaptive icon
// layers. The seal is only scaled and placed, never redrawn.
// The iOS Dark and Tinted appearances come from the flat seal
// (design/assets/brand/seal-flat/seal-flat-small.svg); the light (Any) appearance stays the
// painted seal. The Android themed layer is a vector drawable from tools/gen_seal_silhouette.py.
// Run from the repo root: node web/scripts/build-app-icon.mjs
import { mkdirSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import sharp from 'sharp';

const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..', '..');
const MARK_SVG = join(ROOT, 'design/logo/foreversixty-mark.svg');
const FLAT_SVG = join(ROOT, 'design/assets/brand/seal-flat/seal-flat-small.svg');
const OUT = join(ROOT, 'apps/shared/brand');
const SIZE = 1024;
/** tokens.css --color-bg, the near-black the site and the seal's own disc sit on. */
const GROUND = '#07090d';
/** The seal's ring spans 88% of its canvas. Android's adaptive mask keeps a centred circle of
 *  66/108 of the canvas, so the foreground seal is sized to sit inside it. */
const ANDROID_SEAL_RATIO = 0.6;
/** The painted seal's outer edge (ring radius plus half its stroke) spans 93% of its canvas;
 *  the flat seal fills its own, so it is drawn at this share to sit where the painted one does. */
const FLAT_SEAL_RATIO = 0.93;
/** Tinted icons are greyscale: the levels stretch the flat seal's outline (#170E09, luma 16) to
 *  black and its gold (#C9A465, luma 168) to white, so the tint lands on the ring and the LX. */
const TINT_LEVELS = { black: 16, white: 168 };
const TRANSPARENT = { r: 0, g: 0, b: 0, alpha: 0 };

const renderSeal = (diameter) =>
  sharp(MARK_SVG, { density: 1200 }).resize(diameter, diameter).png().toBuffer();

const renderFlatSeal = (diameter) =>
  sharp(FLAT_SVG, { density: 1200 }).resize(diameter, diameter).png().toBuffer();

const transparentCanvas = () =>
  sharp({ create: { width: SIZE, height: SIZE, channels: 4, background: TRANSPARENT } });

async function iosAny(file) {
  const seal = await renderSeal(SIZE);
  await sharp({ create: { width: SIZE, height: SIZE, channels: 3, background: GROUND } })
    .composite([{ input: seal }])
    .removeAlpha()
    .png()
    .toFile(file);
}

/** Apple's dark icons sit on a transparent ground; the system draws the dark backdrop. */
async function iosDark(file) {
  const seal = await renderFlatSeal(Math.round(SIZE * FLAT_SEAL_RATIO));
  await transparentCanvas()
    .composite([{ input: seal, gravity: 'center' }])
    .png()
    .toFile(file);
}

/** Apple's tinted icons are a greyscale image on a transparent ground; the system tints it. */
async function iosTinted(file) {
  const { black, white } = TINT_LEVELS;
  const gain = 255 / (white - black);
  const seal = await sharp(await renderFlatSeal(Math.round(SIZE * FLAT_SEAL_RATIO)))
    .greyscale()
    .linear(gain, -black * gain)
    .png()
    .toBuffer();
  await transparentCanvas()
    .composite([{ input: seal, gravity: 'center' }])
    .png()
    .toFile(file);
}

/** The iOS appearances: the painted seal for Any, the flat seal for Dark and Tinted. */
const IOS_ICONS = [
  { filename: 'AppIcon-1024.png', render: iosAny },
  { filename: 'AppIcon-Dark-1024.png', luminosity: 'dark', render: iosDark },
  { filename: 'AppIcon-Tinted-1024.png', luminosity: 'tinted', render: iosTinted },
];

async function ios() {
  // A whole asset catalog, so the iOS target compiles it in place (actool skips symlinks).
  const catalog = join(OUT, 'ios', 'Brand.xcassets');
  const dir = join(catalog, 'AppIcon.appiconset');
  mkdirSync(dir, { recursive: true });
  writeFileSync(
    join(catalog, 'Contents.json'),
    `${JSON.stringify({ info: { author: 'xcode', version: 1 } }, null, 2)}\n`,
  );
  for (const { filename, render } of IOS_ICONS) {
    await render(join(dir, filename));
  }
  const contents = {
    images: IOS_ICONS.map(({ filename, luminosity }) => ({
      ...(luminosity && { appearances: [{ appearance: 'luminosity', value: luminosity }] }),
      filename,
      idiom: 'universal',
      platform: 'ios',
      size: '1024x1024',
    })),
    info: { author: 'xcode', version: 1 },
  };
  writeFileSync(join(dir, 'Contents.json'), `${JSON.stringify(contents, null, 2)}\n`);
}

async function android() {
  const dir = join(OUT, 'android');
  mkdirSync(dir, { recursive: true });
  const diameter = Math.round((SIZE * ANDROID_SEAL_RATIO) / 0.88);
  const seal = await renderSeal(diameter);
  await transparentCanvas()
    .composite([{ input: seal, gravity: 'center' }])
    .png()
    .toFile(join(dir, 'ic_launcher_foreground.png'));
  await sharp({ create: { width: SIZE, height: SIZE, channels: 3, background: GROUND } })
    .png()
    .toFile(join(dir, 'ic_launcher_background.png'));
  // The monochrome layer is the app's res/drawable/ic_launcher_monochrome.xml
  // (tools/gen_seal_silhouette.py), which launchers tint for themed icons on Android 13+.
  writeFileSync(
    join(dir, 'ic_launcher.xml'),
    `<?xml version="1.0" encoding="utf-8"?>
<adaptive-icon xmlns:android="http://schemas.android.com/apk/res/android">
    <background android:drawable="@mipmap/ic_launcher_background" />
    <foreground android:drawable="@mipmap/ic_launcher_foreground" />
    <monochrome android:drawable="@drawable/ic_launcher_monochrome" />
</adaptive-icon>
`,
  );
}

await ios();
await android();
