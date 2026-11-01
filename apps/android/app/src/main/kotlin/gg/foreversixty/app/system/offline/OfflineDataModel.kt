package gg.foreversixty.app.system.offline

import android.content.SharedPreferences
import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.core.content.edit
import gg.foreversixty.app.R
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.raids.logic.GameClass
import java.io.File
import java.time.Instant
import java.time.ZoneId

/** The two Updates switches (both on by default) and when the data was last checked, kept on the phone. */
interface OfflineSettings {
    var keepClassesReady: Boolean
    var wifiOnly: Boolean
    var checkedAt: Instant?
}

class PrefsOfflineSettings(private val prefs: SharedPreferences) : OfflineSettings {
    override var keepClassesReady: Boolean
        get() = prefs.getBoolean(KEEP, true)
        set(value) = prefs.edit { putBoolean(KEEP, value) }

    override var wifiOnly: Boolean
        get() = prefs.getBoolean(WIFI, true)
        set(value) = prefs.edit { putBoolean(WIFI, value) }

    override var checkedAt: Instant?
        get() = prefs.getLong(CHECKED, 0L).takeIf { it > 0L }?.let(Instant::ofEpochMilli)
        set(value) = prefs.edit { if (value == null) remove(CHECKED) else putLong(CHECKED, value.toEpochMilli()) }

    companion object {
        const val FILE = "forever_offline"
        private const val KEEP = "offline.keepClassesReady"
        private const val WIFI = "offline.wifiOnly"
        private const val CHECKED = "offline.checkedAt"
    }
}

class MemoryOfflineSettings(
    override var keepClassesReady: Boolean = true,
    override var wifiOnly: Boolean = true,
    override var checkedAt: Instant? = null,
) : OfflineSettings

/** The item icon image cache: its size and its removal. Live is the Coil disk cache's folder. */
interface IconCache {
    fun size(): IconCacheSize
    fun clear()
}

class FolderIconCache(private val folder: File, private val onClear: () -> Unit = {}) : IconCache {
    override fun size(): IconCacheSize {
        val files = folder.walkTopDown().filter { it.isFile }.toList()
        return IconCacheSize(files.sumOf { it.length() }, files.maxOfOrNull { it.lastModified() }?.let(Instant::ofEpochMilli))
    }

    override fun clear() {
        onClear()
        folder.listFiles()?.forEach { it.deleteRecursively() }
    }
}

/** What `Update now` found. */
enum class UpdateOutcome { Updated, UpToDate, Failed }

/** What the screen says under `Update now` after it ran. */
enum class UpdateNote { UpToDate, Failed }

/**
 * You › Offline data (spec F15.4): the report, the two switches, `Update now` and `Remove
 * downloaded data`. [update] refreshes the active build and the saved files (the app graph wires the
 * guides' sync and the store's swap).
 */
class OfflineDataModel(
    private val statics: StaticDataStore,
    private val icons: IconCache,
    val settings: OfflineSettings,
    private val connectivity: Connectivity,
    private val update: suspend () -> UpdateOutcome,
    private val now: () -> Instant = Instant::now,
) {
    var report: OfflineReport? by mutableStateOf(null)
        private set
    var updating: Boolean by mutableStateOf(false)
        private set
    var note: UpdateNote? by mutableStateOf(null)
        private set
    var keepClassesReady: Boolean by mutableStateOf(settings.keepClassesReady)
        private set
    var wifiOnly: Boolean by mutableStateOf(settings.wifiOnly)
        private set
    var checkedAt: Instant? by mutableStateOf(settings.checkedAt)
        private set

    /** `Update now` needs a connection; offline it is disabled with its reason. */
    val online: Boolean get() = connectivity.isConnected()

    suspend fun load() {
        report = OfflineData.report(statics.files(), icons.size())
        checkedAt = settings.checkedAt
    }

    fun keepClassesReady(on: Boolean) {
        settings.keepClassesReady = on
        keepClassesReady = on
    }

    fun wifiOnly(on: Boolean) {
        settings.wifiOnly = on
        wifiOnly = on
    }

    suspend fun updateNow() {
        if (updating) return
        updating = true
        note = null
        val outcome = update()
        if (outcome != UpdateOutcome.Failed) {
            settings.checkedAt = now()
        }
        note = when (outcome) {
            UpdateOutcome.Updated -> null
            UpdateOutcome.UpToDate -> UpdateNote.UpToDate
            UpdateOutcome.Failed -> UpdateNote.Failed
        }
        updating = false
        load()
    }

    suspend fun remove() {
        statics.deleteAll()
        icons.clear()
        note = null
        load()
    }
}

/** The screen's words, from the copy catalog (`offlineData.*`). */
class OfflineDataCopy(private val res: Resources, private val zone: ZoneId = ZoneId.systemDefault()) {
    fun title(category: OfflineCategory): String = res.getString(
        when (category) {
            OfflineCategory.Guides -> R.string.offlineData_guides
            OfflineCategory.Bis -> R.string.offlineData_bis
            OfflineCategory.Talents -> R.string.offlineData_talents
            OfflineCategory.Search -> R.string.offlineData_search
            OfflineCategory.Icons -> R.string.offlineData_icons
        },
    )

    /** `Shaman, Warrior · 2.4 MB · Updated today`; icons show their size only. */
    fun meta(row: OfflineRow, now: Instant): String {
        val size = OfflineData.size(row.bytes)
        val updated = row.updated
        if (row.category == OfflineCategory.Icons || updated == null) return size
        val whenText = OfflineData.updated(updated, now, zone, res.getString(R.string.offlineData_today), res.getString(R.string.offlineData_yesterday))
        val classes = row.classes.map(::className)
        return if (classes.isEmpty()) {
            res.getString(R.string.offlineData_rowMetaNoWhat, size, whenText)
        } else {
            res.getString(R.string.offlineData_rowMeta, classes.joinToString(", "), size, whenText)
        }
    }

    fun section(report: OfflineReport): String = res.getString(R.string.offlineData_section, OfflineData.size(report.totalBytes))

    /** `Checked today at 9:41 pm. …`, or the never-checked footer. */
    fun footer(checkedAt: Instant?, now: Instant): String {
        checkedAt ?: return res.getString(R.string.offlineData_footerNever)
        val day = OfflineData.updated(checkedAt, now, zone, res.getString(R.string.offlineData_today), res.getString(R.string.offlineData_yesterday))
        val whenText = res.getString(R.string.offlineData_checkedAt, day, OfflineData.time(checkedAt, zone))
        return res.getString(R.string.offlineData_footer, whenText)
    }

    private fun className(slug: String): String = GameClass.of(slug)?.let { res.getString(it.label) } ?: slug
}
