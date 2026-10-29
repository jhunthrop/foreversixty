// @vitest-environment jsdom
import { afterEach, describe, expect, it } from 'vitest';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import { discordBlock } from '../../test-support/notifications-api';
import type { DiscordNotifications, Me } from './api';
import {
  ACCOUNT_RESERVE_KEY,
  ACCOUNT_RESERVE_MAX_PX,
  DISCORD_FLOORS,
  clearAccountReserve,
  discordHintState,
  isTransientAccountState,
  readAccountReserve,
  recordAccountReserve,
  writeAccountReserve,
  type HintState,
} from './layout';
import { blockState, type DiscordBlockKey } from './notifications';

afterEach(() => localStorage.clear());

function me(entitled: boolean, discord?: DiscordNotifications | null): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ notifications: entitled }),
    ...(discord === undefined ? {} : { notifications: { discord } }),
  };
}

function meConnect(connect: boolean, discord: DiscordNotifications | null): Me {
  return { ...me(true, discord), notifications: { connect, discord } };
}

const failed = discordBlock({ disabled_reason: 'failed five times', failures: 5 });
const planLapsed = discordBlock({ disabled_reason: 'plan lapsed' });
const noEntitlements: Me = { ...me(true, null), entitlements: undefined };

describe('discordHintState (spec 4.1 precedence)', () => {
  it.each<[string, Me | null, HintState]>([
    ['no cached /v1/me', null, 'none'],
    ['no notifications key, entitled', me(true), 'none'],
    ['no notifications key, free', me(false), 'none'],
    ['lapsed: saved, not entitled', me(false, discordBlock()), 'lapsed'],
    ['lapsed wins over a failed-five-times reason', me(false, failed), 'lapsed'],
    ['lapsed wins over a plan-lapsed reason', me(false, planLapsed), 'lapsed'],
    ['off: entitled, failed five times', me(true, failed), 'off'],
    ['off: entitled, any other reason', me(true, planLapsed), 'off'],
    ['saved', me(true, discordBlock()), 'saved'],
    ['form: entitled, nothing saved', me(true, null), 'form'],
    ['free: not entitled, nothing saved', me(false, null), 'free'],
    ['connect: entitled, nothing saved, connect on', meConnect(true, null), 'connect'],
    ['connect false reads form', meConnect(false, null), 'form'],
    ['connect on does not change saved', meConnect(true, discordBlock()), 'saved'],
    ['connect on does not change free', { ...meConnect(true, null), entitlements: undefined }, 'free'],
    ['free: no entitlements at all', noEntitlements, 'free'],
  ])('%s -> %s', (_name, account, expected) => {
    expect(discordHintState(account)).toBe(expected);
  });

  it('agrees with blockState on every fixture the block is tested with', () => {
    const toHint: Record<DiscordBlockKey, HintState> = {
      off: 'none',
      free: 'free',
      empty: 'form',
      saved: 'saved',
      disabled: 'off',
      lapsed: 'lapsed',
    };
    const fixtures = [
      me(false),
      me(true),
      me(false, null),
      me(true, null),
      me(true, discordBlock()),
      me(true, failed),
      me(false, discordBlock()),
      me(false, failed),
      noEntitlements,
    ];
    for (const account of fixtures) {
      expect(discordHintState(account)).toBe(toHint[blockState(account).state]);
    }
  });

  it('has a floor for every hint state, and none for none', () => {
    expect(Object.keys(DISCORD_FLOORS).sort()).toEqual([
      'connect',
      'form',
      'free',
      'lapsed',
      'none',
      'off',
      'saved',
    ]);
    expect(DISCORD_FLOORS.none).toEqual({ lg: 0, stack: 0 });
    expect(DISCORD_FLOORS.form).toEqual({ lg: 735, stack: 765 });
    expect(DISCORD_FLOORS.connect).toEqual({ lg: 340, stack: 360 });
  });
});

describe('isTransientAccountState', () => {
  afterEach(() => {
    document.body.innerHTML = '';
  });

  it.each(['discord-cancel', 'discord-remove-confirm'])('is true while %s is on the page', (testid) => {
    document.body.innerHTML = `<button data-testid="${testid}"></button>`;
    expect(isTransientAccountState(document)).toBe(true);
  });

  it('is false on a settled page', () => {
    document.body.innerHTML = '<button data-testid="discord-save"></button>';
    expect(isTransientAccountState(document)).toBe(false);
  });

  it('keeps recordAccountReserve from storing a passing state', () => {
    const el = (): Element => {
      const node = document.createElement('div');
      node.getBoundingClientRect = () => ({ height: 700 }) as DOMRect;
      return node;
    };
    document.body.innerHTML = '<button data-testid="discord-cancel"></button>';
    recordAccountReserve(el(), el());
    expect(localStorage.getItem(ACCOUNT_RESERVE_KEY)).toBeNull();
  });
});

describe('account reserve', () => {
  function stored(): unknown {
    return JSON.parse(localStorage.getItem(ACCOUNT_RESERVE_KEY) ?? 'null');
  }

  it('is null with nothing stored', () => {
    expect(readAccountReserve()).toBeNull();
  });

  it('keeps the other bucket when one is written', () => {
    writeAccountReserve('lg', 812.4, 1010);
    expect(readAccountReserve()).toBeNull(); // one bucket is not a full reserve yet
    writeAccountReserve('stack', 640, 1500);
    expect(readAccountReserve()).toEqual({
      lg: { left: 812, rail: 1010 },
      stack: { left: 640, rail: 1500 },
    });
    writeAccountReserve('lg', 900, 1100);
    expect(readAccountReserve()).toEqual({
      lg: { left: 900, rail: 1100 },
      stack: { left: 640, rail: 1500 },
    });
  });

  it('stores the documented shape', () => {
    writeAccountReserve('stack', 640, 1500);
    expect(stored()).toEqual({ v: 1, stack: { left: 640, rail: 1500 } });
  });

  it.each([0, -3, Number.NaN, Number.POSITIVE_INFINITY, ACCOUNT_RESERVE_MAX_PX + 1])(
    'skips a write of %s',
    (bad) => {
      writeAccountReserve('lg', 800, 1000);
      writeAccountReserve('lg', bad, 1000);
      writeAccountReserve('lg', 800, bad);
      expect(stored()).toEqual({ v: 1, lg: { left: 800, rail: 1000 } });
    },
  );

  it('accepts the maximum', () => {
    writeAccountReserve('lg', ACCOUNT_RESERVE_MAX_PX, ACCOUNT_RESERVE_MAX_PX);
    writeAccountReserve('stack', 1, 1);
    expect(readAccountReserve()?.lg.left).toBe(ACCOUNT_RESERVE_MAX_PX);
  });

  it.each([
    ['not JSON', 'nope'],
    ['wrong version', JSON.stringify({ v: 2, lg: { left: 1, rail: 1 }, stack: { left: 1, rail: 1 } })],
    ['zero height', JSON.stringify({ v: 1, lg: { left: 0, rail: 1 }, stack: { left: 1, rail: 1 } })],
    ['oversized', JSON.stringify({ v: 1, lg: { left: 1, rail: 4001 }, stack: { left: 1, rail: 1 } })],
    ['string height', JSON.stringify({ v: 1, lg: { left: '1', rail: 1 }, stack: { left: 1, rail: 1 } })],
    ['missing bucket', JSON.stringify({ v: 1, lg: { left: 1, rail: 1 } })],
  ])('reads %s as null', (_name, raw) => {
    localStorage.setItem(ACCOUNT_RESERVE_KEY, raw);
    expect(readAccountReserve()).toBeNull();
  });

  it('clears the key', () => {
    writeAccountReserve('lg', 800, 1000);
    clearAccountReserve();
    expect(localStorage.getItem(ACCOUNT_RESERVE_KEY)).toBeNull();
  });
});
