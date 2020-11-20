package main

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"strings"
)

func main() {
	var seg = gojieba.NewJieba()
	defer seg.Free()
	var useHmm = true
	var separator = "|"
	var resWords []string
	var sentence = "比特币"
	resWords = seg.CutAll(sentence)
	fmt.Printf("%s\t全模式：%s \n", sentence, strings.Join(resWords, separator))
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t精确模式：%s \n", sentence, strings.Join(resWords, separator))
	//var addWord = "万里长"
	//seg.AddWord(addWord)
	//fmt.Printf("添加新词：%s\n", addWord)
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t精确模式：%s \n", sentence, strings.Join(resWords, separator))
	sentence = "北京鲜花速递"
	resWords = seg.Cut(sentence, useHmm)
	fmt.Printf("%s\t新词识别：%s \n", sentence, strings.Join(resWords, separator))
	sentence = "北京鲜花速递"
	resWords = seg.CutForSearch(sentence, useHmm)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))
	sentence = "北京市朝阳公园"
	resWords = seg.Tag(sentence)
	fmt.Println(sentence, "\t词性标注：", strings.Join(resWords, separator))
	sentence = "鲁迅先生"
	resWords = seg.CutForSearch(sentence, !useHmm)
	fmt.Println(sentence, "\t搜索引擎模式：", strings.Join(resWords, separator))
	words := seg.Tokenize(sentence, gojieba.SearchMode, !useHmm)
	fmt.Println(sentence, "\tTokenize Search Mode 搜索引擎模式：", words)
	words = seg.Tokenize(sentence, gojieba.DefaultMode, !useHmm)
	fmt.Println(sentence, "\tTokenize Default Mode搜索引擎模式：", words)
	word2 := seg.ExtractWithWeight(sentence, 5)
	fmt.Println(sentence, "\tExtract：", word2)
	return
}
