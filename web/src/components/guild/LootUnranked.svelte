<!-- web/src/components/guild/LootUnranked.svelte -->
<!-- The drops nobody ranks, one quiet line at the bottom of the Loot tab (guild loot spec 2026-11-05
     §3.5): quest drops, the bag, the skinned scale and a slotted drop nobody can use yet. A quest
     item still goes to a person, so an officer can award a quest drop or the bag to a character the
     quest is for: Award opens the same ask-first confirm, with a picker of those characters. -->
<script lang="ts">
  import type { GuildLootItem } from '../../lib/guild/api';
  import { UNDO_AWARD, awardedLine, unrankedLabel, unrankedLine } from '../../lib/guild/loot-view';
  import { button } from '../../lib/ui/button';
  import Select from '../ui/Select.svelte';
  import RaidDisclosure from './RaidDisclosure.svelte';
  import RaidLootItem from './RaidLootItem.svelte';

  let {
    items,
    officer,
    readOnly,
    gateDescribedBy,
    busyItemId,
    onAward,
    onUndo,
  }: {
    items: readonly GuildLootItem[];
    officer: boolean;
    /** The officer's tools are off: the controls stay, disabled. */
    readOnly: boolean;
    gateDescribedBy: string | undefined;
    busyItemId: number | null;
    onAward: (item: GuildLootItem, characterKey: string, itemId: number) => void;
    onUndo: (item: GuildLootItem) => void;
  } = $props();

  /** The drop whose Award is open, and who is picked. */
  let picking = $state<{ itemId: number; key: string } | null>(null);

  function open(item: GuildLootItem): void {
    picking = { itemId: item.item_id, key: item.eligible[0]?.character_key ?? '' };
  }

  function confirm(item: GuildLootItem): void {
    const chosen = item.eligible.find((row) => row.character_key === picking?.key);
    picking = null;
    if (chosen !== undefined) onAward(item, chosen.character_key, chosen.item_id);
  }

  function focusOnMount(node: HTMLElement): void {
    node.focus();
  }
</script>

<RaidDisclosure
  heading={unrankedLine(items.length)}
  headingId="guild-loot-unranked-heading"
  headingClass="text-text text-[14px] font-normal"
  testid="guild-loot-unranked"
>
  <ul class="flex flex-col">
    {#each items as item (item.item_id)}
      <li
        class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 gap-y-1 border-t py-2"
        data-testid={`guild-loot-unranked-${item.item_id}`}
      >
        <RaidLootItem
          award={{
            item_id: item.item_id,
            item_name: item.name,
            quality: item.quality,
            icon: item.icon,
            character_key: '',
            name: '',
            class: item.eligible[0]?.class ?? '',
            by_name: '',
            at: '',
          }}
          iconSize={24}
          nameSize="text-[14px]"
        />
        <span class="text-muted text-[13px]" data-testid={`guild-loot-unranked-label-${item.item_id}`}
          >{unrankedLabel(item)}</span
        >
        <span class="ml-auto flex items-center gap-2">
          {#if item.awarded_to !== null}
            <span class="text-muted text-[12px]" data-testid={`guild-loot-awarded-${item.item_id}`}
              >{awardedLine(item.awarded_to)}</span
            >
            {#if officer}
              <button
                class="inline-flex h-11 items-center text-[12px] underline disabled:cursor-default disabled:no-underline lg:h-7"
                onclick={() => onUndo(item)}
                disabled={readOnly || busyItemId === item.item_id}
                aria-describedby={gateDescribedBy}
                aria-label={`Undo the award of ${item.name}`}
                data-testid={`guild-loot-row-undo-${item.item_id}`}>{UNDO_AWARD}</button
              >
            {/if}
          {:else if officer && item.eligible.length > 0}
            <button
              class={button('secondary', 'sm')}
              onclick={() => open(item)}
              disabled={readOnly || busyItemId === item.item_id || picking !== null}
              aria-describedby={gateDescribedBy}
              aria-label={`Award ${item.name}`}
              data-testid={`guild-loot-award-${item.item_id}`}
            >
              Award
            </button>
          {/if}
        </span>
        <!-- Escape backs out of the confirm from any control in it; the group itself takes no action. -->
        <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
        {#if picking !== null && picking.itemId === item.item_id}
          <div
            class="flex basis-full flex-wrap items-end gap-x-4 gap-y-2 pt-1"
            role="group"
            aria-label={`Award ${item.name}`}
            onkeydown={(event) => {
              if (event.key === 'Escape') picking = null;
            }}
            data-testid={`guild-loot-award-confirm-${item.item_id}`}
          >
            <div class="w-64 text-left">
              <Select
                id={`guild-loot-pick-${item.item_id}`}
                size="compact"
                label="Award to"
                options={item.eligible.map((row) => ({ value: row.character_key, label: row.name }))}
                value={picking.key}
                onchange={(event) => {
                  picking = {
                    itemId: item.item_id,
                    key: (event.currentTarget as HTMLSelectElement).value,
                  };
                }}
              />
            </div>
            <button
              class={button('primary', 'sm')}
              onclick={() => confirm(item)}
              data-testid={`guild-loot-award-yes-${item.item_id}`}
            >
              Award
            </button>
            <button
              class="{button('ghost', 'sm')} focus:outline-gold focus:outline-2 focus:outline-offset-2"
              onclick={() => (picking = null)}
              use:focusOnMount
              data-testid={`guild-loot-award-cancel-${item.item_id}`}
            >
              Cancel
            </button>
          </div>
        {/if}
      </li>
    {/each}
  </ul>
</RaidDisclosure>
