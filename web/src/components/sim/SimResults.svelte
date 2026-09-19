<!-- web/src/components/sim/SimResults.svelte -->
<!-- The sim's results are the report's results. Every table here is the component the logs
     report already mounts, taking the same Summary the adapter produced, which is what the
     design means by "no second renderer": an improvement to the damage table reaches the
     simulator for free and cannot drift away from it.
     TimeChart is the one report component deliberately left out. Its series is a real
     fight's second-by-second damage; a sim's series is an average over three thousand
     iterations, and brushing an average would produce numbers that mean nothing. The
     Distribution tab is what stands in its place, and the strip says so. -->
<script lang="ts">
  import AuraTable from '../report/AuraTable.svelte';
  import ActorTable from '../report/ActorTable.svelte';
  import CastTable from '../report/CastTable.svelte';
  import ResourceGraphs from '../report/ResourceGraphs.svelte';
  import TimelinesView from '../report/TimelinesView.svelte';
  import { rowLink } from '../../lib/report/format';
  import { fullWindow } from '../../lib/report/window';
  import type { ActionNames } from '../../lib/sim/action-names';
  import { simCopy } from '../../lib/sim/copy';
  import { namedSummary, summarySentence } from '../../lib/sim/sentence';
  import type { Estimate } from '../../lib/sim/types';
  import type { Summary } from '../../lib/report/types';
  import DpsDistribution from './DpsDistribution.svelte';

  let {
    summary,
    estimate,
    iterationsRun,
    actionNames,
  }: {
    summary: Summary;
    estimate: Estimate;
    iterationsRun: number;
    /** The build's name table (Task 23). Null until it loads; rows then read as keys. */
    actionNames: ActionNames | null;
  } = $props();

  const TABS = [
    { id: 'damage', label: simCopy.tabDamage },
    { id: 'buffs', label: simCopy.tabBuffs },
    { id: 'debuffs', label: simCopy.tabDebuffs },
    { id: 'casts', label: simCopy.tabCasts },
    { id: 'resources', label: simCopy.tabResources },
    { id: 'timeline', label: simCopy.tabTimeline },
    { id: 'distribution', label: simCopy.tabDistribution },
  ] as const;

  let tab = $state<(typeof TABS)[number]['id']>('damage');

  // A sim's roster carries its own class name, so TimelinesView never falls back to this
  // lookup. It is created once rather than per render so the prop is referentially stable.
  const emptyClassMap = new Map<string, string>();

  // Names resolved once, here, and passed down instead of `summary`: the report components
  // render `ability.name`, `aura.name` and `cast.spell_name` as they find them, and a sim's
  // are engine action keys. `named` is a copy -- the stored SimResult keeps the keys.
  const named = $derived(namedSummary(summary, actionNames));
  const sentence = $derived(summarySentence(summary, actionNames));
  const buffs = $derived(named.auras.filter((track) => track.type === 'BUFF'));
  const debuffs = $derived(named.auras.filter((track) => track.type === 'DEBUFF'));

  const pill = `${rowLink} shrink-0 px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9`;
</script>

<p
  class="text-strong px-[18px] text-[15px] md:px-0"
  role="status"
  aria-live="polite"
  data-testid="sim-sentence"
>
  {sentence}
</p>

<section class="flex flex-col gap-3" data-testid="sim-results">
  <!-- No -mx-[18px] bleed here, unlike Header.astro's nav: that trick cancels a parent's
       own px-[18px] padding, and this section's ancestors up to <main> carry none -- every
       block in the sim island insets itself. Applying it anyway pushed this container 18px
       past the viewport on each side, a real page-level horizontal scroll the phone audit
       caught. px-[18px] alone already both insets the first tab to the page gutter and
       gives the last tab the same padding once scrolled to, with the container's own box
       never exceeding the viewport it already fills. -->
  <div
    class="flex flex-nowrap gap-x-1 overflow-x-auto px-[18px] md:flex-wrap md:px-0"
    role="tablist"
    aria-label={simCopy.resultsTablist}
  >
    {#each TABS as entry (entry.id)}
      <button
        type="button"
        role="tab"
        class={`${pill} ${tab === entry.id ? 'text-gold' : 'text-nav'}`}
        aria-selected={tab === entry.id}
        onclick={() => (tab = entry.id)}
        data-testid={`sim-tab-${entry.id}`}
      >
        {entry.label}
      </button>
    {/each}
  </div>

  <div class="border-line bg-raised rounded-panel mx-[18px] border md:mx-0">
    {#if tab === 'damage'}
      <ActorTable actors={named.damage_done} durationMs={named.duration_ms} metricLabel="DPS" />
    {:else if tab === 'buffs'}
      <AuraTable tracks={buffs} durationMs={named.duration_ms} kind="BUFF" />
    {:else if tab === 'debuffs'}
      <AuraTable tracks={debuffs} durationMs={named.duration_ms} kind="DEBUFF" />
    {:else if tab === 'casts'}
      <CastTable rows={named.casts} durationMs={named.duration_ms} />
    {:else if tab === 'resources'}
      <ResourceGraphs tracks={named.resources} durationMs={named.duration_ms} />
    {:else if tab === 'timeline'}
      <!-- One iteration's real cast sequence, not an average: the engine's result carries a
           sample cast log, so the lane view means the same thing here as on a logged fight. -->
      <TimelinesView summary={named} window={fullWindow(named.duration_ms)} classOf={emptyClassMap} />
    {:else}
      <DpsDistribution {estimate} {iterationsRun} />
    {/if}
  </div>
</section>
