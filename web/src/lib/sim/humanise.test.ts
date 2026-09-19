import { describe, expect, it } from 'vitest';
import { humanise } from './humanise';

describe('humanise', () => {
  it('is empty for an empty string', () => {
    expect(humanise('')).toBe('');
  });

  it('capitalises a single word with no underscore', () => {
    expect(humanise('attack')).toBe('Attack');
  });

  it('splits underscores into spaces and capitalises only the first word', () => {
    expect(humanise('rage_gain')).toBe('Rage gain');
    expect(humanise('attack_power')).toBe('Attack power');
  });

  it('leaves an already-capitalised first letter alone', () => {
    expect(humanise('Attack_power')).toBe('Attack power');
  });
});
