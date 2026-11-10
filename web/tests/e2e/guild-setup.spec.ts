// web/tests/e2e/guild-setup.spec.ts
// The guild setup checklist (design/specs/2026-11-09-guild-onboarding.md), API stubbed: the
// steps and their live states, the first to-do step's one primary, the bot install that returns
// to the checklist, Skip and Undo through the mark routes, the paste lines, the plan rail with no
// subscribe button, axe, and the settings row. FS_CAPTURE=<dir> also writes the review captures.
import type { Page } from '@playwright/test';
import { apiUnreachable, expect, test, GUILD_OVERVIEW_READS } from './support/test';
import type { GuildSetup, SetupState, SetupStepKey } from '../../src/lib/guild/setup-api';
import { setupCopy } from '../../src/lib/guild/setup-copy';
import { SETTINGS_GUILD } from '../../src/test-support/billing-fixtures';
import { setupFixture } from '../../src/test-support/guild-setup-fixture';
import {
  E2E_API_BASE,
  envelope,
  guildChannel,
  settingsBody,
  stubGuildDiscord,
} from './support/guild-discord';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';

const SETUP_PATH = '/guild/us/hardcore/the-last-watch/setup';
const SETTINGS_PATH = '/guild/us/hardcore/the-last-watch/settings';
const CAPTURE = process.env.FS_CAPTURE;

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

interface SetupStub {
  marks: { method: string; step: string; body: string }[];
}

/** Routes the setup page's reads for The Last Watch: the guild, settings and the checklist. */
async function stubSetup(page: Page, initial: GuildSetup): Promise<SetupStub> {
  let setup = initial;
  const seen: SetupStub = { marks: [] };
  await stubGuildDiscord(page, { discord: guildChannel({ connected: false }), claimHolder: true });
  await page.route(`**/v1/guilds/${SETTINGS_GUILD.id}/settings`, (route) =>
    route.fulfill(
      envelope({
        ...settingsBody(null, false, true),
        default_visibility: 'public',
        invite: {
          rotated_at: '2026-11-09T20:00:00Z',
          url: '/guild/invite/7Hk2QaLw',
          made_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
        },
      }),
    ),
  );
  await page.route(`**/v1/guilds/${SETTINGS_GUILD.id}/setup`, (route) => route.fulfill(envelope(setup)));
  await page.route(`**/v1/guilds/${SETTINGS_GUILD.id}/setup/*`, (route) => {
    const request = route.request();
    const step = request.url().split('/').pop() as SetupStepKey;
    seen.marks.push({ method: request.method(), step, body: request.postData() ?? '' });
    const state: SetupState =
      request.method() === 'DELETE' ? 'todo' : (JSON.parse(request.postData() ?? '{}').state as SetupState);
    setup = {
      ...setup,
      steps: setup.steps.map((s) => (s.key === step ? { ...s, state, marked: state !== 'todo' } : s)),
    };
    return route.fulfill(envelope(setup));
  });
  return seen;
}

async function capture(page: Page, name: string): Promise<void> {
  if (CAPTURE === undefined) return;
  await page.evaluate(() => document.fonts.ready);
  await page.screenshot({ path: `${CAPTURE}/${name}.png`, fullPage: true });
}

const MIXED: Partial<Record<SetupStepKey, SetupState>> = {
  officer_ranks: 'done',
  discord_bot: 'done',
  invite_link: 'done',
  addon_sync: 'done',
  report_visibility: 'done',
  linked_role: 'skipped',
};

function mixedSetup(): GuildSetup {
  return setupFixture(MIXED, {
    plan: 'guild',
    ranks: { ...setupFixture().ranks, officer_max_rank_index: 2, officers: 8 },
    discord: {
      bot_available: true,
      installed: true,
      server_name: 'The Last Watch',
      channel_name: 'raid-signups',
      recruit_channel_name: null,
      webhook_connected: false,
      linked_members: 19,
      members: 31,
    },
    recruiting: {
      open: true,
      pitch: true,
      nights: 'Tue, Thu 8 pm',
      wanted: ['Holy Priest', 'Restoration Shaman', 'Protection Warrior'],
      role_set: true,
    },
    sync: { synced: 21, accounts: 31, window_days: 14 },
    visibility: 'guild',
  });
}

test.describe('the checklist', () => {
  test('right after a claim: every step to do but the link, the first one the page one fill', async ({
    page,
  }) => {
    await stubSetup(page, setupFixture());
    await gotoHydrated(page, SETUP_PATH, 'guild-setup-steps');
    await expect(page.getByTestId('setup-done-figure')).toContainText('1');
    await expect(page.getByTestId('setup-step-officer_ranks-state')).toHaveText('To do');
    await expect(page.getByTestId('setup-step-invite_link-state')).toHaveText('Done');
    await expect(page.locator('[aria-label="Step 1: Next"]')).toBeVisible();
    await expect(page.getByTestId('setup-discord-add')).toHaveAttribute(
      'href',
      `${E2E_API_BASE}/v1/guilds/${SETTINGS_GUILD.id}/discord/bot/install?return=setup`,
    );
    await expect(page.getByTestId('setup-plan')).toContainText(setupCopy.plan.freeHeading);
    await expect(page.getByTestId('setup-plan').getByRole('button')).toHaveCount(0);
    await expectNoAxeViolations(page, '[data-testid="guild-setup"]');
    await capture(page, `setup-todo-${test.info().project.name}`);
  });

  test('Skip marks a step, Undo takes it back', async ({ page }) => {
    const seen = await stubSetup(page, setupFixture());
    await gotoHydrated(page, SETUP_PATH, 'guild-setup-steps');
    await page.getByTestId('setup-step-linked_role-skip').click();
    await expect(page.getByTestId('setup-step-linked_role-state')).toHaveText('Skipped');
    await page.getByTestId('setup-step-linked_role-undo').click();
    await expect(page.getByTestId('setup-step-linked_role-state')).toHaveText('To do');
    expect(seen.marks.map((m) => [m.method, m.step])).toEqual([
      ['PUT', 'linked_role'],
      ['DELETE', 'linked_role'],
    ]);
  });

  test('Keep and I posted it mark the step done', async ({ page }) => {
    const seen = await stubSetup(page, setupFixture({}));
    await gotoHydrated(page, SETUP_PATH, 'guild-setup-steps');
    await expect(page.getByTestId('setup-invite-paste-text')).toHaveValue(
      /^Join The Last Watch on Forever Sixty and sign up for raids: http\S+\/guild\/invite\/7Hk2QaLw$/,
    );
    await page.getByTestId('setup-invite-posted').click();
    await expect(page.getByTestId('setup-step-invite_link-state')).toHaveText('Done');
    await page.getByTestId('setup-ranks-keep').click();
    await expect(page.getByTestId('setup-step-officer_ranks-state')).toHaveText('Done');
    expect(seen.marks.every((m) => m.body.includes('"done"'))).toBe(true);
  });

  test('part way: done steps fold to their facts, the to-do ones keep their controls', async ({ page }) => {
    await stubSetup(page, mixedSetup());
    await gotoHydrated(page, SETUP_PATH, 'guild-setup-steps');
    await expect(page.getByTestId('setup-done-figure')).toContainText('5');
    await expect(page.getByTestId('setup-discord-bot-body')).toContainText(
      'Posts in #raid-signups on The Last Watch · 19 of 31 members linked Discord',
    );
    await expect(page.getByTestId('setup-ranks-input')).toHaveCount(0);
    await expect(page.getByTestId('setup-schedule-plan')).toBeVisible();
    await expect(page.getByTestId('setup-step-linked_role-state')).toHaveText('Skipped');
    await expectNoAxeViolations(page, '[data-testid="guild-setup"]');
    await capture(page, `setup-mixed-${test.info().project.name}`);
  });

  test('every step done: the figure says 8 and the rail names the plan', async ({ page }) => {
    const all = Object.fromEntries(setupFixture().steps.map((s) => [s.key, 'done'])) as Record<
      SetupStepKey,
      'done'
    >;
    const allDone = setupFixture(all);
    await stubSetup(page, {
      ...mixedSetup(),
      steps: allDone.steps,
      done: 8,
      finished: true,
      plan: 'raid_guild',
    });
    await gotoHydrated(page, SETUP_PATH, 'guild-setup-steps');
    await expect(page.getByTestId('setup-done-figure')).toContainText('8');
    await expect(page.getByTestId('setup-plan-current')).toHaveText('Raid guild');
    await capture(page, `setup-done-${test.info().project.name}`);
  });
});

test('settings heads the page with the setup row until every step is settled', async ({ page }) => {
  await stubSetup(page, setupFixture());
  await gotoHydrated(page, SETTINGS_PATH, 'guild-settings');
  const row = page.getByTestId('guild-settings-setup-row');
  await expect(row).toContainText('Setup · 1 of 8 done');
  await expect(row.getByRole('link')).toHaveAttribute('href', SETUP_PATH);
  await capture(page, `settings-row-${test.info().project.name}`);
});
