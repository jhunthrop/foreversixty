// web/src/lib/guild/report-spell-icons.ts
// A log's spell icons, through the site's one spell icon path: `/data/<build>/icons/<stem>.webp`.
// The build's id table (`spellicons.json`) carries a player spell's own id only, so a spell goes
// down a chain until it has an icon (never an empty socket once the tables load; logs lane review
// 2026-11-09): a melee swing is the client's attack icon; then the client's own override (Auto
// Shot); then the id table; then a same-named spell's icon (`spellicons-by-name.json`, written by
// web/scripts/spell-icons-publish.mjs); then its school's base spell; then a plain rune (never the attack icon on a spell).
import fallbacks from '../../data/spell-icon-fallbacks.json';
import { isRealIcon } from '../icons/placeholder';
import { dataUrl, loadOptional } from '../planner/load';
import { iconUrl } from '../ui/assets';

/** A spell's icon url: by id, then by name and school when the caller has them. */
export type SpellIconOf = (spellId: number | null, name?: string, school?: number) => string | null;

/** Before the tables load: the socket, at its final size. */
export const NO_SPELL_ICONS: SpellIconOf = () => null;

const MELEE_NAMES: ReadonlySet<string> = new Set(['', 'Melee', 'Auto Attack']);
const SCHOOLS: Readonly<Record<string, string>> = fallbacks.schools;
const CLIENT: Readonly<Record<string, string>> = fallbacks.client;

/** A school mask's base icon: its lowest school bit (a fire-frost spell reads as fire). */
function schoolStem(school: number | undefined): string | undefined {
  if (school === undefined || school <= 0) return undefined;
  return SCHOOLS[String(school & -school)];
}

/** The icon stem a spell ends on; never undefined. Exported for the unit test. */
export function spellIconStem(
  table: Readonly<Record<string, string>>,
  byName: Readonly<Record<string, string>>,
  spellId: number,
  name = '',
  school?: number,
): string {
  if (spellId === 0 && MELEE_NAMES.has(name)) return fallbacks.melee;
  const real = (stem: string | undefined): string | undefined =>
    stem !== undefined && isRealIcon(stem) ? stem : undefined;
  return (
    real(CLIENT[String(spellId)]) ??
    real(table[String(spellId)]) ??
    real(name === '' ? undefined : byName[name]) ??
    schoolStem(school) ??
    (spellId === 0 ? fallbacks.melee : fallbacks.spell)
  );
}

/** The icon url of a spell in `build`. With `chain` (the report) it goes down the whole chain and
 *  is never null once loaded; without it (the guild report's death causes, as drawn) the id table
 *  alone, null for an id it lacks. */
export function spellIconOf(
  build: string,
  table: Readonly<Record<string, string>>,
  byName: Readonly<Record<string, string>> = {},
  chain = false,
): SpellIconOf {
  return (spellId, name, school) => {
    if (spellId === null) return null;
    if (!chain) {
      return iconUrl(build, table[String(spellId)]);
    }
    return iconUrl(build, spellIconStem(table, byName, spellId, name, school));
  };
}

/** The one load: both tables, each optional (a build without them still draws the fallbacks). */
export async function loadSpellIcons(
  build: string,
  { chain = false }: { chain?: boolean } = {},
): Promise<SpellIconOf> {
  const [table, byName] = await Promise.all([
    loadOptional<Record<string, string>>(dataUrl(build, 'spellicons.json'), {}),
    loadOptional<Record<string, string>>(dataUrl(build, 'spellicons-by-name.json'), {}),
  ]);
  return spellIconOf(build, table, byName, chain);
}
