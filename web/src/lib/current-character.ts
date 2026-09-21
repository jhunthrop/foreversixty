// web/src/lib/current-character.ts
// The site's one "current character" pointer: written on every successful character load
// (sources.ts), read by CurrentCharacterChip.svelte and by a bare /sim or /planner load to
// restore it. A convenience, never the source of truth -- every read and write is wrapped in
// try/catch and this module never touches `window`/`localStorage` at import time, only
// inside the functions below, so importing it is safe during SSR/prerender (Global
// Constraint: "never read it during SSR/prerender").
import { SIM_TABS, tabHref, type SimTabEntry } from './sim/tabs';
import { defaultSimState, withSimState } from './sim/url';

export type CurrentCharacterSource = 'addon' | 'build' | 'fight' | 'armory' | 'code';

export interface CurrentCharacter {
  source: CurrentCharacterSource;
  /** For 'addon' and 'code' the FS1 string itself; otherwise the existing ref format. */
  ref: string;
  /** What the chip shows: "Simfury · Fury Warrior". Derived once at load, never parsed back. */
  label: string;
  classSlug: string;
  savedAt: string; // ISO
}

export type SimTabId = SimTabEntry['id'];

const STORAGE_KEY = 'fs.currentCharacter';
/** An FS1 export with a full bank is several KB; this stays well clear of it while refusing
 *  a pathological write (the spec's own cap). */
const MAX_STORED_BYTES = 16_384;

function storageOf(storage: Storage | undefined): Storage | null {
  if (storage !== undefined) return storage;
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

export function readCurrent(storage?: Storage): CurrentCharacter | null {
  const target = storageOf(storage);
  if (target === null) return null;
  try {
    const raw = target.getItem(STORAGE_KEY);
    if (raw === null) return null;
    const parsed: unknown = JSON.parse(raw);
    if (!isCurrentCharacter(parsed)) return null;
    return parsed;
  } catch {
    return null;
  }
}

function isCurrentCharacter(value: unknown): value is CurrentCharacter {
  if (typeof value !== 'object' || value === null) return false;
  const candidate = value as Record<string, unknown>;
  return (
    typeof candidate.source === 'string' &&
    typeof candidate.ref === 'string' &&
    typeof candidate.label === 'string' &&
    typeof candidate.classSlug === 'string' &&
    typeof candidate.savedAt === 'string'
  );
}

export function writeCurrent(value: CurrentCharacter, storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    const serialised = JSON.stringify(value);
    if (serialised.length > MAX_STORED_BYTES) return;
    target.setItem(STORAGE_KEY, serialised);
  } catch {
    // Private browsing, quota exceeded, or a disabled storage API: the pointer is a
    // convenience, so a failed write is silently skipped rather than surfaced.
  }
}

export function clearCurrent(storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    target.removeItem(STORAGE_KEY);
  } catch {
    // Same as writeCurrent: nothing to surface.
  }
}

/** The saved-build permalink; not a page under web/src/pages (API-served, share.ts's own
 *  `cardUrlFor` comment: `PUBLIC_BASE_URL + "/b/" + id`). */
function buildPermalink(id: string): string {
  return `/b/${id}`;
}

export function plannerHrefFor(current: CurrentCharacter): string {
  if (current.source === 'addon' || current.source === 'code') {
    return `/planner?code=${encodeURIComponent(current.ref)}`;
  }
  if (current.source === 'build') return buildPermalink(current.ref);
  // 'fight' | 'armory': the pointer alone carries no gear or talents for these kinds -- the
  // same honest class-only fallback sim/character.ts's own plannerHrefFor uses when it has
  // no talent index to encode a full FS1 code from.
  return `/planner?class=${encodeURIComponent(current.classSlug)}`;
}

export function simHrefFor(current: CurrentCharacter, tab: SimTabId = 'quick-sim'): string {
  const entry = SIM_TABS.find((row) => row.id === tab) ?? SIM_TABS[0];
  if (current.source === 'addon' || current.source === 'code') {
    return tabHref(entry.href, withSimState(defaultSimState(), { code: current.ref }));
  }
  return tabHref(entry.href, withSimState(defaultSimState(), { source: current.source, ref: current.ref }));
}
