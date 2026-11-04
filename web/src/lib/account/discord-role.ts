// web/src/lib/account/discord-role.ts
// The account page's Discord Premium role row: the API's verification start (also the Role
// Connections Verification URL Discord opens) and what a return to /account?discord_role=<flag>
// says. `start` is a visitor who signed in on the way and goes straight on to the start route.
import { API_BASE_URL } from '../planner/config';
import { discordRoleCopy } from './discord-role-copy';
import { takeQueryFlag } from './discord-link-return';

export type DiscordRoleFlag = 'linked' | 'denied' | 'failed' | 'start';

export interface DiscordRoleOutcome {
  message: string;
  role: 'status' | 'alert';
  /** Go on to the start route instead of drawing a message. */
  navigate: boolean;
}

const OUTCOMES: Readonly<Record<DiscordRoleFlag, DiscordRoleOutcome>> = {
  linked: { message: discordRoleCopy.returnLinked, role: 'status', navigate: false },
  denied: { message: discordRoleCopy.returnDenied, role: 'status', navigate: false },
  failed: { message: discordRoleCopy.returnFailed, role: 'alert', navigate: false },
  start: { message: '', role: 'status', navigate: true },
};

const FLAG_PARAM = 'discord_role';
const FLAGS = Object.keys(OUTCOMES) as readonly DiscordRoleFlag[];

/** The API route that starts the role link: a plain navigation. */
export function discordRoleStartUrl(apiBase: string = API_BASE_URL): string {
  return `${apiBase}/v1/discord/linked-role`;
}

export const discordRoleReturn = (flag: DiscordRoleFlag): DiscordRoleOutcome => OUTCOMES[flag];

/** Reads `?discord_role=` once and drops it from the address bar. */
export function takeDiscordRoleFlag(
  location: Location = window.location,
  history: History = window.history,
): DiscordRoleFlag | null {
  return takeQueryFlag(FLAG_PARAM, FLAGS, location, history);
}
