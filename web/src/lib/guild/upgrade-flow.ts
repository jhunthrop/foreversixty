// web/src/lib/guild/upgrade-flow.ts
// The Move to Raid guild flow of guild settings #billing as a small controller: preview, commit, and
// what each refusal does (billing-errors.ts). Plain TypeScript with injected calls and timers, so the
// stale-response guard, the single recheck timer and the unmount cleanup are tested without a page.
import type { GuildBillingSettings, UpgradePreview } from '../billing/contract';
import { BILLING_API_FAILED } from '../billing/api';
import { guildBillingCopy as copy } from '../billing/guild-billing-copy';
import { seatsCopy } from '../billing/seats-copy';
import { upgradeOutcome } from './billing-errors';

/** Stripe may still be settling an unconfirmed change: the plan is read once more after this. */
export const UNCONFIRMED_RECHECK_MS = 5000;

/** `unconfirmed-stale`: Stripe is confirming and the plan on screen could not be re-read. */
export type UpgradeNotice = 'none' | 'moved' | 'unconfirmed' | 'unconfirmed-stale';
export type FocusTarget = 'plan' | 'failure';

/** A focus request is handled once: only a counter newer than the last handled one moves focus. */
export function focusIsPending(request: number, lastHandled: number): boolean {
  return request > lastHandled;
}

export interface UpgradeFlowState {
  confirming: boolean;
  preview: UpgradePreview | null;
  committing: boolean;
  error: string;
  notice: UpgradeNotice;
  /** Bumps each time the confirm closes onto a reloaded plan, so the view can move focus. */
  focusRequest: number;
  /** Where that focus goes: the plan line, or the reload failure when the plan could not be read. */
  focusTarget: FocusTarget;
}

export interface UpgradeFlowDeps {
  previewUpgrade: () => Promise<UpgradePreview>;
  upgrade: (prorationDate: number) => Promise<GuildBillingSettings>;
  /** Drops the cached guild pages. */
  invalidate: () => void;
  /** Reads the billing view again; answers whether it succeeded. */
  reload: () => Promise<boolean>;
  /** The view took the plan the upgrade answered with. */
  adopt: (view: GuildBillingSettings) => void;
  onState: (state: UpgradeFlowState) => void;
  setTimer?: (run: () => void, ms: number) => unknown;
  clearTimer?: (handle: unknown) => void;
}

export const INITIAL_UPGRADE_STATE: UpgradeFlowState = {
  confirming: false,
  preview: null,
  committing: false,
  error: '',
  notice: 'none',
  focusRequest: 0,
  focusTarget: 'plan',
};

export interface UpgradeFlow {
  open: () => Promise<void>;
  close: () => void;
  commit: () => Promise<void>;
  /** Cancels the recheck and ignores every answer still in flight. */
  dispose: () => void;
}

export function createUpgradeFlow(deps: UpgradeFlowDeps): UpgradeFlow {
  const setTimer = deps.setTimer ?? ((run, ms) => window.setTimeout(run, ms));
  const clearTimer = deps.clearTimer ?? ((handle) => window.clearTimeout(handle as number));
  let state = INITIAL_UPGRADE_STATE;
  let generation = 0;
  let recheckTimer: unknown = null;

  const set = (patch: Partial<UpgradeFlowState>): void => {
    state = { ...state, ...patch };
    deps.onState(state);
  };

  /** Every entry point starts a new generation: callbacks of an older one skip. */
  function begin(): number {
    generation += 1;
    if (recheckTimer !== null) clearTimer(recheckTimer);
    recheckTimer = null;
    return generation;
  }

  async function reloadAndRefocus(patch: Partial<UpgradeFlowState>, token: number): Promise<void> {
    deps.invalidate();
    const ok = await deps.reload();
    if (token !== generation) return;
    const notice = patch.notice === 'unconfirmed' && !ok ? 'unconfirmed-stale' : patch.notice;
    set({
      ...patch,
      ...(notice === undefined ? {} : { notice }),
      confirming: false,
      focusRequest: state.focusRequest + 1,
      focusTarget: ok ? 'plan' : 'failure',
    });
  }

  /** An answer that arrives after a newer action is never dropped: the plan is read again. */
  async function reloadStale(): Promise<void> {
    deps.invalidate();
    await deps.reload();
  }

  function armRecheck(token: number): void {
    recheckTimer = setTimer(() => {
      recheckTimer = null;
      if (token !== generation) return;
      void deps.reload().then((ok) => {
        if (ok && token === generation) set({ notice: 'none' });
      });
    }, UNCONFIRMED_RECHECK_MS);
  }

  async function open(): Promise<void> {
    const token = begin();
    set({ confirming: true, preview: null, error: '', notice: 'none', committing: false });
    try {
      const fresh = await deps.previewUpgrade();
      if (token === generation) set({ preview: fresh });
    } catch (thrown) {
      if (token !== generation) return;
      if (upgradeOutcome(thrown) === 'reload-quiet') await reloadAndRefocus({ error: '' }, token);
      else set({ error: copy.upgrade.failed });
    }
  }

  function close(): void {
    begin();
    set({ confirming: false, error: '', committing: false });
  }

  async function failCommit(thrown: unknown, token: number): Promise<void> {
    const outcome = upgradeOutcome(thrown);
    if (outcome === 'reopen') {
      await open();
      if (state.preview !== null) set({ error: copy.upgrade.previewExpired });
    } else if (outcome === 'declined') {
      set({ error: copy.upgrade.failed });
    } else if (outcome === 'generic') {
      set({ error: BILLING_API_FAILED });
    } else if (outcome === 'in-progress') {
      set({ error: seatsCopy.whoPays.inProgress });
    } else if (outcome === 'reload-quiet') {
      await reloadAndRefocus({ error: '' }, token);
    } else {
      await reloadAndRefocus({ notice: 'unconfirmed', error: '' }, token);
      if (token === generation) armRecheck(token);
    }
  }

  async function commit(): Promise<void> {
    const preview = state.preview;
    if (preview === null) return;
    const token = begin();
    set({ committing: true, error: '', notice: 'none' });
    try {
      const view = await deps.upgrade(preview.proration_date);
      if (token !== generation) return await reloadStale();
      deps.adopt(view);
      set({
        confirming: false,
        notice: 'moved',
        focusRequest: state.focusRequest + 1,
        focusTarget: 'plan',
      });
    } catch (thrown) {
      if (token === generation) await failCommit(thrown, token);
      else if (upgradeOutcome(thrown) !== 'declined') await reloadStale();
    } finally {
      if (token === generation) set({ committing: false });
    }
  }

  function dispose(): void {
    begin();
  }

  return { open, close, commit, dispose };
}
