<!-- web/src/components/home/HomeSpecWeek.svelte -->
<!-- Your spec this week: the spec's grade, its move since last week and its place, with the tier list
     at the right; under it the spec's board this week (/v1/home movers): the biggest climbers once a
     weekly board exists, else the top three, each to their character page. Full width under the
     feed, or beside Guild sims when the guild has a ladder. Boxed, as every row set. -->
<script lang="ts">
  import type { Movers, SpecWeek } from '../../lib/home/home-answer';
  import { characterLink, homeBCopy, homeBHrefs, specWeekView } from '../../lib/home/home-b-view';
  import { factsCopy, moverDelta, moversTitle } from '../../lib/home/home-facts';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Figure from '../ui/Figure.svelte';
  import GradeTile from '../ui/GradeTile.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { week, movers }: { week: SpecWeek; movers: Movers | null } = $props();

  const CREST_PX = 28;
  const view = $derived(specWeekView(week));
  const board = $derived(movers !== null && movers.rows.length > 0 ? movers : null);
</script>

<Section
  title={homeBCopy.specWeek}
  action={{ label: homeBCopy.tierList, href: homeBHrefs.tiers }}
  boxed
  testid="home-spec-week"
>
  <div>
    <Row title={view.title} sub={view.sub} open>
      {#snippet lead()}
        {#if week.grade}<GradeTile letter={week.grade} bare testid="home-grade" />{/if}
      {/snippet}
    </Row>
    {#if board !== null}
      <div data-testid="home-movers">
        <h3 class="home-spec-board">{moversTitle(board)}</h3>
        {#each board.rows as row (row.key)}
          {@const delta = moverDelta(row)}
          <Row
            title={row.name}
            titleTone={{ classSlug: row.class as ClassSlug }}
            rank={String(row.rank)}
            href={characterLink(row.key, row.name)}
            density="slim"
            open
            testid="home-mover"
          >
            {#snippet lead()}<Crest cls={row.class} size={CREST_PX} />{/snippet}
            {#snippet trail()}
              {#if delta !== null}<Figure
                  value={delta}
                  unit={factsCopy.places}
                  tone="kill"
                  size="sm"
                  align="end"
                />{/if}
            {/snippet}
          </Row>
        {/each}
      </div>
    {/if}
  </div>
</Section>

<style>
  .home-spec-board {
    padding: 8px 0 2px;
    font-size: 13px;
    line-height: 18px;
    font-weight: 600;
    color: var(--color-muted);
  }
</style>
