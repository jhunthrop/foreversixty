<!-- web/src/components/BisTooltipHost.svelte -->
<!-- One shared tooltip for a whole /bis/<class>/<spec> page (lane bis-tooltip-host,
     2026-09-29). That page used to wrap every row in its own ItemHover island
     (`client:idle`) -- 11 bands x 17 slots x 2 factions, plus runner-ups and the "what
     changed" diff, roughly 840 `<astro-island>`s hydrating before first paint. Each one
     serialised its own tooltip model into its props too: 2.5 MB of HTML, and CI's own
     Lighthouse mobile run put total blocking time at 135ms against the 100ms budget
     (lighthouserc.json) -- every deploy blocked on it.

     The page now renders every row as plain server markup (icon, quality-coloured name, a
     `data-bis-item="<id>"` attribute, `tabindex="0"`) with no per-row island at all, and
     this ONE `client:idle` island owns hover/focus/tap open-close for the entire page,
     delegated from a single document-level listener set rather than one Svelte instance
     per row -- exactly one ItemTooltip.svelte panel is ever mounted, moved between rows
     instead of remounted 840 times over.

     Item data rides along as a sibling `<script type="application/json" id="bis-tooltip-
     models">` tag, not this island's own props: Astro serialises a framework island's props
     into an HTML *attribute* (quote-escaped), which very nearly doubles a payload this size
     (a spec page names on the order of a thousand item ids across its bands) -- a sibling
     script tag holds the identical JSON unescaped, and this component reads it once, on
     mount, instead of paying that escaping cost.

     The open/close state machine, the lazy `import()` of the heavy panel, and the edge-
     aware viewport shift are ItemHover.svelte's own idea (components/ItemHover.svelte,
     lib/items/tooltip-position.ts) widened from "one host per item" to "one host for the
     whole page." Unlike ItemHover, this component makes no runtime fetch for a model this
     page's models map does not carry -- every id here already came from a build-time
     `readItemTooltipModel` read (lib/items/lookup.ts), the same the page used to hand
     ItemHover directly, so there is nothing left to fetch. -->
<script lang="ts">
  import { mount, unmount, onMount, type Component } from 'svelte';
  import type { ItemTooltipModel } from '../lib/items/tooltip';
  import { positionTooltipPanel } from '../lib/items/tooltip-position';

  let {
    build,
    modelsElementId = 'bis-tooltip-models',
    resolve = undefined,
  }: {
    build: string;
    modelsElementId?: string;
    /** A page with no build-time models (the character profile's Gear section, L8.6 fix round
     *  1): the model for an id the map lacks, fetched at runtime (lib/items/lookup.ts's
     *  fetchItemTooltipModel) and kept; null when the build does not carry the id. /bis passes
     *  none and keeps its build-time-only rule. */
    resolve?: (itemId: number) => Promise<ItemTooltipModel | null>;
  } = $props();

  // Fixed rather than per-row: exactly one panel is ever mounted, so every row's
  // aria-describedby (when it is the open one) can point at the same id.
  const TOOLTIP_ID = 'bis-item-tooltip';

  type TooltipPanel = Component<{ model: ItemTooltipModel; build: string; id: string }>;

  let models = new Map<number, ItemTooltipModel>();
  let activeHost: HTMLElement | null = null;
  let panelComponent: TooltipPanel | null = null;
  let panelModuleLoad: Promise<TooltipPanel> | null = null;
  let mountedInstance: object | null = null;
  // Bumped every time the panel closes (or is about to move to a different host), so a
  // dynamic import() that was already in flight for a now-stale open() never mounts after
  // the fact -- without this, a fast pointerenter -> pointerleave -> pointerenter on the
  // SAME row (all before the one shared `panelModuleLoad` promise first resolves) leaves
  // two "open" calls racing the same promise, and both would otherwise mount their own
  // panel instance into the row.
  let openToken = 0;
  /** When a pointer was last pressed: the focus that press gives a host is the pointer's, not
   *  the keyboard's, so it opens nothing (the click decides; L8.6 fix round 2 ruling 2). */
  let pressedAt = -Infinity;
  /** Where the mouse last moved to, and whether the page has scrolled since: a page scroll or
   *  an anchor landing slides a row under a stationary pointer, and the browser's fake
   *  pointerenter for that must not open a tooltip (L8.6 fix round 2 ruling 2). */
  let lastX = NaN;
  let lastY = NaN;
  let scrolledSinceMove = false;

  function loadModels(): void {
    const script = document.getElementById(modelsElementId);
    const text = script?.textContent;
    if (text === undefined || text === null || text === '') return;
    try {
      const parsed = JSON.parse(text) as Record<string, ItemTooltipModel>;
      models = new Map(Object.entries(parsed).map(([id, model]) => [Number(id), model]));
    } catch {
      models = new Map(); // malformed JSON never breaks the page over a tooltip
    }
  }

  function hostFor(target: EventTarget | null): HTMLElement | null {
    if (!(target instanceof Element)) return null;
    return target.closest<HTMLElement>('[data-bis-item]');
  }

  /** The row/pill the pointer or focus is moving TO, so a move between two elements still
   *  inside the same host (its icon, its name) never reads as leaving it. */
  function relatedHostFor(event: PointerEvent | FocusEvent): HTMLElement | null {
    const related = 'relatedTarget' in event ? event.relatedTarget : null;
    return hostFor(related as EventTarget | null);
  }

  function unmountPanel(): void {
    openToken += 1;
    if (mountedInstance !== null) unmount(mountedInstance);
    mountedInstance = null;
    activeHost?.removeAttribute('aria-describedby');
    activeHost = null;
  }

  /** The heavy panel (ItemTooltip.svelte) is imported with a dynamic import() and mounted
   *  imperatively, never a static import -- the exact pattern ItemHover.svelte and
   *  GearPanel's own BisSlotPopover trigger already use, so its weight never rides on this
   *  island's own base bundle. */
  function ensurePanelLoaded(): Promise<TooltipPanel> {
    if (panelComponent !== null) return Promise.resolve(panelComponent);
    panelModuleLoad ??= import('./ItemTooltip.svelte').then((imported) => {
      panelComponent = imported.default;
      return panelComponent;
    });
    return panelModuleLoad;
  }

  function openHost(host: HTMLElement): void {
    const itemId = Number(host.dataset.bisItem);
    const known = models.get(itemId);
    if (known === undefined && resolve === undefined) return; // no build-time model for this id -- nothing to show
    if (activeHost === host) return; // already open, or already opening, this exact host
    unmountPanel();
    activeHost = host;
    host.setAttribute('aria-describedby', TOOLTIP_ID);
    const token = ++openToken;
    const model: Promise<ItemTooltipModel | null> =
      known !== undefined
        ? Promise.resolve(known)
        : (resolve as (itemId: number) => Promise<ItemTooltipModel | null>)(itemId).catch(() => null);
    void Promise.all([ensurePanelLoaded(), model]).then(([component, resolved]) => {
      if (token !== openToken) return; // superseded (a close, or a later open) before this resolved
      if (resolved === null) return; // the build does not carry the id: no tooltip
      mountedInstance = mount(component, { target: host, props: { model: resolved, build, id: TOOLTIP_ID } });
      positionTooltipPanel(host);
    });
  }

  function closeHost(host: HTMLElement): void {
    if (activeHost !== host) return;
    unmountPanel();
  }

  /** Hover opens for a mouse only (L8.6 fix round 2 ruling 2): a touch or pen opens on its tap,
   *  and a row that slid under a stationary mouse (a scroll, an anchor landing) waits for the
   *  mouse to move. */
  function onPointerEnter(event: PointerEvent): void {
    if (event.pointerType !== 'mouse' || scrolledSinceMove) return;
    const host = hostFor(event.target);
    if (host !== null) openHost(host);
  }

  function onPointerMove(event: PointerEvent): void {
    if (event.pointerType !== 'mouse') return;
    if (event.clientX === lastX && event.clientY === lastY) return; // the browser's fake move after a scroll
    lastX = event.clientX;
    lastY = event.clientY;
    if (!scrolledSinceMove) return;
    scrolledSinceMove = false;
    // The mouse moved after the scroll: what it now rests on may open.
    const host = hostFor(document.elementFromPoint(event.clientX, event.clientY));
    if (host !== null) openHost(host);
  }

  function onPointerDown(): void {
    pressedAt = performance.now();
  }

  function onScroll(): void {
    scrolledSinceMove = true;
  }

  function onPointerLeave(event: PointerEvent): void {
    if (event.pointerType !== 'mouse') return;
    const host = hostFor(event.target);
    if (host === null || host === relatedHostFor(event)) return; // still inside the same host
    closeHost(host);
  }

  /** Keyboard focus opens; the focus a pointer press gives (a tap, a click) leaves it to the
   *  click. */
  function onFocusIn(event: FocusEvent): void {
    if (performance.now() - pressedAt < 1000) return;
    const host = hostFor(event.target);
    if (host !== null) openHost(host);
  }

  function onFocusOut(event: FocusEvent): void {
    const host = hostFor(event.target);
    if (host === null || host === relatedHostFor(event)) return;
    closeHost(host);
  }

  /** Touch has no hover: a tap toggles instead (a second tap on the same name closes it), the
   *  same rule ItemHover.svelte's own `toggle` follows; a tap or click anywhere else closes the
   *  open one. `preventDefault` keeps a tap from also activating anything the row happens to
   *  sit inside. */
  function onClick(event: MouseEvent): void {
    const host = hostFor(event.target);
    if (host === null) {
      if (activeHost !== null) closeHost(activeHost);
      return;
    }
    event.preventDefault();
    if (activeHost === host) closeHost(host);
    else openHost(host);
  }

  function onKeyDown(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || activeHost === null) return;
    event.stopPropagation();
    closeHost(activeHost);
  }

  onMount(() => {
    loadModels();
    // pointerenter/pointerleave never bubble, so only a capture-phase listener at the
    // document root sees them for every row on the page -- the same trick mouseenter/
    // mouseleave delegation always needs, since capture always fires regardless of a given
    // event's own bubbles flag.
    document.addEventListener('pointerenter', onPointerEnter, true);
    document.addEventListener('pointerleave', onPointerLeave, true);
    document.addEventListener('pointermove', onPointerMove, true);
    document.addEventListener('pointerdown', onPointerDown, true);
    document.addEventListener('scroll', onScroll, { capture: true, passive: true });
    document.addEventListener('focusin', onFocusIn);
    document.addEventListener('focusout', onFocusOut);
    document.addEventListener('click', onClick);
    document.addEventListener('keydown', onKeyDown);
    // A keyboard user who tabbed onto a row before this island hydrated (the island is
    // client:idle, and under CPU contention idle can come late) must still get the
    // tooltip: open for the element that already holds focus, then mark the document so
    // tests and any other code can wait for the host instead of racing it.
    const alreadyFocused = hostFor(document.activeElement);
    if (alreadyFocused !== null) openHost(alreadyFocused);
    document.documentElement.dataset.tooltipHost = 'ready';
    return () => {
      delete document.documentElement.dataset.tooltipHost;
      document.removeEventListener('pointerenter', onPointerEnter, true);
      document.removeEventListener('pointerleave', onPointerLeave, true);
      document.removeEventListener('pointermove', onPointerMove, true);
      document.removeEventListener('pointerdown', onPointerDown, true);
      document.removeEventListener('scroll', onScroll, { capture: true });
      document.removeEventListener('focusin', onFocusIn);
      document.removeEventListener('focusout', onFocusOut);
      document.removeEventListener('click', onClick);
      document.removeEventListener('keydown', onKeyDown);
      unmountPanel();
    };
  });
</script>
