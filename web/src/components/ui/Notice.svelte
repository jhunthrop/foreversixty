<!-- web/src/components/ui/Notice.svelte -->
<!-- One standing sentence under the page header (web layout spec 2026-11-05 §3A.2): a guild's
     standing line, a "before raids open" state. Text and one optional link, no box. `lead` puts a
     small mark before the sentence (the sim's Worth knowing lines lead with the spell's icon). -->
<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  import type { Snippet } from 'svelte';
  import { LINK } from '../../lib/ui/primitives';

  let {
    text,
    action,
    tone = 'gold',
    tail,
    lead,
    short,
    testid,
  }: {
    text: string;
    /** The same note in one phone line (below 768), the long text from 768 (the planner's arrival note). */
    short?: string;
    action?: { label: string; href: string };
    tone?: 'gold' | 'neutral';
    /** Words that finish the sentence after the text (a signed-out "Sign in to ..." clause). */
    tail?: Snippet;
    /** A mark before the text (a 20 px spell icon). */
    lead?: Snippet;
    testid?: string;
  } = $props();
</script>

<p
  class={[
    'ui-notice flex flex-wrap items-baseline gap-x-[14px] gap-y-[6px] pt-[14px] pb-1 text-[15px] leading-6',
    tone === 'gold' ? 'text-gold' : 'text-text',
  ]}
  data-testid={testid}
>
  {#if lead}<span class="inline-flex shrink-0 self-center">{@render lead()}</span>{/if}
  <span
    >{#if tail}{`${text} `}{@render tail()}{:else if short}<span class="ui-notice-long">{text}</span><span
        class="ui-notice-short">{short}</span
      >{:else}{text}{/if}</span
  >
  {#if action}<a class={LINK} href={action.href}>{action.label} ›</a>{/if}
</p>

<style>
  .ui-notice-short {
    display: none;
  }
  @media (max-width: 767px) {
    .ui-notice-long {
      display: none;
    }
    .ui-notice-short {
      display: inline;
    }
  }
</style>
