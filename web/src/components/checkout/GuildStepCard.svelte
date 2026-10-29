<!-- web/src/components/checkout/GuildStepCard.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.2: one guild's card on the checkout guild step. The
     left column is the guild's identity (the ringed HD crest, name, realm), the viewer's own
     characters in it, and the guild's facts; the right column (320 px from 768, below the facts
     after a rule on a phone) is the step block, passed in as `block`. A card with no standing
     (an older /v1/me) draws the identity row and the block only. Not a control: no hover or
     focus state of its own. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { MeCharacter, MeGuild } from '../../lib/account/me-types';
  import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
  import { guildFacts, youRow } from '../../lib/billing/guild-card-model';
  import { guildMarkSrc } from '../../lib/guild/mark';
  import { CARD } from '../../lib/ui/card';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  const FACT_SEPARATOR = ' · ';

  let {
    guild,
    characters,
    mainKey = undefined,
    standing,
    block,
  }: {
    guild: MeGuild;
    characters: readonly MeCharacter[];
    mainKey?: string;
    /** False for an older /v1/me: identity only, no you row or facts. */
    standing: boolean;
    block: Snippet;
  } = $props();

  const you = $derived(youRow(characters, guild.id, mainKey));
  const facts = $derived(guildFacts(guild));
  const realmLine = $derived(
    guild.ruleset_label ? `${guild.ruleset_label} · ${guild.region.toUpperCase()}` : null,
  );
  const markSrc = $derived(guildMarkSrc(guild));
  const nameId = $derived(`guild-${guild.id}-name`);
</script>

<li
  id={`guild-${guild.id}`}
  aria-labelledby={nameId}
  class={`${CARD} grid scroll-mt-24 grid-cols-1 items-start gap-[14px] p-4 md:grid-cols-[minmax(0,1fr)_320px] md:gap-6 md:p-5`}
  data-testid="checkout-guild-card"
>
  <div class="flex min-w-0 flex-col">
    <div class="flex min-w-0 items-center gap-3">
      <span
        class="border-gold flex h-10 w-10 shrink-0 items-center justify-center overflow-hidden rounded-full border-2"
        aria-hidden="true"
        data-testid="guild-crest"
      >
        {#if markSrc !== null}
          <img src={markSrc} alt="" width="40" height="40" class="h-full w-full object-cover" />
        {/if}
      </span>
      <span class="flex min-w-0 flex-col">
        <h2 id={nameId} tabindex="-1" class="heading-name text-strong m-0 truncate text-[16px] outline-none">
          {guild.name}
        </h2>
        {#if realmLine !== null}
          <span class="text-muted text-[13px]">{realmLine}</span>
        {/if}
      </span>
    </div>

    {#if standing}
      {#if you.entries.length > 0}
        <ul
          class="m-0 mt-2 flex list-none flex-wrap items-center gap-2 p-0 text-[14px]"
          data-testid="guild-you"
        >
          {#each you.entries as entry (entry.key)}
            <!-- Row 1b at 360: the name and the rank word never break inside; on a narrow phone the
                 verify mark drops to its own line rather than splitting words. -->
            <li class="flex flex-wrap items-center gap-x-2 gap-y-1">
              <span class="inline-flex items-center gap-2 whitespace-nowrap">
                <ClassCrestRing characterClass={entry.characterClass} size={20} />
                <span class="font-semibold" style={`color: ${classColorVar(entry.characterClass)}`}>
                  {entry.name}
                </span>
              </span>
              <span class="inline-flex items-center gap-2 whitespace-nowrap">
                <span class="text-muted"><span aria-hidden="true">·</span><span class="sr-only">,</span></span
                >
                <span class="text-text">{entry.rankWord}</span>
              </span>
              <span class="inline-flex items-center gap-2 whitespace-nowrap">
                <span class="text-muted"><span aria-hidden="true">·</span><span class="sr-only">,</span></span
                >
                {#if entry.verified}
                  <span class="text-kill inline-flex items-center gap-1">
                    <svg viewBox="0 0 14 14" width="12" height="12" aria-hidden="true">
                      <path
                        d="M2.5 7.5l3 3 6-7"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                    </svg>
                    {copy.verifyMark.verified}
                  </span>
                {:else}
                  <span class="text-muted">{copy.verifyMark.notVerified}</span>
                {/if}
              </span>
            </li>
          {/each}
          {#if you.more !== null}
            <li class="text-muted">{you.more}</li>
          {/if}
        </ul>
      {/if}
      <p class="text-muted m-0 mt-[6px] text-[13px]" data-testid="guild-facts">
        {#each facts as fact, index (index)}
          {#if index > 0}<span class="text-muted"
              ><span aria-hidden="true">{FACT_SEPARATOR}</span><span class="sr-only">,</span></span
            >{/if}{#if fact.figure !== null}<span class="numeric">{fact.figure}</span>{/if}{fact.rest}
        {/each}
      </p>
    {/if}
  </div>

  <div
    class="border-line-soft flex min-w-0 flex-col gap-[10px] border-t pt-[14px] md:border-t-0 md:pt-0"
    data-testid="guild-step-block"
  >
    {@render block()}
  </div>
</li>
