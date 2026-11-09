<!-- web/src/components/guild/GuildLoot.svelte -->
<!-- The Loot tab as the loot council (web layout spec 2026-11-05 §7.2): "This dropped. Who gets it,
     and why?" Page split, built only from the ui/ primitives. The control bar: the raid night. Main:
     the one Plate of that night's drops (Open or Awarded), and on the current night a picker that
     opens any drop of the loot table. Rail, officers: the picked drop's candidates with gain, what
     it replaces, BiS for them, attendance, loot this tier, MS/OS, Votes and Award; members: their
     own MS / OS / None per character that can use it (never anyone's votes). Then the night's
     decision record, with Export history for officers. Votes and Award take the Raid guild plan:
     without it the shipped gate sentence leads and the controls stay on screen, disabled. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { fetchItemTooltipModel } from '../../lib/items/lookup';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import { LootCouncilState, type LootCouncilDeps } from '../../lib/guild/loot-council-state.svelte';
  import {
    NO_DROPS_MEMBER,
    NO_DROPS_OFFICER,
    NO_DROPS_PAST,
    PICK_PROMPT,
    dropRows,
    nightDay,
    nightOptions,
    nightSummary,
    openable,
    ownUsers,
    type DropRow,
  } from '../../lib/guild/loot-council-view';
  import { LOOT_LOAD_ERROR } from '../../lib/guild/loot-view';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import ControlBar from '../ui/ControlBar.svelte';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import DecisionRecord from './loot/DecisionRecord.svelte';
  import LootCandidates from './loot/LootCandidates.svelte';
  import LootDrops from './loot/LootDrops.svelte';
  import MemberFlag from './loot/MemberFlag.svelte';

  let {
    guildId,
    officer,
    officerTools,
    raidGuildTools,
    ownKeys = [],
    gateFacts,
    timeZone,
    onRefused,
    deps,
  }: {
    guildId: number;
    officer: boolean;
    /** False when the guild's officer tools are off (gate spec §3.1). */
    officerTools: boolean;
    /** False for an officer of a Guild-tier guild: Loot is part of Raid guild. */
    raidGuildTools: boolean;
    /** The viewer's own characters: a member flags only these. */
    ownKeys?: readonly string[];
    gateFacts: GateFacts;
    /** The guild's raid zone, for the decision record's day. */
    timeZone?: string;
    onRefused: (gate: 'plan' | 'raid-guild') => void;
    /** Tests and captures inject the API. */
    deps?: LootCouncilDeps;
  } = $props();

  const build = activeBuild.build;
  /** Up to this many nights read as a Segmented; more take the dropdown. */
  const SEGMENTED_NIGHTS = 3;
  const GATE_ID = 'officer-tools-gate-loot';

  // The tab mounts once per guild and role; the controller keeps its own state from there.
  // svelte-ignore state_referenced_locally
  const council = new LootCouncilState({ guildId, officer, onRefused: (gate) => onRefused(gate), deps });
  $effect(() => {
    void council.load();
  });

  const canAct = $derived(officerTools && raidGuildTools);
  const night = $derived(council.night);
  const rows = $derived(night === null ? [] : dropRows(night, council.table));
  const picks = $derived(night === null ? [] : openable(night, council.table, officer ? null : ownKeys));
  const options = $derived(council.nights === null ? [] : nightOptions(council.nights.nights));
  /** The page's one control bar: the raid night, a Segmented up to three nights, the dropdown past. */
  const nightGroup = $derived({
    label: 'Raid night',
    param: 'night',
    kind: options.length <= SEGMENTED_NIGHTS ? ('segmented' as const) : ('select' as const),
    options,
    value: council.nightKey ?? '',
    fieldWidth: 192,
  });
  const tableItem = (itemId: number) => council.table.find((item) => item.item_id === itemId);
  const shipsIcon = (itemId: number): boolean => tableItem(itemId)?.kind === 'gear';
  const emptyText = $derived(
    night === null || !night.current ? NO_DROPS_PAST : officer ? NO_DROPS_OFFICER : NO_DROPS_MEMBER,
  );
  const memberFlags = $derived.by(() => {
    const itemId = council.selected?.itemId;
    const out: Record<string, 'ms' | 'os'> = {};
    for (const [key, value] of Object.entries(council.flags)) {
      const [item, character] = key.split(/:(.*)/s);
      if (Number(item) === itemId && character !== undefined) out[character] = value;
    }
    return out;
  });

  /** The class whose item file carries the drop, for its tooltip. */
  function classOf(itemId: number): string {
    const item = tableItem(itemId);
    return classSlugFromName(item?.candidates[0]?.class ?? item?.eligible[0]?.class ?? 'warrior');
  }

  /** Below 1024 the rail follows the Plate: bring the picked drop's rail into view. */
  function pick(row: DropRow): void {
    council.select(row);
    if (typeof window !== 'undefined' && window.innerWidth < 1024)
      requestAnimationFrame(() =>
        document.getElementById('guild-loot-rail')?.scrollIntoView({ block: 'start' }),
      );
  }
</script>

<div data-testid="guild-loot-tab">
  {#if officer && !canAct}<GuildPlanGate tab="loot" facts={gateFacts} />{/if}
  {#if council.status === 'loading'}
    <Skeleton lines={6} minHeight="min-h-[360px]" testid="guild-loot-skeleton" />
  {:else if council.status === 'failed' || night === null}
    <LoadError message={LOOT_LOAD_ERROR} testid="guild-loot-error" onRetry={() => void council.load()} />
  {:else}
    <Page layout="split" panels testid="guild-loot-page">
      {#snippet bar()}
        <ControlBar
          groups={[nightGroup]}
          label="Loot council choices"
          id="guild-loot-bar"
          testid="guild-loot-nights"
          onchange={(_, value) => council.chooseNight(value)}
        />
      {/snippet}
      <LootDrops
        heading="Onyxia drops"
        meta={`${nightDay(night.date)} · ${nightSummary(night)}`}
        {rows}
        openable={picks}
        raids={council.raids}
        selected={council.selected?.itemId ?? null}
        {build}
        {emptyText}
        onSelect={pick}
        onOpen={council.open}
      />
      {#snippet rail()}
        <div id="guild-loot-rail" class="flex scroll-mt-[96px] flex-col gap-[22px]">
          {#if council.selected === null}
            <Section
              title={officer ? 'Candidates' : 'Your call'}
              panel
              frame="framed"
              testid="guild-loot-prompt"
            >
              <p class="text-muted text-[14px]">{PICK_PROMPT}</p>
            </Section>
          {:else if officer}
            <LootCandidates
              itemName={council.selected.name}
              drop={council.drop}
              status={council.dropStatus}
              {build}
              shipsIcon={council.selected.kind === 'gear'}
              {timeZone}
              {canAct}
              describedBy={canAct ? undefined : GATE_ID}
              busy={council.busy}
              confirmingKey={council.confirmingKey}
              notice={council.notice}
              onVote={(key, vote) => void council.vote(key, vote)}
              onAsk={council.ask}
              onAward={(key, msOs, reason) => void council.award(key, msOs, reason)}
              onCancel={council.cancel}
              onUndo={() => void council.undo()}
              onRetry={() => void council.readDrop()}
            />
          {:else}
            <MemberFlag
              itemName={council.selected.name}
              award={council.selected.award}
              users={ownUsers(tableItem(council.selected.itemId), ownKeys)}
              flags={memberFlags}
              pendingKey={council.flagPending}
              {build}
              shipsIcon={council.selected.kind === 'gear'}
              {timeZone}
              notice={council.notice}
              onFlag={(key, msOs) => void council.flag(key, msOs)}
            />
          {/if}
          <DecisionRecord
            decisions={night.decisions}
            {build}
            {shipsIcon}
            {timeZone}
            exportFor={officer ? guildId : null}
          />
        </div>
      {/snippet}
    </Page>
  {/if}
  <BisTooltipHost
    {build}
    modelsElementId="guild-loot-tooltip-models"
    resolve={(id) => fetchItemTooltipModel(build, classOf(id), id)}
  />
</div>
