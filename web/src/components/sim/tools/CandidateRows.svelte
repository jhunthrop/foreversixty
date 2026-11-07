<!-- web/src/components/sim/tools/CandidateRows.svelte -->
<!-- One slot's candidates: a checkbox, an icon, a name, an item level and a copy-and-modify
     menu, grouped by where each came from. A locked slot disables every checkbox in it
     rather than hiding the rows -- the player has to see what they locked away. -->
<script lang="ts">
  import ItemHover from '../../ItemHover.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Row from '../../ui/Row.svelte';
  import { rarityClassFor } from '../../../lib/planner/items';
  import { button } from '../../../lib/ui/button';
  import type { Slot } from '../../../lib/planner/types';
  import { candidateKey, type CandidateRow } from '../../../lib/sim/candidates';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { suffixesForItem, type EnchantRow, type SuffixRow } from '../../../lib/sim/enchants';
  import EnchantList from './EnchantList.svelte';

  let {
    rows,
    slot,
    locked,
    enchants,
    suffixes,
    classSlug,
    treeVersion,
    ontoggle,
    oncopy,
  }: {
    rows: readonly CandidateRow[];
    slot: Slot;
    locked: boolean;
    enchants: readonly EnchantRow[];
    suffixes: readonly SuffixRow[];
    classSlug: string;
    treeVersion: string;
    ontoggle: (key: string) => void;
    oncopy: (key: string, patch: { enchant?: number; suffix?: number }) => void;
  } = $props();

  /** Which row's copy-and-modify menu is open, by key. Only ever one. */
  let openKey = $state<string | null>(null);

  const ORIGIN_LABELS: Record<string, string> = {
    equipped: bulkCopy.equipped,
    bag: bulkCopy.bags,
    bank: bulkCopy.bank,
    search: bulkCopy.fromSearch,
  };

  /**
   * `drop:<source-id>` and `set:<name>` both carry a wire-only value after the colon, never
   * a display string, so both fall back to a generic word instead of leaking it into the
   * UI -- `set:` has no candidate rows yet (Task 14's NamedSets adds the first), but the
   * `Origin` type already declares it, so this handles it before anything produces one.
   */
  function originLabel(origin: string): string {
    const known = ORIGIN_LABELS[origin];
    if (known !== undefined) return known;
    if (origin.startsWith('drop:')) return bulkCopy.pinned;
    if (origin.startsWith('set:')) return bulkCopy.fromSet;
    return origin;
  }

  /**
   * The row's own test id. A copy carries its enchant and suffix, so the original and every
   * copy of it are separately addressable -- which is the whole point of copy-and-modify.
   */
  function testId(row: CandidateRow): string {
    const base = `sim-candidate-${row.slot}-${row.item.id}`;
    const parts = [row.enchant > 0 ? `e${row.enchant}` : '', row.suffix > 0 ? `s${row.suffix}` : ''];
    return [base, ...parts.filter((part) => part !== '')].join('-');
  }
</script>

<div class="flex flex-col">
  {#each rows as row (candidateKey(row))}
    {@const key = candidateKey(row)}
    <Row title={row.item.name} open wrapTrail testid={testId(row)}>
      {#snippet lead()}
        <!-- The label is the 44 px hit target around the 20 px mark. -->
        <label class="-mx-3 inline-flex min-h-11 min-w-11 items-center justify-center">
          <input
            type="checkbox"
            class="accent-gold h-5 w-5"
            aria-label={row.item.name}
            checked={row.checked}
            disabled={locked || !row.known}
            onchange={() => ontoggle(key)}
          />
        </label>
        <ItemIcon icon={row.item.icon} quality={row.item.quality} build={treeVersion} size={32} />
      {/snippet}
      {#snippet titleContent()}
        <ItemHover itemId={row.item.id} {classSlug} build={treeVersion}>
          <span class={rarityClassFor(row.item.quality)}>{row.item.name}</span>
        </ItemHover>
      {/snippet}
      {#snippet subContent()}
        {originLabel(row.origin)}
        <!-- The engine's SimDatabase does not carry this item id (sim-items.ts's
             `isKnownItem`, set on the row when it was added). The checkbox is disabled
             rather than the row hidden -- the item is real, and vanishing it would look like
             data loss rather than naming a known limitation. -->
        {#if !row.known}
          <span data-testid={`sim-candidate-unknown-${row.slot}-${row.item.id}`}>
            · {bulkCopy.notInSimulator}
          </span>
        {/if}
      {/snippet}
      {#snippet trail()}
        <span class="text-muted text-[13px] tabular-nums">{row.item.item_level}</span>
        <button
          type="button"
          class={button('secondary', 'sm')}
          aria-expanded={openKey === key}
          disabled={locked || !row.known}
          onclick={() => (openKey = openKey === key ? null : key)}
        >
          {bulkCopy.copyAndModify}
        </button>
      {/snippet}
    </Row>
    <!-- A locked slot's checkbox is disabled and `toCandidates` drops anything ticked on
           it, so a copy made here would be checked-but-inert with nothing telling the
           player. The trigger is disabled for the same reason; this closes the panel too,
           for the case where a slot is locked while it is already open. -->
    {#if openKey === key && !locked}
      {@const itemSuffixes = suffixesForItem(suffixes, row.item)}
      <div class="flex w-full flex-col gap-4 pb-2 md:flex-row">
        <div class="flex-1">
          <p class="text-muted text-[13px] font-bold">{bulkCopy.withEnchant}</p>
          <EnchantList
            rows={enchants}
            {slot}
            {classSlug}
            item={row.item}
            {treeVersion}
            onpick={(enchant) => {
              oncopy(key, { enchant });
              openKey = null;
            }}
          />
        </div>
        {#if itemSuffixes.length > 0}
          <div class="flex-1">
            <p class="text-muted text-[13px] font-bold">{bulkCopy.withSuffix}</p>
            <ul class="flex flex-col">
              {#each itemSuffixes as suffix (suffix.id)}
                <li>
                  <button
                    type="button"
                    class="text-text hover:text-gold flex min-h-11 w-full items-center text-left text-[14px]"
                    data-testid={`sim-suffix-${row.slot}-${row.item.id}-${suffix.id}`}
                    onclick={() => {
                      oncopy(key, { suffix: suffix.id });
                      openKey = null;
                    }}>{suffix.name}</button
                  >
                </li>
              {/each}
            </ul>
          </div>
        {/if}
      </div>
    {/if}
  {/each}
</div>
