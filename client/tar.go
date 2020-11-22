package client

type Tar struct {
	c *Client
}

// exec executes an ExecFunc using 'youtube-dl'.
func (tr *Tar) exec(args ...string) ([]byte, error) {
	return tr.c.exec("tar", args...)
}

func (tr *Tar) CompressWithPIGZ(fileName, folderToCompress string) error {
	cmds := []string{"--use-compress-program=pigz", "-cf", fileName, folderToCompress}
	_, err := tr.exec(cmds...)
	if err != nil {
		return err
	}
	return nil
}
