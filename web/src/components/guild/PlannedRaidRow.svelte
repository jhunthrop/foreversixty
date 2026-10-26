<!-- web/src/components/guild/PlannedRaidRow.svelte -->
<!-- One planned raid on the Raids tab (raids spec 3.1): one link to the raid page. Desktop is the
     six-cell grid (date block, title and zone, when, count, chips, chevron); below md the date
     block leads a three-line stack. `now` comes from the tab's one minute tick, so rows run no
     timer of their own. Gold shows once a row: the countdown, or the Not signed up chip. -->
<script lang="ts">
  import type { GuildRaidView } from '../../lib/guild/api';
  import { plannedRowView } from '../../lib/guild/planned-row';
  import RaidChip from './RaidChip.svelte';

  let {
    raid,
    href,
    now,
    viewerTimezone,
    locale = 'en-US',
  }: { raid: GuildRaidView; href: string; now: Date; viewerTimezone: string; locale?: string } = $props();

  const view = $derived(plannedRowView(raid, now, viewerTimezone, locale));
  /** The Not signed up chip is the row's gold mark, so the countdown steps down to muted. */
  const countdownTone = $derived(
    view.viewer !== null && view.viewer.state === null ? 'text-muted' : 'text-gold',
  );
  const titleTone = $derived(view.cancelled ? 'text-muted' : 'text-strong');
  const countTone = $derived(view.cancelled ? 'text-muted' : 'text-strong');
  const testid = $derived(`planned-raid-${raid.id}`);
  /** As an expression: Svelte trims the edge spaces of literal text inside a span. */
  const SEP = ' · ';
</script>

{#snippet dateBlock(box: string)}
  <span
    class="border-line rounded-control flex flex-col items-center justify-center border {box}"
    aria-hidden="true"
  >
    <span class="label text-muted leading-none">{view.weekday}</span>
    <span class="text-strong font-mono text-[22px] leading-tight tabular-nums">{view.day}</span>
    <span class="label text-muted leading-none">{view.month}</span>
  </span>
{/snippet}

<!-- The status chip leads the viewer chip, the mock's order (`Cancelled`, then `In · Tank`). -->
{#snippet chips(layout: string)}
  {#if view.status !== null}
    <RaidChip kind="status" status={view.status} testid={`${testid}-${layout}-status-chip`} />
  {/if}
  {#if view.viewer !== null}
    <RaidChip
      kind="viewer"
      state={view.viewer.state}
      role={view.viewer.role}
      testid={`${testid}-${layout}-viewer-chip`}
    />
  {/if}
{/snippet}

{#snippet bar(width: string)}
  {#if view.fill !== null}
    <span class="bg-line-soft block h-2 overflow-hidden rounded-full {width}" aria-hidden="true">
      <span class="bg-gold block h-full" style:width={`${view.fill}%`}></span>
    </span>
  {/if}
{/snippet}

<a
  {href}
  class="group border-line-soft hover:bg-card-top focus-visible:outline-gold block border-b transition-colors duration-[120ms] focus-visible:outline-2 focus-visible:-outline-offset-2 active:bg-[#171d2b]"
  data-testid={testid}
>
  <span
    class="hidden min-h-16 grid-cols-[56px_minmax(0,1.5fr)_minmax(0,1fr)_minmax(0,1fr)_176px_16px] items-center gap-4 px-3.5 py-[7px] md:grid"
    data-testid={`${testid}-desktop`}
  >
    {@render dateBlock('h-14 w-14')}
    <span class="min-w-0">
      <span
        class="block truncate text-[15px] font-semibold underline-offset-4 pointer-fine:group-hover:underline {titleTone}"
        title={view.title}>{view.title}</span
      >
      {#if view.zoneLine !== null}<span class="text-muted block truncate text-[13px]">{view.zoneLine}</span
        >{/if}
    </span>
    <span class="min-w-0">
      <span class="sr-only">{`${view.fullDate}, `}</span><span
        class="text-text block font-mono text-[14px] tabular-nums">{view.time}</span
      >
      {#if view.guildTime !== null}<span class="text-muted block text-[12px]">{view.guildTime}</span>{/if}
      <span class="block font-mono text-[12px] tabular-nums {countdownTone}">{view.countdown}</span>
    </span>
    <span class="flex min-w-0 flex-col gap-1">
      <span class="font-mono text-[15px] tabular-nums {countTone}">{view.count}</span>
      {@render bar('w-24')}
      <span class="text-muted truncate text-[12px]">{view.breakdown}</span>
    </span>
    <span class="flex flex-wrap items-center justify-end gap-1.5">{@render chips('desktop')}</span>
    <span class="text-muted group-hover:text-gold text-[16px]" aria-hidden="true">›</span>
  </span>

  <span
    class="grid min-h-24 grid-cols-[48px_minmax(0,1fr)] items-center gap-3 px-3.5 py-3 md:hidden"
    data-testid={`${testid}-phone`}
  >
    {@render dateBlock('h-[52px] w-12')}
    <span class="flex min-w-0 flex-col gap-1">
      <span class="flex min-w-0 items-center justify-between gap-2">
        <span class="truncate text-[15px] font-semibold {titleTone}" title={view.title}>{view.title}</span>
        <span class="flex shrink-0 items-center gap-1.5">{@render chips('phone')}</span>
      </span>
      <span class="text-[12px]">
        <span class="sr-only">{`${view.fullDate}, `}</span><span class="text-text font-mono tabular-nums"
          >{view.time}</span
        >{#if view.guildTime !== null}<span class="text-muted">{SEP}{view.guildTime}</span>{/if}<span
          class="font-mono tabular-nums {countdownTone}">{SEP}{view.countdown}</span
        >
      </span>
      <span class="flex min-w-0 items-center gap-2 text-[12px]">
        <span class="font-mono text-[13px] tabular-nums {countTone}">{view.count}</span>
        {@render bar('w-16 shrink-0')}
        <span class="text-muted truncate">{view.breakdown}</span>
      </span>
    </span>
  </span>
</a>
