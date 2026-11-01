package gg.foreversixty.app.planner

import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.RaceClassCombo
import gg.foreversixty.app.model.RaceInfo
import gg.foreversixty.app.model.Talent
import gg.foreversixty.app.model.TalentTree
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.talents.TalentTestData
import java.io.File
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.int
import kotlinx.serialization.json.intOrNull
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive

/** The planner tests' data: the golden vectors, read in place, and the active build's real files. */
object PlannerTestData {
    fun golden(): JsonObject =
        Json.parseToJsonElement(File(BisTestData.repo, "apps/shared/fixtures/planner-rules.json").readText()).jsonObject

    /** The vectors' slim trees as talent files: only what the rules read is in them. */
    fun goldenIndexes(golden: JsonObject): Map<String, TalentIndex> =
        golden.getValue("classes").jsonArray.associate { element ->
            val cls = element.jsonObject
            val slug = cls.getValue("class_slug").jsonPrimitive.content
            val trees = cls.getValue("trees").jsonArray.map { it.jsonObject.toTree() }
            slug to TalentIndex(TalentsFile(build = golden.getValue("build").jsonPrimitive.content, classId = 0, classSlug = slug, trees = trees))
        }

    private fun JsonObject.toTree() = TalentTree(
        id = getValue("id").jsonPrimitive.int,
        name = getValue("name").jsonPrimitive.content,
        position = getValue("position").jsonPrimitive.int,
        background = "",
        icon = "",
        talents = getValue("talents").jsonArray.map { it.jsonObject.toTalent() },
    )

    private fun JsonObject.toTalent() = Talent(
        id = getValue("id").jsonPrimitive.int,
        name = getValue("name").jsonPrimitive.content,
        icon = "",
        maxRank = getValue("max_rank").jsonPrimitive.int,
        tier = getValue("tier").jsonPrimitive.int,
        column = getValue("column").jsonPrimitive.int,
        prereqTalentId = getValue("prereq_talent_id").jsonPrimitive.intOrNull,
        prereqRank = getValue("prereq_rank").jsonPrimitive.intOrNull,
        spellId = 0,
    )

    fun shaman(): TalentsFile = TalentTestData.shaman()
    fun warrior(): TalentsFile = TalentTestData.warrior()

    fun races(): List<RaceInfo> = dataFile("races.json", ListSerializer(RaceInfo.serializer()))
    fun combos(): List<RaceClassCombo> = dataFile("combos.json", ListSerializer(RaceClassCombo.serializer()))
    fun classes(): List<ClassInfo> = dataFile("classes.json", ListSerializer(ClassInfo.serializer()))

    private fun <T> dataFile(name: String, strategy: kotlinx.serialization.DeserializationStrategy<T>): T =
        ApiJson.decodeFromString(strategy, File(BisTestData.repo, "data/builds/${BisTestData.BUILD}/$name").readText())

    /** The spec's §4.4 Elemental raid build as one legal order. */
    fun elementalOrder(): List<Int> {
        val index = TalentIndex(shaman())
        return gg.foreversixty.app.planner.logic.PlannerRules.orderFromRanks(
            index,
            gg.foreversixty.app.talents.TalentRules.ranksOfBand(TalentTestData.SHAMAN_GUIDE),
        ).order
    }

    fun sharedBuild(title: String? = "Elemental raid", views: Long = 12) = Build(
        id = "k7q2m9xa",
        classId = shaman().classId,
        raceId = ORC,
        treeVersion = BisTestData.BUILD,
        pointOrder = elementalOrder(),
        title = title,
        createdAt = "2026-11-04T19:00:00Z",
        views = views,
    )

    const val ORC = 2
}
