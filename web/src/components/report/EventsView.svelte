<!-- web/src/components/report/EventsView.svelte -->
<!-- Every cast, aura application and removal, death and killing hit the summary itself
     timestamps, merged and time-ordered (src/lib/report/events.ts). The complete event
     stream -- every field of every line, not just what the summary keeps -- is
     events.parquet, which is the Queries view's job (Task 16): two to ten megabytes and a
     DuckDB instance for a question this view answers for free. -->
<script lang="ts">
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import { EVENT_KINDS, filterEvents, summaryEvents, type EventKind } from '../../lib/report/events';
  import type { Summary } from '../../lib/report/types';

  let {
    summary,
    classOf,
    inScope = () => true,
    off = [],
    search = '',
    onPatch = () => {},
  }: {
    summary: Summary;
    classOf: Map<string, string>;
    /** The page's Source scope: a line stays when anyone it involves is in scope. */
    inScope?: (guid: string) => boolean;
    /** The kinds switched off and the find text, from the url, so a copied link keeps them. */
    off?: string[];
    search?: string;
    onPatch?: (patch: { eventsOff?: string[]; find?: string }) => void;
  } = $props();

  const kinds = $derived(
    new Set<EventKind>(EVENT_KINDS.map((kind) => kind.id).filter((kind) => !off.includes(kind))),
  );

  /** See FilterBar.svelte: the label around a checkbox is its 44px target, not the box. */
  const check = 'accent-gold';

  const all = $derived(summaryEvents(summary).filter((event) => event.guids.some(inScope)));
  const matching = $derived(filterEvents(all, kinds, search));
  /** How many rows are on the page: a wipe's stream runs to thousands. */
  const PAGE = 200;
  let limit = $state(PAGE);
  const shown = $derived(matching.slice(0, limit));
  $effect(() => {
    // Back to the first page whenever the filter changes.
    void [kinds.size, search];
    limit = PAGE;
  });

  function toggle(kind: EventKind): void {
    onPatch({ eventsOff: kinds.has(kind) ? [...off, kind] : off.filter((entry) => entry !== kind) });
  }
</script>

<div class="flex flex-col gap-3" data-testid="events-view">
  <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
    {#each EVENT_KINDS as kind (kind.id)}
      <label class="flex min-h-11 items-center gap-2 text-[13px] md:min-h-0">
        <input class={check} type="checkbox" checked={kinds.has(kind.id)} onchange={() => toggle(kind.id)} />
        {kind.label}
      </label>
    {/each}
    <label class="label text-muted flex items-center gap-2" for="events-search">
      Find
      <input
        id="events-search"
        class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
        type="search"
        value={search}
        oninput={(event) => onPatch({ find: (event.currentTarget as HTMLInputElement).value })}
        data-testid="events-search"
      />
    </label>
  </div>

  <p class="text-muted text-[12px]">
    <span class="tabular font-mono">{matching.length}</span> events the summary timestamps: casts, auras going up
    and down, the hits before each death, and the deaths. The complete event stream, every field of every line,
    is in Queries.
  </p>

  {#if shown.length === 0}
    <p class="text-muted text-[14px]" data-testid="table-empty">Nothing matches.</p>
  {:else}
    <ul class="flex flex-col" data-testid="event-list">
      {#each shown as event, index (`${event.atMs}-${index}`)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[68px_minmax(0,1fr)_auto] items-center gap-3 border-b px-2 py-1 text-[14px]"
        >
          <span class="text-muted tabular font-mono text-[12px]">{formatDuration(event.atMs)}</span>
          <span class="truncate" style={`color: ${classColorVar(classOf.get(event.guid))}`}>{event.text}</span
          >
          <span class="tabular text-right font-mono text-[13px]">
            {event.amount === undefined ? '' : formatAmount(event.amount)}
          </span>
        </li>
      {/each}
    </ul>
    {#if matching.length > shown.length}
      <button
        type="button"
        class="border-line-warm rounded-control text-text inline-flex h-11 items-center self-start border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
        data-testid="events-more"
        onclick={() => (limit += PAGE)}
      >
        Show {Math.min(PAGE, matching.length - shown.length)} more · {matching.length - shown.length} not yet shown
      </button>
    {/if}
  {/if}
</div>
