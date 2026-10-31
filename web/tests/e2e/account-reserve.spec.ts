// web/tests/e2e/account-reserve.spec.ts
// The /account loading reserve (experience spec 2026-10-25-discord-notifications 9.1): for a
// visitor with the session hint the loading page already has the ready page's shape. Warm (a
// previous ready visit left the cached /v1/me and the stored heights) is gated per `data-discord`
// state at 390 and 1440; cold (nothing stored) is recorded and soft-checked against the spec's
// notes; a visitor without the hint sees no skeleton. /v1/me and /v1/devices are held, then
// released, by hand rather than delayed by a timer, so the loading frame is always on screen.
// Also: the gold checkbox rule resolves to the gold token on three surfaces (9.2).
import { expect, test } from './support/test';
import type { HintState } from '../../src/lib/account/layout';
import { discordBlock, type DiscordMeStages } from './support/discord-notifications';
import {
  addSessionHint,
  createHold,
  footerTravel,
  measureHeldLoad,
  seedWarm,
  stubMember,
  type LoadingMeasure,
} from './support/account-reserve';

const HEIGHT = 1000;
const WARM_MAX_CLS = { 390: 0.05, 1440: 0.03 } as const;
const COLD_NOTE_CLS = { 390: 0.15, 1440: 0.08 } as const;
const MAX_FOOTER_TRAVEL_PX = 8;
const WIDTHS = [390, 1440] as const;

const FIVE_FAILURES = discordBlock({
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
  failures: 5,
});

/** The seven `data-discord` states, each with the member that produces it. */
const STATES: ReadonlyArray<{ state: HintState; stages: Omit<DiscordMeStages, 'hold'> }> = [
  { state: 'none', stages: { initial: null, featureOff: true } },
  { state: 'free', stages: { initial: null, entitled: false } },
  { state: 'connect', stages: { initial: null, connect: true } },
  { state: 'form', stages: { initial: null } },
  { state: 'saved', stages: { initial: discordBlock() } },
  { state: 'off', stages: { initial: FIVE_FAILURES } },
  { state: 'lapsed', stages: { initial: discordBlock(), entitled: false } },
];

function describeMeasure(measure: LoadingMeasure): string {
  return `cls ${measure.cls.toFixed(4)} footer ${measure.footerLoading} -> ${measure.footerReady} (${footerTravel(measure)} px)`;
}

test.describe('warm: a cached /v1/me and stored heights from a previous ready visit', () => {
  for (const width of WIDTHS) {
    for (const { state, stages } of STATES) {
      test(`${state} @${width}: shift and footer movement stay inside the gates`, async ({
        page,
        baseURL,
      }) => {
        const hold = createHold();
        await page.setViewportSize({ width, height: HEIGHT });
        await addSessionHint(page, baseURL);
        await stubMember(page, stages, hold);
        await seedWarm(page);
        const measure = await measureHeldLoad(page, hold);
        test
          .info()
          .annotations.push({ type: 'warm', description: `${state} @${width}: ${describeMeasure(measure)}` });
        expect(measure.discord).toBe(state);
        expect(measure.cls).toBeLessThanOrEqual(WARM_MAX_CLS[width]);
        expect(footerTravel(measure)).toBeLessThanOrEqual(MAX_FOOTER_TRAVEL_PX);
      });
    }
  }
});

test.describe('cold: the session hint and nothing stored', () => {
  for (const width of WIDTHS) {
    for (const { state, stages } of STATES) {
      test(`${state} @${width}: recorded, soft-checked against the spec's note`, async ({
        page,
        baseURL,
      }) => {
        const hold = createHold();
        await page.setViewportSize({ width, height: HEIGHT });
        await addSessionHint(page, baseURL);
        await stubMember(page, stages, hold);
        const measure = await measureHeldLoad(page, hold);
        test
          .info()
          .annotations.push({ type: 'cold', description: `${state} @${width}: ${describeMeasure(measure)}` });
        // Nothing is stored, so the hint is `none` whatever the member turns out to be.
        expect(measure.discord).toBe('none');
        expect.soft(measure.cls).toBeLessThanOrEqual(COLD_NOTE_CLS[width]);
      });
    }
  }
});

test.describe('a visitor without the session hint', () => {
  for (const width of WIDTHS) {
    test(`@${width}: no skeleton is drawn while /v1/me is pending`, async ({ page }) => {
      const hold = createHold();
      await page.setViewportSize({ width, height: HEIGHT });
      await stubMember(page, { initial: null }, hold);
      hold.arm();
      await page.goto('/account', { waitUntil: 'commit' });
      await expect(page.locator('html')).not.toHaveAttribute('data-session', '1');
      await expect(page.locator('html')).not.toHaveAttribute('data-discord', /.*/);
      await expect(page.locator('[data-testid$="-skeleton"]:visible')).toHaveCount(0);
      hold.release();
      await expect(page.getByTestId('account')).toBeVisible();
    });
  }
});
