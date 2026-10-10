// web/src/lib/account/home-hero.svelte.ts
// The one place the home page's islands read the session and settle on a hero character
// (spec 2026-09-24 §2.1/§2.3: HomeAccountPanel's hero and HomeNextSteps' four cards must
// agree on who "the main character" is, and must not each fetch /v1/me or the rating
// separately). Built on createQueryState the same way SessionNav.svelte, Account.svelte and
// the pre-refactor HomeAccountPanel.svelte already do -- the caching layer's query() cache
// (lib/data/query.ts) dedupes the underlying /v1/me and rating requests across both callers
// by URL, so two `createHomeHero()` instances share one network read each, not two.
import { fetchMeOnce, type Me, type MeCharacter } from './api';
import { createQueryState } from '../data/query.svelte';
import { readCurrent, onCurrentCharacterChange } from '../current-character';
import { selectedCharacter } from './hero-character';
import { API_BASE_URL } from '../planner/config';
import { fetchCharacterRating } from '../rankings/api';
import type { CharacterRating } from '../rating/types';
import { parseCharacterPath, type CharacterPath } from '../characters';

export interface HomeHeroHandle {
  readonly me: Me | null;
  readonly ready: boolean;
  readonly hero: MeCharacter | null;
  readonly heroPath: CharacterPath | null;
  readonly rating: CharacterRating | null;
}

export function createHomeHero(): HomeHeroHandle {
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
  // 'ready' mirrors createQueryState's two terminal statuses, matching every other caller's
  // own convention (see HomeAccountPanel's pre-refactor comment, carried over unchanged).
  const ready = $derived(session.status === 'ready' || session.status === 'failed');

  // Bumped on CURRENT_CHARACTER_CHANGED (the header selector, a sibling island), so `hero`
  // re-reads the stored pointer.
  let pointerVersion = $state(0);
  $effect(() => {
    const onChanged = (): void => {
      pointerVersion += 1;
    };
    return onCurrentCharacterChange(onChanged);
  });

  function readCurrentIfReady(): ReturnType<typeof readCurrent> {
    void pointerVersion;
    if (!ready || me === null) return null;
    return readCurrent();
  }

  const hero = $derived<MeCharacter | null>(me === null ? null : selectedCharacter(readCurrentIfReady(), me));

  const heroPath = $derived(hero === null ? null : parseCharacterPath(`/character/${hero.key}`));

  // The rating figure, chained off the hero rather than blocking it (spec 2026-09-23's own
  // rule, carried over unchanged): never shown until it resolves with a real sample.
  let rating = $state<CharacterRating | null>(null);
  $effect(() => {
    const path = heroPath;
    rating = null;
    if (path === null) return;
    void fetchCharacterRating(path)
      .then((result) => {
        if (path !== heroPath) return;
        rating = result;
      })
      .catch(() => {
        rating = null;
      });
  });

  return {
    get me() {
      return me;
    },
    get ready() {
      return ready;
    },
    get hero() {
      return hero;
    },
    get heroPath() {
      return heroPath;
    },
    get rating() {
      return rating;
    },
  };
}
