// web/src/lib/account/character-descriptor.ts
// The muted descriptor line the design system's Character row shows under a name --
// "Night Elf Hunter · Level 25 · Living Flame (PvP US)" (brief 2026-09-22 §B3) -- and the
// class-initial fallback square shown when a character has no avatar_url. One place for
// both so CharacterList.svelte's rows and Account.svelte's hero band read the exact same
// text and letter rather than two hand-rolled copies drifting apart.
import { classColorVar } from '../report/format';
import { classDisplayName } from '../sim/spec-label';
import { rulesetLabel } from '../characters';
import type { MeCharacter } from './api';
import { characterListCopy } from './character-list-copy';

/**
 * Home rebuild spec §3.B.1's signed-in hero descriptor: "Level 24 Troll Marksmanship
 * Hunter" in the mock -- a DIFFERENT word order from `characterDescriptor`'s own "Race
 * Class · Level N · Realm (Ruleset REGION)" line, and mixed with markup
 * (`HomeAccountPanel.svelte` inserts a `FactionMark` icon and the guild tag around this
 * text, plus the faction word in the faction's own colour) that a single plain-text
 * function cannot produce anyway. Kept as its own small helper, not an extension of
 * `characterDescriptor` (which three other surfaces -- Account.svelte's hero band,
 * CharacterList.svelte and sim/LandingState.svelte -- render as-is today, unrelated to this
 * lane and out of scope to reshape): reusing its own building blocks
 * (`classDisplayName`/`rulesetLabel` are for the OTHER function's own realm/ruleset clause,
 * not needed here) keeps this honest and small rather than overloading one function with
 * two unrelated output shapes behind a flag.
 *
 * The mock's line also names a spec ("Marksmanship Hunter") that `MeCharacter` carries no
 * field for -- spec is only inferable today from a sim input fetch this lane does not wire
 * up for every character chip, so this omits it rather than guess (tenet 8): "Level 24
 * Troll Hunter", not a fabricated "Marksmanship".
 */
export function homeHeroLevelRaceClassLine(character: MeCharacter): string {
  const className = character.class === undefined ? undefined : classDisplayName(character.class);
  const parts = [
    character.level === undefined ? undefined : characterListCopy.levelPrefix(character.level),
    character.race,
    className,
  ].filter((part): part is string => part !== undefined);
  return parts.join(' ');
}

/** "Night Elf Hunter · Level 25 · Living Flame (PvP US)"; every part but the ruleset/
 *  region location is omitted when the character has never been through the Battle.net
 *  import. */
export function characterDescriptor(character: MeCharacter): string {
  const className = character.class === undefined ? undefined : classDisplayName(character.class);
  const raceClass = [character.race, className]
    .filter((part): part is string => part !== undefined)
    .join(' ');
  const location =
    character.realm === undefined
      ? locationLabel(character)
      : `${character.realm} (${locationLabel(character)})`;
  const segments = [
    raceClass,
    character.level === undefined ? '' : characterListCopy.levelPrefix(character.level),
    location,
  ];
  return segments.filter((segment) => segment !== '').join(' · ');
}

function locationLabel(character: MeCharacter): string {
  return `${rulesetLabel(character.ruleset)} ${character.region.toUpperCase()}`;
}

/** The smallest shape `classSquare`/`classIconUrl` need -- every `MeCharacter` satisfies it, and
 *  so does `CharacterPortrait.svelte`'s own narrower `PortraitCharacter` prop type (spec
 *  2026-09-23 §2.1: "type it as the smallest shape and let MeCharacter satisfy it"). */
export interface ClassIconSubject {
  name: string;
  class?: string;
}

/** The class-coloured fallback square's letter and colour, for a character with no
 *  avatar_url -- the class's first letter (design/DESIGN-SYSTEM.md's Character row) in
 *  the class's own colour, or the site's own text colour for a character with no class on
 *  file yet. */
export function classSquare(character: ClassIconSubject): { letter: string; color: string } {
  const label = character.class === undefined ? character.name : classDisplayName(character.class);
  return { letter: label.charAt(0).toUpperCase(), color: classColorVar(character.class) };
}

/** Blizzard's own 56 px class icons, on the CDN that also serves character avatars. The
 *  path is per game version; Classic Era's set is the one every vanilla class has today,
 *  and Forever's, once it exists, is one constant away. */
const CLASS_ICON_BASE = 'https://render.worldofwarcraft.com/classic1x-us/icons/56/classicon_';
const CLASS_ICON_SLUGS: ReadonlySet<string> = new Set([
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
]);

/** The class icon a character without an avatar shows; undefined for an unknown class. */
export function classIconUrl(character: ClassIconSubject): string | undefined {
  const slug = character.class?.toLowerCase();
  return slug !== undefined && CLASS_ICON_SLUGS.has(slug) ? `${CLASS_ICON_BASE}${slug}.jpg` : undefined;
}
