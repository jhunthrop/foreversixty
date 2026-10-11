// web/src/lib/rankings/defaults.ts
// Where the board starts for whoever is looking (design/specs/2026-10-11-rankings.md section 3,
// decision 4): the viewer's own class, spec, ruleset and role's metric, the current phase, and
// the encounter they most recently ranked on. A link that names a filter always wins over the
// default for that filter. Replaces the old pointer prefill, which knew only class and ruleset.
import type { MeCharacter } from '../account/api';
import { parseCharacterKey, type CharacterPath } from '../characters';
import { classSlugFromName } from '../report/tree-sizes';
import type { SpecRole } from '../sim/specs';
import type { CharacterFight, CharacterPage, EncounterOption } from './api';
import { phaseAt } from './phases';
import { specFor } from './scope';
import { AIMED_KEYS, type RankingsDefaults, type RankingsState } from './url';

/** The metric a role is judged on: damage dealt, healing done, damage taken by a tank. */
const ROLE_METRIC: Readonly<Record<SpecRole, string>> = {
  dps: 'dps',
  healer: 'hps',
  tank: 'damage_taken',
};

/** Every default, for a signed-out visitor (`character` null: nothing aimed, damage) or a character. */
export function rankingsDefaults(character: MeCharacter | null, now: Date): RankingsDefaults {
  const phase = phaseAt(now.toISOString());
  if (character === null) return { class: '', spec: '', ruleset: '', phase, metric: 'dps' };
  const classSlug = character.class === undefined ? '' : classSlugFromName(character.class);
  const known = classSlug === '' ? undefined : specFor(classSlug, character.spec ?? '');
  return {
    class: classSlug,
    spec: classSlug === '' ? '' : (known?.name ?? character.spec ?? ''),
    ruleset: character.ruleset,
    phase,
    metric: ROLE_METRIC[known?.role ?? 'dps'],
  };
}

/** The state with each default filled in where the URL (`search`) did not name that filter. */
export function applyDefaults(
  state: RankingsState,
  search: string,
  defaults: RankingsDefaults,
): RankingsState {
  const named = new URLSearchParams(search);
  // A link that names a class but no spec means that whole class, not the viewer's own spec of it.
  const wholeClass = named.has('class') && !named.has('spec');
  const filled = [...AIMED_KEYS, 'metric' as const].filter(
    (key) => !named.has(key) && !(wholeClass && key === 'spec'),
  );
  return filled.reduce<RankingsState>((next, key) => ({ ...next, [key]: defaults[key] }), state);
}

/** A header switch re-aims the board at the new character, back on page one; other filters stay. */
export function reaim(state: RankingsState, defaults: RankingsDefaults): RankingsState {
  return { ...state, ...defaults, page: 1 };
}

/** The reads the default encounter is chosen from, injectable so a test needs no network. */
export interface EncounterReads {
  encounters(): Promise<{ rows: EncounterOption[] }>;
  character(path: CharacterPath): Promise<Pick<CharacterPage, 'history'>>;
}

/** The character's page path from its key (`us/normal/obnoxious-yell`); null for a key that is not one. */
export function characterPathOf(character: MeCharacter): CharacterPath | null {
  const parsed = parseCharacterKey(character.key);
  const slug = character.key.split('/')[2];
  return parsed === null || slug === undefined || slug === '' ? null : { ...parsed, slug };
}

/**
 * The encounters the API lists and the one a board opens on for this viewer. A character whose
 * page does not load, or that has no kill, opens on the first listed encounter: the default is
 * a convenience and never blocks the board.
 */
export async function resolveDefaultEncounter(
  character: MeCharacter | null,
  reads: EncounterReads,
): Promise<{ slug: string; encounters: EncounterOption[] }> {
  const path = character === null ? null : characterPathOf(character);
  const [listed, page] = await Promise.all([
    reads.encounters().then((result) => result.rows),
    path === null ? Promise.resolve(null) : reads.character(path).catch(() => null),
  ]);
  return { slug: pickDefaultEncounter(page?.history ?? [], listed), encounters: listed };
}

/**
 * The encounter a board opens on: the one the character most recently killed and so is ranked
 * on, else the first encounter the API lists (until it orders by ranked kills), else none. Only
 * an encounter the API lists is named, so the slug is always one it accepts.
 */
export function pickDefaultEncounter(
  history: readonly Pick<CharacterFight, 'encounter_id' | 'kill'>[],
  encounters: readonly EncounterOption[],
): string {
  for (const fight of history) {
    if (fight.kill === false) continue;
    const listed = encounters.find((encounter) => encounter.id === fight.encounter_id);
    if (listed !== undefined) return listed.slug;
  }
  return encounters[0]?.slug ?? '';
}
