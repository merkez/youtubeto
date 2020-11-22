package client

type YoutubeDL struct {
	c *Client
}

// exec executes an ExecFunc using 'youtube-dl'.
func (ytdl *YoutubeDL) exec(args ...string) ([]byte, error) {
	return ytdl.c.exec("youtube-dl", args...)
}

// DownloadWithOutputName generates Folder named with Playlist name
// downloads videos under given playlist url to Folder
func (ytdl *YoutubeDL) DownloadWithOutputName(folderName, url string) error {
	cmds := []string{"-o", folderName + "/%(playlist_index)s - %(title)s.%(ext)s", url}
	_, err := ytdl.exec(cmds...)
	return err
}
