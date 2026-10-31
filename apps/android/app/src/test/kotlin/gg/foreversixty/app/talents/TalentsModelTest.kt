package gg.foreversixty.app.talents

import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.BuildTalents
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Test
import java.io.IOException

/** How the Talents screen loads (spec F6 §4, §6). iOS: TalentsModelTests.swift. */
class TalentsModelTest {
    private val warrior = TalentTestData.warrior()
    private val fury = TalentTestData.fury()
    private val catalog = BisTestData.catalog()
    private val yell = HomeCharacter.from(BisTestData.me().characters.first { it.name == "Obnoxious Yell" })

    private class FakeSource(
        val talents: ApiResult<CachedFile<TalentsFile>>,
        val older: ApiResult<CachedFile<TalentsFile>>,
        val bis: ApiResult<CachedFile<BisFile>>,
    ) : TalentsSource {
        override suspend fun activeBuild() = BisTestData.BUILD
        override suspend fun talents(build: String, classSlug: String) = if (build == BisTestData.BUILD) talents else older
        override suspend fun bis(build: String, spec: String) = bis
    }

    private fun <T> fresh(value: T): ApiResult<CachedFile<T>> = ApiResult.Success(CachedFile(value, BisTestData.NOW, false), "")
    private val notFound = ApiResult.Failure(ApiError.Status(404))
    private val offline = ApiResult.Failure(ApiError.Network(IOException("offline")))

    private fun model(
        character: HomeCharacter = yell,
        talents: ApiResult<CachedFile<TalentsFile>> = fresh(warrior),
        older: ApiResult<CachedFile<TalentsFile>> = notFound,
        bis: ApiResult<CachedFile<BisFile>> = fresh(fury),
    ) = TalentsModel(FakeSource(talents, older, bis), TalentCharacter.from(character), catalog)

    private fun loaded(model: TalentsModel): TalentsPhase.Loaded {
        runBlocking { model.load() }
        return model.phase.value as? TalentsPhase.Loaded ?: error("not loaded: ${model.phase.value}")
    }

    @Test
    fun `the fixture character compares with the 50 to 59 guide build`() {
        val phase = loaded(model())
        assertEquals(TalentGuideBuild.Build(TalentTestData.BAND_50, 50), phase.presenter.guide)
        assertEquals(listOf(17, 32, 0), phase.presenter.tabs(TalentBuildKind.Yours).map { it.points })
        assertNull(phase.offlineSinceMillis)
    }

    @Test
    fun `at 60 the raid entry is the guide`() {
        val phase = loaded(model(yell.copy(level = 60)))
        assertEquals(TalentGuideBuild.Build(TalentTestData.BAND_60, 60), phase.presenter.guide)
    }

    @Test
    fun `saved copies set the offline stamp`() {
        val phase = loaded(model(talents = ApiResult.Success(CachedFile(warrior, BisTestData.NOW, true), "")))
        assertEquals(BisTestData.NOW, phase.offlineSinceMillis)
    }

    @Test
    fun `a talent file that cannot load is the error state`() {
        val model = model(talents = offline)
        runBlocking { model.load() }
        assertEquals(TalentsPhase.Failed, model.phase.value)
    }

    @Test
    fun `a missing BiS file is no build, a failed one is the guide-failed line`() {
        assertEquals(TalentGuideBuild.None, loaded(model(bis = notFound)).presenter.guide)
        assertEquals(TalentGuideBuild.Failed, loaded(model(bis = offline)).presenter.guide)
    }

    @Test
    fun `under 20 and with no spec there is no guide to read`() {
        assertEquals(TalentGuideBuild.None, loaded(model(yell.copy(level = 15))).presenter.guide)
        assertEquals(TalentGuideBuild.NoSpec, loaded(model(yell.copy(specName = null))).presenter.guide)
    }

    @Test
    fun `an export of zeros is no export`() {
        val zeros = yell.copy(talents = BuildTalents(listOf("0", "0", "0"), listOf(0, 0, 0)))
        assertEquals(TalentExport.None, loaded(model(zeros)).presenter.export)
    }

    @Test
    fun `an export from another build draws only when its trees are laid out the same`() {
        val older = yell.copy(dataBuild = "1.60.1.70532")
        assertNotNull(loaded(model(older, older = fresh(warrior))).presenter.yourRanks)
        assertEquals(TalentExport.Stale, loaded(model(older, older = fresh(TalentTestData.shaman()))).presenter.export)
        assertEquals(TalentExport.Stale, loaded(model(older)).presenter.export)
    }
}
