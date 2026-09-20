<!-- web/src/components/sim/tools/CandidateRows.svelte -->
<!-- One slot's candidates: a checkbox, an icon, a name, an item level and a copy-and-modify
     menu, grouped by where each came from. A locked slot disables every checkbox in it
     rather than hiding the rows -- the player has to see what they locked away. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
  import { dataUrl } from '../../../lib/planner/load';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
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

  function originLabel(origin: string): string {
    return ORIGIN_LABELS[origin] ?? (origin.startsWith('drop:') ? bulkCopy.pinned : origin);
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

<ul class="flex flex-col">
  {#each rows as row (candidateKey(row))}
    <li
      class="border-line-soft flex flex-wrap items-center gap-2 border-b px-2 py-1 last:border-b-0"
      data-testid={testId(row)}
    >
      <label class="flex min-h-11 flex-1 items-center gap-3">
        <input
          type="checkbox"
          class="h-5 w-5"
          checked={row.checked}
          disabled={locked}
          onchange={() => ontoggle(candidateKey(row))}
        />
        <img
          src={dataUrl(treeVersion, `icons/${row.item.icon}.webp`)}
          alt=""
          width="28"
          height="28"
          loading="lazy"
          decoding="async"
          class="rounded-control border-line h-7 w-7 border object-cover"
        />
        <span class={`text-[14px] font-semibold ${rarityClassFor(row.item.quality)}`}>
          {row.item.name}
        </span>
        <span class="text-muted text-[12px]">{originLabel(row.origin)}</span>
        <span class="tabular text-muted ml-auto font-mono text-[12px]">{row.item.item_level}</span>
      </label>
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
        aria-expanded={openKey === candidateKey(row)}
        onclick={() => (openKey = openKey === candidateKey(row) ? null : candidateKey(row))}
      >
        {bulkCopy.copyAndModify}
      </button>
      {#if openKey === candidateKey(row)}
        <div class="flex w-full flex-col gap-2 pb-2 md:flex-row">
          <div class="flex-1">
            <p class="text-muted px-2 text-[12px]">{bulkCopy.withEnchant}</p>
            <EnchantList
              rows={enchants}
              {slot}
              {classSlug}
              {treeVersion}
              onpick={(enchant) => {
                oncopy(candidateKey(row), { enchant });
                openKey = null;
              }}
            />
          </div>
          {#if suffixesForItem(suffixes, row.item).length > 0}
            <div class="flex-1">
              <p class="text-muted px-2 text-[12px]">{bulkCopy.withSuffix}</p>
              <ul class="border-line bg-raised rounded-panel flex flex-col border p-2">
                {#each suffixesForItem(suffixes, row.item) as suffix (suffix.id)}
                  <li>
                    <button
                      type="button"
                      class="text-text flex min-h-11 w-full items-center px-2 text-left text-[13px]"
                      data-testid={`sim-suffix-${row.slot}-${row.item.id}-${suffix.id}`}
                      onclick={() => {
                        oncopy(candidateKey(row), { suffix: suffix.id });
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
    </li>
  {/each}
  {#if rows.length === 0}
    <li class="text-muted px-2 py-2 text-[13px]">{bulkCopy.noCandidates}</li>
  {/if}
</ul>
