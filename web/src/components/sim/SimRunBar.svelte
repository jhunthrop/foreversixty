<!-- web/src/components/sim/SimRunBar.svelte -->
<!-- The run line (spec 2026-11-06-sim §4.1, mock web-sim-you): the ControlBar with Fight, Length,
     Targets, Buffs and Precision, then "More settings" (the overlay, ruling S7) and the page's one
     primary Button. Sticky under the site header; on phone the summary line
     "{style} {length} · {buffs}", Change, and the primary on the same line. Every group and More
     settings are inert at 50% while a run is in flight (the shipped rule); the primary stays live
     (Stop). Precision reads in runs, the one count every part of the page uses. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { PrimaryAction } from '../../lib/sim/answer';
  import { simCopy } from '../../lib/sim/copy';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { PRECISIONS, type PrecisionId } from '../../lib/sim/precision';
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
  import type { ControlGroup } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import ControlBar from '../ui/ControlBar.svelte';

  let {
    settings,
    precisionId,
    spec,
    level,
    primary,
    busy,
    onsettings,
    onprecision,
    onprimary,
    onmore,
    precisionOptions,
    primaryLabel,
    primaryTestid = 'sim-run-button',
    testid = 'sim-settings',
    extra,
  }: {
    settings: SimSettings;
    precisionId: PrecisionId;
    spec: string;
    level: number;
    primary: PrimaryAction;
    /** A run is in flight on either lane. */
    busy: boolean;
    onsettings: (next: SimSettings) => void;
    onprecision: (value: PrecisionId) => void;
    onprimary: () => void;
    onmore: () => void;
    /** The precision choices, when the page's runs are counted differently (the bulk tools). */
    precisionOptions?: readonly { value: string; label: string }[];
    /** The primary's words, over the quick sim's own for `primary`. */
    primaryLabel?: string;
    primaryTestid?: string;
    testid?: string;
    /** A second action before More settings (the bulk tools' "Run on our servers"). */
    extra?: Snippet;
  } = $props();

  const PARAM = {
    style: 'style',
    length: 'length',
    targets: 'targets',
    buffs: 'buffs',
    precision: 'precision',
  } as const;

  const styleId = $derived(styleIdOf(settings));
  const groups = $derived<ControlGroup[]>([
    {
      label: copy.bar.fight,
      param: PARAM.style,
      kind: 'select',
      value: styleId,
      options: [
        ...(styleId === '' ? [{ value: '', label: simCopy.styleCustom }] : []),
        ...FIGHT_STYLES.map((style) => ({
          value: style.id,
          label: simCopy.styleLabel[style.id] ?? style.id,
          description: simCopy.fightStyleOptions[style.id],
        })),
      ],
    },
    {
      label: copy.bar.length,
      param: PARAM.length,
      kind: 'select',
      value: String(settings.encounter.duration_sec),
      options: [
        ...(DURATIONS.includes(settings.encounter.duration_sec) ? [] : [settings.encounter.duration_sec]),
        ...DURATIONS,
      ].map((seconds) => ({ value: String(seconds), label: durationLabel(seconds) })),
    },
    {
      label: copy.bar.targets,
      param: PARAM.targets,
      kind: 'select',
      value: String(settings.encounter.targets),
      options: Array.from({ length: MAX_TARGETS }, (_, index) => ({
        value: String(index + 1),
        label: String(index + 1),
      })),
    },
    {
      label: copy.bar.buffs,
      param: PARAM.buffs,
      kind: 'segmented',
      value: settings.preset,
      options: BUFF_PRESETS.map((preset) => ({
        value: preset.id,
        label: copy.bar.buffOptions[preset.id] ?? preset.label,
      })),
    },
    {
      label: copy.bar.precision,
      param: PARAM.precision,
      kind: 'select',
      value: precisionId,
      options:
        precisionOptions ??
        PRECISIONS.map((id) => ({ value: id, label: copy.bar.precisionOptions[id] ?? id })),
    },
  ]);

  const buffsWord = $derived(copy.bar.buffOptions[settings.preset] ?? settings.preset);
  const summary = $derived(
    `${simCopy.styleLabel[styleId] ?? simCopy.styleCustom} ${durationLabel(settings.encounter.duration_sec)} · ${buffsWord}`,
  );

  function onchange(param: string, value: string): void {
    if (param === PARAM.precision) {
      onprecision(value as PrecisionId);
      return;
    }
    if (param === PARAM.style) {
      const style = fightStyle(value);
      if (style !== null) onsettings(withStyle(settings, style.id));
      return;
    }
    if (param === PARAM.length) onsettings(withDuration(settings, Number(value)));
    else if (param === PARAM.targets) onsettings(withTargets(settings, Number(value)));
    else if (param === PARAM.buffs) onsettings(withPreset(settings, value as BuffPresetId, spec, level));
  }

  const LABEL: Readonly<Record<PrimaryAction, string>> = {
    run: copy.bar.runSim,
    'run-again': copy.bar.runAgain,
    'try-again': copy.bar.tryAgain,
    stop: copy.bar.stop,
    'loading-engine': copy.bar.loadingEngine,
    server: copy.bar.runningOnServers,
    disabled: copy.bar.runSim,
  };
  const primaryTone = $derived(
    primary === 'stop' || primary === 'loading-engine' || primary === 'server' ? 'secondary' : 'primary',
  );
  const primaryDisabled = $derived(
    primary === 'loading-engine' || primary === 'server' || primary === 'disabled',
  );
</script>

<ControlBar
  {groups}
  label={copy.bar.label}
  {summary}
  labelsFrom={1920}
  {busy}
  busyTitle={copy.bar.busyTitle}
  {onchange}
  id="sim-bar"
  {testid}
>
  {#snippet end()}
    {#if extra}{@render extra()}{/if}
    <button type="button" class="sim-more text-gold" onclick={onmore} data-testid="sim-settings-more"
      >{copy.bar.moreSettings}</button
    >
  {/snippet}
  {#snippet action()}
    <Button tone={primaryTone} disabled={primaryDisabled} onclick={onprimary} testid={primaryTestid}
      >{primaryLabel ?? LABEL[primary]}</Button
    >
  {/snippet}
</ControlBar>

<style>
  .sim-more {
    font-size: 14px;
    font-weight: 600;
    min-height: 44px;
    background: none;
    border: 0;
    padding: 0;
  }
  .sim-more:hover {
    color: var(--color-gold-hover);
  }
  .sim-more:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
</style>
