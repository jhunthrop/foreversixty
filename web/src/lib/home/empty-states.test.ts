// web/src/lib/home/empty-states.test.ts
// P.1: a landing list with nothing in it says so in words (never a skeleton), with the exact
// sentence and the ghost button the review asked for.
import { describe, expect, it } from 'vitest';
import { homeTopGuildsCopy } from '../rankings/copy';

describe('landing empty states (P.1)', () => {
  it('Top guilds empty reads the claim sentence and offers the guild page', () => {
    expect(homeTopGuildsCopy.empty).toBe('No guilds ranked yet. Claim yours from the Guild menu.');
    expect(homeTopGuildsCopy.emptyAction).toBe('Open the guild page');
    expect(homeTopGuildsCopy.emptyHref).toBe('/guild');
  });
});
