<!-- web/src/components/sim/SimGearCard.svelte -->
<!-- This run's gear (spec 2026-11-04-sim-results 3.8): an open section holding the shared paperdoll
     in `worn` mode in one Plate (web layout spec 2026-11-05 §4.2), each slot's client enchant on
     line 2 in green (else its item level), then the stats the engine computed, the set lines and
     a link to Best in slot. It shows inputs only: a gain or a pick is /bis's job. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { loadItems, loadSets } from '../../lib/planner/load';
  import type { Item, ItemSet } from '../../lib/planner/types';
  import type { ActionNames } from '../../lib/sim/action-names';
  import {
    enchantNames,
    gearMeta,
    gearSlots,
    classSpecPath,
    setLines,
    statsBlock,
  } from '../../lib/sim/gear-card';
  import { loadEnchants } from '../../lib/sim/enchants';
  import { gearCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import type { SimResult } from '../../lib/sim/types';
  import { gearCopy } from '../../lib/gear/copy';
  import GearDoll from '../character/GearDoll.svelte';
  import ActionLink from '../ui/ActionLink.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';

  const props: { result: SimResult; names: ActionNames | null } = $props();

  const build = activeBuild.build;
  const request = $derived(props.result.request);
  const character = $derived(request.character);
  const hasGear = $derived(character.gear.length > 0);

  let items = $state.raw<ReadonlyMap<number, Item>>(new Map());
  let itemsState = $state<'loading' | 'ready' | 'failed'>('loading');
  let enchants = $state.raw<ReadonlyMap<number, string>>(new Map());
  let sets = $state.raw<readonly ItemSet[]>([]);

  $effect(() => {
    const classSlug = character.class;
    loadItems(build, classSlug).then(
      (file) => {
        items = new Map(file.items.map((item) => [item.id, item]));
        itemsState = 'ready';
      },
      (error: unknown) => {
        console.error('sim gear: the item file did not load', error);
        itemsState = 'failed';
      },
    );
    loadEnchants(build).then(
      (rows) => (enchants = enchantNames(rows)),
      (error: unknown) => console.error('sim gear: the enchant table did not load', error),
    );
    loadSets(build).then(
      (rows) => (sets = rows),
      (error: unknown) => console.error('sim gear: the set table did not load', error),
    );
  });

  const slots = $derived(
    gearSlots({
      gear: character.gear,
      items,
      failed: itemsState === 'failed',
      enchants,
    }),
  );
  const stats = $derived(statsBlock(props.result));
  const lines = $derived(setLines(character.gear, sets));
</script>

<div class="mx-[18px] md:mx-0">
  <Section title={copy.title} meta={hasGear ? gearMeta(character.gear) : undefined} testid="sim-gear-card">
    {#if !hasGear}
      <p class="text-muted text-[14px]" data-testid="sim-gear-empty">
        {copy.noGear}
      </p>
    {:else}
      <div class="flex flex-col gap-4">
        <Plate label={gearCopy.listLabel} testid="sim-gear-paperdoll">
          <GearDoll
            {slots}
            mode="worn"
            {build}
            classSlug={character.class}
            loading={itemsState === 'loading'}
            failed={itemsState === 'failed'}
            label={gearCopy.listLabel}
            testid="sim-gear-doll"
          />
        </Plate>
        {#if stats !== null}
          <section
            class="border-line-soft rounded-panel measure-wide border bg-[#0b0e14] px-4 py-3"
            data-testid="sim-final-stats"
          >
            <h3 class="quiet-label mb-3">{stats.head}</h3>
            <dl class="grid grid-cols-2 gap-x-4 gap-y-3 md:grid-cols-[1fr_2fr_1fr_1fr_1fr]">
              {#if stats.hit !== null}
                <div class="col-span-2 md:order-2 md:col-span-1" data-testid="sim-stat-hit">
                  <dt class="text-muted text-[13px]">{stats.hit.label}</dt>
                  <dd class="text-strong numeric m-0 text-[15px] font-semibold">{stats.hit.value}</dd>
                  {#if stats.hit.reference !== null}
                    <dd class="text-muted m-0 mt-1 text-[13px]" data-testid="sim-stat-hit-measured">
                      {stats.hit.reference.note}
                    </dd>
                  {/if}
                </div>
              {/if}
              {#each stats.cells as cell, at (cell.id)}
                <div class={at === 0 ? 'md:order-1' : 'md:order-3'} data-testid={`sim-stat-${cell.id}`}>
                  <dt class="text-muted text-[13px]">{cell.label}</dt>
                  <dd class="text-strong numeric m-0 text-[15px] font-semibold">{cell.value}</dd>
                </div>
              {/each}
            </dl>
          </section>
        {/if}
      </div>
      <div class="flex flex-wrap items-center justify-between gap-x-4 gap-y-1 py-1.5">
        <div class="text-muted min-w-0 py-2 text-[13px]" data-testid="sim-gear-sets">
          {#if lines.length === 0}
            <p>{copy.noSet}</p>
          {:else}
            {#each lines as line (line.id)}
              <p data-testid="sim-gear-set">
                {line.head}{#each line.bonuses as bonus (bonus.text)}
                  · <span class={bonus.active ? 'text-rarity-uncommon' : ''}>{bonus.text}</span>
                {/each}
              </p>
            {/each}
          {/if}
        </div>
        <ActionLink href={`/bis/${classSpecPath(request.spec)}`}>{copy.bisLink}</ActionLink>
      </div>
    {/if}
  </Section>
</div>
