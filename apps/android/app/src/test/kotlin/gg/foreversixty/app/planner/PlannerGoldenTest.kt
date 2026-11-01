package gg.foreversixty.app.planner

import android.content.Context
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.planner.logic.Fs1
import gg.foreversixty.app.planner.logic.PlannerCopy
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.TalentIndex
import gg.foreversixty.app.talents.TalentRules
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.int
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/**
 * Replays every vector of apps/shared/fixtures/planner-rules.json, generated from the web planner
 * (spec F13 "Rules live in one place"): CI fails the moment the app disagrees with the site.
 */
@RunWith(AndroidJUnit4::class)
class PlannerGoldenTest {
    private val copy = PlannerCopy(ApplicationProvider.getApplicationContext<Context>().resources)
    private val golden = PlannerTestData.golden()
    private val indexes: Map<String, TalentIndex> = PlannerTestData.goldenIndexes(golden)

    private fun index(vector: JsonObject) = requireNotNull(indexes[vector.string("class_slug")])

    @Test fun `the constants match`() {
        val constants = golden.getValue("constants").jsonObject
        assertEquals(constants.int("max_points"), TalentRules.MAX_POINTS)
        assertEquals(constants.int("points_per_tier"), TalentRules.POINTS_PER_TIER)
        assertEquals(constants.int("first_point_level"), TalentRules.BASE_LEVEL + 1)
        assertEquals(golden.string("none_left"), copy.refusal(gg.foreversixty.app.planner.logic.Refusal.NoneLeft))
    }

    @Test fun `points available at every level`() = vectors("points_available") { v ->
        assertEquals(v.toString(), v.int("points"), PlannerRules.pointsAvailable(v.int("level")))
    }

    @Test fun `every add`() = vectors("add") { v ->
        val refusal = PlannerRules.refuseAdd(index(v), v.ints("order"), v.int("talent"), v.int("level"))
        assertEquals(v.toString(), v.reason(), refusal?.let(copy::refusal))
    }

    @Test fun `every remove`() = vectors("remove") { v ->
        val refusal = PlannerRules.refuseRemove(index(v), v.ints("order"), v.int("talent"))
        assertEquals(v.toString(), v.reason(), refusal?.let(copy::refusal))
    }

    @Test fun `every validation`() = vectors("validate") { v ->
        val want = v.getValue("errors").jsonArray.map { it.jsonObject.string("field") to it.jsonObject.string("message") }
        val got = PlannerRules.validate(index(v), v.ints("order")).map { it.field to copy.refusal(it.refusal) }
        assertEquals(v.toString(), want, got)
    }

    @Test fun `every split`() = vectors("split") { v ->
        assertEquals(v.toString(), v.ints("totals"), PlannerRules.pointsPerTree(index(v), v.ints("order")))
    }

    @Test fun `every FS1 code to an order`() = vectors("fs1") { v ->
        val decoded = requireNotNull(Fs1.decode(v.string("code")))
        assertEquals(v.string("class_slug"), decoded.classSlug)
        assertEquals(v.string("race_slug"), decoded.raceSlug)
        assertEquals(v.getValue("tree_ranks").jsonArray.map { tree -> tree.jsonArray.map { it.jsonPrimitive.int } }, decoded.treeRanks)
        val result = PlannerRules.orderFromRanks(index(v), decoded.treeRanks)
        assertEquals(v.toString(), v.ints("order"), result.order)
        assertEquals(v.toString(), v.ints("dropped"), result.dropped)
    }

    private fun vectors(kind: String, check: (JsonObject) -> Unit) {
        val all = golden.getValue(kind).jsonArray.map { it.jsonObject }
        assertTrue("no $kind vectors", all.isNotEmpty())
        all.forEach(check)
    }
}

private fun JsonObject.string(key: String) = getValue(key).jsonPrimitive.content
private fun JsonObject.int(key: String) = getValue(key).jsonPrimitive.int
private fun JsonObject.ints(key: String) = getValue(key).jsonArray.map { it.jsonPrimitive.int }
private fun JsonObject.reason(): String? = get("reason")?.jsonPrimitive?.takeIf { it.isString }?.content
