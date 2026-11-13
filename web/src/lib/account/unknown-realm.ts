// web/src/lib/account/unknown-realm.ts
// A Battle.net character on a realm Blizzard has not typed for us is imported under the ruleset
// `unknown` (api/internal/bnetimport/unknown_realm.go) and /setup asks its player once which realm
// type it is on. This module owns the question's words, which characters it asks about, and the
// answer's request. The four choices are RULESETS (lib/characters.ts), the site's one list.
import type { Me, MeCharacter } from './me-types';
import { RULESETS, UNKNOWN_RULESET, type Ruleset } from '../characters';

export { UNKNOWN_RULESET };

/** The answer's route: POST {key, ruleset}. */
export const REALM_ANSWER_PATH = '/v1/me/realm-answers';

export const unknownRealmCopy = {
  /** The question, one per character, around its class-coloured name: "Which realm type is
   *  <name> on?" */
  questionLead: 'Which realm type is',
  questionTail: 'on?',
  /** Why we ask, in game words. */
  why: 'Blizzard does not list this realm yet. Rankings and BiS follow your answer.',
  /** RulesetOf in the API reads an RP PvP realm as Roleplay. */
  rpPvp: 'RP PvP realms count as Roleplay.',
  failed: 'That did not save.',
} as const;

/** The choices, in the site's order. */
export const REALM_CHOICES: readonly { id: Ruleset; label: string }[] = RULESETS;

/** The account's characters still waiting for an answer, in the order /v1/me lists them. */
export function awaitingRealm(me: Pick<Me, 'characters'> | null): MeCharacter[] {
  return (me?.characters ?? []).filter((character) => character.ruleset === UNKNOWN_RULESET);
}
