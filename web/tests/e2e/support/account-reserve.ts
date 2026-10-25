// web/tests/e2e/support/account-reserve.ts
// Drives /account through its loading state for the specs that gate the loading reserve (experience
// spec 2026-10-25-discord-notifications 9.1): the session hint, a previous ready visit that leaves
// the cached /v1/me and the stored heights behind (warm), and a load held in its loading state,
// then released, with the layout shift and the footer's travel measured across it.
import { expect, type Page } from '@playwright/test';
import { ACCOUNT_RESERVE_KEY } from '../../../src/lib/account/layout';
import {
  routeDiscordNotifications,
  signInMember,
  type DiscordMeStages,
  type DiscordRoutes,
} from './discord-notifications';
import { gotoHydrated } from './hydrated';
import { drainShifts, recordShifts, sumShifts } from './layout-shift';

/** A switch for the stubbed `/v1/me` and `/v1/devices`: open until `arm()`, then closed until `release()`. */
export interface Hold {
  /** Passed to the stub as `stages.hold`. */
  wait: () => Promise<void>;
  arm: () => void;
  release: () => void;
}

export function createHold(): Hold {
  let gate: Promise<void> = Promise.resolve();
  let open = (): void => undefined;
  return {
    wait: () => gate,
    arm: () => {
      gate = new Promise<void>((resolve) => {
        open = resolve;
      });
    },
    release: () => open(),
  };
}

/** The readable half of the session cookie: what makes `<html data-session="1">` and the skeletons draw. */
export async function addSessionHint(page: Page, baseURL: string | undefined): Promise<void> {
  if (baseURL === undefined) throw new Error('the project has no baseURL to set the session cookie on');
  await page.context().addCookies([{ name: 'fs_csrf', value: 'e2e', url: baseURL }]);
}

export interface LoadingMeasure {
  /** `<html data-discord>` while the page was loading. */
  discord: string | null;
  /** Summed layout shift over the whole load, loading frame to ready page. */
  cls: number;
  footerLoading: number;
  footerReady: number;
}

const footerTop = (page: Page): Promise<number> =>
  page.evaluate(() => (document.querySelector('footer')?.getBoundingClientRect().top ?? 0) + window.scrollY);

/** Stubs the member, in `stages`, behind `hold`; the shift observer is installed for every load after. */
export async function stubMember(
  page: Page,
  stages: Omit<DiscordMeStages, 'hold'>,
  hold: Hold,
): Promise<DiscordRoutes> {
  await recordShifts(page);
  const seen = await routeDiscordNotifications(page);
  await signInMember(page, { ...stages, hold: hold.wait }, seen);
  return seen;
}

/** One ready visit: the page writes the cached /v1/me and the heights it will reserve next time. */
export async function seedWarm(page: Page): Promise<void> {
  await gotoHydrated(page, '/account', 'account');
  await expect
    .poll(() => page.evaluate((key) => localStorage.getItem(key), ACCOUNT_RESERVE_KEY))
    .not.toBeNull();
}

/**
 * Loads /account with `/v1/me` held, lets `onLoading` photograph the loading frame, releases the
 * answer, and measures the shift and the footer's travel until the ready page is still.
 */
export async function measureHeldLoad(
  page: Page,
  hold: Hold,
  onLoading: () => Promise<void> = async () => undefined,
): Promise<LoadingMeasure> {
  hold.arm();
  await page.goto('/account', { waitUntil: 'commit' });
  await expect(page.getByTestId('account-left-skeleton')).toBeVisible();
  const discord = await page.evaluate(() => document.documentElement.dataset.discord ?? null);
  const footerLoading = await footerTop(page);
  await onLoading();
  hold.release();
  await expect(page.getByTestId('account')).toBeVisible();
  await expect(page.getByTestId('account-left-skeleton')).toHaveCount(0);
  const cls = sumShifts(await drainShifts(page));
  return { discord, cls, footerLoading, footerReady: await footerTop(page) };
}

export const footerTravel = (measure: LoadingMeasure): number =>
  Math.abs(measure.footerReady - measure.footerLoading);
