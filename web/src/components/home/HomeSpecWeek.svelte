<!-- web/src/components/home/HomeSpecWeek.svelte -->
<!-- Your spec this week, for a character with no guild ladder to show: the spec's grade, its move
     since last week and its place, with the tier list at the right. Boxed, as every row set. -->
<script lang="ts">
  import type { SpecWeek } from '../../lib/home/home-answer';
  import { homeBCopy, homeBHrefs, specWeekView } from '../../lib/home/home-b-view';
  import GradeTile from '../ui/GradeTile.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { week }: { week: SpecWeek } = $props();

  const view = $derived(specWeekView(week));
</script>

<Section
  title={homeBCopy.specWeek}
  action={{ label: homeBCopy.tierList, href: homeBHrefs.tiers }}
  boxed
  testid="home-spec-week"
>
  <Row title={view.title} sub={view.sub} open>
    {#snippet lead()}
      {#if week.grade}<GradeTile letter={week.grade} testid="home-grade" />{/if}
    {/snippet}
  </Row>
</Section>
