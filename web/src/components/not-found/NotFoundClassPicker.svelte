<!-- web/src/components/not-found/NotFoundClassPicker.svelte -->
<!-- The 404's "Pick your class" (pages spec 2026-11-08 §4.4, ruling SHOULD; round 5 mock b4-404): one
     panel with its heading inside, nine class Rows in three columns from 1024, the ringed crest, the class name in its colour
     linking its guides, and each spec linking its own guide. Rendered at build time (no runtime). -->
<script lang="ts">
  import classes from '../../data/classes.json';
  import { SPECS } from '../../lib/sim/specs';
  import { notFoundCopy } from '../../lib/not-found/copy';
  import { titleColor, type ClassSlug } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Row from '../ui/Row.svelte';
  import RowGrid from '../ui/RowGrid.svelte';
  import Section from '../ui/Section.svelte';

  const CREST_PX = 44;
  /** Between two spec links, outside both. */
  const SPEC_SEPARATOR = ', ';
  const rows = classes.map((c) => ({
    ...c,
    specs: SPECS.filter((spec) => spec.class_slug === c.slug).toSorted((a, b) => a.tree_index - b.tree_index),
  }));
</script>

<Section title={notFoundCopy.pick} panel flush testid="not-found-classes">
  <RowGrid cols={3}>
    {#each rows as c (c.slug)}
      <Row title={c.name} testid={`not-found-class-${c.slug}`}>
        {#snippet lead()}<Crest cls={c.name} size={CREST_PX} />{/snippet}
        {#snippet titleContent()}<a
            class="nf-class"
            href={`/guides/${c.slug}`}
            style:color={titleColor({ classSlug: c.slug as ClassSlug })}>{c.name}</a
          >{/snippet}
        {#snippet subContent()}{#each c.specs as spec, i (spec.spec_slug)}{#if i > 0}{SPEC_SEPARATOR}{/if}<a
              class="nf-spec"
              href={`/guides/${c.slug}/${spec.spec_slug}`}>{spec.name}</a
            >{/each}{/snippet}
      </Row>
    {/each}
  </RowGrid>
</Section>

<style>
  /* 24 px tall targets (WCAG 2.5.8): the class and each spec are separate links on one row. */
  .nf-class,
  .nf-spec {
    display: inline-block;
    line-height: 24px;
  }
  .nf-class:hover {
    text-decoration: underline;
  }
  .nf-spec {
    color: var(--color-muted);
  }
  .nf-spec:hover {
    color: var(--color-gold-hover);
  }
  .nf-class:focus-visible,
  .nf-spec:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
</style>
