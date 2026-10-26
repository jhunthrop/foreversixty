<!-- web/src/components/guild/GuildPlanGate.svelte -->
<!-- Officer tools gate spec §3.2: the one sentence a tab shows in officer-read-only, above its
     content. Its id is what each disabled Award, Undo and Nudge names in aria-describedby. The
     link is the inline-link recipe (DESIGN-SYSTEM.md, Accent). It never wraps; below md a 44px
     ::after centred on the link (an inline box is only its font's height, about 16px, so a
     fixed height, not 12px insets, is what reaches 44px) is its hit area on a phone,
     while the focus ring stays on the link and hugs the text. A tab with its own sentence (the
     Raids tab, raids spec 6.1) passes the words before the link as `sentence`. -->
<script lang="ts">
  import { guildGateCopy } from '../../lib/guild/copy';

  let {
    tab,
    settingsHref,
    sentence = guildGateCopy.sentence,
    withId = true,
  }: {
    tab: 'loot' | 'readiness' | 'raids' | 'discord';
    settingsHref: string;
    sentence?: string;
    /** False for a second copy of the sentence on one page (the id is the first copy's). */
    withId?: boolean;
  } = $props();
</script>

<p
  id={withId ? `officer-tools-gate-${tab}` : undefined}
  class="text-text text-[13px] leading-5"
  data-testid={withId ? `guild-${tab}-plan-gate` : `guild-${tab}-plan-gate-repeat`}
>
  {sentence}<a
    class="text-gold hover:text-gold-hover active:text-gold-hover focus-visible:outline-gold relative whitespace-nowrap underline after:absolute after:inset-x-0 after:top-1/2 after:h-11 after:-translate-y-1/2 after:content-[''] focus-visible:outline-2 focus-visible:outline-offset-2 md:after:hidden"
    href={settingsHref}>{guildGateCopy.link}</a
  >.
</p>
