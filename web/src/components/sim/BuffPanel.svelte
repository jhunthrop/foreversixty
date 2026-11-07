<!-- web/src/components/sim/BuffPanel.svelte -->
<!-- Design 4.3, behind "Custom". Twelve sections, each the site's own <details>, each
     closed on arrival with the count of what is on in its summary: the whole vocabulary is
     several hundred rows and a flat list of them is not a control, it is a wall.

     Nothing here knows an id. The rows are buffs.ts's catalogue, which is IDS.md's own
     list read out of src/data/generated/sim-ids.json; the names and icons are the build's,
     with buff-names.ts's humanised fallback when the build has no table. A row this page
     could not file would still be reachable in the request drawer (Task 15), which is the
     escape hatch the design names.

     It sits inside the More settings Sheet (sim spec 2026-11-06): a heading, then one
     Disclosure per group and plain rows inside it, no boxes. -->
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
  import { imbueSections, withImbue } from '../../lib/sim/imbue-rows';
  import ImbueRows from './ImbueRows.svelte';
  import { simCopy } from '../../lib/sim/copy';
  import type { SimSettings } from '../../lib/sim/settings';
  import CooldownRows from './CooldownRows.svelte';
  import Disclosure from '../ui/Disclosure.svelte';
  import Select from '../ui/Select.svelte';

  let {
    settings,
    build,
    classSlug,
    names,
    disabled,
    onchange,
  }: {
    settings: SimSettings;
    /** The character's data build, for the icon paths. */
    build: string;
    /** The character's class: the weapon section offers only what it can use. */
    classSlug: string;
    names: BuffNames | null;
    disabled: boolean;
    onchange: (next: SimSettings) => void;
  } = $props();

  const GRADES: BuffGrade[] = ['off', 'on', 'improved'];

  const selection = $derived({ buffs: settings.buffs, consumables: settings.consumables });

  function set(row: { id: string; kind: 'buff' | 'consumable' }, grade: BuffGrade): void {
    onchange({ ...settings, ...withGrade(selection, row, grade) });
  }

  const imbues = $derived(imbueSections(classSlug));
  const imbueIds = $derived(
    imbues.flatMap((section) => section.kinds.flatMap((kind) => kind.rows.map((row) => row.id))),
  );

  function countIn(group: BuffGroupId): number {
    if (group === 'weapon-imbue') return imbueIds.filter((id) => settings.consumables.includes(id)).length;
    return selectedIn(selection, group).length;
  }

  function sizeOf(group: BuffGroupId): number {
    return group === 'weapon-imbue' ? imbueIds.length : rowsIn(group).length;
  }

  function pickImbue(id: string, on: boolean): void {
    onchange({ ...settings, ...withImbue(selection, id, on) });
  }
</script>

<div class="flex flex-col" data-testid="sim-buff-panel">
  <h3 class="text-label tracking-label text-muted font-bold uppercase">{simCopy.buffPanel}</h3>
  <p class="text-small text-muted mt-1 mb-2">{simCopy.buffPanelNote}</p>

  {#each BUFF_GROUPS as group (group)}
    {@const rows = rowsIn(group)}
    {#if rows.length > 0}
      <Disclosure
        summary={simCopy.buffGroupLabel[group] ?? group}
        meta={`${countIn(group)}/${sizeOf(group)}`}
        testid={`sim-buff-group-${group}`}
      >
        {#if group === 'weapon-imbue'}
          <ImbueRows
            sections={imbues}
            selected={settings.consumables}
            {build}
            {names}
            {disabled}
            onpick={pickImbue}
          />
        {:else}
          <ul class="grid grid-cols-1 gap-x-4 gap-y-1 md:grid-cols-2">
            {#each rows as row (row.id)}
              {@const label = buffLabel(row.id, names)}
              {@const icon = buffIcon(build, row.id, names)}
              {@const grade = gradeOf(listFor(selection, row.kind), row.id)}
              <li class="flex min-h-11 items-center gap-2 text-[14px] md:min-h-9">
                {#if icon !== null}
                  <!-- The name beside it carries the row, so the icon is decorative. -->
                  <img src={icon} alt="" width="20" height="20" class="rounded-[2px]" loading="lazy" />
                {/if}
                {#if row.graded}
                  <div class="flex min-w-0 flex-1 items-center gap-2">
                    <span class="truncate">{label}</span>
                    <Select
                      id={`sim-buff-${row.id}-field`}
                      class="ml-auto w-[130px]"
                      size="dense"
                      align="end"
                      {disabled}
                      label={simCopy.gradeFor(label)}
                      value={grade}
                      options={GRADES.map((option) => ({ value: option, label: simCopy.gradeLabel[option] }))}
                      testid={`sim-buff-${row.id}`}
                      onchange={(event) =>
                        set(row, (event.currentTarget as HTMLSelectElement).value as BuffGrade)}
                    />
                  </div>
                {:else}
                  <label class="flex min-w-0 flex-1 items-center gap-2">
                    <input
                      type="checkbox"
                      class="h-5 w-5 shrink-0"
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
        {/if}
      </Disclosure>
    {/if}
  {/each}

  <CooldownRows
    {settings}
    {names}
    {disabled}
    ids={[...selectedIn(selection, 'potion'), ...selectedIn(selection, 'explosive')]}
    {onchange}
  />
</div>
