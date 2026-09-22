import { describe, expect, it } from 'vitest';
import { BUSY_CLASS } from './busy';

describe('BUSY_CLASS', () => {
  it('is the one shared busy look: dimmed, progress cursor, nothing else', () => {
    expect(BUSY_CLASS).toBe('opacity-60 cursor-progress');
  });
});
