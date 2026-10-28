// web/tests/e2e/support/subscribe-stubs.ts
// What the subscribe-and-join specs share: `/v1/me` answered from a cell (so a claim, a join or a
// save can change what the next read sees), the signed-out answer, the account page's devices
// read, the `fs_csrf` hint that says a session exists, and the 44 px tap rule.
import { expect, type Locator, type Page } from '@playwright/test';
import type { Me } from '../../../src/lib/account/me-types';
import { envelope } from './guild-billing';
import { MIN_TARGET_PX } from './short-controls';

/** The signed-out answer every `/v1/me` stub shares. */
export const SIGNED_OUT = {
  status: 401,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
};

/** The `/v1/me` answers a spec can steer: swap the account (null is signed out), hold, delay or fail the read. */
export interface MeCell {
  set: (me: Me | null) => void;
  /** Holds every read until the returned release runs; the re-read after a claim is the one to hold. */
  hold: () => () => void;
  failReads: (failing: boolean) => void;
  reads: () => number;
}

/** What a stub answers from: the account now, or a function read on every request. */
type MeSource = Me | null | (() => Me | null);

/** Stubs `/v1/me` from a cell; `delayMs` slows every read. A read function stays valid for older call sites. */
export async function routeMeCell(page: Page, initial: MeSource, delayMs = 0): Promise<MeCell> {
  let source = initial;
  let gate: Promise<void> | null = null;
  let failing = false;
  let reads = 0;
  await page.route('**/v1/me', async (route) => {
    reads += 1;
    if (delayMs > 0) await new Promise((resolve) => setTimeout(resolve, delayMs));
    if (gate !== null) await gate;
    if (failing) return route.fulfill(envelope({ code: 'internal', message: 'api sentence' }, 500));
    const me = typeof source === 'function' ? source() : source;
    return route.fulfill(me === null ? SIGNED_OUT : envelope(me));
  });
  return {
    set: (next) => {
      source = next;
    },
    hold: () => {
      let release = (): void => undefined;
      gate = new Promise<void>((resolve) => {
        release = resolve;
      });
      return () => {
        release();
        gate = null;
      };
    },
    failReads: (next) => {
      failing = next;
    },
    reads: () => reads,
  };
}

export async function routeDevices(page: Page): Promise<void> {
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

/** The cookie hint the site reads to know a session may exist, so it asks `/v1/me` at all. */
export async function plantSessionHint(page: Page): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'test', url: 'http://localhost' }]);
}

export async function expectTapHeight(control: Locator): Promise<void> {
  const box = await control.boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(MIN_TARGET_PX);
}
