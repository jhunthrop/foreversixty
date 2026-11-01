package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.auth.session.StoredValue
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.SerializationException
import kotlinx.serialization.json.Json

/** One class's draft as the phone keeps it (spec F13.1 Draft). */
@Serializable
data class PlannerDraft(
    @SerialName("class") val classSlug: String,
    @SerialName("race") val raceSlug: String?,
    @SerialName("level") val level: Int,
    @SerialName("order") val order: List<Int>,
)

/**
 * The drafts, one per class, saved as the player edits and read when `Plan a build` reopens, so a
 * draft survives the process being killed. [valueFor] hands back the class's stored value
 * (private prefs in the app, memory in tests). A draft that does not decode reads as none.
 */
class PlannerDrafts(private val valueFor: (String) -> StoredValue) {
    private val json = Json { ignoreUnknownKeys = true }

    fun load(classSlug: String): PlannerDraft? {
        val text = valueFor(classSlug).load() ?: return null
        return try {
            json.decodeFromString(PlannerDraft.serializer(), text).takeIf { it.classSlug == classSlug }
        } catch (e: SerializationException) {
            // SerializationException is the decoder's IllegalArgumentException: a bad draft is no draft.
            null
        }
    }

    fun save(draft: PlannerDraft) = valueFor(draft.classSlug).save(json.encodeToString(PlannerDraft.serializer(), draft))

    companion object {
        /** The prefs key of a class's draft. */
        fun key(classSlug: String): String = "planner_draft_$classSlug"
    }
}
