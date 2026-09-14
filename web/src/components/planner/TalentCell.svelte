<!-- web/src/components/planner/TalentCell.svelte -->
<!-- One talent. 44px minimum per design/DESIGN-SYSTEM.md. Click adds a point, right-click or
     a long press removes the last one, Enter and Backspace do the same from the keyboard
     (handled by TreeGrid, which owns the roving tabindex). Hover or focus opens the tooltip
     with the current rank's text and the next rank's. -->
<script lang="ts">
  import { dataUrl } from '../../lib/planner/load';
  import { canAddPoint } from '../../lib/planner/rules';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { Talent } from '../../lib/planner/types';

  let {
    store,
    talent,
    focused,
    onfocuscell,
  }: {
    store: PlannerStore;
    talent: Talent;
    focused: boolean;
    onfocuscell: () => void;
  } = $props();

  const LONG_PRESS_MS = 500;

  const rank = $derived(store.ranks.get(talent.id) ?? 0);
  const maxed = $derived(rank >= talent.max_rank);
  const available = $derived(
    store.talentIndex !== null && canAddPoint(store.talentIndex, store.order, talent.id).ok,
  );
  const iconSrc = $derived(dataUrl(store.treeVersion, `icons/${talent.icon}.webp`));
  const tooltipId = $derived(`talent-tip-${talent.id}`);

  let open = $state(false);
  let iconBroken = $state(false);
  let pressTimer: ReturnType<typeof setTimeout> | undefined;
  // True once the gesture in progress has removed its point. Two things would otherwise
  // remove a second one: the click that ends a press (on touch, releasing a long press
  // always fires one), and the platform's own contextmenu, which touch-and-hold raises at
  // about the same threshold as the timer below. Whichever path fires first sets this, and
  // it is what makes the other a no-op, so one gesture is always exactly one removal.
  let gestureRemoved = false;

  const borderClass = $derived(
    maxed ? 'border-gold' : rank > 0 ? 'border-gold-deep' : available ? 'border-line' : 'border-line-soft',
  );

  function startPress(event: PointerEvent): void {
    // Only a primary press can become a long press. A secondary button raises contextmenu
    // on its own, so arming the timer for it too would remove twice for one right-click.
    if (event.button !== 0) return;
    gestureRemoved = false;
    pressTimer = setTimeout(() => {
      disarmPress();
      gestureRemoved = true;
      store.removePoint(talent.id);
    }, LONG_PRESS_MS);
  }

  /** Cancels a pending long press. Leaves `gestureRemoved` alone: the click still follows. */
  function disarmPress(): void {
    clearTimeout(pressTimer);
    pressTimer = undefined;
  }

  /** The press ended somewhere else, so no click follows and nothing is left to suppress. */
  function abandonPress(): void {
    disarmPress();
    gestureRemoved = false;
  }

  function removeOnContextMenu(event: MouseEvent): void {
    event.preventDefault();
    if (gestureRemoved) return;
    // A contextmenu raised while a press is pending is that press's own platform gesture,
    // and on touch a click still follows it; a right-click arms no timer and no click
    // follows it, so only the first case has anything left to suppress.
    const duringPress = pressTimer !== undefined;
    disarmPress();
    gestureRemoved = duringPress;
    store.removePoint(talent.id);
  }

  function add(): void {
    if (gestureRemoved) {
      gestureRemoved = false;
      return;
    }
    store.addPoint(talent.id);
  }
</script>

<div class="relative">
  <button
    type="button"
    tabindex={focused ? 0 : -1}
    aria-describedby={open ? tooltipId : undefined}
    aria-label={`${talent.name}, rank ${rank} of ${talent.max_rank}`}
    data-testid={`talent-${talent.id}`}
    data-rank={rank}
    class={`rounded-control bg-card-top relative flex h-11 w-11 items-center justify-center border md:h-12 md:w-12 ${borderClass} ${rank === 0 && !available ? 'opacity-50' : ''}`}
    onclick={add}
    oncontextmenu={removeOnContextMenu}
    onfocus={() => {
      open = true;
      onfocuscell();
    }}
    onblur={() => (open = false)}
    onmouseenter={() => (open = true)}
    onmouseleave={() => {
      open = false;
      abandonPress();
    }}
    onpointerdown={startPress}
    onpointerup={disarmPress}
    onpointercancel={abandonPress}
  >
    {#if iconBroken}
      <span class="text-muted font-display text-[13px] font-bold" aria-hidden="true">
        {talent.name.slice(0, 2)}
      </span>
    {:else}
      <img
        src={iconSrc}
        alt=""
        width="40"
        height="40"
        loading="lazy"
        decoding="async"
        class="rounded-control h-10 w-10 object-cover"
        onerror={() => (iconBroken = true)}
      />
    {/if}
    <span
      class={`tabular rounded-pill bg-bg absolute -right-1 -bottom-1 border px-1 font-mono text-[11px] leading-[14px] ${maxed ? 'border-gold text-gold' : 'border-line text-text'}`}
    >
      {rank}/{talent.max_rank}
    </span>
  </button>

  {#if open}
    <div
      id={tooltipId}
      role="tooltip"
      class="border-line bg-raised rounded-panel absolute top-full left-0 z-30 mt-2 flex w-[260px] flex-col gap-2 border p-3 shadow-[0_12px_30px_rgba(0,0,0,.45)]"
    >
      <span class="text-strong font-display text-[14px] font-bold">{talent.name}</span>
      <span class="tabular text-muted font-mono text-[12px]">
        Rank {rank} of {talent.max_rank}
      </span>
      {#if rank > 0}
        <p class="text-text text-[13px] leading-snug">{talent.ranks[rank - 1].description}</p>
      {/if}
      {#if rank < talent.max_rank}
        <p class="text-muted text-[13px] leading-snug">
          <span class="label text-muted">Next rank</span>
          {talent.ranks[rank].description}
        </p>
      {/if}
    </div>
  {/if}
</div>
