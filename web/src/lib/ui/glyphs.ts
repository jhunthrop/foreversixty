// web/src/lib/ui/glyphs.ts
// The site's UI glyphs on one 24 px grid, 2 px stroke, round caps (polish spec 4.2: IconTile,
// SettingsRow, Button and ActionLink all draw from this one table). Each value is the inner SVG
// markup, so a glyph is drawn by ui/Glyph.svelte and nowhere else. The Discord mark is a solid
// shape, listed in FILLED_GLYPHS.

/** The three the dropdown's lazy chunk draws, exported alone so it carries only these. */
export const CHECK_GLYPH = '<path d="M5 12.5l4.5 4.5L19 7.5"/>';
export const CLOSE_GLYPH = '<path d="M6 6l12 12M18 6L6 18"/>';
export const SEARCH_GLYPH = '<circle cx="11" cy="11" r="7"/><path d="M20 20l-4-4"/>';

export const GLYPHS = {
  crown: '<path d="M3 8l4.5 4L12 5l4.5 7L21 8l-2 11H5L3 8z"/>',
  server:
    '<rect x="3" y="4" width="18" height="7" rx="2"/><rect x="3" y="13" width="18" height="7" rx="2"/><path d="M7 7.5h.01M7 16.5h.01"/>',
  chart: '<path d="M4 4v16h16"/><path d="M8 15l4-5 3 3 5-7"/>',
  bell: '<path d="M6 17v-6a6 6 0 0112 0v6l1.5 2h-15L6 17z"/><path d="M10 21h4"/>',
  star: '<path d="M12 3.5l2.6 5.4 5.9.8-4.3 4.1 1 5.9L12 16.9l-5.2 2.8 1-5.9-4.3-4.1 5.9-.8L12 3.5z"/>',
  card: '<rect x="3" y="5" width="18" height="14" rx="2"/><path d="M3 10h18M7 15h4"/>',
  hash: '<path d="M9 4L7 20M17 4l-2 16M4 9h16M3 15h16"/>',
  external: '<path d="M14 4h6v6M20 4l-9 9"/><path d="M18 14v4a2 2 0 01-2 2H6a2 2 0 01-2-2V8a2 2 0 012-2h4"/>',
  send: '<path d="M21 3L10 14"/><path d="M21 3l-7 18-4-7-7-4 18-7z"/>',
  refresh:
    '<path d="M20 11a8 8 0 00-14.5-4.5L4 8"/><path d="M4 4v4h4"/><path d="M4 13a8 8 0 0014.5 4.5L20 16"/><path d="M20 20v-4h-4"/>',
  lock: '<rect x="5" y="11" width="14" height="9" rx="2"/><path d="M8 11V8a4 4 0 018 0v3"/>',
  device: '<rect x="3" y="4" width="18" height="12" rx="2"/><path d="M8 20h8M12 16v4"/>',
  signout: '<path d="M9 4H6a2 2 0 00-2 2v12a2 2 0 002 2h3"/><path d="M16 8l4 4-4 4M20 12H9"/>',
  chevron: '<path d="M9 6l6 6-6 6"/>',
  guild:
    '<path d="M12 3l8 3v6c0 4.500-3.200 7.800-8 9-4.800-1.200-8-4.500-8-9V6l8-3z"/><path d="M12 8v8M8.500 11h7"/>',
  user: '<circle cx="12" cy="8" r="4"/><path d="M4 21a8 8 0 0116 0"/>',
  discord:
    '<path d="M18.6 5.6A16 16 0 0014.800 4.400l-.5 1a14 14 0 00-4.600 0l-.5-1A16 16 0 005.400 5.600C3 9.200 2.300 12.700 2.600 16.200A16 16 0 007.400 18.700l1-1.600a10 10 0 01-1.600-.8l.4-.3a11 11 0 009.600 0l.4.300a10 10 0 01-1.600.8l1 1.600a16 16 0 004.800-2.500c.4-4-.7-7.500-2.800-10.600zM9.300 14.300c-.9 0-1.600-.8-1.600-1.800s.7-1.800 1.600-1.800 1.600.8 1.600 1.800-.7 1.800-1.600 1.800zm5.400 0c-.9 0-1.600-.8-1.600-1.800s.7-1.800 1.600-1.800 1.600.8 1.600 1.800-.7 1.800-1.600 1.800z" stroke="none"/>',
  // Site alignment spec 2026-10-29 S4: the card-head glyphs.
  scale:
    '<path d="M12 3v18M7 21h10M5 7h14"/><path d="M5 7l-3 6a3 3 0 006 0L5 7zM19 7l-3 6a3 3 0 006 0l-3-6z"/>',
  bolt: '<path d="M13 2L4 14h7l-1 8 9-12h-7l1-8z"/>',
  list: '<path d="M9 6h11M9 12h11M9 18h11"/><path d="M4 6h.01M4 12h.01M4 18h.01"/>',
  pin: '<path d="M12 21s7-6.2 7-11.5A7 7 0 005 9.5C5 14.8 12 21 12 21z"/><circle cx="12" cy="9.5" r="2.5"/>',
  sword: '<path d="M14.5 17.5L3 6V3h3l11.5 11.5"/><path d="M13 19l6-6M16 16l4 4M19 21l2-2"/>',
  tree: '<circle cx="12" cy="5" r="2"/><circle cx="6" cy="19" r="2"/><circle cx="18" cy="19" r="2"/><path d="M12 7v5M12 12l-6 5M12 12l6 5"/>',
  book: '<path d="M4 5a2 2 0 012-2h13v16H6a2 2 0 00-2 2V5z"/><path d="M4 19a2 2 0 012-2h13"/>',
  upload: '<path d="M12 16V4M7 9l5-5 5 5"/><path d="M4 16v3a1 1 0 001 1h14a1 1 0 001-1v-3"/>',
  download: '<path d="M12 4v12M7 11l5 5 5-5"/><path d="M4 16v3a1 1 0 001 1h14a1 1 0 001-1v-3"/>',
  share:
    '<circle cx="6" cy="12" r="2.5"/><circle cx="18" cy="6" r="2.5"/><circle cx="18" cy="18" r="2.5"/><path d="M8.2 10.8l7.6-3.6M8.2 13.2l7.6 3.6"/>',
  layers: '<path d="M12 3l9 5-9 5-9-5 9-5z"/><path d="M3 13l9 5 9-5"/>',
  target: '<circle cx="12" cy="12" r="8"/><circle cx="12" cy="12" r="4"/><path d="M12 12h.01"/>',
  trophy:
    '<path d="M8 4h8v5a4 4 0 01-8 0V4z"/><path d="M8 6H4v1a4 4 0 004 4M16 6h4v1a4 4 0 01-4 4M12 13v4M8 21h8"/>',
  check: CHECK_GLYPH,
  close: CLOSE_GLYPH,
  search: SEARCH_GLYPH,
  news: '<rect x="3" y="4" width="18" height="16" rx="2"/><path d="M7 8h10M7 12h10M7 16h6"/>',
} as const;

export type GlyphName = keyof typeof GLYPHS;

/** Glyphs drawn as a solid shape rather than a 2 px stroke. */
export const FILLED_GLYPHS: ReadonlySet<GlyphName> = new Set<GlyphName>(['discord']);
