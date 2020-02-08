package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)



func BenchmarkStringPlus(b *testing.B){
	for i:=0;i<b.N;i++{
		s :="举杯邀明月"+"对影成三人"
		s +="会当凌绝顶"+"一览众山小"
	}
}
func BenchmarkStringFmt(b *testing.B){
	for i:=0;i<b.N;i++{
		fmt.Sprintf("%s,%s,%s,%s","举杯邀明月","对影成三人","会当凌绝顶","一览众山小")
	}
}

func BenchmarkStringJoin(b *testing.B){
	for i:=0;i<b.N;i++{
		s :=[]string{"举杯邀明月","对影成三人","会当凌绝顶","一览众山小"}
		strings.Join(s,"")
	}
}

func BenchmarkStringBuffer(b *testing.B){
	for i:=0;i<b.N;i++{
		var b bytes.Buffer
		b.WriteString("举杯邀明月")
		b.WriteString("对影成三人")
		b.WriteString("会当凌绝顶")
		b.WriteString("一览众山小")
	}
}

func BenchmarkStringBuilder(b *testing.B){
	for i:=0;i<b.N;i++{
		var s strings.Builder
		s.WriteString("举杯邀明月")
		s.WriteString("对影成三人")
		s.WriteString("会当凌绝顶")
		s.WriteString("一览众山小")
	}
}