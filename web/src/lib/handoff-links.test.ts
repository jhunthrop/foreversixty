import { describe, expect, it } from 'vitest';
import { plannerCodeHref, simCodeHref, simFightHref } from './handoff-links';

describe('handoff-links', () => {
  it('builds a planner link carrying the FS1 code, URL-encoded', () => {
    expect(plannerCodeHref('FS1:1.60:warrior:human:0/0/0:')).toBe(
      '/planner?code=FS1%3A1.60%3Awarrior%3Ahuman%3A0%2F0%2F0%3A',
    );
  });

  it('builds a simulator link carrying the FS1 code, URL-encoded', () => {
    expect(simCodeHref('FS1:1.60:warrior:human:0/0/0:')).toBe(
      '/sim?code=FS1%3A1.60%3Awarrior%3Ahuman%3A0%2F0%2F0%3A',
    );
  });

  it('builds a fight-sourced sim link for the whole fight when no guid is given', () => {
    expect(simFightHref('abc2defg2hij', 3)).toBe('/sim?source=fight&ref=abc2defg2hij%3A3');
  });

  it('builds a fight-sourced sim link scoped to one combatant when a guid is given', () => {
    expect(simFightHref('abc2defg2hij', 3, 'Player-4184-000000A1')).toBe(
      '/sim?source=fight&ref=abc2defg2hij%3A3%3APlayer-4184-000000A1',
    );
  });
});
