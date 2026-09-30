// web/src/lib/faction-mark.ts
// The one place a faction emblem's image path is built (design/DESIGN-SYSTEM.md's Faction
// section): `FactionMark.astro` renders it for every Astro page, and a Svelte island (which
// cannot import a `.astro` file -- Vite's Astro integration only compiles `.astro` files
// reached from another `.astro` entry point, never from a framework component) inlines the
// identical `<img>` itself, built from this one path function, rather than a second,
// drifting copy of the URL string.
//
// Review round 2: WebP, lossy quality 90 with alpha, resampled to 72px from the original
// scratchpad 256px sources -- 3-3.6KB apiece, down from 7-10KB as PNG at the same 72px.
// The PNGs are gone from public/ entirely -- this is the only path any caller reads.
export type Faction = 'alliance' | 'horde';

export function factionMarkSrc(faction: Faction): string {
  return `/icons/hd/faction/${faction}.webp`;
}
