// web/src/lib/ui/assets.ts
// The one module that names the site's design assets (owner 2026-11-09: one texture set, one art
// set, one seal set, referenced through one manifest). The files live in web/public/art and are
// written by tools/art/build_art.py from the owner's sources in design/assets; a page or
// component imports a constant from here and never hard-codes a path.
// Generated art is the owner's explicit choice (2026-11-09): background only, never captioned as
// a screenshot (design/specs/2026-11-09-directions.md, round 5).

const DIR = '/art';

/** One art file at the widths the build writes: the phone width, then the desktop width. */
export interface ResponsiveArt {
  readonly phone: string;
  readonly desktop: string;
  /** The width below which the phone file is enough (CSS px). */
  readonly breakpoint: number;
}

const art = (name: string, phone: number, desktop: number): ResponsiveArt => ({
  phone: `${DIR}/${name}-${phone}.webp`,
  desktop: `${DIR}/${name}-${desktop}.webp`,
  breakpoint: 768,
});

/** The painted earth ground, periodic at GROUND_TILE_PX (art.ground_tile): tiled at 1:1. */
export const GROUND_TILE = `${DIR}/ground.webp`;
export const GROUND_TILE_PX = 1600;

/** The painted section band, navy into teal cloud: stretched with cover, never tiled. */
export const BAND_ART = art('band', 1200, 2400);
/** The home hero, the camp road at dusk. */
export const HERO_ART = art('hero', 1200, 2400);
/** Behind the raid plaque: the cave mouth and its lava glow. */
export const CAVE_ART = art('cave', 800, 1600);
/** The link preview's background, 1200 x 630 (lib/og.ts). */
export const OG_ART = `${DIR}/og-art.jpg`;

/** The LX seal, v3 (design/assets/logo/lx-seal-v3.png): one logo everywhere, nav, footer and
 *  favicons alike (owner 2026-11-09). Its bold rim holds down to 16 px, so there is no flat seal. */
export const SEAL = {
  painted88: `${DIR}/seal-88.webp`,
  painted128: `${DIR}/seal-128.webp`,
} as const;

/** The 88 file (176 px) covers up to 88 CSS px at 2x; above it the 128. */
const PAINTED_88_MAX_PX = 88;

/** The seal file for a rendered size in CSS px (2x density covered by each file). */
export function sealFor(size: number): string {
  return size <= PAINTED_88_MAX_PX ? SEAL.painted88 : SEAL.painted128;
}

/** The paperdoll frame (paperdoll_asset.py: real alpha, the owner's band in the centre). */
export const PAPERDOLL_FRAME = {
  x1: `${DIR}/paperdoll-frame@1x.webp`,
  x2: `${DIR}/paperdoll-frame@2x.webp`,
  /** The frame's intrinsic size at 2x, for the aspect ratio. */
  width: 1774,
  height: 2222,
} as const;

/** The favicons, at the site root where browsers look for them. */
export const FAVICON = {
  ico: '/favicon.ico',
  png32: '/favicon-32.png',
  appleTouch: '/apple-touch-icon.png',
} as const;

const url = (file: string): string => `url('${file}')`;

/**
 * The art as CSS custom properties, set once on <html> by Base.astro, so the stylesheets
 * (global.css, nav.css, the primitives) read the manifest instead of naming files. A page
 * without Base.astro (the API's /b/{id}) gets the flat token colours underneath.
 */
export const ART_CSS_VARS: Readonly<Record<string, string>> = {
  '--art-ground': url(GROUND_TILE),
  '--art-ground-size': `${GROUND_TILE_PX}px`,
  '--art-band': url(BAND_ART.desktop),
  '--art-band-phone': url(BAND_ART.phone),
  '--art-hero': url(HERO_ART.desktop),
  '--art-hero-phone': url(HERO_ART.phone),
  '--art-cave': url(CAVE_ART.desktop),
  '--art-cave-phone': url(CAVE_ART.phone),
};

/** ART_CSS_VARS as an inline style attribute value. */
export function artStyle(): string {
  return Object.entries(ART_CSS_VARS)
    .map(([name, value]) => `${name}:${value}`)
    .join(';');
}
