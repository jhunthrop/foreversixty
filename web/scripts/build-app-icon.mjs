// web/scripts/build-app-icon.mjs
// Renders the app icon set to apps/shared/brand/: the iOS AppIcon asset-catalog set and the Android
// launcher icon (adaptive layers, the round alias and the legacy density PNGs).
// The light (Any) iOS icon and the Android icon are one tile: the owner's painted blue ground
// (design/assets/art/app-icon-background.jpeg, inside its brass frame, as the companion's macOS
// icon) with the painted seal v4 (design/assets/logo/lx-seal-v4.png, tools/art/cut_seal.py) centred
// on it. The seal is only scaled and placed, never redrawn.
// The iOS Dark and Tinted appearances come from the flat seal
// (design/assets/brand/seal-flat/seal-flat-small.svg). The Android themed layer is a vector
// drawable from tools/gen_seal_silhouette.py.
// Run from the repo root: node web/scripts/build-app-icon.mjs. Deterministic: run twice, same files.
import { mkdirSync, rmSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import sharp from 'sharp';

const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..', '..');
const GROUND_ART = join(ROOT, 'design/assets/art/app-icon-background.jpeg');
const SEAL_PNG = join(ROOT, 'design/assets/logo/lx-seal-v4.png');
const FLAT_SVG = join(ROOT, 'design/assets/brand/seal-flat/seal-flat-small.svg');
const OUT = join(ROOT, 'apps/shared/brand');
/** The iOS icon: one 1024 px image the system scales to every home-screen size. */
const SIZE = 1024;
/** The painted blue inside the ground art's brass frame, as a share of its side (the companion's
 *  macOS tile, companion/tools/icons/render.py MAC_GROUND_INNER): the frame never shows. */
const GROUND_INNER = [0.09, 0.91];
/** The seal's diameter as a share of the visible tile. iOS masks the square to its squircle, which
 *  keeps a centred circle of almost the whole side, so the coin's rim stays whole with a margin of
 *  ground round it. The companion's macOS tile uses 0.78 of its body; the phone tile matches it. */
const SEAL_SHARE = 0.78;
/** Android's adaptive canvas is 108 dp; the launcher shows the middle 72 dp under its mask and
 *  guarantees only the middle 66 dp (the safe zone). */
const ADAPTIVE = { canvas: 108, visible: 72, safe: 66 };
/** The largest density's adaptive layer: 108 dp at xxxhdpi (4x). Nothing asks for more. */
const ADAPTIVE_PX = ADAPTIVE.canvas * 4;
/** The legacy launcher icon, 48 dp, at each density bucket. */
const LEGACY_DENSITIES = { mdpi: 48, hdpi: 72, xhdpi: 96, xxhdpi: 144, xxxhdpi: 192 };
/** The legacy square icon's corner radius as a share of its side (4 dp of 48, Material's square
 *  product-icon keyline). */
const LEGACY_CORNER = 1 / 12;
/** The flat seal fills its own canvas edge to edge; on the dark and tinted icons it keeps a 7.5%
 *  margin inside the mask so it sits at the painted icon's optical size. */
const FLAT_SEAL_RATIO = 0.85;
/** Tinted icons are greyscale: the levels stretch the flat seal's outline (#170E09, luma 16) to
 *  black and its gold (#C9A465, luma 168) to white, so the tint lands on the ring and the LX; the
 *  curve sets the turquoise disc (luma 112) to a mid grey, apart from both the LX and the outline. */
const TINT_LEVELS = { black: 16, white: 168, curve: 1.6 };
/** Rec. 601 luma weights, the ones the levels above are measured in. */
const LUMA = [0.299, 0.587, 0.114];
const TRANSPARENT = { r: 0, g: 0, b: 0, alpha: 0 };
const PNG = { compressionLevel: 9, effort: 10 };

const androidSealShare = (SEAL_SHARE * ADAPTIVE.visible) / ADAPTIVE.canvas;
if (androidSealShare > ADAPTIVE.safe / ADAPTIVE.canvas) {
  throw new Error(`the Android seal (${androidSealShare}) leaves the 66 dp safe zone`);
}

/** The painted ground, its frame cropped away, at `px` square, opaque. */
async function renderGround(px) {
  const { width, height } = await sharp(GROUND_ART).metadata();
  const [from, to] = GROUND_INNER;
  return sharp(GROUND_ART)
    .extract({
      left: Math.round(width * from),
      top: Math.round(height * from),
      width: Math.round(width * (to - from)),
      height: Math.round(height * (to - from)),
    })
    .resize(px, px, { kernel: 'lanczos3' })
    .removeAlpha()
    .png()
    .toBuffer();
}

const renderSeal = (diameter) =>
  sharp(SEAL_PNG).resize(diameter, diameter, { kernel: 'lanczos3' }).png().toBuffer();

const renderFlatSeal = (diameter) =>
  sharp(FLAT_SVG, { density: 1200 }).resize(diameter, diameter).png().toBuffer();

const transparentCanvas = (px = SIZE) =>
  sharp({ create: { width: px, height: px, channels: 4, background: TRANSPARENT } });

/** The tile at `px`: the ground with the seal centred at `share` of the side. Opaque. */
async function renderTile(px, share = SEAL_SHARE) {
  const seal = await renderSeal(Math.round(px * share));
  return sharp(await renderGround(px))
    .composite([{ input: seal, gravity: 'center' }])
    .removeAlpha()
    .png()
    .toBuffer();
}

/** The App Store's icon: a full-bleed square, opaque (no alpha channel at all). */
async function iosAny(file) {
  await sharp(await renderTile(SIZE)).removeAlpha().png(PNG).toFile(file);
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
  const { black, white, curve } = TINT_LEVELS;
  const { data, info } = await sharp(await renderFlatSeal(Math.round(SIZE * FLAT_SEAL_RATIO)))
    .ensureAlpha()
    .raw()
    .toBuffer({ resolveWithObject: true });
  const grey = Buffer.alloc(data.length);
  for (let i = 0; i < data.length; i += 4) {
    const luma = LUMA[0] * data[i] + LUMA[1] * data[i + 1] + LUMA[2] * data[i + 2];
    const level = Math.round(Math.min(1, Math.max(0, (luma - black) / (white - black))) ** curve * 255);
    grey.set([level, level, level, data[i + 3]], i);
  }
  const seal = await sharp(grey, { raw: info }).png().toBuffer();
  await transparentCanvas()
    .composite([{ input: seal, gravity: 'center' }])
    .png()
    .toFile(file);
}

/** The iOS appearances: the painted tile for Any, the flat seal for Dark and Tinted. */
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

/** The adaptive icon wrapper; the round alias is the same layers (the launcher applies its mask). */
const ADAPTIVE_XML = `<?xml version="1.0" encoding="utf-8"?>
<adaptive-icon xmlns:android="http://schemas.android.com/apk/res/android">
    <background android:drawable="@mipmap/ic_launcher_background" />
    <foreground android:drawable="@mipmap/ic_launcher_foreground" />
    <monochrome android:drawable="@drawable/ic_launcher_monochrome" />
</adaptive-icon>
`;

/** A white mask at `px`: the circle, or the rounded square at `corner` of the side. */
const maskSvg = (px, shape) =>
  Buffer.from(
    shape === 'round'
      ? `<svg xmlns='http://www.w3.org/2000/svg' width='${px}' height='${px}'><circle cx='${px / 2}' cy='${px / 2}' r='${px / 2}' fill='#fff'/></svg>`
      : `<svg xmlns='http://www.w3.org/2000/svg' width='${px}' height='${px}'><rect width='${px}' height='${px}' rx='${px * LEGACY_CORNER}' fill='#fff'/></svg>`,
  );

/** A legacy launcher PNG (pre-adaptive launchers, and any that ask for the bitmap): the visible
 *  tile, masked to its shape. Drawn large and reduced, for a clean edge at 48 px. */
async function legacyIcon(px, shape) {
  const big = px * 4;
  const tile = await sharp(await renderTile(big))
    .ensureAlpha()
    .composite([{ input: maskSvg(big, shape), blend: 'dest-in' }])
    .png()
    .toBuffer();
  return sharp(tile).resize(px, px, { kernel: 'lanczos3' }).png(PNG).toBuffer();
}

async function android() {
  const dir = join(OUT, 'android');
  mkdirSync(dir, { recursive: true });
  // The ground fills the whole 108 dp canvas, so the parallax and every mask shape stay painted.
  await sharp(await renderGround(ADAPTIVE_PX)).png(PNG).toFile(join(dir, 'ic_launcher_background.png'));
  const seal = await renderSeal(Math.round(ADAPTIVE_PX * androidSealShare));
  await transparentCanvas(ADAPTIVE_PX)
    .composite([{ input: seal, gravity: 'center' }])
    .png(PNG)
    .toFile(join(dir, 'ic_launcher_foreground.png'));
  // The monochrome layer is the app's res/drawable/ic_launcher_monochrome.xml
  // (tools/gen_seal_silhouette.py), which launchers tint for themed icons on Android 13+.
  writeFileSync(join(dir, 'ic_launcher.xml'), ADAPTIVE_XML);
  writeFileSync(join(dir, 'ic_launcher_round.xml'), ADAPTIVE_XML);
  const legacy = join(dir, 'legacy');
  rmSync(legacy, { recursive: true, force: true });
  for (const [density, px] of Object.entries(LEGACY_DENSITIES)) {
    const bucket = join(legacy, `mipmap-${density}`);
    mkdirSync(bucket, { recursive: true });
    writeFileSync(join(bucket, 'ic_launcher.png'), await legacyIcon(px, 'square'));
    writeFileSync(join(bucket, 'ic_launcher_round.png'), await legacyIcon(px, 'round'));
  }
}

await ios();
await android();
