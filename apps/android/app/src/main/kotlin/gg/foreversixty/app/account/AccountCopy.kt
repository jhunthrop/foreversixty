package gg.foreversixty.app.account

import androidx.compose.runtime.Composable
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.DeletionConsequence
import gg.foreversixty.app.auth.session.DeviceRow
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.auth.session.SignInFailure
import gg.foreversixty.app.auth.session.SignedOutNotice
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Billing
import gg.foreversixty.app.phrase
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale

/** The words for F1's states, all from the copy catalog. Dates: `30 Oct`, `27 Oct 2027`. */
object AccountCopy {
    @Composable
    fun notice(notice: SignedOutNotice): String = when (notice) {
        is SignedOutNotice.SignInFailed -> signInFailure(notice.failure)
        SignedOutNotice.SessionEnded -> stringResource(R.string.signin_signedOut)
        is SignedOutNotice.AccountDeleted -> stringResource(R.string.account_delete_done, longDate(notice.completeBy))
    }

    /** Red is for failures only (spec §5.0): a completed deletion reads neutral. */
    fun isFailure(notice: SignedOutNotice): Boolean = when (notice) {
        is SignedOutNotice.SignInFailed, SignedOutNotice.SessionEnded -> true
        is SignedOutNotice.AccountDeleted -> false
    }

    /** `US · PvP`: Forever has no realms, so a character is its region, ruleset and name. */
    @Composable
    fun place(character: AccountCharacter): String =
        stringResource(R.string.account_character_where, character.region.uppercase(), ruleset(character.ruleset))

    /** The site's ruleset labels (web/src/lib/characters.ts); an unknown one is echoed. */
    @Composable
    fun ruleset(id: String): String = when (id) {
        "normal" -> stringResource(R.string.account_ruleset_normal)
        "pvp" -> stringResource(R.string.account_ruleset_pvp)
        "rp" -> stringResource(R.string.account_ruleset_rp)
        "hardcore" -> stringResource(R.string.account_ruleset_hardcore)
        else -> id
    }

    @Composable
    fun consequence(consequence: DeletionConsequence): String = when (consequence) {
        DeletionConsequence.PremiumPlan -> stringResource(R.string.account_delete_lose_premium)
        DeletionConsequence.FounderPrice -> stringResource(R.string.account_delete_lose_founder)
        is DeletionConsequence.GuildPlan -> stringResource(R.string.account_delete_lose_guildPlan, consequence.guild)
        is DeletionConsequence.Seat -> stringResource(R.string.account_delete_lose_seat, consequence.guild)
    }

    /** `Version 0.1.0 (1) · data 1.60.1.70658`, without the data part before the config answers. */
    @Composable
    fun version(version: String, build: Int, data: String?): String = if (data == null) {
        stringResource(R.string.account_version_noData, version, build.toString())
    } else {
        stringResource(R.string.you_version, version, build.toString(), data)
    }

    @Composable
    fun signInFailure(failure: SignInFailure): String = stringResource(
        when (failure) {
            SignInFailure.BattleNet -> R.string.signin_error_bnet
            SignInFailure.Expired -> R.string.signin_error_expired
            SignInFailure.Offline -> R.string.signin_error_offline
        },
    )

    @Composable
    fun age(age: RelativeAge): String = age.phrase(LocalResources.current)

    @Composable
    fun deviceName(row: DeviceRow): String = if (row.isThisDevice) stringResource(R.string.devices_this) else row.name

    @Composable
    fun deviceMeta(row: DeviceRow): String? = when (val meta = row.meta) {
        is DeviceRow.Meta.SignedIn -> stringResource(R.string.devices_signedIn, shortDate(meta.at))
        is DeviceRow.Meta.LastUsed -> stringResource(R.string.devices_lastUsed, age(meta.age))
        DeviceRow.Meta.None -> null
    }

    /** `Orc Fury Warrior · Level 58` (switcher rows). */
    @Composable
    fun switcherDetail(character: AccountCharacter): String {
        val kind = listOfNotNull(character.race, character.spec, character.className).joinToString(" ") { words(it) }
        val level = character.level ?: return kind
        return stringResource(R.string.account_switcher_detail, kind, level)
    }

    /** `Fury · 58` (the selector's short line); the class when no spec is saved. */
    @Composable
    fun selectorDetail(character: AccountCharacter): String? {
        val level = character.level ?: return null
        val spec = (character.spec ?: character.className)?.let(::words).orEmpty()
        return stringResource(R.string.account_selector_detail, spec, level)
    }

    @Composable
    fun selectorLabel(character: AccountCharacter): String = stringResource(
        R.string.selector_a11y,
        character.name,
        character.spec?.let(::words).orEmpty(),
        character.className?.let(::words).orEmpty(),
        character.level ?: 0,
    )

    /** The site's `accountPlanCopy` lines: `Premium · renews 27 Oct 2027`, or `ends` once cancelled. */
    @Composable
    fun planLine(billing: Billing?): String {
        val end = billing?.currentPeriodEnd?.let { runCatching { Instant.parse(it) }.getOrNull() }
        return when {
            billing == null || end == null -> stringResource(R.string.you_plan_free)
            billing.cancelAtPeriodEnd -> stringResource(R.string.plan_ends, longDate(end))
            else -> stringResource(R.string.plan_renews, longDate(end))
        }
    }

    /** The API's lower-case words (`night elf`, `warrior`) as display words. */
    fun words(raw: String): String = raw.split(' ').joinToString(" ") { it.replaceFirstChar(Char::uppercaseChar) }

    fun shortDate(at: Instant): String = SHORT.format(at)

    fun longDate(at: Instant): String = LONG.format(at)

    /** The catalog is English (en.json); dates match it. */
    private val SHORT = DateTimeFormatter.ofPattern("d MMM", Locale.UK).withZone(ZoneId.systemDefault())
    private val LONG = DateTimeFormatter.ofPattern("d MMM yyyy", Locale.UK).withZone(ZoneId.systemDefault())
}
