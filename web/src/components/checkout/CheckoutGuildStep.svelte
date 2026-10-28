<!-- web/src/components/checkout/CheckoutGuildStep.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.1 body and §5.7: the guild step under the band. The
     centred 880 column (gutters 48 / 18, 32 / 20 px top), the lead sentence, one card per guild
     in orderGuildCards order, and the foot. `me` null is loading: the lead slot and Skeleton
     cards hold the ready height, their count from the session's fs.meGuildCount hint. `failed`
     swaps the body for LoadError; after a retry succeeds, focus moves to the lead. Empty A (no
     characters) and Empty B (no guild) draw the set-up panel, the page's one fill.

     The card body is Task 16's: each card here draws only the identity rows and `Open the
     guild`, which is also exactly what an older /v1/me (no `subscribe`) gets (§8). Task 16
     replaces the `<li>` below with GuildStepCard + GuildStepBlock; Task 18's `ready` card calls
     `onSwitch` for `Switch to Guild` / `Switch to Raid guild`. -->
<script lang="ts">
  import type { Me, MeGuild } from '../../lib/account/api';
  import { guildHref } from '../../lib/characters';
  import { premiumSource } from '../../lib/account/plan-line';
  import { readMeGuildCount } from '../../lib/account/premium-hint';
  import { checkoutGuildCopy as copy, type GuildTier } from '../../lib/billing/checkout-guild-copy';
  import { GUILD_LEAD_RESERVE, GUILD_SKELETON_CARD } from '../../lib/billing/checkout-layout';
  import { priceFor } from '../../lib/billing/plan-costs';
  import { formatPrice, type Interval } from '../../lib/billing/plans';
  import { leadSentence, orderGuildCards, type StepGuild } from '../../lib/billing/subscribe-step';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { checkoutHref } from '../premium/model';
  import LoadError from '../ui/LoadError.svelte';

  let {
    me,
    failed,
    interval,
    founder,
    selfUrl,
    onRetry,
  }: {
    /** Null while /v1/me is in flight. */
    me: Me | null;
    failed: boolean;
    plan: GuildTier;
    interval: Interval;
    founder: boolean;
    /** This checkout's own path and query: the `next` of every Get set up link. */
    selfUrl: string;
    onRetry: () => void;
    /** Rewrites the page's plan in place (spec §5.3); Task 18's `ready` card calls it. */
    onSwitch: (to: GuildTier) => void;
  } = $props();

  const MIN_SKELETON_CARDS = 1;
  /** The design system's card recipe (--raised, 1 px --line, 6 px radius, gold top highlight). */
  const CARD = 'bg-raised border-line rounded-panel border shadow-[inset_0_1px_0_rgba(229,185,85,.35)]';
  const TEXT_LINK =
    'text-text hover:text-strong focus-visible:outline-gold inline-flex min-h-11 items-center underline focus-visible:outline-2 focus-visible:outline-offset-2';

  let leadEl = $state<HTMLParagraphElement | undefined>(undefined);
  /** Plain, not reactive: set while the error shows, so the next ready render moves focus once. */
  let retrying = false;

  // The step rules key on the guild page slug; an older API sends none, and only the
  // hand-off link (Task 16) reads it, never the lead or the order.
  const guilds = $derived<(MeGuild & StepGuild)[]>(
    (me?.guilds ?? []).map((guild) => ({ ...guild, slug: guild.slug ?? '' })),
  );
  const cards = $derived(orderGuildCards(guilds));
  const lead = $derived(leadSentence(guilds));
  const skeletonCards = $derived(Math.max(MIN_SKELETON_CARDS, readMeGuildCount()));
  const setupHref = $derived(`/setup?next=${encodeURIComponent(selfUrl)}`);
  const empty = $derived(
    me === null || me.guilds.length > 0 ? null : me.characters.length === 0 ? copy.emptyA : copy.emptyB,
  );
  const showsPremiumLine = $derived(me !== null && premiumSource(me, new Date()) === 'none');
  const premiumPrice = $derived(formatPrice(priceFor('premium', interval, founder), interval));

  $effect(() => {
    if (failed) {
      retrying = true;
      return;
    }
    if (!retrying || me === null || leadEl === undefined) return;
    retrying = false;
    leadEl.focus();
  });

  function realmLine(guild: MeGuild): string | null {
    return guild.ruleset_label ? `${guild.ruleset_label} · ${guild.region.toUpperCase()}` : null;
  }

  function guildPageHref(guild: MeGuild): string {
    if (guild.slug) return `/guild/${guild.region}/${guild.ruleset}/${encodeURIComponent(guild.slug)}`;
    return guildHref(guild.region, guild.ruleset, guild.name);
  }
</script>

<div class="px-[18px] pt-5 pb-16 md:px-12 md:pt-8" data-testid="checkout-guild-step">
  <div class="mx-auto flex w-full max-w-[880px] flex-col">
    {#if failed}
      <LoadError message={copy.loadError.title} {onRetry} testid="checkout-guild-error" />
    {:else if me === null}
      <p class={`m-0 mb-4 ${GUILD_LEAD_RESERVE}`} aria-hidden="true"></p>
      <ul class="m-0 flex list-none flex-col gap-3 p-0" aria-busy="true" data-testid="checkout-guild-loading">
        <li class="sr-only">{copy.loading}</li>
        {#each Array.from({ length: skeletonCards }, (_, index) => index) as index (index)}
          <li class={`skeleton-block rounded-panel ${GUILD_SKELETON_CARD}`} aria-hidden="true"></li>
        {/each}
      </ul>
    {:else if empty !== null}
      <div
        class={`${CARD} mx-auto flex w-full max-w-[640px] flex-col gap-2 p-5`}
        data-testid={empty === copy.emptyA ? 'checkout-guild-empty-a' : 'checkout-guild-empty-b'}
      >
        <p class="text-strong m-0 text-[15px]">{empty.title}</p>
        <p class="text-text m-0 text-[14px]">{empty.body}</p>
        <a href={setupHref} class={`${PRIMARY_BUTTON_FIXED} mt-2 w-fit px-4`}>{empty.button}</a>
      </div>
    {:else}
      <p
        bind:this={leadEl}
        tabindex="-1"
        class="text-text m-0 mb-4 max-w-[720px] text-[14px] outline-none md:text-[15px]"
        data-testid="checkout-guild-lead"
      >
        {lead}
      </p>
      <ul class="m-0 flex list-none flex-col gap-3 p-0" data-testid="checkout-guild-cards">
        {#each cards as guild (guild.id)}
          <li
            id={`guild-${guild.id}`}
            aria-labelledby={`guild-${guild.id}-name`}
            class={`${CARD} flex scroll-mt-24 flex-col gap-3 p-4 md:p-5`}
            data-testid="checkout-guild-card"
          >
            <div class="flex min-w-0 items-center gap-3">
              <span
                class="border-gold h-10 w-10 shrink-0 overflow-hidden rounded-full border-2"
                aria-hidden="true"
              >
                {#if guild.crest_url}
                  <img
                    src={guild.crest_url}
                    alt=""
                    width="40"
                    height="40"
                    class="h-full w-full object-cover"
                  />
                {/if}
              </span>
              <span class="flex min-w-0 flex-col">
                <h2
                  id={`guild-${guild.id}-name`}
                  class="font-display text-strong m-0 truncate text-[16px] font-bold"
                >
                  {guild.name}
                </h2>
                {#if realmLine(guild) !== null}
                  <span class="text-muted text-[13px]">{realmLine(guild)}</span>
                {/if}
              </span>
            </div>
            <a href={guildPageHref(guild)} class={`${TEXT_LINK} w-fit text-[14px]`}>{copy.openGuild}</a>
          </li>
        {/each}
      </ul>
    {/if}

    {#if me !== null && !failed && empty !== copy.emptyA}
      <div class="text-muted mt-6 flex flex-col gap-1 text-[13px]" data-testid="checkout-guild-foot">
        <p class="m-0">
          {copy.foot.missingGuild}
          <a href={setupHref} class={INLINE_LINK}>{copy.foot.getSetUp}</a>
        </p>
        {#if showsPremiumLine}
          <p class="m-0">
            {copy.foot.premium(premiumPrice)}
            <a href={checkoutHref('premium', interval)} class={INLINE_LINK}>{copy.foot.getPremium}</a>
          </p>
        {/if}
      </div>
    {/if}
  </div>
</div>
