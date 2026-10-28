// web/src/lib/billing/card-checkout.test.ts
import { describe, expect, it, vi } from 'vitest';
import { BillingApiError } from './api';
import {
  guildArrival,
  isReadyFill,
  payAction,
  readyCount,
  readyLabel,
  runCardCheckout,
  runCardPortal,
  upgradeHref,
} from './card-checkout';
import { checkoutGuildCopy as copy } from './checkout-guild-copy';
import type { StepGuild } from './subscribe-step';

const REQUEST = { plan: 'guild', interval: 'yearly', guildId: 7 } as const;
const failing = (status: number) => vi.fn().mockRejectedValue(new BillingApiError('api words', status));

describe('readyLabel', () => {
  const price = '$120 a year';
  it('names the trial when unused, the price when used, again when lapsed', () => {
    expect(readyLabel({ subscribe: { step: 'ready', trial_used: false } }, price)).toBe(
      'Start the 7-day trial',
    );
    expect(readyLabel({ subscribe: { step: 'ready', trial_used: true } }, price)).toBe(
      'Subscribe · $120 a year',
    );
    expect(
      readyLabel({ subscribe: { step: 'ready', trial_used: true }, covered: { state: 'lapsed' } }, price),
    ).toBe('Subscribe again · $120 a year');
  });
});

describe('the one fill', () => {
  it('fills only a lone ready card', () => {
    expect(isReadyFill(1)).toBe(true);
    expect(isReadyFill(2)).toBe(false);
    const guild = (id: number, step: 'ready' | 'on_plan'): StepGuild =>
      ({ id, subscribe: { step } }) as StepGuild;
    expect(readyCount([guild(1, 'ready'), guild(2, 'on_plan'), guild(3, 'ready')])).toBe(2);
  });
});

describe('payAction', () => {
  const covered = (tier: 'guild' | 'raid_guild', you_pay: boolean) => ({ covered: { tier, you_pay } });
  it('manage for a payer at or above the chosen tier, move from Guild to a chosen Raid guild', () => {
    expect(payAction(covered('guild', true), 'guild')).toBe('manage');
    expect(payAction(covered('raid_guild', true), 'guild')).toBe('manage');
    expect(payAction(covered('raid_guild', true), 'raid_guild')).toBe('manage');
    expect(payAction(covered('guild', true), 'raid_guild')).toBe('move');
  });
  it('nothing for a viewer who does not pay', () => {
    expect(payAction(covered('guild', false), 'raid_guild')).toBeNull();
    expect(payAction({}, 'guild')).toBeNull();
  });
  it('the upgrade link opens the settings billing confirm', () => {
    expect(upgradeHref('/guild/us/forever/x')).toBe('/guild/us/forever/x/settings?upgrade=1#billing');
  });
});

describe('runCardCheckout', () => {
  it('redirects to the returned url with the request unchanged', async () => {
    const startCheckout = vi.fn().mockResolvedValue({ checkout_url: 'https://stripe.test/c' });
    expect(await runCardCheckout({ startCheckout }, REQUEST, 'Olympus')).toEqual({
      kind: 'redirect',
      url: 'https://stripe.test/c',
    });
    expect(startCheckout).toHaveBeenCalledWith(REQUEST);
  });
  it('503 says purchases are not open, no reread', async () => {
    expect(await runCardCheckout({ startCheckout: failing(503) }, REQUEST, 'Olympus')).toEqual({
      kind: 'refused',
      line: copy.checkoutErrors.notOpen,
      rereads: false,
      retry: false,
    });
  });
  it('409 names the guild and rereads', async () => {
    expect(await runCardCheckout({ startCheckout: failing(409) }, REQUEST, 'Olympus')).toEqual({
      kind: 'refused',
      line: 'Olympus is already on a plan.',
      rereads: true,
      retry: false,
    });
  });
  it('403 rereads and never shows the API sentence', async () => {
    const outcome = await runCardCheckout({ startCheckout: failing(403) }, REQUEST, 'Olympus');
    expect(outcome).toEqual({ kind: 'refused', line: null, rereads: true, retry: false });
  });
  it('anything else is the generic line with Try again', async () => {
    for (const status of [500, 0]) {
      expect(await runCardCheckout({ startCheckout: failing(status) }, REQUEST, 'O')).toEqual({
        kind: 'refused',
        line: copy.checkoutErrors.generic,
        rereads: false,
        retry: true,
      });
    }
  });
});

describe('runCardPortal', () => {
  it('posts the guild id and redirects', async () => {
    const openPortal = vi.fn().mockResolvedValue({ portal_url: 'https://stripe.test/p' });
    expect(await runCardPortal({ openPortal }, 7)).toEqual({
      kind: 'redirect',
      url: 'https://stripe.test/p',
    });
    expect(openPortal).toHaveBeenCalledWith(7);
  });
  it('any failure says Billing did not open', async () => {
    expect(await runCardPortal({ openPortal: failing(500) }, 7)).toEqual({
      kind: 'refused',
      line: 'Billing did not open. Try again.',
      rereads: false,
      retry: false,
    });
  });
});

describe('guildArrival', () => {
  const guilds = [
    { id: 1, subscribe: { step: 'ready' } },
    { id: 2, subscribe: { step: 'ask_leader' } },
  ] as StepGuild[];
  it('launches a ready guild, floats another, ignores unknown or absent ids', () => {
    expect(guildArrival(guilds, 1)).toEqual({ launchId: 1 });
    expect(guildArrival(guilds, 2)).toEqual({ firstId: 2 });
    expect(guildArrival(guilds, 99)).toEqual({});
    expect(guildArrival(guilds, undefined)).toEqual({});
  });
});
