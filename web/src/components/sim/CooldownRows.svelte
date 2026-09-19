<!-- web/src/components/sim/CooldownRows.svelte -->
<!-- Design 4.3's last sentence, inside the same panel. One row per scheduled thing: the
     mode, and a seconds field only where the mode needs one -- an always-visible number
     input next to "On cooldown" would invite a time that is never read. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { COOLDOWN_MODES, rowsFor, withCooldown, type CooldownMode } from '../../lib/sim/cooldowns';
  import { buffLabel, type BuffNames } from '../../lib/sim/buff-names';
  import type { SimSettings } from '../../lib/sim/settings';

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

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 border px-2 text-[12px] md:min-h-9';
</script>

{#if rows.length > 0}
  <details class="border-line-soft rounded-panel border" data-testid="sim-cooldowns">
    <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-3 md:min-h-9">
      {simCopy.cooldownTiming}
    </summary>
    <div class="flex flex-col gap-2 p-3 pt-0">
      <p class="text-muted text-[12px]">{simCopy.cooldownNote}</p>
      <ul class="flex flex-col gap-1">
        {#each rows as row (row.id)}
          {@const label = buffLabel(row.id, names)}
          <li class="flex min-h-11 flex-wrap items-center gap-2 text-[13px] md:min-h-9">
            <label class="flex min-w-0 flex-1 items-center gap-2">
              <span class="min-w-0 flex-1 truncate">{label}</span>
              <select
                class={control}
                {disabled}
                aria-label={simCopy.cooldownModeFor(label)}
                value={row.mode}
                onchange={(event) => set(row.id, event.currentTarget.value as CooldownMode, row.atSec)}
                data-testid={`sim-cooldown-${row.id}`}
              >
                {#each COOLDOWN_MODES as mode (mode)}
                  <option value={mode}>{simCopy.cooldownModeLabel[mode]}</option>
                {/each}
              </select>
            </label>
            {#if row.mode === 'at-time'}
              <input
                type="number"
                min="0"
                max={settings.encounter.duration_sec}
                step="1"
                class={`${control} w-20`}
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
  </details>
{/if}
