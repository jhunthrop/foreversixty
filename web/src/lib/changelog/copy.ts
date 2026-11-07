// web/src/lib/changelog/copy.ts
// Every string /changelog renders that is not an entry's own content (pages spec 2026-11-08
// section 6, gap PG1): the page's title and lead, the Show filter, the month heads, the
// returning visitor's tag, the empty line, the Coming up block and the site entries' action links.
import type { ChangelogShow } from './entries';

export const changelogCopy = {
  docTitle: 'Changelog',
  description: 'Everything that changed in World of Warcraft: Forever and on this site, dated and sourced.',
  title: 'What changed',
  showLabel: 'Show',
  show: { all: 'All', game: 'Game', site: 'Site' } satisfies Record<ChangelogShow, string>,
  /** The phone bar's line: `Game changes · 25`. */
  summaryWords: { all: 'All changes', game: 'Game changes', site: 'Site changes' } satisfies Record<
    ChangelogShow,
    string
  >,
  summary: (show: ChangelogShow, count: number): string => `${changelogCopy.summaryWords[show]} · ${count}`,
  barLabel: 'Changelog choices',
  changes: (n: number): string => (n === 1 ? '1 change' : `${n} changes`),
  newTag: 'New',
  empty: {
    all: 'No changes yet.',
    game: 'No game changes yet.',
    site: 'No site changes yet.',
  } satisfies Record<ChangelogShow, string>,
  comingUp: 'Coming up',
  /** A site entry ends in a link to the page that changed (ruling Q5), by the page's path.
   *  A site source whose path is not here fails the build: name its page first. */
  siteActions: {
    '/guild': 'Open the guild page',
    '/premium': 'Open the Premium page',
    '/sim': 'Open the simulator',
    '/sim/specs': 'Open spec support',
    '/account': 'Open your account',
    '/character': 'Open your character',
  } as Readonly<Record<string, string>>,
} as const;
