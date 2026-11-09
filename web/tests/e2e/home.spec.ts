// web/tests/e2e/home.spec.ts
// Home, round 5 (design/specs/2026-11-09-directions.md, on home B's states): what each state offers
// and where it leads. The layout is the parity gate's (parity-home.spec.ts); this spec is the behaviour.
import { expect, test } from './support/test';
import { homeEnvelope, openHome } from './support/home-page';

// The fixtures' raid starts at 8 pm in the guild's own zone; read it as a player there would.
test.use({ timezoneId: 'America/Chicago' });

const SIXTY = { me: 'me-60.json', feed: 'feed-60.json', home: 'home-60.json', guild: 'guild-home-60.json' };

test('signed out: Pick your class with Sign in, the nine classes, the feed, how to start and the first raids', async ({
  page,
}) => {
  await openHome(page, { feed: 'feed-out.json' });
  const out = page.getByTestId('home-signed-out');
  await expect(out.getByRole('heading', { level: 1 })).toHaveText('Pick your class');
  await expect(page.getByTestId('home-sign-in')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=/,
  );
  await expect(page.getByTestId('home-class-picker').getByRole('link')).toHaveCount(9);
  const plaque = out.getByTestId('home-raids-countdown');
  await expect(plaque.getByRole('heading', { level: 2 })).toHaveText('First raids open');
  await expect(plaque.getByRole('timer')).toHaveAccessibleName('32 days, 6 hours');
  await expect(plaque.getByRole('listitem')).toHaveText(['Barrow Deeps', 'Hyjal Summit', "Onyxia's Lair"]);
  await expect(out.getByTestId('home-feed-row')).toHaveCount(7);
  await expect(out.getByTestId('home-feed').getByRole('radiogroup')).toHaveCount(0);
  await expect(out.getByTestId('home-feed-row').first()).toContainText('4 min ago');
  await expect(page.getByTestId('home-how-to-start').getByRole('link')).toHaveCount(2);
  await expect(page.getByTestId('home-signed-in-root')).toBeHidden();
});

test('signed in at 34: the sim, the first raids with Find a guild, the feed, upgrades and the grade', async ({
  page,
}) => {
  await openHome(page, { me: 'me-34.json', feed: 'feed-34.json', home: 'home-34.json' });
  await expect(page.getByTestId('home-signed-out')).toBeHidden();
  const head = page.getByTestId('home-head');
  await expect(head.getByRole('heading', { level: 1 })).toHaveText('Felsnap');
  await expect(head).toContainText('Level 34 Dwarf Beast Mastery Hunter');
  await expect(page.getByTestId('home-figure-sim')).toHaveText(/58\.4\s*DPS/);
  const plaque = page.getByTestId('home-signed-in-root').getByTestId('home-raids-countdown');
  await expect(plaque.getByRole('link', { name: /Guilds/ })).toHaveAttribute('href', '/guild');
  await expect(page.getByTestId('home-feed').getByRole('radiogroup')).toHaveCount(0);
  await expect(page.getByTestId('home-upgrade')).toHaveCount(5);
  await expect(page.getByTestId('home-spec-week')).toContainText(
    'Up from B last week, #4 of 23 damage specs',
  );
  await expect(page.getByTestId('home-guild-sims')).toHaveCount(0);
});

test('signed in at 60 in a guild: the raid plaque, the feed, the sims ladder with its tier line', async ({
  page,
}) => {
  await openHome(page, SIXTY);
  const plaque = page.getByTestId('home-raid');
  await expect(plaque.getByRole('heading', { level: 2 })).toHaveText("Onyxia's Lair");
  await expect(plaque).toContainText('Olympus XXVII · Wednesday 9 December · 8 pm');
  await expect(plaque.getByRole('timer')).toHaveAccessibleName('33 days, 8 hours');
  await expect(plaque).toContainText('3 tanks');
  await expect(page.getByTestId('home-raid-late')).toHaveAttribute('aria-pressed', 'true');
  await expect(page.getByTestId('home-raid-in')).toHaveAttribute('aria-pressed', 'false');
  await expect(page.getByTestId('home-raid-prep').getByRole('listitem')).toHaveCount(2);
  const feed = page.getByTestId('home-signed-in-root').getByTestId('home-feed');
  await expect(feed).not.toContainText('posted Onyxia');
  await expect(feed.getByTestId('home-feed-row').first()).toContainText('4 min ago');
  await expect(feed.getByTestId('home-feed-row').filter({ hasText: 'Thalrin' })).toContainText(
    'simmed 987.1 DPS, #2 of 14',
  );
  await expect(page.getByTestId('home-guild-sim')).toHaveCount(6);
  await expect(page.getByTestId('home-guild-sim').last()).toContainText('Thandrel');
  await expect(page.getByTestId('home-tier-line')).toContainText('Tier S');
  await expect(page.getByTestId('home-guild-sims')).not.toContainText('Guild page');
});

test("the plaque signs up in place: I'm in saves and presses", async ({ page }) => {
  await openHome(page, SIXTY);
  let body: unknown = null;
  const raid = (homeEnvelope('guild-home-60.json').json.data as { next_raid: object }).next_raid;
  await page.route('**/v1/guilds/2/raids/41/signups/**', async (route) => {
    body = route.request().postDataJSON();
    await route.fulfill({ json: { ok: true, data: { raid, moved: null }, error: null, request_id: 'home' } });
  });
  await page.getByTestId('home-raid-in').click();
  await expect(page.getByTestId('home-raid-in')).toHaveAttribute('aria-pressed', 'true');
  await expect(page.getByTestId('home-raid-late')).toHaveAttribute('aria-pressed', 'false');
  expect(body).toMatchObject({ state: 'in', role: 'dps' });
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
  await expect(page.getByTestId('home-upgrade')).toHaveCount(5);
  await page.unroute('**/v1/feed?**');
  await page.route('**/v1/feed?**', (route) => route.fulfill(homeEnvelope('feed-34.json')));
  await error.getByRole('button', { name: 'Try again' }).click();
  await expect(root.getByTestId('home-feed-row')).toHaveCount(7);
});

test('signed in with no character: check again, paste, the addon, and the classes', async ({ page }) => {
  await openHome(page, { me: 'me-new.json' });
  await expect(page.getByTestId('home-signed-in-root').getByRole('heading', { level: 1 })).toHaveText(
    'Get your character in',
  );
  await expect(page.getByTestId('home-get-in')).toContainText('Battle.net lists a character');
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
