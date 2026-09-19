<!-- web/src/components/sim/SettingsBar.svelte -->
<!-- The five things that change the number, and nothing else. Variation and the encounter
     profile are fixed at the contract's defaults and stated in the footnote rather than
     exposed: a control nobody moves is a control that costs a phone row. The rotation is
     text, not a disabled select, because the APL builder is deferred and a greyed-out
     control would promise it. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import {
    BUFF_PRESETS,
    DURATIONS,
    MAX_TARGETS,
    durationLabel,
    executePhaseOn,
    withDuration,
    withExecutePhase,
    withPreset,
    withTargets,
    type BuffPresetId,
    type SimSettings,
  } from '../../lib/sim/settings';
  import { specRow } from '../../lib/sim/spec-label';

  let {
    settings,
    spec,
    disabled,
    onchange,
  }: { settings: SimSettings; spec: string; disabled: boolean; onchange: (next: SimSettings) => void } =
    $props();

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 min-w-[7rem] border px-3 text-[14px] font-semibold md:min-h-9';
  const targets = Array.from({ length: MAX_TARGETS }, (_, i) => i + 1);
  // The spec's own display name, from the data lane's generated list -- never the slug
  // with its first letter raised, which turns `beast-mastery` into `Beast-mastery`. An
  // unknown slug falls back to itself, which is what specRow returning null means.
  const rotationName = $derived(specRow(spec)?.name ?? spec);
  const rotationLabel = $derived(`${simCopy.rotationPrefix} ${rotationName}`);
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-4 md:mx-0"
  data-testid="sim-settings"
>
  <div class="flex flex-wrap items-end gap-4 md:gap-5">
    <label class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.fightLength}</span>
      <select
        class={control}
        {disabled}
        value={String(settings.encounter.duration_sec)}
        onchange={(event) => onchange(withDuration(settings, Number(event.currentTarget.value)))}
        data-testid="sim-duration"
      >
        {#each DURATIONS as seconds (seconds)}
          <option value={String(seconds)}>{durationLabel(seconds)}</option>
        {/each}
      </select>
    </label>

    <label class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.targets}</span>
      <select
        class={control}
        {disabled}
        value={String(settings.encounter.targets)}
        onchange={(event) => onchange(withTargets(settings, Number(event.currentTarget.value)))}
        data-testid="sim-targets"
      >
        {#each targets as count (count)}
          <option value={String(count)}>{count}</option>
        {/each}
      </select>
    </label>

    <label class="flex min-h-11 flex-col gap-1 md:min-h-0">
      <span class="label text-muted">{simCopy.executePhase}</span>
      <span class="flex min-h-11 items-center md:min-h-9">
        <input
          type="checkbox"
          class="accent-gold h-5 w-5"
          {disabled}
          checked={executePhaseOn(settings)}
          onchange={(event) => onchange(withExecutePhase(settings, event.currentTarget.checked))}
          data-testid="sim-execute"
        />
      </span>
    </label>

    <label class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.buffs}</span>
      <select
        class={control}
        {disabled}
        value={settings.preset}
        onchange={(event) => onchange(withPreset(settings, event.currentTarget.value as BuffPresetId))}
        data-testid="sim-preset"
      >
        {#each BUFF_PRESETS as preset (preset.id)}
          <option value={preset.id} title={preset.id === 'custom' ? simCopy.customPresetNote : undefined}>
            {preset.label}
          </option>
        {/each}
      </select>
    </label>

    <div class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.rotation}</span>
      <span class="flex min-h-11 items-center gap-2 text-[14px] md:min-h-9">
        <span class="text-strong font-semibold" data-testid="sim-rotation">{rotationLabel}</span>
        <a href={`/sim/specs#${spec}`} class="text-[13px]" data-testid="sim-rotation-link">
          {simCopy.rotationLink}
        </a>
      </span>
    </div>
  </div>

  <p class="text-muted text-[12px]" data-testid="sim-settings-footnote">{simCopy.settingsFootnote}</p>
</section>
