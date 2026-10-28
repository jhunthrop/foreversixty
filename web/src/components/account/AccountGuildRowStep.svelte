<!-- web/src/components/account/AccountGuildRowStep.svelte -->
<!-- Spec 6.5: the two things a guild row adds. `mark` is the muted "Not verified" after the guild
     name; `line` is the no-plan sentence with its link to the checkout guild step. The words and
     the link come from lib/account/guild-row-line. -->
<script lang="ts">
  import { guildRowNextStep, notVerifiedMark, type GuildRowFacts } from '../../lib/account/guild-row-line';

  let { guild, part }: { guild: GuildRowFacts; part: 'mark' | 'line' } = $props();

  const mark = $derived(notVerifiedMark(guild));
  const next = $derived(guildRowNextStep(guild));
</script>

{#if part === 'mark' && mark !== null}
  <span class="text-muted text-[13px]" data-testid="account-guild-not-verified">{mark}</span>
{:else if part === 'line' && next !== null}
  <p class="text-muted basis-full text-[13px] leading-5" data-testid="account-guild-no-plan">
    {next.text}
    <a
      class="text-muted hover:text-strong active:text-strong underline"
      href={next.href}
      aria-label={next.label}>{next.action}</a
    >
  </p>
{/if}
