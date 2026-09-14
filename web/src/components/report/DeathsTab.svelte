<!-- web/src/components/report/DeathsTab.svelte -->
<!-- The deaths recap. One card per death, in the order they happened: the killing blow,
     the last ten damage events with the health they left, the auras that were up and the
     ones that had just fallen off, a button that sets the window to the twenty seconds
     before it, and the link into the planner.
     Cards rather than a table at every width: a death is read top to bottom, and the
     last-ten list is the whole point of the view. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import { plannerLinkFor } from '../../lib/report/planner-link';
  import type { CombatantRow, Death } from '../../lib/report/types';
  import { deathWindow, type TimeWindow } from '../../lib/report/window';

  let {
    deaths,
    durationMs,
    combatants,
    classOf,
    dataBuild,
    treeSizesFor,
    onWindow,
  }: {
    deaths: Death[];
    durationMs: number;
    combatants: CombatantRow[];
    classOf: Map<string, string>;
    dataBuild: string;
    /** Talents per tree for a class, or [] when the planner has no data for it. */
    treeSizesFor: (className: string | undefined) => number[];
    onWindow: (window: TimeWindow) => void;
  } = $props();

  const ordered = $derived([...deaths].sort((a, b) => a.at_ms - b.at_ms));

  function linkFor(death: Death): ReturnType<typeof plannerLinkFor> {
    const combatant = combatants.find((row) => row.guid === death.guid);
    if (combatant === undefined) return null;
    const className = death.class ?? classOf.get(death.guid);
    return plannerLinkFor({ dataBuild, className, treeSizes: treeSizesFor(className), combatant });
  }
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nobody died in this window.</p>
{:else}
  <ul class="flex flex-col gap-4" data-testid="deaths-tab">
    {#each ordered as death (`${death.guid}-${death.at_ms}`)}
      {@const link = linkFor(death)}
      <li
        class="border-line rounded-panel bg-raised flex flex-col gap-3 border p-3"
        data-testid={`death-${death.guid}`}
      >
        <div class="flex flex-wrap items-baseline gap-x-3 gap-y-1">
          <span
            class="text-[15px] font-semibold"
            style={`color: ${classColorVar(death.class ?? classOf.get(death.guid))}`}
          >
            {splitUnitName(death.name).name}
          </span>
          <span class="text-muted font-mono tabular text-[13px]">{formatDuration(death.at_ms)}</span>
          {#if death.killing_blow}
            <span class="text-[13px]">
              killed by {splitUnitName(death.killing_blow.source_name).name} ·
              {death.killing_blow.spell_name === '' ? 'Melee' : death.killing_blow.spell_name} ·
              {formatAmount(death.killing_blow.amount)}
              {#if death.killing_blow.overkill}
                <span class="text-muted"> ({formatAmount(death.killing_blow.overkill)} overkill)</span>
              {/if}
            </span>
          {/if}
          {#if death.release_ms}
            <span class="text-muted text-[13px]">
              released after {formatDuration(death.release_ms - death.at_ms)}
            </span>
          {/if}
        </div>

        <table class="w-full text-[13px]">
          <caption class="label text-muted text-left">Last hits</caption>
          <tbody>
            {#each death.last as hit (`${hit.at_ms}-${hit.spell_id}`)}
              <tr class="border-line-soft border-b">
                <td class="text-muted font-mono tabular py-1 pr-3">{formatDuration(hit.at_ms)}</td>
                <td class="py-1 pr-3">{hit.spell_name === '' ? 'Melee' : hit.spell_name}</td>
                <td class="text-muted truncate py-1 pr-3">{splitUnitName(hit.source_name).name}</td>
                <td class="font-mono tabular py-1 pr-3 text-right">{formatAmount(hit.amount)}</td>
                <td class="w-[30%] py-1">
                  {#if hit.max_hp}
                    <span
                      class="bg-line-soft block h-[6px] w-full"
                      title={`${hit.hp_after ?? 0} of ${hit.max_hp}`}
                    >
                      <span
                        class="bg-gold block h-full"
                        style={`width: ${Math.max(0, Math.min(100, ((hit.hp_after ?? 0) / hit.max_hp) * 100))}%`}
                      ></span>
                    </span>
                  {/if}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>

        {#if death.auras_held.length > 0 || death.auras_lost.length > 0}
          <p class="text-[13px]">
            {#if death.auras_held.length > 0}
              <span class="label text-muted">Up</span>
              {death.auras_held.map((aura) => aura.name).join(', ')}
            {/if}
            {#if death.auras_lost.length > 0}
              <span class="label text-muted ml-3">Just lost</span>
              {death.auras_lost.map((aura) => aura.name).join(', ')}
            {/if}
          </p>
        {/if}

        <div class="flex flex-wrap gap-2">
          <button
            type="button"
            class="border-line-warm rounded-control text-text inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
            data-testid="death-window"
            onclick={() => onWindow(deathWindow(death.at_ms, durationMs))}
          >
            The 20s before this
          </button>
          {#if link}
            <a
              class="border-line-warm-strong rounded-control text-strong inline-flex h-11 items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:h-9"
              href={link.href}
              data-testid="death-build-link"
            >
              {link.label}
            </a>
          {/if}
        </div>
      </li>
    {/each}
  </ul>
{/if}
