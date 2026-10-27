// web/src/lib/billing/plans-page-copy.test.ts
import { describe, expect, it } from 'vitest';
import { specCell } from '../../test-support/spec-table';
import { plansPageCopy as copy } from './plans-page-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
const row = (label: string): string[] => specCell(SPEC, '### 6.1', label);

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
    expect(row('Founder row')).toContain(copy.founder.before);
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

  it('Free summary', () => {
    const { name, price, line, srOnly } = copy.free.summary;
    expect([name, price, line, srOnly]).toEqual(row('Free summary'));
  });

  it('Premium card', () => {
    const [name, audience, ...features] = row('Premium card');
    expect([name, audience]).toEqual([copy.premium.name, copy.premium.audience]);
    expect(features.slice(0, copy.premium.features.length)).toEqual(copy.premium.features);
    expect(row('Premium card').at(-1)).toBe(copy.premium.ships('{Month}'));
  });

  it('Guild card', () => {
    const { name, audience, features } = copy.guild;
    expect([name, audience, ...features]).toEqual(row('Guild card'));
  });

  it('Raid guild card', () => {
    const { name, audience, features } = copy.raidGuild;
    expect([name, audience, ...features]).toEqual(row('Raid guild card'));
  });

  it('Price block', () => {
    expect(row('Price block')).toEqual([
      copy.price.suffix.yearly,
      copy.price.suffix.monthly,
      '{list}',
      '{perMonth} a month, billed yearly',
      '$3.33',
      '$10',
      'Or {yearly price}.',
      copy.price.founder,
    ]);
    expect(copy.price.perMonth(40)).toBe('$3.33 a month, billed yearly');
    expect(copy.price.orYearly('$40 a year')).toBe('Or $40 a year.');
  });

  it('Buttons', () => {
    expect(row('Buttons')).toEqual(['Subscribe · {price}', 'Subscribe a guild · {price}']);
    expect(copy.buttons.subscribe('{price}')).toBe('Subscribe · {price}');
    expect(copy.buttons.subscribeGuild('{price}')).toBe('Subscribe a guild · {price}');
  });

  it('Trial', () => {
    expect([copy.trial]).toEqual(row('Trial'));
  });

  it('History promise', () => {
    expect([copy.historyPromise]).toEqual(row('History promise'));
  });

  it('How it works', () => {
    const [heading, ...sentences] = row('How it works');
    expect(heading).toBe(copy.howItWorks);
    expect(copy.rules).toHaveLength(8);
    expect(copy.rules.flatMap((rule) => [rule.term, rule.sentence])).toEqual(sentences);
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
