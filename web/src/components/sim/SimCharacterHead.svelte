<!-- web/src/components/sim/SimCharacterHead.svelte -->
<!-- The /sim head with a character (spec 2026-11-06-sim §4.3, ruling S1): the ringed class crest,
     the character's name as the title, the meta line (race, spec, level, talents, item level) and
     where the gear came from, the answer at the right with its actions beside it, the tool tabs
     under it. There is no
     "Change source": the header's selector changes the character (one way to switch). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Crest from '../ui/Crest.svelte';
  import { classTone } from '../../lib/ui/primitives';
  import PageHeader from '../ui/PageHeader.svelte';

  let {
    name,
    classSlug,
    lines,
    tabs,
    answer,
    actions,
  }: {
    name: string;
    classSlug: string;
    /** The meta lines, each its own span. */
    lines: readonly string[];
    /** The tool tabs (SimToolTabs): /sim only, never /sim/<id> (ruling S5). */
    tabs?: Snippet;
    answer: Snippet;
    /** The answer's actions (save, open, the next tool), beside it on the head's row. */
    actions?: Snippet;
  } = $props();
</script>

<PageHeader title={name} tone={classTone(classSlug)} testid="sim-head" {answer} {actions} {tabs}>
  {#snippet crest()}<Crest cls={classSlug} size={56} testid="sim-character-class-icon" />{/snippet}
  {#snippet meta()}
    {#each lines as line, index (index)}<span data-testid={index === 0 ? 'sim-head-meta' : 'sim-head-source'}
        >{line}</span
      >{/each}
  {/snippet}
</PageHeader>
