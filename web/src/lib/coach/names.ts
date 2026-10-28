// web/src/lib/coach/names.ts
// The build's name tables the coach panel reads: spell and item names per class (the sim's name
// files) and the spell icon stems. All three are static, cacheable build files; a missing file
// leaves the panel with plain text and no icons rather than failing it.
import { dataUrl, fetchJson } from '../planner/load';
import { buildDictionary, type NameDictionary } from './name-tokens';

interface ClassNames {
  spell: Record<string, string>;
  item: Record<string, string>;
}

export interface CoachNames {
  dictionary: NameDictionary;
  spellName: (id: number) => string;
  iconUrl: IconUrl;
}

export const EMPTY_NAMES: CoachNames = {
  dictionary: new Map(),
  spellName: () => '',
  iconUrl: () => null,
};

async function optional<T>(url: string): Promise<T | null> {
  try {
    return await fetchJson<T>(url);
  } catch {
    return null;
  }
}

export type IconUrl = (spellId: number | null) => string | null;

/** Spell icon urls alone: the row strip needs no names, only the icon table. */
export async function loadSpellIconUrls(build: string): Promise<IconUrl> {
  const icons = await optional<Record<string, string>>(dataUrl(build, 'spellicons.json'));
  return (spellId) => {
    const stem = spellId === null ? undefined : icons?.[String(spellId)];
    return stem === undefined ? null : dataUrl(build, `icons/${stem}.webp`);
  };
}

export async function loadCoachNames(build: string, classSlug: string): Promise<CoachNames> {
  const [names, iconUrl] = await Promise.all([
    optional<ClassNames>(dataUrl(build, `simnames/${classSlug}.json`)),
    loadSpellIconUrls(build),
  ]);
  const spells = names?.spell ?? {};
  return {
    dictionary: buildDictionary(spells, names?.item ?? {}),
    spellName: (id) => spells[String(id)] ?? '',
    iconUrl,
  };
}
