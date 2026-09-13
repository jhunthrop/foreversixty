import { describe, expect, it } from 'vitest';
import { assertLinksAreReal, placeholderKeys } from './links';

const real = {
  discordInvite: 'https://discord.gg/foreversixty',
  githubRepo: 'https://github.com/foreversixty/forever',
};

describe('placeholderKeys', () => {
  it('names every key that is still a placeholder', () => {
    expect(placeholderKeys({ ...real, discordInvite: 'https://discord.gg/PLACEHOLDER' })).toEqual([
      'discordInvite',
    ]);
    expect(placeholderKeys(real)).toEqual([]);
  });
});

describe('assertLinksAreReal', () => {
  it('names the file and the offending keys', () => {
    expect(() =>
      assertLinksAreReal({ ...real, githubRepo: 'https://github.com/PLACEHOLDER/forever' }),
    ).toThrow(/web\/src\/data\/links\.json still contains PLACEHOLDER links: githubRepo/);
  });

  it('passes once every link is real', () => {
    expect(() => assertLinksAreReal(real)).not.toThrow();
  });
});
