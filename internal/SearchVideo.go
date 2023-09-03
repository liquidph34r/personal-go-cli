package video

import "flag"

type SearchVideo struct {
	// Define the flags
	title     string
	order     string
	mediaType string
	videoDef  string
}

func GetSearchVideoFlags(args []string) *SearchVideo {
	searchVideo := new(SearchVideo)
	searchVideo.title = *flag.String("title", "", "The title of the video")
	searchVideo.order = *flag.String("order", "relevance", "The order of the videos")
	searchVideo.mediaType = *flag.String("mediaType", "video", "The type of the media")
	searchVideo.videoDef = *flag.String("videoDef", "any", "The definition of the video")

	return searchVideo
}
