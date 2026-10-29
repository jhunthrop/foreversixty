<!-- web/src/components/planner/PlannerFact.svelte -->
<!-- One fact of the planner's facts rail as a tile (polish spec 5.1): the same recipe as
     `ui/StatTile.svelte` -- quiet ground, figure in the heading face over a micro label -- with
     two things StatTile has no slot for: the figure arrives as a snippet (so each figure keeps
     its own test id and tone, and the DPS tile can hold its "Show" button) and an optional note
     line sits under the label. `min-w-0` lets the tile shrink inside its grid column, which is
     what keeps the rail inside the phone viewport. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    label,
    figure,
    note,
    figureClass = '',
  }: {
    label: string;
    figure: Snippet;
    note?: Snippet;
    /** Extra classes for the figure row (the DPS tile reserves a button's height). */
    figureClass?: string;
  } = $props();
</script>

<div
  class="rounded-control border-line-soft flex min-w-0 flex-col gap-1 border bg-white/[.03] px-[14px] py-3"
>
  <div class="figure-stat text-strong flex items-center {figureClass}">{@render figure()}</div>
  <span class="text-muted text-[11px] font-bold tracking-[.08em] uppercase">{label}</span>
  {#if note !== undefined}{@render note()}{/if}
</div>
