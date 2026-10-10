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

/** The row-descriptor size of the emblem (design/DESIGN-SYSTEM.md: 16px in a row descriptor). */
export const FACTION_MARK_SIZE = 16;

export function factionName(faction: Faction): string {
  return faction === 'alliance' ? 'Alliance' : 'Horde';
}

export function factionMarkSrc(faction: Faction): string {
  return `/icons/hd/faction/${faction}.webp`;
}

// Guild header art round (design/specs/2026-10-04-guild-page.md §12.2, owner's final build
// decision): the flat iconic faction logo (Alliance lion-in-shield, Horde tusked "H") --
// FactionMark.astro's own second provenance note -- is the ONE emblem the guild header
// shows, both in the identity-row ring (FactionCrest.svelte) and the band's watermark.
// `factionMarkSrc`'s 72px unit-frame shield/disc above is untouched everywhere else on the
// site (nav chip, row descriptors, ClassHeader's faction toggle) -- this is a second,
// visually distinct asset, not a larger export of the first.
export function factionLogoSrc(faction: Faction): string {
  return `/icons/hd/faction/${faction}-logo-512.webp`;
}

// design/DESIGN-SYSTEM.md "Faction," the bar variant (full-saturation swatch, the same
// choice ClassCrest.astro's own ring makes from classes.json's full-saturation `color`
// field) -- the one place this swatch is defined for a TypeScript caller; gen_guild_header.py's
// FACTION_BAR dict is this same pair, kept in sync by hand since the mock is Python.
export const FACTION_BAR_COLOR: Record<Faction, string> = {
  alliance: '#2f6fd6',
  horde: '#c0392b',
};
