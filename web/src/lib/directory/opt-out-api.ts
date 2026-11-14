// web/src/lib/directory/opt-out-api.ts
// The "Remove my character" page's three calls (api/internal/directory/optout_handler.go).
import { requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';

export type OptOutRegion = 'us' | 'eu';
export type OptOutState = 'removed' | 'waiting' | 'none';

export interface OptOutCharacter {
  region: OptOutRegion;
  name: string;
}

export interface OptOutCode {
  code: string;
  expires_at: string;
}

const BASE = '/v1/directory/opt-out';

async function post<T>(path: string, who: OptOutCharacter, apiBase: string): Promise<T> {
  const { data } = await requestEnvelope<T>(`${BASE}/${path}`, apiBase, {
    method: 'POST',
    body: { region: who.region, name: who.name.trim() },
  });
  if (data === null) throw new Error(`opt-out ${path}: no data`);
  return data;
}

/** A code for the character's public guild note; a new one replaces the last. */
export function startOptOutCode(who: OptOutCharacter, apiBase: string = API_BASE_URL): Promise<OptOutCode> {
  return post<OptOutCode>('code', who, apiBase);
}

/** Where the character's opt-out stands. */
export async function optOutStatus(
  who: OptOutCharacter,
  apiBase: string = API_BASE_URL,
): Promise<OptOutState> {
  return (await post<{ state: OptOutState }>('status', who, apiBase)).state;
}

/** Removes one of the signed-in account's own Battle.net characters at once. */
export async function removeOwnCharacter(
  who: OptOutCharacter,
  apiBase: string = API_BASE_URL,
): Promise<void> {
  await post<{ state: OptOutState }>('account', who, apiBase);
}

/** A full character name as the form takes it: first name and surname, letters only. */
export function validCharacterName(name: string): boolean {
  return /^\p{L}{2,12}( \p{L}{2,12})?$/u.test(name.trim());
}
