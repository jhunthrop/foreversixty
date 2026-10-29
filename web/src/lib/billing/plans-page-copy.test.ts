// web/src/lib/billing/plans-page-copy.test.ts
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { fillSlots, specCell, specKeyCell } from '../../test-support/spec-table';
import { coachCopy } from '../coach/copy';
import { CAPS, LANE_ITERATION_CEILING } from '../sim/precision';
import {
  FOUNDER_PRICE_UNTIL,
  GUILD_TIER_MEMBER_LIMIT,
  GUILD_TRIAL_DAYS,
  PLANS,
  RAID_SIZES,
  REFUND_WINDOW_DAYS,
  RENEWAL_REMINDER_DAYS,
  TRIAL_REMINDER_DAYS,
  formatAmount,
  perMonthLine,
} from './plans';
import { plansPageCopy as copy, perMemberLine, perRaiderLine } from './plans-page-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
const row = (label: string): string[] => specCell(SPEC, '### 6.1', label);
const POLISH_SPEC = new URL('../../../../design/specs/2026-10-28-polish.md', import.meta.url);
const B5 = '### B5. Copy, verbatim';
const key = (k: string): string[] => specKeyCell(SPEC, B5, k);
const fill = fillSlots;
const count = (n: number): string => n.toLocaleString('en-US');

function allStrings(value: unknown): string[] {
  if (typeof value === 'string') return [value];
  if (Array.isArray(value)) return value.flatMap(allStrings);
  if (typeof value === 'object' && value !== null) return Object.values(value).flatMap(allStrings);
  return [];
}

describe('plansPageCopy against spec 6.1', () => {
  it('Page', () => {
    expect([copy.title, copy.description]).toEqual(row('Page'));
  });

  it('Band', () => {
    expect([copy.band.eyebrow, copy.band.heading, copy.band.lede]).toEqual(row('Band'));
  });

  it('Founder row', () => {
    const [label, open] = row('Founder row');
    expect(label).toBe(copy.founder.label);
    expect(open).toBe(copy.founder.open);
  });

  it('Interval', () => {
    expect([copy.interval.labels.yearly, copy.interval.labels.monthly, copy.interval.groupLabel]).toEqual(
      row('Interval'),
    );
  });

  it('Free card', () => {
    expect(row('Free card')).toEqual([
      copy.free.name,
      copy.free.audience,
      copy.free.price,
      copy.free.note,
      ...copy.free.features,
    ]);
  });

  it('Premium card name and audience (the §6.1 features are overridden by B5)', () => {
    const [name, audience] = row('Premium card');
    expect([name, audience]).toEqual([copy.premium.name, copy.premium.audience]);
  });

  it('Price block', () => {
    expect(row('Price block')).toEqual([
      copy.price.suffix.yearly,
      copy.price.suffix.monthly,
      'Founder price until 6 Jan 2027, then {list}.',
      'Founder price until 6 Jan 2027, then {list} a month.',
      '{perMonth} a month, billed yearly',
      formatAmount(PLANS.premium.founderYearly / 12),
      formatAmount(PLANS.guild.founderYearly / 12),
      'Or {yearly price}.',
    ]);
    expect(copy.price.founderUntil('{list}')).toBe('Founder price until 6 Jan 2027, then {list}.');
    expect(copy.price.perMonth(PLANS.premium.founderYearly)).toBe(perMonthLine(PLANS.premium.founderYearly));
    expect(copy.price.orYearly('{yearly price}')).toBe('Or {yearly price}.');
  });

  it('Buttons', () => {
    expect(row('Buttons')).toEqual(['Subscribe · {price}', 'Subscribe a guild · {price}']);
    expect(copy.buttons.subscribe('{price}')).toBe('Subscribe · {price}');
    expect(copy.buttons.subscribeGuild('{price}')).toBe('Subscribe a guild · {price}');
  });

  it('History promise', () => {
    expect([copy.historyPromise]).toEqual(row('History promise'));
  });

  it('Foot', () => {
    expect([copy.foot]).toEqual(row('Foot'));
  });

  it('FAQ edits', () => {
    const [cancelQuestion, cancelAnswer] = row('FAQ edits');
    expect(copy.faq[0]).toEqual({ question: cancelQuestion, answer: cancelAnswer });
    expect(copy.faq.find((entry) => entry.question === 'Are there ads?')?.answer).toContain(
      'Premium and the guild plans',
    );
  });
});

describe('plansPageCopy guard', () => {
  it('has no retention talk, no comparison caps and no exclamation marks', () => {
    for (const text of allStrings(copy)) {
      expect(text).not.toMatch(/retention|compare more than/i);
      expect(text).not.toContain('!');
    }
  });
});

describe('plansPageCopy Premium against spec B5', () => {
  it('premium.support', () => expect(copy.premium.support).toBe(key('premium.support')[0]));
  it('the coach line is the first Premium benefit', () => {
    expect(copy.premium.features[0]?.text).toBe(key('premium.features[0]')[0]);
    expect(copy.premium.features[0]?.text).toMatch(/^Coach: after every logged raid, /);
  });
  it('premium.features[1] from CAPS', () =>
    expect(copy.premium.features[1]?.text).toBe(
      fill(key('premium.features[1]')[0] ?? '', {
        'CAPS.server': count(CAPS.server),
        'CAPS.browser': count(CAPS.browser),
      }),
    ));
  it('premium.features[2]', () => expect(copy.premium.features[2]?.text).toBe(key('premium.features[2]')[0]));
  it('premium.features[3] from LANE_ITERATION_CEILING', () =>
    expect(copy.premium.features[3]?.text).toBe(
      fill(key('premium.features[3]')[0] ?? '', {
        'LANE_ITERATION_CEILING.server': count(LANE_ITERATION_CEILING.server),
        'LANE_ITERATION_CEILING.browser': count(LANE_ITERATION_CEILING.browser),
      }),
    ));
  it.each([4, 5, 6])('premium.features[%i]', (i) =>
    expect(copy.premium.features[i]?.text).toBe(key(`premium.features[${i}]`)[0]),
  );
  it('has exactly seven features', () => expect(copy.premium.features).toHaveLength(7));
  it('premium.scope says the coach covers every spec', () =>
    expect(copy.premium.scope).toBe(
      'Sims cover damage specs; healing and tanking specs are not simulated yet. The coach reviews every spec; damage specs get the most findings.',
    ));
  it('premium.scope matches the spec', () => expect(copy.premium.scope).toBe(key('premium.scope')[0]));
  it('nothing is held any more', () => expect('featuresHeld' in copy.premium).toBe(false));
  it('old Premium features are gone', () =>
    expect(allStrings(copy)).not.toContain('History charts of your score and parses'));
});

describe('plansPageCopy guild region against spec B5', () => {
  const limit = String(GUILD_TIER_MEMBER_LIMIT);
  const featureSlots = {
    GUILD_TIER_MEMBER_LIMIT: limit,
    'PLANS.guild.seats': String(PLANS.guild.seats),
    'PLANS.raid_guild.seats': String(PLANS.raid_guild.seats),
    'RAID_SIZES[1]': String(RAID_SIZES[1]),
    'RAID_SIZES[2]': String(RAID_SIZES[2]),
  };
  const trialSlots = {
    GUILD_TRIAL_DAYS: String(GUILD_TRIAL_DAYS),
    'GUILD_TRIAL_DAYS + 1': String(GUILD_TRIAL_DAYS + 1),
  };
  const rule = (term: string): string | undefined => copy.rules.find((r) => r.term === term)?.sentence;

  it('band guild link', () =>
    expect([copy.band.guildLead, copy.band.guildLink]).toEqual(key('band.guildLead')));
  it.each(['heading', 'lede', 'members'] as const)('guildRegion.%s', (k) =>
    expect(copy.guildRegion[k]).toBe(key(`guildRegion.${k}`)[0]),
  );
  it('guild.audience', () =>
    expect(copy.guild.audience).toBe(
      fill(key('guild.audience')[0] ?? '', { GUILD_TIER_MEMBER_LIMIT: limit }),
    ));
  it('guild.features', () =>
    expect(copy.guild.features).toEqual(key('guild.features').map((s) => fill(s, featureSlots))));
  it('raidGuild.audience is unchanged', () =>
    expect(copy.raidGuild.audience).toBe(key('raidGuild.audience')[0]));
  it('raidGuild.features, loot council first', () => {
    expect(copy.raidGuild.features).toEqual(key('raidGuild.features').map((s) => fill(s, featureSlots)));
    expect(copy.raidGuild.features[0]).toMatch(/^Loot council:/);
    expect(copy.raidGuild.features[0]).not.toBe('Everything in Guild');
  });
  it.each([
    ['yearly', 1],
    ['monthly', 3],
  ] as const)('perMember %s founder matches the B5 example', (interval, idx) =>
    expect(perMemberLine(interval, true)).toBe(key('price.perMember(interval, founder)')[idx]),
  );
  it.each([
    ['yearly', 1],
    ['monthly', 3],
  ] as const)('perRaider %s founder matches the B5 example', (interval, idx) =>
    expect(perRaiderLine(interval, true)).toBe(key('price.perRaider(interval, founder)')[idx]),
  );
  it('perRaider monthly list carries $0.63', () =>
    expect(perRaiderLine('monthly', false)).toBe(
      '$0.63 per raider a month in a 40-person raid, $1.25 in a 20, $2.50 in a 10.',
    ));
  it('trial', () => expect(copy.trial).toBe(fill(key('trial')[0] ?? '', trialSlots)));
  it('rule order: Verified members first, Paying together after Size', () => {
    const terms = copy.rules.map((r) => r.term);
    expect(terms[0]).toBe('Verified members');
    expect(terms[terms.indexOf('Size') + 1]).toBe('Paying together');
  });
  it('Verified members rule is B5 verbatim and has no Battle.net', () => {
    expect([copy.rules[0]?.term, copy.rules[0]?.sentence]).toEqual(key('rules[Verified members]'));
    expect(copy.rules[0]?.sentence).not.toMatch(/Battle\.net/);
  });
  it('Paying together rule', () =>
    expect(['Paying together', rule('Paying together')]).toEqual(key('rules[Paying together]')));
  it('Trial rule today', () => expect(rule('Trial')).toBe(fill(key('rules[Trial]')[0] ?? '', trialSlots)));
  it('trialEmail is B5 verbatim', () =>
    expect(copy.trialEmail).toBe(
      fill(key('trialEmail')[0] ?? '', { TRIAL_REMINDER_DAYS: String(TRIAL_REMINDER_DAYS) }),
    ));
  it('trial line and Trial rule make no email promise until the trial email ships', () => {
    expect(copy.trial).not.toContain('email you');
    expect(rule('Trial')).not.toContain('email you');
  });
  it('the unchanged rules and the plan names still match §6.1, in order', () => {
    const [heading, ...cells] = row('How it works');
    expect(heading).toBe(copy.howItWorks);
    const spec61 = Array.from({ length: cells.length / 2 }, (_, i) => ({
      term: cells[2 * i],
      sentence: cells[2 * i + 1],
    })).filter((r) => r.term !== 'Trial');
    const b5Terms = ['Verified members', 'Paying together', 'Trial'];
    expect(copy.rules).toHaveLength(10);
    expect(copy.rules.filter((r) => !b5Terms.includes(r.term))).toEqual(spec61);
    expect(copy.guild.name).toBe(row('Guild card')[0]);
    expect(copy.raidGuild.name).toBe(row('Raid guild card')[0]);
  });
  it('howItWorks', () => expect(copy.howItWorks).toBe(key('howItWorks')[0]));
  it('the sample coach review follows polish spec 3.6', () => {
    const cell = (k: string): string[] => specCell(POLISH_SPEC, '### 3.6', `\`coachSample.${k}`);
    expect(copy.coachSample.tag).toBe(cell('tag')[0]);
    expect(copy.coachSample.label).toBe(cell('label')[0]);
    expect(copy.coachSample.caption).toBe(cell('caption')[0]);
    expect(copy.coachSample.ariaLabel).toBe(cell('ariaLabel')[0]);
    const [kill] = cell('context(');
    const slots = { spec: 'Fury', cls: 'Warrior', zone: 'Molten Core', size: '40' };
    expect(copy.coachSample.context('Fury', 'Warrior', 'Molten Core', 40, true)).toBe(
      fill(kill ?? '', slots),
    );
    expect(copy.coachSample.context('Fury', 'Warrior', 'Molten Core', 40, false)).toBe(
      fill(kill ?? '', slots).replace('kill', 'wipe'),
    );
    expect(copy.coachSample.fixFirst).toBe(coachCopy.labelFixFirst);
    expect(copy.coachSample.then).toBe(coachCopy.labelThen);
    expect(copy.coachSample.doingWell).toBe(coachCopy.labelWell);
  });
  it('has no server-run example strings', () => {
    expect(copy).not.toHaveProperty('examples');
  });
});

describe('typed figures come from the plan constants', () => {
  const specText = readFileSync(SPEC, 'utf8');
  const ruleSentence = (term: string): string => copy.rules.find((r) => r.term === term)?.sentence ?? '';

  it.each([
    ['founder date', `Founder prices until ${FOUNDER_PRICE_UNTIL}`, copy.founder.open],
    [
      'renewal reminder',
      `We email ${RENEWAL_REMINDER_DAYS} days before every yearly renewal`,
      ruleSentence('Renewals'),
    ],
    ['refund window', `Ask within ${REFUND_WINDOW_DAYS} days of a first charge`, ruleSentence('Refunds')],
  ])('%s: the spec states the constant and the copy renders it', (_name, phrase, rendered) => {
    expect(specText).toContain(phrase);
    expect(rendered).toContain(phrase);
  });

  it('the copy source writes none of the three figures as a literal', () => {
    const source = readFileSync(new URL('./plans-page-copy.ts', import.meta.url), 'utf8');
    expect(source).not.toMatch(/6 January 2027|email 7 days|within 14 days/);
  });
});
