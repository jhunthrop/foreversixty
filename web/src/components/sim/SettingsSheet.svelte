<!-- web/src/components/sim/SettingsSheet.svelte -->
<!-- The six controls that change a fight without changing what most players are asking.
     A <details> disclosure, the same collapsible report/Glossary.svelte uses: closed on
     first render, so it adds no height to the settings bar's own reservation, and opening
     it is a user gesture, which is the one kind of layout shift the CLS budget excludes.

     Variation, target level, armor and type are the player's and never touch the fight
     style's name. Execute phase and the dummy are the style's, so setting either detaches
     the encounter from its style -- settings.ts's `detached`, not this component's. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import {
    MAX_TARGET_ARMOR,
    TARGET_ARMOR_BY_LEVEL,
    TARGET_LEVELS,
    TARGET_TYPES,
    VARIATIONS,
    executePhaseOn,
    styleIdOf,
    withDummy,
    withExecutePhase,
    withTargetArmor,
    withTargetLevel,
    withTargetType,
    withVariation,
    type SimSettings,
  } from '../../lib/sim/settings';
  import HelpNote from './HelpNote.svelte';

  let {
    settings,
    disabled,
    onchange,
  }: { settings: SimSettings; disabled: boolean; onchange: (next: SimSettings) => void } = $props();

  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 border px-3 text-[14px] font-semibold md:min-h-9';
  const percent = (value: number): string => `${Math.round(value * 100)}%`;
  // Contract A8's figure for whichever level is chosen, so an empty armor field says what
  // the engine will use instead of nothing at all.
  const armorPreset = $derived(
    simCopy.targetArmorPreset(
      (TARGET_ARMOR_BY_LEVEL[settings.encounter.target_level ?? 63] ?? 0).toLocaleString('en-US'),
    ),
  );
  // Task 7 (newcomer MINOR 213-216): the field's own current state, not the placeholder
  // repeated as a tooltip -- `target_armor: 0` is settings.ts's "use the level's preset".
  const armorEmpty = $derived(settings.encounter.target_armor === 0);
  const armorHelp = $derived(
    armorEmpty
      ? simCopy.targetArmorEmptyHelp(armorPreset)
      : simCopy.targetArmorSetHelp(
          (settings.encounter.target_armor ?? 0).toLocaleString('en-US'),
          armorPreset,
        ),
  );
  // Task 7 (newcomer MINOR 204-207): "Execute phase" said nothing about whether the current
  // fight style actually carries one. `styleIdOf` reads the same style the settings bar's
  // own select shows -- "Custom" once a style-owned field (this one included) is detached.
  const currentStyleId = $derived(styleIdOf(settings));
  const currentStyleLabel = $derived(
    currentStyleId === '' ? simCopy.styleCustom : (simCopy.styleLabel[currentStyleId] ?? currentStyleId),
  );
  const executeOn = $derived(executePhaseOn(settings));
  const executePercent = $derived(String(Math.round(settings.encounter.execute_ratio * 100)));
</script>

<details class="border-line-soft rounded-panel border" data-testid="sim-settings-more">
  <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-3 md:min-h-9">
    {simCopy.moreSettings}
  </summary>

  <div class="grid grid-cols-2 gap-3 p-3 pt-0 md:grid-cols-4">
    <div class="flex min-w-0 flex-col gap-1">
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
      </label>
      <HelpNote label={simCopy.variation} id="sim-variation" {disabled}>
        <p>{simCopy.variationNote}</p>
      </HelpNote>
    </div>

    <div class="flex min-w-0 flex-col gap-1">
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
      <HelpNote label={simCopy.targetLevel} id="sim-target-level" {disabled}>
        <p>{simCopy.targetLevelHelp}</p>
      </HelpNote>
    </div>

    <div class="flex min-w-0 flex-col gap-1">
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
          value={settings.encounter.target_armor === 0 ? '' : String(settings.encounter.target_armor)}
          onchange={(event) => onchange(withTargetArmor(settings, Number(event.currentTarget.value)))}
          data-testid="sim-target-armor"
        />
      </label>
      <HelpNote label={simCopy.targetArmor} id="sim-target-armor" {disabled}>
        <p>{armorHelp}</p>
      </HelpNote>
    </div>

    <div class="flex min-w-0 flex-col gap-1">
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
      <HelpNote label={simCopy.targetType} id="sim-target-type" {disabled}>
        <p>{simCopy.targetTypeHelp}</p>
      </HelpNote>
    </div>

    <div class="flex flex-col gap-1">
      <label class="flex min-h-11 items-center gap-2 text-[13px]">
        <input
          type="checkbox"
          class="accent-gold h-5 w-5"
          {disabled}
          checked={executeOn}
          onchange={(event) => onchange(withExecutePhase(settings, event.currentTarget.checked))}
          data-testid="sim-execute"
        />
        <span class="text-muted">{simCopy.executePhase}</span>
      </label>
      <HelpNote label={simCopy.executePhase} id="sim-execute" {disabled}>
        <p>{simCopy.executePhaseHelp(executeOn, executePercent, currentStyleLabel)}</p>
      </HelpNote>
    </div>

    <div class="flex flex-col gap-1">
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
      <HelpNote label={simCopy.dummyTarget} id="sim-dummy" {disabled}>
        <p>{simCopy.dummyNote}</p>
      </HelpNote>
    </div>
  </div>
</details>
