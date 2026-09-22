// web/src/test-support/billing-fixtures.ts
// e2e stubs for the billing routes, following account/api.ts's requestEnvelope envelope
// shape and the sim lane's existing sim-api.ts pattern of one small helper per route rather
// than each spec hand-rolling page.route calls.
import type { Page } from '@playwright/test';
import type { EntitlementsView } from '../lib/account/api';

export function mockEntitlements(overrides: Partial<EntitlementsView> = {}): EntitlementsView {
  return {
    server_sims: false,
    retention: false,
    multi_compare: false,
    history: false,
    notifications: false,
    officer_views: false,
    roster_check: false,
    supporter_mark: false,
    billing: null,
    ...overrides,
  };
}

function envelope(data: unknown, status: number): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
  };
}

export async function stubCheckout(page: Page, response: { status: number; data?: unknown }): Promise<void> {
  await page.route('**/v1/billing/checkout', (route) =>
    route.fulfill(envelope(response.data ?? null, response.status)),
  );
}

export async function stubPortal(page: Page, response: { status: number; data?: unknown }): Promise<void> {
  await page.route('**/v1/billing/portal', (route) =>
    route.fulfill(envelope(response.data ?? null, response.status)),
  );
}
