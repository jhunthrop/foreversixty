<!-- web/src/components/sim/ImbueRows.svelte -->
<!-- The custom panel's weapon section (round 10 wow-player review): one heading per hand the
     character's class can enchant, the rogue's beside-the-poison slot, and inside each the
     rows by kind. Each row is the item's own full name; ticking one replaces whatever its
     slot held, because a slot is one enchant (imbue-rows.ts). -->
<script lang="ts">
  import { buffIcon, imbueLabel, type BuffNames } from '../../lib/sim/buff-names';
  import { simCopy } from '../../lib/sim/copy';
  import type { ImbueSection } from '../../lib/sim/imbue-rows';

  let {
    sections,
    selected,
    build,
    names,
    disabled,
    onpick,
  }: {
    sections: ImbueSection[];
    /** The consumables list the settings carry. */
    selected: readonly string[];
    build: string;
    names: BuffNames | null;
    disabled: boolean;
    onpick: (id: string, on: boolean) => void;
  } = $props();
</script>

<div class="flex flex-col gap-3 p-3 pt-0">
  {#each sections as section (section.slot)}
    <section data-testid={`sim-imbue-slot-${section.slot}`}>
      <h3 class="text-label text-nav mb-1 font-bold">
        {simCopy.imbueSlotLabel[section.slot] ?? section.slot}
      </h3>
      {#each section.kinds as group (group.kind)}
        <p class="text-muted mt-1 text-[12px]">{simCopy.imbueKindLabel[group.kind] ?? group.kind}</p>
        <ul class="grid grid-cols-1 gap-1 md:grid-cols-2">
          {#each group.rows as row (row.id)}
            {@const label = imbueLabel(row.id, names)}
            {@const icon = buffIcon(build, row.id, names)}
            <li class="flex min-h-11 items-center gap-2 text-[13px] md:min-h-9">
              {#if icon !== null}
                <!-- The name beside it carries the row, so the icon is decorative. -->
                <img src={icon} alt="" width="20" height="20" class="rounded-[2px]" loading="lazy" />
              {/if}
              <label class="flex min-w-0 flex-1 items-center gap-2">
                <input
                  type="checkbox"
                  class="h-5 w-5 shrink-0"
                  {disabled}
                  checked={selected.includes(row.id)}
                  onchange={(event) => onpick(row.id, event.currentTarget.checked)}
                  data-testid={`sim-buff-${row.id}`}
                />
                <span>{label}</span>
              </label>
            </li>
          {/each}
        </ul>
      {/each}
    </section>
  {/each}
</div>
