// web/src/lib/nav-panels-dom.ts
// The header's hover panels (spec 2026-10-15-grouped-header.md section 4), as plain DOM wiring
// over the markup Header.astro renders: hover with open and close delays, the chevron toggle,
// a first tap that opens on touch, and the keyboard walk. No island and no framework; the
// header's processed script calls wireNavPanels once. Only one panel is ever open.
import type { SessionGuild } from './character-selector/session-guild';
import { GUILD_GROUP_INDEX, GUILD_OVERVIEW_TAB, NAV_GROUPS } from './nav';

export const OPEN_DELAY_MS = 120;
export const CLOSE_DELAY_MS = 180;
const TOUCH_QUERY = '(hover: none)';
const ITEM_SELECTOR = '[role="menuitem"]';
/** Keeps a wide panel (Guides) this far inside the viewport's right edge on narrow bars. */
const VIEWPORT_MARGIN_PX = 16;

export interface NavPanelOptions {
  openDelay?: number;
  closeDelay?: number;
  /** True on a device with no hover: the first tap on a door opens it, the second navigates. */
  isTouch?: () => boolean;
}

export interface NavPanels {
  open(key: string): void;
  close(): void;
  destroy(): void;
}

interface Door {
  readonly key: string;
  readonly group: HTMLElement;
  readonly link: HTMLElement;
  readonly toggle: HTMLElement;
  readonly panel: HTMLElement;
  openTimer: ReturnType<typeof setTimeout> | undefined;
  closeTimer: ReturnType<typeof setTimeout> | undefined;
}

const browserIsTouch = (): boolean =>
  typeof window.matchMedia === 'function' && window.matchMedia(TOUCH_QUERY).matches;

function readDoor(group: HTMLElement): Door | null {
  const key = group.dataset.group;
  const link = group.querySelector<HTMLElement>('.site-door');
  const toggle = group.querySelector<HTMLElement>('.site-door-toggle');
  const panel = group.querySelector<HTMLElement>('.site-panel');
  if (key === undefined || link === null || toggle === null || panel === null) return null;
  return { key, group, link, toggle, panel, openTimer: undefined, closeTimer: undefined };
}

const itemsOf = (door: Door): HTMLElement[] => [...door.panel.querySelectorAll<HTMLElement>(ITEM_SELECTOR)];

function cancelTimers(door: Door): void {
  clearTimeout(door.openTimer);
  clearTimeout(door.closeTimer);
  door.openTimer = undefined;
  door.closeTimer = undefined;
}

function paint(door: Door, open: boolean): void {
  door.panel.hidden = !open;
  door.link.setAttribute('aria-expanded', String(open));
  door.toggle.setAttribute('aria-expanded', String(open));
}

/** Slides a panel left just enough to sit inside the viewport; a panel that fits stays put. */
function clampToViewport(panel: HTMLElement): void {
  panel.style.translate = '';
  const { width, right } = panel.getBoundingClientRect();
  if (width === 0) return;
  const overflow = right - (document.documentElement.clientWidth - VIEWPORT_MARGIN_PX);
  if (overflow > 0) panel.style.translate = `${-overflow}px 0`;
}

export function wireNavPanels(header: HTMLElement, options: NavPanelOptions = {}): NavPanels {
  const { openDelay = OPEN_DELAY_MS, closeDelay = CLOSE_DELAY_MS, isTouch = browserIsTouch } = options;
  const doors = [...header.querySelectorAll<HTMLElement>('.site-group[data-group]')].flatMap((group) => {
    const door = readDoor(group);
    return door === null ? [] : [door];
  });
  const lifetime = new AbortController();
  const listen = <K extends keyof HTMLElementEventMap>(
    target: HTMLElement,
    type: K,
    handler: (event: HTMLElementEventMap[K]) => void,
  ): void => target.addEventListener(type, handler, { signal: lifetime.signal });

  let openKey: string | null = null;
  const isOpen = (door: Door): boolean => openKey === door.key;

  function show(key: string | null): void {
    for (const door of doors) {
      cancelTimers(door);
      paint(door, door.key === key);
    }
    openKey = key;
    const opened = doors.find((door) => door.key === key);
    if (opened !== undefined) clampToViewport(opened.panel);
  }

  function closeToDoor(door: Door): void {
    show(null);
    door.link.focus();
  }

  function focusItem(door: Door, index: number): void {
    const items = itemsOf(door);
    items.at(index % items.length)?.focus();
  }

  function onEnter(door: Door, event: MouseEvent): void {
    if ((event as PointerEvent).pointerType === 'touch') return;
    clearTimeout(door.closeTimer);
    if (isOpen(door)) return;
    door.openTimer = setTimeout(() => show(door.key), openDelay);
  }

  function onLeave(door: Door, event: MouseEvent): void {
    if ((event as PointerEvent).pointerType === 'touch') return;
    clearTimeout(door.openTimer);
    if (!isOpen(door)) return;
    door.closeTimer = setTimeout(() => show(null), closeDelay);
  }

  function onLinkClick(door: Door, event: MouseEvent): void {
    if (isOpen(door) || !isTouch()) return;
    event.preventDefault();
    show(door.key);
  }

  function onToggleClick(door: Door, event: MouseEvent): void {
    event.preventDefault();
    show(isOpen(door) ? null : door.key);
  }

  function onLinkKeydown(door: Door, event: KeyboardEvent): void {
    const plainEnter =
      event.key === 'Enter' && !(event.ctrlKey || event.metaKey || event.shiftKey || event.altKey);
    if (event.key !== 'ArrowDown' && !(plainEnter && !isOpen(door))) return;
    event.preventDefault();
    show(door.key);
    focusItem(door, 0);
  }

  function onPanelKeydown(door: Door, event: KeyboardEvent): void {
    const step = event.key === 'ArrowDown' ? 1 : event.key === 'ArrowUp' ? -1 : 0;
    if (step === 0) return;
    event.preventDefault();
    const items = itemsOf(door);
    const at = items.findIndex((item) => item === document.activeElement);
    focusItem(door, (at + step + items.length) % items.length);
  }

  function onGroupKeydown(door: Door, event: KeyboardEvent): void {
    if (event.key === 'Escape' && isOpen(door)) {
      closeToDoor(door);
      return;
    }
    if (event.target === door.link) onLinkKeydown(door, event);
    else if (door.panel.contains(event.target as Node)) onPanelKeydown(door, event);
  }

  function onFocusOut(door: Door, event: FocusEvent): void {
    const next = event.relatedTarget;
    if (isOpen(door) && next instanceof Node && !door.group.contains(next)) show(null);
  }

  for (const door of doors) {
    listen(door.group, 'pointerenter', (event) => onEnter(door, event));
    listen(door.group, 'pointerleave', (event) => onLeave(door, event));
    listen(door.link, 'click', (event) => onLinkClick(door, event));
    listen(door.toggle, 'click', (event) => onToggleClick(door, event));
    listen(door.group, 'keydown', (event) => onGroupKeydown(door, event));
    listen(door.group, 'focusout', (event) => onFocusOut(door, event));
  }
  document.addEventListener(
    'click',
    (event) => {
      const inside =
        event.target instanceof Node && doors.some((door) => door.group.contains(event.target as Node));
      if (!inside) show(null);
    },
    { signal: lifetime.signal },
  );

  return {
    open: show,
    close: () => show(null),
    destroy: () => {
      lifetime.abort();
      show(null);
    },
  };
}

const guildGroup = NAV_GROUPS[GUILD_GROUP_INDEX];

/**
 * A member-only door says something else without a guild (`data-guildless-*`, rendered by
 * HeaderPanel from GUILDLESS_SWAPS): the member's own text is kept on first swap so the next
 * session guild can put it back.
 */
function swapText(link: HTMLElement, guildless: boolean): void {
  const label = link.dataset.guildlessLabel;
  if (label === undefined) return;
  const title = link.querySelector<HTMLElement>('.site-panel-title') ?? link;
  const desc = link.querySelector<HTMLElement>('.site-panel-desc');
  if (link.dataset.memberLabel === undefined) {
    link.dataset.memberLabel = title.textContent ?? '';
    link.dataset.memberDescription = desc?.textContent ?? '';
  }
  title.textContent = guildless ? label : link.dataset.memberLabel;
  if (desc !== null && link.dataset.guildlessDescription !== undefined) {
    desc.textContent = guildless ? link.dataset.guildlessDescription : (link.dataset.memberDescription ?? '');
  }
}

/**
 * Points the Guild panel's rows at the session's guild and names it, or restores the
 * find-or-claim defaults (`/guild`, "Your guild") for `null`.
 */
export function applySessionGuild(header: HTMLElement, guild: SessionGuild | null): void {
  const base = guild?.href ?? guildGroup.href;
  for (const link of header.querySelectorAll<HTMLElement>('[data-guild-tab]')) {
    const tab = link.dataset.guildTab;
    swapText(link, guild === null);
    if (guild === null && link.dataset.guildlessHref !== undefined) {
      link.setAttribute('href', link.dataset.guildlessHref);
    } else if (tab === GUILD_OVERVIEW_TAB) {
      link.setAttribute('href', base);
    } else {
      // Settings is a page of its own (`<guild>/settings`); every other tab is a hash.
      link.setAttribute('href', tab === 'settings' ? `${base}/settings` : `${base}#${tab}`);
    }
  }
  for (const name of header.querySelectorAll<HTMLElement>('[data-guild-name]')) {
    name.textContent = guild?.name ?? guildGroup.rows[0].label;
  }
}
