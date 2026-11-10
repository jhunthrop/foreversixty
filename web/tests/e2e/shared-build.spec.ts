// web/tests/e2e/shared-build.spec.ts
import { expect, test } from './support/test';

import { bandCompareCopy, plannerCharacterCardCopy, plannerCopy } from '../../src/lib/planner/copy';
import { READ_ONLY_REASON } from '../../src/lib/planner/store.svelte';
import { ACTIVE_BUILD } from './support/active-build';
import { classRadio, pointsSpent, spentIs } from './support/planner';
import { plannerAccount, routePlannerState } from './support/planner-page';

// The real /b/:id is server-rendered by the Go API and reaches the browser through the
// Worker proxy (Task 12). Playwright serves the static site only, so the page is stubbed
// with exactly the markup the interface contract specifies -- including the island script
// tag, which resolves against the preview server and loads the bundle this task builds.
const RECORD = {
  id: 'k7x2qm4a',
  class_id: 1,
  race_id: 5,
  tree_version: ACTIVE_BUILD,
  point_order: [1001, 1001, 1001],
  gear: {},
  title: 'Arms leveling',
  created_at: '2026-09-14T03:12:44Z',
  views: 12,
};

const pageFor = (build: string): string => `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Arms leveling · 3/0 · Forever Sixty</title>
<link rel="stylesheet" href="/planner-island.css"></head>
<body><main id="main">
<div id="planner" data-build='${build}' data-tree-version="${RECORD.tree_version}"></div>
<script type="module" src="/planner-island.js"></script>
</main></body></html>`;

const sharedPage = pageFor(JSON.stringify(RECORD));

test('the island bundle and its stylesheet are published at fixed paths', async ({ request }) => {
  const script = await request.get('/planner-island.js');
  expect(script.status()).toBe(200);
  expect(script.headers()['content-type']).toContain('javascript');
  expect((await script.body()).length).toBeGreaterThan(1000);

  const styles = await request.get('/planner-island.css');
  expect(styles.status()).toBe(200);
  expect(styles.headers()['content-type']).toContain('css');
});

test('a shared build opens read-only under its own title, and Edit a copy makes it editable', async ({
  page,
}) => {
  await page.route('**/b/k7x2qm4a', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: sharedPage }),
  );
  await page.goto('/b/k7x2qm4a');

  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('tree-points-161')).toHaveText('3');
  // The head (premium spec 2026-11-08 section 8): the record's title, then Shared build, the
  // planner's meta, the split and the views.
  await expect(page.getByTestId('page-title')).toHaveText('Arms leveling');
  await expect(page.getByTestId('planner-shared-meta')).toHaveText(plannerCopy.sharedMeta);
  await expect(page.getByTestId('planner-meta')).toContainText('Level 12');
  await expect(page.getByTestId('planner-split')).toHaveText('3/0');
  await expect(page.getByTestId('planner-views')).toHaveText(plannerCopy.sharedViews(12));
  // No class and race bar (decision B3); Edit a copy in the notes line (B2); In the game leads the
  // rail and the level build is the Guide build.
  await expect(page.getByTestId('planner-bar')).toHaveCount(0);
  await expect(page.getByTestId('planner-shared-note')).toContainText(plannerCopy.sharedLinkNote);
  await expect(page.getByTestId('planner-fork-button')).toHaveText(plannerCopy.fork);
  await expect(page.getByTestId('band-compare').getByRole('heading')).toHaveText(
    bandCompareCopy.sharedHeading,
  );
  await expect(page.getByTestId('planner-reset')).toHaveCount(0);

  // A read-only cell is aria-disabled (locked), but a click still reaches the planner, which
  // refuses and says why: `force` clicks it as a viewer would, past Playwright's enabled check.
  await page.getByTestId('talent-1002').click({ force: true });
  await expect(page.getByTestId('planner-status')).toHaveText(READ_ONLY_REASON);
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '0');

  await page.getByTestId('planner-fork-button').click();
  await expect(page.getByTestId('planner-shared-note')).toHaveCount(0);
  await page.getByTestId('talent-1002').click();
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '1');
  await expect(pointsSpent(page)).toHaveText(spentIs(4));
  await expect(page.getByTestId('planner-reset')).toBeVisible();
});

test('a shared build with no title is headed by its race and class', async ({ page }) => {
  const untitled = pageFor(JSON.stringify({ ...RECORD, id: 'untitled1', title: undefined }));
  await page.route('**/b/untitled1', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: untitled }),
  );
  await page.goto('/b/untitled1');
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('page-title')).toHaveText(/^\S+( \S+)* Warrior$/);
});

// Fix round 2, item 7 (wow-player review finding 3): a populated shared build -- real
// points spent, a real differs count against the build's own band -- confirming cells are
// genuinely inert for a viewer (no click spends or removes a point on someone else's link)
// rather than merely looking disabled. The fixture's hunter class pairs with its own
// published `bis/hunter-marksmanship.json` (the fixture warrior's two trees, Arms/Fury,
// have no published band file at all); 5 points in Lethal Shots (4001, tier 0, no prereq --
// legal in one click each) against that file's own all-zero band 20 string, with 4001 itself
// the one marked cell. The count is the build's own points spent where the band spends none
// (never more than the build holds), so the line reads 5 and Matches is reserved for a build
// with no marked cell at all.
const POPULATED_RECORD = {
  id: 'hunterpop1',
  class_id: 3,
  race_id: 2,
  tree_version: ACTIVE_BUILD,
  point_order: [4001, 4001, 4001, 4001, 4001],
  gear: {},
  title: 'Shared Marksmanship build',
  created_at: '2026-10-01T00:00:00Z',
  views: 4,
};
const populatedSharedPage = pageFor(JSON.stringify(POPULATED_RECORD));

test('a populated read-only build keeps every cell inert, in both directions', async ({ page }) => {
  await page.route('**/b/hunterpop1', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: populatedSharedPage }),
  );
  await page.goto('/b/hunterpop1');

  await expect(page.getByTestId('talent-4001')).toHaveAttribute('data-rank', '5');
  await expect(page.getByTestId('tree-points-302')).toHaveText('5');

  // Against the build's own band (the fixture's is empty, 0 points): all 5 points are past the
  // band's total, so none of them differ -- Matches, "+5 beyond the build" -- while the cell still
  // carries its rank-differs marker (the gold dot, TalentButton's `diff`). Load hidden (nothing
  // to overwrite into on a read-only build), Edit a copy offered instead, and no See what differs.
  await expect(page.getByTestId('band-compare-differs')).toHaveText(/^Matches the .+ guide build$/);
  await expect(page.getByTestId('band-compare-beyond')).toHaveText(bandCompareCopy.beyondGuide(5));
  await expect(page.getByTestId('talent-4001')).toHaveClass(/\bdiff\b/);
  await expect(page.getByTestId('band-compare-load')).toHaveCount(0);
  await expect(page.getByTestId('band-compare-see-differs')).toHaveCount(0);
  await expect(page.getByTestId('planner-fork-button')).toBeVisible();

  // Left-click an unspent cell: no point is added.
  await page.getByTestId('talent-4002').click({ force: true });
  await expect(page.getByTestId('talent-4002')).toHaveAttribute('data-rank', '0');
  await expect(pointsSpent(page)).toHaveText(spentIs(5));

  // Right-click an already-spent cell: no point is removed either.
  await page.getByTestId('talent-4001').click({ button: 'right' });
  await expect(page.getByTestId('talent-4001')).toHaveAttribute('data-rank', '5');
  await expect(pointsSpent(page)).toHaveText(spentIs(5));
});

// The API writes data-build itself, so malformed JSON there means the API is broken, not the
// visitor. The island logs it and mounts anyway rather than leaving a blank page: an empty
// planner someone can use beats nothing at all, and it cannot be read-only -- there is no
// build to be read-only about.
test('a build attribute that is not JSON still mounts an editable planner', async ({ page }) => {
  await page.route('**/b/broken', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: pageFor('{"class_id":') }),
  );
  await page.goto('/b/broken');

  await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '0');
  await expect(classRadio(page, 'warrior')).not.toHaveAttribute('aria-disabled', 'true');
  await expect(page.getByTestId('planner-fork-button')).toHaveCount(0);
  // Editable: a click spends a point, and Reset (hidden at 0 spent) comes with it.
  await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '1');
  await expect(page.getByTestId('planner-reset')).toBeVisible();
});

// boot() resolves the record's class and race ids through loadReference before it mounts.
// When that fetch fails it swallows the error deliberately, because Planner runs the same
// call again and has a state for it -- so the visitor gets the retry rather than a blank div.
test('a reference fetch that fails still mounts, and Planner offers the retry', async ({ page }) => {
  await page.route('**/data/*/classes.json', (route) => route.abort('failed'));
  await page.route('**/b/k7x2qm4a', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: sharedPage }),
  );
  await page.goto('/b/k7x2qm4a');

  await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();
  // The planner's own tree-load retry has its own accessible name (LoadError's
  // `retryAriaLabel`), so it stays unambiguous next to any other "Try again" on the page.
  await expect(page.getByRole('button', { name: 'Reload talents' })).toBeVisible();
});

// Ruling MUST 5 (premium spec 2026-11-08): signed in on another class, the shared page says so in
// one plain line above Copy for the game, which stays the primary; there is no disabled send.
test("another class's shared build says so in one line, and Copy for the game is the primary", async ({
  page,
}) => {
  await routePlannerState(page, 'felsnap');
  await page.route('**/b/k7x2qm4a', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: sharedPage }),
  );
  await page.goto('/b/k7x2qm4a');

  const name = plannerAccount().characters[0].name;
  await expect(page.getByTestId('planner-send-other')).toHaveText(
    plannerCharacterCardCopy.otherClass('Warrior', name, 'Hunter'),
  );
  await expect(page.getByTestId('planner-send')).toHaveCount(0);
  await expect(page.getByTestId('copy-for-game')).toBeVisible();
});
