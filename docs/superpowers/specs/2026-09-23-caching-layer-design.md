# The caching layer: HTTP semantics on the API, one data module in the browser

**Date:** 2026-09-23
**Status:** APPROVED by the owner ("ship it"). Two lanes in parallel: `cache-api` (section 2)
and `cache-web` (section 3). Section 1 binds both; section 4 is the frozen contract; section
5 is the owner's edge step.

## 0. Today

- Some public API reads send `Cache-Control: public, max-age=N` (rankings, ratings, spec
  lists, builds, site pages, the reports feed). Per-user reads send `private, no-store` or
  nothing. No route sends an `ETag`, so a revalidation is always a full response.
- The API is served directly from Cloud Run's front end (`server: Google Frontend`); no
  edge cache sits in front of it.
- Every envelope carries a per-request `request_id`, so two identical answers have
  different bodies.
- In the browser, `/v1/me` is deduped per page and, since today, kept as a snapshot across
  loads (`web/src/lib/account/session-cache.ts`). Every other read (guild module: 15 call
  sites; simulator: 10; character, report, billing, ratings) fetches fresh per island per
  load, sharing nothing beyond the browser's HTTP cache.
- All web API modules go through one `requestEnvelope(path, apiBase, init)` in
  `web/src/lib/account/api.ts` (guild, sim and billing wrap it with their own error copy).

## 1. Global constraints (both lanes)

1. **HTTP is the freshness authority.** Client TTLs are ceilings never longer than the
   server's `max-age`; the server's headers alone decide what the edge and the browser may
   keep.
2. **Private data stays private.** Anything for a signed-in user is `private`, is cached in
   the browser only while the session cookie's readable half (`fs_csrf`) is present, and is
   cleared on sign-out and on a 401. Nothing that is a token or a secret is ever cached.
3. **Mutations are never cached** and always return the resource as it now stands.
4. **Nothing new animates; the states model holds**: a stale answer renders instantly as the
   ready state, a first visit shows the skeleton, a failed revalidation keeps the stale
   answer and logs nothing visible.
5. **Lane ownership.** `cache-api` owns `api/**` and `api/openapi.yaml`. `cache-web` owns
   `web/**`. The contract in section 4 is frozen; the web lane stubs it.
6. Lane house rules (`.superpowers/journeys/lane-common-*.md`): sonnet only, never opus,
   including the final review; commit rules; scoped checks; no broad `pkill`; stop servers;
   Lighthouse quoted before the web lane's final report.

## 2. Lane `cache-api`: ETags and a header policy on every read

### 2.1 ETags, centrally

`httpx.write` (`api/internal/httpx/envelope.go`) gains conditional-request handling for
successful `GET` responses:

- Compute `ETag: W/"<first 16 hex of sha256 of the marshalled data field>"` over `data`
  only, never the whole envelope (the `request_id` differs per request).
- When the request carries `If-None-Match` and a listed tag matches (handle the `W/` prefix
  and comma lists), answer `304 Not Modified` with the same `ETag` and `Cache-Control`
  headers and no body.
- `Vary: Cookie, Authorization` on every private response, so a shared cache never keys a
  user's answer for another.

`WriteOK`'s signature does not change; the handler that already set `Cache-Control` keeps
it. Tests: a 200 with a tag, a 304 on match, a 200 on mismatch, weak-tag and list parsing,
`POST` never tagged.

### 2.2 The header policy

One helper pair in `httpx`: `CachePublic(w, maxAge, staleWhileRevalidate time.Duration)`
and `CachePrivate(w)` (`private, no-cache` plus the `Vary`). Every `GET` route mounted in
the API applies one of them (audit all 40; the lane lists each route and its class in the
plan). Classes, with values:

| Class | Routes | Header |
|---|---|---|
| Live public boards | rankings, ratings, reports feed, guild public page, character public page | `public, max-age=30, stale-while-revalidate=300` |
| Public per-object reads | a public report's meta and fights, public sim results, a public character's sim input | `public, max-age=60, stale-while-revalidate=600` |
| Reference | spec lists, builds, static data, site pages | keep today's values (an hour) |
| Private | `/v1/me`, devices, your reports, your sims, guild settings/home for members, private reports and sims | `private, no-cache` + `Vary` |

A route whose answer depends on the caller (a report that is public for one viewer and
private for another) is private. `httpx/list.go`'s `private, no-store` becomes
`CachePrivate` so 304s work for lists too. The lane confirms with a test per class that the
mounted route sends the expected header.

### 2.3 Trusted proxy hops

`trustedProxyHops` (rate limiting reads the client IP through `X-Forwarded-For`) becomes
configuration `TRUSTED_PROXY_HOPS` (default the current value) with a README line, so the
owner's edge step in section 5 is a config change, not a deploy of code.

## 3. Lane `cache-web`: one data module, every island through it

### 3.1 `web/src/lib/data/query.ts`

```ts
type Scope = 'public' | 'private';
interface QueryOptions<T> { scope: Scope; ttlMs: number; version?: number; parse?: (raw: unknown) => T }
interface QueryState<T> { data: T | null; status: 'idle' | 'loading' | 'ready' | 'failed'; error: string; stale: boolean }
export function query<T>(key: string, load: () => Promise<T>, options: QueryOptions<T>): Promise<T>
export function subscribe<T>(key: string, fn: (state: QueryState<T>) => void): () => void
export function setQueryData<T>(key: string, data: T): void
export function invalidate(prefix: string): void
export function forgetPrivate(): void
```

- `key` is `${apiBase}${path}`; the module is scope-aware, not URL-aware: the API modules
  compute keys.
- **In-page dedupe**: one in-flight promise per key, shared.
- **Stale while revalidate**: a stored entry younger than `ttlMs` resolves immediately and a
  revalidation runs in the background; when the live answer differs (structural compare),
  subscribers are called with the new state. An entry older than the TTL is not shown;
  the load runs as a first visit does.
- **Persistence by scope**: `public` entries persist in `localStorage` under `fs.q.<hash>`
  with `{v, savedAt, data}`; `private` entries persist the same way but are read only while
  `sessionHinted()` (`session-cache.ts`) and are removed by `forgetPrivate()`, which `signOut`
  calls and which a 401 from any private read triggers. A `version` mismatch drops the entry.
  Storage failures are swallowed (private mode, quota); the module works memory-only then.
- **Size cap**: at most 64 persisted entries, oldest evicted; an entry over 256 KB is not
  persisted.
- **ETags**: `requestEnvelope` remembers the `ETag` it last saw per key and sends
  `If-None-Match`; a 304 resolves to the stored data and refreshes `savedAt`. This is the
  cheap revalidation path the API lane enables.
- `session-cache.ts` and the special-case `fetchMeOnce` collapse into this module:
  `fetchMeOnce` becomes `query('/v1/me', …, {scope: 'private', ttlMs: 10 min})` and
  `ME_UPDATED` becomes a subscription; the tests move with them.

### 3.2 Reads migrated onto it

Every read function in `web/src/lib/{account,guild,sim,billing,rankings,report}/api.ts` and
`report/load.ts`, `planner/load.ts` goes through `query` with its class:

| Class | TTL | Scope |
|---|---|---|
| Session and the account's own lists (`/v1/me`, devices, my reports, my sims) | 10 min | private |
| Guild home/settings for a member | 5 min | private |
| Public character, guild page, ratings, rankings, reports feed | 60 s | public |
| A public report's meta and fights, a public sim's result | 5 min | public |
| Spec lists, builds, reference | 1 h | public |

Mutations (`PATCH /v1/me`, pairing, revoke, claim, settings, main character) call
`setQueryData` with the returned resource where the API returns it and `invalidate` for
their dependents otherwise. The account page's `setMain` writes the returned `Me` into the
`/v1/me` entry.

### 3.3 Islands render from it

A small Svelte 5 helper `createQueryState(key, load, options)` in
`web/src/lib/data/query.svelte.ts` returns rune-backed `{data, status, error, stale, refresh}`
by calling `query` and subscribing. The islands that render account or page data
(`HomeAccountPanel`, `Account`, `SessionNav`, `Guild*`, `Character`, `Rankings`,
`RecentReports`, `MyReports`, `HomeTopGuilds`, the sim landing's characters list) use it
in place of their own fetch-and-set effects. Their skeletons show only when `status ===
'loading'` with no data; a stale answer is rendered as ready. Retry buttons call `refresh`.

### 3.4 Tests and evidence

Unit tests for the module (dedupe, stale-while-revalidate, TTL, scope rules, version drop,
eviction, 304 handling with a stubbed `fetch`). The existing account, guild, rankings and
home specs pass; one new Playwright spec per class shows the second load rendering before
the API answers (the same shape as `home-panel.spec.ts`'s returning-visitor test). Lighthouse
quoted per URL: CLS and TBT must not move.

## 4. The frozen contract

- Successful `GET` responses carry `ETag: W/"…"` computed over `data`; `If-None-Match` with
  a matching tag answers `304` with no body and the same cache headers.
- Cache headers are exactly section 2.2's per class.
- Nothing about envelope shapes changes.

## 5. The owner's edge step (after both lanes deploy)

Proxy `api.foreversixty.gg` through Cloudflare (orange-cloud the DNS record) so public
reads are served from the edge under their `Cache-Control`, and set `TRUSTED_PROXY_HOPS` to
count the extra hop so per-IP rate limits keep seeing the client. The coordinator does this
with the owner's go-ahead once the headers are live; it is reversible in one click.
