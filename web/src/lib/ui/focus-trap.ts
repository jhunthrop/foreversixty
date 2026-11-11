// web/src/lib/ui/focus-trap.ts
// The one focus trap (component audit 2026-11-10 cluster 14): Tab and Shift+Tab wrap inside a
// dialog, a sheet or the character selector, so focus never escapes to the page underneath. Every
// modal surface calls `trapTab` from its keydown handler; FOCUSABLE is the one selector
// (state-guard.test.ts fails on a second).

export const FOCUSABLE =
  'a[href],button:not([disabled]):not([tabindex="-1"]),input:not([disabled]),select:not([disabled]),' +
  'textarea:not([disabled]),[tabindex]:not([tabindex="-1"])';

/** The panel's tabbable controls, in document order. */
export function focusables(panel: HTMLElement): HTMLElement[] {
  return [...panel.querySelectorAll<HTMLElement>(FOCUSABLE)];
}

/**
 * Wraps Tab inside `panel`: from the last control to the first, and Shift+Tab from the first (or
 * from `start`, a heading that takes focus on open) to the last. Returns true when it moved focus.
 */
export function trapTab(event: KeyboardEvent, panel: HTMLElement, start?: HTMLElement | null): boolean {
  if (event.key !== 'Tab') return false;
  const items = focusables(panel);
  const first = items[0];
  const last = items[items.length - 1];
  if (first === undefined || last === undefined) return false;
  const active = document.activeElement;
  if (event.shiftKey && (active === first || (start != null && active === start))) {
    event.preventDefault();
    last.focus();
    return true;
  }
  if (!event.shiftKey && active === last) {
    event.preventDefault();
    first.focus();
    return true;
  }
  return false;
}
