package gg.foreversixty.app.firstrun

import android.content.SharedPreferences
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.core.content.edit
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.raids.logic.NotificationPermission

/**
 * The first run after sign-in (flows spec F8, iOS `FirstRunModel`): Pick your main, then, for a guild
 * member, the notification ask at its moment of value (E6). Never at launch, never twice.
 */
class FirstRunModel(private val deps: Dependencies) {
    sealed interface Step {
        data object PickMain : Step

        data class Permission(val guild: String) : Step
    }

    class Dependencies(
        /** `PATCH /v1/me` with the main's key; false when it did not save. */
        val setMain: suspend (String) -> Boolean,
        val permission: NotificationPermission,
        val asked: FirstRunAskStore,
    )

    var step: Step? by mutableStateOf(null)
        private set

    /** The row the button continues with; the highest level until a tap. */
    var selection: String? by mutableStateOf(null)

    var saving: Boolean by mutableStateOf(false)
        private set

    var failed: Boolean by mutableStateOf(false)
        private set

    /** The account loaded or changed: asks for a main when it has none. */
    fun update(me: Me?) {
        me ?: return
        val first = candidates(me).firstOrNull()
        if (me.mainCharacterKey == null && first != null) {
            if (step != Step.PickMain) selection = first.key
            step = Step.PickMain
        } else if (step == Step.PickMain) {
            step = null
        }
    }

    suspend fun continueWithMain(me: Me?) {
        val key = selection ?: return
        val character = me?.characters?.firstOrNull { it.key == key } ?: return
        saving = true
        failed = false
        val saved = deps.setMain(key)
        saving = false
        if (!saved) {
            failed = true
            return
        }
        step = permissionStep(character)
    }

    /** The ask follows the main into a guild, once per install, while the system has not decided. */
    private fun permissionStep(character: AccountCharacter): Step? {
        val guild = character.guild?.name ?: return null
        if (deps.asked.wasAsked() || !deps.permission.isUndetermined()) return null
        return Step.Permission(guild)
    }

    /** `Turn on notifications`: [request] shows the system's Allow. */
    fun turnOnNotifications(request: () -> Unit) {
        deps.asked.markAsked()
        step = null
        request()
    }

    fun notNow() {
        deps.asked.markAsked()
        step = null
    }

    companion object {
        /** The characters Pick your main offers: level 10 and up, the highest first. */
        const val MINIMUM_LEVEL = 10

        fun candidates(me: Me): List<AccountCharacter> =
            me.characters.filter { (it.level ?: 0) >= MINIMUM_LEVEL }.sortedByDescending { it.level ?: 0 }
    }
}

/** Whether the first-run ask was shown on this install (gone with the app's data). */
interface FirstRunAskStore {
    fun wasAsked(): Boolean
    fun markAsked()
}

class PrefsFirstRunAskStore(private val prefs: SharedPreferences) : FirstRunAskStore {
    override fun wasAsked(): Boolean = prefs.getBoolean(ASKED, false)

    override fun markAsked() = prefs.edit { putBoolean(ASKED, true) }

    companion object {
        const val FILE = "forever_first_run"
        private const val ASKED = "permission_asked"
    }
}
