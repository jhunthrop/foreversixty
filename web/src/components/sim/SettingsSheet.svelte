<!-- web/src/components/sim/SettingsSheet.svelte -->
<!-- The six controls that change a fight without changing what most players are asking.
     A <details> disclosure, the same collapsible report/Glossary.svelte uses: closed on
     first render, so it adds no height to the settings bar's own reservation, and opening
     it is a user gesture, which is the one kind of layout shift the CLS budget excludes.

     Variation, target level, armor and type are the player's and never touch the fight
     style's name. Execute phase and the dummy are the style's, so setting either detaches
     the encounter from its style -- settings.ts's `detached`, not this component's. -->
<script lang="ts">
  import { simCopy, toolFixCopy } from '../../lib/sim/copy';
  import {
    MAX_TARGET_ARMOR,
    TARGET_LEVELS,
    TARGET_TYPES,
    VARIATIONS,
    executePhaseOn,
    targetArmorField,
    withDummy,
    withExecutePhase,
    withTargetArmor,
    withTargetLevel,
    withTargetType,
    withVariation,
    type SimSettings,
  } from '../../lib/sim/settings';

  let {
    settings,
    disabled,
    onchange,
  }: { settings: SimSettings; disabled: boolean; onchange: (next: SimSettings) => void } = $props();

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 border px-3 text-[14px] font-semibold md:min-h-9';
  const percent = (value: number): string => `${Math.round(value * 100)}%`;
  // tank MAJOR, review.md:227-229: armor 0 means the level's preset, not an empty field --
  // `targetArmorField` decides what the input DISPLAYS; the wire value settings.ts sends
  // stays 0 until the player types something else.
  const armorField = $derived(targetArmorField(settings.encounter));
  // Contract A8's figure for whichever level is chosen, so an empty armor field says what
  // the engine will use instead of nothing at all. Fix round 1 Minor 1: derived from
  // `armorField.preset` -- settings.ts's own single source of truth for this number --
  // rather than a second, independent `TARGET_ARMOR_BY_LEVEL` lookup that only agreed with
  // it by coincidence (it used a different fallback, `?? 63` vs. `DEFAULT_TARGET_LEVEL`).
  const armorPreset = $derived(simCopy.targetArmorPreset(armorField.preset.toLocaleString('en-US')));
</script>

<details class="border-line-soft rounded-panel border" data-testid="sim-settings-more">
  <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-3 md:min-h-9">
    {simCopy.moreSettings}
  </summary>

  <div class="grid grid-cols-2 gap-3 p-3 pt-0 md:grid-cols-4">
    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.variation}</span>
      <select
        class={control}
        {disabled}
        value={String(settings.encounter.variation)}
        onchange={(event) => onchange(withVariation(settings, Number(event.currentTarget.value)))}
        data-testid="sim-variation"
      >
        {#each VARIATIONS as value (value)}
          <option value={String(value)}>{percent(value)}</option>
        {/each}
      </select>
      <!-- Task 5 (newcomer MAJOR, review.md:360-363): was a hover-only `title`, invisible on
           a phone. Visible text instead, the same treatment Task 4 already gave the
           target-armor field in this same grid -- one more field, the same house pattern. -->
      <p class="text-muted text-[12px]" data-testid="sim-variation-note">{simCopy.variationNote}</p>
    </label>

    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.targetLevel}</span>
      <select
        class={control}
        {disabled}
        value={String(settings.encounter.target_level ?? 63)}
        onchange={(event) => onchange(withTargetLevel(settings, Number(event.currentTarget.value)))}
        data-testid="sim-target-level"
      >
        {#each TARGET_LEVELS as level (level)}
          <option value={String(level)}>{level}</option>
        {/each}
      </select>
    </label>

    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.targetArmor}</span>
      <input
        type="number"
        min="0"
        max={MAX_TARGET_ARMOR}
        step="1"
        class={control}
        {disabled}
        placeholder={armorPreset}
        value={armorField.value}
        onchange={(event) => onchange(withTargetArmor(settings, Number(event.currentTarget.value)))}
        data-testid="sim-target-armor"
      />
      <!-- Fix round, Minor 1: `armorField.level`, not a second `settings.encounter.
           target_level ?? DEFAULT_TARGET_LEVEL` lookup -- `targetArmorField` already
           computed the level `preset` came from, so this reads that same value instead of
           re-deriving one that only agreed with it by coincidence. -->
      <p class="text-muted text-[12px]" data-testid="sim-target-armor-note">
        {toolFixCopy.targetArmorNote(armorField.level, armorField.preset)}
      </p>
    </label>

    <label class="flex min-w-0 flex-col gap-1">
      <span class="label text-muted">{simCopy.targetType}</span>
      <select
        class={control}
        {disabled}
        value={settings.encounter.target_type ?? ''}
        onchange={(event) => onchange(withTargetType(settings, event.currentTarget.value))}
        data-testid="sim-target-type"
      >
        <option value="">{simCopy.targetTypeAny}</option>
        {#each TARGET_TYPES as id (id)}
          <option value={id}>{simCopy.targetTypeLabel[id] ?? id}</option>
        {/each}
      </select>
    </label>

    <label class="flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="accent-gold h-5 w-5"
        {disabled}
        checked={executePhaseOn(settings)}
        onchange={(event) => onchange(withExecutePhase(settings, event.currentTarget.checked))}
        data-testid="sim-execute"
      />
      <span class="text-muted">{simCopy.executePhase}</span>
    </label>

    <div class="flex min-w-0 flex-col gap-1">
      <label class="flex min-h-11 items-center gap-2 text-[13px]">
        <input
          type="checkbox"
          class="accent-gold h-5 w-5"
          {disabled}
          checked={settings.encounter.dummy === true}
          onchange={(event) => onchange(withDummy(settings, event.currentTarget.checked))}
          data-testid="sim-dummy"
        />
        <span class="text-muted">{simCopy.dummyTarget}</span>
      </label>
      <!-- Task 5 (newcomer MAJOR, review.md:360-363): was a hover-only `title`, invisible on
           a phone. Visible text instead, the same treatment Task 4 already gave the
           target-armor field above. -->
      <p class="text-muted text-[12px]" data-testid="sim-dummy-note">{simCopy.dummyNote}</p>
    </div>
  </div>
</details>
