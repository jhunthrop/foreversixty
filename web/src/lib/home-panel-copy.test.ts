// web/src/lib/home-panel-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeHeroCardsCopy } from './home-panel-copy';

describe('homeHeroCardsCopy', () => {
  // Review round 1 item 7: one sentence per not-yet-available card, naming the real
  // reason -- shared verbatim between the Best in slot card and the "Your upgrades"
  // panel, since both name the identical live-worn-gear gap.
  it('carries the exact "Best in slot" / "Your upgrades" not-available sentence', () => {
    expect(homeHeroCardsCopy.bestInSlotNotAvailable).toBe(
      'Not available yet: the addon does not send worn gear.',
    );
  });

  it('carries the exact "Talents" not-available sentence', () => {
    expect(homeHeroCardsCopy.talentsNotAvailable).toBe('Not available yet: the addon does not send talents.');
  });
});
