// web/src/lib/account/deleted-arrival.ts
// The one-time arrival after an account deletion: /account sends the visitor to `/?deleted=1`,
// and the home page reads the flag once, takes it out of the address bar (so a reload or a shared
// link never shows the notice again) and says the account is gone. The same read-once-then-drop
// pattern as /account's `?signed_in=1`.

const PARAM = 'deleted';

/** Where /account lands the visitor after a deletion. */
export const DELETED_ARRIVAL_HREF = `/?${PARAM}=1`;

/** Whether this visit arrived from a deletion; the flag is dropped from the URL either way. */
export function takeDeletedArrival(win: Pick<Window, 'location' | 'history'> = window): boolean {
  const url = new URL(win.location.href);
  if (!url.searchParams.has(PARAM)) return false;
  const arrived = url.searchParams.get(PARAM) === '1';
  url.searchParams.delete(PARAM);
  win.history.replaceState(win.history.state, '', url.pathname + url.search + url.hash);
  return arrived;
}
