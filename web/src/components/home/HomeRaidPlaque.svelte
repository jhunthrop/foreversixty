<!-- web/src/components/home/HomeRaidPlaque.svelte -->
<!-- The raid plaque (round 5, the page's one ornate panel on the cave art): the guild's next raid,
     its day and time, the clock in days and hours, signed up and late, the role counts, your place in the guild (item level among
     your spec, raids attended), the guild's own sign-up control (In, Late, Bench, Out), and Before
     the pull at the right. A choice saves at once through the one sign-up save
     (lib/guild/signup-save.ts); with no role to sign as, the raid page instead. -->
<script lang="ts">
  import { raidCopy } from '../../lib/guild/raid-copy';
  import type { GuildRaidRole, GuildSignupState, NextRaidView } from '../../lib/guild/api';
  import { pressSignup } from '../../lib/guild/signup-save';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import { homeBCopy } from '../../lib/home/home-b-view';
  import { countdownFigures, raidWhen, plaqueCopy, roleCounts, signupFigures } from '../../lib/home/plaque';
  import { LINK } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import BigFigures from '../ui/BigFigures.svelte';
  import Frame from '../ui/Frame.svelte';
  import type { HomeFact } from '../../lib/home/home-facts';
  import HomeFacts from './HomeFacts.svelte';
  import HomePlaqueTitle from './HomePlaqueTitle.svelte';
  import SignupSegments from '../guild/SignupSegments.svelte';
  import { viewerZone } from '../../lib/dates';
  import { countdownText } from '../../lib/ui/countdown';

  let {
    guildId,
    guild,
    raid,
    needs,
    standing,
    characterKey,
    characterName,
    role,
    raidHref,
  }: {
    guildId: number;
    guild: string;
    raid: NextRaidView;
    /** Before the pull: the character's own list from the guild's standing. */
    needs: readonly string[];
    /** You in the guild: your item level place among your spec, raids attended. */
    standing: readonly HomeFact[];
    characterKey: string;
    /** The sign-up control's accessible name: `{name}'s sign-up`. */
    characterName: string;
    /** The role a sign-up carries; null, the raid page signs up instead. */
    role: GuildRaidRole | null;
    raidHref: string | null;
  } = $props();

  const clock = minuteClock(() => undefined);
  const viewerTz = viewerZone();

  /** The raid after a choice saved on this page, over the raid as it loaded. */
  let savedHere = $state<NextRaidView | null>(null);
  let pending = $state<GuildSignupState | null>(null);
  let failedLine = $state('');

  const shown = $derived(savedHere ?? raid);
  const counts = $derived(shown.counts);
  const chosen = $derived(shown.own?.state ?? null);
  const when = $derived(raidWhen(raid.starts_at, viewerTz));
  const parts = $derived(countdownFigures(new Date(raid.starts_at), clock.now));
  const timer = $derived(parts.length === 0 ? undefined : countdownText(new Date(raid.starts_at), clock.now));
  const locked = $derived(shown.status === 'locked');

  async function choose(state: GuildSignupState): Promise<void> {
    if (role === null || pending !== null) return;
    pending = state;
    failedLine = '';
    const outcome = await pressSignup(guildId, shown, characterKey, {
      state,
      role,
      note: shown.own?.note ?? '',
    });
    pending = null;
    if (outcome.next) savedHere = outcome.next;
    if (!outcome.ok) failedLine = outcome.line;
  }
</script>

<!-- Utilities, not a scoped <style>: this island's scoped CSS never reached the page (it is
     only reachable through HomeSignedIn's client graph); the theme utilities are global. -->
<Frame ornate art="cave" region={REGION.countdowns} testid="home-raid">
  <div class="grid gap-x-[34px] gap-y-4 min-[900px]:grid-cols-[minmax(0,1fr)_auto]">
    <div class="min-w-0">
      <HomePlaqueTitle title={raid.zone_name} sub={`${guild} · ${when}`} />
      <div class="mt-4">
        <BigFigures groups={[parts]} {timer} testid="home-raid-figures" />
      </div>
      <p class="text-text mt-[14px] flex flex-wrap gap-x-[18px] gap-y-[6px] text-[14px] leading-5">
        {#each signupFigures(counts) as row (row.label)}<span data-testid="home-raid-signups"
            ><b class="text-strong">{row.value}</b> {row.label}</span
          >{/each}{#each roleCounts(counts) as row (row.word)}<span
            ><b class="text-strong">{row.n}</b> {row.word}</span
          >{/each}
      </p>
      {#if standing.length > 0}
        <div class="mt-[6px]"><HomeFacts facts={standing} testid="home-raid-standing" /></div>
      {/if}
      <div class="mt-[18px] flex flex-wrap items-center gap-[10px]">
        {#if role !== null && !locked}
          <SignupSegments
            value={chosen}
            {pending}
            name={characterName}
            size="strip"
            disabled={pending !== null}
            onChange={(state) => void choose(state)}
            testid="home-raid-signup"
          />
          <span class="text-muted text-[14px]" data-testid="home-raid-role"
            >{raidCopy.signingAs(raidCopy.roleLabels[role])}</span
          >
        {:else if raidHref !== null}
          <a class={LINK} href={raidHref}>{homeBCopy.signUpLink} ›</a>
        {/if}
      </div>
      {#if failedLine !== ''}<p class="text-wipe mt-2 text-[14px]" role="alert">{failedLine}</p>{/if}
    </div>
    {#if needs.length > 0}
      <div
        class="border-frame-line border-t pt-[14px] min-[900px]:max-w-[360px] min-[900px]:self-start min-[900px]:border-t-0 min-[900px]:border-l min-[900px]:pt-0 min-[900px]:pl-7"
        data-testid="home-raid-prep"
      >
        <h3 class="font-heading text-cream text-[19px] leading-[26px] font-normal">
          {plaqueCopy.beforePull}
        </h3>
        <ul class="mt-2">
          {#each needs as need (need)}
            <li
              class="text-text before:border-frame-line flex items-baseline gap-[10px] py-[3px] text-[15px] leading-[22px] before:size-[11px] before:flex-none before:translate-y-px before:border before:content-['']"
            >
              {need}
            </li>
          {/each}
        </ul>
      </div>
    {/if}
  </div>
</Frame>
