package gg.foreversixty.app.search.logic

import android.content.SharedPreferences
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.core.content.edit
import gg.foreversixty.app.model.ApiJson
import kotlinx.serialization.SerializationException
import kotlinx.serialization.builtins.ListSerializer

/** Where Recent is kept: one JSON string on this phone. */
interface RecentStorage {
    fun read(): String?
    fun write(value: String?)
}

class SharedRecentStorage(private val prefs: SharedPreferences) : RecentStorage {
    override fun read(): String? = prefs.getString(KEY, null)
    override fun write(value: String?) = prefs.edit { putString(KEY, value) }

    companion object {
        const val FILE = "forever_search"
        private const val KEY = "recent"
    }
}

/** For tests and previews. */
class MemoryRecentStorage(private var value: String? = null) : RecentStorage {
    override fun read(): String? = value
    override fun write(value: String?) {
        this.value = value
    }
}

/**
 * Recent (spec F12.1): the last 20 things opened from Search on this phone, newest first. Opening
 * one again moves it to the top; sign out clears it.
 */
class RecentSearches(private val storage: RecentStorage) {
    var entries: List<SearchTarget> by mutableStateOf(load())
        private set

    fun add(target: SearchTarget) {
        save((listOf(target) + entries.filter { it.key != target.key }).take(CAP))
    }

    fun clear() = save(emptyList())

    private fun save(next: List<SearchTarget>) {
        entries = next
        storage.write(if (next.isEmpty()) null else ApiJson.encodeToString(SERIALIZER, next))
    }

    private fun load(): List<SearchTarget> {
        val json = storage.read() ?: return emptyList()
        return try {
            ApiJson.decodeFromString(SERIALIZER, json).take(CAP)
        } catch (e: SerializationException) {
            emptyList()
        } catch (e: IllegalArgumentException) {
            emptyList()
        }
    }

    companion object {
        const val CAP = 20
        private val SERIALIZER = ListSerializer(SearchTarget.serializer())
    }
}
