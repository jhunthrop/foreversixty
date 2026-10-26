import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacterGuild } from '../../lib/account/api';
import CharacterGuildLine from './CharacterGuildLine.svelte';

const GUILD: MeCharacterGuild = { id: 1, name: 'The Last Watch', rank: 'member', verified: true };

function text(guild: MeCharacterGuild): string {
  const { body } = render(CharacterGuildLine, { props: { guild } });
  return body
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

describe('CharacterGuildLine attendance', () => {
  it('appends the raid count in raid mode', () => {
    expect(text({ ...GUILD, attendance: { present: 7, nights: 8 } })).toContain('· came 7 of 8 raids');
  });

  it('appends nothing without raid attendance', () => {
    expect(text(GUILD)).not.toContain('came');
  });
});
