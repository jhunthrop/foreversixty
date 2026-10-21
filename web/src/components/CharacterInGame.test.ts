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
    // L8.5 ruling 5: the addon's copy mode has no button; `/fs copy`, then Ctrl+C.
    expect(scoreCopy.inGamePlannerLead).toBe(
      'This page updates by itself each time you log out with the addon and the companion running. Without the companion: /fs copy in game, then',
    );
    expect(scoreCopy.inGamePasteHref).toBe('/setup');
    expect(owner).toContain(`href="${scoreCopy.inGamePasteHref}"`);
    expect(owner).toContain(scoreCopy.inGamePasteLink);
  });

  it('never says "Last seen in game": the export time is not an online time (L8.4 fix round 2 item 5)', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED, owned: true } });
    expect(body).not.toContain('Last seen in game');
    expect(body).not.toContain('character-in-game-last-seen');
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
    // The name 14px bold.
    expect(body).toContain(
      'class="text-[14px] font-bold" style="color: var(--color-class-priest)">Elyra Duskvale<',
    );
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

  // Fix round 2 item 6, from the real 1.60 client (ingame-reference-client-player-tooltip.jpg):
  // name, <Guild>, "Level 60 Night Elf", the class on its own line, the faction on its
  // own line when known, all in the client's light grey-white; then the blank line and the
  // addon's score line.
  it('keeps the client’s order: name, guild, "Level 60 Night Elf", class, faction, a blank line, the score', () => {
    const { body } = render(CharacterInGame, {
      props: {
        ...BASE,
        forever: SCORED,
        guild: 'Cx66 Olympus',
        level: 60,
        race: 'Night Elf',
        faction: 'alliance',
      },
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
      at('data-testid="character-in-game-descriptor">Level 60 Night Elf<'),
      at('data-testid="character-in-game-class">Priest<'),
      at('data-testid="character-in-game-faction">Alliance<'),
      at('<span class="h-[20px]" aria-hidden="true"></span>'),
      at('data-testid="character-in-game-score"'),
    ];
    expect([...order].sort((a, b) => a - b)).toEqual(order);
  });

  it('never says "(Player)", and title-cases a class the API sends lower-case (W.9)', () => {
    const { body } = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, className: 'hunter', level: 24, race: 'Orc', faction: 'horde' },
    });
    expect(body).not.toContain('(Player)');
    expect(body).toMatch(/data-testid="character-in-game-descriptor">Level 24 Orc</);
    expect(body).toMatch(/data-testid="character-in-game-class">Hunter</);
  });

  it('draws the guild, level, class and faction lines in the client’s light grey-white', () => {
    const { body } = render(CharacterInGame, {
      props: {
        ...BASE,
        forever: SCORED,
        guild: 'Cx66 Olympus',
        level: 60,
        race: 'Night Elf',
        faction: 'alliance',
      },
    });
    for (const id of ['guild', 'descriptor', 'class', 'faction']) {
      expect(body, id).toContain(`class="text-[#dfe0e5]" data-testid="character-in-game-${id}"`);
    }
  });

  it('leaves the faction line out when the API does not know it', () => {
    const { body } = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: 60, race: 'Night Elf', faction: null },
    });
    expect(body).toMatch(/data-testid="character-in-game-class">Priest</);
    expect(body).not.toContain('character-in-game-faction');
  });

  it('writes the level line from whichever of level and race is known, and leaves the block out with neither', () => {
    const levelOnly = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: 60, race: null },
    }).body;
    expect(levelOnly).toMatch(/data-testid="character-in-game-descriptor">Level 60</);
    const raceOnly = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: null, race: 'Night Elf' },
    }).body;
    expect(raceOnly).toMatch(/data-testid="character-in-game-descriptor">Night Elf</);
    const neither = render(CharacterInGame, {
      props: { ...BASE, forever: SCORED, level: null, race: null, faction: 'alliance' },
    }).body;
    expect(neither).not.toContain('character-in-game-descriptor');
    // Never "Priest" alone under the name.
    expect(neither).not.toContain('character-in-game-class');
    expect(neither).not.toContain('character-in-game-faction');
  });

  it('drops the blank line when there is neither a guild nor a level line', () => {
    const bare = tooltipOf(render(CharacterInGame, { props: { ...BASE, forever: SCORED } }).body);
    expect(bare).not.toContain('h-[20px]');
    const guildOnly = tooltipOf(
      render(CharacterInGame, { props: { ...BASE, forever: SCORED, guild: 'Cx66 Olympus' } }).body,
    );
    expect(guildOnly).toContain('<span class="h-[20px]" aria-hidden="true"></span>');
    const levelOnly = tooltipOf(
      render(CharacterInGame, { props: { ...BASE, forever: SCORED, level: 60 } }).body,
    );
    expect(levelOnly).toContain('<span class="h-[20px]" aria-hidden="true"></span>');
  });

  it('caps the box at the client’s width and wraps a long line, never clips it', () => {
    const long = 'The Exceedingly Long Guild Name Of Olympus';
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED, guild: long } });
    const frame = (
      /<div class="([^"]*)"[^>]*data-testid="character-in-game-tooltip"/.exec(body)?.[1] ?? ''
    ).split(' ');
    expect(frame).toEqual(expect.arrayContaining(['max-w-[min(100%,320px)]', 'break-words']));
    for (const clip of ['overflow-hidden', 'truncate', 'whitespace-nowrap', 'max-w-full']) {
      expect(frame, clip).not.toContain(clip);
    }
    expect(tooltipOf(body)).toContain(`&lt;${long}`);
  });

  it('puts the line under the mock below 900px and beside it from 900px', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED } });
    expect(body).toContain('flex flex-col gap-3 min-[900px]:flex-row');
  });

  // L8.4 fix round 5 item 2: the inline links take their gold from the base `a` rule, so the
  // design system's hover and pressed --gold-hover apply; a text-gold utility would pin them.
  it('leaves the inline links’ colour to the base rule: no text-gold utility (fix round 5 item 2)', () => {
    const { body } = render(CharacterInGame, { props: { ...BASE, forever: SCORED, owned: true } });
    const links = [...body.matchAll(/<a class="([^"]*)"/g)].map((match) => match[1].split(' '));
    expect(links).toHaveLength(2);
    for (const link of links) {
      expect(link).toContain('underline');
      expect(link).not.toContain('text-gold');
    }
  });
});
