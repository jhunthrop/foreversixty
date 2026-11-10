import { describe, expect, it } from 'vitest';
import { homeTopGuildsCopy } from './copy';

describe('homeTopGuildsCopy.bosses', () => {
  it('reads "1 boss" for one and "bosses" otherwise', () => {
    expect(homeTopGuildsCopy.bosses(1)).toBe('1 boss');
    expect(homeTopGuildsCopy.bosses(0)).toBe('0 bosses');
    expect(homeTopGuildsCopy.bosses(12)).toBe('12 bosses');
    expect(homeTopGuildsCopy.bossesDown(1)).toBe('1 boss down');
    expect(homeTopGuildsCopy.bossesDown(4)).toBe('4 bosses down');
  });
});
