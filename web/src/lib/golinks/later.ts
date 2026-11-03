// web/src/lib/golinks/later.ts
// The tracked links' and the visit beacon's work that waits for the page: Header.astro's loader
// imports this after first paint (lib/after-paint.ts), so none of it sits ahead of the largest
// paint. Two jobs, neither of which sets a cookie or reads anything about the visitor:
//  1. On a signed-in page, send the slug arrival.ts kept to the API once; it records it as the
//     account's first link only when the account is new. Kept while signed out, for the sign-in.
//  2. Load the cookieless Cloudflare Web Analytics beacon, on the production host only.
import { csrfToken } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import {
  ARRIVAL_KEY,
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

/** The kept slug while it is fresh; a stale or malformed one is dropped. */
function keptArrival(storage: Storage, now: number): KeptArrival | null {
  const raw = storage.getItem(ARRIVAL_KEY);
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
  storage.removeItem(ARRIVAL_KEY);
  return null;
}

/** Sends the kept slug when the visitor is signed in; resolves when the answer is handled. */
export async function sendArrival(apiBase: string = API_BASE_URL): Promise<void> {
  let kept: KeptArrival | null;
  try {
    kept = keptArrival(localStorage, Date.now());
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
      body: JSON.stringify({ slug: kept.slug }),
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
