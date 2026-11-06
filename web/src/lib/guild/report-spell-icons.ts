// web/src/lib/guild/report-spell-icons.ts
// The death causes' spell icons, through the site's one spell icon path: the build's spell id to
// icon stem table (`spellicons.json`, optional) and `/data/<build>/icons/<stem>.webp`. An id the
// table does not carry (most boss abilities, a melee swing) has no icon, never a broken image.
import { isRealIcon } from '../icons/placeholder';
import { dataUrl, loadOptional } from '../planner/load';

export type SpellIconOf = (spellId: number | null) => string | null;

export const NO_SPELL_ICONS: SpellIconOf = () => null;

/** The icon url of a spell id in `icons`, or null. */
export function spellIconOf(build: string, icons: Readonly<Record<string, string>>): SpellIconOf {
  return (spellId) => {
    if (spellId === null) return null;
    const stem = icons[String(spellId)];
    return stem !== undefined && isRealIcon(stem) ? dataUrl(build, `icons/${stem}.webp`) : null;
  };
}

/** The one load; a build without the table draws every cause without an icon. */
export async function loadSpellIcons(build: string): Promise<SpellIconOf> {
  return spellIconOf(
    build,
    await loadOptional<Record<string, string>>(dataUrl(build, 'spellicons.json'), {}),
  );
}
