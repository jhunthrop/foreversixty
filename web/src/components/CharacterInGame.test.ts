// web/src/components/CharacterInGame.test.ts
// Task L8.2 ruling 5 and fix round 1 item 5: the character page's "In game" section -- a mock
// of the game tooltip carrying the addon's own score line, one line on where it shows, and for
// the owner only, how to copy the character in and when the site last saw its export.
// Plain data in: Character.svelte resolves the score, ownership and the export time.
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
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: null,
};

const BASE = { name: 'Elyra Duskvale', className: 'Priest' };

function tooltipOf(body: string): string {
  const match = /data-testid="character-in-game-tooltip"[\s\S]*?<\/div>/.exec(body);
  if (match === null) throw new Error('tooltip mock not found');
  return match[0];
}

describe('CharacterInGame', () => {
  it('mocks the game tooltip: the name in its class colour, then the addon’s own score line', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED } });
    const tooltip = tooltipOf(body);
    expect(tooltip).toContain('color: var(--color-class-priest)');
    expect(tooltip).toContain('Elyra Duskvale');
    // "Forever Score 1830 · Epic · as of 4 Nov", the fixture's `line` and `asOf` words.
    const text = tooltip.replace(/<[^>]+>/g, '');
    expect(text).toContain(`${labels.line.replace('%d', '1830')} · ${labels.tiers.epic.label} · as of 4 Nov`);
    expect(tooltip).toMatch(/data-testid="score-badge"[^>]*>1830</);
    // The number at body size takes the lightened epic text colour; "as of" is muted.
    expect(tooltip).toContain('color: #b866f5');
    expect(tooltip).toMatch(/class="text-muted">as of 4 Nov</);
  });

  it('says where the line shows, with a Get set up link', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED } });
    expect(body).toContain(scoreCopy.inGameShownOn);
    expect(body).toContain(`href="${scoreCopy.inGameSetupHref}"`);
    expect(body).toContain(scoreCopy.inGameSetupLink);
  });

  it('with no score, the mock carries the no-logs line in muted, never a zero', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: null } });
    const tooltip = tooltipOf(body);
    expect(tooltip).toMatch(new RegExp(`class="text-muted"[^>]*>${labels.noLogs}<`));
    expect(tooltip).not.toContain('score-badge');
    expect(body).not.toMatch(/>\s*0\s*</);
  });

  it('shows the planner line only to the owner', () => {
    const visitor = render(CharacterInGame, { props: { ...BASE, forever: SCORED } }).body;
    expect(visitor).not.toContain('character-in-game-planner');
    expect(visitor).not.toContain(scoreCopy.inGamePlannerLead);

    const owner = render(CharacterInGame, { props: { ...BASE, forever: SCORED, owned: true } }).body;
    expect(owner).toContain('data-testid="character-in-game-planner"');
    expect(scoreCopy.inGamePlannerLead).toBe('In game, type /fs, press Copy my character, then');
    expect(owner).toContain(`href="${scoreCopy.inGamePasteHref}"`);
    expect(owner).toContain(scoreCopy.inGamePasteLink);
  });

  it('shows when the site last saw the export only to the owner, and only with a time', () => {
    const local = new Date(2026, 10, 4, 14, 3).toISOString();
    const notOwner = render(CharacterInGame, { props: { ...BASE, forever: SCORED, lastSeenAt: local } }).body;
    expect(notOwner).not.toContain('character-in-game-last-seen');

    const noTime = render(CharacterInGame, { props: { ...BASE, forever: SCORED, owned: true } }).body;
    expect(noTime).not.toContain('character-in-game-last-seen');

    const owner = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, owned: true, lastSeenAt: local },
    }).body;
    expect(owner).toContain('data-testid="character-in-game-last-seen"');
    expect(owner).toContain('Last seen in game 4 Nov 14:03');
  });
});
