// web/tests/e2e/report-duplicates.spec.ts
import { readFile } from 'node:fs/promises';
import path from 'node:path';

import { expect, test } from '@playwright/test';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');

// Real fights repeat timestamps: a multi-target spell lands two casts in one millisecond,
// a trinket and its proc apply on the same tick, a DoT tick and a cleave hit a dying player
// in the same instant. The timelines and deaths views keyed their rows on the timestamp
// alone, so the first real wipe uploaded threw a duplicate-key error and rendered nothing.
test('a fight whose events share timestamps still renders its timelines and deaths', async ({ page }) => {
  const summary = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const cast = summary.casts[0];
  cast.sequence = [...cast.sequence, cast.sequence[0], cast.sequence[0]];
  const aura = summary.auras[0];
  aura.segments = [...aura.segments, { ...aura.segments[0] }];
  const death = summary.deaths[0];
  death.last = [...death.last, { ...death.last[0] }];
  summary.deaths = [...summary.deaths, { ...death, guid: death.guid, at_ms: death.at_ms + 1 }];

  const errors: string[] = [];
  page.on('pageerror', (error) => errors.push(error.message));
  await page.route('**/fights/3/summary.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary) }),
  );

  await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
  await expect(page.locator('[data-testid^="lane-"]').first()).toBeVisible();

  await page.goto('/reports/fixture2abcd?fight=3&tab=deaths');
  await expect(page.getByTestId('deaths-tab')).toBeVisible();

  expect(errors).toEqual([]);
});
