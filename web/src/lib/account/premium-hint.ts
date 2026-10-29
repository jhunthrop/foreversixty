// The two hints the session leaves in storage so the next paint can reserve the right box
// before `/v1/me` answers (spec 4.4): who pays for Premium and how many guild cards to expect.
// A convenience, never a requirement: every read and write survives a throwing storage.
import type { PremiumSource } from '../billing/contract';
import { storage } from '../layout-reserve';

export const PREMIUM_SOURCE_KEY = 'fs.premiumSource';
export const ME_GUILD_COUNT_KEY = 'fs.meGuildCount';
export type StoredPremiumSource = 'own' | 'seat' | 'grant';

/** The three words the session stores; Header.astro's inline script accepts only these. */
export const STORED_SOURCES: readonly string[] = ['own', 'seat', 'grant'];
/** Spec §5.7: the skeleton always draws at least one card. */
export const MIN_GUILD_COUNT = 1;
export const MAX_GUILD_COUNT = 4;
/** Read when nothing usable is stored; checkout.astro's head script clamps the same way. */
export const DEFAULT_GUILD_COUNT = 2;

function remove(key: string): void {
  try {
    storage()?.removeItem(key);
  } catch {
    // storage is a hint only
  }
}

function setAttribute(doc: Document, value: StoredPremiumSource | null): void {
  if (value === null) delete doc.documentElement.dataset.premium;
  else doc.documentElement.dataset.premium = value;
}

function storedSource(source: PremiumSource): StoredPremiumSource | null {
  if (source === 'own_then_seat') return 'own';
  return source === 'own' || source === 'seat' || source === 'grant' ? source : null;
}

export function writePremiumHint(source: PremiumSource, doc: Document = document): void {
  const stored = storedSource(source);
  setAttribute(doc, stored);
  if (stored === null) return remove(PREMIUM_SOURCE_KEY);
  try {
    storage()?.setItem(PREMIUM_SOURCE_KEY, stored);
  } catch {
    // storage is a hint only
  }
}

export function readPremiumHint(): StoredPremiumSource | null {
  try {
    const value = storage()?.getItem(PREMIUM_SOURCE_KEY) ?? null;
    if (value === null) return null;
    if (STORED_SOURCES.includes(value)) return value as StoredPremiumSource;
  } catch {
    return null;
  }
  remove(PREMIUM_SOURCE_KEY);
  return null;
}

export function clearPremiumHint(doc: Document = document): void {
  setAttribute(doc, null);
  remove(PREMIUM_SOURCE_KEY);
}

export function clearMeGuildCount(): void {
  remove(ME_GUILD_COUNT_KEY);
}

export function writeMeGuildCount(count: number): void {
  try {
    storage()?.setItem(ME_GUILD_COUNT_KEY, String(count));
  } catch {
    // storage is a hint only
  }
}

export function readMeGuildCount(): number {
  try {
    const stored = storage()?.getItem(ME_GUILD_COUNT_KEY) ?? null;
    const count = stored === null || stored.trim() === '' ? NaN : Math.round(Number(stored));
    if (!Number.isFinite(count)) return DEFAULT_GUILD_COUNT;
    return Math.min(MAX_GUILD_COUNT, Math.max(MIN_GUILD_COUNT, count));
  } catch {
    return DEFAULT_GUILD_COUNT;
  }
}

/** The session's one write of the guild-count hint: pass `/v1/me`'s answer through. */
export function rememberMeHints<T extends { guilds: readonly unknown[] } | null>(me: T): T {
  if (me !== null) writeMeGuildCount(me.guilds.length);
  return me;
}
