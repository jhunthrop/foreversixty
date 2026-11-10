<!-- web/src/components/character/CharacterGear.svelte -->
<!-- The character's gear (round 5, directions mock b2-character): the page's framed panel holding
     the one shared paperdoll on the owner's frame (ui/PaperdollFrame). The title bar names the
     spec and the level; the medallion is the character's Battle.net head render, else the seal.
     The centre is the client tooltip ground: the Upgrades (the band's picks by their priced gain)
     until a slot is chosen, then that slot: what is worn, the band's pick with its source and two
     runners-up. A phone draws the frame's list form, the Upgrades under it. `cells` decides the
     words (lib/character/doll.ts); this lays them out. -->
<script lang="ts">
  import { isTie } from '../../lib/bis/doll-view';
  import type { BisBand } from '../../lib/bis/types';
  import { bandSlotFor, runnersUp } from '../../lib/character/gear-context';
  import { characterPageCopy } from '../../lib/character/page-view';
  import { UPGRADES_SHOWN, type UpgradeRowView } from '../../lib/character/upgrades';
  import { gearCopy } from '../../lib/gear/copy';
  import type { GearItemTable } from '../../lib/gear/qualities';
  import type { CharacterGearSlot } from '../../lib/gear/types';
  import { signedFigure, type ItemRef, LINK } from '../../lib/ui/primitives';
  import type { DollCell } from '../../lib/ui/paperdoll';
  import ItemRow from '../ui/ItemRow.svelte';
  import Overlay from '../ui/Overlay.svelte';
  import PaperdollFrame from '../ui/PaperdollFrame.svelte';
  import Section from '../ui/Section.svelte';

  let {
    sectionTitle = characterPageCopy.gear,
    title,
    meta,
    avatar = null,
    cells,
    gear,
    band,
    table,
    upgrades,
    unit,
    build,
    bisHref,
    loading = false,
    note = null,
  }: {
    /** The panel's heading: "Gear", or the unread state's "Best in slot for level 30 to 39". */
    sectionTitle?: string;
    /** The frame's title bar: the spec and class ("Holy Priest"). */
    title: string;
    /** After the title, small: "Level 60". */
    meta?: string;
    /** The character's Battle.net head render (`avatar_url`), else null for the seal. */
    avatar?: string | null;
    cells: readonly DollCell[];
    gear: readonly CharacterGearSlot[];
    band: BisBand | null;
    table: GearItemTable | null;
    /** Every pick with a priced gain, largest first (lib/character/upgrades.ts); the centre
     *  shows UPGRADES_SHOWN and counts the rest, so the rows and the band's total agree. */
    upgrades: readonly UpgradeRowView[];
    /** The band's rate word ("DPS", "HPS"). */
    unit: string;
    build: string;
    bisHref: string | null;
    /** The items file and the band are on their way: the frame waits empty, never a white name. */
    loading?: boolean;
    /** One muted line under the doll (items not in our data, a failed items file). */
    note?: string | null;
  } = $props();

  let selected = $state<string | null>(null);

  function ref(id: number, name: string): ItemRef {
    const facts = table?.get(id);
    return { id, name, quality: facts?.quality ?? 1, icon: facts?.icon ?? undefined };
  }

  const bySlot = $derived(new Map(gear.map((slot) => [slot.slot, slot])));
  /** What the rows past the centre's six add, so the list and the band's total agree on screen. */
  const hiddenGain = $derived(upgrades.slice(UPGRADES_SHOWN).reduce((sum, row) => sum + row.value, 0));
</script>

{#snippet gainText(value: number)}
  <span class="text-gain text-[12px] leading-4 font-bold whitespace-nowrap tabular-nums"
    >{signedFigure(value)} {unit}</span
  >
{/snippet}

{#snippet upgradeList()}
  <Overlay
    kind="tooltip"
    role="group"
    label={characterPageCopy.upgrades}
    class="w-full max-w-[340px] px-3 py-[10px]"
    testid="character-upgrades"
  >
    <h3 class="text-strong mb-1 text-[13px] leading-5 font-semibold">{characterPageCopy.upgrades}</h3>
    {#each upgrades.slice(0, UPGRADES_SHOWN) as row (row.slot)}
      <ItemRow item={row.item} {build} source={row.label} iconSize={28} bare>
        {#snippet trail()}{@render gainText(row.value)}{/snippet}
      </ItemRow>
    {/each}
    {#if upgrades.length > UPGRADES_SHOWN && bisHref !== null}<a
        class="{LINK} mt-1 inline-block text-[13px]"
        href={bisHref}
        data-testid="character-upgrades-more"
        >{characterPageCopy.moreUpgrades(
          upgrades.length - UPGRADES_SHOWN,
          `${signedFigure(hiddenGain)} ${unit}`,
        )} ›</a
      >{/if}
  </Overlay>
{/snippet}

{#snippet slotDetail(cell: DollCell)}
  {@const api = bySlot.get(cell.slot as CharacterGearSlot['slot'])}
  {@const pick = api?.pick ?? null}
  {@const row = bandSlotFor(band, cell.slot, pick?.id ?? null)}
  <Overlay
    kind="tooltip"
    role="group"
    label={cell.label}
    class="w-full max-w-[340px] px-3 py-[10px]"
    testid="character-inspector"
  >
    <h3 class="text-strong mb-1 text-[13px] leading-5 font-semibold">{cell.label}</h3>
    {#if cell.item}<ItemRow
        item={cell.item}
        {build}
        source={pick !== null && pick.id === cell.item.id ? pick.source : undefined}
        iconSize={28}
        bare
      />{/if}
    {#if pick !== null && pick.id !== cell.item?.id}
      <p class="text-muted mt-2 text-[12px] leading-4 font-semibold">{characterPageCopy.gearAction}</p>
      <ItemRow
        item={ref(pick.id, pick.name)}
        {build}
        source={pick.chance === null ? pick.source : `${pick.source} · ${gearCopy.chance(pick.chance)}`}
        iconSize={28}
        bare
        testid="character-inspector-pick"
      />
      {#each runnersUp(row) as alt (alt.id)}
        <ItemRow
          item={ref(alt.id, alt.name)}
          {build}
          source={isTie(alt.delta) ? alt.source : `${alt.source} · ${signedFigure(alt.delta)} ${unit}`}
          iconSize={28}
          bare
        />
      {/each}
    {/if}
    <span class="mt-2 flex flex-wrap gap-x-4">
      {#if upgrades.length > 0}<button
          type="button"
          class="{LINK} text-[13px]"
          onclick={() => (selected = null)}
          data-testid="character-upgrades-back">{characterPageCopy.allUpgrades}</button
        >{/if}
      {#if bisHref !== null}<a class="{LINK} text-[13px]" href={bisHref}>{characterPageCopy.openBis} ›</a
        >{/if}
    </span>
  </Overlay>
{/snippet}

{#snippet centre(cell: DollCell | undefined)}
  {#if cell !== undefined}{@render slotDetail(cell)}{:else if upgrades.length > 0}<div
      class="contents @min-[980px]/gear:hidden"
    >
      {@render upgradeList()}
    </div>{/if}
{/snippet}

<!-- Beside the doll on a wide panel: every upgrade, largest first, so the panel's width holds the
     whole list rather than ground either side of the doll. -->
{#snippet upgradeColumn()}
  <div class="flex h-full min-w-0 flex-col gap-1" role="group" aria-label={characterPageCopy.upgrades}>
    <h3 class="text-strong text-[15px] leading-6 font-semibold">{characterPageCopy.upgrades}</h3>
    {#each upgrades as row (row.slot)}
      <ItemRow item={row.item} {build} source={row.label} iconSize={32} bare>
        {#snippet trail()}{@render gainText(row.value)}{/snippet}
      </ItemRow>
    {/each}
    {#if bisHref !== null}<a class="{LINK} mt-auto self-start pt-2 text-[13px]" href={bisHref}
        >{characterPageCopy.openBis} ›</a
      >{/if}
  </div>
{/snippet}

<Section
  title={sectionTitle}
  panel
  frame="framed"
  action={bisHref === null ? undefined : { label: characterPageCopy.gearAction, href: bisHref }}
  testid="character-gear-section"
>
  <!-- The frame keeps to 640 px, centred. The Upgrades sit in its centre; where it gives way to its
       list form (under MIN_FRAME_PX, 423 px, lib/ui/paperdoll-frame.ts) they follow the list; on a
       panel of 980 px and up (the wide page frame) they stand beside it, every row, and the centre
       keeps the chosen slot. -->
  <div class="@container/gear">
    <div
      class="flex justify-center @min-[980px]/gear:grid @min-[980px]/gear:grid-cols-[640px_minmax(0,1fr)] @min-[980px]/gear:items-start @min-[980px]/gear:gap-8"
    >
      <div class="@container w-full max-w-[640px]" data-testid="character-gear" aria-busy={loading}>
        <PaperdollFrame
          cells={loading ? [] : cells}
          {build}
          {title}
          {meta}
          {avatar}
          bind:selected
          {centre}
          label={gearCopy.listLabel}
          testid="character-doll"
        />
        {#if upgrades.length > 0 && !loading}<div class="hidden @max-[423px]:mt-3 @max-[423px]:block">
            {@render upgradeList()}
          </div>{/if}
      </div>
      {#if upgrades.length > 0 && !loading}<div
          class="hidden min-w-0 @min-[980px]/gear:block @min-[980px]/gear:self-stretch"
          data-testid="character-upgrades-side"
        >
          {@render upgradeColumn()}
        </div>{/if}
    </div>
  </div>
  {#if note !== null}
    <p class="text-muted pt-3 text-[13px] leading-[18px]" data-testid="character-doll-note">{note}</p>
  {/if}
</Section>
