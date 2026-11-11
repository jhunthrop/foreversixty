package gg.foreversixty.app.push

import android.app.NotificationManager
import android.content.Context
import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.NotificationSettings

/**
 * The drawer's Notifications row: `All on`, `All off`, else the groups that are off, in their order,
 * sentence case (`Gear and sims, guild ladder, tiers off`). A group is on only when the account's
 * choice for it (`GET /v1/me/notifications`, every event of the group on for push, as iOS reads it)
 * and its Android channel ([PushGroup]) are both on (round 7 row 4).
 */
object NotificationsLine {
    /** Without the account's settings: the groups a player can get. Account is always on, Officer only for an officer. */
    fun shown(officer: Boolean): List<PushGroup> =
        PushGroup.entries.filter { it != PushGroup.Account && (officer || it != PushGroup.Officer) }

    /**
     * The account's choice per group, the groups the player gets events of in their order: on when
     * every event of the group is on for push.
     */
    fun account(settings: NotificationSettings): Map<PushGroup, Boolean> {
        val byGroup = settings.events.groupBy { PushGroup.of(it.group, it.kind) }
        return PushGroup.entries.filter { it != PushGroup.Account }
            .mapNotNull { group -> byGroup[group]?.let { events -> group to events.all { it.channels.push } } }
            .toMap()
    }

    /**
     * The line for the [shown] groups of which [on] are on (round 6 row 1): `All on`, `All off`, else
     * what the player misses, `Logs, tiers off`. The golden and the drawer both build it here.
     */
    fun of(res: Resources, shown: List<PushGroup>, on: List<PushGroup>): String {
        val off = shown.filterNot { it in on }
        val first = off.firstOrNull() ?: return res.getString(R.string.menu_notifications_allOn)
        if (off.size == shown.size) return res.getString(R.string.menu_notifications_off)
        val rest = off.drop(1).map { res.getString(it.title).lowercase() }
        val groups = (listOf(res.getString(first.title)) + rest).joinToString(LIST)
        return res.getString(R.string.menu_notifications_someOff, groups)
    }

    /** The drawer's line: the account's groups when they loaded ([account]), else [shown] for [officer]. */
    fun line(res: Resources, context: Context, account: Map<PushGroup, Boolean>?, officer: Boolean): String {
        val shown = account?.keys?.toList() ?: shown(officer)
        return line(res, shown, account, channelsOn(context, shown))
    }

    /** The line for [shown] with the account's choices and the channels on: a group is on when both are. */
    fun line(res: Resources, shown: List<PushGroup>, account: Map<PushGroup, Boolean>?, channels: List<PushGroup>): String =
        of(res, shown, shown.filter { it in channels && account?.get(it) != false })

    /**
     * Which of [shown]'s channels are on now: the app's notifications allowed and the group's channel
     * not off (the channels exist from launch, [PushGroup.createChannels]; one not made yet is on).
     */
    fun channelsOn(context: Context, shown: List<PushGroup>): List<PushGroup> {
        val manager = context.getSystemService(NotificationManager::class.java) ?: return emptyList()
        if (!manager.areNotificationsEnabled()) return emptyList()
        return shown.filter { group ->
            manager.getNotificationChannel(group.channelId)?.importance != NotificationManager.IMPORTANCE_NONE
        }
    }

    private const val LIST = ", "
}
