<!-- web/src/components/chart/PlayersControl.svelte -->
<!-- The chart head's Players control (controls and charts spec §A4.5, §B5.3): who the chart draws,
     on the shared multi-select. A Compare group (the raid median, a class median), then Tanks,
     Healers and Damage with crests, class colours, the You tag and each player's figure; an
     always-drawn player is on and disabled. The foot holds the presets (Top 4 + bottom 4, Top 8,
     Bottom 8) and Custom, lit while the reader's own set is drawn, with the drawn count. It opens
     bottom-end in the top layer; on a coarse pointer it opens the legend sheet (onCoarse). This
     file only maps the chart's model onto Select. -->
<script lang="ts">
  import { chartCopy } from '../../lib/chart/copy';
  import type { DrawPreset, PlayersChoice } from '../../lib/chart/drawn';
  import type { MedianInfo } from '../../lib/chart/rate';
  import type { PlayerEntry } from '../../lib/chart/types';
  import type { FootAction, SelectOption } from '../../lib/ui/listbox-model';
  import { classOption } from '../../lib/ui/select-options';
  import Select from '../ui/Select.svelte';

  let {
    id,
    entries,
    medians = [],
    choice,
    presets,
    seen = false,
    formatValue,
    onToggle,
    onMedian,
    onPreset,
    onCoarse,
  }: {
    id: string;
    entries: readonly PlayerEntry[];
    medians?: readonly MedianInfo[];
    /** What is drawn now: a preset, or the reader's own set. */
    choice: PlayersChoice;
    /** The presets this chart offers: all three on the damage chart, Top 8 on threat. */
    presets: readonly DrawPreset[];
    /** The entries are everyone in range: the field says "players seen". */
    seen?: boolean;
    formatValue: (value: number) => string;
    onToggle: (id: string) => void;
    onMedian?: (id: string) => void;
    onPreset: (preset: DrawPreset) => void;
    onCoarse: () => void;
  } = $props();

  const ROLE_ORDER = ['tank', 'healer', 'dps', 'other'] as const;

  const options = $derived<SelectOption[]>([
    ...medians.map((median) => ({
      value: median.id,
      label: median.label,
      description: median.description,
      group: chartCopy.groups.compare,
    })),
    ...ROLE_ORDER.flatMap((role) =>
      entries
        .filter((entry) => entry.role === role)
        .map((entry) => ({
          ...classOption({
            value: entry.id,
            label: entry.label,
            className: entry.className,
            group: chartCopy.groups[role],
            you: entry.you,
            meta: formatValue(entry.figure),
            description: entry.pinned ? chartCopy.alwaysDrawn : undefined,
          }),
          disabled: entry.pinned,
        })),
    ),
  ]);
  const drawn = $derived([
    ...medians.filter((median) => median.on).map((median) => median.id),
    ...entries.filter((entry) => entry.drawn).map((entry) => entry.id),
  ]);
  const fieldText = $derived(
    chartCopy.playersField(entries.filter((entry) => entry.drawn).length, entries.length, choice, seen),
  );
  const footActions = $derived<FootAction[]>([
    ...presets.map((preset) => ({
      id: preset,
      label: chartCopy.presets[preset],
      pressed: choice === preset,
    })),
    { id: 'custom', label: chartCopy.presets.custom, pressed: choice === 'custom', inert: true },
  ]);
  const medianIds = $derived(new Set(medians.map((median) => median.id)));

  /** One toggle at a time reaches the chart: the id whose state differs from what is drawn. */
  function onchange(event: Event): void {
    const picked = new Set(
      [...(event.currentTarget as HTMLSelectElement).selectedOptions].map((o) => o.value),
    );
    const before = new Set(drawn);
    const toggled = [...picked, ...before].find((value) => picked.has(value) !== before.has(value));
    if (toggled === undefined) return;
    if (medianIds.has(toggled)) onMedian?.(toggled);
    else onToggle(toggled);
  }

  function onfoot(action: string): void {
    const preset = presets.find((candidate) => candidate === action);
    if (preset !== undefined) onPreset(preset);
  }
</script>

<Select
  {id}
  class="{seen ? 'w-[260px]' : 'w-[196px]'} max-md:w-full"
  size="dense"
  align="end"
  multiple
  values={drawn}
  {options}
  label={chartCopy.players}
  fieldLabel={fieldText}
  searchPlaceholder={chartCopy.filterPlayers}
  {footActions}
  footCount={chartCopy.drawnCount(drawn.length)}
  testid="players"
  {onchange}
  {onfoot}
  oncoarseopen={onCoarse}
/>
