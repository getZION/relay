package media

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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
	logrus.Info("Someone is trying to upload a media file...")

	var params s3ConnectionParams
	envconfig.Process("", &params)

	request := UploadRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
		return err
	}

	// Log.Debug().
	// Str("Type", request.Type).
	// Str("Attestation", request.Attestation).
	// Msg("So we got...")

	switch request.Type {

	case "image/png":
		// Log.Debug().
		// Msg("PNG...")

		// Log.Debug().
		// Str("request.Bytes", request.Bytes).
		// Msg("request.Bytes")

		// imgBytes := request.Bytes
		imgBytes := "iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAApgAAAKYB3X3/OAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAANCSURBVEiJtZZPbBtFFMZ/M7ubXdtdb1xSFyeilBapySVU8h8OoFaooFSqiihIVIpQBKci6KEg9Q6H9kovIHoCIVQJJCKE1ENFjnAgcaSGC6rEnxBwA04Tx43t2FnvDAfjkNibxgHxnWb2e/u992bee7tCa00YFsffekFY+nUzFtjW0LrvjRXrCDIAaPLlW0nHL0SsZtVoaF98mLrx3pdhOqLtYPHChahZcYYO7KvPFxvRl5XPp1sN3adWiD1ZAqD6XYK1b/dvE5IWryTt2udLFedwc1+9kLp+vbbpoDh+6TklxBeAi9TL0taeWpdmZzQDry0AcO+jQ12RyohqqoYoo8RDwJrU+qXkjWtfi8Xxt58BdQuwQs9qC/afLwCw8tnQbqYAPsgxE1S6F3EAIXux2oQFKm0ihMsOF71dHYx+f3NND68ghCu1YIoePPQN1pGRABkJ6Bus96CutRZMydTl+TvuiRW1m3n0eDl0vRPcEysqdXn+jsQPsrHMquGeXEaY4Yk4wxWcY5V/9scqOMOVUFthatyTy8QyqwZ+kDURKoMWxNKr2EeqVKcTNOajqKoBgOE28U4tdQl5p5bwCw7BWquaZSzAPlwjlithJtp3pTImSqQRrb2Z8PHGigD4RZuNX6JYj6wj7O4TFLbCO/Mn/m8R+h6rYSUb3ekokRY6f/YukArN979jcW+V/S8g0eT/N3VN3kTqWbQ428m9/8k0P/1aIhF36PccEl6EhOcAUCrXKZXXWS3XKd2vc/TRBG9O5ELC17MmWubD2nKhUKZa26Ba2+D3P+4/MNCFwg59oWVeYhkzgN/JDR8deKBoD7Y+ljEjGZ0sosXVTvbc6RHirr2reNy1OXd6pJsQ+gqjk8VWFYmHrwBzW/n+uMPFiRwHB2I7ih8ciHFxIkd/3Omk5tCDV1t+2nNu5sxxpDFNx+huNhVT3/zMDz8usXC3ddaHBj1GHj/As08fwTS7Kt1HBTmyN29vdwAw+/wbwLVOJ3uAD1wi/dUH7Qei66PfyuRj4Ik9is+hglfbkbfR3cnZm7chlUWLdwmprtCohX4HUtlOcQjLYCu+fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII="

		unbased, _ := base64.StdEncoding.DecodeString(imgBytes)
		// Log.Debug().
		// 	Bytes("unbased", unbased).
		// 	Msg("unbased...")

		res := bytes.NewReader(unbased)
		// byteSize := res.Size()
		// Log.Debug().
		// 	Int64("byteSize", byteSize).
		// 	Msg("byteSize...")
		pngI, errPng := png.Decode(res)

		if errPng == nil {
			// os.Chdir()
			// file, errFile := os.Open("./test2.png")
			file, errFile := os.OpenFile("test.png", os.O_WRONLY|os.O_CREATE, 0777)
			defer file.Close()

			if errFile != nil {
				// Log.Debug().
				// Err(errFile).
				// Msg("err file...")
				return errFile
			}

			errEncode := png.Encode(file, pngI)
			if errEncode != nil {
				// Log.Debug().
				// 	Err(errEncode).
				// 	Msg("err encode...")
				return errEncode
			}

			fileStat, errStat := file.Stat()
			if errStat != nil {
				// Log.Debug().
				// 	Err(errStat).
				// 	Msg("err stat...")
				return errStat
			}
			fileSize := fileStat.Size()

			// Log.Debug().
			// 	Int64("fileSize", fileSize).
			// 	Str("filesizzeee", fmt.Sprint(fileSize)).
			// 	Msg("File size")

			// create an AWS session which can be
			// reused if we're uploading many files
			s, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-1"),
				Credentials: credentials.NewStaticCredentials(
					params.AccessId,     // id
					params.AccessSecret, // secret
					""),                 // token can be left blank for now
			})
			if err != nil {
				// fmt.Fprintf(w, "Could not upload file")
				// Log.Debug().
				// 	Err(err).
				// 	Msg("Could not upload file")
				return err
			}

			userDid := request.Attestation

			objectUrl, err := UploadFileToS3(s, file, fileSize, userDid)
			if err != nil {
				// fmt.Fprintf(w, "Could not upload file")
				// Log.Debug().
				// 	Err(err).
				// 	Msg("Could not upload file 2")
				return err
			} else {
				// Log.Debug().
				// 	Str("objectUrl:", objectUrl).
				// 	Msg("Image uploaded successfully")
				// ctx.Response()
				// ctx.Response().Write(objectUrl)
				// set some headers and status code first
				ctx.Response().SetStatusCode(200)
				ctx.Response().SetBody([]byte(objectUrl))

				tr := UploadResponse{true, objectUrl}
				ctx.JSON(tr)

				// Log.Debug().
				// 	Msg("returning after setting JSON response!")

				return nil
				// ctx.Response().SetContentType("foo/bar")
				// ctx.SetStatusCode(fasthttp.StatusOK)

				// then write the first part of body
				// fmt.Fprintf(ctx, "this is the first part of body\n")

				// then set more headers
				// ctx.Response.Header.Set("Foo-Bar", "baz")

				// then write more body
				// fmt.Fprintf(ctx, "this is the second part of body\n")

				// then override already written body
				// ctx.SetBody([]byte("this is completely new body contents"))

				// then update status code
				// ctx.SetStatusCode(fasthttp.StatusNotFound)

			}

			// fmt.Fprintf(w, "Image uploaded successfully: %v", fileName)

			// png.Encode(f, pngI)
			// fmt.Println("Png generated")
			// user := model.User{}
			// user.ProfilePic = "storage/" + strconv.Itoa(session.Values["id"].(int)) + "/" + uid.String()
			// session.Values["profile_pic"] = user.ProfilePic
			// session.Save(r, w)
			// database.Connection.Id(session.Values["id"].(int)).Update(&user)
		} else {
			// Log.Debug().
			// 	Str("png bounds hm", pngI.Bounds().String()).
			// 	Msg("this is what")
			fmt.Println(errPng.Error())
		}

	case "image/jpeg":
		// Log.Debug().
		// 	Msg("JPEG...")
	default:
		// Log.Debug().
		// 	Msg("No!")
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
