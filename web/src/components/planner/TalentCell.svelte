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
  // Releasing a long press also fires a click, and on touch that is the only way the press
  // ends. Without this flag the click would add straight back the point the press removed.
  let removedByLongPress = false;

  const borderClass = $derived(
    maxed ? 'border-gold' : rank > 0 ? 'border-gold-deep' : available ? 'border-line' : 'border-line-soft',
  );

  function startPress(): void {
    removedByLongPress = false;
    pressTimer = setTimeout(() => {
      removedByLongPress = true;
      store.removePoint(talent.id);
    }, LONG_PRESS_MS);
  }

  /** The press ended on the cell: a click follows, and it is the one the flag suppresses. */
  function endPress(): void {
    clearTimeout(pressTimer);
    pressTimer = undefined;
  }

  /** The press ended somewhere else, so no click follows and nothing is left to suppress. */
  function abandonPress(): void {
    endPress();
    removedByLongPress = false;
  }

  function add(): void {
    if (removedByLongPress) {
      removedByLongPress = false;
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
    oncontextmenu={(event) => {
      event.preventDefault();
      store.removePoint(talent.id);
    }}
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
    onpointerup={endPress}
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
