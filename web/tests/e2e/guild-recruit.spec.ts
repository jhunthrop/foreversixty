// web/tests/e2e/guild-recruit.spec.ts
// Recruitment, applications and the member profile (web layout spec 2026-11-05 §7.2) on the mocked
// API (support/guild-recruit.ts): a visitor reads the card and is asked to sign in; a player applies,
// sees their application and withdraws; a closed guild says so; officers vote, accept (and read what
// Discord did) and edit the card; an officer reads and writes notes on a profile; a member reading
// themselves never meets a note; Action needed opens the applications; nothing scrolls sideways.
import { expect, test, apiUnreachable, GUILD_OVERVIEW_READS } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { GUILD_URL, envelope, expectNoAxeViolations, stubGuildCentre } from './support/officer-tools';
import {
  APPLICANT_KEY,
  APPLICATIONS_URL,
  RECRUIT_URL,
  profileUrl,
  stubRecruit,
} from './support/guild-recruit';
import { GUILD_ID } from '../../src/fixtures/guild/mock-guild';
import { buildMockApplication } from '../../src/fixtures/guild/mock-recruit';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

test.describe('the recruiting card', () => {
  test('a visitor reads the card and the officers, and is asked to sign in', async ({ page }) => {
    await stubRecruit(page, { viewer: 'visitor' });
    await gotoHydrated(page, RECRUIT_URL, 'guild-recruit-card');
    await expect(page.getByTestId('guild-recruit-head')).toContainText('Join Olympus XXVII');
    await expect(page.getByTestId('guild-recruit-open')).toHaveText('Recruiting');
    await expect(page.getByTestId('guild-recruit-facts')).toContainText(
      'Tue and Thu, 8 to 11 pm server time',
    );
    // Looking for (batch 3): each spec in its class colour with the officers' note under it.
    const wanted = page.getByTestId('guild-recruit-wanted');
    for (const words of [
      'Holy Paladin',
      'Two for the heal rotation',
      'Protection Warrior',
      'Off tank',
      'Mage, any spec',
    ]) {
      await expect(wanted).toContainText(words);
    }
    await expect(page.getByTestId('guild-recruit-officers')).toContainText(
      'Protection Warrior · Guild master',
    );
    await expect(page.getByTestId('guild-recruit-sign-in')).toBeVisible();
    await expect(page.getByTestId('guild-recruit-form')).toHaveCount(0);
    await expectNoAxeViolations(page, '[data-testid="guild-recruit-page"]');
  });

  test('a player applies with their own character, the opt-in off, then withdraws', async ({ page }) => {
    const writes = await stubRecruit(page, { viewer: 'applicant' });
    await gotoHydrated(page, RECRUIT_URL, 'guild-recruit-form');
    // The only character on the guild's realm is picked already; the one on EU is not offered.
    await expect(page.getByTestId('guild-recruit-discord-dm')).not.toBeChecked();
    await page.getByTestId('guild-recruit-message').fill('  Two tiers of heals.  ');
    await page.getByTestId('guild-recruit-availability').fill('Tue, Thu');
    await page.getByTestId('guild-recruit-submit').click();
    await expect(page.getByTestId('guild-recruit-mine-row')).toContainText(
      "Ashenveil's application is with the officers.",
    );
    expect(writes[0]).toMatchObject({
      method: 'POST',
      path: `/v1/guilds/${GUILD_ID}/applications`,
      body: {
        character_key: APPLICANT_KEY,
        message: 'Two tiers of heals.',
        availability: 'Tue, Thu',
        discord_dm: false,
      },
    });
    await page.getByTestId('guild-recruit-withdraw').click();
    await expect(page.getByTestId('guild-recruit-form')).toBeVisible();
    expect(writes[1]).toMatchObject({
      method: 'POST',
      path: `/v1/guilds/${GUILD_ID}/applications/41/withdraw`,
    });
  });

  test('a refusal reads in words under the form', async ({ page }) => {
    await stubRecruit(page, { viewer: 'applicant', applyRefusal: 'already_applied' });
    await gotoHydrated(page, RECRUIT_URL, 'guild-recruit-form');
    await page.getByTestId('guild-recruit-submit').click();
    await expect(page.getByTestId('guild-recruit-error')).toHaveText(
      'You already have an application open with this guild.',
    );
  });

  test('an open application shows its status on arrival', async ({ page }) => {
    await stubRecruit(page, { viewer: 'applicant', mine: [buildMockApplication()] });
    await gotoHydrated(page, RECRUIT_URL, 'guild-recruit-mine');
    await expect(page.getByTestId('guild-recruit-mine-status')).toHaveText('Waiting');
    await expect(page.getByTestId('guild-recruit-form')).toHaveCount(0);
  });

  test('a closed guild says so in one line and lists no wants', async ({ page }) => {
    await stubRecruit(page, { viewer: 'applicant', open: false });
    await gotoHydrated(page, RECRUIT_URL, 'guild-recruit-closed');
    await expect(page.getByTestId('guild-recruit-closed')).toHaveText(
      'Olympus XXVII is not recruiting right now.',
    );
    await expect(page.getByTestId('guild-recruit-open')).toHaveText('Not recruiting');
    await expect(page.getByTestId('guild-recruit-form')).toHaveCount(0);
    await expect(page.getByTestId('guild-recruit-wanted')).toHaveCount(0);
  });

  test('at 390 who they look for leads, the form follows, the card after, and nothing scrolls sideways', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 390, height: 844 });
    await stubRecruit(page, { viewer: 'applicant' });
    await gotoHydrated(page, RECRUIT_URL, 'guild-recruit-form');
    const wanted = await page.getByTestId('guild-recruit-wanted').boundingBox();
    const form = await page.getByTestId('guild-recruit-form').boundingBox();
    const card = await page.getByTestId('guild-recruit-card').boundingBox();
    expect(form!.y).toBeGreaterThan(wanted!.y);
    expect(card!.y).toBeGreaterThan(form!.y);
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
  });
});

test.describe('applications, officers', () => {
  test('votes, accepts, and reads what happened on Discord', async ({ page }) => {
    const writes = await stubRecruit(page, { viewer: 'officer' });
    await gotoHydrated(page, APPLICATIONS_URL, 'guild-applications-plate');
    const second = page.getByTestId('guild-application-42');
    await expect(second).toContainText('Destruction Warlock · Level 60 · Item level 60 · Top DPS 465');
    await second.getByTestId('guild-application-votes-yes').click();
    await expect(second.getByTestId('guild-application-votes-yes')).toHaveAttribute('aria-pressed', 'true');
    expect(writes[0]).toMatchObject({ method: 'PUT', body: { vote: 'yes' } });

    await second.getByTestId('guild-application-accept').click();
    await expect(page.getByTestId('guild-applications-outcome')).toHaveText(
      'Quillbright accepted. Discord role granted. Told them on Discord.',
    );
    await expect(page.getByTestId('guild-application-42')).toHaveCount(0);
    await expect(page.getByTestId('guild-application-decided-42')).toContainText('Discord role granted');
    await expect(page.getByTestId('guild-application-decided-39')).toContainText('by Kraggor');
    await expectNoAxeViolations(page, '[data-testid="guild-applications-page"]');
  });

  test('edits the card: closes it and sets the Discord role', async ({ page }) => {
    const writes = await stubRecruit(page, { viewer: 'officer' });
    await gotoHydrated(page, APPLICATIONS_URL, 'guild-recruit-edit');
    await page.getByTestId('guild-recruit-edit-wanted-open').click();
    await page.getByTestId('guild-recruit-edit-role').fill('12');
    await expect(page.getByTestId('guild-recruit-edit-save')).toBeDisabled();
    await page.getByTestId('guild-recruit-edit-role').fill('987654321098765432');
    await page.getByTestId('guild-recruit-edit-open').click();
    await page.getByTestId('guild-recruit-edit-wanted-remove').first().click();
    await page.getByTestId('guild-recruit-edit-save').click();
    await expect(page.getByTestId('guild-recruit-edit-line')).toHaveText('Saved.');
    expect(writes.at(-1)).toMatchObject({
      method: 'PATCH',
      body: { open: false, discord_role_id: '987654321098765432' },
    });
    expect((writes.at(-1)?.body as { wanted: unknown[] }).wanted).toHaveLength(2);
  });

  test('a member is told only officers see applications', async ({ page }) => {
    await stubRecruit(page, { viewer: 'member' });
    await gotoHydrated(page, APPLICATIONS_URL, 'guild-applications-forbidden');
    await expect(page.getByTestId('guild-application-41')).toHaveCount(0);
  });

  test('Action needed counts the applications and Open goes to them', async ({ page }) => {
    await stubGuildCentre(page, {
      viewer: 'officer',
      officerTools: true,
      raidGuildTools: true,
      claimHolder: true,
    });
    await page.route(`**/v1/guilds/${GUILD_ID}/officer/actions`, (route) =>
      route.fulfill(
        envelope({
          approvals: { count: 0, names: [] },
          applications: { count: 2, names: ['Ashenveil', 'Quillbright'] },
          not_signed: { raid_id: null, count: 0, names: [] },
          drops_to_award: null,
        }),
      ),
    );
    await gotoHydrated(page, GUILD_URL, 'guild-overview-actions');
    await expect(page.getByTestId('guild-overview-actions')).toContainText('2 applications');
    await expect(page.getByTestId('guild-open-applications')).toHaveAttribute('href', APPLICATIONS_URL);
  });
});

test.describe('the member profile', () => {
  test('an officer reads the raider and adds and deletes a note', async ({ page }) => {
    const writes = await stubRecruit(page, { viewer: 'officer' });
    await gotoHydrated(page, profileUrl('u:1001'), 'guild-profile-notes');
    await expect(page.getByTestId('guild-profile-head')).toContainText('Zulmara');
    await expect(page.getByTestId('guild-profile-figure-attendance')).toContainText('3 of 4');
    await expect(page.getByTestId('guild-profile-attendance-grid-row')).toHaveCount(1);
    await expect(page.getByTestId('guild-profile-attendance-line')).toHaveText('25 of the last 30 raids');
    await expect(page.getByTestId('guild-profile-loot-row')).toContainText("Dragonstalker's Helm");
    await expect(page.getByTestId('guild-profile-note')).toHaveCount(2);
    await expect(page.getByTestId('guild-profile-note').first()).toContainText('Officers only');

    await page.getByTestId('guild-profile-note-body').fill('Swapped to Survival for Onyxia.');
    await page.getByTestId('guild-profile-note-save').click();
    await expect(page.getByTestId('guild-profile-note')).toHaveCount(3);
    await expect(page.getByTestId('guild-profile-note-body')).toHaveValue('');
    expect(writes[0]).toMatchObject({ method: 'POST', body: { body: 'Swapped to Survival for Onyxia.' } });

    await page.getByTestId('guild-profile-note-delete').first().click();
    await expect(page.getByTestId('guild-profile-note')).toHaveCount(2);
    expect(writes[1].method).toBe('DELETE');
    await expectNoAxeViolations(page, '[data-testid="guild-profile-page"]');
  });

  test('a member reading themselves never meets a note', async ({ page }) => {
    await stubRecruit(page, { viewer: 'member' });
    await gotoHydrated(page, profileUrl('me'), 'guild-profile-gear-gap');
    await expect(page.getByTestId('guild-profile-head')).toContainText('Zulmara');
    await expect(page.getByTestId('guild-profile-notes')).toHaveCount(0);
    await expect(page.getByTestId('guild-profile-page')).not.toContainText(/note/i);
  });

  test('a member opening another raider is turned away', async ({ page }) => {
    await stubRecruit(page, { viewer: 'member' });
    await gotoHydrated(page, profileUrl('u:1000'), 'guild-profile-refused');
    await expect(page.getByTestId('guild-profile-refused')).toHaveText(
      "Only officers see another raider's profile.",
    );
  });

  test('the roster names link to profiles: an officer to anyone', async ({ page }) => {
    await stubGuildCentre(page, {
      viewer: 'officer',
      officerTools: true,
      raidGuildTools: true,
      claimHolder: true,
    });
    await gotoHydrated(page, `${GUILD_URL}#roster`, 'guild-roster-tab');
    const link = page.locator('a[data-testid="guild-roster-name"]', { hasText: 'Zulmara' });
    await expect(link).toHaveAttribute('href', /\/members\/u%3A\d+$/);
  });
});
