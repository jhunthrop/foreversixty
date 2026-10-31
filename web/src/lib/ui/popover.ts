// web/src/lib/ui/popover.ts
// The popover core every menu and list shares (design/specs/2026-10-30-controls-charts.md §A3,
// §C): the panel is shown in the browser's top layer through the Popover API, so no card,
// sticky row or chart can paint over it, and placed by Floating UI: 4 px from its anchor,
// flipped above when less than 240 px fits below, kept 8 px inside the viewport, sized to the
// room it has, closed when its anchor scrolls out of the viewport or a click lands outside it. One
// open overlay at a time (overlay.ts).
import { computePosition, offset, shift, size, type Middleware, type Placement } from '@floating-ui/dom';
import {
  LIST_MAX_WIDTH_PX,
  LIST_OFFSET_PX,
  VIEWPORT_PADDING_PX,
  chooseSide,
  listMaxHeight,
  type Side,
} from './listbox-model';
import { claimOverlay } from './overlay';

export type Align = 'start' | 'end';

export interface PopoverOptions {
  /** The element the panel hangs from. */
  anchor: HTMLElement;
  /** A `popover="manual"` element; a descendant of `container` in the DOM. */
  panel: HTMLElement;
  /** Clicks inside it do not close the panel (the field's wrapper). */
  container: HTMLElement;
  /** `end` for a field in a right slot (§A3.7). */
  align: Align;
  /** At least the anchor's width (lists); a menu sizes to its content. */
  matchWidth: boolean;
  /** Called once when the core closes the panel itself (outside click, scrolled away,
   *  another overlay opened). Not called by the handle's own close(). */
  onDismiss: () => void;
}

export interface PopoverHandle {
  close: () => void;
  /** Re-measures now (after the panel's content changed). */
  update: () => void;
}

/** Below, or above when below is short and above is roomier (§A4.2). The panel's natural
 *  height is its scrollHeight, read before any max-height is applied. */
function sideRule(align: Align, naturalHeight: () => number): Middleware {
  return {
    name: 'sideRule',
    fn({ rects, placement }) {
      const below = window.innerHeight - (rects.reference.y + rects.reference.height) - LIST_OFFSET_PX;
      const above = rects.reference.y - LIST_OFFSET_PX;
      const side = chooseSide(below, above, naturalHeight());
      const wanted: Placement = `${side}-${align}`;
      return wanted === placement ? {} : { reset: { placement: wanted } };
    },
  };
}

/** The room on a side, from the anchor to the viewport edge less the offset. */
function roomOn(side: Side, anchor: DOMRect): number {
  return side === 'bottom'
    ? window.innerHeight - anchor.bottom - LIST_OFFSET_PX
    : anchor.top - LIST_OFFSET_PX;
}

/** True once the anchor has scrolled out of the viewport: the panel closes (§A6.3). */
function outOfView(anchor: HTMLElement): boolean {
  const rect = anchor.getBoundingClientRect();
  return rect.bottom < 0 || rect.top > window.innerHeight || rect.right < 0 || rect.left > window.innerWidth;
}

/** Re-places the panel when the page scrolls or resizes, or either box changes size. Floating
 *  UI's own autoUpdate does the same with a layout-shift observer this panel does not need,
 *  at about 1 KB gzipped of the lazy chunk's budget (§A11.2). */
function watchLayout(anchor: HTMLElement, panel: HTMLElement, update: () => void): () => void {
  const onScroll = (event: Event): void => {
    if (!(event.target instanceof Node) || !panel.contains(event.target)) update();
  };
  window.addEventListener('scroll', onScroll, { capture: true, passive: true });
  window.addEventListener('resize', update, { passive: true });
  const observer = typeof ResizeObserver === 'function' ? new ResizeObserver(update) : null;
  observer?.observe(anchor);
  observer?.observe(panel);
  return () => {
    window.removeEventListener('scroll', onScroll, { capture: true });
    window.removeEventListener('resize', update);
    observer?.disconnect();
  };
}

export function openPopover(options: PopoverOptions): PopoverHandle {
  const { anchor, panel, container, align, matchWidth, onDismiss } = options;
  let closed = false;
  let natural = 0;

  panel.style.removeProperty('--popover-room');
  panel.showPopover();
  natural = panel.scrollHeight;

  const place = async (): Promise<void> => {
    const {
      x,
      y,
      placement: placed,
    } = await computePosition(anchor, panel, {
      strategy: 'fixed',
      placement: `bottom-${align}`,
      middleware: [
        offset(LIST_OFFSET_PX),
        sideRule(align, () => natural),
        size({
          padding: VIEWPORT_PADDING_PX,
          apply({ rects, placement, availableWidth }) {
            panel.style.maxWidth = `${Math.min(LIST_MAX_WIDTH_PX, availableWidth)}px`;
            const side = placement.startsWith('top') ? 'top' : 'bottom';
            const room = roomOn(side, anchor.getBoundingClientRect());
            panel.style.setProperty('--popover-room', `${listMaxHeight(room, Infinity)}px`);
            if (matchWidth) panel.style.minWidth = `${rects.reference.width}px`;
          },
        }),
        shift({ padding: VIEWPORT_PADDING_PX }),
      ],
    });
    if (closed) return;
    if (outOfView(anchor)) {
      dismiss();
      return;
    }
    panel.dataset.side = placed.startsWith('top') ? 'top' : 'bottom';
    Object.assign(panel.style, { left: `${x}px`, top: `${y}px` });
  };

  const stopWatching = watchLayout(anchor, panel, () => void place());
  void place();

  const onPointerDown = (event: PointerEvent): void => {
    if (!(event.target instanceof Node) || !container.contains(event.target)) dismiss();
  };
  document.addEventListener('pointerdown', onPointerDown, true);

  const release = claimOverlay(() => dismiss());

  function teardown(): void {
    if (closed) return;
    closed = true;
    stopWatching();
    document.removeEventListener('pointerdown', onPointerDown, true);
    release();
    try {
      panel.hidePopover();
    } catch {
      // Already hidden (removed from the DOM with its field): nothing to undo.
    }
  }

  function dismiss(): void {
    if (closed) return;
    teardown();
    onDismiss();
  }

  return {
    close: teardown,
    update: () => {
      panel.style.removeProperty('--popover-room');
      natural = panel.scrollHeight;
      void place();
    },
  };
}
