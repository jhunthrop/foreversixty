import { expect, test } from '@playwright/test';

test('the planner opens on the default class with an empty build', async ({ page }) => {
  const errors: string[] = [];
  page.on('console', (m) => {
    if (m.type() === 'error') errors.push(m.text());
  });
  await page.goto('/planner');
  await expect(page.getByTestId('planner-level')).toHaveText('9');
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(
    page.getByText(
      'Classic Era trees shown until the beta client exports; Forever revamped talents replace them then.',
    ),
  ).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Fury' })).toBeVisible();
  expect(errors).toEqual([]);
});

test('?class and ?race preselect the build', async ({ page }) => {
  await page.goto('/planner?class=warrior&race=dwarf');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(page.getByLabel('Race')).toHaveValue('dwarf');
});

test('a failed talent fetch shows the reason and a working retry', async ({ page }) => {
  let attempts = 0;
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    attempts += 1;
    if (attempts === 1) return route.abort('failed');
    return route.continue();
  });
  await page.goto('/planner');
  await expect(page.getByText('Talent data did not load')).toBeVisible();
  await page.getByRole('button', { name: 'Retry' }).click();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
});

test('clicking a talent spends points and the counters follow', async ({ page }) => {
  await page.goto('/planner');
  const improvedHeroicStrike = page.getByTestId('talent-1001');
  await improvedHeroicStrike.click();
  await improvedHeroicStrike.click();
  await improvedHeroicStrike.click();
  await expect(improvedHeroicStrike).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('planner-spent')).toHaveText('3/51');
  await expect(page.getByTestId('planner-split')).toHaveText('3/0');
  await expect(page.getByTestId('planner-level')).toHaveText('12');
});

test('a locked tier refuses the point and says why', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1003').click();
  await expect(page.getByTestId('planner-refusal')).toHaveText('Tier 1 of Arms needs 5 points in Arms first');
  await expect(page.getByTestId('talent-1003')).toHaveAttribute('data-rank', '0');
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
});

test('a full talent refuses a fourth point with its own reason', async ({ page }) => {
  await page.goto('/planner');
  for (let i = 0; i < 4; i += 1) await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('planner-refusal')).toHaveText(
    'Improved Heroic Strike is already at 3 of 3 points',
  );
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
});

test('right-click removes the last point in a talent', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1001').click({ button: 'right' });
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '1');
});

test('removal is refused when a later point depends on it', async ({ page }) => {
  await page.goto('/planner');
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1002').click();
  await page.getByTestId('talent-1002').click();
  await page.getByTestId('talent-1003').click();
  await page.getByTestId('talent-1001').click({ button: 'right' });
  await expect(page.getByTestId('planner-refusal')).toHaveText('Tier 1 of Arms needs 5 points in Arms first');
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
});

test('the tooltip shows the current and the next rank', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1001').hover();
  const tip = page.getByRole('tooltip');
  await expect(tip).toContainText('Improved Heroic Strike');
  await expect(tip).toContainText('Rank 1 of 3');
  await expect(tip).toContainText('Reduces the rage cost of Heroic Strike by 1.');
  await expect(tip).toContainText('Reduces the rage cost of Heroic Strike by 2.');
});

test('a keyboard-only visitor can spend and remove points', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').focus();
  await page.keyboard.press('Enter');
  await page.keyboard.press('Enter');
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '2');
  await page.keyboard.press('ArrowRight');
  await expect(page.getByTestId('talent-1002')).toBeFocused();
  await page.keyboard.press('Enter');
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '1');
  await page.keyboard.press('Backspace');
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '0');
  await expect(page.getByTestId('planner-spent')).toHaveText('2/51');
});
