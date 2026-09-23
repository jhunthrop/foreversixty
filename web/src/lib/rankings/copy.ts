// web/src/lib/rankings/copy.ts
// Every visible string the /rankings encounter picker uses, in the site's own honest-copy
// voice: never demands a choice the visitor cannot make.
export const encounterPickerCopy = {
  heading: 'Choose an encounter',
  failed: 'Encounters did not load.',
  noneYet: 'No encounters have been ranked yet.',
  tryGuildsProgress: 'The Guilds board ranks guild progress across every encounter and needs none.',
  guildsButton: 'Guilds',
  progressButton: 'Progress',
} as const;

/** The home page's Rankings panel (spec 2026-09-23 §2 item 3): the top five guilds by
 *  progression, from `GET /v1/rankings/guilds?kind=progress` (needs no encounter). */
export const homeTopGuildsCopy = {
  failed: 'Rankings did not load.',
  empty: 'No kills ranked yet.',
  bossesUnit: 'bosses',
} as const;
