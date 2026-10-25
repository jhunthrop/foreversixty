import { describe, expect, it } from 'vitest';
import { historyCopy } from './history-copy';

describe('historyCopy (spec section 6, verbatim)', () => {
  it('has the fixed strings', () => {
    expect(historyCopy.heading).toBe('History');
    expect(historyCopy.caption).toBe('Includes your private and guild reports. Only you see this section.');
    expect(historyCopy.scoreHeading).toBe('Score over time');
    expect(historyCopy.parseHeading).toBe('Parse by boss');
    expect(historyCopy.loadingSr).toBe('Reading your history…');
    expect(historyCopy.loadFailed).toBe('Your history did not load.');
    expect(historyCopy.empty).toBe('No history yet. Each rated fight you upload adds a point.');
    expect(historyCopy.parseOneKill).toBe('One kill so far. A line needs two.');
    expect(historyCopy.parseNoKills).toBe('No kills on record yet. Parse charts start with your first kill.');
    expect(`${historyCopy.freeOwnerSentence} ${historyCopy.freeOwnerLink}.`).toBe(
      'History charts draw your score and parses over time. See plans.',
    );
    expect(historyCopy.showFewer).toBe('Show fewer');
    expect(historyCopy.scorePlotLabel).toBe('Score chart. Use the arrow keys to move between points.');
    expect(historyCopy.scoreListLabel).toBe('Fights behind the score');
    expect(historyCopy.srScoreCaption).toBe('Forever Score by date, oldest first');
    expect(historyCopy.srParseCaption).toBe('Parse by kill, oldest first');
    expect(historyCopy.step.score).toEqual({ earlier: 'Earlier point', later: 'Later point' });
    expect(historyCopy.step.parse).toEqual({ earlier: 'Earlier kill', later: 'Later kill' });
  });

  it('writes the score note with and without rings', () => {
    expect(historyCopy.scoreNote(false)).toBe("One point a day: the day's last rated fight.");
    expect(historyCopy.scoreNote(true)).toBe(
      "One point a day: the day's last rated fight. Rings are private or guild reports.",
    );
  });

  it('writes counts and details', () => {
    expect(historyCopy.pillName('Golemagg', 5)).toBe('Golemagg, 5 kills');
    expect(historyCopy.pillName('Gehennas', 1)).toBe('Gehennas, 1 kill');
    expect(historyCopy.pillBoss('Golemagg', 'Hard')).toBe('Golemagg · Hard');
    expect(historyCopy.scoreDetail('11 Dec 2026', 2)).toBe('11 Dec 2026 · 2 fights');
    expect(historyCopy.scoreDetail('11 Dec 2026', 1)).toBe('11 Dec 2026 · 1 fight');
    expect(historyCopy.parseDetail('11 Dec 2026', 5, 5)).toBe('11 Dec 2026 · kill 5 of 5');
    expect(historyCopy.scoreReadoutText(1830, 'Epic', '11 Dec 2026 · 2 fights')).toBe(
      '1830 Epic · 11 Dec 2026 · 2 fights',
    );
  });

  it('writes the net change, gaps, labels and buttons', () => {
    expect(historyCopy.netChange(1090, '12 Oct')).toBe('Up 1090 since 12 Oct');
    expect(historyCopy.netChange(-40, '12 Oct')).toBe('Down 40 since 12 Oct');
    expect(historyCopy.netChange(0, '12 Oct')).toBe('No change since 12 Oct');
    expect(historyCopy.gapNote('20 Jul', '9 Aug', 'No rated fights')).toBe(
      'Gap 20 Jul to 9 Aug: No rated fights',
    );
    expect(historyCopy.srGapRow('26 Oct', '17 Nov', 'No rated fights')).toBe(
      'Gap from 26 Oct to 17 Nov: No rated fights',
    );
    expect(historyCopy.parsePlotLabel('Golemagg')).toBe(
      'Parse chart for Golemagg. Use the arrow keys to move between kills.',
    );
    expect(historyCopy.parseListLabel('Golemagg')).toBe('Kills of Golemagg');
    expect(historyCopy.showAll(31, 'fights')).toBe('Show all 31 fights');
    expect(historyCopy.showAll(12, 'kills')).toBe('Show all 12 kills');
    expect(historyCopy.scoreOnePoint('11 Dec')).toBe(
      'First score recorded 11 Dec. A line needs a second point.',
    );
  });

  it('names the visibility words', () => {
    expect(historyCopy.visibilityWord('private')).toBe('Private');
    expect(historyCopy.visibilityWord('guild')).toBe('Guild');
    expect(historyCopy.visibilityWord('public')).toBeNull();
  });
});
