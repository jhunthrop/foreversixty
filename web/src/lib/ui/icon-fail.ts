// web/src/lib/ui/icon-fail.ts
// A spell icon whose file fails to load shows nothing (review 2026-11-10: an empty frame reads as an
// unticked checkbox), on every page and before any island hydrates. Its box stays, invisible, so
// nothing on the page moves when an image fails late. ui/SpellIcon and its .astro twin mark a frame
// that holds no column with ICON_HIDE_ON_FAIL; Base.astro inlines
// ICON_FAIL_SCRIPT at the top of <head>, so its listener is in place before the first image can
// fail. An image's error does not bubble, so the listener captures. A few dozen bytes: every page
// carries it.

/** The attribute on a frame that goes away with its failed image. */
export const ICON_HIDE_ON_FAIL = 'data-icon-hide-on-fail';

export const ICON_FAIL_SCRIPT =
  'addEventListener("error",function(e){var f=e.target.closest&&e.target.closest("[' +
  ICON_HIDE_ON_FAIL +
  ']");if(f&&e.target.tagName==="IMG")f.style.visibility="hidden"},true)';
