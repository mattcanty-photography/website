package website

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getPortfolio(portfolioID string) (portfolio, error) {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(os.Getenv("PHOTO_BUCKET_NAME")),
		Prefix:    aws.String(fmt.Sprintf("portfolios/%s/albums/", portfolioID)),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return portfolio{}, err
	}

	var albums []album

	for _, o := range resp.CommonPrefixes {
		albumID := strings.Split(*o.Prefix, "/")[3]
		coverID, err := getCoverPhotoID(portfolioID, albumID)
		if err != nil {
			return portfolio{}, err
		}

		albums = append(albums, album{
			ID:          albumID,
			PortfolioID: portfolioID,
			CoverID:     coverID,
		})
	}

	return portfolio{ID: portfolioID, Albums: albums}, nil
}

func getAlbum(portfolioID string, albumID string) (album, error) {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	prefix := fmt.Sprintf("portfolios/%s/albums/%s/thumbs/", portfolioID, albumID)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:     aws.String(os.Getenv("PHOTO_BUCKET_NAME")),
		Prefix:     aws.String(prefix),
		StartAfter: aws.String(prefix),
	})
	if err != nil {
		return album{}, err
	}

	var photos []photo

	for _, o := range resp.Contents {
		keyElems := strings.Split(*o.Key, "/")
		photos = append(photos, photo{
			ID: keyElems[5],
			Album: album{
				ID:          keyElems[3],
				PortfolioID: keyElems[1],
			},
		})
	}

	return album{ID: albumID, Photos: photos}, nil
}

func getCoverPhotoID(portfolioID string, albumID string) (string, error) {
	bucket := os.Getenv("PHOTO_BUCKET_NAME")
	prefix := fmt.Sprintf("portfolios/%s/albums/%s/thumbs/", portfolioID, albumID)

	svc := s3.New(session.New(), &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:     aws.String(bucket),
		Prefix:     aws.String(prefix),
		StartAfter: aws.String(prefix),
		MaxKeys:    aws.Int64(1),
	})
	if err != nil {
		return "", err
	}

	return strings.Split(*resp.Contents[0].Key, "/")[5], nil
}
