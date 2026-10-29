<!-- web/src/components/guild/PlannedRaidRow.svelte -->
<!-- One planned raid on the Raids tab (raids spec 3.1): one link to the raid page. Desktop is the
     six-cell grid (date block, title and zone, when, count, chips, chevron); below md the date
     block leads a three-line stack. `now` comes from the tab's one minute tick, so rows run no
     timer of their own. Gold shows once a row: the countdown, or the Not signed up chip. -->
<script lang="ts">
  import { viewerLocale } from '../../lib/guild/raid-time';
  import type { GuildRaidView } from '../../lib/guild/api';
  import { plannedRowView } from '../../lib/guild/planned-row';
  import { SEP } from '../../lib/guild/raid-copy';
  import RaidChip from './RaidChip.svelte';

  let {
    raid,
    href,
    now,
    viewerTimezone,
    locale = viewerLocale(),
  }: { raid: GuildRaidView; href: string; now: Date; viewerTimezone: string; locale?: string } = $props();

  const view = $derived(plannedRowView(raid, now, viewerTimezone, locale));
  /** The Not signed up chip is the row's gold mark, so the countdown steps down to muted; a
   *  cancelled or started raid has no gold at all. */
  const countdownTone = $derived(
    view.quiet || (view.viewer !== null && view.viewer.state === null) ? 'text-muted' : 'text-gold',
  );
  const fillTone = $derived(view.quiet ? 'bg-muted' : 'bg-gold');
  const titleTone = $derived(view.cancelled ? 'text-muted' : 'text-strong');
  const countTone = $derived(view.cancelled ? 'text-muted' : 'text-strong');
  const testid = $derived(`planned-raid-${raid.id}`);
</script>

{#snippet dateBlock(box: string)}
  <span
    class="border-line rounded-control flex flex-col items-center justify-center border {box}"
    aria-hidden="true"
  >
    <span class="label text-muted leading-none">{view.weekday}</span>
    <span class="text-strong numeric text-[22px] leading-tight">{view.day}</span>
    <span class="label text-muted leading-none">{view.month}</span>
  </span>
{/snippet}

<!-- The status chip leads the viewer chip (spec 3.1). Phone line 1 holds one chip: the status
     when there is one, else the viewer's. -->
{#snippet chips(layout: 'desktop' | 'phone')}
  {#if view.status !== null}
    <RaidChip kind="status" status={view.status} testid={`${testid}-${layout}-status-chip`} />
  {/if}
  {#if view.viewer !== null && (layout === 'desktop' || view.status === null)}
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
      <span class="{fillTone} block h-full" style:width={`${view.fill}%`}></span>
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
      <span class="sr-only">{`${view.fullDate}, `}</span><span class="text-text numeric block text-[14px]"
        >{view.time}</span
      >
      {#if view.guildTime !== null}<span class="text-muted block text-[12px]">{view.guildTime}</span>{/if}
      <span class="numeric block text-[12px] {countdownTone}">{view.countdown}</span>
    </span>
    <span class="flex min-w-0 flex-col gap-1">
      <span class="flex items-center gap-2" data-testid={`${testid}-desktop-count`}
        ><span class="numeric text-[15px] {countTone}">{view.count}</span>{@render bar('w-24 shrink-0')}</span
      >
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
      <span class="flex min-w-0 items-start justify-between gap-2">
        <span class="line-clamp-2 text-[15px] font-semibold {titleTone}" title={view.title}>{view.title}</span
        >
        <span class="flex shrink-0 items-center gap-1.5">{@render chips('phone')}</span>
      </span>
      <span class="text-[12px]">
        <span class="sr-only">{`${view.fullDate}, `}</span><span class="text-text numeric">{view.time}</span
        >{#if view.guildTime !== null}<span class="text-muted">{SEP}{view.guildTime}</span>{/if}<span
          class="numeric {countdownTone}">{SEP}{view.countdown}</span
        >
      </span>
      <span class="flex min-w-0 items-center gap-2">
        <span class="numeric text-[13px] {countTone}">{view.count}</span>
        {@render bar('w-16 shrink-0')}
      </span>
      <span class="text-muted block text-[12px] leading-4" data-testid={`${testid}-phone-breakdown`}
        >{view.breakdown}</span
      >
    </span>
  </span>
</a>
