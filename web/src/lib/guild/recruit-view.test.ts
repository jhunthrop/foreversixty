import { describe, expect, it } from 'vitest';
import {
  buildMockApplication,
  buildMockDecision,
  buildMockOfficerApplications,
  buildMockProfile,
  buildMockRecruitCard,
} from '../../fixtures/guild/mock-recruit';
import type { Me } from '../account/me-types';
import { GuildApiError } from './api';
import {
  applicantFacts,
  appliedDecision,
  decisionSub,
  decisionTitle,
  discordLines,
  draftOf,
  patchOf,
  roleIdError,
  specOptions,
  splitApplications,
  wantedLabel,
  withRow,
} from './applications-view';
import { attendanceView, gearGapLines, lootRows, noteSub, profileFigures, profileHead } from './profile-view';
import {
  applicationLine,
  applyErrorLine,
  eligibleCharacters,
  officerRows,
  progressLine,
  recruitFacts,
  recruitState,
  upcomingNights,
  wantedRows,
  wantedSummary,
} from './recruit-view';

const TZ = 'America/Chicago';
const LOCALE = 'en-US';

function me(overrides: Partial<Me> = {}): Me {
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
      { key: 'eu/pvp/other', region: 'eu', ruleset: 'pvp', name: 'Other', class: 'mage', level: 60 },
    ],
    guilds: [],
    ...overrides,
  } as Me;
}

describe('the public card', () => {
  const card = buildMockRecruitCard();

  it('states the realm, faction and size, and the next nights in the viewer clock', () => {
    expect(recruitFacts(card, TZ, LOCALE).map((f) => [f.label, f.value])).toEqual([
      ['Region', 'US PvP'],
      ['Faction', 'Horde'],
      ['Members', '24'],
      ['Raids', card.raid_nights.text],
      ['Next raid nights', 'Wed 9 Dec, 8 pm, Mon 14 Dec, 8 pm'],
      ['Progress', '1 of 1 named bosses down · Tier 1'],
    ]);
    expect(upcomingNights(card, TZ, LOCALE)).toEqual(['Wed 9 Dec, 8 pm', 'Mon 14 Dec, 8 pm']);
    expect(progressLine(card)).toBe('1 of 1 named bosses down · Tier 1');
    expect(progressLine({ ...card, progress: { ...card.progress, named_encounters: 0 } })).toBe('');
  });

  it('names what they look for, a class alone reading any spec', () => {
    expect(wantedRows(card.wanted).map((r) => [r.title, r.note, r.classSlug])).toEqual([
      ['Holy Paladin', 'Two for the heal rotation', 'paladin'],
      ['Protection Warrior', 'Off tank', 'warrior'],
      ['Mage, any spec', '', 'mage'],
    ]);
    expect(wantedSummary(card.wanted)).toBe('Holy Paladin, Protection Warrior, Mage');
    expect(wantedSummary([])).toBe('Every class and spec welcome');
  });

  it('lists the officers with their spec, class and rank in words', () => {
    expect(officerRows(card.officers).map((row) => row.sub)).toEqual([
      'Protection Warrior · Guild master',
      'Combat Rogue · Officer',
    ]);
    const hunter = officerRows([{ name: 'Bow', class: 'hunter', spec: 'BeastMastery', rank: 'leader' }]);
    expect(hunter[0].sub).toBe('Beast Mastery Hunter · Guild master');
  });
});

describe('the application form', () => {
  const card = buildMockRecruitCard();

  it('offers only the viewer characters on the guild realm', () => {
    expect(eligibleCharacters(me(), card).map((c) => c.value)).toEqual(['us/pvp/ashenveil']);
    expect(eligibleCharacters(me(), card)[0].sub).toBe('Level 60 Paladin');
  });

  it('picks the state: signed out, closed, member, own application, no character, the form', () => {
    expect(recruitState(card, null, []).kind).toBe('signed-out');
    expect(recruitState(buildMockRecruitCard(false), null, []).kind).toBe('closed');
    expect(recruitState(buildMockRecruitCard(false), me(), []).kind).toBe('closed');
    const member = me({ guilds: [{ id: card.guild.id, verified: true } as Me['guilds'][number]] });
    expect(recruitState(card, member, []).kind).toBe('member');
    expect(recruitState(card, me(), [buildMockApplication()]).kind).toBe('application');
    // A decided application on a closed card still shows its outcome; on an open one, apply again.
    expect(recruitState(buildMockRecruitCard(false), me(), [buildMockApplication('declined')]).kind).toBe(
      'application',
    );
    expect(recruitState(card, me(), [buildMockApplication('withdrawn')]).kind).toBe('form');
    expect(recruitState(card, me({ characters: [] }), []).kind).toBe('no-character');
  });

  it('words the status and the refusals', () => {
    expect(applicationLine(buildMockApplication(), 'Olympus XXVII')).toBe(
      "Ashenveil's application is with the officers.",
    );
    expect(applicationLine(buildMockApplication('accepted'), 'Olympus XXVII')).toBe(
      'Ashenveil was accepted. Welcome to Olympus XXVII.',
    );
    expect(applyErrorLine(new GuildApiError('x', 409, 'already_applied'))).toBe(
      'You already have an application open with this guild.',
    );
    expect(applyErrorLine(new GuildApiError('x', 409, 'not_recruiting'))).toBe(
      'This guild stopped taking applications.',
    );
    expect(applyErrorLine(new GuildApiError('x', 409, 'already_member'))).toBe(
      'You are already in this guild.',
    );
    expect(applyErrorLine(new GuildApiError('x', 400, 'invalid', { message: 'message is too long' }))).toBe(
      'message is too long',
    );
    expect(applyErrorLine(new Error('boom'))).toBe('That did not send. Try again.');
  });
});

describe('the officers list', () => {
  const list = buildMockOfficerApplications();

  it('reads every figure with its unit, and splits pending from decided', () => {
    expect(applicantFacts(list.rows[0].character)).toBe(
      'Holy Paladin · Level 60 · Item level 63 · Top HPS 465',
    );
    const { pending, decided } = splitApplications(list);
    expect(pending.map((r) => r.id)).toEqual([41, 42]);
    expect(decided.map((r) => r.id)).toEqual([39]);
  });

  it('says what happened on Discord in plain words', () => {
    const accepted = list.rows[2];
    expect(decisionTitle(accepted)).toBe('Hexmarrow accepted');
    expect(decisionSub(accepted)).toBe(
      '3 yes, 0 no of 3 officers · by Kraggor · Discord role granted · No Discord message: they did not ask for one',
    );
    expect(discordLines({ status: 'accepted', role_outcome: 'no_role_set', dm_outcome: 'sent' })).toEqual([
      'No Discord role set, so none was granted',
      'Told them on Discord',
    ]);
    expect(
      discordLines({ status: 'declined', role_outcome: 'not_accepted', dm_outcome: 'not_linked' }),
    ).toEqual(['No Discord message: they have not linked Discord']);
    expect(
      discordLines({ status: 'accepted', role_outcome: 'discord_refused', dm_outcome: null })[0],
    ).toContain('Manage Roles');
  });

  it('folds a decision answer into its row and replaces it without mutating', () => {
    const row = appliedDecision(buildMockDecision(list.rows[1]));
    expect(row).toMatchObject({ status: 'accepted', role_outcome: 'granted', dm_outcome: 'sent' });
    const declined = appliedDecision(buildMockDecision(list.rows[0], false));
    expect(declined).toMatchObject({ role_outcome: 'not_accepted', dm_outcome: 'not_opted_in' });
    const next = withRow(list, row);
    expect(next.pending).toBe(1);
    expect(list.rows[1].status).toBe('pending');
  });

  it('edits the card: draft, role id rule, the patch, spec choices', () => {
    const draft = draftOf(buildMockRecruitCard(), '123456789012345678');
    expect(draft.wanted[0]).toEqual({
      class: 'paladin',
      spec: 'paladin-holy',
      note: 'Two for the heal rotation',
    });
    expect(roleIdError('')).toBe('');
    expect(roleIdError('1234')).toBe('A role id is 17 to 20 digits.');
    expect(roleIdError(' 123456789012345678 ')).toBe('');
    expect(patchOf({ ...draft, pitch: '  Hi  ', roleId: ' ' })).toMatchObject({
      pitch: 'Hi',
      discord_role_id: '',
    });
    expect(specOptions('paladin').map((o) => o.value)).toEqual([
      '',
      'paladin-holy',
      'paladin-protection',
      'paladin-retribution',
    ]);
    expect(wantedLabel({ class: 'paladin', spec: 'paladin-holy', note: '' })).toBe('Holy Paladin');
    expect(wantedLabel({ class: 'mage', spec: null, note: '' })).toBe('Mage');
  });
});

describe('the member profile', () => {
  const profile = buildMockProfile('Zulmara', true);

  it('heads with the spec, level and rank, and only the figures with a source', () => {
    const head = profileHead(profile);
    expect(head.name).toBe('Zulmara');
    expect(head.classSlug).toBe('hunter');
    expect(head.facts[0]).toBe('Marksmanship Hunter');
    const figures = profileFigures(profile).map((f) => [f.label, f.value]);
    expect(figures[0]).toEqual(['Attendance', '3 of 4']);
    expect(figures.at(-1)).toEqual(['Items this tier', '1']);
    expect(profileFigures({ ...profile, attendance: null, member: null }).map((f) => f.key)).toEqual([
      'items',
    ]);
  });

  it('writes the gear gap in the BiS unit, and the counts alone for a healer', () => {
    const base = profile.readiness!;
    const at = (gap: NonNullable<typeof base.gear_gap>) =>
      gearGapLines({ ...profile, readiness: { ...base, gear_gap: gap } });
    expect(at({ upgrades: 3, gain_dps: 12.44, not_sim_checked: 1, unit: 'dps' })?.map((l) => l.text)).toEqual(
      ['3 upgrades on their BiS list', '+12.4 DPS over what they wear', '1 not sim-checked'],
    );
    expect(at({ upgrades: 2, gain_dps: 0, not_sim_checked: 0, unit: null })?.map((l) => l.text)).toEqual([
      '2 upgrades on their BiS list',
    ]);
    expect(at({ upgrades: 0, gain_dps: 0, not_sim_checked: 0, unit: 'dps' })?.[0].text).toBe(
      'Nothing on their BiS list beats what they wear.',
    );
    expect(gearGapLines({ ...profile, readiness: null })).toBeNull();
  });

  it('draws the last 4 raids oldest first with the 30-raid line, and the reports mode line', () => {
    const view = attendanceView(profile, TZ, LOCALE)!;
    expect(view.columns).toEqual(['Mon 23 Nov', 'Wed 25 Nov', 'Mon 30 Nov', 'Wed 2 Dec']);
    expect(view.rows[0].cells.map((c) => c.text)).toEqual(['✓', '✓', '×', 'L']);
    expect(view.line).toBe('25 of the last 30 raids');
    const reports = attendanceView(
      {
        ...profile,
        attendance: { source: 'reports', last_4: { present: 2, nights: 4 }, last_30: null, nights: [] },
      },
      TZ,
      LOCALE,
    )!;
    expect(reports).toEqual({ columns: [], rows: [], line: '2 of the last 4 logged nights' });
  });

  it('lists loot newest first and names who wrote each note', () => {
    expect(lootRows(profile, TZ, LOCALE)[0]).toMatchObject({
      name: "Dragonstalker's Helm",
      source: 'Awarded Mon 30 Nov',
    });
    expect(noteSub(profile.notes![0], TZ, LOCALE)).toBe('Kraggor · Sat 5 Dec');
    expect(noteSub({ ...profile.notes![0], author: null }, TZ, LOCALE)).toBe('An officer · Sat 5 Dec');
  });
});
