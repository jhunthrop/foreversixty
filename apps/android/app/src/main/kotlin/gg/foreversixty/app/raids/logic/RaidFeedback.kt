package gg.foreversixty.app.raids.logic

import android.Manifest
import android.content.Context
import android.content.SharedPreferences
import android.content.pm.PackageManager
import android.os.Build
import android.view.HapticFeedbackConstants
import android.view.View
import androidx.core.content.ContextCompat
import androidx.core.content.edit
import gg.foreversixty.app.model.ApiJson
import java.io.File
import java.time.Duration
import java.time.Instant
import kotlinx.serialization.KSerializer
import kotlinx.serialization.SerializationException
import kotlinx.serialization.json.buildJsonObject
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import kotlinx.serialization.json.long
import kotlinx.serialization.json.put

/** The sign-up's haptics (spec 7): a tick on a press, confirm on a save, reject on a refusal. */
interface RaidHaptics {
    fun selection()
    fun success()
    fun error()
}

/** The system haptics on the view the press came from; none below API 30 for confirm and reject. */
class ViewRaidHaptics(private val view: View) : RaidHaptics {
    override fun selection() {
        val tick = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.UPSIDE_DOWN_CAKE) {
            HapticFeedbackConstants.SEGMENT_TICK
        } else {
            HapticFeedbackConstants.CLOCK_TICK
        }
        view.performHapticFeedback(tick)
    }

    override fun success() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) view.performHapticFeedback(HapticFeedbackConstants.CONFIRM)
    }

    override fun error() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) view.performHapticFeedback(HapticFeedbackConstants.REJECT)
    }
}

object NoRaidHaptics : RaidHaptics {
    override fun selection() = Unit
    override fun success() = Unit
    override fun error() = Unit
}

/** Whether POST_NOTIFICATIONS is still undecided (Android 13+); 9 to 12 need no ask (spec F4.4). */
fun interface NotificationPermission {
    fun isUndetermined(): Boolean
}

class SystemNotificationPermission(private val context: Context) : NotificationPermission {
    override fun isUndetermined(): Boolean =
        Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU &&
            ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED
}

/** The reminder ask's rule (spec F4.4): at most twice per device, the second at least 30 days later. */
data class ReminderAskPolicy(val shown: Int, val lastShown: Instant?) {
    fun mayAsk(now: Instant): Boolean = shown < MAX_ASKS && (lastShown == null || Duration.between(lastShown, now) >= GAP)

    companion object {
        const val MAX_ASKS = 2
        val GAP: Duration = Duration.ofDays(30)
    }
}

interface ReminderAskStore {
    fun policy(): ReminderAskPolicy
    fun recordShown(at: Instant)
}

class PrefsReminderAskStore(private val prefs: SharedPreferences) : ReminderAskStore {
    override fun policy() = ReminderAskPolicy(
        prefs.getInt(COUNT, 0),
        prefs.getLong(LAST, 0L).takeIf { it > 0L }?.let(Instant::ofEpochMilli),
    )

    override fun recordShown(at: Instant) = prefs.edit {
        putInt(COUNT, prefs.getInt(COUNT, 0) + 1)
        putLong(LAST, at.toEpochMilli())
    }

    private companion object {
        const val COUNT = "raids.reminderAsk.count"
        const val LAST = "raids.reminderAsk.last"
    }
}

/**
 * The last good answer per endpoint, on disk (spec 5.0 Caching): a private file per key. A cached
 * answer draws at once and sits under the offline strip when a refresh cannot connect.
 */
interface RaidCache {
    fun <T> read(key: String, serializer: KSerializer<T>): Cached<T>?
    fun <T> write(key: String, serializer: KSerializer<T>, value: T, at: Instant)

    /** Sign-out: nothing of the last account's guild stays on the device. */
    fun removeAll()
}

data class Cached<T>(val value: T, val savedAt: Instant)

class FileRaidCache(private val directory: File) : RaidCache {
    private fun file(key: String) = File(directory, key.replace('/', '_') + ".json")

    override fun removeAll() {
        directory.deleteRecursively()
    }

    override fun <T> read(key: String, serializer: KSerializer<T>): Cached<T>? = try {
        val entry = ApiJson.parseToJsonElement(file(key).readText()).jsonObject
        val savedAt = entry.getValue(SAVED_AT).jsonPrimitive.long
        Cached(ApiJson.decodeFromJsonElement(serializer, entry.getValue(VALUE)), Instant.ofEpochMilli(savedAt))
    } catch (_: java.io.IOException) {
        null
    } catch (_: SerializationException) {
        null
    } catch (_: IllegalArgumentException) {
        null
    } catch (_: NoSuchElementException) {
        null
    }

    override fun <T> write(key: String, serializer: KSerializer<T>, value: T, at: Instant) {
        try {
            directory.mkdirs()
            val entry = buildJsonObject {
                put(SAVED_AT, at.toEpochMilli())
                put(VALUE, ApiJson.encodeToJsonElement(serializer, value))
            }
            file(key).writeText(entry.toString())
        } catch (_: java.io.IOException) {
            // A cache that cannot write only costs the next cold start its instant draw.
        }
    }

    private companion object {
        const val SAVED_AT = "saved_at"
        const val VALUE = "value"
    }
}

object NoRaidCache : RaidCache {
    override fun <T> read(key: String, serializer: KSerializer<T>): Cached<T>? = null
    override fun <T> write(key: String, serializer: KSerializer<T>, value: T, at: Instant) = Unit
    override fun removeAll() = Unit
}
