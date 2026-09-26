// web/tests/e2e/real-data.spec.ts
// The pre-deploy smoke against the real pipeline output. Every other spec in this directory
// runs on src/fixtures/planner (see playwright.config.ts), because they assert on talent
// names, item ids and counts that the real data moves under them on every regeneration. This
// one runs only under FOREVER_DATA=real -- `npm run test:e2e:real` -- and asserts on
// structure plus the names it reads back out of the synced data itself, never on hardcoded
// Forever facts, so regenerating data/builds cannot turn it red on its own.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';
import { collectPageErrors } from './support/console';
import { ACTIVE_BUILD } from './support/active-build';
import { shareBuild } from './support/planner';
import { POINTS_PER_TIER } from '../../src/lib/planner/types';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');

/** Classic's three trees per class -- the shape the planner lays out its columns for. */
const TREES_PER_CLASS = 3;

// The suite defaults to the fixture, so an unset FOREVER_DATA means the served build is the
// fixture and these assertions would be measuring the wrong thing.
test.skip(
  process.env.FOREVER_DATA !== 'real',
  'real-data smoke; run it with FOREVER_DATA=real (npm run test:e2e:real)',
);

function readSynced<T>(...segments: string[]): T {
  const { build } = readJson<{ build: string }>(path.join(WEB_ROOT, 'src/data/active-build.json'));
  return readJson<T>(path.join(WEB_ROOT, 'public/data', build, ...segments));
}

function readJson<T>(file: string): T {
  return JSON.parse(readFileSync(file, 'utf8')) as T;
}

/** The tree names the planner currently has laid out, in column order. */
function renderedTreeNames(page: Page): Promise<string[]> {
  return page
    .locator('[data-testid="tree-columns"] [role="grid"]')
    .evaluateAll((nodes) =>
      nodes.map((node) => (node.getAttribute('aria-label') ?? '').replace(/ talents$/, '')),
    );
}

/** The tree names the synced talent file lists for a class, in position order. */
function syncedTreeNames(slug: string): string[] {
  const talents = readSynced<{ trees: { name: string; position: number }[] }>('talents', `${slug}.json`);
  return [...talents.trees].sort((a, b) => a.position - b.position).map((tree) => tree.name);
}

/**
 * Opens /planner and waits for the reference data to land, returning the class it opened on.
 * The class <select> carries no options until `store.classes` is populated, so reading its
 * value straight after the navigation can come back empty -- the trees appearing is the
 * signal that the reference and talent fetches have both resolved.
 */
async function openPlanner(page: Page): Promise<string> {
  await page.goto('/planner');
  await expect.poll(() => renderedTreeNames(page)).toHaveLength(TREES_PER_CLASS);
  const slug = await page.getByLabel('Class').inputValue();
  expect(slug).not.toBe('');
  return slug;
}

test('the planner opens on the default class with its three real trees', async ({ page }) => {
  const errors = collectPageErrors(page);

  const slug = await openPlanner(page);
  const expected = syncedTreeNames(slug);
  expect(expected).toHaveLength(TREES_PER_CLASS);
  expect(await renderedTreeNames(page)).toEqual(expected);
  expect(errors).toEqual([]);
});

test('switching class lays out the new class real trees', async ({ page }) => {
  const from = await openPlanner(page);
  expect(await renderedTreeNames(page)).toEqual(syncedTreeNames(from));

  const classes = readSynced<{ slug: string }[]>('classes.json');
  // Tree names repeat across classes (Holy, Protection, Restoration), so the switch is only
  // observable when the two name lists differ. Picking the first class that differs keeps
  // this deterministic without naming a class the data may one day reshape.
  const to = classes
    .map((row) => row.slug)
    .find(
      (candidate) =>
        candidate !== from && syncedTreeNames(candidate).join('|') !== syncedTreeNames(from).join('|'),
    );
  expect(to, 'no second class with a distinguishable set of tree names').toBeDefined();

  await page.getByLabel('Class').selectOption(to!);
  await expect(page.getByLabel('Class')).toHaveValue(to!);
  await expect.poll(() => renderedTreeNames(page)).toEqual(syncedTreeNames(to!));
});

test('the planner offers the share panel', async ({ page }) => {
  await openPlanner(page);
  await expect(page.getByTestId('share-panel')).toBeAttached();
  await expect(page.getByRole('button', { name: 'Share' })).toBeAttached();
});

/** The synced talent file for a class, trees in position order. */
interface SyncedTalent {
  id: number;
  name: string;
  tier: number;
  column: number;
  max_rank: number;
  prereq_talent_id: number | null;
  prereq_rank: number | null;
  spell_id: number;
}
interface SyncedTree {
  id: number;
  name: string;
  position: number;
  background: string;
  talents: SyncedTalent[];
}

function syncedTrees(slug: string): SyncedTree[] {
  const file = readSynced<{ trees: SyncedTree[] }>('talents', `${slug}.json`);
  return [...file.trees].sort((a, b) => a.position - b.position);
}

/**
 * Spends `points` in a tree by always taking the leftmost cell that is open. "Open" means
 * addable -- `available` (untouched) or `filled` (already ranked, but not yet maxed) -- not
 * only untouched: the real client's trees mix in authentic multi-rank talents (Arms' tier 0
 * alone is Improved Heroic Strike/3, Deflection/5, Improved Rend/3), and a tier needs
 * POINTS_PER_TIER points spent in the tree as a whole before the next one unlocks. Clicking
 * only ever-untouched cells caps a tier at one point per talent in it, which is fewer than
 * POINTS_PER_TIER whenever a tier holds fewer talents than that -- true of every tier here --
 * so the tree would wedge itself long before 51 points. The fixture's two-tree warrior never
 * exercised this because every one of its talents happens to be single-rank.
 */
async function spend(page: Page, treeId: number, points: number): Promise<void> {
  for (let i = 0; i < points; i += 1) {
    await page
      .locator(`[data-testid="tree-${treeId}"] :is([data-state="available"], [data-state="filled"])`)
      .first()
      .click();
  }
}

test('the trees are laid out in the order the client draws them', async ({ page }) => {
  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  expect(syncedTrees('warrior').map((tree) => tree.name)).toEqual(['Arms', 'Fury', 'Protection']);
});

test('the banner names the build the trees were read from', async ({ page }) => {
  await openPlanner(page);
  await expect(
    page.getByText(`Talent trees read from the game client, build ${ACTIVE_BUILD}.`),
  ).toBeVisible();
});

test('a 31/20/0 warrior is spent and read back as 31/20/0', async ({ page }) => {
  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  const [arms, fury] = syncedTrees('warrior');

  await spend(page, arms.id, 31);
  await spend(page, fury.id, 20);

  await expect(page.getByTestId('planner-split')).toHaveText('31/20/0');
  await expect(page.getByTestId('planner-spent')).toHaveText('51/51');
  await expect(page.getByTestId('planner-remaining')).toHaveText('0');
  await expect(page.getByTestId('planner-level')).toHaveText('60');
  await expect(page.getByTestId(`tree-points-${arms.id}`)).toHaveText('31');
  await expect(page.getByTestId(`tree-points-${fury.id}`)).toHaveText('20');
});

test('a prerequisite link is drawn for every prerequisite the client has', async ({ page }) => {
  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  const pairs = syncedTrees('warrior').flatMap((tree) =>
    tree.talents.filter((talent) => talent.prereq_talent_id !== null).map((talent) => ({ tree, talent })),
  );
  expect(pairs.length).toBeGreaterThan(0);
  for (const { talent } of pairs) {
    await expect(page.getByTestId(`connector-${talent.prereq_talent_id}-${talent.id}`)).toBeAttached();
  }

  // The first one turns gold once its prerequisite is maxed. The prerequisite's own tier
  // has to be unlocked first -- POINTS_PER_TIER points spent anywhere in the tree, per
  // tier -- before its talent cell will even accept a click.
  const { tree, talent } = pairs[0];
  const prereq = tree.talents.find((t) => t.id === talent.prereq_talent_id)!;
  const link = page.getByTestId(`connector-${prereq.id}-${talent.id}`);
  await expect(link).toHaveAttribute('data-met', 'false');
  await spend(page, tree.id, POINTS_PER_TIER * prereq.tier);
  for (let i = 0; i < prereq.max_rank; i += 1) {
    await page.getByTestId(`talent-${prereq.id}`).click();
  }
  await expect(link).toHaveAttribute('data-met', 'true');
  await expect(page.getByTestId(`talent-${prereq.id}`)).toHaveAttribute('data-state', 'maxed');
});

test('every tree ships its own processed art', async ({ page }) => {
  await openPlanner(page);
  for (const slug of readSynced<{ slug: string }[]>('classes.json').map((row) => row.slug)) {
    for (const tree of syncedTrees(slug)) {
      const response = await page.request.get(`/data/${ACTIVE_BUILD}/trees/${tree.background}.webp`);
      expect(response.status(), `${slug} ${tree.name}`).toBe(200);
      expect(response.headers()['content-type']).toContain('image/webp');
    }
  }
});

test('a build shared on the new trees reopens from its link, on desktop and on a phone', async ({ page }) => {
  let body: { point_order: number[]; tree_version: string } | undefined;
  await page.route('**/v1/builds', async (route) => {
    body = route.request().postDataJSON();
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { id: 'real1234', url: 'https://foreversixty.gg/b/real1234' },
        error: null,
        request_id: 'req-1',
      }),
    });
  });

  await openPlanner(page);
  await page.getByLabel('Class').selectOption('warrior');
  await expect.poll(() => renderedTreeNames(page)).toEqual(['Arms', 'Fury', 'Protection']);
  const [arms, fury] = syncedTrees('warrior');
  await spend(page, arms.id, 31);
  await spend(page, fury.id, 20);
  await page.getByLabel('Title').fill('Arms PvE');
  await shareBuild(page);
  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/real1234');

  expect(body!.tree_version).toBe(ACTIVE_BUILD);
  expect(body!.point_order).toHaveLength(51);

  // /b/:id is rendered by the Go API in production; the static preview gets the
  // same markup the interface contract specifies, carrying the body just posted.
  const record = {
    id: 'real1234',
    class_id: 1,
    race_id: 1,
    tree_version: ACTIVE_BUILD,
    point_order: body!.point_order,
    gear: {},
    title: 'Arms PvE',
    created_at: '2026-09-17T09:00:00Z',
    views: 1,
  };
  const shared = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Arms PvE · 31/20/0 · Forever Sixty</title>
<link rel="stylesheet" href="/planner-island.css"></head>
<body><main id="main">
<div id="planner" data-build='${JSON.stringify(record)}' data-tree-version="${ACTIVE_BUILD}"></div>
<script type="module" src="/planner-island.js"></script>
</main></body></html>`;
  await page.route('**/b/real1234', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: shared }),
  );

  await page.goto('/b/real1234');
  await expect(page.getByTestId('planner-split')).toHaveText('31/20/0');
  await expect(page.getByLabel('Class')).toBeDisabled();

  await page.setViewportSize({ width: 390, height: 800 });
  await page.reload();
  await expect(page.getByTestId('planner-split')).toHaveText('31/20/0');
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
});
