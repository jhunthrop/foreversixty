// @vitest-environment jsdom
// web/src/lib/account/discord-return.test.ts
// Reading the outcome Discord's callback left in the address bar: once, cleaned, and a
// connected outcome alone re-reads /v1/me.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import * as accountApi from './api';
import { consumeReturn } from './discord-return.svelte';
import { notificationsCopy } from './notifications-copy';

let fresh: ReturnType<typeof vi.spyOn>;

beforeEach(() => {
  fresh = vi.spyOn(accountApi, 'fetchMeFresh').mockResolvedValue(null);
});
afterEach(() => {
  vi.restoreAllMocks();
  window.history.replaceState({}, '', '/');
});

describe('consumeReturn', () => {
  it('connected: the sentence, focus on the test button, one fresh read, the flag gone', () => {
    window.history.replaceState({}, '', '/account?discord=connected#discord');
    const outcome = consumeReturn();
    expect(outcome?.status).toEqual({ ok: true, message: notificationsCopy.returnConnected });
    expect(outcome?.focus).toContain('discord-test');
    expect(fresh).toHaveBeenCalledTimes(1);
    expect(window.location.pathname + window.location.search + window.location.hash).toBe('/account#discord');
    expect(consumeReturn()).toBeNull();
    expect(fresh).toHaveBeenCalledTimes(1);
  });

  it.each(['denied', 'failed'])('%s: no fresh read', (flag) => {
    window.history.replaceState({}, '', `/account?discord=${flag}#discord`);
    expect(consumeReturn()).not.toBeNull();
    expect(fresh).not.toHaveBeenCalled();
  });

  it('no flag: nothing', () => {
    window.history.replaceState({}, '', '/account#discord');
    expect(consumeReturn()).toBeNull();
  });
});
