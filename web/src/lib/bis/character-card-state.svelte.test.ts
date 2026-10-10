// @vitest-environment jsdom
// web/src/lib/bis/character-card-state.svelte.test.ts
// The BiS character card follows the header selector without a reload.
import { describe, expect, it, vi } from 'vitest';
import { pointerForCharacter } from '../account/main-character';
import { withEffectRoot } from '../data/query.svelte';
import { CURRENT_CHARACTER_CHANGED, writeCurrent } from '../current-character';

function flush(): Promise<void> {
  return Promise.resolve().then(() => Promise.resolve());
}

const ARMS = {
  key: 'us/normal/armsy',
  region: 'us',
  ruleset: 'normal',
  name: 'Armsy',
  class: 'Warrior',
  spec: 'Arms',
};
const BOW = {
  key: 'us/normal/bow-jackzon',
  region: 'us',
  ruleset: 'normal',
  name: 'Bow Jackzon',
  class: 'Hunter',
  spec: 'Beast Mastery',
};

vi.mock('../account/api', async (importOriginal) => {
  const actual = await importOriginal<typeof import('../account/api')>();
  return {
    ...actual,
    fetchMeOnce: vi.fn().mockResolvedValue({
      user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
      characters: [ARMS, BOW],
      guilds: [],
    }),
  };
});

const { createCharacterCardState } = await import('./character-card-state.svelte');

describe('createCharacterCardState', () => {
  it('re-selects the character when the header selector writes a new pointer', async () => {
    writeCurrent(pointerForCharacter(ARMS));
    let handle: ReturnType<typeof createCharacterCardState>;
    const cleanup = withEffectRoot(() => {
      handle = createCharacterCardState();
    });
    await flush();
    await flush();
    expect(handle!.character?.name).toBe('Armsy');

    writeCurrent(pointerForCharacter(BOW));
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await flush();
    expect(handle!.character?.name).toBe('Bow Jackzon');
    cleanup();
  });
});
