// The client-side half of the same contract the API's auth.safeNext (bnet.go:142) enforces
// server-side: a `next` value must be a same-site path, with no scheme and no host. Used
// wherever a page reads `next` out of its own query string before handing it to
// battlenetStartUrl or building a /login link -- so a crafted `?next=` can only ever send a
// visitor back to a path on this site, never off it.
export function safeNextPath(raw: string | null | undefined, fallback: string): string {
  if (raw === null || raw === undefined || raw === '') return fallback;
  if (!raw.startsWith('/')) return fallback;
  if (raw.startsWith('//') || raw.startsWith('/\\')) return fallback;
  // A scheme (e.g. "javascript:") is only a real risk before the first '/', but the leading
  // '/' check above already rejects "javascript:alert(1)" as not starting with '/'. This
  // second check catches a same-site-looking path that still smuggles a colon-scheme via an
  // encoded or backslash variant a browser might normalise.
  if (/^\/[a-zA-Z][a-zA-Z0-9+.-]*:/.test(raw)) return fallback;
  return raw;
}
