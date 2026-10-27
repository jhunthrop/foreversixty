<!-- web/src/components/checkout/CheckoutGuildChooser.svelte -->
<!-- Spec 3.1.7: a guild plan with no guild id asks "Which guild?" -- one 56px row per guild
     the account may subscribe. `guilds` null means loading; `failed` shows the load error.
     `crests` maps guild id to its mark (guildMarkSrc); the ring shows it, or stays a plain
     gold ring while a guild's mark is unknown. The h2 takes focus when the chooser shows. -->
<script lang="ts">
  import type { MeGuild } from '../../lib/account/api';
  import { billableGuilds, guildFitsRaidTier } from '../../lib/billing/checkout-params';
  import { checkoutCopy } from '../../lib/billing/copy';
  import Skeleton from '../ui/Skeleton.svelte';
  import LoadError from '../ui/LoadError.svelte';

  const SKELETON_ROWS = 3;

  let headingEl = $state<HTMLHeadingElement | undefined>(undefined);

  let {
    guilds,
    failed = false,
    crests = {},
    onChoose,
    onRetry,
  }: {
    guilds: readonly MeGuild[] | null;
    failed?: boolean;
    crests?: Readonly<Record<number, string | null>>;
    onChoose: (guildId: number) => void;
    onRetry: () => void;
  } = $props();

  $effect(() => headingEl?.focus());

  const rows = $derived(guilds === null ? [] : billableGuilds(guilds));

  function fitLabel(verifiedCount: number): string {
    return guildFitsRaidTier(verifiedCount)
      ? checkoutCopy.fitsRaidGuild(verifiedCount)
      : checkoutCopy.guildPlanLabel;
  }
</script>

<section class="flex flex-col gap-3" aria-labelledby="checkout-which-guild" data-testid="checkout-chooser">
  <h2
    id="checkout-which-guild"
    bind:this={headingEl}
    tabindex="-1"
    class="section-title text-[15px] outline-none"
  >
    {checkoutCopy.whichGuild}
  </h2>
  {#if failed}
    <LoadError message={checkoutCopy.guildsDidNotLoad} {onRetry} testid="checkout-chooser-error" />
  {:else if guilds === null}
    <Skeleton lines={SKELETON_ROWS} rowHeight="h-14" testid="checkout-chooser-loading" />
  {:else if rows.length === 0}
    <p class="text-muted text-[14px]" data-testid="checkout-chooser-none">
      {checkoutCopy.noBillableGuild}
      <a href="/guild" class="text-gold underline">{checkoutCopy.findYourGuild}</a>
    </p>
  {:else}
    <ul class="flex flex-col gap-2">
      {#each rows as guild (guild.id)}
        <li>
          <button
            type="button"
            class="border-line-warm hover:bg-card-top focus-visible:outline-gold rounded-control flex h-14 w-full items-center gap-3 border px-3 text-left focus-visible:outline-2 focus-visible:-outline-offset-2"
            data-testid="checkout-guild-row"
            onclick={() => onChoose(guild.id)}
          >
            <span
              class="border-gold h-8 w-8 shrink-0 overflow-hidden rounded-full border-2"
              aria-hidden="true"
            >
              {#if crests[guild.id]}
                <img src={crests[guild.id]} alt="" class="h-full w-full object-cover" />
              {/if}
            </span>
            <span class="flex min-w-0 flex-1 flex-col">
              <span class="text-strong truncate text-[15px]">{guild.name}</span>
              <span class="text-muted text-[13px]">{checkoutCopy.verifiedMembers(guild.verified_count)}</span>
            </span>
            <span class="text-muted shrink-0 text-[13px]">{fitLabel(guild.verified_count)}</span>
          </button>
        </li>
      {/each}
    </ul>
  {/if}
</section>
