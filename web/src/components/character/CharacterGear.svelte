<!-- web/src/components/character/CharacterGear.svelte -->
<!-- The character page's one Plate (character spec 2026-11-06 C1, C4): the shared Paperdoll in
     `worn` mode inside the `Gear` Section, every slot an icon in its quality frame, the item name
     and its item level; the owner's cells carry the gain to the band's pick. Selecting a slot opens
     the Inspector (at the side from 1760, in place below): the worn item, the band's pick with its
     source and drop chance, two runners-up and one link to best in slot. `cells` decides the words
     (lib/character/doll.ts); this lays them out. -->
<script lang="ts">
  import type { BisBand } from '../../lib/bis/types';
  import { bandSlotFor, runnersUp } from '../../lib/character/gear-context';
  import { characterPageCopy } from '../../lib/character/page-view';
  import { gearCopy } from '../../lib/gear/copy';
  import type { GearItemTable } from '../../lib/gear/qualities';
  import type { CharacterGearSlot } from '../../lib/gear/types';
  import { rarityClassFor } from '../../lib/planner/items';
  import { signedFigure, type ItemRef, LINK } from '../../lib/ui/primitives';
  import type { DollCell } from '../../lib/ui/paperdoll';
  import Inspector from '../ui/Inspector.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import Paperdoll from '../ui/Paperdoll.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';

  let {
    title = characterPageCopy.gear,
    meta,
    cells,
    gear,
    band,
    table,
    unit,
    build,
    bisHref,
    loading = false,
    failed = false,
    note = null,
  }: {
    title?: string;
    meta?: string;
    cells: readonly DollCell[];
    gear: readonly CharacterGearSlot[];
    band: BisBand | null;
    table: GearItemTable | null;
    /** The band's rate word for a runner-up's delta ("DPS", "HPS"). */
    unit: string;
    build: string;
    bisHref: string | null;
    loading?: boolean;
    failed?: boolean;
    /** One muted line under the doll (items not in our data, a failed items file). */
    note?: string | null;
  } = $props();

  let selected = $state<string | null>(null);

  function ref(id: number, name: string): ItemRef {
    const facts = table?.get(id);
    return { id, name, quality: facts?.quality ?? 1, icon: facts?.icon ?? undefined };
  }

  const bySlot = $derived(new Map(gear.map((slot) => [slot.slot, slot])));
</script>

<Section
  {title}
  {meta}
  action={bisHref === null ? undefined : { label: characterPageCopy.gearAction, href: bisHref }}
  testid="character-gear-section"
>
  <Plate label={characterPageCopy.gear} testid="character-gear">
    <Paperdoll
      {cells}
      mode="worn"
      {build}
      {loading}
      {failed}
      bind:selected
      label={gearCopy.listLabel}
      testid="character-doll"
    >
      {#snippet detail(cell, placement)}
        {@const api = bySlot.get(cell.slot as CharacterGearSlot['slot'])}
        {@const pick = api?.pick ?? null}
        {@const row = bandSlotFor(band, cell.slot, pick?.id ?? null)}
        <Inspector selection={cell.label} {placement} testid="character-inspector" sections={[]}>
          {#snippet head()}
            {#if placement === 'side' && cell.item}
              <p class="text-label tracking-label text-muted font-bold uppercase">{cell.label}</p>
              <p class="text-[18px] leading-[22px] font-bold {rarityClassFor(cell.item.quality)}">
                {cell.item.name}
              </p>
              {#if cell.source}<p class="text-muted text-[13px] leading-[18px]">{cell.source}</p>{/if}
            {/if}
            {#if pick !== null}
              <p class="text-label tracking-label text-muted mt-3 font-bold uppercase">Best in slot</p>
              <ItemRow
                item={ref(pick.id, pick.name)}
                {build}
                source={pick.chance === null
                  ? pick.source
                  : `${pick.source} · ${gearCopy.chance(pick.chance)}`}
                bare
                testid="character-inspector-pick"
              />
              {#each runnersUp(row) as alt (alt.id)}
                <ItemRow
                  item={ref(alt.id, alt.name)}
                  {build}
                  source={alt.source}
                  gain={{ value: signedFigure(alt.delta), unit, tone: alt.delta < 0 ? 'wipe' : 'muted' }}
                  iconSize={32}
                  bare
                />
              {/each}
            {/if}
            {#if bisHref !== null}<a class="{LINK} mt-2 inline-block text-[14px]" href={bisHref}
                >Open in best in slot ›</a
              >{/if}
          {/snippet}
        </Inspector>
      {/snippet}
    </Paperdoll>
    {#if note !== null}
      <p class="text-muted px-4 py-3 text-[13px] leading-[18px]" data-testid="character-doll-note">{note}</p>
    {/if}
  </Plate>
</Section>
