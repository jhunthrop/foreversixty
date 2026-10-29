import { describe, expect, it } from 'vitest';
import { bandCompareCopy } from './copy';

describe('bandCompareCopy.noLevelBuild', () => {
  it('names the level and points at the class’s other specs', () => {
    expect(bandCompareCopy.noLevelBuild(9, 'Arms', ['Fury', 'Protection'])).toBe(
      'No level 9 build for Arms. Pick Fury or Protection.',
    );
  });

  it('handles one, and no, other spec', () => {
    expect(bandCompareCopy.noLevelBuild(9, 'Arms', ['Fury'])).toBe('No level 9 build for Arms. Pick Fury.');
    expect(bandCompareCopy.noLevelBuild(9, 'Arms', [])).toBe('No level 9 build for Arms.');
  });
});
