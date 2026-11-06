<!-- web/src/components/sim/SimCastList.svelte -->
<!-- The timeline's cast list (spec 2026-11-04-sim-results 3.7; sim spec 2026-11-06 §6 Timeline foot):
     one line that counts the casts, says what Swings and ticks adds while it is Off, and opens the
     presses in view as open Rows ("List every cast"). It is the table a screen reader reads, on
     /sim/<id> and after a run on /sim alike. -->
<script lang="ts">
  import { LINK } from '../../lib/ui/primitives';
  import Row from '../ui/Row.svelte';
  import type { ActionNames } from '../../lib/sim/action-names';
  import { resolveActionName } from '../../lib/sim/action-names';
  import type { ActionVisuals } from '../../lib/sim/action-visuals';
  import {
    baseActionKey,
    preciseClock,
    pressesInView,
    type TimelineEvent,
    type TimelineModel,
    type ZoomWindow,
  } from '../../lib/sim/timeline';
  import { timelineCopy } from '../../lib/sim/timeline-copy';
  import ActionIcon from './ActionIcon.svelte';

  let {
    model,
    window: view,
    closeUp,
    swings,
    names,
    visuals,
    visualsLoaded,
  }: {
    model: TimelineModel;
    window: ZoomWindow;
    closeUp: boolean;
    /** Swings and ticks is On: the line has nothing to add about them. */
    swings: boolean;
    names: ActionNames | null;
    visuals: ActionVisuals;
    visualsLoaded: boolean;
  } = $props();

  let expanded = $state(false);

  const inView = $derived(pressesInView(model, view));
  const total = $derived(model.presses.length);
  const countLine = $derived(
    closeUp ? timelineCopy.castsInView(inView.length, total) : timelineCopy.castsInRun(total),
  );
  const resourceLabel = $derived(model.resource?.label ?? '');

  function pressName(press: TimelineEvent): string {
    const name = resolveActionName(baseActionKey(press.action), names);
    return press.queued ? timelineCopy.queued(name) : name;
  }
</script>

<div data-testid="sim-cast-list">
  <p class="text-muted pt-[10px] text-[13px] leading-5" data-testid="sim-cast-count">
    {swings ? countLine : `${countLine} ${timelineCopy.castsTail(model.periodicName)}`}
    {#if inView.length > 0}<button
        type="button"
        class="{LINK} text-[13px]"
        aria-expanded={expanded}
        onclick={() => (expanded = !expanded)}
        data-testid="sim-cast-toggle">{expanded ? timelineCopy.hideList : timelineCopy.listEvery} ›</button
      >{/if}
  </p>
  {#if expanded}
    <ol class="flex flex-col pt-2">
      {#each inView as press, index (`${press.atMs}-${press.action}-${index}`)}
        <li data-testid="sim-cast-row">
          <Row title={pressName(press)} density="slim" open>
            {#snippet lead()}<span class="text-muted w-[72px] text-[14px] tabular-nums"
                >{preciseClock(press.atMs)}</span
              ><span class="inline-flex h-6 w-6"
                ><ActionIcon url={visuals(press.action).icon} loaded={visualsLoaded} size={24} /></span
              >{/snippet}
            {#snippet trail()}{#if press.resource !== null}<span class="text-muted text-[14px] tabular-nums"
                  >{resourceLabel} {Math.round(press.resource)}</span
                >{/if}{/snippet}
          </Row>
        </li>
      {/each}
    </ol>
  {/if}
</div>
