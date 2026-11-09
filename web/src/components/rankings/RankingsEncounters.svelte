<!-- web/src/components/rankings/RankingsEncounters.svelte -->
<!-- The ranked encounters as Rows (directions mock b1-rankings): the boss's name as the link, its
     zone and raid size under it, and its ranked kills at the right. The picker's list, and the
     "ranked elsewhere" list under an empty board. -->
<script lang="ts">
  import type { RankedEncounter } from '../../lib/rankings/api';
  import { encounterPickerCopy } from '../../lib/rankings/copy';
  import Row from '../ui/Row.svelte';

  let {
    encounters,
    hrefFor,
    testid,
  }: {
    encounters: readonly RankedEncounter[];
    hrefFor: (slug: string) => string;
    testid?: string;
  } = $props();
</script>

<div class="px-5 pb-2 max-md:px-[14px]" data-testid={testid}>
  {#each encounters as encounter (encounter.id)}
    <Row
      title={encounter.name}
      sub={encounterPickerCopy.where(encounter.zone, encounter.size) || undefined}
      href={hrefFor(encounter.slug)}
    >
      {#snippet titleContent()}<span class="text-gold">{encounter.name}</span>{/snippet}
      {#snippet trail()}<span class="text-strong text-[14px] tabular-nums"
          >{encounterPickerCopy.ranked(encounter.kills)}</span
        >{/snippet}
    </Row>
  {/each}
</div>
