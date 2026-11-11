<!-- web/src/components/guild/report/ReportDeaths.svelte -->
<!-- The officer's deaths by raider (lane h-edges, guild-officer round 1: "who needs a word"): each
     raider who died, in the night's order (never worst first), every death of theirs (what killed
     them, which pull, when) and, past one, how many on wipes and on kills. Two columns from 1280.
     Officers only; a member's own deaths are Your night's. -->
<script lang="ts">
  import type { RaidReport } from '../../../lib/guild/attendance-contract';
  import { raiderDeaths } from '../../../lib/guild/report-night-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { classTone } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let { report }: { report: RaidReport } = $props();

  const rows = $derived(raiderDeaths(report));
</script>

{#if rows.length > 0}
  <Section
    title={copy.deaths.head}
    meta={copy.deaths.meta(rows.length)}
    anchor="raid-deaths"
    panel
    testid="raid-deaths"
  >
    <ul class="grid gap-x-8 xl:grid-cols-2">
      {#each rows as row (row.key)}
        <li>
          <Row
            title={row.name}
            titleTone={classTone(row.class)}
            sub={row.deaths}
            open
            density="slim"
            testid="raid-deaths-row"
          >
            {#snippet lead()}<Crest cls={row.class} size={24} />{/snippet}
            {#snippet trail()}{#if row.count !== ''}<span class="text-text text-[13px] tabular-nums"
                  >{row.count}</span
                >{/if}{/snippet}
          </Row>
        </li>
      {/each}
    </ul>
  </Section>
{/if}
