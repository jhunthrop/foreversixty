// web/src/lib/golinks/later.ts
// The tracked links' and the visit beacon's work that waits for the page: Header.astro's loader
// imports this after first paint (lib/after-paint.ts), so none of it sits ahead of the largest
// paint. Two jobs, neither of which sets a cookie or reads anything about the visitor:
//  1. Keep the first click arrival.ts saw (promoteArrival), and on a signed-in page send it with its
//     time to the API once; it records it as the account's first link only when the account is new.
//     Kept while signed out, for the sign-in.
//  2. Load the cookieless Cloudflare Web Analytics beacon, on the production host only.
import { csrfToken } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import {
  ARRIVAL_KEY,
  ARRIVAL_NEXT_KEY,
  ARRIVAL_PATH,
  ARRIVAL_TTL_MS,
  BEACON_HOST,
  BEACON_SKIP_PREFIXES,
  BEACON_SRC,
  BEACON_TOKEN,
  isSlug,
} from './config';

interface KeptArrival {
  slug: string;
  at: number;
}

/** The click kept under key while it is fresh; a stale or malformed one is dropped. */
function keptAt(storage: Storage, key: string, now: number): KeptArrival | null {
  const raw = storage.getItem(key);
  if (raw === null) return null;
  try {
    const kept = JSON.parse(raw) as Partial<KeptArrival> | null;
    const age = typeof kept?.at === 'number' ? now - kept.at : -1;
    if (typeof kept?.slug === 'string' && isSlug(kept.slug) && age >= 0 && age < ARRIVAL_TTL_MS) {
      return { slug: kept.slug, at: kept.at as number };
    }
  } catch {
    /* malformed: dropped below */
  }
  storage.removeItem(key);
  return null;
}

/**
 * First touch: the newest click the capture kept becomes the arrival only while no fresh earlier
 * one is kept, so an account is credited to the first link that brought it (within the TTL).
 */
export function promoteArrival(storage: Storage, now: number): KeptArrival | null {
  const next = keptAt(storage, ARRIVAL_NEXT_KEY, now);
  storage.removeItem(ARRIVAL_NEXT_KEY);
  const first = keptAt(storage, ARRIVAL_KEY, now);
  if (first !== null || next === null) return first;
  storage.setItem(ARRIVAL_KEY, JSON.stringify(next));
  return next;
}

/** Sends the kept slug when the visitor is signed in; resolves when the answer is handled. */
export async function sendArrival(apiBase: string = API_BASE_URL): Promise<void> {
  let kept: KeptArrival | null;
  try {
    kept = promoteArrival(localStorage, Date.now());
  } catch {
    return;
  }
  const csrf = csrfToken();
  if (kept === null || csrf === '') return;
  try {
    const response = await fetch(`${apiBase}${ARRIVAL_PATH}`, {
      method: 'POST',
      credentials: 'include',
      headers: { 'content-type': 'application/json', 'x-csrf-token': csrf },
      body: JSON.stringify({ slug: kept.slug, at: kept.at }),
    });
    // Sent (recorded or not), or refused for good. A 401 or 403 means the session is gone, so
    // the slug waits for the sign-in.
    if (response.ok || response.status === 400) localStorage.removeItem(ARRIVAL_KEY);
  } catch {
    /* offline: tried again on the next page */
  }
}

/** Adds the beacon on the production host, except on the pages that must load no third party. */
export function loadBeacon(host: string = BEACON_HOST): void {
  if (location.hostname !== host) return;
  if (BEACON_SKIP_PREFIXES.some((prefix) => location.pathname.startsWith(prefix))) return;
  if (document.querySelector(`script[src="${BEACON_SRC}"]`) !== null) return;
  const beacon = document.createElement('script');
  beacon.type = 'module';
  beacon.src = BEACON_SRC;
  beacon.setAttribute('data-cf-beacon', JSON.stringify({ token: BEACON_TOKEN, spa: false }));
  document.head.appendChild(beacon);
}

export function runLater(): void {
  loadBeacon();
  void sendArrival();
}
