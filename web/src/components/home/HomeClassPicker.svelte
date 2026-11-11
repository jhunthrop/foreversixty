<!-- web/src/components/home/HomeClassPicker.svelte -->
<!-- "Pick your class" (home spec 2026-11-06 H8): nine Rows, the crest, the class name in its colour
     and its specs, three columns from 1024, each to the class guide. Signed out it leads under the
     title band that names it; signed in with no character it is its own boxed section. -->
<script lang="ts">
  import classes from '../../data/classes.json';
  import { SPECS } from '../../lib/sim/specs';
  import { homeCopy, homeHrefs } from '../../lib/home/home-view';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Row from '../ui/Row.svelte';
  import RowGrid from '../ui/RowGrid.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';

  /** `headless`: the page's title already says "Pick your class" (signed out), so the grid stands
   *  on its own plate with no heading of its own. */
  let { headless = false }: { headless?: boolean } = $props();

  const CREST_PX = 44;
  const rows = classes.map((c) => ({
    ...c,
    specs: SPECS.filter((spec) => spec.class_slug === c.slug)
      .map((spec) => spec.name)
      .join(', '),
  }));
</script>

{#snippet grid()}
  <RowGrid cols={3}>
    {#each rows as c (c.slug)}
      <Row
        title={c.name}
        titleTone={{ classSlug: c.slug as ClassSlug }}
        sub={c.specs}
        href={homeHrefs.guide(c.slug)}
        open
        testid={`home-class-${c.slug}`}
      >
        {#snippet lead()}<Crest cls={c.name} size={CREST_PX} />{/snippet}
      </Row>
    {/each}
  </RowGrid>
{/snippet}

{#if headless}
  <Plate flush={false} label={homeCopy.pickClass} testid="home-class-picker">{@render grid()}</Plate>
{:else}
  <Section title={homeCopy.pickClass} boxed testid="home-class-picker">{@render grid()}</Section>
{/if}
