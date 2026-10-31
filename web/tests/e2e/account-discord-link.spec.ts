// web/tests/e2e/account-discord-link.spec.ts
// The You panel's Discord link row (experience spec 2026-10-27-discord-bot §3.2 and §9 items 8
// and 9): unlinked default and suggested, each `discord_link` return flag, the start flag that
// goes on to the API, unlink asking in place, `Opening Discord…` and the 10 s stuck line under
// fake timers, the row absent when unavailable, the `#you` anchor, axe, 44 px at 360 and the
// warm layout shift at 390 and 1440. /v1/me, /v1/devices, /v1/reports and the unlink route are
// stubbed; the start route is the API's and is only observed.
import { createRequire } from 'node:module';
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { discordLinkCopy as copy } from '../../src/lib/account/discord-link-copy';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { E2E_API_BASE, envelope } from './support/discord-notifications';
import { addSessionHint, createHold, measureHeldLoad, seedWarm, type Hold } from './support/account-reserve';
import { gotoHydrated } from './support/hydrated';
import { recordShifts } from './support/layout-shift';

const START_URL = `${E2E_API_BASE}/v1/discord/link/start`;
const UNLINK_GLOB = '**/v1/discord/link';
const DISPLAY_NAME = 'zulmara';
const GUILD_NAME = 'Olympus XXVII';
const MAX_SHIFT = 0.02;
const STUCK_MS = 10_000;
const MIN_TARGET_PX = 44;
const PHONE = { width: 390, height: 844 };
const SMALL_PHONE = { width: 360, height: 700 };
const DESKTOP = { width: 1440, height: 1000 };

interface LinkState {
  available?: boolean;
  linked?: boolean;
  suggested?: string | null;
  hold?: Hold;
  /** The unlink route answers this status; 204 by default. */
  unlinkStatus?: number;
}

interface Seen {
  unlinks: number;
}

function meBody(state: LinkState, seen: Seen): unknown {
  const linked = state.linked === true && seen.unlinks === 0;
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    entitlements: mockEntitlements({ notifications: true }),
    ...(state.available === false
      ? {}
      : {
          discord_link_available: true,
          discord_link: linked ? { display_name: DISPLAY_NAME } : null,
          discord_link_suggested: linked ? null : (state.suggested ?? null),
        }),
  };
}

async function stubAccount(page: Page, state: LinkState): Promise<Seen> {
  const seen: Seen = { unlinks: 0 };
  await page.route('**/v1/me', async (route) => {
    await state.hold?.wait();
    await route.fulfill(envelope(meBody(state, seen)));
  });
  await page.route('**/v1/reports**', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 20 })),
  );
  await page.route('**/v1/devices', async (route) => {
    await state.hold?.wait();
    await route.fulfill(envelope([]));
  });
  await page.route(UNLINK_GLOB, (route) => {
    const status = state.unlinkStatus ?? 204;
    if (status === 204) seen.unlinks += 1;
    return route.fulfill(
      status === 204
        ? { status }
        : {
            status,
            contentType: 'application/json',
            body: JSON.stringify({ ok: false, data: null, error: { message: 'broken' }, request_id: 'r' }),
          },
    );
  });
  return seen;
}

async function openAccount(page: Page, state: LinkState, path = '/account#you'): Promise<Seen> {
  const seen = await stubAccount(page, state);
  await gotoHydrated(page, path, 'account');
  return seen;
}

const row = (page: Page) => page.getByTestId('account-discord-link');
const link = (page: Page) => page.getByTestId('discord-link-start');
const unlink = (page: Page) => page.getByTestId('discord-unlink');
const heightOf = async (locator: Locator): Promise<number> => (await locator.boundingBox())?.height ?? 0;
const status = (page: Page) => page.getByTestId('discord-link-status');

test.describe('unlinked', () => {
  test('the default lead with a secondary button that goes to the start route', async ({ page }) => {
    await openAccount(page, {});
    await expect(row(page)).toContainText(copy.lead);
    await expect(link(page)).toHaveText(copy.link);
    await expect(link(page)).toHaveAttribute('href', START_URL);
    await expect(link(page)).not.toHaveClass(/linear-gradient/);
  });

  test('suggested: the guild lead and the one fill', async ({ page }) => {
    await openAccount(page, { suggested: GUILD_NAME });
    await expect(row(page)).toContainText(copy.leadSuggested(GUILD_NAME));
    await expect(link(page)).toHaveClass(/linear-gradient/);
  });

  test('absent when the deployment has the bot off', async ({ page }) => {
    await openAccount(page, { available: false });
    await expect(page.getByTestId('account-you')).toBeVisible();
    await expect(row(page)).toHaveCount(0);
  });
});

test.describe('return flags', () => {
  const cases = [
    { flag: 'denied', message: copy.returnDenied, role: 'status', focus: 'discord-link-start' },
    { flag: 'failed', message: copy.returnFailed, role: 'alert', focus: 'discord-link-start' },
    { flag: 'conflict', message: copy.returnConflict, role: 'alert', focus: 'discord-link-start' },
  ] as const;
  for (const { flag, message, role, focus } of cases) {
    test(`${flag}: the line, its role, focus, and the flag dropped`, async ({ page }) => {
      await openAccount(page, {}, `/account?discord_link=${flag}#you`);
      await expect(status(page)).toHaveAttribute('role', role);
      await expect(status(page)).toContainText(message);
      await expect(page.getByTestId(focus)).toBeFocused();
      await expect(page).toHaveURL(/\/account#you$/);
      if (flag === 'conflict') {
        await expect(status(page)).toContainText(copy.conflictHelp);
        await expect(status(page).locator('p').first()).toHaveClass(/text-wipe/);
      }
    });
  }

  test('linked: the linked row, the line, focus on Unlink', async ({ page }) => {
    await openAccount(page, { linked: true }, '/account?discord_link=linked#you');
    await expect(status(page)).toHaveAttribute('role', 'status');
    await expect(status(page)).toHaveText(copy.returnLinked);
    await expect(unlink(page)).toBeFocused();
    await expect(page.getByTestId('discord-linked-as')).toContainText(`${copy.linkedLead} ${DISPLAY_NAME}`);
    await expect(page).toHaveURL(/\/account#you$/);
  });

  test('start with a session goes on to the start route', async ({ page }) => {
    await page.route(START_URL, (route) => route.fulfill({ status: 200, body: 'Discord' }));
    const started = page.waitForRequest(START_URL);
    await stubAccount(page, {});
    await page.goto('/account?discord_link=start#you');
    expect((await started).url()).toBe(START_URL);
  });
});

test.describe('unlink', () => {
  test('asks in place; Escape cancels', async ({ page }) => {
    await openAccount(page, { linked: true });
    await unlink(page).click();
    await expect(page.getByTestId('discord-unlink-question')).toHaveText(copy.unlinkQuestion);
    await expect(page.getByTestId('discord-linked-as')).toBeVisible();
    await expect(page.getByTestId('discord-unlink-cancel')).toBeFocused();
    expect(
      await page.getByTestId('discord-unlink-cancel').evaluate((el) => el.matches(':focus-visible')),
    ).toBe(true);
    await page.keyboard.press('Escape');
    await expect(page.getByTestId('discord-unlink-question')).toHaveCount(0);
    await expect(unlink(page)).toBeFocused();
  });

  test('success shows Unlinked. and focuses Link Discord', async ({ page }) => {
    const seen = await openAccount(page, { linked: true });
    await unlink(page).click();
    await page.getByTestId('discord-unlink-confirm').click();
    await expect(status(page)).toHaveText(copy.unlinked);
    await expect(link(page)).toBeFocused();
    expect(seen.unlinks).toBe(1);
    await expect(status(page)).toHaveText('', { timeout: 6000 });
  });

  test('a failure says so and the question stays', async ({ page }) => {
    await openAccount(page, { linked: true, unlinkStatus: 500 });
    await unlink(page).click();
    await page.getByTestId('discord-unlink-confirm').click();
    await expect(status(page)).toHaveAttribute('role', 'alert');
    await expect(status(page)).toHaveText(copy.unlinkFailed);
    await expect(page.getByTestId('discord-unlink-question')).toBeVisible();
  });
});

test('Opening Discord… and the 10 s stuck line under fake timers', async ({ page }) => {
  await page.clock.install();
  await openAccount(page, {});
  // Keep the page where it is: the click's own handler has run by the time this one does.
  await page.evaluate(() =>
    window.addEventListener('click', (event) => {
      if (event.target instanceof Element && event.target.closest('[data-testid="discord-link-start"]'))
        event.preventDefault();
    }),
  );
  await link(page).click();
  await expect(link(page)).toHaveText(copy.opening);
  await expect(link(page)).toHaveAttribute('aria-busy', 'true');
  await page.clock.fastForward(STUCK_MS);
  await expect(link(page)).toHaveText(copy.link);
  await expect(status(page)).toHaveText(copy.stuck);
});

test.describe('layout', () => {
  test('#you scrolls the panel to the top', async ({ page }) => {
    await page.setViewportSize({ width: DESKTOP.width, height: 400 });
    await openAccount(page, {});
    const panel = page.locator('#you');
    await expect(panel).toHaveCSS('scroll-margin-top', '80px');
    await expect.poll(() => page.evaluate(() => window.scrollY)).toBeGreaterThan(0);
    const top = (await panel.boundingBox())?.y ?? Infinity;
    expect(top).toBeGreaterThanOrEqual(0);
    expect(top).toBeLessThanOrEqual(100);
  });

  test('every control is 44 px tall at 360', async ({ page }) => {
    await page.setViewportSize(SMALL_PHONE);
    await openAccount(page, {});
    expect(await heightOf(link(page))).toBeGreaterThanOrEqual(MIN_TARGET_PX);
  });

  test('the linked row and its question are 44 px tall at 360', async ({ page }) => {
    await page.setViewportSize(SMALL_PHONE);
    await openAccount(page, { linked: true });
    expect(await heightOf(unlink(page))).toBeGreaterThanOrEqual(MIN_TARGET_PX);
    await unlink(page).click();
    for (const id of ['discord-unlink-confirm', 'discord-unlink-cancel']) {
      expect(await heightOf(page.getByTestId(id))).toBeGreaterThanOrEqual(MIN_TARGET_PX);
    }
  });
});

interface AxeResults {
  violations: { id: string; nodes: { target: unknown }[] }[];
}

const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

async function expectNoAxeViolations(page: Page): Promise<void> {
  await page.evaluate(() => Promise.all(document.getAnimations().map((animation) => animation.finished)));
  await page.addScriptTag({ path: AXE_SOURCE });
  const results = await page.evaluate(async (): Promise<AxeResults> => {
    const axe = (window as unknown as { axe: { run: (context: string) => Promise<AxeResults> } }).axe;
    return axe.run('[data-testid="account"]');
  });
  expect(results.violations.map((violation) => violation.id)).toEqual([]);
}

test.describe('axe', () => {
  test.use({ reducedMotion: 'reduce' });
  const states: { name: string; state: LinkState; path?: string }[] = [
    { name: 'unlinked', state: {} },
    { name: 'suggested', state: { suggested: GUILD_NAME } },
    { name: 'linked', state: { linked: true } },
    { name: 'conflict', state: {}, path: '/account?discord_link=conflict#you' },
  ];
  for (const { name, state, path } of states) {
    test(`${name} has no violations`, async ({ page }) => {
      await openAccount(page, state, path);
      await expectNoAxeViolations(page);
    });
  }

  test('the unlink question has no violations', async ({ page }) => {
    await openAccount(page, { linked: true });
    await unlink(page).click();
    await expectNoAxeViolations(page);
  });
});

test.describe('warm layout shift', () => {
  const states: { hint: string; state: Omit<LinkState, 'hold'> }[] = [
    { hint: 'unlinked', state: {} },
    { hint: 'suggested', state: { suggested: GUILD_NAME } },
    { hint: 'linked', state: { linked: true } },
  ];
  for (const viewport of [PHONE, DESKTOP]) {
    for (const { hint, state } of states) {
      test(`${hint} @${viewport.width} stays within ${MAX_SHIFT}`, async ({ page, baseURL }) => {
        const hold = createHold();
        await page.setViewportSize(viewport);
        await addSessionHint(page, baseURL);
        await recordShifts(page);
        await stubAccount(page, { ...state, hold });
        await seedWarm(page);
        const measure = await measureHeldLoad(page, hold);
        await expect(page.locator('html')).toHaveAttribute('data-discord-link', hint);
        test
          .info()
          .annotations.push({ type: 'warm', description: `${hint} @${viewport.width}: ${measure.cls}` });
        expect(measure.cls).toBeLessThanOrEqual(MAX_SHIFT);
      });
    }
  }
});
