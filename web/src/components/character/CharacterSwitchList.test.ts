// web/src/components/character/CharacterSwitchList.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import CharacterSwitchList from './CharacterSwitchList.svelte';

const CHARACTERS: MeCharacter[] = [
  { key: 'us/normal/simfury', region: 'us', ruleset: 'normal', name: 'Simfury', class: 'warrior' },
  { key: 'us/normal/roland', region: 'us', ruleset: 'normal', name: 'Roland', class: 'mage' },
];

describe('CharacterSwitchList', () => {
  it('lists every character with a switch action, each carrying its own key in the testid', () => {
    const { body } = render(CharacterSwitchList, {
      props: { characters: CHARACTERS, currentKey: null, onswitch: () => {} },
    });
    for (const character of CHARACTERS) {
      expect(body).toContain(`data-testid="current-character-bar-switch-${character.key}"`);
      expect(body).toContain(character.name);
    }
  });

  it('marks the current character so it is visibly distinct from the others', () => {
    const { body } = render(CharacterSwitchList, {
      props: { characters: CHARACTERS, currentKey: 'us/normal/simfury', onswitch: () => {} },
    });
    expect(body).toContain('data-testid="current-character-bar-switch-current"');
  });

  it('renders no marker at all when nothing is current', () => {
    const { body } = render(CharacterSwitchList, {
      props: { characters: CHARACTERS, currentKey: null, onswitch: () => {} },
    });
    expect(body).not.toContain('data-testid="current-character-bar-switch-current"');
  });
});
