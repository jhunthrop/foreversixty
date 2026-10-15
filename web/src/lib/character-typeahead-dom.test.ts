// @vitest-environment jsdom
// web/src/lib/character-typeahead-dom.test.ts
// The typeahead on both lookup boxes (design/specs/2026-10-15-character-typeahead.md): the
// debounce, the abort of a superseded request, the rows, the keys and the quiet failure.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { CharacterSearch, CharacterSuggestion } from './character-search-api';
import { DEBOUNCE_MS, wireTypeahead } from './character-typeahead-dom';

const BOLT: CharacterSuggestion = {
  key: 'us/normal/bolt',
  region: 'us',
  ruleset: 'normal',
  name: 'Bolt',
  class: 'Mage',
  level: 60,
  spec: 'Fire',
};
const BOROS: CharacterSuggestion = {
  key: 'eu/hardcore/boros',
  region: 'eu',
  ruleset: 'hardcore',
  name: 'Boros',
  class: 'Warrior',
  level: 41,
  spec: null,
};
const BO: CharacterSuggestion = {
  key: 'us/pvp/bo',
  region: 'us',
  ruleset: 'pvp',
  name: 'Bo',
  class: null,
  level: null,
  spec: null,
};

function mount(withFilters = false): { form: HTMLFormElement; input: HTMLInputElement } {
  document.body.innerHTML = `
    <form>
      <div data-typeahead>
        <input id="q" name="name" type="search" role="combobox" aria-autocomplete="list"
          aria-expanded="false" aria-controls="q-listbox" />
        <ul id="q-listbox" role="listbox" hidden></ul>
      </div>
      ${
        withFilters
          ? '<select name="region"><option value="eu" selected>EU</option></select><select name="ruleset"><option value="hardcore" selected>Hardcore</option></select>'
          : ''
      }
      <button type="submit">Look up</button>
    </form>`;
  return {
    form: document.querySelector('form') as HTMLFormElement,
    input: document.querySelector('input') as HTMLInputElement,
  };
}

function type(input: HTMLInputElement, text: string): void {
  input.value = text;
  input.dispatchEvent(new Event('input', { bubbles: true }));
}

function press(input: HTMLInputElement, key: string): KeyboardEvent {
  const event = new KeyboardEvent('keydown', { key, bubbles: true, cancelable: true });
  input.dispatchEvent(event);
  return event;
}

const options = (): HTMLAnchorElement[] =>
  Array.from(document.querySelectorAll<HTMLAnchorElement>('a[role="option"]'));
const listbox = (): HTMLElement => document.getElementById('q-listbox') as HTMLElement;

async function settle(): Promise<void> {
  await vi.advanceTimersByTimeAsync(DEBOUNCE_MS);
}

let navigate: ReturnType<typeof vi.fn<(href: string) => void>>;

function wire(search: CharacterSearch, withFilters = false) {
  const parts = mount(withFilters);
  navigate = vi.fn<(href: string) => void>();
  wireTypeahead(parts.form, { search, navigate });
  parts.input.focus();
  return parts;
}

const answering = (rows: CharacterSuggestion[]) => vi.fn<CharacterSearch>(async () => rows);

beforeEach(() => vi.useFakeTimers());
afterEach(() => {
  vi.useRealTimers();
  document.body.innerHTML = '';
});

describe('asking the endpoint', () => {
  it('waits for a pause of 150ms and then asks once, with the box’s filters', async () => {
    const search = answering([BOLT]);
    const { input } = wire(search, true);
    type(input, 'b');
    type(input, 'bo');
    await vi.advanceTimersByTimeAsync(DEBOUNCE_MS - 1);
    expect(search).not.toHaveBeenCalled();
    await vi.advanceTimersByTimeAsync(1);
    expect(search).toHaveBeenCalledTimes(1);
    expect(search.mock.calls[0][0]).toBe('bo');
    expect(search.mock.calls[0][1]).toMatchObject({ region: 'eu', ruleset: 'hardcore' });
  });

  it('searches every realm when the box has no region or ruleset field', async () => {
    const search = answering([BOLT]);
    const { input } = wire(search);
    type(input, 'bo');
    await settle();
    expect(search.mock.calls[0][1].region).toBeUndefined();
    expect(search.mock.calls[0][1].ruleset).toBeUndefined();
  });

  it('does not ask under two characters, and closes and clears the list', async () => {
    const search = answering([BOLT]);
    const { input } = wire(search);
    type(input, 'bo');
    await settle();
    expect(options()).toHaveLength(1);
    type(input, 'b');
    await settle();
    expect(search).toHaveBeenCalledTimes(1);
    expect(listbox().hidden).toBe(true);
    expect(listbox().children).toHaveLength(0);
    expect(input.getAttribute('aria-expanded')).toBe('false');
  });

  it('aborts the superseded request and ignores its answer', async () => {
    const signals: AbortSignal[] = [];
    const resolvers: ((rows: CharacterSuggestion[]) => void)[] = [];
    const search = vi.fn<CharacterSearch>(
      (_text, { signal }) =>
        new Promise((resolve) => {
          signals.push(signal);
          resolvers.push(resolve);
        }),
    );
    const { input } = wire(search);
    type(input, 'bo');
    await settle();
    type(input, 'bol');
    expect(signals[0].aborted).toBe(true);
    await settle();
    resolvers[0]([BOROS]);
    resolvers[1]([BOLT]);
    await vi.advanceTimersByTimeAsync(0);
    expect(options().map((option) => option.textContent)).toEqual([expect.stringContaining('Bolt')]);
  });

  it('leaves the list as it was while a request is in flight', async () => {
    const pending = new Promise<CharacterSuggestion[]>(() => undefined);
    const search = vi.fn<CharacterSearch>().mockResolvedValueOnce([BOLT]).mockReturnValueOnce(pending);
    const { input } = wire(search);
    type(input, 'bo');
    await settle();
    type(input, 'bol');
    await settle();
    expect(options()).toHaveLength(1);
    expect(listbox().hidden).toBe(false);
  });
});

describe('the rows', () => {
  it('renders each as a link to the character with a crest and the realm line', async () => {
    const { input } = wire(answering([BOLT, BOROS, BO]));
    type(input, 'bo');
    await settle();
    const rows = options();
    expect(rows.map((row) => row.getAttribute('href'))).toEqual([
      '/character/us/normal/bolt',
      '/character/eu/hardcore/boros',
      '/character/us/pvp/bo',
    ]);
    expect(rows[0].querySelector('img')?.getAttribute('src')).toBe('/icons/hd/crests/mage.webp');
    expect(rows[0].style.getPropertyValue('--c')).toBe('#3fc7eb');
    expect(rows[0].querySelector('.typeahead-name')?.textContent).toBe('Bolt');
    expect(rows[0].querySelector('.typeahead-detail')?.textContent).toBe('Fire Mage · 60');
    expect(rows[0].querySelector('.typeahead-realm')?.textContent).toBe('US Normal');
    expect(rows[1].querySelector('.typeahead-detail')?.textContent).toBe('Warrior · 41');
    expect(rows[1].querySelector('.typeahead-realm')?.textContent).toBe('EU Hardcore');
  });

  it('gives a character with no class a neutral ring and no detail', async () => {
    const { input } = wire(answering([BO]));
    type(input, 'bo');
    await settle();
    const [row] = options();
    expect(row.querySelector('img')).toBeNull();
    expect(row.querySelector('.typeahead-crest-none')).not.toBeNull();
    expect(row.classList.contains('typeahead-row-plain')).toBe(true);
    expect(row.querySelector('.typeahead-detail')).toBeNull();
    expect(row.querySelector('.typeahead-realm')?.textContent).toBe('US PvP');
  });

  it('opens the list and wires aria-expanded, and gives the list its name', async () => {
    const { input } = wire(answering([BOLT]));
    type(input, 'bo');
    await settle();
    expect(listbox().hidden).toBe(false);
    expect(listbox().getAttribute('aria-label')).toBe('Characters');
    expect(input.getAttribute('aria-expanded')).toBe('true');
  });
});

describe('keys', () => {
  async function open() {
    const parts = wire(answering([BOLT, BOROS, BO]));
    type(parts.input, 'bo');
    await settle();
    return parts;
  }
  const active = (input: HTMLInputElement): string | null => input.getAttribute('aria-activedescendant');

  it('moves the highlight with ArrowDown and ArrowUp, wrapping both ways', async () => {
    const { input } = await open();
    const ids = options().map((row) => row.id);
    expect(press(input, 'ArrowDown').defaultPrevented).toBe(true);
    expect(active(input)).toBe(ids[0]);
    press(input, 'ArrowDown');
    press(input, 'ArrowDown');
    expect(active(input)).toBe(ids[2]);
    press(input, 'ArrowDown');
    expect(active(input)).toBe(ids[0]);
    press(input, 'ArrowUp');
    expect(active(input)).toBe(ids[2]);
    expect(options()[2].getAttribute('aria-selected')).toBe('true');
    expect(options()[0].getAttribute('aria-selected')).toBe('false');
  });

  it('starts at the last row on ArrowUp from nothing', async () => {
    const { input } = await open();
    press(input, 'ArrowUp');
    expect(active(input)).toBe(options()[2].id);
  });

  it('opens the highlighted row’s page on Enter without submitting the form', async () => {
    const { form, input } = await open();
    const submit = vi.fn();
    form.addEventListener('submit', submit);
    press(input, 'ArrowDown');
    press(input, 'ArrowDown');
    const enter = press(input, 'Enter');
    expect(enter.defaultPrevented).toBe(true);
    expect(navigate).toHaveBeenCalledWith('/character/eu/hardcore/boros');
    expect(submit).not.toHaveBeenCalled();
  });

  it('leaves Enter to the form when nothing is highlighted', async () => {
    const { input } = await open();
    const enter = press(input, 'Enter');
    expect(enter.defaultPrevented).toBe(false);
    expect(navigate).not.toHaveBeenCalled();
  });

  it('closes on Escape and on Tab, and ArrowDown opens it again', async () => {
    const { input } = await open();
    press(input, 'ArrowDown');
    expect(press(input, 'Escape').defaultPrevented).toBe(true);
    expect(listbox().hidden).toBe(true);
    expect(active(input)).toBeNull();
    press(input, 'ArrowDown');
    expect(listbox().hidden).toBe(false);
    expect(press(input, 'Tab').defaultPrevented).toBe(false);
    expect(listbox().hidden).toBe(true);
  });

  it('closes on blur', async () => {
    const { input } = await open();
    input.blur();
    expect(listbox().hidden).toBe(true);
  });

  it('highlights the row under the pointer', async () => {
    const { input } = await open();
    options()[1].dispatchEvent(new MouseEvent('mouseover', { bubbles: true }));
    expect(active(input)).toBe(options()[1].id);
  });

  it('keeps the box focused when a row is pressed, so the click still lands', async () => {
    await open();
    const down = new MouseEvent('mousedown', { bubbles: true, cancelable: true });
    options()[0].dispatchEvent(down);
    expect(down.defaultPrevented).toBe(true);
  });
});

describe('when nothing matches or the read fails', () => {
  it('shows one muted, unselectable row and lets Enter submit the typed name', async () => {
    const { input } = wire(answering([]));
    type(input, 'zz');
    await settle();
    expect(options()).toHaveLength(0);
    const row = listbox().querySelector('[aria-disabled="true"]') as HTMLElement;
    expect(row.textContent).toBe('No character named "zz" yet');
    expect(row.getAttribute('role')).toBe('option');
    press(input, 'ArrowDown');
    expect(input.getAttribute('aria-activedescendant')).toBeNull();
    expect(press(input, 'Enter').defaultPrevented).toBe(false);
    expect(navigate).not.toHaveBeenCalled();
  });

  it('closes the list quietly when the read fails', async () => {
    const search = vi.fn<CharacterSearch>().mockResolvedValueOnce([BOLT]).mockRejectedValueOnce(new Error());
    const { input } = wire(search);
    type(input, 'bo');
    await settle();
    type(input, 'bol');
    await settle();
    expect(listbox().hidden).toBe(true);
    expect(input.getAttribute('aria-expanded')).toBe('false');
    expect(press(input, 'Enter').defaultPrevented).toBe(false);
  });
});

describe('wiring', () => {
  it('wires a form only once', async () => {
    const search = answering([BOLT]);
    const { form, input } = wire(search);
    wireTypeahead(form, { search, navigate });
    type(input, 'bo');
    await settle();
    expect(search).toHaveBeenCalledTimes(1);
  });
});
