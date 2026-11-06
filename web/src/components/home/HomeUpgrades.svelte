<!-- web/src/components/home/HomeUpgrades.svelte -->
<!-- "Your upgrades" (home spec 2026-11-06 H3): the page's one Plate, the top three BiS picks for the
     character's band with the DPS each adds over what it wears, then a link to the rest on BiS. Its
     own skeleton while the band loads and its own LoadError when it fails. -->
<script lang="ts">
  import type { Item } from '../../lib/planner/types';
  import { SLOT_LABELS } from '../../lib/planner/types';
  import type { CachedUpgrades } from '../../lib/home/upgrades-loader';
  import {
    UPGRADES_SHOWN,
    bandWords,
    gainText,
    homeCopy,
    upgradeLines,
    type HomeLink,
  } from '../../lib/home/home-view';
  import ItemRow from '../ui/ItemRow.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';

  let {
    status,
    upgrades,
    bis,
    build,
    onRetry,
  }: {
    status: 'loading' | 'ready' | 'failed' | 'none';
    upgrades: CachedUpgrades | null;
    bis: HomeLink | null;
    build: string;
    onRetry: () => void;
  } = $props();

  const SKELETON_ROWS = 3;
  const items = $derived(new Map<number, Item>((upgrades?.items ?? []).map((item) => [item.id, item])));
  const list = $derived(upgrades?.result.upgrades ?? []);
  const band = $derived(upgrades === null ? '' : bandWords(upgrades.band));
  const action = $derived(
    bis === null || list.length === 0
      ? undefined
      : { label: homeCopy.upgradesAll(list.length), href: bis.href },
  );
</script>

{#if status !== 'none'}
  <Section
    title={homeCopy.upgrades}
    meta={band === '' ? undefined : homeCopy.upgradesMeta(band)}
    {action}
    testid="home-upgrades"
  >
    {#if status === 'loading'}
      <Plate flush={false} label={homeCopy.upgrades}>
        <div class="flex flex-col gap-3" role="status" aria-busy="true">
          {#each Array.from({ length: SKELETON_ROWS }, (_, index) => index) as index (index)}
            <span class="skeleton-block block h-[42px]" aria-hidden="true"></span>
          {/each}
        </div>
      </Plate>
    {:else if status === 'failed'}
      <Plate flush={false} label={homeCopy.upgrades}>
        <LoadError message={homeCopy.upgradesFailed} {onRetry} testid="home-upgrades-error" />
      </Plate>
    {:else if list.length === 0}
      <p class="text-muted text-[15px] leading-6">{homeCopy.upgradesNone(band)}</p>
    {:else}
      <Plate label={homeCopy.upgrades}>
        {#each list.slice(0, UPGRADES_SHOWN) as upgrade (upgrade.slot)}
          {@const pick = items.get(upgrade.pick.item_id)}
          {@const lines = upgradeLines(upgrade)}
          <ItemRow
            item={{
              id: upgrade.pick.item_id,
              name: upgrade.pick.item_name,
              quality: pick?.quality ?? 1,
              icon: pick?.icon,
            }}
            {build}
            slotLabel={SLOT_LABELS[upgrade.slot]}
            source={lines.source}
            gain={upgrade.gainDps === null
              ? undefined
              : { value: gainText(upgrade.gainDps), unit: 'DPS', caption: lines.caption }}
            testid="home-upgrade"
          />
        {/each}
      </Plate>
    {/if}
  </Section>
{/if}
