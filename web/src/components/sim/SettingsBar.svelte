<!-- web/src/components/sim/SettingsBar.svelte -->
<!-- The five primary controls that change the number, plus the fight-style picker that
     writes several of them at once. The six secondary controls -- variation, target level,
     armor, type, execute phase and the dummy -- live in SettingsSheet.svelte's disclosure,
     closed on arrival so a control nobody moves costs no phone row. The rotation is text,
     not a disabled select, because the APL builder is deferred and a greyed-out control
     would promise it. -->
<script lang="ts">
  import { simCopy, toolFixCopy } from '../../lib/sim/copy';
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
  import { FIGHT_STYLES, fightStyle, targetsSummary } from '../../lib/sim/styles';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import SettingsSheet from './SettingsSheet.svelte';

  let {
    settings,
    spec,
    disabled,
    onchange,
  }: { settings: SimSettings; spec: string; disabled: boolean; onchange: (next: SimSettings) => void } =
    $props();

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
  // tank MAJOR, review.md:325-327: a timeline style's `targets` field is only the ramp's
  // opening count, so TARGETS reads the ramp itself under one of those styles rather than
  // a number the run does not describe.
  const targetsInfo = $derived(targetsSummary(settings.encounter));
  // Fix round 1: which of the two TARGETS notes is true depends only on `targetsInfo`'s
  // `attached` flag; `targetsTimelineNoteFor` makes that choice, not an `{#if}`/`{:else}`
  // here.
  const targetsNote = $derived(
    targetsInfo.timeline ? toolFixCopy.targetsTimelineNoteFor(targetsInfo.attached) : '',
  );

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
             style by hand (settings.ts's `detached`); it is never a thing to choose, so it
             is hidden the rest of the time rather than offered as a tenth style. -->
        {#if styleId === ''}
          <option value="">{simCopy.styleCustom}</option>
        {/if}
        {#each FIGHT_STYLES as style (style.id)}
          <option value={style.id}>{simCopy.styleLabel[style.id] ?? style.id}</option>
        {/each}
      </select>
    </label>

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

    {#if targetsInfo.timeline}
      <!-- Read-only: the fight style (or, once detached, the ramp it left behind) owns the
           target count here, and editing a ramp with a single-number select would be a lie
           either way. A <div>, not a <label> (fix round 1 Minor 2): a <label> forms no
           accessible-name association with a non-form element, and the rotation block just
           below has the right pattern for exactly this -- a visible label span stacked over
           a value span, no form control involved. Carries the same data-testid the <select>
           below carries so existing tests and e2e still find the control regardless of
           which branch renders. -->
      <div class="flex flex-col gap-1">
        <span class="label text-muted">{simCopy.targets}</span>
        <span class={`${control} flex items-center`} data-testid="sim-targets">
          {toolFixCopy.targetsTimeline(targetsInfo.first, targetsInfo.max)}
        </span>
      </div>
    {:else}
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
    {/if}

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
          <option value={preset.id}>
            {preset.label}
          </option>
        {/each}
      </select>
    </label>

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

  {#if targetsNote !== ''}
    <p class="text-muted text-[12px]" data-testid="sim-targets-note">{targetsNote}</p>
  {/if}

  <SettingsSheet {settings} {disabled} {onchange} />
</section>
