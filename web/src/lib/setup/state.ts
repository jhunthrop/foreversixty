// web/src/lib/setup/state.ts
// The one read /setup's islands share: who is signed in and which companions the account holds.
// The progress chips and the three done lines all ask for it; the answer is fetched once.
import { fetchMeOnce, listCompanions, type Device, type Me } from '../account/api';
import type { SetupState } from './progress';

let pending: Promise<SetupState> | null = null;

/** Signed out is `{ me: null, devices: [] }`; a device list that would not load stays null, so no companion line is claimed. */
export function loadSetupState(): Promise<SetupState> {
  pending ??= (async (): Promise<SetupState> => {
    const me: Me | null = await fetchMeOnce().catch(() => null);
    if (me === null) return { me: null, devices: [] };
    const devices: Device[] | null = await listCompanions().catch(() => null);
    return { me, devices };
  })();
  return pending;
}
