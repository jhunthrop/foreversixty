// web/src/lib/creators/model.test.ts
import { describe, expect, it } from 'vitest';
import { PLANS } from '../billing/plans';
import { PROGRAM_TERMS, programCopy } from './copy';
import { earningsRows, firstYearEarnings } from './model';

describe('a player in a creator first year', () => {
  it('is the bounty plus half of a yearly payment', () => {
    expect(firstYearEarnings('premium', 'yearly', true)).toBe(
      PROGRAM_TERMS.bountyUsd + PLANS.premium.founderYearly / 2,
    );
  });

  it('is the bounty plus half of twelve monthly payments', () => {
    expect(firstYearEarnings('premium', 'monthly', false)).toBe(
      PROGRAM_TERMS.bountyUsd + (PLANS.premium.monthly * 12) / 2,
    );
  });

  it('labels each row with its plan and price and gives the dollars', () => {
    const rows = earningsRows(true);
    expect(rows.map((r) => r.label)).toEqual([
      `Premium, $${PLANS.premium.founderYearly} a year`,
      `Premium, $${PLANS.premium.founderMonthly} a month`,
      `Raid guild, $${PLANS.raid_guild.founderYearly} a year`,
    ]);
    expect(rows[0]?.value).toBe(`$${PROGRAM_TERMS.bountyUsd + PLANS.premium.founderYearly / 2}`);
  });
});

describe('the creators copy', () => {
  it('has no em dashes and no exclamation marks', () => {
    const text = JSON.stringify(programCopy);
    expect(text).not.toContain('—');
    expect(text).not.toContain('!');
  });

  it('every section figure lane names its section', () => {
    const ids = programCopy.sections.map((s) => s.id);
    for (const lane of programCopy.lanes) expect(ids).toContain(lane.section);
  });
});
