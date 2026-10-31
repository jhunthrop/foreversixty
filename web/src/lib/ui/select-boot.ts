// web/src/lib/ui/select-boot.ts
// The small, eager half of the shared dropdown (design/specs/2026-10-30-controls-charts.md
// §A3.4, §A9): it decides whether a field gets the listbox (a fine pointer that hovers, and
// the Popover API), and loads listbox.ts with Floating UI as one lazy chunk after first
// paint. Until then, and always on a coarse pointer, the native <select> under the styled
// face is the control. A hybrid laptop re-runs the test when the pointer changes; an open
// list closes as it flips back to native.
import { FINE_POINTER_QUERY, popoverSupported } from './listbox-model';
import type { Listbox } from './listbox';

type ListboxModule = typeof import('./listbox');

let loading: Promise<ListboxModule> | null = null;
const attached = new Set<HTMLElement>();
let pointerQuery: MediaQueryList | null = null;

function loadListbox(): Promise<ListboxModule> {
  loading ??= import('./listbox');
  return loading;
}

function afterIdle(run: () => void): void {
  if (typeof requestIdleCallback === 'function') requestIdleCallback(run, { timeout: 1500 });
  else setTimeout(run, 1);
}

/** True when this device should get the listbox rather than the native picker (D3). */
export function wantsListbox(): boolean {
  return (
    typeof window !== 'undefined' &&
    typeof window.matchMedia === 'function' &&
    window.matchMedia(FINE_POINTER_QUERY).matches &&
    popoverSupported()
  );
}

async function apply(field: HTMLElement): Promise<Listbox | undefined> {
  const module = await loadListbox();
  if (!attached.has(field) || !field.isConnected) return undefined;
  if (wantsListbox()) return module.enhance(field);
  module.listboxOf(field)?.destroy();
  return undefined;
}

function watchPointer(): void {
  if (pointerQuery !== null || typeof window.matchMedia !== 'function') return;
  pointerQuery = window.matchMedia(FINE_POINTER_QUERY);
  pointerQuery.addEventListener('change', () => {
    for (const field of attached) void apply(field);
  });
}

export interface AttachedField {
  /** Re-reads the select after the page changed its options, value or state. */
  refresh: () => void;
  /** Detaches: the native control comes back. */
  destroy: () => void;
}

/** Attaches the listbox to a field after first paint, on a fine pointer only. */
export function attachListbox(field: HTMLElement): AttachedField {
  attached.add(field);
  watchPointer();
  if (wantsListbox()) afterIdle(() => void apply(field));
  return {
    refresh: () => {
      if (loading === null) return;
      void loading.then((module) => module.listboxOf(field)?.refresh());
    },
    destroy: () => {
      attached.delete(field);
      if (loading === null) return;
      void loading.then((module) => module.listboxOf(field)?.destroy());
    },
  };
}

/** A static (server-rendered, never hydrated) field's face follows its select's value. */
export function paintFace(field: HTMLElement): void {
  const select = field.querySelector('select');
  const label = field.querySelector<HTMLElement>('.sel-label');
  if (select === null || label === null) return;
  const option = select.selectedOptions[0];
  const placeholder = field.dataset.placeholder === 'true' && (option === undefined || option.value === '');
  label.textContent = option?.textContent?.trim() ?? '';
  label.classList.toggle('is-placeholder', placeholder);
  const colour = option?.dataset.colour;
  label.style.color = colour !== undefined && /^--[a-z0-9-]+$/.test(colour) ? `var(${colour})` : '';
  const icon = field.querySelector<HTMLImageElement>('.sel-face .sel-icon img');
  if (icon !== null && option?.dataset.icon !== undefined && option.dataset.icon !== '')
    icon.src = option.dataset.icon;
}

/** Boots every static field under `root` (Select.astro's script). */
export function bootStaticFields(root: ParentNode = document): void {
  for (const field of root.querySelectorAll<HTMLElement>('[data-listbox][data-static]')) {
    const handle = attachListbox(field);
    field.addEventListener('change', () => {
      paintFace(field);
      handle.refresh();
    });
  }
}
