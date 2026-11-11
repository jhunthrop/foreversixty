// web/src/lib/ui/regions.ts
// The mock parity gate's region names (design/parity.md). The same `data-region` values are
// emitted by the mock library (design/mocks/ws_lib.py, ws_prims.py) and by the primitives here,
// so tests/e2e/support/parity.ts can pair every region of a mock with the built page's and
// compare what each holds and where it sits. One name, one place: a primitive never spells a
// region as a string literal.

export const REGION = {
  /** The site header (Header.astro): wordmark, search, account chip, the doors. */
  siteHeader: 'site-header',
  /** The site footer (Footer.astro). */
  siteFooter: 'site-footer',
  /** PageHeader: the whole page head. */
  head: 'head',
  /** PageHeader: crest, title and meta line. */
  headTitle: 'head-title',
  /** PageHeader: the one meta line under the title. */
  headMeta: 'head-meta',
  /** PageHeader: the answer at the right (BiS: the set's figure, its trust line and links). */
  headAnswer: 'head-answer',
  /** ControlBar. */
  controls: 'controls',
  /** The page's summary line under the bar (BiS: the comparison sentence and the view toggle). */
  summary: 'summary',
  /** Countdown panels under the head (home: the first raids and the guild's next event). */
  countdowns: 'countdowns',
  /** One-line standing notes between the summary and the Plate (BiS: source rule, hit cap, weapon). */
  notes: 'notes',
  /** Plate: the page's one primary object. */
  plate: 'plate',
  /** Paperdoll: the slot cells. */
  slots: 'slots',
  /** PaperdollFrame: the medallion socket (the Battle.net head render, else the seal). */
  paperdollMedallion: 'paperdoll-medallion',
  /** PaperdollFrame: the centre panel (the selected slot's tooltip). */
  paperdollCentre: 'paperdoll-centre',
  /** Inspector at the side (from 1760). */
  inspector: 'inspector',
  /** Inspector expanded in place under its cell (below 1760). */
  inspectorInline: 'inspector-inline',
  /** The page's closing one-liner (BiS: the guide line). */
  footerNote: 'footer-note',
  /** A document's prose before its first heading (About's text; pages
   *  spec 2026-11-08 section 3A). */
  intro: 'intro',
  /** A modal confirm over the page (the account deletion dialog; mock ws_prims `Dialog`). */
  dialog: 'dialog',
} as const;

export type RegionName = (typeof REGION)[keyof typeof REGION];

/** A Section's region: `section-<slug of its title>` ("What to stack" -> `section-what-to-stack`).
 *  The mock library's `region_slug` (design/mocks/ws_lib.py) is the same function. */
export function sectionRegion(title: string): string {
  const slug = title
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '');
  return `section-${slug}`;
}
