<!-- web/src/components/account/AccountGuildPlanLine.svelte -->
<!-- The plan line on an account page guild row: what the guild has, then a link to its settings. -->
<script lang="ts">
  import type { GuildBillingView } from '../../lib/account/api';
  import { guildSettingsHref } from '../../lib/characters';
  import { accountGuildPlanLine } from '../../lib/guild/settings-view';

  let {
    plan,
    region,
    ruleset,
    name,
  }: { plan: GuildBillingView; region: string; ruleset: string; name: string } = $props();
</script>

<p class="text-muted basis-full text-[13px] leading-5" data-testid="account-guild-plan">
  {accountGuildPlanLine(plan)} ·
  <a class="hover:text-strong active:text-strong underline" href={guildSettingsHref(region, ruleset, name)}
    >Guild settings</a
  >
</p>
