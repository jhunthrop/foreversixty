// web/tests/e2e/guild-billing-seats.spec.ts
// Launch pricing spec section 9, guild settings #billing: case 12 (seats), 13 (transfer) and 15
// (a read-only officer). API stubbed.
import { expect, test } from './support/test';
import type { GuildBillingSettings } from '../../src/lib/billing/contract';
import {
  READ_ONLY_OFFICER_READ,
  TRANSFER_DECLINED,
  TRANSFER_OFFERED,
  TRANSFER_RECEIVED,
  emptySeat,
  guildBillingRead,
  heldSeat,
  pendingSeat,
  planView,
  returnedSeat,
} from '../../src/test-support/billing-fixtures';
import { GUILD_ID, envelope, openLive, stubCandidates } from './support/guild-billing';
import { choose, fieldControl } from './support/select';

const PENDING_OFFEREE = { name: 'Mirelle', class: 'priest', offered_at: '2026-11-10T12:00:00Z' };
const SEATED = (seats: GuildBillingSettings['seats']): GuildBillingSettings =>
  guildBillingRead({ billing: planView(), seats });

test('case 12: give to a non-payer holds it with its lock date', async ({ page }) => {
  const cell = await openLive(page, SEATED([emptySeat(1)]));
  await stubCandidates(page);
  await page.route(`**/v1/guilds/${GUILD_ID}/seats/1/give`, (route) => {
    cell.current = SEATED([heldSeat(1, { movable_at: '2099-01-02T12:00:00Z' })]);
    return route.fulfill(envelope(cell.current.seats[0]));
  });
  await page.getByTestId('guild-seat-1-give').click();
  await expect(fieldControl(page, 'guild-seat-1-select')).toBeFocused();
  await expect(page.getByTestId('guild-seat-1-pays-note')).toHaveCount(0);
  await page.getByTestId('guild-seat-1-submit').click();
  await expect(page.getByTestId('guild-seat-1-name')).toHaveText('Kraggor');
  await expect(page.getByTestId('guild-seat-1-note')).toContainText('Can move from 2 Jan');
  await expect(page.getByTestId('guild-seat-1-move')).toHaveAttribute('aria-disabled', 'true');
});

test('case 12: give to a payer shows the payer note, then Waiting, and Withdraw empties it', async ({
  page,
}) => {
  const cell = await openLive(page, SEATED([emptySeat(1)]));
  await stubCandidates(page);
  await page.route(`**/v1/guilds/${GUILD_ID}/seats/1/give`, (route) => {
    cell.current = SEATED([emptySeat(1, { offeree: PENDING_OFFEREE })]);
    return route.fulfill(envelope(cell.current.seats[0]));
  });
  await page.route(`**/v1/guilds/${GUILD_ID}/seats/1/withdraw`, (route) => {
    cell.current = SEATED([emptySeat(1)]);
    return route.fulfill(envelope(cell.current.seats[0]));
  });
  await page.getByTestId('guild-seat-1-give').click();
  await choose(page.getByTestId('guild-seat-1-select'), { label: 'Mirelle · Priest · pays for Premium' });
  await expect(page.getByTestId('guild-seat-1-pays-note')).toBeVisible();
  await page.getByTestId('guild-seat-1-submit').click();
  await expect(page.getByTestId('guild-seat-1')).toContainText('Waiting');
  await page.getByTestId('guild-seat-1-withdraw').click();
  await expect(page.getByTestId('guild-seat-1-give')).toBeVisible();
  await expect(page.getByTestId('guild-seat-1-give')).toBeFocused();
});

test('case 12: a held seat moving to a payer names the holder, then Moving to with Waiting, and only Withdraw', async ({
  page,
}) => {
  await openLive(page, SEATED([pendingSeat(1, { movable_at: '2000-01-01T12:00:00Z' })]));
  await expect(page.getByTestId('guild-seat-1-name')).toHaveText('Kraggor');
  await expect(page.getByTestId('guild-seat-1-note')).toContainText('Moving to Mirelle · offered 10 Nov');
  await expect(page.getByTestId('guild-seat-1-note')).toContainText('Waiting');
  await expect(page.getByTestId('guild-seat-1-move')).toHaveCount(0);
  await expect(page.getByTestId('guild-seat-1-withdraw')).toBeVisible();
});

test('case 12: a 409 seat_locked and a returned seat each show their line; Escape cancels', async ({
  page,
}) => {
  await openLive(page, SEATED([heldSeat(1, { movable_at: '2000-01-01T12:00:00Z' }), returnedSeat(2)]));
  await stubCandidates(page);
  await page.route(`**/v1/guilds/${GUILD_ID}/seats/1/move`, (route) =>
    route.fulfill(
      envelope(
        { code: 'seat_locked', message: 'locked', fields: { movable_at: '2026-12-02T12:00:00Z' } },
        409,
      ),
    ),
  );
  await expect(page.getByTestId('guild-seat-2-note')).toContainText('Kraggor left the guild.');
  await page.getByTestId('guild-seat-1-move').click();
  await page.getByTestId('guild-seat-1-submit').click();
  await expect(page.getByTestId('guild-seat-1-error')).toHaveText('That seat can move from 2 Dec.');
  await page.keyboard.press('Escape');
  await expect(page.getByTestId('guild-seat-1-picker')).toHaveCount(0);
  await expect(page.getByTestId('guild-seat-1-move')).toBeFocused();
});

test('case 13: send, wait and withdraw puts the button back', async ({ page }) => {
  const read = guildBillingRead({
    billing: planView(),
    transfer_candidates: [{ user_id: 40, name: 'Mirelle' }],
  });
  const cell = await openLive(page, read);
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/transfer`, (route) => {
    cell.current = { ...read, transfer: TRANSFER_OFFERED };
    return route.fulfill({ status: 204 });
  });
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/transfer/withdraw`, (route) => {
    cell.current = read;
    return route.fulfill({ status: 204 });
  });
  await page.getByTestId('guild-transfer-open').click();
  await expect(fieldControl(page, 'guild-transfer-select')).toBeFocused();
  await expect(page.getByTestId('guild-transfer-consequence')).toContainText(
    'Nothing changes until they accept.',
  );
  await expect(page.getByTestId('guild-transfer-stays')).toHaveText(
    'Seats, raids and records stay with the guild; Mirelle manages seats from then.',
  );
  await page.getByTestId('guild-transfer-send').click();
  await expect(page.getByTestId('guild-transfer-waiting')).toContainText('Waiting for them.');
  await expect(page.getByTestId('guild-transfer-withdraw')).toBeFocused();
  await page.getByTestId('guild-transfer-withdraw').click();
  await expect(page.getByTestId('guild-transfer-open')).toBeVisible();
  await expect(page.getByTestId('guild-transfer-open')).toBeFocused();
});

test('case 13: a decline shows to the sender and the button returns', async ({ page }) => {
  await openLive(
    page,
    guildBillingRead({
      billing: planView(),
      transfer: TRANSFER_DECLINED,
      transfer_candidates: [{ user_id: 40, name: 'Mirelle' }],
    }),
  );
  await expect(page.getByTestId('guild-transfer-declined')).toHaveText('Mirelle declined.');
  await expect(page.getByTestId('guild-transfer-open')).toBeVisible();
});

test('case 13: the receiver sees the offer and Take over the plan opens checkout in transfer mode', async ({
  page,
}) => {
  let body: unknown = null;
  await openLive(
    page,
    guildBillingRead({
      billing: planView({ you_are_billing_contact: false }),
      transfer: TRANSFER_RECEIVED,
      can: [],
    }),
  );
  await page.route('**/v1/billing/checkout', (route) => {
    body = route.request().postDataJSON();
    return route.fulfill(envelope({ checkout_url: 'https://checkout.stripe.com/transfer' }));
  });
  await page.route('https://checkout.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await expect(page.getByTestId('guild-transfer-offer-text')).toContainText('offers you the guild plan');
  await page.getByTestId('guild-transfer-take').click();
  await page.waitForURL('https://checkout.stripe.com/**');
  expect(body).toEqual({ plan: 'guild', interval: 'monthly', guild_id: GUILD_ID, intent: 'transfer' });
});

test('case 15: a read-only officer sees the seats with no buttons, and Records', async ({ page }) => {
  await openLive(page, READ_ONLY_OFFICER_READ);
  await expect(page.getByTestId('guild-seat-1-name')).toBeVisible();
  await expect(page.getByTestId('guild-seats').locator('button')).toHaveCount(0);
  await expect(page.getByTestId('guild-whopays')).toHaveCount(0);
  await expect(page.getByTestId('guild-records')).toBeVisible();
});
