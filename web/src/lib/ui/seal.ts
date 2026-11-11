// web/src/lib/ui/seal.ts
// The LX seal's files (design/DESIGN-SYSTEM.md "Logo"), apart from the asset manifest so a light
// island (the crest's fallback in the select list) never pulls in the manifest's data loader.
const DIR = '/art';

/** The LX seal. The painted seal v4 (design/assets/logo/lx-seal-v4.png) at 64 px and up: the footer and
 *  the medallion. At 44 px and below the painted seal turns to mud, so the nav takes the flat seal
 *  (owner 2026-11-10, design/assets/brand/seal-flat): the small seal with its swash from 25 to 44 px, the
 *  tiny bold seal at 24 px and below. The flat seals are SVG, crisp at every density. */
export const SEAL = {
  flatTiny: `${DIR}/seal-flat-tiny-bold.svg`,
  flatSmall: `${DIR}/seal-flat-small.svg`,
  painted88: `${DIR}/seal-88.webp`,
  painted128: `${DIR}/seal-128.webp`,
} as const;

/** The tiny bold seal up to 24 CSS px. */
const FLAT_TINY_MAX_PX = 24;
/** The flat seal ends at 44 CSS px; the painted seal takes over above it. */
const FLAT_MAX_PX = 44;
/** The 88 file (176 px) covers up to 88 CSS px at 2x; above it the 128. */
const PAINTED_88_MAX_PX = 88;

/** The seal file for a rendered size in CSS px (each painted file covers 2x density). */
export function sealFor(size: number): string {
  if (size <= FLAT_TINY_MAX_PX) return SEAL.flatTiny;
  if (size <= FLAT_MAX_PX) return SEAL.flatSmall;
  return size <= PAINTED_88_MAX_PX ? SEAL.painted88 : SEAL.painted128;
}
