package gg.foreversixty.app.menu

import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Me
import java.time.Instant

/**
 * What the profile menu shows (spec 2026-11-08 section 3): You, the other characters to switch to
 * (level 10 and up, the rest behind `Show low-level characters`), highest level first.
 */
data class MenuModel(
    val you: AccountCharacter?,
    val isMain: Boolean,
    val updated: RelativeAge?,
    val switchTo: List<AccountCharacter>,
    val lowLevel: List<AccountCharacter>,
) {
    companion object {
        /** Characters under this level fold behind `Show low-level characters (N)`. */
        const val LOW_LEVEL = 10

        fun of(me: Me, current: AccountCharacter?, now: Instant): MenuModel {
            val others = me.characters.filter { it.key != current?.key }.sortedByDescending { it.level ?: 0 }
            val (low, high) = others.partition { (it.level ?: 0) < LOW_LEVEL }
            return MenuModel(
                you = current,
                isMain = current != null && current.key == me.mainCharacterKey,
                updated = me.bnetImportedAt?.let { runCatching { Instant.parse(it) }.getOrNull() }?.let { RelativeAge.of(it, now) },
                switchTo = high,
                lowLevel = low,
            )
        }
    }
}

/** The Battle.net stamp's Refresh: at rest, the call out, or the last one failed. */
enum class MenuRefresh {
    Idle,
    Refreshing,
    Failed,
    ;

    companion object {
        /** Runs [refresh] (true when it read the account) with the stamp showing each step through [show]. */
        suspend fun run(show: (MenuRefresh) -> Unit, refresh: suspend () -> Boolean) {
            show(Refreshing)
            show(if (refresh()) Idle else Failed)
        }
    }
}

/** The menu's rows that push a screen, under the tab the menu was opened on. */
enum class MenuDestination {
    Guides,
    Tiers,
    PatchNotes,
    Sims,
    Logs,
    Planner,
    Rankings,
    Notifications,
    Plan,
    Devices,
    OfflineData,
    Settings,
    ;

    companion object {
        /** The Tools rows a signed-out player sees (the menu-out mock). */
        val signedOutTools = listOf(Guides, Tiers, PatchNotes, Planner, Rankings)
        val tools = listOf(Guides, Tiers, PatchNotes, Sims, Logs, Planner, Rankings)
        val account = listOf(Notifications, Plan, Devices, OfflineData, Settings)
    }
}
