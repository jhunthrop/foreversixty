// web/src/lib/sim/character-bootstrap.ts
// Decision logic (plus its own thin orchestration, `runBootstrapRestore`) for a /sim*
// page's own bootstrap fallback (ToolsView.svelte's and SimView.svelte's onMount), split
// out here because neither onMount is itself unit-testable -- Task 4, the
// current-character spec (2026-09-21) section 1: every sim tab loads the current character
// with no re-paste, precedence a whole share-link request, then `?code=`, then
// `?source=&ref=`, then the stored pointer, then nothing. `decideBootstrap` and
// `settleRestore` never touch `window` or storage -- the caller reads both and passes in
// already-parsed values; `runBootstrapRestore` is the one exception, and only ever to
// clear a dead pointer (`clearCurrent`), never to read one.
//
// Named for the concern (deciding, and dispatching, what character to bootstrap), not for
// either mounting page: Task 5 generalised this from `tools-bootstrap.ts` so SimView.svelte
// (which also has to weigh `?req=`, a whole request that wins over everything else -- design
// 8) can share it rather than re-deriving the same precedence and the same settle rule.
import { parseCharacterPath, type CharacterPath } from '../characters';
import { clearCurrent, type CurrentCharacter } from '../current-character';
import type { LootSource } from './loot';
import type { SourceKind } from './types';

/** The URL pieces a mounting page's own bootstrap reads, already bounded by
 *  `parseSimState`/`URLSearchParams` -- this module never reads the query string itself. */
export interface BootstrapUrl {
  code: string;
  source: SourceKind | '';
  ref: string;
  /** True when the URL also carries a whole share-link request (SimView's own `?req=`,
   *  design 8) -- that request wins over everything (`store.svelte.ts`'s own `ready`
   *  already applies it before this function is ever consulted), so a stored pointer is
   *  never restored over it. Optional and omitted by the tools island, which has no `?req=`
   *  of its own vocabulary. */
  hasRequest?: boolean;
}

export type BootstrapLoad =
  | { kind: 'code'; code: string }
  | { kind: 'addon'; code: string }
  | { kind: 'build'; id: string }
  | { kind: 'fight'; ref: string }
  | { kind: 'stored'; path: CharacterPath }
  | { kind: 'none' };

/** What to load, plus whether that came from the stored pointer rather than the URL --
 *  the mounting page uses `restored` to show the chip's "Restored your last character"
 *  line. */
export type BootstrapDecision = BootstrapLoad & { restored: boolean };

const NOTHING: BootstrapDecision = { kind: 'none', restored: false };

/** `source`/`ref` (a URL query or the stored pointer, `armory`'s ref an armory character
 *  key) resolved to a `stored` load -- shared so the URL and pointer paths cannot drift. */
function toStoredLoad(ref: string, restored: boolean): BootstrapDecision {
  const path = parseCharacterPath(`/character/${ref}`);
  return path === null ? NOTHING : { kind: 'stored', path, restored };
}

/** The stored pointer alone, once the URL has offered neither a whole request nor
 *  `?code=` nor `?source=&ref=`.
 *
 * 'addon' and 'code' both store the FS1 string itself (CurrentCharacter's own doc comment),
 * but they are not the same load: 'addon' means the pointer came from an addon export
 * paste, and BootstrapLoad already carries a distinct `{kind: 'addon'}` for exactly that --
 * loadAddon (store.svelte.ts) tags the resulting character `{kind: 'addon', ...}` the same
 * way a fresh paste does, where loadCode's `fromManualCode` always tags 'manual'. Folding
 * both into `{kind: 'code'}` here lost that distinction on restore: a character restored
 * from an addon paste came back tagged 'manual' and its source badge read "Entered by
 * hand", even though nothing was ever typed by hand (2026-09-21 result-page review round
 * 3, newcomer's own finding).
 */
function fromStoredPointer(stored: CurrentCharacter): BootstrapDecision {
  if (stored.source === 'addon') return { kind: 'addon', code: stored.ref, restored: true };
  if (stored.source === 'code') return { kind: 'code', code: stored.ref, restored: true };
  if (stored.source === 'build') return { kind: 'build', id: stored.ref, restored: true };
  if (stored.source === 'fight') return { kind: 'fight', ref: stored.ref, restored: true };
  // The only source left is 'armory': the same character-key ref `?source=armory&ref=<key>`
  // carries, parsed the same way -- `store.loadStored` (both stores) takes a structured
  // `CharacterPath`, not this pointer's bare ref string.
  return toStoredLoad(stored.ref, true);
}

/**
 * What a /sim* page should load, and whether that came from the stored pointer rather than
 * the URL. The mounting page's own `onMount` calls this once and switches on the result --
 * `decision.restored` alone tells the caller whether it is safe to actually start the load
 * this returns: false means either there is nothing to do, or the URL itself already won
 * (and the store's own init-time bootstrap is already handling it), so firing this decision
 * a second time would duplicate that load.
 */
export function decideBootstrap(url: BootstrapUrl, stored: CurrentCharacter | null): BootstrapDecision {
  if (url.hasRequest === true) return NOTHING;
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

/** The one method shape every /sim* store (`store.svelte.ts`, `bulk-store.svelte.ts`)
 *  exposes for a `BootstrapDecision` to run against -- `startBootstrapLoad` and
 *  `runBootstrapRestore` below are written once, against this shape, rather than once per
 *  mounting island. */
export interface CharacterLoaders {
  loadCode(code: string): Promise<void>;
  loadAddon(code: string): Promise<void>;
  loadBuild(id: string): Promise<void>;
  loadFight(ref: string): Promise<void>;
  loadStored(path: CharacterPath): Promise<void>;
  setMessage(text: string | null): void;
}

/** `decision.kind`'s own loader, or null for `'none'` -- the one place a `BootstrapDecision`
 *  is mapped onto a store's loaders, so a mounting page's own bootstrap reads as "decide,
 *  load, settle" rather than a second copy of this switch. */
export function startBootstrapLoad(
  loaders: CharacterLoaders,
  decision: BootstrapDecision,
): Promise<void> | null {
  if (decision.kind === 'code') return loaders.loadCode(decision.code);
  if (decision.kind === 'addon') return loaders.loadAddon(decision.code);
  if (decision.kind === 'build') return loaders.loadBuild(decision.id);
  if (decision.kind === 'fight') return loaders.loadFight(decision.ref);
  if (decision.kind === 'stored') return loaders.loadStored(decision.path);
  return null;
}

export interface RestoreSettlement {
  /** Whether `clearCurrent()` should run: a restore that settled with no character loaded
   *  -- a dead or stale pointer, not one worth trying again on the next visit. */
  clearPointer: boolean;
  /** The `restored` flag the mounting page should show its chip with. */
  restored: boolean;
  /** Whether the store's own `message` should be cleared too, so a first-time-looking
   *  page is shown rather than an error the player did nothing to cause. */
  clearMessage: boolean;
}

/**
 * What to do once a /sim* page's own bootstrap load has settled (fix round 1, Task 4's
 * review, Important: a dead stored pointer was showing "Restored" beside an error and would
 * be retried forever). A URL-driven load (`decision.restored` false) keeps today's behaviour
 * on failure -- nothing here is cleared, and the store's own message shows, exactly as it did
 * before this pointer existed. A load that came from the stored pointer only claims
 * "restored" once it actually produced a character.
 *
 * Spec 2026-09-25 section 4.2 ("the landing list only when [the character] does not [have a
 * build]"): an `'armory'` pointer (`decision.kind === 'stored'`) names a real character,
 * independent of whether the simulator can run it -- a build-less restore here is
 * `/sim`-specific absence, not a dead reference, and the same pointer is still exactly what
 * Planner, Logs and Rankings want. It alone survives a build-less restore (`clearPointer:
 * false`); every other kind (`code`, `addon`, `build`, `fight`) names an artifact that will
 * never resolve differently on a retry and is still forgotten, unchanged from before this
 * ruling. Either way the store's own refusal message is cleared (`clearMessage: true`): the
 * player never asked for this specific background load and should see the ordinary empty
 * state (or, for a survived armory pointer, the ordinary landing list), not an error for a
 * restore they cannot see or act on.
 */
export function settleRestore(decision: BootstrapDecision, characterLoaded: boolean): RestoreSettlement {
  if (!decision.restored) return { clearPointer: false, restored: false, clearMessage: false };
  if (characterLoaded) return { clearPointer: false, restored: true, clearMessage: false };
  const survivesEmptyLoad = decision.kind === 'stored';
  return { clearPointer: !survivesEmptyLoad, restored: false, clearMessage: true };
}

export interface RunBootstrapRestoreOptions {
  storage?: Storage;
  /**
   * True (the default) when the mounting store already ran the URL's own bootstrap by
   * itself before this is ever called -- SimView.svelte's `await store.ready`, which is
   * `store.svelte.ts`'s own `init.code`/`init.request`/`init.source`/`init.ref` handling,
   * resolved at construction. This function then fires only the stored-pointer fallback
   * (a URL-driven decision never re-dispatches, which would double-load the same
   * character). False for a store with no such init-time bootstrap of its own
   * (`bulk-store.svelte.ts`: its `BulkStoreInit.source`/`ref` are read by nothing at
   * construction) -- ToolsView.svelte passes false because this call is the ONLY place a
   * `?code=`/`?source=&ref=` load ever starts for the tools island, so it must actually
   * start one rather than assume something else already did (fix round: the assumption
   * this flag replaces silently dropped every direct-URL load on /sim/gear, /sim/drops,
   * /sim/talents and /sim/weights).
   */
  storeHandlesUrl?: boolean;
}

/**
 * The whole "decide, load, settle" sequence a /sim* page's own bootstrap fallback runs, in
 * one call -- both ToolsView.svelte and SimView.svelte call this rather than each
 * reimplementing the same three-step orchestration around `decideBootstrap`,
 * `startBootstrapLoad` and `settleRestore`. `characterLoaded` is a getter, not a plain
 * boolean, because it has to be read *after* the load settles, not when this is called.
 * Returns the `restored` flag the caller's own chip should render with. This is the one
 * function in the module that touches storage, and only ever to clear it
 * (`settleRestore`'s own `clearPointer`) -- `stored` is read by the caller, exactly as
 * `decideBootstrap` alone already required.
 */
export async function runBootstrapRestore(
  loaders: CharacterLoaders,
  url: BootstrapUrl,
  stored: CurrentCharacter | null,
  characterLoaded: () => boolean,
  options: RunBootstrapRestoreOptions = {},
): Promise<boolean> {
  const { storage, storeHandlesUrl = true } = options;
  const decision = decideBootstrap(url, stored);
  if (decision.kind === 'none') return false;
  if (!decision.restored && storeHandlesUrl) return false;
  const load = startBootstrapLoad(loaders, decision);
  if (load !== null) await load;
  const outcome = settleRestore(decision, characterLoaded());
  if (outcome.clearPointer) clearCurrent(storage);
  if (outcome.clearMessage) loaders.setMessage(null);
  return outcome.restored;
}

/**
 * Fix round 1, Important #1: SimView.svelte's `LandingState` shows a signed-in member's own
 * characters, each pickable through `store.loadStored`, which -- like every `adopt()` call
 * -- has no per-load generation guard (`LandingState.svelte`'s own comment on its button's
 * `disabled`). Before this, a background restore's own load raced a player's own pick with
 * nothing to stop it: `store.phase` (which gates `SourceSwitcher`) is set the moment the
 * restore's loader starts, but `LandingState` never looks at `store.phase` -- it only looks
 * at its own separately-tracked `busyKey` prop, which the restore never touched. A player
 * fast enough to click "Sim it" while the restore's own fetch was still in flight could have
 * their own pick clobbered by whichever load settled last, then mislabelled `restored: true`
 * or handed the restore's own refusal message.
 *
 * Never a bare string literal in the caller: `<region>/<ruleset>/<slug>` (LandingState's
 * own `busyKey === character.key` check, `pathOf`) always has exactly two `/`s, so this
 * sentinel -- with none -- can never collide with a real character key. SimView.svelte's
 * `restoreFromPointer` holds `landingBusyKey` at this value only once `decideBootstrap`
 * says a restore is actually happening (`.restored` true), and releases it in a `finally`
 * around `runBootstrapRestore` -- the same decision that function itself runs, so the two
 * can never disagree about whether one is in flight.
 */
export const RESTORE_BUSY_KEY = 'restoring-current-character';
