package main

import (
	"os"
	"errors"
	"archive/zip"
	"path/filepath"
	"strings"
	"io"
)

const target = "downloads.zip"

func archive(dir string) (string, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return "", err
	}

	baseDir := filepath.Base(dir)

	if !info.IsDir() {
		return "", errors.New("dir should be as dir")
	}

	zipFile, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, dir))

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Store
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err

	})

	if err != nil {
		return "", err
	}
	return target, nil
}
