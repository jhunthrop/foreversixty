// web/src/lib/ui/listbox-model.ts
// The shared dropdown's pure rules (design/specs/2026-10-30-controls-charts.md §A): the option
// data the native <option> carries, the keyboard moves over enabled rows, typeahead, the
// filter match, and where the list opens. No DOM writes here; listbox.ts draws what these
// decide, and Select.svelte reads the same option shape.

export type IconKind = 'crest' | 'spec' | 'spell' | 'item';
export type Quality = 'poor' | 'common' | 'uncommon' | 'rare' | 'epic' | 'legendary';

export interface OptionIcon {
  kind: IconKind;
  src: string;
  quality?: Quality;
}

/** One option, as call sites pass it to Select and as the enhancer reads it back. */
export interface SelectOption {
  value: string;
  label: string;
  group?: string;
  description?: string;
  /** Right-aligned muted figure, for example `×8` or a DPS. */
  meta?: string;
  icon?: OptionIcon;
  /** A token name, for example `--color-class-mage`. Class names only. */
  colour?: string;
  disabled?: boolean;
  /** The viewer's own character: the `You` tag. */
  you?: boolean;
  /** The option's accessible name when its label is a mark, not words. */
  ariaLabel?: string;
}

/** An option as the open list holds it: its data plus its selected state. */
export interface ListRow extends SelectOption {
  selected: boolean;
}

/** The typeahead buffer clears after this long without a key (§A6.3). */
export const TYPEAHEAD_MS = 500;
/** PageUp and PageDown move by this many rows (§A6.3). */
export const PAGE_ROWS = 10;
/** Below this much room under the field, the list flips above it when more fits there (§A4.2). */
export const FLIP_THRESHOLD_PX = 240;
/** The list's tallest (§A4.2). */
export const LIST_MAX_HEIGHT_PX = 360;
/** The list's widest (§A4.2). */
export const LIST_MAX_WIDTH_PX = 420;
/** The gap from the field and the margin from every viewport edge (§A4.2). */
export const LIST_OFFSET_PX = 4;
export const VIEWPORT_PADDING_PX = 8;
/** Icons requested ahead of the first open (§A11.3). */
export const PRELOAD_ICON_COUNT = 12;

const QUALITIES: readonly Quality[] = ['poor', 'common', 'uncommon', 'rare', 'epic', 'legendary'];
const ICON_KINDS: readonly IconKind[] = ['crest', 'spec', 'spell', 'item'];

function isIconKind(value: string | undefined): value is IconKind {
  return value !== undefined && (ICON_KINDS as readonly string[]).includes(value);
}

function isQuality(value: string | undefined): value is Quality {
  return value !== undefined && (QUALITIES as readonly string[]).includes(value);
}

/** The data attributes an `<option>` carries for the enhancer (§A3). */
export function optionAttributes(option: SelectOption): Record<string, string | undefined> {
  return {
    'data-icon': option.icon?.src,
    'data-icon-kind': option.icon?.kind,
    'data-quality': option.icon?.quality,
    'data-description': option.description,
    'data-meta': option.meta,
    'data-colour': option.colour,
    'data-you': option.you ? 'true' : undefined,
  };
}

/** Reads an `<option>` element back into the option shape (the inverse of optionAttributes). */
export function readOption(element: HTMLOptionElement): ListRow {
  const data = element.dataset;
  const kind = data.iconKind;
  const parent = element.parentElement;
  const group = parent instanceof HTMLOptGroupElement ? parent.label : undefined;
  return {
    value: element.value,
    label: element.textContent?.trim() ?? '',
    group,
    description: data.description,
    meta: data.meta,
    icon:
      data.icon !== undefined && isIconKind(kind)
        ? { kind, src: data.icon, quality: isQuality(data.quality) ? data.quality : undefined }
        : undefined,
    colour: data.colour,
    disabled: element.disabled,
    you: data.you === 'true',
    ariaLabel: element.getAttribute('aria-label') ?? undefined,
    selected: element.selected,
  };
}

/** Every option of a select, in order, skipping the placeholder's empty first option. */
export function readOptions(select: HTMLSelectElement, placeholder: boolean): ListRow[] {
  const rows = [...select.options].map(readOption);
  return placeholder && rows[0]?.value === '' ? rows.slice(1) : rows;
}

/** The option the closed field shows: the picked one, else undefined (placeholder). */
export function pickedOption<T extends SelectOption>(options: readonly T[], value: string): T | undefined {
  return options.find((option) => option.value === value);
}

/** The next enabled row from `from` by `step` (±1, ±PAGE_ROWS), clamped, never wrapping (§A6.3). */
export function moveActive(rows: readonly ListRow[], from: number, step: number): number {
  if (rows.length === 0) return -1;
  const direction = Math.sign(step) || 1;
  const target = Math.max(0, Math.min(rows.length - 1, from + step));
  for (let index = target; index >= 0 && index < rows.length; index += direction) {
    if (!rows[index].disabled) return index;
  }
  // Nothing enabled past the target: walk back towards `from` for the last enabled row.
  for (let index = target - direction; index >= 0 && index < rows.length; index -= direction) {
    if (!rows[index].disabled) return index;
  }
  return from;
}

/** The first or last enabled row. */
export function edgeRow(rows: readonly ListRow[], edge: 'first' | 'last'): number {
  const indices = rows.map((_, index) => index);
  const ordered = edge === 'first' ? indices : indices.reverse();
  return ordered.find((index) => !rows[index].disabled) ?? -1;
}

/** Folds case and accents so `Lúcio` matches `luc` (§A4.4). */
export function foldText(text: string): string {
  return text
    .normalize('NFD')
    .replace(/\p{Diacritic}/gu, '')
    .toLocaleLowerCase();
}

/** The first enabled row whose label starts with the typed run, searching from after
 *  `from` and wrapping, so a repeated letter cycles through its rows. -1 when none. */
export function typeaheadMatch(rows: readonly ListRow[], run: string, from: number): number {
  const needle = foldText(run);
  if (needle === '') return -1;
  // A run of one repeated letter cycles from the next row; a longer run keeps the current
  // row while it still fits.
  const cycles = [...needle].every((char) => char === needle[0]);
  const start = cycles ? from + 1 : from;
  const probe = cycles ? needle[0] : needle;
  for (let offset = 0; offset < rows.length; offset += 1) {
    const index = (((start + offset) % rows.length) + rows.length) % rows.length;
    const row = rows[index];
    if (!row.disabled && foldText(row.label).startsWith(probe)) return index;
  }
  return -1;
}

/** The typeahead buffer: keys within TYPEAHEAD_MS extend the run, a pause starts a new one. */
export function createTypeahead(now: () => number = () => Date.now()) {
  let run = '';
  let last = -Infinity;
  return {
    push(char: string): string {
      const time = now();
      run = time - last > TYPEAHEAD_MS ? char : run + char;
      last = time;
      return run;
    },
    reset(): void {
      run = '';
      last = -Infinity;
    },
  };
}

/** Where a match sits in a label, for the gold run (§A4.3 Filter match). */
export interface MatchSpan {
  start: number;
  end: number;
}

export function matchSpan(label: string, query: string): MatchSpan | undefined {
  const needle = foldText(query.trim());
  if (needle === '') return undefined;
  const start = foldText(label).indexOf(needle);
  return start === -1 ? undefined : { start, end: start + needle.length };
}

/** The rows a filter keeps: a case- and accent-insensitive substring match on the label. */
export function filterRows<T extends ListRow>(rows: readonly T[], query: string): T[] {
  const needle = foldText(query.trim());
  if (needle === '') return [...rows];
  return rows.filter((row) => foldText(row.label).includes(needle));
}

/** The row the open list scrolls to: the picked one, or the first enabled one. */
export function initialActive(rows: readonly ListRow[]): number {
  const picked = rows.findIndex((row) => row.selected && !row.disabled);
  return picked === -1 ? edgeRow(rows, 'first') : picked;
}

/** The scrollTop that puts a row in the middle of the list, or at the top near the start (§A4.2). */
export function centredScrollTop(rowTop: number, rowHeight: number, viewportHeight: number): number {
  return Math.max(0, Math.round(rowTop + rowHeight / 2 - viewportHeight / 2));
}

export type Side = 'bottom' | 'top';

/** Below, unless less than FLIP_THRESHOLD_PX fits there (or less than the list needs) and
 *  more fits above (§A4.2). */
export function chooseSide(spaceBelow: number, spaceAbove: number, listHeight: number): Side {
  const needed = Math.min(FLIP_THRESHOLD_PX, listHeight);
  return spaceBelow < needed && spaceAbove > spaceBelow ? 'top' : 'bottom';
}

/** The list's max height on the chosen side: min(cap, available − 16) (§A4.2). The rows cap
 *  at 360 px; the whole panel, with its filter and foot, caps at the room alone. */
export function listMaxHeight(available: number, cap: number = LIST_MAX_HEIGHT_PX): number {
  return Math.max(0, Math.min(cap, available - 2 * VIEWPORT_PADDING_PX));
}

/** The pointer test for the listbox: a fine pointer that can hover (D3). */
export const FINE_POINTER_QUERY = '(pointer: fine) and (hover: hover)';

/** True where the Popover API exists (§A3.8): without it the native select stays. */
export function popoverSupported(): boolean {
  return typeof HTMLElement !== 'undefined' && typeof HTMLElement.prototype.showPopover === 'function';
}

/** A token name only (`--color-class-mage`), so an option's colour can never inject CSS. */
const TOKEN = /^--[a-z0-9-]+$/;
/** Qualities whose border colour is too dark for text carry a lighter `-text` twin. */
const QUALITY_TEXT_TWIN: ReadonlySet<Quality> = new Set<Quality>(['rare', 'epic']);

/** The label's colour: an item's quality text colour, else the option's class token. */
export function labelColour(row: Pick<SelectOption, 'colour' | 'icon'>): string | undefined {
  const quality = row.icon?.kind === 'item' ? row.icon.quality : undefined;
  if (quality !== undefined) {
    return `var(--color-rarity-${quality}${QUALITY_TEXT_TWIN.has(quality) ? '-text' : ''})`;
  }
  return row.colour !== undefined && TOKEN.test(row.colour) ? `var(${row.colour})` : undefined;
}

/** An item icon's border colour (§A4.1). */
export function qualityBorder(icon: OptionIcon | undefined): string | undefined {
  return icon?.quality === undefined ? undefined : `var(--color-rarity-${icon.quality})`;
}

/** The event the multi-select foot's button dispatches on the field (`Top 8`, §A4.5). */
export const FOOT_EVENT = 'selectfoot';

/** A crest's class ring: the option's class token (§A4.1). */
export function crestRing(icon: OptionIcon | undefined, colour: string | undefined): string | undefined {
  return icon?.kind === 'crest' && colour !== undefined && TOKEN.test(colour) ? `var(${colour})` : undefined;
}
