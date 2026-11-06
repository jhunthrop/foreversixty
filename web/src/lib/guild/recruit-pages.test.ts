// The three views' states (recruit, applications, member profile) on injected reads and writes.
import { describe, expect, it, vi } from 'vitest';
import { buildMockGuildPage } from '../../fixtures/guild/mock-guild';
import {
  buildMockApplication,
  buildMockDecision,
  buildMockNotes,
  buildMockOfficerApplications,
  buildMockProfile,
  buildMockRecruitCard,
} from '../../fixtures/guild/mock-recruit';
import type { Me } from '../account/me-types';
import { GuildApiError } from './api';
import { ApplicationsPage } from './applications-page.svelte';
import { loadGuildView, viewRole } from './guild-view-load';
import { ProfilePage } from './profile-page.svelte';
import { RecruitPage } from './recruit-page.svelte';

const PATH = { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' } as const;

function me(rank: 'member' | 'officer' | 'leader' | null, verified = true): Me {
  return {
    user: { id: 9, battletag: 'Ash#1', email: null, role: 'user', anonymize: false },
    characters: [
      { key: 'us/pvp/ashenveil', region: 'us', ruleset: 'pvp', name: 'Ashenveil', class: 'paladin' },
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
              verified,
              verified_count: 1,
            },
          ],
  } as Me;
}

const view = (who: Me | null) => ({
  loadGuild: async () => buildMockGuildPage(),
  loadMe: async () => who,
});

describe('who is asking', () => {
  it('reads the role from /v1/me by realm and slug', async () => {
    expect(viewRole(null, null)).toBe('signed-out');
    expect((await loadGuildView(PATH, view(me(null)))).role).toBe('visitor');
    expect((await loadGuildView(PATH, view(me('member', false)))).role).toBe('unverified');
    expect((await loadGuildView(PATH, view(me('member')))).role).toBe('member');
    expect((await loadGuildView(PATH, view(me('leader')))).role).toBe('officer');
    const failedMe = {
      loadGuild: async () => buildMockGuildPage(),
      loadMe: () => Promise.reject(new Error('x')),
    };
    expect((await loadGuildView(PATH, failedMe)).role).toBe('signed-out');
  });
});

describe('the recruit page', () => {
  function page(who: Me | null, overrides: Record<string, unknown> = {}) {
    return new RecruitPage(
      {
        ...view(who),
        loadCard: async () => buildMockRecruitCard(),
        loadMine: async () => ({ rows: [] }),
        apply: async () => buildMockApplication(),
        withdraw: async () => buildMockApplication('withdrawn'),
        ...overrides,
      },
      PATH,
    );
  }

  it('signed out reads no applications and asks to sign in', async () => {
    const loadMine = vi.fn();
    const p = page(null, { loadMine });
    await p.start();
    expect(p.load).toBe('ready');
    expect(p.state?.kind).toBe('signed-out');
    expect(loadMine).not.toHaveBeenCalled();
  });

  it('applies, shows the pending application, withdraws, and offers the form again', async () => {
    const apply = vi.fn(async () => buildMockApplication());
    const p = page(me(null), { apply });
    await p.start();
    expect(p.state?.kind).toBe('form');
    const input = { character_key: 'us/pvp/ashenveil', message: 'hi', availability: '', discord_dm: false };
    await p.apply(input);
    expect(apply).toHaveBeenCalledWith(2024, input);
    expect(p.state?.kind).toBe('application');
    await p.withdraw();
    expect(p.mine[0].status).toBe('withdrawn');
    expect(p.state?.kind).toBe('form');
  });

  it('puts a refusal in words and keeps the form', async () => {
    const p = page(me(null), {
      apply: async () => {
        throw new GuildApiError('Conflict', 409, 'already_applied');
      },
    });
    await p.start();
    await p.apply({ character_key: 'us/pvp/ashenveil', message: '', availability: '', discord_dm: false });
    expect(p.error).toBe('You already have an application open with this guild.');
    expect(p.state?.kind).toBe('form');
  });

  it('fails the load when the card does not answer', async () => {
    const p = page(null, { loadCard: () => Promise.reject(new Error('down')) });
    await p.start();
    expect(p.load).toBe('failed');
  });
});

describe('the applications page', () => {
  function page(who: Me | null, overrides: Record<string, unknown> = {}) {
    return new ApplicationsPage(
      {
        ...view(who),
        loadCard: async () => buildMockRecruitCard(),
        loadList: async () => buildMockOfficerApplications(),
        vote: async () => ({ yes: 2, no: 1, voters: 3, mine: 'no' as const }),
        decide: async (_g: number, appId: number, decision: string) =>
          buildMockDecision(
            buildMockOfficerApplications().rows.find((r) => r.id === appId)!,
            decision === 'accept',
          ),
        saveCard: async () => ({ card: buildMockRecruitCard(false), discord_role_id: null }),
        ...overrides,
      },
      PATH,
    );
  }

  it('turns a member away before any officer read', async () => {
    const loadList = vi.fn();
    const p = page(me('member'), { loadList });
    await p.start();
    expect(p.load).toBe('forbidden');
    expect(loadList).not.toHaveBeenCalled();
  });

  it('reads a 403 from the server as forbidden too', async () => {
    const p = page(me('officer'), {
      loadList: () => Promise.reject(new GuildApiError('no', 403, 'forbidden')),
    });
    await p.start();
    expect(p.load).toBe('forbidden');
  });

  it('votes, then accepts: the row joins the record and the page says what Discord did', async () => {
    const p = page(me('officer'));
    await p.start();
    await p.vote(42, 'no');
    expect(p.list?.rows.find((r) => r.id === 42)?.votes.mine).toBe('no');
    await p.decide(42, 'accept');
    expect(p.list?.pending).toBe(1);
    expect(p.list?.rows.map((r) => r.id)).toEqual([41, 42, 39]);
    expect(p.lastDecision).toMatchObject({
      id: 42,
      status: 'accepted',
      role_outcome: 'granted',
      dm_outcome: 'sent',
    });
  });

  it('keeps a failed decision on its row', async () => {
    const p = page(me('officer'), { decide: () => Promise.reject(new Error('x')) });
    await p.start();
    await p.decide(41, 'decline');
    expect(p.rowError).toEqual({ id: 41, text: 'That did not save. Try again.' });
    expect(p.list?.pending).toBe(2);
  });

  it('saves the card and takes the role id the API answers', async () => {
    const saveCard = vi.fn(async () => ({ card: buildMockRecruitCard(false), discord_role_id: null }));
    const p = page(me('officer'), { saveCard });
    await p.start();
    const ok = await p.saveCard({ open: false, pitch: ' x ', nights: '', wanted: [], roleId: '' });
    expect(ok).toBe(true);
    expect(saveCard).toHaveBeenCalledWith(2024, {
      open: false,
      pitch: 'x',
      raid_nights: '',
      wanted: [],
      discord_role_id: '',
    });
    expect(p.card?.open).toBe(false);
    expect(p.list?.discord_role_id).toBeNull();
    expect(p.saveLine).toEqual({ tone: 'ok', text: 'Saved.' });
  });
});

describe('the member profile page', () => {
  function page(who: Me | null, account: string, overrides: Record<string, unknown> = {}) {
    return new ProfilePage(
      {
        ...view(who),
        loadProfile: async () => buildMockProfile('Zulmara', true),
        addNote: async (_g: number, _a: string, body: string) => ({ ...buildMockNotes()[0], id: 99, body }),
        deleteNote: async () => ({ status: 'removed' as const }),
        ...overrides,
      },
      PATH,
      account,
    );
  }

  it('a member reading themselves holds no notes at all', async () => {
    const p = page(me('member'), 'me', { loadProfile: async () => buildMockProfile('Zulmara') });
    await p.start();
    expect(p.load).toBe('ready');
    expect(p.notes).toBeNull();
    expect(await p.addNote('sneaky')).toBe(false);
  });

  it('an officer adds a note on the profile account and deletes one', async () => {
    const addNote = vi.fn(async (_g: number, _a: string, body: string) => ({
      ...buildMockNotes()[0],
      id: 99,
      body,
    }));
    const p = page(me('officer'), 'u:1001', { addNote });
    await p.start();
    expect(await p.addNote('  Brings his own flasks  ')).toBe(true);
    expect(addNote).toHaveBeenCalledWith(2024, p.profile?.account_key, 'Brings his own flasks');
    expect(p.notes?.[0].id).toBe(99);
    await p.deleteNote(7);
    expect(p.notes?.map((n) => n.id)).toEqual([99, 6]);
  });

  it('reads signed out, forbidden and not found from the API', async () => {
    const signedOut = page(null, 'me');
    await signedOut.start();
    expect(signedOut.load).toBe('signed-out');
    for (const [status, load] of [
      [403, 'forbidden'],
      [404, 'not-found'],
      [500, 'failed'],
    ] as const) {
      const p = page(me('member'), 'u:5', {
        loadProfile: () => Promise.reject(new GuildApiError('x', status)),
      });
      await p.start();
      expect(p.load).toBe(load);
    }
  });
});
