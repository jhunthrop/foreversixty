// web/src/lib/home-panel-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeHeroCardsCopy, homePanelCopy } from './home-panel-copy';

describe('homeHeroCardsCopy', () => {
  // Review round 1 item 7: one sentence per not-yet-available card, naming the real
  // reason -- shared verbatim between the Best in slot card and the "Your upgrades"
  // panel, since both name the identical live-worn-gear gap. Reworded (home rebuild spec
  // §4, "Empty states stay honest") once `MeCharacter.build.gear`/`.talents` existed to
  // compare against: the gap is now this one character's own export, not a site-wide
  // limitation, so the sentence names the character's own next step.
  it('carries the exact "Best in slot" / "Your upgrades" not-available sentence', () => {
    expect(homeHeroCardsCopy.bestInSlotNotAvailable).toBe(
      'Not available yet: no gear export for this character. Open the addon once to send it.',
    );
  });

  it('carries the exact "Talents" not-available sentence', () => {
    expect(homeHeroCardsCopy.talentsNotAvailable).toBe(
      'Not available yet: no talent export for this character. Open the addon once to send it.',
    );
  });
});

describe('the Simulator card copy', () => {
  it('puts the unit on both figures: "26 DPS now · 37.5 DPS at band best in slot"', () => {
    const together = `${homeHeroCardsCopy.simulatorNowFigure('26')} · ${homeHeroCardsCopy.simulatorBandLine(37.46)}`;
    expect(together).toBe('26 DPS now · 37.5 DPS at band best in slot');
  });
});

describe('homePanelCopy', () => {
  it('points at the header selector as plain text, and names the current character once', () => {
    expect(homePanelCopy.changeCharacterHint).toBe('Change character: top right');
    expect(homePanelCopy.currentCharacterEyebrow).toBe('Current character');
    expect(homePanelCopy.yourCharacterEyebrow).toBe('Your character');
  });
});
