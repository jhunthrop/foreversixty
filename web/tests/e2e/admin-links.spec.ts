// web/tests/e2e/admin-links.spec.ts
// /admin/links is the owner's tracked links: the funnel per link, each step's share of the one
// before, and a form that makes a new link. Anyone else sees one line, never the figures. And a
// /go/ landing's ?golink= is kept for the first sign-in and taken out of the address bar.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';

const fulfil = (body: unknown, status = 200, fields: Record<string, string> | null = null) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? body : null,
    error: status < 400 ? null : { code: 'x', message: 'refused', fields },
    request_id: 'r',
  }),
});

const me = (role: string) => ({
  user: { id: 78, battletag: 'Owner#1894', email: null, role, anonymize: false },
  characters: [],
  guilds: [],
});

const link = (slug: string, target: string, label: string, figures: number[]) => ({
  slug,
  target,
  label,
  created_at: '2026-11-02T00:00:00Z',
  clicks: figures[0],
  bot_clicks: figures[1],
  sign_ins: figures[2],
  linked: figures[3],
  premium: figures[4],
});

const LINKS = [
  link('reddit', '/', 'Reddit', [40, 6, 10, 5, 1]),
  link('reddit-sim', '/sim', 'Reddit: sim post', [120, 2, 12, 9, 3]),
  link('wago', '/', 'Wago', [0, 0, 0, 0, 0]),
];

/** A link's row in whichever layout is showing: the table from 768, the list below it. */
const linkRow = (page: Page, slug: string) =>
  page.locator(`[data-testid="links-row-${slug}"]:visible, [data-testid="links-item-${slug}"]:visible`);

async function asOwner(page: Page, links = LINKS): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(me('admin'))));
  await page.route('**/v1/admin/golinks', (route) => route.fulfill(fulfil({ links })));
}

test('the owner sees every link busiest first, with each step’s share and the totals', async ({ page }) => {
  await asOwner(page);
  await page.goto('/admin/links');
  await expect(page.getByRole('heading', { level: 1, name: 'Links' })).toBeVisible();
  const rows = page.locator('[data-testid^="links-row-"]');
  await expect(rows).toHaveCount(3);
  await expect(rows.first()).toHaveAttribute('data-testid', 'links-row-reddit-sim');
  const sim = linkRow(page, 'reddit-sim');
  await expect(sim).toContainText('/go/reddit-sim');
  await expect(sim).toContainText('/sim');
  await expect(sim).toContainText('2 bots');
  await expect(sim).toContainText('10%'); // 12 of 120 signed up
  await expect(sim).toContainText('75%'); // 9 of 12 linked a character
  await expect(sim).toContainText('33%'); // 3 of 9 hold Premium
  await expect(page.getByTestId('links-totals')).toContainText('160');
  await expect(page.getByTestId('links-totals')).toContainText('8 bots');
});

test('the copy button puts the short link on the clipboard', async ({ page, context }) => {
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await asOwner(page);
  await page.goto('/admin/links');
  await linkRow(page, 'reddit').getByTestId('links-copy-reddit').click();
  await expect(linkRow(page, 'reddit').getByTestId('links-copy-reddit')).toHaveText('Copied');
  expect(await page.evaluate(() => navigator.clipboard.readText())).toBe('https://foreversixty.gg/go/reddit');
});

test('a new link is saved and joins the table', async ({ page }) => {
  await asOwner(page);
  let posted: unknown = null;
  await page.route('**/v1/admin/golinks', async (route) => {
    if (route.request().method() === 'POST') {
      posted = route.request().postDataJSON();
      await route.fulfill(fulfil({ ...(posted as object), created_at: '2026-11-02T01:00:00Z' }, 201));
      return;
    }
    const extra = posted === null ? [] : [link('reddit-tbc', '/guides', 'Reddit: TBC post', [0, 0, 0, 0, 0])];
    await route.fulfill(fulfil({ links: [...LINKS, ...extra] }));
  });
  await page.goto('/admin/links');
  await page.getByTestId('links-slug').fill('reddit-tbc');
  await page.getByTestId('links-target').fill('/guides');
  await page.getByTestId('links-label').fill('Reddit: TBC post');
  await page.getByTestId('links-create').click();
  await expect(page.getByTestId('links-created')).toHaveText('foreversixty.gg/go/reddit-tbc is live.');
  expect(posted).toEqual({ slug: 'reddit-tbc', target: '/guides', label: 'Reddit: TBC post' });
  await expect(linkRow(page, 'reddit-tbc')).toBeVisible();
  await expect(page.getByTestId('links-slug')).toHaveValue('');
});

test('a refused field says why beside the field', async ({ page }) => {
  await asOwner(page);
  await page.route('**/v1/admin/golinks', async (route) => {
    if (route.request().method() === 'POST') {
      await route.fulfill(fulfil(null, 400, { target: 'a path on this site, like /guides' }));
      return;
    }
    await route.fulfill(fulfil({ links: LINKS }));
  });
  await page.goto('/admin/links');
  await page.getByTestId('links-slug').fill('x');
  await page.getByTestId('links-target').fill('https://elsewhere.example');
  await page.getByTestId('links-label').fill('X');
  await page.getByTestId('links-create').click();
  await expect(page.locator('#link-target-error')).toHaveText('a path on this site, like /guides');
  await expect(page.getByTestId('links-target')).toHaveAttribute('aria-invalid', 'true');
});

test('anyone but the owner sees one line and no figures', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(me('user'))));
  const asked: string[] = [];
  await page.route('**/v1/admin/golinks', (route) => {
    asked.push(route.request().url());
    return route.fulfill(fulfil(null, 403));
  });
  await page.goto('/admin/links');
  await expect(page.getByTestId('links-not-owner')).toContainText('This page is for the site’s owner.');
  await expect(page.locator('[data-testid^="links-row-"]')).toHaveCount(0);
  expect(asked).toEqual([]);
});

test('signed out, it asks for the owner’s sign-in and comes back here', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/admin/links');
  await expect(
    page.getByTestId('links-signin').getByRole('link', { name: 'Sign in with Battle.net' }),
  ).toHaveAttribute('href', /next=%2Fadmin%2Flinks/);
});

test('the page is never indexed', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/admin/links');
  await expect(page.locator('meta[name="robots"]')).toHaveAttribute('content', 'noindex');
});

test('a /go/ landing keeps the slug, cleans the address bar and sends it once signed in', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/guides?golink=reddit-guides');
  await expect(page).toHaveURL(/\/guides$/);
  // The click is kept after first paint (lib/golinks/later.ts), so the read waits for it.
  await expect
    .poll(async () => {
      const kept = await page.evaluate(() => localStorage.getItem('fs.arrival.next'));
      return (JSON.parse(kept ?? '{}') as { slug?: string }).slug;
    })
    .toBe('reddit-guides');

  let sent: unknown = null;
  await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me/arrival', async (route) => {
    sent = { body: route.request().postDataJSON(), csrf: route.request().headers()['x-csrf-token'] };
    await route.fulfill(fulfil({ recorded: true }));
  });
  await page.goto('/guides');
  await expect.poll(() => sent).toMatchObject({ body: { slug: 'reddit-guides' }, csrf: 'tok' });
  await expect.poll(() => page.evaluate(() => localStorage.getItem('fs.arrival'))).toBeNull();
});

test('no page loads the visit beacon off the production host', async ({ page }) => {
  const beacon: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('cloudflareinsights')) beacon.push(request.url());
  });
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/');
  await page.waitForLoadState('load');
  await page.waitForTimeout(1500);
  expect(beacon).toEqual([]);
});
