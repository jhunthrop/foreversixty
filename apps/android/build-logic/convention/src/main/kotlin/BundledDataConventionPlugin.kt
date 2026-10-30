import com.android.build.api.variant.AndroidComponentsExtension
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
 * The repo files the app carries in every variant, each from its one source (lane F3, BiS):
 * - `assets/bundled/specs.json` from data/curated/specs.json, the spec catalogue the CDN does not
 *   publish, `assets/bundled/active-build.json` and `dates.json` from web/src/data, the build and the
 *   key dates a first offline launch reads (neither is on the CDN yet);
 * - `res/drawable-nodpi/crest_<class>.webp` and `faction_<faction>.webp` from web/public/icons/hd,
 *   so a crest or a faction mark never waits on the network.
 */
class BundledDataConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val repo = rootProject.layout.projectDirectory.dir(REPO_ROOT)
        val assets = tasks.register("copyBundledData", CopyBundledData::class.java) {
            repoDir.set(repo.dir("data/curated"))
            webDataDir.set(repo.dir("web/src/data"))
            outputDir.set(layout.buildDirectory.dir("generated/bundledData"))
        }
        val images = tasks.register("copyBrandImages", CopyBrandImages::class.java) {
            iconsDir.set(repo.dir("web/public/icons/hd"))
            outputDir.set(layout.buildDirectory.dir("generated/brandImages"))
        }
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            variant.sources.assets?.addGeneratedSourceDirectory(assets, CopyBundledData::outputDir)
            variant.sources.res?.addGeneratedSourceDirectory(images, CopyBrandImages::outputDir)
        }
    }

    companion object {
        /** From apps/android, the Gradle root, to the repository root. */
        const val REPO_ROOT = "../.."
    }
}

abstract class CopyBundledData @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val repoDir: DirectoryProperty

    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val webDataDir: DirectoryProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.sync {
            from(repoDir) { include("specs.json") }
            from(webDataDir) { include("active-build.json", "dates.json") }
            into(outputDir.dir(ASSET_DIR))
        }
    }

    companion object {
        const val ASSET_DIR = "bundled"
    }
}

abstract class CopyBrandImages @Inject constructor(
    private val fs: FileSystemOperations,
) : DefaultTask() {
    @get:InputDirectory
    @get:PathSensitive(PathSensitivity.RELATIVE)
    abstract val iconsDir: DirectoryProperty

    @get:OutputDirectory
    abstract val outputDir: DirectoryProperty

    @TaskAction
    fun copy() {
        fs.sync {
            from(iconsDir.dir("crests")) {
                include("*.webp")
                rename { "crest_$it" }
            }
            from(iconsDir.dir("faction")) {
                include("horde.webp", "alliance.webp")
                rename { "faction_$it" }
            }
            into(outputDir.dir("drawable-nodpi"))
        }
    }
}
