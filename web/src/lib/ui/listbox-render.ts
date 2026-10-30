// web/src/lib/ui/listbox-render.ts
// The open list's DOM (design/specs/2026-10-30-controls-charts.md §A4.2 to §A4.5): the
// panel, the filter row, the option rows with their icon, You tag, meta, description and
// check, the group heads, the no-match row and the multi-select foot. Built from data only;
// listbox.ts owns state and events.
import { BUTTON_GHOST, BUTTON_SM } from './button';
import { selectCopy } from './copy';
import { CHECK_GLYPH, CLOSE_GLYPH, SEARCH_GLYPH } from './glyphs';

const LIST_GLYPHS = { check: CHECK_GLYPH, close: CLOSE_GLYPH, search: SEARCH_GLYPH } as const;
type ListGlyph = keyof typeof LIST_GLYPHS;
import { crestRing, labelColour, matchSpan, qualityBorder, type ListRow } from './listbox-model';

const SVG_NS = 'http://www.w3.org/2000/svg';

export function glyph(name: ListGlyph, size: number, className: string): SVGSVGElement {
  const svg = document.createElementNS(SVG_NS, 'svg');
  svg.setAttribute('class', className);
  svg.setAttribute('width', String(size));
  svg.setAttribute('height', String(size));
  svg.setAttribute('viewBox', '0 0 24 24');
  svg.setAttribute('fill', 'none');
  svg.setAttribute('stroke', 'currentColor');
  svg.setAttribute('stroke-width', '2');
  svg.setAttribute('stroke-linecap', 'round');
  svg.setAttribute('stroke-linejoin', 'round');
  svg.setAttribute('aria-hidden', 'true');
  // A constant from glyphs.ts, never user text.
  svg.innerHTML = LIST_GLYPHS[name];
  return svg;
}

function el<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  className: string,
  text?: string,
): HTMLElementTagNameMap[K] {
  const node = document.createElement(tag);
  node.className = className;
  if (text !== undefined) node.textContent = text;
  return node;
}

/** The 20 px icon slot: a ringed crest, or a square spec, spell or item icon (§A4.1). */
export function iconSlot(icon: NonNullable<ListRow['icon']>, colour: string | undefined): HTMLElement {
  const slot = el('span', 'sel-icon');
  slot.dataset.kind = icon.kind;
  const border = qualityBorder(icon);
  if (border !== undefined) slot.style.setProperty('--sel-quality', border);
  const ring = crestRing(icon, colour);
  if (ring !== undefined) slot.style.setProperty('--sel-ring', ring);
  const img = el('img', '');
  img.src = icon.src;
  img.alt = '';
  img.width = 20;
  img.height = 20;
  img.decoding = 'async';
  slot.append(img);
  return slot;
}

/** The label with the filter's matched run in gold (§A4.3 Filter match). */
function labelText(label: string, query: string, id: string, colour: string | undefined): HTMLElement {
  const span = el('span', 'sel-row-label');
  span.id = id;
  if (colour !== undefined) span.style.color = colour;
  const match = matchSpan(label, query);
  if (match === undefined) {
    span.textContent = label;
    return span;
  }
  span.append(
    label.slice(0, match.start),
    el('mark', 'sel-match', label.slice(match.start, match.end)),
    label.slice(match.end),
  );
  return span;
}

export interface RowContext {
  listId: string;
  index: number;
  query: string;
  multiple: boolean;
}

export function optionRow(row: ListRow, context: RowContext): HTMLElement {
  const id = `${context.listId}-opt-${context.index}`;
  const node = el('div', 'sel-row');
  node.id = id;
  node.setAttribute('role', 'option');
  node.setAttribute('aria-selected', row.selected ? 'true' : 'false');
  if (row.disabled) node.setAttribute('aria-disabled', 'true');
  node.dataset.index = String(context.index);
  node.dataset.value = row.value;

  if (context.multiple) {
    const box = el('span', 'sel-box');
    box.append(glyph('check', 12, 'sel-box-check'));
    node.append(box);
  }
  if (row.icon !== undefined) node.append(iconSlot(row.icon, row.colour));

  const text = el('span', 'sel-row-text');
  const line = el('span', 'sel-row-line');
  const labelId = `${id}-label`;
  line.append(labelText(row.label, context.query, labelId, labelColour(row)));
  const named = [labelId];
  if (row.you) {
    const tag = el('span', 'sel-you', selectCopy.you);
    tag.id = `${id}-you`;
    named.push(tag.id);
    line.append(tag);
  }
  text.append(line);
  if (row.description !== undefined) {
    const description = el('span', 'sel-row-desc', row.description);
    description.id = `${id}-desc`;
    node.setAttribute('aria-describedby', description.id);
    text.append(description);
  }
  node.append(text);
  if (row.ariaLabel !== undefined) node.setAttribute('aria-label', row.ariaLabel);
  else node.setAttribute('aria-labelledby', named.join(' '));

  if (row.meta !== undefined) node.append(el('span', 'sel-row-meta', row.meta));
  if (!context.multiple && row.selected) node.append(glyph('check', 16, 'sel-check'));
  return node;
}

/** Rows grouped under their `<optgroup>` heads, in order; ungrouped rows sit bare. */
export function listContent(
  rows: readonly ListRow[],
  indices: readonly number[],
  context: Omit<RowContext, 'index'>,
): DocumentFragment {
  const fragment = document.createDocumentFragment();
  let group: HTMLElement | null = null;
  let groupName: string | undefined;
  indices.forEach((index, position) => {
    const row = rows[index];
    if (row.group !== groupName || position === 0) {
      groupName = row.group;
      group = null;
      if (row.group !== undefined) {
        group = el('div', 'sel-grp');
        group.setAttribute('role', 'group');
        const head = el('div', 'sel-grp-head', row.group);
        head.id = `${context.listId}-grp-${index}`;
        head.setAttribute('role', 'presentation');
        group.setAttribute('aria-labelledby', head.id);
        group.append(head);
        fragment.append(group);
      }
    }
    const node = optionRow(row, { ...context, index });
    (group ?? fragment).append(node);
  });
  return fragment;
}

export interface PanelParts {
  panel: HTMLElement;
  list: HTMLElement;
  scroller: HTMLElement;
  empty: HTMLElement;
  input: HTMLInputElement | null;
  clear: HTMLButtonElement | null;
  count: HTMLElement | null;
  foot: HTMLButtonElement | null;
  footCount: HTMLElement | null;
}

export interface PanelConfig {
  listId: string;
  labelledBy: string;
  searchable: boolean;
  searchPlaceholder: string;
  multiple: boolean;
  dense: boolean;
  footLabel: string | undefined;
}

function filterRow(
  config: PanelConfig,
): Pick<PanelParts, 'input' | 'clear' | 'count'> & { row: HTMLElement } {
  const row = el('div', 'sel-filter');
  const box = el('span', 'sel-filter-box');
  const input = el('input', 'sel-filter-input');
  input.type = 'text';
  input.placeholder = config.searchPlaceholder;
  input.autocomplete = 'off';
  input.spellcheck = false;
  input.setAttribute('role', 'combobox');
  input.setAttribute('aria-autocomplete', 'list');
  input.setAttribute('aria-expanded', 'true');
  input.setAttribute('aria-controls', config.listId);
  input.setAttribute('aria-labelledby', config.labelledBy);
  const clear = el('button', 'sel-filter-clear');
  clear.type = 'button';
  clear.tabIndex = -1;
  clear.hidden = true;
  clear.setAttribute('aria-label', selectCopy.clearFilter);
  clear.append(glyph('close', 14, ''));
  box.append(glyph('search', 16, 'sel-filter-glyph'), input, clear);
  const count = el('span', 'sel-count');
  count.setAttribute('aria-live', 'polite');
  row.append(box, count);
  return { row, input, clear, count };
}

/** The panel: a `popover="manual"` surface holding the filter row, the scrolling listbox
 *  with its fades, the no-match row and the multi-select foot. */
export function buildPanel(config: PanelConfig): PanelParts {
  const panel = el('div', 'sel-pop');
  panel.setAttribute('popover', 'manual');
  panel.dataset.dense = config.dense ? 'true' : 'false';
  const filter = config.searchable ? filterRow(config) : null;
  if (filter !== null) panel.append(filter.row);

  const body = el('div', 'sel-body');
  const scroller = el('div', 'sel-scroll');
  // The rows are reached through aria-activedescendant, so the scrolling region is focusable
  // itself for axe's scrollable-region rule. Tab never lands on it: Tab closes the list first.
  scroller.tabIndex = 0;
  const list = el('div', 'sel-list');
  list.id = config.listId;
  list.setAttribute('role', 'listbox');
  list.setAttribute('aria-labelledby', config.labelledBy);
  if (config.multiple) list.setAttribute('aria-multiselectable', 'true');
  const empty = el('div', 'sel-empty');
  empty.hidden = true;
  scroller.append(list, empty);
  body.append(scroller, el('span', 'sel-fade sel-fade-top'), el('span', 'sel-fade sel-fade-bottom'));
  panel.append(body);

  let foot: HTMLButtonElement | null = null;
  let footCount: HTMLElement | null = null;
  if (config.multiple) {
    const row = el('div', 'sel-foot');
    if (config.footLabel !== undefined) {
      foot = el('button', `${BUTTON_GHOST} ${BUTTON_SM} sel-foot-button`, config.footLabel);
      foot.type = 'button';
      row.append(foot);
    }
    footCount = el('span', 'sel-foot-count');
    row.append(footCount);
    panel.append(row);
  }
  return {
    panel,
    list,
    scroller,
    empty,
    input: filter?.input ?? null,
    clear: filter?.clear ?? null,
    count: filter?.count ?? null,
    foot,
    footCount,
  };
}
