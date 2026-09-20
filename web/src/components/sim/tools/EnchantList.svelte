<!-- web/src/components/sim/tools/EnchantList.svelte -->
<!-- The enchants one slot allows, as a menu behind "copy and modify" and as the per-slot
     checklist design 3.1.4 asks for. "Keep current" and "None" are rows, not a separate
     control, because they are two of the answers to the same question. -->
<script lang="ts">
  import { dataUrl } from '../../../lib/planner/load';
  import {
    ENCHANTS_PER_SLOT_CAP,
    KEEP_CURRENT_ENCHANT,
    NO_ENCHANT,
    enchantsForSlot,
    type EnchantRow,
  } from '../../../lib/sim/enchants';
  import { bulkCopy } from '../../../lib/sim/copy';

  let {
    rows,
    slot,
    classSlug,
    treeVersion,
    onpick,
  }: {
    rows: readonly EnchantRow[];
    slot: string;
    classSlug: string;
    treeVersion: string;
    onpick: (enchantId: number) => void;
  } = $props();

  const allowed = $derived(enchantsForSlot(rows, slot, classSlug).slice(0, ENCHANTS_PER_SLOT_CAP));
</script>

<ul class="border-line bg-raised rounded-panel flex flex-col border p-2" data-testid={`sim-enchants-${slot}`}>
  <li>
    <button
      type="button"
      class="text-text flex min-h-11 w-full items-center px-2 text-left text-[13px]"
      onclick={() => onpick(KEEP_CURRENT_ENCHANT)}>{bulkCopy.keepCurrentEnchant}</button
    >
  </li>
  <li>
    <button
      type="button"
      class="text-text flex min-h-11 w-full items-center px-2 text-left text-[13px]"
      onclick={() => onpick(NO_ENCHANT)}>{bulkCopy.noEnchant}</button
    >
  </li>
  {#each allowed as enchant (enchant.id)}
    <li class="border-line-soft border-t">
      <button
        type="button"
        class="text-text flex min-h-11 w-full items-center gap-2 px-2 text-left text-[13px]"
        data-testid={`sim-enchant-${slot}-${enchant.id}`}
        onclick={() => onpick(enchant.id)}
      >
        <img
          src={dataUrl(treeVersion, `icons/${enchant.icon}.webp`)}
          alt=""
          width="20"
          height="20"
          loading="lazy"
          decoding="async"
          class="rounded-control border-line h-5 w-5 border object-cover"
        />
        {enchant.name}
      </button>
    </li>
  {/each}
  {#if allowed.length === 0}
    <li class="text-muted px-2 py-2 text-[13px]">{bulkCopy.noEnchantsAvailable}</li>
  {/if}
</ul>
<p class="text-muted px-2 text-[12px]">{bulkCopy.enchantCap(ENCHANTS_PER_SLOT_CAP)}</p>
