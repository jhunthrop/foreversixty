// Recruitment, applications and the member profile rendered with svelte/server (web layout spec
// 2026-11-05 §7.2), each view on a seeded state: what a visitor, an applicant, an officer and a
// member reading themselves see, and that a member's own profile carries no trace of notes.
import { render } from 'svelte/server';
import type { Component } from 'svelte';
import { describe, expect, it } from 'vitest';
import { buildMockGuildPage } from '../../../fixtures/guild/mock-guild';
import {
  buildMockApplication,
  buildMockOfficerApplications,
  buildMockProfile,
  buildMockRecruitCard,
} from '../../../fixtures/guild/mock-recruit';
import type { Me } from '../../../lib/account/me-types';
import { ApplicationsPage } from '../../../lib/guild/applications-page.svelte';
import { ProfilePage } from '../../../lib/guild/profile-page.svelte';
import { RecruitPage } from '../../../lib/guild/recruit-page.svelte';
import ApplyForm from './ApplyForm.svelte';
import GuildApplications from './GuildApplications.svelte';
import GuildMemberProfile from './GuildMemberProfile.svelte';
import GuildRecruit from './GuildRecruit.svelte';
import RecruitEdit from './RecruitEdit.svelte';

const PATH = { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' } as const;
const TZ = 'America/Chicago';
const noop = () => {};

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

function me(rank: 'member' | 'officer' | null): Me {
  return {
    user: { id: 9, battletag: 'Ash#1', email: null, role: 'user', anonymize: false },
    characters: [
      {
        key: 'us/pvp/ashenveil',
        region: 'us',
        ruleset: 'pvp',
        name: 'Ashenveil',
        class: 'paladin',
        level: 60,
      },
    ],
    guilds:
      rank === null
        ? []
        : [
            {
              id: 2024,
              region: 'us',
              ruleset: 'pvp',
              name: 'Olympus XXVII',
              rank,
              verified: true,
              verified_count: 1,
            },
          ],
  } as Me;
}

const view = (who: Me | null) => ({ loadGuild: async () => buildMockGuildPage(), loadMe: async () => who });

async function recruitPage(
  who: Me | null,
  open = true,
  mine = [] as ReturnType<typeof buildMockApplication>[],
) {
  const p = new RecruitPage(
    {
      ...view(who),
      loadCard: async () => buildMockRecruitCard(open),
      loadMine: async () => ({ rows: mine }),
      apply: async () => buildMockApplication(),
      withdraw: async () => buildMockApplication('withdrawn'),
    },
    PATH,
  );
  await p.start();
  return html(GuildRecruit, { path: PATH, page: p, viewerTimezone: TZ, locale: 'en-US' });
}

describe('the recruit page', () => {
  it('shows a visitor the card, the officers and a sign-in prompt', async () => {
    const out = await recruitPage(null);
    expect(out).toContain('Join Olympus XXVII');
    expect(out).toContain('Recruiting');
    expect(out).toContain('Tue and Thu, 8 to 11 pm server time');
    expect(out).toContain('Holy Paladin');
    expect(out).toContain('Mage, any spec');
    expect(out).toContain('data-testid="guild-recruit-officers"');
    expect(out).toContain('Sign in with Battle.net to apply');
    expect(out).not.toContain('data-testid="guild-recruit-form"');
    // A visitor gets no guild tabs.
    expect(out).not.toContain('guild-tab-roster');
  });

  it('a closed guild says so in one line and lists no wants', async () => {
    const out = await recruitPage(me(null), false);
    expect(out).toContain('Olympus XXVII is not recruiting right now.');
    expect(out).toContain('Not recruiting');
    expect(out).not.toContain('Holy Paladin');
    expect(out).not.toContain('data-testid="guild-recruit-form"');
  });

  it('a signed-in player gets the form; with one open, its status and Withdraw', async () => {
    const form = await recruitPage(me(null));
    expect(form).toContain('data-testid="guild-recruit-form"');
    expect(form).toContain('Message me on Discord about this application');
    expect(form).not.toMatch(/data-testid="guild-recruit-discord-dm"[^>]*checked/);
    const pending = await recruitPage(me(null), true, [buildMockApplication()]);
    expect(pending).toContain("Ashenveil's application is with the officers.");
    expect(pending).toContain('data-testid="guild-recruit-withdraw"');
  });

  it('a member is told they are in, with the guild tabs', async () => {
    const out = await recruitPage(me('member'));
    expect(out).toContain('You are in Olympus XXVII.');
    expect(out).toContain('guild-tab-roster');
  });

  it('the form preselects the only character and counts the message', () => {
    const out = html(ApplyForm, {
      characters: [
        { value: 'us/pvp/ashenveil', label: 'Ashenveil', sub: 'Level 60 Paladin', classSlug: 'paladin' },
      ],
      busy: false,
      error: 'You already have an application open with this guild.',
      onApply: noop,
    });
    expect(out).toContain('0 / 1000');
    expect(out).toContain('role="alert"');
    expect(out).toContain('maxlength="1000"');
  });
});

describe('the applications page', () => {
  async function applications(who: Me) {
    const p = new ApplicationsPage(
      {
        ...view(who),
        loadCard: async () => buildMockRecruitCard(),
        loadList: async () => buildMockOfficerApplications(),
        vote: async () => ({ yes: 0, no: 0, voters: 3, mine: null }),
        decide: async () => {
          throw new Error('unused');
        },
        saveCard: async () => ({ card: buildMockRecruitCard(), discord_role_id: null }),
      },
      PATH,
    );
    await p.start();
    return html(GuildApplications, { path: PATH, page: p });
  }

  it('lists the open applicants with votes, Accept and Decline, and the record', async () => {
    const out = await applications(me('officer'));
    expect(out).toContain('data-testid="guild-application-41"');
    expect(out).toContain('Holy Paladin · Level 60 · Item level 63 · Top HPS 465');
    expect(out).toContain('2 yes');
    expect(out).toContain('of 3 officers');
    expect(out).toContain('data-testid="guild-application-accept"');
    expect(out).toContain('Hexmarrow accepted');
    expect(out).not.toContain('Discord role granted');
    expect(out).toContain('data-testid="guild-recruit-edit"');
    // The rail reads the card's fields, each with Edit; the form waits behind it.
    expect(out).toContain('data-testid="guild-recruit-edit-wanted-open"');
    expect(out).toContain('Tue and Thu, 8 to 11 pm server time');
    expect(out).not.toContain('value="123456789012345678"');
  });

  it('a member is told only officers see it, with no applicant on the page', async () => {
    const out = await applications(me('member'));
    expect(out).toContain('Only officers see applications.');
    expect(out).not.toContain('Ashenveil');
  });

  it('the card edit carries the role id help line', () => {
    const out = html(RecruitEdit, {
      card: buildMockRecruitCard(),
      roleId: null,
      recruitHref: '/guild/us/pvp/olympus-xxvii/recruit',
      saving: false,
      line: null,
      onSave: noop,
      editing: true,
    });
    expect(out).toContain('Copy Role ID');
    expect(out).toContain('role="switch"');
    expect(out).toContain('data-testid="guild-recruit-edit-wanted-line"');
  });
});

describe('the member profile', () => {
  async function profile(who: Me, officer: boolean) {
    const p = new ProfilePage(
      {
        ...view(who),
        loadProfile: async () => buildMockProfile('Zulmara', officer),
        addNote: async () => {
          throw new Error('unused');
        },
        deleteNote: async () => ({ status: 'removed' as const }),
      },
      PATH,
      officer ? 'u:1001' : 'me',
    );
    await p.start();
    return html(GuildMemberProfile, {
      path: PATH,
      account: 'me',
      page: p,
      viewerTimezone: TZ,
      locale: 'en-US',
    });
  }

  it('an officer reads the figures, the sections and the notes', async () => {
    const out = await profile(me('officer'), true);
    expect(out).toContain('Zulmara');
    expect(out).toContain('3 of 4');
    expect(out).toContain('data-testid="guild-profile-attendance-grid"');
    expect(out).toContain('25 of the last 30 raids');
    expect(out).toContain('Dragonstalker');
    expect(out).toContain('Officer notes');
    expect(out).toContain('data-testid="guild-profile-notes"');
    expect(out).toContain('data-testid="guild-profile-ready-tag"');
    expect(out).toContain('Guild median 61');
    expect(out).toContain('Ask about the Thursday conflict.');
    expect(out).toContain('data-testid="guild-profile-note-form"');
  });

  it('a member reading themselves sees no notes and no word of them', async () => {
    const out = await profile(me('member'), false);
    expect(out).toContain('Zulmara');
    expect(out).not.toMatch(/note/i);
    expect(out).not.toContain('guild-profile-notes');
  });
});
