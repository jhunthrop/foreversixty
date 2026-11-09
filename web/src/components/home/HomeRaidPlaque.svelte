<!-- web/src/components/home/HomeRaidPlaque.svelte -->
<!-- The raid plaque (round 5, the page's one ornate panel on the cave art): the guild's next raid,
     its day and time, the clock in days and hours, signed up and late, the role counts, your place in the guild (item level among
     your spec, raids attended), I'm in /
     Late / Can't come (the saved choice pressed), and Before the pull at the right. A choice saves
     at once through the raid's sign-up endpoint; with no role to sign as, the raid page instead. -->
<script lang="ts">
  import type { GuildRaidRole, GuildSignupState, NextRaidView } from '../../lib/guild/api';
  import { putSignup } from '../../lib/guild/raids-api';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import { homeBCopy } from '../../lib/home/home-b-view';
  import {
    PLAQUE_CHOICES,
    choiceOf,
    plaqueClock,
    plaqueClockLabel,
    raidWhen,
    plaqueCopy,
    roleCounts,
    signupFigures,
    type PlaqueChoice,
  } from '../../lib/home/plaque';
  import { LINK } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import BigFigures from '../ui/BigFigures.svelte';
  import Button from '../ui/Button.svelte';
  import Frame from '../ui/Frame.svelte';
  import type { HomeFact } from '../../lib/home/home-facts';
  import HomeFacts from './HomeFacts.svelte';
  import HomePlaqueTitle from './HomePlaqueTitle.svelte';

  let {
    guildId,
    guild,
    raid,
    needs,
    standing,
    characterKey,
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
    /** The role a sign-up carries; null, the raid page signs up instead. */
    role: GuildRaidRole | null;
    raidHref: string | null;
  } = $props();

  const clock = minuteClock(() => undefined);
  const viewerTz = Intl.DateTimeFormat().resolvedOptions().timeZone;

  /** What a choice saved on this page, over the raid as it loaded. */
  let savedHere = $state<{ state: GuildSignupState; counts: NextRaidView['counts'] } | null>(null);
  let busy = $state<PlaqueChoice | null>(null);
  let failed = $state(false);

  const counts = $derived(savedHere?.counts ?? raid.counts);
  const chosen = $derived(choiceOf(savedHere?.state ?? raid.own?.state ?? null));
  const when = $derived(raidWhen(raid.starts_at, viewerTz));
  const parts = $derived(plaqueClock(new Date(raid.starts_at), clock.now));
  const timer = $derived(parts.length === 0 ? undefined : plaqueClockLabel(parts));
  const locked = $derived(raid.status === 'locked');

  async function choose(choice: PlaqueChoice): Promise<void> {
    if (role === null || busy !== null) return;
    busy = choice;
    failed = false;
    try {
      const result = await putSignup(guildId, raid.id, characterKey, {
        state: choice,
        role,
        note: raid.own?.note ?? '',
      });
      savedHere = { state: choice, counts: result.raid.counts };
    } catch {
      failed = true;
    } finally {
      busy = null;
    }
  }
</script>

<!-- Utilities, not a scoped <style>: this island's scoped CSS never reached the page (it is
     only reachable through HomeSignedIn's client graph); the theme utilities are global. -->
<Frame ornate art="cave" region={REGION.countdowns} testid="home-raid">
  <div class="grid gap-x-[34px] gap-y-4 min-[900px]:grid-cols-2">
    <div class="min-w-0">
      <HomePlaqueTitle title={raid.zone_name} sub={`${guild} · ${when}`} />
      <div class="mt-4">
        <BigFigures groups={[parts, signupFigures(counts)]} {timer} testid="home-raid-figures" />
      </div>
      <p class="text-text mt-[14px] flex flex-wrap gap-x-[18px] gap-y-[6px] text-[14px] leading-5">
        {#each roleCounts(counts) as row (row.word)}<span><b class="text-strong">{row.n}</b> {row.word}</span
          >{/each}
      </p>
      {#if standing.length > 0}
        <div class="mt-[6px]"><HomeFacts facts={standing} testid="home-raid-standing" /></div>
      {/if}
      <div class="mt-[18px] flex flex-wrap items-center gap-[10px]">
        {#if role !== null && !locked}
          {#each PLAQUE_CHOICES as choice (choice)}
            <Button
              tone={choice === 'in' ? 'primary' : 'secondary'}
              size="sm"
              pressed={chosen === choice}
              disabled={busy !== null}
              onclick={() => choose(choice)}
              testid={`home-raid-${choice}`}>{plaqueCopy.choices[choice]}</Button
            >
          {/each}
        {:else if raidHref !== null}
          <a class={LINK} href={raidHref}>{homeBCopy.signUpLink} ›</a>
        {/if}
      </div>
      {#if failed}<p class="text-wipe mt-2 text-[14px]" role="alert">{plaqueCopy.signUpFailed}</p>{/if}
    </div>
    {#if needs.length > 0}
      <div
        class="border-frame-line border-t pt-[14px] min-[900px]:border-t-0 min-[900px]:border-l min-[900px]:pt-0 min-[900px]:pl-7"
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
