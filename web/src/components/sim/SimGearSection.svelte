<!-- web/src/components/sim/SimGearSection.svelte -->
<!-- The character's gear (sim spec 2026-11-06 §4.3, §4.5; spec 2026-11-04-sim-results 3.8): the
     shared paperdoll frame in `worn` mode, as BiS, the character page and the reports draw it
     (owner ruling 2026-11-11: one gear look), its title bar the spec and the level; under the
     frame's 423 px minimum (every phone) it is the frame's own list form, each slot's client
     enchant after its slot in green (else its item level). The saved-runs strip (CharacterStrip)
     keeps the list form: a compact strip, not a gear panel.
     After a run it is open (no Plate, ruling S3) with the five stats the engine computed, the
     run's own measured miss rate, the set line and the Best in slot link. On the landing and while
     running (`plated`) the doll is the page's one Plate and the section's action is the link. It
     shows inputs only: a gain or a pick is /bis's job. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import activeBuild from '../../data/active-build.json';
  import { gearCopy } from '../../lib/gear/copy';
  import { loadItems, loadSets } from '../../lib/planner/load';
  import type { Item, ItemSet } from '../../lib/planner/types';
  import { loadEnchants } from '../../lib/sim/enchants';
  import { classSpecPath, enchantNames, gearSlots, setLines, statsBlock } from '../../lib/sim/gear-card';
  import { gearCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { SimRequest, SimResult } from '../../lib/sim/types';
  import { LINK } from '../../lib/ui/primitives';
  import GearDoll from '../character/GearDoll.svelte';
  import Figure from '../ui/Figure.svelte';
  import Section from '../ui/Section.svelte';

  let {
    request,
    result = null,
    plated = false,
  }: {
    request: SimRequest;
    /** The run: its stats and measured misses. Absent before a run. */
    result?: SimResult | null;
    /** The landing's doll: the page's one Plate, the BiS link as the section's action. */
    plated?: boolean;
  } = $props();

  const build = activeBuild.build;
  const character = $derived(request.character);
  /** The frame's width cap, the character page's (CharacterGear): the stats and set lines run under it. */
  const DOLL_WIDTH = 'w-full max-w-[640px]';
  const hasGear = $derived(character.gear.length > 0);
  const bisHref = $derived(`/bis/${classSpecPath(request.spec)}`);

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
    gearSlots({ gear: character.gear, items, failed: itemsState === 'failed', enchants }),
  );
  const stats = $derived(result === null ? null : statsBlock(result));
  /** Attack power, Hit, then the rest in the engine's order (spec 4.5). */
  const figures = $derived.by(() => {
    if (stats === null) return [];
    const [first, ...rest] = stats.cells;
    const lead = first?.id === 'attack-power' ? [first] : [];
    const tail = first?.id === 'attack-power' ? rest : stats.cells;
    return [...lead, ...(stats.hit === null ? [] : [stats.hit]), ...tail];
  });
  /** The run's measured misses only: the boss and the fight are the page's own. */
  const measured = $derived(stats?.hit?.reference?.note ?? null);
  const lines = $derived(setLines(character.gear, sets));
  /** The frame's title bar, as the character page's: the spec and class, then the level. */
  const dollTitle = $derived(specLabel(request.spec));
  const dollMeta = $derived(copy.dollLevel(character.level));
</script>

{#snippet doll()}
  <GearDoll
    {slots}
    mode="worn"
    {build}
    classSlug={character.class}
    loading={itemsState === 'loading'}
    failed={itemsState === 'failed'}
    title={dollTitle}
    meta={dollMeta}
    form="auto"
    head={false}
    label={gearCopy.listLabel}
    testid="sim-gear-doll"
  />
{/snippet}

<Section
  panel
  title={copy.title}
  action={plated && hasGear ? { label: copy.bisAction, href: bisHref } : undefined}
  testid="sim-gear-card"
>
  {#if !hasGear}
    <EmptyState message={copy.noGear} testid="sim-gear-empty" />
  {:else if plated}
    <div class={DOLL_WIDTH} data-testid="sim-gear-paperdoll">{@render doll()}</div>
  {:else}
    <div class={DOLL_WIDTH} data-testid="sim-gear-paperdoll">{@render doll()}</div>
    {#if result !== null && stats === null}
      <p class="text-muted pt-4 text-[13px] leading-5" data-testid="sim-stats-older-run">
        {copy.statsNeedNewerRun}
      </p>
    {:else if stats !== null}
      <div class="flex flex-col gap-2 pt-4" data-testid="sim-final-stats">
        <div class="grid grid-cols-2 gap-x-6 gap-y-5 sm:grid-cols-3 lg:grid-cols-5">
          {#each figures as figure (figure.id)}
            <Figure label={figure.label} value={figure.value} size="sm" testid={`sim-stat-${figure.id}`} />
          {/each}
        </div>
        {#if measured}<p class="text-muted text-[13px] leading-5" data-testid="sim-stat-hit-measured">
            {measured}
          </p>{/if}
      </div>
    {/if}
    <div class="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-2 pt-4">
      {#if lines.length > 0}
        <span class="text-muted min-w-0 text-[13px] leading-5" data-testid="sim-gear-sets">
          {#each lines as line (line.id)}
            <span class="block" data-testid="sim-gear-set"
              >{line.head}{#each line.bonuses as bonus (bonus.text)}
                · <span class={bonus.active ? 'text-rarity-uncommon' : ''}>{bonus.text}</span>
              {/each}</span
            >
          {/each}
        </span>
      {/if}
      <a class="{LINK} ml-auto text-[14px]" href={bisHref}>{copy.bisLink} ›</a>
    </div>
  {/if}
</Section>
