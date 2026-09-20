<!-- web/src/components/sim/tools/NamedSets.svelte -->
<!-- Design 3.1.7 and the decision that killed Gear Compare: a whole-set alternative is one
     candidate that replaces every slot at once. Sets come from the addon export (part A's
     FS1 v2 `sets=` section) or from a second export string pasted here. -->
<script lang="ts">
  import { decodeFS1, type FS1GearSlot, type FS1Set } from '../../../lib/planner/fs1';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { GearSet } from '../../../lib/sim/bulk-types';
  import type { SimCharacter } from '../../../lib/sim/character';
  import { bulkCopy } from '../../../lib/sim/copy';

  let {
    character,
    sets,
    onadd,
    onremove,
  }: {
    character: SimCharacter;
    sets: readonly GearSet[];
    onadd: (set: GearSet) => void;
    onremove: (name: string) => void;
  } = $props();

  let code = $state('');
  let name = $state('');
  let error = $state('');

  /**
   * A decoder gear entry (the addon export's `FS1GearSlot`, or `decodeFS1`'s own return
   * shape for a pasted string) as the engine's `GearSlot` -- contract 10.5's `item_id`,
   * enchant and suffix carried through unchanged, the same conversion `character.ts` makes
   * for the character's own equipped gear.
   */
  function toGearSlot(entry: FS1GearSlot): GearSet['gear'][number] {
    return {
      slot: entry.slot,
      item_id: entry.itemId,
      ...(entry.enchant === undefined ? {} : { enchant: entry.enchant }),
      ...(entry.suffix === undefined ? {} : { suffix: entry.suffix }),
    };
  }

  function toGearSet(set: FS1Set): GearSet {
    return { name: set.name, gear: set.gear.map(toGearSlot) };
  }

  /** The export's own named sets (part A's FS1 v2 decoder), offered as one-click adds. */
  const exported = $derived<GearSet[]>(character.sets.map(toGearSet));

  /** Every row to show: the export's own sets first, then any ticked set the export did not
   *  already name -- a set the player pasted in here themselves. A set present in both is
   *  shown once, from `exported`: `add()` below hands `onadd` the identical conversion, so
   *  the two never disagree about what that name means. */
  const rows = $derived<GearSet[]>([
    ...exported,
    ...sets.filter((set) => !exported.some((entry) => entry.name === set.name)),
  ]);

  function isAdded(setName: string): boolean {
    return sets.some((entry) => entry.name === setName);
  }

  function add(): void {
    error = '';
    const decoded = decodeFS1(code.trim());
    if (!decoded.ok) {
      error = bulkCopy.setsBadCode;
      return;
    }
    const setName = name.trim() === '' ? `Set ${sets.length + 1}` : name.trim();
    onadd(toGearSet({ name: setName, gear: decoded.build.gearSlots }));
    code = '';
    name = '';
  }
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0"
  data-testid="sim-named-sets"
>
  <h3 class="section-title text-[14px]">{bulkCopy.setsTitle}</h3>
  <p class="text-muted text-[12px]">{bulkCopy.setsIntro}</p>

  {#each rows as set (set.name)}
    <div class="flex min-h-11 items-center gap-2 text-[13px]" data-testid={`sim-set-${set.name}`}>
      <span class="text-text flex-1">{set.name}</span>
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
        onclick={() => (isAdded(set.name) ? onremove(set.name) : onadd(set))}
      >
        {isAdded(set.name) ? bulkCopy.setsRemove : bulkCopy.setsAdd}
      </button>
    </div>
  {/each}

  <div class="flex flex-wrap gap-2">
    <input
      type="text"
      bind:value={name}
      placeholder="Name"
      aria-label="Set name"
      data-testid="sim-set-name"
      class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 w-40 border px-3 text-[14px]"
    />
    <input
      type="text"
      bind:value={code}
      placeholder={bulkCopy.setsPaste}
      aria-label={bulkCopy.setsPaste}
      data-testid="sim-set-input"
      class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 min-w-[200px] flex-1 border px-3 text-[14px]"
    />
    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
      data-testid="sim-set-add"
      onclick={add}>{bulkCopy.setsAdd}</button
    >
  </div>
  {#if error !== ''}
    <p class="text-muted text-[13px]" role="alert" data-testid="sim-set-error">{error}</p>
  {/if}
</section>
