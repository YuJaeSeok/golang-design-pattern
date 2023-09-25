package templatemethod

import "testing"

func ExampleNewHttpDownloader() {
	var downloader Downloader = NewHttpDownloader()
	downloader.Download("http://example.com/abc.zip")
}

func ExampleNewFTPDownloader() {
	var downloader Downloader = NewFTPDownloader()
	downloader.Download("ftp://example.com/abc.zip")
}

func TestNewTemplate(t *testing.T) {
	ExampleNewHttpDownloader()
	ExampleNewFTPDownloader()
}
