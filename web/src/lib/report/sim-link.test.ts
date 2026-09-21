// web/src/lib/report/sim-link.test.ts
import { describe, expect, it } from 'vitest';
import { simLinkFor } from './sim-link';

describe('simLinkFor', () => {
  it('builds a fight-sourced sim link scoped to one combatant, labelled Sim', () => {
    const link = simLinkFor('abc2defg2hij', 3, 'Player-4184-000000A1');
    expect(link.href).toBe('/sim?source=fight&ref=abc2defg2hij%3A3%3APlayer-4184-000000A1');
    expect(link.label).toBe('Sim');
  });
});
