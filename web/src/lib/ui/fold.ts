// web/src/lib/ui/fold.ts
// ReadingSection's `fold`: a section that reads open from 768 folds under its heading on a phone.
// Its body is a `<details open data-fold>`, so a page with no script shows everything; this inline
// script closes them below FOLD_BELOW_PX as the page parses. The folded sections sit far below a
// phone's first screen, so closing them moves nothing the reader sees.

/** The width below which a folded section starts closed (the site's phone breakpoint). */
export const FOLD_BELOW_PX = 768;

/** The inline script a page with folded sections runs after them. */
export const foldScript = `if (matchMedia('(max-width: ${FOLD_BELOW_PX - 1}px)').matches) for (const d of document.querySelectorAll('details[data-fold]')) d.open = false;`;
