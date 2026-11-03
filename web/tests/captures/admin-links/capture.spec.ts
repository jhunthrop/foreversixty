// web/tests/captures/admin-links/capture.spec.ts
// Review screenshots of /admin/links, the owner's tracked links: a launch-week funnel for the
// owner at 1440 and 390, and the line anyone else sees. Not a test of behaviour
// (tests/e2e/admin-links.spec.ts is); it writes PNGs to design/reviews/2026-11-02-admin-links/.
import path from 'node:path';
import { test, type Page } from '@playwright/test';
import type { GoLinkFunnel } from '../../../src/lib/golinks/api';

const OUT = path.resolve(import.meta.dirname, '../../../../design/reviews/2026-11-02-admin-links');

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const me = (role: string) => ({
  user: { id: 78, battletag: 'hunthrop#1894', email: null, role, anonymize: false },
  characters: [],
  guilds: [],
});

function link(slug: string, target: string, label: string, figures: readonly number[]): GoLinkFunnel {
  const [clicks, bot_clicks, sign_ins, linked, premium] = figures;
  return {
    slug,
    target,
    label,
    created_at: '2026-11-02T00:00:00Z',
    clicks,
    bot_clicks,
    sign_ins,
    linked,
    premium,
  };
}

const LINKS: GoLinkFunnel[] = [
  link('reddit', '/', 'Reddit, any subreddit', [412, 38, 61, 44, 9]),
  link('reddit-classicwow', '/', 'Reddit: r/classicwow', [1280, 96, 143, 101, 17]),
  link('reddit-wowforever', '/', 'Reddit: r/wowforever', [2210, 120, 260, 198, 31]),
  link('reddit-wowclassic', '/', 'Reddit: r/wowclassic', [190, 14, 18, 11, 1]),
  link('reddit-guides', '/guides', 'Reddit post linking the guides', [655, 40, 52, 39, 4]),
  link('reddit-sim', '/sim', 'Reddit post linking the sim', [870, 31, 97, 80, 14]),
  link('discord', '/', 'Discord', [340, 512, 71, 66, 12]),
  link('youtube', '/', 'YouTube', [96, 4, 8, 6, 0]),
  link('curseforge', '/', 'CurseForge', [58, 2, 9, 9, 2]),
  link('wago', '/', 'Wago', [21, 1, 3, 3, 0]),
  link('twitter', '/', 'Twitter', [0, 0, 0, 0, 0]),
];

async function open(page: Page, role: string): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me(role))));
  await page.route('**/v1/admin/golinks', (route) => route.fulfill(envelope({ links: LINKS })));
  await page.goto('/admin/links');
  await page.waitForSelector(
    role === 'admin' ? '[data-testid="links-form"]' : '[data-testid="links-not-owner"]',
  );
  await page.waitForTimeout(600);
}

test.describe('desktop', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test('owner at 1440', async ({ page }) => {
    await open(page, 'admin');
    await page.screenshot({ path: path.join(OUT, 'admin-links-1440.png'), fullPage: true });
  });
  test('not the owner at 1440', async ({ page }) => {
    await open(page, 'user');
    await page.screenshot({ path: path.join(OUT, 'admin-links-not-owner-1440.png'), fullPage: true });
  });
});

test.describe('phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, isMobile: true, hasTouch: true, deviceScaleFactor: 2 });
  test('owner at 390', async ({ page }) => {
    await open(page, 'admin');
    await page.screenshot({ path: path.join(OUT, 'admin-links-390.png'), fullPage: true });
  });
});
