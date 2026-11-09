<!-- web/src/components/home/HomeUpgrades.svelte -->
<!-- Your upgrades (round 5: the page's second featured panel, the plain frame): the five biggest
     upgrades for the character's band, each the item in its quality colour, its slot and source,
     and what it adds, with Best in slot at the right of the title. Left out with none to show. -->
<script lang="ts">
  import type { Item } from '../../lib/planner/types';
  import type { CachedUpgrades } from '../../lib/home/upgrades-loader';
  import { UPGRADES_SHOWN, gainText, upgradeLine, type HomeLink } from '../../lib/home/home-view';
  import { homeBCopy } from '../../lib/home/home-b-view';
  import { LINK } from '../../lib/ui/primitives';
  import { sectionRegion } from '../../lib/ui/regions';
  import Frame from '../ui/Frame.svelte';
  import ItemRow from '../ui/ItemRow.svelte';

  let { upgrades, bis, build }: { upgrades: CachedUpgrades | null; bis: HomeLink | null; build: string } =
    $props();

  const items = $derived(new Map<number, Item>((upgrades?.items ?? []).map((item) => [item.id, item])));
  const list = $derived((upgrades?.result.upgrades ?? []).slice(0, UPGRADES_SHOWN));
</script>

{#if list.length > 0}
  <Frame region={sectionRegion(homeBCopy.upgrades)} testid="home-upgrades">
    <div class="home-upgrades-head">
      <h2 class="text-section">{homeBCopy.upgrades}</h2>
      {#if bis !== null}<a class="{LINK} text-[14px]" href={bis.href}>{homeBCopy.bisLink} ›</a>{/if}
    </div>
    {#each list as upgrade (upgrade.slot)}
      {@const pick = items.get(upgrade.pick.item_id)}
      <ItemRow
        item={{
          id: upgrade.pick.item_id,
          name: upgrade.pick.item_name,
          quality: pick?.quality ?? 1,
          icon: pick?.icon,
        }}
        {build}
        iconSize={36}
        source={upgradeLine(upgrade)}
        gain={upgrade.gainDps === null
          ? undefined
          : { value: gainText(upgrade.gainDps), unit: 'DPS', tone: 'kill' }}
        testid="home-upgrade"
      />
    {/each}
  </Frame>
{/if}

<style>
  .home-upgrades-head {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 12px;
    padding-bottom: 6px;
  }
</style>
