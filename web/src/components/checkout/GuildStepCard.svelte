<!-- web/src/components/checkout/GuildStepCard.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.2: one guild's Row on the checkout guild step (premium
     spec 2026-11-08 section 6: a Row of the `Which guild?` Plate). The lead is the ringed HD crest,
     the body the guild's name, realm, the viewer's own characters in it and the guild's facts; the
     trail (320 px from 768, under the body on a phone) is the step block, passed in as `block`. A card with no standing
     (an older /v1/me) draws the identity row and the block only. Not a control: no hover or
     focus state of its own. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import Icon from '../ui/Icon.svelte';
  import { ICON_CHECK } from '../../lib/ui/ui-icons';
  import type { Snippet } from 'svelte';
  import type { MeCharacter, MeGuild } from '../../lib/account/me-types';
  import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
  import { guildFacts, youRow } from '../../lib/billing/guild-card-model';
  import { guildMarkSrc } from '../../lib/guild/mark';
  import Row from '../ui/Row.svelte';

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

<Row
  title={guild.name}
  item
  id={`guild-${guild.id}`}
  labelledby={nameId}
  wrapTrail
  testid="checkout-guild-card"
>
  {#snippet lead()}
    <span
      class="border-gold flex h-10 w-10 shrink-0 items-center justify-center overflow-hidden rounded-full border-2"
      aria-hidden="true"
      data-testid="guild-crest"
    >
      {#if markSrc !== null}
        <img src={markSrc} alt="" width="40" height="40" class="h-full w-full object-cover" />
      {/if}
    </span>
  {/snippet}
  {#snippet titleContent()}
    <h2 id={nameId} tabindex="-1" class="m-0 truncate text-[16px] leading-5 font-bold outline-none">
      {guild.name}
    </h2>
  {/snippet}
  {#snippet subContent()}
    {#if realmLine !== null}<span class="block">{realmLine}</span>{/if}
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
                <CharacterName name={entry.name} cls={entry.characterClass} crestSize={20} />
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
                    <Icon icon={ICON_CHECK} size={12} />
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
      <span class="mt-[6px] block" data-testid="guild-facts">
        {#each facts as fact, index (index)}
          {#if index > 0}<span class="text-muted"
              ><span aria-hidden="true">{FACT_SEPARATOR}</span><span class="sr-only">,</span></span
            >{/if}{#if fact.figure !== null}<span class="numeric">{fact.figure}</span>{/if}{fact.rest}
        {/each}
      </span>
    {/if}
  {/snippet}
  {#snippet trail()}
    <div class="flex w-full min-w-0 flex-col gap-[10px] md:w-[320px]" data-testid="guild-step-block">
      {@render block()}
    </div>
  {/snippet}
</Row>
