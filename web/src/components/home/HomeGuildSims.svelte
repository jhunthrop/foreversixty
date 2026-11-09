<!-- web/src/components/home/HomeGuildSims.svelte -->
<!-- Guild sims (round 5: under the feed, two columns from 900 px of its own width): the guild's ladder for
     the viewer's role, band and preset, the viewer's own row pinned last when it is below the cut.
     Each row carries a bar in its class colour against the top of the ladder and, beside the figure,
     its gap to your sim (ahead or behind). Boxed: the names are class colours. -->
<script lang="ts">
  import type { GuildSimRow, GuildSims } from '../../lib/home/home-answer';
  import {
    GUILD_SIMS_SHOWN,
    characterLink,
    figureText,
    homeBCopy,
    homeBHrefs,
  } from '../../lib/home/home-b-view';
  import { simGap, simShare } from '../../lib/home/home-facts';
  import { titleColor, type ClassSlug } from '../../lib/ui/primitives';
  import { LINK } from '../../lib/ui/primitives';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { sims }: { sims: GuildSims } = $props();

  const CREST_PX = 28;
  const ladder = $derived(sims.rows.slice(0, GUILD_SIMS_SHOWN));
  const youInLadder = $derived(ladder.some((row) => row.you));
  const pinned = $derived(youInLadder ? null : sims.you);
  const you = $derived(sims.rows.find((row) => row.you) ?? sims.you);
  const top = $derived(sims.rows[0]?.value ?? 0);
</script>

{#snippet simRow(row: GuildSimRow)}
  {@const gap = simGap(row, you)}
  <Row
    title={row.name}
    titleTone={{ classSlug: row.class as ClassSlug }}
    rank={String(row.rank)}
    href={characterLink(row.key, row.name)}
    density="slim"
    selected={row.you}
    open
    testid="home-guild-sim"
  >
    {#snippet lead()}<ClassCrestRing characterClass={row.class} size={CREST_PX} />{/snippet}
    {#snippet trail()}
      <span class="home-sims-trail">
        <span class="home-sims-bar"
          ><Meter
            value={simShare(row, top)}
            tone="class"
            color={titleColor({ classSlug: row.class as ClassSlug })}
          /></span
        >
        <span class="home-sims-gap" data-testid="home-guild-sim-gap">{gap ?? ''}</span>
        <span class="home-sims-figure"
          ><Figure value={figureText(row.value)} unit={sims.unit} size="sm" align="end" /></span
        >
      </span>
    {/snippet}
  </Row>
{/snippet}

<Section title={homeBCopy.guildSims} boxed testid="home-guild-sims">
  <div class="home-sims-box">
    <div class="home-sims-ladder">
      {#each ladder as row (row.key)}{@render simRow(row)}{/each}
      {#if pinned !== null}{@render simRow(pinned)}{/if}
    </div>
  </div>
  {#if pinned === null && !youInLadder}
    <p class="home-sims-line">
      {homeBCopy.youNoSim} · <a class={LINK} href={homeBHrefs.sim}>{homeBCopy.runSim} ›</a>
    </p>
  {/if}
</Section>

<style>
  .home-sims-box {
    container-type: inline-size;
  }
  .home-sims-ladder > :global(*) {
    break-inside: avoid;
  }
  .home-sims-trail {
    display: flex;
    align-items: center;
    gap: 14px;
  }
  .home-sims-bar {
    display: block;
    width: 120px;
  }
  .home-sims-gap {
    width: 96px;
    font-size: 13px;
    line-height: 18px;
    text-align: right;
    color: var(--color-muted);
    font-variant-numeric: tabular-nums;
  }
  .home-sims-figure {
    display: flex;
    justify-content: flex-end;
    min-width: 104px;
  }
  .home-sims-line {
    padding: 12px 0 8px;
    font-size: 15px;
    line-height: 22px;
    color: var(--color-text);
  }
  @container (min-width: 900px) {
    .home-sims-ladder {
      columns: 2;
      column-gap: 48px;
    }
  }
  /* Two narrow columns: the bar shortens and the gap gives way to it. */
  @media (max-width: 1099px) {
    .home-sims-bar {
      width: 56px;
    }
    .home-sims-gap {
      display: none;
    }
  }
  /* A phone keeps the name and the figure; the bar would sit on a long name. */
  @media (max-width: 767px) {
    .home-sims-bar {
      display: none;
    }
  }
</style>
