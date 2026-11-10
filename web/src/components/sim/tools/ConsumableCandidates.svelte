<!-- web/src/components/sim/tools/ConsumableCandidates.svelte -->
<!-- Design 3.1.5, ratified as BulkSpec.Consumables by contract 10.1 A5: the consumables the
     settings panel applies as a setting can instead be tried one at a time as candidates,
     each ticked one replacing CharacterSpec.Consumes wholesale for its combination.
     The id list is the preset's own, not a second copy of IDS.md -- settings.ts is the one
     place this lane names a consumable id -- and the names and icons come from
     simbuffs.json (contract 10.4) rather than from a de-underscored id. -->
<script lang="ts">
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { presetConsumables } from '../../../lib/sim/settings';
  import { buffIcon, buffName, type SimBuffFile } from '../../../lib/sim/sim-buffs';
  import ToolSection from './ToolSection.svelte';

  let {
    picked,
    buffs,
    treeVersion,
    spec,
    level,
    ontoggle,
  }: {
    picked: readonly string[];
    buffs: SimBuffFile;
    treeVersion: string;
    /** The character's spec key and level: the offer is that spec's raid set at that level. */
    spec: string;
    level: number;
    ontoggle: (id: string) => void;
  } = $props();

  const offered = $derived([...new Set(presetConsumables('raid-buffed', spec, level))]);
</script>

<ToolSection title={bulkCopy.tryEach} testid="sim-consumables">
  <ul class="grid grid-cols-1 gap-x-6 pt-1 md:grid-cols-2 xl:grid-cols-3">
    {#each offered as id (id)}
      <li>
        <label class="text-text flex min-h-11 items-center gap-2 text-[14px]">
          <input
            type="checkbox"
            class="accent-gold h-5 w-5"
            data-testid={`sim-consumable-${id}`}
            checked={picked.includes(id)}
            onchange={() => ontoggle(id)}
          />
          {#if buffIcon(buffs, id) !== ''}
            <ItemIcon icon={buffIcon(buffs, id)} build={treeVersion} size={28} />
          {/if}
          {buffName(buffs, id)}
        </label>
      </li>
    {/each}
  </ul>
</ToolSection>
