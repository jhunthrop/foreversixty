// web/src/lib/character-selector/config.ts
// Every number the nav character selector (design/specs/2026-10-09-nav-character-selector.md)
// is built from, in one place so the stale rule, the list sizes and the storage key are
// never restated in a component.

const SECONDS_PER_DAY = 86_400;

/** Decision 11: stale is more than this many times the player's own median gap between syncs. */
export const STALE_GAP_MULTIPLIER = 3;
/** ...clamped to at least this many days (a nightly raider) ... */
export const STALE_MIN_DAYS = 3;
/** ...and at most this many (a weekend player). */
export const STALE_MAX_DAYS = 14;
/** With fewer than three syncs on record the API sends no median; everyone gets this. */
export const STALE_FALLBACK_DAYS = 7;

export const STALE_MIN_SEC = STALE_MIN_DAYS * SECONDS_PER_DAY;
export const STALE_MAX_SEC = STALE_MAX_DAYS * SECONDS_PER_DAY;
export const STALE_FALLBACK_SEC = STALE_FALLBACK_DAYS * SECONDS_PER_DAY;

/** Decision 10: `fs.characterOrder`, a convenience list of the keys chosen on this browser. */
export const CHARACTER_ORDER_STORAGE_KEY = 'fs.characterOrder';
export const MAX_CHOSEN_KEYS = 20;

/** Section 3.C: six rows fit without scrolling; a seventh scrolls under a fade. */
export const VISIBLE_ROWS = 6;
/** Section 3.C: more than this many characters adds the "Filter by name" input. */
export const FILTER_AFTER_CHARACTERS = 8;

/** Dispatched on `window` by the phone Menu's first row; the selector island opens its sheet. */
export const OPEN_SELECTOR_EVENT = 'fs:open-character-selector';

/** The pages whose state lives in the URL and cannot re-render in place on a choice. */
export const URL_STATE_PAGES = ['/planner', '/sim'] as const;
