<!-- web/src/components/home/HomeUpgrades.svelte -->
<!-- Your upgrades (on the panel under its heading, as the feed beside it): the five biggest
     upgrades for the character's band, each the item in its quality colour, its slot and source,
     the piece it replaces ("over Mantle of Thieves", in that piece's quality colour) and what the
     pick adds in the band's own unit (DPS, HPS, a tank's score), with Best in slot at the right of the title; under them the whole gap (every
     measured upgrade's gain), how many slots are already best in slot, and See all N to the band's
     best-in-slot list when there are more than five. Left out with none to show. -->
<script lang="ts">
  import ItemName from '../ui/ItemName.svelte';
  import type { Item } from '../../lib/planner/types';
  import type { CachedUpgrades } from '../../lib/home/upgrades-loader';
  import type { SlotUpgrade } from '../../lib/home/upgrades';
  import { UPGRADES_SHOWN, upgradeLine, type HomeLink } from '../../lib/home/home-view';
  import { homeBCopy } from '../../lib/home/home-b-view';
  import { factsCopy } from '../../lib/home/home-facts';
  import { upgradeGain, upgradesFacts } from '../../lib/home/upgrades-view';
  import { LINK } from '../../lib/ui/primitives';
  import ItemRow from '../ui/ItemRow.svelte';
  import Section from '../ui/Section.svelte';
  import HomeFacts from './HomeFacts.svelte';

  let { upgrades, bis, build }: { upgrades: CachedUpgrades | null; bis: HomeLink | null; build: string } =
    $props();

  const items = $derived(new Map<number, Item>((upgrades?.items ?? []).map((item) => [item.id, item])));
  const all = $derived(upgrades?.result.upgrades ?? []);
  const unit = $derived(upgrades?.result.scoreUnit ?? 'dps');
  const list = $derived(all.slice(0, UPGRADES_SHOWN));
  const facts = $derived(upgradesFacts(upgrades?.result));
  const more = $derived(bis !== null && all.length > UPGRADES_SHOWN);

  /** The piece worn in the slot, as the item table knows it; undefined for an id it does not know. */
  function worn(upgrade: SlotUpgrade): Item | undefined {
    return upgrade.wornItemId === undefined ? undefined : items.get(upgrade.wornItemId);
  }
</script>

{#if list.length > 0}
  <Section
    title={homeBCopy.upgrades}
    action={bis === null ? undefined : { label: homeBCopy.bisLink, href: bis.href }}
    boxed
    testid="home-upgrades"
  >
    <div class="home-upgrades-rows">
      {#each list as upgrade (upgrade.slot)}
        {@const pick = items.get(upgrade.pick.item_id)}
        {@const piece = worn(upgrade)}
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
          gain={upgradeGain(upgrade, unit)}
          testid="home-upgrade"
        >
          {#snippet detail()}
            <span data-testid="home-upgrade-over">
              {#if upgrade.wornItemId === undefined}{factsCopy.emptySlot}{:else if piece !== undefined}{factsCopy.over}
                <ItemName
                  name={piece.name}
                  quality={piece.quality}
                  id={piece.id}
                  class="font-semibold"
                />{/if}
            </span>
          {/snippet}
        </ItemRow>
      {/each}
    </div>
    {#if facts.length > 0 || more}
      <div class="home-upgrades-foot">
        <HomeFacts {facts} testid="home-upgrades-total" />
        {#if more && bis !== null}<a
            class="{LINK} shrink-0 text-[14px]"
            href={bis.href}
            data-testid="home-upgrades-all">{factsCopy.seeAll(all.length)} ›</a
          >{/if}
      </div>
    {/if}
  </Section>
{/if}

<style>
  .home-upgrades-foot {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 12px;
    padding-top: 12px;
  }
  /* One line per upgrade's name on the desktop rail (an ellipsis past it); the worn piece under it
     wraps, so "over Lieutenant Commander's Plate Helmet" is never cut. A phone wraps both. */
  @media (min-width: 768px) {
    .home-upgrades-rows :global([data-testid^='item-hover-']) {
      display: block;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
    }
  }
</style>
