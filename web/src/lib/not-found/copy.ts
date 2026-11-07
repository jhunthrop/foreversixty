// web/src/lib/not-found/copy.ts
// Every string the 404 page renders that no shared copy module carries (pages spec 2026-11-08
// section 6, gap PG1). The answer's strings are the guides' (lib/guides/copy.ts,
// lib/account/answer-copy.ts); the lookup's are the home lookup's (lib/home-landing-copy.ts).
export const notFoundCopy = {
  docTitle: 'Page not found',
  description: 'No page exists at this address.',
  title: 'No page here',
  job: 'Nothing exists at this address. It may have moved as the game changes.',
  pick: 'Pick your class',
  didYouMean: 'Did you mean',
  /** The signed-in answer's third, smaller link (ruling Q6). */
  yourCharacter: 'Your character',
  /** A suggestion's page name. */
  guideTitle: (spec: string, className: string): string => `${spec} ${className} guide`,
  classGuideTitle: (className: string): string => `${className} guides`,
  bisTitle: (spec: string, className: string): string => `${spec} ${className} best in slot`,
} as const;
