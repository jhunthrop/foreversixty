// web/src/lib/account/discord-link-api.ts
// The account page's two Discord link calls (experience spec 2026-10-27-discord-bot §3.2 and
// the contract in §9): the start route is a navigation, the unlink is a DELETE that answers 204.
import { API_BASE_URL } from '../planner/config';
import { meKey, requestEnvelope } from './api';
import type { Me } from './api';
import { setQueryData } from '../data/query';

const LINK_PATH = '/v1/discord/link';

/** The API route that starts the link: a plain navigation, so a middle-click and a long-press work. */
export function discordLinkStartUrl(apiBase: string = API_BASE_URL): string {
  return `${apiBase}${LINK_PATH}/start`;
}

/**
 * Unlinks, then writes the unlinked account into the cached `/v1/me` as a new object (the given
 * `me` is untouched), so the row re-renders in place with no loading state. Rejects when the
 * call fails; the cache is then unchanged.
 */
export async function unlinkDiscord(me: Me, apiBase: string = API_BASE_URL): Promise<void> {
  await requestEnvelope<null>(LINK_PATH, apiBase, { method: 'DELETE' });
  setQueryData(meKey(apiBase), { ...me, discord_link: null } satisfies Me);
}
