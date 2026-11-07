// web/src/lib/ui/control-bar.ts
// The ControlBar's behaviour (web layout spec 2026-11-05 §3.4): the URL is the truth. Every choice
// reads from and writes to the query string with history.replaceState, so a reload, a share and
// the back button restore it. A segmented group is a radio group: arrow keys move and select,
// Home and End jump, Tab leaves. `bootControlBars` wires the server-rendered bar on a static page
// (ControlBar.astro); an island passes `onchange` to ControlBar.svelte and calls the same pure
// helpers. Each change is announced as a bubbling `controlbar:change` event on the bar.

/** The event a static page listens for: which parameter changed and its new value. */
export const CONTROL_BAR_EVENT = 'controlbar:change';

export interface ControlBarChange {
  param: string;
  value: string;
}

/** The query string with `param` set to `value`, or removed when `value` is its default. */
export function withParam(search: string, param: string, value: string, defaultValue?: string): string {
  const params = new URLSearchParams(search);
  if (value === '' || value === defaultValue) params.delete(param);
  else params.set(param, value);
  const text = params.toString();
  return text === '' ? '' : `?${text}`;
}

/** The value a group reads on load: the query's, when it names one of the group's options. */
export function paramValue(
  search: string,
  param: string,
  allowed: readonly string[],
  fallback: string,
): string {
  const value = new URLSearchParams(search).get(param);
  return value !== null && allowed.includes(value) ? value : fallback;
}

/** The radio a key moves to among `count` enabled radios from `index`, or null for other keys. */
export function radioKeyTarget(key: string, index: number, count: number): number | null {
  if (count === 0) return null;
  switch (key) {
    case 'ArrowRight':
    case 'ArrowDown':
      return (index + 1) % count;
    case 'ArrowLeft':
    case 'ArrowUp':
      return (index - 1 + count) % count;
    case 'Home':
      return 0;
    case 'End':
      return count - 1;
    default:
      return null;
  }
}

const RADIO = '[role="radio"]:not([aria-disabled="true"])';

function enabledRadios(group: Element): HTMLElement[] {
  return [...group.querySelectorAll<HTMLElement>(RADIO)];
}

/** Marks `radio` checked in its group (roving tabindex) and returns the change it makes. */
export function checkRadio(radio: HTMLElement): ControlBarChange | null {
  const group = radio.closest<HTMLElement>('[role="radiogroup"][data-param]');
  const param = group?.dataset.param;
  const value = radio.dataset.value;
  if (!group || param === undefined || value === undefined) return null;
  for (const other of group.querySelectorAll<HTMLElement>('[role="radio"]')) {
    const on = other === radio;
    other.setAttribute('aria-checked', String(on));
    other.tabIndex = on ? 0 : -1;
  }
  return { param, value };
}

function applyChange(bar: HTMLElement, change: ControlBarChange, defaultValue?: string): void {
  const search = withParam(window.location.search, change.param, change.value, defaultValue);
  window.history.replaceState(
    window.history.state,
    '',
    `${window.location.pathname}${search}${window.location.hash}`,
  );
  bar.dispatchEvent(new CustomEvent<ControlBarChange>(CONTROL_BAR_EVENT, { detail: change, bubbles: true }));
  closeControlBar(bar);
}

/** Collapses the phone bar (a choice was made): the toggle says so to assistive tech. */
export function closeControlBar(bar: HTMLElement): void {
  bar.removeAttribute('data-open');
  bar.querySelector('[data-control-bar-toggle]')?.setAttribute('aria-expanded', 'false');
}

function onRadioActivate(bar: HTMLElement, radio: HTMLElement, event: Event): void {
  if (radio.getAttribute('aria-disabled') === 'true') {
    event.preventDefault();
    return;
  }
  if (radio.dataset.navigate === 'true') return; // a path option (the spec) is a real link
  event.preventDefault();
  const group = radio.closest<HTMLElement>('[role="radiogroup"]');
  const change = checkRadio(radio);
  if (change) applyChange(bar, change, group?.dataset.default);
}

function onRadioKey(bar: HTMLElement, event: KeyboardEvent): void {
  const radio = (event.target as Element | null)?.closest<HTMLElement>('[role="radio"]');
  const group = radio?.closest('[role="radiogroup"]');
  if (!radio || !group) return;
  const radios = enabledRadios(group);
  const target = radioKeyTarget(event.key, radios.indexOf(radio), radios.length);
  if (target === null) return;
  event.preventDefault();
  radios[target].focus();
  radios[target].click();
}

/** Wires one server-rendered bar: radios, selects, and the phone toggle. Idempotent. */
export function bootControlBar(bar: HTMLElement): void {
  if (bar.dataset.booted === 'true') return;
  bar.dataset.booted = 'true';
  bar.addEventListener('click', (event) => {
    const target = event.target as Element | null;
    const toggle = target?.closest<HTMLElement>('[data-control-bar-toggle]');
    if (toggle) {
      const open = bar.toggleAttribute('data-open');
      toggle.setAttribute('aria-expanded', String(open));
      return;
    }
    const radio = target?.closest<HTMLElement>('[role="radio"]');
    if (radio) onRadioActivate(bar, radio, event);
  });
  bar.addEventListener('keydown', (event) => onRadioKey(bar, event));
  bar.addEventListener('change', (event) => {
    const select = event.target as HTMLSelectElement | null;
    const param = select?.closest<HTMLElement>('[data-param]')?.dataset.param;
    if (!select || select.tagName !== 'SELECT' || param === undefined) return;
    const holder = select.closest<HTMLElement>('[data-param]');
    applyChange(bar, { param, value: select.value }, holder?.dataset.default);
  });
}

/** Boots every bar under `root`. */
export function bootControlBars(root: ParentNode = document): void {
  for (const bar of root.querySelectorAll<HTMLElement>('[data-control-bar]')) bootControlBar(bar);
}
