package main

import (
	"github.com/Vito89/heaputil"
	"testing"
)

func TestShowStatisticMessageSend(t *testing.T) {
	var expectedString = `Today we have top of the list:
<1 place> heaputil.KeyValue{key:"key1", val:10}
<2 place> heaputil.KeyValue{key:"key2", val:1}
<3 place> heaputil.KeyValue{key:"key0", val:0}`
	dict := map[string]int{"key0": 0, "key1": 10, "key2": 1}
	heapArray := heaputil.GetHeap(dict)

	actualString := showStatisticMessageSend(heapArray)

	if actualString != expectedString {
		t.Errorf("Wanted:\n%v\nbut got:\n%v", expectedString, actualString)
	}
}

func TestShowStatisticMessageSendEmpty(t *testing.T) {
	var expectedString = `Today we have empty top :(`
	dict := map[string]int{}
	heapArray := heaputil.GetHeap(dict)

	actualString := showStatisticMessageSend(heapArray)

	if actualString != expectedString {
		t.Errorf("Wanted:\n%v\nbut got:\n%v", expectedString, actualString)
	}
}
