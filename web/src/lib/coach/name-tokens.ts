// web/src/lib/coach/name-tokens.ts
// Names inline (experience spec 3.1, tenets 2 and 3): in titles, why and doing-well lines every
// spell and item name the build's dictionary holds exactly becomes a token (icon and name; items
// open the item tooltip). Longest match first, whole words only, punctuation stays outside the
// token, and a name the dictionary does not hold stays plain text.

export type NameKind = 'spell' | 'item';
export interface NameEntry {
  kind: NameKind;
  id: number;
}
export type NameDictionary = ReadonlyMap<string, NameEntry>;

export type TextPart = { text: string } | { text: string; entry: NameEntry };

const MAX_WORDS = 6;
const WORD = /\S+/g;
const EDGE_PUNCTUATION = /[.,;:!?)"']+$/;
const LEAD_PUNCTUATION = /^[("']+/;

interface Word {
  start: number;
  end: number;
}

/** A word's span without the punctuation that wraps it. */
function trimmed(text: string, word: Word, isLast: boolean): Word {
  const raw = text.slice(word.start, word.end);
  const lead = raw.length - raw.replace(LEAD_PUNCTUATION, '').length;
  const tail = isLast ? raw.length - raw.replace(EDGE_PUNCTUATION, '').length : 0;
  return { start: word.start + lead, end: word.end - tail };
}

function matchAt(
  text: string,
  words: readonly Word[],
  index: number,
  dictionary: NameDictionary,
): { end: number; next: number; entry: NameEntry; start: number } | null {
  for (let length = Math.min(MAX_WORDS, words.length - index); length >= 1; length -= 1) {
    const first = trimmed(text, words[index]!, false);
    const last = trimmed(text, words[index + length - 1]!, true);
    const entry = dictionary.get(text.slice(first.start, last.end));
    if (entry !== undefined) return { start: first.start, end: last.end, next: index + length, entry };
  }
  return null;
}

export function tokenizeNames(text: string, dictionary: NameDictionary): TextPart[] {
  if (dictionary.size === 0) return text === '' ? [] : [{ text }];
  const words = [...text.matchAll(WORD)].map((m) => ({ start: m.index, end: m.index + m[0].length }));
  const parts: TextPart[] = [];
  let cursor = 0;
  let index = 0;
  while (index < words.length) {
    const hit = matchAt(text, words, index, dictionary);
    if (hit === null) {
      index += 1;
      continue;
    }
    if (hit.start > cursor) parts.push({ text: text.slice(cursor, hit.start) });
    parts.push({ text: text.slice(hit.start, hit.end), entry: hit.entry });
    cursor = hit.end;
    index = hit.next;
  }
  if (cursor < text.length) parts.push({ text: text.slice(cursor) });
  return parts;
}

/** A dictionary from id-to-name tables; the first id of a name wins. */
export function buildDictionary(
  spells: Readonly<Record<string, string>>,
  items: Readonly<Record<string, string>>,
): NameDictionary {
  const dictionary = new Map<string, NameEntry>();
  const add = (kind: NameKind, table: Readonly<Record<string, string>>): void => {
    for (const [id, name] of Object.entries(table)) {
      if (name !== '' && !dictionary.has(name)) dictionary.set(name, { kind, id: Number(id) });
    }
  };
  add('spell', spells);
  add('item', items);
  return dictionary;
}
