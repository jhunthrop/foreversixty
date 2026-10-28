import { describe, expect, it } from 'vitest';
import { parseInterval, parsePlan } from './plan-params';

describe('parsePlan', () => {
  it.each(['premium', 'guild', 'raid_guild'] as const)('keeps %s', (key) => {
    expect(parsePlan(key)).toBe(key);
  });

  it.each([null, undefined, '', 'Guild', 'raid-guild', 'enterprise', 'guild '])('drops %j', (raw) => {
    expect(parsePlan(raw)).toBeUndefined();
  });
});

describe('parseInterval', () => {
  it.each(['monthly', 'yearly'] as const)('keeps %s', (key) => {
    expect(parseInterval(key)).toBe(key);
  });

  it.each([null, undefined, '', 'weekly', 'Yearly', 'year'])('drops %j', (raw) => {
    expect(parseInterval(raw)).toBeUndefined();
  });
});
