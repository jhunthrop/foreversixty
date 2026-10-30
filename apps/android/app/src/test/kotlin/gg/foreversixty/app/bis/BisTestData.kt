package gg.foreversixty.app.bis

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.Me
import kotlinx.serialization.DeserializationStrategy
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.json.jsonObject
import java.io.File

/**
 * The repo's files, read in place: the contract fixtures (on the test classpath), the spec
 * catalogue and the generator's real output for the active build. A missing file fails the test
 * that asked for it, never the run.
 */
object BisTestData {
    /** From apps/android/app, Gradle's working directory for the module's tests, to the repo root. */
    val repo = File("../../..")
    const val BUILD = "1.60.1.70658"

    fun fixtureText(name: String): String =
        requireNotNull(javaClass.classLoader?.getResource("fixtures/app/$name.json")) { "no fixture $name" }.readText()

    fun <T> fixture(name: String, strategy: DeserializationStrategy<T>): T = ApiJson.decodeFromString(strategy, fixtureText(name))

    fun me(): Me {
        val root = ApiJson.parseToJsonElement(fixtureText("me")).jsonObject
        return ApiJson.decodeFromJsonElement(Me.serializer(), requireNotNull(root["data"]))
    }

    fun catalog(): SpecCatalog = SpecCatalog.decode(File(repo, "data/curated/specs.json").readText())

    fun bis(): BisFile = fixture("cdn.bis.warrior-fury", BisFile.serializer())
    fun items(): ItemsFile = fixture("cdn.items.warrior", ItemsFile.serializer())
    fun classes(): List<ClassInfo> = fixture("cdn.classes", ListSerializer(ClassInfo.serializer()))

    /** The fixture's level-60 Horde raid list for Fury. */
    val raid = BisSelection("warrior", "warrior-fury", BisRange(60, BisRange.RAID), "horde")

    /** A fixed clock: Fri 30 Oct 2026, 21:41 UTC. */
    const val NOW = 1_793_396_460_000L
}
