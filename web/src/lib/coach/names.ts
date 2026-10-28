// web/src/lib/coach/names.ts
// The build's name tables the coach panel reads: spell and item names per class (the sim's name
// files) and the spell icon stems. All three are static, cacheable build files; a missing file
// leaves the panel with plain text and no icons rather than failing it.
import { dataUrl, fetchJson } from '../planner/load';
import { buildDictionary, type NameDictionary } from './name-tokens';
import type { CoachAnchor } from './types';

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

let warned = false;

/**
 * Says once, in development only, that a name table is missing: the panel then shows plain text
 * and no icons, which is easy to mistake for a rendering bug.
 */
function warnFallback(url: string): void {
  if (!import.meta.env.DEV || warned) return;
  warned = true;
  console.warn(`coach: ${url} did not load; spell and item names fall back to plain text`);
}

async function optional<T>(url: string): Promise<T | null> {
  try {
    return await fetchJson<T>(url);
  } catch {
    warnFallback(url);
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
    dictionary: buildDictionary(spells, names?.item ?? {}, (id) => iconUrl(id) !== null),
    spellName: (id) => spells[String(id)] ?? '',
    iconUrl,
  };
}

/**
 * The name tables with the answer's own anchor spell names added: a spell the build's class name
 * table lacks (a talent such as Bloodthirst) still becomes a token in the prose and names its chip.
 * The tables win where both name a spell.
 */
export function withAnchorNames(
  names: CoachNames,
  anchors: Readonly<Record<string, CoachAnchor>>,
): CoachNames {
  const named = new Map<number, string>();
  for (const anchor of Object.values(anchors)) {
    if (anchor.spell_id !== undefined && anchor.spell_name !== undefined)
      named.set(anchor.spell_id, anchor.spell_name);
  }
  if (named.size === 0) return names;
  const dictionary = new Map(names.dictionary);
  for (const [id, name] of named) if (!dictionary.has(name)) dictionary.set(name, { kind: 'spell', id });
  return {
    ...names,
    dictionary,
    spellName: (id) => names.spellName(id) || (named.get(id) ?? ''),
  };
}
