<!-- web/src/components/sim/TimelineTickTip.svelte -->
<!-- A tick's tooltip (spec 2026-11-04-sim-results 3.7): the icon, name and rank; the time and the
     resource after it; the next one on the lane; and, when the run recorded it, the outcome. -->
<script lang="ts">
  import type { ActionVisuals } from '../../lib/sim/action-visuals';
  import { preciseClock, seconds, tickFollow, type TickLane } from '../../lib/sim/timeline';
  import { timelineCopy } from '../../lib/sim/timeline-copy';
  import Overlay from '../ui/Overlay.svelte';
  import ActionIcon from './ActionIcon.svelte';

  import { grouped } from '../../lib/format/number';
  let {
    lane,
    index,
    resourceLabel,
    visuals,
    visualsLoaded,
  }: {
    lane: TickLane;
    index: number;
    resourceLabel: string | null;
    visuals: ActionVisuals;
    visualsLoaded: boolean;
  } = $props();

  const event = $derived(lane.events[index]);
  const visual = $derived(visuals(lane.action));
  const follow = $derived(tickFollow(lane, index));
</script>

{#if event !== undefined}
  <Overlay testid="timeline-tick-tip">
    <div class="px-3 py-[10px] text-[13px] leading-5">
      <div class="mb-1 flex items-center gap-2">
        <ActionIcon url={visual.icon} loaded={visualsLoaded} size={24} />
        <span class="text-strong text-[15px] font-semibold">{lane.name}</span>
        {#if visual.rank > 0}<span class="text-muted">{timelineCopy.rank(visual.rank)}</span>{/if}
      </div>
      <p class="text-text">
        {#if resourceLabel !== null && event.resource !== null}
          {preciseClock(event.atMs)} · {resourceLabel.toLowerCase()} after
          <strong class="text-strong font-semibold">{Math.round(event.resource)}</strong>
        {:else}
          {preciseClock(event.atMs)}
        {/if}
      </p>
      <p class="text-muted">
        {follow.nextInMs === null
          ? timelineCopy.tipLast(follow.name)
          : timelineCopy.tipNext(follow.name, seconds(follow.nextInMs))}
      </p>
      {#if event.outcome !== undefined}
        <p class="text-text" data-testid="timeline-tick-outcome">
          {timelineCopy.tipOutcome(event.outcome, event.amount === undefined ? null : grouped(event.amount))}
        </p>
      {/if}
    </div>
  </Overlay>
{/if}
