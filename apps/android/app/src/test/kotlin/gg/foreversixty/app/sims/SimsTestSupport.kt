package gg.foreversixty.app.sims

import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.ItemsFile
import gg.foreversixty.app.model.SimList
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimResult
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.model.SimStarted
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.ok
import gg.foreversixty.app.sims.logic.SimGearLook
import gg.foreversixty.app.sims.logic.SimsContent
import gg.foreversixty.app.sims.logic.SimsService
import java.io.File
import java.time.Instant
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.buildJsonArray
import kotlinx.serialization.json.buildJsonObject
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.put

/**
 * F5's test data, the same as iOS's SimsTestData: the contract fixtures, plus a level-60 Fury set
 * from the build's own BiS list for the result boards. Every helper throws an ordinary exception.
 */
object SimsTestData {
    const val BUILD = "1.60.1.70658"
    val now: Instant = Instant.parse("2026-12-15T05:22:10Z")
    private val repo = File("../../..")

    /** Wearing Cape of the Black Baron where Howler's Furs is best: the Top Gear board's one change. */
    val furyGear = listOf(
        "head" to 12640, "neck" to 22150, "shoulder" to 272105, "back" to 13340, "chest" to 252483,
        "wrist" to 19587, "hands" to 272095, "waist" to 20216, "legs" to 22385, "feet" to 21995,
        "finger1" to 234034, "finger2" to 19325, "trinket1" to 19289, "trinket2" to 249469,
        "main_hand" to 21520, "off_hand" to 21522, "ranged" to 22656,
    )
    const val HOWLERS_FURS = 272414

    fun specs(): SpecCatalog = SpecCatalog.decode(File(repo, "data/curated/specs.json").readText())

    fun list(): SimList = RaidFixture.load(SimList.serializer(), "sims.list")

    fun row(id: String): SimRow = requireNotNull(list().rows.firstOrNull { it.simId == id }) { "no row $id" }

    val browserRow = SimRow(
        simId = "p2k7d0wq5nme", spec = "warrior-fury", kind = "run", dps = 851.9, headline = "851.9 DPS", state = "done",
        createdAt = "2026-12-12T20:40:00Z", title = "Bare 60", lane = "browser", seen = true,
    )

    val topGearRow = SimRow(
        simId = "g5tq8n2rk1wa", spec = "warrior-fury", kind = "gear", headline = "+4 DPS from Howler's Furs", state = "done",
        createdAt = "2026-12-15T04:02:00Z", title = "Cloak check", lane = "server", finishedAt = "2026-12-15T04:05:40Z", seen = true,
    )

    fun content(offline: Instant? = null, server: List<SimRow>? = null, mine: List<SimRow>? = null): SimsContent {
        val rows = list().rows
        val mineRows = mine ?: (rows + browserRow)
        return SimsContent(server ?: rows, mineRows, mineRows.size, 1, offline)
    }

    /** The fixture result with the Fury set, 10,000 iterations, and for [topGear] the swap into Howler's Furs. */
    fun result(topGear: Boolean = false, untitled: Boolean = false): SimResult {
        val root = ApiJson.parseToJsonElement(RaidFixture.text("sims.result")).jsonObject
        val data = requireNotNull(root["data"]).jsonObject
        val request = requireNotNull(data["request"]).jsonObject
        val character = requireNotNull(request["character"]).jsonObject
        val gear = buildJsonArray { furyGear.forEach { (slot, id) -> add(buildJsonObject { put("slot", slot); put("item_id", id) }) } }
        val patched = data.toMutableMap().apply {
            put("request", JsonObject(request + ("character" to JsonObject(character + ("gear" to gear)))))
            put("dps", buildJsonObject { put("mean", 877.54); put("stddev", 61.2); put("error", 1.22); put("min", 702.1); put("max", 1041.8) })
            put("iterations_run", JsonPrimitive(10000))
            if (untitled) remove("title")
            if (topGear) {
                put(
                    "combos",
                    JsonArray(
                        listOf(
                            buildJsonObject {
                                put("substitutions", buildJsonArray { add(buildJsonObject { put("kind", "item"); put("slot", "back"); put("item_id", HOWLERS_FURS) }) })
                                put("dps", buildJsonObject { put("mean", 881.6) })
                                put("delta", buildJsonObject { put("mean", 4.07) })
                                put("group", 0)
                            },
                        ),
                    ),
                )
            }
        }
        return ApiJson.decodeFromJsonElement(SimResult.serializer(), JsonObject(patched))
    }

    /** The build's warrior items, as the gear grid's names and icons. */
    fun looks(): Map<Int, SimGearLook> {
        val items = ApiJson.decodeFromString(ItemsFile.serializer(), File(repo, "data/builds/$BUILD/items/warrior.json").readText()).items
        val wanted = furyGear.map { it.second }.toSet() + HOWLERS_FURS
        return items.filter { it.id in wanted }.associate { it.id to SimGearLook(it.name, it.icon, it.quality) }
    }
}

/** A scripted sims service: each call answers from its lambda; reruns and seen marks are recorded. */
class FakeSimsService : SimsService {
    var listAnswer: (Int, String?) -> ApiResult<SimList> = { _, lane ->
        val list = SimsTestData.list()
        ok(if (lane == "server") list.copy(rows = list.rows.filter { it.lane == "server" }) else list)
    }
    var progressAnswer: (String) -> ApiResult<SimProgress> = { ok(RaidFixture.load(SimProgress.serializer(), "sims.progress")) }
    var rerunAnswer: (String) -> ApiResult<SimStarted> = { ok(RaidFixture.load(SimStarted.serializer(), "sims.rerun")) }
    var resultAnswer: (String) -> ApiResult<SimResult> = { ok(RaidFixture.load(SimResult.serializer(), "sims.result")) }
    val lists = mutableListOf<Pair<Int, String?>>()
    val reruns = mutableListOf<String>()
    val seen = mutableListOf<String>()
    val polled = mutableListOf<String>()

    override suspend fun list(page: Int, lane: String?): ApiResult<SimList> {
        lists += page to lane
        return listAnswer(page, lane)
    }
    override suspend fun result(id: String) = resultAnswer(id)
    override suspend fun progress(id: String): ApiResult<SimProgress> {
        polled += id
        return progressAnswer(id)
    }
    override suspend fun rerun(id: String): ApiResult<SimStarted> {
        reruns += id
        return rerunAnswer(id)
    }
    override suspend fun markSeen(id: String): ApiResult<Unit> {
        seen += id
        return ok(Unit)
    }
}
