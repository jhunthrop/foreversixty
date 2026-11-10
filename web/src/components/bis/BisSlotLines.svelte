<!-- web/src/components/bis/BisSlotLines.svelte -->
<!-- What the selected slot is worth (design/specs/2026-11-09-directions.md, Paperdoll): "+0.3 DPS vs
     Cape of the Brotherhood" (against the worn item, signed in), why a slot was re-picked, or why
     an off hand stays empty. It ends the item's client tooltip in the frame's centre; signed in it
     sits in the bar under the frame with Open in the sim, so the character's model shows. Words
     only: the caller decides the surface, so quality colours always sit on a dark panel. -->
<script lang="ts">
  import { rarityClassFor } from '../../lib/planner/items';
  import type { DollCell } from '../../lib/ui/paperdoll';
  import { bisCopy } from '../../lib/bis/copy';
  import { isMeasuredGain } from '../../lib/bis/frame-view';

  let { cell }: { cell: DollCell | undefined } = $props();

  const gain = $derived(cell?.gain);
  const repick = $derived(cell?.note?.label === 'You wear' ? undefined : cell?.note);
  /** A measured gain leads with its figure; a tie, an only option or a worn item says its basis
   *  alone ("No close alternative", "−1.8 DPS vs ...", "You have this"). */
  const measured = $derived(isMeasuredGain(gain?.value));

  function sentence(words: string): string {
    return words.charAt(0).toUpperCase() + words.slice(1);
  }
</script>

{#if cell?.item && gain}
  <p data-testid="bis-slot-gain" class:bis-lines-basis={!measured}>
    {#if measured}<b class="bis-lines-fig">{gain.value}{gain.unit ? ` ${gain.unit}` : ''}</b>{/if}
    {#if gain.basis}{measured ? gain.basis.prefix : sentence(gain.basis.prefix)}
      {#if gain.basis.item}<span class={rarityClassFor(gain.basis.item.quality)}>{gain.basis.item.name}</span
        >{/if}{/if}
  </p>
{:else if cell?.coveredBy}
  <p data-testid="bis-slot-gain">{bisCopy.offHandCovered(cell.coveredBy.name)}</p>
{/if}
{#if repick}
  <p data-testid="bis-slot-repick">
    {repick.lead}
    {#if repick.item}<span class={rarityClassFor(repick.item.quality)}>{repick.item.name}</span>{/if}
    {#if repick.tail}{repick.tail}{/if}
  </p>
{/if}

<style>
  /* Only a measured gain takes the tooltip's gain colour; a tie's next alternative ("−1.8 DPS vs
     ...") is a reference, not a gain, so it reads muted (the item name keeps its quality). */
  .bis-lines-basis {
    color: var(--color-muted);
  }
  .bis-lines-fig {
    font-weight: 600;
    font-variant-numeric: tabular-nums;
  }
</style>
