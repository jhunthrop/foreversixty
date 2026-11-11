package gg.foreversixty.app.push

import android.app.NotificationManager
import android.content.Context
import android.content.res.Resources
import gg.foreversixty.app.R

/**
 * The drawer's Notifications row (round 5 a1): the groups a player can get that are on, in their
 * order, sentence case (`Raids, gear and sims, logs on`), else `All off`. On Android each group is
 * its own channel ([PushGroup]), so the state is the channels' (and the app's) own switches.
 */
object NotificationsLine {
    /** The groups the row names: Account is always on, Officer only for an officer. */
    fun shown(officer: Boolean): List<PushGroup> =
        PushGroup.entries.filter { it != PushGroup.Account && (officer || it != PushGroup.Officer) }

    /** The line for the groups [on]; the golden and the drawer both build it here. */
    fun of(res: Resources, on: List<PushGroup>): String {
        val first = on.firstOrNull() ?: return res.getString(R.string.menu_notifications_off)
        val rest = on.drop(1).map { res.getString(it.title).lowercase() }
        val groups = (listOf(res.getString(first.title)) + rest).joinToString(LIST)
        return res.getString(R.string.menu_notifications_on, groups)
    }

    /**
     * Which of [shown]'s groups are on now: the app's notifications allowed and the group's channel not
     * off (the channels exist from launch, [PushGroup.createChannels]; one not made yet is on by default).
     */
    fun on(context: Context, officer: Boolean): List<PushGroup> {
        val manager = context.getSystemService(NotificationManager::class.java) ?: return emptyList()
        if (!manager.areNotificationsEnabled()) return emptyList()
        return shown(officer).filter { group ->
            manager.getNotificationChannel(group.channelId)?.importance != NotificationManager.IMPORTANCE_NONE
        }
    }

    private const val LIST = ", "
}
