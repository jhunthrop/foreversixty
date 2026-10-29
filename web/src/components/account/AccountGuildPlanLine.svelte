<!-- web/src/components/account/AccountGuildPlanLine.svelte -->
<!-- The plan row on an account page guild block: `Guild settings` as a link row whose value is what
     the guild has (the plan words, the date, the seat note). -->
<script lang="ts">
  import type { GuildBillingView } from '../../lib/account/api';
  import { guildRowPlanLine } from '../../lib/account/plan-line';
  import { guildSettingsHref } from '../../lib/characters';
  import { settingsWords } from '../../lib/guild/settings-view';
  import SettingsRow from '../ui/SettingsRow.svelte';

  let {
    plan,
    region,
    ruleset,
    name,
  }: { plan: GuildBillingView; region: string; ruleset: string; name: string } = $props();
</script>

<SettingsRow
  glyph="guild"
  label={settingsWords.guildSettingsLink}
  value={guildRowPlanLine(plan)}
  href={guildSettingsHref(region, ruleset, name)}
  valueTestid="account-guild-plan"
/>
