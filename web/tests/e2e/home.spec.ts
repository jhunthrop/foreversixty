// web/tests/e2e/home.spec.ts
// Home, layout B (design/specs/2026-11-07-home-b.md): what each state offers and where it leads.
// The layout is the parity gate's (parity-home.spec.ts); this spec is the behaviour.
import { expect, test } from './support/test';
import { homeEnvelope, openHome } from './support/home-page';

test('signed out: the title, Sign in, the nine classes, the raid countdown, the feed and how to start', async ({
  page,
}) => {
  await openHome(page, { feed: 'feed-out.json' });
  const out = page.getByTestId('home-signed-out');
  await expect(out.getByRole('heading', { level: 1 })).toHaveText('Play your class better.');
  await expect(page.getByTestId('home-sign-in')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=/,
  );
  await expect(page.getByTestId('home-class-picker').getByRole('link')).toHaveCount(9);
  await expect(out.getByTestId('home-raids-countdown').getByRole('timer')).toHaveAccessibleName(
    '32 days, 6 hours, 0 min',
  );
  await expect(out.getByTestId('home-feed-row')).toHaveCount(7);
  await expect(out.getByTestId('home-feed').getByRole('radiogroup')).toHaveCount(0);
  await expect(out.getByTestId('home-feed-row').first()).toContainText('4 min ago');
  await expect(page.getByTestId('home-how-to-start').getByRole('link')).toHaveCount(2);
  await expect(page.getByTestId('home-signed-in-root')).toBeHidden();
});

test('signed in at 34: the three figures, the feed with its filters, upgrades, the grade and the movers', async ({
  page,
}) => {
  await openHome(page, { me: 'me-34.json', feed: 'feed-34.json', home: 'home-34.json' });
  await expect(page.getByTestId('home-signed-out')).toBeHidden();
  await expect(page.getByTestId('home-head').getByRole('heading', { level: 1 })).toHaveText('Felsnap');
  await expect(page.getByTestId('home-figure-sim')).toContainText('58.4');
  await expect(page.getByTestId('home-figure-sim')).toContainText('Simmed 6 Nov');
  await expect(page.getByTestId('home-figure-band')).toContainText('To band 40');
  await expect(page.getByTestId('home-guild-countdown')).toContainText('Find a guild');
  const filters = page.getByTestId('home-feed-filter').getByRole('radio');
  await expect(filters).toHaveText(['All', 'Your spec', 'Logs']);
  await expect(page.getByTestId('home-upgrade')).toHaveCount(3);
  await expect(page.getByTestId('home-spec-week')).toContainText(
    'Up from B last week, #4 of 23 damage specs',
  );
  await expect(page.getByTestId('home-mover')).toHaveCount(3);
  await expect(page.getByTestId('home-guild-sims')).toHaveCount(0);

  const asked = page.waitForRequest((request) => request.url().includes('filter=spec'));
  await filters.nth(1).click();
  const url = new URL((await asked).url());
  expect(url.searchParams.get('spec')).toBe('hunter-beast-mastery');
  expect(url.searchParams.get('band')).toBe('30');
});

test('signed in at 60 in a guild: the guild event, its sims ladder and the Your guild filter', async ({
  page,
}) => {
  await openHome(page, {
    me: 'me-60.json',
    feed: 'feed-60.json',
    home: 'home-60.json',
    guild: 'guild-home-60.json',
  });
  const guild = page.getByTestId('home-guild-countdown');
  await expect(guild).toContainText("Onyxia's Lair");
  await expect(guild).toContainText('you: late');
  await expect(guild.getByRole('link', { name: /Sign up on the raid page/ })).toHaveAttribute(
    'href',
    /\/guild\/us\/normal\/.+\/raids\/41$/,
  );
  await expect(page.getByTestId('home-feed-filter').getByRole('radio')).toHaveCount(4);
  await expect(page.getByTestId('home-guild-sim')).toHaveCount(6);
  await expect(page.getByTestId('home-guild-sim').last()).toContainText('Thandrel');
  await expect(page.getByTestId('home-figure-band')).toContainText('On best in slot');
  await expect(page.getByTestId('home-movers')).toContainText('Top 3 for Fury Warrior, 60');
});

test('a quiet feed fills with real dated site events, the raid opening first', async ({ page }) => {
  await openHome(page, { me: 'me-34.json', feed: 'feed-empty.json', home: 'home-34.json' });
  const rows = page.getByTestId('home-signed-in-root').getByTestId('home-feed-row');
  await expect(rows).toHaveCount(5);
  await expect(rows.first()).toContainText('First raids open');
  await expect(rows.first()).toContainText('Wed 9 Dec');
});

test('the feed fails on its own and comes back with Try again', async ({ page }) => {
  await openHome(page, { me: 'me-34.json', feed: 'fail', home: 'home-34.json' });
  const root = page.getByTestId('home-signed-in-root');
  const error = root.getByTestId('home-feed-error');
  await expect(error).toContainText('We could not load what is happening on Forever.');
  await expect(page.getByTestId('home-upgrade')).toHaveCount(3);
  await page.unroute('**/v1/feed?**');
  await page.route('**/v1/feed?**', (route) => route.fulfill(homeEnvelope('feed-34.json')));
  await error.getByRole('button', { name: 'Try again' }).click();
  await expect(root.getByTestId('home-feed-row')).toHaveCount(7);
});

test('signed in with no character: check again, paste, the addon, and the classes', async ({ page }) => {
  await openHome(page, { me: 'me-new.json' });
  await expect(page.getByTestId('home-next-step')).toContainText('No character yet.');
  await expect(page.getByTestId('home-check-again')).toBeVisible();
  await expect(page.getByTestId('home-get-in').getByRole('link')).toHaveCount(2);
  await expect(
    page.getByTestId('home-signed-in-root').getByTestId('home-class-picker').getByRole('link'),
  ).toHaveCount(9);
});

test('a stale session hint gives the signed-out page back', async ({ page }) => {
  await page.route('**/v1/feed?**', (route) => route.fulfill(homeEnvelope('feed-empty.json')));
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.goto('/');
  await expect(page.getByTestId('home-signed-out')).toBeVisible();
  await expect(page.locator('html[data-home-out]')).toHaveCount(1);
});
