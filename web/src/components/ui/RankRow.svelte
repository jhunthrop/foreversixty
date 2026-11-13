<!-- web/src/components/ui/RankRow.svelte -->
<!-- One ranked row, a Row preset (tiers spec 2026-11-06 §4.4): rank, a crest, the name as a link in
     the class colour with a sub line, a bar scaled to the leader in the class colour (a tie note
     under it), the figures, the gap and a strip of next-step links. The row itself is no link: its
     next steps are links inside it (the Row rule, a link or an action, never both). Hover is the S2
     ground and the row's slice of the rail in the class colour; `selected` the gold ground and a
     gold slice (in a `compact` RankList, the accent tint and a cyan slice). Below 1024 it is the four-area phone grid. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { titleColor } from '../../lib/ui/primitives';
  import type { RankFigure, RankGap, RankLink, RankNote } from '../../lib/ui/rank-list';
  import Crest from './Crest.svelte';
  import Meter from './Meter.svelte';
  import './rank-list.css';

  let {
    rank = '',
    rankColor,
    label,
    name,
    href,
    classSlug,
    sub,
    note,
    tie,
    fraction,
    barTone = 'class',
    figures,
    gap,
    links = [],
    selected = false,
    key,
    tag,
    nameTestid,
    testid,
  }: {
    /** "3", "=2", or empty on a list that claims no order. */
    rank?: string;
    /** The rank's own colour (a log ranking's percentile on the parse scale); muted without. */
    rankColor?: string;
    /** The row's accessible name when it stands apart (a pinned "Your best kill"). */
    label?: string;
    name: string;
    href: string;
    /** The crest, the name's and the bar's colour; without one no crest, and the name and bar are
     *  the strong text and gold. */
    classSlug?: string;
    sub?: string;
    /** A third line under the name ("Stat weights less certain"). */
    note?: RankNote;
    /** A line under the bar ("≈ tie with Arms"). */
    tie?: RankNote;
    /** The bar, 0 to 1. */
    fraction: number;
    /** The bar's fill: the class colour (a log ranking's Details! bar) or the cyan accent (the tier
     *  list's relative bar, round 5). */
    barTone?: 'class' | 'accent';
    /** The sorted figure first. */
    figures: readonly RankFigure[];
    gap: RankGap;
    links?: readonly RankLink[];
    selected?: boolean;
    /** `data-rank-key`: what a script finds the row by to mark it. */
    key?: string;
    /** After the name (a "Your spec" Tag). */
    tag?: Snippet;
    nameTestid?: string;
    testid?: string;
    /** Never rendered: an Astro caller's whitespace between named slots arrives as children. */
    children?: Snippet;
  } = $props();

  const tone = $derived(classSlug === undefined ? 'var(--color-gold)' : titleColor({ classSlug }));
</script>

{#snippet figure(item: RankFigure, first: boolean)}
  <span
    class={['ui-rr-fig', item.phoneOnly && 'md:hidden']}
    data-first={first ? 'true' : undefined}
    data-best={item.best ? 'true' : undefined}
    title={item.best ? item.bestTitle : undefined}
    >{#if !item.bare}<span class="ui-rr-fig-label" aria-hidden="true">{item.label}</span>{/if}<span
      class="sr-only">{`${item.label}: `}</span
    >{#if item.href}<a
        class="ui-rr-fig-value ui-rr-fig-link"
        href={item.href}
        title={item.title}
        aria-label={item.ariaLabel}
        style:color={item.color}
        data-testid={item.testid}>{item.value}</a
      >{:else}<span
        class="ui-rr-fig-value"
        title={item.title}
        style:color={item.color}
        data-testid={item.testid}>{item.value}</span
      >{/if}{#if item.unit}<small>{item.unit}</small>{/if}</span
  >
{/snippet}

<li
  class="ui-rank-row ui-rank-cols"
  style:--rr-tone={tone}
  aria-current={selected ? 'true' : undefined}
  aria-label={label}
  data-rank-key={key}
  data-testid={testid}
>
  <span class="ui-rr-rank" style:color={rankColor} data-testid="rank-rank">{rank}</span>
  <span class="ui-rr-lead"
    >{#if classSlug !== undefined}<Crest cls={classSlug} size={36} class="ui-rr-crest" />{/if}</span
  >
  <span class="ui-rr-name">
    <span class="ui-rr-title"
      ><a class="ui-rr-link" {href} data-testid={nameTestid ?? 'rank-name'}>{name}</a
      >{#if tag}{@render tag()}{/if}</span
    >
    <!-- One line: a guild and spec that do not fit end in an ellipsis, the whole line on hover. -->
    {#if sub}<span class="ui-rr-sub block truncate" title={sub}>{sub}</span>{/if}
    {#if note}<span class="ui-rr-note" title={note.title} data-testid="rank-note">{note.text}</span>{/if}
  </span>
  <span class="ui-rr-bar">
    {#if barTone === 'accent'}<Meter value={fraction} />{:else}<Meter
        value={fraction}
        tone="class"
        color={tone}
      />{/if}
    {#if tie}<span class="ui-rr-tie" title={tie.title} data-testid="rank-tie">{tie.text}</span>{/if}
  </span>
  {#if figures[0]}{@render figure(figures[0], true)}{/if}
  <span class="ui-rr-extra"
    >{#each figures.slice(1) as item (item.label)}{@render figure(item, false)}{/each}</span
  >
  <span class="ui-rr-gap" data-leader={gap.leader ? 'true' : undefined} title={gap.title}
    ><span class="sr-only">{`${gap.label}: `}</span><span data-testid="rank-gap">{gap.text}</span></span
  >
  {#if links.length > 0}
    <span class="ui-rr-links">
      {#each links as link (link.href)}
        <a
          href={link.href}
          data-primary={link.primary ? 'true' : undefined}
          aria-label={link.ariaLabel}
          title={link.title}
          data-testid={link.testid}>{link.label}</a
        >
      {/each}
    </span>
  {/if}
</li>
