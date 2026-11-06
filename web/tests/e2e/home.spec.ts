// web/tests/e2e/home.spec.ts
// Home (design/specs/2026-11-06-home.md): what each state offers and where it leads. The
// layout is the parity gate's (parity-home.spec.ts); this spec is the behaviour.
import { expect, test } from './support/test';
import { openHome } from './support/home-page';

test('signed out: the title, Sign in, the nine classes, one class at every band, how to start', async ({
  page,
}) => {
  await openHome(page, {});
  const out = page.getByTestId('home-signed-out');
  await expect(out.getByRole('heading', { level: 1 })).toHaveText('Play your class better.');
  await expect(page.getByTestId('home-sign-in')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=/,
  );
  const classes = page.getByTestId('home-class-picker').getByRole('link');
  await expect(classes).toHaveCount(9);
  await expect(page.getByTestId('home-class-warrior')).toHaveAttribute('href', '/guides/warrior');
  await expect(page.getByTestId('home-example-row')).not.toHaveCount(0);
  await expect(page.getByTestId('home-how-to-start').getByRole('link')).toHaveCount(2);
  await expect(page.getByTestId('home-signed-in-root')).toBeHidden();
});

test('signed in at 34: the character, the next step, three upgrades, talents and the last sim', async ({
  page,
}) => {
  await openHome(page, { me: 'me-34.json', sims: 'sims-34.json', reports: 'reports-34.json' });
  await expect(page.getByTestId('home-signed-out')).toBeHidden();
  await expect(page.getByTestId('home-head').getByRole('heading', { level: 1 })).toHaveText('Felsnap');
  await expect(page.getByTestId('home-next-step')).toContainText(
    'Next: spend your 2 talent points, then get',
  );
  await expect(page.getByTestId('home-upgrade')).toHaveCount(3);
  await expect(page.getByTestId('home-answer')).toContainText('DPS');
  await expect(page.getByTestId('home-talents').getByRole('link')).toHaveAttribute(
    'href',
    '/planner?character=us%2Fnormal%2Ffelsnap',
  );
  await expect(
    page.getByTestId('home-sim').getByRole('link', { name: /Beast Mastery Hunter/ }),
  ).toHaveAttribute('href', '/sim/k2h9f');
  await expect(page.getByTestId('home-logs')).toContainText('Upload a log');
  await expect(page.getByTestId('home-raid')).toHaveCount(0);
  await expect(page.getByTestId('home-guild')).toHaveCount(0);
});

test('signed in at 60 in a guild: the raid leads, with what to fix before it', async ({ page }) => {
  await openHome(page, {
    me: 'me-60.json',
    sims: 'sims-60.json',
    reports: 'reports-60.json',
    guild: 'guild-home-60.json',
  });
  await expect(page.getByTestId('home-next-step')).toContainText("Onyxia's Lair is on");
  await expect(page.getByTestId('home-next-step')).toContainText('you are Late');
  await expect(page.getByTestId('home-raid-needs').getByRole('img')).toHaveCount(2);
  await expect(page.getByTestId('home-raid').getByRole('link', { name: /Onyxia/ })).toHaveAttribute(
    'href',
    /\/guild\/us\/normal\/.+\/raids\/41$/,
  );
  await expect(page.getByTestId('home-talents')).toHaveCount(0);
  await expect(page.getByTestId('home-logs').getByRole('link')).toHaveCount(4);
});

test('signed in with no character: check again, paste, the addon, and the classes', async ({ page }) => {
  await openHome(page, { me: 'me-new.json' });
  await expect(page.getByTestId('home-next-step')).toContainText('No character yet.');
  await expect(page.getByTestId('home-check-again')).toBeVisible();
  await expect(page.getByTestId('home-get-in').getByRole('link')).toHaveCount(2);
  await expect(page.getByTestId('home-class-picker').getByRole('link')).toHaveCount(9);
});

test('the upgrades fail on their own and come back with Try again', async ({ page }) => {
  await openHome(page, { me: 'me-34.json', failBis: true });
  const error = page.getByTestId('home-upgrades-error');
  await expect(error).toContainText('We could not load your upgrades.');
  await page.unroute('**/data/*/bis/hunter-beast-mastery.json');
  await page.route('**/data/*/bis/hunter-beast-mastery.json', (route) =>
    route.fulfill({ path: 'tests/e2e/fixtures/home/bis-hunter-beast-mastery.json' }),
  );
  await error.getByRole('button', { name: 'Try again' }).click();
  await expect(page.getByTestId('home-upgrade')).toHaveCount(3);
});

test('a stale session hint gives the signed-out page back', async ({ page }) => {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.goto('/');
  await expect(page.getByTestId('home-signed-out')).toBeVisible();
  await expect(page.locator('html[data-home-out]')).toHaveCount(1);
});
