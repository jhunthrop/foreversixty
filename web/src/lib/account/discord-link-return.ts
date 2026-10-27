// web/src/lib/account/discord-link-return.ts
// What a return to /account?discord_link=<flag>#you says and does (experience spec
// 2026-10-27-discord-bot §3.2): the callback's four outcomes, and `start`, which a visitor who
// signed in on the way from Discord lands with and which goes straight on to the start route.
import { discordLinkCopy } from './discord-link-copy';

export type DiscordLinkFlag = 'linked' | 'denied' | 'failed' | 'conflict' | 'start';

export interface DiscordLinkOutcome {
  /** The status line; empty for `start`, which leaves the page. */
  message: string;
  role: 'status' | 'alert';
  /** The control that takes focus once it exists. */
  focus: string;
  /** The 13 px muted line under the status line (the conflict only). */
  help?: string;
  /** Hold the linked height until the fresh `/v1/me` answers: the held one predates the link. */
  refresh: boolean;
  /** Go on to the start route instead of drawing a message. */
  navigate: boolean;
}

export const LINK_BUTTON_SELECTOR = '[data-testid="discord-link-start"]';
export const UNLINK_BUTTON_SELECTOR = '[data-testid="discord-unlink"]';

const OUTCOMES: Readonly<Record<DiscordLinkFlag, DiscordLinkOutcome>> = {
  linked: {
    message: discordLinkCopy.returnLinked,
    role: 'status',
    focus: UNLINK_BUTTON_SELECTOR,
    refresh: true,
    navigate: false,
  },
  denied: {
    message: discordLinkCopy.returnDenied,
    role: 'status',
    focus: LINK_BUTTON_SELECTOR,
    refresh: false,
    navigate: false,
  },
  failed: {
    message: discordLinkCopy.returnFailed,
    role: 'alert',
    focus: LINK_BUTTON_SELECTOR,
    refresh: false,
    navigate: false,
  },
  conflict: {
    message: discordLinkCopy.returnConflict,
    role: 'alert',
    focus: LINK_BUTTON_SELECTOR,
    help: discordLinkCopy.conflictHelp,
    refresh: false,
    navigate: false,
  },
  start: {
    message: '',
    role: 'status',
    focus: LINK_BUTTON_SELECTOR,
    refresh: false,
    navigate: true,
  },
};

export const discordLinkReturn = (flag: DiscordLinkFlag): DiscordLinkOutcome => OUTCOMES[flag];

const FLAG_PARAM = 'discord_link';
const FLAGS = Object.keys(OUTCOMES) as readonly DiscordLinkFlag[];

/** The `discord_link` flag in a query string; any other value is no flag. */
export function discordLinkFlag(search: string): DiscordLinkFlag | null {
  const value = new URLSearchParams(search).get(FLAG_PARAM);
  return FLAGS.find((flag) => flag === value) ?? null;
}

/**
 * Reads the `?discord_link=` flag once and drops it from the address bar, keeping the path, any
 * other query and the hash, so a reload shows no second status line. A value that is not a flag
 * is removed too and reads as no flag.
 */
export function takeDiscordLinkFlag(
  location: Location = window.location,
  history: History = window.history,
): DiscordLinkFlag | null {
  const params = new URLSearchParams(location.search);
  if (!params.has(FLAG_PARAM)) return null;
  const flag = discordLinkFlag(location.search);
  params.delete(FLAG_PARAM);
  const query = params.toString();
  history.replaceState(
    history.state,
    '',
    `${location.pathname}${query === '' ? '' : `?${query}`}${location.hash}`,
  );
  return flag;
}
