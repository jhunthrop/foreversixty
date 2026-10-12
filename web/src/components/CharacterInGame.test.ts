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
    expect(scoreCopy.inGamePlannerLead).toBe(
      'To update this page: in game, type /fs, press Copy my character, then',
    );
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

  it('is the game’s tooltip frame: translucent dark, 1px grey-blue border, 4px radius, 10px padding (round 2 item 3)', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED } });
    const frame = /<div class="([^"]*)"[^>]*data-testid="character-in-game-tooltip"/.exec(body)?.[1] ?? '';
    expect(frame.split(' ')).toEqual(
      expect.arrayContaining([
        'border',
        'border-[#5d6b82]',
        'bg-[rgba(8,10,16,0.92)]',
        'rounded-[4px]',
        'p-[10px]',
      ]),
    );
    expect(frame).not.toContain('gold');
    // The name 14px bold, then a blank line before the addon's line.
    expect(body).toContain(
      'class="text-[14px] font-bold" style="color: var(--color-class-priest)">Elyra Duskvale<',
    );
    expect(tooltipOf(body)).toContain('<span class="h-[20px]" aria-hidden="true"></span>');
  });

  it('adds the game’s "Level 60 Human Priest" line only when level and race are known', () => {
    const without = render(CharacterInGame, { props: { ...BASE, forever: SCORED } }).body;
    expect(without).not.toContain('character-in-game-descriptor');
    const withLine = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: 60, race: 'Human' },
    }).body;
    expect(withLine).toMatch(/data-testid="character-in-game-descriptor">Level 60 Human Priest</);
  });

  it('adds the guild line under the name only when the character has a guild (Task L8.4 ruling 2)', () => {
    const without = render(CharacterInGame, { props: { ...BASE, forever: SCORED } }).body;
    expect(without).not.toContain('character-in-game-guild');
    const nullGuild = render(CharacterInGame, { props: { ...BASE, forever: SCORED, guild: null } }).body;
    expect(nullGuild).not.toContain('character-in-game-guild');
    const withGuild = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, guild: 'Cx66 Olympus' },
    }).body;
    // The client's own form: the name in angle brackets, escaped in the markup.
    expect(scoreCopy.tooltipGuild('Cx66 Olympus')).toBe('<Cx66 Olympus>');
    expect(tooltipOf(withGuild)).toMatch(/data-testid="character-in-game-guild">&lt;Cx66 Olympus(&gt;|>)</);
  });

  it('keeps the game’s order: name, guild, level race class, a blank line, the score', () => {
    const { body } = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, guild: 'Cx66 Olympus', level: 60, race: 'Night Elf' },
    });
    const tooltip = tooltipOf(body);
    const at = (needle: string): number => {
      const index = tooltip.indexOf(needle);
      expect(index, needle).toBeGreaterThan(-1);
      return index;
    };
    const order = [
      at('>Elyra Duskvale<'),
      at('data-testid="character-in-game-guild"'),
      at('data-testid="character-in-game-descriptor">Level 60 Night Elf Priest<'),
      at('<span class="h-[20px]" aria-hidden="true"></span>'),
      at('data-testid="character-in-game-score"'),
    ];
    expect([...order].sort((a, b) => a - b)).toEqual(order);
  });

  it('writes the level line from whichever of level and race is known, and leaves it out with neither', () => {
    const levelOnly = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: 60, race: null },
    }).body;
    expect(levelOnly).toMatch(/data-testid="character-in-game-descriptor">Level 60 Priest</);
    const raceOnly = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: null, race: 'Night Elf' },
    }).body;
    expect(raceOnly).toMatch(/data-testid="character-in-game-descriptor">Night Elf Priest</);
    const neither = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: null, race: null },
    }).body;
    expect(neither).not.toContain('character-in-game-descriptor');
  });

  it('puts the line under the mock below 900px and beside it from 900px', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED } });
    expect(body).toContain('flex flex-col gap-3 min-[900px]:flex-row');
  });

  it('shows "Last seen in game" without a time for an owner whose export time cannot be read (fix round 4 item 3)', () => {
    const { body } = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, owned: true, lastSeenAt: 'not a time' },
    });
    expect(body).toContain(
      'class="text-muted text-[14px] leading-[22px]" data-testid="character-in-game-last-seen">Last seen in game<',
    );
  });
});
