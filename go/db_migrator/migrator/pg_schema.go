package migrator

import (
	"bytes"
	"codeforge/config"
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const headerComment = "-- Generated by pgdumpschema using pg_dump -s\n\n"

// Dump dumps a database schema based on the given config.
func Dump() error {

	if config.Cfg.PgDBName == "" {
		return errors.New("must supply a database name")
	}
	program := config.Cfg.PGDump
	if program == "" {
		program = "pg_dump"
	}
	args := []string{"-s", "-d", config.Cfg.PgDBName, "-w"}
	// If the pg dump command begins with docker, we are most likely attempting to run
	// docker exec <container name> pg_dump ...
	// In this case, we need to set program to docker, and prepend exec, container_name, and pg_dump to args
	if strings.HasPrefix(config.Cfg.PGDump, "docker") {
		program = "docker"
		args = append(strings.Split(config.Cfg.PGDump, " ")[1:], args...)
	}
	for _, schema := range config.Cfg.Schemas {
		args = append(args, "-n", schema)
	}
	if config.Cfg.PgHost != "" {
		args = append(args, "-h", config.Cfg.PgHost)
	}
	if config.Cfg.PgPort != 0 {
		args = append(args, "-p", strconv.Itoa(config.Cfg.PgPort))
	}
	if config.Cfg.PgUser != "" {
		args = append(args, "-U", config.Cfg.PgUser)
	}
	cmd := exec.Command(program, args...)
	if config.Cfg.PgPassword != "" {
		cmd.Env = append(os.Environ(), "PGPASSWORD="+config.Cfg.PgPassword)
	}
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return err
	}
	output := buf.Bytes()
	if !config.Cfg.NoClean {
		output = clean(output)
	}
	writer := config.Cfg.Writer
	if writer == nil {
		writer = os.Stdout
	}
	if !config.Cfg.NoHeader {
		_, err = io.WriteString(writer, headerComment)
		if err != nil {
			return err
		}
	}
	_, err = writer.Write(output)

	return err
}

var (
	newlineBytes = []byte("\n")
	commentBytes = []byte("--")
)

// Clean up pg_dump's schema output (removing comments and multiple
// blank lines).
func clean(input []byte) []byte {
	prevBlank := true
	output := make([]byte, 0, len(input))
	lines := bytes.Split(input, newlineBytes)
	for _, line := range lines {
		if bytes.HasPrefix(line, commentBytes) {
			// Ignore SQL comment lines (pg_dump )
			continue
		}
		curBlank := len(bytes.TrimSpace(line)) == 0
		if prevBlank && curBlank {
			// Ignore multiple blank lines in a row
			continue
		}
		prevBlank = curBlank
		output = append(output, line...)
		output = append(output, '\n')
	}
	return output
}
