package flogos3tofile

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	// f, err := os.Open("/Users/avanderg@tibco.com/working/image_recog_hands_on/IMG_20190214_152236108.jpg")
	// if err != nil {
	// 	log.Fatal("trouble loading test file")
	// }

	// defer f.Close()

	// fileInfo, _ := f.Stat()
	// var size int64 = fileInfo.Size()
	// bytes := make([]byte, size)

	// // read file into bytes
	// buffer := bufio.NewReader(f)
	// _, err = buffer.Read(bytes)

	// //THIS TEST IS BROKEN CONVERT FILE TO []byte

	// settings := &Settings{ResamplingFilter: "NearestNeighbor"}

	// iCtx := test.NewActivityInitContext(settings, nil)
	// act, err := New(iCtx)
	// assert.Nil(t, err)
	// input := &Input{MaxDimSize: 256, File: bytes}
	// tc := test.NewActivityContext(act.Metadata())

	// tc.SetInputObject(input)

	// done, err := act.Eval(tc)
	// assert.True(t, done)
	// assert.Nil(t, err)

	// output := &Output{}
	// tc.GetOutputObject(output)
	// // fmt.Println(output.ResizedImage)
}
