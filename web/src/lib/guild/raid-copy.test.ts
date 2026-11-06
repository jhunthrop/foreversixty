import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { specCell, specString } from '../../test-support/spec-table';
import {
  bandLine,
  bandLineParts,
  cancelActionOf,
  cancelConfirmOf,
  cancelledLineOf,
  cancelQuestionOf,
  compCopy,
  countdown,
  countLine,
  countRatio,
  endedLineOf,
  eyebrowOf,
  needsBefore,
  needsDeadline,
  nextRaidDeadline,
  notesEmptyOf,
  raidCopy,
  showAll,
  showsBandLine,
  whenPhrase,
  yourTimeLine,
  zonedClock,
} from './raid-copy';

const SPEC = new URL('../../../../design/specs/2026-10-26-guild-raids.md', import.meta.url);
const HEADING = '### 6.1 On the page';
const SPEC2 = new URL('../../../../design/specs/2026-10-28-discord-guided-cards.md', import.meta.url);
const HEADING2 = '## P.7 Copy, verbatim';
const str2 = (key: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC2, HEADING2, key, index, fill);
const cell = (label: string): string[] => specCell(SPEC, HEADING, label);
const str = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);

const SPEC4 = new URL('../../../../design/specs/2026-11-05-guild-raid-comp.md', import.meta.url);
const HEADING4 = '## 7. Copy, verbatim';
const cell4 = (label: string): string[] => specCell(SPEC4, HEADING4, label);
const str4 = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC4, HEADING4, label, index, fill);

const SPEC3 = new URL('../../../../design/specs/2026-11-05-guild-events.md', import.meta.url);
const HEADING3 = '## 5. Copy, verbatim';
const cell3 = (label: string): string[] => specCell(SPEC3, HEADING3, label);
const str3 = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC3, HEADING3, label, index, fill);

const CHICAGO = 'America/Chicago';
const NEW_YORK = 'America/New_York';
// Wednesday 9 Dec 2026, 10:00 am Chicago.
const NOW = new Date('2026-12-09T16:00:00Z');
const at = (isoUtc: string): Date => new Date(isoUtc);

const KINDS = ['raid', 'dungeon', 'event'] as const;

describe('raidCopy constants, verbatim from spec 6.1', () => {
  it('names the tab button and cap', () => {
    expect(raidCopy.planButton).toBe(str3('Tab button'));
    expect(raidCopy.cap).toBe(str('Cap'));
  });

  it('has every sub-head, the roster heads as the raid comp spec amends them', () => {
    const s = raidCopy.subheads;
    expect([s.planned, s.raidNights, s.yourSignup, s.signups, s.notes]).toEqual(
      cell('Sub-heads').slice(0, 5),
    );
    expect([s.tanks, s.healers, s.melee, s.ranged, s.late, s.bench, s.out, s.notSigned]).toEqual(
      cell4('Roster heads'),
    );
  });

  it('has the segments, placeholders and strip status', () => {
    const g = raidCopy.segments;
    expect([g.in, g.late, g.bench, g.out]).toEqual(cell('Segments'));
    expect(raidCopy.placeholders.note).toBe(str('Placeholders', 0));
    expect([raidCopy.placeholders.chooseWhat, raidCopy.placeholders.eventTitle]).toEqual(
      cell3('Placeholders'),
    );
    expect(raidCopy.stripStatus.saved).toBe(str('Strip status', 0));
    expect(raidCopy.stripStatus.failed).toBe(str('Strip status', 1));
    expect(raidCopy.stripStatus.movedToBench('Lightbrand')).toBe(
      str('Strip status', 2, { Name: 'Lightbrand' }),
    );
    expect(raidCopy.stripStatus.left(120)).toBe(str('Strip status', 3, { n: '120' }));
    expect(raidCopy.waitingCharacter).toBe(str('Waiting character'));
  });

  it('has the empty states', () => {
    expect(raidCopy.emptyPlanned.member).toBe(str3('Empty Planned', 0));
    expect(raidCopy.emptyPlanned.officer).toBe(str3('Empty Planned', 1));
    expect(raidCopy.bandNone.member).toBe(str('Band line, none planned', 0));
    expect(raidCopy.bandNone.officer).toBe(str('Band line, none planned', 1));
    expect(raidCopy.bandNone.planLink).toBe(str('Band line, none planned', 2));
    expect(notesEmptyOf('raid', false)).toBe(str('Notes empty', 0));
    expect(notesEmptyOf('raid', true)).toBe(str('Notes empty', 1));
    expect(notesEmptyOf('dungeon', true)).toBe('No notes yet. Edit the run to add some.');
    expect(notesEmptyOf('event', false)).toBe('No notes for this event.');
    const e = raidCopy.emptyGroup;
    expect([e.tank, e.healer, e.melee, e.ranged]).toEqual(cell4('Empty groups'));
  });

  it('has the raid states, races and poll line', () => {
    const r = raidCopy.raidStates;
    // The ended, done and open-night lines moved to the attendance copy (its spec 3.1).
    const [locked] = cell('Raid states');
    expect(r.locked).toBe(locked);
    expect(KINDS.map(cancelledLineOf)).toEqual(cell3('Cancelled'));
    expect(KINDS.map(endedLineOf)).toEqual(cell3('Ended'));
    // The raids spec's own copy table carries every kind's ended line too (guild loop round 3).
    expect(cell('Raid states')).toEqual(expect.arrayContaining(KINDS.map(endedLineOf)));
    const c = raidCopy.races;
    expect([c.locked, c.cancelled, c.officerToolsEnded]).toEqual(cell('Races'));
    expect(raidCopy.poll).toBe(str('Poll'));
  });

  it('has the officer row and the cancel confirm', () => {
    const o = raidCopy.officerRow;
    expect([o.edit, o.lock, o.locking, o.unlock, o.copy, o.copied, 'Cancel raid', o.failed]).toEqual(
      cell('Officer row'),
    );
    expect(KINDS.map(cancelActionOf)).toEqual(cell3('Officer row'));
    const k = raidCopy.cancelConfirm;
    const [, postToDiscord, , confirming, keep] = cell('Cancel confirm');
    expect([k.postToDiscord, k.confirming, k.keep]).toEqual([postToDiscord, confirming, keep]);
    // The spec's template holds pipes, which split a table cell, so it is read as quoted text.
    const question = 'Cancel this {raid|run|event}? Members see it as cancelled.';
    const confirm = 'Cancel the {raid|run|event}';
    const spec = readFileSync(SPEC3, 'utf8');
    expect(spec).toContain(`\`${question}\``);
    expect(spec).toContain(`\`${confirm}\``);
    const words = { raid: 'raid', dungeon: 'run', event: 'event' } as const;
    for (const kind of KINDS) {
      expect(cancelQuestionOf(kind)).toBe(question.replace('{raid|run|event}', words[kind]));
      expect(cancelConfirmOf(kind)).toBe(confirm.replace('{raid|run|event}', words[kind]));
    }
    expect(KINDS.map((kind) => eyebrowOf(kind).toUpperCase())).toEqual(cell3('Eyebrow'));
    expect(raidCopy.copyFailed).toBe(str('Copy failed'));
  });

  it('has the form labels, helpers and errors', () => {
    const f = raidCopy.form;
    expect(f.what).toBe(str3('Form label', 0));
    expect(f.groupSize).toBe(str3('Form label', 1));
    expect(f.noLimit).toBe(str3('Form label', 2));
    expect([f.groups.event, f.groups.dungeons, f.groups.raids]).toEqual(cell3('Groups'));
    expect(f.opens("Onyxia's Lair", 'Wed 9 Dec')).toBe(
      str3('Helper', 0, { Raid: "Onyxia's Lair", 'Wed 9 Dec': 'Wed 9 Dec' }),
    );
    expect(f.submitCreate).toBe(str3('Submit', 0));
    expect(f.submitCreating).toBe(str3('Submit', 1));
    expect(f.submitEdit).toBe(str3('Submit', 2));
    expect(f.submitEditing).toBe(str3('Submit', 3));
    const [, ...oldLabels] = cell('Form labels');
    expect(
      oldLabels.filter(
        (label) => label !== 'Post the raid' && !label.endsWith('…') && label !== 'Save changes',
      ),
    ).toEqual([
      f.title,
      f.date,
      f.time,
      f.duration,
      f.guildZone,
      f.notes,
      f.repeat,
      f.postToDiscord,
      f.postToDiscordIfMoved,
      f.cancel,
    ]);
    expect(f.forYou('Thu 10 Dec, 9:00 pm')).toBe(str('Form helpers', 0));
    expect(f.forYouSame).toBe(str('Form helpers', 1));
    expect(f.markdownHelper).toBe(str('Form helpers', 2));
    expect(f.onlyFit(2)).toBe(str('Form helpers', 3, { n: '2' }));
    expect(f.connectDiscord).toBe(str('Form helpers', 4));
    expect(f.connectDiscordLink).toBe('Settings');
    const e = f.errors;
    const [, ...oldErrors] = cell('Form errors');
    expect([e.date, e.time, e.cap, e.failed]).toEqual(oldErrors);
    expect(e.what).toBe(str3('Errors', 0));
    expect(e.eventTitle).toBe(str3('Errors', 1));
    expect(e.opens("Onyxia's Lair", 'Wed 9 Dec')).toBe(
      str3('Errors', 2, { Raid: "Onyxia's Lair", 'Wed 9 Dec': 'Wed 9 Dec' }),
    );
  });

  it('form strings match P.7', () => {
    const f = raidCopy.form;
    const ch = { channel: '#raid-signups' };
    const day = { Weekday: 'Thursday', d: '3', Month: 'December' };
    expect(f.remindDayBefore).toBe(str2('`form.remindDayBefore`'));
    expect(f.postNow('#raid-signups')).toBe(str2('`form.postNow`', 0, ch));
    expect(f.postNow(null)).toBe(str2('`form.postNow`', 1));
    expect(f.postLater('#raid-signups', 'Thursday 3 December')).toBe(
      str2('`form.postLater`', 0, { ...ch, ...day }),
    );
    expect(f.postLater(null, 'Thursday 3 December')).toBe(str2('`form.postLater`', 1, day));
    expect(f.postRepeat(true)).toBe(str2('`form.postRepeat`', 0));
    expect(f.postRepeat(false)).toBe(str2('`form.postRepeat`', 1));
    expect(f.postOff).toBe(str2('`form.postOff`'));
    expect(f.remind).toBe(str2('`form.remind`'));
    expect(f.remindNeedsPost).toBe(str2('`form.remindNeedsPost`'));
    expect(f.posted('#raid-signups')).toBe(str2('`form.posted`', 0, ch));
    expect(f.reminded).toBe(str2('`form.reminded`'));
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

  it('has the raid page lines of 3.2, 3.7 and 4, each quoted in the spec', () => {
    const spec = readFileSync(SPEC, 'utf8');
    const quoted = (words: string): boolean => spec.includes(`\`${words}\``);
    expect(quoted(raidCopy.page.eyebrow.toUpperCase())).toBe(true);
    expect(quoted(raidCopy.page.loading)).toBe(true);
    expect(quoted(raidCopy.page.signupFor('Lightbrand'))).toBe(true);
    expect(quoted(raidCopy.page.notSignedMark)).toBe(true);
  });
});

describe('whenPhrase', () => {
  const when = (starts: string, tz = CHICAGO, viewerTz = CHICAGO, now = NOW): string =>
    whenPhrase(at(starts), now, 'en-US', tz, viewerTz);

  it('says Tonight for the guild’s own day, in the guild’s clock and zone', () => {
    expect(when('2026-12-10T02:00:00Z')).toBe('Tonight 8 pm CST');
  });

  it('says Tomorrow for the next day', () => {
    expect(when('2026-12-11T02:00:00Z')).toBe('Tomorrow 8 pm CST');
  });

  it('names the weekday within six days', () => {
    expect(when('2026-12-12T02:00:00Z')).toBe('Friday 8 pm CST');
    expect(when('2026-12-15T02:00:00Z')).toBe('Monday 8 pm CST');
  });

  it('spells the date from the seventh day', () => {
    expect(when('2026-12-18T02:00:00Z')).toBe('Thu 17 Dec, 8 pm CST');
  });

  it('puts the guild clock first and the viewer’s own after it when the offsets differ (spec 3.4)', () => {
    // A Chicago guild's 8 pm raid, read in New York.
    expect(when('2026-12-11T02:00:00Z', CHICAGO, NEW_YORK)).toBe('Tomorrow 8 pm CST (9 pm EST your time)');
    expect(when('2026-12-11T02:00:00Z', CHICAGO, NEW_YORK)).toBe(
      `Tomorrow 8 pm CST${str4('Band `{when}`', 1)}`,
    );
  });

  it('says Today for a start before 5 pm on the guild clock, Tonight from 5 pm', () => {
    // NOW is 10 am in Chicago on Wed 9 Dec.
    expect(when('2026-12-09T18:49:00Z')).toBe('Today 12:49 pm CST');
    expect(when('2026-12-09T22:59:00Z')).toBe('Today 4:59 pm CST');
    expect(when('2026-12-09T23:00:00Z')).toBe('Tonight 5 pm CST');
    expect(when('2026-12-09T12:49:00Z', CHICAGO, CHICAGO, at('2026-12-09T08:00:00Z'))).toBe(
      'Today 6:49 am CST',
    );
    expect(str('`{when}`', 0)).toBe('Today 10 am');
  });

  it('writes minutes off the hour', () => {
    expect(when('2026-12-10T02:30:00Z')).toBe('Tonight 8:30 pm CST');
  });

  it('uses a 24-hour clock where the locale does', () => {
    expect(whenPhrase(at('2026-12-10T02:00:00Z'), NOW, 'de-DE', CHICAGO, CHICAGO)).toBe('Tonight 20:00 CST');
  });
});

describe('yourTimeLine', () => {
  it('is null when the offsets agree', () => {
    expect(yourTimeLine(at('2026-12-11T02:00:00Z'), CHICAGO, CHICAGO)).toBeNull();
  });

  it('reads the viewer’s start and zone, the spec’s words', () => {
    expect(yourTimeLine(at('2026-12-11T02:00:00Z'), CHICAGO, NEW_YORK)).toBe(
      str4('When, guild first', 1, { start: '9:00 pm', ZONE: 'EST' }),
    );
  });

  it('names the viewer’s weekday when their date is another', () => {
    // 8 pm Thursday in Chicago is 2 am Friday in London.
    expect(yourTimeLine(at('2026-12-11T02:00:00Z'), CHICAGO, 'Europe/London')).toBe(
      'Fri 2:00 am GMT your time',
    );
  });

  it('glues a zone to its clock', () => {
    expect(zonedClock(at('2026-12-11T02:00:00Z'), CHICAGO)).toBe('8:00 pm CST');
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
    expect(countLine(0, null)).toBe('Nobody signed up yet');
    expect(countLine(0, 40)).toBe('Nobody signed up yet');
  });

  it('writes the count clauses, singular at one (raid comp spec 7)', () => {
    expect([compCopy.tanks(1), compCopy.tanks(3), compCopy.healers(1), compCopy.healers(7)]).toEqual([
      '1 tank',
      '3 tanks',
      '1 healer',
      '7 healers',
    ]);
    expect([compCopy.melee(7), compCopy.ranged(5)]).toEqual(['7 melee', '5 ranged']);
  });

  it('names the late roles, a single part that is all of them without its number', () => {
    expect(compCopy.late(1, 0, 1)).toBe('1 late (healer)');
    expect(compCopy.late(2, 1, 1)).toBe('2 late (1 tank, 1 healer)');
    expect(compCopy.late(3, 0, 1)).toBe('3 late (1 healer)');
    expect(compCopy.late(2, 0, 0)).toBe('2 late');
    expect(compCopy.late(2, 0, 2)).toBe('2 late (healers)');
  });

  it('offers to show them all', () => {
    expect(showAll(7)).toBe(str('Show more', 0, { n: '7' }));
  });

  it('names the needs-line deadline', () => {
    expect(needsBefore('Thursday')).toBe('before Thursday');
    expect(needsBefore(null)).toBe('before the next raid');
  });

  it('takes the deadline weekday from the next raid in the viewer zone', () => {
    // 02:00 UTC Friday is Thursday evening in Chicago.
    expect(nextRaidDeadline('2026-12-11T02:00:00Z', 'en-US', CHICAGO)).toBe('before Thursday');
    expect(nextRaidDeadline('2026-12-11T02:00:00Z', 'en-US', 'Europe/Berlin')).toBe('before Friday');
    expect(nextRaidDeadline(undefined, 'en-US', CHICAGO)).toBe('before the next raid');
  });

  it('names the raid the API sends, on the guild calendar, whatever leads the band', () => {
    const onyxia = { title: "Onyxia's Lair", starts_at: '2026-12-11T02:00:00Z' };
    const run = { guild_timezone: CHICAGO } as Parameters<typeof needsDeadline>[1];
    expect(needsDeadline(onyxia, run, 'en-US', 'Europe/Berlin')).toBe("before Onyxia's Lair, Thu 10 Dec");
    expect(needsDeadline(onyxia, null, 'en-US', CHICAGO)).toBe("before Onyxia's Lair, Thu 10 Dec");
    expect(needsDeadline(null, null, 'en-US', CHICAGO)).toBe('before the next raid');
    expect(needsDeadline(undefined, undefined, 'en-US', CHICAGO)).toBe('before the next raid');
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
      str('Band line', 4, { when: 'Tonight 8 pm', n: '23', size: '40' }),
    );
    expect(bandLine({ ...base, state: null })).toBe('Tonight 8 pm, you have not signed up yet. 23 of 40 in.');
    expect(bandLine({ ...base, n: 1, size: 5, state: null })).toBe(
      'Tonight 8 pm, you have not signed up yet. 1 of 5 in.',
    );
  });

  it('says how many are in while unsigned and the size is unknown', () => {
    expect(bandLine({ ...base, size: null, state: null })).toBe(
      `Tonight 8 pm, you have not signed up yet. ${str('Band line', 5, { n: '23' })}`,
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
    expect(specCell(SPEC, HEADING, 'Band line')[7]).toBe('you are In on {Name} as {Role}');
  });
});

describe('bandLineParts', () => {
  const base = { when: 'Tonight 8 pm', n: 23, size: 40 as number | null };

  it('splits the sentence around its state word', () => {
    expect(bandLineParts({ ...base, state: 'in', role: 'dps' })).toEqual({
      head: 'Tonight 8 pm, you are ',
      word: 'In',
      tail: ' as DPS, 23 of 40 signed up.',
    });
    expect(bandLineParts({ ...base, state: 'bench', role: 'dps' })).toEqual({
      head: 'Tonight 8 pm, you are on the ',
      word: 'Bench',
      tail: ', 23 of 40 signed up.',
    });
    expect(bandLineParts({ ...base, state: 'late', role: 'tank', characterName: 'Late' })).toEqual({
      head: 'Tonight 8 pm, you are ',
      word: 'Late',
      tail: ' on Late as Tank, 23 of 40 signed up.',
    });
  });

  it('says nobody rather than 0 signed up', () => {
    expect(bandLineParts({ ...base, n: 0, state: 'out' }).tail).toBe(', nobody signed up yet.');
  });

  it('has no state word while unsigned', () => {
    const parts = bandLineParts({ ...base, state: null });
    expect(parts.word).toBe('');
    expect(parts.head + parts.tail).toBe(bandLine({ ...base, state: null }));
  });
});

describe('showsBandLine', () => {
  it('is for verified members and officers only (spec 3.4, R1)', () => {
    expect(showsBandLine(true, 'member')).toBe(true);
    expect(showsBandLine(true, 'officer')).toBe(true);
    expect(showsBandLine(false, 'member')).toBe(false);
    expect(showsBandLine(false, 'officer')).toBe(false);
    expect(showsBandLine(true, 'public')).toBe(false);
  });
});
