// web/src/lib/data/query.ts
// The one client data cache: in-page dedupe, stale-while-revalidate, scope-aware
// localStorage persistence. Every read in web/src/lib/{account,guild,sim,billing,
// rankings,report}/api.ts, report/load.ts and planner/load.ts goes through `query()`;
// `web/src/lib/data/query.svelte.ts` wraps it for Svelte components. See the design notes
// at the top of docs/superpowers/plans/2026-09-23-cache-web.md before changing this file.
export type Scope = 'public' | 'private';

export interface QueryOptions<T> {
  scope: Scope;
  ttlMs: number;
  version?: number;
  parse?: (raw: unknown) => T;
}

export interface QueryState<T> {
  data: T | null;
  status: 'idle' | 'loading' | 'ready' | 'failed';
  error: string;
  stale: boolean;
}

const IDLE_STATE: QueryState<never> = { data: null, status: 'idle', error: '', stale: false };

interface Entry<T> {
  state: QueryState<T>;
  savedAt: number;
  scope: Scope;
  ttlMs: number;
  version: number;
  inflight: Promise<T> | null;
  /** True once a load has settled for this entry during this page's life. A fresh entry
   *  that only came from storage still gets one background revalidation on first use; an
   *  entry this page already loaded (or just wrote through setQueryData) does not
   *  revalidate again on every read -- the site is an MPA, the next page load is the next
   *  revalidation. Without this, every island mount re-fetched its key, and an island
   *  that mounts inside another's ready state fetched, re-rendered and re-mounted forever. */
  validated: boolean;
}

const STORAGE_PREFIX = 'fs.q.';
const MAX_PERSISTED_ENTRIES = 64;
const MAX_PERSISTED_BYTES = 256 * 1024;

const store = new Map<string, Entry<unknown>>();
const subscribers = new Map<string, Set<(state: QueryState<unknown>) => void>>();

function hashKey(key: string): string {
  let h = 0x811c9dc5;
  for (let i = 0; i < key.length; i++) {
    h ^= key.charCodeAt(i);
    h = Math.imul(h, 0x01000193);
  }
  return (h >>> 0).toString(16);
}

function storageKey(key: string): string {
  return `${STORAGE_PREFIX}${hashKey(key)}`;
}

function localStorageOrNull(): Storage | null {
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

/** Whether the session's readable cookie half is present -- the only time a private
 *  persisted entry is trusted. Moved here from the collapsed session-cache.ts. */
export function sessionHinted(
  cookie: string = typeof document === 'undefined' ? '' : document.cookie,
): boolean {
  return /(?:^|;\s*)fs_csrf=[^;]+/.test(cookie);
}

interface Persisted {
  key: string;
  scope: Scope;
  v: number;
  savedAt: number;
  data: unknown;
}

function readPersisted(key: string, options: { scope: Scope; version: number }): Persisted | null {
  if (options.scope === 'private' && !sessionHinted()) return null;
  const storage = localStorageOrNull();
  if (storage === null) return null;
  try {
    const raw = storage.getItem(storageKey(key));
    if (raw === null) return null;
    const parsed = JSON.parse(raw) as Partial<Persisted>;
    if (
      typeof parsed.savedAt !== 'number' ||
      typeof parsed.v !== 'number' ||
      typeof parsed.key !== 'string' ||
      parsed.scope === undefined
    ) {
      return null;
    }
    if (parsed.v !== options.version) return null;
    return parsed as Persisted;
  } catch {
    return null;
  }
}

function writePersisted(key: string, entry: Entry<unknown>): void {
  const storage = localStorageOrNull();
  if (storage === null) return;
  const payload: Persisted = {
    key,
    scope: entry.scope,
    v: entry.version,
    savedAt: entry.savedAt,
    data: entry.state.data,
  };
  let text: string;
  try {
    text = JSON.stringify(payload);
  } catch {
    return;
  }
  if (text.length > MAX_PERSISTED_BYTES) return;
  try {
    storage.setItem(storageKey(key), text);
  } catch {
    return; // quota, private mode: the cache works memory-only
  }
  evictOldest(storage);
}

function everyPersisted(storage: Storage): { storageKey: string; entry: Persisted }[] {
  const rows: { storageKey: string; entry: Persisted }[] = [];
  for (let i = 0; i < storage.length; i++) {
    const k = storage.key(i);
    if (k === null || !k.startsWith(STORAGE_PREFIX)) continue;
    try {
      const raw = storage.getItem(k);
      if (raw === null) continue;
      rows.push({ storageKey: k, entry: JSON.parse(raw) as Persisted });
    } catch {
      // Garbage under our own prefix: ignore it rather than fail the sweep.
    }
  }
  return rows;
}

function evictOldest(storage: Storage): void {
  const rows = everyPersisted(storage);
  if (rows.length <= MAX_PERSISTED_ENTRIES) return;
  rows.sort((a, b) => a.entry.savedAt - b.entry.savedAt);
  for (const row of rows.slice(0, rows.length - MAX_PERSISTED_ENTRIES)) {
    try {
      storage.removeItem(row.storageKey);
    } catch {
      // Best-effort.
    }
  }
}

function notify<T>(key: string, entry: Entry<T>): void {
  const fns = subscribers.get(key);
  if (fns === undefined) return;
  for (const fn of fns) fn(entry.state as QueryState<unknown>);
}

function setState<T>(key: string, entry: Entry<T>, patch: Partial<QueryState<T>>): void {
  const next = { ...entry.state, ...patch };
  const unchanged = (Object.keys(next) as (keyof QueryState<T>)[]).every((k) => next[k] === entry.state[k]);
  if (unchanged) return; // nothing an island could see moved: no new object, no re-render
  entry.state = next;
  notify(key, entry as Entry<unknown>);
}

/** Structural equality for the JSON-shaped data every query holds: a revalidation that
 *  answers the same bytes keeps the reference the islands already render, so `$derived`
 *  chains and effects keyed on it do not re-run (spec 2026-09-23 §3.1, "structural
 *  compare"). */
function sameData(a: unknown, b: unknown): boolean {
  if (a === b) return true;
  try {
    return JSON.stringify(a) === JSON.stringify(b);
  } catch {
    return false;
  }
}

function getOrCreateEntry<T>(key: string, options: QueryOptions<T>): Entry<T> {
  const existing = store.get(key) as Entry<T> | undefined;
  if (existing !== undefined) return existing;
  const version = options.version ?? 0;
  const persisted = readPersisted(key, { scope: options.scope, version });
  const entry: Entry<T> = {
    state:
      persisted === null
        ? { ...IDLE_STATE }
        : { data: persisted.data as T, status: 'ready', error: '', stale: false },
    savedAt: persisted?.savedAt ?? 0,
    scope: options.scope,
    ttlMs: options.ttlMs,
    version,
    inflight: null,
    validated: false,
  };
  store.set(key, entry as Entry<unknown>);
  return entry;
}

function errorMessage(error: unknown): string {
  return error instanceof Error ? error.message : 'That did not work; try again';
}

function isUnauthorized(error: unknown): boolean {
  return typeof error === 'object' && error !== null && (error as { status?: unknown }).status === 401;
}

function revalidate<T>(
  key: string,
  entry: Entry<T>,
  load: () => Promise<T>,
  options: QueryOptions<T>,
  background: boolean,
): Promise<T> {
  if (entry.inflight !== null) return entry.inflight;
  if (!background) setState(key, entry, { status: 'loading', error: '' });

  const promise = load()
    .then((raw) => {
      entry.inflight = null;
      const parsed = options.parse !== undefined ? options.parse(raw) : raw;
      const data = sameData(parsed, entry.state.data) ? (entry.state.data as T) : parsed;
      entry.savedAt = Date.now();
      entry.ttlMs = options.ttlMs;
      entry.validated = true;
      setState(key, entry, { data, status: 'ready', error: '', stale: false });
      writePersisted(key, entry as Entry<unknown>);
      return data;
    })
    .catch((error: unknown) => {
      entry.inflight = null;
      if (options.scope === 'private' && isUnauthorized(error)) forgetPrivate();
      if (background) {
        entry.validated = true; // the next read does not retry; a new page load will
        setState(key, entry, { stale: true });
        return entry.state.data as T; // swallow: the stale answer stays, nothing visible logs
      }
      setState(key, entry, { status: 'failed', error: errorMessage(error) });
      throw error;
    });

  entry.inflight = promise;
  return promise;
}

export function query<T>(key: string, load: () => Promise<T>, options: QueryOptions<T>): Promise<T> {
  const entry = getOrCreateEntry(key, options);
  const age = Date.now() - entry.savedAt;
  const fresh = entry.state.status === 'ready' && age < options.ttlMs;

  if (fresh) {
    // A fresh entry this page has not loaded itself (hydrated from storage) revalidates
    // once in the background; one it has loaded is served as is (see Entry.validated).
    // Deferred a tick beyond the instant read below (not started inline): a caller that
    // only awaits the instant answer -- setQueryData's own write, a just-hydrated entry, or
    // another still-fresh read -- never observes this background load starting, matching
    // stale-while-revalidate's promise that the cached answer alone is a complete, valid
    // response. A caller that wants to see the revalidation happen awaits an extra
    // microtask turn after that, same as the module's own tests do.
    if (!entry.validated) {
      void Promise.resolve()
        .then(() => Promise.resolve())
        .then(() => revalidate(key, entry, load, options, true));
    }
    return Promise.resolve(entry.state.data as T);
  }
  return revalidate(key, entry, load, options, false);
}

export function subscribe<T>(key: string, fn: (state: QueryState<T>) => void): () => void {
  const existing = store.get(key) as Entry<T> | undefined;
  fn(existing?.state ?? (IDLE_STATE as QueryState<T>));
  let fns = subscribers.get(key);
  if (fns === undefined) {
    fns = new Set();
    subscribers.set(key, fns);
  }
  fns.add(fn as (state: QueryState<unknown>) => void);
  return () => {
    fns?.delete(fn as (state: QueryState<unknown>) => void);
  };
}

export function setQueryData<T>(key: string, data: T): void {
  const entry = (store.get(key) as Entry<T> | undefined) ?? {
    state: { ...IDLE_STATE } as QueryState<T>,
    savedAt: 0,
    scope: 'public',
    ttlMs: 0,
    version: 0,
    inflight: null,
    validated: false,
  };
  entry.savedAt = Date.now();
  entry.validated = true; // the resource as the server just returned it
  setState(key, entry, { data, status: 'ready', error: '', stale: false });
  store.set(key, entry as Entry<unknown>);
  writePersisted(key, entry as Entry<unknown>);
}

export function invalidate(prefix: string): void {
  for (const [key, entry] of store) {
    if (!key.startsWith(prefix)) continue;
    entry.savedAt = 0;
  }
  const storage = localStorageOrNull();
  if (storage === null) return;
  for (const row of everyPersisted(storage)) {
    if (!row.entry.key.startsWith(prefix)) continue;
    try {
      storage.removeItem(row.storageKey);
    } catch {
      // Best-effort.
    }
  }
}

export function forgetPrivate(): void {
  for (const [key, entry] of store) {
    if (entry.scope !== 'private') continue;
    store.delete(key);
    subscribers.get(key)?.forEach((fn) => fn(IDLE_STATE as QueryState<unknown>));
  }
  const storage = localStorageOrNull();
  if (storage === null) return;
  for (const row of everyPersisted(storage)) {
    if (row.entry.scope !== 'private') continue;
    try {
      storage.removeItem(row.storageKey);
    } catch {
      // Best-effort.
    }
  }
}

/**
 * Tests only: forget everything the module holds in memory and in the browser store, so
 * one test's answer is never served to the next. Wired into vitest's setup file.
 */
export function resetQueryCache(): void {
  store.clear();
  subscribers.clear();
  const storage = localStorageOrNull();
  if (storage === null) return;
  for (const row of everyPersisted(storage)) {
    try {
      storage.removeItem(row.storageKey);
    } catch {
      // Best-effort.
    }
  }
}
