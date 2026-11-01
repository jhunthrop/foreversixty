package gg.foreversixty.app.system.logic

import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Instant

/**
 * `Sign up In` from the launcher (spec F15.1): the current character's next raid (the Which raid
 * rule), answered In through the raids lane's own sign-up call ([RaidService.signup], the same PUT
 * the raid page, the notification and the widget make), keeping the character's note. It never
 * opens the app except for the spec ask (no spec saved) and two raids within 24 hours.
 */
class ShortcutSignup(
    private val raids: RaidService,
    private val signedIn: suspend () -> Boolean,
    /** The session's picked character key; null follows the account's main. */
    private val currentCharacter: () -> String?,
    private val isOn: () -> Boolean,
    private val now: () -> Instant = Instant::now,
) {
    suspend fun run(): SignupShortcutResult {
        if (!isOn()) return SignupShortcutResult.Off
        if (!signedIn()) return SignupShortcutResult.SignedOut
        val me = when (val answer = raids.me()) {
            is ApiResult.Success -> answer.data
            is ApiResult.Failure -> return failure(answer.error)
        }
        return signUp(me)
    }

    private suspend fun signUp(me: Me): SignupShortcutResult {
        val character = me.characters.firstOrNull { it.key == currentCharacter() } ?: me.mainCharacter
            ?: return SignupShortcutResult.NoCharacter
        val guild = character.guild?.takeIf { it.verified } ?: return SignupShortcutResult.NoGuild(character.name)
        val planned = when (val answer = raids.raids(guild.id)) {
            is ApiResult.Success -> answer.data.planned.orEmpty()
            is ApiResult.Failure -> return failure(answer.error)
        }
        return when (val pick = IntentRaid.pick(planned, now())) {
            RaidPick.None -> SignupShortcutResult.NoRaid
            is RaidPick.Which -> SignupShortcutResult.Which
            is RaidPick.One -> put(guild.id, pick.raid, character.key)
        }
    }

    private suspend fun put(guild: Long, raid: RaidView, characterKey: String): SignupShortcutResult {
        val viewer = raid.viewer.characters.firstOrNull { it.characterKey == characterKey && it.verified }
            ?: return SignupShortcutResult.Failed
        if (viewer.raidSpec == null && (viewer.specOptions?.size ?: 0) >= 2) return SignupShortcutResult.AskSpec(guild, raid.id)
        val note = raid.signups.firstOrNull { it.characterKey == characterKey }?.note.orEmpty()
        return when (val answer = raids.signup(guild, raid.id, characterKey, SignupRequest(state = SignupState.In.wire, note = note))) {
            is ApiResult.Success -> SignupShortcutResult.Signed(raid.zoneName)
            is ApiResult.Failure -> failure(answer.error)
        }
    }

    private fun failure(error: ApiError): SignupShortcutResult = when {
        error is ApiError.Network -> SignupShortcutResult.Offline
        error is ApiError.Server && error.isUnauthorized -> SignupShortcutResult.SignedOut
        error is ApiError.Server && error.code == LOCKED -> SignupShortcutResult.Locked
        else -> SignupShortcutResult.Failed
    }

    private companion object {
        const val LOCKED = "raid_locked"
    }
}
