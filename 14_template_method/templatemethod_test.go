/*
   @Time : 20-9-1 上午10:24
   @Author : young
   @File : templatemethod_test.go
   @Software: GoLand
*/
package _4_template_method

import "testing"

func TestNewHTTPDownloader(t *testing.T) {
	httpLoader := NewHTTPDownloader()
	httpLoader.Download("www.facebook.com")
}

func TestNewTCPDownloader(t *testing.T) {
	tcpLoader := NewTCPDownloader()
	tcpLoader.Download("www.google.com")
}
