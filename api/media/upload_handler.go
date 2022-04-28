package media

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/gofiber/fiber/v2"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type UploadRequest struct {
	Attestation string `json:"attestation,omitempty"`
	Bytes       string `json:"bytes,omitempty"`
	// Type        string `json:"type,omitempty"`
}

type UploadResponse struct {
	Success  bool   `json:"success"`
	ImageUrl string `json:"imageUrl"`
}

type s3ConnectionParams struct {
	KeyBucketRegion string `envconfig:"S3_IMAGES_BUCKET_REGION"`
	KeyBucketName   string `envconfig:"S3_IMAGES_BUCKET_NAME"`
	AccessId        string `envconfig:"S3_ACCESS_ID"`
	AccessSecret    string `envconfig:"S3_ACCESS_SECRET"`
}

func UploadHandler(ctx *fiber.Ctx) error {
	logrus.Info("Trying to upload a media file...!")

	var params s3ConnectionParams
	envconfig.Process("", &params)

	request := UploadRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return err
	}

	return nil
}

// UploadFileToS3 saves a file to aws bucket and returns the url to
// the file and an error if there's any
func UploadFileToS3(s *session.Session, file *os.File, size int64, userDid string) (string, error) {

	bucketName := "zion-images-switch"

	buffer := make([]byte, size)
	file.Read(buffer)

	userKey := "user-" + userDid

	// create a unique file name for the file
	tempFileName := userKey + "/" + bson.NewObjectId().Hex() + ".png" //filepath.Ext(fileHeader.Filename)
	// tempFileName := "pictures/" + bson.NewObjectId().Hex() + ".png" //filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}
	// Log.Debug().
	// 	Str("wat is this", wat.String()).
	// 	Msg("success - what is this?")

	objectUrl := "https://" + bucketName + ".s3.amazonaws.com/" + tempFileName

	// Log.Debug().
	// 	Str("objectUrl", objectUrl).
	// 	Msg("objectUrl")

	return objectUrl, err
}
