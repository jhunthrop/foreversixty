<!-- web/src/components/guild/loot/DecisionRow.svelte -->
<!-- One award's decision record (web layout spec 2026-11-05 §7.2): the item to whom, then main or off
     spec, the officers' votes when it was awarded (officers only), by whom and when, and the reason
     on its own line. An officer gets Undo, which takes the award back and reopens the drop. -->
<script lang="ts">
  import type { LootDecision } from '../../../lib/guild/loot-council-api';
  import { decisionSub, decisionTitle } from '../../../lib/guild/loot-council-view';
  import Button from '../../ui/Button.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Row from '../../ui/Row.svelte';

  let {
    decision,
    build,
    shipsIcon,
    timeZone,
    onUndo,
    undoDisabled = false,
  }: {
    decision: LootDecision;
    build: string;
    /** The build ships the item's icon (a gear drop); else the empty socket, never a 404. */
    shipsIcon: boolean;
    timeZone?: string;
    /** Officers: take the award back. */
    onUndo?: () => void;
    undoDisabled?: boolean;
  } = $props();
</script>

<Row title={decisionTitle(decision)} open testid={`guild-loot-decision-${decision.award_id}`}>
  {#snippet lead()}<ItemIcon
      icon={shipsIcon && decision.icon !== '' ? decision.icon : undefined}
      quality={decision.quality}
      {build}
      size={32}
    />{/snippet}
  {#snippet subContent()}
    <span class="block">{decisionSub(decision, timeZone)}</span>
    {#if decision.reason !== null}<span class="text-text block" data-testid="guild-loot-decision-reason"
        >{decision.reason}</span
      >{/if}
  {/snippet}
  {#snippet trail()}
    {#if onUndo}<Button
        tone="ghost"
        size="sm"
        disabled={undoDisabled}
        onclick={onUndo}
        testid="guild-loot-undo">Undo</Button
      >{/if}
  {/snippet}
</Row>
