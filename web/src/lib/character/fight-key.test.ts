import { describe, expect, it } from 'vitest';
import { fightKey } from './fight-key';

describe('fightKey', () => {
  it('is report_id:fight_index', () => {
    expect(fightKey('abcd1234efgh', 7)).toBe('abcd1234efgh:7');
  });
});
