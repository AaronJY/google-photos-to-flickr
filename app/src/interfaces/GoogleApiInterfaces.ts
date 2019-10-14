export interface Photo {
    cameraMake: string;
    cameraModel: string;
    focalLength?: number;
    apertureFNumber?: number;
    isoEquivalent?: number;
}

export interface Video {
    fps: number;
    status: string;
}

export interface MediaMetadata {
    creationTime: Date;
    width: string;
    height: string;
    photo: Photo;
    video: Video;
}

export interface MediaItem {
    id: string;
    productUrl: string;
    baseUrl: string;
    mimeType: string;
    mediaMetadata: MediaMetadata;
    filename: string;
}

// TODO: Need better name. Lookup Google docs
export interface GoogleResponseRootObject {
    mediaItems: MediaItem[];
    nextPageToken: string;
}