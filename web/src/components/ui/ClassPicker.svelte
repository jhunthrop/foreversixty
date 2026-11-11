<!-- web/src/components/ui/ClassPicker.svelte -->
<!-- The class picker (component audit 2026-11-10 cluster 26; home spec 2026-11-06 H8, pages spec
     2026-11-08 §4.4): nine Rows in three columns from 1024, the ringed crest, the class name in its
     colour and its specs. Each row links its class guide; with `specLinks` the class name and each
     spec are their own links instead. The caller frames it (a Plate or a Section). -->
<script lang="ts">
  import { pickerClasses, specSummary, type PickerClass } from '../../lib/class-picker';
  import Crest from './Crest.svelte';
  import Row from './Row.svelte';
  import RowGrid from './RowGrid.svelte';

  let {
    classes = pickerClasses(),
    specLinks = false,
    testidPrefix,
  }: {
    classes?: readonly PickerClass[];
    specLinks?: boolean;
    /** Each row's testid before its class slug. */
    testidPrefix: string;
  } = $props();

  const CREST_PX = 44;
  /** Between two specs, outside the links. */
  const SPEC_SEPARATOR = ', ';
</script>

<RowGrid cols={3}>
  {#each classes as c (c.slug)}
    {#if specLinks}
      <Row title={c.name} titleTone={{ classSlug: c.slug }} testid={`${testidPrefix}${c.slug}`}>
        {#snippet lead()}<Crest cls={c.slug} size={CREST_PX} />{/snippet}
        {#snippet titleContent()}<a class="cpick-class" href={c.href}>{c.name}</a>{/snippet}
        {#snippet subContent()}{#each c.specs as spec, i (spec.key)}{#if i > 0}{SPEC_SEPARATOR}{/if}<a
              class="cpick-spec"
              href={spec.href}>{spec.name}</a
            >{/each}{/snippet}
      </Row>
    {:else}
      <Row
        title={c.name}
        titleTone={{ classSlug: c.slug }}
        sub={specSummary(c, SPEC_SEPARATOR)}
        href={c.href}
        open
        testid={`${testidPrefix}${c.slug}`}
      >
        {#snippet lead()}<Crest cls={c.slug} size={CREST_PX} />{/snippet}
      </Row>
    {/if}
  {/each}
</RowGrid>

<style>
  /* 24 px tall targets (WCAG 2.5.8): the class and each spec are separate links on one row. */
  .cpick-class,
  .cpick-spec {
    display: inline-block;
    line-height: 24px;
  }
  .cpick-class:hover {
    text-decoration: underline;
  }
  .cpick-spec {
    color: var(--color-muted);
  }
  .cpick-spec:hover {
    color: var(--color-gold-hover);
  }
  .cpick-class:focus-visible,
  .cpick-spec:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
</style>
