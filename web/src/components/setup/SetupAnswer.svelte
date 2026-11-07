<!-- web/src/components/setup/SetupAnswer.svelte -->
<!-- /setup's head answer once the account has a character (setup spec U2): the header selector's
     current character as the 40 px ringed crest (one crest language), the name in its class colour
     over "Level 34 Beast Mastery Hunter", when Battle.net last brought it in, and the two places to
     take it next. Right-aligned from 900, stacked under the meta below. Never switches the
     character. -->
<script lang="ts">
  import type { SetupAnswer } from '../../lib/setup/answer';
  import { classColorVar } from '../../lib/report/format';
  import { LINK } from '../../lib/ui/primitives';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  const CREST_PX = 40;

  let { answer }: { answer: SetupAnswer } = $props();
</script>

<span class="setup-answer-who">
  <ClassCrestRing characterClass={answer.classSlug ?? ''} size={CREST_PX} />
  <span class="min-w-0">
    <b
      class="block text-[18px] leading-[22px] font-bold"
      style:color={classColorVar(answer.classSlug ?? undefined)}>{answer.name}</b
    >
    <span class="text-small text-muted block">{answer.levelLine}</span>
  </span>
</span>
{#if answer.imported !== null}<span class="text-small text-muted">{answer.imported}</span>{/if}
<span class="setup-answer-links">
  <a class={LINK} href={answer.open.href} data-testid="setup-answer-open">{answer.open.label} ›</a>
  {#if answer.bis !== null}<a class={LINK} href={answer.bis.href} data-testid="setup-answer-bis"
      >{answer.bis.label} ›</a
    >{/if}
</span>

<style>
  .setup-answer-who {
    display: flex;
    align-items: center;
    gap: 12px;
  }
  .setup-answer-links {
    display: flex;
    flex-wrap: wrap;
    gap: 16px;
  }
  /* On a touch screen each link's hit area is 44 px tall: 12 px of padding above and below, taken
     back by the same negative margin, so the answer does not grow (as Section's action link). */
  @media (pointer: coarse) {
    .setup-answer-links a {
      display: inline-block;
      margin-block: -12px;
      padding-block: 12px;
    }
  }
  @media (min-width: 900px) {
    /* The answer column holds 300 px from 900, as the mock's. */
    :global([data-region='head-answer']:has(> .setup-answer-who)) {
      min-width: 300px;
    }
    .setup-answer-who {
      justify-content: flex-end;
      text-align: right;
    }
    .setup-answer-links {
      justify-content: flex-end;
    }
  }
</style>
