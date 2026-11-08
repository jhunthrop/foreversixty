package gg.foreversixty.app.account.notifications

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.NotificationChoice
import gg.foreversixty.app.model.NotificationSettings
import gg.foreversixty.app.model.NotificationsPatch
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.push.PushGroup
import java.time.LocalTime
import java.time.format.DateTimeFormatter
import java.util.Locale

/** `GET` and `PATCH /v1/me/notifications`. */
interface NotificationSettingsService {
    suspend fun load(): ApiResult<NotificationSettings>
    suspend fun save(patch: NotificationsPatch): ApiResult<NotificationSettings>
}

class LiveNotificationSettingsService(private val api: ApiClient) : NotificationSettingsService {
    override suspend fun load(): ApiResult<NotificationSettings> = api.send(ApiEndpoints.getNotificationSettings())

    override suspend fun save(patch: NotificationsPatch): ApiResult<NotificationSettings> =
        api.send(ApiEndpoints.updateNotificationSettings(patch))
}

/**
 * The settings as the screen shows them (spec 2026-11-08 F4, mock `notification-settings`): one
 * switch per group the player gets events of, in the groups' order (Raids and Guild ladder for a
 * guild member, Officer for an officer, as the API's event list says; Account has no switch), the
 * raid-day reminder and the officer DM opt-in.
 */
data class NotificationGroupsSummary(
    val rows: List<Row>,
    /** The raid-day reminder, `15:00` in the player's zone. */
    val reminder: String,
    val officerDms: Boolean,
) {
    data class Row(val group: PushGroup, val on: Boolean)

    companion object {
        const val DEFAULT_REMINDER = "15:00"

        /** The groups that carry a switch, in the settings' order. */
        val SWITCHED = PushGroup.entries - PushGroup.Account

        fun of(settings: NotificationSettings): NotificationGroupsSummary {
            val byGroup = settings.events.groupBy { event -> PushGroup.entries.firstOrNull { it.wire == event.group } }
            val rows = SWITCHED.mapNotNull { group ->
                byGroup[group]?.takeIf { it.isNotEmpty() }?.let { events -> Row(group, events.all { it.channels.push }) }
            }
            return NotificationGroupsSummary(rows, settings.raidReminder ?: DEFAULT_REMINDER, settings.officerDms ?: false)
        }

        /** The PATCH that turns a group's push on or off: every kind in it. */
        fun patch(group: PushGroup, on: Boolean, settings: NotificationSettings): NotificationsPatch =
            NotificationsPatch(
                events = settings.events.filter { it.group == group.wire }.map { NotificationChoice(it.kind, PUSH, on) },
            )

        /** The settings with a PATCH applied, drawn before the server answers. */
        fun applying(patch: NotificationsPatch, settings: NotificationSettings): NotificationSettings {
            val choices = patch.events.orEmpty().filter { it.channel == PUSH }.associate { it.kind to it.enabled }
            return settings.copy(
                raidReminder = patch.raidReminder ?: settings.raidReminder,
                officerDms = patch.officerDms ?: settings.officerDms,
                events = settings.events.map { event ->
                    choices[event.kind]?.let { event.copy(channels = event.channels.copy(push = it)) } ?: event
                },
            )
        }

        /** `15:00` as the player reads it: `3:00 pm`. */
        fun clock(hhmm: String): String = runCatching {
            CLOCK.format(LocalTime.parse(hhmm)).replace("AM", "am").replace("PM", "pm")
        }.getOrDefault(hhmm)

        /** The hours Change offers: noon to 8 pm, when a raid-day nudge is useful (as on iOS). */
        val HOURS: List<String> = (FIRST_HOUR..LAST_HOUR).map { "%02d:00".format(Locale.US, it) }

        private const val PUSH = "push"
        private const val FIRST_HOUR = 12
        private const val LAST_HOUR = 20
        private val CLOCK = DateTimeFormatter.ofPattern("h:mm a", Locale.US)
    }
}

/** The screen's state. */
sealed interface NotificationSettingsState {
    data object Loading : NotificationSettingsState
    data object Failed : NotificationSettingsState
    data class Loaded(val settings: NotificationSettings) : NotificationSettingsState
}

/**
 * Notifications (the menu's row): loads the choices and saves each change at once; a refused save
 * puts the old choice back and says so.
 */
class NotificationSettingsFeature(private val service: NotificationSettingsService) {
    var state: NotificationSettingsState by mutableStateOf(NotificationSettingsState.Loading)
        private set
    var saveFailed: Boolean by mutableStateOf(false)
        private set

    suspend fun load() {
        when (val result = service.load()) {
            is ApiResult.Success -> state = NotificationSettingsState.Loaded(result.data)
            is ApiResult.Failure -> if (state !is NotificationSettingsState.Loaded) state = NotificationSettingsState.Failed
        }
    }

    suspend fun setGroup(group: PushGroup, on: Boolean) = save { NotificationGroupsSummary.patch(group, on, it) }

    suspend fun setReminder(hhmm: String) = save { NotificationsPatch(raidReminder = hhmm) }

    suspend fun setOfficerDms(on: Boolean) = save { NotificationsPatch(officerDms = on) }

    private suspend fun save(patchOf: (NotificationSettings) -> NotificationsPatch) {
        val settings = (state as? NotificationSettingsState.Loaded)?.settings ?: return
        val patch = patchOf(settings)
        saveFailed = false
        state = NotificationSettingsState.Loaded(NotificationGroupsSummary.applying(patch, settings))
        when (val result = service.save(patch)) {
            is ApiResult.Success -> state = NotificationSettingsState.Loaded(result.data)
            is ApiResult.Failure -> {
                state = NotificationSettingsState.Loaded(settings)
                saveFailed = true
            }
        }
    }
}
