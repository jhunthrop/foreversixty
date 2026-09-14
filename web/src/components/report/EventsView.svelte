<!-- web/src/components/report/EventsView.svelte -->
<!-- Every cast, aura application and removal, death and killing hit the summary itself
     timestamps, merged and time-ordered (src/lib/report/events.ts). The complete event
     stream -- every field of every line, not just what the summary keeps -- is
     events.parquet, which is the Queries view's job (Task 16): two to ten megabytes and a
     DuckDB instance for a question this view answers for free. -->
<script lang="ts">
  import { SvelteSet } from 'svelte/reactivity';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import { EVENT_KINDS, filterEvents, summaryEvents, type EventKind } from '../../lib/report/events';
  import type { Summary } from '../../lib/report/types';

  let { summary, classOf }: { summary: Summary; classOf: Map<string, string> } = $props();

  // A SvelteSet, not a plain one: toggle() below adds and removes single entries in place,
  // which is exactly the per-entry tracking SvelteSet exists for, unlike the Maps this
  // component's siblings rebuild wholesale (see ReportView.svelte's `treeSizes` and
  // `summaries` for that other case and why they stay plain). SvelteSet is already
  // reactive on its own, so it is not also wrapped in $state.
  const kinds = new SvelteSet<EventKind>(EVENT_KINDS.map((kind) => kind.id));
  let search = $state('');

  const all = $derived(summaryEvents(summary));
  const shown = $derived(filterEvents(all, kinds, search));

  function toggle(kind: EventKind): void {
    if (kinds.has(kind)) kinds.delete(kind);
    else kinds.add(kind);
  }
</script>

<div class="flex flex-col gap-3" data-testid="events-view">
  <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
    {#each EVENT_KINDS as kind (kind.id)}
      <label class="flex min-h-11 items-center gap-2 text-[13px] md:min-h-0">
        <input type="checkbox" checked={kinds.has(kind.id)} onchange={() => toggle(kind.id)} />
        {kind.label}
      </label>
    {/each}
    <label class="label text-muted flex items-center gap-2" for="events-search">
      Find
      <input
        id="events-search"
        class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
        type="search"
        bind:value={search}
        data-testid="events-search"
      />
    </label>
  </div>

  <p class="text-muted text-[12px]">
    Everything the fight summary timestamps. The complete event stream, every field of every line, is
    in Queries.
  </p>

  {#if shown.length === 0}
    <p class="text-muted text-[14px]" data-testid="table-empty">Nothing matches.</p>
  {:else}
    <ul class="flex flex-col" data-testid="event-list">
      {#each shown as event, index (`${event.atMs}-${index}`)}
        <li class="border-line-soft grid min-h-11 grid-cols-[68px_minmax(0,1fr)_auto] items-center gap-3 border-b px-2 py-1 text-[14px]">
          <span class="text-muted font-mono tabular text-[12px]">{formatDuration(event.atMs)}</span>
          <span class="truncate" style={`color: ${classColorVar(classOf.get(event.guid))}`}>{event.text}</span>
          <span class="font-mono tabular text-right text-[13px]">
            {event.amount === undefined ? '' : formatAmount(event.amount)}
          </span>
        </li>
      {/each}
    </ul>
  {/if}
</div>
