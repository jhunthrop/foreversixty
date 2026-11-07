<!-- web/src/components/report/ReportBar.svelte -->
<!-- The report's one control bar (logs spec 2026-11-07 section 6.1, decision R2): Fight (the
     night's boss pulls numbered per boss, then the trash that killed something; the pulls that
     killed nothing behind the list's last row), Source (the scope words, then the players under
     their roles) and Window (the whole fight, the slices, each phase, each death's 20 s, the
     brush's own range). The three actor tabs add Target, Ability and a Filters button that opens
     the switches in a popover. Group labels show from 1920; below 768 the bar is one summary line
     and Change. Every choice lands in the url through `onPatch`, `onWindow` or `onFilters`. -->
<script lang="ts">
  import { clockLabel } from '../../lib/chart/time-axis';
  import { splitUnitName } from '../../lib/characters';
  import { createCoachViewerState } from '../../lib/coach/viewer-state.svelte';
  import { ownedKey } from '../../lib/coach/viewer';
  import { reportCopy } from '../../lib/report/copy';
  import { abilityOptions, targetOptionId, targetOptions, type ReportFilters } from '../../lib/report/filters';
  import { fightLabel, fightOptions, idleToggleLabel, pullLength } from '../../lib/report/frame';
  import { pullNumbers } from '../../lib/report/fights';
  import type { GroupScope } from '../../lib/report/group';
  import { unitLabel } from '../../lib/report/names.svelte';
  import type { Actor, FightEntry } from '../../lib/report/types';
  import { SOURCE_ENEMIES, SOURCE_EVERYONE, SOURCE_FRIENDLIES, type ReportState } from '../../lib/report/url';
  import { isFullWindow, type TimeWindow, type WindowPreset } from '../../lib/report/window';
  import { loadSpellIconUrls, type IconUrl } from '../../lib/coach/names';
  import activeBuild from '../../data/active-build.json';
  import { selectCopy } from '../../lib/ui/copy';
  import type { SelectOption } from '../../lib/ui/listbox-model';
  import { classOption, iconOption } from '../../lib/ui/select-options';
  import ControlBar from '../ui/ControlBar.svelte';
  import Select from '../ui/Select.svelte';
  import ReportFilterPopover from './ReportFilterPopover.svelte';

  let {
    state: reportState,
    fights,
    roster,
    scope,
    window: current,
    durationMs,
    presets,
    actorTab,
    filters,
    filterActors,
    afterDeathAvailable,
    onPatch,
    onWindow,
    onFilters,
  }: {
    state: ReportState;
    fights: readonly FightEntry[];
    roster: readonly { guid: string; name: string; class?: string; role?: string }[];
    scope: GroupScope;
    window: TimeWindow;
    durationMs: number;
    presets: readonly WindowPreset[];
    /** Damage done, Healing or Damage taken: the bar adds Target, Ability and Filters. */
    actorTab: boolean;
    filters: ReportFilters;
    /** The tab's rows before the filters: what Target and Ability list. */
    filterActors: readonly Actor[];
    afterDeathAvailable: boolean;
    onPatch: (patch: Partial<ReportState>) => void;
    onWindow: (window: TimeWindow | null) => void;
    onFilters: (filters: ReportFilters) => void;
  } = $props();

  /** The value of the Fight list's last row, which shows or hides the pulls that killed nothing. */
  const IDLE_ROW = 'idle';
  /** A list longer than this gets the filter field (controls spec §A10 rows 1 and 6). */
  const SEARCHABLE_OVER = 15;

  let showIdle = $state(false);
  const idleLabel = $derived(idleToggleLabel(fights));
  const fightChoices = $derived<SelectOption[]>([
    ...fightOptions(fights, reportState.fight, showIdle),
    ...(idleLabel === ''
      ? []
      : [
          {
            value: IDLE_ROW,
            label: showIdle ? idleLabel.replace(/^Show/, 'Hide') : idleLabel,
            group: ' ',
          },
        ]),
  ]);

  const ROLE_GROUPS: readonly { role: string; group: string }[] = [
    { role: 'tank', group: reportCopy.sourceGroups.tank },
    { role: 'healer', group: reportCopy.sourceGroups.healer },
    { role: 'dps', group: reportCopy.sourceGroups.dps },
  ];
  const groupOf = (role: string | undefined): string =>
    ROLE_GROUPS.find((entry) => entry.role === role)?.group ?? reportCopy.sourceGroups.dps;
  const coachViewer = createCoachViewerState();
  const sourceChoices = $derived.by((): SelectOption[] => {
    const players = ROLE_GROUPS.flatMap(({ group }) =>
      roster
        .filter((unit) => groupOf(unit.role) === group)
        .map((unit) =>
          classOption({
            value: unit.guid,
            label: unitLabel(unit.name).text,
            className: unit.class,
            group,
            you: coachViewer.viewer.signedIn && ownedKey(coachViewer.viewer, unit.name) !== null,
          }),
        ),
    );
    // A crowded open-world log defaults to the logger's group and offers everyone it saw.
    const friendlies = scope.scoped
      ? [
          {
            value: SOURCE_FRIENDLIES,
            label: scope.noun === 'raid' ? reportCopy.yourRaid : reportCopy.yourGroup,
          },
          { value: SOURCE_EVERYONE, label: reportCopy.everyoneInRange },
        ]
      : [
          scope.crowded
            ? {
                value: reportState.source === SOURCE_EVERYONE ? SOURCE_EVERYONE : SOURCE_FRIENDLIES,
                label: reportCopy.everyoneInRange,
              }
            : { value: SOURCE_FRIENDLIES, label: reportCopy.allFriendlies },
        ];
    return [...friendlies, { value: SOURCE_ENEMIES, label: reportCopy.allEnemies }, ...players];
  });

  const whole = $derived(isFullWindow(current, durationMs));
  const pickedPreset = $derived(
    presets.findIndex((preset) =>
      preset.window === null
        ? whole
        : preset.window.startMs === current.startMs && preset.window.endMs === current.endMs,
    ),
  );
  const windowChoices = $derived([
    ...presets.map((preset, index) => ({ value: String(index), label: preset.label })),
    ...(pickedPreset < 0
      ? [{ value: 'custom', label: `${clockLabel(current.startMs)} to ${clockLabel(current.endMs)}` }]
      : []),
  ]);

  let iconUrl = $state<IconUrl>(() => null);
  $effect(() => {
    if (!actorTab) return;
    void loadSpellIconUrls(activeBuild.build).then((urls) => {
      iconUrl = urls;
    });
  });
  const targets = $derived(actorTab ? targetOptions([...filterActors]) : []);
  const targetChoices = $derived([
    { value: '', label: reportCopy.everyTarget },
    ...targets.map((option) => ({ value: option.id, label: splitUnitName(option.name).name })),
  ]);
  const abilityChoices = $derived([
    { value: '', label: reportCopy.everyAbility },
    ...(actorTab ? abilityOptions([...filterActors]) : []).map((option) =>
      iconOption(option.id, option.name, iconUrl(Number(option.id)), 'spell'),
    ),
  ]);

  const fight = $derived(fights.find((entry) => entry.index === reportState.fight) ?? null);
  const sourceText = $derived(
    sourceChoices.find((option) => option.value === reportState.source)?.label ?? reportCopy.allFriendlies,
  );
  /** The phone line: "Trash 12:43 · Your group". */
  const summary = $derived(
    fight === null
      ? `${reportCopy.pickFight} · ${sourceText}`
      : `${fight.name} ${pullLength(fight.duration_ms)} · ${sourceText}`,
  );
  const pulls = $derived(pullNumbers(fights));

  function value(event: Event): string {
    return (event.currentTarget as HTMLSelectElement).value;
  }
</script>

{#snippet group(label: string, id: string)}
  <span class="report-bar-label text-label tracking-label text-muted font-bold uppercase" id={`${id}-bar-label`}
    >{label}</span
  >
{/snippet}

<ControlBar groups={[]} label="Report choices" {summary} labelsFrom={1920} testid="report-bar">
  {#snippet tail()}
    <div class="report-bar-group" data-testid="report-fight-group">
      {@render group(reportCopy.fight, 'report-fight')}
      <Select
        id="report-fight"
        label={reportCopy.fight}
        size="dense"
        fit
        value={String(reportState.fight)}
        placeholder={fight === null ? reportCopy.pickFight : fightLabel(fight, pulls.get(fight.index)?.pull)}
        options={fightChoices}
        testid="report-fight"
        onchange={(event) => {
          const picked = value(event);
          if (picked === IDLE_ROW) showIdle = !showIdle;
          else if (picked !== '' && picked !== 'none') onPatch({ fight: Number(picked) });
        }}
      />
    </div>
    <div class="report-bar-group">
      {@render group(reportCopy.source, 'report-source')}
      <Select
        id="report-source"
        label={reportCopy.source}
        size="dense"
        fit
        value={reportState.source}
        options={sourceChoices}
        searchable={roster.length > SEARCHABLE_OVER}
        searchPlaceholder={selectCopy.filterPlayers}
        testid="source-scope"
        onchange={(event) => onPatch({ source: value(event) })}
      />
    </div>
    <div class="report-bar-group">
      {@render group(reportCopy.window, 'window-select')}
      <Select
        id="window-select"
        label={reportCopy.window}
        size="dense"
        fit
        options={windowChoices}
        value={pickedPreset < 0 ? 'custom' : String(pickedPreset)}
        testid="window-select"
        onchange={(event) => {
          const picked = value(event);
          if (picked !== 'custom') onWindow(presets[Number(picked)]?.window ?? null);
        }}
      />
    </div>
    {#if actorTab}
      <div class="report-bar-group">
        {@render group(reportCopy.target, 'filter-target')}
        <Select
          id="filter-target"
          label={reportCopy.target}
          size="dense"
          fit
          value={targetOptionId([...filterActors], filters.target)}
          options={targetChoices}
          searchable={targets.length > SEARCHABLE_OVER}
          searchPlaceholder={selectCopy.filterTargets}
          testid="filter-target"
          onchange={(event) => onFilters({ ...filters, target: value(event) })}
        />
      </div>
      <div class="report-bar-group">
        {@render group(reportCopy.ability, 'filter-ability')}
        <Select
          id="filter-ability"
          label={reportCopy.ability}
          size="dense"
          fit
          value={filters.ability === null ? '' : String(filters.ability)}
          options={abilityChoices}
          searchable={abilityChoices.length > SEARCHABLE_OVER}
          searchPlaceholder={selectCopy.filterAbilities}
          testid="filter-ability"
          onchange={(event) => {
            const raw = value(event);
            onFilters({ ...filters, ability: raw === '' ? null : Number(raw) });
          }}
        />
      </div>
      <ReportFilterPopover
        {filters}
        showBossOnly={reportState.tab !== 'healing'}
        {afterDeathAvailable}
        onChange={onFilters}
      />
    {/if}
  {/snippet}
</ControlBar>

<style>
  .report-bar-group {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  @media (min-width: 768px) {
    .report-bar-group {
      flex-direction: row;
      align-items: center;
      gap: 8px;
    }
    .report-bar-label {
      display: none;
    }
  }
  @media (min-width: 1920px) {
    .report-bar-label {
      display: block;
    }
  }
</style>
