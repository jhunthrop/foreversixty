package gg.foreversixty.app.talents

import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.TalentsFile
import java.io.File

/**
 * The F6 tests' data: the contract fixtures (the real warrior talent file and the Fury BiS bands
 * of the active build), the active build's shaman file read in place, and named builds.
 * iOS: TalentTestData.swift.
 */
object TalentTestData {
    fun warrior(): TalentsFile = BisTestData.fixture("cdn.talents.warrior", TalentsFile.serializer())
    fun fury(): BisFile = BisTestData.bis()
    fun shaman(): TalentsFile = ApiJson.decodeFromString(
        TalentsFile.serializer(),
        File(BisTestData.repo, "data/builds/${BisTestData.BUILD}/talents/shaman.json").readText(),
    )

    /** The Fury BiS file's 50 to 59 Horde guide build (41 points). */
    const val BAND_50 = "34200000000000000-05153105022011501-000000000000000000"
    /** The Fury BiS file's level-60 raid guide build (51 points). */
    const val BAND_60 = "34320003002000000-05153105022011501-200000000000000000"
    /** Obnoxious Yell's build in `me.json`: the level-60 path less Improved Bloodrage, 49 points. */
    val yoursTrees: List<List<Int>> = listOf("34320003002", "05153105022011501", "0").map(TalentRules::ranksOfTree)

    /** The spec's §4.4 vector: Elemental Shaman, level 60, Horde, band 60 raid. */
    const val SHAMAN_GUIDE = "5510331300001350-030000000000000000-0533520000000000"
    const val SHAMAN_YOURS = "5512311100001350-030000000000000000-0433520000000000"
}
