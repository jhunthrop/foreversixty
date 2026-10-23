// web/src/lib/header-wiring.ts
// The header's behaviour scripts: the compact lookup and the "/" key, the four doors' hover
// panels and the Guild panel following the session's guild. The header is complete markup
// without them (the Menu is plain inline script, the doors are links), so Header.astro
// loads this module after first paint -- see its loader -- instead of putting this import
// chain ahead of the page's largest contentful paint on every page (W.16).
import { isTypingTarget, wireLookupForm } from './character-search-dom';
import { SESSION_GUILD_EVENT } from './character-selector/config';
import type { SessionGuild } from './character-selector/session-guild';
import { readPublishedSessionGuild } from './character-selector/session-guild-dom';
import { applySessionGuild, wireNavPanels } from './nav-panels-dom';

/** Marks the header once its panels and lookup are wired; specs wait for it, the way they
 *  wait for the tooltip host. */
export const HEADER_WIRED_ATTRIBUTE = 'data-header-wired';

export function wireHeader(): void {
  // The compact lookup and the "/" key (HeaderSearch.astro). "/" outside a text field
  // focuses the home page's own box where there is one; otherwise it opens the bar's panel
  // (from 1024px) or the Menu (below), cursor in the box.
  const header = document.querySelector<HTMLElement>('[data-site-header]');
  // The four doors' hover panels, and the Guild panel following the session's guild: the
  // stamp covers a selector that settled before this ran, the event covers the rest.
  if (header !== null) {
    wireNavPanels(header);
    applySessionGuild(header, readPublishedSessionGuild(header));
    window.addEventListener(SESSION_GUILD_EVENT, (event) =>
      applySessionGuild(header, (event as CustomEvent<SessionGuild | null>).detail),
    );
  }
  for (const form of header?.querySelectorAll<HTMLFormElement>('form[data-character-search]') ?? []) {
    wireLookupForm(form);
  }
  const toggle = header?.querySelector<HTMLButtonElement>('[data-header-search-toggle]') ?? null;
  const panel = header?.querySelector<HTMLElement>('[data-header-search-panel]') ?? null;
  const setPanel = (open: boolean): void => {
    if (toggle === null || panel === null) return;
    panel.hidden = !open;
    toggle.setAttribute('aria-expanded', String(open));
    if (open) panel.querySelector<HTMLInputElement>('input[name="name"]')?.focus();
  };
  toggle?.addEventListener('click', () => setPanel(panel !== null && panel.hidden === true));
  document.addEventListener('click', (event) => {
    if (panel === null || panel.hidden) return;
    const target = event.target;
    if (target instanceof Node && (panel.contains(target) || toggle?.contains(target))) return;
    setPanel(false);
  });
  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape' && panel !== null && !panel.hidden) {
      setPanel(false);
      toggle?.focus();
      return;
    }
    if (event.key !== '/' || event.ctrlKey || event.metaKey || event.altKey) return;
    if (isTypingTarget(event.target)) return;
    const primary = document.querySelector<HTMLInputElement>(
      'form[data-character-search-primary] input[name="name"]',
    );
    event.preventDefault();
    if (primary !== null && primary.offsetParent !== null) {
      primary.focus();
      return;
    }
    if (window.matchMedia('(min-width: 1024px)').matches) {
      setPanel(true);
      return;
    }
    const menuButton = header?.querySelector<HTMLButtonElement>('[data-menu-button]');
    if (menuButton !== null && menuButton !== undefined && !header?.hasAttribute('data-menu-open'))
      menuButton.click();
    header?.querySelector<HTMLInputElement>('[data-header-search-menu] input[name="name"]')?.focus();
  });
  header?.setAttribute(HEADER_WIRED_ATTRIBUTE, '');
}
