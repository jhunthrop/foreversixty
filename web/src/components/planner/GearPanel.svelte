<!-- web/src/components/planner/GearPanel.svelte -->
<!-- The 17-slot grid, the summed stats, and the active set bonuses. Two columns of slots on
     phone, four from md up; every slot button clears 44px. -->
<script lang="ts">
  import { mount, unmount, type Component } from 'svelte';
  import { plannerCopy } from '../../lib/planner/copy';
  import { addonCopy } from '../../lib/addon/copy';
  import { scoreItem, specKeyFor, weightsFor, type WeightsFile } from '../../lib/addon/score';
  import { pointsPerTree } from '../../lib/planner/derive';
  import { rarityClassFor, wornItemLabel } from '../../lib/planner/items';
  import { dataUrl, loadItemNames } from '../../lib/planner/load';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SLOTS, SLOT_LABELS, STAT_LABELS, type Slot, type StatKey } from '../../lib/planner/types';
  import { specLabel } from '../../lib/sim/spec-label';
  import { statLabel } from '../../lib/sim/stats';
  import ItemPicker from './ItemPicker.svelte';

  let { store, weights = [] }: { store: PlannerStore; weights?: WeightsFile } = $props();

  // Names for worn items the per-class file leaves out (a keepsake ring, a totem), fetched
  // once and only when a slot needs one -- the same lookup the simulator's strip makes.
  let outsideNames = $state<Record<string, string>>({});
  const needsOutsideNames = $derived(
    SLOTS.some((slot) => {
      const id = store.gear[slot];
      return id !== undefined && !store.itemIndex.has(id);
    }),
  );
  $effect(() => {
    if (!needsOutsideNames) return;
    const build = store.treeVersion;
    void loadItemNames(build).then((file) => {
      outsideNames = file.names;
    });
  });

  let openSlot = $state<Slot | null>(null);

  // The BiS hover popover (design step 1 of the bis-hover-web lane brief). Mouse and
  // keyboard get it from hover/focus, exactly like TalentCell's own tooltip; a tap on touch
  // focuses the button too (every mobile browser does this for a plain <button>), which
  // shows the popover the same way, alongside the item picker the tap always opened.
  //
  // Loaded and mounted imperatively with svelte's own mount()/unmount() (already bundled --
  // planner-island.ts's own boot calls mount(Planner, ...)), not a reactive `{#if}` around a
  // `$state`-held component reference: a dynamically resolved component tag compiles to
  // Svelte's generic dynamic-component runtime, which cost roughly as much as the popover
  // itself saved by moving out of the boot chunk. mount()/unmount() sidesteps that runtime
  // entirely -- see this lane's final report for the measured before/after.
  let hoveredSlot = $state<Slot | null>(null);
  const popoverHosts: Partial<Record<Slot, HTMLDivElement>> = {};
  let PopoverComponent: Component<{ store: PlannerStore; slot: Slot; spec: string; id: string }> | null =
    null;
  let popoverModuleLoad: Promise<unknown> | null = null;
  let mountedSlot: Slot | null = null;
  let mountedInstance: object | null = null;

  function unmountPopover(): void {
    if (mountedInstance === null) return;
    unmount(mountedInstance);
    mountedInstance = null;
    mountedSlot = null;
  }

  /**
   * Mounts the popover into `slot`'s own host div, replacing whichever slot's instance was
   * showing (only one is ever visible: hover and focus both move `hoveredSlot`, never add to
   * it). `store` is passed by reference -- the popover's own $derived/$effect read its
   * getters directly, the same reactivity every other reader of `store` gets, imperative
   * mount or not -- but `spec` is a plain string snapshot at mount time; a talent edit made
   * while a slot happens to still be hovered will not retarget an open popover, which the
   * hover/hide lifecycle here makes a narrow enough window to accept.
   */
  function mountPopoverFor(slot: Slot): void {
    if (mountedSlot === slot) return;
    unmountPopover();
    if (PopoverComponent === null) return;
    const host = popoverHosts[slot];
    if (host === undefined) return;
    mountedInstance = mount(PopoverComponent, {
      target: host,
      props: { store, slot, spec: specKey, id: `bis-hover-${slot}` },
    });
    mountedSlot = slot;
  }

  function showPopover(slot: Slot): void {
    hoveredSlot = slot;
    if (PopoverComponent !== null) {
      mountPopoverFor(slot);
      return;
    }
    if (popoverModuleLoad !== null) return;
    popoverModuleLoad = import('./BisSlotPopover.svelte').then((mod) => {
      PopoverComponent = mod.default;
      if (hoveredSlot !== null) mountPopoverFor(hoveredSlot);
    });
  }

  function hidePopoverIfShown(slot: Slot): void {
    if (hoveredSlot !== slot) return;
    hoveredSlot = null;
    unmountPopover();
  }

  // GearPanel itself can unmount with a popover still showing (navigating away from the
  // planner); mount() instances live outside the normal component tree and need their own
  // teardown.
  $effect(() => () => unmountPopover());

  /** Closes the popover only once focus has left the whole slot (button + popover), not
   *  when it moves from the button onto the popover's own "See the full list" link --
   *  `focusout` bubbles and carries `relatedTarget`, unlike `blur`, so that distinction is
   *  checkable here. */
  function onSlotFocusOut(event: FocusEvent, slot: Slot): void {
    const related = event.relatedTarget as Node | null;
    const container = event.currentTarget as HTMLElement;
    if (related && container.contains(related)) return;
    hidePopoverIfShown(slot);
  }

  const totals = $derived(
    (Object.entries(store.statTotals) as [StatKey, number][]).sort(([a], [b]) =>
      STAT_LABELS[a].localeCompare(STAT_LABELS[b]),
    ),
  );

  const specKey = $derived(
    store.talentIndex === null
      ? ''
      : specKeyFor(store.classSlug, pointsPerTree(store.talentIndex, store.order)),
  );
  const specWeights = $derived(weightsFor(weights, specKey));
  const setScore = $derived(
    specWeights === undefined
      ? null
      : SLOTS.reduce((total, slot) => {
          const id = store.gear[slot];
          const item = id === undefined ? undefined : store.itemIndex.get(id);
          return item === undefined ? total : total + scoreItem(item, specWeights.weights);
        }, 0),
  );
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-4 border p-4 md:mx-0"
  data-testid="gear-panel"
>
  <h2 class="section-title text-[15px]">Gear</h2>

  <div class="grid grid-cols-2 gap-2 md:grid-cols-4">
    {#each SLOTS as slot (slot)}
      {@const equippedId = store.gear[slot]}
      {@const item = equippedId === undefined ? undefined : store.itemIndex.get(equippedId)}
      <div
        class="relative"
        role="group"
        onmouseenter={() => showPopover(slot)}
        onmouseleave={() => hidePopoverIfShown(slot)}
        onfocusout={(event) => onSlotFocusOut(event, slot)}
      >
        <button
          type="button"
          class="border-line rounded-control bg-card-top flex min-h-11 w-full items-center gap-2 border px-2 py-1 text-left"
          data-testid={`slot-${slot}`}
          aria-label={item ? `${SLOT_LABELS[slot]}: ${item.name}` : `${SLOT_LABELS[slot]}: empty`}
          aria-describedby={hoveredSlot === slot ? `bis-hover-${slot}` : undefined}
          disabled={store.readOnly}
          onclick={() => (openSlot = openSlot === slot ? null : slot)}
          onfocus={() => showPopover(slot)}
        >
          {#if item}
            <!-- The aria-label above names the slot and the item, so the icon is decorative. -->
            <img
              src={dataUrl(store.treeVersion, `icons/${item.icon}.webp`)}
              alt=""
              width="28"
              height="28"
              loading="lazy"
              decoding="async"
              class="rounded-control border-line h-7 w-7 border object-cover"
            />
          {:else}
            <span class="rounded-control border-line-soft h-7 w-7 border" aria-hidden="true"></span>
          {/if}
          <span class="flex min-w-0 flex-col">
            <span class="label text-muted">{SLOT_LABELS[slot]}</span>
            <span
              class={`truncate text-[13px] font-semibold ${item ? rarityClassFor(item.quality) : 'text-muted'}`}
            >
              {wornItemLabel(item?.name, equippedId, outsideNames, plannerCopy)}
            </span>
          </span>
        </button>

        <div bind:this={popoverHosts[slot]}></div>
      </div>
    {/each}
  </div>

  {#if openSlot}
    <ItemPicker {store} slot={openSlot} weights={specWeights?.weights} onclose={() => (openSlot = null)} />
  {/if}

  {#if specWeights === undefined}
    <p class="text-muted text-[13px]" data-testid="gear-no-weights">{addonCopy.weightsMissing}</p>
  {:else}
    <p class="text-[13px]" data-testid="gear-set-score">
      {addonCopy.scoreColumn}: {setScore?.toFixed(1)}
    </p>
    <details data-testid="gear-weights">
      <summary class="text-muted text-[13px]">{addonCopy.weightsTitle(specLabel(specKey))}</summary>
      <p class="text-muted text-[13px]">{addonCopy.weightsAreOpinions}</p>
      <ul class="text-muted text-[13px]">
        {#each Object.entries(specWeights.weights).sort( ([a], [b]) => a.localeCompare(b) ) as [stat, weight] (stat)}
          <li>{statLabel(stat)}: {weight}</li>
        {/each}
      </ul>
      <ul class="text-[13px]">
        {#each specWeights.sources as source (source.url)}
          <li><a class="underline" href={source.url} rel="noopener">{source.label}</a></li>
        {/each}
      </ul>
    </details>
  {/if}

  <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
    <div class="flex flex-col gap-1" data-testid="gear-totals">
      <h3 class="label text-muted">Totals</h3>
      {#if totals.length === 0}
        <p class="text-muted text-[13px]">Nothing equipped.</p>
      {:else}
        <dl class="flex flex-col">
          {#each totals as [key, value] (key)}
            <div class="border-line-soft flex justify-between gap-3 border-b py-1 last:border-b-0">
              <dt class="text-muted text-[13px]">{STAT_LABELS[key]}</dt>
              <dd class="tabular text-strong font-mono text-[13px]">{value}</dd>
            </div>
          {/each}
        </dl>
      {/if}
    </div>

    <div class="flex flex-col gap-1" data-testid="gear-sets">
      <h3 class="label text-muted">Sets</h3>
      {#if store.activeSets.length === 0}
        <p class="text-muted text-[13px]">No set pieces equipped.</p>
      {:else}
        {#each store.activeSets as active (active.set.id)}
          <div class="border-line-soft flex flex-col gap-1 border-b py-2 last:border-b-0">
            <span class="text-strong text-[13px] font-semibold">
              {active.set.name}
              <span class="tabular text-gold font-mono">
                ({active.pieces}/{active.set.item_ids.length})
              </span>
            </span>
            {#each active.active as bonus (bonus.pieces)}
              <span class="text-muted text-[13px]">
                ({bonus.pieces}) {bonus.description}
              </span>
            {/each}
          </div>
        {/each}
      {/if}
    </div>
  </div>
</section>
