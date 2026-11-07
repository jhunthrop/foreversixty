<!-- web/src/components/sim/tools/NamedSets.svelte -->
<!-- Design 3.1.7 and the decision that killed Gear Compare: a whole-set alternative is one
     candidate that replaces every slot at once. Sets come from the addon export (part A's
     FS1 v2 `sets=` section) or from a second export string pasted here. -->
<script lang="ts">
  import { decodeFS1, type FS1GearSlot, type FS1Set } from '../../../lib/planner/fs1';
  import type { GearSet } from '../../../lib/sim/bulk-types';
  import type { SimCharacter } from '../../../lib/sim/character';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { dedupeByName } from '../../../lib/sim/dedupe';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';
  import Row from '../../ui/Row.svelte';
  import ToolSection from './ToolSection.svelte';

  let {
    character,
    sets,
    onadd,
    onremove,
    plate = false,
  }: {
    character: SimCharacter;
    sets: readonly GearSet[];
    onadd: (set: GearSet) => void;
    onremove: (name: string) => void;
    /** The block sits on the page's one Plate. */
    plate?: boolean;
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

  /** The export's own named sets (part A's FS1 v2 decoder), offered as one-click adds.
   *  Deduplicated on its own: nothing stops an addon export from naming two sets alike,
   *  and two rows sharing one `sim-set-<name>` test id is a Playwright strict-mode failure
   *  waiting to happen the moment a test looks for it, not just a confusing screen. */
  const exported = $derived<GearSet[]>(dedupeByName(new Set<string>(), character.sets.map(toGearSet)));

  /**
   * Every row to show: whatever `sets` actually holds -- the truth of what `BulkSpec.sets`
   * will carry -- claims a name first, then any export-offered set not already added. `sets`
   * has to win a name collision, not `exported`: a player who pastes an edited set under a
   * name the export already offered means the edit, and `exported`'s own row for that name
   * is nothing but a now-stale suggestion -- keeping both under the identical test id would
   * be the same strict-mode failure `exported`'s own dedup above guards against, and keeping
   * the wrong one would show a suggestion the outgoing request has already moved past.
   */
  const rows = $derived<GearSet[]>(dedupeByName(new Set<string>(), [...sets, ...exported]));

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

<ToolSection title={bulkCopy.setsTitle} testid="sim-named-sets" {plate}>
  <p class="text-muted text-[13px]">{bulkCopy.setsIntro}</p>

  {#if rows.length > 0}
    <div class="pt-1">
      {#each rows as set (set.name)}
        <Row title={set.name} open density="slim" testid={`sim-set-${set.name}`}>
          {#snippet trail()}
            <Button size="sm" onclick={() => (isAdded(set.name) ? onremove(set.name) : onadd(set))}
              >{isAdded(set.name) ? bulkCopy.setsRemove : bulkCopy.setsAdd}</Button
            >
          {/snippet}
        </Row>
      {/each}
    </div>
  {/if}

  <div class="flex flex-wrap items-end gap-3 pt-3">
    <div class="w-40">
      <Field id="sim-set-name" label="Name" placeholder="Name" testid="sim-set-name" bind:value={name} />
    </div>
    <Field
      id="sim-set-input"
      label={bulkCopy.setsPaste}
      placeholder={bulkCopy.setsPaste}
      grow
      invalid={error !== ''}
      testid="sim-set-input"
      bind:value={code}
    />
    <Button size="md" testid="sim-set-add" onclick={add}>{bulkCopy.setsAdd}</Button>
  </div>
  {#if error !== ''}
    <p class="text-wipe pt-2 text-[13px]" role="alert" data-testid="sim-set-error">{error}</p>
  {/if}
</ToolSection>
