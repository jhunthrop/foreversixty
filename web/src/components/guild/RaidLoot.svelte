<!-- web/src/components/guild/RaidLoot.svelte -->
<!-- Loot (attendance spec 3.4, members and officers): grouped by boss under a label head with its
     mono count; one 44 px row per award: ItemHover (28 px icon), a muted arrow, the awardee's crest
     and name, and `by {officer}`. Sorted by quality then name. None reads by rank: an officer is told
     where to award it (the Loot tab links), a member that none was recorded. -->
<script lang="ts">
  import type { RaidLootGroup } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, byName } from '../../lib/guild/attendance-copy';
  import { OWN_ROW_WASH, ROW_DIVIDER } from '../../lib/guild/raid-report-styles';
  import { lootOrder } from '../../lib/guild/report-view';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import RaidBlock from './RaidBlock.svelte';
  import RaidLootItem from './RaidLootItem.svelte';

  let {
    groups,
    officer,
    lootHref,
    ownKeys = [],
  }: {
    groups: readonly RaidLootGroup[];
    officer: boolean;
    /** The guild's Loot tab. */
    lootHref: string;
    ownKeys?: readonly string[];
  } = $props();

  const filled = $derived(groups.filter((group) => group.awards.length > 0));
  const copy = attendanceCopy.loot;
</script>

<RaidBlock
  heading={attendanceCopy.report.loot}
  count={filled.reduce((sum, group) => sum + group.awards.length, 0)}
  testid="raid-loot"
>
  {#if filled.length === 0}
    <p class="text-muted text-[14px]" data-testid="raid-loot-none">
      {#if officer}{copy.officerEmptyLead}<a href={lootHref} class="text-text underline"
          >{copy.officerEmptyLink}</a
        >{copy.officerEmptyTail}{:else}{copy.memberEmpty}{/if}
    </p>
  {:else}
    {#each filled as group (group.encounter_id)}
      <div class="flex flex-col" data-testid="raid-loot-group">
        {#if filled.length > 1}
          <h4 class="label text-muted pb-1">
            {group.boss} <span class="numeric">{group.awards.length}</span>
          </h4>
        {/if}
        <ul>
          {#each lootOrder(group.awards) as award (`${award.item_id}-${award.character_key}-${award.at}`)}
            <!-- From 768: item, arrow, awardee, `by` right-aligned on one grid. Below: the item, then
                 the arrow, awardee and `by` on line 2. -->
            <li
              class="{ROW_DIVIDER} grid min-h-11 grid-cols-[16px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-1 py-1.5 md:grid-cols-[minmax(0,1fr)_16px_minmax(0,1fr)_auto]"
              style={ownKeys.includes(award.character_key) ? `background:${OWN_ROW_WASH}` : undefined}
              data-testid="raid-loot-row"
            >
              <span class="col-span-3 min-w-0 md:col-span-1"
                ><RaidLootItem {award} iconSize={28} nameSize="text-[14px]" /></span
              >
              <span class="text-muted" aria-hidden="true">→</span>
              <span class="flex min-w-0 items-center gap-2">
                <ClassCrestRing characterClass={award.class} size={20} />
                <span class="truncate text-[13px] font-semibold" style={`color:${classColorVar(award.class)}`}
                  >{award.name}</span
                >
              </span>
              <span class="text-muted text-right text-[12px] whitespace-nowrap">{byName(award.by_name)}</span>
            </li>
          {/each}
        </ul>
      </div>
    {/each}
  {/if}
</RaidBlock>
