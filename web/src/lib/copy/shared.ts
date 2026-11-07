// web/src/lib/copy/shared.ts
// The words the site and both apps show (apps/shared/copy/shared.json, mobile redesign spec 2.2).
// The web copy modules read their shared words through these two functions and keep their own
// names and exports, so no call site changes. Keys are typed from the JSON itself: a key that is
// not in the catalog does not compile.
import catalog from '../../../../apps/shared/copy/shared.json';

type Strings = typeof catalog.strings;
export type SharedKey = keyof Strings;

type Values = Readonly<Record<string, string | number>>;
interface TextEntry {
  text: string;
}
interface PluralEntry {
  count: string;
  one: string;
  other: string;
}

const PLACEHOLDER = /\{([a-zA-Z][a-zA-Z0-9]*)\}/g;

function entry(key: SharedKey): unknown {
  return (catalog.strings as Record<string, unknown>)[key];
}

/** `{name}` placeholders filled from `values`; a placeholder without a value fails loudly. */
function fill(key: SharedKey, template: string, values: Values): string {
  return template.replace(PLACEHOLDER, (_, name: string) => {
    const value = values[name];
    if (value === undefined) throw new Error(`shared copy ${key}: no value for {${name}}`);
    return String(value);
  });
}

/** A shared string, its placeholders filled. */
export function sharedText(key: SharedKey, values: Values = {}): string {
  const raw = entry(key);
  if (typeof raw === 'string') return fill(key, raw, values);
  if (raw !== null && typeof raw === 'object' && 'text' in raw)
    return fill(key, (raw as TextEntry).text, values);
  throw new Error(`shared copy ${key}: not a text entry`);
}

/** A shared plural: its `one` form at 1, else `other` with the count filled in. */
export function sharedPlural(key: SharedKey, count: number): string {
  const raw = entry(key);
  if (raw === null || typeof raw !== 'object' || !('count' in raw)) {
    throw new Error(`shared copy ${key}: not a plural entry`);
  }
  const plural = raw as PluralEntry;
  return fill(key, count === 1 ? plural.one : plural.other, { [plural.count]: count });
}
