// web/src/lib/ui/toc-jump.ts
// "On this page" jumps to a section (guides spec 2026-11-08 §4.3, pages spec 2026-11-08 §4.1):
// the bar's select below 1760 and the open phone bar's rows. The phone bar closes, so the section
// lands below it, the address takes its hash, and focus moves to the section as a link's would.
import { closeControlBar } from './control-bar';

/** Scrolls to the section `id` and focuses it, closing the bar `from` sits in. */
export function jumpToSection(from: HTMLElement, id: string): void {
  const section = id === '' ? null : document.getElementById(id);
  if (section === null) return;
  const bar = from.closest<HTMLElement>('[data-control-bar]');
  if (bar !== null) closeControlBar(bar);
  history.replaceState(history.state, '', `#${section.id}`);
  section.scrollIntoView();
  if (!section.hasAttribute('tabindex')) section.setAttribute('tabindex', '-1');
  section.focus({ preventScroll: true });
}

/** Wires a contents select (by id) and a phone list of `#anchor` links (by selector). */
export function wireTocJump(selectId: string, phoneSelector: string): void {
  const field = document.getElementById(selectId) as HTMLSelectElement | null;
  field?.addEventListener('change', () => jumpToSection(field, field.value));
  document.querySelector(phoneSelector)?.addEventListener('click', (event) => {
    const link = (event.target as HTMLElement).closest<HTMLAnchorElement>('a[href^="#"]');
    if (link === null) return;
    event.preventDefault();
    jumpToSection(link, link.hash.slice(1));
  });
}
