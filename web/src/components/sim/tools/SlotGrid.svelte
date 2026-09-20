<!-- web/src/components/sim/tools/SlotGrid.svelte -->
<!-- Design 3.1.2: one section per slot, each with its candidates and a lock. Rings and
     trinkets get one section, not two: sim/bulk tries them in both slots and the grid
     saying otherwise would be a promise the engine does not keep. -->
<script lang="ts">
  import { SLOTS, SLOT_LABELS, type Slot } from '../../../lib/planner/types';
  import type { CandidateRow } from '../../../lib/sim/candidates';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { EnchantRow, SuffixRow } from '../../../lib/sim/enchants';
  import CandidateRows from './CandidateRows.svelte';

  let {
    rows,
    locked,
    enchants,
    suffixes,
    classSlug,
    treeVersion,
    ontoggle,
    oncopy,
    onlock,
  }: {
    rows: readonly CandidateRow[];
    locked: readonly string[];
    enchants: readonly EnchantRow[];
    suffixes: readonly SuffixRow[];
    classSlug: string;
    treeVersion: string;
    ontoggle: (key: string) => void;
    oncopy: (key: string, patch: { enchant?: number; suffix?: number }) => void;
    onlock: (slot: Slot) => void;
  } = $props();

  /**
   * finger2 and trinket2 never get a section of their own -- uiSlotsOf collapses a ring or
   * trinket to the first alias slot, so no row ever carries these two, and the grid never
   * offers a lock for a section that can never have candidates.
   */
  const SHOWN: readonly Slot[] = SLOTS.filter((slot) => slot !== 'finger2' && slot !== 'trinket2');

  const bySlot = $derived(new Map(SHOWN.map((slot) => [slot, rows.filter((row) => row.slot === slot)])));
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-slot-grid">
  {#each SHOWN as slot (slot)}
    {@const slotRows = bySlot.get(slot) ?? []}
    {#if slotRows.length > 0}
      <div class="border-line rounded-panel border p-3">
        <header class="flex flex-wrap items-center justify-between gap-2">
          <h3 class="section-title text-[14px]">{SLOT_LABELS[slot]}</h3>
          <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
            <input
              type="checkbox"
              class="h-5 w-5"
              data-testid={`sim-lock-${slot}`}
              checked={locked.includes(slot)}
              onchange={() => onlock(slot)}
            />
            {locked.includes(slot) ? bulkCopy.lockedSlot : bulkCopy.lockSlot}
          </label>
        </header>
        <CandidateRows
          rows={slotRows}
          {slot}
          locked={locked.includes(slot)}
          {enchants}
          {suffixes}
          {classSlug}
          {treeVersion}
          {ontoggle}
          {oncopy}
        />
      </div>
    {/if}
  {/each}
</section>
