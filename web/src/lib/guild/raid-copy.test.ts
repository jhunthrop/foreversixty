import { describe, expect, it } from 'vitest';
import { specCell, specString } from '../../test-support/spec-table';
import type { GuildRaidView } from './api';
import {
  bandLine,
  copyList,
  countLine,
  countRatio,
  countdown,
  guildTimeLabel,
  needsBefore,
  raidCopy,
  roleBreakdown,
  showAll,
  whenPhrase,
} from './raid-copy';

const SPEC = new URL('../../../../design/specs/2026-10-26-guild-raids.md', import.meta.url);
const HEADING = '### 6.1 On the page';
const cell = (label: string): string[] => specCell(SPEC, HEADING, label);
const str = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);

const CHICAGO = 'America/Chicago';
const NEW_YORK = 'America/New_York';
// Wednesday 9 Dec 2026, 10:00 am Chicago.
const NOW = new Date('2026-12-09T16:00:00Z');
const at = (isoUtc: string): Date => new Date(isoUtc);

describe('raidCopy constants, verbatim from spec 6.1', () => {
  it('names the tab button, gate and cap', () => {
    expect(raidCopy.planButton).toBe(str('Tab button'));
    expect(raidCopy.gate).toBe(str('Gate sentence'));
    expect(raidCopy.gateLink).toBe('Subscribe the guild');
    expect(raidCopy.cap).toBe(str('Cap'));
  });

  it('has every sub-head, in the spec order', () => {
    const s = raidCopy.subheads;
    expect([
      s.planned,
      s.raidNights,
      s.yourSignup,
      s.signups,
      s.notes,
      s.tanks,
      s.healers,
      s.dps,
      s.late,
      s.bench,
      s.out,
      s.notSigned,
    ]).toEqual(cell('Sub-heads'));
  });

  it('has the segments, placeholders and strip status', () => {
    const g = raidCopy.segments;
    expect([g.in, g.late, g.bench, g.out]).toEqual(cell('Segments'));
    expect([raidCopy.placeholders.note, raidCopy.placeholders.chooseRaid]).toEqual(cell('Placeholders'));
    expect(raidCopy.stripStatus.saved).toBe(str('Strip status', 0));
    expect(raidCopy.stripStatus.failed).toBe(str('Strip status', 1));
    expect(raidCopy.stripStatus.movedToBench('Lightbrand')).toBe(
      str('Strip status', 2, { Name: 'Lightbrand' }),
    );
    expect(raidCopy.stripStatus.left(120)).toBe(str('Strip status', 3, { n: '120' }));
    expect(raidCopy.waitingCharacter).toBe(str('Waiting character'));
  });

  it('has the empty states', () => {
    expect(raidCopy.emptyPlanned.member).toBe(str('Empty Planned', 0));
    expect(raidCopy.emptyPlanned.officer).toBe(str('Empty Planned', 1));
    expect(raidCopy.bandNone.member).toBe(str('Band line, none planned', 0));
    expect(raidCopy.bandNone.officer).toBe(str('Band line, none planned', 1));
    expect(raidCopy.bandNone.planLink).toBe(str('Band line, none planned', 2));
    expect(raidCopy.notesEmpty.member).toBe(str('Notes empty', 0));
    expect(raidCopy.notesEmpty.officer).toBe(str('Notes empty', 1));
    const e = raidCopy.emptyRole;
    expect([e.tank, e.healer, e.dps]).toEqual(cell('Empty role'));
  });

  it('has the raid states, races and poll line', () => {
    const r = raidCopy.raidStates;
    expect([r.locked, r.ended, r.cancelled, r.done, r.openNight]).toEqual(cell('Raid states'));
    const c = raidCopy.races;
    expect([c.locked, c.cancelled, c.officerToolsEnded]).toEqual(cell('Races'));
    expect(raidCopy.poll).toBe(str('Poll'));
  });

  it('has the officer row and the cancel confirm', () => {
    const o = raidCopy.officerRow;
    expect([o.edit, o.lock, o.locking, o.unlock, o.copy, o.copied, o.cancel, o.failed]).toEqual(
      cell('Officer row'),
    );
    const k = raidCopy.cancelConfirm;
    expect([k.question, k.postToDiscord, k.confirm, k.confirming, k.keep]).toEqual(cell('Cancel confirm'));
    expect(raidCopy.copyFailed).toBe(str('Copy failed'));
  });

  it('has the form labels, helpers and errors', () => {
    const f = raidCopy.form;
    expect([
      f.raid,
      f.title,
      f.date,
      f.time,
      f.duration,
      f.guildZone,
      f.notes,
      f.repeat,
      f.postToDiscord,
      f.postToDiscordIfMoved,
      f.submitCreate,
      f.submitCreating,
      f.submitEdit,
      f.submitEditing,
      f.cancel,
    ]).toEqual(cell('Form labels'));
    expect(f.markdownHelper).toBe(str('Form helpers', 2));
    expect(f.onlyFit(2)).toBe(str('Form helpers', 3, { n: '2' }));
    expect(f.connectDiscord).toBe(str('Form helpers', 4));
    expect(f.connectDiscordLink).toBe('Settings');
    const e = f.errors;
    expect([e.raid, e.date, e.time, e.cap, e.failed]).toEqual(cell('Form errors'));
  });

  it('has the not-visible, load-error, back-link and changelog lines', () => {
    expect(raidCopy.notVisible.members('Olympus XXVII')).toBe(
      str('Not visible', 0, { Guild: 'Olympus XXVII' }),
    );
    expect(raidCopy.notVisible.unverified).toBe(str('Not visible', 3));
    expect(raidCopy.loadErrors.raid).toBe(str('Load errors', 0));
    expect(raidCopy.loadErrors.planned).toBe(str('Load errors', 1));
    expect(raidCopy.backLink('Olympus XXVII')).toBe(str('Back link', 0, { Guild: 'Olympus XXVII' }));
    expect(raidCopy.changelog).toBe(str('Changelog'));
  });
});

describe('whenPhrase', () => {
  const when = (starts: string, tz = CHICAGO, viewerTz = CHICAGO, now = NOW): string =>
    whenPhrase(at(starts), now, 'en-US', tz, viewerTz);

  it('says Tonight for the viewer’s own day', () => {
    expect(when('2026-12-10T02:00:00Z')).toBe('Tonight 8 pm');
  });

  it('says Tomorrow for the next day', () => {
    expect(when('2026-12-11T02:00:00Z')).toBe('Tomorrow 8 pm');
  });

  it('names the weekday within six days', () => {
    expect(when('2026-12-12T02:00:00Z')).toBe('Friday 8 pm');
    expect(when('2026-12-15T02:00:00Z')).toBe('Monday 8 pm');
  });

  it('spells the date from the seventh day', () => {
    expect(when('2026-12-18T02:00:00Z')).toBe('Thu 17 Dec, 8 pm');
  });

  it('appends the guild clock when the offsets differ', () => {
    expect(when('2026-12-10T02:00:00Z', NEW_YORK, CHICAGO)).toBe('Tonight 8 pm (9 pm guild time)');
  });

  it('writes minutes off the hour', () => {
    expect(when('2026-12-10T02:30:00Z')).toBe('Tonight 8:30 pm');
  });

  it('uses a 24-hour clock where the locale does', () => {
    expect(whenPhrase(at('2026-12-10T02:00:00Z'), NOW, 'de-DE', CHICAGO, CHICAGO)).toBe('Tonight 20:00');
  });
});

describe('countdown', () => {
  const start = at('2026-12-10T02:00:00Z');
  const end = at('2026-12-10T05:00:00Z');
  const cd = (isoNow: string): string => countdown(start, end, at(isoNow));

  it('counts days, then days and hours', () => {
    expect(cd('2026-12-08T02:00:00Z')).toBe('in 2 days');
    expect(cd('2026-12-08T22:00:00Z')).toBe('in 1 day 4 h');
  });

  it('counts hours and minutes, then minutes', () => {
    expect(cd('2026-12-09T20:40:00Z')).toBe('in 5 h 20 min');
    expect(cd('2026-12-10T01:48:00Z')).toBe('in 12 min');
  });

  it('says starting now inside the first minute either side', () => {
    expect(cd('2026-12-10T01:59:30Z')).toBe('starting now');
    expect(cd('2026-12-10T02:00:30Z')).toBe('starting now');
  });

  it('says how long ago it started, then ended', () => {
    expect(cd('2026-12-10T02:20:00Z')).toBe('started 20 min ago');
    expect(cd('2026-12-10T07:00:00Z')).toBe('ended 2 h ago');
  });
});

describe('counts', () => {
  it('writes the count line and the short ratio', () => {
    expect(countLine(23, 40)).toBe(str('Count', 0, { n: '23', size: '40' }));
    expect(countRatio(23, 40)).toBe(str('Count', 1, { n: '23', size: '40' }));
    expect(countLine(23, null)).toBe('23 signed up');
  });

  it('breaks the roles down, singular at one', () => {
    expect(roleBreakdown(3, 7, 13)).toBe('3 tanks · 7 healers · 13 DPS');
    expect(roleBreakdown(1, 1, 0)).toBe('1 tank · 1 healer · 0 DPS');
  });

  it('labels the guild clock', () => {
    expect(guildTimeLabel(at('2026-12-10T02:00:00Z'), CHICAGO, 'en-US')).toBe(str('Guild time'));
  });

  it('offers to show them all', () => {
    expect(showAll(7)).toBe(str('Show more', 0, { n: '7' }));
  });

  it('names the needs-line deadline', () => {
    expect(needsBefore('Thursday')).toBe('before Thursday');
    expect(needsBefore(null)).toBe('before the next raid');
  });
});

describe('bandLine', () => {
  const base = { when: 'Tonight 8 pm', n: 23, size: 40 as number | null };

  it('follows the In sentence', () => {
    expect(bandLine({ ...base, state: 'in', role: 'dps' })).toBe(
      str('Band line', 0, { when: 'Tonight 8 pm', Role: 'DPS', n: '23', size: '40' }),
    );
  });

  it('follows the Late, Bench and Out sentences', () => {
    expect(bandLine({ ...base, state: 'late', role: 'tank' })).toBe(
      'Tonight 8 pm, you are Late as Tank, 23 of 40 signed up.',
    );
    expect(bandLine({ ...base, state: 'bench', role: 'tank' })).toBe(
      'Tonight 8 pm, you are on the Bench, 23 of 40 signed up.',
    );
    expect(bandLine({ ...base, state: 'out', role: 'tank' })).toBe(
      'Tonight 8 pm, you are Out, 23 of 40 signed up.',
    );
  });

  it('follows the unsigned sentence', () => {
    expect(bandLine({ ...base, state: null })).toBe(
      'Tonight 8 pm, you have not signed up, 23 of 40 signed up.',
    );
    expect(specCell(SPEC, HEADING, 'Band line')[4]).toBe(
      '{when}, you have not signed up, {n} of {size} signed up.',
    );
  });

  it('drops the size when unknown', () => {
    expect(bandLine({ ...base, size: null, state: 'in', role: 'healer' })).toBe(
      'Tonight 8 pm, you are In as Healer, 23 signed up.',
    );
  });

  it('names the character when the account has more than one', () => {
    expect(bandLine({ ...base, state: 'in', role: 'dps', characterName: 'Lightbrand' })).toBe(
      'Tonight 8 pm, you are In on Lightbrand as DPS, 23 of 40 signed up.',
    );
    expect(specCell(SPEC, HEADING, 'Band line')[6]).toBe('you are In on {Name} as {Role}');
  });
});

describe('copyList', () => {
  const view = {
    title: 'Onyxia’s Lair',
    starts_at: '2026-12-10T02:00:00Z',
    signups: [
      {
        character_key: 'a',
        name: 'Zed',
        class: 'warrior',
        spec: 'protection',
        role: 'tank',
        state: 'in',
        note: '',
        own: false,
      },
      {
        character_key: 'b',
        name: 'Amy',
        class: 'priest',
        spec: 'holy',
        role: 'healer',
        state: 'in',
        note: '',
        own: false,
      },
      {
        character_key: 'c',
        name: 'Bo',
        class: 'mage',
        spec: 'fire',
        role: 'dps',
        state: 'late',
        note: '',
        own: false,
      },
      {
        character_key: 'd',
        name: 'Cy',
        class: 'rogue',
        spec: 'combat',
        role: 'dps',
        state: 'out',
        note: '',
        own: false,
      },
    ],
    not_signed: [{ character_key: 'e', name: 'Di', class: 'druid', spec: 'balance' }],
  } as unknown as GuildRaidView;

  it('lists each group on its own line, empty groups left out', () => {
    expect(copyList(view, CHICAGO)).toBe(
      [
        'Onyxia’s Lair · Wed 9 Dec, 8:00 pm CST guild time',
        'Tanks (1): Zed',
        'Healers (1): Amy',
        'Late (1): Bo',
        'Out (1): Cy',
        'Not signed up (1): Di',
      ].join('\n'),
    );
  });
});
