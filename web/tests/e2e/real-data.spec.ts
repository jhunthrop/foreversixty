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

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');

/** Classic's three trees per class -- the shape the planner lays out its columns for. */
const TREES_PER_CLASS = 3;
/** The nine classes and nine races the reference page exists to cross. */
const CLASS_COUNT = 9;
const RACE_COUNT = 9;

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
  const errors: string[] = [];
  page.on('console', (message) => {
    if (message.type() === 'error') errors.push(message.text());
  });

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

test('the reference page crosses all nine classes with all nine races', async ({ page }) => {
  await page.goto('/classes');
  // One "Open in the planner" link per class, one table row per race.
  await expect(page.locator('a[href^="/planner?class="]:not([href*="race="])')).toHaveCount(CLASS_COUNT);
  await expect(page.locator('table tbody tr')).toHaveCount(RACE_COUNT);
  // The nine class columns plus the leading "Race" header.
  await expect(page.locator('table thead th')).toHaveCount(CLASS_COUNT + 1);
});
