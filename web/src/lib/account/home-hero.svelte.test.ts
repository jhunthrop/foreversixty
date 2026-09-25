// @vitest-environment jsdom
// web/src/lib/account/home-hero.svelte.test.ts
import { describe, expect, it, vi } from 'vitest';
import { withEffectRoot } from '../data/query.svelte';
import { CURRENT_CHARACTER_CHANGED, writeCurrent } from '../current-character';

function flush(): Promise<void> {
  return Promise.resolve().then(() => Promise.resolve());
}

const CHAR = {
  key: 'us/normal/kiloz',
  region: 'us',
  ruleset: 'normal',
  name: 'Kiloz',
  class: 'Warrior',
  level: 60,
};

vi.mock('../account/api', async (importOriginal) => {
  const actual = await importOriginal<typeof import('../account/api')>();
  return {
    ...actual,
    fetchMeOnce: vi.fn().mockResolvedValue({
      user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
      characters: [CHAR],
      guilds: [],
    }),
  };
});

vi.mock('../rankings/api', async (importOriginal) => {
  const actual = await importOriginal<typeof import('../rankings/api')>();
  return {
    ...actual,
    fetchCharacterRating: vi.fn().mockResolvedValue({
      player_key: CHAR.key,
      sample_size: 0,
      trend: [],
      best_component: '',
      worst_component: '',
      latest: null,
    }),
  };
});

// A dynamic import, done after the fixture const above and the two vi.mock calls: a static
// `import { createHomeHero } from './home-hero.svelte'` is evaluated before any of this
// file's own top-level `const`s run (ESM module-graph order), so the `../account/api` mock
// factory above would close over `CHAR` while it is still in its temporal dead zone. Same
// fix `lib/addon-export.test.ts` already uses for the same reason.
const { createHomeHero } = await import('./home-hero.svelte');

describe('createHomeHero', () => {
  it('resolves the hero from main-character fallback once the session loads', async () => {
    let handle: ReturnType<typeof createHomeHero>;
    const cleanup = withEffectRoot(() => {
      handle = createHomeHero();
    });
    expect(handle!.ready).toBe(false);
    await flush();
    await flush();
    expect(handle!.ready).toBe(true);
    expect(handle!.hero?.key).toBe(CHAR.key);
    cleanup();
  });

  it('re-derives the hero when a sibling island writes a new current-character pointer', async () => {
    let handle: ReturnType<typeof createHomeHero>;
    const cleanup = withEffectRoot(() => {
      handle = createHomeHero();
    });
    await flush();
    await flush();
    writeCurrent({
      source: 'armory',
      ref: CHAR.key,
      label: 'Kiloz',
      classSlug: 'warrior',
      savedAt: new Date().toISOString(),
    });
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await flush();
    expect(handle!.hero?.key).toBe(CHAR.key);
    cleanup();
  });
});
