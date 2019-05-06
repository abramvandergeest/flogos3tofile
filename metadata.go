package flogos3tofile

import (
	"os"

	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	ResamplingFilter string `md:"resamplingFilter"`
}

type Input struct {
	Bucket       string `md:"bucket,required"`
	Key string         `md:"key.required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	r.Bucket,_ = coerce.ToString(values["bucket"])
	r.Key,_ = coerce.ToString(values["key"])

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"bucket":       r.Bucket,
		"key": r.Key,
	}
}

type Output struct {
	File *os.File `md:"file"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.File = values["file"].(*os.File)
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"file": o.File,
	}
}
