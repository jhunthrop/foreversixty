<!-- web/src/components/sim/SettingsBar.svelte -->
<!-- The five primary controls that change the number, plus the fight-style picker that
     writes several of them at once. The six secondary controls -- variation, target level,
     armor, type, execute phase and the dummy -- live in SettingsSheet.svelte's disclosure,
     closed on arrival so a control nobody moves costs no phone row. The rotation is text,
     not a disabled select, because the APL builder is deferred and a greyed-out control
     would promise it. -->
<script lang="ts">
  import type { BuffNames } from '../../lib/sim/buff-names';
  import { simCopy } from '../../lib/sim/copy';
  import { presetSummary } from '../../lib/sim/preset-summary';
  import {
    BUFF_PRESETS,
    DURATIONS,
    MAX_TARGETS,
    durationLabel,
    styleIdOf,
    withDuration,
    withPreset,
    withStyle,
    withTargets,
    type BuffPresetId,
    type SimSettings,
  } from '../../lib/sim/settings';
  import { FIGHT_STYLES, fightStyle } from '../../lib/sim/styles';
  import { referenceStatOf, specDisplayName } from '../../lib/sim/spec-label';
  import Disclosure from './Disclosure.svelte';
  import HelpNote from './HelpNote.svelte';
  import SettingsSheet from './SettingsSheet.svelte';

  let {
    settings,
    spec,
    disabled,
    onchange,
    /** Task 4: named rows for "what's in it". Null before the build's table has loaded --
     *  buffLabel's own humanised fallback still names every row, so the disclosure is never
     *  wrong to open early. */
    names = null,
  }: {
    settings: SimSettings;
    spec: string;
    disabled: boolean;
    onchange: (next: SimSettings) => void;
    names?: BuffNames | null;
  } = $props();

  const referenceStat = $derived(referenceStatOf(spec));
  // Custom has its own always-visible panel (BuffPanel.svelte) for exactly this, so the
  // disclosure only exists for the two static presets it actually summarises.
  const presetGroups = $derived(
    settings.preset === 'custom' ? [] : presetSummary(settings.preset, referenceStat, names),
  );

  // min-w-0 on phone, not min-w-[7rem]: at 412px, five selects each demanding 112px force
  // the row into an uneven two-column wrap with the gutter closed to 4px. Restored at md
  // and up, where the row has the width to spare and the fixed minimum keeps every select
  // the same size regardless of its shortest option's own text.
  const control =
    'border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 md:min-w-[7rem] border px-3 text-[14px] font-semibold md:min-h-9';
  const targets = Array.from({ length: MAX_TARGETS }, (_, i) => i + 1);
  const rotationLabel = $derived(`${simCopy.rotationPrefix} ${specDisplayName(spec)}`);
  const styleId = $derived(styleIdOf(settings));
  // Only the two movement styles carry a note (copy.ts's styleNote). Everything else
  // renders nothing at all rather than an empty paragraph that would reserve a line.
  const styleNote = $derived(simCopy.styleNote[styleId] ?? '');

  // The <select>'s value is a plain string; FightStyleId is a literal union, so a raw cast
  // would let an id outside the contract's nine reach withStyle/applyFightStyle, which
  // (styles.ts) returns the encounter unchanged for an id it does not recognise -- the one
  // branch of that function that is not a fresh object. `fightStyle` is styles.ts's own
  // safe, null-returning lookup for exactly this; re-deriving it here with a second
  // `FIGHT_STYLES.find` would be the same rule kept in two places.
  function selectStyle(id: string): void {
    const style = fightStyle(id);
    if (style === null) return;
    onchange(withStyle(settings, style.id));
  }
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-4 md:mx-0"
  data-testid="sim-settings"
>
  <div class="flex flex-wrap items-end gap-4 md:gap-5">
    <div class="flex flex-col gap-1">
      <label class="flex flex-col gap-1">
        <span class="label text-muted">{simCopy.fightStyle}</span>
        <select
          class={control}
          {disabled}
          value={styleId}
          onchange={(event) => selectStyle(event.currentTarget.value)}
          data-testid="sim-style"
        >
          <!-- The empty option exists only while the encounter has been detached from a
               style by hand (settings.ts's `detached`); it is never a thing to choose, so
               it is hidden the rest of the time rather than offered as a tenth style. -->
          {#if styleId === ''}
            <option value="">{simCopy.styleCustom}</option>
          {/if}
          {#each FIGHT_STYLES as style (style.id)}
            <option value={style.id}>{simCopy.styleLabel[style.id] ?? style.id}</option>
          {/each}
        </select>
      </label>
      <HelpNote label={simCopy.fightStyle} id="sim-style" {disabled}>
        <p>{simCopy.fightStyleHelp}</p>
        <dl class="flex flex-col gap-1" data-testid="sim-style-help-options">
          {#each FIGHT_STYLES as style (style.id)}
            <div>
              <dt class="text-text font-semibold">{simCopy.styleLabel[style.id] ?? style.id}</dt>
              <dd>{simCopy.fightStyleOptions[style.id] ?? ''}</dd>
            </div>
          {/each}
        </dl>
      </HelpNote>
    </div>

    <div class="flex flex-col gap-1">
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
      <HelpNote label={simCopy.fightLength} id="sim-duration" {disabled}>
        <p>{simCopy.fightLengthHelp}</p>
      </HelpNote>
    </div>

    <div class="flex flex-col gap-1">
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
      <HelpNote label={simCopy.targets} id="sim-targets" {disabled}>
        <p>{simCopy.targetsHelp}</p>
      </HelpNote>
    </div>

    <div class="flex flex-col gap-1">
      <label class="flex flex-col gap-1">
        <span class="label text-muted">{simCopy.buffs}</span>
        <select
          class={control}
          {disabled}
          value={settings.preset}
          onchange={(event) =>
            onchange(withPreset(settings, event.currentTarget.value as BuffPresetId, referenceStat))}
          data-testid="sim-preset"
        >
          {#each BUFF_PRESETS as preset (preset.id)}
            <option value={preset.id}>
              {preset.label}
            </option>
          {/each}
        </select>
      </label>
      <div class="flex flex-wrap items-center gap-3">
        <HelpNote label={simCopy.buffs} id="sim-preset" {disabled}>
          <p>{simCopy.buffsHelp}</p>
        </HelpNote>
        {#if settings.preset !== 'custom'}
          <Disclosure
            label={simCopy.whatsInIt}
            id="sim-preset-summary"
            {disabled}
            triggerClass="label text-nav text-[12px] underline decoration-dotted underline-offset-2"
            panelClass="border-line-soft rounded-panel flex flex-col gap-2 border p-3"
            triggerTestId="sim-preset-summary-trigger"
            panelTestId="sim-preset-summary-panel"
          >
            {#if presetGroups.length === 0}
              <p class="text-muted text-[12px]">{simCopy.whatsInItEmpty}</p>
            {/if}
            {#each presetGroups as group (group.group)}
              <div>
                <p class="label text-muted text-[11px]">
                  {simCopy.buffGroupLabel[group.group] ?? group.group}
                </p>
                <ul class="text-text flex flex-wrap gap-x-3 gap-y-1 text-[12px]">
                  {#each group.rows as row (row.id)}
                    <li data-testid={`sim-preset-summary-${row.id}`}>{row.label}</li>
                  {/each}
                </ul>
              </div>
            {/each}
          </Disclosure>
        {/if}
      </div>
    </div>

    <div class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.rotation}</span>
      <span class="flex min-h-11 items-center gap-2 text-[14px] md:min-h-9">
        <span class="text-strong font-semibold" data-testid="sim-rotation">{rotationLabel}</span>
        <!-- 68px wide already clears the 44px hit-target floor, but its own line-height
             does not; the row around it already reserves min-h-11 (44px) on mobile, so
             giving the link itself the same min-height fills that already-reserved space
             rather than growing the row further. -->
        <a
          href={`/sim/specs#${spec}`}
          class="inline-flex min-h-11 items-center text-[13px] md:min-h-0"
          data-testid="sim-rotation-link"
        >
          {simCopy.rotationLink}
        </a>
      </span>
    </div>
  </div>

  {#if styleNote !== ''}
    <p class="text-muted text-[12px]" data-testid="sim-style-note">{styleNote}</p>
  {/if}

  <SettingsSheet {settings} {disabled} {onchange} />
</section>
