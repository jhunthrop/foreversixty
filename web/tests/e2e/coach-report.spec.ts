// web/tests/e2e/coach-report.spec.ts
// The coach on the report page (design/specs/2026-10-28-log-coach.md sections 3.1, 3.2, 4.2, 9),
// API stubbed with the coach's own output (src/fixtures/coach, through support/coach-stubs.ts) on the
// worked Onyxia pull. Every asserted sentence is read from lib/coach/copy.ts or the fixtures.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { coachCopy, fill } from '../../src/lib/coach/copy';
import type { CoachAnchorKind, CoachFull } from '../../src/lib/coach/types';
import {
  FIX_FIRST_TITLE,
  GUID,
  LINK_SENTENCE,
  PANEL_URL,
  ROWS_URL,
  STATES,
  coachFixture,
  routeCoach,
  routeRatingKeys,
  signInAs,
  writingThenModel,
} from './support/coach-stubs';
import { gotoHydrated } from './support/hydrated';
import { drainShifts, recordShifts, sumShifts } from './support/layout-shift';
import { expectNoAxeViolations } from './support/officer-tools';
import { shortControls } from './support/short-controls';

const MAX_PANEL_CLS = 0.02;
const PHONE = { width: 360, height: 800 };
const DESKTOP = { width: 1280, height: 800 };
/** Spec 1: at 1280 x 800 the header, the opening and the whole Fix first block end above y 640. */
const FOLD_Y = 640;
/** Spec 3.2 and 7: the phone strip's reserve by entitlement. */
const STRIP_PHONE_PX = { premium: 52, free: 88 } as const;
/** Where each anchor kind's chip lands in the report. */
const LANDS: Readonly<Record<CoachAnchorKind, RegExp>> = {
  cast: /tab=casts/,
  aura: /tab=(buffs|debuffs)/,
  buffs: /tab=buffs/,
  death: /tab=deaths/,
  healing: /tab=healing/,
  time: /view=timelines/,
  resource: /view=timelines/,
};
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
    .getByTestId('coach-card')
    .evaluate((node) => node.getBoundingClientRect().bottom + window.scrollY);
}

test.describe('free owner', () => {
  test.beforeEach(async ({ page }) => {
    await signInAs(page, 'free');
    await routeCoach(page, STATES['teaser-several']);
  });

  test('the row line reads the free teaser and links to /premium', async ({ page }) => {
    await page.goto(ROWS_URL);
    await expect(page.getByTestId('coach-row-open')).toHaveText(
      fill(coachCopy.rowFree, { n: 3, title: FIX_FIRST_TITLE }).replace(` · ${coachCopy.upgrade}`, ''),
    );
    await expect(page.getByTestId('coach-row-premium')).toHaveAttribute('href', '/premium#coach');
    await expect(page.getByTestId('coach-row-premium')).toHaveText(coachCopy.upgrade);
  });

  test('opening the panel shows the teaser only', async ({ page }) => {
    await page.goto(ROWS_URL);
    await page.getByTestId('coach-row-open').click();
    await expect(page).toHaveURL(/tab=coach/);
    await expect(page.getByTestId('coach-teaser-count')).toHaveText(fill(coachCopy.teaserCount, { n: 3 }));
    await expect(page.getByTestId('coach-teaser-first')).toHaveText(
      fill(coachCopy.teaserFirst, { title: FIX_FIRST_TITLE }),
    );
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleModel, { Spec: 'Fury' }));
    await expect(page.getByTestId('coach-upgrade')).toHaveAttribute('href', '/premium#coach');
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
    await expect(line).toContainText(FIX_FIRST_TITLE);
    expect(seen.panels).toHaveLength(0);
    await line.click();
    await expect(page).toHaveURL(/tab=coach/);
    await expect(page.getByTestId('coach-fix-title')).toContainText(FIX_FIRST_TITLE);
  });

  test('the model review renders every region with the pack titles', async ({ page }) => {
    await routeCoach(page, STATES['ready-model']);
    await openPanel(page);
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleModel, { Spec: 'Fury' }));
    await expect(page.getByTestId('coach-fix-first')).toContainText(coachCopy.labelFixFirst);
    await expect(page.getByTestId('coach-fix-first')).toContainText(coachCopy.labelPractice);
    await expect(page.getByTestId('coach-then')).toContainText(coachCopy.labelThen);
    await expect(page.getByTestId('coach-then-item')).toHaveCount(
      (STATES['ready-model'].panel as CoachFull).review.thenFix.length,
    );
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
    const review = STATES['ready-model'].panel as CoachFull;
    const cited = [review.review.fixFirst, ...review.review.thenFix].flatMap((item) => item?.anchors ?? []);
    expect(hrefs).toHaveLength(cited.length);
    cited.forEach((id, index) => {
      const kind = review.anchors[id]!.kind;
      expect(hrefs[index], `anchor ${id} (${kind})`).toMatch(LANDS[kind]);
    });
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
    await routeCoach(
      page,
      writingThenModel(() => ready),
    );
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

    const reason = (state: keyof typeof STATES): string =>
      (STATES[state].panel as CoachFull).review.couldntCheck[0]!;
    const fury = (template: string): string => fill(template, { Spec: 'Fury' });
    const cases: [keyof typeof STATES, string, string][] = [
      ['preparing', fury(coachCopy.preparing), fury(coachCopy.titleModel)],
      ['not-reviewable', reason('not-reviewable'), fury(coachCopy.titleRules)],
      ['not-reviewable-short', reason('not-reviewable-short'), fury(coachCopy.titleRules)],
      ['not-reviewable-spec', reason('not-reviewable-spec'), coachCopy.titleGeneric],
      ['error', coachCopy.error, fury(coachCopy.titleModel)],
      ['not-available', coachCopy.notAvailable, coachCopy.titleGeneric],
      ['no-findings', coachCopy.rulesNothing, fury(coachCopy.titleRules)],
    ];
    for (const [state, sentence, title] of cases) {
      await page.unroute(/\/v1\/reports\/fixture2abcd\/fights\/3\/coach/);
      await routeCoach(page, STATES[state]);
      await page.goto(PANEL_URL());
      await expect(page.getByTestId('coach-panel')).toContainText(sentence);
      await expect(page.getByTestId('coach-title')).toHaveText(title);
    }
  });

  test('raid buffs only: no Fix first, the raid buff under Then', async ({ page }) => {
    const answer = STATES['raid-buffs-only'].panel as CoachFull;
    await routeCoach(page, STATES['raid-buffs-only']);
    await openPanel(page);
    await expect(page.getByTestId('coach-fix-first')).toHaveCount(0);
    await expect(page.getByTestId('coach-then-item').first()).toContainText(answer.review.thenFix[0]!.title);
  });

  test('at 1280 x 800 Fix first ends above the fold, with no chart on the tab', async ({
    page,
    isMobile,
  }) => {
    // The fold is the fine-pointer desktop's: a touch device at this width draws the report's
    // dense tool-bar fields at 44 px, not 36 (controls-charts spec §A4.1), and is not this budget.
    test.skip(isMobile, 'the fine-pointer desktop layout');
    await page.setViewportSize(DESKTOP);
    await routeCoach(page, STATES['ready-model']);
    await openPanel(page);
    await expect(page.getByTestId('chart-collapsed')).toHaveCount(0);
    await expect(page.locator('canvas')).toHaveCount(0);
    const why = page.getByTestId('coach-fix-first').locator('p').nth(1);
    const bottom = await why.evaluate((node) => node.getBoundingClientRect().bottom + window.scrollY);
    expect(bottom).toBeLessThan(FOLD_Y);
  });

  test('the rail never runs past the report column', async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 900 });
    await routeCoach(page, STATES['ready-model']);
    await openPanel(page);
    const width = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(width).toBeLessThanOrEqual(1440);
    const card = await page.getByTestId('coach-card').boundingBox();
    const rail = await page.getByTestId('coach-rail').boundingBox();
    expect(rail!.x + rail!.width).toBeLessThanOrEqual(card!.x + card!.width + 1);
  });

  test('a healer reads their own review and the healing anchor', async ({ page }) => {
    const answer = coachFixture('holy-model');
    await signInAs(page, 'premium', ['sunwick']);
    await routeCoach(page, STATES.healer);
    await openPanel(page, PANEL_URL(GUID.sunwick));
    await expect(page.getByTestId('coach-title')).toHaveText(fill(coachCopy.titleModel, { Spec: 'Holy' }));
    await expect(page.getByTestId('coach-fix-title')).toContainText(answer.review.fixFirst!.title);
    await expect(page.getByTestId('coach-footer')).toContainText(answer.basis_labels.at(-1)!);
    await expect(page.locator('[data-testid="coach-chips"] a[href*="tab=healing"]')).toHaveCount(1);
  });

  test('whole night: one muted line, no request', async ({ page }) => {
    const seen = await routeCoach(page, STATES['ready-model']);
    await page.goto(PANEL_URL().replace('fight=3', 'fight=0'));
    await expect(page.getByTestId('coach-tab-night')).toHaveText(coachCopy.night);
    expect(seen.panels).toHaveLength(0);
    // The test ends before the summary stub's fetch settles; its handler must not outlive the page.
    await page.unrouteAll({ behavior: 'ignoreErrors' });
  });

  test('the phone strip holds its reserve by entitlement: 52 px Premium, 88 px free', async ({ browser }) => {
    for (const kind of ['premium', 'free'] as const) {
      const page = await browser.newPage({ viewport: PHONE });
      await signInAs(page, kind);
      await routeCoach(page, STATES[kind === 'premium' ? 'ready-model' : 'teaser-several']);
      await page.goto(ROWS_URL);
      const line = page.getByTestId('coach-row-line');
      await expect(line).toBeVisible();
      expect(Math.round((await line.boundingBox())!.height)).toBe(STRIP_PHONE_PX[kind]);
      await page.close();
    }
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
    await page.getByTestId(`actor-${GUID.morrowlyn}`).getByRole('button').first().click();
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
