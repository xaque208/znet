	"path/filepath"
			return nil, err
	// FileDiff is added/deleted file
	// No further collection of hunks needed
	if fd.NewName == "" {
		return fd, nil
	}

		fd.OrigTime = origTime
		fd.NewTime = newTime
// timestamps). Or which starts with "Only in " with dir path and filename.
// "Only in" message is supported in POSIX locale: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/diff.html#tag_20_34_10
	if r.fileHeaderLine != nil {
		if isOnlyMessage, source, filename := parseOnlyInMessage(r.fileHeaderLine); isOnlyMessage {
			return filepath.Join(string(source), string(filename)),
				"", nil, nil, nil
		}
	}

	return fmt.Sprintf("overflowed into next file: %s", string(e))
		// Reached message that file is added/deleted
		if isOnlyInMessage, _, _ := parseOnlyInMessage(line); isOnlyInMessage {
			r.fileHeaderLine = line // pass to readOneFileHeader (see fileHeaderLine field doc)
			return xheaders, nil
		}

	case lineCount == 6 && strings.HasPrefix(fd.Extended[5], "Binary files ") && strings.HasPrefix(fd.Extended[2], "rename from ") && strings.HasPrefix(fd.Extended[3], "rename to "):
		names := strings.SplitN(fd.Extended[0][len("diff --git "):], " ", 2)
		fd.OrigName = names[0]
		fd.NewName = names[1]
		return true

	// ErrBadOnlyInMessage is when a file have a malformed `only in` message
	// Should be in format `Only in {source}: {filename}`
	ErrBadOnlyInMessage = errors.New("bad 'only in' message")

			// If the line starts with `---` and the next one with `+++` we're
			// looking at a non-extended file header and need to abort.
			if bytes.HasPrefix(line, []byte("---")) {
				ok, err := peekPrefix(r.reader, "+++")
				if err != nil {
					return r.hunk, err
				}
				if ok {
					return r.hunk, &ParseError{r.line, r.offset, &ErrBadHunkLine{Line: line}}
				}
			}

			// If the line starts with the hunk prefix, this hunk is complete.
				// But we've already read in the next hunk's
// peekPrefix peeks into the given reader to check whether the next
// bytes match the given prefix.
func peekPrefix(reader *bufio.Reader, prefix string) (bool, error) {
	next, err := reader.Peek(len(prefix))
	if err != nil {
		if err == io.EOF {
			return false, nil
		}
		return false, err
	}
	return bytes.HasPrefix(next, []byte(prefix)), nil
}

// parseOnlyInMessage checks if line is a "Only in {source}: {filename}" and returns source and filename
func parseOnlyInMessage(line []byte) (bool, []byte, []byte) {
	if !bytes.HasPrefix(line, onlyInMessagePrefix) {
		return false, nil, nil
	}
	line = line[len(onlyInMessagePrefix):]
	idx := bytes.Index(line, []byte(": "))
	if idx < 0 {
		return false, nil, nil
	}
	return true, line[:idx], line[idx+2:]
}
