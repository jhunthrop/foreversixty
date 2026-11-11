<!-- web/src/components/character/GearDoll.svelte -->
<!-- A worn set on the shared paperdoll frame (ui/PaperdollFrame; web layout spec 2026-11-05 §4.2
     `worn` on the sim, `report` on a report's Gear at the pull): the slots mapped to cells by
     lib/gear/doll-cells.ts, then the note (an item not in our data, or a failed items file) once
     under the doll.
     - The frame (`form="auto"`): the selected slot's client tooltip stands in the centre, the worn
       enchant on its green line as the client prints it, ending on the worn doll's upgrade ("+1.8
       DPS vs <pick>", "Best in slot") when there is one. It opens on the first upgrade, else the
       first item. Under the frame's minimum it is the list form.
     - The list (either form): every slot as "Head · Item level 61 · Crusader" (the enchant green)
       and its gain; the page's tooltip host opens each item by the row it was tapped on. -->
<script lang="ts">
  import ItemName from '../ui/ItemName.svelte';
  import { dollCellsFromSlots, type WornDollMode } from '../../lib/gear/doll-cells';
  import { paperdollNote, unknownCount, type PaperdollSlot } from '../../lib/gear/paperdoll';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import { defaultSelection, type DollCell } from '../../lib/ui/paperdoll';
  import { FACT_SEPARATOR, type GainFigure } from '../../lib/ui/primitives';
  import type ItemTooltipComponent from '../ItemTooltip.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import Overlay from '../ui/Overlay.svelte';
  import PaperdollFrame from '../ui/PaperdollFrame.svelte';

  let {
    slots,
    mode,
    build,
    classSlug = '',
    loading = false,
    failed = false,
    title,
    meta,
    figure,
    avatar = null,
    form = 'auto',
    head = true,
    label,
    testid,
  }: {
    slots: readonly PaperdollSlot[];
    mode: WornDollMode;
    /** The data build the icons and tooltips are read from. */
    build: string;
    /** The class, for its slot words ("Relic") and its items file. */
    classSlug?: string;
    /** The items file is pending: the frame waits empty, never a white name. */
    loading?: boolean;
    /** The items file failed: plain names, and the note says so. */
    failed?: boolean;
    /** The frame's title bar: the character's name. */
    title: string;
    /** After the title, small: the spec. */
    meta?: string;
    /** At the bar's right: the figure and its unit. */
    figure?: { value: string; unit: string };
    avatar?: string | null;
    form?: 'auto' | 'list';
    head?: boolean;
    label: string;
    testid: string;
  } = $props();

  const cells = $derived(loading ? [] : dollCellsFromSlots(slots, mode, classSlug));
  const note = $derived(loading ? null : paperdollNote(unknownCount(slots), failed));
  let chosen = $state<string | null>(null);
  const shownSlot = $derived(
    chosen !== null && cells.some((cell) => cell.slot === chosen) ? chosen : defaultSelection(cells),
  );
  const shownCell = $derived(cells.find((cell) => cell.slot === shownSlot));

  /** The client's compound enchant ("Defense +7/Stamina +10/Block Value +15") breaks only at its
   *  slashes in the frame's narrow centre: each part keeps its words together. */
  const ZERO_WIDTH_SPACE = String.fromCharCode(0x200b);
  const NO_BREAK_SPACE = String.fromCharCode(0xa0);
  function breakable(enchant: string | undefined): string | undefined {
    if (enchant === undefined || !enchant.includes('/')) return enchant;
    return enchant
      .split('/')
      .map((part) => part.replaceAll(' ', NO_BREAK_SPACE))
      .join(`/${ZERO_WIDTH_SPACE}`);
  }

  /** The client tooltips fetched so far, by item id; null where the items file has no row. The
   *  tooltip and its lookup load on first use, off the report and sim islands' startup. */
  let tips = $state<Record<number, ItemTooltipModel | null>>({});
  let ItemTooltip = $state.raw<typeof ItemTooltipComponent | null>(null);
  $effect(() => {
    const item = shownCell?.item;
    if (form === 'list' || !item || shownCell?.unknown || item.id in tips || classSlug === '') return;
    const id = item.id;
    void Promise.all([import('../../lib/items/lookup'), import('../ItemTooltip.svelte')])
      .then(([{ fetchItemTooltipModel }, tooltip]) => {
        ItemTooltip = tooltip.default;
        return fetchItemTooltipModel(build, classSlug, id);
      })
      .then(
        (model) => (tips = { ...tips, [id]: model }),
        (error: unknown) => {
          console.error('gear doll: the item tooltip did not load', error);
          tips = { ...tips, [id]: null };
        },
      );
  });
</script>

{#snippet gainLine(gain: GainFigure)}
  <p data-testid={`${testid}-gain`}>
    <b class="tabular-nums">{gain.value}{gain.unit ? ` ${gain.unit}` : ''}</b>
    {#if gain.basis}{gain.basis.prefix}
      {#if gain.basis.item}<ItemName
          name={gain.basis.item.name}
          quality={gain.basis.item.quality}
        />{/if}{/if}
  </p>
{/snippet}

{#snippet centre(cell: DollCell | undefined)}
  {#if cell?.item}
    {@const tip = tips[cell.item.id]}
    {@const gain = cell.gain}
    {#if tip && ItemTooltip && gain}
      <ItemTooltip model={tip} {build} floating={false} enchant={breakable(cell.enchant)}>
        {#snippet tail()}{@render gainLine(gain)}{/snippet}
      </ItemTooltip>
    {:else if tip && ItemTooltip}
      <ItemTooltip model={tip} {build} floating={false} enchant={breakable(cell.enchant)} />
    {:else}
      <Overlay kind="tooltip" role="group" label={cell.label} class="w-full px-3 py-[10px]">
        <ItemRow
          item={cell.item}
          {build}
          source={[cell.source, cell.enchant].filter(Boolean).join(FACT_SEPARATOR)}
          unknown={cell.unknown}
          iconSize={28}
          bare
        />
      </Overlay>
    {/if}
  {/if}
{/snippet}

<PaperdollFrame
  {cells}
  {build}
  {title}
  {meta}
  {figure}
  {avatar}
  {form}
  {head}
  facts
  listCentre={false}
  selected={form === 'list' ? null : shownSlot}
  onselect={(slot) => (chosen = slot)}
  centre={form === 'list' ? undefined : centre}
  {label}
  {testid}
/>
{#if note !== null}
  <p class="text-muted pt-3 text-[13px] leading-[18px]" data-testid={`${testid}-note`}>{note}</p>
{/if}
