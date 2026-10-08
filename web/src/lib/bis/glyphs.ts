// web/src/lib/bis/glyphs.ts
// The BiS page's stroke glyphs (slot and source kind), one path per glyph on a 0 0 24 24
// box. They live here, not inline in SlotIcon.astro / SourceKindIcon.astro, because the
// page draws hundreds of them (353 rows on the hunter page): GlyphSheet.astro emits each
// path ONCE as a <symbol> and every icon is a <use> reference, which cut the page HTML that
// Lighthouse's largest-contentful-paint budget was measuring (2026-09-29).

/** Slot glyphs; finger1/finger2 and trinket1/trinket2 share `finger`/`trinket`. */
export const SLOT_GLYPH_PATHS: Record<string, string> = {
  head: 'M6 11a6 6 0 0 1 12 0v3H6v-3ZM5 14h14v2a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2Z',
  neck: 'M7 4c0 3 2 5 5 5s5-2 5-5 M12 9v3 M12 12a4 4 0 1 0 0.01 0Z',
  shoulder: 'M4 10a8 8 0 0 1 16 0v2h-5l-3-2-3 2H4v-2Z',
  back: 'M12 4 5 8v12l7-3 7 3V8l-7-4Z',
  chest: 'M8 4 4 7v13h16V7l-4-3-4 2-4-2Z',
  wrist: 'M5 10h14v4H5z M8 10v4M16 10v4',
  hands:
    'M7 21v-7a2 2 0 1 1 4 0v-3a1.5 1.5 0 1 1 3 0v2a1.5 1.5 0 1 1 3 0v2a1.5 1.5 0 1 1 3 0v3a4 4 0 0 1-4 4H7Z',
  waist: 'M4 11h16v2H4z M10 10h4v4h-4z',
  legs: 'M8 4h8l1 16h-4l-1-9-1 9H7L8 4Z',
  feet: 'M7 4h4v9l4 3v4H7a2 2 0 0 1-2-2v-9a5 5 0 0 1 2-5Z',
  finger: 'M12 21a6 6 0 1 0 0-12 6 6 0 0 0 0 12ZM10 9l1-6h2l1 6',
  trinket: 'M12 3 3 9l9 12 9-12-9-6Z',
  main_hand: 'M6 20 18 8M15 5l4 4-2 2-4-4 2-2ZM4 22l3-1 1-3-3 1-1 3Z',
  off_hand: 'M12 3 4 6v6c0 5 3.5 7.5 8 9 4.5-1.5 8-4 8-9V6l-8-3Z',
  ranged: 'M5 4a16 16 0 0 1 0 16M9 12h9m-3-3 3 3-3 3',
};

/** Source-kind glyphs; raid shares `dungeon`, world shares `zone`. */
export const SOURCE_KIND_GLYPH_PATHS: Record<string, string> = {
  quest: 'M6 3h9l3 3v15H6V3ZM15 3v3h3 M9 10h6M9 13h6M9 16h4',
  dungeon:
    'M12 3a6 6 0 0 1 6 6c0 3-2 4-2 6H8c0-2-2-3-2-6a6 6 0 0 1 6-6ZM9 17h6l-1 4H10l-1-4ZM9.5 9a1 1 0 1 0 0-2 1 1 0 0 0 0 2ZM14.5 9a1 1 0 1 0 0-2 1 1 0 0 0 0 2Z',
  crafted: 'M14 3 21 10l-2.5 2.5L11 5 14 3ZM11 5 4 12l1.5 4.5L10 18l7-7 M4.5 19.5l3-1-2-2-1 3Z',
  vendor: 'M12 2a4 4 0 0 1 4 4v2H8V6a4 4 0 0 1 4-4ZM6 8h12l1 12H5L6 8ZM10 12v2M14 12v2',
  rep: 'M6 3h12v10l-6 8-6-8V3ZM9 8h6M9 11h6',
  zone: 'M12 21s7-7.5 7-12a7 7 0 1 0-14 0c0 4.5 7 12 7 12ZM12 12a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5Z',
  pvp: 'M4 5l8 8M20 5l-8 8M4 19l6-6M20 19l-6-6M4 5l3-1M20 5l-3-1M4 19l3 1M20 19l-3 1',
};

/** The set-bonus note's glyph: two linked pieces. */
export const SET_BONUS_GLYPH_PATH = 'M9 15a4 4 0 1 1 0-8h2M15 9a4 4 0 1 1 0 8h-2M9 12h6';

/** The verified-pick check mark. */
export const VERIFIED_GLYPH_PATH = 'M5 13l4 4L19 7';

export function slotGlyphKey(slot: string): string {
  if (slot === 'finger1' || slot === 'finger2') return 'finger';
  if (slot === 'trinket1' || slot === 'trinket2') return 'trinket';
  return slot in SLOT_GLYPH_PATHS ? slot : 'trinket';
}

export function sourceKindGlyphKey(kind: string): string {
  if (kind === 'raid') return 'dungeon';
  // world_drop reuses the same glyph `world` itself reuses (`zone`) --
  // both are "go kill something out in the world", not a distinct icon
  // worth its own path (world-drop-pool lane, 2026-09-29).
  if (kind === 'world' || kind === 'world_drop') return 'zone';
  return kind in SOURCE_KIND_GLYPH_PATHS ? kind : 'zone';
}
