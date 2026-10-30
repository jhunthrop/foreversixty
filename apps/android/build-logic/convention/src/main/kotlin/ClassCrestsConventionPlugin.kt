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
 * The ringed HD class crests (spec 4.3) from the site's one copy, web/public/icons/hd/crests, as
 * `@drawable/crest_warrior` and so on. Nothing is checked in twice (lane F4).
 */
class ClassCrestsConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val copy = tasks.register("copyClassCrests", CopyClassCrests::class.java) {
            sourceDir.set(rootProject.layout.projectDirectory.dir("../../web/public/icons/hd/crests"))
            outputDir.set(layout.buildDirectory.dir("generated/classCrests"))
        }
        extensions.getByType(AndroidComponentsExtension::class.java).onVariants { variant ->
            variant.sources.res?.addGeneratedSourceDirectory(copy, CopyClassCrests::outputDir)
        }
    }
}

abstract class CopyClassCrests @Inject constructor(
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
            from(sourceDir) { include("*.webp") }
            into(outputDir.dir("drawable-nodpi"))
            rename { "crest_$it" }
        }
    }
}
