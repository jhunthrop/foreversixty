// Talents start at level 10: a character below it has no point to spend, so the head says
// "0" points left and "0 of 0 spent at level 9" instead of the 51 a finished build holds. A
// level-21 character has earned 12 of them. Bug B2 (spec 2026-11-08 decisions P3, P4): the
// points a character may spend are its level's, never stretched to fit a longer build; at the
// cap the unspent cells grey and say why, and a further click plans a point ahead, marked as
// such on the cell, on the spent line and in Point order.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { plannerCopy } from '../../src/lib/planner/copy';
import { plannedLine, pointsSpent, spentIs } from './support/planner';

import { meAddonFixture } from '../../src/fixtures/me-addon';

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

function me(level: number) {
  const own = meAddonFixture.characters[0];
  return {
    ...meAddonFixture,
    main_character_key: 'us/pvp/obnoxious-yell',
    characters: [
      {
        ...own,
        key: 'us/pvp/obnoxious-yell',
        region: 'us',
        ruleset: 'pvp',
        name: 'Obnoxious Yell',
        class: 'Warrior',
        spec: undefined,
        level,
        build: { ...own.build, talents: { trees: ['0', '0', '0'], points: [0, 0, 0] } },
      },
    ],
  };
}

async function openAs(page: Page, level: number) {
  await page.route('**/v1/me', (route) => route.fulfill(ok(me(level))));
  await page.route('**/v1/devices', (route) => route.fulfill(ok([])));
  await page.goto('/planner');
}

test.describe('the planner counts only the points a character has earned', () => {
  test.skip(({ isMobile }) => isMobile, 'The head is asserted at desktop width');
  test.use({ viewport: { width: 1280, height: 900 } });

  test('a level 9 character has no points', async ({ page }) => {
    await openAs(page, 9);
    await expect(page.getByTestId('planner-meta')).toContainText('Level 9');
    await expect(pointsSpent(page)).toHaveText(spentIs(0));
    await expect(page.getByTestId('planner-answer')).toContainText(plannerCopy.pointsOfUnit(0));
  });

  test('a level 21 character has twelve', async ({ page }) => {
    await openAs(page, 21);
    await expect(page.getByTestId('planner-meta')).toContainText('Level 21');
    await expect(pointsSpent(page)).toHaveText(spentIs(0));
    await expect(page.getByTestId('planner-answer')).toContainText(plannerCopy.pointsOfUnit(12));
  });

  test('at the level cap the rest grey out, and a further point is planned ahead', async ({ page }) => {
    await openAs(page, 21);
    await expect(page.getByTestId('planner-answer')).toContainText(plannerCopy.pointsOfUnit(12));
    // Twelve legal points in Arms: Improved Heroic Strike 3, Deflection 5, Improved Rend 3 (tier 1
    // opens at 5) and Tactical Mastery 1 (needs Deflection 2).
    for (const [id, ranks] of [
      [1001, 3],
      [1002, 5],
      [1003, 3],
      [1004, 1],
    ] as const) {
      for (let i = 0; i < ranks; i += 1) await page.getByTestId(`talent-${id}`).click();
    }
    await expect(pointsSpent(page)).toHaveText(spentIs(12));
    await expect(plannedLine(page)).toHaveCount(0);

    // P3: an open cell with no point left to give reads locked, and its tooltip says why.
    const impale = page.getByTestId('talent-1006');
    await expect(impale).toHaveAttribute('data-state', 'lock');
    await impale.hover();
    await expect(page.getByTestId('talent-tip-1006')).toContainText(plannerCopy.noPointsLeft(21));

    // P4: a click still spends it, as a point planned past the character's level (the cell is
    // aria-disabled, so `force` clicks it as a player would, past Playwright's enabled check).
    await impale.click({ force: true });
    await expect(impale).toHaveAttribute('data-rank', '1');
    await expect(impale).toHaveAccessibleName(`Impale, rank 1 of 2, ${plannerCopy.plannedCell(21)}`);
    await expect(pointsSpent(page)).toHaveText(spentIs(12));
    await expect(plannedLine(page)).toHaveText(plannerCopy.plannedAhead(1));
    // Point order dims it under the gold line naming the character's level.
    await expect(page.getByTestId('order-point-11')).toHaveAttribute('data-kind', 'spent');
    await expect(page.getByTestId('order-point-12')).toHaveAttribute('data-kind', 'planned');
    await expect(page.getByTestId('planner-point-order')).toContainText(
      plannerCopy.levelLine('Obnoxious Yell', 21),
    );
  });
});
