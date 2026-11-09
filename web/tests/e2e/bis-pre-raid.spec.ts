// web/tests/e2e/bis-pre-raid.spec.ts
// The level 60 Pre-raid toggle (lib/bis/pre-raid.ts, owner 2026-11-08): it shows by itself once a
// level 60 pick comes from a raid or PvP, re-picks those slots from their by_kind entries, and has
// its own URL, /bis/<class>/<spec>/60/pre-raid, canonical to /60 while the plain list is pre-raid.
// The page's full data is served with its level 60 head pick made a raid drop (or every pick made
// a dungeon drop), so the cases hold whatever tonight's ranking holds.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { openBis } from './support/bis-page';

const SPEC_PATH = '/bis/warrior/fury';

interface Row {
  slot: string;
  filter: { item_id: number; source_kind: string; kinds?: string[]; by_kind?: Record<string, unknown> };
}
interface FullData {
  models: Record<string, { band: number; rows: Row[] }>;
}

/** Serves the spec's full data with every level 60 row passed through `change`. */
async function serveSixty(page: Page, change: (row: Row) => void): Promise<void> {
  await page.route(`**${SPEC_PATH}.json`, async (route) => {
    const response = await route.fetch();
    const data = (await response.json()) as FullData;
    for (const model of Object.values(data.models)) {
      if (model.band === 60) model.rows.forEach(change);
    }
    await route.fulfill({ response, json: data });
  });
}

const raidHead = (row: Row): void => {
  if (row.slot !== 'head') return;
  row.filter.source_kind = 'raid';
  row.filter.kinds = ['raid'];
};
const noRaid = (row: Row): void => {
  if (row.filter.source_kind === 'raid' || row.filter.source_kind === 'pvp') {
    row.filter.source_kind = 'dungeon';
    row.filter.kinds = ['dungeon'];
  }
};

test('a level 60 list with a raid pick offers Pre-raid only, which moves to its own page and back', async ({
  page,
}) => {
  await serveSixty(page, raidHead);
  await openBis(page, `${SPEC_PATH}/60`);
  const toggle = page.getByTestId('bis-pre-raid-toggle');
  await expect(toggle).toHaveText('Pre-raid only ›');
  await expect(toggle).toHaveAttribute('href', `${SPEC_PATH}/60/pre-raid`);
  await toggle.click();
  await expect(page).toHaveURL(new RegExp(`${SPEC_PATH}/60/pre-raid$`));
  await expect(page.locator('.pdf-frame .pdf-meta')).toHaveText('Pre-raid 60 · Alliance');
  await expect(toggle).toHaveText('Include raid and PvP gear ›');
  await toggle.click();
  await expect(page).toHaveURL(new RegExp(`${SPEC_PATH}/60$`));
});

test('a level 60 list with no raid or PvP pick has no toggle: it already is the pre-raid list', async ({
  page,
}) => {
  await serveSixty(page, noRaid);
  await openBis(page, `${SPEC_PATH}/60`);
  await expect(page.getByTestId('bis-band-alliance-60')).toBeVisible();
  await expect(page.getByTestId('bis-pre-raid-toggle')).toHaveCount(0);
});

test.describe('without JavaScript, as a crawler fetches', () => {
  test.use({ javaScriptEnabled: false });

  test('/60/pre-raid has its own title and breadcrumb, server-rendered on the pre-raid filter', async ({
    page,
  }) => {
    await page.goto(`${SPEC_PATH}/60/pre-raid`);
    await expect(page).toHaveTitle(/^Fury Warrior Pre-Raid BiS, WoW Forever/);
    await expect(page.locator('.pdf-frame .pdf-meta')).toHaveText('Pre-raid 60 · Alliance');
    const crumbs = JSON.parse(
      (await page.locator('script[type="application/ld+json"]').first().textContent()) ?? '{}',
    ) as { itemListElement: { name: string; item: string }[] };
    expect(crumbs.itemListElement.map((crumb) => crumb.name)).toEqual([
      'Home',
      'Best in slot',
      'Fury Warrior BiS',
      'Fury Warrior BiS, Level 60',
      'Fury Warrior Pre-Raid BiS',
    ]);
    expect(crumbs.itemListElement.at(-1)?.item).toBe(`https://foreversixty.gg${SPEC_PATH}/60/pre-raid`);
  });

  test('its canonical is /60 while the plain list is pre-raid, its own URL once it is not', async ({
    page,
  }) => {
    await page.goto(`${SPEC_PATH}/60/pre-raid`);
    const canonical = await page.locator('link[rel="canonical"]').getAttribute('href');
    await page.goto(`${SPEC_PATH}/60`);
    const plainIsPreRaid = /Pre-Raid/.test(await page.title());
    expect(canonical).toBe(`https://foreversixty.gg${SPEC_PATH}/60${plainIsPreRaid ? '' : '/pre-raid'}`);
  });
});
