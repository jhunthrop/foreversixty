<!-- web/src/components/sim/BuffPanel.svelte -->
<!-- Design 4.3, behind "Custom". Twelve sections, each the site's own <details>, each
     closed on arrival with the count of what is on in its summary: the whole vocabulary is
     several hundred rows and a flat list of them is not a control, it is a wall.

     Nothing here knows an id. The rows are buffs.ts's catalogue, which is IDS.md's own
     list read out of src/data/generated/sim-ids.json; the names and icons are the build's,
     with buff-names.ts's humanised fallback when the build has no table. A row this page
     could not file would still be reachable in the request drawer (Task 15), which is the
     escape hatch the design names. -->
<script lang="ts">
  import {
    BUFF_GROUPS,
    gradeOf,
    listFor,
    rowsIn,
    selectedIn,
    withGrade,
    type BuffGrade,
    type BuffGroupId,
  } from '../../lib/sim/buffs';
  import { buffIcon, buffLabel, type BuffNames } from '../../lib/sim/buff-names';
  import { simCopy } from '../../lib/sim/copy';
  import type { SimSettings } from '../../lib/sim/settings';

  let {
    settings,
    build,
    names,
    disabled,
    onchange,
  }: {
    settings: SimSettings;
    /** The character's data build, for the icon paths. */
    build: string;
    names: BuffNames | null;
    disabled: boolean;
    onchange: (next: SimSettings) => void;
  } = $props();

  const GRADES: BuffGrade[] = ['off', 'on', 'improved'];

  const selection = $derived({ buffs: settings.buffs, consumables: settings.consumables });

  function set(row: { id: string; kind: 'buff' | 'consumable' }, grade: BuffGrade): void {
    onchange({ ...settings, ...withGrade(selection, row, grade) });
  }

  function countIn(group: BuffGroupId): number {
    return selectedIn(selection, group).length;
  }
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-buff-panel"
>
  <h2 class="section-title text-[15px]">{simCopy.buffPanel}</h2>
  <p class="text-muted text-[12px]">{simCopy.buffPanelNote}</p>

  {#each BUFF_GROUPS as group (group)}
    {@const rows = rowsIn(group)}
    {#if rows.length > 0}
      <details class="border-line-soft rounded-panel border" data-testid={`sim-buff-group-${group}`}>
        <summary class="label text-nav flex min-h-11 cursor-pointer items-center gap-2 px-3 md:min-h-9">
          <span>{simCopy.buffGroupLabel[group] ?? group}</span>
          <span class="tabular text-muted font-mono text-[12px]">{countIn(group)}/{rows.length}</span>
        </summary>
        <ul class="grid grid-cols-1 gap-1 p-3 pt-0 md:grid-cols-2">
          {#each rows as row (row.id)}
            {@const label = buffLabel(row.id, names)}
            {@const icon = buffIcon(build, row.id, names)}
            {@const grade = gradeOf(listFor(selection, row.kind), row.id)}
            <li class="flex min-h-11 items-center gap-2 text-[13px] md:min-h-9">
              {#if icon !== null}
                <!-- The name beside it carries the row, so the icon is decorative. -->
                <img src={icon} alt="" width="20" height="20" class="rounded-[2px]" loading="lazy" />
              {/if}
              {#if row.graded}
                <label class="flex min-w-0 flex-1 items-center gap-2">
                  <span class="truncate">{label}</span>
                  <select
                    class="border-line-warm rounded-control bg-raised text-text ml-auto min-h-11 border px-2 text-[12px] md:min-h-9"
                    {disabled}
                    aria-label={simCopy.gradeFor(label)}
                    value={grade}
                    onchange={(event) => set(row, event.currentTarget.value as BuffGrade)}
                    data-testid={`sim-buff-${row.id}`}
                  >
                    {#each GRADES as option (option)}
                      <option value={option}>{simCopy.gradeLabel[option]}</option>
                    {/each}
                  </select>
                </label>
              {:else}
                <label class="flex min-w-0 flex-1 items-center gap-2">
                  <input
                    type="checkbox"
                    class="accent-gold h-5 w-5 shrink-0"
                    {disabled}
                    checked={grade !== 'off'}
                    onchange={(event) => set(row, event.currentTarget.checked ? 'on' : 'off')}
                    data-testid={`sim-buff-${row.id}`}
                  />
                  <span class="truncate">{label}</span>
                </label>
              {/if}
            </li>
          {/each}
        </ul>
      </details>
    {/if}
  {/each}
</section>
