// web/src/lib/content-pages/copy.ts
// The strings the content pages (legal and About) add around their markdown (pages spec
// 2026-11-08 section 6.2): the phone contents count and the legal tabs' accessible name. The
// contents label itself is the guides' (`guidesCopy.tocBarLabel`, On this page).
export const contentPagesCopy = {
  sections: (n: number): string => (n === 1 ? '1 section' : `${n} sections`),
  legalTabsLabel: 'Legal pages',
} as const;
