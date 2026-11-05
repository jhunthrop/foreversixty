<!-- web/src/components/sim/SimCastList.svelte -->
<!-- The timeline's cast list (spec 2026-11-04-sim-results 3.7): the presses in view, five rows
     until "Show all". It is the table a screen reader reads. The live /sim page keeps SampleLog's
     full per-resource table until that page adopts this layout. -->
<script lang="ts">
  import { button } from '../../lib/ui/button';
  import type { ActionNames } from '../../lib/sim/action-names';
  import { resolveActionName } from '../../lib/sim/action-names';
  import type { ActionVisuals } from '../../lib/sim/action-visuals';
  import {
    baseActionKey,
    COLLAPSED_CAST_ROWS,
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
    names,
    visuals,
    visualsLoaded,
  }: {
    model: TimelineModel;
    window: ZoomWindow;
    closeUp: boolean;
    names: ActionNames | null;
    visuals: ActionVisuals;
    visualsLoaded: boolean;
  } = $props();

  let expanded = $state(false);

  const inView = $derived(pressesInView(model, view));
  const total = $derived(model.presses.length);
  const shown = $derived(expanded ? inView : inView.slice(0, COLLAPSED_CAST_ROWS));
  const canToggle = $derived(inView.length > COLLAPSED_CAST_ROWS);
  const countLine = $derived(
    closeUp ? timelineCopy.castsInView(inView.length, total) : timelineCopy.castsAll(inView.length, total),
  );
  const resourceLabel = $derived(model.resource?.label ?? '');

  function pressName(press: TimelineEvent): string {
    const name = resolveActionName(baseActionKey(press.action), names);
    return press.queued ? timelineCopy.queued(name) : name;
  }
</script>

<div class="border-line-soft border-t" data-testid="sim-cast-list">
  <div class="flex min-h-[52px] flex-wrap items-center justify-between gap-x-4 gap-y-2 px-5 py-2">
    <p class="quiet-label !normal-case" data-testid="sim-cast-count">
      {countLine}. {timelineCopy.castsTail(model.periodicName)}
    </p>
    {#if canToggle}
      <button
        type="button"
        class={button('secondary', 'sm')}
        onclick={() => (expanded = !expanded)}
        data-testid="sim-cast-toggle"
      >
        {expanded ? timelineCopy.showFewer : timelineCopy.showAll(inView.length)}
      </button>
    {/if}
  </div>
  <ol class="flex flex-col">
    {#each shown as press, index (`${press.atMs}-${press.action}-${index}`)}
      <li
        class="border-line-soft grid h-10 grid-cols-[72px_24px_minmax(0,1fr)_auto] items-center gap-x-[10px] border-t px-5 text-[14px]"
        data-testid="sim-cast-row"
      >
        <span class="numeric text-muted">{preciseClock(press.atMs)}</span>
        <span class="inline-flex h-6 w-6"
          ><ActionIcon url={visuals(press.action).icon} loaded={visualsLoaded} size={24} /></span
        >
        <span class="text-text truncate">{pressName(press)}</span>
        {#if press.resource !== null}
          <span class="numeric text-muted">{resourceLabel} {Math.round(press.resource)}</span>
        {/if}
      </li>
    {/each}
  </ol>
</div>
