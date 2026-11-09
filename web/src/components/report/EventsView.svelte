<!-- web/src/components/report/EventsView.svelte -->
<!-- Every cast, aura application and removal, death and killing hit the summary itself
     timestamps, merged and time-ordered (src/lib/report/events.ts). The complete event
     stream -- every field of every line, not just what the summary keeps -- is
     events.parquet, which is the Queries view's job (Task 16): two to ten megabytes and a
     DuckDB instance for a question this view answers for free.
     Drawn to logs spec 2026-11-07 section 6.2 (`web-report-events`): Section `Events` holding
     the one Plate, a StatTable of lines (the spell's icon, the sentence with each player's name
     in their class colour, Time, Amount), 200 a page. The Show chips are the rail's
     (EventsRail); `chips` draws them here until the frame places the rail. -->
<script lang="ts">
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import {
    EVENT_KINDS,
    filterEvents,
    sentenceParts,
    streamEvents,
    summaryEvents,
    type EventKind,
    type SummaryEvent,
  } from '../../lib/report/events';
  import type { StreamLine } from '../../lib/report/exact';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import type { Summary } from '../../lib/report/types';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { LINK } from '../../lib/ui/primitives';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import EventChips from './EventChips.svelte';

  let {
    summary,
    classOf,
    inScope = () => true,
    off = [],
    search = '',
    onPatch = () => {},
    loadStream = undefined,
    names = new Map<string, string>(),
    chips = true,
  }: {
    summary: Summary;
    classOf: Map<string, string>;
    /** The page's Source scope: a line stays when anyone it involves is in scope. */
    inScope?: (guid: string) => boolean;
    /** The kinds switched off and the find text, from the url, so a copied link keeps them. */
    off?: string[];
    search?: string;
    onPatch?: (patch: { eventsOff?: string[]; find?: string }) => void;
    /** Loads every hit and heal in the window from the fight's events; absent over the night. */
    loadStream?: () => Promise<StreamLine[]>;
    /** GUID to unit name, so an aura line can say who applied it. */
    names?: ReadonlyMap<string, string>;
    /** Draw the Show chips above the table; false once the frame draws EventsRail. */
    chips?: boolean;
  } = $props();

  const icons = new SpellIcons();

  let stream = $state<StreamLine[] | null>(null);
  let streaming = $state(false);
  let streamError = $state('');
  // A new window is a new stream: the loaded one answered the old question.
  $effect(() => {
    void summary;
    stream = null;
    streamError = '';
  });
  async function runStream(): Promise<void> {
    if (loadStream === undefined) return;
    streaming = true;
    streamError = '';
    try {
      stream = await loadStream();
    } catch (thrown) {
      streamError = thrown instanceof Error ? thrown.message : 'The events did not load.';
    } finally {
      streaming = false;
    }
  }

  const kinds = $derived(
    new Set<EventKind>(EVENT_KINDS.map((kind) => kind.id).filter((kind) => !off.includes(kind))),
  );

  // With the stream loaded, the summary's few hits and heals (the ones before a death) give way
  // to every one of them; the casts, auras and deaths stay the summary's.
  const all = $derived(
    [
      ...summaryEvents(summary, names).filter(
        (event) => stream === null || (event.kind !== 'damage' && event.kind !== 'heal'),
      ),
      ...(stream === null ? [] : streamEvents(stream)),
    ]
      .filter((event) => event.guids.some(inScope))
      .sort((a, b) => a.atMs - b.atMs),
  );
  const matching = $derived(filterEvents(all, kinds, search));
  /** How many rows are on the page: a wipe's stream runs to thousands. */
  const PAGE = 200;
  let limit = $state(PAGE);
  const shown = $derived(matching.slice(0, limit));
  const unshown = $derived(matching.length - shown.length);
  $effect(() => {
    // Back to the first page whenever the filter changes.
    void [kinds.size, search];
    limit = PAGE;
  });

  const streamedCount = $derived(
    all.filter((event) => event.kind === 'damage' || event.kind === 'heal').length,
  );

  const COLUMNS: readonly StatColumn[] = [
    { id: 'time', label: 'Time' },
    { id: 'amount', label: 'Amount' },
  ];
  const DASH: StatFigure = { value: '–', muted: true };

  function figuresOf(event: SummaryEvent): StatFigure[] {
    return [
      { value: formatDuration(event.atMs) },
      event.amount === undefined ? DASH : { value: formatAmount(event.amount) },
    ];
  }

  /** A player's name in their class colour; an enemy's in the line's own colour. */
  const nameColor = (guid: string): string | undefined =>
    classOf.has(guid) ? classColorVar(classOf.get(guid)) : undefined;

  function onFind(event: Event): void {
    const input = event.target;
    if (input instanceof HTMLInputElement) onPatch({ find: input.value });
  }
</script>

<Section
  title="Events"
  meta={`${matching.length.toLocaleString()} events in this window, in order`}
  testid="events-view"
  panel
>
  <div class="flex flex-col gap-3">
    <div class="flex flex-wrap items-center gap-x-6 gap-y-2">
      {#if chips}<EventChips {off} {onPatch} />{/if}
      <!-- The field reports each keystroke to the url; its own value follows the url back. -->
      <div class="w-full md:ml-auto md:w-[240px]" oninput={onFind}>
        <Field
          id="events-search"
          label="Find"
          labelHidden
          type="search"
          placeholder="Find a name, spell or amount"
          value={search}
          testid="events-search"
        />
      </div>
    </div>

    <div>
      {#if shown.length === 0}
        <p class="text-muted px-4 py-4 text-[14px]" data-testid="table-empty">
          No rows matched.{stream === null && search.trim() !== ''
            ? ' Load every hit and heal below to search them all.'
            : ''}
        </p>
      {:else}
        <StatTable columns={COLUMNS} nameHead="Event" bar={false} label="Events" testid="event-list">
          {#each shown as event, index (`${event.atMs}-${index}`)}
            <StatRow figures={figuresOf(event)} bar={false}>
              {#snippet icon()}{#if event.spellId !== undefined}<SpellIcon
                    url={icons.of(event.spellId, event.spellName)}
                    spellId={event.spellId}
                    size={24}
                    tip
                  />{/if}{/snippet}
              {#snippet name()}
                <!-- Wraps on a phone: the part a reader wants is the end of the line, "on Hobolol
                     by Deadclasslol", and a truncated line cuts exactly that. -->
                <span class="text-text text-[14px] font-normal"
                  >{#each sentenceParts(event) as part, partIndex (partIndex)}{#if part.guid !== undefined && nameColor(part.guid) !== undefined}<span
                        class="font-semibold"
                        style:color={nameColor(part.guid)}>{part.text}</span
                      >{:else}{part.text}{/if}{/each}</span
                >
              {/snippet}
            </StatRow>
          {/each}
        </StatTable>
        {#if unshown > 0}
          <div class="px-4 py-3">
            <Button size="sm" testid="events-more" onclick={() => (limit += PAGE)}
              >Show {Math.min(PAGE, unshown)} more · {unshown.toLocaleString()} not yet shown</Button
            >
          </div>
        {/if}
      {/if}
    </div>

    <p class="text-muted text-[13px] leading-5">
      {#if stream === null}
        From the summary: casts, auras going up and down, the hits and heals before each death, and the
        deaths. An aura applied while the same aura is already up on that player is folded into the running
        one, so its applier is named on the Buffs table and not here.{#if loadStream !== undefined}
          <button
            type="button"
            class={`${LINK} inline-flex min-h-11 items-center md:min-h-0 ${streaming ? BUSY_CLASS : ''}`}
            data-testid="events-stream"
            disabled={streaming}
            aria-busy={streaming}
            onclick={() => void runStream()}>Load every hit and heal in this window</button
          >
          from the fight’s events.{/if}
      {:else}
        Casts, auras and deaths from the summary, and <span class="tabular-nums">{streamedCount}</span>
        hits and heals from the fight’s events{stream.length === all.length ? '' : ' in scope'}.
      {/if}
      {#if streamError !== ''}<span class="text-wipe" role="alert">{streamError}</span>{/if}
      Every field of every line is in Queries.
    </p>
  </div>
</Section>
