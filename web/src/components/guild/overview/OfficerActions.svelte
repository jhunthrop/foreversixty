<!-- web/src/components/guild/overview/OfficerActions.svelte -->
<!-- Action needed (batch 3 mock, officers): the page's framed panel, one Row per thing waiting, each
     with its one secondary button. "Ping on Discord" posts one message under the raid's post naming
     who has not signed (owner decision D16; never a DM); its answer replaces the row's sub line. -->
<script lang="ts">
  import { GuildApiError } from '../../../lib/guild/api';
  import { guildApplicationsHref } from '../../../lib/characters';
  import { pingRaid, type GuildOfficerActions } from '../../../lib/guild/overview-api';
  import { actionRows, pingLine } from '../../../lib/guild/overview-role-view';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import type { GuildTabId } from '../GuildTabs.svelte';

  let {
    guildId,
    actions,
    day,
    canPing,
    onSelectTab,
    guildPath,
  }: {
    guildId: number;
    /** The guild page's path: Open goes to its applications. */
    guildPath: string;
    actions: GuildOfficerActions;
    day: string | null;
    /** The raid has a Discord post and the throttle allows a ping. */
    canPing: boolean;
    onSelectTab: (tab: GuildTabId) => void;
  } = $props();

  const rows = $derived(actionRows(actions, day));
  let pinging = $state(false);
  let pingAnswer = $state<string | null>(null);

  async function ping(): Promise<void> {
    const raidId = actions.not_signed.raid_id;
    if (raidId === null || pinging) return;
    pinging = true;
    try {
      pingAnswer = pingLine(await pingRaid(guildId, raidId));
    } catch (error) {
      pingAnswer = error instanceof GuildApiError ? error.message : 'That did not go through. Try again.';
    } finally {
      pinging = false;
    }
  }
</script>

{#if rows.length > 0}
  <Section title="Action needed" panel frame="framed" testid="guild-overview-actions">
    {#each rows as row (row.key)}
      <Row
        title={row.title}
        sub={row.key === 'not_signed' && pingAnswer !== null ? pingAnswer : row.sub}
        open
      >
        {#snippet trail()}
          {#if row.key === 'approvals'}
            <Button tone="secondary" size="sm" onclick={() => onSelectTab('roster')}>Review</Button>
          {:else if row.key === 'applications'}
            <Button
              tone="secondary"
              size="sm"
              href={guildApplicationsHref(guildPath)}
              testid="guild-open-applications">Open</Button
            >
          {:else if row.key === 'not_signed'}
            <Button
              tone="secondary"
              size="sm"
              disabled={!canPing || pinging}
              onclick={() => void ping()}
              testid="guild-ping-discord">Ping on Discord</Button
            >
          {:else}
            <Button tone="secondary" size="sm" onclick={() => onSelectTab('loot')}>Open loot</Button>
          {/if}
        {/snippet}
      </Row>
    {/each}
  </Section>
{/if}
