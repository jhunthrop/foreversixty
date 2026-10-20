// web/src/lib/report/roster-class.ts
// The "Spec" cell of a roster row. A log without COMBATANT_INFO has no spec for anyone; the
// engine still reads a class off the spells a player cast (`class_source: "inferred"`), and
// that is worth showing. A column never says "Unknown".
import type { Unit } from './types';

export const INFERRED_CLASS_TITLE = 'Class read from the spells cast; the log carried no spec';

export interface UnitClass {
  class: string;
  source: string;
}

/** GUID to the class (and where it came from) the report's units carry. */
export function unitClassMap(units: readonly Unit[]): Map<string, UnitClass> {
  const out = new Map<string, UnitClass>();
  for (const unit of units) {
    if (unit.class) out.set(unit.guid, { class: unit.class, source: unit.class_source ?? '' });
  }
  return out;
}

export interface SpecCell {
  /** The class to colour and crest the row with, when one is known. */
  class: string | undefined;
  text: string;
  title: string | undefined;
  /** True for the dash: nothing is known. */
  muted: boolean;
  /** True when the text is the class standing in for a spec, drawn in the class colour. */
  classColoured: boolean;
}

export function specCell(
  row: { guid: string; class?: string; class_source?: string; spec?: string },
  units: ReadonlyMap<string, UnitClass> = new Map(),
): SpecCell {
  const unit = units.get(row.guid);
  const klass = row.class || unit?.class || undefined;
  const source = row.class ? (row.class_source ?? unit?.source ?? '') : (unit?.source ?? '');
  if (row.spec) return { class: klass, text: row.spec, title: undefined, muted: false, classColoured: false };
  if (klass)
    return {
      class: klass,
      text: klass,
      title: source === 'inferred' ? INFERRED_CLASS_TITLE : undefined,
      muted: false,
      classColoured: true,
    };
  return { class: undefined, text: '—', title: undefined, muted: true, classColoured: false };
}
