// web/tests/e2e/support/sim-page.ts
// Driving /sim on the flat primitives (spec 2026-11-06-sim §4, §5.2): the paste Plate, the run
// line (ControlBar `sim-settings`, which folds behind its "Change" toggle on phone) and the More
// settings sheet. A spec about something else (a result, a request, the tool tabs) reaches the
// page's controls through these, so it holds at every width.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

/** The fixture's level-33 Orc Fury Warrior as an addon export: needs no API stub. */
export const SIM_FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

/** Opens a /sim URL and waits for the island to mount over the static landing: the paste box
 *  before that is the server-rendered form, and the mount replaces it. */
export async function gotoSim(page: Page, url = '/sim'): Promise<void> {
  await page.goto(url);
  await expect(page.getByTestId('sim-view')).toBeAttached();
}

/**
 * A result's secondary section (settings, rotation, resource, buffs) in view and open. Below
 * 768 px those start folded to their title (LazySection `fold`), so on a phone this presses the
 * fold first; on a desktop the section is already there.
 */
export async function openSimSection(page: Page, id: string): Promise<void> {
  const section = page.getByTestId(`sim-section-${id}`);
  const fold = page.getByTestId(`sim-section-${id}-fold`);
  await expect(section.or(fold).first()).toBeVisible({ timeout: 15_000 });
  if (await fold.isVisible()) await fold.click();
  await section.scrollIntoViewIfNeeded();
}

/** The loaded character's head: its meta line exists only once a character is up. */
export function simCharacter(page: Page) {
  return page.getByTestId('sim-head-meta');
}

/** Pastes an export into the landing's Plate and presses its Run, which loads and runs it; with
 *  `settle` (the default) waits for that first run to finish ("Run again"). */
export async function pasteAndRun(
  page: Page,
  code: string = SIM_FURY,
  { settle = true, url = '/sim' }: { settle?: boolean; url?: string } = {},
): Promise<void> {
  await gotoSim(page, url);
  await page.getByTestId('sim-cold-paste-input').fill(code);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(simCharacter(page)).toBeVisible();
  if (settle) await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 15_000 });
}

/** Opens the run line's groups when they are folded behind "Change" (phone). */
export async function openRunLine(page: Page): Promise<void> {
  // The quick sim's run line, or a tool page's (BulkRunBar on the same SimRunBar).
  const toggle = page
    .locator('[data-testid="sim-settings"], [data-testid="sim-run-bulk-bar"]')
    .locator('[data-control-bar-toggle]');
  if ((await toggle.isVisible()) && (await toggle.getAttribute('aria-expanded')) !== 'true') {
    await toggle.click();
  }
}

/** Folds the run line back to its one line when a phone step left it open (a no-op at 1440). */
export async function closeRunLine(page: Page): Promise<void> {
  const toggle = page
    .locator('[data-testid="sim-settings"], [data-testid="sim-run-bulk-bar"]')
    .locator('[data-control-bar-toggle]');
  if ((await toggle.isVisible()) && (await toggle.getAttribute('aria-expanded')) === 'true') {
    await toggle.click();
  }
}

/** Picks a Buffs preset on the run line's segmented control. */
export async function chooseBuffs(page: Page, preset: 'raid-buffed' | 'bare' | 'custom'): Promise<void> {
  await openRunLine(page);
  await page.getByTestId(`control-buffs-${preset}`).click();
  await expect(page.getByTestId(`control-buffs-${preset}`)).toHaveAttribute('aria-checked', 'true');
}

/** Opens the More settings sheet. */
export async function openMoreSettings(page: Page): Promise<void> {
  await openRunLine(page);
  await page.getByTestId('sim-settings-more').click();
  await expect(page.getByTestId('sim-more-settings')).toBeVisible();
}

/** Opens More settings and the request JSON inside it (RequestDrawer, its own <details>). */
export async function openRequest(page: Page): Promise<void> {
  if (!(await page.getByTestId('sim-more-settings').isVisible())) await openMoreSettings(page);
  if ((await page.getByTestId('sim-request-drawer').count()) === 0) {
    await page.getByTestId('sim-request-toggle').click();
  }
  const drawer = page.getByTestId('sim-request-drawer');
  if ((await drawer.getAttribute('open')) === null) await page.getByTestId('sim-request-summary').click();
  await expect(page.getByTestId('sim-request-json')).toBeVisible();
}

/** Closes the More settings sheet with Done. */
export async function closeMoreSettings(page: Page): Promise<void> {
  await page.getByTestId('sim-more-done').click();
  await expect(page.getByTestId('sim-more-settings')).toHaveCount(0);
}

/**
 * A prerendered saved run (/sim/simfixtureab, /sim/simfullfixt1) paints the page in its HTML and
 * the island mounts the same page over it: wait for the island before driving it, or a click lands
 * on markup that is about to be replaced.
 */
export async function gotoSavedSim(page: Page, path: string): Promise<void> {
  await page.goto(path);
  await expect(page.getByTestId('sim-view')).toBeVisible();
}

/** The sim's Gear doll (SimGearSection on the shared PaperdollFrame): the frame from 423 px (the
 *  desktop project), its list form under that (the mobile project). */
export function simGearDoll(page: Page) {
  return page.getByTestId('sim-gear-doll');
}

/** The doll's slot in the form the width shows: the frame's socket or the list's row. */
function shownGearSlot(page: Page, slot: string) {
  return simGearDoll(page).locator(
    `[data-testid="paperdoll-frame-slot-${slot}"]:visible, [data-testid="paperdoll-list-slot-${slot}"]:visible`,
  );
}

const escapeRegExp = (text: string): string => text.replaceAll(/[.*+?^${}()|[\]\\]/g, '\\$&');

/** Waits for the slot to show, and with `item` reads its name where the form carries it: the
 *  socket's accessible name ("Head: Lionheart Helm") or the list row's text. */
export async function expectSimGearSlot(page: Page, slot: string, item?: string): Promise<void> {
  const shown = shownGearSlot(page, slot);
  await expect(shown).toBeVisible();
  if (item === undefined) return;
  const inFrame = ((await shown.getAttribute('data-testid')) ?? '').startsWith('paperdoll-frame-slot-');
  if (inFrame) await expect(shown).toHaveAccessibleName(new RegExp(`: ${escapeRegExp(item)}`));
  else await expect(shown).toContainText(item);
}
