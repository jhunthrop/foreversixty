// web/src/lib/ui/assets.ts
// The one module that names the site's design assets (owner 2026-11-09: one texture set, one art
// set, one seal set, referenced through one manifest). The files live in web/public/art and are
// written by tools/art/build_art.py from the owner's sources in design/assets; a page or
// component imports a constant from here and never hard-codes a path.
// Generated art is the owner's explicit choice (2026-11-09): background only, never captioned as
// a screenshot (design/specs/2026-11-09-directions.md, round 5).

import { isRealIcon } from '../icons/placeholder';
import { dataUrl } from '../planner/load';
import type { BandArt } from './primitives';
// The phone files' bytes as data URLs, built in (assets.test.ts holds each to its file).
import bandPhoneInline from '/art/band-1200.webp?inline';
import heroPhoneInline from '/art/hero-1200.webp?inline';

const DIR = '/art';

/**
 * The one url of a game icon: the build's HD icon file (data/builds/<build>/icons/<stem>.webp),
 * or null when there is none to draw (no stem, an empty one, or the client's red question mark
 * placeholder), so every surface draws its socket or nothing instead (2026-11-04: the "?" leaked
 * wherever a caller skipped the check). No page or component builds an icon path by hand
 * (icon-guard.test.ts).
 */
export function iconUrl(build: string, stem: string | null | undefined): string | null {
  return isRealIcon(stem) ? dataUrl(build, `icons/${stem}.webp`) : null;
}

/** iconUrl for a stem a generated catalogue guarantees (a spec's tab icon, a role's icon): it
 *  throws on a missing or placeholder stem, so a broken catalogue fails the build, never a page. */
export function catalogueIconUrl(build: string, stem: string): string {
  const url = iconUrl(build, stem);
  if (url === null) throw new Error(`catalogue icon "${stem}" is not a real icon`);
  return url;
}

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
/**
 * Each band art's phone file as a data URL, the page's own band pixels inside the page itself
 * (2026-11-13, band LCP). The band is the largest paint on most pages, and as a CSS background
 * its file was found only after the stylesheet and then queued behind the document and the
 * faces: real visitors saw it seconds after the title. Inline it paints with the first frame.
 * The desktop file is not inlined but drawn over it (ui/band.css), so a wide screen gets the
 * same art at once and the sharper copy when it lands. A page inlines only its own art
 * (artStyle): the band is 3.5 KB, the hero 20 KB and only home draws it.
 */
export const ART_INLINE: Readonly<Record<BandArt, string>> = { band: bandPhoneInline, hero: heroPhoneInline };
/** Behind the raid plaque: the cave mouth and its lava glow. */
export const CAVE_ART = art('cave', 800, 1600);
/** The link preview's background, 1200 x 630 (lib/og.ts). */
export const OG_ART = `${DIR}/og-art.jpg`;

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
  /** The tiny bold flat seal, the tab's icon wherever SVG favicons work. */
  svg: '/favicon.svg',
  /** The flat seal rendered at 16, 32 and 48 px, for the rest. */
  ico: '/favicon.ico',
  appleTouch: '/apple-touch-icon.png',
} as const;

const url = (file: string): string => `url('${file}')`;

/**
 * The art as CSS custom properties, set once on <html> by Base.astro, so the stylesheets
 * (global.css, nav.css, the primitives) read the manifest instead of naming files. The page's
 * own band art (`art`) has its phone file inlined (ART_INLINE). A page without Base.astro (the
 * API's /b/{id}) carries the same properties in its copied chrome.
 */
export function artCssVars(art: BandArt): Readonly<Record<string, string>> {
  const phone = (key: BandArt, file: ResponsiveArt): string => url(key === art ? ART_INLINE[key] : file.phone);
  return {
    '--art-ground': url(GROUND_TILE),
    '--art-ground-size': `${GROUND_TILE_PX}px`,
    '--art-band': url(BAND_ART.desktop),
    '--art-band-phone': phone('band', BAND_ART),
    '--art-hero': url(HERO_ART.desktop),
    '--art-hero-phone': phone('hero', HERO_ART),
    '--art-cave': url(CAVE_ART.desktop),
    '--art-cave-phone': url(CAVE_ART.phone),
  };
}

/** artCssVars as an inline style attribute value. */
export function artStyle(art: BandArt): string {
  return Object.entries(artCssVars(art))
    .map(([name, value]) => `${name}:${value}`)
    .join(';');
}
