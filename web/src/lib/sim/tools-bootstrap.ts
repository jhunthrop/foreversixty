// web/src/lib/sim/tools-bootstrap.ts
// Pure decision logic for the tools island's own bootstrap (ToolsView.svelte's onMount),
// split out here because that onMount is not itself unit-testable -- Task 4, the
// current-character spec (2026-09-21) section 1: every sim tab loads the current character
// with no re-paste, precedence `?code=`, then `?source=&ref=`, then the stored pointer,
// then nothing. Nothing here touches `window` or `localStorage`; the caller reads both and
// passes in already-parsed values.
import { parseCharacterPath, type CharacterPath } from '../characters';
import type { CurrentCharacter } from '../current-character';
import type { LootSource } from './loot';
import type { SourceKind } from './types';

/** The three URL pieces ToolsView.svelte's own bootstrap reads, already bounded by
 *  `parseSimState`/`URLSearchParams` -- this module never reads the query string itself. */
export interface ToolsBootstrapUrl {
  code: string;
  source: SourceKind | '';
  ref: string;
}

export type ToolsBootstrapLoad =
  | { kind: 'code'; code: string }
  | { kind: 'addon'; code: string }
  | { kind: 'build'; id: string }
  | { kind: 'fight'; ref: string }
  | { kind: 'stored'; path: CharacterPath }
  | { kind: 'none' };

/** What to load, plus whether that came from the stored pointer rather than the URL --
 *  `ToolsView.svelte` uses `restored` to show the "Restored your last character" line. */
export type ToolsBootstrapDecision = ToolsBootstrapLoad & { restored: boolean };

const NOTHING: ToolsBootstrapDecision = { kind: 'none', restored: false };

/** `source`/`ref` (a URL query or the stored pointer, `armory`'s ref an armory character
 *  key) resolved to a `stored` load -- shared so the URL and pointer paths cannot drift. */
function toStoredLoad(ref: string, restored: boolean): ToolsBootstrapDecision {
  const path = parseCharacterPath(`/character/${ref}`);
  return path === null ? NOTHING : { kind: 'stored', path, restored };
}

/** The stored pointer alone, once the URL has offered neither `?code=` nor `?source=&ref=`. */
function fromStoredPointer(stored: CurrentCharacter): ToolsBootstrapDecision {
  if (stored.source === 'addon' || stored.source === 'code') {
    return { kind: 'code', code: stored.ref, restored: true };
  }
  if (stored.source === 'build') return { kind: 'build', id: stored.ref, restored: true };
  if (stored.source === 'fight') return { kind: 'fight', ref: stored.ref, restored: true };
  // The only source left is 'armory': the same character-key ref `?source=armory&ref=<key>`
  // carries, parsed the same way -- `store.loadStored` (bulk-store.svelte.ts) takes a
  // structured `CharacterPath`, not this pointer's bare ref string.
  return toStoredLoad(stored.ref, true);
}

/**
 * What the tools island should load, and whether that came from the stored pointer rather
 * than the URL. `ToolsView.svelte`'s `onMount` calls this once and switches on the result.
 */
export function decideToolsBootstrap(
  url: ToolsBootstrapUrl,
  stored: CurrentCharacter | null,
): ToolsBootstrapDecision {
  if (url.code !== '') return { kind: 'code', code: url.code, restored: false };
  if (url.source !== '' && url.ref !== '') {
    if (url.source === 'addon') return { kind: 'addon', code: url.ref, restored: false };
    if (url.source === 'build') return { kind: 'build', id: url.ref, restored: false };
    if (url.source === 'fight') return { kind: 'fight', ref: url.ref, restored: false };
    // 'manual' has no ref-shaped loader either (store.svelte.ts's own `bootstrapSource`
    // agrees): a link naming it bootstraps nothing rather than guessing at one.
    return url.source === 'armory' ? toStoredLoad(url.ref, false) : NOTHING;
  }
  return stored === null ? NOTHING : fromStoredPointer(stored);
}

/** Characters of a `?instance=` slug the query string is checked against before it ever
 *  reaches a `LootSource.id` comparison -- the query is attacker-controlled. Every real
 *  zone slug (`molten-core`, `hall-of-thanes`, …) is well under this. */
const MAX_INSTANCE_SLUG_LENGTH = 64;
const INSTANCE_SLUG_PATTERN = /^[a-z0-9-]+$/;

function isInstanceSlug(value: string): boolean {
  return value !== '' && value.length <= MAX_INSTANCE_SLUG_LENGTH && INSTANCE_SLUG_PATTERN.test(value);
}

/**
 * The `LootSource.id` (`'<kind>:<zone-slug>'`, loot.ts) whose zone matches `?instance=`, or
 * null when the slug is empty, malformed, or matches nothing -- `/sim/drops`'s own
 * preselect (`ToolsView.svelte`'s instance `$effect`).
 */
export function sourceIdForInstance(sources: readonly LootSource[], slug: string): string | null {
  if (!isInstanceSlug(slug)) return null;
  return sources.find((source) => source.id.endsWith(`:${slug}`))?.id ?? null;
}
