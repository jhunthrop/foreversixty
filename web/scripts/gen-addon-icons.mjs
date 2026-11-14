// web/scripts/gen-addon-icons.mjs
// Writes the in-game addon's textures (addon/ForeverSixty/media) from the site's own sources, so the
// addon draws the site's art and never a hand-made copy (design/specs/2026-11-14-addon-brand.md
// sections 3 to 5):
// - icons/<name>.tga: ADDON_ICONS from the site's icon set (design/assets/icons/ui, read through
//   gen-ui-icons.mjs), a white glyph on transparent that the addon tints with SetVertexColor.
// - seal.tga: the flat seal (web/public/art/seal-flat-small.svg) in its own colours.
// - filigree.tga: the lead panel's corner (design/assets/ornaments/corner-filigree.svg), white.
// - ground.tga: the painted earth tile (web/public/art/ground.webp), whole, so it stays seamless.
// - band-<mood>.tga: the header art per tab (BAND_CROPS), cave.tga: the raid plate's backdrop.
// - crests/<class>.tga and crests/ring.tga: the site's ringed HD crests and one tintable ring.
// - MANIFEST.json: every file written with its size and sha256, and the sha256 of every source and of
//   this script. `--check` re-hashes the sources, this script and the files on disk against it
//   without rendering (the renderers' bytes differ by platform, the hashes of what is committed do
//   not) and lists anything stale, missing or unlisted. The addon's guard test
//   (addon/tests/media_guard_spec.lua) reads the same manifest.
// Every texture is a 32-bit uncompressed TGA with power-of-two sides: what the 1.60 client reads.
// Run: node scripts/gen-addon-icons.mjs [--check]
import { createHash } from 'node:crypto';
import { mkdirSync, readdirSync, readFileSync, rmSync, statSync, writeFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { Resvg } from '@resvg/resvg-js';
import sharp from 'sharp';
import { LINE_DIR, ICON_DIR, parseCorner, readIcons } from './gen-ui-icons.mjs';

const WEB = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const ROOT = path.resolve(WEB, '..');
export const MEDIA = path.join(ROOT, 'addon/ForeverSixty/media');
export const MANIFEST = path.join(MEDIA, 'MANIFEST.json');
const ART = path.join(WEB, 'public/art');
const CRESTS = path.join(WEB, 'public/icons/hd/crests');
const SEAL_SVG = path.join(ART, 'seal-flat-small.svg');
const CORNER_SVG = path.join(ROOT, 'design/assets/ornaments/corner-filigree.svg');
const SELF = fileURLToPath(import.meta.url);

/** The site glyphs the addon draws (spec section 3): the six tab icons, then the row glyphs. */
export const ADDON_ICONS = [
  'character', 'planner', 'bis', 'sim', 'guild', 'manage',
  'close', 'chevron-right', 'chevron-down', 'check', 'upgrade', 'skull', 'target', 'refresh', 'bolt',
  'pin', 'search', 'spinner',
];
/** The nine classes, the addon's lower-case class slugs. */
export const CLASSES = ['warrior', 'paladin', 'hunter', 'rogue', 'priest', 'shaman', 'mage', 'warlock', 'druid'];
/** Each header mood: its source in web/public/art, the crop as fractions (left, top, right, bottom)
 *  and whether it is mirrored. The approved mock's crops (design/mocks/addon_brand.py BAND_CROPS). */
export const BAND_CROPS = {
  me: { source: 'band-2400.webp', box: [0, 0.3, 1, 0.52], mirror: false },
  builds: { source: 'band-2400.webp', box: [0, 0.48, 1, 0.7], mirror: true },
  gear: { source: 'paperdoll-frame@2x.webp', box: [0.22, 0.42, 0.78, 0.5], mirror: false },
  rotation: { source: 'band-2400.webp', box: [0, 0.62, 1, 0.84], mirror: false },
  cave: { source: 'cave-1600.webp', box: [0.3, 0.12, 1, 0.4], mirror: false },
};
/** The raid plate's backdrop: the cave mouth and its glow, in source pixels (x, y, width, height). */
const CAVE_PLATE = { source: 'cave-1600.webp', region: [480, 60, 1120, 560] };
/** Texture sides, in px. */
const SIZE = { icon: 32, seal: 64, filigree: 64, ground: 256, bandW: 512, bandH: 64, caveW: 256, caveH: 128, crest: 128 };
/** The ring: its outer radius and thickness at the 128 px crest (the 64 px ring's 31 and 4, doubled). */
const RING = { outer: 62, width: 8 };

const isPow2 = (n) => n > 0 && (n & (n - 1)) === 0;

/** RGBA pixels (top row first) to a 32-bit uncompressed, top-left-origin TGA. */
export function tga(width, height, rgba) {
  if (!isPow2(width) || !isPow2(height)) throw new Error(`${width}x${height}: the client reads power-of-two sides only`);
  if (rgba.length !== width * height * 4) throw new Error(`${width}x${height}: ${rgba.length} bytes is not RGBA`);
  const out = Buffer.alloc(18 + rgba.length);
  out[2] = 2; // uncompressed true colour
  out.writeUInt16LE(width, 12);
  out.writeUInt16LE(height, 14);
  out[16] = 32;
  out[17] = 0x20; // top-left origin, as the shipped Pillow files
  for (let i = 0; i < rgba.length; i += 4) {
    out[18 + i] = rgba[i + 2];
    out[18 + i + 1] = rgba[i + 1];
    out[18 + i + 2] = rgba[i];
    out[18 + i + 3] = rgba[i + 3];
  }
  return out;
}

function svgPixels(svg, side) {
  const image = new Resvg(svg, { fitTo: { mode: 'width', value: side }, background: 'rgba(0,0,0,0)' }).render();
  if (image.width !== side || image.height !== side) throw new Error(`svg rendered ${image.width}x${image.height}`);
  return Buffer.from(image.pixels);
}

/** One site glyph as white on transparent, drawn as ui/Icon.svelte draws it. */
function glyphSvg(icon) {
  const paint =
    icon.kind === 'fill'
      ? 'fill="#fff"'
      : 'fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"';
  return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${icon.box} ${icon.box}" ${paint}>${icon.body}</svg>`;
}

async function rasterTga(input, width, height) {
  const { data } = await input.ensureAlpha().raw().toBuffer({ resolveWithObject: true });
  return tga(width, height, data);
}

async function bandTga({ source, box, mirror }) {
  const image = sharp(path.join(ART, source));
  const { width, height } = await image.metadata();
  const [l, t, r, b] = box;
  let cut = image
    .extract({ left: Math.round(l * width), top: Math.round(t * height),
      width: Math.round((r - l) * width), height: Math.round((b - t) * height) })
    .resize(SIZE.bandW, SIZE.bandH, { fit: 'fill', kernel: 'lanczos3' });
  if (mirror) cut = cut.flop();
  return rasterTga(cut, SIZE.bandW, SIZE.bandH);
}

/** Every source this generator reads, for the manifest. */
function sources() {
  const { custom, line } = readIcons();
  const families = [...custom.map((icon) => [icon, ICON_DIR]), ...line.map((icon) => [icon, LINE_DIR])];
  const icons = ADDON_ICONS.map((name) => {
    const found = families.find(([icon]) => icon.name === name);
    if (found === undefined) throw new Error(`ADDON_ICONS names ${name}, which design/assets/icons/ui lacks`);
    return { icon: found[0], file: path.join(found[1], `${name}.svg`) };
  });
  const files = [
    ...icons.map((entry) => entry.file),
    SEAL_SVG, CORNER_SVG, path.join(ART, 'ground.webp'), path.join(ART, CAVE_PLATE.source),
    ...new Set(Object.values(BAND_CROPS).map((crop) => path.join(ART, crop.source))),
    ...CLASSES.map((slug) => path.join(CRESTS, `${slug}.webp`)),
  ];
  return { icons, files: [...new Set(files)].sort() };
}

const sha = (bytes) => createHash('sha256').update(bytes).digest('hex');
const rel = (file) => path.relative(ROOT, file).split(path.sep).join('/');

/** Every texture by its path under media/, as TGA bytes. */
export async function renderAddonMedia() {
  const { icons } = sources();
  const out = new Map();
  for (const { icon } of icons) {
    out.set(`icons/${icon.name}.tga`, tga(SIZE.icon, SIZE.icon, svgPixels(glyphSvg(icon), SIZE.icon)));
  }
  out.set('seal.tga', tga(SIZE.seal, SIZE.seal, svgPixels(readFileSync(SEAL_SVG, 'utf8'), SIZE.seal)));
  const corner = parseCorner();
  const cornerSvg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${corner.box} ${corner.box}" fill="#fff">${corner.body}</svg>`;
  out.set('filigree.tga', tga(SIZE.filigree, SIZE.filigree, svgPixels(cornerSvg, SIZE.filigree)));
  // The whole tile, so the texture repeats without a seam; the addon draws it at Theme.SIZES.groundRepeat.
  out.set('ground.tga', await rasterTga(
    sharp(path.join(ART, 'ground.webp')).resize(SIZE.ground, SIZE.ground, { kernel: 'lanczos3' }), SIZE.ground, SIZE.ground));
  for (const [mood, crop] of Object.entries(BAND_CROPS)) out.set(`band-${mood}.tga`, await bandTga(crop));
  const [x, y, w, h] = CAVE_PLATE.region;
  out.set('cave.tga', await rasterTga(
    sharp(path.join(ART, CAVE_PLATE.source)).extract({ left: x, top: y, width: w, height: h })
      .resize(SIZE.caveW, SIZE.caveH, { fit: 'fill', kernel: 'lanczos3' }), SIZE.caveW, SIZE.caveH));
  for (const slug of CLASSES) {
    const source = sharp(path.join(CRESTS, `${slug}.webp`));
    const { data } = await source.clone().ensureAlpha().raw().toBuffer({ resolveWithObject: true });
    if (data[3] > 8) throw new Error(`${slug}.webp: the corner is not transparent; not the round crest`);
    out.set(`crests/${slug}.tga`, await rasterTga(source.resize(SIZE.crest, SIZE.crest, { kernel: 'lanczos3' }), SIZE.crest, SIZE.crest));
  }
  const c = SIZE.crest / 2;
  const ring = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${SIZE.crest} ${SIZE.crest}"><circle cx="${c}" cy="${c}" r="${RING.outer - RING.width / 2}" fill="none" stroke="#fff" stroke-width="${RING.width}"/></svg>`;
  out.set('crests/ring.tga', tga(SIZE.crest, SIZE.crest, svgPixels(ring, SIZE.crest)));
  return out;
}

/** The manifest for `files` (path under media/ -> bytes). */
export function manifestFor(files) {
  const entries = {};
  for (const [file, bytes] of [...files].sort(([a], [b]) => a.localeCompare(b))) {
    entries[file] = { width: bytes.readUInt16LE(12), height: bytes.readUInt16LE(14), bytes: bytes.length, sha256: sha(bytes) };
  }
  const { files: inputs } = sources();
  const sourceHashes = Object.fromEntries(inputs.map((file) => [rel(file), sha(readFileSync(file))]));
  return {
    generator: { path: rel(SELF), sha256: sha(readFileSync(SELF)) },
    sources: sourceHashes,
    files: entries,
    totalBytes: Object.values(entries).reduce((sum, entry) => sum + entry.bytes, 0),
  };
}

/** Files under media/ the generator does not own: the fonts and the minimap seal (addon/tools). */
export const NOT_GENERATED = [/^fonts\//, /^minimap\.tga$/, /^MANIFEST\.json$/];

function listMedia(dir = MEDIA, prefix = '') {
  return readdirSync(dir).flatMap((name) => {
    const full = path.join(dir, name);
    return statSync(full).isDirectory() ? listMedia(full, `${prefix}${name}/`) : [`${prefix}${name}`];
  });
}

/** What is wrong with the committed media against the manifest and the sources now: [] when current. */
export function checkAddonMedia() {
  const manifest = JSON.parse(readFileSync(MANIFEST, 'utf8'));
  const problems = [];
  if (manifest.generator.sha256 !== sha(readFileSync(SELF))) problems.push('the generator changed since the media were written');
  const { files: inputs } = sources();
  const now = Object.fromEntries(inputs.map((file) => [rel(file), sha(readFileSync(file))]));
  for (const [file, hash] of Object.entries(now)) {
    if (manifest.sources[file] !== hash) problems.push(`source ${file} changed or is new`);
  }
  for (const file of Object.keys(manifest.sources)) if (!(file in now)) problems.push(`source ${file} is no longer read`);
  for (const [file, entry] of Object.entries(manifest.files)) {
    let bytes;
    try {
      bytes = readFileSync(path.join(MEDIA, file));
    } catch {
      problems.push(`media/${file} is missing`);
      continue;
    }
    if (sha(bytes) !== entry.sha256) problems.push(`media/${file} differs from the manifest`);
    if (!isPow2(entry.width) || !isPow2(entry.height) || bytes[2] !== 2 || bytes[16] !== 32)
      problems.push(`media/${file} is not a power-of-two 32-bit uncompressed TGA`);
  }
  for (const file of listMedia()) {
    if (!(file in manifest.files) && !NOT_GENERATED.some((pattern) => pattern.test(file)))
      problems.push(`media/${file} is not written by this generator`);
  }
  return problems;
}

if (process.argv[1] === SELF) {
  if (process.argv.includes('--check')) {
    const problems = checkAddonMedia();
    for (const problem of problems) console.error(problem);
    if (problems.length > 0) {
      console.error('run: node scripts/gen-addon-icons.mjs');
      process.exit(1);
    }
    console.log('addon media are current');
  } else {
    const files = await renderAddonMedia();
    rmSync(path.join(MEDIA, 'icons'), { recursive: true, force: true });
    for (const [file, bytes] of files) {
      mkdirSync(path.dirname(path.join(MEDIA, file)), { recursive: true });
      writeFileSync(path.join(MEDIA, file), bytes);
    }
    const manifest = manifestFor(files);
    writeFileSync(MANIFEST, `${JSON.stringify(manifest, null, 2)}\n`);
    console.log(`wrote ${files.size} textures, ${Math.round(manifest.totalBytes / 1024)} KB, and media/MANIFEST.json`);
  }
}
