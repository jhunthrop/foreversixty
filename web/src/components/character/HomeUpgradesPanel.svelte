<!-- web/src/components/character/HomeUpgradesPanel.svelte -->
<!-- Home rebuild spec §3.B.3: "Your upgrades", full width -- the worn-vs-BiS table for the
     signed-in hero's own band, built on `lib/home/upgrades.ts`'s pure comparison and
     `lib/home/upgrades-loader.ts`'s fetch. Mounted by `HomeAccountPanel.svelte` into
     `index.astro`'s `home-upgrades-slot`, the same dynamic-import-once-signed-in-and-ready
     trick that island already uses for `HomeSwitchCharacterPanel`/`HomeHeroCards`, so a
     signed-out page never fetches this module, its BiS file, or its item table either. Owns
     its own header row (heading + "Full list for …" aside link) rather than splitting that
     across index.astro's static markup and this island, since the aside link's own text
     names the character's real spec/band and cannot be known before this mounts. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import type { MeCharacter } from '../../lib/account/api';
  import { bisPageHref } from '../../lib/bis/hover';
  import { bisCopy } from '../../lib/bis/copy';
  import { homeHeroCardsCopy, homeUpgradesCopy } from '../../lib/home-panel-copy';
  import { loadBisContextFor, type BisContext } from '../../lib/home/upgrades-loader';
  import { upgradesFor, type UpgradesResult } from '../../lib/home/upgrades';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import { dataUrl } from '../../lib/planner/load';
  import { rarityBorderColorFor } from '../../lib/planner/items';
  import UpgradeRow from './UpgradeRow.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let { hero }: { hero: MeCharacter } = $props();

  type Status = 'loading' | 'no-spec' | 'no-gear' | 'no-list-yet' | 'ready';

  let status = $state<Status>('loading');
  let ctx = $state<BisContext | null>(null);
  let result = $state<UpgradesResult | null>(null);

  $effect(() => {
    const current = hero;
    ctx = null;
    result = null;
    if (current.spec === undefined) {
      status = 'no-spec';
      return;
    }
    if (current.build?.gear === undefined) {
      status = 'no-gear';
      return;
    }
    status = 'loading';
    void loadBisContextFor(current).then((loaded) => {
      if (current !== hero) return;
      if (loaded === null) {
        status = 'no-list-yet';
        return;
      }
      ctx = loaded;
      result = upgradesFor(current, loaded.band, loaded.items);
      status = 'ready';
    });
  });

  const classSlug = $derived(classSlugFromName(hero.class ?? ''));
  const bandLabel = $derived(ctx === null ? '' : bisCopy.bandRangeLabel(ctx.band.band));
  const fullListHref = $derived(
    ctx === null || hero.faction === undefined
      ? undefined
      : bisPageHref(ctx.specKey, hero.faction, ctx.band.band),
  );
  const fullListLabel = $derived(
    ctx === null ? '' : homeUpgradesCopy.fullListLink(specDisplayName(ctx.specKey), bandLabel),
  );
  const shownAlreadyBis = $derived(result === null ? [] : result.alreadyBis.slice(0, 3));
  const extraAlreadyBisCount = $derived(result === null ? 0 : Math.max(0, result.alreadyBis.length - 3));
</script>

<div class="flex items-baseline justify-between" data-testid="home-upgrades-header">
  <h2 class="display text-strong text-[18px] font-bold tracking-[0.10em]">{homeUpgradesCopy.heading}</h2>
  {#if fullListHref !== undefined}
    <a href={fullListHref} class="text-gold text-[12px] font-bold tracking-[0.12em] uppercase">
      {fullListLabel}
    </a>
  {/if}
</div>

<div class="bg-raised border-line rounded-panel border px-5 py-[18px]" data-testid="home-upgrades-body">
  {#if status === 'loading'}
    <Skeleton lines={4} rowHeight="h-5" testid="home-upgrades-skeleton" />
  {:else if status === 'no-spec'}
    <p class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" data-testid="home-upgrades-no-spec">
      <span class="label text-gold">{homeHeroCardsCopy.pickASpec}</span>
      <a href="/planner" class="text-nav text-[13px] font-semibold">{homeHeroCardsCopy.pickASpecLine}</a>
    </p>
  {:else if status === 'no-gear'}
    <p
      class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]"
      data-testid="home-upgrades-not-yet-available"
    >
      <span class="text-muted">{homeHeroCardsCopy.bestInSlotNotAvailable}</span>
    </p>
  {:else if status === 'no-list-yet'}
    <p class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" data-testid="home-upgrades-no-list-yet">
      <span class="text-muted">{homeHeroCardsCopy.bestInSlotNoListYet}</span>
    </p>
  {:else if result !== null}
    {#if result.upgrades.length === 0}
      <p class="text-muted mb-3 text-[14px]" data-testid="home-upgrades-all-match">
        {homeUpgradesCopy.everySlotMatches(bandLabel)}
      </p>
    {:else}
      <div class="upgrades-table-header" data-testid="home-upgrades-table">
        <span class="upgrades-table-header-slot">{homeUpgradesCopy.slotHeader}</span>
        <span>{homeUpgradesCopy.youWearHeader}</span>
        <span></span>
        <span>{homeUpgradesCopy.bestInSlotHeader}</span>
        <span class="upgrades-table-header-gain">{homeUpgradesCopy.gainHeader}</span>
      </div>
      <ul class="flex flex-col" data-testid="home-upgrades-list">
        {#each result.upgrades as upgrade (upgrade.slot)}
          <UpgradeRow
            {upgrade}
            wornItem={upgrade.wornItemId === undefined ? undefined : ctx?.items.get(upgrade.wornItemId)}
            pickItem={ctx?.items.get(upgrade.pick.item_id)}
            build={activeBuild.build}
            {classSlug}
          />
        {/each}
      </ul>
    {/if}
    {#if result.alreadyBis.length > 0}
      <p
        class="border-line-soft mt-3 flex flex-wrap items-center gap-2 border-t pt-3 text-[13px]"
        data-testid="home-upgrades-already-bis"
      >
        <span class="text-muted font-semibold">{homeUpgradesCopy.alreadyBestInSlotPrefix}</span>
        {#each shownAlreadyBis as entry, index (entry.slot)}
          {#if index > 0}<span class="text-muted">·</span>{/if}
          {@const item = ctx?.items.get(entry.itemId)}
          <img
            src={dataUrl(activeBuild.build, `icons/${item?.icon ?? ''}.webp`)}
            alt=""
            width="18"
            height="18"
            loading="lazy"
            decoding="async"
            class="rounded-control border object-cover"
            style={`border-color:${rarityBorderColorFor(item?.quality ?? 1)}`}
          />
          <span class="text-strong">{entry.itemName}</span>
        {/each}
        {#if extraAlreadyBisCount > 0}
          <span class="text-muted">· {homeUpgradesCopy.moreSlots(extraAlreadyBisCount)}</span>
        {/if}
      </p>
    {/if}
  {/if}
</div>

<style>
  .upgrades-table-header {
    display: grid;
    grid-template-columns: 90px 1fr 20px 1fr 90px;
    column-gap: 12px;
    padding-bottom: 8px;
    border-bottom: 1px solid var(--color-line-soft, #222);
    font-size: 11px;
    font-weight: 700;
    text-transform: uppercase;
    color: var(--color-muted, #999);
  }
  .upgrades-table-header-gain {
    text-align: right;
  }
  @media (max-width: 720px) {
    .upgrades-table-header {
      display: none;
    }
  }
</style>
