<!-- web/src/components/sim/tools/SourcePicker.svelte -->
<!-- Design 6.1. A raid card is one checkbox for every boss; each boss is its own. A source
     gated on a phase that has not opened is hidden until "show unreleased content", and
     then carries the date it opens rather than pretending it is live.
     There are no zone icons to draw: src/content/zones/*.md carries a title, a continent
     and a level range, and nothing image-shaped (contract 10.7) -- this picker labels by
     kind and name and draws none. -->
<script lang="ts">
  import { bulkCopy } from '../../../lib/sim/copy';
  import {
    groupSources,
    isOpen,
    itemsOfSource,
    professionSplit,
    SOURCE_KIND_LABELS,
    type LootFile,
    type LootSource,
  } from '../../../lib/sim/loot';
  import { PHASE_LATER, openDateLabel, phaseLabel, type PhaseRow } from '../../../lib/sim/phase';

  let {
    loot,
    phases,
    shownKinds,
    showUpcoming,
    picked,
    professions,
    now,
    ontogglekind,
    ontoggleupcoming,
    ontoggle,
  }: {
    loot: LootFile;
    phases: readonly PhaseRow[];
    shownKinds: readonly string[];
    showUpcoming: boolean;
    picked: readonly string[];
    professions: readonly string[];
    now: Date;
    ontogglekind: (kind: string) => void;
    ontoggleupcoming: (value: boolean) => void;
    ontoggle: (sourceId: string, bossId?: string) => void;
  } = $props();

  const groups = $derived(groupSources(loot));

  function visible(source: LootSource): boolean {
    return shownKinds.includes(source.kind) && (showUpcoming || isOpen(phases, source, now));
  }

  function isPicked(sourceId: string, bossId = ''): boolean {
    return picked.includes(`${sourceId}|${bossId}`);
  }

  /**
   * "First raids, 9 December 2026" for a dated phase, and the no-date sentence for the
   * literal `"later"` -- contract 10.4's "shows as unreleased without a date".
   */
  function gateLabel(source: LootSource): string {
    if (source.opens === undefined || isOpen(phases, source, now)) return '';
    if (source.opens === PHASE_LATER) return bulkCopy.opensLater;
    const date = openDateLabel(phases, source.opens);
    return date === '' ? bulkCopy.notOpenYet : bulkCopy.opensOn(phaseLabel(source.opens), date);
  }
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-3 border p-3 md:mx-0"
  data-testid="sim-source-picker"
>
  <div class="flex flex-wrap items-center gap-3">
    {#each groups as group (group.kind)}
      <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
        <input
          type="checkbox"
          class="h-5 w-5"
          data-testid={`sim-kind-${group.kind}`}
          checked={shownKinds.includes(group.kind)}
          onchange={() => ontogglekind(group.kind)}
        />
        {group.label}
      </label>
    {/each}
    <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid="sim-upcoming"
        checked={showUpcoming}
        onchange={(event) => ontoggleupcoming(event.currentTarget.checked)}
      />
      {bulkCopy.showUpcoming}
    </label>
  </div>

  {#each groups as group (group.kind)}
    {@const shown = group.sources.filter(visible)}
    {#if shown.length > 0}
      <div class="flex flex-col gap-1">
        <h3 class="section-title text-[13px]">{SOURCE_KIND_LABELS[group.kind]}</h3>
        {#if group.kind === 'crafted'}
          {@const split = professionSplit(shown, professions)}
          {#if split.mine.length > 0}
            <p class="text-muted text-[12px]">{bulkCopy.sourcesMyProfessions}</p>
          {:else}
            <p class="text-muted text-[12px]">{bulkCopy.sourcesProfessionsUnknown}</p>
          {/if}
        {/if}
        {#if group.kind === 'quest'}
          <p class="text-muted text-[12px]">{bulkCopy.sourcesQuestNote}</p>
        {/if}
        {#each shown as source (source.id)}
          <div class="border-line-soft flex flex-col gap-1 border-b py-2 last:border-b-0">
            <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
              <input
                type="checkbox"
                class="h-5 w-5"
                data-testid={`sim-source-${source.id}`}
                checked={isPicked(source.id)}
                onchange={() => ontoggle(source.id)}
              />
              <span class="flex-1">{source.name}</span>
              {#if (source.bosses ?? []).length > 0}
                <span class="text-muted text-[12px]">{bulkCopy.sourcesWholeRaid}</span>
              {/if}
              <span class="tabular text-muted font-mono text-[12px]">
                {itemsOfSource(source).length}
              </span>
            </label>
            {#if gateLabel(source) !== ''}
              <p class="text-muted pl-7 text-[12px]">{gateLabel(source)}</p>
            {/if}
            {#each source.bosses ?? [] as boss (boss.id)}
              <label class="text-text flex min-h-11 items-center gap-2 pl-7 text-[13px]">
                <input
                  type="checkbox"
                  class="h-5 w-5"
                  data-testid={`sim-source-${boss.id}`}
                  checked={isPicked(source.id, boss.id)}
                  onchange={() => ontoggle(source.id, boss.id)}
                />
                <!-- Contract 10.4: "a boss with no name in either database is emitted with
                     an empty name". The id is the honest stand-in; inventing one is not. -->
                <span class="flex-1">{boss.name === '' ? boss.id : boss.name}</span>
                <span class="tabular text-muted font-mono text-[12px]">{boss.items.length}</span>
              </label>
            {/each}
            {#if (source.trash ?? []).length > 0}
              <p class="text-muted pl-7 text-[12px]">{bulkCopy.sourcesTrash}</p>
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  {/each}
</section>
