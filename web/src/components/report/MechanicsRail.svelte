<!-- web/src/components/report/MechanicsRail.svelte -->
<!-- The Mechanics tab's rail under This pull (logs spec 2026-11-07 section 6.2): Section `Clean`,
     a Row per judged player no avoidable mechanic landed on (crest, name in the class colour,
     `No avoidable damage`). Whole fight, like the tab: it reads the same unscoped summary the
     tab does. A pull with no table (every trash pull), or a table that judges nobody, draws
     nothing. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { cleanPlayers } from '../../lib/report/mechanics';
  import { classSlugOf } from '../../lib/report/planner-link';
  import type { Summary } from '../../lib/report/types';
  import { classTone } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { summary, classOf }: { summary: Summary; classOf: Map<string, string> } = $props();

  const judged = $derived(summary.mechanics?.table_found === true);
  const clean = $derived(cleanPlayers(summary));
</script>

{#if judged}
  <Section title="Clean" meta="No avoidable damage" testid="mechanics-rail" panel>
    {#if clean.length === 0}
      <p class="text-muted text-[14px]">Everyone took avoidable damage.</p>
    {:else}
      {#each clean as player (player.guid)}
        {@const slug = classSlugOf(classOf.get(player.guid)) ?? ''}
        <Row title={splitUnitName(player.name).name} titleTone={classTone(slug)} density="slim" open>
          {#snippet lead()}<Crest cls={slug} size={24} />{/snippet}
        </Row>
      {/each}
    {/if}
  </Section>
{/if}
