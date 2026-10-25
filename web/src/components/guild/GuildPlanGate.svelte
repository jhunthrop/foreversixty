<!-- web/src/components/guild/GuildPlanGate.svelte -->
<!-- Officer tools gate spec §3.2: the one sentence a tab shows in officer-read-only, above its
     content. Its id is what each disabled Award, Undo and Nudge names in aria-describedby. The
     link is the inline-link recipe (DESIGN-SYSTEM.md, Accent); py-3 below md gives it a 44px
     hit area on a phone without growing the 20px line box. -->
<script lang="ts">
  import { guildGateCopy } from '../../lib/guild/copy';

  let { tab, settingsHref }: { tab: 'loot' | 'readiness'; settingsHref: string } = $props();
</script>

<p
  id={`officer-tools-gate-${tab}`}
  class="text-text text-[13px] leading-5"
  data-testid={`guild-${tab}-plan-gate`}
>
  {guildGateCopy.sentence}<a
    class="text-gold hover:text-gold-hover active:text-gold-hover focus-visible:outline-gold py-3 underline focus-visible:outline-2 focus-visible:outline-offset-2 md:py-0"
    href={settingsHref}>{guildGateCopy.link}</a
  >.
</p>
