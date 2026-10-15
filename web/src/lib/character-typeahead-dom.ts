// web/src/lib/character-typeahead-dom.ts
// Typeahead on the character lookup boxes (design/specs/2026-10-15-character-typeahead.md): the
// home page's full box and the header's compact one. The name input is a combobox over a
// listbox of real links; the form's own submit is untouched, so Enter with nothing highlighted,
// or a failed read, behaves exactly as it did before. Plain DOM, no island: the listeners cost
// nothing until a key is pressed, and the read code is imported on the first query only.
import type { CharacterSearch, CharacterSuggestion } from './character-search-api';
import { characterHref, rulesetLabel } from './characters';
import classes from '../data/classes.json';
import { classCrestSrc } from './class-crest';
import { homeCharacterSearchCopy as copy } from './home-landing-copy';

/** The pause after the last keystroke before the endpoint is asked. */
export const DEBOUNCE_MS = 150;
/** The endpoint answers nothing under this many characters. */
export const MIN_QUERY_LENGTH = 2;
/** Rows asked for, and the most the list shows. */
export const SUGGESTION_LIMIT = 8;

export interface TypeaheadDeps {
  search: CharacterSearch;
  /** Opens a row's page from the keyboard (a click is the link itself). */
  navigate: (href: string) => void;
}

const defaultDeps: TypeaheadDeps = {
  // Imported on the first query, so the read code and its transport stay off every page load.
  search: async (text, options) => (await import('./character-search-api')).searchCharacters(text, options),
  navigate: (href) => window.location.assign(href),
};

const classBySlug = new Map(classes.map((entry) => [entry.slug, entry]));

/** The class a suggestion names (`Mage`, or its slug), or undefined for null and the unknown. */
function classOf(row: CharacterSuggestion): (typeof classes)[number] | undefined {
  return row.class === null ? undefined : classBySlug.get(row.class.toLowerCase());
}

function element<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  className: string,
  text?: string,
): HTMLElementTagNameMap[K] {
  const node = document.createElement(tag);
  node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
}

function crestFor(klass: (typeof classes)[number] | undefined): HTMLElement {
  if (klass === undefined) return element('span', 'typeahead-crest typeahead-crest-none');
  const crest = element('img', 'typeahead-crest');
  crest.src = classCrestSrc(klass.slug);
  crest.alt = '';
  crest.width = 28;
  crest.height = 28;
  return crest;
}

/** "{Spec} {Class} · {Level}", with whatever of it is known. */
function detailText(row: CharacterSuggestion): string {
  const kind = [row.spec, row.class].filter((part): part is string => part !== null).join(' ');
  return [kind, row.level === null ? '' : String(row.level)].filter((part) => part !== '').join(' · ');
}

function renderRow(row: CharacterSuggestion, id: string): HTMLLIElement {
  const klass = classOf(row);
  const item = element('li', 'typeahead-item');
  item.setAttribute('role', 'presentation');
  const link = element('a', klass === undefined ? 'typeahead-row typeahead-row-plain' : 'typeahead-row');
  link.id = id;
  link.href = characterHref(row.region, row.ruleset, row.name);
  link.tabIndex = -1;
  link.setAttribute('role', 'option');
  link.setAttribute('aria-selected', 'false');
  if (klass !== undefined) link.style.setProperty('--c', klass.color);
  const text = element('span', 'typeahead-text');
  text.append(element('span', 'typeahead-name', row.name));
  const meta = element('span', 'typeahead-meta');
  const detail = detailText(row);
  if (detail !== '') meta.append(element('span', 'typeahead-detail', detail));
  meta.append(element('span', 'typeahead-realm', `${row.region.toUpperCase()} ${rulesetLabel(row.ruleset)}`));
  text.append(meta);
  link.append(crestFor(klass), text);
  item.append(link);
  return item;
}

function renderEmpty(typed: string): HTMLLIElement {
  const item = element('li', 'typeahead-empty', copy.typeaheadEmpty(typed));
  item.setAttribute('role', 'option');
  item.setAttribute('aria-disabled', 'true');
  return item;
}

/** The `region` / `ruleset` fields a form has, as search filters; a form without them searches all. */
function filtersOf(form: HTMLFormElement): { region?: string; ruleset?: string } {
  const value = (name: string): string | undefined => {
    const field = form.elements.namedItem(name);
    return field instanceof HTMLSelectElement || field instanceof HTMLInputElement ? field.value : undefined;
  };
  return { region: value('region'), ruleset: value('ruleset') };
}

/** Wires the suggestion list of one `form[data-character-search]` that has a `[data-typeahead]`. */
export function wireTypeahead(form: HTMLFormElement, deps: TypeaheadDeps = defaultDeps): void {
  const input = form.querySelector<HTMLInputElement>('input[name="name"]');
  const list = form.querySelector<HTMLUListElement>('[data-typeahead] [role="listbox"]');
  if (input === null || list === null || form.dataset.typeaheadWired === '1') return;
  form.dataset.typeaheadWired = '1';
  list.setAttribute('aria-label', copy.typeaheadLabel);

  let rows: CharacterSuggestion[] = [];
  let highlighted = -1;
  let pause: ReturnType<typeof setTimeout> | undefined;
  let inFlight: AbortController | undefined;

  const links = (): HTMLAnchorElement[] => Array.from(list.querySelectorAll('a[role="option"]'));

  const highlight = (index: number): void => {
    highlighted = index;
    links().forEach((link, at) => link.setAttribute('aria-selected', String(at === index)));
    if (index < 0) input.removeAttribute('aria-activedescendant');
    else input.setAttribute('aria-activedescendant', links()[index].id);
    links()[index]?.scrollIntoView?.({ block: 'nearest' });
  };

  const setOpen = (open: boolean): void => {
    list.hidden = !open;
    input.setAttribute('aria-expanded', String(open));
    if (!open) highlight(-1);
  };

  const cancelPending = (): void => {
    clearTimeout(pause);
    inFlight?.abort();
    inFlight = undefined;
  };

  const clear = (): void => {
    cancelPending();
    rows = [];
    list.replaceChildren();
    setOpen(false);
  };

  const show = (typed: string, found: CharacterSuggestion[]): void => {
    rows = found.slice(0, SUGGESTION_LIMIT);
    list.replaceChildren(
      ...(rows.length === 0
        ? [renderEmpty(typed)]
        : rows.map((row, at) => renderRow(row, `${list.id}-option-${at}`))),
    );
    setOpen(true);
  };

  const ask = async (typed: string): Promise<void> => {
    const controller = new AbortController();
    inFlight = controller;
    try {
      const found = await deps.search(typed, {
        ...filtersOf(form),
        limit: SUGGESTION_LIMIT,
        signal: controller.signal,
      });
      // A late answer never reopens a list the player has already left.
      if (!controller.signal.aborted && document.activeElement === input) show(typed, found);
    } catch {
      if (!controller.signal.aborted) clear();
    }
  };

  const schedule = (): void => {
    cancelPending();
    const typed = input.value.trim();
    if (typed.length < MIN_QUERY_LENGTH) {
      clear();
      return;
    }
    pause = setTimeout(() => void ask(typed), DEBOUNCE_MS);
  };

  const step = (direction: 1 | -1): void => {
    const count = links().length;
    if (count === 0) return;
    if (list.hidden) setOpen(true);
    const from = highlighted < 0 ? (direction === 1 ? -1 : 0) : highlighted;
    highlight((from + direction + count) % count);
  };

  input.addEventListener('input', schedule);
  form.addEventListener('submit', clear);
  form.querySelectorAll('select').forEach((select) => select.addEventListener('change', schedule));
  input.addEventListener('blur', () => setOpen(false));
  list.addEventListener('mousedown', (event) => event.preventDefault());
  list.addEventListener('mouseover', (event) => {
    const link = (event.target as Element).closest<HTMLAnchorElement>('a[role="option"]');
    if (link !== null) highlight(links().indexOf(link));
  });
  input.addEventListener('keydown', (event) => {
    if (event.key === 'ArrowDown' || event.key === 'ArrowUp') {
      event.preventDefault();
      step(event.key === 'ArrowDown' ? 1 : -1);
    } else if (event.key === 'Enter' && highlighted >= 0) {
      event.preventDefault();
      deps.navigate(links()[highlighted].getAttribute('href') ?? '');
    } else if (event.key === 'Escape' && !list.hidden) {
      event.preventDefault();
      event.stopPropagation();
      setOpen(false);
    } else if (event.key === 'Tab') {
      setOpen(false);
    }
  });
}
