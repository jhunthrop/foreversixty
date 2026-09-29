<!-- web/src/components/ItemHover.svelte -->
<!-- The reusable wrapper: an anchor (the default icon+name pill, or `children` for a
     caller's own markup) that opens ItemTooltip.svelte on hover or focus, a touch tap
     toggles it, ESC closes it, and only one is ever open across the whole page. Works two
     ways, same as lib/items/lookup.ts's two loaders:

       * runtime (planner/sim islands): `model` absent, so opening fetches the tooltip's
         model once per item and caches it (lookup.ts's fetchItemTooltipModel).
       * build time (an SSR page like /bis): `model` given, already built off disk
         (lookup.ts's readItemTooltipModel) -- opening mounts the panel with no fetch at all.

     The heavy panel (ItemTooltip.svelte) is imported with a dynamic import() and mounted
     imperatively (svelte's own mount()/unmount()), never a static import -- the exact
     pattern GearPanel.svelte's own BisSlotPopover trigger already uses, and for the same
     reason: this component itself IS statically imported everywhere an item appears, so
     only its own small shell -- not the panel's full layout -- may ride on every page's
     base bundle. -->
<script module lang="ts">
  import type { ItemTooltipModel } from '../lib/items/tooltip';

  // Page-wide "only one open" rule: a mouse hover already only ever touches one anchor, but
  // a tap-to-open on touch, or tabbing past several before one is closed, would otherwise
  // leave more than one visibly open at once. Module-scope state is shared by every
  // ItemHover instance the page mounts.
  let activeHoverId = $state<string | null>(null);
</script>

<script lang="ts">
  import { mount, unmount, type Component, type Snippet } from 'svelte';
  import { fetchItemTooltipModel } from '../lib/items/lookup';
  import { rarityClassFor } from '../lib/planner/items';
  import { dataUrl } from '../lib/planner/load';

  let {
    itemId,
    classSlug,
    build,
    model,
    class: pillClass = '',
    children,
  }: {
    itemId: number;
    classSlug: string;
    build: string;
    model?: ItemTooltipModel;
    class?: string;
    children?: Snippet;
  } = $props();

  const hoverId = `item-hover-${itemId}-${Math.random().toString(36).slice(2, 8)}`;
  const tooltipId = `${hoverId}-tip`;
  const isOpen = $derived(activeHoverId === hoverId);

  let resolvedModel = $state<ItemTooltipModel | null>(model ?? null);
  // Which build/class/item resolvedModel was fetched for, so a prop change (a different
  // itemId reusing the same instance -- SubstitutionChips re-keys its rows by item id, so
  // this should not happen in practice, but nothing here assumes it can't) re-fetches
  // instead of showing a stale model.
  let loadedFor = '';

  let host: HTMLDivElement | undefined = $state();
  let PanelComponent: Component<{ model: ItemTooltipModel; build: string; id: string }> | null = null;
  let panelModuleLoad: Promise<unknown> | null = null;
  let mountedInstance: object | null = null;

  function unmountPanel(): void {
    if (mountedInstance === null) return;
    unmount(mountedInstance);
    mountedInstance = null;
  }

  /** Clamps the panel inside the viewport, the same edge-aware move TalentCell's own
   *  tooltip makes: near the right edge of a phone, hanging the panel from the anchor's
   *  left edge would otherwise push it off-screen and scroll the page sideways. Applied as
   *  a CSS variable on the host rather than a prop into the mounted panel -- mount()'s
   *  props are a point-in-time snapshot (GearPanel's own popover host makes the identical
   *  trade), and remeasuring only needs the DOM, not a re-mount. */
  function positionPanel(): void {
    if (host === undefined) return;
    const anchor = host.getBoundingClientRect();
    const viewport = document.documentElement.clientWidth;
    const margin = 8;
    const panelWidth = 288; // w-72
    const maxLeft = viewport - margin - panelWidth - anchor.left;
    host.style.setProperty('--item-tooltip-shift', `${Math.min(0, maxLeft)}px`);
  }

  function mountPanel(): void {
    if (resolvedModel === null || PanelComponent === null || host === undefined) return;
    unmountPanel();
    mountedInstance = mount(PanelComponent, {
      target: host,
      props: { model: resolvedModel, build, id: tooltipId },
    });
    positionPanel();
  }

  function ensureModel(): void {
    if (model !== undefined) {
      resolvedModel = model;
      return;
    }
    const key = `${build}::${classSlug}::${itemId}`;
    if (loadedFor === key) return;
    loadedFor = key;
    void fetchItemTooltipModel(build, classSlug, itemId).then((loaded) => {
      if (loadedFor !== key) return;
      resolvedModel = loaded;
      if (isOpen) mountPanel();
    });
  }

  function ensurePanelLoaded(): void {
    if (PanelComponent !== null) {
      mountPanel();
      return;
    }
    if (panelModuleLoad !== null) return;
    panelModuleLoad = import('./ItemTooltip.svelte').then((imported) => {
      PanelComponent = imported.default;
      if (isOpen) mountPanel();
    });
  }

  function open(): void {
    activeHoverId = hoverId;
    ensureModel();
    ensurePanelLoaded();
  }

  function close(): void {
    if (activeHoverId !== hoverId) return;
    activeHoverId = null;
    unmountPanel();
  }

  /** Touch has no hover: a tap toggles instead. `preventDefault` keeps a tap on this pill
   *  from also activating a `<label>` it happens to sit inside (CandidateRows.svelte's
   *  checkbox row) -- a label's own click-forwarding is default browser behaviour, not a
   *  bubbling listener, so stopPropagation alone would not stop it. */
  function toggle(event: MouseEvent): void {
    event.preventDefault();
    if (isOpen) close();
    else open();
  }

  function onkeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || !isOpen) return;
    event.stopPropagation();
    close();
  }

  // ItemHover can unmount with its panel still showing (the row it names scrolls out of a
  // virtualised list, a route change) -- mount() instances live outside the normal
  // component tree and need their own teardown, the same as GearPanel's popover.
  $effect(() => () => unmountPanel());
</script>

<div class="relative inline-flex" bind:this={host}>
  <span
    class={`inline-flex items-center gap-1 ${pillClass}`}
    tabindex="0"
    role="button"
    aria-haspopup="true"
    aria-describedby={isOpen ? tooltipId : undefined}
    onmouseenter={open}
    onmouseleave={close}
    onfocus={open}
    onblur={close}
    onclick={toggle}
    {onkeydown}
  >
    {#if children}
      {@render children()}
    {:else if resolvedModel}
      <img
        src={dataUrl(build, `icons/${resolvedModel.icon}.webp`)}
        alt=""
        width="20"
        height="20"
        loading="lazy"
        decoding="async"
        class="rounded-control border-line h-5 w-5 border object-cover"
      />
      <span class={rarityClassFor(resolvedModel.quality)}>{resolvedModel.name}</span>
    {:else}
      <span class="text-muted">Item {itemId}</span>
    {/if}
  </span>
</div>
