import { describe, expect, it } from 'vitest';
import { connectCopy } from './connect-copy';

describe('connectCopy success page', () => {
  it('names the next step in game, and the two places to go', () => {
    expect(connectCopy.successBody).toBe(
      'Connected. Go back to the companion; it is ready. Now log out once in game, and your characters appear on the site within a minute.',
    );
    expect(connectCopy.yourAccount).toBe('Your account');
    expect(connectCopy.getSetUp).toBe('Get set up');
  });
});
