<!-- web/src/components/guild/report/PullDeaths.svelte -->
<!-- One pull's deaths (web layout spec 2026-11-05 §7.2 "Raid report"), the body of the pull's
     Inspector: per death, who (class colour) and how far into the pull, the killing blow (the
     spell's icon where the build has one, the ability, its source, the amount and the overkill),
     then the hits before it, oldest first. A pull the log has not been read for says so in one
     muted line. A member's report carries only their own deaths, so nothing here filters. -->
<script lang="ts">
  import CharacterName from '../../ui/CharacterName.svelte';
  import type { RaidReportPull } from '../../../lib/guild/attendance-contract';
  import { deathModels, type HitModel } from '../../../lib/guild/report-night-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import SpellIcon from '../../ui/SpellIcon.svelte';

  let {
    pull,
    iconOf,
  }: {
    pull: RaidReportPull;
    /** The spell's icon url in the active build, or null. */
    iconOf: (spellId: number | null) => string | null;
  } = $props();

  const deaths = $derived(pull.death_causes === null ? null : deathModels(pull.death_causes));
</script>

{#snippet hit(model: HitModel, strong: boolean)}
  <span class="flex min-w-0 items-center gap-2">
    <SpellIcon url={iconOf(model.spellId)} size={20} tip={model.spellId ?? undefined} />
    <span class="min-w-0">
      <span class={strong ? 'text-strong font-bold' : 'text-text'}>{model.ability}</span><span
        class="text-muted">{` from ${model.source}`}</span
      ><span class="text-text tabular-nums">, {model.amount}</span>
    </span>
  </span>
{/snippet}

{#if deaths === null}
  <p class="text-muted" data-testid="raid-pull-not-read">{copy.pulls.notRead}</p>
{:else}
  <ol class="flex flex-col gap-4" data-testid="raid-pull-deaths">
    {#each deaths as death (death.key)}
      <li class="flex min-w-0 flex-col gap-1" data-testid="raid-pull-death">
        <p>
          <CharacterName name={death.name} cls={death.class} weight="inherit" class="font-bold" />
          <span class="text-muted tabular-nums">{death.at}</span>
        </p>
        {#if death.killingBlow === null}
          <p class="text-muted" data-testid="raid-pull-no-hit">{copy.pulls.noHit}</p>
        {:else}
          <div data-testid="raid-pull-killing-blow">
            <span class="sr-only">{copy.pulls.killingBlow}: </span>{@render hit(death.killingBlow, true)}
          </div>
        {/if}
        {#if death.before.length > 0}
          <ol class="flex flex-col gap-1 pl-7 text-[13px]" aria-label={copy.pulls.lastHits}>
            {#each death.before as model, index (index)}
              <li class="flex items-baseline gap-2" data-testid="raid-pull-last-hit">
                <span class="text-muted w-9 shrink-0 tabular-nums">{model.clock}</span>
                {@render hit(model, false)}
              </li>
            {/each}
          </ol>
        {/if}
      </li>
    {/each}
  </ol>
{/if}
