<!-- web/src/components/account/AccountGuildPlanLine.svelte -->
<!-- The plan row of a guild on /account (premium spec 2026-11-08 section 7.4): `Guild settings` as a
     slim row whose trail links to them with what the guild has (the plan words, the date, the seat
     note). -->
<script lang="ts">
  import type { GuildBillingView } from '../../lib/account/api';
  import { guildRowPlanLine } from '../../lib/account/plan-line';
  import { guildSettingsHref } from '../../lib/characters';
  import { settingsWords } from '../../lib/guild/settings-view';
  import { LINK } from '../../lib/ui/primitives';
  import Row from '../ui/Row.svelte';

  let {
    plan,
    region,
    ruleset,
    name,
  }: { plan: GuildBillingView; region: string; ruleset: string; name: string } = $props();
</script>

<Row title={settingsWords.guildSettingsLink} density="slim" open wrapTrail>
  {#snippet trail()}
    <a
      class="{LINK} text-[15px]"
      href={guildSettingsHref(region, ruleset, name)}
      data-testid="account-guild-plan">{guildRowPlanLine(plan)} ›</a
    >
  {/snippet}
</Row>
