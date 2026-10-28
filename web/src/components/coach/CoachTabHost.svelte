<!-- web/src/components/coach/CoachTabHost.svelte -->
<!-- ReportView's one mount for the Coach tab: the whole-night line, or the lazily imported
     CoachTab behind a skeleton that reserves the tab's height (640 px desktop, 420 px phone).
     The panel, review body and teaser ship in their own chunk and load only when `tab=coach`. -->
<script lang="ts">
  import { coachCopy } from '../../lib/coach/copy';
  import type { RosterUnit } from '../../lib/coach/player-pick';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import { REPORT_LAZY_MIN_H } from '../../lib/report/layout';
  import type { FightEntry } from '../../lib/report/types';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    reportId,
    fight,
    firstFight,
    roster,
    source,
    nightMode,
    onSelectPlayer,
  }: {
    reportId: string;
    fight: FightEntry | null;
    firstFight: number;
    roster: RosterUnit[];
    source: string;
    nightMode: boolean;
    onSelectPlayer: (guid: string) => void;
  } = $props();

  const lazy = createLazyComponent(() => import('./CoachTab.svelte'));
  $effect(() => {
    if (!nightMode) lazy.load();
  });
</script>

{#if nightMode}
  <p class="text-muted text-[14px]" data-testid="coach-tab-night">{coachCopy.night}</p>
{:else if lazy.current}
  <lazy.current {reportId} {fight} {firstFight} {roster} {source} {onSelectPlayer} />
{:else if lazy.error !== ''}
  <LoadError message={lazy.error} onRetry={() => lazy.load()} testid="coach-tab-error" />
{:else}
  <Skeleton minHeight={REPORT_LAZY_MIN_H.coach} testid="coach-tab-skeleton" />
{/if}
