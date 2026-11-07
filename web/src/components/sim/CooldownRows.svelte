<!-- web/src/components/sim/CooldownRows.svelte -->
<!-- Design 4.3's last sentence, inside the same panel. One row per scheduled thing: the
     mode, and a seconds field only where the mode needs one -- an always-visible number
     input next to "On cooldown" would invite a time that is never read. A Disclosure of plain
     rows inside the More settings Sheet, no box. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { COOLDOWN_MODES, rowsFor, withCooldown, type CooldownMode } from '../../lib/sim/cooldowns';
  import { buffLabel, type BuffNames } from '../../lib/sim/buff-names';
  import type { SimSettings } from '../../lib/sim/settings';
  import { FIELD_CONTROL, FIELD_FOCUS, FIELD_HEIGHT, fieldEdge } from '../../lib/ui/field';
  import Disclosure from '../ui/Disclosure.svelte';
  import Select from '../ui/Select.svelte';

  let {
    settings,
    ids,
    names,
    disabled,
    onchange,
  }: {
    settings: SimSettings;
    /** The ids a row is offered for: whatever is ticked in Potions and Explosives. */
    ids: string[];
    names: BuffNames | null;
    disabled: boolean;
    onchange: (next: SimSettings) => void;
  } = $props();

  const rows = $derived(rowsFor(ids, settings.cooldowns, settings.encounter));

  function set(id: string, mode: CooldownMode, atSec: number): void {
    onchange({
      ...settings,
      cooldowns: withCooldown(settings.cooldowns, id, mode, atSec, settings.encounter),
    });
  }

  const control = [FIELD_CONTROL, FIELD_HEIGHT, FIELD_FOCUS, fieldEdge(false), 'w-20'].join(' ');
</script>

{#if rows.length > 0}
  <Disclosure summary={simCopy.cooldownTiming} testid="sim-cooldowns">
    <div class="flex flex-col gap-2">
      <p class="text-small text-muted">{simCopy.cooldownNote}</p>
      <ul class="flex flex-col gap-1">
        {#each rows as row (row.id)}
          {@const label = buffLabel(row.id, names)}
          <li class="flex min-h-11 flex-wrap items-center gap-2 text-[14px] md:min-h-9">
            <div class="flex min-w-0 flex-1 items-center gap-2">
              <span class="min-w-0 flex-1 truncate">{label}</span>
              <Select
                id={`sim-cooldown-${row.id}-field`}
                class="w-[150px]"
                size="dense"
                align="end"
                {disabled}
                label={simCopy.cooldownModeFor(label)}
                value={row.mode}
                options={COOLDOWN_MODES.map((mode) => ({
                  value: mode,
                  label: simCopy.cooldownModeLabel[mode],
                }))}
                testid={`sim-cooldown-${row.id}`}
                onchange={(event) =>
                  set(row.id, (event.currentTarget as HTMLSelectElement).value as CooldownMode, row.atSec)}
              />
            </div>
            {#if row.mode === 'at-time'}
              <input
                type="number"
                min="0"
                max={settings.encounter.duration_sec}
                step="1"
                class={control}
                {disabled}
                aria-label={simCopy.cooldownAtFor(label)}
                value={String(row.atSec)}
                onchange={(event) => set(row.id, 'at-time', Number(event.currentTarget.value))}
                data-testid={`sim-cooldown-at-${row.id}`}
              />
            {/if}
          </li>
        {/each}
      </ul>
    </div>
  </Disclosure>
{/if}
