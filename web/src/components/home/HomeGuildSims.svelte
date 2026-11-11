<!-- web/src/components/home/HomeGuildSims.svelte -->
<!-- Guild sims (round 5: under the feed, one column, a ladder read top to bottom): the guild's ladder for
     the viewer's role, band and preset, the viewer's own row pinned last when it is below the cut.
     Each row carries a bar in its class colour against the top of the ladder (as wide as the panel
     allows) and, under the name, its gap to your sim (`12.8 DPS ahead of you`). Boxed: the names are class colours. -->
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
  import Crest from '../ui/Crest.svelte';
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
  {@const gap = simGap(row, you, sims.unit)}
  <Row
    title={row.name}
    sub={gap ?? undefined}
    titleTone={{ classSlug: row.class as ClassSlug }}
    rank={String(row.rank)}
    href={characterLink(row.key, row.name)}
    density="slim"
    selected={row.you}
    open
    testid="home-guild-sim"
  >
    {#snippet lead()}<Crest cls={row.class} size={CREST_PX} />{/snippet}
    {#snippet trail()}
      <span class="home-sims-trail">
        <span class="home-sims-bar"
          ><Meter
            value={simShare(row, top)}
            tone="class"
            color={titleColor({ classSlug: row.class as ClassSlug })}
          /></span
        >
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
  /* The gap under the name stays on one line. */
  .home-sims-ladder :global([data-testid='home-guild-sim'] .text-muted) {
    white-space: nowrap;
  }
  .home-sims-trail {
    display: flex;
    align-items: center;
    gap: 14px;
  }
  /* The bar grows with the panel, so a wide screen reads the ladder rather than empty space. */
  .home-sims-bar {
    display: block;
    width: clamp(56px, 36cqi, 360px);
  }
  /* A phone panel: the name, its gap and the figure keep the row; the bar gives way. */
  @container (max-width: 519px) {
    .home-sims-bar {
      display: none;
    }
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
</style>
