// web/tests/e2e/coach-report.spec.ts
// The coach on the report page (design/specs/2026-10-28-log-coach.md sections 3.1, 3.2, 4.2, 9),
// API stubbed with the worked-example packs in support/coach-stubs.ts. Every asserted sentence is
// read from lib/coach/copy.ts, never typed.
import { expect, test, type Page } from '@playwright/test';
import { coachCopy, fill } from '../../src/lib/coach/copy';
import {
  ANCHOR_LANDS,
  ANCHORS,
  GUID,
  LINK_SENTENCE,
  PANEL_URL,
  ROWS_URL,
  STATES,
  routeCoach,
  routeRatingKeys,
  signInAs,
  workedReview,
  workedRules,
} from './support/coach-stubs';
import { gotoHydrated } from './support/hydrated';
import { drainShifts, recordShifts, sumShifts } from './support/layout-shift';
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';

const MAX_PANEL_CLS = 0.02;
const PHONE = { width: 360, height: 800 };
const FORBIDDEN_TEXT = /!|heal more/i;
const leadOf = (template: string): string => template.split('{title}')[0]!.trim();
const panel = (page: Page) => page.getByTestId('coach-panel');

async function openPanel(page: Page, url = PANEL_URL()): Promise<void> {
  await gotoHydrated(page, url, 'coach-panel');
}

async function labelY(page: Page, testId: string): Promise<number> {
  return page
    .getByTestId(testId)
    .locator('h3')
    .first()
    .evaluate((node) => node.getBoundingClientRect().top + window.scrollY);
}

async function cardBottom(page: Page): Promise<number> {
  return page
    .getByTestId('coach-panel')
    .locator('> section')
    .first()
    .evaluate((node) => node.getBoundingClientRect().bottom + window.scrollY);
}

test.describe('free owner', () => {
  test.beforeEach(async ({ page }) => {
    await signInAs(page, 'free');
    await routeCoach(page, STATES['teaser-two']);
  });

  test('the row line reads the free teaser and links to /premium', async ({ page }) => {
    await page.goto(ROWS_URL);
    const line = page.getByTestId('coach-row-line');
    await expect(line).toContainText('2 things to fix');
    await expect(line).toContainText('Bloodthirst: 9 casts lost');
    await expect(page.getByTestId('coach-row-premium')).toHaveAttribute('href', '/premium');
    await expect(page.getByTestId('coach-row-premium')).toHaveText(coachCopy.upgrade);
  });

  test('opening the panel shows the teaser only', async ({ page }) => {
    await page.goto(ROWS_URL);
    await page.getByTestId('coach-row-open').click();
    await expect(page).toHaveURL(/tab=coach/);
    await expect(page.getByTestId('coach-teaser-count')).toHaveText(fill(coachCopy.teaserCount, { n: 2 }));
    await expect(page.getByTestId('coach-teaser-first')).toContainText('Bloodthirst: 9 casts lost');
    await expect(page.getByTestId('coach-upgrade')).toHaveAttribute('href', '/premium');
    await expect(page.getByTestId('coach-review-body')).toHaveCount(0);
  });
});

test.describe('Premium owner', () => {
  test.beforeEach(async ({ page }) => {
    await signInAs(page, 'premium');
  });

  test('the row line is the Fix first title, read-only, and opens the Coach tab', async ({ page }) => {
    const seen = await routeCoach(page, STATES['ready-model']);
    await page.goto(ROWS_URL);
    const line = page.getByTestId('coach-row-line');
    await expect(line).toContainText(leadOf(coachCopy.rowPremium));
    await expect(line).toContainText('Bloodthirst: 9 casts lost');
    expect(seen.panels).toHaveLength(0);
    await line.click();
    await expect(page).toHaveURL(/tab=coach/);
    await expect(page.getByTestId('coach-fix-title')).toContainText('Bloodthirst: 9 casts lost');
  });

  test('the model review renders every region with the pack titles', async ({ page }) => {
    await routeCoach(page, STATES['ready-model']);
    await openPanel(page);
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleModel, { Spec: 'Fury' }));
    await expect(page.getByTestId('coach-fix-first')).toContainText(coachCopy.labelFixFirst);
    await expect(page.getByTestId('coach-fix-first')).toContainText(coachCopy.labelPractice);
    await expect(page.getByTestId('coach-then')).toContainText(coachCopy.labelThen);
    await expect(page.getByTestId('coach-then-item')).toHaveCount(3);
    await expect(page.getByTestId('coach-well')).toContainText(coachCopy.labelWell);
    await expect(page.getByTestId('coach-cannot')).toContainText(coachCopy.labelCannot);
    await expect(page.getByTestId('coach-footer')).toContainText('Reviewed');
    await expect(page.getByTestId('coach-history')).toHaveText(coachCopy.history);
    const text = (await page.getByTestId('coach-tab').textContent()) ?? '';
    expect(text).not.toMatch(FORBIDDEN_TEXT);
  });

  test('every anchor chip lands on its tab', async ({ page }) => {
    await routeCoach(page, STATES['ready-model']);
    await openPanel(page);
    const hrefs = await page
      .getByTestId('coach-chips')
      .locator('a')
      .evaluateAll((nodes) => nodes.map((node) => node.getAttribute('href') ?? ''));
    expect(hrefs.length).toBeGreaterThanOrEqual(Object.keys(ANCHORS).length - 1);
    // The worked review is a warrior's, so it cites every anchor except the healing one (the
    // healer test below covers that chip).
    const cited = Object.entries(ANCHOR_LANDS).filter(([id]) => id !== 'healing');
    for (const [id, lands] of cited) {
      const matches = hrefs.filter((href) => lands.test(href));
      expect(matches.length, `anchor ${id} has a chip`).toBeGreaterThan(0);
    }
    await page.getByTestId('coach-chips').first().locator('a').first().click();
    await expect(page).toHaveURL(/tab=casts/);
    await expect(panel(page)).toHaveCount(0);
  });

  test('the rule-only review claims no persona and prints the link line word for word', async ({ page }) => {
    await routeCoach(page, STATES['ready-rules']);
    await openPanel(page);
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleRules, { Spec: 'Fury' }));
    await expect(page.getByTestId('coach-link-lines').first()).toContainText(LINK_SENTENCE);
  });

  test('writing swaps to ready in place: the labels and the card bottom do not move', async ({ page }) => {
    test.slow();
    let ready = false;
    await routeCoach(page, { panel: () => (ready ? workedReview() : workedRules({ status: 'writing' })) });
    await recordShifts(page);
    await page.goto(PANEL_URL());
    await expect(page.getByTestId('coach-status')).toContainText(coachCopy.writing);
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleRules, { Spec: 'Fury' }));
    await drainShifts(page);
    const before = {
      then: await labelY(page, 'coach-then'),
      well: await labelY(page, 'coach-well'),
      bottom: await cardBottom(page),
    };
    ready = true;
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleModel, { Spec: 'Fury' }), {
      timeout: 15_000,
    });
    const after = {
      then: await labelY(page, 'coach-then'),
      well: await labelY(page, 'coach-well'),
      bottom: await cardBottom(page),
    };
    expect(after).toEqual(before);
    expect(sumShifts(await drainShifts(page))).toBeLessThanOrEqual(MAX_PANEL_CLS);
  });

  test('busy shows the rule-only review and the busy line', async ({ page }) => {
    await routeCoach(page, STATES.busy);
    await openPanel(page);
    await expect(page.getByTestId('coach-busy')).toHaveText(coachCopy.busy);
    await expect(page.getByTestId('coach-fix-first')).toBeVisible();
  });

  test('a starter persona shows its hint and early-version title', async ({ page }) => {
    await routeCoach(page, STATES.starter);
    await openPanel(page);
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleStarter, { Spec: 'Fury' }));
    await expect(page.getByTestId('coach-starter-hint')).toHaveText(
      fill(coachCopy.starterHint, { Spec: 'Fury' }),
    );
  });

  test('the loading, preparing, not reviewable, error and not available states carry their copy', async ({
    page,
  }) => {
    let release: () => void = () => undefined;
    const gate = new Promise<void>((resolve) => {
      release = resolve;
    });
    await routeCoach(page, { ...STATES['ready-model'], gate });
    await page.goto(PANEL_URL());
    await expect(page.getByTestId('coach-status')).toContainText(coachCopy.loading);
    release();
    await expect(page.getByTestId('coach-fix-first')).toBeVisible();

    const cases: [keyof typeof STATES, string][] = [
      ['preparing', coachCopy.preparing.split('{Spec}')[1]!.split('.')[0]!],
      ['not-reviewable', coachCopy.notReviewable],
      ['error', coachCopy.error],
      ['not-available', coachCopy.notAvailable],
    ];
    for (const [state, sentence] of cases) {
      await page.unroute(/\/v1\/reports\/fixture2abcd\/fights\/3\/coach/);
      await routeCoach(page, STATES[state]);
      await page.goto(PANEL_URL());
      await expect(page.getByTestId('coach-panel')).toContainText(sentence);
    }
  });

  test('a healer reads the healing anchor', async ({ page }) => {
    await signInAs(page, 'premium', ['sunwick']);
    await routeCoach(page, STATES.healer);
    await openPanel(page, PANEL_URL(GUID.sunwick));
    await expect(page.getByTestId('coach-fix-title')).toContainText('Heal: 38% overheal');
    await expect(page.getByTestId('coach-chips').first().locator('a')).toHaveAttribute('href', /tab=healing/);
  });

  test('whole night: one muted line, no request', async ({ page }) => {
    const seen = await routeCoach(page, STATES['ready-model']);
    await page.goto(PANEL_URL().replace('fight=3', 'fight=0'));
    await expect(page.getByTestId('coach-tab-night')).toHaveText(coachCopy.night);
    expect(seen.panels).toHaveLength(0);
  });

  test('axe finds nothing, and no control is under 44 px at 360', async ({ page }) => {
    await page.setViewportSize(PHONE);
    await routeCoach(page, STATES['ready-model']);
    await openPanel(page);
    await expectNoAxeViolations(page, '[data-testid="coach-panel"]');
    expect(await shortControls(panel(page))).toEqual([]);
    await page.goto(ROWS_URL);
    await expect(page.getByTestId('coach-row-line')).toBeVisible();
    expect(await shortControls(page.getByTestId('coach-row-line'))).toEqual([]);
  });
});

test.describe('officer', () => {
  test('a member row carries the Coach review button, which opens the officer view', async ({ page }) => {
    await signInAs(page, 'officer');
    await routeRatingKeys(page);
    await routeCoach(page, STATES.officer);
    await page.goto(ROWS_URL);
    await expect(page.getByTestId('coach-row-line')).toHaveCount(0);
    await page.getByTestId(`actor-${GUID.thalgrit}`).getByRole('button').first().click();
    await page.getByTestId('coach-officer-open').click();
    await expect(page).toHaveURL(/tab=coach/);
    await expect(page.getByTestId('coach-officer-pill')).toHaveText(coachCopy.officerPill);
    await expect(page.getByTestId('coach-officer-note')).toBeVisible();
    await expect(page.getByTestId('coach-history')).toHaveCount(0);
  });
});

test.describe('strangers', () => {
  for (const kind of ['stranger', 'signedOut'] as const) {
    test(`${kind}: no coach line and no Coach tab entry`, async ({ page }) => {
      await signInAs(page, kind);
      const seen = await routeCoach(page, STATES['ready-model']);
      await page.goto(ROWS_URL);
      await expect(page.getByTestId('actor-table')).toBeVisible();
      await expect(page.getByTestId('coach-row-line')).toHaveCount(0);
      await expect(page.getByTestId('coach-panel')).toHaveCount(0);
      await expect(page.getByTestId('tab-coach')).toHaveCount(0);
      await expect(page.getByTestId('tab-select').locator('option[value="coach"]')).toHaveCount(0);
      expect(seen.lines).toHaveLength(0);
    });
  }

  test('signed out on a Discord link: the sign-in line', async ({ page }) => {
    await signInAs(page, 'signedOut');
    await routeCoach(page, STATES['ready-model']);
    await page.goto(PANEL_URL());
    await expect(page.getByTestId('coach-signed-out')).toHaveText(coachCopy.signedOut);
    await expect(page.getByTestId('coach-sign-in')).toHaveText(coachCopy.signIn);
  });
});
