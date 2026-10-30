import DesignSystem
import Networking
import SwiftUI
import UIKit

/// Where game icons come from: `/data/<build>/icons/<icon>.webp` on the site (contract "Static
/// files"). Nil in snapshot tests and previews, so the tile draws without a request.
struct IconSource: Equatable, Sendable {
    let baseURL: URL
    let build: String

    func url(_ icon: String) -> URL {
        baseURL.appending(path: "data/\(build)/icons/\(icon).webp")
    }
}

extension EnvironmentValues {
    @Entry var iconSource: IconSource?
}

/// Loads icons once and keeps them: a 200 MB disk cache and a memory cache (spec §11). A miss
/// or a 404 leaves the tile empty, never a broken-image glyph.
actor IconLoader {
    static let shared = IconLoader()
    static let diskCapacity = 200 * 1024 * 1024
    static let memoryCapacity = 20 * 1024 * 1024

    private let session: URLSession
    private let memory = NSCache<NSURL, UIImage>()

    init() {
        let configuration = URLSessionConfiguration.default
        configuration.urlCache = URLCache(
            memoryCapacity: Self.memoryCapacity,
            diskCapacity: Self.diskCapacity,
            directory: FileManager.default.urls(for: .cachesDirectory, in: .userDomainMask).first?
                .appending(path: "icons")
        )
        configuration.requestCachePolicy = .returnCacheDataElseLoad
        session = URLSession(configuration: configuration)
    }

    func image(_ url: URL) async -> UIImage? {
        if let hit = memory.object(forKey: url as NSURL) { return hit }
        guard let (data, response) = try? await session.data(from: url),
              (response as? HTTPURLResponse)?.statusCode == 200,
              let image = UIImage(data: data)
        else { return nil }
        memory.setObject(image, forKey: url as NSURL)
        return image
    }
}

/// A square game icon from the CDN, faded in over 120 ms once loaded (spec §8 Item icon).
struct RemoteIcon: View {
    let icon: String?
    @Environment(\.iconSource) private var source
    @Environment(\.accessibilityReduceMotion) private var reduceMotion
    @State private var image: UIImage?

    static let fade = 0.12

    var body: some View {
        ZStack {
            if let image {
                Image(uiImage: image).resizable().interpolation(.high).transition(.opacity)
            }
        }
        .task(id: url) {
            guard let url else { return }
            let loaded = await IconLoader.shared.image(url)
            withAnimation(reduceMotion ? nil : .easeOut(duration: Self.fade)) { image = loaded }
        }
    }

    private var url: URL? {
        guard let icon, !icon.isEmpty, let source else { return nil }
        return source.url(icon)
    }
}

/// The quality-ringed item tile (spec §4.3): a 1.5 pt quality border, a 1 pt inner black line
/// and a 10 pt glow at 22%. It draws at once at its real size; the icon fades in.
struct ItemIcon: View {
    let icon: String?
    let quality: Int
    var size: CGFloat = 44

    static let glowOpacity = 0.22
    static let glowRadius: CGFloat = 10
    static let border: CGFloat = 1.5

    var body: some View {
        let shape = RoundedRectangle(cornerRadius: Theme.Radius.pill)
        RemoteIcon(icon: icon)
            .frame(width: size, height: size)
            .background(Theme.Palette.cardTop)
            .clipShape(shape)
            .overlay(shape.inset(by: Self.border).stroke(.black, lineWidth: 1))
            .overlay(shape.stroke(GamePalette.qualityBorder(quality), lineWidth: Self.border))
            .shadow(color: GamePalette.qualityBorder(quality).opacity(Self.glowOpacity), radius: Self.glowRadius / 2)
            .accessibilityHidden(true)
    }
}

/// A spec's talent-tab icon (spec §4.3): square, the pill radius, a 1 pt line.
struct SpecIcon: View {
    let icon: String
    var size: CGFloat = 28

    var body: some View {
        let shape = RoundedRectangle(cornerRadius: Theme.Radius.pill)
        RemoteIcon(icon: icon)
            .frame(width: size, height: size)
            .background(Theme.Palette.cardTop)
            .clipShape(shape)
            .overlay(shape.stroke(Theme.Palette.iconLine, lineWidth: 1))
            .accessibilityHidden(true)
    }
}
