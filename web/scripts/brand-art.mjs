// web/scripts/brand-art.mjs
// The brand's raster pieces, drawn one way for every generator that needs them: the app icon set
// (build-app-icon.mjs), the Discord images (design/community/discord/build-assets.mjs) and the
// store listing images (design/listing/build-store.mjs).
// - The painted seal v4 (design/assets/logo/lx-seal-v4.png, cut by tools/art/cut_seal.py) at 64 px
//   and up. It is only scaled and placed, never redrawn.
// - The flat seal (design/assets/brand/seal-flat, tools/art/flat_seal.py) at 44 px and under.
// - The tile: the owner's painted blue ground (design/assets/art/app-icon-background.jpeg) with its
//   brass frame cropped away and the painted seal centred on it. It is the square identity image:
//   the app icon, the store icons, the Discord server icon and the bot avatar.
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import sharp from 'sharp';

export const ROOT = join(dirname(fileURLToPath(import.meta.url)), '..', '..');
const GROUND_ART = join(ROOT, 'design/assets/art/app-icon-background.jpeg');
const SEAL_PNG = join(ROOT, 'design/assets/logo/lx-seal-v4.png');
/** The flat seal's drawings: `small` for 24 to 44 px, `tiny` (bold, fewer details) for 20 px and under. */
const FLAT_SVG = {
  small: join(ROOT, 'design/assets/brand/seal-flat/seal-flat-small.svg'),
  tiny: join(ROOT, 'design/assets/brand/seal-flat/seal-flat-tiny-bold.svg'),
};
/** The tile's master size: the App Store's 1024 px icon. */
export const TILE_SIZE = 1024;
/** The painted blue inside the ground art's brass frame, as a share of its side (the companion's
 *  macOS tile, companion/tools/icons/render.py MAC_GROUND_INNER): the frame never shows. */
const GROUND_INNER = [0.09, 0.91];
/** The seal's diameter as a share of the visible tile. iOS masks the square to its squircle, which
 *  keeps a centred circle of almost the whole side, so the coin's rim stays whole with a margin of
 *  ground round it. The companion's macOS tile uses 0.78 of its body; the phone tile matches it.
 *  A circle mask (Discord, the Play launcher) keeps the same margin. */
export const SEAL_SHARE = 0.78;
/** The largest size the flat seal is ever shown at (owner 2026-11-10); the painted seal from 64 up. */
export const FLAT_SEAL_MAX_PX = 44;
export const PNG = { compressionLevel: 9, effort: 10 };
const TRANSPARENT = { r: 0, g: 0, b: 0, alpha: 0 };

/** The painted ground, its frame cropped away, at `px` square, opaque. */
export async function renderGround(px) {
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

/** The painted seal at `diameter`, on transparency. */
export const renderSeal = (diameter) =>
  sharp(SEAL_PNG).resize(diameter, diameter, { kernel: 'lanczos3' }).png().toBuffer();

/** The flat seal at `diameter`, on transparency, rendered from its SVG (never a downscale). */
export const renderFlatSeal = (diameter, drawing = 'small') =>
  sharp(FLAT_SVG[drawing], { density: 1200 }).resize(diameter, diameter).png().toBuffer();

/** The seal a placement `px` wide shows: flat at FLAT_SEAL_MAX_PX and under, painted above. */
export const renderSealFor = (px) => (px <= FLAT_SEAL_MAX_PX ? renderFlatSeal(px) : renderSeal(px));

export const transparentCanvas = (px = TILE_SIZE) =>
  sharp({ create: { width: px, height: px, channels: 4, background: TRANSPARENT } });

/** The tile at `px`: the ground with the seal centred at `share` of the side. Opaque. */
export async function renderTile(px, share = SEAL_SHARE) {
  const seal = await renderSeal(Math.round(px * share));
  return sharp(await renderGround(px))
    .composite([{ input: seal, gravity: 'center' }])
    .removeAlpha()
    .png()
    .toBuffer();
}
