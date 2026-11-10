<!-- web/src/components/sim/tools/EnchantList.svelte -->
<!-- The enchants one slot allows, as a menu behind "copy and modify" and as the per-slot
     checklist design 3.1.4 asks for. "Keep current" and "None" are rows, not a separate
     control, because they are two of the answers to the same question. -->
<script lang="ts">
  import { iconUrl } from '../../../lib/ui/assets';
  import SpellIcon from '../../ui/SpellIcon.svelte';
  import {
    ENCHANTS_PER_SLOT_CAP,
    KEEP_CURRENT_ENCHANT,
    NO_ENCHANT,
    enchantsForSlot,
    type EnchantRow,
  } from '../../../lib/sim/enchants';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { Item } from '../../../lib/planner/types';

  let {
    rows,
    slot,
    classSlug,
    item,
    treeVersion,
    onpick,
  }: {
    rows: readonly EnchantRow[];
    slot: string;
    classSlug: string;
    /** The item the enchant would go on: only enchants the client lets go on it are offered. */
    item?: Item;
    treeVersion: string;
    onpick: (enchantId: number) => void;
  } = $props();

  /** One answer's line: 44 px, the body ink, gold on hover. */
  const CHOICE = 'text-text hover:text-gold flex min-h-11 w-full items-center gap-2 text-left text-[14px]';

  const allowed = $derived(enchantsForSlot(rows, slot, classSlug, item).slice(0, ENCHANTS_PER_SLOT_CAP));
</script>

<!-- A plain menu of choices: no box of its own (it opens under a candidate row), one 44 px line per
     answer, the hover in gold. -->
<ul class="flex flex-col" data-testid={`sim-enchants-${slot}`}>
  <li>
    <button type="button" class={CHOICE} onclick={() => onpick(KEEP_CURRENT_ENCHANT)}
      >{bulkCopy.keepCurrentEnchant}</button
    >
  </li>
  <li>
    <button type="button" class={CHOICE} onclick={() => onpick(NO_ENCHANT)}>{bulkCopy.noEnchant}</button>
  </li>
  {#each allowed as enchant (enchant.id)}
    <li>
      <button
        type="button"
        class={CHOICE}
        data-testid={`sim-enchant-${slot}-${enchant.id}`}
        onclick={() => onpick(enchant.id)}
      >
        <SpellIcon url={iconUrl(treeVersion, enchant.icon)} size={28} />
        {enchant.name}
      </button>
    </li>
  {/each}
  {#if allowed.length === 0}
    <li class="text-muted py-2 text-[13px]">{bulkCopy.noEnchantsAvailable}</li>
  {/if}
</ul>
<p class="text-muted text-[12px]">{bulkCopy.enchantCap(ENCHANTS_PER_SLOT_CAP)}</p>
