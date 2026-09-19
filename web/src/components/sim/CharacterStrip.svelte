<!-- web/src/components/sim/CharacterStrip.svelte -->
<!-- Who is being simulated. Read-only by design: gear is changed in the planner, which the
     "Open in planner" link goes to, and a second editor here would be a second set of rules
     to keep in step with lib/planner/rules.ts.
     The slot grid's markup is deliberately a copy of GearPanel.svelte's rather than a reuse
     of the component: that one is a button grid bound to PlannerStore and opens an item
     picker, and none of that belongs on this page. The shared parts -- SLOTS, SLOT_LABELS,
     rarityClassFor, dataUrl -- are imported, so only the layout is duplicated. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import { rarityClassFor } from '../../lib/planner/items';
  import { dataUrl } from '../../lib/planner/load';
  import { SLOTS, SLOT_LABELS, type Item, type RaceRow, type Slot } from '../../lib/planner/types';
  import type { SimCharacter } from '../../lib/sim/character';
  import { SIM_LEVEL, needsRace, plannerHrefFor } from '../../lib/sim/character';
  import { simCopy } from '../../lib/sim/copy';
  import { sourcePill } from '../../lib/sim/sources';
  import { specLabel } from '../../lib/sim/spec-label';

  let {
    character,
    items,
    races = [],
    gearKnown = true,
    onchange,
    onrace = () => {},
  }: {
    character: SimCharacter;
    items: Map<number, Item>;
    /** The build's races, for the one-time picker. Only the fight source ever needs it. */
    races?: RaceRow[];
    /** False on a saved sim, whose stored request has no gear list. */
    gearKnown?: boolean;
    onchange: () => void;
    /** The player answering the race question; the store replaces the character. */
    onrace?: (slug: string) => void;
  } = $props();

  const colour = $derived(classColorVar(character.class_slug));
  // A combat log records no race, so a character from one arrives with PENDING_RACE and
  // the descriptor asks instead of asserting. Rendering an empty string there would be the
  // page quietly claiming a raceless character, which is the thing Task 7 refuses to do.
  const pending = $derived(needsRace(character));
  // A build with fewer points than SIM_LEVEL - BASE_LEVEL implies is a part-levelled
  // character: the engine still sims it at 60 (there is no level control on this page), so
  // the line says how many points it actually spends rather than letting "· 60" imply a
  // full-levelled build that ran the tree dry.
  const levelSuffix = $derived(
    character.talent_level < SIM_LEVEL ? ` · ${character.point_order.length} talent points` : '',
  );
  const descriptor = $derived(
    (pending
      ? `${specLabel(character.spec)} · ${SIM_LEVEL}`
      : `${character.race_slug.replace(/-/g, ' ')} ${specLabel(character.spec)} · ${SIM_LEVEL}`) +
      levelSuffix,
  );
  const split = $derived(character.point_order.length);
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-4 border p-4 md:mx-0"
  data-testid="sim-character"
>
  <div class="flex flex-wrap items-center gap-3">
    <span class="rounded-control h-9 w-9 shrink-0" style={`background: ${colour}`} aria-hidden="true"></span>
    <span class="text-[17px] font-semibold" style={`color: ${colour}`} data-testid="sim-character-name">
      {character.name}
    </span>
    <span class="text-muted text-[13px] capitalize" data-testid="sim-character-descriptor">
      {descriptor}
    </span>
    {#if pending}
      <!-- One question, asked once. The run control stays disabled until it is answered,
           because sim/request.ParseRace refuses an empty race at the boundary and a failed
           run is a worse answer than a question. -->
      <label class="flex items-center gap-2">
        <span class="text-strong text-[13px]" data-testid="sim-race-prompt">{simCopy.pickRace}</span>
        <select
          class="border-line-warm rounded-control bg-raised text-text min-h-11 border px-3 text-[14px] font-semibold md:min-h-9"
          value=""
          onchange={(event) => onrace(event.currentTarget.value)}
          data-testid="sim-race-picker"
        >
          <option value="" disabled>{simCopy.pickRacePlaceholder}</option>
          {#each races as race (race.slug)}
            <option value={race.slug}>{race.name}</option>
          {/each}
        </select>
      </label>
    {/if}
    <span class="pill pill-site" data-testid="sim-source-pill">{sourcePill(character.source)}</span>
    <button
      type="button"
      class="border-line-warm rounded-control text-nav label ml-auto min-h-11 border px-3 md:min-h-9"
      onclick={onchange}
      data-testid="sim-change-source"
    >
      {simCopy.changeSource}
    </button>
  </div>

  {#if !gearKnown}
    <p class="text-muted text-[13px]" data-testid="sim-no-gear">{simCopy.savedNoGear}</p>
  {:else}
    <div class="grid grid-cols-2 gap-2 md:grid-cols-4" data-testid="sim-gear">
      {#each SLOTS as slot (slot)}
        {@const equippedId = character.gear[slot as Slot]}
        {@const item = equippedId === undefined ? undefined : items.get(equippedId)}
        <div
          class="border-line rounded-control bg-card-top flex min-h-11 items-center gap-2 border px-2 py-1"
          data-testid={`sim-slot-${slot}`}
        >
          {#if item}
            <img
              src={dataUrl(character.tree_version, `icons/${item.icon}.webp`)}
              alt=""
              width="28"
              height="28"
              loading="lazy"
              decoding="async"
              class="rounded-control border-line h-7 w-7 border object-cover"
            />
          {:else}
            <span class="rounded-control border-line-soft h-7 w-7 border" aria-hidden="true"></span>
          {/if}
          <span class="flex min-w-0 flex-col">
            <span class="label text-muted">{SLOT_LABELS[slot]}</span>
            <span
              class={`truncate text-[13px] font-semibold ${item ? rarityClassFor(item.quality) : 'text-muted'}`}
            >
              {item ? item.name : 'Empty'}
            </span>
          </span>
        </div>
      {/each}
    </div>
  {/if}

  <div class="flex flex-wrap items-baseline gap-3 text-[13px]">
    <span class="text-muted label">Talents</span>
    <!-- A saved sim's stored request carries the engine's talent *string*, not the
         planner's point order (Task 17), so its `point_order` is always empty and `split`
         is 0 -- not because the build spent nothing, but because this page cannot see the
         order. Showing "0 points" would claim the build spent none, so the count is left
         out entirely rather than printed wrong. -->
    {#if split > 0}
      <span class="tabular text-strong font-mono" data-testid="sim-talent-count">{split} points</span>
    {/if}
    <a class="ml-auto" href={plannerHrefFor(character)} data-testid="sim-open-planner">
      {simCopy.openInPlanner}
    </a>
  </div>
</section>
