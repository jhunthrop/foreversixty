// web/src/lib/ui/listbox.ts
// The shared dropdown's enhancer (design/specs/2026-10-30-controls-charts.md §A3, §A6). On a
// fine pointer it takes a Select field (`[data-listbox]`), hides its native control (the
// `hidden` attribute keeps the select in form submission) and adds a trigger button and a
// listbox shown in the top layer. The native <select> stays the source of truth: committing
// sets its value and dispatches `input` and `change`, so every call site's handler runs
// unchanged. Framework-free: the static pages' fields use it with no Svelte runtime.
//
// Focus (§A6.1): a plain list keeps focus on the trigger, the one combobox, which owns
// aria-activedescendant; a searchable or multi-select list moves focus into its filter input,
// which is then the one combobox.
import { selectCopy } from './copy';
import {
  FOOT_EVENT,
  PAGE_ROWS,
  PRELOAD_ICON_COUNT,
  centredScrollTop,
  createTypeahead,
  edgeRow,
  filterRows,
  initialActive,
  moveActive,
  readFootActions,
  readOptions,
  typeaheadMatch,
  type ListRow,
} from './listbox-model';
import { buildPanel, footContent, listContent, type PanelParts } from './listbox-render';
import { openPopover, type PopoverHandle } from './popover';

export interface Listbox {
  open: () => void;
  close: () => void;
  /** Re-reads the select (options, value, disabled, error) after the page changed it. */
  refresh: () => void;
  /** Puts the native control back. */
  destroy: () => void;
}

const instances = new WeakMap<HTMLElement, Listbox>();
let uid = 0;

interface Parts {
  field: HTMLElement;
  select: HTMLSelectElement;
  natives: HTMLElement[];
  trigger: HTMLButtonElement;
  triggerText: HTMLSpanElement;
  nameSource: HTMLElement | null;
  labelledBy: string;
}

function ensureId(element: HTMLElement, prefix: string): string {
  if (element.id === '') {
    uid += 1;
    element.id = `${prefix}-${uid}`;
  }
  return element.id;
}

/** The ids that name the field: its visible <label>, else a hidden span holding its aria-label. */
function nameIds(field: HTMLElement, select: HTMLSelectElement): { ids: string; source: HTMLElement | null } {
  const label = select.labels?.[0];
  if (label !== undefined) return { ids: ensureId(label, 'sel-label'), source: null };
  const source = document.createElement('span');
  source.hidden = true;
  source.textContent = select.getAttribute('aria-label') ?? '';
  field.append(source);
  return { ids: ensureId(source, 'sel-name'), source };
}

function buildTrigger(field: HTMLElement, select: HTMLSelectElement): Parts {
  const natives = [...field.querySelectorAll<HTMLElement>(':scope > .sel-native')];
  const trigger = document.createElement('button');
  trigger.type = 'button';
  trigger.className = 'sel-trigger';
  const selectId = ensureId(select, 'sel');
  trigger.id = `${selectId}-trigger`;
  trigger.setAttribute('aria-haspopup', 'listbox');
  trigger.setAttribute('aria-expanded', 'false');
  const triggerText = document.createElement('span');
  triggerText.className = 'sr-only';
  trigger.append(triggerText);
  const { ids, source } = nameIds(field, select);
  const testid = select.dataset.testid;
  if (testid !== undefined) trigger.dataset.testid = `${testid}-trigger`;
  const hadFocus = natives.some((native) => native === document.activeElement);
  field.append(trigger);
  natives.forEach((native) => {
    native.hidden = true;
  });
  if (hadFocus) trigger.focus();
  return { field, select, natives, trigger, triggerText, nameSource: source, labelledBy: ids };
}

/** Builds and wires one enhanced field. */
function selectOf(field: HTMLElement): HTMLSelectElement {
  const select = field.querySelector<HTMLSelectElement>('select');
  if (select === null) throw new Error('A [data-listbox] field needs its <select>.');
  return select;
}

function create(field: HTMLElement): Listbox {
  const select = selectOf(field);
  const parts = buildTrigger(field, select);
  const { trigger, triggerText } = parts;
  const listId = `${select.id}-list`;
  trigger.setAttribute('aria-controls', listId);

  const config = () => {
    const data = field.dataset;
    const multiple = data.multiple === 'true';
    return {
      multiple,
      searchable: multiple || data.searchable === 'true',
      placeholder: data.placeholder === 'true',
      align: data.align === 'end' ? ('end' as const) : ('start' as const),
      dense: data.size === 'dense',
    };
  };

  let rows: ListRow[] = [];
  let visible: number[] = [];
  let active = -1;
  let query = '';
  let panel: PanelParts | null = null;
  let handle: PopoverHandle | null = null;
  let iconsPreloaded = false;
  /** The last pointer type: a touch grows the rows to 44 px until a mouse moves (§A4.3). */
  let lastPointer = 'mouse';
  const typeahead = createTypeahead();

  const isOpen = (): boolean => handle !== null;
  const blocked = (): boolean => select.disabled || field.dataset.loading === 'true';

  function syncTrigger(): void {
    const { searchable } = config();
    const face = field.querySelector('.sel-label')?.textContent?.trim() ?? '';
    triggerText.textContent = face;
    trigger.disabled = blocked();
    if (searchable) trigger.removeAttribute('role');
    else trigger.setAttribute('role', 'combobox');
    trigger.setAttribute('aria-labelledby', `${parts.labelledBy} ${trigger.id}`);
    for (const name of ['aria-describedby', 'aria-invalid', 'aria-busy']) {
      const value = select.getAttribute(name);
      if (value === null) trigger.removeAttribute(name);
      else trigger.setAttribute(name, value);
    }
    if (parts.nameSource !== null) parts.nameSource.textContent = select.getAttribute('aria-label') ?? '';
  }

  /** The element that holds aria-activedescendant: the filter input, else the trigger. */
  function focusHolder(): HTMLElement {
    return panel?.input ?? trigger;
  }

  function rowNode(index: number): HTMLElement | null {
    return panel?.list.querySelector<HTMLElement>(`[data-index="${index}"]`) ?? null;
  }

  function setActive(index: number, scroll: 'nearest' | 'centre' | 'none'): void {
    rowNode(active)?.classList.remove('is-active');
    active = index;
    const node = rowNode(index);
    if (node === null) {
      focusHolder().removeAttribute('aria-activedescendant');
      return;
    }
    node.classList.add('is-active');
    focusHolder().setAttribute('aria-activedescendant', node.id);
    if (scroll === 'nearest') node.scrollIntoView({ block: 'nearest' });
    if (scroll === 'centre' && panel !== null) {
      panel.scroller.scrollTop = centredScrollTop(
        node.offsetTop,
        node.offsetHeight,
        panel.scroller.clientHeight,
      );
    }
  }

  function updateFades(): void {
    if (panel === null) return;
    const { scroller, panel: surface } = panel;
    surface.dataset.fadeTop = scroller.scrollTop > 0 ? 'true' : 'false';
    surface.dataset.fadeBottom =
      scroller.scrollTop + scroller.clientHeight < scroller.scrollHeight - 1 ? 'true' : 'false';
  }

  function render(): void {
    if (panel === null) return;
    const { multiple } = config();
    const kept = filterRows(
      rows.map((row, index) => ({ ...row, index })),
      query,
    );
    visible = kept.map((row) => row.index);
    panel.list.replaceChildren(listContent(rows, visible, { listId, query, multiple }));
    panel.empty.hidden = visible.length > 0;
    panel.empty.textContent = visible.length > 0 ? '' : selectCopy.noMatch(query.trim());
    if (panel.count !== null)
      panel.count.textContent = query.trim() === '' ? '' : selectCopy.matchCount(visible.length);
    if (panel.clear !== null) panel.clear.hidden = query === '';
    if (panel.footCount !== null) panel.footCount.textContent = field.dataset.footCount ?? '';
    panel.footActions?.replaceChildren(footContent(readFootActions(field.dataset.foot)));
  }

  /** Moves among the visible rows by `step` from the active one. */
  function step(by: number): void {
    const shown = visible.map((index) => rows[index]);
    const from = Math.max(0, visible.indexOf(active));
    const next = active === -1 ? edgeRow(shown, by > 0 ? 'first' : 'last') : moveActive(shown, from, by);
    if (next !== -1) setActive(visible[next], 'nearest');
  }

  function edge(which: 'first' | 'last'): void {
    const next = edgeRow(
      visible.map((index) => rows[index]),
      which,
    );
    if (next !== -1) setActive(visible[next], 'nearest');
  }

  function dispatchChange(): void {
    select.dispatchEvent(new Event('input', { bubbles: true }));
    select.dispatchEvent(new Event('change', { bubbles: true }));
  }

  /** Commits a row: single sets the value and closes; multi toggles and stays open (§A6.3). */
  function commit(index: number): void {
    const row = rows[index];
    if (row === undefined || row.disabled) return;
    if (config().multiple) {
      const option = [...select.options].find((candidate) => candidate.value === row.value);
      if (option === undefined) return;
      option.selected = !option.selected;
      dispatchChange();
      reload();
      return;
    }
    const changed = select.value !== row.value;
    close(true);
    if (changed) {
      select.value = row.value;
      field.removeAttribute('data-invalid');
      dispatchChange();
    }
  }

  /** Re-reads the options into an open list, keeping the active row (a toggle, or the page
   *  changing the options under it). */
  function reload(): void {
    if (!isOpen() || panel === null) return;
    if (blocked()) {
      close(false);
      return;
    }
    const keep = rows[active]?.value;
    rows = readOptions(select, config().placeholder);
    render();
    const index = rows.findIndex((row) => row.value === keep);
    setActive(index === -1 ? initialActive(rows) : index, 'none');
    updateFades();
    handle?.update();
  }

  function onScroll(): void {
    updateFades();
  }

  function onListPointerMove(event: PointerEvent): void {
    if (event.pointerType === 'mouse') panel?.panel.removeAttribute('data-touch');
    const row = (event.target as Element | null)?.closest<HTMLElement>('[role="option"]');
    if (row === null || row === undefined || row.getAttribute('aria-disabled') === 'true') return;
    const index = Number(row.dataset.index);
    if (index !== active) setActive(index, 'none');
  }

  function onListClick(event: MouseEvent): void {
    const row = (event.target as Element | null)?.closest<HTMLElement>('[role="option"]');
    if (row === null || row === undefined) return;
    commit(Number(row.dataset.index));
  }

  /** Focus never leaves the trigger or the filter for a row (§A6.1). */
  function keepFocus(event: Event): void {
    if (!(event.target instanceof HTMLInputElement)) event.preventDefault();
  }

  function onFilterInput(): void {
    if (panel?.input == null) return;
    query = panel.input.value;
    render();
    const first = visible.find((index) => !rows[index].disabled);
    setActive(first ?? -1, 'none');
    if (panel.scroller.scrollTop !== 0) panel.scroller.scrollTop = 0;
    updateFades();
    handle?.update();
  }

  function onClear(): void {
    if (panel?.input == null) return;
    panel.input.value = '';
    onFilterInput();
    panel.input.focus();
  }

  function onFoot(event: MouseEvent): void {
    const button = (event.target as Element | null)?.closest<HTMLButtonElement>('button[data-foot-id]');
    if (button === null || button === undefined) return;
    field.dispatchEvent(
      new CustomEvent(FOOT_EVENT, { bubbles: true, detail: { id: button.dataset.footId } }),
    );
  }

  function mountPanel(): PanelParts {
    const { searchable, multiple, dense } = config();
    const built = buildPanel({
      listId,
      labelledBy: parts.labelledBy,
      searchable,
      searchPlaceholder: field.dataset.searchPlaceholder ?? '',
      multiple,
      dense,
    });
    built.scroller.addEventListener('scroll', onScroll, { passive: true });
    built.list.addEventListener('pointermove', onListPointerMove);
    built.list.addEventListener('click', onListClick);
    built.panel.addEventListener('mousedown', keepFocus);
    built.input?.addEventListener('input', onFilterInput);
    built.input?.addEventListener('keydown', onKeydown);
    built.clear?.addEventListener('click', onClear);
    built.footActions?.addEventListener('click', onFoot);
    field.append(built.panel);
    return built;
  }

  function open(initial?: 'first' | 'last'): void {
    if (isOpen() || blocked()) return;
    const { placeholder, align } = config();
    rows = readOptions(select, placeholder);
    query = '';
    panel?.panel.remove();
    panel = mountPanel();
    if (lastPointer === 'touch') panel.panel.dataset.touch = 'true';
    render();
    trigger.setAttribute('aria-expanded', 'true');
    field.dataset.open = 'true';
    handle = openPopover({
      anchor: field,
      panel: panel.panel,
      container: field,
      align,
      matchWidth: true,
      onDismiss: () => close(false),
    });
    const start = initial === undefined ? initialActive(rows) : edgeRow(rows, initial);
    setActive(start, 'centre');
    updateFades();
    panel.input?.focus();
  }

  /** Closes the list. `refocus` returns focus to the trigger when it was inside the list. */
  function close(refocus: boolean): void {
    if (!isOpen()) return;
    const hadFocus = panel?.panel.contains(document.activeElement) ?? false;
    handle?.close();
    handle = null;
    typeahead.reset();
    trigger.setAttribute('aria-expanded', 'false');
    trigger.removeAttribute('aria-activedescendant');
    delete field.dataset.open;
    panel?.panel.remove();
    panel = null;
    active = -1;
    if (refocus || hadFocus) trigger.focus();
  }

  function typeInto(char: string): void {
    const run = typeahead.push(char);
    const shown = visible.map((index) => rows[index]);
    const from = Math.max(0, visible.indexOf(active));
    const match = typeaheadMatch(shown, run, from);
    if (match !== -1) setActive(visible[match], 'nearest');
  }

  function isPrintable(event: KeyboardEvent): boolean {
    return event.key.length === 1 && !event.ctrlKey && !event.metaKey && !event.altKey;
  }

  function onClosedKey(event: KeyboardEvent): void {
    const { searchable } = config();
    const key = event.key;
    if (['ArrowDown', 'ArrowUp', 'Enter', ' '].includes(key)) {
      event.preventDefault();
      open();
    } else if (key === 'Home' || key === 'End') {
      event.preventDefault();
      open(key === 'Home' ? 'first' : 'last');
    } else if (isPrintable(event)) {
      event.preventDefault();
      open();
      if (searchable && panel?.input != null) {
        panel.input.value = event.key;
        onFilterInput();
      } else {
        typeInto(event.key);
      }
    }
  }

  /** Escape: clear a filter first, then close; never reaches a surrounding dialog (§A3.6). */
  function onEscape(event: KeyboardEvent): void {
    event.preventDefault();
    event.stopPropagation();
    if (panel?.input != null && panel.input.value !== '') {
      panel.input.value = '';
      onFilterInput();
      return;
    }
    close(true);
  }

  function onOpenKey(event: KeyboardEvent): void {
    const { searchable, multiple } = config();
    const key = event.key;
    const moves: Record<string, () => void> = {
      ArrowDown: () => step(1),
      ArrowUp: () => step(-1),
      PageDown: () => step(PAGE_ROWS),
      PageUp: () => step(-PAGE_ROWS),
    };
    if (key === 'Escape') return onEscape(event);
    if (key === 'ArrowUp' && event.altKey) {
      event.preventDefault();
      if (multiple) close(true);
      else commit(active);
      return;
    }
    if (moves[key] !== undefined) {
      event.preventDefault();
      moves[key]();
      return;
    }
    if ((key === 'Home' || key === 'End') && !searchable) {
      event.preventDefault();
      edge(key === 'Home' ? 'first' : 'last');
      return;
    }
    if (key === 'Enter') {
      event.preventDefault();
      commit(active);
      return;
    }
    if (key === 'Tab') {
      // Tab commits only the keyboard-active row, never one a parked pointer highlights (§A6.2).
      if (!multiple && active !== -1) {
        const row = rows[active];
        if (row !== undefined && !row.disabled && select.value !== row.value) {
          select.value = row.value;
          dispatchChange();
        }
      }
      // Focus back on the trigger first, so the browser's own Tab moves on from the field.
      close(true);
      return;
    }
    if (key === ' ' && multiple && event.ctrlKey) {
      event.preventDefault();
      commit(active);
      return;
    }
    if (searchable) return; // Every other key types into the filter.
    if (key === ' ') {
      event.preventDefault();
      commit(active);
      return;
    }
    if (isPrintable(event)) {
      event.preventDefault();
      typeInto(event.key);
    }
  }

  function onKeydown(event: KeyboardEvent): void {
    if (isOpen()) onOpenKey(event);
    else if (event.currentTarget === trigger) onClosedKey(event);
  }

  function onTriggerClick(): void {
    if (isOpen()) close(true);
    else open();
  }

  function onPointerDown(event: PointerEvent): void {
    lastPointer = event.pointerType;
  }
  function onPointerMove(event: PointerEvent): void {
    if (event.pointerType === 'mouse') lastPointer = 'mouse';
  }

  /** Icons of the first rows load before the first open, so the list never flashes (§A11.3). */
  function preloadIcons(): void {
    if (iconsPreloaded) return;
    iconsPreloaded = true;
    [...select.options]
      .map((option) => option.dataset.icon)
      .filter((src): src is string => src !== undefined && src !== '')
      .slice(0, PRELOAD_ICON_COUNT)
      .forEach((src) => {
        new Image().src = src;
      });
  }

  /** A hidden required select puts its error and focus on the trigger (§A3.9). */
  function onInvalid(event: Event): void {
    event.preventDefault();
    field.dataset.invalid = 'true';
    trigger.focus();
  }

  function onLabelClick(event: MouseEvent): void {
    event.preventDefault();
    trigger.focus();
  }

  // The page's own select.focus() (a form putting focus on its first bad field) lands on the
  // trigger, since a hidden select cannot take focus.
  const nativeFocus = Object.getOwnPropertyDescriptor(select, 'focus');
  select.focus = (options?: FocusOptions) => trigger.focus(options);
  const labels = [...(select.labels ?? [])];
  trigger.addEventListener('keydown', onKeydown);
  trigger.addEventListener('click', onTriggerClick);
  field.addEventListener('pointerdown', onPointerDown);
  field.addEventListener('pointermove', onPointerMove);
  field.addEventListener('pointerenter', preloadIcons);
  field.addEventListener('focusin', preloadIcons);
  select.addEventListener('invalid', onInvalid);
  labels.forEach((label) => label.addEventListener('click', onLabelClick));
  field.dataset.enhanced = 'true';
  syncTrigger();

  const instance: Listbox = {
    open: () => open(),
    close: () => close(false),
    refresh: () => {
      syncTrigger();
      reload();
    },
    destroy: () => {
      close(false);
      trigger.remove();
      parts.nameSource?.remove();
      parts.natives.forEach((native) => {
        native.hidden = false;
      });
      if (nativeFocus === undefined) delete (select as { focus?: unknown }).focus;
      else Object.defineProperty(select, 'focus', nativeFocus);
      field.removeEventListener('pointerdown', onPointerDown);
      field.removeEventListener('pointermove', onPointerMove);
      field.removeEventListener('pointerenter', preloadIcons);
      field.removeEventListener('focusin', preloadIcons);
      select.removeEventListener('invalid', onInvalid);
      labels.forEach((label) => label.removeEventListener('click', onLabelClick));
      delete field.dataset.enhanced;
      instances.delete(field);
    },
  };
  return instance;
}

/** Enhances a field once; a second call returns the same instance. */
export function enhance(field: HTMLElement): Listbox {
  const existing = instances.get(field);
  if (existing !== undefined) return existing;
  const instance = create(field);
  instances.set(field, instance);
  return instance;
}

/** The instance on a field, if it is enhanced. */
export function listboxOf(field: HTMLElement): Listbox | undefined {
  return instances.get(field);
}
