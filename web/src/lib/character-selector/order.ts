// web/src/lib/character-selector/order.ts
// The list order (decision 10): the current character, then the characters chosen on this
// browser (most recent choice first), then the rest newest sync first. A sync therefore
// reorders only the never-chosen tail, so the list never reshuffles under the player's hand.
import type { MeCharacter } from '../account/api';
import { CHARACTER_ORDER_STORAGE_KEY, MAX_CHOSEN_KEYS } from './config';

function storageOf(storage: Storage | undefined): Storage | null {
  if (storage !== undefined) return storage;
  try {
    return globalThis.localStorage ?? null;
  } catch {
    return null;
  }
}

/** The keys chosen on this browser, most recent first. Anything malformed reads as empty. */
export function readChosenKeys(storage?: Storage): string[] {
  const target = storageOf(storage);
  if (target === null) return [];
  try {
    const parsed: unknown = JSON.parse(target.getItem(CHARACTER_ORDER_STORAGE_KEY) ?? '[]');
    if (!Array.isArray(parsed)) return [];
    return parsed.filter((entry): entry is string => typeof entry === 'string').slice(0, MAX_CHOSEN_KEYS);
  } catch {
    return [];
  }
}

/** `key` moved to the front of `chosen`, de-duplicated and capped. A new list; no mutation. */
export function withChosenFirst(chosen: readonly string[], key: string): string[] {
  return [key, ...chosen.filter((existing) => existing !== key)].slice(0, MAX_CHOSEN_KEYS);
}

export function recordChosenKey(key: string, storage?: Storage): void {
  const target = storageOf(storage);
  if (target === null) return;
  try {
    target.setItem(CHARACTER_ORDER_STORAGE_KEY, JSON.stringify(withChosenFirst(readChosenKeys(target), key)));
  } catch {
    // Private browsing or a full quota: the order is a convenience, never surfaced.
  }
}

function capturedMs(character: MeCharacter): number {
  const parsed = Date.parse(character.build?.captured_at ?? '');
  return Number.isNaN(parsed) ? Number.NEGATIVE_INFINITY : parsed;
}

/** `characters` in selector order. `currentKey` is null when nothing in the list is current. */
export function orderCharacters(
  characters: readonly MeCharacter[],
  currentKey: string | null,
  chosenKeys: readonly string[],
): MeCharacter[] {
  const current = characters.filter((character) => character.key === currentKey);
  const chosen = chosenKeys
    .map((key) => characters.find((character) => character.key === key))
    .filter((character): character is MeCharacter => character !== undefined && character.key !== currentKey);
  const placed = new Set([...current, ...chosen].map((character) => character.key));
  const rest = characters
    .filter((character) => !placed.has(character.key))
    .sort((a, b) => capturedMs(b) - capturedMs(a));
  return [...current, ...chosen, ...rest];
}
