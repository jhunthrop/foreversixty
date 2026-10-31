package gg.foreversixty.app.widgets

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Widget
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.widgets.logic.WidgetSnapshot
import java.io.File
import java.time.Instant
import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.jsonObject

/** The shared golden vectors (apps/shared/widgets/figures.json) and snapshots built from the fixtures. */
object WidgetTestData {
    /** From apps/android/app (the unit tests' working directory) to the shared vectors. */
    private val vectors: File = generateSequence(File("").absoluteFile) { it.parentFile }
        .map { File(it, "apps/shared/widgets/figures.json") }
        .firstOrNull { it.exists() }
        ?: error("apps/shared/widgets/figures.json not found above ${File("").absolutePath}")

    val figures: JsonObject by lazy { ApiJson.parseToJsonElement(vectors.readText()).jsonObject }

    /** The fixture's data, with [patch] merged over data.next_raid. */
    fun widget(fixture: String, patch: JsonObject? = null): Widget {
        val data = RaidFixture.text(fixture).let { ApiJson.parseToJsonElement(it).jsonObject.getValue("data").jsonObject }
        val merged = if (patch == null) data else {
            val raid = data.getValue("next_raid").jsonObject
            JsonObject(data + ("next_raid" to JsonObject(raid + patch)))
        }
        return ApiJson.decodeFromJsonElement(Widget.serializer(), merged as JsonElement)
    }

    fun snapshot(fixture: String = "widget.get", fetchedAt: String = "2026-12-08T23:30:00Z", patch: JsonObject? = null) =
        WidgetSnapshot(widget(fixture, patch), fetchedAt)

    /** A next_raid patch with that many signed up (the Live Activity's 38 of 40). */
    fun signedUp(n: Int) = JsonObject(mapOf("signed_up" to kotlinx.serialization.json.JsonPrimitive(n)))

    /** A next_raid patch for a cancelled raid. */
    val cancelled = JsonObject(mapOf("status" to kotlinx.serialization.json.JsonPrimitive("cancelled"), "can_sign_up" to kotlinx.serialization.json.JsonPrimitive(false)))

    /** The board's moment: Tuesday 8 December, 5:39 pm in Chicago, 2 days 2 h before the pull. */
    val boardNow: Instant = Instant.parse("2026-12-08T23:39:00Z")
}
