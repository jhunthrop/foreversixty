// @vitest-environment jsdom
// web/src/lib/ui/listbox.test.ts
// The enhancer on the real Select markup (design/specs/2026-10-30-controls-charts.md §A3, §A6):
// keyboard, typeahead, the one-combobox focus rule, Escape that closes only the list, the
// two highlights, the top layer, positioning and the flip. jsdom has no Popover API or
// layout, so the popover calls are recorded and the boxes are given. The field is built to
// Select.svelte's markup contract (Select.test.ts pins that markup on the server render; this
// environment compiles the component for the client, which svelte/server cannot render).
import { afterEach, beforeAll, beforeEach, describe, expect, it, vi } from 'vitest';
import { enhance, type Listbox } from './listbox';
import { optionAttributes, type SelectOption } from './listbox-model';
import { openOverlayCount } from './overlay';

const PLAYERS: SelectOption[] = [
  { value: 'garrosk', label: 'Garrosk', group: 'Damage', colour: '--color-class-warrior' },
  { value: 'velyn', label: 'Velyn', group: 'Damage', disabled: true },
  { value: 'morrigan', label: 'Morrigan', group: 'Damage' },
  { value: 'ysolde', label: 'Ysolde', group: 'Damage', description: 'Fire mage' },
  { value: 'gorehowl', label: 'Gorehowl', group: 'Damage', meta: '×8' },
  { value: 'lightningbolltz', label: 'Lightningbolltz', group: 'Damage', you: true },
];

const shown = new Set<Element>();
let instances: Listbox[] = [];

beforeAll(() => {
  HTMLElement.prototype.showPopover = function showPopover(this: HTMLElement) {
    shown.add(this);
  };
  HTMLElement.prototype.hidePopover = function hidePopover(this: HTMLElement) {
    shown.delete(this);
  };
  class Observer {
    observe(): void {}
    unobserve(): void {}
    disconnect(): void {}
  }
  vi.stubGlobal('ResizeObserver', Observer);
  vi.stubGlobal('IntersectionObserver', Observer);
  Element.prototype.scrollIntoView = () => {};
  // jsdom lays nothing out: give the viewport its size so Floating UI's clipping rect is real.
  Object.defineProperty(document.documentElement, 'clientWidth', {
    configurable: true,
    get: () => window.innerWidth,
  });
  Object.defineProperty(document.documentElement, 'clientHeight', {
    configurable: true,
    get: () => window.innerHeight,
  });
  Object.defineProperty(HTMLElement.prototype, 'scrollHeight', { configurable: true, get: () => 300 });
});

beforeEach(() => {
  document.body.innerHTML = '';
  shown.clear();
});

afterEach(() => {
  instances.forEach((instance) => instance.destroy());
  instances = [];
});

interface FieldProps {
  id?: string;
  value?: string;
  values?: string[];
  size?: string;
  align?: string;
  searchable?: boolean;
  searchPlaceholder?: string;
  multiple?: boolean;
  fieldLabel?: string;
  foot?: string;
  footCount?: string;
  placeholder?: string;
  required?: boolean;
}

/** Select.svelte's markup: the wrapper's data, the native select with option data, the face. */
function fieldMarkup(props: FieldProps, id: string): HTMLElement {
  const field = document.createElement('div');
  field.className = 'sel';
  field.dataset.listbox = '';
  field.dataset.size = props.size ?? 'dense';
  field.dataset.align = props.align ?? 'start';
  if (props.searchable || props.multiple) field.dataset.searchable = 'true';
  if (props.searchPlaceholder !== undefined) field.dataset.searchPlaceholder = props.searchPlaceholder;
  if (props.multiple) field.dataset.multiple = 'true';
  if (props.placeholder !== undefined) field.dataset.placeholder = 'true';
  if (props.foot !== undefined) field.dataset.foot = props.foot;
  if (props.footCount !== undefined) field.dataset.footCount = props.footCount;
  const select = document.createElement('select');
  select.className = 'sel-native';
  select.id = id;
  select.multiple = props.multiple ?? false;
  select.required = props.required ?? false;
  if (props.placeholder !== undefined) select.append(new Option(props.placeholder, ''));
  let group: HTMLOptGroupElement | null = null;
  for (const option of PLAYERS) {
    const node = new Option(option.label, option.value);
    node.disabled = option.disabled ?? false;
    for (const [name, value] of Object.entries(optionAttributes(option))) {
      if (value !== undefined) node.setAttribute(name, value);
    }
    if (option.group !== undefined && group?.label !== option.group) {
      group = document.createElement('optgroup');
      group.label = option.group;
      select.append(group);
    }
    (group ?? select).append(node);
  }
  const values = props.multiple ? (props.values ?? []) : [props.value ?? 'morrigan'];
  for (const option of select.options) option.selected = values.includes(option.value);
  const face = document.createElement('span');
  face.className = 'sel-face';
  const label = document.createElement('span');
  label.className = 'sel-label';
  label.textContent = props.multiple
    ? (props.fieldLabel ?? '')
    : (select.selectedOptions[0]?.textContent ?? '');
  face.append(label);
  field.append(select, face);
  return field;
}

function mount(props: FieldProps = {}): {
  field: HTMLElement;
  select: HTMLSelectElement;
  trigger: HTMLButtonElement;
  listbox: Listbox;
} {
  const id = props.id ?? 'source';
  const host = document.createElement('div');
  host.innerHTML = `<label for="${id}">Source</label>`;
  const field = fieldMarkup(props, id);
  host.append(field);
  document.body.append(host);
  const listbox = enhance(field);
  instances.push(listbox);
  return {
    field,
    select: field.querySelector('select')!,
    trigger: field.querySelector<HTMLButtonElement>('.sel-trigger')!,
    listbox,
  };
}

function key(target: Element, name: string, init: KeyboardEventInit = {}): KeyboardEvent {
  const event = new KeyboardEvent('keydown', { key: name, bubbles: true, cancelable: true, ...init });
  target.dispatchEvent(event);
  return event;
}

const activeLabel = (holder: Element): string | undefined => {
  const id = holder.getAttribute('aria-activedescendant');
  return id === null
    ? undefined
    : (document.getElementById(id)?.querySelector('.sel-row-label')?.textContent ?? '');
};

const panelOf = (field: HTMLElement): HTMLElement | null => field.querySelector('.sel-pop');

describe('enhancing a field', () => {
  it('hides the native select (still in the form) and adds a select-only combobox', () => {
    const { select, trigger } = mount();
    expect(select.hidden).toBe(true);
    expect(trigger.getAttribute('role')).toBe('combobox');
    expect(trigger.getAttribute('aria-haspopup')).toBe('listbox');
    expect(trigger.getAttribute('aria-expanded')).toBe('false');
    expect(trigger.getAttribute('aria-labelledby')).toBe(`${select.labels[0].id} ${trigger.id}`);
    expect(trigger.textContent).toBe('Morrigan');
  });

  it('destroy puts the native select back', () => {
    const { field, select, listbox } = mount();
    listbox.destroy();
    instances = [];
    expect(select.hidden).toBe(false);
    expect(field.querySelector('.sel-trigger')).toBeNull();
  });
});

describe('the top layer', () => {
  it('opens a popover="manual" list that is a child of the field, shown with showPopover', () => {
    const { field, trigger } = mount();
    key(trigger, 'ArrowDown');
    const panel = panelOf(field)!;
    expect(panel.getAttribute('popover')).toBe('manual');
    expect(field.contains(panel)).toBe(true);
    expect(shown.has(panel)).toBe(true);
    expect(trigger.getAttribute('aria-expanded')).toBe('true');
    expect(openOverlayCount()).toBe(1);
  });

  it('one open overlay at a time: opening a second list closes the first', () => {
    const first = mount({ id: 'a' });
    const second = mount({ id: 'b' });
    key(first.trigger, 'ArrowDown');
    key(second.trigger, 'ArrowDown');
    expect(panelOf(first.field)).toBeNull();
    expect(first.trigger.getAttribute('aria-expanded')).toBe('false');
    expect(openOverlayCount()).toBe(1);
  });

  it('a click outside closes with no change', () => {
    const { field, select } = mount();
    field.querySelector<HTMLButtonElement>('.sel-trigger')!.click();
    document.body.dispatchEvent(new Event('pointerdown', { bubbles: true }));
    expect(panelOf(field)).toBeNull();
    expect(select.value).toBe('morrigan');
  });
});

describe('keyboard on a plain list (focus stays on the trigger)', () => {
  it('opens on the picked row and moves by the arrows, skipping disabled rows, no wrap', () => {
    const { trigger } = mount();
    key(trigger, 'ArrowDown');
    expect(activeLabel(trigger)).toBe('Morrigan');
    key(trigger, 'ArrowUp');
    expect(activeLabel(trigger)).toBe('Garrosk');
    key(trigger, 'ArrowUp');
    expect(activeLabel(trigger)).toBe('Garrosk');
    key(trigger, 'End');
    expect(activeLabel(trigger)).toBe('Lightningbolltz');
    key(trigger, 'ArrowDown');
    expect(activeLabel(trigger)).toBe('Lightningbolltz');
    key(trigger, 'Home');
    expect(activeLabel(trigger)).toBe('Garrosk');
  });

  it('typeahead lands on the first row starting with the typed run', () => {
    const { trigger } = mount();
    key(trigger, 'ArrowDown');
    key(trigger, 'y');
    expect(activeLabel(trigger)).toBe('Ysolde');
  });

  it('a printable key on the closed field opens it and runs typeahead', () => {
    const { trigger } = mount();
    key(trigger, 'l');
    expect(trigger.getAttribute('aria-expanded')).toBe('true');
    expect(activeLabel(trigger)).toBe('Lightningbolltz');
  });

  it('Enter commits the active row once, closes, and keeps focus on the trigger', () => {
    const { select, trigger } = mount();
    const changes = vi.fn();
    select.addEventListener('change', changes);
    trigger.focus();
    key(trigger, 'ArrowDown');
    key(trigger, 'ArrowDown');
    key(trigger, 'Enter');
    expect(select.value).toBe('ysolde');
    expect(changes).toHaveBeenCalledTimes(1);
    expect(trigger.getAttribute('aria-expanded')).toBe('false');
    expect(document.activeElement).toBe(trigger);
  });

  it('committing the already-picked row fires nothing', () => {
    const { select, trigger } = mount();
    const changes = vi.fn();
    select.addEventListener('change', changes);
    key(trigger, 'ArrowDown');
    key(trigger, 'Enter');
    expect(changes).not.toHaveBeenCalled();
  });

  it('Escape closes only the list: it never reaches a surrounding dialog', () => {
    const { select, trigger } = mount();
    const dialogEscape = vi.fn();
    const onDocumentKey = (event: KeyboardEvent): void => {
      if (event.key === 'Escape') dialogEscape();
    };
    document.addEventListener('keydown', onDocumentKey);
    key(trigger, 'ArrowDown');
    key(trigger, 'ArrowDown');
    const event = key(trigger, 'Escape');
    document.removeEventListener('keydown', onDocumentKey);
    expect(event.defaultPrevented).toBe(true);
    expect(dialogEscape).not.toHaveBeenCalled();
    expect(trigger.getAttribute('aria-expanded')).toBe('false');
    expect(select.value).toBe('morrigan');
  });
});

describe('two highlights, one commit rule', () => {
  it('a pointer resting over a row does not move the keyboard-active row; Tab commits that one', () => {
    const { field, select, trigger } = mount();
    key(trigger, 'ArrowDown');
    const ysolde = [...field.querySelectorAll('[role="option"]')].find((node) =>
      node.textContent?.includes('Ysolde'),
    )!;
    ysolde.dispatchEvent(new Event('pointerover', { bubbles: true }));
    key(trigger, 'Tab');
    expect(select.value).toBe('morrigan');
  });

  it('a moving pointer carries the keyboard-active row with it', () => {
    const { field, select, trigger } = mount();
    key(trigger, 'ArrowDown');
    const ysolde = [...field.querySelectorAll('[role="option"]')].find((node) =>
      node.textContent?.includes('Ysolde'),
    )!;
    ysolde.dispatchEvent(new PointerEvent('pointermove', { bubbles: true, pointerType: 'mouse' }));
    expect(activeLabel(trigger)).toBe('Ysolde');
    key(trigger, 'Tab');
    expect(select.value).toBe('ysolde');
  });

  it('a click commits the clicked row', () => {
    const { field, select, trigger } = mount();
    trigger.click();
    const row = [...field.querySelectorAll<HTMLElement>('[role="option"]')].find((node) =>
      node.textContent?.includes('Gorehowl'),
    )!;
    row.click();
    expect(select.value).toBe('gorehowl');
    expect(panelOf(field)).toBeNull();
  });
});

describe('rows', () => {
  it('carry the gold check, the You tag, meta, and the description as aria-describedby', () => {
    const { field, trigger } = mount();
    key(trigger, 'ArrowDown');
    const rows = [...field.querySelectorAll<HTMLElement>('[role="option"]')];
    const morrigan = rows.find((node) => node.textContent?.includes('Morrigan'))!;
    expect(morrigan.getAttribute('aria-selected')).toBe('true');
    expect(morrigan.querySelector('.sel-check')).not.toBeNull();
    const ysolde = rows.find((node) => node.textContent?.includes('Ysolde'))!;
    const describedBy = ysolde.getAttribute('aria-describedby')!;
    expect(document.getElementById(describedBy)?.textContent).toBe('Fire mage');
    expect(ysolde.getAttribute('aria-labelledby')).not.toContain(describedBy);
    expect(
      rows.find((node) => node.textContent?.includes('Lightningbolltz'))?.querySelector('.sel-you'),
    ).not.toBeNull();
    expect(
      rows.find((node) => node.textContent?.includes('Gorehowl'))?.querySelector('.sel-row-meta')
        ?.textContent,
    ).toBe('×8');
    expect(rows.find((node) => node.textContent?.includes('Velyn'))?.getAttribute('aria-disabled')).toBe(
      'true',
    );
    expect(field.querySelector('[role="group"] .sel-grp-head')?.textContent).toBe('Damage');
    expect(field.querySelector('.sel-pop')?.getAttribute('data-dense')).toBe('true');
  });
});

describe('a searchable list (focus moves into the filter)', () => {
  const searchable = { searchable: true, searchPlaceholder: 'Filter players' };

  it('the filter input is the one combobox; the trigger is a plain button', () => {
    const { field, trigger } = mount(searchable);
    trigger.focus();
    key(trigger, 'ArrowDown');
    const input = field.querySelector<HTMLInputElement>('.sel-filter-input')!;
    expect(document.activeElement).toBe(input);
    expect(trigger.hasAttribute('role')).toBe(false);
    expect(input.getAttribute('role')).toBe('combobox');
    expect(document.querySelectorAll('[role="combobox"]')).toHaveLength(1);
    expect(activeLabel(input)).toBe('Morrigan');
    expect(trigger.hasAttribute('aria-activedescendant')).toBe(false);
  });

  it('filters, counts and marks the match; no match shows the no-match row', () => {
    const { field, trigger } = mount(searchable);
    key(trigger, 'ArrowDown');
    const input = field.querySelector<HTMLInputElement>('.sel-filter-input')!;
    input.value = 'go';
    input.dispatchEvent(new Event('input'));
    expect(field.querySelectorAll('[role="option"]')).toHaveLength(1);
    expect(field.querySelector('.sel-count')?.textContent).toBe('1 match');
    expect(field.querySelector('.sel-match')?.textContent).toBe('Go');
    input.value = 'naxx';
    input.dispatchEvent(new Event('input'));
    expect(field.querySelector('.sel-empty')?.textContent).toBe('Nothing matches “naxx”.');
  });

  it('Escape clears the filter first, then closes and returns focus to the trigger', () => {
    const { field, trigger } = mount(searchable);
    key(trigger, 'ArrowDown');
    const input = field.querySelector<HTMLInputElement>('.sel-filter-input')!;
    input.value = 'go';
    input.dispatchEvent(new Event('input'));
    key(input, 'Escape');
    expect(input.value).toBe('');
    expect(trigger.getAttribute('aria-expanded')).toBe('true');
    key(input, 'Escape');
    expect(trigger.getAttribute('aria-expanded')).toBe('false');
    expect(document.activeElement).toBe(trigger);
  });

  it('Space types into the filter rather than committing', () => {
    const { field, select, trigger } = mount(searchable);
    key(trigger, 'ArrowDown');
    const event = key(field.querySelector('.sel-filter-input')!, ' ');
    expect(event.defaultPrevented).toBe(false);
    expect(select.value).toBe('morrigan');
  });
});

describe('the multi-select', () => {
  const multi = {
    multiple: true,
    values: ['garrosk'],
    fieldLabel: 'Top 8 of 6',
    foot: JSON.stringify([
      { id: 'top', label: 'Top 8' },
      { id: 'custom', label: 'Custom', pressed: true, inert: true },
    ]),
    footCount: '1 drawn',
  };

  it('toggles with Enter and stays open; each toggle fires change', () => {
    const { field, select, trigger } = mount(multi);
    const changes = vi.fn();
    select.addEventListener('change', changes);
    key(trigger, 'ArrowDown');
    const input = field.querySelector('.sel-filter-input')!;
    expect(field.querySelector('[role="listbox"]')?.getAttribute('aria-multiselectable')).toBe('true');
    key(input, 'Enter');
    expect([...select.selectedOptions].map((o) => o.value)).toEqual([]);
    key(input, 'ArrowDown');
    key(input, 'Enter');
    expect([...select.selectedOptions].map((o) => o.value)).toEqual(['morrigan']);
    expect(changes).toHaveBeenCalledTimes(2);
    expect(trigger.getAttribute('aria-expanded')).toBe('true');
  });

  it('a foot preset dispatches the foot event with its id; the inert tag lights and never fires', () => {
    const { field, trigger } = mount(multi);
    const foot = vi.fn();
    field.addEventListener('selectfoot', foot);
    key(trigger, 'ArrowDown');
    field.querySelector<HTMLElement>('[data-foot-id="custom"]')!.click();
    expect(foot).not.toHaveBeenCalled();
    expect(field.querySelector('[data-foot-id="custom"]')?.getAttribute('data-lit')).toBe('true');
    field.querySelector<HTMLButtonElement>('button[data-foot-id="top"]')!.click();
    expect(foot).toHaveBeenCalledTimes(1);
    expect((foot.mock.calls[0][0] as CustomEvent).detail).toEqual({ id: 'top' });
    expect(field.querySelector('.sel-foot-count')?.textContent).toBe('1 drawn');
  });
});

describe('positioning', () => {
  function placeField(field: HTMLElement, top: number): void {
    field.getBoundingClientRect = () =>
      ({
        x: 100,
        y: top,
        top,
        left: 100,
        right: 300,
        bottom: top + 36,
        width: 200,
        height: 36,
        toJSON: () => ({}),
      }) as DOMRect;
  }

  it('opens below, 4 px from the field', async () => {
    const { field, trigger } = mount();
    placeField(field, 100);
    key(trigger, 'ArrowDown');
    await vi.waitFor(() => expect(panelOf(field)?.dataset.side).toBe('bottom'));
    expect(panelOf(field)?.style.top).toBe('140px');
    expect(panelOf(field)?.style.minWidth).toBe('200px');
  });

  it('flips above when less than 240 px fits below and more fits above', async () => {
    const { field, trigger } = mount();
    placeField(field, window.innerHeight - 80);
    key(trigger, 'ArrowDown');
    await vi.waitFor(() => expect(panelOf(field)?.dataset.side).toBe('top'));
    // The whole panel takes the room above less 16 px; the rows cap at 360 px in CSS.
    expect(panelOf(field)?.style.getPropertyValue('--popover-room')).toBe(
      `${window.innerHeight - 80 - 4 - 16}px`,
    );
  });

  it('a right-slot field opens bottom-end', async () => {
    const { field, trigger } = mount({ align: 'end' });
    placeField(field, 100);
    key(trigger, 'ArrowDown');
    await vi.waitFor(() => expect(panelOf(field)?.dataset.side).toBe('bottom'));
    // jsdom lays the panel out 0 px wide, so bottom-end puts its left edge at the field's right.
    expect(panelOf(field)?.style.left).toBe('300px');
  });
});

describe('validation', () => {
  it('a hidden invalid select puts the error state and focus on the trigger', () => {
    const { field, select, trigger } = mount({ required: true, placeholder: 'Choose', value: '' });
    const event = new Event('invalid', { cancelable: true });
    select.dispatchEvent(event);
    expect(event.defaultPrevented).toBe(true);
    expect(field.dataset.invalid).toBe('true');
    expect(document.activeElement).toBe(trigger);
  });
});
