import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { BILLING_API_FAILED, BillingApiError } from '../billing/api';
import { guildBillingCopy } from '../billing/guild-billing-copy';
import { seatsCopy } from '../billing/seats-copy';
import { GUILD_TIER_28_READ, UPGRADE_PREVIEW } from '../../test-support/billing-fixtures';
import {
  INITIAL_UPGRADE_STATE,
  UNCONFIRMED_RECHECK_MS,
  focusIsPending,
  createUpgradeFlow,
  type UpgradeFlowDeps,
  type UpgradeFlowState,
} from './upgrade-flow';

const refusal = (status: number, code: string | null): BillingApiError =>
  new BillingApiError('x', status, code);

interface Harness {
  flow: ReturnType<typeof createUpgradeFlow>;
  state: () => UpgradeFlowState;
  deps: { [K in 'previewUpgrade' | 'upgrade' | 'invalidate' | 'reload' | 'adopt']: ReturnType<typeof vi.fn> };
}

function harness(overrides: Partial<UpgradeFlowDeps> = {}): Harness {
  let state = INITIAL_UPGRADE_STATE;
  const deps = {
    previewUpgrade: vi.fn(async () => UPGRADE_PREVIEW),
    upgrade: vi.fn(async () => GUILD_TIER_28_READ),
    invalidate: vi.fn(),
    reload: vi.fn(async () => true),
    adopt: vi.fn(),
  };
  const flow = createUpgradeFlow({ ...deps, ...overrides, onState: (next) => (state = next) });
  return { flow, state: () => state, deps };
}

beforeEach(() => vi.useFakeTimers());
afterEach(() => vi.useRealTimers());

// The flow reads window.setTimeout by default; fake timers patch the global the tests run in.
vi.stubGlobal('window', globalThis);

describe('preview', () => {
  it('loads the preview into the open confirm', async () => {
    const h = harness();
    await h.flow.open();
    expect(h.state()).toMatchObject({ confirming: true, preview: UPGRADE_PREVIEW });
  });

  it('ignores a preview that lands after the confirm was closed and reopened', async () => {
    let first!: (value: typeof UPGRADE_PREVIEW) => void;
    const previews = [
      new Promise<typeof UPGRADE_PREVIEW>((resolve) => (first = resolve)),
      Promise.resolve({ ...UPGRADE_PREVIEW, amount_due_cents: 999 }),
    ];
    const h = harness({ previewUpgrade: async () => previews.shift()! });
    const opening = h.flow.open();
    h.flow.close();
    await h.flow.open();
    first(UPGRADE_PREVIEW);
    await opening;
    expect(h.state().preview?.amount_due_cents).toBe(999);
  });

  it('not_upgradable on the preview closes, invalidates and reloads with no line', async () => {
    const h = harness({
      previewUpgrade: async () => {
        throw refusal(409, 'not_upgradable');
      },
    });
    await h.flow.open();
    expect(h.deps.invalidate).toHaveBeenCalled();
    expect(h.deps.reload).toHaveBeenCalled();
    expect(h.state()).toMatchObject({ confirming: false, error: '', notice: 'none' });
  });
});

describe('commit', () => {
  it('adopts the answered plan and says it moved', async () => {
    const h = harness();
    await h.flow.open();
    await h.flow.commit();
    expect(h.deps.upgrade).toHaveBeenCalledWith(UPGRADE_PREVIEW.proration_date);
    expect(h.deps.adopt).toHaveBeenCalledWith(GUILD_TIER_28_READ);
    expect(h.state()).toMatchObject({ confirming: false, notice: 'moved', committing: false });
    expect(h.state().focusRequest).toBeGreaterThan(0);
  });

  it('a declined card keeps the confirm and says nothing was charged', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(409, 'upgrade_failed');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state()).toMatchObject({ confirming: true, error: guildBillingCopy.upgrade.failed });
  });

  it('any other 4xx reads as the generic failure, not as nothing charged', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(403, 'forbidden');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state().error).toBe(BILLING_API_FAILED);
  });

  it('an open transfer checkout keeps the confirm and says a change is in progress', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(409, 'transfer_in_progress');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state()).toMatchObject({ confirming: true, error: seatsCopy.whoPays.inProgress });
  });

  it('an open transfer offer is named in the in-progress line', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(409, 'transfer_in_progress');
      },
      openTransferTo: () => 'Mirelle',
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state().error).toBe(seatsCopy.whoPays.inProgressOffer('Mirelle'));
  });

  it('not_upgradable closes, invalidates and reloads with no line', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(409, 'not_upgradable');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.deps.invalidate).toHaveBeenCalled();
    expect(h.state()).toMatchObject({ confirming: false, error: '', notice: 'none' });
  });

  it('preview_expired re-opens the preview with the prices-changed line', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(409, 'preview_expired');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.deps.previewUpgrade).toHaveBeenCalledTimes(2);
    expect(h.state()).toMatchObject({ confirming: true, error: guildBillingCopy.upgrade.previewExpired });
  });
});

describe('unconfirmed', () => {
  const failing = (): Partial<UpgradeFlowDeps> => ({
    upgrade: async () => {
      throw refusal(502, 'upgrade_unconfirmed');
    },
  });

  it('closes, reloads, shows the status and re-reads once after 5 s, then clears the status', async () => {
    const h = harness(failing());
    await h.flow.open();
    await h.flow.commit();
    expect(h.state()).toMatchObject({ confirming: false, notice: 'unconfirmed' });
    expect(h.deps.reload).toHaveBeenCalledTimes(1);
    await vi.advanceTimersByTimeAsync(UNCONFIRMED_RECHECK_MS);
    expect(h.deps.reload).toHaveBeenCalledTimes(2);
    expect(h.state().notice).toBe('none');
  });

  it('keeps the status when the delayed re-read fails', async () => {
    const h = harness(failing());
    await h.flow.open();
    await h.flow.commit();
    h.deps.reload.mockResolvedValue(false);
    await vi.advanceTimersByTimeAsync(UNCONFIRMED_RECHECK_MS);
    expect(h.state().notice).toBe('unconfirmed');
  });

  it('a network failure (status 0) counts as unconfirmed; a plain throw is the generic failure', async () => {
    const net = harness({
      upgrade: async () => {
        throw refusal(0, null);
      },
    });
    await net.flow.open();
    await net.flow.commit();
    expect(net.state().notice).toBe('unconfirmed');
    const odd = harness({
      upgrade: async () => {
        throw new TypeError('bug');
      },
    });
    await odd.flow.open();
    await odd.flow.commit();
    expect(odd.state()).toMatchObject({ notice: 'none', error: BILLING_API_FAILED });
  });

  it('opening again cancels the pending re-read and clears the status', async () => {
    const h = harness(failing());
    await h.flow.open();
    await h.flow.commit();
    await h.flow.open();
    expect(h.state().notice).toBe('none');
    await vi.advanceTimersByTimeAsync(UNCONFIRMED_RECHECK_MS);
    expect(h.deps.reload).toHaveBeenCalledTimes(1);
  });

  it('dispose (unmount) cancels the pending re-read', async () => {
    const h = harness(failing());
    await h.flow.open();
    await h.flow.commit();
    h.flow.dispose();
    await vi.advanceTimersByTimeAsync(UNCONFIRMED_RECHECK_MS);
    expect(h.deps.reload).toHaveBeenCalledTimes(1);
  });

  it('a commit after a close clears the earlier status and its timer', async () => {
    const h = harness(failing());
    await h.flow.open();
    await h.flow.commit();
    h.flow.close();
    await vi.advanceTimersByTimeAsync(UNCONFIRMED_RECHECK_MS);
    expect(h.deps.reload).toHaveBeenCalledTimes(1);
  });
});

describe('fix round 2', () => {
  it('a commit that lands after a newer action still invalidates and reloads, and is not adopted', async () => {
    let land!: (view: typeof GUILD_TIER_28_READ) => void;
    const h = harness({ upgrade: () => new Promise((resolve) => (land = resolve)) });
    await h.flow.open();
    const committing = h.flow.commit();
    await h.flow.open();
    land(GUILD_TIER_28_READ);
    await committing;
    expect(h.deps.adopt).not.toHaveBeenCalled();
    expect(h.deps.invalidate).toHaveBeenCalled();
    expect(h.deps.reload).toHaveBeenCalledTimes(1);
  });

  it('a stale failure that is not a clean decline also reloads', async () => {
    let fail!: (error: unknown) => void;
    const h = harness({ upgrade: () => new Promise((_, reject) => (fail = reject)) });
    await h.flow.open();
    const committing = h.flow.commit();
    h.flow.close();
    fail(refusal(502, 'upgrade_unconfirmed'));
    await committing;
    expect(h.deps.reload).toHaveBeenCalledTimes(1);
  });

  it('focus is pending once per request', () => {
    expect(focusIsPending(1, 0)).toBe(true);
    expect(focusIsPending(1, 1)).toBe(false);
    expect(focusIsPending(0, 0)).toBe(false);
  });

  it('a failed reload sends focus to the failure and uses the stale unconfirmed line', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(502, 'upgrade_unconfirmed');
      },
      reload: vi.fn(async () => false),
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state()).toMatchObject({ notice: 'unconfirmed-stale', focusTarget: 'failure' });
    expect(h.state().focusRequest).toBeGreaterThan(0);
  });

  it('a good reload sends focus to the plan', async () => {
    const h = harness({
      upgrade: async () => {
        throw refusal(409, 'not_upgradable');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state().focusTarget).toBe('plan');
  });

  it('preview_expired shows the prices-changed line only if the new preview loaded', async () => {
    let previews = 0;
    const h = harness({
      previewUpgrade: async () => {
        previews += 1;
        if (previews === 1) return UPGRADE_PREVIEW;
        throw refusal(500, null);
      },
      upgrade: async () => {
        throw refusal(409, 'preview_expired');
      },
    });
    await h.flow.open();
    await h.flow.commit();
    expect(h.state().preview).toBeNull();
    expect(h.state().error).toBe(guildBillingCopy.upgrade.failed);
  });

  it('states the stale-unconfirmed addition', () => {
    expect(guildBillingCopy.upgrade.unconfirmedStale).toBe(
      'Stripe is still confirming the move. Reload in a minute to see the plan.',
    );
  });
});
