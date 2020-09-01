/*
   @Time : 20-9-1 上午10:23
   @Author : young
   @File : templatemethod.go
   @Software: GoLand
*/
package _4_template_method

import "fmt"

type doAction interface {
	prepare()
	download()
	save()
	end()
}

type Template struct {
	doAction
	url string
}

func newTemplate(tmp doAction) Template {
	return Template{
		doAction: tmp,
	}
}

func (t *Template) prepare() {
	fmt.Println("this is default prepare")
}

func (t *Template) download() {
	fmt.Println("this is default download")
}

func (t *Template) save() {
	fmt.Println("this is default save.")
}

func (t *Template) end() {
	fmt.Println("this is default end.")
}

type DownLoader interface {
	Download(url string)
}

func (t *Template) Download(url string) {
	// step 1
	t.url = url
	fmt.Println("[default] step1: prepare.")
	t.doAction.prepare()

	// step 2
	fmt.Println("[default] step2: download.")
	t.doAction.download()

	// step 3
	fmt.Println("[default] step3: save.")
	t.doAction.save()

	// step 4
	fmt.Println("[default] step4: end.")
	t.doAction.end()
}

type HttpDownloader struct {
	Template
}

func NewHTTPDownloader() DownLoader {
	downloader := &HttpDownloader{}

	template := newTemplate(downloader)
	downloader.Template = template
	return downloader
}

func (the *HttpDownloader) prepare() {
	fmt.Printf("[http]: this url : %s\n", the.url)
}

func (the *HttpDownloader) download() {
	fmt.Println("[http]: this is download.")
}

//func(the *HttpDownloader)save(){
//	fmt.Println("[http]: this is save.")
//}

func (the *HttpDownloader) end() {
	fmt.Println("[http]: this is end.")
}


// tcp loader

type TCPDownloader struct {
	Template
}

func NewTCPDownloader() DownLoader {
	downloader := &TCPDownloader{}

	template := newTemplate(downloader)
	downloader.Template = template
	return downloader
}

func (the *TCPDownloader) prepare() {
	fmt.Printf("[tcp]: this url : %s\n", the.url)
}

func (the *TCPDownloader) download() {
	fmt.Println("[tcp]: this is download.")
}

//func(the *TCPDownloader)save(){
//	fmt.Println("[tcp]: this is http save.")
//}

func (the *TCPDownloader) end() {
	fmt.Println("[tcp]: this is end.")
}
