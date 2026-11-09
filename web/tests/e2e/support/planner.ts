// web/tests/e2e/support/planner.ts
// Moves shared by the planner specs (spec 2026-11-08), so the phone layout and the page's
// controls are spelled out in one place: the ControlBar's class crests and race select (one
// line that opens on a phone), the head's figures, the tree tabs below 768.
import { expect, type Locator, type Page } from '@playwright/test';
import { MAX_POINTS } from '../../../src/lib/planner/types';
import { choose } from './select';

/** Tailwind's `md`. Below it the planner shows one tree at a time behind the tree tabs. */
const MD_BREAKPOINT = 768;

const isPhone = (page: Page): boolean => (page.viewportSize()?.width ?? MD_BREAKPOINT) < MD_BREAKPOINT;

/**
 * The fixture warrior one point short of a whole build: Arms full (24), Fury at 26 with
 * Flurry (talent 2007) empty. A spec that wants a whole build starts here and spends the last
 * point with `finishBuild`.
 */
export const NEARLY_FINISHED_BUILD = `/planner?code=${encodeURIComponent(
  'FS1:1.15.9.69722:warrior:human:3535125/5555510/0:',
)}`;
export const LAST_TALENT = 'talent-2007';

/** The fixture warrior's trees, by their tab index (`tree-tab-<index>`). */
const TREE_TAB = { Arms: 0, Fury: 1 } as const;
export type FixtureTree = keyof typeof TREE_TAB;

/**
 * A tree's panel by its name (`role="tabpanel"`, named by the tree). Below 768 only the open
 * tree's panel is shown; `includeHidden` finds the others in the document.
 */
export function treePanel(page: Page, name: string, includeHidden = false): Locator {
  return page.getByRole('tabpanel', { name, exact: true, includeHidden });
}

/** Below md the planner shows one tree at a time, so a talent's tree is brought up first. */
export async function showTree(page: Page, tree: FixtureTree): Promise<void> {
  if (isPhone(page)) await page.getByTestId(`tree-tab-${TREE_TAB[tree]}`).click();
}

/** `LAST_TALENT` is in the second tree. */
export async function showLastTalent(page: Page): Promise<void> {
  await showTree(page, 'Fury');
  await expect(page.getByTestId(LAST_TALENT)).toBeVisible();
}

/** A ui/Figure's number (its testid also holds the label above it). */
export function figureValue(page: Page, testid: string): Locator {
  return page.getByTestId(testid).locator('span.tabular-nums');
}

/** The head's points spent: its figure reads `51` then its unit, `of 51 points`. */
export function pointsSpent(page: Page): Locator {
  return figureValue(page, 'planner-points');
}

/** `pointsSpent` reading `spent` points (the number, then the unit glued after it). */
export function spentIs(spent: number): RegExp {
  return new RegExp(`^${spent}of \\d+ points$`);
}

/** The line under the points: the points a character has planned past its level. */
export function plannedLine(page: Page): Locator {
  return page.getByTestId('planner-spent-line');
}

/** Spends the last point: the build is whole. */
export async function finishBuild(page: Page): Promise<void> {
  await showLastTalent(page);
  await page.getByTestId(LAST_TALENT).click();
  await expect(pointsSpent(page)).toHaveText(spentIs(MAX_POINTS));
}

/**
 * Waits until the planner island has hydrated. Before it does, a class crest is a plain
 * `?class=` link (a click reloads the page) and the phone bar's toggle has no handler: Astro
 * marks an island not yet hydrated with `ssr` (support/hydrated.ts).
 */
export async function plannerHydrated(page: Page): Promise<void> {
  await expect(page.getByTestId('planner-bar')).toBeAttached();
  await expect(page.locator('astro-island[ssr]:has([data-testid="planner-bar"])')).toHaveCount(0);
}

/** Opens the ControlBar on a phone, where it is one line until "Change" is pressed. */
export async function openPlannerBar(page: Page): Promise<void> {
  await plannerHydrated(page);
  const toggle = page.getByTestId('planner-bar').locator('[data-control-bar-toggle]');
  if (!(await toggle.isVisible())) return;
  if ((await toggle.getAttribute('aria-expanded')) !== 'true') await toggle.click();
  await expect(toggle).toHaveAttribute('aria-expanded', 'true');
}

/** The class crest radio for `slug` in the ControlBar. */
export function classRadio(page: Page, slug: string): Locator {
  return page.getByTestId(`control-class-${slug}`);
}

/** Picks a class from the ControlBar's crests (opening the bar first on a phone). */
export async function chooseClass(page: Page, slug: string): Promise<void> {
  await openPlannerBar(page);
  await classRadio(page, slug).click();
}

/** The class the bar has checked. */
export async function expectClass(page: Page, slug: string): Promise<void> {
  await expect(classRadio(page, slug)).toHaveAttribute('aria-checked', 'true');
}

/** The ControlBar's race field: its native select, the source of truth for its value. */
export function raceSelect(page: Page): Locator {
  return page.locator('select[data-testid="control-race"]');
}

/** Picks a race from the ControlBar (opening the bar first on a phone). */
export async function chooseRace(page: Page, slug: string): Promise<void> {
  await openPlannerBar(page);
  await choose(raceSelect(page), slug);
}

/**
 * Share, all the way to a saved build. Task 11: the Share button (`share-open`) no longer
 * saves on its own click -- it opens a confirm step naming what becomes public, and only
 * "Share anyway" (`share-confirm-proceed`) actually calls `saveBuild`. Every planner spec
 * that lands on a saved link goes through here, so the confirm step is exercised (not
 * bypassed) by every one of them rather than each spec growing its own two-click copy.
 */
export async function shareBuild(page: Page): Promise<void> {
  await page.getByTestId('share-open').click();
  await expect(page.getByTestId('share-confirm')).toBeVisible();
  await page.getByTestId('share-confirm-proceed').click();
}
