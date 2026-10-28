// web/tests/e2e/support/subscribe-stubs.ts
// What the subscribe-and-join specs share: `/v1/me` answered from a cell (so a claim, a join or a
// save can change what the next read sees), the signed-out answer, the account page's devices
// read, the `fs_csrf` hint that says a session exists, and the 44 px tap rule.
import { expect, type Locator, type Page } from '@playwright/test';
import type { Me } from '../../../src/lib/account/me-types';
import { envelope } from './guild-billing';

export const MIN_TAP_PX = 44;

const SIGNED_OUT = {
  status: 401,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
};

/** Answers `/v1/me` with whatever `read` returns now; null is a signed-out visitor. */
export async function routeMeCell(page: Page, read: () => Me | null): Promise<void> {
  await page.route('**/v1/me', (route) => {
    const me = read();
    return route.fulfill(me === null ? SIGNED_OUT : envelope(me));
  });
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
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(MIN_TAP_PX);
}
