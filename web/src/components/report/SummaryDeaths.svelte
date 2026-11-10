<!-- web/src/components/report/SummaryDeaths.svelte -->
<!-- The Summary rail's Deaths (logs spec 2026-11-07 6.2, restored in the owner's review of
     2026-11-09): one Row per death in time order, the time, the crest, the name in its class
     colour and "{ability} · {source} · {amount}" (a melee swing reads Melee), with the Deaths ›
     link. Nobody died: one short line. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import type { Summary } from '../../lib/report/types';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import ClassIcon from './ClassIcon.svelte';

  /** The log's null unit: a fall, a hazard, or a source the log did not name. */
  const NULL_GUID = /^0+$/;

  let { summary, tabHref }: { summary: Summary; tabHref: string } = $props();

  const deaths = $derived([...summary.deaths].sort((a, b) => a.at_ms - b.at_ms));

  function cause(death: Summary['deaths'][number]): string {
    const blow = death.killing_blow;
    if (blow === undefined) return '';
    const spell =
      NULL_GUID.test(blow.source_guid) && blow.spell_name === ''
        ? 'a fall, a hazard or an untracked source'
        : blow.spell_name === ''
          ? 'Melee'
          : blow.spell_name;
    return `${spell} · ${splitUnitName(blow.source_name).name} · ${formatAmount(blow.amount)}`;
  }
</script>

<Section
  title="Deaths"
  action={deaths.length === 0 ? undefined : { label: 'Deaths', href: tabHref, arrow: true }}
  panel
  testid="panel-deaths"
>
  {#if deaths.length === 0}
    <p class="text-muted text-[14px] leading-5" data-testid="panel-deaths-none">Nobody died.</p>
  {:else}
    {#each deaths as death, i (`${death.guid}-${death.at_ms}-${i}`)}
      <Row title={splitUnitName(death.name).name} sub={cause(death)} open density="slim">
        {#snippet lead()}<span class="text-muted w-9 text-right text-[13px] tabular-nums"
            >{formatDuration(death.at_ms)}</span
          ><ClassIcon className={death.class} size={24} />{/snippet}
        {#snippet titleContent()}<span style:color={classColorVar(death.class)}
            >{splitUnitName(death.name).name}</span
          >{/snippet}
      </Row>
    {/each}
  {/if}
</Section>
