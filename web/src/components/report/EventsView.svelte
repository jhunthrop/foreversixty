<!-- web/src/components/report/EventsView.svelte -->
<!-- Every cast, aura application and removal, death and killing hit the summary itself
     timestamps, merged and time-ordered (src/lib/report/events.ts). The complete event
     stream -- every field of every line, not just what the summary keeps -- is
     events.parquet, which is the Queries view's job: two to ten megabytes and a DuckDB
     instance for a question this view answers for free.
     Drawn to round 5 (design/mocks/directions/report_tail.py, b1-report-events): the panel
     `Events` with its count, the Show chips and Find on one row under the head, a StatTable of
     lines (the spell's icon, the sentence with each player's name in their class colour, Time,
     Amount), 100 a page, and one foot line that loads every hit and heal. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import activeBuild from '../../data/active-build.json';
  import { isUnnamed, loadSharedSpellNames, nameAuras, type SpellNames } from '../../lib/report/aura-names';
  import { classColorVar } from '../../lib/report/format';
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
  import { LINK, RULE_TOP } from '../../lib/ui/primitives';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import { clockDuration } from '../../lib/format/duration';
  import { compact, grouped } from '../../lib/format/number';
  import Button from '../ui/Button.svelte';
  import Spinner from '../ui/Spinner.svelte';
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
  } = $props();

  const icons = new SpellIcons();

  /** The build's spell names, for an aura up at the pull that no line of the log named (B2). */
  let spellNames = $state<SpellNames>({});
  $effect(() => {
    void loadSharedSpellNames(activeBuild.build).then((loaded) => {
      spellNames = loaded;
    });
  });
  /** The summary with its auras named as the Buffs tab names them; one no table knows is dropped. */
  const named = $derived<Summary>({
    ...summary,
    auras: nameAuras(summary.auras, spellNames).filter((track) => !isUnnamed(track.name)),
  });

  let stream = $state<StreamLine[] | null>(null);
  let streaming = $state(false);
  let streamError = $state('');
  // A new window is a new stream: the loaded one answered the old question.
  $effect(() => {
    void summary;
    stream = null;
    streamError = '';
  });
  // The hits and heals are what a log is read for: on a single pull they load as the tab opens,
  // once per window, and the button stays for a retry after a failed load.
  let asked: Summary | null = null;
  $effect(() => {
    if (loadStream === undefined || stream !== null || streaming || asked === summary) return;
    asked = summary;
    void runStream();
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
      ...summaryEvents(named, names).filter(
        (event) => stream === null || (event.kind !== 'damage' && event.kind !== 'heal'),
      ),
      ...(stream === null ? [] : streamEvents(stream)),
    ]
      .filter((event) => event.guids.some(inScope))
      .sort((a, b) => a.atMs - b.atMs),
  );
  const matching = $derived(filterEvents(all, kinds, search));
  /** How many rows are on the page: a wipe's stream runs to thousands. */
  const PAGE = 50;
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

  /** Amount only while a line on the page carries one: a column of dashes says nothing. */
  const withAmounts = $derived(shown.some((event) => event.amount !== undefined));
  const columns = $derived<StatColumn[]>([
    { id: 'time', label: 'Time' },
    ...(withAmounts ? [{ id: 'amount', label: 'Amount' }] : []),
  ]);
  const DASH: StatFigure = { value: '–', muted: true };

  function figuresOf(event: SummaryEvent): StatFigure[] {
    const time = { value: clockDuration(event.atMs) };
    if (!withAmounts) return [time];
    return [time, event.amount === undefined ? DASH : { value: compact(event.amount) }];
  }

  /** A player's name in their class colour; an enemy's in the line's own colour. */
  const nameColor = (guid: string): string | undefined =>
    classOf.has(guid) ? classColorVar(classOf.get(guid)) : undefined;

  function onFind(event: Event): void {
    const input = event.target;
    if (input instanceof HTMLInputElement) onPatch({ find: input.value });
  }
</script>

<Section title="Events" meta={`${grouped(matching.length)} lines`} testid="events-view" panel flush>
  <div class="flex flex-col">
    <div class="flex flex-wrap items-center gap-x-4 gap-y-2 px-4 pb-3 md:px-5">
      <EventChips {off} {onPatch} />
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
        <EmptyState inset message="No rows matched." testid="table-empty" />
      {:else}
        <StatTable {columns} nameHead="Event" bar={false} dense label="Events" testid="event-list">
          {#each shown as event, index (`${event.atMs}-${index}`)}
            <StatRow figures={figuresOf(event)} bar={false}>
              {#snippet icon()}{#if event.spellId !== undefined}<SpellIcon
                    socket
                    url={icons.of(event.spellId, event.spellName)}
                    size={20}
                    tip={event.spellId}
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
          <div class="text-muted flex items-center gap-[14px] px-4 py-3 text-[13px] md:px-5">
            <Button size="sm" testid="events-more" onclick={() => (limit += PAGE)}
              >Show {Math.min(PAGE, unshown)} more</Button
            ><span class="tabular-nums">{grouped(unshown)} more</span>
          </div>
        {/if}
      {/if}
    </div>

    <p class="text-muted {RULE_TOP} px-4 pt-[10px] pb-[14px] text-[13px] leading-5 md:px-5">
      {#if stream === null}
        Hits and heals: only those before a death.{#if loadStream !== undefined}<button
            type="button"
            class={`${LINK} ml-1 inline-flex min-h-11 items-center gap-2 md:min-h-0`}
            data-testid="events-stream"
            disabled={streaming}
            aria-busy={streaming}
            onclick={() => void runStream()}
            >{#if streaming}<Spinner />Loading every hit and heal…{:else}Load every hit and heal{/if}</button
          >{/if}
      {:else}
        Every hit and heal in the window: <span class="tabular-nums">{grouped(streamedCount)}</span>.
      {/if}
      {#if streamError !== ''}<span class="text-wipe" role="alert">{streamError}</span>{/if}
    </p>
  </div>
</Section>
