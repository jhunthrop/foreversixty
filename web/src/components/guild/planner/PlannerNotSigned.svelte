<!-- web/src/components/guild/planner/PlannerNotSigned.svelte -->
<!-- Officers: the raiders with no sign-up yet, one per account (its main), and Ping on Discord: one
     message under the raid's post naming them (owner decision D16; never a DM). The answer, or why
     the ping cannot go out yet, is the line under the head. An officer may answer for a raider here
     with the same in-row control the roster uses. -->
<script lang="ts">
  import type { GuildRaidView, GuildSignupState } from '../../../lib/guild/api';
  import { kindCountsRoles, kindOf } from '../../../lib/guild/raid-copy';
  import { notSignedRows, pingBlockedLine } from '../../../lib/guild/raid-planner-view';
  import type { RosterRow } from '../../../lib/guild/raid-roster';
  import type { PlannerFacts } from '../../../lib/guild/planner-facts';
  import Button from '../../ui/Button.svelte';
  import Section from '../../ui/Section.svelte';
  import PlannerRow from './PlannerRow.svelte';

  let {
    raid,
    now,
    clock,
    officerTools,
    closed = false,
    busyKey = null,
    pinging = false,
    pingAnswer = null,
    facts = null,
    onPing,
    onSetState,
  }: {
    /** Each raider's figures by character key; null draws none. */
    facts?: ReadonlyMap<string, PlannerFacts> | null;
    raid: GuildRaidView;
    now: Date;
    /** The next-ping time in words (`2:10 am`). */
    clock: (at: Date) => string;
    officerTools: boolean;
    closed?: boolean;
    busyKey?: string | null;
    pinging?: boolean;
    pingAnswer?: string | null;
    onPing: () => void;
    onSetState: (row: RosterRow, state: GuildSignupState) => Promise<boolean>;
  } = $props();

  const rows = $derived(notSignedRows(raid));
  const blocked = $derived(pingBlockedLine(raid, now, clock));
  const roles = $derived(kindCountsRoles(kindOf(raid)));
  const line = $derived(pingAnswer ?? blocked);
</script>

{#if rows.length > 0}
  <Section title="Not signed up" panel meta={String(rows.length)} testid="raid-not-signed">
    {#snippet head()}
      <Button
        size="sm"
        disabled={!officerTools || closed || blocked !== null || pinging}
        onclick={onPing}
        testid="raid-ping-discord">Ping on Discord</Button
      >
    {/snippet}
    {#if line !== null}
      <p class="text-muted mb-2 text-[13px]" role="status" data-testid="raid-ping-line">{line}</p>
    {/if}
    <ul>
      {#each rows as entry (entry.character_key)}
        <PlannerRow
          row={entry.roster}
          sub={entry.sub}
          {roles}
          officer
          disabled={closed}
          busy={busyKey === entry.character_key}
          facts={facts?.get(entry.character_key) ?? null}
          {onSetState}
        />
      {/each}
    </ul>
  </Section>
{/if}
