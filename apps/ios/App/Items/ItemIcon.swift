import DesignTokens
import SwiftUI
import UIKit

/// Where item icons come from. Live: `/data/{build}/icons/{icon}.webp` on the site, through a
/// URLCache with a 200 MB disk cap (spec §11) and a decoded-image memory cache. Tests and
/// captures inject a source that reads the repo's own icon files.
protocol ItemIconSource: Sendable {
    /// An image already in memory, so a cached icon draws in the first frame.
    func cached(icon: String, build: String) -> UIImage?
    func load(icon: String, build: String) async -> UIImage?
}

/// NSCache and URLSession are thread-safe; nothing else here mutates.
final class RemoteItemIconSource: ItemIconSource, @unchecked Sendable {
    static let shared = RemoteItemIconSource()

    private let session: URLSession
    private let siteURL: URL
    private let images = NSCache<NSString, UIImage>()

    init(siteURL: URL = URL(string: "https://foreversixty.gg")!, diskCapacity: Int = 200 * 1024 * 1024) {
        let configuration = URLSessionConfiguration.default
        configuration.urlCache = URLCache(memoryCapacity: 16 * 1024 * 1024, diskCapacity: diskCapacity, directory: Self.cacheDirectory)
        configuration.requestCachePolicy = .returnCacheDataElseLoad
        session = URLSession(configuration: configuration)
        self.siteURL = siteURL
    }

    private static var cacheDirectory: URL? {
        FileManager.default.urls(for: .cachesDirectory, in: .userDomainMask).first?.appending(path: "item-icons")
    }

    func url(icon: String, build: String) -> URL {
        siteURL.appending(path: "data/\(build)/icons/\(icon).webp")
    }

    func cached(icon: String, build: String) -> UIImage? {
        images.object(forKey: key(icon, build))
    }

    func load(icon: String, build: String) async -> UIImage? {
        if let image = cached(icon: icon, build: build) { return image }
        guard let (data, response) = try? await session.data(from: url(icon: icon, build: build)),
              (response as? HTTPURLResponse)?.statusCode == 200,
              let image = UIImage(data: data)
        else { return nil }
        images.setObject(image, forKey: key(icon, build))
        return image
    }

    private func key(_ icon: String, _ build: String) -> NSString { "\(build)/\(icon)" as NSString }
}

private struct ItemIconSourceKey: EnvironmentKey {
    static let defaultValue: any ItemIconSource = RemoteItemIconSource.shared
}

extension EnvironmentValues {
    var itemIconSource: any ItemIconSource {
        get { self[ItemIconSourceKey.self] }
        set { self[ItemIconSourceKey.self] = newValue }
    }
}

/// A square game icon in its quality border (spec §4.3): 1.5 pt quality border, a 1 pt inner
/// black line and a 10 pt quality glow at 22%. The tile draws at once on `#131824`; the image
/// fades in over 120 ms, and a missing icon leaves the tile, never a broken-image glyph (§8).
struct ItemIconView: View {
    let icon: String?
    let quality: Int?
    let build: String
    let size: CGFloat
    var dashed = false

    @Environment(\.itemIconSource) private var source
    @State private var image: UIImage?

    var body: some View {
        let color = ItemQuality(value: quality).border
        let shape = RoundedRectangle(cornerRadius: FSRadius.pill)
        ZStack {
            FSColor.cardTop
            if let shown = image ?? icon.flatMap({ source.cached(icon: $0, build: build) }) {
                Image(uiImage: shown).resizable().scaledToFill().transition(.opacity)
            }
        }
        .frame(width: size, height: size)
        .clipShape(shape)
        .overlay(shape.inset(by: 1.5).stroke(Color.black, lineWidth: 1))
        .overlay(dashed ? AnyView(shape.stroke(FSColor.line, style: StrokeStyle(lineWidth: 1.5, dash: [4, 3])))
            : AnyView(shape.stroke(color, lineWidth: 1.5)))
        .shadow(color: dashed ? .clear : color.opacity(0.22), radius: 5)
        .accessibilityHidden(true)
        .task(id: icon) {
            guard let icon, image == nil else { return }
            let loaded = await source.load(icon: icon, build: build)
            withAnimation(.easeIn(duration: 0.12)) { image = loaded }
        }
    }
}
