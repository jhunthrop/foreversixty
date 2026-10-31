package gg.foreversixty.app.home

import android.content.res.Resources
import android.graphics.BitmapFactory
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.graphics.painter.BitmapPainter
import androidx.compose.ui.graphics.painter.Painter
import gg.foreversixty.app.items.ItemIconPainterSource
import gg.foreversixty.app.items.ItemTooltipSources
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.FinishedSim
import gg.foreversixty.app.model.ItemSet
import gg.foreversixty.app.model.ItemsFile
import java.io.File
import java.time.Instant
import java.time.ZoneOffset
import kotlinx.serialization.KSerializer
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonNull
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.jsonObject

/** The contract fixtures (copied onto the test classpath) and the repo's icons, for the F2 tests. */
object HomeFixtures {
    /** The home fixture's character: Velaren, an Orc Beast Mastery Hunter. */
    val velaren = HomeCharacter(
        key = "us/normal/velaren", region = "us", ruleset = "normal", name = "Velaren", race = "orc",
        className = "hunter", specName = "Beast Mastery", level = 60, guild = "Olympus XXVII", faction = "horde",
    )

    /** 2026-10-29 23:15 UTC: two hours after the fixture's sync. */
    val now: Instant = Instant.parse("2026-10-29T23:15:00Z")

    const val BUILD = "1.60.1.70658"

    fun text(name: String): String {
        val resource = HomeFixtures::class.java.classLoader?.getResource("fixtures/app/$name.json")
        return checkNotNull(resource) { "no fixture $name" }.readText()
    }

    /** An enveloped API fixture's `data`, optionally switched to an empty [state]. */
    fun home(name: String = "character.home", state: String? = null): CharacterHome {
        val root = ApiJson.parseToJsonElement(text(name)).jsonObject
        var data = checkNotNull(root["data"]).jsonObject
        if (state != null) {
            data = JsonObject(
                data + mapOf<String, JsonElement>(
                    "state" to JsonPrimitive(state), "upgrades" to JsonArray(emptyList()),
                    "already_best" to JsonArray(emptyList()), "total_gain" to JsonNull, "upgrade_count" to JsonPrimitive(0),
                ),
            )
        }
        return ApiJson.decodeFromJsonElement(CharacterHome.serializer(), data)
    }

    fun <T> staticFile(name: String, serializer: KSerializer<T>): T = ApiJson.decodeFromString(serializer, text(name))

    fun tooltipSources() = ItemTooltipSources(
        items = staticFile("cdn.items.hunter", ItemsFile.serializer()).items.associateBy { it.id },
        sets = staticFile("cdn.sets", ListSerializer(ItemSet.serializer())),
        bis = staticFile("cdn.bis.hunter-beast-mastery", BisFile.serializer()),
    )

    fun presenter(res: Resources, sims: List<FinishedSim> = emptyList()) =
        HomePresenter(velaren, sims, HomeTime(now, ZoneOffset.UTC), res)

    /** The repo root, from the app module's working directory (apps/android/app). */
    private val repoRoot: File = File(System.getProperty("user.dir").orEmpty()).absoluteFile.parentFile?.parentFile?.parentFile
        ?: File(".")

    /** Item icons from the repo's own build folder, so a screenshot shows the real art. */
    val repoIcons = ItemIconPainterSource { icon, build -> repoIcon(icon, build) }

    @Composable
    private fun repoIcon(icon: String, build: String): Painter? = remember(icon, build) {
        val file = File(repoRoot, "data/builds/$build/icons/$icon.webp")
        if (!file.exists()) return@remember null
        BitmapFactory.decodeFile(file.path)?.let { BitmapPainter(it.asImageBitmap()) }
    }
}
