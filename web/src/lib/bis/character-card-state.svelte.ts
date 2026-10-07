// web/src/lib/bis/character-card-state.svelte.ts
// The Character card's own session read (bis rebuild spec §4.B/§5) -- the same
// `createQueryState` + `heroCharacter`/`mainCharacter` selection `lib/account/home-hero.
// svelte.ts` uses (so this card agrees with every other surface about which character is
// "yours"), but exposing `status`/`error`/`refresh` too, which that shared handle does not
// -- the card's own `Skeleton`/`LoadError` states need them. A small module of its own
// rather than widening `home-hero.svelte.ts` itself, which the home rebuild lane owns.
import { fetchMeOnce, type Me, type MeCharacter } from '../account/api';
import { createQueryState } from '../data/query.svelte';
import { readCurrent } from '../current-character';
import { selectedCharacter } from '../account/hero-character';
import { API_BASE_URL } from '../planner/config';

export interface CharacterCardState {
  readonly status: 'idle' | 'loading' | 'ready' | 'failed';
  readonly error: string;
  readonly me: Me | null;
  readonly character: MeCharacter | null;
  refresh(): void;
}

export function createCharacterCardState(): CharacterCardState {
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
  const character = $derived<MeCharacter | null>(me === null ? null : selectedCharacter(readCurrent(), me));

  return {
    get status() {
      return session.status;
    },
    get error() {
      return session.error;
    },
    get me() {
      return me;
    },
    get character() {
      return character;
    },
    refresh(): void {
      session.refresh();
    },
  };
}
