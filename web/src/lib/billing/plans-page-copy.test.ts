// web/src/lib/billing/plans-page-copy.test.ts
import { describe, expect, it } from 'vitest';
import { specCell, specKeyCell, specLineAfter } from '../../test-support/spec-table';
import { exampleCaption, serverRunExample } from './premium-examples';
import { CAPS, LANE_ITERATION_CEILING } from '../sim/precision';
import { GUILD_TIER_MEMBER_LIMIT, GUILD_TRIAL_DAYS, PLANS, formatAmount, perMonthLine } from './plans';
import { plansPageCopy as copy, perMemberLine, perRaiderLine, serverRunCaption } from './plans-page-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
const row = (label: string): string[] => specCell(SPEC, '### 6.1', label);
const B5 = '### B5. Copy, verbatim';
const key = (k: string): string[] => specKeyCell(SPEC, B5, k);
const fill = (text: string, slots: Record<string, string>): string =>
  Object.entries(slots).reduce((t, [slot, value]) => t.replaceAll(`{${slot}}`, value), text);
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
      copy.free.signIn,
      ...copy.free.features,
      copy.free.promise,
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
      '{list} after 6 Jan 2027',
      '{list} a month after 6 Jan 2027',
      '{perMonth} a month, billed yearly',
      formatAmount(PLANS.premium.founderYearly / 12),
      formatAmount(PLANS.guild.founderYearly / 12),
      'Or {yearly price}.',
      copy.price.founder,
    ]);
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
  it('premium.features[0] from CAPS', () =>
    expect(copy.premium.features[0]).toBe(
      fill(key('premium.features[0]')[0] ?? '', {
        'CAPS.server': count(CAPS.server),
        'CAPS.browser': count(CAPS.browser),
      }),
    ));
  it('premium.features[1]', () => expect(copy.premium.features[1]).toBe(key('premium.features[1]')[0]));
  it('premium.features[2] from LANE_ITERATION_CEILING', () =>
    expect(copy.premium.features[2]).toBe(
      fill(key('premium.features[2]')[0] ?? '', {
        'LANE_ITERATION_CEILING.server': count(LANE_ITERATION_CEILING.server),
        'LANE_ITERATION_CEILING.browser': count(LANE_ITERATION_CEILING.browser),
      }),
    ));
  it.each([3, 4, 5])('premium.features[%i]', (i) =>
    expect(copy.premium.features[i]).toBe(key(`premium.features[${i}]`)[0]),
  );
  it('has exactly six features', () => expect(copy.premium.features).toHaveLength(6));
  it('premium.scope', () => expect(copy.premium.scope).toBe(key('premium.scope')[0]));
  it('held coach line is the spec line and is not in premium.features', () => {
    expect(copy.premium.featuresHeld.coach).toBe(specLineAfter(SPEC, '**Held until built'));
    expect(copy.premium.features).not.toContain(copy.premium.featuresHeld.coach);
    expect(copy.premium.features.join(' ')).not.toMatch(/coach/i);
  });
  it('old Premium features are gone', () =>
    expect(allStrings(copy)).not.toContain('History charts of your score and parses'));
});

describe('plansPageCopy guild region against spec B5', () => {
  const limit = String(GUILD_TIER_MEMBER_LIMIT);
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
    expect(copy.guild.features).toEqual(
      key('guild.features').map((s) => fill(s, { GUILD_TIER_MEMBER_LIMIT: limit })),
    ));
  it('raidGuild.audience is unchanged', () =>
    expect(copy.raidGuild.audience).toBe(key('raidGuild.audience')[0]));
  it('raidGuild.features, loot council first', () => {
    expect(copy.raidGuild.features).toEqual(key('raidGuild.features'));
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
  it('howItWorks', () => expect(copy.howItWorks).toBe(key('howItWorks')[0]));
  it('example label and link', () => {
    expect(copy.examples.serverRun.label).toBe(key('examples.serverRun.label')[0]);
    expect(copy.examples.serverRun.link).toBe(key('examples.serverRun.link')[0]);
  });
  it('caption.preset and caption.ownGear follow the B5 templates', () => {
    const f = {
      characterName: 'Kraggor',
      level: '60',
      specLabel: 'Fury Warrior',
      combinations: '4,812',
      ranAt: '27 Oct 2026',
    };
    const slots = { ...f, 'ranAt as d MMM yyyy': f.ranAt };
    expect(copy.examples.serverRun.caption.preset(f)).toBe(
      fill(key('examples.serverRun.caption.preset')[0] ?? '', slots),
    );
    expect(copy.examples.serverRun.caption.ownGear(f)).toBe(
      fill(key('examples.serverRun.caption.ownGear')[0] ?? '', slots),
    );
  });
  it('serverRunCaption reads the data file through the templates', () => {
    expect(serverRunCaption()).toBe(exampleCaption(serverRunExample, copy.examples.serverRun.caption));
  });
});
