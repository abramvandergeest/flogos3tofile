package flogos3tofile

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	
	"os"
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	activity.Register(&Activity{}, New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: ResamplingFilter = %s", s.ResamplingFilter)

	act := &Activity{settings: s} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings *Settings
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	// a.settings

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	fmt.Println(input.Bucket,input.Key)

    file, err := os.Create(input.Key)
    if err != nil {
		ctx.Logger().Errorf("Unable to open file %q, %v,",err)
    }

	defer file.Close()
	
	    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )

    downloader := s3manager.NewDownloader(sess)

    numBytes, err := downloader.Download(file,
        &s3.GetObjectInput{
            Bucket: aws.String(input.Bucket),
            Key:    aws.String(input.Key),
        })
    if err != nil {
        ctx.Logger().Errorf("Unable to download item %q, %v", input.Key, err)
    }

    ctx.Logger().Infof("Downloaded %s %d bytes", file.Name(), numBytes)

	outFile, err := os.Open(input.Key)
	if err != nil {
        ctx.Logger().Errorf("Unable to load file %q, %v", input.Key, err)
    }

	ctx.Logger().Infof("file opened for output %q", file.Name())

	output := &Output{File: outFile}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
