// web/src/lib/legal-notice.ts
// The site's full legal notice: the trademark sentence and the operator line. Round 5 moved them out
// of the one-row footer (which keeps FOOTER_FAN_LINE) to /about and /terms, whose content pages carry
// them verbatim (src/content/pages/legal-notice.test.ts holds them to these strings).

export const LEGAL_TRADEMARK =
  'Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.';

export const LEGAL_OPERATOR = 'Forever Sixty is a registered trade name of Commish LLC. © Commish LLC.';

/** The pages that carry the notice, by content file. */
export const LEGAL_NOTICE_PAGES = ['about.md', 'terms.md'] as const;
