<!-- web/src/components/home/HomeGuildSims.svelte -->
<!-- Guild sims (round 5: full width under the feed, two columns from 700): the guild's ladder for
     the viewer's role, band and preset, the viewer's own row pinned last when it is below the cut,
     then the spec's tier line ("Tier A · #2 of 20 · Tier list"). Boxed: the names are class colours. -->
<script lang="ts">
  import type { GuildSimRow, GuildSims, SpecWeek } from '../../lib/home/home-answer';
  import {
    GUILD_SIMS_SHOWN,
    characterLink,
    figureText,
    guildSimsMeta,
    homeBCopy,
    homeBHrefs,
  } from '../../lib/home/home-b-view';
  import { tierLine } from '../../lib/home/plaque';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import { LINK } from '../../lib/ui/primitives';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { sims, week, guildHref }: { sims: GuildSims; week: SpecWeek | null; guildHref: string | null } =
    $props();

  const CREST_PX = 28;
  const ladder = $derived(sims.rows.slice(0, GUILD_SIMS_SHOWN));
  const pinned = $derived(ladder.some((row) => row.you) ? null : sims.you);
  const tier = $derived(tierLine(week));
</script>

{#snippet simRow(row: GuildSimRow)}
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
    {#snippet trail()}<Figure
        value={figureText(row.value)}
        unit={sims.unit}
        size="sm"
        align="end"
      />{/snippet}
  </Row>
{/snippet}

<Section
  title={homeBCopy.guildSims}
  meta={guildSimsMeta(sims)}
  action={guildHref === null ? undefined : { label: homeBCopy.guildPage, href: guildHref }}
  boxed
  testid="home-guild-sims"
>
  <div class="home-sims-ladder">
    {#each ladder as row (row.key)}{@render simRow(row)}{/each}
    {#if pinned !== null}{@render simRow(pinned)}{/if}
  </div>
  {#if pinned === null && !ladder.some((row) => row.you)}
    <p class="home-sims-line">
      {homeBCopy.youNoSim} · <a class={LINK} href={homeBHrefs.sim}>{homeBCopy.runSim} ›</a>
    </p>
  {/if}
  {#if tier !== null}
    <p class="home-sims-line" data-testid="home-tier-line">
      {tier} · <a class={LINK} href={homeBHrefs.tiers}>{homeBCopy.tierList} ›</a>
    </p>
  {/if}
</Section>

<style>
  .home-sims-ladder > :global(*) {
    break-inside: avoid;
  }
  .home-sims-line {
    padding: 12px 0 8px;
    font-size: 15px;
    line-height: 22px;
    color: var(--color-text);
  }
  @media (min-width: 700px) {
    .home-sims-ladder {
      columns: 2;
      column-gap: 48px;
    }
  }
</style>
