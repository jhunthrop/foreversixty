// web/src/lib/setup/state.ts
// The one read /setup's island makes: who is signed in and which companions the account holds.
// A 401 is signed out; any other failure is `error`, never "signed out" (setup spec B1).
import { fetchMeOnce, listCompanions, type Device, type Me } from '../account/api';
import type { SessionState } from './steps';

export interface SetupRead {
  session: Exclude<SessionState, 'loading'>;
  me: Me | null;
  /** Null when the device list did not load: no companion line is claimed. */
  devices: readonly Device[] | null;
}

export async function loadSetupState(fetchMe: () => Promise<Me | null> = fetchMeOnce): Promise<SetupRead> {
  let me: Me | null;
  try {
    me = await fetchMe();
  } catch {
    return { session: 'error', me: null, devices: null };
  }
  if (me === null) return { session: 'out', me: null, devices: [] };
  const devices = await listCompanions().catch(() => null);
  return { session: 'in', me, devices };
}
