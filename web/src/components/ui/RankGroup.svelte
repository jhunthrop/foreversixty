<!-- web/src/components/ui/RankGroup.svelte -->
<!-- One letter group of a ranked list (tiers spec 2026-11-06 §4.4): from 1100 a 184 px label column
     (the letter in its tone over the band and the count) with a 3 px rail in the tone at the list's
     x 0; below 1100 a header strip above its rows. It takes the plate's ground: no tint, no surface. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import './rank-list.css';

  let {
    letter,
    tone,
    band,
    count,
    label,
    title,
    testid,
    children,
  }: {
    letter: string;
    /** The letter's colour and its rail's (a quality colour token); a compact list's letter is cream. */
    tone?: string;
    /** "Within 5% of the top"; a compact list's letter column carries the letter alone. */
    band?: string;
    /** "8 specs". */
    count?: string;
    /** The group's accessible name ("Tier B"). */
    label: string;
    /** The letter's hover title ("10 to 20% behind #1"), where the column carries the letter alone. */
    title?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<li class="ui-rank-group" style:--rg-tone={tone} data-group={letter} data-testid={testid}>
  <div class="ui-rank-group-head" data-testid="rank-group-head">
    <span class="ui-rg-letter" {title} data-testid="rank-group-letter"
      ><span class="sr-only">{`${label}: `}</span><span aria-hidden="true">{letter}</span></span
    >
    {#if band !== undefined}<span class="ui-rg-band">{band}</span>{/if}
    {#if count !== undefined}<span class="ui-rg-count">{count}</span>{/if}
  </div>
  <ol class="ui-rank-group-rows" aria-label={label}>{@render children()}</ol>
</li>
