// web/src/lib/home-panel-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeHeroCardsCopy, homePanelCopy } from './home-panel-copy';

describe('homeHeroCardsCopy', () => {
  // Review round 1 item 7: one sentence per not-yet-available card, naming the real
  // reason -- shared verbatim between the header and the "Your upgrades"
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
  it('puts the unit on both figures: "26 DPS now · 37.5 DPS at best in slot for your level range"', () => {
    const together = `${homeHeroCardsCopy.simulatorNowFigure('26')} · ${homeHeroCardsCopy.simulatorBandLine(37.46)}`;
    expect(together).toBe('26 DPS now · 37.5 DPS at best in slot for your level range');
  });
});

describe('homePanelCopy', () => {
  it('carries the header answer words from the character header spec', () => {
    expect(homePanelCopy.headerAnswerLabel).toBe('Upgrades for your level range');
    expect(homePanelCopy.headerNote('60')).toBe('Best in slot for 60, together');
    expect(homePanelCopy.headerAllMatch('60')).toBe('Every slot matches the 60 list.');
    expect(homePanelCopy.headerAction).toBe('See your upgrades →');
  });

  it('names the slots left out of the gain, in the singular for one', () => {
    expect(homePanelCopy.headerNotSimChecked(1)).toBe('one slot not sim-checked');
    expect(homePanelCopy.headerNotSimChecked(3)).toBe('3 slots not sim-checked');
  });
});
