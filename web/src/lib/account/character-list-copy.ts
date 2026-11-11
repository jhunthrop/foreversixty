import { ctaCopy } from '../addon/cta-copy';
import { compactAgo } from '../dates';
import { fillShared } from '../copy/shared';
import { sharedCharactersUpdatedFrom } from '../copy/shared/characters.generated';
import { uiCopy } from '../ui/copy';
// web/src/lib/account/character-list-copy.ts
// Every visible string CharacterList.svelte and Account.svelte's ?refreshed=1 toast use.
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

function characterCount(count: number): string {
  return `${count} ${count === 1 ? 'character' : 'characters'}`;
}

const EMPTY_LEAD =
  'No Forever characters on this Battle.net account yet. Make one in game, then refresh from Battle.net.';
const EMPTY_BY_EMAIL_LEAD = 'No characters yet.';
const EMPTY_BY_EMAIL_FILL = `Sign in with Battle.net to bring your Forever characters here, or ${ctaCopy.installAddonAndCompanionInline}.`;
const EMPTY_KOREA = `Playing in Korea? ${ctaCopy.installAddonAndCompanion}, or paste an export.`;

export const characterListCopy = {
  heading: 'Characters',
  /** The row's equipped item level, as the client's character sheet names it. */
  itemLevel: (level: number): string => `Item level ${level}`,
  /** Battle.net is the primary path (2026-11-03): a Battle.net sign-in imports every Forever
   *  character, so an empty list means no Forever character on that account yet, an email
   *  sign-in, or a region Blizzard does not serve (KR); the addon covers the last two. */
  empty: `${EMPTY_LEAD} ${EMPTY_KOREA}`,
  /** /setup splits the line: the lead in the warning colour, the Korea route muted (rulings Q2). */
  emptyLead: EMPTY_LEAD,
  emptyKorea: EMPTY_KOREA,
  emptyByEmail: `${EMPTY_BY_EMAIL_LEAD} ${EMPTY_BY_EMAIL_FILL}`,
  emptyByEmailLead: EMPTY_BY_EMAIL_LEAD,
  emptyByEmailFill: EMPTY_BY_EMAIL_FILL,
  getSetUp: 'Get set up',
  getSetUpHref: '/setup',
  pasteAnExport: 'Paste an export',
  pasteHref: '/setup#paste',
  refreshFromBattlenet: 'Refresh from Battle.net',
  /** The account's Battle.net freshness: "Updated 2 h ago from Battle.net" (users.bnet_imported_at,
   *  which every sign-in, visit refresh and background refresh moves). */
  updatedFrom: (at: string, now?: Date): string =>
    fillShared(sharedCharactersUpdatedFrom, { ago: compactAgo(new Date(at), now) }),
  verified: 'Verified',
  /** The main character's pill, and the control on every alt. */
  main: 'Main',
  current: 'Current',
  /** A character row's link to its page (with `›`). */
  open: 'Open',
  setAsMain: 'Set as main',
  setAsMainFailed: 'Could not set your main just now.',
  levelPrefix: (level: number): string => `Level ${level}`,
  // --- Battle.net refresh (design/specs/2026-11-04-account-characters.md section 6) ---------
  refreshing: 'Refreshing…',
  refreshingAnnounce: 'Refreshing from Battle.net.',
  refreshError: 'Battle.net did not answer. Nothing changed.',
  tryAgain: uiCopy.retry,
  updatedJustNow: (count: number, added: number): string =>
    `Updated just now · ${characterCount(count)}${added > 0 ? ` · ${added} new` : ''}`,
  noChanges: (count: number): string => `No changes · ${characterCount(count)}`,
  /** spec 2026-09-22 §3.3: the export "how" explained once, here, for every row. */
  introBattlenetLine: 'Your characters come from Battle.net and update on their own.',
  introInstallAddonLink: ctaCopy.installAddonAndCompanion,
  introAddonTail: 'for bags, bank and live sync. In Korea, or Battle.net is down?',
  introPasteLink: 'Paste an export',
  introPasteTail: 'instead.',
  /** The account rows' and the sim landing rows' build-source pill (spec 2026-09-22 §3.1). */
  battlenetSource: 'Battle.net',
  addonSource: 'Addon',
  noBuildYet: 'No gear yet',
} as const;
