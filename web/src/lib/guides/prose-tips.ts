// web/src/lib/guides/prose-tips.ts
// Marks the spell, talent, racial and item names inside a guide's rendered prose so each one
// opens the game's tooltip, the way every name does on the rest of the site. The names come
// from the client data (lib/guides/guide-tips.ts's `proseNames`), never a list typed per guide.
//
// Rules: a name matches whole and with its own capitals ("Execute", never "execute"), the
// longest name wins where two overlap ("Improved Heroic Strike" over "Heroic Strike"), and
// nothing inside a heading, a link, code or an already marked name is touched. Every mention
// is marked, as a game guide's links are.
import { rarityClassFor } from '../planner/items';
import { escapeHtml } from '../html';

export type ProseTarget = { key: string } | { itemId: number; quality: number };

/** The tags whose text is never marked: headings keep their anchors, links their own target. */
const SKIP_TAGS = new Set([
  'a',
  'h1',
  'h2',
  'h3',
  'h4',
  'h5',
  'h6',
  'code',
  'pre',
  'script',
  'style',
  'button',
]);
/** The marker this module writes, so a second pass never marks inside the first. */
const MARK_CLASS = 'gtip';

const APOSTROPHE = "(?:'|’|&#39;|&#x27;|&apos;|&rsquo;)";

function escapeRegExp(text: string): string {
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

function namePattern(name: string): string {
  return escapeRegExp(name).split("'").join(APOSTROPHE).replace(/&/g, '(?:&|&amp;)');
}

/** The text a match spelled, back in the name's own form, for the lookup. */
function canonical(text: string): string {
  return text.replace(/’|&#39;|&#x27;|&apos;|&rsquo;/g, "'").replace(/&amp;/g, '&');
}

/** The marked span for one mention. */
export function markFor(target: ProseTarget, text: string): string {
  if ('key' in target) {
    return `<span class="${MARK_CLASS}" tabindex="0" data-game-tip="${escapeHtml(target.key)}">${text}</span>`;
  }
  const quality = rarityClassFor(target.quality);
  return `<span class="${MARK_CLASS} ${MARK_CLASS}-item ${quality}" tabindex="0" data-bis-item="${target.itemId}">${text}</span>`;
}

function matcher(names: ReadonlyMap<string, ProseTarget>): RegExp | undefined {
  const sorted = [...names.keys()].filter((name) => name.trim() !== '').sort((a, b) => b.length - a.length);
  if (sorted.length === 0) return undefined;
  const alternatives = sorted.map(namePattern).join('|');
  return new RegExp(`(?<![\\p{L}\\p{N}'’])(?:${alternatives})(?![\\p{L}\\p{N}]|${APOSTROPHE}[\\p{L}])`, 'gu');
}

function tagName(tag: string): { name: string; closing: boolean; selfClosing: boolean } | undefined {
  const match = /^<\s*(\/?)\s*([a-zA-Z][a-zA-Z0-9]*)/.exec(tag);
  if (match === null) return undefined;
  return { name: match[2]!.toLowerCase(), closing: match[1] === '/', selfClosing: /\/\s*>$/.test(tag) };
}

/** `html` with every known name outside a heading, link or code marked. */
export function linkProse(html: string, names: ReadonlyMap<string, ProseTarget>): string {
  const pattern = matcher(names);
  if (pattern === undefined || html === '') return html;
  // Open skip tags, innermost last; a marked span counts as one so its text stays as it is.
  const skipping: string[] = [];
  return html
    .split(/(<[^>]*>)/)
    .map((part) => {
      if (part.startsWith('<')) {
        const tag = tagName(part);
        if (tag === undefined || tag.selfClosing) return part;
        const marked = tag.name === 'span' && part.includes(`class="${MARK_CLASS}`);
        if (!tag.closing && (SKIP_TAGS.has(tag.name) || marked)) skipping.push(tag.name);
        else if (!tag.closing && skipping.length > 0 && tag.name === 'span') skipping.push('span');
        else if (tag.closing && skipping.length > 0 && skipping[skipping.length - 1] === tag.name)
          skipping.pop();
        return part;
      }
      if (skipping.length > 0 || part === '') return part;
      return part.replace(pattern, (text) => {
        const target = names.get(canonical(text));
        return target === undefined ? text : markFor(target, text);
      });
    })
    .join('');
}
