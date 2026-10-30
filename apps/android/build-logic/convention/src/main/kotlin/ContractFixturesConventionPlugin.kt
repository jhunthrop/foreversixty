import com.android.build.api.variant.AndroidComponentsExtension
import com.android.build.api.variant.HasHostTests
import org.gradle.api.DefaultTask
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.api.file.DirectoryProperty
import org.gradle.api.file.FileSystemOperations
import org.gradle.api.tasks.InputDirectory
import org.gradle.api.tasks.OutputDirectory
import org.gradle.api.tasks.PathSensitive
import org.gradle.api.tasks.PathSensitivity
import org.gradle.api.tasks.TaskAction
import javax.inject.Inject

/**
 * Lane C's contract fixtures (docs/contracts/fixtures/app at the repo root) are the single
 * source of truth for every API shape. This plugin copies them, unchanged, to `fixtures/app/`
 * on the JVM unit tests' classpath and, in the application module, into the debug APK's assets,
 * so the app's fixture mode and every test read exactly the files the Go contract test pins.
 * Release builds never carry them.
 */
class ContractFixturesConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val source = rootProject.layout.projectDirectory.dir(CONTRACT_FIXTURES_DIR)
        val copy = tasks.register("copyContractFixtures", CopyContractFixtures::class.java) {
            sourceDir.set(source)
            outputDir.set(layout.buildDirectory.dir("generated/contractFixtures"))
        }
        val isApp = pluginManager.hasPlugin("com.android.application")
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            if (isApp && variant.buildType == "debug") {
                variant.sources.assets?.addGeneratedSourceDirectory(copy, CopyContractFixtures::outputDir)
            }
            (variant as? HasHostTests)?.hostTests?.values?.forEach { hostTest ->
                hostTest.sources.resources?.addGeneratedSourceDirectory(copy, CopyContractFixtures::outputDir)
            }
        }
    }

    companion object {
        /** From apps/android, the Gradle root, to the repo's contract fixtures. */
        const val CONTRACT_FIXTURES_DIR = "../../docs/contracts/fixtures/app"
    }
}

abstract class CopyContractFixtures @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val sourceDir: DirectoryProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.sync {
            from(sourceDir) { include("**/*.json") }
            into(outputDir.dir(ASSET_PREFIX))
        }
    }

    companion object {
        /** Where the fixtures sit inside assets and the test classpath. */
        const val ASSET_PREFIX = "fixtures/app"
    }
}
