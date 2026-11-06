<!-- web/src/components/character/GearDoll.svelte -->
<!-- A worn set as the one shared paperdoll (web layout spec 2026-11-05 §4.2: `worn` on the
     character page and the sim, `report` on a report's "At pull"): the slots mapped to cells by
     lib/gear/doll-cells.ts, then the note (an item not in our data, or a failed items file) once
     under the doll. It draws no box; the caller's Plate holds it. Every item carries
     `data-bis-item`, so the page's tooltip host (BisTooltipHost) opens the real item. -->
<script lang="ts">
  import { dollCellsFromSlots, type WornDollMode } from '../../lib/gear/doll-cells';
  import { paperdollNote, unknownCount, type PaperdollSlot } from '../../lib/gear/paperdoll';
  import Paperdoll from '../ui/Paperdoll.svelte';

  let {
    slots,
    mode,
    build,
    classSlug = '',
    loading = false,
    failed = false,
    label,
    testid,
  }: {
    slots: readonly PaperdollSlot[];
    mode: WornDollMode;
    /** The data build the icons are read from. */
    build: string;
    /** The class, for its slot words ("Relic"). */
    classSlug?: string;
    /** The items file is pending: the doll's skeleton rows. */
    loading?: boolean;
    /** The items file failed: plain names, and the note says so. */
    failed?: boolean;
    label: string;
    testid: string;
  } = $props();

  const cells = $derived(dollCellsFromSlots(slots, mode, classSlug));
  const note = $derived(loading ? null : paperdollNote(unknownCount(slots), failed));
</script>

<Paperdoll {cells} {mode} {build} {loading} {label} {testid} />
{#if note !== null}
  <p class="text-muted px-4 py-3 text-[13px] leading-[18px]" data-testid={`${testid}-note`}>{note}</p>
{/if}
