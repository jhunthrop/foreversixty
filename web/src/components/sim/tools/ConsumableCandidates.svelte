<!-- web/src/components/sim/tools/ConsumableCandidates.svelte -->
<!-- Design 3.1.5, ratified as BulkSpec.Consumables by contract 10.1 A5: the consumables the
     settings panel applies as a setting can instead be tried one at a time as candidates,
     each ticked one replacing CharacterSpec.Consumes wholesale for its combination.
     The id list is the preset's own, not a second copy of IDS.md -- settings.ts is the one
     place this lane names a consumable id -- and the names and icons come from
     simbuffs.json (contract 10.4) rather than from a de-underscored id. -->
<script lang="ts">
  import { dataUrl } from '../../../lib/planner/load';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { PRESET_CONSUMABLES } from '../../../lib/sim/settings';
  import { buffIcon, buffName, type SimBuffFile } from '../../../lib/sim/sim-buffs';

  let {
    picked,
    buffs,
    treeVersion,
    ontoggle,
  }: {
    picked: readonly string[];
    buffs: SimBuffFile;
    treeVersion: string;
    ontoggle: (id: string) => void;
  } = $props();

  const offered = $derived([...new Set(PRESET_CONSUMABLES['raid-buffed'])]);
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0"
  data-testid="sim-consumables"
>
  <h3 class="section-title text-[14px]">{bulkCopy.tryEach}</h3>
  <p class="text-muted text-[12px]">{bulkCopy.consumableCandidates}</p>
  <ul class="flex flex-wrap gap-3">
    {#each offered as id (id)}
      <li>
        <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
          <input
            type="checkbox"
            class="h-5 w-5"
            data-testid={`sim-consumable-${id}`}
            checked={picked.includes(id)}
            onchange={() => ontoggle(id)}
          />
          {#if buffIcon(buffs, id) !== ''}
            <img
              src={dataUrl(treeVersion, `icons/${buffIcon(buffs, id)}.webp`)}
              alt=""
              width="20"
              height="20"
              loading="lazy"
              decoding="async"
              class="rounded-control border-line h-5 w-5 border object-cover"
            />
          {/if}
          {buffName(buffs, id)}
        </label>
      </li>
    {/each}
  </ul>
</section>
