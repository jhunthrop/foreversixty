package gg.foreversixty.app.widgets.data

import android.content.Context
import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.booleanPreferencesKey
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import androidx.datastore.preferences.preferencesDataStore
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.widgets.logic.SignupOverlay
import gg.foreversixty.app.widgets.logic.WidgetSnapshot
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.flow.map
import kotlinx.serialization.KSerializer
import kotlinx.serialization.SerializationException
import kotlinx.serialization.builtins.MapSerializer
import kotlinx.serialization.builtins.serializer

private val Context.widgetData: DataStore<Preferences> by preferencesDataStore(name = "forever_widgets")

/**
 * The shared container the Glance widgets, the tile and the ongoing notifications read (spec §5.6
 * Refresh): one snapshot per character a widget is set to ("" is the session's current character),
 * the pending or refused widget press, each widget's character setting, the Live surfaces switches
 * and the live registrations' ids. Nothing here is a secret; it is cleared at sign-out.
 */
class WidgetStore(private val data: DataStore<Preferences>) {
    constructor(context: Context) : this(context.applicationContext.widgetData)

    suspend fun snapshot(character: String = CURRENT): WidgetSnapshot? = read(snapshotKey(character), WidgetSnapshot.serializer())

    suspend fun saveSnapshot(character: String, snapshot: WidgetSnapshot) = write(snapshotKey(character), WidgetSnapshot.serializer(), snapshot)

    suspend fun overlay(): SignupOverlay? = read(OVERLAY, SignupOverlay.serializer())

    /** The snapshot and the pending press as they change: a placed widget redraws from these. */
    fun watch(character: String): Flow<Pair<WidgetSnapshot?, SignupOverlay?>> = data.data.map { prefs ->
        decode(prefs[snapshotKey(character)], WidgetSnapshot.serializer()) to decode(prefs[OVERLAY], SignupOverlay.serializer())
    }

    suspend fun setOverlay(overlay: SignupOverlay?) = write(OVERLAY, SignupOverlay.serializer(), overlay)

    /** The character a widget shows; [CURRENT] (the session's) unless its configuration picked one. */
    suspend fun characterFor(appWidgetId: Int): String = data.data.first()[widgetKey(appWidgetId)] ?: CURRENT

    suspend fun setCharacter(appWidgetId: Int, character: String) {
        data.edit { it[widgetKey(appWidgetId)] = character }
    }

    /** Every character some widget is set to, plus the current one. */
    suspend fun characters(): Set<String> =
        data.data.first().asMap().filterKeys { it.name.startsWith(WIDGET_PREFIX) }.values.filterIsInstance<String>().toSet() + CURRENT

    suspend fun liveSetting(kind: LiveKind): Boolean = data.data.first()[settingKey(kind)] ?: true

    suspend fun setLiveSetting(kind: LiveKind, on: Boolean) {
        data.edit { it[settingKey(kind)] = on }
    }

    /** The server's id for the live registration of an object (`raid:7`), to DELETE when it ends. */
    suspend fun registrations(): Map<String, Long> = read(REGISTRATIONS, REGISTRATIONS_SERIALIZER).orEmpty()

    suspend fun setRegistrations(map: Map<String, Long>) = write(REGISTRATIONS, REGISTRATIONS_SERIALIZER, map)

    /** Sign-out: everything goes; the widgets then say `Sign in to see your next raid.` */
    suspend fun clear() {
        data.edit { it.clear() }
    }

    private suspend fun <T> read(key: Preferences.Key<String>, serializer: KSerializer<T>): T? = decode(data.data.first()[key], serializer)

    private fun <T> decode(raw: String?, serializer: KSerializer<T>): T? {
        if (raw == null) return null
        return try {
            ApiJson.decodeFromString(serializer, raw)
        } catch (_: SerializationException) {
            null
        } catch (_: IllegalArgumentException) {
            null
        }
    }

    private suspend fun <T> write(key: Preferences.Key<String>, serializer: KSerializer<T>, value: T?) {
        data.edit { prefs ->
            if (value == null) prefs.remove(key) else prefs[key] = ApiJson.encodeToString(serializer, value)
        }
    }

    companion object {
        const val CURRENT = ""
        private const val WIDGET_PREFIX = "widget_character_"
        private val OVERLAY = stringPreferencesKey("signup_overlay")
        private val REGISTRATIONS = stringPreferencesKey("live_registrations")
        private val REGISTRATIONS_SERIALIZER = MapSerializer(String.serializer(), Long.serializer())

        private fun snapshotKey(character: String) = stringPreferencesKey("snapshot_$character")
        private fun widgetKey(appWidgetId: Int) = stringPreferencesKey("$WIDGET_PREFIX$appWidgetId")
        private fun settingKey(kind: LiveKind) = booleanPreferencesKey("live_setting_${kind.wire}")
    }
}

/** The two live surfaces (W3-C13 `kind`), each with its own You › Live surfaces switch. */
enum class LiveKind(val wire: String) {
    Raid("raid"),
    Sim("sim"),
}
