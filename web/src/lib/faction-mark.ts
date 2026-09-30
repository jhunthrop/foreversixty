// web/src/lib/faction-mark.ts
// The one place a faction emblem's image path is built (design/DESIGN-SYSTEM.md's Faction
// section): `FactionMark.astro` renders it for every Astro page, and a Svelte island (which
// cannot import a `.astro` file -- Vite's Astro integration only compiles `.astro` files
// reached from another `.astro` entry point, never from a framework component) inlines the
// identical `<img>` itself, built from this one path function, rather than a second,
// drifting copy of the URL string.
export type Faction = 'alliance' | 'horde';

export function factionMarkSrc(faction: Faction): string {
  return `/icons/hd/faction/${faction}.png`;
}
