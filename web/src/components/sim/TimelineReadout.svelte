<!-- web/src/components/sim/TimelineReadout.svelte -->
<!-- The crosshair's readout (spec 2026-11-04-sim-results 3.7): the time, the resource and what is
     up on you and on the target, then the five most recent casts. Pinned, it gains a close
     button. Positioned by its parent: floating beside the line on a desktop, full width under the
     brush on a phone. -->
<script lang="ts">
  import type { ActionVisuals } from '../../lib/sim/action-visuals';
  import { seconds, type Readout } from '../../lib/sim/timeline';
  import { timelineCopy } from '../../lib/sim/timeline-copy';
  import Icon from '../ui/Icon.svelte';
  import { ICON_CLOSE } from '../../lib/ui/ui-icons';
  import Divided from '../ui/Divided.svelte';
  import Overlay from '../ui/Overlay.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';

  let {
    readout,
    pinned,
    visuals,
    visualsLoaded,
    onclose,
  }: {
    readout: Readout;
    pinned: boolean;
    visuals: ActionVisuals;
    visualsLoaded: boolean;
    onclose: () => void;
  } = $props();

  const SEPARATOR = ' · ';
</script>

<Overlay role={pinned ? 'dialog' : 'tooltip'} label={readout.time} testid="timeline-readout">
  <div class="p-3" data-pinned={pinned ? 'true' : undefined}>
    <div class="flex items-start justify-between gap-2">
      <span
        class="font-heading text-strong text-[18px] leading-6 font-extrabold"
        data-testid="timeline-readout-time">{readout.time}</span
      >
      {#if pinned}
        <button
          type="button"
          class="close text-muted hover:text-strong focus-visible:outline-gold -mt-3 -mr-3 inline-flex h-11 w-11 shrink-0 items-center justify-center focus-visible:outline-2"
          aria-label={timelineCopy.closeReadout}
          onclick={(event) => {
            event.stopPropagation();
            onclose();
          }}
          data-testid="timeline-readout-close"
        >
          <Icon icon={ICON_CLOSE} size={16} />
        </button>
      {/if}
    </div>
    <p class="text-muted text-[13px] leading-[19px]" data-testid="timeline-readout-state">
      {#if readout.resource !== null}
        {readout.resource.label}
        <strong class="text-strong font-semibold">{readout.resource.value}</strong>{SEPARATOR}
      {/if}{timelineCopy.readoutOnYou(readout.onYou)}
    </p>
    {#if readout.onTarget.length > 0}
      <p class="text-muted text-[13px] leading-[19px]" data-testid="timeline-readout-target">
        {timelineCopy.readoutOnTarget(readout.onTarget)}
      </p>
    {/if}
    {#if readout.casts.length > 0}
      <div class="mt-2">
        <Divided gap={0}>
          <ul class="flex w-full flex-col">
            {#each readout.casts as cast, index (`${cast.action}-${cast.agoMs}-${index}`)}
              <li class="flex h-[26px] items-center gap-2 text-[14px]" data-testid="timeline-readout-cast">
                <SpellIcon url={visuals(cast.action).icon} loading={!visualsLoaded} size={20} />
                <span class="text-text min-w-0 truncate">{cast.name}</span>
                <span class="numeric text-muted ml-auto shrink-0 text-[13px]"
                  >{timelineCopy.secondsAgo(seconds(cast.agoMs))}</span
                >
              </li>
            {/each}
          </ul>
        </Divided>
      </div>
    {/if}
  </div>
</Overlay>
