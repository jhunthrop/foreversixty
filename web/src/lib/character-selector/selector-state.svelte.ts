// web/src/lib/character-selector/selector-state.svelte.ts
// The reactive half of the nav selector, shared by its two islands (the bar's selector and
// the phone Menu's first row): the cached `/v1/me`, the stored pointer and the keys chosen on
// this browser, folded into the view model of rows.ts. Islands share no store, so each island
// builds its own copy of this; the `/v1/me` read is one request through the client cache.
import { ME_QUERY_VERSION, ME_TTL_MS, fetchMeOnce, meKey, type Me } from '../account/api';
import { createQueryState } from '../data/query.svelte';
import { sessionHinted } from '../data/query';
import { onCurrentCharacterChange, readCurrent, type CurrentCharacter } from '../current-character';
import { API_BASE_URL } from '../planner/config';
import { readChosenKeys } from './order';
import { accountNameOf, buildSelectorModel, closedKindOf, type ClosedKind, type SelectorModel } from './rows';
import { guildPagePointer, memberGuildOf, type SessionGuild } from './session-guild';

export type LoadStatus = 'loading' | 'ready' | 'failed';

/** What the closed control shows: the server-rendered link, a character, or the sign-in slot. */
export type { ClosedKind };

export interface SelectorState {
  /** False until the island has hydrated and read the browser: the server-rendered link state. */
  readonly mounted: boolean;
  readonly model: SelectorModel;
  readonly loadStatus: LoadStatus;
  readonly closed: ClosedKind;
  /** The battletag (or the email's local part) of a signed-in account; empty until `/v1/me` answered. */
  readonly accountName: string;
  /** The readable half of the session cookie is present: someone is probably signed in. */
  readonly sessionHint: boolean;
  /** The session's first verified guild once `/v1/me` answered; null signed out or without one. */
  readonly sessionGuild: SessionGuild | null;
  /** Re-reads the pointer, the chosen keys and the clock (the list is about to be shown). */
  refresh(): void;
  /** Asks `/v1/me` again (the list's "Try again"). */
  reload(): void;
}

function loadStatusOf(status: 'idle' | 'loading' | 'ready' | 'failed'): LoadStatus {
  if (status === 'ready') return 'ready';
  return status === 'failed' ? 'failed' : 'loading';
}

export function createSelectorState(): SelectorState {
  const session = createQueryState<Me | null>(meKey(API_BASE_URL), () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: ME_TTL_MS,
    version: ME_QUERY_VERSION,
  });

  let mounted = $state(false);
  let pointer = $state<CurrentCharacter | null>(null);
  let chosenKeys = $state<string[]>([]);
  let nowMs = $state(Date.now());
  let sessionHint = $state(false);

  function refresh(): void {
    pointer = readCurrent();
    chosenKeys = readChosenKeys();
    nowMs = Date.now();
  }

  $effect(() => {
    refresh();
    sessionHint = sessionHinted();
    mounted = true;
    return onCurrentCharacterChange(refresh);
  });

  // A character's age is measured against the moment its data arrived, not the moment this
  // island was created.
  $effect(() => {
    void session.data;
    nowMs = Date.now();
  });

  const loadStatus = $derived(loadStatusOf(session.status));
  // On a guild page the header shows the account's character in that guild (a view only).
  const shownPointer = $derived(
    mounted && loadStatus === 'ready'
      ? guildPagePointer(session.data, window.location.pathname, pointer, performance.timeOrigin)
      : pointer,
  );
  const model = $derived(
    buildSelectorModel({
      me: loadStatus === 'ready' ? session.data : undefined,
      pointer: shownPointer,
      chosenKeys,
      nowMs,
    }),
  );

  const accountName = $derived(loadStatus === 'ready' && session.data ? accountNameOf(session.data) : '');

  const sessionGuild = $derived(loadStatus === 'ready' ? memberGuildOf(session.data) : null);

  // Until the browser has been read, and while a signed-in visitor's character is still
  // unknown, the closed control stays the pre-paint link (its skeleton variant); a visitor
  // with nothing to show gets the sign-in slot at once.
  const closed = $derived(closedKindOf({ mounted, model, loadStatus, sessionHint }));

  return {
    get closed() {
      return closed;
    },
    get accountName() {
      return accountName;
    },
    get mounted() {
      return mounted;
    },
    get model() {
      return model;
    },
    get loadStatus() {
      return loadStatus;
    },
    get sessionHint() {
      return sessionHint;
    },
    get sessionGuild() {
      return sessionGuild;
    },
    refresh,
    reload: () => session.refresh(),
  };
}
