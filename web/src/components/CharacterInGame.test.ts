// web/src/components/CharacterInGame.test.ts
// Task L8.2 ruling 5: the character page's "In game" section, three lines of reference copy.
// Plain data in: Character.svelte resolves the score and, for the owner, the last export time.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import { scoreCopy } from '../lib/score/copy';
import type { ForeverScore } from '../lib/score/score';
import CharacterInGame from './CharacterInGame.svelte';

const SCORED: ForeverScore = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 640, parses: 510, attendance: 180, discipline: 88 },
  experience: null,
};

describe('CharacterInGame', () => {
  it('says the addon shows the same number and tier in game', () => {
    const { body } = render(CharacterInGame, { props: { forever: SCORED, lastSeenAt: null } });
    expect(body).toContain('data-testid="character-in-game"');
    expect(body).toContain(scoreCopy.inGameHeading);
    expect(body).toContain(scoreCopy.inGameScoreLead);
    expect(body).toMatch(/data-testid="score-badge"[^>]*>1830</);
    expect(body).toContain(labels.tiers.epic.label);
  });

  it('with no score, says so in the no-logs words, never a zero', () => {
    const { body } = render(CharacterInGame, { props: { forever: null, lastSeenAt: null } });
    expect(body).toContain(scoreCopy.inGameNoScore);
    expect(body).not.toContain('score-badge');
    expect(body).not.toMatch(/>\s*0\s*</);
  });

  it('tells the player how to put the character into the planner, with the paste link', () => {
    const { body } = render(CharacterInGame, { props: { forever: SCORED, lastSeenAt: null } });
    expect(body).toContain(scoreCopy.inGamePlannerLead);
    expect(body).toContain(`href="${scoreCopy.inGamePasteHref}"`);
    expect(body).toContain(scoreCopy.inGamePasteLink);
  });

  it('shows when the site last saw the export only when it has a time (the owner, signed in)', () => {
    const without = render(CharacterInGame, { props: { forever: SCORED, lastSeenAt: null } }).body;
    expect(without).not.toContain('character-in-game-last-seen');

    const local = new Date(2026, 10, 4, 14, 3).toISOString();
    const withTime = render(CharacterInGame, { props: { forever: SCORED, lastSeenAt: local } }).body;
    expect(withTime).toContain('data-testid="character-in-game-last-seen"');
    expect(withTime).toContain('Last seen in game 4 Nov 14:03');
  });
});
