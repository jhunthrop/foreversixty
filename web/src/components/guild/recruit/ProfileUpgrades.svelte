<!-- web/src/components/guild/recruit/ProfileUpgrades.svelte -->
<!-- The member profile's Next upgrades (lane h-edges, guild-officer round 7): the gear gap named
     slot by slot, the biggest known gain first. Each piece is an ItemRow, as every item on the
     page: its icon in the quality frame and the quality-coloured name from the raider's class item
     file's own row (fetched once, cached), the slot and where it comes from under it, the gain at the right,
     the client tooltip on hover, focus and long-press. Until the class file answers, a row draws
     the socket and the epic colour the BiS picks are. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import activeBuild from '../../../data/active-build.json';
  import type { UpgradeRowModel } from '../../../lib/guild/profile-view';
  import { fetchItemRow, fetchItemTooltipModel } from '../../../lib/items/lookup';
  import type { Item } from '../../../lib/planner/types';
  import type { ItemTooltipModel } from '../../../lib/items/tooltip';
  import { profileCopy as copy } from '../../../lib/guild/recruit-copy';
  import BisTooltipHost from '../../BisTooltipHost.svelte';
  import ItemRow from '../../ui/ItemRow.svelte';
  import Section from '../../ui/Section.svelte';

  /** The quality a row draws before its class file answers: a BiS pick at 60 is almost always epic. */
  const QUALITY_UNTIL_READ = 4;

  let { rows, classSlug }: { rows: readonly UpgradeRowModel[]; classSlug: string } = $props();

  const build = activeBuild.build;
  let items = $state<ReadonlyMap<number, Item>>(new Map());

  const resolve = (itemId: number): Promise<ItemTooltipModel | null> =>
    classSlug === '' ? Promise.resolve(null) : fetchItemTooltipModel(build, classSlug, itemId);

  onMount(() => {
    if (classSlug === '') return;
    void Promise.all(
      rows.map(
        async (row) =>
          [row.itemId, await fetchItemRow(build, classSlug, row.itemId).catch(() => null)] as const,
      ),
    ).then((found) => {
      items = new Map(found.filter((entry): entry is [number, Item] => entry[1] !== null));
    });
  });
</script>

<Section title={copy.upgradesHead} panel testid="guild-profile-upgrades">
  <ul>
    {#each rows as row (row.key)}
      {@const item = items.get(row.itemId)}
      <li>
        <ItemRow
          item={{
            id: row.itemId,
            name: row.name,
            quality: item?.quality ?? QUALITY_UNTIL_READ,
            icon: item?.icon,
          }}
          {build}
          iconSize={32}
          source={row.sub}
          gain={row.gain ?? undefined}
          testid="guild-profile-upgrade"
        />
      </li>
    {/each}
  </ul>
  <BisTooltipHost {build} modelsElementId="guild-profile-upgrade-models" {resolve} />
</Section>
